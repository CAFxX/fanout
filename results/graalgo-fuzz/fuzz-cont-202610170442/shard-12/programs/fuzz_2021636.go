package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint, p1 int8) int8 {
    v4 := GenS1{F0: -9223372036854775808, F1: (p0 ^ (p0 << ((p0) & 63)))}
    var v5 uint = p0
    p1 = (p1 >> ((62) & 63))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (^(p1))
}
func genFn1(p0 uint16) int16 {
    p0 -= p0
    v6 := []int{147, 32768, 127}
    switch -1 {
    case 127:
        for fi7 := 0; fi7 < 4; fi7++ {
            p0 += (^(p0))
            v8 := GenS1{F0: -952, F1: 18446744073709551615}
            _ = fi7
            _ = v8
        }
    case -809:
        for fi9 := 0; fi9 < 8; fi9++ {
            p0 += 6
            p0 += p0
            _ = fi9
        }
    default:
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(fi10, fi10)
            var v11 uint8 = 59
            _ = fi10
            _ = v11
        }
    }
    if (!true) {
        var v12 uint16 = p0
        _ = v12
    } else {
        var v13 int64 = -1
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint8 = 255
            var v16 uint = 65535
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    }
    v17 := GenS2{F0: 158}
    _ = p0
    _ = v17
    _ = v6
    return 364
}

func main() {
    v1 := []string{("a" + ("" + "foo bar")), (((("hello" + ("foo bar" + "foo bar")) + "hello") + "\n\t") + ""), "hello"}
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v18 int32 = 2147483647
    var v19 string = "hello"
    fmt.Println(v18, v18)
    var v20 int64 = 1
    var v21 uint32 = 256
    if (849 >= 2) {
        v22 := v2[4294967295]
        v23 := GenS3{F0: 338}
        switch int(v20) {
        case -1:
            var v24 float32 = (1.5 / 3.14159)
            _ = v24
        case 127:
            for fi25 := 0; fi25 < 19; fi25++ {
                for fi26 := 0; fi26 < 4; fi26++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi26
                }
                _ = fi25
            }
        }
        _ = v22
        _ = v23
    } else {
        fmt.Println(v18, v19)
    }
    v21 += 4294967295
    v18 = (v18 | (v18 - 189))
    var v27 int8 = 100
    var v28 int32 = (v18 << ((13) & 63))
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
}
