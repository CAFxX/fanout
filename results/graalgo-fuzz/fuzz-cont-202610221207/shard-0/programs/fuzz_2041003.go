package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 uint }

func main() {
    var v1 string = ("a" + "\n\t")
    v2 := make(map[uint32]int)
    var v3 float64 = -2.25
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 13; fi5++ {
            v6 := GenS1{F0: -1, F1: 65535, F2: 18446744073709551615}
            var v7 uint8 = 1
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    delete(v2, 0)
    if (!(false || (!true))) {
        if (!(!(!((42 == 0) || false)))) {
            fmt.Println(v1, v3)
        }
        if false {
            v1 = v1
            v1 = "\n\t"
            v3 = -2.25
        }
        if (0 >= -1) {
            if true {
                v8 := GenS2{F0: ""}
                switch ("a" + "hello") {
                case "":
                    var v9 int32 = 2147483647
                    _ = v9
                default:
                    v1 = v1
                }
                for fi10 := 0; fi10 < 14; fi10++ {
                    var v11 int16 = 0
                    _ = fi10
                    _ = v11
                }
                _ = v8
            } else {
                v1 = ((v1 + "foo bar") + v1)
                v1 = v1
                var v12 float64 = float64((0.1 - 1e10))
                _ = v12
            }
            var v13 uint8 = 120
            for fi14 := 0; fi14 < 6; fi14++ {
                v13 = (v13 & 255)
                _ = fi14
            }
            _ = v13
        } else {
            if true {
                if (2 < 165) {
                    v15 := []int{0, -195, -9223372036854775808}
                    var v16 bool = (!(!true))
                    for fi17 := 0; fi17 < 20; fi17++ {
                        if (32767 <= 32767) {
                            var v18 int32 = -694
                            _ = len(v15)
                            _ = v18
                        } else {
                            var v19 uint64 = 2
                            var v20 int = fi17
                            _ = v19
                            _ = v20
                        }
                        _ = fi17
                    }
                    _ = v15
                    _ = v16
                }
                v2[2147483647] = 2147483647
            } else {
                fmt.Println(v3)
                var v21 bool = (507 <= 4294967295)
                v22 := []int{-2147483648, 1, -671}
                _ = v21
                _ = v22
            }
            v3 += (1.5 / 3.14159)
            var v23 uint16 = 570
            _ = v23
        }
    } else {
        var v24 int8 = 42
        switch "hello" {
        case "\n\t":
            for fi25 := 0; fi25 < 15; fi25++ {
                v26 := v2[42]
                fmt.Println(v1, v1, v24)
                _ = fi25
                _ = v26
            }
        case "":
            if true {
                if ((-2.25 - (v3 / ((v3 / (0.1 * v3)) * (v3 / v3)))) < v3) {
                    _ = len(v2)
                }
                if (!false) {
                    if false {
                        var v27 int = 969
                        _ = v27
                    }
                    var v28 float64 = v3
                    _ = v28
                } else {
                    var v29 bool = (!(128 >= 255))
                    var v30 bool = ((!(!v29)) || false)
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v32 := GenS3{F0: 128, F1: -1}
                        _ = fi31
                        _ = v32
                    }
                    _ = v29
                    _ = v30
                }
            } else {
                v33 := GenS4{F0: 0}
                v24 = (^(v24))
                _ = v33
            }
        default:
            switch 989 {
            case 0:
                var v34 uint64 = 812
                _ = v34
            case 32768:
                switch -686 {
                case 32768:
                    v35 := []int{285, 513}
                    _ = v35
                }
            default:
                v36 := GenS5{F0: 719, F1: (42 != -2147483648)}
                _ = v36
            }
        }
        switch -2147483648 {
        case -9223372036854775808:
            v24 = (v24 << ((688) & 63))
        case 42:
            var v37 string = ("\n\t" + (((v1 + v1) + v1) + v1))
            _ = v37
        default:
            for fi38 := 0; fi38 < 14; fi38++ {
                v3 = 1e10
                _ = fi38
            }
        }
        _ = v24
    }
    v39 := GenS6{F0: false}
    var v40 uint64 = 42
    var v41 float32 = float32(3.14159)
    for fi42 := 0; fi42 < 13; fi42++ {
        v43 := GenS2{F0: "a"}
        fmt.Println(v3, v40)
        _ = fi42
        _ = v43
    }
    v44 := GenS7{F0: (32767 <= 673), F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v44
}
