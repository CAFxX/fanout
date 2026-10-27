package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint64; F1 int64; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint }
func genFn0(p0 int64) string {
    p0 = p0
    if (42 > 0) {
        if false {
            v5 := []int{}
            v6 := GenS1{F0: -2147483648, F1: 65, F2: 2147483647}
            _ = v5
            _ = v6
        }
        p0 = p0
        fmt.Println(p0)
    }
    var v7 int64 = (p0 & (p0 * (^(p0))))
    _ = p0
    _ = v7
    return "a"
}
func genFn1(p0 int64, p1 uint64) uint16 {
    fmt.Println(p1)
    var v8 int32 = -1
    fmt.Println(v8, v8)
    _ = p0
    _ = p1
    _ = v8
    return uint16(p1)
}

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    var v3 uint64 = 18446744073709551615
    var v4 bool = (128 >= 2)
    var v9 uint = 18446744073709551615
    for fi10 := 0; fi10 < 6; fi10++ {
        v1 = ("hello" <= "a")
        v11 := GenS2{F0: uint16(v9), F1: uint8(v3)}
        var v12 int64 = 386
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        switch fi13 {
        case 60:
            for fi14 := 0; fi14 < 18; fi14++ {
                fmt.Println(v1, fi13)
                var v15 uint8 = 127
                v9 = (v9 + (v9 ^ v9))
                switch fi13 {
                case 2147483647:
                    var v16 float32 = 0.1
                    fmt.Println(v3)
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
        case -9223372036854775808:
            var v17 int16 = 346
            _ = v17
        case 2147483647:
            switch ((("a" + "\n\t") + "") + "\n\t") {
            case "a":
                v9 += v9
            case "hello":
                for fi18 := 0; fi18 < 2; fi18++ {
                    var v19 int = -9223372036854775808
                    v1 = (!((v4 || v1) && false))
                    v20 := GenS3{F0: 100, F1: false}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case "":
                v3 = v3
            default:
                var v21 uint = (v9 % ((v9) | 1))
                _ = v21
            }
        default:
            var v22 float32 = (1.5 / float32((((3.14159 * ((1.5 / (3.14159 * float64(0.1))) / 1.0)) / 1e10) * 3.14159)))
            _ = v22
        }
        if v4 {
            var v23 int64 = 850
            for fi24 := 0; fi24 < 6; fi24++ {
                if v4 {
                    fmt.Println(v4, v9)
                    var v25 uint64 = v3
                    _ = v25
                }
                _ = fi24
            }
            _ = v23
        }
        switch ("a" + "") {
        case "":
            if (genFn1(1, v3) >= 92) {
                switch 9223372036854775807 {
                case -250:
                    v26 := GenS4{F0: v3, F1: -9223372036854775808, F2: -53}
                    _ = v26
                case 278:
                    if v1 {
                        for fi27 := 0; fi27 < 3; fi27++ {
                            if (1 >= 2) {
                                fmt.Println(v3, v4)
                            }
                            var v28 uint16 = genFn1(9223372036854775807, v3)
                            switch "foo bar" {
                            case "":
                                v4 = v1
                            }
                            _ = fi27
                            _ = v28
                        }
                        var v29 int = fi13
                        _ = v29
                    }
                case 724:
                    if v4 {
                        v30 := GenS5{F0: uint16(v3)}
                        _ = v30
                    }
                default:
                    if (v1 && true) {
                        var v31 float64 = (3.14159 + 1.5)
                        v4 = false
                        _ = v31
                    } else {
                        switch "a" {
                        case "":
                            switch fi13 {
                            case -1:
                                var v32 string = (("hello" + "a") + "foo bar")
                                _ = v32
                            case -9223372036854775808:
                                v3 += (v3 ^ v3)
                            case 1:
                                var v33 float64 = (0.0 * 3.14159)
                                _ = v33
                            }
                        case "hello":
                            if (32767 >= 42) {
                                v34 := GenS2{F0: 0, F1: 2}
                                v35 := []int{(^(fi13)), (fi13 & 0)}
                                for fi36 := 0; fi36 < 5; fi36++ {
                                    switch 9223372036854775807 {
                                    case -194:
                                        var v37 float32 = float32((1.5 * float64((float32(1e10) - 3.14159))))
                                        _ = v37
                                    case 42:
                                        var v38 int32 = 42
                                        _ = v38
                                    default:
                                        v4 = (-1 != -802)
                                        fmt.Println(fi36)
                                        var v39 int = -362
                                        _ = v39
                                    }
                                    switch ("a" + "hello") {
                                    case "a":
                                        v40 := v2[fi13]
                                        _ = v40
                                    default:
                                        v3 = 0
                                    }
                                    var v41 uint32 = uint32(v9)
                                    _ = fi36
                                    _ = v41
                                }
                                _ = v34
                                _ = v35
                            }
                        }
                    }
                }
                var v42 int16 = -32768
                _ = v42
            } else {
                fmt.Println(v3)
                v2[32767] = (fi13 >> ((0) & 63))
            }
        case "foo bar":
            var v43 int16 = 1
            _ = v43
        case "\n\t":
            fmt.Println(v4, v1)
        }
        _ = fi13
    }
    switch -648 {
    case -2147483648:
        switch -1 {
        case 32767:
            var v44 float32 = 3.14159
            _ = v44
        case 973:
            var v45 uint16 = genFn1(-1, 1)
            _ = v45
        }
    case 32768:
        v4 = false
    default:
        v1 = (v4 && true)
    }
    v1 = (128 >= 2)
    v3 = v3
    fmt.Println(v4)
    if (0.1 >= 0.0) {
        switch ("hello" + ("a" + "a")) {
        case "\n\t":
            v46 := GenS6{F0: 4294967295}
            _ = v46
        case "":
            v1 = ((true || true) && v4)
        case "foo bar":
            v47 := []int{-548, -448, 127, 127}
            _ = v47
        }
        v1 = v1
    } else {
        v9 = 787
    }
    v48 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v9
}
