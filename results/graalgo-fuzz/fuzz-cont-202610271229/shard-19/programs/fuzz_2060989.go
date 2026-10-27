package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 string; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64; F1 uint64; F2 int32 }
func genFn0(p0 uint32, p1 int16) uint {
    v5 := []int{}
    if (!(-2147483648 >= 0)) {
        var v6 int64 = 0
        p1 += p1
        _ = v6
    }
    p0 = 1
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1(p0 uint, p1 uint16) int {
    var v7 string = (((("foo bar" + "\n\t") + ("hello" + ("" + ""))) + "\n\t") + "hello")
    var v8 float32 = 0.1
    if (-833 == 0) {
        var v9 int16 = -1
        var v10 string = ((((v7 + "a") + v7) + (("hello" + ((v7 + v7) + (v7 + v7))) + (v7 + "a"))) + (v7 + ("a" + "foo bar")))
        var v11 int64 = 0
        _ = v10
        _ = v11
        _ = v9
    } else {
        if true {
            if ((699 == uint64(p1)) && (false || true)) {
                switch ("hello" + v7) {
                case "foo bar":
                    fmt.Println(p0)
                case "\n\t":
                    var v12 uint64 = 42
                    _ = v12
                case "":
                    p0 -= genFn0(173, 42)
                default:
                    fmt.Println(p1, p1, p1)
                }
                var v13 uint32 = 514
                _ = v13
            }
            switch v7 {
            case "hello":
                var v14 int16 = 1000
                _ = v14
            default:
                v15 := GenS1{F0: 0}
                _ = v15
            }
            v16 := []int{9223372036854775807, 0, -1, 9223372036854775807}
            _ = v16
        }
    }
    var v17 uint8 = 80
    v18 := GenS2{F0: true, F1: "\n\t", F2: ((1.5 >= (0.0 / 1.5)) || (!true))}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 108
    var v3 bool = (true && (!true))
    v4 := []uint32{2147483647}
    v19 := GenS3{F0: -9223372036854775808}
    switch 2147483647 {
    case 127:
        v4 = append(v4, 2147483647)
    case -497:
        v2 -= (^(v2))
    case -2147483648:
        if (-501 >= 0) {
            var v20 string = ("a" + "hello")
            _ = v20
        }
    }
    if v3 {
        v4 = append(v4, 662)
        var v21 uint8 = 104
        _ = v21
    } else {
        var v22 uint = genFn0(0, 32767)
        for fi23 := 0; fi23 < 4; fi23++ {
            if (0.1 != ((1.5 / 1e10) + 0.0)) {
                v2 = v2
            } else {
                _ = len(v4)
            }
            _ = fi23
        }
        _ = v22
    }
    fmt.Println(v3)
    if (!false) {
        var v24 uint64 = 670
        if false {
            v25 := GenS1{F0: 127}
            v2 += v2
            v26 := GenS1{F0: 43}
            _ = v25
            _ = v26
        } else {
            switch ("hello" + "foo bar") {
            case "\n\t":
                var v27 string = "\n\t"
                _ = v27
            case "hello":
                var v28 bool = (42 >= 128)
                _ = v28
            case "":
                v29 := GenS4{F0: (v2 * v2)}
                _ = v29
            }
        }
        v30 := GenS1{F0: 127}
        _ = v24
        _ = v30
    }
    if v3 {
        var v31 float32 = (float32(3.14159) / (3.14159 * 3.14159))
        _ = v31
    } else {
        var v32 string = ("foo bar" + "a")
        v3 = true
        for fi33 := 0; fi33 < 11; fi33++ {
            if v3 {
                v34 := v1[2147483648]
                v3 = ((v3 && (!(!(v34 >= v34)))) && (true || (938 <= 65535)))
                fmt.Println(v32, fi33)
                _ = v34
            } else {
                var v35 int64 = 42
                switch v32 {
                case "hello":
                    v36 := GenS5{F0: genFn0(0, 1000), F1: 147, F2: 909}
                    _ = v36
                case "foo bar":
                    _ = len(v1)
                case "":
                    if (195 > 42) {
                        v37 := v1[2147483648]
                        _ = v37
                    } else {
                        v32 = (("\n\t" + v32) + v32)
                        if v3 {
                            v2 = (v2 << ((2) & 63))
                        }
                        var v38 uint = v2
                        _ = v38
                    }
                }
                v39 := GenS3{F0: (-(v35))}
                _ = v35
                _ = v39
            }
            v2 = genFn0(2, 1000)
            _ = fi33
        }
        _ = v32
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        v1[1] = 2147483648
        if (1 < 2147483647) {
            if ("foo bar" != "\n\t") {
                var v41 int8 = 101
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v42 uint8 = 255
                _ = v41
                _ = v42
            } else {
                v43 := GenS6{F0: 619}
                v1[2147483648] = 1
                _ = v43
            }
            var v44 string = "a"
            v2 = v2
            _ = v44
        }
        v2 = 166
        _ = fi40
    }
    fmt.Println(v3, v2)
    for fi45 := 0; fi45 < 19; fi45++ {
        v2 += v2
        switch ("" + "\n\t") {
        case "":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "a":
            v3 = v3
        case "\n\t":
            v1[uint32(v2)] = 353
        }
        _ = fi45
    }
    var v46 uint32 = 786
    v3 = v3
    var v47 bool = (("a" + ("hello" + (("a" + "foo bar") + "hello"))) == "hello")
    switch 42 {
    case 1:
        v48 := v1[(^(v46))]
        _ = v48
    case 2147483647:
        v49 := GenS7{F0: 0, F1: 627, F2: 1}
        _ = v49
    default:
        for fi50 := 0; fi50 < 2; fi50++ {
            v3 = (1 == 2147483647)
            v3 = false
            var v51 uint = (v2 >> ((0) & 63))
            _ = fi50
            _ = v51
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
}
