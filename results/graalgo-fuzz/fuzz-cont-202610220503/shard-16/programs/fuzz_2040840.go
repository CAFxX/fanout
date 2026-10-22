package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int32 }
type GenS2 struct { F0 int16; F1 int32; F2 uint8 }
func genFn0(p0 uint32, p1 uint64, p2 string) uint {
    if (-14 == 127) {
        p1 -= 42
        v4 := []int{127, 42, 4294967295, int(p1)}
        p0 = p0
        _ = v4
    } else {
        v5 := GenS1{F0: 1, F1: "foo bar", F2: -523}
        _ = v5
    }
    v6 := GenS1{F0: -1, F1: p2, F2: 678}
    v7 := []int{}
    var v8 uint64 = (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint32 = 0
    var v2 bool = (!(v1 < v1))
    var v3 int16 = -1
    v9 := GenS1{F0: (v3 / (((^(v3))) | 1)), F1: "\n\t", F2: -1}
    v3 -= -1
    var v10 float32 = 3.14159
    if (752 >= 42) {
        var v11 uint16 = 255
        v12 := GenS1{F0: v3, F1: "a", F2: -554}
        v2 = (42 <= -237)
        _ = v11
        _ = v12
    } else {
        var v13 bool = ((87 == 1) && v2)
        v10 -= (3.14159 - (-2.25 - ((v10 * v10) - v10)))
        _ = v13
    }
    switch "a" {
    case "":
        v14 := []int{}
        _ = v14
    case "\n\t":
        if (0 <= 427) {
            v10 += v10
            var v15 string = ""
            v10 += v10
            _ = v15
        } else {
            if ((v3 % ((v3) | 1)) != v3) {
                v3 -= v3
                v16 := GenS2{F0: (v3 | v3), F1: 42, F2: 255}
                v17 := []int{0, 2147483647, 2147483647, -9223372036854775808}
                _ = v16
                _ = v17
            }
        }
    default:
        v3 -= (v3 ^ v3)
    }
    if (((1 != 0) && (-1 == 17)) || true) {
        if ((1 > 337) && (!(42 == 921))) {
            v10 -= v10
            v10 = (1.5 * (1e10 / v10))
            for fi18 := 0; fi18 < 11; fi18++ {
                for fi19 := 0; fi19 < 6; fi19++ {
                    switch ((("" + "\n\t") + "\n\t") + (("\n\t" + "") + "")) {
                    case "":
                        v20 := []int{fi19, 4294967295, (fi18 % ((-409) | 1))}
                        _ = v20
                    case "a":
                        v10 += v10
                    case "hello":
                        var v21 int16 = (v3 + -756)
                        _ = v21
                    default:
                        var v22 uint32 = 4294967295
                        _ = v22
                    }
                    _ = fi19
                }
                var v23 uint32 = 0
                var v24 int64 = -9223372036854775808
                _ = fi18
                _ = v23
                _ = v24
            }
        }
        switch 32767 {
        case -9223372036854775808:
            for fi25 := 0; fi25 < 12; fi25++ {
                fmt.Println(v2, v10, v10)
                v26 := []int{9223372036854775807, (fi25 % (((fi25 & (fi25 << ((1) & 63)))) | 1))}
                _ = fi25
                _ = v26
            }
        default:
            for fi27 := 0; fi27 < 9; fi27++ {
                v3 = v3
                v28 := []int{(fi27 % ((-301) | 1)), fi27, (fi27 >> ((uint(v1)) & 63)), 346}
                _ = fi27
                _ = v28
            }
        }
    }
    switch 127 {
    case 0:
        var v29 uint32 = v1
        _ = v29
    case -43:
        for fi30 := 0; fi30 < 5; fi30++ {
            v31 := []int{(fi30 / ((-906) | 1)), (fi30 - fi30), fi30, fi30}
            _ = fi30
            _ = v31
        }
    }
    v32 := []int{32768, 4294967295, -1, 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v9
}
