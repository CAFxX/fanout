package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int8 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int64, p1 int8) string {
    p1 = (p1 ^ p1)
    var v6 int8 = p1
    var v7 float64 = (3.14159 + 3.14159)
    v8 := []int{32767, 32767}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return ""
}
func genFn1(p0 int32) int {
    p0 -= p0
    v9 := GenS1{F0: 774, F1: -573, F2: 51}
    _ = p0
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 int32 = 42
    var v2 int = 42
    var v3 uint16 = 0
    var v4 int64 = 1
    v5 := make(map[uint32]string)
    v10 := GenS2{F0: 0, F1: false}
    for fi11 := 0; fi11 < 9; fi11++ {
        if true {
            v3 += v3
            v12 := []int{(fi11 | (v2 >> ((114) & 63))), (fi11 ^ (v2 * 42)), (fi11 << ((65535) & 63)), v2}
            _ = v12
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v2 = (v2 + v2)
        _ = fi13
    }
    var v14 int16 = 439
    v3 += v3
    var v15 uint32 = 4294967295
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 2; fi17++ {
            switch ("hello" + "\n\t") {
            case "":
                v18 := GenS3{F0: 1}
                _ = v18
            case "\n\t":
                var v19 uint16 = v3
                _ = v19
            case "foo bar":
                v20 := GenS1{F0: uint64(v3), F1: -1, F2: -95}
                _ = v20
            default:
                var v21 int32 = -2147483648
                v21 += v21
                if true {
                    v22 := GenS3{F0: 736}
                    v21 = v21
                    v15 -= 59
                    _ = v22
                }
                _ = v21
            }
            _ = fi17
        }
        fmt.Println(v2, v3)
        v23 := GenS4{F0: 2147483647, F1: 456}
        _ = fi16
        _ = v23
    }
    var v24 uint8 = 2
    fmt.Println(v2, v15)
    v24 -= v24
    for fi25 := 0; fi25 < 19; fi25++ {
        fmt.Println(v14)
        var v26 uint32 = uint32(v24)
        fmt.Println(v26, v4)
        _ = fi25
        _ = v26
    }
    var v27 bool = (!((v1 < (^(v1))) || false))
    for fi28 := 0; fi28 < 18; fi28++ {
        var v29 float64 = 0.0
        v30 := GenS5{F0: v15}
        fmt.Println(v4)
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 int64 = (^(v4))
    switch (^(v2)) {
    case -1:
        v15 = v15
    default:
        for fi32 := 0; fi32 < 2; fi32++ {
            switch -9223372036854775808 {
            case -9223372036854775808:
                for fi33 := 0; fi33 < 17; fi33++ {
                    delete(v5, v15)
                    fmt.Println(v27)
                    _ = fi33
                }
            case 127:
                v34 := GenS5{F0: (v15 + v15)}
                _ = v34
            case -860:
                fmt.Println(fi32, v15)
            }
            v4 -= v31
            _ = fi32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
