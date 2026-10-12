package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 bool) uint64 {
    var v6 bool = p0
    if (42 > -188) {
        p0 = (!("a" <= ""))
        var v7 int = 104
        if (!(!(false && (0 > 32767)))) {
            var v8 int64 = 9223372036854775807
            var v9 uint32 = 2147483648
            _ = v8
            _ = v9
        }
        _ = v7
    } else {
        for fi10 := 0; fi10 < 6; fi10++ {
            v6 = v6
            fmt.Println(v6, v6, fi10)
            v11 := GenS1{F0: 0}
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := []int{32767, 4294967295}
            v6 = false
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v15 := []int{}
        fmt.Println(fi14, p0, fi14)
        _ = fi14
        _ = v15
    }
    p0 = false
    if (65535 < 18446744073709551615) {
        v16 := []int{128, -328, 186}
        for fi17 := 0; fi17 < 12; fi17++ {
            switch -2147483648 {
            case 4294967295:
                v18 := GenS1{F0: 714}
                _ = v18
            }
            if len(v16) > 0 {
                v16[0] = (fi17 ^ fi17)
            }
            v6 = (!v6)
            _ = fi17
        }
        _ = v16
    }
    _ = p0
    _ = v6
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 0
    var v3 uint16 = 42
    var v4 uint = 107
    var v5 uint64 = uint64(v2)
    var v19 uint16 = v3
    v5 += v5
    for fi20 := 0; fi20 < 8; fi20++ {
        if (-2147483648 <= -910) {
            switch -2147483648 {
            case 32768:
                fmt.Println(v19, v2, fi20)
            }
        }
        for fi21 := 0; fi21 < 8; fi21++ {
            v4 += (v4 % (((v4 | v4)) | 1))
            v2 = v19
            v22 := GenS1{F0: v5}
            _ = fi21
            _ = v22
        }
        _ = len(v1)
        _ = fi20
    }
    v4 = (^(v4))
    if (!(!(108 != 180))) {
        var v23 int64 = -919
        _ = v23
    } else {
        v2 = (v3 % (((v2 ^ uint16(v4))) | 1))
        var v24 int = 42
        _ = v24
    }
    v25 := GenS2{F0: 32767}
    var v26 int = 2147483647
    v27 := []int{v26, 9223372036854775807, (v26 - v26), -9223372036854775808}
    v19 = 255
    v19 -= 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
