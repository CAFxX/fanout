package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 bool = (!false)
    v4 := []string{"hello"}
    var v5 uint32 = 2147483648
    v6 := GenS1{F0: 0}
    _ = len(v2)
    v7 := GenS2{F0: 42}
    var v8 bool = (!v3)
    fmt.Println(v5, v3)
    v8 = (!v8)
    v5 = 479
    fmt.Println(v5, v5, v5)
    switch "" {
    case "foo bar":
        fmt.Println(v5)
    default:
        for fi9 := 0; fi9 < 7; fi9++ {
            if (!true) {
                switch ("foo bar" + "\n\t") {
                case "a":
                    switch "\n\t" {
                    case "hello":
                        var v10 uint16 = 255
                        _ = v10
                    default:
                        var v11 bool = (((0.1 + (0.1 * (0.1 - ((3.14159 / float32(1e10)) - (1e10 - float32((0.0 * 3.14159))))))) < 3.14159) && (true || ((^(v5)) > v5)))
                        _ = v11
                    }
                case "":
                    if (238 <= 1) {
                        switch ("" + "\n\t") {
                        case "\n\t":
                            var v12 bool = true
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 14; fi13++ {
                            v14 := GenS2{F0: 42}
                            var v15 uint64 = 79
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    }
                default:
                    fmt.Println(v3, v3)
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi16 := 0; fi16 < 16; fi16++ {
                    switch 394 {
                    case 32768:
                        if len(v4) > 0 {
                            v4[0] = ("a" + "")
                        }
                    case 42:
                        if (!(!v8)) {
                            _ = len(v2)
                            var v17 float64 = float64(3.14159)
                            v5 -= (v5 ^ (^(v5)))
                            _ = v17
                        }
                    case -1:
                        v4 = append(v4, "hello")
                    }
                    _ = fi16
                }
            } else {
                var v18 string = "hello"
                var v19 float64 = (3.14159 * 1e10)
                var v20 uint32 = v5
                _ = v18
                _ = v19
                _ = v20
            }
            _ = fi9
        }
    }
    var v21 uint16 = 65535
    var v22 string = "a"
    v23 := GenS1{F0: 2}
    v4 = append(v4, v22)
    var v24 int8 = -1
    var v25 int64 = 42
    fmt.Println(v8, v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
