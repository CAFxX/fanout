package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int16 }
type GenS2 struct { F0 bool; F1 uint; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint32; F2 int }
type GenS5 struct { F0 int; F1 int8 }
func genFn0(p0 int16) string {
    p0 = (p0 << ((0) & 63))
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := []int{1, 32767, (fi5 | 4294967295)}
        v7 := GenS1{F0: 212, F1: 341, F2: p0}
        p0 -= (p0 ^ p0)
        _ = fi5
        _ = v6
        _ = v7
    }
    if (true || ((-(p0)) >= p0)) {
        var v8 uint = 1
        v9 := []int{-2147483648, 2147483647, -515}
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 128, F1: 255, F2: (p0 >> ((65535) & 63))}
    _ = p0
    _ = v10
    return ("foo bar" + ("\n\t" + "hello"))
}
func genFn1(p0 int) int16 {
    var v11 float64 = (1.5 * -2.25)
    var v12 float32 = float32(float64((1e10 * -2.25)))
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint64 = 817
        var v15 uint = 65535
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = p0
    _ = v11
    _ = v12
    return 393
}

func main() {
    v1 := []string{"", ""}
    v2 := make(map[string]bool)
    v3 := []int{130, 4294967295, -989}
    v4 := []int{128}
    var v16 int = 32768
    if false {
        for fi17 := 0; fi17 < 20; fi17++ {
            for fi18 := 0; fi18 < 5; fi18++ {
                if (!(-1 == 42)) {
                    switch "\n\t" {
                    case "":
                        var v19 bool = false
                        _ = v19
                    case "a":
                        switch 0 {
                        case 718:
                            v16 += fi17
                        case -1:
                            v16 += 42
                        case 128:
                            if true {
                                v4 = append(v4, fi17)
                                v16 = (v16 | (fi17 - (fi18 * 180)))
                                v16 += (v16 << ((0) & 63))
                                if true {
                                    v20 := GenS2{F0: true, F1: 1, F2: -298}
                                    _ = v20
                                } else {
                                    fmt.Println(fi17, fi17, fi17)
                                }
                            } else {
                                v16 += 0
                                if (-81 >= 32767) {
                                    var v21 uint = 358
                                    _ = v21
                                }
                            }
                        default:
                            fmt.Println(fi18, v16, fi18)
                            v16 = v16
                        }
                    case "foo bar":
                        var v22 bool = false
                        _ = v22
                    default:
                        v16 = fi18
                        var v23 bool = true
                        v4 = append(v4, fi18)
                        v16 += v16
                        _ = v23
                    }
                    v24 := GenS3{F0: 339}
                    delete(v2, ("\n\t" + "foo bar"))
                    _ = v24
                }
                v25 := GenS4{F0: (v16 - fi18), F1: 752, F2: v16}
                _ = fi18
                _ = v25
            }
            for fi26 := 0; fi26 < 11; fi26++ {
                v16 += fi26
                v16 += (fi17 * v16)
                _ = fi26
            }
            _ = fi17
        }
    }
    v27 := GenS3{F0: 42}
    switch "foo bar" {
    case "hello":
        for fi28 := 0; fi28 < 3; fi28++ {
            _ = len(v2)
            v16 -= (fi28 ^ 127)
            _ = fi28
        }
    case "":
        v29 := GenS5{F0: v16, F1: 92}
        _ = v29
    case "foo bar":
        if (!false) {
            if (!true) {
                v16 -= (v16 - v16)
                var v30 uint8 = 0
                _ = v30
            }
        } else {
            v16 = v16
            v2["foo bar"] = (434 == 2)
        }
    default:
        var v31 float32 = (((1.5 + 3.14159) + float32(-2.25)) - float32(1.5))
        _ = v31
    }
    if (575 > 644) {
        for fi32 := 0; fi32 < 18; fi32++ {
            if (!true) {
                v16 = v16
                fmt.Println(fi32, fi32)
                for fi33 := 0; fi33 < 16; fi33++ {
                    var v34 uint64 = 42
                    var v35 int8 = 127
                    v34 = (v34 / (((v34 / ((2) | 1))) | 1))
                    var v36 uint32 = 2147483648
                    _ = fi33
                    _ = v34
                    _ = v35
                    _ = v36
                }
            }
            var v37 string = ((("\n\t" + ((("hello" + ("a" + "a")) + "") + "\n\t")) + "a") + "a")
            var v38 int8 = -63
            _ = fi32
            _ = v37
            _ = v38
        }
        if (true && false) {
            v39 := GenS1{F0: 2, F1: 0, F2: genFn1((-(v16)))}
            v16 = -2147483648
            _ = v39
        } else {
            switch "hello" {
            case "\n\t":
                var v40 uint16 = 0
                _ = v40
            case "a":
                fmt.Println(v16)
            case "hello":
                fmt.Println(v16)
            }
        }
    } else {
        v2["hello"] = (!((true && false) || true))
        v16 += v16
    }
    fmt.Println(v16)
    var v41 bool = false
    var v42 uint8 = 214
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
