package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int16; F2 int16 }

func main() {
    v1 := []uint32{}
    var v2 uint8 = 0
    var v3 int64 = -885
    v4 := make(map[string]uint32)
    if true {
        if (255 != 1) {
            v2 = (v2 * v2)
            v2 += (^(v2))
        }
    }
    v5 := v4["\n\t"]
    v6 := GenS1{F0: 725}
    v3 += v3
    v5 = (v5 + (^(v5)))
    v5 = 65535
    for fi7 := 0; fi7 < 17; fi7++ {
        if true {
            switch fi7 {
            case 701:
                v2 += 128
            case 423:
                if true {
                    v2 = (v2 % ((v2) | 1))
                    if (465 != 231) {
                        var v8 int32 = -1
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch "foo bar" {
                        case "a":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = v8
                    }
                }
            case 42:
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := GenS2{F0: -1, F1: 32767, F2: -32768}
                    v3 = v3
                    var v11 int64 = 1
                    v5 -= 42
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            default:
                var v12 int64 = -1
                _ = v12
            }
        } else {
            v2 -= (^(v2))
            var v13 uint = 2
            v13 = v13
            _ = v13
        }
        _ = fi7
    }
    v2 += v2
    delete(v4, "foo bar")
    var v14 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
