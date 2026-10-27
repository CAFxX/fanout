package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint8) int32 {
    var v5 int64 = -9223372036854775808
    if false {
        p0 -= p0
    } else {
        var v6 float64 = 3.14159
        _ = v6
    }
    v7 := []int{2147483647, -1}
    var v8 uint = 2
    p0 += 2
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint64 = 42
    var v2 int16 = -32768
    var v3 float32 = -2.25
    var v4 int64 = 1
    v9 := []int{32768}
    if true {
        v2 += (v2 ^ (v2 * (-(v2))))
    }
    switch -9223372036854775808 {
    case 1:
        fmt.Println(v4, v4, v3)
    case -9223372036854775808:
        if len(v9) > 0 {
            v9[0] = -1
        }
    case -2147483648:
        if false {
            var v10 int64 = v4
            for fi11 := 0; fi11 < 17; fi11++ {
                switch "hello" {
                case "\n\t":
                    v2 = 0
                }
                v9 = append(v9, fi11)
                v9 = append(v9, fi11)
                _ = fi11
            }
            var v12 float32 = 0.1
            _ = v10
            _ = v12
        }
    default:
        v2 = 32767
    }
    var v13 int64 = v4
    switch ("hello" + "a") {
    case "\n\t":
        v4 = 1
    default:
        if ((2 < 2) || (v2 <= int16(v1))) {
            if false {
                switch ("a" + ("a" + "foo bar")) {
                case "foo bar":
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                default:
                    if len(v9) > 0 {
                        v9[0] = 42
                    }
                }
                var v14 string = ("foo bar" + "a")
                _ = v14
            } else {
                v1 = 968
                if true {
                    var v15 uint8 = 127
                    _ = v15
                }
            }
            if len(v9) > 0 {
                v9[0] = -838
            }
            v13 = v4
        } else {
            var v16 uint = 2
            switch 0 {
            case 42:
                var v17 int64 = 42
                _ = v17
            case -375:
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            }
            _ = v16
        }
    }
    var v18 uint16 = 32767
    v19 := GenS1{F0: ("a" + "foo bar"), F1: 680}
    v18 -= v18
    v4 -= -1
    for fi20 := 0; fi20 < 4; fi20++ {
        v2 += 1
        _ = fi20
    }
    switch 32768 {
    case -784:
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 bool = (!(32768 <= fi21))
            _ = fi21
            _ = v22
        }
    case 2147483647:
        fmt.Println(v13, v13)
    case 762:
        if true {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            if true {
                var v23 int16 = 350
                var v24 uint8 = 1
                var v25 int = 127
                _ = v23
                _ = v24
                _ = v25
            } else {
                v18 = (^(v18))
            }
            var v26 float32 = v3
            _ = v26
        }
    default:
        if (v4 < 155) {
            v13 += -348
            if (42 > 4294967295) {
                if true {
                    _ = len(v9)
                    var v27 int8 = 1
                    _ = v27
                }
                fmt.Println(v4, v1)
            } else {
                if len(v9) > 0 {
                    v9[0] = 0
                }
                v9 = append(v9, 32768)
            }
        } else {
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 uint32 = 2147483647
                v30 := GenS2{F0: 2}
                if (0.1 <= 1e10) {
                    v1 -= (v1 / (((v1 + v1)) | 1))
                    if false {
                        fmt.Println(v3, v3)
                        if (((1.5 * (-2.25 - ((1.5 + float64(0.0)) + 0.0))) + 1.5) != -2.25) {
                            v2 = v2
                        } else {
                            v4 += 742
                            v31 := GenS1{F0: "\n\t", F1: -2147483648}
                            _ = v31
                        }
                    }
                }
                _ = fi28
                _ = v29
                _ = v30
            }
            v1 += (v1 / ((v1) | 1))
        }
    }
    var v32 string = "foo bar"
    var v33 int32 = 0
    var v34 float32 = (v3 - v3)
    v9 = append(v9, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v9
}
