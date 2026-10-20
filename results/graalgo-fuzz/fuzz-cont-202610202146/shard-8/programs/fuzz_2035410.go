package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int8; F2 uint }
type GenS5 struct { F0 string; F1 int8; F2 int16 }
func genFn0(p0 uint32, p1 uint32) uint32 {
    v7 := GenS1{F0: 0}
    if true {
        v8 := GenS2{F0: 17}
        if ((!(127 != -62)) && true) {
            fmt.Println(p0, p1)
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := []int{fi9, (-(fi9)), int(p0)}
                v11 := GenS2{F0: 204}
                _ = fi9
                _ = v10
                _ = v11
            }
        } else {
            fmt.Println(p0, p0)
            v12 := GenS2{F0: 0}
            _ = v12
        }
        _ = v8
    } else {
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := []int{fi13, fi13, -9223372036854775808, 2147483647}
            switch "hello" {
            case "foo bar":
                var v15 uint32 = (p0 & (p1 ^ p1))
                _ = v15
            default:
                v16 := GenS2{F0: 158}
                _ = v16
            }
            var v17 string = ("\n\t" + "\n\t")
            _ = fi13
            _ = v14
            _ = v17
        }
        v18 := GenS2{F0: 151}
        v19 := GenS3{F0: (true || ((0 == 0) && false))}
        _ = v18
        _ = v19
    }
    var v20 float64 = ((0.0 / 3.14159) + 0.1)
    _ = p0
    _ = p1
    _ = v20
    _ = v7
    return (p1 >> ((4294967295) & 63))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    var v5 uint8 = 240
    v6 := []uint32{}
    var v21 uint16 = 255
    v22 := GenS1{F0: 0}
    v21 = v21
    fmt.Println(v5, v21)
    switch ((("" + "foo bar") + "") + (("\n\t" + "foo bar") + ("\n\t" + ("" + ("foo bar" + "hello"))))) {
    case "hello":
        _ = len(v1)
    case "\n\t":
        var v23 int64 = -696
        _ = v23
    case "foo bar":
        var v24 uint64 = 42
        _ = v24
    }
    var v25 float64 = 0.1
    var v26 int = -2147483648
    v27 := GenS2{F0: (v5 + (v5 / ((2) | 1)))}
    if (65535 < 255) {
        for fi28 := 0; fi28 < 17; fi28++ {
            v26 -= (-(fi28))
            var v29 uint16 = (v21 + v21)
            v29 -= (^(v21))
            _ = fi28
            _ = v29
        }
        v30 := GenS4{F0: 0, F1: 42, F2: 1}
        var v31 int8 = 69
        _ = v30
        _ = v31
    } else {
        if len(v6) > 0 {
            v6[0] = 962
        }
        v5 = v5
        for fi32 := 0; fi32 < 20; fi32++ {
            var v33 uint32 = 2147483648
            fmt.Println(v5, v25)
            v33 = v33
            _ = fi32
            _ = v33
        }
    }
    if ((0.1 * (1e10 - v25)) != (((v25 + ((3.14159 + (v25 / 1.5)) / v25)) * v25) + v25)) {
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 uint = 0
            _ = fi34
            _ = v35
        }
        for fi36 := 0; fi36 < 5; fi36++ {
            switch 127 {
            case 532:
                v25 = 3.14159
            }
            v25 += v25
            _ = fi36
        }
        v37 := GenS5{F0: "a", F1: 108, F2: -539}
        _ = v37
    }
    for fi38 := 0; fi38 < 8; fi38++ {
        var v39 int = fi38
        v40 := GenS3{F0: (!true)}
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
