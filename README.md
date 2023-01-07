# nanoleaf_custom

This code is making an HTTP GET request to the URL "http://192.168.X.X:16021/api/v1/AUTHTOKEN/panelLayout/layout", and then unmarshalling the response body into a struct of type "Response". The "Response" struct has three fields: NumPanels, SideLength, and PositionData. The "PositionData" field is a slice of "PositionData" structs.

The "PositionData" struct has four fields: PanelID, X, Y, and O. The "for" loop iterates over each "PositionData" struct in the "PositionData" slice, and generates a random color for each panel. It then converts the random color values and the PanelID to strings, and makes an HTTP PUT request to the URL "http://192.168.X.X:16021/api/v1/AUTHTOKEN/effects" with a payload containing the panel ID, color values, and some other data.

The function "makeHTTPCall" is used to make the HTTP request with the specified method, URL, and payload. It creates a new HTTP client, creates a new HTTP request with the specified method, URL, and payload, sends the request, and reads the response body. If there are any errors, it will print the error and return.
