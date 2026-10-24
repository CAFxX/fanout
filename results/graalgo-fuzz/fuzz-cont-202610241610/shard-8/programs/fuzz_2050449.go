package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int16) int {
    p0 -= (^(p0))
    var v4 uint16 = 407
    _ = p0
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 string = "hello"
    var v2 int32 = 464
    v3 := []string{}
    var v5 int64 = -1
    switch 228 {
    case 888:
        switch "foo bar" {
        case "foo bar":
            var v6 int32 = (v2 - (v2 << ((173) & 63)))
            _ = v6
        case "":
            v1 = v1
        default:
            if false {
                var v7 uint = 65535
                _ = v7
            } else {
                var v8 int16 = 32767
                fmt.Println(v1, v2, v2)
                var v9 uint64 = 4294967295
                _ = v8
                _ = v9
            }
        }
    default:
        for fi10 := 0; fi10 < 17; fi10++ {
            v5 = (v5 | -402)
            fmt.Println(v5, fi10)
            _ = fi10
        }
    }
    v5 = v5
    if (((v1 + "foo bar") + (v1 + (((v1 + (v1 + v1)) + (("foo bar" + (((v1 + "hello") + "foo bar") + v1)) + v1)) + v1))) >= "hello") {
        v2 -= v2
        fmt.Println(v2, v2)
    }
    v5 -= v5
    if len(v3) > 0 {
        v3[0] = (v1 + (v1 + v1))
    }
    var v11 uint = 1
    if len(v3) > 0 {
        v3[0] = ("foo bar" + v1)
    }
    _ = len(v3)
    var v12 uint8 = 0
    if (false || true) {
        v5 -= v5
        _ = len(v3)
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if false {
        var v13 int32 = (-(v2))
        for fi14 := 0; fi14 < 19; fi14++ {
            v3 = append(v3, v1)
            v2 += 115
            _ = fi14
        }
        _ = v13
    }
    v15 := GenS1{F0: (v2 & (^(v2)))}
    switch 127 {
    case -9223372036854775808:
        switch genFn0(int16(v11)) {
        case 128:
            if len(v3) > 0 {
                v3[0] = ((v1 + "") + ((v1 + v1) + ""))
            }
        case -9223372036854775808:
            _ = len(v3)
        }
    default:
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 int8 = 28
            if true {
                var v18 uint16 = 513
                fmt.Println(v1)
                _ = v18
            } else {
                var v19 float32 = (0.0 * ((0.1 * 1e10) + float32((3.14159 - (((0.1 + (((float64(1.5) / 1.5) / ((((0.0 * 3.14159) + 3.14159) * 0.1) - (0.0 - 3.14159))) * float64((1.5 * ((-2.25 / (3.14159 - 0.0)) * -2.25))))) * (0.0 / 1.0)) * 1.5)))))
                _ = v19
            }
            if true {
                for fi20 := 0; fi20 < 9; fi20++ {
                    v12 -= (^(v12))
                    var v21 int8 = v17
                    _ = fi20
                    _ = v21
                }
            }
            _ = fi16
            _ = v17
        }
    }
    v22 := GenS1{F0: (-(v2))}
    if len(v3) > 0 {
        v3[0] = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v5
}
