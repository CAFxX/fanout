package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32; F1 uint }

func main() {
    var v1 string = "a"
    var v2 int8 = -105
    var v3 uint16 = 840
    var v4 uint16 = v3
    var v5 float32 = (0.0 - (0.0 - float32(1.5)))
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(fi6, v3)
        _ = fi6
    }
    v7 := GenS1{F0: 255, F1: 343}
    v8 := []int{32767}
    if true {
        var v9 uint = 2
        switch -334 {
        case 229:
            v1 = v1
        case -1:
            switch 4294967295 {
            case -320:
                v10 := GenS1{F0: (v4 << (((v9 / (((v9 / ((v9) | 1))) | 1))) & 63)), F1: 4294967295}
                _ = v10
            default:
                for fi11 := 0; fi11 < 7; fi11++ {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        var v13 uint32 = uint32(v9)
                        _ = fi12
                        _ = v13
                    }
                    _ = fi11
                }
            }
        default:
            _ = len(v8)
        }
        var v14 uint16 = (v4 + v3)
        _ = v14
        _ = v9
    } else {
        if false {
            if (1 >= -2147483648) {
                var v15 uint = 957
                var v16 bool = false
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 int32 = -2147483648
                v19 := GenS2{F0: (fi17 != (fi17 << ((2) & 63))), F1: 245}
                _ = fi17
                _ = v18
                _ = v19
            }
            switch 127 {
            case -2147483648:
                fmt.Println(v3, v4, v1)
            default:
                if (((0 > 232) || false) && (!true)) {
                    for fi20 := 0; fi20 < 5; fi20++ {
                        v2 += (v2 + v2)
                        var v21 uint32 = 2147483648
                        switch (fi20 - fi20) {
                        case 0:
                            if true {
                                v8 = append(v8, (fi20 % (((fi20 % (((fi20 % (((fi20 | (fi20 >> ((995) & 63)))) | 1))) | 1))) | 1)))
                                fmt.Println(v3, v2)
                            } else {
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    v4 = 1
                                    switch ((v1 + ("hello" + "a")) + ((("hello" + "hello") + "") + v1)) {
                                    case "a":
                                        v5 = v5
                                    case "hello":
                                        var v23 float32 = ((float32(-2.25) - v5) * v5)
                                        _ = v23
                                    }
                                    fmt.Println(v5, v4, v4)
                                    var v24 uint64 = 1
                                    _ = len(v8)
                                    if false {
                                        var v25 int = -186
                                        _ = v25
                                    } else {
                                        v24 = (v24 << ((990) & 63))
                                        switch (fi20 - 128) {
                                        case 1:
                                            switch v1 {
                                            case "hello":
                                                var v26 string = v1
                                                _ = v26
                                            default:
                                                switch int(v3) {
                                                case 42:
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                case 1:
                                                    fmt.Println(fi22, v1)
                                                case 2147483647:
                                                    v4 = (v3 - 42)
                                                }
                                            }
                                        case -9223372036854775808:
                                            var v27 float32 = 0.0
                                            _ = v27
                                        default:
                                            v28 := GenS3{F0: v2}
                                            _ = v28
                                        }
                                    }
                                    _ = fi22
                                    _ = v24
                                }
                                v1 = v1
                            }
                        case 32768:
                            switch "a" {
                            case "hello":
                                for fi29 := 0; fi29 < 6; fi29++ {
                                    v21 = 2147483648
                                    v30 := GenS1{F0: v3, F1: 522}
                                    fmt.Println(v2)
                                    _ = fi29
                                    _ = v30
                                }
                            default:
                                switch fi20 {
                                case 32767:
                                    fmt.Println(v2, v1, v3)
                                case -758:
                                    switch 716 {
                                    case 127:
                                        if true {
                                            v5 += 0.1
                                            switch (v1 + "\n\t") {
                                            case "":
                                                if true {
                                                    if (833 < v3) {
                                                        v31 := GenS4{F0: ("" + (v1 + "a"))}
                                                        v5 -= -2.25
                                                        _ = v31
                                                    }
                                                    fmt.Println(v3, v3, fi20)
                                                    v32 := GenS4{F0: (v1 + "a")}
                                                    _ = v32
                                                } else {
                                                    fmt.Println(v21)
                                                }
                                            }
                                            v1 = ((v1 + "foo bar") + v1)
                                        } else {
                                            var v33 float32 = v5
                                            var v34 uint64 = 0
                                            _ = v33
                                            _ = v34
                                        }
                                    case 32767:
                                        var v35 uint16 = 2
                                        _ = v35
                                    default:
                                        var v36 float64 = -2.25
                                        _ = v36
                                    }
                                }
                            }
                        case -612:
                            switch "foo bar" {
                            case "a":
                                if ((178 >= 100) && false) {
                                    v37 := GenS5{F0: 0}
                                    v3 = (v3 & v4)
                                    _ = v37
                                }
                            }
                        default:
                            v2 -= v2
                        }
                        _ = fi20
                        _ = v21
                    }
                    var v38 uint = 0
                    var v39 int16 = 32767
                    _ = v38
                    _ = v39
                }
            }
        }
        switch "" {
        case "hello":
            v1 = v1
        }
    }
    var v40 int8 = 127
    fmt.Println(v5, v2)
    if len(v8) > 0 {
        v8[0] = 127
    }
    v41 := GenS1{F0: v4, F1: 0}
    v1 = "hello"
    v42 := GenS6{F0: -773, F1: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v7
    _ = v8
}
