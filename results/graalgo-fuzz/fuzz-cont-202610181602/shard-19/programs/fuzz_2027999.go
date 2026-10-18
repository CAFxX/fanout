package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 string; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 string }
func genFn0(p0 int) uint16 {
    p0 += (p0 + -135)
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 667, F1: ((65535 != 1) && true)}
        var v8 int8 = 39
        for fi9 := 0; fi9 < 9; fi9++ {
            if false {
                var v10 uint = 18446744073709551615
                p0 += 127
                _ = v10
            } else {
                v11 := []int{fi9, (fi9 ^ (-(fi9)))}
                _ = len(v11)
                v8 = (v8 % ((v8) | 1))
                if (0 != 0) {
                    var v12 int = 87
                    fmt.Println(v8, p0, v8)
                    v13 := GenS2{F0: "hello", F1: 4294967295, F2: 2}
                    _ = v12
                    _ = v13
                }
                if (18446744073709551615 == 18446744073709551615) {
                    v14 := GenS1{F0: 0, F1: true}
                    var v15 float32 = float32(1e10)
                    var v16 int32 = -1
                    var v17 int = 0
                    fmt.Println(p0)
                    if len(v11) > 0 {
                        v11[0] = -2147483648
                    }
                    v17 = -9223372036854775808
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    v8 = (v8 << ((18446744073709551615) & 63))
                    v8 = (v8 << ((1) & 63))
                    if ((0.1 - 0.0) == 3.14159) {
                        v18 := GenS2{F0: ("foo bar" + ""), F1: 2, F2: 255}
                        _ = len(v11)
                        _ = v18
                    } else {
                        v19 := GenS3{F0: 0, F1: 65535}
                        fmt.Println(v8, v8)
                        _ = v19
                    }
                }
                _ = v11
            }
            fmt.Println(v8, fi9)
            _ = fi9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v20 := GenS2{F0: ("foo bar" + "foo bar"), F1: 548, F2: 127}
    v21 := GenS3{F0: 32767, F1: 2147483647}
    if (!((2147483647 > -230) && (-9223372036854775808 == 922))) {
        fmt.Println(p0)
        var v22 uint32 = 2147483647
        _ = v22
    }
    _ = p0
    _ = v20
    _ = v21
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{9223372036854775807, 127}
    var v3 int8 = 42
    v4 := []string{"\n\t", "", "hello"}
    v5 := make(map[uint32]string)
    _ = len(v5)
    v23 := GenS4{F0: ("foo bar" + "foo bar")}
    if false {
        v3 = (v3 | -128)
        v24 := GenS4{F0: ("hello" + "a")}
        if (127 != v3) {
            fmt.Println(v3)
        } else {
            delete(v1, "foo bar")
        }
        _ = v24
    } else {
        v3 += (v3 >> ((614) & 63))
        var v25 string = "a"
        _ = v25
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        v3 = (v3 | v3)
        _ = fi26
    }
    v3 = v3
    var v27 uint64 = 2
    if (false && true) {
        v27 = 0
    }
    v28 := GenS1{F0: 455, F1: false}
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
