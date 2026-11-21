package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 bool) int8 {
    switch "hello" {
    case "a":
        for fi4 := 0; fi4 < 19; fi4++ {
            switch "hello" {
            case "hello":
                p0 = p0
            default:
                if p0 {
                    var v5 int16 = -57
                    _ = v5
                } else {
                    fmt.Println(p0)
                }
            }
            p0 = p0
            _ = fi4
        }
    case "":
        v6 := GenS1{F0: 0}
        _ = v6
    default:
        fmt.Println(p0)
    }
    if (2147483647 >= 9223372036854775807) {
        var v7 int64 = 427
        _ = v7
    }
    _ = p0
    return -128
}
func genFn1(p0 uint64, p1 string, p2 uint) uint16 {
    if (!(4294967295 >= 774)) {
        var v8 uint8 = 255
        _ = v8
    } else {
        fmt.Println(p1)
    }
    v9 := []int{32768}
    fmt.Println(p1)
    v10 := GenS2{F0: 12}
    switch -1 {
    case 1:
        var v11 int8 = -1
        _ = v11
    case 0:
        fmt.Println(p2, p1, p0)
    default:
        if ((3.14159 != 3.14159) && false) {
            p0 += 786
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 573
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int16 = 32767
    var v3 float32 = 3.14159
    switch 9223372036854775807 {
    case 4294967295:
        if (v3 != v3) {
            v3 += v3
        } else {
            var v12 bool = (0 <= 357)
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 uint8 = 127
                _ = fi13
                _ = v14
            }
            if (v12 && v12) {
                v3 += (3.14159 - (v3 * (v3 * (float32(((1e10 / -2.25) * 0.0)) + 1e10))))
                v15 := []int{-1, -1}
                for fi16 := 0; fi16 < 13; fi16++ {
                    v12 = v12
                    if len(v15) > 0 {
                        v15[0] = fi16
                    }
                    _ = fi16
                }
                _ = v15
            }
            _ = v12
        }
    default:
        fmt.Println(v2)
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        if true {
            fmt.Println(v3, fi17)
        }
        _ = fi17
    }
    switch 2147483647 {
    case 128:
        v18 := []int{-2147483648, 2147483647, -212}
        _ = v18
    case 168:
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 20; fi20++ {
                v21 := []int{(fi19 << ((2) & 63)), fi20, fi19, fi20}
                var v22 int64 = 1
                fmt.Println(v3, v3)
                _ = fi20
                _ = v21
                _ = v22
            }
            v23 := GenS1{F0: 0}
            _ = fi19
            _ = v23
        }
    case -9:
        var v24 int = -188
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        for fi26 := 0; fi26 < 19; fi26++ {
            v3 -= v3
            v1 = (^(v1))
            v2 = (v2 - (-(v2)))
            _ = fi26
        }
        var v27 int8 = 63
        v27 = v27
        _ = fi25
        _ = v27
    }
    v2 += 181
    v28 := []int{}
    v29 := GenS1{F0: -1}
    if (true && false) {
        v30 := GenS3{F0: -49}
        var v31 float32 = -2.25
        var v32 uint64 = (v1 ^ v1)
        _ = v30
        _ = v31
        _ = v32
    }
    switch 1 {
    case -930:
        v33 := GenS3{F0: genFn0((!(genFn1(18446744073709551615, "a", 2) < 323)))}
        _ = v33
    default:
        var v34 int8 = genFn0(true)
        _ = v34
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        var v36 uint8 = 42
        fmt.Println(v2)
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
}
