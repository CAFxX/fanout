package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int64; F2 uint16 }
func genFn0() int8 {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 int64 = 1
        _ = fi6
        _ = v7
    }
    var v8 uint = 949
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 bool = (339 < -334)
        _ = fi9
        _ = v10
    }
    var v11 int16 = -32768
    _ = v11
    _ = v8
    return 127
}
func genFn1() uint8 {
    if false {
        var v12 float32 = 3.14159
        v12 += (v12 / (((0.0 - (1e10 * (0.0 + (3.14159 + (v12 - (float32(-2.25) - (0.1 - (v12 + v12)))))))) / ((float32(-2.25) + 1.5) - ((float32((1e10 / 0.1)) - v12) - 0.1))) / v12))
        _ = v12
    }
    v13 := []int{450}
    _ = v13
    return 159
}

func main() {
    var v1 uint8 = 120
    v2 := make(map[uint32]string)
    var v3 uint = 4294967295
    var v4 int32 = 2147483647
    var v5 float32 = 1.5
    for fi14 := 0; fi14 < 2; fi14++ {
        v1 = genFn1()
        for fi15 := 0; fi15 < 7; fi15++ {
            fmt.Println(v4, v3)
            v16 := GenS1{F0: 42, F1: 733}
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    var v17 float64 = 0.0
    v18 := []int{2147483647, 734, 130}
    v19 := GenS2{F0: ""}
    v20 := GenS1{F0: -468, F1: 32767}
    v21 := GenS2{F0: ("foo bar" + "\n\t")}
    fmt.Println(v3)
    switch 127 {
    case -473:
        if len(v18) > 0 {
            v18[0] = 9223372036854775807
        }
    case -9223372036854775808:
        if false {
            for fi22 := 0; fi22 < 1; fi22++ {
                switch (-(fi22)) {
                case 42:
                    if len(v18) > 0 {
                        v18[0] = fi22
                    }
                case 9223372036854775807:
                    if true {
                        fmt.Println(v3, v3, v3)
                        switch "a" {
                        case "":
                            v23 := GenS2{F0: "a"}
                            _ = v23
                        case "hello":
                            v3 = 809
                        }
                        switch ("foo bar" + (("\n\t" + ("a" + "")) + "\n\t")) {
                        case "foo bar":
                            var v24 uint8 = 239
                            _ = v24
                        case "\n\t":
                            v25 := GenS1{F0: v4, F1: -680}
                            _ = v25
                        default:
                            for fi26 := 0; fi26 < 12; fi26++ {
                                if len(v18) > 0 {
                                    v18[0] = (-(fi26))
                                }
                                _ = fi26
                            }
                        }
                    }
                case 32768:
                    _ = len(v18)
                default:
                    var v27 float32 = 1.5
                    _ = v27
                }
                for fi28 := 0; fi28 < 3; fi28++ {
                    v4 -= v4
                    var v29 uint = v3
                    _ = fi28
                    _ = v29
                }
                v18 = append(v18, -388)
                _ = fi22
            }
        } else {
            if false {
                v30 := GenS3{F0: 9223372036854775807, F1: 9223372036854775807, F2: 547}
                _ = v30
            }
        }
    }
    v31 := GenS2{F0: ""}
    var v32 float32 = (v5 * 1e10)
    var v33 int8 = -1
    fmt.Println(v32)
    v34 := GenS3{F0: 9223372036854775807, F1: 42, F2: 42}
    var v35 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
