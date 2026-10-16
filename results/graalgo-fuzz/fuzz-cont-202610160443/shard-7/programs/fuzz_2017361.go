package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0() int32 {
    var v7 uint16 = 32767
    switch "hello" {
    case "foo bar":
        var v8 float32 = (-2.25 / (-2.25 + 3.14159))
        _ = v8
    case "\n\t":
        var v9 uint16 = 0
        _ = v9
    case "":
        var v10 int32 = 829
        _ = v10
    default:
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int = fi11
            var v13 int32 = -401
            if true {
                v12 -= -1
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    v14 := GenS1{F0: 42}
    v15 := []int{592, 0}
    _ = v14
    _ = v15
    _ = v7
    return 0
}
func genFn1(p0 uint8) uint8 {
    for fi16 := 0; fi16 < 5; fi16++ {
        p0 += 208
        for fi17 := 0; fi17 < 7; fi17++ {
            p0 = (p0 >> ((1) & 63))
            switch "" {
            case "hello":
                var v18 uint32 = 0
                _ = v18
            }
            fmt.Println(fi17)
            _ = fi17
        }
        _ = fi16
    }
    p0 = p0
    _ = p0
    return (p0 ^ (p0 * p0))
}

func main() {
    var v1 uint16 = 65535
    v2 := []int{140, -2147483648, -2147483648}
    v3 := make(map[uint32]string)
    v4 := []uint32{2147483648, 0, 2147483647}
    var v5 uint16 = 88
    v6 := make(map[int]bool)
    v1 -= (v5 + v5)
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 int64 = 9223372036854775807
        switch "\n\t" {
        case "hello":
            for fi21 := 0; fi21 < 20; fi21++ {
                if (0 == (v1 ^ v1)) {
                    var v22 int64 = v20
                    fmt.Println(v20, v5, v1)
                    _ = len(v3)
                    switch "" {
                    case "foo bar":
                        if ((true || (!true)) && ((genFn1(42) != 1) || false)) {
                            _ = len(v4)
                        }
                    }
                    _ = v22
                } else {
                    var v23 float64 = (0.0 * (1.5 / 1e10))
                    v20 -= (v20 / ((0) | 1))
                    _ = v23
                }
                v24 := GenS1{F0: 71}
                _ = fi21
                _ = v24
            }
        case "\n\t":
            fmt.Println(v5, fi19, v20)
        case "":
            if false {
                v6[int(v20)] = (float32(((0.0 - float64((1e10 * 0.0))) * 1.5)) > 3.14159)
                v20 = 166
                switch (fi19 >> ((921) & 63)) {
                case -899:
                    var v25 string = ""
                    _ = v25
                case 1:
                    if (255 < 255) {
                        _ = len(v4)
                    }
                }
            } else {
                switch fi19 {
                case 32767:
                    for fi26 := 0; fi26 < 3; fi26++ {
                        fmt.Println(fi19, v5)
                        _ = fi26
                    }
                default:
                    v27 := GenS2{F0: 798}
                    _ = v27
                }
                _ = len(v4)
                var v28 float64 = 1e10
                _ = v28
            }
        default:
            var v29 float64 = 1e10
            _ = v29
        }
        delete(v3, 0)
        _ = fi19
        _ = v20
    }
    if false {
        var v30 uint8 = 2
        _ = v30
    } else {
        var v31 int64 = 1
        var v32 uint64 = 42
        _ = v31
        _ = v32
    }
    v33 := v3[2]
    for fi34 := 0; fi34 < 13; fi34++ {
        var v35 bool = false
        switch ("a" + (("" + "") + "hello")) {
        case "foo bar":
            for fi36 := 0; fi36 < 11; fi36++ {
                var v37 int16 = -32768
                _ = fi36
                _ = v37
            }
        case "a":
            fmt.Println(v35, v35, v1)
        case "hello":
            fmt.Println(v1, fi34, v5)
        default:
            v35 = v35
        }
        var v38 uint32 = 407
        _ = fi34
        _ = v35
        _ = v38
    }
    v5 -= (v5 / ((0) | 1))
    if (79 > 42) {
        var v39 uint8 = genFn1(160)
        _ = v39
    }
    var v40 int32 = 1
    if len(v2) > 0 {
        v2[0] = -1
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi41 := 0; fi41 < 19; fi41++ {
        if false {
            var v42 int32 = (v40 - v40)
            _ = v42
        } else {
            var v43 uint16 = v5
            _ = v43
        }
        if (!true) {
            fmt.Println(v5, fi41, v40)
            for fi44 := 0; fi44 < 14; fi44++ {
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
                _ = fi44
            }
        }
        v40 += -2147483648
        _ = fi41
    }
    var v45 uint32 = 0
    for fi46 := 0; fi46 < 1; fi46++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi46
    }
    v5 = v5
    if (816 >= 931) {
        var v47 int8 = -118
        _ = v47
    }
    fmt.Println(v45, v5, v45)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v40
    _ = v45
    _ = v5
    _ = v6
}
