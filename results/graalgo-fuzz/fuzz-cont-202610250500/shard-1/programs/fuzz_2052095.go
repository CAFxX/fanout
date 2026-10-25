package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint }
type GenS2 struct { F0 uint32; F1 uint16; F2 int8 }
func genFn0(p0 uint) uint {
    if false {
        p0 = p0
    } else {
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 int64 = 0
            _ = fi6
            _ = v7
        }
    }
    v8 := GenS1{F0: 92, F1: 1, F2: (p0 / (((p0 % ((p0) | 1))) | 1))}
    _ = p0
    _ = v8
    return (p0 / (((^(p0))) | 1))
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (!(!(!(-740 == 42))))
    v3 := make(map[uint32]int)
    var v4 uint16 = 32767
    var v5 uint8 = 255
    v9 := []int{2147483647, 4294967295}
    for fi10 := 0; fi10 < 11; fi10++ {
        v2 = v2
        _ = fi10
    }
    fmt.Println(v5, v2, v2)
    v3[65535] = 128
    for fi11 := 0; fi11 < 14; fi11++ {
        if v2 {
            v2 = false
            switch "hello" {
            case "\n\t":
                if true {
                    var v12 uint16 = v4
                    v2 = true
                    _ = v12
                } else {
                    var v13 int32 = 42
                    if v2 {
                        fmt.Println(v4, v5, v4)
                        fmt.Println(v5)
                    } else {
                        for fi14 := 0; fi14 < 20; fi14++ {
                            v15 := GenS1{F0: (v5 / (((^(v5))) | 1)), F1: 4294967295, F2: 255}
                            if len(v9) > 0 {
                                v9[0] = fi11
                            }
                            v4 = (v4 - (v4 / ((v4) | 1)))
                            var v16 string = ""
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        var v17 uint16 = v4
                        _ = v17
                    }
                    _ = v13
                }
            case "":
                v18 := v3[2147483647]
                _ = v18
            case "hello":
                v4 = 42
            default:
                var v19 int16 = -32768
                _ = v19
            }
            var v20 string = ("a" + "foo bar")
            _ = v20
        }
        if len(v9) > 0 {
            v9[0] = (fi11 >> ((896) & 63))
        }
        for fi21 := 0; fi21 < 14; fi21++ {
            v4 = (v4 / ((42) | 1))
            fmt.Println(v4)
            v4 = 620
            _ = fi21
        }
        _ = fi11
    }
    v5 = 2
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 6; fi23++ {
            v4 += 1
            v24 := GenS1{F0: v5, F1: 2, F2: 18446744073709551615}
            fmt.Println(fi23, v5, fi22)
            var v25 int = fi22
            if false {
                switch fi23 {
                case 2147483647:
                    var v26 int16 = 1
                    _ = v26
                default:
                    v2 = ((1 > 4294967295) || false)
                }
                v5 = v5
            } else {
                var v27 float32 = float32(((((1e10 - (1.5 + (0.0 * 1.5))) - (1e10 + -2.25)) - 0.0) - 3.14159))
                _ = v27
            }
            _ = fi23
            _ = v24
            _ = v25
        }
        if v2 {
            v28 := GenS2{F0: 481, F1: (v4 >> ((2) & 63)), F2: 19}
            _ = v28
        }
        _ = fi22
    }
    var v29 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
