package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint; F1 int32; F2 uint64 }
type GenS3 struct { F0 string; F1 uint32 }

func main() {
    var v1 int = 176
    var v2 int16 = 887
    var v3 int64 = 0
    v4 := make(map[int]bool)
    v5 := make(map[int]uint32)
    fmt.Println(v1)
    var v6 int8 = -128
    v7 := GenS1{F0: 2147483647, F1: 838}
    v6 += v6
    if false {
        switch v1 {
        case 4294967295:
            fmt.Println(v3)
        case 32768:
            fmt.Println(v3, v2)
        case -799:
            v1 += -2147483648
        }
        for fi8 := 0; fi8 < 6; fi8++ {
            v2 = (v2 - v2)
            _ = fi8
        }
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint8 = 42
            _ = fi9
            _ = v10
        }
    } else {
        v2 += (v2 | (-(v2)))
        v11 := GenS1{F0: -206, F1: 18446744073709551615}
        for fi12 := 0; fi12 < 9; fi12++ {
            v6 += v6
            for fi13 := 0; fi13 < 11; fi13++ {
                _ = len(v4)
                switch "" {
                case "":
                    var v14 int8 = -128
                    _ = v14
                default:
                    var v15 uint8 = 255
                    _ = v15
                }
                var v16 bool = (-2.25 > ((float64(-2.25) - ((0.0 - (((3.14159 / 0.1) / 1.5) - 1.5)) + 1e10)) * 0.1))
                _ = fi13
                _ = v16
            }
            delete(v4, v1)
            _ = fi12
        }
        _ = v11
    }
    var v17 bool = ((v2 | v2) >= (-(v2)))
    if v17 {
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 bool = (!(4294967295 > 0))
            v20 := GenS1{F0: -2147483648, F1: 18446744073709551615}
            _ = fi18
            _ = v19
            _ = v20
        }
    } else {
        if (!false) {
            if ("" == ("a" + ("\n\t" + "\n\t"))) {
                v21 := GenS1{F0: 1, F1: 2}
                _ = v21
            } else {
                var v22 int64 = -716
                _ = v22
            }
            v6 -= 100
            v17 = v17
        } else {
            if false {
                switch v1 {
                case -233:
                    for fi23 := 0; fi23 < 11; fi23++ {
                        for fi24 := 0; fi24 < 5; fi24++ {
                            var v25 int = 32768
                            _ = fi24
                            _ = v25
                        }
                        v26 := GenS2{F0: 245, F1: 0, F2: 0}
                        _ = fi23
                        _ = v26
                    }
                default:
                    switch int(v2) {
                    case -9223372036854775808:
                        for fi27 := 0; fi27 < 6; fi27++ {
                            v3 = 252
                            _ = len(v4)
                            _ = fi27
                        }
                    case 32767:
                        for fi28 := 0; fi28 < 4; fi28++ {
                            var v29 string = ("hello" + "")
                            delete(v4, fi28)
                            v2 = -1
                            _ = fi28
                            _ = v29
                        }
                    case 9223372036854775807:
                        if false {
                            var v30 uint = 4294967295
                            _ = len(v5)
                            v1 += v1
                            _ = v30
                        } else {
                            if false {
                                if (((true && (0 > 262)) && ((v6 * (v6 & (-(v6)))) < v6)) && false) {
                                    for fi31 := 0; fi31 < 17; fi31++ {
                                        for fi32 := 0; fi32 < 13; fi32++ {
                                            var v33 bool = (false || v17)
                                            var v34 string = (("\n\t" + "") + ("foo bar" + "\n\t"))
                                            if v33 {
                                                var v35 uint8 = 0
                                                _ = v35
                                            }
                                            var v36 int16 = 208
                                            var v37 bool = true
                                            _ = fi32
                                            _ = v33
                                            _ = v34
                                            _ = v36
                                            _ = v37
                                        }
                                        v38 := GenS1{F0: -261, F1: 143}
                                        _ = fi31
                                        _ = v38
                                    }
                                }
                            }
                            var v39 int64 = 42
                            _ = v39
                        }
                    default:
                        var v40 string = ((("foo bar" + "\n\t") + "a") + "a")
                        _ = v40
                    }
                }
                var v41 uint8 = 239
                for fi42 := 0; fi42 < 6; fi42++ {
                    v2 = (v2 >> ((42) & 63))
                    _ = fi42
                }
                _ = v41
            } else {
                v2 = (v2 << ((65535) & 63))
                v1 -= (-(v1))
            }
            switch 1 {
            case -9223372036854775808:
                var v43 float64 = 0.1
                _ = v43
            case -1:
                v44 := GenS3{F0: ((("\n\t" + ((((("foo bar" + "hello") + "hello") + "\n\t") + "") + ("" + "foo bar"))) + "") + ""), F1: 713}
                _ = v44
            case 0:
                delete(v4, (v1 >> ((970) & 63)))
            default:
                var v45 int64 = v3
                _ = v45
            }
        }
        fmt.Println(v6)
    }
    var v46 string = "hello"
    v46 = (v46 + "a")
    v1 = v1
    fmt.Println(v1)
    var v47 uint = 363
    var v48 uint64 = 0
    fmt.Println(v48)
    v2 = (v2 + (v2 / ((v2) | 1)))
    delete(v4, (v1 << ((uint(v48)) & 63)))
    var v49 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
