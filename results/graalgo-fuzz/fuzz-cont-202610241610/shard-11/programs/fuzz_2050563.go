package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 string }

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]string)
    var v3 int16 = 468
    var v4 float64 = (-2.25 - 1e10)
    v5 := make(map[uint32]uint32)
    v1 -= v1
    v6 := GenS1{F0: "a"}
    v3 -= v3
    var v7 uint8 = 0
    fmt.Println(v1)
    switch 4294967295 {
    case 1:
        for fi8 := 0; fi8 < 16; fi8++ {
            if ((65535 != 2) && false) {
                if false {
                    v9 := v2[65535]
                    for fi10 := 0; fi10 < 1; fi10++ {
                        v7 = (v7 * v7)
                        _ = fi10
                    }
                    _ = v9
                } else {
                    fmt.Println(v7)
                    _ = len(v2)
                    delete(v2, 4294967295)
                }
                for fi11 := 0; fi11 < 12; fi11++ {
                    v3 = v3
                    v12 := GenS1{F0: "\n\t"}
                    v13 := GenS1{F0: "a"}
                    var v14 float64 = (float64((float32(float64(float32(v4))) / (1e10 - (1.5 * 3.14159)))) - v4)
                    fmt.Println(fi11)
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                v4 = v4
            } else {
                var v15 uint32 = 4294967295
                var v16 int8 = 63
                _ = v15
                _ = v16
            }
            _ = fi8
        }
    case 825:
        if (9223372036854775807 > 1) {
            v17 := GenS1{F0: ("foo bar" + "hello")}
            _ = v17
        }
    default:
        var v18 int16 = v3
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 int64 = 221
        _ = fi19
        _ = v20
    }
    var v21 int8 = 1
    v21 -= (v21 % (((-(v21))) | 1))
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := GenS1{F0: "hello"}
        v7 = (v1 % ((0) | 1))
        v1 = (v1 >> ((534) & 63))
        _ = fi22
        _ = v23
    }
    var v24 uint32 = 660
    if true {
        fmt.Println(v3)
        var v25 uint16 = 255
        switch "a" {
        case "foo bar":
            v7 -= (v1 ^ (v1 * (v7 >> ((65535) & 63))))
        case "\n\t":
            var v26 string = "hello"
            _ = v26
        }
        _ = v25
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 bool = true
        if (-660 == 658) {
            for fi29 := 0; fi29 < 13; fi29++ {
                _ = len(v2)
                v30 := GenS2{F0: v1, F1: "a"}
                _ = fi29
                _ = v30
            }
            for fi31 := 0; fi31 < 9; fi31++ {
                v32 := []int{(fi27 << ((255) & 63)), 4294967295, (fi31 & (fi27 | (fi27 - 32767)))}
                switch -582 {
                case 127:
                    v28 = v28
                case -1:
                    if (v28 && false) {
                        v33 := GenS1{F0: ("foo bar" + "foo bar")}
                        v34 := GenS2{F0: (^(v7)), F1: ""}
                        _ = v33
                        _ = v34
                    } else {
                        v28 = (!(!((v1 + (v7 - (v7 & 0))) <= 42)))
                    }
                }
                var v35 int = 4294967295
                _ = fi31
                _ = v32
                _ = v35
            }
            if ("\n\t" == (("hello" + (("\n\t" + ("foo bar" + ("hello" + ("foo bar" + "")))) + ("\n\t" + "hello"))) + "a")) {
                for fi36 := 0; fi36 < 9; fi36++ {
                    v37 := v5[4294967295]
                    fmt.Println(v3)
                    v38 := GenS2{F0: v7, F1: "a"}
                    _ = fi36
                    _ = v37
                    _ = v38
                }
            }
        }
        _ = fi27
        _ = v28
    }
    if true {
        v4 -= v4
        fmt.Println(v24)
        v1 = v1
    } else {
        v39 := GenS1{F0: "foo bar"}
        _ = v39
    }
    v1 = v1
    fmt.Println(v21)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
