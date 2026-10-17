package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 bool; F1 int16 }
func genFn0(p0 uint, p1 int, p2 uint16) uint {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int8 = -121
        if false {
            v7 := GenS1{F0: false, F1: (v6 & (v6 * (v6 * v6)))}
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    v8 := GenS2{F0: 42, F1: 1}
    var v9 int64 = -1
    v10 := GenS3{F0: 85, F1: 18446744073709551615}
    var v11 uint16 = 861
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint8) bool {
    p0 = p0
    p0 = (p0 * p0)
    _ = p0
    return false
}

func main() {
    v1 := []uint32{2147483647, 398}
    v2 := []string{}
    v3 := make(map[string]bool)
    var v4 int16 = 32767
    fmt.Println(v4)
    var v12 bool = (false || (287 < 42))
    var v13 int = 128
    fmt.Println(v13, v4)
    v4 = v4
    var v14 uint8 = 1
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS2{F0: v4, F1: 42}
        if ("" <= "a") {
            switch "hello" {
            case "\n\t":
                v4 += v4
            }
            for fi17 := 0; fi17 < 12; fi17++ {
                fmt.Println(fi15, v13)
                _ = fi17
            }
            var v18 int32 = -752
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    v4 -= v4
    v14 -= v14
    for fi19 := 0; fi19 < 4; fi19++ {
        if (!(!false)) {
            v3["foo bar"] = ((((false || (919 >= 4294967295)) || (!(false || ((1.5 - 1.5) <= -2.25)))) && false) || v12)
            v20 := GenS4{F0: 255, F1: fi19}
            v21 := GenS5{F0: false, F1: -32768}
            _ = v20
            _ = v21
        } else {
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 bool = (1.5 >= float64(3.14159))
                var v24 float64 = 0.1
                var v25 uint16 = 2
                if ((v23 && true) || false) {
                    switch fi19 {
                    case 642:
                        var v26 float64 = ((v24 + v24) * (v24 + ((float64((1e10 - 1.5)) + v24) - 0.0)))
                        _ = v26
                    case 127:
                        if v23 {
                            fmt.Println(v23, v4, v25)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if v23 {
                                _ = len(v3)
                                fmt.Println(v13)
                                var v27 uint8 = (v14 >> ((65535) & 63))
                                fmt.Println(v13, v4, v13)
                                if v12 {
                                    fmt.Println(fi22, v25, fi19)
                                    v25 += 32767
                                }
                                _ = v27
                            }
                        }
                    case -9223372036854775808:
                        fmt.Println(v25, v12)
                    default:
                        v28 := GenS5{F0: v23, F1: v4}
                        _ = v28
                    }
                } else {
                    var v29 int16 = (v4 + -32768)
                    fmt.Println(v24, v14, v24)
                    _ = v29
                }
                v25 -= v25
                _ = len(v3)
                switch fi22 {
                case 909:
                    var v30 string = "a"
                    _ = v30
                default:
                    var v31 uint8 = 21
                    _ = v31
                }
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            fmt.Println(v13)
            v32 := GenS5{F0: true, F1: v4}
            _ = v32
        }
        if (v4 <= (v4 & v4)) {
            if (v12 && v12) {
                switch "hello" {
                case "a":
                    switch (("hello" + "hello") + ("hello" + "foo bar")) {
                    case "a":
                        var v33 bool = ((-(fi19)) >= (v13 * -2147483648))
                        _ = v33
                    }
                case "":
                    v14 -= (v14 / ((v14) | 1))
                }
            }
        }
        fmt.Println(fi19, v14)
        _ = fi19
    }
    v34 := GenS2{F0: 42, F1: 18446744073709551615}
    fmt.Println(v13, v14)
    v13 -= (v13 % (((v13 % ((v13) | 1))) | 1))
    v35 := GenS5{F0: v12, F1: v4}
    v13 = v13
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
