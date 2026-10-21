package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 uint }
func genFn0(p0 int) uint {
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: 2147483647, F1: "a", F2: ("" + "hello")}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS2{F0: 360}
        if false {
            v10 := GenS2{F0: 94}
            p0 = (fi8 << ((493) & 63))
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    p0 = p0
    var v11 int = (p0 / ((p0) | 1))
    fmt.Println(p0, v11)
    _ = p0
    _ = v11
    return 84
}
func genFn1(p0 int, p1 int64, p2 uint8) uint8 {
    fmt.Println(p2)
    if (!true) {
        for fi12 := 0; fi12 < 15; fi12++ {
            if (p1 == p1) {
                fmt.Println(fi12, p1)
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 int = 42
                    v15 := []int{(fi13 | (^(fi13))), fi13, fi13}
                    if (!true) {
                        _ = len(v15)
                        p1 = (p1 - (p1 ^ 9223372036854775807))
                    } else {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        var v16 uint32 = 65535
                        _ = v16
                    }
                    v17 := GenS1{F0: 316, F1: "foo bar", F2: ("\n\t" + "")}
                    p1 -= p1
                    v15 = append(v15, fi13)
                    var v18 int64 = (^(p1))
                    p1 = v18
                    var v19 int = fi12
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v17
                    _ = v18
                    _ = v19
                }
                var v20 float64 = 0.1
                _ = v20
            } else {
                v21 := []int{-1, (fi12 * p0), fi12}
                _ = v21
            }
            _ = fi12
        }
        if (((-2.25 - ((1e10 * 0.0) - -2.25)) - 3.14159) < (0.0 - float32((3.14159 * 3.14159)))) {
            v22 := []int{(p0 / (((p0 % ((p0) | 1))) | 1)), 9223372036854775807}
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            switch "\n\t" {
            case "foo bar":
                if len(v22) > 0 {
                    v22[0] = p0
                }
            }
            _ = v22
        } else {
            var v23 uint8 = p2
            v24 := []int{(-(p0))}
            _ = v23
            _ = v24
        }
    }
    var v25 uint8 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return (p2 - p2)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 uint32 = 2
    var v4 int32 = 40
    var v5 float32 = 3.14159
    if (("\n\t" == (("hello" + "a") + "")) && true) {
        v4 = v4
        v26 := GenS1{F0: v3, F1: "\n\t", F2: ("" + "\n\t")}
        _ = v26
    }
    var v27 float32 = 1.5
    for fi28 := 0; fi28 < 15; fi28++ {
        v29 := GenS2{F0: 42}
        v30 := GenS2{F0: 18446744073709551615}
        var v31 int16 = 1000
        _ = fi28
        _ = v29
        _ = v30
        _ = v31
    }
    fmt.Println(v4, v5, v27)
    for fi32 := 0; fi32 < 2; fi32++ {
        v4 = v4
        var v33 int16 = -1
        _ = fi32
        _ = v33
    }
    var v34 int64 = -142
    var v35 int16 = 32767
    if (248 <= 454) {
        fmt.Println(v35)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
