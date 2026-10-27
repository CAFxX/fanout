package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 int }
func genFn0(p0 uint8, p1 uint32) int {
    switch -374 {
    case 9223372036854775807:
        v4 := []int{388}
        _ = v4
    case 2147483647:
        v5 := GenS1{F0: 0, F1: 127, F2: -128}
        _ = v5
    default:
        v6 := GenS1{F0: p1, F1: int(p1), F2: 124}
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        if (2147483647 != -849) {
            fmt.Println(fi7, p1, p0)
            var v8 uint = 892
            _ = v8
        } else {
            v9 := GenS1{F0: (p1 + (p1 & (p1 * 0))), F1: 2147483647, F2: -59}
            if false {
                var v10 uint = 0
                _ = v10
            } else {
                var v11 uint = 42
                _ = v11
            }
            _ = v9
        }
        p0 = 128
        _ = fi7
    }
    var v12 uint = 65535
    v13 := []int{1, 4294967295}
    v13 = append(v13, -1)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    v1 := []int{}
    var v2 int32 = -2147483648
    v3 := []int{}
    fmt.Println(v2)
    switch -2147483648 {
    case 32767:
        switch -130 {
        case -9223372036854775808:
            v2 = v2
        }
    case -2147483648:
        var v14 int64 = -255
        _ = v14
    case 0:
        v2 -= (v2 | v2)
    default:
        v2 += v2
    }
    switch -68 {
    case 42:
        if ((float32(-2.25) != -2.25) || ("hello" == ("a" + "foo bar"))) {
            if len(v1) > 0 {
                v1[0] = 128
            }
        }
    case 32767:
        for fi15 := 0; fi15 < 14; fi15++ {
            if false {
                var v16 int32 = (v2 >> ((4294967295) & 63))
                var v17 uint64 = 860
                _ = v16
                _ = v17
            } else {
                var v18 int64 = 24
                v19 := GenS1{F0: 1, F1: (fi15 | fi15), F2: -128}
                _ = v18
                _ = v19
            }
            var v20 float64 = -2.25
            _ = fi15
            _ = v20
        }
    }
    fmt.Println(v2)
    if (("foo bar" < "a") && (32767 >= 925)) {
        for fi21 := 0; fi21 < 18; fi21++ {
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 uint16 = 0
                v23 += 0
                v24 := GenS1{F0: 880, F1: -2147483648, F2: 100}
                _ = fi22
                _ = v23
                _ = v24
            }
            var v25 bool = (fi21 == (fi21 + genFn0(193, 2147483647)))
            v2 = 1
            _ = fi21
            _ = v25
        }
        if (42 < -1) {
            v2 += (v2 & (v2 * (v2 | v2)))
        }
    } else {
        if len(v3) > 0 {
            v3[0] = -545
        }
        v26 := GenS1{F0: 1, F1: -2147483648, F2: 1}
        _ = v26
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v27 string = ("" + ("a" + "\n\t"))
        for fi28 := 0; fi28 < 16; fi28++ {
            fmt.Println(v2, v27)
            _ = fi28
        }
        _ = v27
    } else {
        fmt.Println(v2)
        v29 := GenS1{F0: 967, F1: 32768, F2: 1}
        v2 += (v2 % ((-1) | 1))
        _ = v29
    }
    v30 := GenS1{F0: 65535, F1: -601, F2: 1}
    var v31 uint32 = 2
    switch 35 {
    case 32767:
        v32 := GenS2{F0: -128, F1: 24, F2: 4294967295}
        _ = v32
    }
    v31 += (v31 & (v31 << ((uint(v31)) & 63)))
    fmt.Println(v31)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
