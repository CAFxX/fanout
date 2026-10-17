package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 int64 }

func main() {
    v1 := []string{"", "hello"}
    var v2 int64 = -573
    var v3 int64 = v2
    v4 := make(map[uint32]uint32)
    v3 = (v2 - (-(v2)))
    switch "hello" {
    case "a":
        v2 -= v3
    }
    fmt.Println(v3, v2)
    v3 += v2
    for fi5 := 0; fi5 < 11; fi5++ {
        if false {
            for fi6 := 0; fi6 < 10; fi6++ {
                if ((!false) && (true || (!((v3 >= v3) || (!(2147483648 > 544)))))) {
                    v2 = (v2 >> ((494) & 63))
                    var v7 string = ("a" + "")
                    _ = v7
                } else {
                    if true {
                        if false {
                            var v8 int64 = (v2 % ((v2) | 1))
                            fmt.Println(fi6)
                            fmt.Println(v8, v2, v3)
                            _ = v8
                        } else {
                            fmt.Println(fi5, v2)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v1 = append(v1, ("a" + "foo bar"))
                        }
                        v9 := GenS1{F0: 1, F1: 100}
                        v2 = v3
                        v10 := GenS2{F0: false}
                        _ = v10
                        _ = v9
                    } else {
                        v11 := GenS2{F0: (false || (!true))}
                        if (-1 > 42) {
                            v3 = (v2 ^ (v3 % (((v3 * v3)) | 1)))
                            var v12 float64 = 1e10
                            _ = v12
                        } else {
                            var v13 int32 = -1
                            _ = v13
                        }
                        if (false && true) {
                            v2 = v2
                        } else {
                            v2 = (v3 * 9223372036854775807)
                        }
                        _ = v11
                    }
                }
                v2 = (v2 - (v2 << ((2) & 63)))
                _ = fi6
            }
            v2 = (^(v2))
        } else {
            v4[2147483648] = 1
            v3 -= v3
            v14 := GenS1{F0: 42, F1: 42}
            _ = v14
        }
        _ = fi5
    }
    var v15 int16 = 0
    v16 := GenS3{F0: (v3 << ((473) & 63)), F1: 0}
    for fi17 := 0; fi17 < 2; fi17++ {
        if true {
            if true {
                switch fi17 {
                case -543:
                    v18 := v4[574]
                    _ = v18
                default:
                    var v19 int8 = 0
                    _ = v19
                }
                var v20 float64 = float64(3.14159)
                for fi21 := 0; fi21 < 17; fi21++ {
                    v2 += v2
                    var v22 uint32 = 1
                    v3 = -9223372036854775808
                    _ = fi21
                    _ = v22
                }
                _ = v20
            } else {
                fmt.Println(v3, v3)
                fmt.Println(v2)
            }
        }
        v4[743] = 516
        _ = fi17
    }
    var v23 float64 = (0.1 - 1e10)
    var v24 float64 = ((v23 * 1.5) / 1e10)
    v1 = append(v1, "foo bar")
    v2 -= v3
    v25 := GenS4{F0: 854}
    if true {
        v26 := GenS5{F0: 100, F1: -1}
        v27 := v4[0]
        _ = v26
        _ = v27
    }
    fmt.Println(v23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
