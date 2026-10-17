package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 string }
func genFn0(p0 uint16, p1 uint16) int {
    v6 := []int{368, 4294967295}
    p0 = 255
    _ = p0
    _ = p1
    _ = v6
    return 1
}
func genFn1(p0 uint16, p1 uint8, p2 uint64) int {
    p2 -= (p2 / ((uint64(p1)) | 1))
    var v7 int16 = -1
    v8 := GenS1{F0: 18446744073709551615, F1: 153, F2: ""}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32768
}

func main() {
    var v1 int64 = 0
    v2 := []uint32{2}
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v5 int32 = -1
    v9 := GenS1{F0: 585, F1: 900, F2: ""}
    for fi10 := 0; fi10 < 7; fi10++ {
        v1 += (v1 * (v1 * v1))
        var v11 uint8 = 1
        var v12 int32 = (v5 >> ((396) & 63))
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        if true {
            v13 := GenS1{F0: 2, F1: 1, F2: ("a" + ("a" + ("a" + ("hello" + ""))))}
            if (!(-32768 < -1)) {
                var v14 uint16 = 123
                _ = v14
            }
            v15 := v4[("foo bar" + ("a" + "\n\t"))]
            _ = v13
            _ = v15
        } else {
            if (65535 == 1) {
                v16 := GenS1{F0: 0, F1: 164, F2: ""}
                v5 = 1
                _ = v16
            }
            for fi17 := 0; fi17 < 5; fi17++ {
                if (false || ((0 != 0) || false)) {
                    if (0 <= 18446744073709551615) {
                        v5 += v5
                    }
                }
                _ = fi17
            }
        }
        var v18 uint64 = 230
        var v19 int = -2147483648
        _ = v18
        _ = v19
    } else {
        fmt.Println(v1, v1)
        var v20 float32 = 0.1
        for fi21 := 0; fi21 < 5; fi21++ {
            for fi22 := 0; fi22 < 18; fi22++ {
                v20 -= (float32(0.0) + 1.5)
                v23 := GenS1{F0: 224, F1: 400, F2: ""}
                var v24 uint = 2
                _ = fi22
                _ = v23
                _ = v24
            }
            v25 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: ("foo bar" + "hello")}
            _ = fi21
            _ = v25
        }
        _ = v20
    }
    v3["hello"] = "hello"
    v1 -= v1
    var v26 float32 = -2.25
    v2 = append(v2, 1)
    v27 := GenS1{F0: 42, F1: 346, F2: "\n\t"}
    for fi28 := 0; fi28 < 12; fi28++ {
        fmt.Println(v5, v5, v5)
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
