package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int32; F1 uint32; F2 uint64 }
func genFn0(p0 uint32) uint8 {
    if false {
        v7 := []int{459}
        _ = v7
    }
    v8 := []int{0}
    v9 := GenS1{F0: 194, F1: 255}
    if true {
        var v10 uint16 = uint16(p0)
        var v11 uint32 = p0
        _ = v10
        _ = v11
    } else {
        if len(v8) > 0 {
            v8[0] = 625
        }
        _ = len(v8)
    }
    var v12 int8 = -1
    _ = p0
    _ = v12
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 string, p1 uint32, p2 uint16) bool {
    p2 -= (p2 % (((p2 / (((p2 << ((4294967295) & 63))) | 1))) | 1))
    var v13 float32 = 0.1
    fmt.Println(p0, p0)
    v14 := []int{32768}
    var v15 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return false
}

func main() {
    var v1 uint64 = 539
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    var v4 uint32 = 2147483647
    var v5 int32 = 0
    v6 := []string{""}
    var v16 int64 = 0
    if (1 == 42) {
        var v17 uint16 = 287
        _ = v17
    }
    switch -1 {
    case 448:
        v1 -= v1
    case 4294967295:
        switch ("a" + ("" + "foo bar")) {
        case "\n\t":
            fmt.Println(v4, v1)
        default:
            v5 = (-(v5))
        }
    }
    var v18 uint32 = (v4 >> ((650) & 63))
    v1 += v1
    if ((v1 | 885) >= v1) {
        var v19 bool = (true && (false && (65535 == 0)))
        _ = v19
    }
    _ = len(v6)
    switch 0 {
    case 106:
        if (!true) {
            switch -9223372036854775808 {
            case 128:
                if true {
                    v20 := GenS1{F0: 0, F1: 622}
                    fmt.Println(v1)
                    v18 = v18
                    _ = v20
                }
            default:
                fmt.Println(v18, v16, v16)
            }
            v3[117] = ""
            if true {
                v21 := GenS1{F0: 255, F1: 42}
                var v22 int = 0
                var v23 int64 = v16
                _ = v21
                _ = v22
                _ = v23
            }
        }
    }
    switch -530 {
    case 1:
        for fi24 := 0; fi24 < 16; fi24++ {
            switch 0 {
            case 128:
                var v25 float32 = 1.5
                _ = v25
            case 555:
                v6 = append(v6, ("" + "foo bar"))
            default:
                for fi26 := 0; fi26 < 4; fi26++ {
                    fmt.Println(v1, v4)
                    v27 := GenS2{F0: -1, F1: fi24}
                    _ = fi26
                    _ = v27
                }
            }
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 uint16 = 2
                _ = fi28
                _ = v29
            }
            v30 := GenS1{F0: 235, F1: 65535}
            _ = fi24
            _ = v30
        }
    case 755:
        if (127 < 0) {
            var v31 int16 = 1000
            v6 = append(v6, "")
            _ = v31
        }
    case 42:
        if len(v6) > 0 {
            v6[0] = "hello"
        }
    default:
        v32 := GenS3{F0: -478, F1: (v4 * (v4 * 0)), F2: (^(v1))}
        _ = v32
    }
    v18 = (^(v4))
    var v33 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
