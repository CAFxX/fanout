package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 bool; F2 uint64 }
type GenS3 struct { F0 int64; F1 int8; F2 uint32 }
func genFn0(p0 uint64, p1 int32, p2 int32) uint32 {
    p0 += (p0 & (p0 ^ p0))
    p2 += 2147483647
    p2 += p2
    var v4 int32 = (-(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    switch -366 {
    case 128:
        var v5 uint32 = 65535
        _ = v5
    case 9223372036854775807:
        var v6 bool = false
        _ = v6
    case -462:
        if true {
            var v7 uint = 0
            _ = v7
        } else {
            var v8 uint32 = 42
            _ = v8
        }
    default:
        var v9 int16 = -1
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        switch ("a" + "hello") {
        case "a":
            fmt.Println(fi10)
        case "hello":
            v11 := GenS1{F0: (-(fi10))}
            _ = v11
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            if (18446744073709551615 > 4294967295) {
                v13 := []int{fi12, (fi12 / (((fi12 & 127)) | 1)), (^(fi12)), fi12}
                var v14 int16 = 42
                if true {
                    fmt.Println(v14)
                }
                _ = v13
                _ = v14
            } else {
                var v15 uint64 = 42
                fmt.Println(v15)
                v16 := []int{(fi12 << ((1) & 63)), -9223372036854775808}
                _ = v15
                _ = v16
            }
            _ = fi12
        }
        var v17 bool = (!true)
        _ = fi10
        _ = v17
    }
    var v18 uint32 = genFn0(2, -2147483648, -813)
    if (4294967295 != -397) {
        v18 = genFn0(783, 700, -180)
        v19 := GenS1{F0: -1}
        _ = v19
    } else {
        for fi20 := 0; fi20 < 18; fi20++ {
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 float64 = 0.0
                v23 := GenS2{F0: -103, F1: false, F2: 4294967295}
                if false {
                    v18 -= (^(v18))
                    var v24 uint = 903
                    var v25 int32 = -270
                    _ = v24
                    _ = v25
                } else {
                    fmt.Println(fi20)
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        switch 127 {
        case -1:
            var v26 string = "foo bar"
            _ = v26
        }
        v27 := GenS3{F0: 1, F1: -128, F2: v18}
        _ = v27
    }
    fmt.Println(v18)
    var v28 string = "\n\t"
    v29 := GenS2{F0: -128, F1: false, F2: 42}
    v28 = v28
    v28 = "\n\t"
    v18 = (^(v18))
    var v30 int16 = -32768
    var v31 bool = (1.5 >= (float64(1.5) + 1e10))
    if (1 < 42) {
        if (2 >= 409) {
            var v32 float64 = float64(0.0)
            v33 := []int{78, int(v30)}
            _ = v32
            _ = v33
        }
    } else {
        v34 := []int{int(v30), 947}
        switch 1 {
        case 9223372036854775807:
            fmt.Println(v31, v31, v18)
        }
        _ = v34
    }
    fmt.Println(v28, v31)
    v30 = v30
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
