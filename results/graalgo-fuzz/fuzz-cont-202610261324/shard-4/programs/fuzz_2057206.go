package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
func genFn0(p0 uint16, p1 int16) int32 {
    if ((!false) || true) {
        if false {
            if true {
                fmt.Println(p0, p1)
                p1 += (p1 >> ((18446744073709551615) & 63))
            } else {
                var v5 uint64 = 102
                switch -1 {
                case 32767:
                    v6 := []int{32768, -2147483648}
                    _ = v6
                default:
                    if (1 > 4294967295) {
                        fmt.Println(p1, p1)
                        p0 = 866
                    }
                }
                _ = v5
            }
        }
        p1 -= 0
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint16 = (p0 - (p0 & p0))
        if true {
            v9 := GenS1{F0: 1, F1: 1}
            var v10 int16 = (p1 + p1)
            _ = v10
            _ = v9
        } else {
            v11 := GenS1{F0: 255, F1: 4294967295}
            v8 += v8
            p0 -= (v8 % (((^(v8))) | 1))
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    var v12 float64 = float64(0.1)
    for fi13 := 0; fi13 < 4; fi13++ {
        v12 = (v12 - v12)
        fmt.Println(p0, v12)
        v14 := []int{(fi13 + (fi13 / (((fi13 >> ((919) & 63))) | 1))), 401, (fi13 - (fi13 - fi13)), fi13}
        _ = fi13
        _ = v14
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v12
    return 0
}

func main() {
    var v1 int8 = 61
    var v2 bool = (false && (v1 > (v1 | v1)))
    v3 := make(map[int]bool)
    var v4 bool = ((1.5 != (1.5 + float32(3.14159))) || true)
    fmt.Println(v1, v2)
    v4 = (false || (1 < 0))
    var v15 uint64 = 4294967295
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 uint64 = v15
        var v18 int16 = 0
        v19 := []int{}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    switch "foo bar" {
    case "hello":
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 int = 32767
            _ = fi20
            _ = v21
        }
    case "a":
        if (!true) {
            if false {
                var v22 int64 = -9223372036854775808
                var v23 int32 = -2147483648
                for fi24 := 0; fi24 < 16; fi24++ {
                    var v25 string = "hello"
                    switch fi24 {
                    case 32767:
                        v26 := GenS1{F0: 193, F1: 2}
                        _ = v26
                    }
                    var v27 uint8 = 127
                    _ = fi24
                    _ = v25
                    _ = v27
                }
                _ = v22
                _ = v23
            } else {
                v2 = (!true)
                var v28 uint8 = 127
                _ = v28
            }
        }
    case "\n\t":
        v29 := []int{42}
        _ = v29
    }
    switch "hello" {
    case "":
        if false {
            v30 := GenS1{F0: 683, F1: 65535}
            v31 := GenS1{F0: 32767, F1: 255}
            _ = v30
            _ = v31
        }
    }
    var v32 uint8 = 77
    v3[4294967295] = true
    v33 := GenS1{F0: 32767, F1: 4294967295}
    v34 := GenS1{F0: 65535, F1: 465}
    v35 := GenS1{F0: 65535, F1: 65535}
    var v36 int32 = -2147483648
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
