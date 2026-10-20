package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int32; F2 uint32 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 string; F1 uint64 }
func genFn0(p0 uint16, p1 bool, p2 int32) uint64 {
    var v4 uint32 = 2147483648
    if p1 {
        var v5 float32 = 3.14159
        _ = v5
    } else {
        v6 := []int{1, 2147483647}
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        p0 = p0
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}
func genFn1(p0 int, p1 bool, p2 int16) int32 {
    if (407 >= 358) {
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS1{F0: 328}
            _ = fi8
            _ = v9
        }
    } else {
        p2 = (p2 & (p2 + p2))
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 int32 = -2147483648
        if (false || (4294967295 >= 42)) {
            var v12 uint16 = 255
            p2 = (p2 | -32768)
            v13 := GenS2{F0: "\n\t", F1: (v11 | v11), F2: 872}
            _ = v12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    var v1 int = 276
    var v2 int32 = -551
    var v3 float64 = float64((1.5 / 1.0))
    var v14 uint = 524
    if (((v14 | 42) > (v14 / ((v14) | 1))) && (false || (!true))) {
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := []int{(v1 + int(v2)), 42, (fi15 + -9223372036854775808)}
            var v17 uint8 = 0
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 uint = v14
        v19 := GenS3{F0: "a", F1: 678}
        _ = v18
        _ = v19
    }
    var v20 string = "a"
    var v21 int8 = 35
    fmt.Println(v21, v3)
    var v22 int64 = 407
    v23 := []int{(v1 & (v1 - v1)), (v1 - v1), v1, 751}
    switch "a" {
    case "hello":
        for fi24 := 0; fi24 < 20; fi24++ {
            v21 -= v21
            if false {
                fmt.Println(fi24, v2)
                v25 := GenS4{F0: (("hello" + "hello") + "foo bar"), F1: 57}
                _ = v25
            } else {
                if len(v23) > 0 {
                    v23[0] = (-(fi24))
                }
            }
            _ = fi24
        }
    case "":
        switch v20 {
        case "":
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 int32 = (v2 & v2)
                for fi28 := 0; fi28 < 1; fi28++ {
                    var v29 bool = false
                    v27 += v2
                    var v30 uint = v14
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = fi26
                _ = v27
            }
        default:
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
        }
    case "foo bar":
        fmt.Println(v22, v20, v1)
    default:
        var v31 uint8 = 179
        _ = v31
    }
    var v32 int = (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v32
}
