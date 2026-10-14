package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS4 struct { F0 int16; F1 bool; F2 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint16 }

func main() {
    v1 := []uint32{866, 1, 65535, 1}
    var v2 bool = true
    var v3 uint32 = 707
    v4 := make(map[int]int)
    v5 := []int{316, 32768, -2147483648}
    var v6 uint64 = 597
    if v2 {
        var v7 bool = v2
        switch (("hello" + "hello") + "hello") {
        case "foo bar":
            v5 = append(v5, 810)
        case "a":
            v8 := GenS1{F0: (v3 <= v3)}
            _ = v8
        case "":
            var v9 float32 = float32(float64((1e10 / 1.5)))
            _ = v9
        default:
            fmt.Println(v2)
        }
        for fi10 := 0; fi10 < 6; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 float32 = ((0.0 - (-2.25 - 1.5)) - 0.1)
                _ = fi11
                _ = v12
            }
            switch fi10 {
            case 32768:
                if false {
                    var v13 bool = v2
                    v13 = v2
                    for fi14 := 0; fi14 < 2; fi14++ {
                        if false {
                            v15 := GenS1{F0: (!false)}
                            fmt.Println(v7)
                            _ = v15
                        }
                        v16 := GenS2{F0: 255}
                        _ = fi14
                        _ = v16
                    }
                    _ = v13
                }
            case 42:
                fmt.Println(fi10, v3, v7)
            case -929:
                for fi17 := 0; fi17 < 17; fi17++ {
                    v3 += (v3 << ((42) & 63))
                    _ = fi17
                }
            }
            v6 -= (^(v6))
            _ = fi10
        }
        _ = v7
    }
    v18 := GenS3{F0: (v3 - (v3 - (v3 | v3))), F1: uint32(v6), F2: 0}
    _ = len(v5)
    for fi19 := 0; fi19 < 17; fi19++ {
        for fi20 := 0; fi20 < 13; fi20++ {
            v5 = append(v5, (^(fi19)))
            fmt.Println(v3, v3, fi20)
            if false {
                v21 := GenS4{F0: 42, F1: true, F2: 825}
                var v22 int16 = 0
                _ = v21
                _ = v22
            } else {
                var v23 uint8 = 127
                _ = v23
            }
            _ = fi20
        }
        v24 := GenS5{F0: int(v3)}
        if true {
            v6 -= v6
            var v25 uint64 = v6
            v26 := GenS6{F0: 0, F1: 32767}
            _ = v25
            _ = v26
        } else {
            v2 = true
            var v27 float32 = (-2.25 * 1e10)
            _ = v27
        }
        _ = fi19
        _ = v24
    }
    fmt.Println(v3, v2)
    v28 := GenS4{F0: -606, F1: ("a" < "\n\t"), F2: 838}
    delete(v4, 0)
    if v2 {
        if (!(uint8(v6) >= 2)) {
            var v29 int = -724
            _ = v29
        } else {
            _ = len(v5)
            if v2 {
                switch 388 {
                case 32767:
                    fmt.Println(v3, v2)
                case 2147483647:
                    fmt.Println(v3, v3)
                case -173:
                    delete(v4, 32767)
                default:
                    var v30 int8 = -128
                    _ = v30
                }
                _ = len(v4)
            } else {
                for fi31 := 0; fi31 < 15; fi31++ {
                    v32 := GenS7{F0: 65535}
                    _ = fi31
                    _ = v32
                }
            }
        }
        v33 := GenS5{F0: -15}
        _ = v33
    } else {
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
        v6 = (^(v6))
        v6 -= (v6 % ((v6) | 1))
    }
    fmt.Println(v6, v3, v6)
    fmt.Println(v6, v6, v3)
    var v34 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
