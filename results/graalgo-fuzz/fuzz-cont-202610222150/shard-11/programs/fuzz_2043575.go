package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32 }
func genFn0(p0 int, p1 uint16) string {
    switch p0 {
    case -9223372036854775808:
        fmt.Println(p0, p0)
    case -853:
        if (int8(p1) >= 1) {
            p0 += p0
            fmt.Println(p0)
        } else {
            if true {
                for fi7 := 0; fi7 < 14; fi7++ {
                    var v8 uint8 = 0
                    _ = fi7
                    _ = v8
                }
                var v9 int16 = -221
                var v10 int = p0
                _ = v10
                _ = v9
            } else {
                if (100 >= 42) {
                    v11 := GenS1{F0: 127, F1: false, F2: 111}
                    _ = v11
                }
                v12 := []int{(p0 & (p0 / ((p0) | 1))), p0}
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = v12
            }
        }
    case -369:
        p1 = p1
    }
    var v13 uint = 255
    for fi14 := 0; fi14 < 15; fi14++ {
        v13 = (v13 + 42)
        _ = fi14
    }
    v15 := GenS1{F0: uint8(v13), F1: false, F2: 255}
    var v16 string = ("" + ((("hello" + "a") + "a") + ("hello" + "")))
    _ = p0
    _ = p1
    _ = v13
    _ = v15
    _ = v16
    return ((v16 + (v16 + v16)) + (v16 + v16))
}
func genFn1(p0 uint16, p1 bool, p2 int32) uint64 {
    var v17 uint16 = p0
    v18 := []int{9223372036854775807, 0, -862}
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 uint = 40
        if len(v18) > 0 {
            v18[0] = (fi19 << (((v20 >> (((v20 - uint(v17))) & 63))) & 63))
        }
        fmt.Println(v17, p1)
        _ = fi19
        _ = v20
    }
    var v21 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v21
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 42
    v3 := make(map[uint32]bool)
    var v4 bool = true
    var v5 string = "foo bar"
    var v6 uint32 = 886
    v5 = v5
    var v22 uint16 = 65535
    v23 := []int{127}
    switch -1 {
    case 127:
        switch ((v5 + ("\n\t" + v5)) + "foo bar") {
        case "":
            var v24 int = 1
            _ = v24
        default:
            v25 := GenS2{F0: v2}
            _ = v25
        }
    default:
        switch v5 {
        case "foo bar":
            v26 := GenS1{F0: 1, F1: (-32768 > -32768), F2: 255}
            _ = v26
        case "\n\t":
            if v4 {
                switch (("hello" + (((v5 + v5) + (v5 + v5)) + "")) + "") {
                case "hello":
                    fmt.Println(v6, v2, v2)
                case "foo bar":
                    for fi27 := 0; fi27 < 7; fi27++ {
                        v4 = v4
                        for fi28 := 0; fi28 < 19; fi28++ {
                            switch "hello" {
                            case "foo bar":
                                var v29 float32 = 1e10
                                if v4 {
                                    if false {
                                        _ = len(v3)
                                        var v30 int32 = 1
                                        switch 0 {
                                        case 127:
                                            v31 := GenS3{F0: (^(v2)), F1: (v6 << ((18446744073709551615) & 63))}
                                            _ = v31
                                        }
                                        _ = v30
                                    }
                                }
                                var v32 bool = v4
                                _ = v29
                                _ = v32
                            case "\n\t":
                                v33 := GenS3{F0: genFn1(255, true, 573), F1: (^(v6))}
                                _ = v33
                            default:
                                var v34 int32 = -256
                                _ = v34
                            }
                            v5 = v5
                            _ = fi28
                        }
                        if (true && true) {
                            switch fi27 {
                            case -1:
                                var v35 float32 = 1e10
                                _ = v35
                            }
                            for fi36 := 0; fi36 < 6; fi36++ {
                                v2 = (v2 * genFn1(uint16(v2), ((1.5 + (0.1 / (((float64(float32(-2.25)) / float64(0.1)) - 0.0) - -2.25))) == float64(1e10)), 42))
                                _ = fi36
                            }
                            var v37 int8 = -88
                            _ = v37
                        }
                        _ = fi27
                    }
                case "":
                    v38 := GenS4{F0: 9223372036854775807}
                    _ = v38
                }
                v5 = v5
            }
        case "a":
            fmt.Println(v5, v6, v2)
        default:
            switch "a" {
            case "hello":
                fmt.Println(v4, v2)
            default:
                fmt.Println(v6, v2, v4)
            }
        }
    }
    var v39 float64 = 0.1
    var v40 uint32 = 4294967295
    if v4 {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
    } else {
        v4 = v4
        v2 = (v2 - (v2 + (v2 << ((325) & 63))))
        var v41 int8 = 41
        _ = v41
    }
    v42 := GenS5{F0: 1}
    fmt.Println(v5, v6, v4)
    fmt.Println(v39, v39, v2)
    var v43 uint16 = v22
    var v44 int8 = 100
    if v4 {
        var v45 int8 = (v44 >> ((911) & 63))
        var v46 uint16 = (v22 & v22)
        _ = v45
        _ = v46
    }
    var v47 uint8 = 253
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v44
    _ = v47
    _ = v5
    _ = v6
}
