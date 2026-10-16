package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := []int{4294967295, 815}
    var v4 float64 = 0.1
    var v5 uint64 = 18446744073709551615
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint64 = 42
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: int16(fi9), F1: -1}
        _ = fi9
        _ = v10
    }
    v5 = v5
    var v11 uint64 = v5
    if true {
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := v2["hello"]
            _ = len(v2)
            fmt.Println(v4)
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: 1, F1: -75}
        if false {
            for fi15 := 0; fi15 < 13; fi15++ {
                fmt.Println(fi15)
                var v16 int16 = 679
                _ = fi15
                _ = v16
            }
            if true {
                switch "\n\t" {
                case "foo bar":
                    v5 -= (^(v5))
                case "hello":
                    v3 = append(v3, 814)
                case "\n\t":
                    switch 42 {
                    case -950:
                        var v17 int32 = 2147483647
                        _ = v17
                    case 42:
                        var v18 uint32 = 42
                        _ = v18
                    }
                default:
                    v4 -= (1e10 + ((float64((1e10 / (((1e10 - 0.1) / 0.1) / -2.25))) - v4) * 3.14159))
                }
                if len(v3) > 0 {
                    v3[0] = 0
                }
                for fi19 := 0; fi19 < 5; fi19++ {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        _ = len(v3)
                        if (true && false) {
                            var v21 int = fi19
                            _ = v21
                        } else {
                            v22 := GenS2{F0: "", F1: -1}
                            if false {
                                v23 := v2[("\n\t" + (("a" + ("hello" + "a")) + ("a" + "")))]
                                var v24 int8 = 100
                                _ = v23
                                _ = v24
                            }
                            _ = v22
                        }
                        var v25 uint32 = 4294967295
                        var v26 string = ""
                        var v27 int = (fi20 & 990)
                        var v28 int32 = -629
                        _ = fi20
                        _ = v25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    var v29 int16 = 42
                    _ = fi19
                    _ = v29
                }
            }
            if true {
                fmt.Println(v4, v5)
                var v30 int64 = -1
                for fi31 := 0; fi31 < 3; fi31++ {
                    v11 = v5
                    var v32 int32 = 267
                    _ = fi31
                    _ = v32
                }
                _ = v30
            }
        }
        _ = v14
    } else {
        v5 -= (v5 >> ((4294967295) & 63))
        for fi33 := 0; fi33 < 2; fi33++ {
            v4 = v4
            if (847 > -207) {
                if (!(false && true)) {
                    switch "foo bar" {
                    case "":
                        var v34 float64 = ((v4 + 3.14159) / v4)
                        _ = v34
                    case "foo bar":
                        var v35 uint16 = 255
                        _ = v35
                    case "hello":
                        _ = len(v6)
                    default:
                        var v36 uint8 = 127
                        _ = v36
                    }
                } else {
                    v1[("\n\t" + "hello")] = "\n\t"
                    var v37 float64 = (((((float64(((0.1 - 1.5) + ((0.0 * 1e10) + 0.1))) - (((v4 / (v4 / v4)) + v4) - 1.5)) + v4) / ((v4 + v4) / float64((3.14159 * float32(1e10))))) / (v4 - v4)) + ((0.0 + (0.1 + (v4 * v4))) / 1.0))
                    _ = v37
                }
            } else {
                var v38 bool = false
                var v39 bool = v38
                _ = v38
                _ = v39
            }
            _ = fi33
        }
        var v40 uint16 = 699
        _ = v40
    }
    var v41 string = ("\n\t" + ("hello" + "a"))
    if len(v3) > 0 {
        v3[0] = 128
    }
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
