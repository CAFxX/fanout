package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint, p1 uint8) uint16 {
    if true {
        var v6 uint16 = 1
        if (!(268 != 42)) {
            if (true && true) {
                v6 = 42
            } else {
                v7 := []int{32768, 1, 9223372036854775807}
                _ = v7
            }
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := GenS1{F0: true}
                _ = fi8
                _ = v9
            }
            switch 32767 {
            case 9223372036854775807:
                var v10 uint = (^(p0))
                _ = v10
            case 911:
                v11 := GenS1{F0: false}
                _ = v11
            }
        } else {
            switch "" {
            case "hello":
                p1 = 128
            case "a":
                p0 = p0
            default:
                var v12 int32 = 127
                _ = v12
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 uint8 = p1
                if true {
                    v15 := GenS1{F0: (-950 > -1)}
                    v16 := GenS1{F0: (true || false)}
                    _ = v15
                    _ = v16
                } else {
                    v17 := []int{(fi13 << ((p0) & 63)), fi13}
                    for fi18 := 0; fi18 < 20; fi18++ {
                        if false {
                            v6 += v6
                            p1 = (v14 / ((p1) | 1))
                        }
                        v17 = append(v17, (fi18 * fi13))
                        _ = len(v17)
                        _ = fi18
                    }
                    _ = v17
                }
                _ = fi13
                _ = v14
            }
        }
        _ = v6
    }
    p1 = p1
    fmt.Println(p0, p0)
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 42
}
func genFn1() int8 {
    var v19 float64 = 0.1
    v19 = 0.1
    _ = v19
    return -55
}

func main() {
    var v1 int64 = -853
    var v2 int8 = 40
    var v3 int8 = v2
    var v4 string = ""
    v5 := []int{42, -989}
    for fi20 := 0; fi20 < 20; fi20++ {
        v21 := GenS1{F0: true}
        v3 += (-(v2))
        _ = fi20
        _ = v21
    }
    switch -960 {
    case 9223372036854775807:
        var v22 float32 = 0.0
        _ = v22
    case 42:
        var v23 int32 = 0
        _ = v23
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v4 = v4
    var v24 uint64 = 42
    v2 -= v3
    v25 := GenS1{F0: ((v3 >> ((959) & 63)) < (v2 & (v2 % (((v3 ^ v3)) | 1))))}
    var v26 uint = 42
    switch 0 {
    case -9223372036854775808:
        fmt.Println(v3, v3)
    case 0:
        v24 = (^(v24))
    case 978:
        fmt.Println(v1)
    }
    v3 = -1
    var v27 uint = (v26 + 65535)
    switch v4 {
    case "foo bar":
        for fi28 := 0; fi28 < 11; fi28++ {
            v29 := GenS2{F0: (v24 % ((0) | 1)), F1: 672}
            _ = fi28
            _ = v29
        }
    case "":
        for fi30 := 0; fi30 < 12; fi30++ {
            for fi31 := 0; fi31 < 14; fi31++ {
                var v32 uint16 = 42
                var v33 int16 = 0
                _ = fi31
                _ = v32
                _ = v33
            }
            v1 = v1
            v26 = (v26 % (((v27 / (((v27 & 65535)) | 1))) | 1))
            _ = fi30
        }
    }
    v3 += (v3 / (((v3 * (v3 * (^(v2))))) | 1))
    v34 := GenS2{F0: v24, F1: 42}
    v4 = "\n\t"
    if ((v4 != ((("a" + (v4 + ("a" + (((v4 + v4) + (("" + v4) + v4)) + "hello")))) + "") + (v4 + (v4 + v4)))) || (!(!((v24 - v24) != v24)))) {
        v1 += v1
        var v35 bool = (1 >= -2147483648)
        _ = v35
    } else {
        for fi36 := 0; fi36 < 2; fi36++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v37 := GenS1{F0: true}
            _ = fi36
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
