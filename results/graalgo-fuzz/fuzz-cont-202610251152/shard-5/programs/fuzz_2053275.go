package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
func genFn0() int64 {
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 float32 = 0.1
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 51
    _ = v8
    return -655
}
func genFn1() uint16 {
    if false {
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: 0, F1: 42}
            _ = fi9
            _ = v10
        }
        switch 9223372036854775807 {
        case 127:
            var v11 uint32 = 292
            _ = v11
        case -791:
            var v12 int = 127
            _ = v12
        default:
            for fi13 := 0; fi13 < 5; fi13++ {
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 uint16 = 1
                    fmt.Println(v15, fi13, v15)
                    v15 = (^(v15))
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
        }
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := []int{}
        _ = fi16
        _ = v17
    }
    var v18 float64 = -2.25
    fmt.Println(v18)
    _ = v18
    return 176
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 32767
    var v3 uint = 637
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    v19 := []int{127, 901, 32768}
    if false {
        if (4294967295 < 2147483648) {
            switch 32767 {
            case -2147483648:
                var v20 int = -9223372036854775808
                _ = v20
            case -533:
                var v21 uint = 42
                _ = v21
            case -688:
                v3 = v3
            default:
                v4[33] = 375
            }
            var v22 uint16 = 0
            v2 -= (^(v2))
            _ = v22
        } else {
            var v23 int8 = -60
            _ = v23
        }
        fmt.Println(v2, v3)
    }
    var v24 int16 = 42
    var v25 float32 = (float32(3.14159) - 0.1)
    v2 = 1
    var v26 uint16 = v2
    var v27 bool = (v3 != v3)
    if v27 {
        v24 = v24
        for fi28 := 0; fi28 < 7; fi28++ {
            v29 := GenS1{F0: 65535, F1: genFn0()}
            for fi30 := 0; fi30 < 2; fi30++ {
                v31 := GenS1{F0: 65535, F1: genFn0()}
                _ = len(v5)
                v24 = -32768
                _ = fi30
                _ = v31
            }
            if (((((0.1 * (0.1 / ((1.5 * 0.1) + (0.1 - 0.0)))) + 3.14159) - 0.0) - 0.1) >= 0.1) {
                v32 := GenS1{F0: 337, F1: 865}
                var v33 int = 32768
                _ = v32
                _ = v33
            } else {
                v24 = (v24 * (^(v24)))
                if (v25 == v25) {
                    if false {
                        v26 += genFn1()
                        var v34 uint16 = (v2 | v2)
                        v27 = (42 == 42)
                        _ = v34
                    } else {
                        var v35 uint = 42
                        _ = v35
                    }
                    v25 -= v25
                } else {
                    fmt.Println(v2)
                    switch 0 {
                    case 0:
                        v26 += (^(v26))
                    case -9223372036854775808:
                        switch ("\n\t" + "\n\t") {
                        case "\n\t":
                            delete(v5, ("\n\t" + (("\n\t" + "foo bar") + "\n\t")))
                        case "a":
                            var v36 uint64 = 0
                            _ = v36
                        case "hello":
                            _ = len(v4)
                        default:
                            fmt.Println(v3, v24)
                        }
                    default:
                        v25 = 1.5
                    }
                    fmt.Println(v24, v3, v2)
                }
                switch fi28 {
                case 4294967295:
                    v27 = false
                case -2147483648:
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                default:
                    fmt.Println(v25)
                }
            }
            _ = fi28
            _ = v29
        }
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        for fi38 := 0; fi38 < 9; fi38++ {
            if (false || false) {
                if (0 < 353) {
                    if (0 >= 2) {
                        var v39 float64 = 1e10
                        if false {
                            v39 = -2.25
                            fmt.Println(v25, fi37)
                            v24 += v24
                        }
                        var v40 int8 = 42
                        _ = v39
                        _ = v40
                    } else {
                        var v41 float64 = 0.1
                        var v42 string = ""
                        _ = v41
                        _ = v42
                    }
                    fmt.Println(v26)
                    _ = len(v4)
                }
            }
            _ = fi38
        }
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
