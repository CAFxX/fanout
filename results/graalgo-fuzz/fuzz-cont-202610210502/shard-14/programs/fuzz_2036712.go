package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint32; F1 int8 }
func genFn0() uint8 {
    switch 128 {
    case -2147483648:
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int8 = -128
            var v9 bool = false
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    if true {
        switch ("hello" + "hello") {
        case "a":
            var v10 float64 = -2.25
            _ = v10
        case "":
            v11 := []int{9223372036854775807, -13, 9223372036854775807, 991}
            _ = v11
        }
    }
    var v12 int8 = 1
    v12 = v12
    _ = v12
    return 128
}

func main() {
    v1 := []int{32768, 555, 42}
    v2 := make(map[string]bool)
    var v3 uint8 = 1
    v4 := []string{(("\n\t" + (("foo bar" + "\n\t") + "")) + "a"), ("hello" + "\n\t")}
    var v5 uint16 = 1
    v6 := []uint32{377, 53, 0, 2147483647}
    v13 := GenS1{F0: 32767, F1: 0, F2: 18446744073709551615}
    for fi14 := 0; fi14 < 15; fi14++ {
        delete(v2, "foo bar")
        v1 = append(v1, fi14)
        v5 -= (v5 & 255)
        _ = fi14
    }
    var v15 int16 = 42
    switch "\n\t" {
    case "foo bar":
        var v16 int32 = -1
        _ = v16
    case "\n\t":
        v17 := GenS2{F0: 65535, F1: 127}
        _ = v17
    }
    var v18 int64 = int64(v5)
    if (int8(v18) == -128) {
        v5 = (v5 / (((v5 >> ((1) & 63))) | 1))
        v18 -= -1
    }
    switch int(v15) {
    case 128:
        v19 := GenS2{F0: 2147483648, F1: 42}
        _ = v19
    case 42:
        for fi20 := 0; fi20 < 14; fi20++ {
            for fi21 := 0; fi21 < 11; fi21++ {
                var v22 int8 = 127
                if false {
                    var v23 uint64 = 18446744073709551615
                    if len(v4) > 0 {
                        v4[0] = ("hello" + "hello")
                    }
                    _ = v23
                }
                v24 := GenS2{F0: 2147483648, F1: (v22 & v22)}
                _ = fi21
                _ = v22
                _ = v24
            }
            _ = fi20
        }
    }
    if false {
        var v25 int16 = (v15 + 32767)
        if (18446744073709551615 < 82) {
            v26 := GenS2{F0: uint32(v5), F1: 100}
            if (!true) {
                var v27 float64 = -2.25
                if ((2147483647 > 1) && false) {
                    for fi28 := 0; fi28 < 7; fi28++ {
                        if true {
                            v29 := v2[""]
                            var v30 uint64 = 1
                            var v31 uint64 = (v30 % ((v30) | 1))
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        v5 += (v5 % (((v5 | v5)) | 1))
                        _ = fi28
                    }
                    v3 += genFn0()
                }
                _ = len(v6)
                _ = v27
            } else {
                var v32 bool = true
                _ = v32
            }
            _ = v26
        } else {
            var v33 float32 = 3.14159
            for fi34 := 0; fi34 < 18; fi34++ {
                for fi35 := 0; fi35 < 15; fi35++ {
                    if false {
                        v15 = -669
                    }
                    _ = fi35
                }
                _ = fi34
            }
            _ = v33
        }
        _ = v25
    } else {
        v36 := GenS3{F0: 4294967295, F1: 93}
        var v37 int8 = 42
        var v38 uint = 574
        _ = v36
        _ = v37
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
