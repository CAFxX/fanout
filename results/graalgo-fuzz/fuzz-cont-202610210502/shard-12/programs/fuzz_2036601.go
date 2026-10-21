package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }
func genFn0() bool {
    var v6 uint8 = 1
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 92, F1: "hello", F2: 100}
        v6 = 36
        _ = fi7
        _ = v8
    }
    var v9 uint = 4294967295
    var v10 bool = ((true && ((float64(1e10) / -2.25) >= 0.0)) && ((((647 > 858) || true) || false) || (4294967295 != v9)))
    _ = v10
    _ = v6
    _ = v9
    return v10
}
func genFn1() uint16 {
    var v11 int = 9223372036854775807
    v11 += v11
    v12 := []int{42, v11, 42, v11}
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := []uint32{987, 2, 4294967295}
    var v2 int64 = 0
    var v3 int = 889
    var v4 int32 = -275
    var v5 uint64 = 1
    fmt.Println(v3, v4)
    if (217 <= 42) {
        for fi13 := 0; fi13 < 8; fi13++ {
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := GenS1{F0: -205, F1: ((("a" + "a") + ("hello" + "\n\t")) + ""), F2: 1}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi14
                _ = v15
            }
            fmt.Println(v2, v2)
            var v16 int = -931
            _ = fi13
            _ = v16
        }
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 float32 = ((-2.25 - 3.14159) * 3.14159)
            fmt.Println(fi17, v4)
            var v19 float32 = (v18 + (v18 * -2.25))
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            fmt.Println(v4, v2, fi20)
            _ = fi20
        }
    } else {
        for fi21 := 0; fi21 < 18; fi21++ {
            fmt.Println(v2, v4)
            _ = fi21
        }
        var v22 int16 = 32767
        switch -773 {
        case 0:
            v2 -= (v2 + v2)
        case 42:
            if false {
                for fi23 := 0; fi23 < 7; fi23++ {
                    v5 = (^(v5))
                    _ = fi23
                }
                fmt.Println(v3)
                v24 := GenS1{F0: v4, F1: "\n\t", F2: -1}
                _ = v24
            } else {
                switch "\n\t" {
                case "":
                    if true {
                        v2 = (v2 % (((^(v2))) | 1))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v2 += v2
                    }
                case "hello":
                    if (uint(v5) <= 2) {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            switch "" {
                            case "foo bar":
                                v5 = 2
                            }
                            var v26 uint8 = 127
                            v1 = append(v1, 42)
                            _ = fi25
                            _ = v26
                        }
                    }
                case "foo bar":
                    v22 -= 32767
                default:
                    var v27 int = v3
                    _ = v27
                }
            }
        case 432:
            fmt.Println(v3)
        default:
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 int8 = -58
                v30 := GenS1{F0: 144, F1: "\n\t", F2: v29}
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        _ = v22
    }
    var v31 uint = 42
    var v32 int = (v3 & v3)
    fmt.Println(v32)
    fmt.Println(v2)
    var v33 int32 = (v4 - 366)
    var v34 uint32 = 2147483647
    var v35 string = "hello"
    var v36 int8 = 42
    var v37 uint = (v31 - v31)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v4, v31, v36)
    }
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    v3 -= (v32 / (((^(v3))) | 1))
    v31 = (^(v37))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
