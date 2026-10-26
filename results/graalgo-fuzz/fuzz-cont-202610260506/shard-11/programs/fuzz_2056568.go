package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint16 }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int; F1 uint8; F2 string }
type GenS5 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS6 struct { F0 int }
func genFn0(p0 int, p1 uint16, p2 int) uint16 {
    for fi6 := 0; fi6 < 9; fi6++ {
        if (!false) {
            p2 -= -914
            var v7 uint16 = 381
            if false {
                v8 := GenS1{F0: true, F1: "", F2: 733}
                v9 := GenS2{F0: 9223372036854775807, F1: (^(p2)), F2: (^(v7))}
                v10 := []int{9223372036854775807, (p2 + (fi6 >> ((4294967295) & 63)))}
                _ = v10
                _ = v8
                _ = v9
            }
            _ = v7
        } else {
            var v11 float64 = 1.5
            _ = v11
        }
        _ = fi6
    }
    var v12 int = (p0 % ((p2) | 1))
    var v13 uint64 = 410
    for fi14 := 0; fi14 < 10; fi14++ {
        fmt.Println(p1)
        p0 = (p0 + 1)
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return p1
}
func genFn1(p0 bool, p1 int8) uint8 {
    fmt.Println(p0, p0)
    var v15 uint64 = 247
    _ = p0
    _ = p1
    _ = v15
    return 195
}

func main() {
    var v1 uint32 = 656
    var v2 uint8 = 255
    v3 := []uint32{778}
    v4 := []string{"\n\t", ("foo bar" + "\n\t"), "\n\t", "\n\t"}
    var v5 uint64 = 927
    var v16 int = -2147483648
    if (1 != 11) {
        var v17 uint64 = 2
        _ = v17
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        fmt.Println(v1)
        _ = fi18
    }
    var v19 uint = uint(v1)
    v19 = 1
    var v20 int16 = 965
    if (1 != -695) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        fmt.Println(v1, v20)
        v20 = v20
    }
    v2 -= 255
    fmt.Println(v5, v20)
    if false {
        v5 = v5
        v21 := GenS3{F0: -1, F1: 4294967295}
        _ = v21
    } else {
        var v22 int16 = v20
        _ = len(v4)
        _ = v22
    }
    v19 = v19
    for fi23 := 0; fi23 < 17; fi23++ {
        fmt.Println(v5, v5)
        if (!false) {
            if (!true) {
                fmt.Println(v19, fi23)
                v3 = append(v3, v1)
            }
            switch "hello" {
            case "a":
                fmt.Println(v19, v2)
            case "foo bar":
                v20 = v20
            case "":
                switch ("hello" + ((("a" + ("\n\t" + ("\n\t" + (("a" + "\n\t") + ("" + (("\n\t" + "\n\t") + "foo bar")))))) + "foo bar") + ((("foo bar" + "\n\t") + ("foo bar" + ("" + "hello"))) + ""))) {
                case "\n\t":
                    var v24 int32 = -1
                    _ = v24
                }
            default:
                for fi25 := 0; fi25 < 2; fi25++ {
                    v2 = v2
                    if len(v3) > 0 {
                        v3[0] = uint32(v19)
                    }
                    if (false && (!false)) {
                        v5 += (v5 * 42)
                        var v26 uint64 = (v5 - (v5 << ((uint(v2)) & 63)))
                        if true {
                            v27 := GenS3{F0: 1, F1: v5}
                            v26 = (v26 - (v26 >> (((^(v19))) & 63)))
                            v28 := GenS4{F0: fi25, F1: genFn1((!true), 1), F2: ("foo bar" + ("hello" + ""))}
                            v29 := GenS2{F0: 0, F1: fi23, F2: 0}
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v30 := GenS5{F0: 65535, F1: v2, F2: -1}
                            var v31 uint = 710
                            _ = v30
                            _ = v31
                        }
                        _ = v26
                    }
                    _ = fi25
                }
            }
        }
        v32 := GenS6{F0: (v16 + (fi23 + fi23))}
        _ = fi23
        _ = v32
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v19, v20)
    var v33 uint8 = 128
    v20 = (v20 + (v20 | (v20 / (((^(v20))) | 1))))
    v34 := GenS6{F0: v16}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
