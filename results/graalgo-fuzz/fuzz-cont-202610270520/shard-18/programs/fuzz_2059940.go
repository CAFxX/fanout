package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float32 = 0.1
    var v2 uint16 = 2
    var v3 uint64 = 4294967295
    var v4 float32 = v1
    var v5 uint64 = (v3 % ((v3) | 1))
    var v6 int16 = -930
    v7 := []int{32768}
    if (true || (127 > 42)) {
        v4 -= 1.5
        for fi8 := 0; fi8 < 6; fi8++ {
            fmt.Println(v4)
            v1 -= v4
            _ = fi8
        }
        switch -1 {
        case 127:
            if (65535 >= 1) {
                v5 += (v3 | v3)
            }
        case -9223372036854775808:
            if (180 > 208) {
                v3 = 42
                if false {
                    v9 := GenS1{F0: 62}
                    v4 += v4
                    _ = v9
                }
                _ = len(v7)
            } else {
                if len(v7) > 0 {
                    v7[0] = -9223372036854775808
                }
                if false {
                    if false {
                        v5 += v3
                        var v10 uint32 = 42
                        for fi11 := 0; fi11 < 11; fi11++ {
                            var v12 int16 = v6
                            v2 = (v2 & 1)
                            for fi13 := 0; fi13 < 7; fi13++ {
                                var v14 bool = false
                                var v15 float64 = (1e10 / 1e10)
                                v15 -= ((((v15 + (1e10 + v15)) / v15) - (3.14159 - v15)) / v15)
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            _ = fi11
                            _ = v12
                        }
                        _ = v10
                    }
                    var v16 string = ("foo bar" + "")
                    _ = v16
                }
                _ = len(v7)
            }
        }
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    if false {
        v6 += 550
        var v17 uint16 = (^(v2))
        switch 429 {
        case 128:
            v2 -= v17
        case 9223372036854775807:
            fmt.Println(v4)
        default:
            var v18 float32 = v4
            _ = v18
        }
        _ = v17
    }
    v19 := GenS1{F0: 42}
    for fi20 := 0; fi20 < 17; fi20++ {
        switch 32767 {
        case -2147483648:
            v1 = 0.1
        case -1:
            switch 79 {
            case 1:
                v21 := GenS1{F0: (^(fi20))}
                _ = v21
            }
        case 479:
            v6 = v6
        }
        _ = fi20
    }
    fmt.Println(v6, v3)
    if true {
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 bool = false
            _ = fi22
            _ = v23
        }
        v5 -= v5
        var v24 int = 378
        _ = v24
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
