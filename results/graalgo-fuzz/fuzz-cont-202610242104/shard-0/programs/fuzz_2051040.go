package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }
func genFn0(p0 string, p1 int8) uint8 {
    p1 = (p1 + (p1 << ((2) & 63)))
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(p0, p1)
        _ = fi7
    }
    _ = p0
    _ = p1
    return 14
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int = -2147483648
    v3 := make(map[uint32]bool)
    var v4 uint64 = v1
    var v5 string = ("" + (("\n\t" + "foo bar") + "foo bar"))
    var v6 int32 = 135
    switch v5 {
    case "a":
        switch "" {
        case "foo bar":
            delete(v3, 0)
        case "":
            delete(v3, 497)
        case "a":
            if true {
                for fi8 := 0; fi8 < 3; fi8++ {
                    for fi9 := 0; fi9 < 10; fi9++ {
                        var v10 uint32 = 4294967295
                        _ = fi9
                        _ = v10
                    }
                    v6 = 1
                    v2 -= fi8
                    _ = fi8
                }
                fmt.Println(v2, v6)
            }
        }
    }
    v5 = v5
    var v11 uint16 = 22
    v12 := GenS1{F0: 921, F1: (v6 << ((2) & 63))}
    var v13 uint8 = 42
    for fi14 := 0; fi14 < 13; fi14++ {
        v5 = ""
        _ = fi14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(v5, v6)
        for fi16 := 0; fi16 < 6; fi16++ {
            if (2 >= 92) {
                v17 := GenS1{F0: 1, F1: -10}
                _ = v17
            } else {
                switch (^(fi16)) {
                case 84:
                    v18 := GenS2{F0: 42}
                    _ = v18
                case 1:
                    v11 -= 32767
                }
                if true {
                    v19 := []int{(fi16 / (((fi15 >> ((300) & 63))) | 1)), (v2 & fi16), (fi15 % (((fi16 | (fi15 * 1))) | 1))}
                    v20 := GenS1{F0: 424, F1: 1}
                    var v21 int = fi16
                    var v22 float32 = 1.5
                    v3[904] = (1 != 0)
                    _ = v19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                v23 := GenS1{F0: 75, F1: (-(v6))}
                _ = v23
            }
            var v24 uint16 = 2
            _ = fi16
            _ = v24
        }
        fmt.Println(v5)
        _ = fi15
    }
    v6 -= v6
    switch (v2 / (((v2 & 128)) | 1)) {
    case 42:
        var v25 int = (-(v2))
        _ = v25
    }
    var v26 int = 127
    switch 0 {
    case 682:
        v26 -= (v26 & v2)
    case -100:
        fmt.Println(v26)
    default:
        v13 -= v13
    }
    if true {
        v27 := GenS3{F0: 2147483647}
        for fi28 := 0; fi28 < 18; fi28++ {
            fmt.Println(fi28, v26, v2)
            _ = fi28
        }
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 bool = false
            _ = fi29
            _ = v30
        }
        _ = v27
    }
    var v31 int8 = 127
    for fi32 := 0; fi32 < 14; fi32++ {
        fmt.Println(v6, fi32)
        v33 := v3[2147483647]
        _ = fi32
        _ = v33
    }
    var v34 int8 = (v31 >> ((4294967295) & 63))
    var v35 float64 = float64((0.1 + 1e10))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
