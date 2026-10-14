package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint }
type GenS2 struct { F0 uint; F1 bool; F2 int64 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int8; F1 uint16; F2 string }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := []string{("hello" + ""), "foo bar"}
    var v4 uint8 = 220
    var v5 uint8 = 147
    v5 = (v5 - 1)
    if len(v3) > 0 {
        v3[0] = ("hello" + "")
    }
    v4 = 255
    switch "a" {
    case "":
        switch -9223372036854775808 {
        case 1:
            for fi6 := 0; fi6 < 5; fi6++ {
                fmt.Println(fi6, fi6)
                if (42 < 434) {
                    switch -737 {
                    case 2147483647:
                        if (1 == 4294967295) {
                            var v7 int8 = -128
                            _ = v7
                        } else {
                            var v8 uint16 = uint16(v5)
                            var v9 uint16 = v8
                            for fi10 := 0; fi10 < 12; fi10++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if (0 >= 2) {
                                    fmt.Println(v9, v5)
                                }
                                v11 := GenS1{F0: v4, F1: 32767, F2: 1}
                                _ = fi10
                                _ = v11
                            }
                            _ = v8
                            _ = v9
                        }
                    case -107:
                        if (1 == 1) {
                            var v12 uint32 = 637
                            _ = v12
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    default:
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                    }
                    var v13 uint = 923
                    _ = v13
                }
                _ = fi6
            }
        default:
            var v14 int = 32768
            _ = v14
        }
    }
    v2[42] = 2147483647
    v3 = append(v3, "foo bar")
    switch ("foo bar" + ("a" + (("foo bar" + "a") + "a"))) {
    case "\n\t":
        fmt.Println(v4, v4)
    default:
        _ = len(v3)
    }
    delete(v2, 713)
    var v15 int8 = -46
    v5 -= (v5 - (^(v4)))
    var v16 uint16 = 32767
    switch -2147483648 {
    case 4294967295:
        v15 += v15
    case 127:
        _ = len(v3)
    }
    v17 := GenS2{F0: 65535, F1: ((false && (!(991 > 2))) && (0 > -1)), F2: -9223372036854775808}
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := GenS3{F0: (v15 | v15), F1: fi18}
        if ((fi18 >> ((0) & 63)) < (-(fi18))) {
            v16 -= v16
        } else {
            var v20 uint16 = (^(v16))
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    if false {
        delete(v2, 1)
    }
    switch -2147483648 {
    case 134:
        for fi21 := 0; fi21 < 6; fi21++ {
            if false {
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 float32 = ((-2.25 / 1.0) * 1.5)
                    fmt.Println(v15, v16)
                    _ = len(v2)
                    _ = fi22
                    _ = v23
                }
                var v24 int16 = 1000
                _ = v24
            }
            fmt.Println(v5, v4)
            if (v4 != v4) {
                fmt.Println(v5)
                var v25 uint = 0
                _ = len(v3)
                _ = v25
            }
            _ = fi21
        }
    case 4294967295:
        var v26 uint8 = uint8(v16)
        _ = v26
    case 42:
        var v27 string = ""
        _ = v27
    }
    v28 := GenS4{F0: 0, F1: v16, F2: "a"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
