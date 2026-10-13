package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint8 }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint = 1
        if (!false) {
            var v9 int = (fi7 >> ((v8) & 63))
            var v10 string = ("" + "\n\t")
            _ = v10
            _ = v9
        } else {
            fmt.Println(fi7, fi7)
        }
        v11 := GenS1{F0: 100, F1: "a", F2: 2}
        _ = fi7
        _ = v11
        _ = v8
    }
    switch "foo bar" {
    case "foo bar":
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 int = fi12
            _ = fi12
            _ = v13
        }
    }
    v14 := []int{42, -1}
    var v15 bool = (!(!(42 > 0)))
    _ = v14
    _ = v15
    return 2147483647
}
func genFn1(p0 int8, p1 int64) uint {
    var v16 uint8 = 181
    var v17 float32 = 1e10
    switch 1 {
    case 890:
        v16 += 0
    default:
        if false {
            if (float32(0.0) > (v17 - v17)) {
                v17 += (v17 + (1e10 - (1.5 + ((((v17 * ((v17 * v17) + v17)) - ((1e10 + ((v17 / float32((1e10 - 3.14159))) + v17)) + 1e10)) / ((1e10 + v17) - 0.0)) / v17))))
            } else {
                p0 -= (p0 * -1)
                var v18 int64 = (p1 % ((p1) | 1))
                _ = v18
            }
        }
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 32
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 1
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    var v5 string = ("" + "hello")
    var v6 int8 = 1
    if false {
        v19 := GenS2{F0: 314, F1: 2147483647}
        if (0 == 0) {
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 int64 = 9223372036854775807
                v6 += (v6 % ((127) | 1))
                _ = fi20
                _ = v21
            }
            v6 += (-(v6))
        } else {
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 int32 = -2147483648
                var v24 float32 = (float32(1e10) * 3.14159)
                if (-2.25 != 1e10) {
                    v25 := []int{}
                    v26 := GenS2{F0: 295, F1: 43}
                    v27 := GenS1{F0: -1, F1: "hello", F2: (^(v2))}
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            if (!(1e10 != 1.5)) {
                fmt.Println(v2)
                delete(v1, 40)
                v2 = (^(v2))
            } else {
                v2 = v2
                var v28 string = (v5 + "foo bar")
                fmt.Println(v6, v2, v6)
                _ = v28
            }
        }
        fmt.Println(v6, v5, v2)
        _ = v19
    }
    var v29 int64 = -9223372036854775808
    v30 := GenS1{F0: (v6 << ((866) & 63)), F1: (("a" + "a") + "foo bar"), F2: 233}
    v29 -= v29
    v29 = int64(v6)
    var v31 string = (("hello" + ("foo bar" + (v5 + (v5 + v5)))) + "a")
    switch -1 {
    case 194:
        v32 := []int{}
        _ = v32
    case 0:
        var v33 uint = 1
        _ = v33
    case 9223372036854775807:
        fmt.Println(v2, v6)
    }
    var v34 float64 = 1.5
    fmt.Println(v29, v2, v6)
    var v35 int = 1
    v36 := GenS2{F0: v2, F1: 32767}
    if (v6 < (v6 - v6)) {
        v37 := GenS3{F0: 18446744073709551615, F1: 217}
        _ = v37
    } else {
        v38 := GenS1{F0: (v6 % (((v6 * (v6 ^ v6))) | 1)), F1: v5, F2: 271}
        v39 := GenS4{F0: 0}
        var v40 float32 = 0.1
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
