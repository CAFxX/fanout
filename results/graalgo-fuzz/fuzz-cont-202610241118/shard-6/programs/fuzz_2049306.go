package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 uint16; F1 int; F2 uint32 }
func genFn0(p0 int8, p1 uint64) bool {
    p0 -= (p0 ^ (p0 & p0))
    v6 := GenS1{F0: ("foo bar" + (("hello" + ("" + "")) + "a")), F1: 222, F2: 4294967295}
    fmt.Println(p1)
    p0 += -1
    var v7 int = 128
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (2 <= 2)
}

func main() {
    var v1 float32 = -2.25
    v2 := []string{("\n\t" + (("hello" + "") + ("foo bar" + "foo bar"))), "a", "", "\n\t"}
    var v3 int = -1
    v4 := make(map[string]bool)
    var v5 float64 = (1.5 - (3.14159 * -2.25))
    fmt.Println(v1, v3, v5)
    fmt.Println(v1, v5)
    if (float32(v5) <= float32((3.14159 - ((3.14159 / v5) * ((v5 + float64(v1)) / v5))))) {
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(fi8, fi8, v1)
            for fi9 := 0; fi9 < 11; fi9++ {
                v4[("foo bar" + (("a" + "") + "a"))] = (!false)
                v5 = (float64(v1) - v5)
                _ = fi9
            }
            _ = fi8
        }
    } else {
        fmt.Println(v3)
        if (!((v1 / (0.1 / (-2.25 * 1.5))) >= (v1 * (v1 / 3.14159)))) {
            if (!(42 <= 32767)) {
                for fi10 := 0; fi10 < 12; fi10++ {
                    v1 += v1
                    v11 := GenS2{F0: 65535, F1: (fi10 * fi10), F2: 1}
                    v5 -= -2.25
                    _ = fi10
                    _ = v11
                }
                var v12 int16 = 42
                if false {
                    for fi13 := 0; fi13 < 7; fi13++ {
                        var v14 float32 = v1
                        _ = fi13
                        _ = v14
                    }
                    var v15 int64 = 42
                    v1 -= v1
                    _ = v15
                }
                _ = v12
            }
            var v16 int = (v3 / ((v3) | 1))
            if false {
                var v17 int32 = 2147483647
                _ = v17
            }
            _ = v16
        }
        if (1e10 <= -2.25) {
            v18 := GenS1{F0: "a", F1: 65535, F2: 102}
            _ = v18
        } else {
            v19 := GenS2{F0: 255, F1: (v3 * v3), F2: 924}
            _ = v19
        }
    }
    var v20 uint8 = 143
    if false {
        fmt.Println(v3, v20)
        var v21 int64 = 9223372036854775807
        if (v21 < 0) {
            v21 -= (v21 % ((0) | 1))
        } else {
            var v22 uint = 255
            if ((!true) && (false || true)) {
                var v23 int64 = -1
                _ = v23
            }
            _ = v22
        }
        _ = v21
    }
    v1 -= 0.1
    if ((-1 != 579) && (!(-1 <= 127))) {
        switch ("hello" + "hello") {
        case "hello":
            switch "\n\t" {
            case "":
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            case "foo bar":
                var v24 uint = 2
                _ = v24
            }
        case "":
            v25 := GenS2{F0: 782, F1: (v3 | v3), F2: 194}
            _ = v25
        case "\n\t":
            fmt.Println(v5)
        }
        fmt.Println(v3, v5)
    }
    v26 := v4[""]
    fmt.Println(v3)
    v1 += v1
    v27 := GenS2{F0: 1, F1: v3, F2: 350}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
