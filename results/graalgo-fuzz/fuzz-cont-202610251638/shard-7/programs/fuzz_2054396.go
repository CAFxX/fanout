package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int64; F1 bool }
func genFn0(p0 uint8, p1 uint64) int8 {
    switch "a" {
    case "":
        v7 := GenS1{F0: (-1 >= 0), F1: 2147483648}
        _ = v7
    default:
        var v8 uint32 = 718
        _ = v8
    }
    if true {
        v9 := []int{-1, 549, -1, -9223372036854775808}
        _ = v9
    }
    var v10 int32 = 2147483647
    var v11 int32 = (^(v10))
    v12 := GenS2{F0: "", F1: uint32(p0)}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return -1
}
func genFn1(p0 uint16) int64 {
    v13 := GenS3{F0: 2147483648}
    fmt.Println(p0)
    p0 = 2
    fmt.Println(p0)
    _ = p0
    _ = v13
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{("foo bar" + "\n\t"), ""}
    var v3 uint = 4294967295
    var v4 uint16 = 583
    var v5 int16 = 42
    var v6 int32 = -1
    if (2 >= 487) {
        switch 127 {
        case 4294967295:
            switch (("\n\t" + "hello") + "\n\t") {
            case "":
                for fi14 := 0; fi14 < 12; fi14++ {
                    if ("foo bar" < "a") {
                        v15 := v1[792]
                        _ = v15
                    } else {
                        var v16 int16 = v5
                        _ = v16
                    }
                    if (!false) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v17 float32 = -2.25
                        _ = v17
                    } else {
                        fmt.Println(v6, v3, fi14)
                        for fi18 := 0; fi18 < 10; fi18++ {
                            v5 -= 42
                            fmt.Println(v3, fi18)
                            _ = fi18
                        }
                        switch "hello" {
                        case "foo bar":
                            v4 = (v4 << ((v3) & 63))
                        case "a":
                            var v19 uint16 = v4
                            _ = v19
                        case "hello":
                            v1[860] = 2147483648
                        default:
                            var v20 int64 = 631
                            _ = v20
                        }
                    }
                    v6 -= (v6 ^ (-(v6)))
                    _ = fi14
                }
            case "a":
                if ((0.0 * (1e10 - float64(1.5))) <= -2.25) {
                    v3 = v3
                }
            case "\n\t":
                var v21 uint8 = 0
                _ = v21
            }
        case 9223372036854775807:
            var v22 uint = v3
            _ = v22
        default:
            if (-1 <= 32768) {
                for fi23 := 0; fi23 < 20; fi23++ {
                    switch ("" + ((("a" + "\n\t") + "\n\t") + "hello")) {
                    case "\n\t":
                        var v24 int16 = (v5 << (((v3 * (v3 + (v3 << (((^(v3))) & 63))))) & 63))
                        _ = v24
                    case "a":
                        var v25 uint32 = 2147483648
                        _ = v25
                    case "foo bar":
                        for fi26 := 0; fi26 < 16; fi26++ {
                            fmt.Println(fi26, v5, v4)
                            v27 := GenS4{F0: v4, F1: 4294967295, F2: 42}
                            v3 -= (v3 ^ (v3 & (v3 % ((748) | 1))))
                            _ = fi26
                            _ = v27
                        }
                    default:
                        v4 -= v4
                    }
                    for fi28 := 0; fi28 < 19; fi28++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch fi28 {
                        case 9223372036854775807:
                            var v29 int = fi23
                            _ = v29
                        default:
                            var v30 uint16 = v4
                            _ = v30
                        }
                        var v31 int16 = 852
                        v2 = append(v2, "hello")
                        _ = fi28
                        _ = v31
                    }
                    _ = fi23
                }
                switch ("foo bar" + "a") {
                case "a":
                    delete(v1, 2147483648)
                case "\n\t":
                    fmt.Println(v4)
                case "":
                    var v32 int64 = 42
                    _ = v32
                default:
                    v5 -= v5
                }
                v33 := GenS2{F0: "foo bar", F1: 65535}
                _ = v33
            }
        }
        v4 = (v4 + 42)
    } else {
        v34 := GenS5{F0: -9223372036854775808, F1: (!true)}
        fmt.Println(v3, v3, v5)
        _ = v34
    }
    v3 -= v3
    switch 32767 {
    case 127:
        _ = len(v2)
    default:
        if len(v2) > 0 {
            v2[0] = ("\n\t" + "\n\t")
        }
    }
    var v35 uint8 = 95
    v4 = (^(v4))
    for fi36 := 0; fi36 < 6; fi36++ {
        if ((v6 * (v6 ^ 42)) != v6) {
            for fi37 := 0; fi37 < 15; fi37++ {
                switch "hello" {
                case "":
                    v35 += v35
                case "\n\t":
                    v35 = v35
                default:
                    v35 -= v35
                    v4 = v4
                }
                if true {
                    v35 = v35
                    var v38 int = (fi37 + fi37)
                    _ = v38
                }
                _ = fi37
            }
        } else {
            v39 := v1[1]
            _ = v39
        }
        _ = fi36
    }
    var v40 uint16 = v4
    var v41 uint16 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
