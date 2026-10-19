package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint32 }
func genFn0(p0 uint64, p1 string, p2 int16) uint {
    var v4 string = "foo bar"
    fmt.Println(v4, p2, v4)
    switch 1 {
    case -2147483648:
        p0 -= (p0 * p0)
    case -9223372036854775808:
        p0 = (p0 ^ p0)
    case 0:
        var v5 int16 = (p2 << ((65535) & 63))
        _ = v5
    default:
        v6 := []int{127, -2147483648, 1, -344}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 int16, p1 uint32, p2 uint) int64 {
    for fi7 := 0; fi7 < 13; fi7++ {
        switch ("hello" + "a") {
        case "\n\t":
            p1 += (p1 / ((p1) | 1))
        case "":
            p1 += p1
        default:
            v8 := GenS1{F0: 32767, F1: 127, F2: 947}
            _ = v8
        }
        _ = fi7
    }
    if false {
        switch "a" {
        case "hello":
            if false {
                var v9 int64 = 9223372036854775807
                fmt.Println(v9)
                _ = v9
            } else {
                switch 4294967295 {
                case 2147483647:
                    v10 := []int{9223372036854775807, 0}
                    _ = v10
                default:
                    v11 := GenS1{F0: 255, F1: 122, F2: 1}
                    _ = v11
                }
                var v12 float32 = float32(((3.14159 / 1.0) - (((0.1 / (3.14159 * 1.5)) + (((0.1 * (float64((((float32(float64(float32((0.1 * -2.25)))) + 1.5) * 1.5) / 3.14159)) + 0.1)) * 0.1) / (1.5 + 1e10))) - 3.14159)))
                _ = v12
            }
        case "\n\t":
            if (!((128 > 127) || false)) {
                if ((4294967295 >= 4294967295) && false) {
                    var v13 int8 = 9
                    _ = v13
                }
                fmt.Println(p2, p1, p2)
            }
        case "foo bar":
            fmt.Println(p2)
        }
        v14 := []int{-391, 127, 32768, 2147483647}
        _ = v14
    } else {
        p1 = (p1 / ((168) | 1))
        v15 := GenS2{F0: 255}
        _ = v15
    }
    v16 := []int{32767, -507, int(p0)}
    switch ("\n\t" + "hello") {
    case "a":
        p1 += 0
    case "\n\t":
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 675
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{("foo bar" + ("hello" + ""))}
    var v3 uint32 = 816
    var v17 uint8 = 42
    fmt.Println(v3, v3)
    switch "\n\t" {
    case "a":
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 uint8 = 128
            v19 -= (v17 & v19)
            _ = fi18
            _ = v19
        }
    case "":
        v3 = (^(v3))
    default:
        v17 += v17
    }
    v20 := GenS1{F0: 65535, F1: 67, F2: 42}
    if true {
        var v21 float32 = 0.1
        _ = v21
    } else {
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 int32 = -1
            for fi24 := 0; fi24 < 13; fi24++ {
                v25 := GenS2{F0: 65535}
                _ = len(v1)
                v3 += v3
                v17 = (^(v17))
                _ = fi24
                _ = v25
            }
            var v26 bool = false
            _ = fi22
            _ = v23
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 uint32 = 65535
        _ = fi27
        _ = v28
    }
    fmt.Println(v3)
    delete(v1, "a")
    _ = len(v1)
    for fi29 := 0; fi29 < 14; fi29++ {
        v30 := GenS3{F0: (fi29 / ((-752) | 1)), F1: (v3 + (v3 - 2147483648))}
        fmt.Println(fi29, fi29)
        _ = fi29
        _ = v30
    }
    var v31 int16 = 1
    v17 = (v17 & v17)
    fmt.Println(v17, v31, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v31
}
