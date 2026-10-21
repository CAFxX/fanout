package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint64 = 2
    v2 := []string{"\n\t", "hello", ""}
    var v3 float64 = 0.0
    var v4 string = "hello"
    if (!false) {
        fmt.Println(v1, v1)
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi5 := 0; fi5 < 7; fi5++ {
                var v6 int64 = 0
                if (1 < 255) {
                    v7 := GenS1{F0: 127}
                    switch (v4 + (((("" + v4) + "a") + "foo bar") + "a")) {
                    case "a":
                        if (!((^(fi5)) <= fi5)) {
                            var v8 int = (fi5 - 270)
                            switch 327 {
                            case 466:
                                var v9 bool = false
                                _ = v9
                            case -918:
                                for fi10 := 0; fi10 < 3; fi10++ {
                                    var v11 uint32 = 0
                                    _ = fi10
                                    _ = v11
                                }
                            default:
                                var v12 float64 = v3
                                _ = v12
                            }
                            _ = v8
                        }
                    }
                    if len(v2) > 0 {
                        v2[0] = (v4 + v4)
                    }
                    _ = v7
                } else {
                    var v13 int8 = 127
                    _ = v13
                }
                _ = fi5
                _ = v6
            }
        } else {
            v2 = append(v2, "foo bar")
        }
    }
    v2 = append(v2, v4)
    fmt.Println(v3, v4, v3)
    var v14 uint8 = 13
    v14 += v14
    for fi15 := 0; fi15 < 9; fi15++ {
        v4 = v4
        _ = fi15
    }
    var v16 int64 = -257
    v16 = (-(v16))
    var v17 uint8 = 132
    v16 = (v16 * v16)
    v14 = (v14 % ((42) | 1))
    fmt.Println(v3, v1)
    var v18 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
