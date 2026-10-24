package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 string; F1 int8 }
func genFn0(p0 bool) uint32 {
    var v5 int = 42
    v6 := []int{(v5 ^ 840), (v5 - (v5 - v5)), 1}
    var v7 string = "foo bar"
    if p0 {
        var v8 int8 = 127
        _ = v8
    } else {
        if (p0 || (1 >= -2147483648)) {
            switch v5 {
            case -9223372036854775808:
                if ("" > ("foo bar" + v7)) {
                    fmt.Println(p0)
                    var v9 uint = 65535
                    _ = v9
                }
            case -2147483648:
                if (1.5 >= ((-2.25 + 3.14159) / 1.0)) {
                    p0 = p0
                    for fi10 := 0; fi10 < 15; fi10++ {
                        v6 = append(v6, fi10)
                        var v11 float32 = 3.14159
                        v11 += 1.5
                        _ = fi10
                        _ = v11
                    }
                    if p0 {
                        v12 := GenS1{F0: v5, F1: 2}
                        _ = v12
                    } else {
                        v7 = "hello"
                    }
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v13 int16 = -32768
                    _ = v13
                }
            case 128:
                if p0 {
                    var v14 float32 = ((0.0 * (1.5 * (((((((0.0 - 3.14159) / 3.14159) * (3.14159 + 0.1)) / 1.0) + (((0.1 - (((1e10 * (((((0.0 / 1e10) / 1.0) / ((1e10 / -2.25) / 1.0)) * (-2.25 / 1.0)) * 1.5)) + 0.0) + 1e10)) / (3.14159 * (float32(((((((1.5 + ((1e10 * (3.14159 - 1e10)) / -2.25)) - ((0.1 * 0.0) / (0.0 - 1.5))) * (0.1 + 1e10)) - 1.5) / 3.14159) * (((float64((float32(float64(((1.5 + 1e10) / 1.5))) / 0.1)) / 3.14159) * (1.5 + 0.0)) / ((0.1 + 1.5) / 1e10)))) / float32((1.5 + 1.5))))) / (3.14159 - 0.1))) - float32(0.0)) + 1.5))) - 0.1)
                    for fi15 := 0; fi15 < 19; fi15++ {
                        v16 := GenS1{F0: v5, F1: 51}
                        _ = fi15
                        _ = v16
                    }
                    fmt.Println(v14, v14)
                    _ = v14
                }
            }
            var v17 bool = p0
            _ = v17
        }
        fmt.Println(v7, p0, v7)
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int, p1 bool) uint32 {
    v18 := []int{p0}
    if len(v18) > 0 {
        v18[0] = (-(p0))
    }
    p0 = 128
    _ = p0
    _ = p1
    _ = v18
    return 946
}

func main() {
    v1 := []int{32767, 10, 128}
    var v2 int8 = -63
    v3 := make(map[uint32]string)
    var v4 int8 = v2
    if false {
        if false {
            if false {
                for fi19 := 0; fi19 < 9; fi19++ {
                    fmt.Println(v4)
                    _ = fi19
                }
                switch "foo bar" {
                case "a":
                    var v20 uint64 = 642
                    _ = v20
                case "\n\t":
                    v3[739] = "foo bar"
                }
            } else {
                v2 = v4
                v4 -= v2
            }
        } else {
            var v21 bool = ("" == "foo bar")
            var v22 int16 = 32767
            _ = len(v3)
            _ = v21
            _ = v22
        }
        var v23 bool = false
        v4 = v4
        _ = v23
    }
    v2 -= v2
    v24 := GenS1{F0: -854, F1: 127}
    v4 = 25
    var v25 uint8 = 0
    v26 := GenS2{F0: "foo bar", F1: (v2 ^ v4)}
    var v27 int8 = v2
    for fi28 := 0; fi28 < 2; fi28++ {
        for fi29 := 0; fi29 < 16; fi29++ {
            v30 := GenS1{F0: 85, F1: 224}
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    var v31 uint16 = 65535
    var v32 string = "a"
    fmt.Println(v27, v25)
    var v33 uint16 = v31
    switch v32 {
    case "":
        var v34 uint16 = 121
        _ = v34
    case "foo bar":
        v32 = (v32 + ((v32 + "foo bar") + (("" + "a") + v32)))
    default:
        switch 42 {
        case -751:
            var v35 uint = 65535
            _ = v35
        case 4294967295:
            switch -415 {
            case 9223372036854775807:
                v36 := GenS2{F0: "hello", F1: 100}
                _ = v36
            default:
                v31 -= v31
            }
        default:
            v31 = uint16(v25)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
