package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int64 }
func genFn0() uint32 {
    v5 := GenS1{F0: "foo bar"}
    v6 := GenS2{F0: "\n\t"}
    var v7 float64 = ((0.1 / (-2.25 - 1e10)) + 3.14159)
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 int = -2147483648
    var v2 int8 = -1
    var v3 int64 = -9223372036854775808
    var v4 int8 = v2
    var v8 uint = 255
    v3 -= (v3 << ((v8) & 63))
    if (((!false) && true) || true) {
        v9 := GenS1{F0: "\n\t"}
        v3 -= 9223372036854775807
        v10 := GenS3{F0: 1}
        _ = v10
        _ = v9
    } else {
        v11 := []int{v1, (v1 % (((v1 / (((v1 + int(v3))) | 1))) | 1))}
        var v12 uint32 = 0
        v12 = (^(v12))
        _ = v11
        _ = v12
    }
    v4 = (v2 - v4)
    v13 := []int{(v1 * v1), (v1 ^ 32768), v1}
    var v14 bool = ((32767 != (-(v1))) && true)
    switch v1 {
    case 4294967295:
        switch "\n\t" {
        case "foo bar":
            if (-2.25 != 0.1) {
                v13 = append(v13, (v1 << ((18446744073709551615) & 63)))
                switch v1 {
                case 128:
                    for fi15 := 0; fi15 < 18; fi15++ {
                        v16 := GenS3{F0: 729}
                        _ = fi15
                        _ = v16
                    }
                }
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            } else {
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 uint = 1
                    var v19 uint32 = genFn0()
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            }
        case "hello":
            switch 4294967295 {
            case 42:
                var v20 string = "hello"
                _ = v20
            case 502:
                switch "a" {
                case "hello":
                    if (int(v8) >= -627) {
                        if v14 {
                            var v21 float32 = float32(3.14159)
                            _ = len(v13)
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 10; fi22++ {
                            v2 -= (v4 % ((-37) | 1))
                            _ = fi22
                        }
                        if (!v14) {
                            var v23 uint32 = 2147483648
                            _ = v23
                        } else {
                            var v24 int16 = 0
                            v25 := GenS3{F0: 1}
                            _ = v24
                            _ = v25
                        }
                    }
                case "a":
                    var v26 uint64 = 1
                    _ = v26
                default:
                    for fi27 := 0; fi27 < 5; fi27++ {
                        if false {
                            v14 = (false || true)
                        } else {
                            v1 = 92
                        }
                        if ((982 != uint32(v8)) && ((false && ((3.14159 <= 1e10) || (((v4 % ((0) | 1)) == (v2 ^ (v2 / ((34) | 1)))) && (42 >= 65535)))) || v14)) {
                            fmt.Println(v2)
                        }
                        _ = fi27
                    }
                }
            }
        case "a":
            v4 = v4
        default:
            if len(v13) > 0 {
                v13[0] = -2147483648
            }
        }
    case 2147483647:
        var v28 uint16 = 887
        _ = v28
    case -586:
        var v29 int8 = (v4 ^ v4)
        _ = v29
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        v14 = (v8 == 179)
        var v31 int16 = 1
        _ = fi30
        _ = v31
    }
    var v32 float32 = ((float32(0.0) / 1e10) + 3.14159)
    v33 := GenS3{F0: 42}
    for fi34 := 0; fi34 < 10; fi34++ {
        v13 = append(v13, (v1 % ((fi34) | 1)))
        switch v1 {
        case 4294967295:
            v3 = (^(v3))
        case 128:
            var v35 uint8 = 57
            _ = v35
        default:
            if (-32768 >= 1) {
                var v36 string = ("a" + "hello")
                v37 := GenS4{F0: 4294967295, F1: (v3 & (-(v3)))}
                v3 += (v3 >> (((^(v8))) & 63))
                _ = v36
                _ = v37
            }
        }
        _ = fi34
    }
    v38 := GenS2{F0: "a"}
    v1 -= v1
    if (!v14) {
        for fi39 := 0; fi39 < 18; fi39++ {
            v14 = (true || v14)
            fmt.Println(fi39, v3)
            _ = fi39
        }
        _ = len(v13)
    } else {
        var v40 int = 2147483647
        fmt.Println(v3)
        _ = v40
    }
    var v41 int = (^(v1))
    switch "\n\t" {
    case "a":
        var v42 uint16 = 255
        _ = v42
    case "\n\t":
        var v43 string = "\n\t"
        _ = v43
    default:
        v44 := GenS2{F0: ("hello" + "hello")}
        _ = v44
    }
    var v45 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v41
    _ = v45
    _ = v8
}
