package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int32 }
func genFn0(p0 int64) bool {
    v5 := GenS1{F0: 290}
    for fi6 := 0; fi6 < 19; fi6++ {
        p0 = -9223372036854775808
        _ = fi6
    }
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 float64 = 0.1
        p0 = -141
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v7
    return false
}
func genFn1(p0 bool, p1 uint16, p2 int8) uint64 {
    if p0 {
        var v10 bool = p0
        if p0 {
            switch 1 {
            case 860:
                p0 = true
            case -9223372036854775808:
                var v11 int8 = (p2 ^ p2)
                _ = v11
            }
            for fi12 := 0; fi12 < 15; fi12++ {
                v10 = ((-2.25 + 3.14159) < 1.5)
                p2 += (-(p2))
                _ = fi12
            }
        }
        var v13 uint32 = 4294967295
        _ = v10
        _ = v13
    } else {
        for fi14 := 0; fi14 < 8; fi14++ {
            p0 = p0
            _ = fi14
        }
        v15 := []int{-2147483648, 4294967295, -522, 32768}
        _ = v15
    }
    if p0 {
        v16 := []int{4294967295, -396, -2147483648, int(p1)}
        if p0 {
            p0 = p0
            for fi17 := 0; fi17 < 17; fi17++ {
                v18 := GenS1{F0: 65535}
                _ = fi17
                _ = v18
            }
        }
        _ = v16
    }
    v19 := GenS1{F0: 225}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 146
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]bool)
    v3 := []int{}
    var v4 float64 = ((3.14159 * (1.5 + -2.25)) * 1e10)
    v20 := GenS1{F0: 721}
    if (1.5 >= v4) {
        _ = len(v2)
        var v21 int = 882
        v21 = v21
        _ = v21
    }
    switch 128 {
    case 0:
        v4 += float64(v1)
    }
    if true {
        switch "foo bar" {
        case "hello":
            v22 := v2[("hello" + "hello")]
            _ = v22
        case "\n\t":
            v1 = (1.5 / (float32(v4) * 3.14159))
        default:
            for fi23 := 0; fi23 < 15; fi23++ {
                _ = len(v2)
                fmt.Println(v4, v4, fi23)
                v1 = v1
                _ = fi23
            }
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            v25 := GenS1{F0: 65535}
            if (2 == 239) {
                fmt.Println(v1)
                if (-2.25 != ((v1 / (v1 - (v1 + v1))) + 1e10)) {
                    v26 := GenS2{F0: genFn1(false, 255, 1), F1: -2147483648}
                    v4 = v4
                    _ = v26
                }
            } else {
                for fi27 := 0; fi27 < 1; fi27++ {
                    v28 := GenS2{F0: genFn1((0 >= 0), 50, 0), F1: -119}
                    _ = fi27
                    _ = v28
                }
                v4 = (-2.25 - ((-2.25 + (0.0 * ((3.14159 / (1.5 + 0.1)) - v4))) / v4))
                if len(v3) > 0 {
                    v3[0] = (fi24 + (fi24 + (-(fi24))))
                }
            }
            var v29 float64 = 0.1
            _ = fi24
            _ = v25
            _ = v29
        }
        v1 += (v1 - float32(-2.25))
    }
    switch "hello" {
    case "foo bar":
        v1 -= -2.25
    case "\n\t":
        v1 = 1e10
    }
    if true {
        if false {
            switch "foo bar" {
            case "hello":
                for fi30 := 0; fi30 < 16; fi30++ {
                    delete(v2, "a")
                    _ = fi30
                }
            case "foo bar":
                var v31 bool = (2147483647 != 523)
                _ = v31
            default:
                var v32 int8 = 1
                _ = v32
            }
        } else {
            v2["hello"] = true
        }
        var v33 uint32 = 1
        _ = v33
    }
    fmt.Println(v1, v4)
    var v34 int32 = -166
    var v35 float64 = v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
