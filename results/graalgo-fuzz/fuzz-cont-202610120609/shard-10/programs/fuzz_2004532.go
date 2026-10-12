package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int32 }
type GenS2 struct { F0 int16; F1 int16; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int8; F1 int; F2 int32 }
type GenS6 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 int; F1 int8; F2 int64 }
type GenS9 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint, p1 uint32, p2 uint16) int {
    v5 := GenS1{F0: 123, F1: (!(42 == 0)), F2: 42}
    v6 := GenS2{F0: -1, F1: -58, F2: 42}
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 uint32 = p1
            v10 := []int{572}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    p2 -= 42
    p2 = (p2 << ((p0) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 32768
}

func main() {
    v1 := []string{("a" + "foo bar"), "foo bar", "a"}
    v2 := make(map[uint32]bool)
    var v3 bool = (false && (!(2 > 762)))
    var v4 int = -762
    for fi11 := 0; fi11 < 16; fi11++ {
        v3 = (-2147483648 == -2147483648)
        _ = fi11
    }
    var v12 uint = 18446744073709551615
    v13 := GenS1{F0: int8(v4), F1: (!v3), F2: -650}
    fmt.Println(v12, v3, v4)
    v4 += v4
    v14 := GenS2{F0: -827, F1: 361, F2: -924}
    v3 = v3
    if v3 {
        for fi15 := 0; fi15 < 6; fi15++ {
            for fi16 := 0; fi16 < 1; fi16++ {
                v12 = 633
                v4 -= (v4 - 127)
                if ((^(v12)) != v12) {
                    switch "hello" {
                    case "\n\t":
                        var v17 int8 = 0
                        _ = len(v2)
                        _ = v17
                    default:
                        v18 := GenS2{F0: 223, F1: 32767, F2: -474}
                        _ = v18
                    }
                    v4 = fi15
                }
                v19 := GenS3{F0: 42}
                _ = fi16
                _ = v19
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(fi15, v12)
            _ = fi15
        }
    }
    v20 := GenS4{F0: v4, F1: 42}
    switch "\n\t" {
    case "a":
        fmt.Println(v3)
    case "":
        v1 = append(v1, "a")
    case "\n\t":
        var v21 float64 = 3.14159
        _ = v21
    default:
        var v22 uint32 = 733
        _ = v22
    }
    v23 := GenS3{F0: 42}
    if v3 {
        v24 := GenS4{F0: (v4 / ((v4) | 1)), F1: 57}
        switch -147 {
        case 4294967295:
            for fi25 := 0; fi25 < 10; fi25++ {
                switch v4 {
                case 647:
                    v12 = v12
                case 127:
                    _ = len(v1)
                case 32768:
                    var v26 uint32 = uint32(v12)
                    _ = v26
                }
                _ = fi25
            }
        }
        var v27 uint8 = 128
        _ = v24
        _ = v27
    } else {
        v28 := GenS5{F0: -128, F1: genFn0(v12, 2147483647, 42), F2: -140}
        if true {
            if v3 {
                var v29 int = -894
                v30 := GenS6{F0: (v3 && true), F1: 255, F2: 0}
                _ = v29
                _ = v30
            }
            if v3 {
                for fi31 := 0; fi31 < 18; fi31++ {
                    var v32 string = ("\n\t" + "a")
                    v4 = fi31
                    _ = fi31
                    _ = v32
                }
            }
            v33 := GenS4{F0: (v4 % ((v4) | 1)), F1: 0}
            _ = v33
        } else {
            v34 := GenS4{F0: (v4 | v4), F1: 58}
            _ = v34
        }
        _ = v28
    }
    var v35 float64 = (0.0 * float64(-2.25))
    for fi36 := 0; fi36 < 2; fi36++ {
        var v37 uint64 = 0
        var v38 int = (fi36 >> ((v12) & 63))
        if v3 {
            v35 = 0.0
            switch ("hello" + (("" + ("foo bar" + "")) + "hello")) {
            case "\n\t":
                var v39 int8 = 127
                _ = v39
            default:
                var v40 uint32 = 347
                _ = v40
            }
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    v4 = v4
    v3 = (1.5 != 3.14159)
    v4 -= (v4 | (v4 << (((v12 | (^(v12)))) & 63)))
    switch ("" + (("" + "hello") + (("" + "a") + "a"))) {
    case "foo bar":
        for fi41 := 0; fi41 < 8; fi41++ {
            var v42 int16 = 42
            var v43 uint32 = 4294967295
            for fi44 := 0; fi44 < 15; fi44++ {
                var v45 uint32 = (^(v43))
                v12 = (^(v12))
                v46 := GenS7{F0: 0, F1: 255}
                var v47 int16 = (v42 % ((v42) | 1))
                if (!v3) {
                    v48 := GenS8{F0: int(v47), F1: int8(fi44), F2: -22}
                    switch "" {
                    case "hello":
                        fmt.Println(v4, v47, v3)
                    case "foo bar":
                        var v49 uint32 = v43
                        _ = v49
                    case "":
                        if true {
                            if ((v42 / ((v47) | 1)) < (v47 / ((1) | 1))) {
                                var v50 float64 = v35
                                fmt.Println(v43, v42)
                                if v3 {
                                    var v51 int64 = 42
                                    v1 = append(v1, ("" + "a"))
                                    _ = v51
                                }
                                if (v3 || (false && (-1 > -375))) {
                                    switch genFn0(v12, v43, 0) {
                                    case -2147483648:
                                        delete(v2, 4294967295)
                                    }
                                }
                                var v52 uint32 = (v43 >> ((v12) & 63))
                                _ = v50
                                _ = v52
                            }
                            v43 = 915
                        }
                    default:
                        v4 = fi41
                    }
                    if (v3 && (101 < 0)) {
                        v12 = v12
                    } else {
                        v53 := GenS9{F0: 18446744073709551615, F1: 919}
                        fmt.Println(v43)
                        if v3 {
                            var v54 int32 = 42
                            fmt.Println(fi44)
                            _ = v54
                        } else {
                            delete(v2, 2147483647)
                            var v55 float64 = 0.1
                            _ = v55
                        }
                        _ = v53
                    }
                    v43 += v45
                    v56 := v2[v45]
                    _ = v48
                    _ = v56
                }
                _ = fi44
                _ = v45
                _ = v46
                _ = v47
            }
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v35
    _ = v4
}
