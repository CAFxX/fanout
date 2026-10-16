package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }
func genFn0() int32 {
    if true {
        for fi4 := 0; fi4 < 12; fi4++ {
            var v5 int64 = -1
            _ = fi4
            _ = v5
        }
    }
    if false {
        var v6 int32 = 1
        _ = v6
    }
    return -1
}
func genFn1(p0 string) uint64 {
    p0 = ("foo bar" + p0)
    var v7 uint8 = 247
    fmt.Println(p0, v7)
    _ = p0
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 int = 2147483647
    var v2 int = (v1 / (((v1 * v1)) | 1))
    var v3 uint8 = 42
    v8 := GenS1{F0: (^(v3))}
    var v9 int32 = 1
    for fi10 := 0; fi10 < 13; fi10++ {
        v9 = (v9 << ((1) & 63))
        var v11 int16 = -789
        v12 := GenS2{F0: 9223372036854775807, F1: 857}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 float64 = 0.1
    v14 := GenS2{F0: 986, F1: 913}
    v3 = v3
    fmt.Println(v9, v13, v2)
    fmt.Println(v13)
    v9 += (v9 << ((255) & 63))
    v15 := []int{(^(v1)), (v2 / (((-(v2))) | 1)), v2}
    _ = len(v15)
    var v16 int = 9223372036854775807
    v13 = float64((3.14159 * 1.5))
    for fi17 := 0; fi17 < 5; fi17++ {
        if true {
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := GenS1{F0: 127}
                fmt.Println(v1)
                v20 := GenS3{F0: (!false)}
                switch "a" {
                case "\n\t":
                    v2 += (fi18 + 9223372036854775807)
                case "hello":
                    v2 = 32767
                case "foo bar":
                    if true {
                        var v21 int64 = 42
                        _ = len(v15)
                        _ = v21
                    } else {
                        if false {
                            _ = len(v15)
                            v22 := GenS2{F0: 9223372036854775807, F1: 65535}
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            _ = v22
                        }
                        if ((2 == 42) && true) {
                            var v23 int64 = 190
                            if (1 == 2) {
                                var v24 int32 = (v9 ^ v9)
                                _ = v24
                            } else {
                                switch "a" {
                                case "":
                                    fmt.Println(v9, v9)
                                case "\n\t":
                                    var v25 float64 = -2.25
                                    _ = v25
                                case "foo bar":
                                    var v26 float64 = 1.5
                                    _ = v26
                                default:
                                    var v27 int16 = 32767
                                    _ = v27
                                }
                                v28 := GenS3{F0: (433 < v23)}
                                _ = v28
                            }
                            _ = v23
                        }
                        switch ("a" + "foo bar") {
                        case "hello":
                            v2 = (^(v2))
                        case "a":
                            v29 := GenS1{F0: v3}
                            _ = v29
                        }
                        v30 := GenS4{F0: (v2 <= fi18)}
                        if (18446744073709551615 == 0) {
                            var v31 float64 = (float64((0.1 - 0.0)) + (3.14159 * float64((-2.25 / -2.25))))
                            _ = v31
                        }
                        _ = v30
                    }
                default:
                    switch ("hello" + "a") {
                    case "hello":
                        if (!true) {
                            if len(v15) > 0 {
                                v15[0] = fi18
                            }
                        } else {
                            var v32 uint16 = 32767
                            switch 32768 {
                            case 42:
                                if len(v15) > 0 {
                                    v15[0] = v1
                                }
                            default:
                                var v33 int64 = 1
                                _ = v33
                            }
                            _ = v32
                        }
                    case "a":
                        var v34 int32 = (v9 ^ (^(v9)))
                        _ = len(v15)
                        _ = v34
                    default:
                        v13 = ((v13 * ((0.1 / 1.0) * v13)) - (v13 / v13))
                    }
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        _ = fi17
    }
    var v35 int = v16
    v36 := GenS4{F0: (((("a" + "foo bar") + "hello") <= "\n\t") || true)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v8
    _ = v9
}
