package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int8 }
func genFn0() uint64 {
    var v4 float32 = 0.0
    v4 -= 1e10
    var v5 uint = 255
    v5 = v5
    var v6 int = -946
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 int, p1 int16, p2 uint32) string {
    switch "a" {
    case "":
        var v7 string = "hello"
        _ = v7
    case "hello":
        p1 += 32767
    default:
        var v8 uint = 65535
        _ = v8
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}

func main() {
    v1 := []uint32{4294967295}
    var v2 float32 = 1.5
    var v3 uint8 = 7
    switch 127 {
    case 9223372036854775807:
        if (false || true) {
            fmt.Println(v3, v3)
            fmt.Println(v2, v3)
            var v9 float32 = v2
            _ = v9
        } else {
            if ((25 < 1) && false) {
                v3 = 127
            }
        }
    case 0:
        switch 32768 {
        case 32768:
            v2 = 0.0
        }
    case 128:
        var v10 int16 = 978
        _ = v10
    }
    v2 = (v2 * (v2 + 3.14159))
    if len(v1) > 0 {
        v1[0] = 209
    }
    v1 = append(v1, 4294967295)
    v2 = -2.25
    v3 += v3
    if true {
        for fi11 := 0; fi11 < 11; fi11++ {
            fmt.Println(v3)
            v12 := GenS1{F0: uint32(v3)}
            v13 := GenS1{F0: 2147483647}
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 uint32 = 2147483647
        _ = v14
    } else {
        switch ("" + "\n\t") {
        case "":
            if true {
                for fi15 := 0; fi15 < 10; fi15++ {
                    if (((false || (0 < -2147483648)) || false) || ((250 > v3) || (!true))) {
                        v1 = append(v1, 1)
                    }
                    _ = fi15
                }
                v16 := GenS1{F0: 386}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v16
            }
        case "hello":
            if false {
                v17 := GenS2{F0: false, F1: 127}
                switch "a" {
                case "":
                    var v18 uint32 = 2
                    _ = v18
                case "hello":
                    fmt.Println(v2, v3)
                }
                var v19 string = "hello"
                _ = v17
                _ = v19
            }
        case "\n\t":
            v3 = (^(v3))
        }
    }
    _ = len(v1)
    v2 += float32(float64(((v2 / (v2 / ((float32(0.0) - -2.25) * (((v2 * v2) * (0.0 / ((((1e10 * 0.0) - v2) * ((v2 + v2) - (0.0 * v2))) + (0.1 - v2)))) + (v2 + v2))))) / -2.25)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
