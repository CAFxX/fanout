package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint16; F1 int; F2 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 int; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint; F2 int32 }
type GenS7 struct { F0 int8; F1 bool }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]int)
    v3 := []int{1, 0, -256}
    v4 := make(map[uint32]uint32)
    v1 -= v1
    for fi5 := 0; fi5 < 7; fi5++ {
        v1 = v1
        for fi6 := 0; fi6 < 1; fi6++ {
            v1 = v1
            v7 := GenS1{F0: 0, F1: "hello"}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v8 := GenS2{F0: 42, F1: "foo bar"}
    var v9 int8 = 42
    switch -93 {
    case 127:
        fmt.Println(v9, v9)
    case -2147483648:
        v10 := GenS2{F0: 0, F1: "\n\t"}
        _ = v10
    default:
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 bool = (true && ((!false) || (!false)))
            for fi13 := 0; fi13 < 8; fi13++ {
                switch fi13 {
                case 128:
                    v14 := GenS2{F0: 42, F1: "a"}
                    _ = v14
                case 127:
                    v3 = append(v3, fi13)
                case -1:
                    var v15 uint = 2
                    _ = v15
                }
                var v16 int64 = -30
                v17 := GenS3{F0: 255, F1: 0, F2: (("hello" + "foo bar") + "a")}
                _ = fi13
                _ = v16
                _ = v17
            }
            _ = fi11
            _ = v12
        }
    }
    if false {
        if (!(float64(v1) >= float64((((1.5 / (1.5 / 3.14159)) - v1) + 0.0)))) {
            _ = len(v3)
            switch 32767 {
            case 32768:
                for fi18 := 0; fi18 < 15; fi18++ {
                    if len(v3) > 0 {
                        v3[0] = (fi18 % ((fi18) | 1))
                    }
                    if len(v3) > 0 {
                        v3[0] = fi18
                    }
                    for fi19 := 0; fi19 < 4; fi19++ {
                        v1 = ((float32(0.0) * 0.1) * (v1 - 0.1))
                        var v20 string = ("\n\t" + "foo bar")
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
            }
        }
        if (0.0 >= 0.1) {
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 uint32 = 981
                v9 = v9
                _ = fi21
                _ = v22
            }
            if (!true) {
                if false {
                    v23 := GenS4{F0: 18446744073709551615}
                    _ = v23
                } else {
                    var v24 uint = 2
                    v25 := GenS5{F0: -755, F1: -2147483648, F2: 0}
                    _ = v24
                    _ = v25
                }
                var v26 uint64 = 4294967295
                _ = v26
            }
            var v27 uint = 255
            _ = v27
        } else {
            switch 0 {
            case 710:
                if true {
                    v28 := GenS5{F0: -163, F1: -2147483648, F2: -72}
                    var v29 int64 = -9223372036854775808
                    _ = len(v3)
                    _ = v28
                    _ = v29
                }
            case 42:
                fmt.Println(v1, v9)
            default:
                v1 += float32((1e10 * float64((float32((3.14159 + 0.0)) - v1))))
            }
        }
        fmt.Println(v1, v9)
    }
    if ((!true) && (!false)) {
        v30 := GenS6{F0: 42, F1: 22, F2: 0}
        v9 -= (v9 & v9)
        _ = v30
    }
    v31 := GenS5{F0: 9223372036854775807, F1: 578, F2: 0}
    if false {
        v32 := GenS7{F0: v9, F1: true}
        v33 := v2[1]
        _ = v32
        _ = v33
    } else {
        var v34 uint = 255
        v1 = -2.25
        fmt.Println(v34)
        _ = v34
    }
    var v35 uint8 = 84
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v8
    _ = v9
}
