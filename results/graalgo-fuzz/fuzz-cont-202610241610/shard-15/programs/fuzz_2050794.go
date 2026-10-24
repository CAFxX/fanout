package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint16; F1 uint }
func genFn0(p0 bool, p1 uint16, p2 uint) uint32 {
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint = 747
        v9 := []int{(fi7 * (fi7 * (fi7 % ((fi7) | 1)))), 127, (-(fi7))}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 32767, F1: 65535}
    p0 = (42 > 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 2147483648
}
func genFn1() int16 {
    for fi11 := 0; fi11 < 10; fi11++ {
        if false {
            var v12 int8 = -128
            for fi13 := 0; fi13 < 12; fi13++ {
                v12 += (v12 >> ((2) & 63))
                var v14 uint16 = 642
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        var v15 int32 = -1
        _ = fi11
        _ = v15
    }
    switch "foo bar" {
    case "a":
        if false {
            var v16 float64 = -2.25
            v16 += v16
            _ = v16
        }
    }
    return -967
}

func main() {
    var v1 int64 = -993
    v2 := []string{"hello", ("foo bar" + ("hello" + "a")), "a", ""}
    var v3 int = 2147483647
    v4 := make(map[uint32]string)
    var v5 uint = 42
    var v6 int64 = v1
    if true {
        fmt.Println(v3, v6)
        v2 = append(v2, "a")
        v6 = (v1 >> ((v5) & 63))
    }
    v5 = 65535
    if (!((v3 % ((v3) | 1)) < v3)) {
        fmt.Println(v1, v3, v1)
        fmt.Println(v6, v5)
    }
    switch (("hello" + ("a" + "hello")) + "\n\t") {
    case "":
        v6 += v1
    case "\n\t":
        _ = len(v4)
    default:
        var v17 int32 = 2147483647
        _ = v17
    }
    v18 := GenS2{F0: (true && (255 <= 127)), F1: (v5 << ((4294967295) & 63))}
    var v19 int64 = -1
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            v6 -= (v1 >> ((255) & 63))
            _ = fi21
        }
        _ = fi20
    }
    if false {
        v5 = (v5 / (((^(v5))) | 1))
        if true {
            for fi22 := 0; fi22 < 10; fi22++ {
                v19 -= v6
                v3 += (v3 / (((fi22 ^ v3)) | 1))
                _ = fi22
            }
            var v23 bool = (127 < -128)
            switch "\n\t" {
            case "hello":
                var v24 bool = v23
                _ = v24
            case "\n\t":
                v23 = v23
            default:
                v6 += (-(v6))
            }
            _ = v23
        } else {
            switch "" {
            case "\n\t":
                fmt.Println(v5, v6, v5)
            case "hello":
                var v25 uint64 = 684
                _ = v25
            default:
                fmt.Println(v3, v5)
            }
            var v26 int32 = 76
            for fi27 := 0; fi27 < 20; fi27++ {
                if false {
                    if (v5 == v5) {
                        v28 := GenS3{F0: 42, F1: 135}
                        v3 += (fi27 << ((v5) & 63))
                        _ = v28
                    } else {
                        v3 = (v3 % ((32768) | 1))
                    }
                    v1 = -9223372036854775808
                }
                _ = fi27
            }
            _ = v26
        }
    } else {
        var v29 string = "foo bar"
        v4[2] = ((v29 + ((("hello" + ((v29 + ("foo bar" + (v29 + "foo bar"))) + v29)) + ((("" + (v29 + v29)) + (v29 + "\n\t")) + v29)) + ("foo bar" + v29))) + v29)
        var v30 int8 = 0
        _ = v29
        _ = v30
    }
    var v31 float32 = 1e10
    v32 := GenS4{F0: 332, F1: v5}
    if (0 <= 643) {
        var v33 uint8 = 0
        switch -2147483648 {
        case 127:
            var v34 float64 = 0.1
            _ = v34
        case 2147483647:
            v19 += v19
        case 614:
            for fi35 := 0; fi35 < 19; fi35++ {
                for fi36 := 0; fi36 < 10; fi36++ {
                    v31 -= -2.25
                    _ = fi36
                }
                var v37 uint32 = 1
                _ = fi35
                _ = v37
            }
        default:
            v38 := GenS1{F0: genFn1(), F1: genFn0((1 <= -128), 1, 42)}
            _ = v38
        }
        _ = v33
    } else {
        v39 := GenS3{F0: 255, F1: genFn0((!(-69 >= 42)), 65535, 1)}
        var v40 int64 = -234
        _ = v39
        _ = v40
    }
    var v41 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
