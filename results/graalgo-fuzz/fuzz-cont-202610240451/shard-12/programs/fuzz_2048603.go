package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 bool; F1 int8; F2 uint8 }
type GenS4 struct { F0 int; F1 uint16; F2 uint }
type GenS5 struct { F0 int32; F1 uint }
func genFn0() bool {
    var v5 bool = (false || (3.14159 >= 1.5))
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 18446744073709551615, F1: (^(fi6)), F2: 127}
        v5 = (!false)
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: 1, F1: 18446744073709551615}
    v9 := []int{2147483647, 9223372036854775807, 4294967295, -9223372036854775808}
    v5 = false
    _ = v5
    _ = v8
    _ = v9
    return v5
}

func main() {
    var v1 uint = 255
    v2 := []string{("hello" + ("hello" + "foo bar")), ("" + ("hello" + "a")), "", ""}
    var v3 uint8 = 255
    var v4 float32 = 1.5
    var v10 uint64 = 42
    switch ("" + "") {
    case "":
        v2 = append(v2, "\n\t")
    case "hello":
        var v11 float32 = v4
        _ = v11
    case "a":
        if false {
            v3 = (v3 ^ (v3 / ((2) | 1)))
            if true {
                if true {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        for fi13 := 0; fi13 < 4; fi13++ {
                            if false {
                                var v14 uint = v1
                                if (true && (!false)) {
                                    var v15 int = (fi12 % ((fi13) | 1))
                                    _ = v15
                                }
                                _ = v14
                            } else {
                                v3 += 127
                                var v16 string = "foo bar"
                                _ = v16
                            }
                            _ = fi13
                        }
                        _ = fi12
                    }
                }
                for fi17 := 0; fi17 < 12; fi17++ {
                    fmt.Println(v3, v3, v10)
                    v18 := GenS1{F0: (v10 << ((v1) & 63)), F1: (fi17 * fi17), F2: 0}
                    for fi19 := 0; fi19 < 15; fi19++ {
                        if false {
                            var v20 uint16 = 32767
                            var v21 bool = false
                            _ = v20
                            _ = v21
                        }
                        _ = fi19
                    }
                    _ = fi17
                    _ = v18
                }
                switch -937 {
                case 2147483647:
                    v1 = v1
                default:
                    v22 := GenS3{F0: true, F1: -128, F2: (^(v3))}
                    _ = v22
                }
            } else {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                for fi23 := 0; fi23 < 1; fi23++ {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v10 = (v10 * 0)
                        var v25 uint8 = uint8(v10)
                        _ = fi24
                        _ = v25
                    }
                    _ = fi23
                }
            }
            v10 = v10
        } else {
            var v26 uint32 = 249
            for fi27 := 0; fi27 < 3; fi27++ {
                switch (fi27 << (((v1 / ((1) | 1))) & 63)) {
                case -451:
                    switch (-(fi27)) {
                    case -245:
                        if true {
                            var v28 int16 = 32767
                            var v29 float64 = (3.14159 - (1e10 / 3.14159))
                            var v30 int64 = 1
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            v2 = append(v2, "\n\t")
                            v31 := GenS2{F0: (v26 << (((^(v1))) & 63)), F1: v10}
                            _ = v31
                        }
                    default:
                        if (v26 >= v26) {
                            var v32 string = (("foo bar" + ("" + (((("foo bar" + "foo bar") + "foo bar") + (("a" + "foo bar") + "a")) + "foo bar"))) + "foo bar")
                            v3 = 1
                            switch fi27 {
                            case 471:
                                var v33 int = (fi27 * fi27)
                                _ = v33
                            default:
                                var v34 uint64 = v10
                                _ = v34
                            }
                            _ = v32
                        }
                    }
                }
                _ = fi27
            }
            fmt.Println(v3)
            _ = v26
        }
    }
    _ = len(v2)
    switch 143 {
    case -1:
        v35 := GenS4{F0: 32768, F1: 2, F2: 42}
        _ = v35
    }
    var v36 float32 = v4
    switch 42 {
    case -59:
        var v37 int8 = 7
        _ = v37
    case -2147483648:
        for fi38 := 0; fi38 < 14; fi38++ {
            if (0.1 <= 3.14159) {
                for fi39 := 0; fi39 < 11; fi39++ {
                    v10 -= (v10 * (v10 ^ (v10 << ((v1) & 63))))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch fi39 {
                    case 4294967295:
                        fmt.Println(v4, v3)
                    case 659:
                        fmt.Println(fi39, fi38, fi38)
                    case 32768:
                        v40 := GenS5{F0: -2147483648, F1: (v1 << ((908) & 63))}
                        _ = v40
                    default:
                        var v41 float64 = 1.5
                        _ = v41
                    }
                    var v42 float32 = (((v4 / (v4 * 1.5)) - 0.1) / (((3.14159 + -2.25) + float32(1.5)) - v4))
                    if false {
                        var v43 float64 = 1.5
                        _ = v43
                    }
                    _ = fi39
                    _ = v42
                }
                fmt.Println(v4, v3, v4)
                if false {
                    v44 := GenS5{F0: -1000, F1: v1}
                    _ = v44
                }
            }
            _ = fi38
        }
    case 2147483647:
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    default:
        fmt.Println(v4, v4, v4)
    }
    var v45 int32 = 861
    if true {
        var v46 string = "foo bar"
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v45
}
