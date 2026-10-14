package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 string; F1 int }

func main() {
    var v1 uint8 = 1
    var v2 bool = (0 != -497)
    var v3 int16 = 95
    v4 := make(map[int]bool)
    var v5 uint = 112
    v6 := GenS1{F0: -1, F1: 65535, F2: 262}
    if true {
        var v7 bool = v2
        if v2 {
            for fi8 := 0; fi8 < 5; fi8++ {
                for fi9 := 0; fi9 < 9; fi9++ {
                    _ = len(v4)
                    v3 = v3
                    _ = fi9
                }
                for fi10 := 0; fi10 < 12; fi10++ {
                    if v2 {
                        v7 = v2
                        var v11 string = ""
                        _ = v11
                    }
                    _ = fi10
                }
                fmt.Println(v5, fi8)
                _ = fi8
            }
        } else {
            v3 = -57
        }
        var v12 int32 = -546
        _ = v12
        _ = v7
    }
    switch 2147483647 {
    case 4294967295:
        v13 := GenS2{F0: (!true), F1: 284}
        _ = v13
    case -173:
        switch -9223372036854775808 {
        case 718:
            fmt.Println(v1, v3, v3)
        }
    case 128:
        fmt.Println(v2, v1)
    default:
        switch 32768 {
        case 718:
            for fi14 := 0; fi14 < 11; fi14++ {
                fmt.Println(v3, v1)
                v3 = 0
                var v15 uint8 = 1
                _ = fi14
                _ = v15
            }
        case 965:
            var v16 float64 = 0.0
            _ = v16
        default:
            v3 = v3
        }
    }
    var v17 uint = v5
    v1 -= (v1 * 42)
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 int8 = 115
        _ = fi18
        _ = v19
    }
    v3 = (v3 + v3)
    switch 32767 {
    case 2147483647:
        v20 := GenS1{F0: v3, F1: 412, F2: 279}
        _ = v20
    }
    v21 := GenS3{F0: -121, F1: 964}
    for fi22 := 0; fi22 < 17; fi22++ {
        v4[fi22] = (!(!true))
        if v2 {
            for fi23 := 0; fi23 < 12; fi23++ {
                v1 = (v1 ^ v1)
                v24 := GenS1{F0: v3, F1: 1, F2: 2147483647}
                v25 := []int{127, fi23}
                var v26 uint16 = 65535
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            v17 = (^(v17))
        } else {
            fmt.Println(fi22, fi22)
            v3 = v3
            if v2 {
                if ((73 >= v1) || false) {
                    v27 := []int{(fi22 * 183), (-(fi22)), 357, fi22}
                    v5 -= v17
                    var v28 int = (-(fi22))
                    _ = v27
                    _ = v28
                } else {
                    if v2 {
                        var v29 bool = v2
                        _ = v29
                    } else {
                        if false {
                            for fi30 := 0; fi30 < 17; fi30++ {
                                v2 = false
                                var v31 uint64 = 18446744073709551615
                                var v32 int8 = 127
                                var v33 uint64 = v31
                                _ = fi30
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            var v34 float32 = -2.25
                            _ = v34
                        }
                    }
                    var v35 float32 = ((3.14159 * (1.5 + 1.5)) * (0.0 + 3.14159))
                    v3 += (v3 ^ int16(fi22))
                    _ = v35
                }
                if v2 {
                    v36 := GenS4{F0: -1, F1: v2}
                    v37 := GenS5{F0: ("" + ("" + "\n\t")), F1: fi22}
                    _ = v36
                    _ = v37
                }
            } else {
                var v38 uint64 = 18446744073709551615
                _ = v38
            }
        }
        var v39 int = fi22
        _ = fi22
        _ = v39
    }
    switch (((((("\n\t" + ("hello" + "a")) + (("hello" + "foo bar") + "foo bar")) + "\n\t") + "a") + (((("" + "\n\t") + ("a" + (("foo bar" + "") + "\n\t"))) + ("foo bar" + ("" + ("" + "foo bar")))) + "")) + "\n\t") {
    case "\n\t":
        var v40 float64 = (3.14159 / (3.14159 / 1e10))
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
