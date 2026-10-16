package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int64; F2 string }
func genFn0(p0 int, p1 uint32) int32 {
    switch p0 {
    case 42:
        var v5 float32 = (((float32(float64(1e10)) / 1e10) + (-2.25 / 3.14159)) / 1.0)
        _ = v5
    }
    v6 := GenS1{F0: 4294967295, F1: 65535}
    p1 = p1
    _ = p0
    _ = p1
    _ = v6
    return -451
}

func main() {
    var v1 int16 = 32767
    var v2 int8 = 100
    var v3 uint16 = 42
    v4 := make(map[string]int)
    v7 := []int{127, 127, 2147483647, 1}
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: 2147483648, F1: 42}
            if (!((v1 - -419) != v1)) {
                if false {
                    var v11 bool = false
                    _ = v11
                } else {
                    switch "hello" {
                    case "\n\t":
                        if true {
                            if false {
                                var v12 uint = 293
                                v2 += (v2 + -1)
                                _ = v12
                            }
                        }
                    }
                }
                v4[(("foo bar" + "foo bar") + ("hello" + "\n\t"))] = (fi8 + fi9)
            }
            var v13 int8 = (v2 << ((4294967295) & 63))
            _ = fi9
            _ = v10
            _ = v13
        }
        fmt.Println(v2, v1)
        v14 := GenS2{F0: 844, F1: int64(v1), F2: "\n\t"}
        _ = fi8
        _ = v14
    }
    v15 := v4["\n\t"]
    v1 += v1
    v16 := GenS2{F0: 1, F1: 42, F2: "a"}
    switch v15 {
    case 2147483647:
        v17 := GenS1{F0: 0, F1: (v3 / (((v3 << ((uint(v3)) & 63))) | 1))}
        _ = v17
    }
    switch "\n\t" {
    case "a":
        for fi18 := 0; fi18 < 1; fi18++ {
            if (2147483648 == 925) {
                var v19 bool = false
                var v20 uint16 = (v3 << ((1) & 63))
                fmt.Println(v1)
                _ = v19
                _ = v20
            }
            switch "a" {
            case "a":
                fmt.Println(fi18, v2, v2)
            case "hello":
                var v21 float64 = 3.14159
                _ = v21
            default:
                fmt.Println(v15, v2, fi18)
            }
            var v22 uint8 = 2
            _ = fi18
            _ = v22
        }
    }
    if false {
        if true {
            for fi23 := 0; fi23 < 4; fi23++ {
                v3 -= v3
                _ = fi23
            }
        }
        if (94 != 0) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v24 int32 = 1
            _ = v24
        }
    }
    switch "hello" {
    case "":
        for fi25 := 0; fi25 < 7; fi25++ {
            delete(v4, "\n\t")
            switch ("hello" + ("a" + "a")) {
            case "hello":
                var v26 float64 = (float64(1e10) + 0.1)
                _ = v26
            case "foo bar":
                for fi27 := 0; fi27 < 10; fi27++ {
                    v28 := GenS1{F0: 984, F1: (v3 << ((454) & 63))}
                    _ = fi27
                    _ = v28
                }
            }
            _ = fi25
        }
    case "\n\t":
        v29 := v4["a"]
        _ = v29
    default:
        v2 += -92
    }
    v30 := GenS1{F0: 42, F1: (v3 * (v3 ^ 32767))}
    v31 := v4["\n\t"]
    v15 = v31
    v31 = (v15 ^ v31)
    v1 = int16(v15)
    v1 = (v1 >> ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v7
}
