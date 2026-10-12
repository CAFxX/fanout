package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int16; F2 int32 }
func genFn0(p0 bool, p1 string, p2 uint) bool {
    p2 = p2
    var v4 float32 = (1.5 / float32((1e10 + -2.25)))
    v5 := GenS1{F0: ("\n\t" + (((p1 + ((p1 + "hello") + p1)) + p1) + p1))}
    v4 += ((((v4 * 1e10) / float32(3.14159)) + ((1e10 / (v4 - v4)) * (v4 - (v4 + (v4 - v4))))) - float32(3.14159))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p0
}

func main() {
    var v1 uint16 = 255
    v2 := []uint32{990, 332, 0}
    var v3 string = "\n\t"
    var v6 int8 = 42
    var v7 float32 = 3.14159
    v6 = -128
    switch ("a" + v3) {
    case "foo bar":
        var v8 int32 = -2147483648
        _ = v8
    default:
        if false {
            v9 := GenS1{F0: ("a" + "\n\t")}
            for fi10 := 0; fi10 < 18; fi10++ {
                v6 -= 23
                _ = fi10
            }
            _ = v9
        }
    }
    if true {
        if ((v3 < v3) || true) {
            switch 42 {
            case 2147483647:
                v11 := GenS2{F0: (v1 & v1), F1: -656, F2: -564}
                _ = v11
            case 128:
                var v12 float32 = (((v7 * (v7 / (((0.1 - (((v7 * -2.25) - 0.1) - 0.0)) / (v7 * (((0.0 - float32(0.0)) - (0.1 - v7)) * v7))) / v7))) - v7) - (v7 * 1.5))
                _ = v12
            case 127:
                if (-380 >= 128) {
                    v3 = v3
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                }
            }
            fmt.Println(v3, v7, v1)
        }
    } else {
        var v13 float64 = -2.25
        _ = v13
    }
    fmt.Println(v3, v3, v1)
    var v14 int8 = -1
    v1 -= v1
    var v15 int16 = 42
    fmt.Println(v7, v15, v3)
    if ((^(v1)) >= 615) {
        switch v3 {
        case "":
            v6 -= 0
        case "foo bar":
            v6 += v6
        case "\n\t":
            var v16 bool = (-9223372036854775808 > -261)
            _ = v16
        }
    }
    if false {
        v15 = 1000
        v6 -= (v6 + v14)
        var v17 int8 = v6
        _ = v17
    }
    switch (v3 + v3) {
    case "hello":
        v3 = (v3 + (("" + v3) + ""))
    case "a":
        var v18 uint32 = 4294967295
        _ = v18
    default:
        switch 0 {
        case 32767:
            switch "\n\t" {
            case "\n\t":
                if (false && true) {
                    v7 += float32(float64(v7))
                    v6 = v14
                    v15 += (v15 ^ v15)
                }
            case "a":
                var v19 uint = 820
                _ = v19
            case "hello":
                _ = len(v2)
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        default:
            v20 := GenS2{F0: v1, F1: v15, F2: 860}
            _ = v20
        }
    }
    v3 = ("" + "")
    v6 = v14
    v7 += ((float32(((float64(v7) * 0.0) + 3.14159)) / v7) * ((v7 + v7) / -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
