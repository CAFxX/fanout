package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int16; F1 uint; F2 int32 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 42
    var v3 uint8 = (v2 & v2)
    if (!(true || (false && (-245 > 127)))) {
        if (1 != 92) {
            if (true && false) {
                v4 := GenS1{F0: 9223372036854775807, F1: 0}
                var v5 float64 = 1e10
                _ = v4
                _ = v5
            } else {
                delete(v1, 9223372036854775807)
                fmt.Println(v3, v2)
                if (("" + ("a" + ((("\n\t" + "foo bar") + ((("a" + "\n\t") + "a") + ((("\n\t" + "\n\t") + (("hello" + "") + ("a" + "\n\t"))) + (("a" + "\n\t") + "")))) + ""))) >= "foo bar") {
                    for fi6 := 0; fi6 < 20; fi6++ {
                        if ((255 != 1) && (false && false)) {
                            var v7 float64 = (-2.25 / float64(-2.25))
                            for fi8 := 0; fi8 < 19; fi8++ {
                                var v9 float32 = 1e10
                                v2 += (v2 << ((18446744073709551615) & 63))
                                v3 = 95
                                v3 += (v2 * v3)
                                _ = fi8
                                _ = v9
                            }
                            for fi10 := 0; fi10 < 9; fi10++ {
                                fmt.Println(v7, v3, v2)
                                fmt.Println(v7)
                                _ = fi10
                            }
                            _ = v7
                        } else {
                            var v11 int32 = -1
                            _ = v11
                        }
                        fmt.Println(fi6, v3)
                        _ = fi6
                    }
                }
            }
            switch "\n\t" {
            case "":
                v2 -= 1
            case "a":
                v12 := []int{1, 2147483647}
                _ = v12
            }
            var v13 uint64 = 1
            _ = v13
        }
        fmt.Println(v2, v3)
    } else {
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := GenS2{F0: -834, F1: ("\n\t" + "a")}
            _ = fi14
            _ = v15
        }
        fmt.Println(v2, v3)
    }
    if (v2 >= (v2 * (^(v3)))) {
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 uint = 2
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS2{F0: 1, F1: "\n\t"}
                _ = fi18
                _ = v19
            }
            _ = fi16
            _ = v17
        }
        if (481 == 211) {
            var v20 int32 = 317
            _ = v20
        }
    } else {
        v21 := GenS1{F0: -9223372036854775808, F1: 2147483647}
        _ = v21
    }
    var v22 int64 = -589
    v23 := GenS3{F0: 32767, F1: 0, F2: 42}
    fmt.Println(v2)
    v24 := GenS4{F0: true, F1: v3}
    if (false || (!true)) {
        var v25 int64 = v22
        for fi26 := 0; fi26 < 11; fi26++ {
            switch (("\n\t" + (("hello" + "hello") + "hello")) + "") {
            case "":
                v27 := GenS5{F0: "hello"}
                _ = v27
            case "foo bar":
                fmt.Println(v22)
            case "hello":
                fmt.Println(fi26, v3)
            default:
                fmt.Println(v25)
            }
            var v28 float32 = float32(3.14159)
            switch "" {
            case "":
                v29 := GenS4{F0: false, F1: v2}
                _ = v29
            case "foo bar":
                v30 := GenS6{F0: 28}
                _ = v30
            case "a":
                switch fi26 {
                case 2147483647:
                    v1[fi26] = ""
                case 9223372036854775807:
                    v31 := GenS7{F0: 1}
                    _ = v31
                default:
                    v32 := v1[fi26]
                    _ = v32
                }
            default:
                v3 -= (v2 >> ((145) & 63))
            }
            _ = fi26
            _ = v28
        }
        _ = v25
    }
    v33 := GenS3{F0: -32768, F1: 382, F2: 0}
    var v34 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v34
}
