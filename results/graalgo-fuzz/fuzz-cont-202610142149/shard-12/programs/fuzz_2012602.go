package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int64 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint16 }
func genFn0() int {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 float32 = (((1e10 + float32(((float64(1e10) + (3.14159 * 0.0)) - (3.14159 + 0.0)))) * 0.1) * ((((-2.25 / 0.1) - float32(0.1)) * 1e10) - 0.1))
        switch (fi7 >> ((1) & 63)) {
        case 910:
            v9 := GenS1{F0: 127, F1: 2147483647, F2: -1}
            _ = v9
        case 0:
            v10 := GenS1{F0: 0, F1: 2147483647, F2: 533}
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    if false {
        var v11 float64 = -2.25
        _ = v11
    } else {
        var v12 int = -9223372036854775808
        _ = v12
    }
    v13 := []int{-970}
    _ = v13
    return -9223372036854775808
}
func genFn1() uint {
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 bool = false
        if ((!v15) && v15) {
            for fi16 := 0; fi16 < 2; fi16++ {
                v15 = (482 <= 65535)
                _ = fi16
            }
            for fi17 := 0; fi17 < 9; fi17++ {
                if v15 {
                    v18 := []int{-1}
                    var v19 int64 = -477
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    _ = v18
                    _ = v19
                } else {
                    fmt.Println(v15)
                    v20 := []int{genFn0(), (fi14 ^ fi14)}
                    v15 = false
                    var v21 int64 = 845
                    _ = v20
                    _ = v21
                }
                switch "\n\t" {
                case "":
                    v15 = (!((!true) || (108 < 32767)))
                default:
                    switch fi17 {
                    case 9223372036854775807:
                        v15 = (false || (114 > 2))
                    case 315:
                        v22 := GenS1{F0: -1, F1: 1, F2: 42}
                        _ = v22
                    case 782:
                        v23 := GenS1{F0: 1, F1: 1, F2: 162}
                        _ = v23
                    }
                    v24 := []int{}
                    _ = v24
                }
                _ = fi17
            }
            switch ("" + "\n\t") {
            case "\n\t":
                var v25 int64 = 42
                _ = v25
            case "":
                switch "hello" {
                case "\n\t":
                    v15 = v15
                }
            }
        } else {
            switch (fi14 % ((42) | 1)) {
            case -1:
                v26 := GenS1{F0: 69, F1: 710, F2: 1}
                _ = v26
            case 42:
                v15 = v15
            }
        }
        v15 = (4294967295 >= 887)
        _ = fi14
        _ = v15
    }
    var v27 int32 = -2147483648
    _ = v27
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 207
    var v3 string = "\n\t"
    v4 := make(map[string]bool)
    v5 := make(map[uint32]bool)
    var v6 uint = 4294967295
    v3 = v3
    fmt.Println(v3)
    var v28 int = -1
    if true {
        v6 += (^(v6))
    } else {
        v3 = "hello"
    }
    v29 := GenS1{F0: -1, F1: 2147483647, F2: -9223372036854775808}
    v3 = ("a" + "")
    for fi30 := 0; fi30 < 6; fi30++ {
        for fi31 := 0; fi31 < 6; fi31++ {
            v32 := GenS1{F0: 100, F1: 42, F2: 1}
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    v5[2147483647] = (v2 > 0)
    var v33 int32 = 0
    v34 := GenS2{F0: (v33 / (((v33 << ((255) & 63))) | 1)), F1: (^(v2)), F2: 647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
