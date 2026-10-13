package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 int8) int16 {
    p0 = (p0 >> ((255) & 63))
    if (2 > 42) {
        v7 := []int{128, 32767, 42}
        v8 := GenS1{F0: "", F1: 42, F2: ""}
        _ = v7
        _ = v8
    }
    _ = p0
    return -1
}
func genFn1(p0 uint16, p1 bool) uint64 {
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            p0 = 65535
            _ = fi10
        }
        _ = fi9
    }
    var v11 int64 = 9223372036854775807
    p0 -= p0
    _ = p0
    _ = p1
    _ = v11
    return 2
}

func main() {
    v1 := []string{"\n\t", ("foo bar" + (((("" + ("\n\t" + "a")) + "") + ("hello" + "hello")) + ("foo bar" + "foo bar"))), "hello"}
    v2 := []uint32{343}
    v3 := []int{-655}
    var v4 int32 = -2147483648
    var v5 float32 = ((-2.25 - (float32(0.0) + 1e10)) - -2.25)
    v6 := make(map[int]int)
    var v12 float32 = v5
    delete(v6, 128)
    v5 = v12
    if (false && true) {
        var v13 int8 = 62
        _ = v13
    } else {
        for fi14 := 0; fi14 < 2; fi14++ {
            for fi15 := 0; fi15 < 7; fi15++ {
                v5 = ((float32(-2.25) / v5) * v5)
                var v16 uint64 = genFn1(1, (30 != 110))
                _ = fi15
                _ = v16
            }
            if ((0.0 * -2.25) > ((0.1 - 1e10) - -2.25)) {
                v12 += ((3.14159 * (float32(1.5) + (v5 * v5))) / (v5 + v5))
                if (true && (16 > 42)) {
                    if (0 > 0) {
                        v17 := GenS2{F0: (v4 | v4)}
                        _ = v17
                    }
                    v4 -= 1
                }
                if false {
                    var v18 int8 = -1
                    _ = v18
                }
            } else {
                fmt.Println(v4, v12, v5)
                v4 -= (v4 & v4)
            }
            v1 = append(v1, "foo bar")
            _ = fi14
        }
        _ = len(v6)
    }
    var v19 uint64 = 0
    v5 = 3.14159
    for fi20 := 0; fi20 < 3; fi20++ {
        v5 = ((v12 * v5) * v12)
        if len(v3) > 0 {
            v3[0] = (^(fi20))
        }
        v12 = 0.1
        _ = fi20
    }
    if (-2.25 == (((float32(float64((v5 - (0.0 - (-2.25 + (float32(-2.25) * (v5 / (v5 * v5)))))))) - v12) + float32(1e10)) * 1e10)) {
        v21 := GenS3{F0: 0}
        switch 9223372036854775807 {
        case 1:
            var v22 uint = 18446744073709551615
            _ = v22
        case 32768:
            for fi23 := 0; fi23 < 11; fi23++ {
                delete(v6, (fi23 | fi23))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi24 := 0; fi24 < 13; fi24++ {
                    fmt.Println(v19, fi23)
                    v6[(fi23 * fi23)] = (^(fi24))
                    _ = fi24
                }
                _ = fi23
            }
        }
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
