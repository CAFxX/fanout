package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 4294967295
    v3 := make(map[uint32]bool)
    var v4 int64 = -1
    _ = len(v3)
    var v5 string = ("" + (("a" + (("\n\t" + "foo bar") + "a")) + (("foo bar" + "\n\t") + "\n\t")))
    if true {
        var v6 int = 42
        _ = v6
    } else {
        for fi7 := 0; fi7 < 11; fi7++ {
            v8 := v3[699]
            v5 = ("\n\t" + "foo bar")
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 127, F1: "\n\t"}
        fmt.Println(v5)
        _ = v9
    }
    if false {
        switch int(v2) {
        case 249:
            v10 := v3[1]
            _ = v10
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            if (false && (42 < 0)) {
                var v12 string = "a"
                _ = v12
            } else {
                v13 := GenS1{F0: 20, F1: v5}
                fmt.Println(v5, fi11)
                _ = v13
            }
            v14 := []int{(fi11 >> ((0) & 63))}
            var v15 string = v5
            _ = fi11
            _ = v14
            _ = v15
        }
    } else {
        _ = len(v3)
        var v16 float64 = -2.25
        _ = v16
    }
    if (42 == 1) {
        if false {
            v17 := GenS1{F0: 42, F1: (v5 + ("" + "\n\t"))}
            _ = v17
        }
        v4 = 9223372036854775807
    } else {
        var v18 uint32 = 551
        var v19 int64 = (-(v4))
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 int8 = 42
        switch "a" {
        case "\n\t":
            if true {
                switch v5 {
                case "foo bar":
                    switch 9223372036854775807 {
                    case 32768:
                        if (v2 == (v2 | (v2 * v2))) {
                            for fi22 := 0; fi22 < 1; fi22++ {
                                v5 = ("foo bar" + v5)
                                v23 := GenS2{F0: -2147483648}
                                if (!(!((0.0 / 3.14159) < 1e10))) {
                                    switch v5 {
                                    case "a":
                                        if (!(3.14159 != 1.5)) {
                                            v24 := []int{}
                                            if len(v24) > 0 {
                                                v24[0] = 841
                                            }
                                            v3[2] = false
                                            _ = v24
                                        } else {
                                            var v25 float32 = 3.14159
                                            _ = v25
                                        }
                                    default:
                                        var v26 uint8 = 1
                                        _ = v26
                                    }
                                } else {
                                    _ = len(v3)
                                    v21 = (v21 + -128)
                                }
                                _ = fi22
                                _ = v23
                            }
                        }
                    case 127:
                        v21 = v21
                    case 697:
                        for fi27 := 0; fi27 < 1; fi27++ {
                            fmt.Println(v2)
                            v21 = (v21 + -1)
                            _ = fi27
                        }
                    default:
                        var v28 int64 = (-(v4))
                        _ = v28
                    }
                case "":
                    for fi29 := 0; fi29 < 1; fi29++ {
                        fmt.Println(v5, v4, fi20)
                        if (42 >= 127) {
                            if true {
                                v30 := GenS3{F0: 347, F1: 65535, F2: 32767}
                                v21 = v21
                                var v31 bool = true
                                v32 := []int{(fi20 * 42), 42, int(v2)}
                                _ = len(v3)
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            v33 := GenS1{F0: 42, F1: (v5 + v5)}
                            v2 = (v2 + (v2 >> (((^(v2))) & 63)))
                            _ = v33
                        } else {
                            v4 = v4
                        }
                        v5 = "foo bar"
                        v34 := v1[v5]
                        _ = fi29
                        _ = v34
                    }
                default:
                    var v35 int64 = (v4 | (v4 >> (((^(v2))) & 63)))
                    _ = v35
                }
                v36 := GenS4{F0: ("a" + v5)}
                _ = v36
            }
        default:
            fmt.Println(v2, v5)
        }
        _ = fi20
        _ = v21
    }
    var v37 uint32 = 65535
    v38 := v3[uint32(v2)]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
