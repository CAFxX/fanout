package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0(p0 int8, p1 string, p2 bool) uint8 {
    p2 = (!false)
    var v4 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 127
}

func main() {
    v1 := []uint32{1}
    var v2 uint64 = 326
    var v3 uint = 1
    fmt.Println(v3, v2)
    v2 -= v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        switch 861 {
        case -927:
            if (298 >= 2) {
                var v5 float32 = 0.1
                _ = v5
            }
        default:
            switch "" {
            case "\n\t":
                if len(v1) > 0 {
                    v1[0] = 2147483648
                }
            }
        }
    }
    switch ((("foo bar" + ((("a" + "hello") + "a") + "hello")) + "foo bar") + ("hello" + "foo bar")) {
    case "hello":
        v2 = (^(v2))
    case "":
        v6 := GenS1{F0: -93, F1: (v3 * v3)}
        _ = v6
    case "foo bar":
        var v7 uint8 = 38
        _ = v7
    }
    switch 127 {
    case -116:
        fmt.Println(v3, v3)
    case 42:
        if false {
            var v8 float32 = (float32(3.14159) + 0.1)
            _ = v8
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            v3 = (v3 - 53)
            var v11 bool = false
            v3 = (v3 % (((v3 % (((v3 * (v3 / ((uint(v2)) | 1)))) | 1))) | 1))
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 uint16 = 42
    if (!false) {
        var v13 float64 = 0.0
        _ = v13
    }
    var v14 float32 = -2.25
    v15 := GenS1{F0: int32(v2), F1: 960}
    v12 -= uint16(v3)
    if (!(!(!(!false)))) {
        if (!false) {
            var v16 uint32 = 65535
            _ = v16
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
