package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 uint16; F1 string; F2 uint64 }
type GenS5 struct { F0 int8 }

func main() {
    var v1 float64 = 1e10
    var v2 uint64 = 543
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    var v5 uint16 = 65535
    fmt.Println(v1, v2, v1)
    v5 += (v5 - (v5 | v5))
    if (1 < 1) {
        v6 := GenS1{F0: "foo bar"}
        _ = v6
    }
    v7 := []int{int(v2)}
    for fi8 := 0; fi8 < 12; fi8++ {
        v2 = 2
        v9 := GenS1{F0: (("\n\t" + "foo bar") + ((("hello" + "") + "a") + "hello"))}
        v7 = append(v7, (fi8 / ((-2147483648) | 1)))
        _ = fi8
        _ = v9
    }
    var v10 int8 = -112
    if ((("a" + (("a" + (((("" + "\n\t") + ("hello" + ("a" + "\n\t"))) + "") + ("foo bar" + ""))) + "hello")) + "") < (("\n\t" + ("\n\t" + "\n\t")) + ("foo bar" + "a"))) {
        var v11 float32 = 1.5
        switch -916 {
        case 128:
            var v12 uint8 = 1
            _ = v12
        case -927:
            v13 := GenS2{F0: -687, F1: 0}
            _ = v13
        }
        _ = v11
    } else {
        var v14 int = 315
        var v15 int8 = (v10 | 100)
        var v16 uint8 = 17
        _ = v14
        _ = v15
        _ = v16
    }
    v5 = (v5 + (^(v5)))
    for fi17 := 0; fi17 < 5; fi17++ {
        switch ("foo bar" + "a") {
        case "foo bar":
            var v18 bool = false
            _ = v18
        }
        v19 := GenS2{F0: -2147483648, F1: 128}
        fmt.Println(fi17, v5)
        _ = fi17
        _ = v19
    }
    var v20 uint = 255
    for fi21 := 0; fi21 < 4; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            fmt.Println(v1)
            _ = fi22
        }
        _ = fi21
    }
    switch 2147483647 {
    case -2147483648:
        v23 := GenS3{F0: -989, F1: 465}
        _ = v23
    case 0:
        v4[0] = ("\n\t" + ("a" + "\n\t"))
    case 2147483647:
        v1 += v1
    default:
        var v24 uint16 = (v5 & (v5 / ((v5) | 1)))
        _ = v24
    }
    if true {
        var v25 uint = v20
        switch -353 {
        case -954:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        default:
            switch "foo bar" {
            case "hello":
                v25 += v25
            case "a":
                v26 := GenS3{F0: 2147483647, F1: v25}
                _ = v26
            case "":
                v25 = (^(v20))
            default:
                v7 = append(v7, 83)
            }
        }
        _ = v25
    }
    var v27 float64 = -2.25
    v4[32768] = "a"
    switch 9223372036854775807 {
    case 6:
        for fi28 := 0; fi28 < 5; fi28++ {
            if false {
                _ = len(v4)
            } else {
                var v29 int = (fi28 - (fi28 >> (((v20 + (v20 ^ 0))) & 63)))
                if ((int64(v20) != -485) || (42 == 32767)) {
                    v30 := GenS2{F0: 2147483647, F1: 255}
                    if (!(!(uint8(v20) >= 1))) {
                        _ = len(v4)
                        if (!false) {
                            switch ("" + ("foo bar" + ("hello" + "\n\t"))) {
                            case "foo bar":
                                var v31 float32 = 1e10
                                _ = v31
                            case "":
                                v10 = (-(v10))
                            case "\n\t":
                                switch (fi28 + v29) {
                                case 127:
                                    v32 := GenS4{F0: 1, F1: "", F2: v2}
                                    _ = v32
                                case -2147483648:
                                    var v33 float64 = float64((1.5 / 1.0))
                                    _ = v33
                                case 42:
                                    _ = len(v3)
                                default:
                                    var v34 int64 = 9223372036854775807
                                    _ = v34
                                }
                            }
                        } else {
                            v1 += -2.25
                        }
                    }
                    var v35 string = ""
                    _ = v30
                    _ = v35
                }
                switch "a" {
                case "":
                    fmt.Println(v2, v1)
                case "a":
                    v36 := GenS5{F0: 42}
                    _ = v36
                }
                _ = v29
            }
            _ = fi28
        }
    case -1:
        for fi37 := 0; fi37 < 6; fi37++ {
            if false {
                v20 = (v20 & (^(v20)))
                var v38 uint8 = 102
                _ = v38
            }
            delete(v4, 4294967295)
            _ = fi37
        }
    default:
        v27 = 3.14159
    }
    v1 -= float64(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
