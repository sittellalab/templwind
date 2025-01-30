# Templwind

Convenient and accessible Templ components for Go developers.

### Features
  - Accessibility
  - CSP compatible
  - Configurability
  - A library of additional helper functions
  - Looks kinda somewhat decent

## Usage

### Installation

`go get github.com/dimmerz92/templwind`

### Example

**In a Templ file**
```templ
package components

import "github.com/dimmerz92/templwind/twui"

templ MyComponent() {
  // ... surrounding content
  @twui.Alert(ui.Alert{Theme: themes.Warning})
  // ... surrounding content
}
```

**In a Go file**
```go
package main

import (
  "net/http"

  "github.com/dimmerz92/templwind/twlib"
  "github.com/dimmerz92/templwind/twui"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		twlib.Render(w, r, 200, twui.Alert())
	})

  log.Fatal(http.ListenAndServe(":8000", nil))
}
```

## Planned Components

### Actions
- [ ] Dialog
- [ ] Modal
- [X] [Theme Toggle](/pkg/twui/theme_toggle.templ)

### Containers (Active)
- [ ] Accordion
- [ ] Carousel
- [ ] DataTable
- [ ] Dropdown Menu
- [ ] Hover Card
- [ ] Sheet
- [ ] Toast

### Containers (Passive)
- [ ] Alert
- [ ] Avatar
- [ ] Badge
- [ ] Card
- [ ] Scrollable
- [ ] Table

### Content
- [X] [Lucide Icons](/pkg/twicons)
- [X] [Spinner](/pkg/twui/spinner.templ)

### Form Controls
- [ ] Button
- [ ] Checkbox
- [ ] Color Input
- [ ] Date Input
- [ ] Datetime Input
- [ ] Email Input
- [ ] File Input
- [ ] Month Input
- [ ] Number Input
- [ ] OTP Input
- [ ] Password Input
- [ ] Radio
- [ ] Range Input
- [ ] Select
- [ ] Search Input
- [ ] Switch
- [ ] Tel Input
- [ ] Textarea
- [ ] Text Input
- [ ] URL Input
- [ ] Week Input

### Navigation
- [ ] Breadcrumb
- [ ] Nav Menu
- [ ] Sidebar
- [ ] Tabs


## Contributing

  - For a small change, just send a PR.
  - For bigger changes open an issue for discussion before sending a PR.
  - PR should have:
    - Documentation
    - Example (If it makes sense)
  - You can also contribute by:
    - Reporting issues
    - Suggesting new features or enhancements
    - Improve/fix documentation

## License

This project is licensed under the [MIT license](/LICENSE).
