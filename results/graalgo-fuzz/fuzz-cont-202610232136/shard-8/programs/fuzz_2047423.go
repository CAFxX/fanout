package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint8 = 108
    v2 := []int{}
    v3 := []uint32{0}
    var v4 int16 = 1
    var v5 float64 = 3.14159
    switch "" {
    case "hello":
        switch 127 {
        case 32768:
            switch "hello" {
            case "\n\t":
                var v6 uint16 = 42
                _ = v6
            default:
                var v7 int32 = 93
                _ = v7
            }
        case 128:
            var v8 int = 4294967295
            _ = v8
        case -9223372036854775808:
            v1 = v1
        }
    }
    var v9 int = 127
    v3 = append(v3, 2147483647)
    var v10 bool = (102 >= -128)
    v1 = 94
    var v11 string = (("foo bar" + "\n\t") + "a")
    var v12 uint8 = v1
    if len(v2) > 0 {
        v2[0] = -885
    }
    switch (v9 / ((v9) | 1)) {
    case -456:
        var v13 bool = (1 == 857)
        _ = v13
    case 0:
        if (!v10) {
            v14 := GenS1{F0: 142}
            if v10 {
                if v10 {
                    v5 = v5
                } else {
                    v5 = (-2.25 * v5)
                    v10 = (0.0 >= (1.5 / v5))
                }
                for fi15 := 0; fi15 < 11; fi15++ {
                    v16 := GenS2{F0: "\n\t"}
                    for fi17 := 0; fi17 < 13; fi17++ {
                        if len(v2) > 0 {
                            v2[0] = (-(fi15))
                        }
                        _ = fi17
                    }
                    _ = fi15
                    _ = v16
                }
                v1 -= v1
            }
            _ = v14
        }
    case 127:
        switch "hello" {
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
        case "":
            v18 := GenS2{F0: (v11 + (v11 + ("a" + (v11 + v11))))}
            _ = v18
        }
    }
    v12 = (v12 + 42)
    for fi19 := 0; fi19 < 14; fi19++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 int64 = 0
            v22 := GenS2{F0: v11}
            var v23 int = (fi20 / (((v9 / (((^(fi19))) | 1))) | 1))
            v21 += (v21 >> ((42) & 63))
            v23 += int(v21)
            v21 += (-(v21))
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
