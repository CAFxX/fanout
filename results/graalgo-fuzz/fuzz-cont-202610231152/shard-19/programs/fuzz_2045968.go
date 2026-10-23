package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
func genFn0(p0 int16, p1 uint64, p2 int) int32 {
    v5 := []int{}
    fmt.Println(p2)
    if len(v5) > 0 {
        v5[0] = p2
    }
    if true {
        for fi6 := 0; fi6 < 13; fi6++ {
            p1 += (p1 >> ((15) & 63))
            _ = fi6
        }
    } else {
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                p2 += (p2 % (((^(p2))) | 1))
                var v9 int8 = 90
                _ = fi8
                _ = v9
            }
            p1 -= p1
            _ = fi7
        }
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 bool) int64 {
    p0 = (genFn0(-1, 1, 2147483647) < 254)
    v10 := []int{-2147483648, 128, -2147483648}
    _ = p0
    _ = v10
    return -58
}

func main() {
    var v1 float64 = -2.25
    var v2 float32 = (0.0 / 3.14159)
    var v3 int16 = 32767
    var v4 int = -217
    switch v4 {
    case 42:
        v11 := GenS1{F0: -1, F1: 0}
        _ = v11
    case 127:
        if (!true) {
            if false {
                v1 -= v1
                var v12 bool = (67 >= 2)
                _ = v12
            }
            var v13 int16 = (v3 % ((v3) | 1))
            switch "a" {
            case "foo bar":
                for fi14 := 0; fi14 < 20; fi14++ {
                    fmt.Println(v13, v3, v4)
                    _ = fi14
                }
            }
            _ = v13
        }
    }
    switch "foo bar" {
    case "\n\t":
        switch (v4 ^ 57) {
        case 964:
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 int32 = 0
                v17 := []int{(fi15 >> ((16) & 63))}
                v4 = int(v3)
                _ = fi15
                _ = v16
                _ = v17
            }
        default:
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                switch "\n\t" {
                case "foo bar":
                    v4 += v4
                }
            case "hello":
                for fi18 := 0; fi18 < 8; fi18++ {
                    fmt.Println(fi18, v4)
                    _ = fi18
                }
            case "":
                v19 := GenS1{F0: 1, F1: 255}
                _ = v19
            }
        }
    case "a":
        if true {
            if (552 > 2147483647) {
                var v20 uint8 = 242
                _ = v20
            }
        } else {
            if false {
                v3 = v3
                for fi21 := 0; fi21 < 13; fi21++ {
                    var v22 uint = 1
                    v23 := GenS1{F0: int64(v3), F1: 127}
                    fmt.Println(v22, v1, v4)
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
        }
    default:
        for fi24 := 0; fi24 < 12; fi24++ {
            switch (-(fi24)) {
            case 0:
                for fi25 := 0; fi25 < 10; fi25++ {
                    v1 -= 3.14159
                    var v26 uint = 566
                    var v27 int16 = (v3 + 1)
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
            _ = fi24
        }
    }
    v3 += v3
    var v28 int8 = 100
    if (18446744073709551615 < 551) {
        var v29 int8 = -85
        v28 = 7
        _ = v29
    } else {
        var v30 bool = (!true)
        var v31 float32 = (((v2 / v2) - v2) - v2)
        _ = v30
        _ = v31
    }
    var v32 int = (v4 << ((904) & 63))
    var v33 int = v32
    switch "\n\t" {
    case "foo bar":
        v4 = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
