package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int }
type GenS2 struct { F0 uint16; F1 int32; F2 int32 }
func genFn0(p0 uint, p1 int16) int32 {
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(p1, p0, fi5)
        for fi6 := 0; fi6 < 4; fi6++ {
            p0 -= (p0 % ((p0) | 1))
            var v7 float32 = 1e10
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v8 := []int{}
    var v9 string = "foo bar"
    fmt.Println(v9, p1, p1)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1(p0 uint) string {
    v10 := GenS1{F0: 4294967295, F1: 602, F2: 299}
    fmt.Println(p0)
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := []int{}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    return ""
}

func main() {
    var v1 int64 = -1
    v2 := []uint32{}
    v3 := []string{"foo bar", "", ("\n\t" + "hello"), ""}
    v4 := []uint32{1, 591}
    switch "\n\t" {
    case "hello":
        fmt.Println(v1)
    }
    if ((float32((1e10 * 3.14159)) - float32(1.5)) >= (((0.1 * 1.5) + -2.25) / (0.1 - 1.5))) {
        v1 = -9223372036854775808
    } else {
        _ = len(v3)
        var v13 float64 = (-2.25 - 1e10)
        _ = v13
    }
    v1 += 632
    var v14 int8 = 0
    var v15 uint8 = 74
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 uint64 = 894
        if (float32((float64((3.14159 - 0.1)) * 0.0)) >= 0.1) {
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 string = "a"
                _ = fi18
                _ = v19
            }
            v15 = (v15 ^ v15)
            fmt.Println(v1, v1, v15)
        } else {
            fmt.Println(v1, v15)
            v4 = append(v4, 65535)
            var v20 string = "hello"
            _ = v20
        }
        v15 += (v15 - (v15 << ((1) & 63)))
        _ = fi16
        _ = v17
    }
    v2 = append(v2, 2)
    v1 -= (^(v1))
    v14 += v14
    for fi21 := 0; fi21 < 18; fi21++ {
        if (!("" > "")) {
            if (2147483647 < 2147483647) {
                for fi22 := 0; fi22 < 14; fi22++ {
                    var v23 int32 = int32(fi22)
                    v15 = 199
                    v14 -= v14
                    v1 -= v1
                    _ = fi22
                    _ = v23
                }
                v24 := GenS2{F0: 599, F1: 2147483647, F2: -170}
                _ = v24
            }
            var v25 float32 = (float32(0.1) - 3.14159)
            v1 -= (-(v1))
            _ = v25
        }
        _ = fi21
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        for fi27 := 0; fi27 < 17; fi27++ {
            v28 := GenS2{F0: 1, F1: 0, F2: 0}
            _ = fi27
            _ = v28
        }
        var v29 uint64 = 18446744073709551615
        _ = fi26
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
