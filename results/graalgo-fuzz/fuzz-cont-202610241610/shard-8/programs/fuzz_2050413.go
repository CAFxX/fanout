package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 string }
type GenS2 struct { F0 int16; F1 bool; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int64; F2 uint16 }
func genFn0() uint64 {
    switch 9223372036854775807 {
    case 358:
        for fi4 := 0; fi4 < 20; fi4++ {
            v5 := GenS1{F0: 4294967295, F1: 255, F2: (("a" + "") + "foo bar")}
            var v6 uint8 = 155
            if true {
                v6 -= 128
                var v7 int8 = 42
                if true {
                    for fi8 := 0; fi8 < 18; fi8++ {
                        v9 := GenS2{F0: -347, F1: (618 == 42), F2: 309}
                        if ((0.1 + (0.1 - 1.5)) >= 1.5) {
                            v10 := []int{fi8, (fi4 % ((1) | 1)), fi8, (fi8 - (^(fi8)))}
                            var v11 bool = (((("a" + ((("foo bar" + "hello") + (("\n\t" + "hello") + ("foo bar" + ("hello" + "foo bar")))) + ("a" + "foo bar"))) <= "a") && true) && (true || true))
                            _ = v10
                            _ = v11
                        }
                        _ = fi8
                        _ = v9
                    }
                } else {
                    var v12 int32 = 534
                    v13 := []int{fi4, fi4, 186}
                    _ = v12
                    _ = v13
                }
                _ = v7
            } else {
                var v14 int16 = -670
                _ = v14
            }
            _ = fi4
            _ = v5
            _ = v6
        }
    case -352:
        var v15 uint16 = 0
        _ = v15
    case -375:
        var v16 int16 = -1
        _ = v16
    default:
        switch 42 {
        case 127:
            v17 := []int{}
            _ = v17
        case -9223372036854775808:
            var v18 uint8 = 2
            _ = v18
        case 32767:
            v19 := GenS2{F0: -32768, F1: ("foo bar" >= (("\n\t" + ("hello" + "foo bar")) + ("a" + "\n\t"))), F2: 1}
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS3{F0: 128}
        _ = fi20
        _ = v21
    }
    var v22 uint8 = 1
    _ = v22
    return 494
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int64 = 1
    var v3 uint8 = 42
    if false {
        v2 = (v2 * (v2 & v2))
    } else {
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 uint64 = genFn0()
            v25 := []int{}
            v26 := GenS4{F0: 0, F1: -9223372036854775808, F2: 2}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
    }
    fmt.Println(v2)
    if true {
        fmt.Println(v1, v1, v3)
        var v27 float64 = float64(((3.14159 - (0.1 * 3.14159)) * 1e10))
        _ = v27
    }
    v28 := GenS4{F0: 185, F1: (v2 << ((255) & 63)), F2: 255}
    v2 = (-(v2))
    switch 32767 {
    case 42:
        fmt.Println(v3)
    case 69:
        v29 := []int{128}
        _ = v29
    case 32767:
        v30 := GenS1{F0: 42, F1: 169, F2: ("a" + (("\n\t" + ("" + "\n\t")) + "foo bar"))}
        _ = v30
    default:
        var v31 uint32 = 582
        _ = v31
    }
    if true {
        var v32 string = ("hello" + ((("foo bar" + "hello") + "\n\t") + (("foo bar" + "hello") + "a")))
        _ = v32
    } else {
        var v33 uint16 = 255
        switch "\n\t" {
        case "\n\t":
            switch 9223372036854775807 {
            case -2147483648:
                if true {
                    if false {
                        fmt.Println(v1)
                    }
                } else {
                    for fi34 := 0; fi34 < 1; fi34++ {
                        v33 = v33
                        _ = fi34
                    }
                    var v35 uint64 = v1
                    var v36 uint16 = (v33 & v33)
                    _ = v35
                    _ = v36
                }
            default:
                v37 := []int{}
                _ = v37
            }
        case "hello":
            for fi38 := 0; fi38 < 8; fi38++ {
                v2 -= 9223372036854775807
                _ = fi38
            }
        case "":
            switch "foo bar" {
            case "a":
                v3 += (v3 >> ((1) & 63))
            case "hello":
                switch 42 {
                case 9223372036854775807:
                    for fi39 := 0; fi39 < 17; fi39++ {
                        fmt.Println(v1, v2)
                        fmt.Println(v33)
                        v3 = (v3 / ((2) | 1))
                        _ = fi39
                    }
                default:
                    var v40 uint32 = 42
                    _ = v40
                }
            }
        }
        _ = v33
    }
    var v41 uint = 255
    v42 := GenS1{F0: (v41 - (v41 - 713)), F1: v41, F2: "a"}
    v41 = uint(v3)
    v43 := []int{127, 4294967295, -168}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v41
    _ = v42
    _ = v43
}
