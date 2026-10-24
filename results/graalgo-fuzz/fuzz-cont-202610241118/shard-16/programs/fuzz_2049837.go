package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 string; F2 uint32 }
type GenS6 struct { F0 bool }

func main() {
    var v1 uint16 = 255
    var v2 uint = 65535
    v3 := []int{}
    v4 := []uint32{634, 65535}
    var v5 string = "foo bar"
    switch 723 {
    case 995:
        var v6 bool = false
        _ = v6
    case -288:
        switch v5 {
        case "":
            v7 := GenS1{F0: 32767, F1: 0, F2: v1}
            _ = v7
        case "a":
            v8 := GenS2{F0: 2147483647, F1: -283}
            _ = v8
        case "\n\t":
            if (104 > int8(v1)) {
                switch "" {
                case "hello":
                    v5 = (v5 + "")
                default:
                    var v9 int8 = int8(v1)
                    _ = v9
                }
                _ = len(v3)
                if len(v3) > 0 {
                    v3[0] = -683
                }
            } else {
                v10 := GenS2{F0: -960, F1: 32767}
                for fi11 := 0; fi11 < 7; fi11++ {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v13 := GenS2{F0: 1, F1: 1000}
                        v14 := GenS1{F0: 42, F1: 2, F2: v1}
                        var v15 int = fi12
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    _ = fi11
                }
                _ = v10
            }
        default:
            var v16 float64 = 0.0
            _ = v16
        }
    case -512:
        v17 := GenS3{F0: 973, F1: ((0.0 > (float32(1.5) - ((1e10 / ((1.5 / float32(3.14159)) + 0.0)) * 0.0))) || false)}
        _ = v17
    }
    v2 -= 171
    var v18 uint = 903
    switch 0 {
    case 0:
        var v19 float32 = float32(0.1)
        _ = v19
    case -458:
        var v20 uint32 = 527
        _ = v20
    case -9223372036854775808:
        v18 = v2
    }
    v18 += v2
    for fi21 := 0; fi21 < 6; fi21++ {
        if false {
            v2 += (^(v2))
            if false {
                v2 = (^(v2))
            }
            for fi22 := 0; fi22 < 12; fi22++ {
                v23 := GenS4{F0: (!(!true))}
                var v24 float32 = -2.25
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        _ = len(v3)
        _ = fi21
    }
    switch 122 {
    case -124:
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 float32 = ((-2.25 + (0.1 + -2.25)) - (((((3.14159 + ((float32(float64(float32(3.14159))) * 1e10) - ((float32(0.1) / 3.14159) / 1e10))) + (0.1 + 1.5)) + (1.5 + (0.0 / ((0.0 / -2.25) + (float32(-2.25) + 0.1))))) / ((1e10 / float32(-2.25)) + -2.25)) * 1e10))
            _ = fi25
            _ = v26
        }
    default:
        for fi27 := 0; fi27 < 20; fi27++ {
            v28 := GenS4{F0: true}
            if false {
                var v29 int32 = -877
                if (!(!(0.0 <= ((0.1 + (1.5 + (-2.25 + 1.5))) - 1e10)))) {
                    var v30 float64 = (0.0 + -2.25)
                    _ = v30
                }
                if false {
                    v1 = (v1 & 0)
                    for fi31 := 0; fi31 < 3; fi31++ {
                        var v32 uint32 = 42
                        var v33 bool = ((54 <= 127) && ((18446744073709551615 <= 18446744073709551615) || (false || (-128 <= -128))))
                        v34 := GenS5{F0: (v29 | (v29 >> ((v18) & 63))), F1: "foo bar", F2: v32}
                        var v35 uint64 = 601
                        v36 := GenS6{F0: v33}
                        v37 := GenS3{F0: v18, F1: ((false && v33) && v33)}
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                }
                _ = v29
            }
            var v38 uint = v18
            _ = fi27
            _ = v28
            _ = v38
        }
    }
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v5 = ("a" + v5)
    }
    switch -940 {
    case 32767:
        var v39 uint32 = 2147483647
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
