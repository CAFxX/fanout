package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int; F1 int64; F2 int64 }
func genFn0(p0 uint64, p1 uint8) uint64 {
    var v6 bool = true
    switch 0 {
    case -643:
        var v7 uint64 = p0
        _ = v7
    case 536:
        var v8 uint = 65535
        _ = v8
    case 568:
        v9 := GenS1{F0: -9223372036854775808, F1: 342, F2: 0}
        _ = v9
    }
    v6 = (v6 && (!true))
    for fi10 := 0; fi10 < 13; fi10++ {
        v6 = v6
        _ = fi10
    }
    v6 = ((!v6) && (!true))
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    var v1 int8 = 42
    var v2 uint8 = 1
    var v3 uint8 = v2
    var v4 string = "a"
    v5 := make(map[int]bool)
    v11 := GenS1{F0: -452, F1: 42, F2: 32767}
    if true {
        v12 := GenS1{F0: -2147483648, F1: 1, F2: 0}
        var v13 uint16 = 1
        v14 := GenS1{F0: 90, F1: 18446744073709551615, F2: int16(v1)}
        _ = v12
        _ = v13
        _ = v14
    } else {
        v2 = v2
        v4 = (v4 + v4)
    }
    switch "\n\t" {
    case "hello":
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS2{F0: 2147483647}
            fmt.Println(v4)
            _ = fi15
            _ = v16
        }
    case "\n\t":
        fmt.Println(v1)
    default:
        var v17 float64 = (-2.25 / 1.0)
        _ = v17
    }
    switch (v4 + "foo bar") {
    case "":
        var v18 uint = 1
        _ = v18
    }
    switch 163 {
    case 32767:
        var v19 uint32 = 4294967295
        _ = v19
    }
    fmt.Println(v2)
    v1 = v1
    v4 = (v4 + "hello")
    if ((((-2.25 * 0.1) - 1.5) - float64((3.14159 * 1e10))) <= (1e10 / (1.5 / 0.1))) {
        v3 += (^(v2))
        v20 := GenS1{F0: 32767, F1: 2, F2: 1000}
        for fi21 := 0; fi21 < 7; fi21++ {
            v22 := []int{fi21, (fi21 + (fi21 & -496)), (fi21 << ((1) & 63))}
            fmt.Println(v2, v4, v2)
            if true {
                fmt.Println(v3)
            }
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        for fi23 := 0; fi23 < 15; fi23++ {
            for fi24 := 0; fi24 < 5; fi24++ {
                switch (-(fi24)) {
                case -9223372036854775808:
                    if false {
                        v25 := GenS3{F0: 4294967295, F1: 4294967295}
                        v1 -= v1
                        _ = v25
                    }
                case 9223372036854775807:
                    switch -9223372036854775808 {
                    case -714:
                        v26 := GenS4{F0: (fi23 ^ (fi23 >> ((4294967295) & 63))), F1: 0, F2: 0}
                        _ = v26
                    case -290:
                        var v27 string = ("" + "a")
                        _ = v27
                    default:
                        if true {
                            v2 = v3
                            v1 = (v1 >> ((42) & 63))
                            v28 := GenS2{F0: 65535}
                            _ = v28
                        }
                    }
                }
                fmt.Println(v2, v1)
                _ = fi24
            }
            _ = fi23
        }
        v2 = 127
        for fi29 := 0; fi29 < 4; fi29++ {
            v30 := GenS2{F0: 228}
            _ = fi29
            _ = v30
        }
    }
    v1 = v1
    v2 -= (v3 % ((42) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
