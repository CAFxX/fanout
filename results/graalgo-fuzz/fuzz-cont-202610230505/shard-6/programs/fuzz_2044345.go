package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint; F2 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 uint; F1 bool; F2 uint }
func genFn0() int64 {
    var v5 int32 = 1
    switch "foo bar" {
    case "foo bar":
        var v6 uint32 = 2147483647
        _ = v6
    case "a":
        v5 += 0
    case "hello":
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: v5}
            _ = fi7
            _ = v8
        }
    default:
        v5 = (v5 * (v5 * v5))
    }
    var v9 int32 = v5
    switch 2147483647 {
    case -178:
        if false {
            v10 := GenS2{F0: 1000}
            v9 -= v9
            _ = v10
        } else {
            if false {
                switch "a" {
                case "hello":
                    var v11 uint = 2
                    _ = v11
                case "a":
                    v5 = 1
                default:
                    switch ("hello" + "hello") {
                    case "foo bar":
                        fmt.Println(v9)
                    case "hello":
                        for fi12 := 0; fi12 < 1; fi12++ {
                            v9 = v5
                            for fi13 := 0; fi13 < 18; fi13++ {
                                var v14 int64 = -489
                                _ = fi13
                                _ = v14
                            }
                            _ = fi12
                        }
                    default:
                        switch 42 {
                        case 4294967295:
                            var v15 int8 = 127
                            _ = v15
                        case -2147483648:
                            v9 += (v5 << ((817) & 63))
                        case -115:
                            var v16 uint16 = 1
                            _ = v16
                        }
                    }
                }
            }
        }
    case 450:
        var v17 int = 42
        _ = v17
    case 9223372036854775807:
        var v18 bool = false
        _ = v18
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 uint32 = 533
            _ = fi20
            _ = v21
        }
        var v22 uint = 833
        var v23 int64 = 750
        _ = fi19
        _ = v22
        _ = v23
    }
    _ = v5
    _ = v9
    return 248
}
func genFn1(p0 int64) uint {
    var v24 int8 = 42
    var v25 float64 = 0.1
    fmt.Println(p0)
    _ = p0
    _ = v24
    _ = v25
    return 860
}

func main() {
    var v1 string = (("" + ("\n\t" + ("a" + "a"))) + (("a" + ("a" + "a")) + "hello"))
    var v2 string = v1
    var v3 uint16 = 255
    var v4 float64 = ((((0.1 - float64((3.14159 - 1.5))) + 0.0) * 0.0) + 1.5)
    v4 += 1e10
    v26 := []int{}
    v3 = (^(v3))
    v3 = v3
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 int = 299
        switch v2 {
        case "\n\t":
            if (1 <= 218) {
                fmt.Println(v28, v1, v2)
                v26 = append(v26, -112)
                var v29 string = (v1 + v1)
                _ = v29
            }
        case "":
            switch v28 {
            case 9223372036854775807:
                var v30 uint32 = 15
                _ = v30
            default:
                v2 = (v2 + "\n\t")
            }
        case "a":
            for fi31 := 0; fi31 < 3; fi31++ {
                var v32 float64 = float64((1e10 / (1.5 - 1e10)))
                v33 := GenS3{F0: 775, F1: 1, F2: 2}
                var v34 uint8 = 2
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
        }
        _ = fi27
        _ = v28
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        if (!((1.5 != (v4 - v4)) && true)) {
            fmt.Println(fi35)
        } else {
            v26 = append(v26, fi35)
            v36 := GenS4{F0: (18446744073709551615 >= 1)}
            _ = v36
        }
        var v37 float32 = 1e10
        _ = fi35
        _ = v37
    }
    v38 := GenS2{F0: 32767}
    switch (v1 + v1) {
    case "foo bar":
        fmt.Println(v3, v4)
    default:
        var v39 string = v2
        _ = v39
    }
    var v40 uint = 18446744073709551615
    v26 = append(v26, 559)
    var v41 bool = (779 != 9223372036854775807)
    for fi42 := 0; fi42 < 10; fi42++ {
        v43 := GenS4{F0: (v4 == -2.25)}
        fmt.Println(v4, fi42)
        v26 = append(v26, (fi42 * fi42))
        _ = fi42
        _ = v43
    }
    switch "hello" {
    case "\n\t":
        v44 := GenS5{F0: 48, F1: 128}
        _ = v44
    case "":
        v40 += (v40 - v40)
    case "a":
        for fi45 := 0; fi45 < 2; fi45++ {
            var v46 int16 = 888
            _ = fi45
            _ = v46
        }
    default:
        var v47 uint8 = 42
        _ = v47
    }
    v2 = v1
    v48 := GenS6{F0: (^(v40)), F1: (!v41), F2: uint(v3)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v48
}
