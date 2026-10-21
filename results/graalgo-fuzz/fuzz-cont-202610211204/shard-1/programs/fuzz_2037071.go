package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int32 }
func genFn0(p0 string, p1 int8, p2 uint) string {
    v5 := []int{440, 42, -912}
    var v6 string = p0
    var v7 int16 = 32767
    fmt.Println(v6, p2)
    switch "" {
    case "\n\t":
        v8 := GenS1{F0: 18446744073709551615, F1: (p1 + p1), F2: true}
        _ = v8
    default:
        v5 = append(v5, int(p1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return ""
}
func genFn1(p0 uint32, p1 int64, p2 bool) int {
    if false {
        var v9 int16 = -32768
        var v10 string = ("foo bar" + "")
        _ = v10
        _ = v9
    }
    v11 := []int{271, 0}
    switch 32767 {
    case -30:
        var v12 float64 = 1.5
        _ = v12
    case 32767:
        v13 := GenS1{F0: 4294967295, F1: 89, F2: (("a" + "foo bar") <= "foo bar")}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = -32768
    var v3 int8 = 92
    var v4 int16 = v2
    v14 := GenS2{F0: 1}
    v4 += (v4 + (v4 << ((26) & 63)))
    switch -16 {
    case -9223372036854775808:
        var v15 float64 = (3.14159 / 1e10)
        _ = v15
    }
    var v16 int = -441
    for fi17 := 0; fi17 < 11; fi17++ {
        v4 = v4
        _ = fi17
    }
    fmt.Println(v4, v2)
    var v18 uint32 = 31
    switch 657 {
    case 2147483647:
        v19 := []int{(v16 & genFn1(1, 232, false)), (v16 | (v16 % (((v16 | 32768)) | 1))), v16}
        _ = v19
    case 127:
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := GenS3{F0: "\n\t"}
            _ = fi20
            _ = v21
        }
    case 0:
        if (!((42 <= 0) || true)) {
            v22 := GenS2{F0: 1}
            v4 -= 130
            _ = v22
        } else {
            fmt.Println(v3)
        }
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS4{F0: 252, F1: int32(v18)}
        fmt.Println(v3, v18)
        if true {
            for fi25 := 0; fi25 < 3; fi25++ {
                fmt.Println(v3, fi25, v4)
                v2 = (^(v2))
                v26 := GenS2{F0: v18}
                _ = fi25
                _ = v26
            }
        }
        _ = fi23
        _ = v24
    }
    v4 = v2
    switch v16 {
    case -2147483648:
        if false {
            var v27 int8 = int8(v4)
            for fi28 := 0; fi28 < 20; fi28++ {
                if (448 == 255) {
                    v29 := GenS3{F0: ""}
                    fmt.Println(v27, v4, v4)
                    _ = v29
                }
                _ = fi28
            }
            fmt.Println(v27, v4)
            _ = v27
        }
    default:
        var v30 int64 = 42
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
