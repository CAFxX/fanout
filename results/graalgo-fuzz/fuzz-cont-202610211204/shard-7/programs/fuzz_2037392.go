package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint64 }
func genFn0() bool {
    for fi5 := 0; fi5 < 17; fi5++ {
        switch fi5 {
        case -1:
            v6 := GenS1{F0: 2147483648, F1: 633, F2: 18446744073709551615}
            _ = v6
        case -316:
            var v7 int16 = -32768
            _ = v7
        default:
            var v8 uint = 6
            _ = v8
        }
        _ = fi5
    }
    v9 := []int{2147483647, -2147483648, -9223372036854775808, 32767}
    _ = v9
    return false
}

func main() {
    var v1 int = -2147483648
    var v2 uint = 65535
    var v3 int8 = 100
    var v4 int8 = 0
    v1 = (v1 & 298)
    var v10 int64 = -1
    v1 = (v1 | 52)
    fmt.Println(v1)
    v2 += v2
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := []int{v1, fi11}
        v10 = int64(v4)
        _ = fi11
        _ = v12
    }
    v2 += v2
    var v13 int32 = 864
    v14 := GenS1{F0: 2147483648, F1: 62, F2: 1}
    if (v10 <= v10) {
        var v15 int16 = 32767
        if false {
            var v16 int8 = 1
            v17 := []int{(v1 / ((448) | 1)), (v1 >> (((v2 & v2)) & 63)), v1}
            _ = v16
            _ = v17
        }
        switch "a" {
        case "hello":
            if ((v10 + int64(v2)) < v10) {
                v15 -= v15
            }
        }
        _ = v15
    } else {
        v4 = v3
        v18 := []int{v1}
        _ = v18
    }
    v19 := GenS1{F0: 42, F1: 18446744073709551615, F2: 4294967295}
    v13 = 34
    if true {
        if (v3 != -128) {
            for fi20 := 0; fi20 < 4; fi20++ {
                for fi21 := 0; fi21 < 18; fi21++ {
                    v22 := []int{}
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    _ = fi21
                    _ = v22
                }
                if false {
                    v13 = 2147483647
                    var v23 int32 = int32(v3)
                    v13 = (v23 * (v13 - (v23 >> ((v2) & 63))))
                    _ = v23
                } else {
                    v24 := []int{(-(fi20)), (fi20 ^ v1)}
                    _ = v24
                }
                fmt.Println(fi20, v1, v1)
                _ = fi20
            }
            v4 = 87
            v4 = v4
        }
    } else {
        var v25 float64 = (((1e10 - 1.5) * ((0.0 / -2.25) + (-2.25 + 1e10))) - 0.0)
        for fi26 := 0; fi26 < 8; fi26++ {
            v27 := []int{314, (fi26 | v1)}
            v13 -= 0
            switch "\n\t" {
            case "hello":
                for fi28 := 0; fi28 < 19; fi28++ {
                    var v29 int16 = -32768
                    _ = fi28
                    _ = v29
                }
            case "":
                v4 = v3
            }
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    var v30 uint8 = uint8(v2)
    v13 = 816
    var v31 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
