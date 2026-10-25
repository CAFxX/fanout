package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int64, p1 bool, p2 uint32) int32 {
    fmt.Println(p2, p1)
    switch 32767 {
    case 128:
        for fi4 := 0; fi4 < 9; fi4++ {
            var v5 float32 = (0.0 + (3.14159 + 0.0))
            p1 = (!p1)
            _ = fi4
            _ = v5
        }
    case 127:
        if ((p0 & p0) < -9223372036854775808) {
            switch "\n\t" {
            case "foo bar":
                v6 := GenS1{F0: 271, F1: 0}
                _ = v6
            case "\n\t":
                for fi7 := 0; fi7 < 4; fi7++ {
                    var v8 uint32 = (p2 ^ p2)
                    var v9 float64 = (3.14159 * 0.0)
                    var v10 uint8 = 2
                    _ = fi7
                    _ = v10
                    _ = v8
                    _ = v9
                }
            case "a":
                var v11 uint = 0
                _ = v11
            }
        }
    case 891:
        if p1 {
            p1 = p1
            p1 = false
            var v12 uint = 536
            _ = v12
        } else {
            p1 = false
            p0 = p0
            var v13 string = "foo bar"
            _ = v13
        }
    default:
        v14 := []int{9223372036854775807, 42, -1, -563}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return -584
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    var v3 int32 = -1
    v15 := GenS1{F0: 42, F1: 156}
    if (v3 >= v3) {
        v16 := []int{9223372036854775807}
        _ = v16
    }
    if true {
        switch -606 {
        case 684:
            v3 = 2147483647
        case 897:
            fmt.Println(v3)
        case 128:
            v3 = v3
        default:
            var v17 int8 = -83
            _ = v17
        }
        v3 = genFn0(-9223372036854775808, (9223372036854775807 < 1), 65535)
    }
    switch "" {
    case "":
        var v18 int = 585
        _ = v18
    }
    var v19 uint64 = 1
    v19 = v19
    switch (("" + (("hello" + "foo bar") + "\n\t")) + "a") {
    case "":
        if (1.5 != 3.14159) {
            v2[42] = ""
            for fi20 := 0; fi20 < 20; fi20++ {
                v19 += v19
                _ = fi20
            }
        } else {
            switch 0 {
            case 865:
                for fi21 := 0; fi21 < 16; fi21++ {
                    if ((1.5 + (float32(((1e10 / -2.25) * 1e10)) + (1.5 - (0.0 * 3.14159)))) == 1.5) {
                        if true {
                            delete(v1, -2147483648)
                            v22 := GenS2{F0: (((!(!false)) && (!true)) || false), F1: -32768}
                            var v23 float64 = 3.14159
                            _ = v22
                            _ = v23
                        }
                        _ = len(v1)
                    } else {
                        v2[fi21] = ""
                        fmt.Println(v3, v19)
                        fmt.Println(v3, fi21)
                    }
                    var v24 uint = 1
                    _ = fi21
                    _ = v24
                }
            case -778:
                var v25 uint32 = 4294967295
                _ = v25
            case -2147483648:
                delete(v1, 9223372036854775807)
            }
        }
    case "a":
        switch "\n\t" {
        case "hello":
            var v26 uint16 = 238
            _ = v26
        case "\n\t":
            if false {
                v27 := GenS1{F0: 554, F1: 0}
                v19 = v19
                _ = v27
            }
        default:
            v28 := GenS3{F0: -1}
            _ = v28
        }
    case "hello":
        v19 -= 472
    default:
        v29 := []int{}
        _ = v29
    }
    v3 += 42
    switch -791 {
    case 311:
        var v30 int8 = -116
        _ = v30
    default:
        v31 := []int{42}
        _ = v31
    }
    var v32 uint8 = 127
    v19 -= uint64(v32)
    switch "a" {
    case "foo bar":
        if true {
            fmt.Println(v32, v3)
        }
    case "":
        var v33 int64 = 1
        _ = v33
    default:
        v34 := []int{32768, 4294967295, 2147483647, 261}
        _ = v34
    }
    var v35 string = ("foo bar" + "a")
    fmt.Println(v3, v32)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v35
}
