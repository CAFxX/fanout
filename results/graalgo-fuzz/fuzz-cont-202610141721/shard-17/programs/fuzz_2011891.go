package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint8 }
func genFn0() uint16 {
    if (false || (1 > 60)) {
        var v4 uint = 0
        fmt.Println(v4)
        if true {
            var v5 int32 = 2147483647
            if (false && false) {
                v6 := []int{127, -9223372036854775808}
                _ = v6
            } else {
                v4 = (v4 * (v4 >> (((v4 ^ v4)) & 63)))
                switch -9223372036854775808 {
                case 128:
                    v5 += v5
                }
                for fi7 := 0; fi7 < 9; fi7++ {
                    v5 += v5
                    var v8 int32 = 1
                    v9 := GenS1{F0: "", F1: 2147483648}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
            _ = v5
        }
        _ = v4
    }
    switch "" {
    case "hello":
        v10 := []int{}
        _ = v10
    case "":
        var v11 uint = 148
        _ = v11
    default:
        for fi12 := 0; fi12 < 5; fi12++ {
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := []int{fi12, fi12, fi12, (fi13 << ((2) & 63))}
                var v15 uint8 = 255
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 int32 = 1
            _ = fi12
            _ = v16
        }
    }
    return 42
}

func main() {
    var v1 uint8 = 153
    var v2 bool = false
    v3 := []int{}
    v3 = append(v3, -1)
    v17 := GenS1{F0: ("" + (("\n\t" + "hello") + "hello")), F1: 2147483647}
    var v18 uint16 = 0
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 uint32 = 1
        _ = fi19
        _ = v20
    }
    var v21 int64 = -908
    var v22 float32 = -2.25
    v18 += (^(v18))
    if (false && ((1 > 18446744073709551615) || ("" >= ("\n\t" + (("" + "foo bar") + "\n\t"))))) {
        fmt.Println(v1, v21)
        v21 -= int64(v18)
    } else {
        if false {
            v23 := GenS1{F0: "\n\t", F1: 2147483648}
            for fi24 := 0; fi24 < 5; fi24++ {
                switch "a" {
                case "":
                    var v25 bool = (-1 >= fi24)
                    _ = v25
                case "\n\t":
                    v22 += v22
                default:
                    fmt.Println(v18, v21, v2)
                }
                v21 += v21
                v18 = v18
                _ = fi24
            }
            v3 = append(v3, 153)
            _ = v23
        }
    }
    if ((2147483648 <= 0) && v2) {
        var v26 string = "a"
        switch v26 {
        case "a":
            var v27 int8 = 127
            _ = v27
        case "\n\t":
            for fi28 := 0; fi28 < 2; fi28++ {
                var v29 uint16 = 232
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi28
                _ = v29
            }
        }
        var v30 uint64 = 4294967295
        _ = v26
        _ = v30
    } else {
        if v2 {
            var v31 float32 = v22
            _ = v31
        } else {
            v2 = v2
            if v2 {
                if v2 {
                    v22 -= ((3.14159 / (3.14159 + v22)) + ((0.0 * 0.0) - 1.5))
                    var v32 float32 = v22
                    var v33 float32 = (1.5 * v32)
                    _ = v32
                    _ = v33
                } else {
                    var v34 uint16 = v18
                    _ = v34
                }
                switch 9223372036854775807 {
                case -906:
                    var v35 int32 = 0
                    _ = v35
                case 127:
                    var v36 int32 = 0
                    _ = v36
                case 42:
                    var v37 uint8 = (^(v1))
                    _ = v37
                default:
                    v38 := GenS1{F0: "\n\t", F1: 4294967295}
                    _ = v38
                }
            } else {
                v18 = v18
            }
        }
    }
    v39 := GenS2{F0: -1, F1: 4294967295, F2: v1}
    v1 = (^(v1))
    fmt.Println(v18, v21)
    fmt.Println(v1, v21)
    switch -68 {
    case 9223372036854775807:
        for fi40 := 0; fi40 < 20; fi40++ {
            var v41 uint32 = 117
            v3 = append(v3, (fi40 + -9223372036854775808))
            _ = fi40
            _ = v41
        }
    }
    switch ("foo bar" + "foo bar") {
    case "hello":
        fmt.Println(v1)
    case "foo bar":
        if false {
            var v42 uint8 = (v1 * v1)
            _ = v42
        } else {
            fmt.Println(v22, v2, v2)
        }
    default:
        var v43 int8 = 0
        _ = v43
    }
    var v44 int8 = -128
    if len(v3) > 0 {
        v3[0] = -176
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v39
    _ = v44
}
