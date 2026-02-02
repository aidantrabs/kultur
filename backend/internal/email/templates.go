package email

import (
    "bytes"
    "html/template"
)

const (
    ColorRed     = "#CE1126"
    ColorRedDark = "#a30d1e"
    ColorBlack   = "#1a1a1a"
    ColorWhite   = "#fafafa"
    ColorGold    = "#D4A84B"
    ColorGray    = "#6b7280"
    ColorBorder  = "#e5e7eb"
)

type TemplateData struct {
    PreviewText    string
    LogoURL        string
    Heading        string
    Body           template.HTML
    ButtonText     string
    ButtonURL      string
    FooterText     string
    UnsubscribeURL string
    Year           int
}

const baseTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="x-apple-disable-message-reformatting">
    <meta name="format-detection" content="telephone=no,address=no,email=no,date=no,url=no">
    <title>{{.Heading}}</title>
    <!--[if mso]>
    <noscript>
        <xml>
            <o:OfficeDocumentSettings>
                <o:PixelsPerInch>96</o:PixelsPerInch>
            </o:OfficeDocumentSettings>
        </xml>
    </noscript>
    <![endif]-->
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body { margin: 0; padding: 0; width: 100%; -webkit-text-size-adjust: 100%; -ms-text-size-adjust: 100%; }
        table { border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; }
        img { border: 0; height: auto; line-height: 100%; outline: none; text-decoration: none; -ms-interpolation-mode: bicubic; }
        a { color: ` + ColorRed + `; text-decoration: none; }
        a:hover { text-decoration: underline; }
        .button:hover { background-color: ` + ColorRedDark + ` !important; }
        @media only screen and (max-width: 600px) {
            .container { width: 100% !important; padding: 0 16px !important; }
            .content { padding: 32px 24px !important; }
            h1 { font-size: 24px !important; }
        }
    </style>
</head>
<body style="margin: 0; padding: 0; background-color: #f3f4f6; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;">
    <!-- Preview Text -->
    <div style="display: none; max-height: 0; overflow: hidden;">
        {{.PreviewText}}
        &#847; &#847; &#847; &#847; &#847; &#847; &#847; &#847; &#847; &#847; &#847; &#847; &#847; &#847; &#847;
    </div>

    <!-- Email Container -->
    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%" style="background-color: #f3f4f6;">
        <tr>
            <td style="padding: 40px 16px;">
                <table role="presentation" cellspacing="0" cellpadding="0" border="0" class="container" style="max-width: 560px; margin: 0 auto; width: 100%;">

                    <!-- Header with Logo -->
                    <tr>
                        <td style="padding: 0 0 24px 0; text-align: center;">
                            <a href="https://kultur-tt.app" style="text-decoration: none;">
                                <img src="https://kultur-tt.app/logo.png" alt="KULTUR" width="80" height="80" style="display: block; margin: 0 auto; border: 0;" />
                            </a>
                        </td>
                    </tr>

                    <!-- Main Content Card -->
                    <tr>
                        <td>
                            <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%" style="background-color: #ffffff; border-radius: 12px; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);">

                                <!-- Red Accent Bar -->
                                <tr>
                                    <td style="height: 4px; background: linear-gradient(90deg, ` + ColorRed + ` 0%, ` + ColorGold + ` 100%); border-radius: 12px 12px 0 0;"></td>
                                </tr>

                                <!-- Content -->
                                <tr>
                                    <td class="content" style="padding: 40px 32px;">
                                        <!-- Heading -->
                                        <h1 style="margin: 0 0 24px 0; font-size: 28px; font-weight: 700; color: ` + ColorBlack + `; line-height: 1.3;">
                                            {{.Heading}}
                                        </h1>

                                        <!-- Body Content -->
                                        <div style="margin: 0 0 32px 0; font-size: 16px; line-height: 1.7; color: #374151;">
                                            {{.Body}}
                                        </div>

                                        {{if .ButtonText}}
                                        <!-- CTA Button -->
                                        <table role="presentation" cellspacing="0" cellpadding="0" border="0" style="margin: 0 auto;">
                                            <tr>
                                                <td style="border-radius: 8px; background-color: ` + ColorRed + `;">
                                                    <a class="button" href="{{.ButtonURL}}" target="_blank" style="display: inline-block; padding: 16px 32px; font-size: 16px; font-weight: 600; color: #ffffff; text-decoration: none; border-radius: 8px;">
                                                        {{.ButtonText}}
                                                    </a>
                                                </td>
                                            </tr>
                                        </table>
                                        {{end}}
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>

                    <!-- Footer -->
                    <tr>
                        <td style="padding: 32px 16px; text-align: center;">
                            {{if .FooterText}}
                            <p style="margin: 0 0 16px 0; font-size: 14px; line-height: 1.6; color: ` + ColorGray + `;">
                                {{.FooterText}}
                            </p>
                            {{end}}

                            <p style="margin: 0 0 8px 0; font-size: 13px; color: ` + ColorGray + `;">
                                <strong style="color: ` + ColorBlack + `;">KULTUR</strong> · Your guide to Trinidad's cultural festivals
                            </p>

                            {{if .UnsubscribeURL}}
                            <p style="margin: 0; font-size: 12px; color: #9ca3af;">
                                <a href="{{.UnsubscribeURL}}" style="color: #9ca3af; text-decoration: underline;">Unsubscribe</a>
                            </p>
                            {{end}}
                        </td>
                    </tr>

                </table>
            </td>
        </tr>
    </table>
</body>
</html>`

var tmpl = template.Must(template.New("email").Parse(baseTemplate))

func RenderTemplate(data TemplateData) (string, error) {
    var buf bytes.Buffer
    if err := tmpl.Execute(&buf, data); err != nil {
        return "", err
    }

    return buf.String(), nil
}
