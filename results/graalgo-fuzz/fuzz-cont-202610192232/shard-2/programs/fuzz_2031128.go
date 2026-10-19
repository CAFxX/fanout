package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32, p1 int8, p2 int) int8 {
    var v6 uint = 0
    switch -9223372036854775808 {
    case 1:
        for fi7 := 0; fi7 < 11; fi7++ {
            for fi8 := 0; fi8 < 20; fi8++ {
                fmt.Println(p1, fi7, fi8)
                v9 := []int{(fi7 % (((fi8 + 127)) | 1))}
                switch p2 {
                case -1:
                    v6 = 255
                case 127:
                    switch (p2 ^ fi7) {
                    case -1:
                        v6 -= (v6 | 693)
                    case 0:
                        v9 = append(v9, 0)
                        var v10 uint64 = 1
                        _ = v10
                    }
                case 0:
                    if false {
                        var v11 int8 = (p1 % ((p1) | 1))
                        if ((!false) && false) {
                            v11 = (v11 + (p1 | (v11 - 42)))
                            p2 = 32767
                            v11 -= (v11 % ((v11) | 1))
                        }
                        _ = v11
                    }
                }
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    case 127:
        var v12 uint8 = 255
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}
func genFn1() string {
    if false {
        var v13 uint16 = 32767
        _ = v13
    }
    switch 32768 {
    case -2147483648:
        var v14 uint8 = 113
        _ = v14
    case 9223372036854775807:
        if (true && (!(!((65 == 1) || (1 < 372))))) {
            v15 := []int{}
            var v16 bool = (-1 >= 42)
            _ = v15
            _ = v16
        }
    default:
        if false {
            switch "foo bar" {
            case "\n\t":
                var v17 int8 = 1
                _ = v17
            case "hello":
                for fi18 := 0; fi18 < 13; fi18++ {
                    for fi19 := 0; fi19 < 1; fi19++ {
                        switch -37 {
                        case 32767:
                            var v20 uint8 = 255
                            _ = v20
                        case 4294967295:
                            var v21 uint8 = 128
                            _ = v21
                        }
                        var v22 int8 = genFn0(4294967295, -99, fi18)
                        v22 -= (v22 >> ((567) & 63))
                        _ = fi19
                        _ = v22
                    }
                    _ = fi18
                }
            }
        }
    }
    return "hello"
}

func main() {
    var v1 int64 = 0
    var v2 float64 = ((3.14159 - 1.5) - 1.5)
    var v3 int16 = 1000
    v4 := make(map[string]int)
    v5 := []uint32{839, 1, 852}
    fmt.Println(v2, v2, v3)
    v23 := GenS1{F0: 513, F1: 1000, F2: 0}
    var v24 uint8 = 98
    v1 = (v1 | v1)
    v2 = float64(((float32(((0.1 + 3.14159) / v2)) * 0.0) + float32((v2 + ((float64((3.14159 + 1e10)) - (v2 + v2)) * 0.0)))))
    var v25 bool = true
    for fi26 := 0; fi26 < 4; fi26++ {
        v27 := GenS2{F0: v3}
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: 2, F1: 191, F2: 297}
    v4["\n\t"] = -2147483648
    if (v25 && v25) {
        var v29 int8 = 42
        var v30 int64 = (v1 ^ v1)
        var v31 int32 = 2147483647
        _ = v29
        _ = v30
        _ = v31
    } else {
        var v32 int16 = -855
        for fi33 := 0; fi33 < 12; fi33++ {
            fmt.Println(v24, v24, v32)
            _ = fi33
        }
        _ = v32
    }
    if (990 > 18446744073709551615) {
        switch "hello" {
        case "":
            v3 = v3
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
