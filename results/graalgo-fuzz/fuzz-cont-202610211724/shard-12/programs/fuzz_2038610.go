package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int64 }
type GenS2 struct { F0 int8; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS4 struct { F0 int }

func main() {
    var v1 uint16 = 602
    v2 := make(map[int]string)
    var v3 float64 = (float64((((-2.25 / -2.25) + 0.0) * (float32(3.14159) - 1e10))) * float64(((((0.0 * float32(0.0)) - float32(1e10)) * 3.14159) / ((3.14159 - (0.0 / float32(((3.14159 + 0.0) + (float64(1.5) * (1e10 - float64(1.5))))))) + (1.5 * 0.1)))))
    v4 := make(map[uint32]string)
    var v5 bool = (370 > 938)
    var v6 bool = v5
    v5 = v5
    var v7 int64 = 9223372036854775807
    if (!v6) {
        for fi8 := 0; fi8 < 7; fi8++ {
            switch "a" {
            case "hello":
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 bool = false
                    delete(v2, 0)
                    _ = fi9
                    _ = v10
                }
            case "":
                switch "hello" {
                case "\n\t":
                    switch -712 {
                    case 2147483647:
                        var v11 int8 = -1
                        _ = v11
                    case -9223372036854775808:
                        v5 = false
                    }
                case "":
                    switch fi8 {
                    case 42:
                        v6 = v6
                    case 32768:
                        var v12 float32 = 1e10
                        _ = v12
                    case 187:
                        var v13 uint8 = uint8(v1)
                        _ = v13
                    }
                default:
                    v14 := []int{(fi8 / ((fi8) | 1)), 1, (fi8 >> ((666) & 63)), (fi8 - fi8)}
                    _ = v14
                }
            default:
                var v15 float64 = v3
                _ = v15
            }
            if v5 {
                var v16 float64 = (float64(0.0) - (((3.14159 * ((1e10 * v3) + float64((3.14159 / (3.14159 * 0.1))))) - v3) + (v3 / (3.14159 - v3))))
                _ = v16
            } else {
                if (1 != 1) {
                    v17 := GenS1{F0: "a", F1: 0, F2: (v7 * 367)}
                    var v18 uint32 = 2147483647
                    _ = v17
                    _ = v18
                } else {
                    v2[-581] = "hello"
                }
            }
            _ = fi8
        }
    } else {
        for fi19 := 0; fi19 < 2; fi19++ {
            if (!(v6 || (((v3 + v3) - float64(-2.25)) <= v3))) {
                fmt.Println(v1, v5)
            } else {
                var v20 float32 = 1.5
                for fi21 := 0; fi21 < 1; fi21++ {
                    v6 = true
                    v22 := GenS2{F0: -97, F1: v1, F2: 1}
                    if v5 {
                        v3 = (v3 - 0.0)
                        if ((true && (v6 && v6)) || false) {
                            var v23 uint32 = 4294967295
                            _ = v23
                        } else {
                            v3 = v3
                            v20 -= (v20 + 3.14159)
                        }
                        v24 := GenS1{F0: "a", F1: 199, F2: (-(v7))}
                        _ = v24
                    }
                    _ = fi21
                    _ = v22
                }
                _ = v20
            }
            _ = fi19
        }
        for fi25 := 0; fi25 < 19; fi25++ {
            _ = len(v2)
            _ = fi25
        }
    }
    var v26 float32 = 0.1
    switch "hello" {
    case "hello":
        if v6 {
            for fi27 := 0; fi27 < 14; fi27++ {
                v28 := GenS1{F0: ("a" + ("foo bar" + "")), F1: 0, F2: v7}
                var v29 uint8 = 127
                v7 -= v7
                _ = fi27
                _ = v28
                _ = v29
            }
            v2[0] = ("" + (((("foo bar" + "") + "a") + ("foo bar" + (("\n\t" + (("" + "a") + "hello")) + ("hello" + "foo bar")))) + "hello"))
        }
    case "":
        v6 = (!v5)
    default:
        v3 = 0.1
    }
    v3 -= 1.5
    v7 = (v7 & (v7 * v7))
    if v5 {
        var v30 float32 = 1e10
        var v31 uint16 = v1
        _ = v30
        _ = v31
    }
    v32 := GenS3{F0: v1, F1: 4294967295, F2: v1}
    v33 := GenS1{F0: (("" + "") + "hello"), F1: 47, F2: (^(v7))}
    v7 = (v7 ^ (^(v7)))
    if (!(0 < (v7 & v7))) {
        v34 := GenS4{F0: 32768}
        v7 = (v7 * v7)
        v7 = v7
        _ = v34
    }
    v5 = v6
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
