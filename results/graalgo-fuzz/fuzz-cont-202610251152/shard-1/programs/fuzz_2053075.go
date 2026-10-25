package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 string; F1 string; F2 uint }
type GenS4 struct { F0 uint8; F1 uint }

func main() {
    var v1 uint64 = 0
    var v2 bool = false
    var v3 uint64 = (^(v1))
    var v4 string = ("" + "foo bar")
    v4 = ((((v4 + ("\n\t" + v4)) + "foo bar") + v4) + v4)
    v3 = 2
    switch -2147483648 {
    case -9223372036854775808:
        var v5 int64 = 42
        _ = v5
    default:
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := GenS1{F0: (^(v3))}
            v8 := []int{fi6}
            fmt.Println(v2, v1)
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    v9 := []int{32768, 127}
    var v10 float32 = 1e10
    fmt.Println(v10)
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS1{F0: v1}
        for fi13 := 0; fi13 < 5; fi13++ {
            switch 127 {
            case 4294967295:
                fmt.Println(v1)
            case 127:
                var v14 uint32 = 65535
                _ = v14
            }
            fmt.Println(v2, v1)
            switch (-(fi11)) {
            case 9223372036854775807:
                var v15 uint32 = 2
                _ = v15
            }
            var v16 int8 = 51
            v17 := GenS1{F0: (v3 << ((4294967295) & 63))}
            v18 := GenS2{F0: (v4 < v4), F1: (1 == 127)}
            _ = fi13
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi11
        _ = v12
    }
    var v19 int32 = -747
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 int32 = 961
        _ = fi20
        _ = v21
    }
    if v2 {
        switch ("hello" + v4) {
        case "hello":
            v22 := GenS3{F0: v4, F1: "", F2: 1}
            _ = v22
        case "a":
            var v23 int32 = (v19 - (v19 % ((v19) | 1)))
            _ = v23
        }
        if ((!(!(32767 >= 32767))) || (v2 && (((((((float64(1.5) - ((1e10 + ((1e10 / -2.25) + (((((-2.25 * float64((v10 + float32(0.1)))) / (0.1 / float64(v10))) / float64(0.1)) - 0.0) + float64(((((0.0 + ((1e10 * v10) * v10)) + v10) * (v10 / (v10 + v10))) * float32(3.14159)))))) - (1.5 / -2.25))) + ((3.14159 / (3.14159 + (((0.1 * (0.1 * (1.5 / (1e10 - (0.0 * 3.14159))))) / (((1.5 / (0.1 + ((float64((v10 * v10)) / 1.0) + ((1e10 - 0.1) / 1.5)))) / -2.25) - float64(v10))) / float64(1e10)))) + 0.0)) + 3.14159) + float64((v10 + 0.1))) * 0.1) - 3.14159) > 1e10))) {
            v2 = (!(0.1 >= 1e10))
        }
        var v24 float32 = (1e10 * v10)
        _ = v24
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    if false {
        var v25 uint16 = 65535
        if ((-46 != -167) || (127 != -1)) {
            v19 -= (^(v19))
            for fi26 := 0; fi26 < 16; fi26++ {
                v25 = 42
                if (true && true) {
                    var v27 float64 = float64(0.0)
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 12; fi28++ {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        v10 = 3.14159
                        _ = fi28
                    }
                    if len(v9) > 0 {
                        v9[0] = fi26
                    }
                    v19 = 2147483647
                }
                if len(v9) > 0 {
                    v9[0] = fi26
                }
                _ = fi26
            }
        } else {
            if v2 {
                fmt.Println(v3)
                v4 = v4
                v4 = (v4 + v4)
            } else {
                v10 = -2.25
                for fi29 := 0; fi29 < 8; fi29++ {
                    var v30 int8 = 106
                    _ = fi29
                    _ = v30
                }
                if v2 {
                    var v31 string = ""
                    _ = v31
                }
            }
            fmt.Println(v19)
            if len(v9) > 0 {
                v9[0] = 128
            }
        }
        _ = v25
    } else {
        v32 := GenS4{F0: 76, F1: 42}
        if len(v9) > 0 {
            v9[0] = -9223372036854775808
        }
        _ = v32
    }
    v2 = v2
    for fi33 := 0; fi33 < 9; fi33++ {
        if len(v9) > 0 {
            v9[0] = 32768
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
