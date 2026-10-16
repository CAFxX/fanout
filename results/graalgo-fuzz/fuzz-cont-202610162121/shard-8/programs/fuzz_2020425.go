package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int32 }
type GenS2 struct { F0 int; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint, p1 uint64) bool {
    p1 -= (^(p1))
    if false {
        for fi5 := 0; fi5 < 16; fi5++ {
            v6 := GenS1{F0: -182, F1: ((("a" + "") + "foo bar") + "\n\t"), F2: 1}
            if true {
                v7 := []int{fi5, (^(fi5)), 32767, fi5}
                var v8 uint16 = uint16(p1)
                _ = v7
                _ = v8
            }
            switch ((("" + "a") + ("" + "hello")) + "\n\t") {
            case "\n\t":
                for fi9 := 0; fi9 < 7; fi9++ {
                    p0 += p0
                    switch (-(fi9)) {
                    case -112:
                        var v10 uint8 = 42
                        _ = v10
                    }
                    _ = fi9
                }
            case "a":
                fmt.Println(fi5, fi5, fi5)
            }
            _ = fi5
            _ = v6
        }
    }
    var v11 uint8 = 0
    var v12 int64 = 0
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return (0 < -2147483648)
}
func genFn1(p0 uint32, p1 uint8, p2 string) uint64 {
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 float32 = -2.25
        _ = fi13
        _ = v14
    }
    if false {
        v15 := GenS2{F0: 42, F1: (p0 >> ((4294967295) & 63)), F2: 1}
        _ = v15
    } else {
        v16 := GenS3{F0: (p1 & (p1 * (^(p1))))}
        v17 := []int{-9223372036854775808}
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return 213
}

func main() {
    var v1 string = "hello"
    var v2 uint32 = 42
    v3 := make(map[uint32]string)
    var v4 int64 = -9223372036854775808
    v18 := []int{32767, -2147483648}
    if (42 > 703) {
        v18 = append(v18, 4294967295)
        v1 = ""
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := GenS2{F0: 32767, F1: 4294967295, F2: 919}
            _ = fi19
            _ = v20
        }
    }
    fmt.Println(v2)
    var v21 int64 = (v4 - v4)
    delete(v3, v2)
    var v22 float32 = 0.1
    fmt.Println(v2, v21, v4)
    var v23 float64 = 0.0
    v23 += 1.5
    v18 = append(v18, 390)
    var v24 float32 = 0.0
    v3[(v2 / ((v2) | 1))] = "\n\t"
    v23 -= ((v23 / (float64(v22) + v23)) / ((float64((v22 - (float32((v23 - float64(v22))) + ((v24 - 3.14159) - (v24 / (v22 - (float32(3.14159) - (0.1 + v22)))))))) - (-2.25 - (v23 / (float64(v24) - v23)))) + 1e10))
    for fi25 := 0; fi25 < 1; fi25++ {
        v23 += v23
        var v26 int32 = -2147483648
        var v27 float64 = v23
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
