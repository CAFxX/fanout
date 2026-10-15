package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
func genFn0(p0 int32) uint32 {
    v4 := []int{32767}
    v4 = append(v4, 32767)
    _ = p0
    _ = v4
    return 2
}

func main() {
    v1 := []uint32{2147483648, 2147483647, 4294967295, 346}
    var v2 int8 = 127
    v3 := make(map[string]uint32)
    if (!(-86 == v2)) {
        if (!(((42 != -9223372036854775808) || (192 == -32768)) || (1e10 < (0.1 / 0.1)))) {
            v2 += v2
            var v5 int64 = 328
            fmt.Println(v5)
            _ = v5
        }
        switch (("" + "") + "") {
        case "foo bar":
            var v6 uint16 = 619
            _ = v6
        case "":
            for fi7 := 0; fi7 < 3; fi7++ {
                fmt.Println(v2, v2)
                var v8 float32 = float32((1.5 + (0.0 / float64(((0.1 * (float32((0.0 + float64(1e10))) - 3.14159)) * 3.14159)))))
                _ = fi7
                _ = v8
            }
        default:
            v9 := GenS1{F0: 1000, F1: 32767}
            _ = v9
        }
        var v10 int64 = 42
        _ = v10
    } else {
        for fi11 := 0; fi11 < 17; fi11++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v12 string = ("hello" + "foo bar")
            _ = fi11
            _ = v12
        }
        v3[("" + "")] = 2
    }
    var v13 uint = 327
    if false {
        v13 = 1
        switch ("\n\t" + "") {
        case "a":
            v14 := GenS1{F0: 0, F1: uint16(v13)}
            _ = v14
        default:
            fmt.Println(v2, v13)
        }
        var v15 string = ("\n\t" + ("" + ""))
        _ = v15
    }
    v2 -= v2
    var v16 uint = (v13 % ((v13) | 1))
    v1 = append(v1, 2147483647)
    fmt.Println(v13, v16, v13)
    switch 599 {
    case -120:
        switch 128 {
        case 1:
            if len(v1) > 0 {
                v1[0] = genFn0(960)
            }
        case -577:
            _ = len(v1)
        }
    case 218:
        var v17 int16 = 1
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
}
