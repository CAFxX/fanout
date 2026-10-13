package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 uint; F1 int16; F2 uint16 }
func genFn0(p0 int8, p1 uint, p2 bool) uint32 {
    fmt.Println(p2, p1, p1)
    if p2 {
        switch "a" {
        case "":
            var v6 int64 = -1
            _ = v6
        case "a":
            for fi7 := 0; fi7 < 19; fi7++ {
                if (float32(1e10) < 1e10) {
                    p1 = (p1 / (((p1 % ((p1) | 1))) | 1))
                    switch (fi7 + 737) {
                    case -101:
                        p1 = 4294967295
                    case -9223372036854775808:
                        v8 := []int{2147483647, fi7}
                        _ = v8
                    default:
                        p2 = (!p2)
                    }
                    p2 = false
                } else {
                    p2 = (p2 || p2)
                    v9 := []int{fi7, (-(fi7)), (-(fi7)), fi7}
                    _ = v9
                }
                var v10 float32 = (0.0 - 3.14159)
                _ = fi7
                _ = v10
            }
        case "foo bar":
            p1 -= 2
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := []int{(fi11 & fi11), (^(fi11)), (fi11 & fi11)}
            _ = fi11
            _ = v12
        }
    } else {
        v13 := []int{439, -2147483648, 32767, -633}
        if p2 {
            if ((p0 + (p0 * (p0 >> ((p1) & 63)))) > -128) {
                if (!true) {
                    for fi14 := 0; fi14 < 12; fi14++ {
                        p2 = p2
                        _ = fi14
                    }
                } else {
                    switch "hello" {
                    case "":
                        var v15 uint32 = 365
                        _ = v15
                    case "hello":
                        var v16 float32 = (float32(1.5) / (0.0 + 1e10))
                        _ = v16
                    }
                    var v17 uint64 = 261
                    _ = v17
                }
                if (!false) {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                p2 = true
            }
            v18 := GenS1{F0: 0}
            v19 := GenS2{F0: 2, F1: 4294967295, F2: 2}
            _ = v18
            _ = v19
        } else {
            var v20 int8 = 100
            p0 = v20
            _ = v20
        }
        if (!((!p2) && true)) {
            if (398 >= 32767) {
                p0 -= (p0 * p0)
                if p2 {
                    p1 -= (p1 - (p1 | p1))
                    var v21 uint8 = 1
                    _ = v21
                }
            } else {
                v22 := GenS3{F0: 1, F1: 2}
                p0 = (p0 + (p0 >> ((4294967295) & 63)))
                p2 = p2
                _ = v22
            }
        }
        _ = v13
    }
    fmt.Println(p1)
    p1 = p1
    var v23 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return (v23 & (^(v23)))
}
func genFn1(p0 string, p1 int64, p2 string) uint16 {
    p0 = ""
    p2 = p2
    if true {
        var v24 bool = false
        fmt.Println(p1, p0, p0)
        if (true && v24) {
            var v25 uint64 = 2
            _ = v25
        }
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 string = "a"
    v2 := []uint32{1, 2, 2147483647}
    v3 := make(map[int]uint32)
    var v4 uint64 = 4294967295
    v5 := []string{v1, (v1 + v1), "\n\t"}
    for fi26 := 0; fi26 < 20; fi26++ {
        v1 = ("hello" + v1)
        fmt.Println(v1)
        var v27 uint8 = 128
        _ = fi26
        _ = v27
    }
    if false {
        var v28 uint16 = genFn1(v1, 900, "")
        fmt.Println(v28, v4)
        _ = v28
    }
    v29 := GenS4{F0: 380}
    var v30 uint16 = 255
    switch int(v4) {
    case 2147483647:
        for fi31 := 0; fi31 < 4; fi31++ {
            switch "hello" {
            case "":
                v30 += (^(v30))
            case "foo bar":
                for fi32 := 0; fi32 < 15; fi32++ {
                    if false {
                        var v33 uint8 = uint8(v4)
                        v4 -= (v4 | (v4 * v4))
                        _ = v33
                    }
                    _ = fi32
                }
            case "hello":
                v1 = (v1 + (((("foo bar" + v1) + "\n\t") + "hello") + (v1 + "foo bar")))
            default:
                for fi34 := 0; fi34 < 7; fi34++ {
                    v1 = v1
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                    var v35 int32 = 42
                    _ = fi34
                    _ = v35
                }
            }
            v30 += 42
            var v36 uint32 = 1
            _ = fi31
            _ = v36
        }
    default:
        fmt.Println(v1)
    }
    v30 = v30
    v37 := GenS5{F0: 42, F1: 225}
    var v38 bool = (-128 > -128)
    var v39 uint16 = (^(v30))
    var v40 int = 1
    var v41 float64 = (0.1 / 1.5)
    v42 := GenS5{F0: -119, F1: 168}
    for fi43 := 0; fi43 < 19; fi43++ {
        _ = len(v2)
        if (!(1 >= 0)) {
            var v44 int64 = 470
            if (v44 < v44) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v45 float64 = 1.5
                _ = v45
            }
            _ = v44
        }
        _ = fi43
    }
    var v46 int8 = 127
    v47 := GenS3{F0: (v40 | v40), F1: v30}
    v48 := GenS6{F0: 42, F1: int16(v46), F2: v39}
    v30 = (v30 - 32767)
    v49 := GenS2{F0: genFn1("hello", -1, v1), F1: 2147483647, F2: 137}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}
