package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 string }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint = 491
    var v2 uint16 = 42
    v3 := make(map[uint32]int)
    v2 = v2
    switch (("\n\t" + "foo bar") + "\n\t") {
    case "hello":
        if (0 > -161) {
            for fi4 := 0; fi4 < 1; fi4++ {
                v2 -= 0
                var v5 bool = (2 == v2)
                v6 := []int{fi4, (^(fi4)), (fi4 % ((fi4) | 1)), (fi4 / (((fi4 - (fi4 ^ (fi4 & (fi4 - (^(fi4))))))) | 1))}
                _ = fi4
                _ = v5
                _ = v6
            }
        }
    default:
        if (0.0 != 0.1) {
            v7 := []int{42, -2147483648, 128, -1}
            _ = v7
        }
    }
    fmt.Println(v2, v2)
    if true {
        if (true || (!(255 != 38))) {
            var v8 int = 0
            var v9 float64 = (-2.25 * (1e10 / (-2.25 * ((((0.1 - 0.1) * float64(1e10)) + float64(0.1)) + (1.5 - ((0.1 / float64(0.1)) * (1e10 / 1.0)))))))
            var v10 uint32 = 2147483647
            _ = v10
            _ = v8
            _ = v9
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 8; fi12++ {
                v1 -= (v1 / (((v1 / ((v1) | 1))) | 1))
                fmt.Println(v2, v1, fi11)
                fmt.Println(v1)
                _ = fi12
            }
            _ = fi11
        }
        for fi13 := 0; fi13 < 13; fi13++ {
            for fi14 := 0; fi14 < 18; fi14++ {
                v15 := []int{fi14, fi13}
                v16 := GenS1{F0: v2}
                if true {
                    if false {
                        if len(v15) > 0 {
                            v15[0] = 128
                        }
                        var v17 int16 = 42
                        var v18 int = 4294967295
                        _ = v17
                        _ = v18
                    }
                    v1 = v1
                    fmt.Println(v1)
                } else {
                    v15 = append(v15, (fi13 ^ (fi13 | 872)))
                }
                fmt.Println(v1)
                v19 := GenS1{F0: 255}
                var v20 int64 = -9223372036854775808
                switch (((("hello" + ("a" + ("foo bar" + "a"))) + ("\n\t" + ("\n\t" + "a"))) + "\n\t") + "\n\t") {
                case "\n\t":
                    var v21 int = (fi14 | fi14)
                    _ = v21
                }
                _ = fi14
                _ = v15
                _ = v16
                _ = v19
                _ = v20
            }
            _ = fi13
        }
    } else {
        v22 := []int{-9223372036854775808, 528}
        for fi23 := 0; fi23 < 13; fi23++ {
            fmt.Println(v1, v2, v1)
            if (((3.14159 * -2.25) != 0.0) || true) {
                var v24 uint16 = v2
                switch -9223372036854775808 {
                case -808:
                    var v25 uint = (v1 / (((v1 / ((255) | 1))) | 1))
                    _ = v25
                case 32767:
                    _ = len(v22)
                }
                _ = v24
            }
            _ = fi23
        }
        _ = v22
    }
    v2 = (v2 % (((v2 >> ((4294967295) & 63))) | 1))
    v1 -= (v1 >> ((253) & 63))
    fmt.Println(v1, v1)
    var v26 bool = false
    fmt.Println(v1, v26, v1)
    var v27 uint = (v1 & (v1 >> (((v1 * 1)) & 63)))
    var v28 uint64 = 1
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := GenS2{F0: 17, F1: 2147483648, F2: ""}
        v2 = (v2 - (^(v2)))
        _ = fi29
        _ = v30
    }
    fmt.Println(v1)
    var v31 uint8 = 1
    v32 := GenS3{F0: 789}
    var v33 float64 = 3.14159
    var v34 bool = v26
    if v34 {
        fmt.Println(v33, v34, v34)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
