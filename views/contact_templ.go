// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/andrenormanlang/common"

func MakeContactFormWithRecaptcha(recaptcha_sitekey string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form id=\"demo-form\" class=\"space-y-4 p-12\" method=\"post\" hx-post=\"/contact-send\" hx-target=\"#contact-form\" hx-trigger=\"verified\"><label for=\"name\" class=\"block text-md font-medium\">Name:</label> <input id=\"name\" class=\"mt-1 block w-full px-3 py-2 border border-pastel-blue dark:border-pastel-blue-900 rounded-md shadow-md focus:outline-none focus:ring-pastel-blue-500 focus:border-pastel-blue-500\" type=\"text\" name=\"name\" required> <label for=\"email\" class=\"block text-md font-medium\">Email:</label> <input id=\"email\" class=\"mt-1 block w-full px-3 py-2 border border-pastel-blue dark:border-pastel-blue-900 rounded-md shadow-md focus:outline-none focus:ring-pastel-blue-500 focus:border-pastel-blue-500\" type=\"email\" name=\"email\" required> <label for=\"message\" class=\"block text-md font-medium\">Message:</label> <textarea id=\"message\" class=\"mt-1 block w-full px-3 py-2 border border-pastel-blue dark:border-pastel-blue-900 rounded-md shadow-md focus:outline-none focus:ring-pastel-blue-500 focus:border-pastel-blue-500\" name=\"message\" rows=\"4\" cols=\"50\" required></textarea><div class=\"text-center\"><button class=\"g-recaptcha text-gray-700 dark:text-gray-900 w-fit inline-flex justify-center py-2 px-4 border border-transparent shadow-md text-sm font-medium rounded-md bg-gray-100 hover:bg-gray-200 dark:bg-gray-200 dark:hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-offset-2 fous:ring-pastel-blue-300\" data-sitekey=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(recaptcha_sitekey))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-callback=\"onSubmit\" data-action=\"submit\">Submit\r</button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func MakeContactForm() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form id=\"demo-form\" class=\"space-y-4 p-12\" method=\"post\" hx-post=\"/contact-send\" hx-target=\"#contact-form\"><label for=\"name\" class=\"block text-md font-medium\">Name:</label> <input id=\"name\" class=\"mt-1 block w-full px-3 py-2 border border-pastel-blue dark:border-pastel-blue-900 rounded-md shadow-md focus:outline-none focus:ring-pastel-blue-500 focus:border-pastel-blue-500\" type=\"text\" name=\"name\" required> <label for=\"email\" class=\"block text-md font-medium\">Email:</label> <input id=\"email\" class=\"mt-1 block w-full px-3 py-2 border border-pastel-blue dark:border-pastel-blue-900 rounded-md shadow-md focus:outline-none focus:ring-pastel-blue-500 focus:border-pastel-blue-500\" type=\"email\" name=\"email\" required> <label for=\"message\" class=\"block text-md font-medium\">Message:</label> <textarea id=\"message\" class=\"mt-1 block w-full px-3 py-2 border border-pastel-blue dark:border-pastel-blue-900 rounded-md shadow-md focus:outline-none focus:ring-pastel-blue-500 focus:border-pastel-blue-500\" name=\"message\" rows=\"4\" cols=\"50\" required></textarea><div class=\"text-center\"><button type=\"submit\" class=\"text-gray-700 dark:text-gray-900 w-fit inline-flex justify-center py-2 px-4 border border-transparent shadow-md text-sm font-medium rounded-md bg-gray-100 hover:bg-gray-200 dark:bg-gray-200 dark:hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-offset-2 fous:ring-pastel-blue-300\">Send Message\r</button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func makeContact(recaptcha_sitekey string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"contact-form\" class=\"rounded-lg shadow-md w-full\"><h2 class=\"text-4xl font-bold mb-6 text-center\">Contact Us</h2>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(recaptcha_sitekey) > 0 {
			templ_7745c5c3_Err = MakeContactFormWithRecaptcha(recaptcha_sitekey).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = MakeContactForm().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(recaptcha_sitekey) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script src=\"https://www.google.com/recaptcha/api.js\"></script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func MakeContactPage(links []common.Link, recaptcha_sitekey string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = MakeLayout("Menu and Contact Form", links, makeContact(recaptcha_sitekey)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
