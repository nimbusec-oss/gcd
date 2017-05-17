// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains CSS functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/nimbusec-oss/gcd/gcdmessage"
)

// CSS rule collection for a single pseudo style.
type CSSPseudoElementMatches struct {
	PseudoType string          `json:"pseudoType"` // Pseudo element type. enum values: first-line, first-letter, before, after, backdrop, selection, first-line-inherited, scrollbar, scrollbar-thumb, scrollbar-button, scrollbar-track, scrollbar-track-piece, scrollbar-corner, resizer, input-list-button
	Matches    []*CSSRuleMatch `json:"matches"`    // Matches of CSS rules applicable to the pseudo style.
}

// Inherited CSS rule collection from ancestor node.
type CSSInheritedStyleEntry struct {
	InlineStyle     *CSSCSSStyle    `json:"inlineStyle,omitempty"` // The ancestor node's inline style, if any, in the style inheritance chain.
	MatchedCSSRules []*CSSRuleMatch `json:"matchedCSSRules"`       // Matches of CSS rules matching the ancestor node in the style inheritance chain.
}

// Match data for a CSS rule.
type CSSRuleMatch struct {
	Rule              *CSSCSSRule `json:"rule"`              // CSS rule in the match.
	MatchingSelectors []int       `json:"matchingSelectors"` // Matching selector indices in the rule's selectorList selectors (0-based).
}

// Data for a simple selector (these are delimited by commas in a selector list).
type CSSValue struct {
	Text  string          `json:"text"`            // Value text.
	Range *CSSSourceRange `json:"range,omitempty"` // Value range in the underlying resource (if available).
}

// Selector list data.
type CSSSelectorList struct {
	Selectors []*CSSValue `json:"selectors"` // Selectors in the list.
	Text      string      `json:"text"`      // Rule selector text.
}

// CSS stylesheet metainformation.
type CSSCSSStyleSheetHeader struct {
	StyleSheetId string  `json:"styleSheetId"`           // The stylesheet identifier.
	FrameId      string  `json:"frameId"`                // Owner frame identifier.
	SourceURL    string  `json:"sourceURL"`              // Stylesheet resource URL.
	SourceMapURL string  `json:"sourceMapURL,omitempty"` // URL of source map associated with the stylesheet (if any).
	Origin       string  `json:"origin"`                 // Stylesheet origin. enum values: injected, user-agent, inspector, regular
	Title        string  `json:"title"`                  // Stylesheet title.
	OwnerNode    int     `json:"ownerNode,omitempty"`    // The backend id for the owner node of the stylesheet.
	Disabled     bool    `json:"disabled"`               // Denotes whether the stylesheet is disabled.
	HasSourceURL bool    `json:"hasSourceURL,omitempty"` // Whether the sourceURL field value comes from the sourceURL comment.
	IsInline     bool    `json:"isInline"`               // Whether this stylesheet is created for STYLE tag by parser. This flag is not set for document.written STYLE tags.
	StartLine    float64 `json:"startLine"`              // Line offset of the stylesheet within the resource (zero based).
	StartColumn  float64 `json:"startColumn"`            // Column offset of the stylesheet within the resource (zero based).
	Length       float64 `json:"length"`                 // Size of the content (in characters).
}

// CSS rule representation.
type CSSCSSRule struct {
	StyleSheetId string           `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	SelectorList *CSSSelectorList `json:"selectorList"`           // Rule selector data.
	Origin       string           `json:"origin"`                 // Parent stylesheet's origin. enum values: injected, user-agent, inspector, regular
	Style        *CSSCSSStyle     `json:"style"`                  // Associated style declaration.
	Media        []*CSSCSSMedia   `json:"media,omitempty"`        // Media list array (for rules involving media queries). The array enumerates media queries starting with the innermost one, going outwards.
}

// CSS coverage information.
type CSSRuleUsage struct {
	StyleSheetId string  `json:"styleSheetId"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	StartOffset  float64 `json:"startOffset"`  // Offset of the start of the rule (including selector) from the beginning of the stylesheet.
	EndOffset    float64 `json:"endOffset"`    // Offset of the end of the rule body from the beginning of the stylesheet.
	Used         bool    `json:"used"`         // Indicates whether the rule was actually used by some element in the page.
}

// Text range within a resource. All numbers are zero-based.
type CSSSourceRange struct {
	StartLine   int `json:"startLine"`   // Start line of range.
	StartColumn int `json:"startColumn"` // Start column of range (inclusive).
	EndLine     int `json:"endLine"`     // End line of range
	EndColumn   int `json:"endColumn"`   // End column of range (exclusive).
}

// No Description.
type CSSShorthandEntry struct {
	Name      string `json:"name"`                // Shorthand name.
	Value     string `json:"value"`               // Shorthand value.
	Important bool   `json:"important,omitempty"` // Whether the property has "!important" annotation (implies <code>false</code> if absent).
}

// No Description.
type CSSCSSComputedStyleProperty struct {
	Name  string `json:"name"`  // Computed style property name.
	Value string `json:"value"` // Computed style property value.
}

// CSS style representation.
type CSSCSSStyle struct {
	StyleSheetId     string               `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	CssProperties    []*CSSCSSProperty    `json:"cssProperties"`          // CSS properties in the style.
	ShorthandEntries []*CSSShorthandEntry `json:"shorthandEntries"`       // Computed values for all shorthands found in the style.
	CssText          string               `json:"cssText,omitempty"`      // Style declaration text (if available).
	Range            *CSSSourceRange      `json:"range,omitempty"`        // Style declaration range in the enclosing stylesheet (if available).
}

// CSS property declaration data.
type CSSCSSProperty struct {
	Name      string          `json:"name"`                // The property name.
	Value     string          `json:"value"`               // The property value.
	Important bool            `json:"important,omitempty"` // Whether the property has "!important" annotation (implies <code>false</code> if absent).
	Implicit  bool            `json:"implicit,omitempty"`  // Whether the property is implicit (implies <code>false</code> if absent).
	Text      string          `json:"text,omitempty"`      // The full property text as specified in the style.
	ParsedOk  bool            `json:"parsedOk,omitempty"`  // Whether the property is understood by the browser (implies <code>true</code> if absent).
	Disabled  bool            `json:"disabled,omitempty"`  // Whether the property is disabled by the user (present for source-based properties only).
	Range     *CSSSourceRange `json:"range,omitempty"`     // The entire property range in the enclosing style declaration (if available).
}

// CSS media rule descriptor.
type CSSCSSMedia struct {
	Text         string           `json:"text"`                   // Media query text.
	Source       string           `json:"source"`                 // Source of the media query: "mediaRule" if specified by a @media rule, "importRule" if specified by an @import rule, "linkedSheet" if specified by a "media" attribute in a linked stylesheet's LINK tag, "inlineSheet" if specified by a "media" attribute in an inline stylesheet's STYLE tag.
	SourceURL    string           `json:"sourceURL,omitempty"`    // URL of the document containing the media query description.
	Range        *CSSSourceRange  `json:"range,omitempty"`        // The associated rule (@media or @import) header range in the enclosing stylesheet (if available).
	StyleSheetId string           `json:"styleSheetId,omitempty"` // Identifier of the stylesheet containing this object (if exists).
	MediaList    []*CSSMediaQuery `json:"mediaList,omitempty"`    // Array of media queries.
}

// Media query descriptor.
type CSSMediaQuery struct {
	Expressions []*CSSMediaQueryExpression `json:"expressions"` // Array of media query expressions.
	Active      bool                       `json:"active"`      // Whether the media query condition is satisfied.
}

// Media query expression descriptor.
type CSSMediaQueryExpression struct {
	Value          float64         `json:"value"`                    // Media query expression value.
	Unit           string          `json:"unit"`                     // Media query expression units.
	Feature        string          `json:"feature"`                  // Media query expression feature.
	ValueRange     *CSSSourceRange `json:"valueRange,omitempty"`     // The associated range of the value text in the enclosing stylesheet (if available).
	ComputedLength float64         `json:"computedLength,omitempty"` // Computed length of media query expression (if applicable).
}

// Information about amount of glyphs that were rendered with given font.
type CSSPlatformFontUsage struct {
	FamilyName   string  `json:"familyName"`   // Font's family name reported by platform.
	IsCustomFont bool    `json:"isCustomFont"` // Indicates if the font was downloaded or resolved locally.
	GlyphCount   float64 `json:"glyphCount"`   // Amount of glyphs that were rendered with this font.
}

// CSS keyframes rule representation.
type CSSCSSKeyframesRule struct {
	AnimationName *CSSValue             `json:"animationName"` // Animation name.
	Keyframes     []*CSSCSSKeyframeRule `json:"keyframes"`     // List of keyframes.
}

// CSS keyframe rule representation.
type CSSCSSKeyframeRule struct {
	StyleSheetId string       `json:"styleSheetId,omitempty"` // The css style sheet identifier (absent for user agent stylesheet and user-specified stylesheet rules) this rule came from.
	Origin       string       `json:"origin"`                 // Parent stylesheet's origin. enum values: injected, user-agent, inspector, regular
	KeyText      *CSSValue    `json:"keyText"`                // Associated key text.
	Style        *CSSCSSStyle `json:"style"`                  // Associated style declaration.
}

// A descriptor of operation to mutate style declaration text.
type CSSStyleDeclarationEdit struct {
	StyleSheetId string          `json:"styleSheetId"` // The css style sheet identifier.
	Range        *CSSSourceRange `json:"range"`        // The range of the style text in the enclosing stylesheet.
	Text         string          `json:"text"`         // New style text.
}

// Details of post layout rendered text positions. The exact layout should not be regarded as stable and may change between versions.
type CSSInlineTextBox struct {
	BoundingBox         *DOMRect `json:"boundingBox"`         // The absolute position bounding box.
	StartCharacterIndex int      `json:"startCharacterIndex"` // The starting index in characters, for this post layout textbox substring.
	NumCharacters       int      `json:"numCharacters"`       // The number of characters in this post layout textbox substring.
}

// Details of an element in the DOM tree with a LayoutObject.
type CSSLayoutTreeNode struct {
	NodeId          int                 `json:"nodeId"`                    // The id of the related DOM node matching one from DOM.GetDocument.
	BoundingBox     *DOMRect            `json:"boundingBox"`               // The absolute position bounding box.
	LayoutText      string              `json:"layoutText,omitempty"`      // Contents of the LayoutText if any
	InlineTextNodes []*CSSInlineTextBox `json:"inlineTextNodes,omitempty"` // The post layout inline text nodes, if any.
	StyleIndex      int                 `json:"styleIndex,omitempty"`      // Index into the computedStyles array returned by getLayoutTreeAndStyles.
}

// A subset of the full ComputedStyle as defined by the request whitelist.
type CSSComputedStyle struct {
	Properties []*CSSCSSComputedStyleProperty `json:"properties"` //
}

// Fired whenever a stylesheet is changed as a result of the client operation.
type CSSStyleSheetChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		StyleSheetId string `json:"styleSheetId"` //
	} `json:"Params,omitempty"`
}

// Fired whenever an active document stylesheet is added.
type CSSStyleSheetAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		Header *CSSCSSStyleSheetHeader `json:"header"` // Added stylesheet metainfo.
	} `json:"Params,omitempty"`
}

// Fired whenever an active document stylesheet is removed.
type CSSStyleSheetRemovedEvent struct {
	Method string `json:"method"`
	Params struct {
		StyleSheetId string `json:"styleSheetId"` // Identifier of the removed stylesheet.
	} `json:"Params,omitempty"`
}

type CSS struct {
	target gcdmessage.ChromeTargeter
}

func NewCSS(target gcdmessage.ChromeTargeter) *CSS {
	c := &CSS{target: target}
	return c
}

// Enables the CSS agent for the given page. Clients should not assume that the CSS agent has been enabled until the result of this command is received.
func (c *CSS) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.enable"})
}

// Disables the CSS agent for the given page.
func (c *CSS) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.disable"})
}

// GetMatchedStylesForNode - Returns requested styles for a DOM node identified by <code>nodeId</code>.
// nodeId -
// Returns -  inlineStyle - Inline style for the specified DOM node. attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%"). matchedCSSRules - CSS rules matching this node, from all applicable stylesheets. pseudoElements - Pseudo style matches for this node. inherited - A chain of inherited styles (from the immediate node parent up to the DOM tree root). cssKeyframesRules - A list of CSS keyframed animations matching this node.
func (c *CSS) GetMatchedStylesForNode(nodeId int) (*CSSCSSStyle, *CSSCSSStyle, []*CSSRuleMatch, []*CSSPseudoElementMatches, []*CSSInheritedStyleEntry, []*CSSCSSKeyframesRule, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getMatchedStylesForNode", Params: paramRequest})
	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			InlineStyle       *CSSCSSStyle
			AttributesStyle   *CSSCSSStyle
			MatchedCSSRules   []*CSSRuleMatch
			PseudoElements    []*CSSPseudoElementMatches
			Inherited         []*CSSInheritedStyleEntry
			CssKeyframesRules []*CSSCSSKeyframesRule
		}
	}

	if resp == nil {
		return nil, nil, nil, nil, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	return chromeData.Result.InlineStyle, chromeData.Result.AttributesStyle, chromeData.Result.MatchedCSSRules, chromeData.Result.PseudoElements, chromeData.Result.Inherited, chromeData.Result.CssKeyframesRules, nil
}

// GetInlineStylesForNode - Returns the styles defined inline (explicitly in the "style" attribute and implicitly, using DOM attributes) for a DOM node identified by <code>nodeId</code>.
// nodeId -
// Returns -  inlineStyle - Inline style for the specified DOM node. attributesStyle - Attribute-defined element style (e.g. resulting from "width=20 height=100%").
func (c *CSS) GetInlineStylesForNode(nodeId int) (*CSSCSSStyle, *CSSCSSStyle, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getInlineStylesForNode", Params: paramRequest})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			InlineStyle     *CSSCSSStyle
			AttributesStyle *CSSCSSStyle
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.InlineStyle, chromeData.Result.AttributesStyle, nil
}

// GetComputedStyleForNode - Returns the computed style for a DOM node identified by <code>nodeId</code>.
// nodeId -
// Returns -  computedStyle - Computed style for the specified DOM node.
func (c *CSS) GetComputedStyleForNode(nodeId int) ([]*CSSCSSComputedStyleProperty, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getComputedStyleForNode", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			ComputedStyle []*CSSCSSComputedStyleProperty
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.ComputedStyle, nil
}

// GetPlatformFontsForNode - Requests information about platform fonts which we used to render child TextNodes in the given node.
// nodeId -
// Returns -  fonts - Usage statistics for every employed platform font.
func (c *CSS) GetPlatformFontsForNode(nodeId int) ([]*CSSPlatformFontUsage, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getPlatformFontsForNode", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Fonts []*CSSPlatformFontUsage
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Fonts, nil
}

// GetStyleSheetText - Returns the current textual content and the URL for a stylesheet.
// styleSheetId -
// Returns -  text - The stylesheet text.
func (c *CSS) GetStyleSheetText(styleSheetId string) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["styleSheetId"] = styleSheetId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getStyleSheetText", Params: paramRequest})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Text string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.Text, nil
}

// CollectClassNames - Returns all class names from specified stylesheet.
// styleSheetId -
// Returns -  classNames - Class name list.
func (c *CSS) CollectClassNames(styleSheetId string) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["styleSheetId"] = styleSheetId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.collectClassNames", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			ClassNames []string
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.ClassNames, nil
}

// SetStyleSheetText - Sets the new stylesheet text.
// styleSheetId -
// text -
// Returns -  sourceMapURL - URL of source map associated with script (if any).
func (c *CSS) SetStyleSheetText(styleSheetId string, text string) (string, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["text"] = text
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setStyleSheetText", Params: paramRequest})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			SourceMapURL string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.SourceMapURL, nil
}

// SetRuleSelector - Modifies the rule selector.
// styleSheetId -
// range -
// selector -
// Returns -  selectorList - The resulting selector list after modification.
func (c *CSS) SetRuleSelector(styleSheetId string, theRange *CSSSourceRange, selector string) (*CSSSelectorList, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["range"] = theRange
	paramRequest["selector"] = selector
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setRuleSelector", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			SelectorList *CSSSelectorList
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.SelectorList, nil
}

// SetKeyframeKey - Modifies the keyframe rule key text.
// styleSheetId -
// range -
// keyText -
// Returns -  keyText - The resulting key text after modification.
func (c *CSS) SetKeyframeKey(styleSheetId string, theRange *CSSSourceRange, keyText string) (*CSSValue, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["range"] = theRange
	paramRequest["keyText"] = keyText
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setKeyframeKey", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			KeyText *CSSValue
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.KeyText, nil
}

// SetStyleTexts - Applies specified style edits one after another in the given order.
// edits -
// Returns -  styles - The resulting styles after modification.
func (c *CSS) SetStyleTexts(edits []*CSSStyleDeclarationEdit) ([]*CSSCSSStyle, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["edits"] = edits
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setStyleTexts", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Styles []*CSSCSSStyle
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Styles, nil
}

// SetMediaText - Modifies the rule selector.
// styleSheetId -
// range -
// text -
// Returns -  media - The resulting CSS media rule after modification.
func (c *CSS) SetMediaText(styleSheetId string, theRange *CSSSourceRange, text string) (*CSSCSSMedia, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["range"] = theRange
	paramRequest["text"] = text
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setMediaText", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Media *CSSCSSMedia
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Media, nil
}

// CreateStyleSheet - Creates a new special "via-inspector" stylesheet in the frame with given <code>frameId</code>.
// frameId - Identifier of the frame where "via-inspector" stylesheet should be created.
// Returns -  styleSheetId - Identifier of the created "via-inspector" stylesheet.
func (c *CSS) CreateStyleSheet(frameId string) (string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["frameId"] = frameId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.createStyleSheet", Params: paramRequest})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			StyleSheetId string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.StyleSheetId, nil
}

// AddRule - Inserts a new rule with the given <code>ruleText</code> in a stylesheet with given <code>styleSheetId</code>, at the position specified by <code>location</code>.
// styleSheetId - The css style sheet identifier where a new rule should be inserted.
// ruleText - The text of a new rule.
// location - Text position of a new rule in the target style sheet.
// Returns -  rule - The newly created rule.
func (c *CSS) AddRule(styleSheetId string, ruleText string, location *CSSSourceRange) (*CSSCSSRule, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["styleSheetId"] = styleSheetId
	paramRequest["ruleText"] = ruleText
	paramRequest["location"] = location
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.addRule", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Rule *CSSCSSRule
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Rule, nil
}

// ForcePseudoState - Ensures that the given node will have specified pseudo-classes whenever its style is computed by the browser.
// nodeId - The element id for which to force the pseudo state.
// forcedPseudoClasses - Element pseudo classes to force when computing the element's style.
func (c *CSS) ForcePseudoState(nodeId int, forcedPseudoClasses []string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 2)
	paramRequest["nodeId"] = nodeId
	paramRequest["forcedPseudoClasses"] = forcedPseudoClasses
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.forcePseudoState", Params: paramRequest})
}

// GetMediaQueries - Returns all media queries parsed by the rendering engine.
// Returns -  medias -
func (c *CSS) GetMediaQueries() ([]*CSSCSSMedia, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getMediaQueries"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Medias []*CSSCSSMedia
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Medias, nil
}

// SetEffectivePropertyValueForNode - Find a rule with the given active property for the given node and set the new value for this property
// nodeId - The element id for which to set property.
// propertyName -
// value -
func (c *CSS) SetEffectivePropertyValueForNode(nodeId int, propertyName string, value string) (*gcdmessage.ChromeResponse, error) {
	paramRequest := make(map[string]interface{}, 3)
	paramRequest["nodeId"] = nodeId
	paramRequest["propertyName"] = propertyName
	paramRequest["value"] = value
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.setEffectivePropertyValueForNode", Params: paramRequest})
}

// GetBackgroundColors -
// nodeId - Id of the node to get background colors for.
// Returns -  backgroundColors - The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load).
func (c *CSS) GetBackgroundColors(nodeId int) ([]string, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["nodeId"] = nodeId
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getBackgroundColors", Params: paramRequest})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			BackgroundColors []string
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.BackgroundColors, nil
}

// GetLayoutTreeAndStyles - For the main document and any content documents, return the LayoutTreeNodes and a whitelisted subset of the computed style. It only returns pushed nodes, on way to pull all nodes is to call DOM.getDocument with a depth of -1.
// computedStyleWhitelist - Whitelist of computed styles to return.
// Returns -  layoutTreeNodes -  computedStyles -
func (c *CSS) GetLayoutTreeAndStyles(computedStyleWhitelist []string) ([]*CSSLayoutTreeNode, []*CSSComputedStyle, error) {
	paramRequest := make(map[string]interface{}, 1)
	paramRequest["computedStyleWhitelist"] = computedStyleWhitelist
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.getLayoutTreeAndStyles", Params: paramRequest})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			LayoutTreeNodes []*CSSLayoutTreeNode
			ComputedStyles  []*CSSComputedStyle
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.LayoutTreeNodes, chromeData.Result.ComputedStyles, nil
}

// Enables the selector recording.
func (c *CSS) StartRuleUsageTracking() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.startRuleUsageTracking"})
}

// TakeCoverageDelta - Obtain list of rules that became used since last call to this method (or since start of coverage instrumentation)
// Returns -  coverage -
func (c *CSS) TakeCoverageDelta() ([]*CSSRuleUsage, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.takeCoverageDelta"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Coverage []*CSSRuleUsage
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Coverage, nil
}

// StopRuleUsageTracking - The list of rules with an indication of whether these were used
// Returns -  ruleUsage -
func (c *CSS) StopRuleUsageTracking() ([]*CSSRuleUsage, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CSS.stopRuleUsageTracking"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			RuleUsage []*CSSRuleUsage
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.RuleUsage, nil
}
