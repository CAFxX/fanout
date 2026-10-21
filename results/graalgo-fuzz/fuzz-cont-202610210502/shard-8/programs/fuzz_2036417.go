package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint32 }
type GenS2 struct { F0 string; F1 int64; F2 uint16 }
type GenS3 struct { F0 string; F1 string }
func genFn0(p0 int64) string {
    p0 = p0
    v5 := GenS1{F0: 0, F1: ("foo bar" + "\n\t"), F2: 65535}
    if (1e10 > float32(((1.5 / 1.0) - 1e10))) {
        v6 := []int{127, 9223372036854775807}
        switch 4294967295 {
        case 128:
            p0 += p0
        default:
            var v7 float32 = (0.0 * 0.1)
            _ = v7
        }
        p0 -= (p0 / ((p0) | 1))
        _ = v6
    }
    fmt.Println(p0)
    switch -842 {
    case 9223372036854775807:
        fmt.Println(p0)
    default:
        p0 += (^(p0))
    }
    _ = p0
    _ = v5
    return ("\n\t" + ("a" + "a"))
}
func genFn1(p0 int8, p1 uint8, p2 int16) uint16 {
    var v8 int32 = -328
    if false {
        if false {
            v9 := []int{int(v8), 32767, 4294967295}
            var v10 float64 = ((-2.25 + ((0.0 - float64(0.0)) * ((-2.25 / 1.5) + ((3.14159 / 3.14159) + 0.0)))) + (float64(3.14159) * -2.25))
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 string = "a"
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 int = (-(fi13))
                    _ = fi13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
            _ = v10
            _ = v9
        }
    } else {
        v15 := []int{2147483647, -465}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 255
}

func main() {
    var v1 bool = (-1 != 45)
    v2 := make(map[uint32]int)
    var v3 uint32 = 2
    v4 := []string{"a", ("hello" + ("" + ("hello" + ("a" + "hello"))))}
    delete(v2, v3)
    _ = len(v2)
    v16 := GenS2{F0: "foo bar", F1: 436, F2: 255}
    if (170 <= 128) {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
    } else {
        v17 := GenS3{F0: "a", F1: ("a" + "a")}
        v18 := GenS1{F0: 2, F1: "", F2: v3}
        _ = v17
        _ = v18
    }
    var v19 float32 = 0.1
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 int16 = 42
        v3 += 1
        for fi22 := 0; fi22 < 7; fi22++ {
            v21 += v21
            var v23 uint = 2
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    if (127 <= -1) {
        v3 = v3
    } else {
        v19 = ((-2.25 / (v19 - 1e10)) / 1e10)
    }
    _ = len(v2)
    var v24 int32 = 0
    v25 := GenS2{F0: "hello", F1: 1, F2: genFn1(int8(v24), 198, 42)}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
