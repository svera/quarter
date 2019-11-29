/*
Package level deals with the reading, organisation and rendering of a level on screen.
The following is an example of a level file, stored as JSON:

    {
        "version": "1",
        "levels": [
            {
                "name": "level 1",
                "layers": [
                    {
                        "name": "background",
                        "image": {
                            "path": "",
                            "extra": {}
                        },
                        "grid": {
                            "assets": {
                                "path": "",
                                "quantity": 5,
                                "offset": {
                                    "x": 0,
                                    "y": 0
                                }
                                "width": 16,
                                "height": 16
                            },
                            "tiles": [
                                {
                                    "asset": 0,
                                    "x": 10,
                                    "y": 10,
                                    "extra": ""
                                }
                            ],
                            "extra": {}
                        },
                        "bounds": [
                            {
                                "type": "box",
                                "dimensions": {
                                    "x": 10,
                                    "y": 10,
                                    "width": 20,
                                    "height": 20
                                },
                                "extra": ""
                            }
                        ]
                    }
                ],
                "extra": ""
            }
        ]
    }
*/
package level
