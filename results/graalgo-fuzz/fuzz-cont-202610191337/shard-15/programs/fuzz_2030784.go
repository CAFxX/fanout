package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 uint8; F2 int8 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int8, p1 uint64, p2 uint32) int32 {
    v4 := []int{4294967295, 9223372036854775807}
    if len(v4) > 0 {
        v4[0] = 32768
    }
    switch (("foo bar" + "hello") + "") {
    case "\n\t":
        switch 32767 {
        case -794:
            p1 = 1
        }
    default:
        var v5 uint64 = p1
        _ = v5
    }
    var v6 int = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return 2147483647
}
func genFn1(p0 bool) uint64 {
    v7 := []int{-992}
    if p0 {
        if (171 != 384) {
            var v8 float64 = 1.5
            _ = v8
        }
        v7 = append(v7, 32767)
    } else {
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 uint32 = 101
            if (p0 || p0) {
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 uint32 = v10
                    v7 = append(v7, (fi9 - fi9))
                    _ = fi11
                    _ = v12
                }
                var v13 float32 = float32(((3.14159 + (0.0 - 1.5)) * 3.14159))
                if len(v7) > 0 {
                    v7[0] = fi9
                }
                _ = v13
            }
            _ = fi9
            _ = v10
        }
    }
    switch -828 {
    case 840:
        fmt.Println(p0)
    case -2147483648:
        for fi14 := 0; fi14 < 16; fi14++ {
            for fi15 := 0; fi15 < 16; fi15++ {
                switch fi14 {
                case 32768:
                    if p0 {
                        p0 = true
                        p0 = true
                    }
                }
                fmt.Println(fi15, fi15, fi14)
                p0 = true
                _ = fi15
            }
            var v16 float32 = (0.1 * 1e10)
            _ = fi14
            _ = v16
        }
    case -1:
        if (!p0) {
            if p0 {
                if (759 >= 65535) {
                    switch "hello" {
                    case "hello":
                        p0 = (-1 <= 32767)
                    case "foo bar":
                        for fi17 := 0; fi17 < 10; fi17++ {
                            v18 := GenS1{F0: fi17}
                            p0 = false
                            p0 = p0
                            _ = fi17
                            _ = v18
                        }
                    case "\n\t":
                        p0 = p0
                    default:
                        if len(v7) > 0 {
                            v7[0] = 42
                        }
                    }
                    switch ("\n\t" + "") {
                    case "foo bar":
                        v19 := GenS2{F0: 42, F1: 2147483648}
                        _ = v19
                    case "a":
                        fmt.Println(p0)
                    }
                }
            } else {
                v7 = append(v7, 987)
            }
        }
    default:
        v20 := GenS3{F0: 1000}
        _ = v20
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 255
    v3 := make(map[int]bool)
    switch ("foo bar" + (("foo bar" + (((("hello" + "foo bar") + "hello") + "hello") + "a")) + "")) {
    case "":
        if (-128 > 42) {
            v2 = v2
            v21 := GenS4{F0: 1000, F1: 1, F2: 64}
            if true {
                switch ("foo bar" + "") {
                case "":
                    v2 = (v2 & (v2 << ((0) & 63)))
                case "hello":
                    v22 := []int{}
                    _ = v22
                case "a":
                    v2 = (v2 << ((694) & 63))
                default:
                    _ = len(v1)
                }
                var v23 uint16 = 32767
                var v24 uint16 = v2
                _ = v23
                _ = v24
            } else {
                v2 -= (v2 * (v2 << ((4294967295) & 63)))
            }
            _ = v21
        }
    default:
        v25 := []int{4294967295, -609}
        _ = v25
    }
    v26 := []int{-9223372036854775808, 32767, -798}
    v2 = (^(v2))
    var v27 uint8 = 255
    v28 := GenS1{F0: -670}
    _ = len(v3)
    v2 += v2
    for fi29 := 0; fi29 < 2; fi29++ {
        _ = len(v3)
        _ = fi29
    }
    if (true || false) {
        fmt.Println(v2)
        var v30 float32 = 0.1
        _ = v30
    } else {
        v2 = (v2 % ((517) | 1))
        for fi31 := 0; fi31 < 12; fi31++ {
            delete(v1, "a")
            var v32 int8 = 1
            var v33 uint32 = 2147483648
            _ = fi31
            _ = v32
            _ = v33
        }
        if false {
            fmt.Println(v2)
            v34 := GenS5{F0: 42}
            _ = v34
        } else {
            switch ("\n\t" + "\n\t") {
            case "hello":
                for fi35 := 0; fi35 < 16; fi35++ {
                    for fi36 := 0; fi36 < 5; fi36++ {
                        var v37 int8 = 1
                        _ = fi36
                        _ = v37
                    }
                    v27 -= 115
                    _ = fi35
                }
            case "":
                fmt.Println(v27, v27)
            }
        }
    }
    var v38 float32 = 0.0
    var v39 int = 9223372036854775807
    var v40 uint = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v40
}
