package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0() int8 {
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        switch 127 {
        case 638:
            if true {
                v7 := []int{2147483647, 32768}
                for fi8 := 0; fi8 < 17; fi8++ {
                    switch "hello" {
                    case "a":
                        v7 = append(v7, (fi8 << ((42) & 63)))
                    case "\n\t":
                        var v9 int8 = 1
                        _ = v9
                    }
                    _ = fi8
                }
                _ = v7
            }
        case 32768:
            var v10 uint32 = 42
            _ = v10
        }
    case "a":
        v11 := []int{}
        _ = v11
    }
    var v12 float64 = (float64(3.14159) + 3.14159)
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int64 = -9223372036854775808
        var v15 float64 = v12
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 int = fi16
            v18 := GenS1{F0: 42, F1: "a"}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v12
    return -126
}

func main() {
    v1 := []int{-817, 0, -9223372036854775808}
    var v2 string = "\n\t"
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    var v5 int8 = 96
    v6 := []uint32{0, 1, 974, 2}
    if false {
        fmt.Println(v2, v2)
    }
    var v19 uint64 = 2
    for fi20 := 0; fi20 < 19; fi20++ {
        switch ("a" + ("\n\t" + v2)) {
        case "hello":
            v19 -= 42
        }
        _ = fi20
    }
    var v21 uint64 = v19
    if (true && ((!false) && false)) {
        v6 = append(v6, 2147483647)
    }
    var v22 int32 = 84
    fmt.Println(v2, v5, v2)
    if false {
        fmt.Println(v22, v5, v21)
        v23 := GenS1{F0: 2, F1: "foo bar"}
        _ = v23
    } else {
        for fi24 := 0; fi24 < 6; fi24++ {
            for fi25 := 0; fi25 < 9; fi25++ {
                switch v2 {
                case "foo bar":
                    var v26 bool = false
                    _ = v26
                default:
                    fmt.Println(fi24, fi24)
                }
                _ = fi25
            }
            fmt.Println(v2, fi24)
            v2 = v2
            _ = fi24
        }
        v27 := GenS1{F0: uint16(v19), F1: (("hello" + v2) + ("\n\t" + ("hello" + (v2 + ""))))}
        _ = v27
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v21 -= v21
    if ((1.5 < ((((-2.25 + 0.1) - 0.0) + (1.5 / ((-2.25 - float64((((float32((1e10 - 0.0)) * 0.1) / 3.14159) + ((-2.25 - (3.14159 / 0.1)) / 1e10)))) * (1e10 * 3.14159)))) - 0.0)) && (false && (1 == 1))) {
        v28 := v4[0]
        _ = v28
    } else {
        v2 = v2
        if (100 != (-(v5))) {
            for fi29 := 0; fi29 < 5; fi29++ {
                switch "\n\t" {
                case "":
                    fmt.Println(v21, v21, v5)
                case "a":
                    var v30 string = v2
                    _ = v30
                case "\n\t":
                    v31 := GenS1{F0: 65535, F1: v2}
                    _ = v31
                }
                v32 := v3[4294967295]
                _ = fi29
                _ = v32
            }
            var v33 float32 = (0.0 / 1e10)
            var v34 int8 = v5
            _ = v33
            _ = v34
        } else {
            v2 = "hello"
            var v35 int16 = -676
            _ = v35
        }
        for fi36 := 0; fi36 < 12; fi36++ {
            _ = len(v4)
            _ = fi36
        }
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        if false {
            if (1e10 == 1e10) {
                v21 -= 42
            } else {
                var v38 uint16 = 32767
                _ = v38
            }
            v39 := GenS1{F0: 32767, F1: v2}
            _ = v39
        }
        _ = fi37
    }
    var v40 int32 = (v22 | (v22 / ((-2147483648) | 1)))
    var v41 uint = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
