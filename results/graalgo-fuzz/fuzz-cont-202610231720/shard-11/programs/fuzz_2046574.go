package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 string }
func genFn0(p0 int16, p1 uint, p2 uint8) string {
    v6 := []int{42, -2147483648, 32767, 42}
    var v7 float64 = ((1.5 - 3.14159) * 3.14159)
    var v8 string = ("a" + ("" + ("foo bar" + "\n\t")))
    if ((42 < 189) && (true && (p1 == 1))) {
        switch (((v8 + "hello") + ("\n\t" + v8)) + v8) {
        case "":
            for fi9 := 0; fi9 < 6; fi9++ {
                _ = len(v6)
                v7 = v7
                _ = fi9
            }
        case "hello":
            fmt.Println(p2, v7, v8)
        default:
            if (-2.25 >= float32((v7 / -2.25))) {
                switch "hello" {
                case "hello":
                    v10 := GenS1{F0: true}
                    _ = v10
                }
                v11 := GenS1{F0: (!(p2 <= p2))}
                p1 += p1
                _ = v11
            } else {
                for fi12 := 0; fi12 < 4; fi12++ {
                    if (-1 >= 127) {
                        v8 = v8
                        p1 += (p1 & 975)
                    } else {
                        for fi13 := 0; fi13 < 11; fi13++ {
                            var v14 uint8 = (p2 ^ p2)
                            v15 := GenS1{F0: true}
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        v16 := GenS2{F0: v8, F1: (^(p1))}
                        v17 := GenS2{F0: v8, F1: p1}
                        _ = v16
                        _ = v17
                    }
                    fmt.Println(fi12, v7, p1)
                    fmt.Println(p0, v7)
                    _ = fi12
                }
            }
        }
        fmt.Println(p2, v7, p0)
    } else {
        var v18 int8 = int8(p2)
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 int8 = 10
    v2 := []int{}
    v3 := []uint32{}
    var v4 uint8 = 2
    var v5 int8 = (v1 * (v1 + v1))
    if (4294967295 == 1) {
        v19 := GenS3{F0: (("" + "a") + (("" + "foo bar") + ((("\n\t" + ("hello" + "foo bar")) + "foo bar") + "\n\t")))}
        _ = v19
    } else {
        v3 = append(v3, 1)
        v3 = append(v3, 2147483647)
        fmt.Println(v1)
    }
    v4 -= (v4 << ((42) & 63))
    for fi20 := 0; fi20 < 10; fi20++ {
        v1 -= (v1 % (((-(v5))) | 1))
        v5 = -128
        v5 += 103
        _ = fi20
    }
    v3 = append(v3, 4294967295)
    v4 = (v4 - v4)
    var v21 uint32 = 4294967295
    for fi22 := 0; fi22 < 19; fi22++ {
        if len(v2) > 0 {
            v2[0] = 0
        }
        switch 32767 {
        case 128:
            v1 -= v5
        }
        _ = fi22
    }
    var v23 int16 = 0
    for fi24 := 0; fi24 < 18; fi24++ {
        fmt.Println(v5, v5)
        for fi25 := 0; fi25 < 12; fi25++ {
            v5 = v1
            var v26 int8 = int8(v23)
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    var v27 int64 = 9223372036854775807
    v4 += v4
    v28 := GenS1{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
