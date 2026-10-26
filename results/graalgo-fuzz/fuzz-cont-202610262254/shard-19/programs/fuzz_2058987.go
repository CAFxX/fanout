package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int16 }
type GenS2 struct { F0 int8; F1 int32; F2 int16 }
func genFn0(p0 int8, p1 int64) uint16 {
    if false {
        var v5 int8 = p0
        _ = v5
    } else {
        for fi6 := 0; fi6 < 14; fi6++ {
            switch fi6 {
            case -217:
                v7 := []int{-1}
                _ = v7
            case 128:
                for fi8 := 0; fi8 < 13; fi8++ {
                    var v9 float32 = 3.14159
                    v10 := []int{-196, (fi6 * -1)}
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case -610:
                if (65535 < 4294967295) {
                    p1 = p1
                    var v11 int64 = -1
                    p0 -= (p0 ^ (^(p0)))
                    _ = v11
                }
            }
            var v12 int64 = -9223372036854775808
            _ = fi6
            _ = v12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 uint16 = 65535
            switch "\n\t" {
            case "foo bar":
                var v15 uint16 = (v14 + v14)
                _ = v15
            case "a":
                if (18446744073709551615 >= 42) {
                    p0 = p0
                }
            }
            fmt.Println(fi13, p0, p1)
            _ = fi13
            _ = v14
        }
        switch -2147483648 {
        case 32767:
            p0 += p0
        case -312:
            switch 32767 {
            case 32767:
                v16 := []int{32768, 32767, 32767, -1}
                _ = v16
            case -9223372036854775808:
                switch "a" {
                case "a":
                    var v17 uint = 2
                    _ = v17
                case "hello":
                    p1 -= (^(p1))
                case "":
                    v18 := GenS1{F0: -2147483648, F1: -2147483648, F2: -1}
                    _ = v18
                }
            default:
                switch "hello" {
                case "foo bar":
                    if true {
                        switch ("a" + "foo bar") {
                        case "a":
                            for fi19 := 0; fi19 < 11; fi19++ {
                                for fi20 := 0; fi20 < 9; fi20++ {
                                    v21 := []int{1, (fi19 % ((0) | 1)), fi20}
                                    _ = fi20
                                    _ = v21
                                }
                                _ = fi19
                            }
                        case "hello":
                            var v22 uint32 = 809
                            _ = v22
                        default:
                            fmt.Println(p1)
                        }
                        switch ("a" + ("\n\t" + ((("hello" + "") + "\n\t") + ("\n\t" + "hello")))) {
                        case "hello":
                            var v23 int8 = (p0 << ((428) & 63))
                            _ = v23
                        case "\n\t":
                            v24 := GenS1{F0: int32(p0), F1: 9223372036854775807, F2: -843}
                            _ = v24
                        case "a":
                            fmt.Println(p0, p1)
                        }
                        var v25 int16 = -131
                        _ = v25
                    }
                }
            }
        case 128:
            var v26 uint16 = 2
            _ = v26
        }
    }
    var v27 int64 = p1
    switch (("\n\t" + "") + "hello") {
    case "":
        var v28 float32 = 0.0
        _ = v28
    case "foo bar":
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 int32 = 2147483647
            _ = fi29
            _ = v30
        }
    default:
        var v31 string = ("hello" + "\n\t")
        _ = v31
    }
    _ = p0
    _ = p1
    _ = v27
    return 89
}
func genFn1() int16 {
    var v32 uint64 = 2
    var v33 string = "hello"
    v32 -= v32
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 uint32 = 65535
        _ = fi34
        _ = v35
    }
    var v36 float32 = 1e10
    _ = v32
    _ = v33
    _ = v36
    return 1
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]bool)
    var v3 bool = (!(535 != 478))
    v4 := []uint32{2147483648, 704}
    fmt.Println(v1)
    var v37 int64 = 9223372036854775807
    if true {
        var v38 float32 = (0.0 + ((float32(3.14159) * (1.5 * 0.1)) - 0.0))
        fmt.Println(v3)
        switch ("" + ("" + "foo bar")) {
        case "foo bar":
            v1 += 860
        case "":
            var v39 float32 = 1e10
            _ = v39
        }
        _ = v38
    } else {
        v37 = (v37 * v37)
        var v40 uint32 = 2
        _ = v40
    }
    switch ("" + "hello") {
    case "hello":
        for fi41 := 0; fi41 < 13; fi41++ {
            if v3 {
                var v42 uint64 = (v1 / ((1) | 1))
                var v43 float32 = 1.5
                _ = v42
                _ = v43
            } else {
                v3 = ((fi41 << ((300) & 63)) > fi41)
                if v3 {
                    v37 = v37
                }
                var v44 float32 = (0.0 / ((0.1 * 1.5) / 1.5))
                _ = v44
            }
            _ = fi41
        }
    case "a":
        delete(v2, 2147483647)
    case "\n\t":
        v45 := GenS2{F0: 110, F1: 42, F2: genFn1()}
        _ = v45
    }
    v37 = (v37 >> ((uint(v1)) & 63))
    v1 -= 2
    v46 := GenS1{F0: 42, F1: -758, F2: genFn1()}
    v37 = v37
    var v47 int8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v46
    _ = v47
}
