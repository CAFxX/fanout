package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint }
func genFn0(p0 int, p1 uint32) int16 {
    v5 := []int{}
    p0 += 968
    switch (^(p0)) {
    case 284:
        fmt.Println(p0)
    case -504:
        v6 := GenS1{F0: 404, F1: 2, F2: 32767}
        _ = v6
    case 1:
        var v7 string = "\n\t"
        _ = v7
    default:
        p0 += (p0 * -793)
    }
    _ = p0
    _ = p1
    _ = v5
    return -32768
}
func genFn1(p0 int32, p1 uint) uint64 {
    p1 = p1
    var v8 uint32 = 2147483647
    if (true && (9223372036854775807 != 32767)) {
        fmt.Println(v8, p1)
    }
    var v9 int = -9223372036854775808
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[int]string)
    var v3 uint64 = 0
    var v4 int32 = 42
    var v10 float64 = 1.5
    var v11 uint = 642
    v1 += 1.5
    var v12 uint32 = 2147483648
    v13 := []int{32767}
    v3 = v3
    v2[88] = "hello"
    for fi14 := 0; fi14 < 10; fi14++ {
        delete(v2, (fi14 * (fi14 - 399)))
        _ = fi14
    }
    switch 42 {
    case 4294967295:
        fmt.Println(v4)
    case -3:
        switch 1 {
        case 950:
            switch "hello" {
            case "":
                var v15 uint64 = 2
                _ = v15
            default:
                v10 = v10
            }
        case -225:
            fmt.Println(v4, v4)
        default:
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := GenS2{F0: false, F1: 255}
                _ = fi16
                _ = v17
            }
        }
    case -9223372036854775808:
        if false {
            v18 := GenS1{F0: 255, F1: 255, F2: 32768}
            _ = v18
        } else {
            switch 111 {
            case -1:
                fmt.Println(v10)
            default:
                v2[-2147483648] = "foo bar"
            }
            var v19 int = int(v12)
            var v20 uint64 = v3
            _ = v19
            _ = v20
        }
    default:
        fmt.Println(v10)
    }
    _ = len(v13)
    if (!(!false)) {
        if ((v3 % ((genFn1(v4, 18446744073709551615)) | 1)) > genFn1(1, (^(v11)))) {
            var v21 int = -306
            var v22 int32 = v4
            _ = v21
            _ = v22
        }
        var v23 int = -1
        for fi24 := 0; fi24 < 13; fi24++ {
            if true {
                if ((v3 & v3) >= 414) {
                    v10 = 1.5
                }
                v25 := GenS1{F0: 1, F1: 236, F2: (fi24 - (v23 / ((v23) | 1)))}
                var v26 uint32 = uint32(v3)
                _ = v25
                _ = v26
            } else {
                v10 -= (3.14159 / float64(v1))
                var v27 bool = (!false)
                v28 := GenS3{F0: 0, F1: 0}
                _ = v27
                _ = v28
            }
            v29 := GenS4{F0: 214}
            _ = fi24
            _ = v29
        }
        _ = v23
    }
    switch -605 {
    case 32768:
        fmt.Println(v4, v12)
    case 766:
        var v30 bool = (-32768 != 0)
        _ = v30
    case 9223372036854775807:
        v12 = (v12 | (v12 % ((v12) | 1)))
    }
    if false {
        var v31 bool = true
        var v32 int32 = -2147483648
        v33 := GenS4{F0: uint(v12)}
        _ = v31
        _ = v32
        _ = v33
    }
    v12 += 980
    for fi34 := 0; fi34 < 18; fi34++ {
        switch fi34 {
        case 9223372036854775807:
            v4 = (v4 | (^(v4)))
        }
        v2[fi34] = "a"
        v11 -= (v11 * v11)
        _ = fi34
    }
    v10 += (v10 - v10)
    delete(v2, 32768)
    v35 := GenS1{F0: 32767, F1: 255, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
