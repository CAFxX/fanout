package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int; F1 int32; F2 int }
type GenS5 struct { F0 uint64; F1 string; F2 int16 }
func genFn0(p0 uint) int64 {
    var v5 uint8 = 224
    v5 += (v5 + (v5 * v5))
    var v6 int32 = -2147483648
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int, p1 int32) uint16 {
    fmt.Println(p1, p0)
    var v7 uint = 18446744073709551615
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = (0 != 2)
    var v3 float64 = (1.5 / 0.1)
    v4 := make(map[uint32]int)
    var v8 float64 = v3
    v8 = v8
    if (9223372036854775807 != 128) {
        fmt.Println(v8)
        switch "a" {
        case "\n\t":
            v2 = (32767 < genFn1(0, -33))
        case "foo bar":
            var v9 uint16 = 1
            _ = v9
        case "a":
            switch "foo bar" {
            case "foo bar":
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := GenS1{F0: 31, F1: 415}
                    _ = fi10
                    _ = v11
                }
            case "hello":
                var v12 float32 = float32(1e10)
                _ = v12
            case "\n\t":
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 float32 = 1e10
                    fmt.Println(fi13, v2)
                    _ = fi13
                    _ = v14
                }
            default:
                v2 = v2
            }
        default:
            switch ("a" + "foo bar") {
            case "foo bar":
                var v15 int8 = 42
                _ = v15
            case "a":
                fmt.Println(v8, v8)
            }
        }
    } else {
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 bool = (-2.25 < -2.25)
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        if (!v2) {
            fmt.Println(v2, v2)
            _ = len(v1)
        }
        var v19 bool = (!(!(173 < 44)))
        _ = v19
    }
    v20 := GenS2{F0: -101}
    var v21 uint64 = 42
    var v22 int = 32767
    v8 = (v8 + (v8 / 1.0))
    for fi23 := 0; fi23 < 19; fi23++ {
        switch 2147483647 {
        case 63:
            v24 := GenS2{F0: -7}
            _ = v24
        case 127:
            for fi25 := 0; fi25 < 18; fi25++ {
                delete(v4, 531)
                _ = fi25
            }
        case -9223372036854775808:
            v2 = v2
        default:
            v26 := []int{127}
            _ = v26
        }
        fmt.Println(v21, v8)
        _ = fi23
    }
    v27 := GenS3{F0: ("" + "foo bar"), F1: 42}
    v28 := GenS4{F0: (-(v22)), F1: -929, F2: (v22 % (((v22 % (((v22 / ((v22) | 1))) | 1))) | 1))}
    v8 = 0.1
    _ = len(v4)
    if (v2 && v2) {
        v22 = (-(v22))
        for fi29 := 0; fi29 < 19; fi29++ {
            v22 = (-(v22))
            _ = fi29
        }
        v8 -= v8
    }
    var v30 uint8 = 255
    for fi31 := 0; fi31 < 11; fi31++ {
        var v32 int8 = -79
        v21 = (v21 ^ (v21 + 835))
        for fi33 := 0; fi33 < 4; fi33++ {
            fmt.Println(v3, v32, fi31)
            if v2 {
                var v34 int32 = 1
                if ((v30 + 171) > (v30 * (v30 & v30))) {
                    switch 2147483647 {
                    case -9223372036854775808:
                        v1[fi33] = ("hello" + "hello")
                    case 9223372036854775807:
                        var v35 uint16 = 255
                        _ = v35
                    }
                    v4[270] = v22
                } else {
                    v36 := v4[uint32(v21)]
                    switch ("foo bar" + ((("\n\t" + "hello") + "hello") + "a")) {
                    case "hello":
                        fmt.Println(v22, v30, v2)
                    case "":
                        var v37 uint64 = (v21 / ((1) | 1))
                        var v38 float64 = v8
                        _ = v37
                        _ = v38
                    case "a":
                        v39 := GenS5{F0: v21, F1: ("" + "hello"), F2: -32768}
                        _ = v39
                    default:
                        v40 := GenS2{F0: v32}
                        _ = v40
                    }
                    v30 = v30
                    _ = v36
                }
                var v41 int64 = -9223372036854775808
                var v42 int32 = (v34 % ((-1) | 1))
                _ = v34
                _ = v41
                _ = v42
            } else {
                if ((1 >= 42) || v2) {
                    v22 += (fi31 - 42)
                }
                v43 := []int{(fi33 % (((^(fi31))) | 1)), (-(v22)), -9223372036854775808}
                _ = len(v43)
                _ = v43
            }
            _ = fi33
        }
        _ = fi31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v8
}
