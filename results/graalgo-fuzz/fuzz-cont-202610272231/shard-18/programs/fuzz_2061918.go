package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint8, p1 int64) int32 {
    p1 = (p1 - p1)
    var v5 string = ((("\n\t" + (((("foo bar" + "") + "\n\t") + ("hello" + ("a" + ""))) + "")) + "foo bar") + "\n\t")
    v6 := GenS1{F0: (p0 | (p0 >> ((42) & 63)))}
    v7 := []int{4294967295, 4294967295, 2147483647, -2147483648}
    if (!false) {
        if len(v7) > 0 {
            v7[0] = -59
        }
        if false {
            fmt.Println(p0, v5, v5)
        } else {
            if len(v7) > 0 {
                v7[0] = -9223372036854775808
            }
            p1 = p1
            p1 += 1
        }
        v8 := GenS2{F0: p0, F1: 255}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 266
}

func main() {
    var v1 uint64 = 0
    var v2 float64 = float64((0.1 - -2.25))
    v3 := []int{}
    v4 := make(map[uint32]string)
    v2 = v2
    v9 := GenS2{F0: 1, F1: 42}
    switch 0 {
    case 9223372036854775807:
        var v10 uint32 = 2147483648
        _ = v10
    case 9:
        v11 := GenS1{F0: 1}
        _ = v11
    case -840:
        if true {
            for fi12 := 0; fi12 < 18; fi12++ {
                v1 -= (v1 / ((v1) | 1))
                if true {
                    fmt.Println(v2, fi12, v2)
                    v2 = v2
                    switch (("foo bar" + "hello") + ("a" + "\n\t")) {
                    case "hello":
                        for fi13 := 0; fi13 < 17; fi13++ {
                            if (4294967295 < 255) {
                                v2 = ((0.0 - v2) + 0.0)
                                var v14 bool = false
                                var v15 uint16 = 681
                                _ = v14
                                _ = v15
                            } else {
                                v3 = append(v3, 671)
                            }
                            var v16 uint8 = 255
                            _ = fi13
                            _ = v16
                        }
                    case "a":
                        var v17 float64 = 3.14159
                        _ = v17
                    }
                }
                if (false && (v2 <= float64((((float32(v2) * (-2.25 * 0.1)) - 1e10) * 1.5)))) {
                    v18 := GenS1{F0: 0}
                    _ = len(v4)
                    _ = v18
                }
                _ = fi12
            }
        }
    }
    v3 = append(v3, 42)
    v19 := GenS1{F0: 40}
    var v20 int8 = 0
    fmt.Println(v2)
    v2 += (v2 * float64(1.5))
    var v21 int16 = 1
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 int16 = (v21 - v21)
        _ = fi22
        _ = v23
    }
    switch -1 {
    case -265:
        v21 = -1
    case 639:
        for fi24 := 0; fi24 < 4; fi24++ {
            v21 = v21
            v1 += 18446744073709551615
            v1 -= v1
            _ = fi24
        }
    default:
        if (true && false) {
            if len(v3) > 0 {
                v3[0] = 32767
            }
        }
    }
    switch ("foo bar" + ("" + "a")) {
    case "hello":
        if (-1 > -1) {
            if false {
                var v25 uint32 = 2
                for fi26 := 0; fi26 < 11; fi26++ {
                    var v27 uint16 = 42
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
            if (509 >= 0) {
                var v28 int8 = (^(v20))
                fmt.Println(v2, v1)
                v29 := v4[2147483648]
                _ = v28
                _ = v29
            }
            for fi30 := 0; fi30 < 1; fi30++ {
                v4[1] = ("hello" + (("hello" + "hello") + "foo bar"))
                _ = fi30
            }
        } else {
            v1 -= v1
            fmt.Println(v1)
            v21 = -235
        }
    case "":
        if (3.14159 > float32(v2)) {
            switch 128 {
            case -328:
                var v31 uint16 = 32767
                _ = v31
            }
        }
    default:
        if (65535 != 2) {
            var v32 int16 = (-(v21))
            _ = v32
        } else {
            v2 -= v2
        }
    }
    delete(v4, uint32(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
