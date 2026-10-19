package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint16 }
func genFn0(p0 uint32, p1 bool, p2 uint8) string {
    fmt.Println(p0)
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            fmt.Println(fi9)
            _ = fi9
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("hello" + ("" + "a"))
}
func genFn1(p0 bool) int8 {
    v10 := GenS1{F0: 2147483647, F1: 100, F2: 1}
    if p0 {
        var v11 uint16 = 637
        v12 := []int{290}
        _ = v11
        _ = v12
    } else {
        var v13 int32 = -352
        v14 := GenS1{F0: 621, F1: -1, F2: 764}
        _ = v13
        _ = v14
    }
    var v15 uint8 = 0
    var v16 string = (("\n\t" + "foo bar") + "hello")
    _ = p0
    _ = v10
    _ = v15
    _ = v16
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 uint8 = 2
    var v5 int = 4294967295
    var v6 float64 = -2.25
    var v17 float32 = 3.14159
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 float64 = (v6 * (v6 * v6))
        var v20 uint32 = 616
        v21 := GenS1{F0: 2, F1: -1, F2: 255}
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    if (!(343 < uint(v4))) {
        v22 := GenS1{F0: 401, F1: 1, F2: 194}
        var v23 int8 = -34
        v6 += float64((0.1 / (v17 - v17)))
        _ = v22
        _ = v23
    }
    if (false && true) {
        for fi24 := 0; fi24 < 16; fi24++ {
            v4 += (v4 / ((0) | 1))
            var v25 int32 = 793
            switch "foo bar" {
            case "foo bar":
                if (-9223372036854775808 <= -1) {
                    for fi26 := 0; fi26 < 11; fi26++ {
                        var v27 uint32 = 42
                        _ = fi26
                        _ = v27
                    }
                    var v28 uint = 987
                    _ = v28
                }
            }
            _ = fi24
            _ = v25
        }
        delete(v2, v5)
        v29 := []int{9223372036854775807, -2147483648, v5, v5}
        _ = v29
    }
    if (!(!true)) {
        v2[(v5 % ((v5) | 1))] = false
    }
    v30 := v3["a"]
    if (448 <= 255) {
        v31 := []int{v30}
        fmt.Println(v4, v5)
        switch v30 {
        case -391:
            var v32 string = ""
            _ = v32
        }
        _ = v31
    }
    v30 = -1
    fmt.Println(v30, v6, v5)
    var v33 int = (^(v5))
    var v34 uint8 = v4
    if ((65535 != 2) || (207 >= 991)) {
        var v35 int8 = 1
        v30 = (v33 ^ 397)
        _ = v35
    } else {
        v36 := []int{2147483647, (v30 << ((1) & 63))}
        _ = v36
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        v17 += ((v17 + ((((float32(((float64((v17 / 3.14159)) * -2.25) + 3.14159)) - (float32(v6) + 1e10)) - 3.14159) - (-2.25 + float32(((v6 / v6) + v6)))) / v17)) * v17)
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
