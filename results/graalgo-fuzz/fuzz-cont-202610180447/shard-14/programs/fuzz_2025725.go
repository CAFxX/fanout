package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int64 }
func genFn0() string {
    v7 := []int{4294967295, -9223372036854775808, -1}
    var v8 uint16 = 255
    switch -725 {
    case 0:
        for fi9 := 0; fi9 < 20; fi9++ {
            fmt.Println(v8, v8)
            var v10 int64 = 0
            _ = fi9
            _ = v10
        }
    default:
        if true {
            switch -310 {
            case 0:
                var v11 int8 = 0
                _ = v11
            }
        }
    }
    v8 = (v8 >> ((0) & 63))
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1(p0 string) int16 {
    var v12 uint64 = 4294967295
    for fi13 := 0; fi13 < 20; fi13++ {
        if false {
            v14 := []int{-2147483648, (-(fi13))}
            _ = v14
        }
        _ = fi13
    }
    fmt.Println(v12, v12)
    var v15 uint64 = 18446744073709551615
    if false {
        v15 -= 0
    }
    _ = p0
    _ = v12
    _ = v15
    return -32768
}

func main() {
    var v1 uint32 = 92
    v2 := make(map[string]string)
    var v3 uint16 = 2
    var v4 int64 = -304
    var v5 int64 = 42
    v6 := []uint32{v1, 998, 42}
    for fi16 := 0; fi16 < 16; fi16++ {
        if (2147483647 != 2147483647) {
            fmt.Println(v1)
        } else {
            var v17 uint8 = 0
            fmt.Println(v3, v1)
            fmt.Println(v5, v4, fi16)
            _ = v17
        }
        _ = fi16
    }
    _ = len(v2)
    var v18 uint8 = 192
    v19 := GenS1{F0: v3, F1: (v18 / (((v18 / (((v18 * (^(v18)))) | 1))) | 1)), F2: int16(v18)}
    if (1 >= int32(v18)) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        v20 := GenS2{F0: 0}
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        v5 += (v4 + v4)
        v22 := GenS3{F0: 42, F1: 1}
        _ = fi21
        _ = v22
    }
    v23 := GenS2{F0: (v3 % ((32767) | 1))}
    if (0 == 42) {
        var v24 int32 = 634
        if (true || false) {
            var v25 uint16 = (v3 - v3)
            _ = v25
        }
        _ = v24
    }
    var v26 int8 = 42
    var v27 uint = 1
    v18 += uint8(v1)
    fmt.Println(v3, v27)
    v27 += (^(v27))
    if (false || ((^(v27)) <= v27)) {
        v26 = (v26 & 42)
    } else {
        switch "a" {
        case "":
            if true {
                var v28 int16 = -828
                _ = v28
            }
        default:
            v27 -= v27
        }
        for fi29 := 0; fi29 < 2; fi29++ {
            v27 += v27
            if true {
                v3 -= (v3 | v3)
            }
            _ = fi29
        }
        if false {
            v27 -= v27
            var v30 bool = false
            _ = v30
        } else {
            v31 := v2["hello"]
            var v32 int32 = 1
            _ = v31
            _ = v32
        }
    }
    if (true && (true && ((!(954 <= 1)) && (true || (4294967295 >= (v1 & (^(v1)))))))) {
        v33 := GenS1{F0: (v3 + v3), F1: (v18 | (v18 % ((66) | 1))), F2: 1000}
        _ = v33
    }
    v34 := GenS4{F0: -1}
    v18 = 255
    v2["hello"] = ("a" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
