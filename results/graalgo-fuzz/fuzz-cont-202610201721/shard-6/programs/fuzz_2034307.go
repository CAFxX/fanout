package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 string; F2 int32 }
type GenS3 struct { F0 bool; F1 string; F2 uint64 }
func genFn0(p0 uint8, p1 uint64) int32 {
    v6 := GenS1{F0: -55}
    if ((-2.25 / 1e10) <= -2.25) {
        p1 += (p1 & (p1 & (p1 >> ((uint(p0)) & 63))))
        v7 := GenS1{F0: -1}
        _ = v7
    } else {
        var v8 uint64 = p1
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int16 = -651
            _ = fi9
            _ = v10
        }
        v11 := []int{4294967295, 0, 1}
        _ = v11
        _ = v8
    }
    if (0 <= 42) {
        p1 = p1
    }
    var v12 string = "\n\t"
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    return 377
}

func main() {
    var v1 float32 = (1e10 + ((((((0.1 + ((0.0 * 3.14159) / 1e10)) + 0.0) * (((-2.25 * (1e10 - 0.0)) * (((float32(3.14159) * (3.14159 / 0.1)) + 1e10) - float32(3.14159))) * 3.14159)) / (3.14159 * ((-2.25 - 0.1) / (((0.1 + 0.1) + 1e10) - 3.14159)))) - 0.0) + 3.14159))
    v2 := make(map[string]string)
    v3 := []uint32{93, 246}
    v4 := []uint32{}
    var v5 float32 = v1
    _ = len(v3)
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS2{F0: 0, F1: "hello", F2: -2147483648}
        switch (fi13 & (-(fi13))) {
        case 124:
            v5 = v5
        default:
            fmt.Println(fi13, v1)
        }
        var v15 int16 = 0
        _ = fi13
        _ = v14
        _ = v15
    }
    v2[""] = ""
    if (1 == 0) {
        if true {
            v16 := GenS3{F0: (!(-1 <= 0)), F1: "foo bar", F2: 2}
            _ = v16
        }
        fmt.Println(v5, v1)
        var v17 int32 = 2147483647
        _ = v17
    } else {
        switch ("a" + "foo bar") {
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = 1
            }
        case "a":
            switch 1 {
            case 635:
                fmt.Println(v5)
            case -2147483648:
                v18 := GenS2{F0: 32767, F1: "hello", F2: -2147483648}
                _ = v18
            case 4294967295:
                var v19 int8 = -1
                _ = v19
            }
        case "":
            for fi20 := 0; fi20 < 20; fi20++ {
                if true {
                    v5 -= v5
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi20
            }
        default:
            var v21 uint64 = 502
            _ = v21
        }
        v22 := v2[""]
        _ = len(v3)
        _ = v22
    }
    v5 = float32(3.14159)
    var v23 string = "a"
    if true {
        for fi24 := 0; fi24 < 16; fi24++ {
            v1 -= v1
            var v25 int = fi24
            v26 := GenS2{F0: 1000, F1: ("foo bar" + v23), F2: 0}
            _ = fi24
            _ = v25
            _ = v26
        }
        v5 -= 3.14159
    }
    v5 = (1e10 / v5)
    for fi27 := 0; fi27 < 6; fi27++ {
        var v28 uint = 18446744073709551615
        fmt.Println(v23, fi27)
        _ = fi27
        _ = v28
    }
    var v29 uint = 950
    var v30 float64 = ((1e10 + 0.0) - 0.1)
    if (-511 <= 42) {
        v31 := GenS2{F0: 42, F1: ((v23 + v23) + ("hello" + (("foo bar" + "hello") + v23))), F2: -302}
        v23 = "\n\t"
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
