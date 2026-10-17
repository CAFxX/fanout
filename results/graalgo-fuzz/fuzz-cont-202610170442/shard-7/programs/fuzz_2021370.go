package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16 }
func genFn0() int64 {
    var v5 bool = (1 == 0)
    v6 := []int{424, 32767, 863, -1}
    switch "" {
    case "a":
        if (v5 && true) {
            for fi7 := 0; fi7 < 3; fi7++ {
                v5 = (v5 && true)
                _ = fi7
            }
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 string = ""
                _ = fi8
                _ = v9
            }
        } else {
            if v5 {
                fmt.Println(v5)
            }
            v5 = (-1 > 42)
        }
    case "foo bar":
        v5 = (642 <= -828)
    default:
        if v5 {
            if v5 {
                if (63 != 0) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v10 int32 = 0
                    for fi11 := 0; fi11 < 2; fi11++ {
                        var v12 bool = (204 > -121)
                        if (2 >= 42) {
                            var v13 string = "hello"
                            _ = v13
                        } else {
                            v12 = (4294967295 != 958)
                            fmt.Println(v10, v12, v10)
                        }
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                } else {
                    _ = len(v6)
                    switch ((("" + "hello") + "\n\t") + "hello") {
                    case "":
                        for fi14 := 0; fi14 < 20; fi14++ {
                            if (65535 <= 0) {
                                v5 = (!v5)
                                if (((("a" + "a") + ("a" + "")) + "\n\t") < "") {
                                    if (!("" == "\n\t")) {
                                        if (-356 < 42) {
                                            var v15 bool = ((255 != 32767) && (1000 != 32767))
                                            for fi16 := 0; fi16 < 20; fi16++ {
                                                var v17 uint8 = 123
                                                v15 = v5
                                                var v18 uint64 = 2
                                                _ = fi16
                                                _ = v17
                                                _ = v18
                                            }
                                            _ = v15
                                        } else {
                                            switch "hello" {
                                            case "a":
                                                v19 := GenS1{F0: 0, F1: 0}
                                                _ = v19
                                            default:
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            }
                                        }
                                        if v5 {
                                            for fi20 := 0; fi20 < 9; fi20++ {
                                                if true {
                                                    switch "foo bar" {
                                                    case "":
                                                        if false {
                                                            v5 = (9223372036854775807 <= -1)
                                                            if len(v6) > 0 {
                                                                v6[0] = fi20
                                                            }
                                                        }
                                                    }
                                                    var v21 float32 = 1.5
                                                    _ = v21
                                                }
                                                _ = fi20
                                            }
                                            v5 = false
                                        }
                                        switch "" {
                                        case "":
                                            for fi22 := 0; fi22 < 2; fi22++ {
                                                v23 := GenS2{F0: -538}
                                                fmt.Println(fi22, fi22)
                                                if false {
                                                    if len(v6) > 0 {
                                                        v6[0] = 2147483647
                                                    }
                                                    v24 := GenS3{F0: 42}
                                                    fmt.Println(fi22, v5, fi14)
                                                    _ = v24
                                                }
                                                _ = fi22
                                                _ = v23
                                            }
                                        }
                                    } else {
                                        for fi25 := 0; fi25 < 7; fi25++ {
                                            _ = len(v6)
                                            _ = fi25
                                        }
                                        v5 = (1.5 >= ((0.0 + ((3.14159 + (0.1 / 1.0)) + (3.14159 * 0.1))) * -2.25))
                                    }
                                    v5 = (((1.5 - ((float32((-2.25 - 0.1)) * 0.1) - -2.25)) * float32((0.1 - 1.5))) != 0.1)
                                    v26 := GenS3{F0: -984}
                                    _ = v26
                                } else {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v5 = (false && v5)
                                    case "":
                                        v5 = ((!(822 < 19)) && v5)
                                    default:
                                        if len(v6) > 0 {
                                            v6[0] = fi14
                                        }
                                    }
                                    v5 = ("foo bar" <= ("\n\t" + "foo bar"))
                                    var v27 uint32 = 1
                                    _ = v27
                                }
                                var v28 int8 = -23
                                _ = v28
                            }
                            v5 = (-1 < 42)
                            _ = fi14
                        }
                    case "\n\t":
                        v6 = append(v6, -585)
                    case "foo bar":
                        switch "foo bar" {
                        case "foo bar":
                            var v29 float32 = (0.1 * (0.1 - 1.5))
                            _ = v29
                        case "\n\t":
                            v5 = (true && (v5 || ((695 < 42) || v5)))
                        case "":
                            var v30 float32 = (-2.25 / (((-2.25 * float32(-2.25)) - float32((1.5 - 0.0))) + -2.25))
                            _ = v30
                        default:
                            var v31 float64 = (0.1 / 1.0)
                            _ = v31
                        }
                    }
                }
                v5 = (!(2 < 908))
            }
            fmt.Println(v5)
        }
    }
    if ((("\n\t" + "hello") != ("hello" + ("\n\t" + ("" + "hello")))) && v5) {
        v5 = (!v5)
    }
    v5 = v5
    _ = v5
    _ = v6
    return -9223372036854775808
}

func main() {
    var v1 float32 = -2.25
    var v2 string = "foo bar"
    v3 := make(map[uint32]bool)
    var v4 string = "a"
    v32 := GenS4{F0: -35}
    v1 = -2.25
    v33 := []int{}
    v2 = (v2 + v2)
    _ = len(v3)
    fmt.Println(v4, v4)
    v34 := GenS1{F0: 30, F1: 2147483648}
    v3[2147483647] = (65535 == 32767)
    v4 = ((v4 + "hello") + ((v4 + ((((((v4 + ("a" + "hello")) + ("a" + v4)) + (v2 + "foo bar")) + (v2 + ("a" + ("hello" + (v2 + v2))))) + "") + "\n\t")) + (v2 + ("hello" + (v4 + "foo bar")))))
    v35 := GenS2{F0: 920}
    v2 = v2
    v1 = (0.1 - float32(1.5))
    v1 = 3.14159
    v36 := GenS3{F0: 2147483647}
    _ = len(v33)
    fmt.Println(v2, v2, v4)
    v4 = "foo bar"
    for fi37 := 0; fi37 < 6; fi37++ {
        fmt.Println(fi37, v2, v2)
        v38 := GenS5{F0: 188}
        _ = len(v33)
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
