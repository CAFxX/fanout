package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint64 }
func genFn0(p0 uint, p1 uint32, p2 uint32) uint64 {
    v6 := GenS1{F0: 425, F1: 32767, F2: 0}
    v7 := GenS1{F0: -9223372036854775808, F1: 65535, F2: 86}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}
func genFn1() string {
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{fi8, (fi8 & fi8), fi8, (fi8 ^ (fi8 << ((1) & 63)))}
        var v10 int = -861
        var v11 int8 = 124
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 float32 = 1.5
    v13 := []int{-498, -278, -204}
    _ = v12
    _ = v13
    return ("hello" + (("a" + ("a" + "\n\t")) + ""))
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]bool)
    var v3 uint16 = 374
    v4 := []int{905, -1, 281, 32768}
    var v5 float32 = (((1e10 - 1e10) * 0.0) - 1e10)
    var v14 uint32 = 42
    v1 += (v1 * (v1 + (-(v1))))
    v15 := GenS1{F0: -9223372036854775808, F1: (v3 * v3), F2: 18446744073709551615}
    v16 := GenS1{F0: v1, F1: 1, F2: 42}
    switch "" {
    case "":
        v1 -= v1
    case "\n\t":
        fmt.Println(v1)
    }
    var v17 float64 = 1e10
    for fi18 := 0; fi18 < 9; fi18++ {
        v2[fi18] = (false && true)
        if true {
            var v19 uint8 = 1
            switch (("hello" + "") + ((("a" + ("a" + "foo bar")) + "") + "foo bar")) {
            case "hello":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "foo bar":
                var v20 int8 = -1
                _ = v20
            }
            if (true || true) {
                var v21 int8 = 1
                if true {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        var v23 uint64 = genFn0(540, (^(v14)), v14)
                        var v24 int8 = v21
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    var v25 int16 = 32767
                    _ = v25
                } else {
                    var v26 uint = 18446744073709551615
                    var v27 uint8 = v19
                    _ = v26
                    _ = v27
                }
                _ = v21
            }
            _ = v19
        } else {
            switch fi18 {
            case 127:
                for fi28 := 0; fi28 < 9; fi28++ {
                    v14 += uint32(v3)
                    v29 := GenS1{F0: v1, F1: (v3 >> ((2) & 63)), F2: 1}
                    if ((32767 == 1) && (fi18 == (fi18 >> ((42) & 63)))) {
                        var v30 float32 = 1e10
                        _ = v30
                    } else {
                        fmt.Println(v5, fi18, v17)
                        v31 := GenS1{F0: v1, F1: v3, F2: 0}
                        _ = v31
                    }
                    fmt.Println(v5, v14)
                    fmt.Println(v3)
                    _ = fi28
                    _ = v29
                }
            case -1:
                v14 -= (v14 % ((v14) | 1))
            case -954:
                v5 += (1e10 / ((1e10 / v5) + ((v5 / -2.25) + v5)))
            }
            v32 := GenS1{F0: (v1 | v1), F1: v3, F2: 1}
            for fi33 := 0; fi33 < 14; fi33++ {
                fmt.Println(fi18)
                v5 += 1.5
                if false {
                    var v34 float32 = float32(v17)
                    _ = v34
                }
                _ = fi33
            }
            _ = v32
        }
        _ = fi18
    }
    switch -538 {
    case -17:
        fmt.Println(v3, v17, v5)
    default:
        fmt.Println(v5, v14)
    }
    var v35 int8 = 64
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
