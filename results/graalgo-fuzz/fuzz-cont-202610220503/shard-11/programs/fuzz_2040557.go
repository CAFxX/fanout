package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint; F1 int8; F2 int64 }
type GenS4 struct { F0 uint; F1 int64; F2 uint32 }
type GenS5 struct { F0 bool; F1 bool; F2 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 string; F2 uint8 }
type GenS8 struct { F0 string }
func genFn0(p0 int16, p1 uint16) int64 {
    fmt.Println(p1)
    p1 -= p1
    var v6 uint64 = 0
    if (!true) {
        v7 := GenS1{F0: p0, F1: 127}
        if (-1 <= 42) {
            fmt.Println(p1, p1)
            var v8 uint8 = 141
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := GenS2{F0: v8, F1: v8}
                p0 += int16(fi9)
                v6 += v6
                _ = fi9
                _ = v10
            }
            _ = v8
        } else {
            var v11 int64 = -972
            _ = v11
        }
        _ = v7
    } else {
        var v12 int8 = 0
        p1 += p1
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v6
    return -9223372036854775808
}
func genFn1(p0 int64, p1 int, p2 uint) string {
    p1 = (p1 & p1)
    if (!(386 >= 390)) {
        v13 := GenS1{F0: -941, F1: 127}
        var v14 bool = false
        _ = v13
        _ = v14
    } else {
        fmt.Println(p1, p1, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    var v3 bool = (32768 == 2147483647)
    v4 := []int{42, -451, 602}
    v5 := []int{1}
    fmt.Println(v3)
    fmt.Println(v3)
    v4 = append(v4, -523)
    v15 := GenS3{F0: 983, F1: 42, F2: -9223372036854775808}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3)
    var v16 uint32 = 2147483648
    _ = len(v5)
    _ = len(v1)
    fmt.Println(v3)
    var v17 uint = uint(v16)
    fmt.Println(v3, v16, v17)
    switch "" {
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            v18 := GenS3{F0: (^(v17)), F1: 1, F2: -9223372036854775808}
            _ = v18
        case "hello":
            v3 = v3
        case "a":
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v3)
            case "":
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 string = ("a" + ("\n\t" + "foo bar"))
                    var v21 int8 = 0
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        }
    case "":
        v22 := GenS4{F0: (v17 + v17), F1: genFn0(-101, 768), F2: v16}
        _ = v22
    case "a":
        var v23 uint8 = 233
        _ = v23
    }
    v17 = (v17 & v17)
    switch "\n\t" {
    case "hello":
        for fi24 := 0; fi24 < 14; fi24++ {
            for fi25 := 0; fi25 < 18; fi25++ {
                v16 -= v16
                v5 = append(v5, (fi25 + -298))
                if v3 {
                    var v26 uint64 = 10
                    switch "" {
                    case "hello":
                        var v27 uint16 = 237
                        _ = v27
                    case "foo bar":
                        if (v3 || (-128 != 42)) {
                            var v28 int64 = genFn0(659, 1)
                            v28 = (-(v28))
                            _ = v28
                        }
                    case "\n\t":
                        if v3 {
                            v3 = (!v3)
                            var v29 bool = v3
                            v30 := GenS5{F0: (135 != 128), F1: (!((v16 - v16) == (v16 >> ((10) & 63)))), F2: v16}
                            if (false || v29) {
                                v31 := GenS6{F0: uint32(v17)}
                                _ = v31
                            } else {
                                v26 -= v26
                                if true {
                                    v29 = (!true)
                                    _ = len(v2)
                                }
                                fmt.Println(v17, fi24)
                            }
                            _ = v29
                            _ = v30
                        }
                    default:
                        if (1.5 == -2.25) {
                            switch fi25 {
                            case 9223372036854775807:
                                switch ("a" + "\n\t") {
                                case "hello":
                                    var v32 uint = v17
                                    _ = v32
                                case "":
                                    var v33 uint32 = (v16 >> (((^(v17))) & 63))
                                    _ = v33
                                }
                            }
                        } else {
                            if v3 {
                                var v34 bool = v3
                                _ = v34
                            } else {
                                var v35 bool = (!v3)
                                v3 = (((float32(3.14159) / 0.1) - float32(-2.25)) < 0.0)
                                _ = v35
                            }
                            switch "hello" {
                            case "":
                                fmt.Println(fi24, v17, v3)
                            case "\n\t":
                                var v36 bool = (v3 || (!(0 <= 0)))
                                _ = v36
                            case "a":
                                v37 := GenS7{F0: v16, F1: (((("foo bar" + "foo bar") + ("\n\t" + "foo bar")) + "") + "foo bar"), F2: 255}
                                var v38 int = 32767
                                v39 := GenS3{F0: 42, F1: 127, F2: 558}
                                if false {
                                    v26 = v26
                                    var v40 uint8 = 107
                                    _ = v40
                                }
                                _ = v37
                                _ = v38
                                _ = v39
                            default:
                                v17 = 61
                            }
                        }
                    }
                    _ = v26
                } else {
                    v17 = (v17 << ((v17) & 63))
                    v41 := GenS8{F0: "hello"}
                    var v42 int16 = 0
                    _ = v41
                    _ = v42
                }
                _ = fi25
            }
            _ = fi24
        }
    case "a":
        for fi43 := 0; fi43 < 9; fi43++ {
            for fi44 := 0; fi44 < 12; fi44++ {
                v5 = append(v5, 127)
                var v45 int = fi44
                _ = fi44
                _ = v45
            }
            _ = fi43
        }
    case "\n\t":
        v3 = (-32768 == 1000)
    }
    v16 -= (v16 * v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
