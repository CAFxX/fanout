package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = -32768
    v3 := []uint32{}
    var v4 float32 = 1e10
    v5 := GenS1{F0: "\n\t", F1: 42, F2: 2}
    v4 += v4
    switch "a" {
    case "":
        if len(v3) > 0 {
            v3[0] = 1
        }
    case "\n\t":
        if true {
            v2 = v2
            switch 42 {
            case -509:
                v6 := GenS2{F0: 0}
                _ = v6
            case 32767:
                if ((("\n\t" > "foo bar") && (int16(v1) != v2)) && false) {
                    for fi7 := 0; fi7 < 14; fi7++ {
                        for fi8 := 0; fi8 < 10; fi8++ {
                            var v9 int = 128
                            _ = fi8
                            _ = v9
                        }
                        var v10 float64 = 1.5
                        for fi11 := 0; fi11 < 20; fi11++ {
                            fmt.Println(fi7, fi11)
                            v3 = append(v3, 447)
                            fmt.Println(v10, v2, fi11)
                            switch "" {
                            case "a":
                                fmt.Println(fi7, v10)
                                var v12 int8 = 38
                                v13 := GenS1{F0: "\n\t", F1: 846, F2: 255}
                                _ = v12
                                _ = v13
                            default:
                                v10 = float64(float32(3.14159))
                            }
                            v10 = (v10 * float64(0.1))
                            _ = fi11
                        }
                        _ = fi7
                        _ = v10
                    }
                    if true {
                        _ = len(v3)
                        var v14 string = "\n\t"
                        _ = v14
                    } else {
                        v4 -= v4
                        for fi15 := 0; fi15 < 7; fi15++ {
                            if false {
                                if (!(2147483647 == -2147483648)) {
                                    fmt.Println(v4)
                                } else {
                                    v4 -= float32((1e10 / float64((v4 - ((((0.1 + v4) / 1.0) * v4) + (v4 * float32(-2.25)))))))
                                    v2 -= (^(v2))
                                    v1 += v1
                                }
                                v3 = append(v3, 0)
                            }
                            _ = fi15
                        }
                        v16 := GenS3{F0: 1}
                        _ = v16
                    }
                    if true {
                        for fi17 := 0; fi17 < 12; fi17++ {
                            var v18 bool = (-420 < -933)
                            v4 -= 1e10
                            if (v18 && v18) {
                                for fi19 := 0; fi19 < 19; fi19++ {
                                    var v20 int8 = 42
                                    _ = fi19
                                    _ = v20
                                }
                                _ = len(v3)
                            }
                            _ = fi17
                            _ = v18
                        }
                    }
                }
            }
        }
    default:
        if false {
            for fi21 := 0; fi21 < 5; fi21++ {
                for fi22 := 0; fi22 < 9; fi22++ {
                    switch ("foo bar" + ("hello" + ("" + "a"))) {
                    case "hello":
                        v23 := GenS3{F0: 884}
                        var v24 bool = (3.14159 > float32(0.0))
                        _ = v23
                        _ = v24
                    default:
                        v25 := GenS3{F0: 209}
                        _ = v25
                    }
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        _ = len(v3)
                    case "":
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                    case "\n\t":
                        var v26 bool = (!(v1 > (v1 % ((v1) | 1))))
                        _ = v26
                    }
                    fmt.Println(fi21, fi21)
                    _ = fi22
                }
                v27 := GenS2{F0: 5}
                var v28 uint32 = 4294967295
                _ = fi21
                _ = v27
                _ = v28
            }
            v29 := GenS1{F0: ((("foo bar" + "hello") + "a") + ""), F1: 0, F2: 42}
            var v30 uint32 = 725
            _ = v29
            _ = v30
        } else {
            if (269 != 0) {
                v2 = v2
            } else {
                v1 = v1
                for fi31 := 0; fi31 < 10; fi31++ {
                    v1 = (v1 * (v1 - v1))
                    _ = fi31
                }
            }
            v1 = -9223372036854775808
            for fi32 := 0; fi32 < 16; fi32++ {
                if ((!(true || true)) && true) {
                    var v33 int = fi32
                    _ = v33
                } else {
                    var v34 uint = 4294967295
                    _ = v34
                }
                for fi35 := 0; fi35 < 2; fi35++ {
                    v36 := GenS4{F0: 255, F1: (fi32 / ((fi32) | 1))}
                    _ = fi35
                    _ = v36
                }
                _ = fi32
            }
        }
    }
    v2 = (^(v2))
    v2 -= (v2 - (v2 - v2))
    fmt.Println(v4, v4, v1)
    v4 -= v4
    v4 -= ((-2.25 * v4) * float32(float64(0.0)))
    var v37 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
