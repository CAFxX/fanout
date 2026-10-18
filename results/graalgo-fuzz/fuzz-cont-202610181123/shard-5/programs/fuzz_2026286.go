package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int64; F1 int; F2 uint }
type GenS6 struct { F0 uint8 }
func genFn0() int32 {
    var v4 uint16 = 65535
    var v5 string = "\n\t"
    var v6 uint64 = uint64(v4)
    _ = v4
    _ = v5
    _ = v6
    return -1
}
func genFn1() int {
    if false {
        var v7 string = ("" + "")
        var v8 uint64 = 4294967295
        v7 = ((v7 + v7) + ((v7 + v7) + (v7 + "hello")))
        _ = v7
        _ = v8
    } else {
        var v9 uint64 = 18446744073709551615
        if ((2 < 1) || (!true)) {
            v9 = (v9 ^ v9)
            var v10 uint = 255
            _ = v10
        }
        switch 377 {
        case -1:
            fmt.Println(v9)
        case -466:
            v11 := GenS1{F0: 2, F1: -9223372036854775808}
            _ = v11
        case 32768:
            v12 := []int{-1}
            _ = v12
        default:
            var v13 int8 = 42
            _ = v13
        }
        _ = v9
    }
    if true {
        var v14 int64 = 956
        var v15 int8 = -128
        _ = v14
        _ = v15
    }
    return -1
}

func main() {
    var v1 int16 = 1
    var v2 int16 = v1
    v3 := make(map[uint32]bool)
    v2 -= 1000
    v16 := GenS2{F0: 325, F1: 2}
    var v17 uint16 = 255
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS2{F0: 0, F1: 18446744073709551615}
        var v20 int = 32768
        if false {
            for fi21 := 0; fi21 < 16; fi21++ {
                v2 = 619
                v22 := []int{v20, genFn1()}
                var v23 int16 = v2
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 float64 = float64(((0.1 / float32(1e10)) * 1.5))
            _ = v24
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            v27 := GenS3{F0: 2147483647}
            _ = fi26
            _ = v27
        }
        fmt.Println(v17, v17)
        v1 += (-(v1))
        _ = fi25
    }
    v28 := GenS4{F0: false, F1: 2, F2: v17}
    var v29 uint64 = 162
    if (((false || (v17 > (v17 * 2))) && (!(!(-2147483648 == 42)))) && true) {
        for fi30 := 0; fi30 < 10; fi30++ {
            fmt.Println(v1)
            if ((fi30 >> ((18446744073709551615) & 63)) <= (fi30 - (fi30 / ((26) | 1)))) {
                fmt.Println(v17, fi30, v2)
            } else {
                switch "hello" {
                case "\n\t":
                    var v31 bool = true
                    _ = v31
                case "foo bar":
                    var v32 bool = (!true)
                    _ = v32
                case "hello":
                    if true {
                        fmt.Println(v17, v1, v29)
                    }
                }
                if (!false) {
                    v33 := GenS3{F0: -1}
                    if false {
                        for fi34 := 0; fi34 < 14; fi34++ {
                            delete(v3, 4294967295)
                            var v35 string = ((("a" + "\n\t") + ("hello" + (("" + "") + (("foo bar" + "") + ("a" + ("foo bar" + "foo bar")))))) + (("\n\t" + "\n\t") + "hello"))
                            _ = fi34
                            _ = v35
                        }
                        v36 := GenS5{F0: 416, F1: (^(fi30)), F2: 255}
                        _ = v36
                    }
                    var v37 int64 = 104
                    _ = v33
                    _ = v37
                }
            }
            _ = fi30
        }
        v38 := []int{32767, -9223372036854775808, 128, 523}
        v39 := GenS1{F0: 2147483648, F1: -417}
        _ = v38
        _ = v39
    }
    v3[832] = true
    for fi40 := 0; fi40 < 11; fi40++ {
        var v41 uint64 = (v29 + v29)
        var v42 int = (fi40 % ((fi40) | 1))
        _ = fi40
        _ = v41
        _ = v42
    }
    fmt.Println(v2, v17)
    v43 := GenS6{F0: 58}
    var v44 float64 = (0.0 + 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v43
    _ = v44
}
