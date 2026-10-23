package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
func genFn0(p0 uint) int64 {
    if true {
        var v6 uint = 95
        v7 := []int{}
        switch "" {
        case "foo bar":
            var v8 uint16 = 32767
            _ = v8
        case "hello":
            fmt.Println(v6, v6)
        default:
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := GenS1{F0: (v6 * v6)}
                _ = fi9
                _ = v10
            }
        }
        _ = v6
        _ = v7
    } else {
        p0 += 0
        fmt.Println(p0)
    }
    var v11 uint8 = 4
    fmt.Println(p0, p0)
    for fi12 := 0; fi12 < 16; fi12++ {
        switch "" {
        case "foo bar":
            fmt.Println(p0)
        case "hello":
            p0 += (p0 & 2)
        case "":
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := []int{fi12}
                v15 := GenS2{F0: true}
                if len(v14) > 0 {
                    v14[0] = (fi13 % (((fi12 % ((42) | 1))) | 1))
                }
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            var v16 uint16 = 0
            _ = v16
        }
        var v17 int32 = -2147483648
        p0 += (p0 ^ p0)
        _ = fi12
        _ = v17
    }
    _ = p0
    _ = v11
    return -1
}
func genFn1(p0 uint16, p1 uint, p2 int16) uint32 {
    if true {
        var v18 bool = true
        _ = v18
    }
    p2 += (p2 << (((p1 % (((p1 * p1)) | 1))) & 63))
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("hello" + ("foo bar" + "")), "a"}
    var v3 int16 = -832
    var v4 uint16 = 714
    var v5 float32 = ((((0.1 + 3.14159) / -2.25) + (0.1 * 0.0)) * -2.25)
    fmt.Println(v5)
    for fi19 := 0; fi19 < 10; fi19++ {
        fmt.Println(fi19, fi19)
        _ = fi19
    }
    if (1 > 18446744073709551615) {
        v5 = (3.14159 + ((v5 / ((-2.25 + -2.25) - (float32((3.14159 + float64((v5 / 3.14159)))) + 0.1))) / 1.5))
    }
    var v20 uint32 = 1
    var v21 int = 0
    v21 -= (v21 << ((255) & 63))
    v22 := GenS2{F0: (((float32((0.1 - (1e10 * 1e10))) >= 1e10) && false) || true)}
    if (4294967295 > 4294967295) {
        v23 := v1[genFn1(v4, 255, v3)]
        if (("hello" < "hello") || (120 < 255)) {
            v3 = (v3 * v3)
            fmt.Println(v21)
        } else {
            for fi24 := 0; fi24 < 5; fi24++ {
                v25 := GenS1{F0: 1}
                var v26 float64 = (float64(((v5 / 1.0) + ((float32(3.14159) + v5) / float32(-2.25)))) * float64(v5))
                _ = fi24
                _ = v25
                _ = v26
            }
            var v27 int64 = -673
            _ = v27
        }
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
