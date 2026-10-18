package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 string }

func main() {
    v1 := []uint32{2, 2147483647, 42, 767}
    var v2 string = "a"
    var v3 uint = 598
    v4 := []int{}
    var v5 uint = v3
    var v6 uint32 = 65535
    switch 127 {
    case -534:
        if false {
            v5 = 211
            var v7 int32 = 2147483647
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v7
        } else {
            fmt.Println(v3)
            v8 := GenS1{F0: 33}
            _ = v8
        }
    case -9223372036854775808:
        switch 394 {
        case 801:
            v6 = 669
        case 4294967295:
            switch v2 {
            case "\n\t":
                v4 = append(v4, -573)
            default:
                v6 = v6
            }
        }
    default:
        v2 = v2
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        v5 += v3
        _ = fi9
    }
    switch "hello" {
    case "foo bar":
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int16 = -32768
            if len(v4) > 0 {
                v4[0] = fi10
            }
            v6 -= v6
            _ = fi10
            _ = v11
        }
    default:
        if (!true) {
            v5 -= v3
            var v12 uint16 = 32767
            _ = v12
        } else {
            v13 := GenS1{F0: 8}
            if len(v1) > 0 {
                v1[0] = (v6 & v6)
            }
            _ = v13
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = "foo bar"
    v2 = ("hello" + v2)
    v5 = 0
    v6 += v6
    v2 = ""
    if ((float64((0.0 + 1e10)) >= 0.0) || (!false)) {
        var v14 int32 = 42
        _ = v14
    }
    var v15 string = v2
    v16 := GenS2{F0: (^(v6)), F1: ((v15 + ((v2 + v15) + (("hello" + ((v15 + v15) + "")) + (("a" + v2) + "")))) + v2)}
    for fi17 := 0; fi17 < 14; fi17++ {
        if true {
            v15 = (v15 + "")
        } else {
            switch (v15 + v15) {
            case "hello":
                v5 = (v5 ^ (^(v5)))
            default:
                switch fi17 {
                case -967:
                    var v18 uint64 = uint64(v6)
                    _ = v18
                default:
                    if ((((float32(0.1) / (0.0 + (0.0 + 1e10))) / -2.25) * 3.14159) == 0.1) {
                        var v19 uint16 = 1
                        _ = v19
                    }
                }
            }
            var v20 int16 = 399
            _ = len(v4)
            _ = v20
        }
        v6 = v6
        _ = fi17
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        v2 = (v15 + (v2 + v15))
        _ = fi21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
