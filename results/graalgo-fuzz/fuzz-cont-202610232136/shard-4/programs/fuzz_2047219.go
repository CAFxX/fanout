package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint8 }
func genFn0(p0 int8, p1 uint) uint {
    fmt.Println(p1, p0)
    if (43 > 95) {
        if true {
            fmt.Println(p0, p0)
        }
    } else {
        for fi7 := 0; fi7 < 16; fi7++ {
            p0 += -128
            switch fi7 {
            case 1:
                p1 -= 1
            case 0:
                v8 := GenS1{F0: 9223372036854775807, F1: p0, F2: uint8(p1)}
                _ = v8
            }
            _ = fi7
        }
        v9 := []int{-1}
        _ = v9
    }
    _ = p0
    _ = p1
    return (^(p1))
}
func genFn1(p0 uint8, p1 int16, p2 uint64) uint16 {
    v10 := []int{32767, 60}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    p0 = (^(p0))
    v11 := GenS1{F0: 42, F1: 0, F2: p0}
    p0 += (p0 >> ((123) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 500
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{105, 2147483647}
    var v3 uint32 = 4294967295
    var v4 uint16 = 2
    v5 := make(map[uint32]bool)
    var v6 float32 = 3.14159
    v12 := GenS1{F0: 42, F1: 0, F2: 236}
    var v13 bool = (genFn0(-50, uint(v4)) == genFn0(0, 4294967295))
    switch "foo bar" {
    case "\n\t":
        v4 = (v4 ^ 32767)
    default:
        v3 = (v3 * (v3 >> ((18446744073709551615) & 63)))
    }
    _ = len(v2)
    switch (("hello" + ("hello" + "\n\t")) + "\n\t") {
    case "hello":
        var v14 string = ("\n\t" + "foo bar")
        _ = v14
    case "\n\t":
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 string = "a"
                if v13 {
                    v13 = (((v3 - (v3 * v3)) <= (v3 % ((879) | 1))) && v13)
                    v4 += (v4 - 672)
                } else {
                    switch v17 {
                    case "foo bar":
                        v4 -= 2
                    case "":
                        v18 := GenS1{F0: 0, F1: -91, F2: 127}
                        _ = v18
                    }
                    v17 = "hello"
                    v2 = append(v2, 2147483647)
                }
                switch (fi16 + (fi15 ^ 127)) {
                case 128:
                    v17 = "a"
                }
                _ = fi16
                _ = v17
            }
            var v19 int = fi15
            var v20 uint32 = (v3 & uint32(v4))
            _ = fi15
            _ = v19
            _ = v20
        }
    }
    if (0.1 < 3.14159) {
        v6 -= v6
    }
    if v13 {
        var v21 int32 = 2147483647
        _ = v21
    } else {
        if v13 {
            v4 += (v4 % (((v4 / ((v4) | 1))) | 1))
        }
        v13 = v13
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            v24 := GenS1{F0: -9223372036854775808, F1: int8(fi23), F2: 127}
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
