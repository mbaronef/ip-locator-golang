package main

import (
	"iplocator/core"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/iplocate/go-iplocate"
)

type AppState struct {
	resultsContainer *fyne.Container
	currentResults   []*iplocate.LookupResponse
	client           *iplocate.Client
	window           fyne.Window
}

func main() {
	ipLocatorApp := app.New()
	window := ipLocatorApp.NewWindow("IP Locator")

	appState, _ := initializeApp(window)

	content := setupUI(appState, ipLocatorApp)
	window.SetContent(content)
	window.Resize(fyne.NewSize(1100, 600))
	window.ShowAndRun()
}

func initializeApp(window fyne.Window) (*AppState, error) {
	appState := &AppState{
		resultsContainer: container.NewVBox(),
		window:           window,
	}

	client, err := core.NewClient()
	if err != nil {
		showError(appState, err.Error())
		return appState, err
	}
	appState.client = client

	return appState, nil
}

func setupUI(appState *AppState, fyneApp fyne.App) *fyne.Container {
	ipEntry := widget.NewEntry()
	ipEntry.SetPlaceHolder("Enter IP address:")

	lookupBtn := lookupButtonEvent(appState, ipEntry)
	selfLookupBtn := selfLookupButtonEvent(appState)
	fileUploadBtn := fileUploadButtonEvent(appState)
	jsonDownloadBtn := jsonDownloadButtonEvent(appState)
	textDownloadBtn := textDownloadButtonEvent(appState)
	exitBtn := exitButtonEvent(fyneApp)

	buttonContainer := container.NewHBox(
		lookupBtn, selfLookupBtn, fileUploadBtn, jsonDownloadBtn, textDownloadBtn, exitBtn,
	)

	resultScroll := container.NewScroll(appState.resultsContainer)
	resultScroll.SetMinSize(fyne.NewSize(950, 400))

	return container.NewVBox(
		ipEntry,
		buttonContainer,
		resultScroll,
	)
}

func createResultCard(result *iplocate.LookupResponse) *widget.Card {
	title := "IP: " + result.IP
	content := core.FormatResult(result)
	return widget.NewCard(title, "", widget.NewLabel(content))
}

func updateResults(appState *AppState, results []*iplocate.LookupResponse) {
	appState.currentResults = results
	appState.resultsContainer.RemoveAll()

	if len(results) == 0 {
		appState.resultsContainer.Add(widget.NewLabel("No results to display"))
		return
	}

	for _, result := range results {
		appState.resultsContainer.Add(createResultCard(result))
	}
}

func showError(appState *AppState, message string) {
	appState.resultsContainer.RemoveAll()
	appState.resultsContainer.Add(widget.NewCard("Error", "", widget.NewLabel(message)))
}

func lookupButtonEvent(appState *AppState, ipEntry *widget.Entry) *widget.Button {
	return widget.NewButton("Lookup", func() {
		ip := ipEntry.Text
		if ip == "" {
			showError(appState, "Please enter an IP address!")
			return
		}

		if err := core.ValidateIPs([]string{ip}); err != nil {
			showError(appState, err.Error())
			return
		}

		res, errs := core.LookupIPs(appState.client, []string{ip})
		if len(errs) > 0 && errs[0] != nil {
			showError(appState, "Error: "+errs[0].Error())
			return
		}

		updateResults(appState, res)
	})
}

func selfLookupButtonEvent(appState *AppState) *widget.Button {
	return widget.NewButton("Self Lookup", func() {
		res, err := core.LookupSelf(appState.client)
		if err != nil {
			showError(appState, "Error: "+err.Error())
			return
		}
		if res == nil {
			showError(appState, "No result returned from self lookup")
			return
		}
		updateResults(appState, []*iplocate.LookupResponse{res})
	})
}

func fileUploadButtonEvent(appState *AppState) *widget.Button {
	return widget.NewButton("Upload File", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				showError(appState, "Error opening file: "+err.Error())
				return
			}
			if reader == nil {
				return
			}
			defer reader.Close()

			ips, err := core.ReadIPsFromFile(reader.URI().Path())
			if err != nil {
				showError(appState, err.Error())
				return
			}

			if err := core.ValidateIPs(ips); err != nil {
				showError(appState, err.Error())
				return
			}

			res, errs := core.LookupIPs(appState.client, ips)
			if len(errs) > 0 && errs[0] != nil {
				showError(appState, "Error: "+errs[0].Error())
				return
			}

			updateResults(appState, res)
		}, appState.window)
	})
}

func jsonDownloadButtonEvent(appState *AppState) *widget.Button {
	return widget.NewButton("Download JSON", func() {
		if len(appState.currentResults) == 0 {
			dialog.ShowInformation("No Data", "No results to download. Please perform a lookup first.", appState.window)
			return
		}

		jsonData, err := core.FormatJSON(appState.currentResults)
		if err != nil {
			dialog.ShowError(err, appState.window)
			return
		}

		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, appState.window)
				return
			}
			if writer == nil {
				return
			}
			defer writer.Close()

			_, err = writer.Write([]byte(jsonData))
			if err != nil {
				dialog.ShowError(err, appState.window)
				return
			}

			dialog.ShowInformation("Success", "JSON file saved successfully!", appState.window)
		}, appState.window)
	})
}

func textDownloadButtonEvent(appState *AppState) *widget.Button {
	return widget.NewButton("Download Text", func() {
		if len(appState.currentResults) == 0 {
			dialog.ShowInformation("No Data", "No results to download. Please perform a lookup first.", appState.window)
			return
		}

		var textData string
		for i, result := range appState.currentResults {
			if i > 0 {
				textData += "\n" + strings.Repeat("=", 50) + "\n\n"
			}
			textData += core.FormatResult(result)
		}

		dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, appState.window)
				return
			}
			if writer == nil {
				return
			}
			defer writer.Close()

			_, err = writer.Write([]byte(textData))
			if err != nil {
				dialog.ShowError(err, appState.window)
				return
			}

			dialog.ShowInformation("Success", "Text file saved successfully!", appState.window)
		}, appState.window)
	})
}

func exitButtonEvent(fyneApp fyne.App) *widget.Button {
	return widget.NewButton("Exit", func() {
		fyneApp.Quit()
	})
}
