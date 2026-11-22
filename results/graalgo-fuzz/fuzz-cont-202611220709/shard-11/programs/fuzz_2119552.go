package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0() uint64 {
    var v7 bool = (true && (("foo bar" + ("hello" + "")) < (("\n\t" + "foo bar") + (("hello" + "foo bar") + "\n\t"))))
    v8 := GenS1{F0: "\n\t", F1: 0}
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int8) uint32 {
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: "", F1: 0}
        for fi11 := 0; fi11 < 4; fi11++ {
            if true {
                var v12 uint8 = 128
                if (0 >= 42) {
                    fmt.Println(p0)
                    v13 := GenS1{F0: "foo bar", F1: 65535}
                    if (true && false) {
                        switch "hello" {
                        case "\n\t":
                            if (0.0 <= 0.1) {
                                var v14 uint16 = 389
                                _ = v14
                            }
                        case "hello":
                            var v15 string = (("hello" + "a") + "foo bar")
                            _ = v15
                        }
                        v16 := GenS2{F0: -32768, F1: 32767}
                        v12 -= v12
                        _ = v16
                    } else {
                        fmt.Println(p0)
                    }
                    fmt.Println(fi11)
                    _ = v13
                } else {
                    fmt.Println(v12, v12, p0)
                }
                fmt.Println(fi9)
                _ = v12
            } else {
                fmt.Println(p0)
                p0 = (p0 << ((2) & 63))
            }
            var v17 uint32 = 1
            _ = fi11
            _ = v17
        }
        _ = fi9
        _ = v10
    }
    switch ("foo bar" + "\n\t") {
    case "hello":
        for fi18 := 0; fi18 < 8; fi18++ {
            if (true || false) {
                p0 = p0
                switch (^(fi18)) {
                case 32768:
                    p0 = (p0 >> ((1) & 63))
                case 127:
                    p0 = (p0 & -1)
                case -967:
                    if ((fi18 | fi18) > fi18) {
                        p0 += 127
                        for fi19 := 0; fi19 < 1; fi19++ {
                            var v20 int = fi19
                            v21 := []int{(fi19 / ((fi19) | 1))}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                default:
                    var v22 uint = 18446744073709551615
                    _ = v22
                }
                switch "foo bar" {
                case "foo bar":
                    var v23 uint8 = 2
                    _ = v23
                default:
                    var v24 int16 = 32767
                    _ = v24
                }
            } else {
                v25 := GenS2{F0: 1, F1: 65535}
                v26 := []int{}
                _ = v25
                _ = v26
            }
            switch "a" {
            case "":
                v27 := GenS1{F0: "a", F1: 980}
                _ = v27
            case "a":
                p0 += p0
            case "hello":
                p0 -= p0
            }
            p0 -= -128
            _ = fi18
        }
    default:
        if true {
            fmt.Println(p0)
        }
    }
    p0 -= -96
    _ = p0
    return 2147483647
}

func main() {
    var v1 int = 0
    var v2 int = 203
    var v3 uint32 = 885
    var v4 uint8 = 42
    v5 := []int{32768, int(v4)}
    v6 := make(map[string]string)
    var v28 float32 = 1e10
    for fi29 := 0; fi29 < 6; fi29++ {
        v5 = append(v5, v1)
        _ = fi29
    }
    v30 := GenS2{F0: 1000, F1: 65535}
    if (v28 > (0.0 + (v28 / v28))) {
        v31 := GenS2{F0: int16(v3), F1: 255}
        v4 += (v4 % ((2) | 1))
        _ = v31
    }
    delete(v6, "\n\t")
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 uint64 = 112
        switch 4294967295 {
        case 1:
            switch "hello" {
            case "":
                fmt.Println(v2)
            case "a":
                v34 := v6["a"]
                _ = v34
            case "foo bar":
                switch (v2 * (-(v2))) {
                case 42:
                    v3 -= (v3 / ((uint32(v33)) | 1))
                default:
                    var v35 uint32 = (v3 * 568)
                    _ = v35
                }
            default:
                for fi36 := 0; fi36 < 10; fi36++ {
                    var v37 float64 = (float64(((0.1 / v28) * float32(0.1))) + (1.5 + (0.0 - (float64((v28 + 0.1)) - 1.5))))
                    switch v2 {
                    case 2147483647:
                        v28 = v28
                    case 0:
                        if false {
                            v3 = (v3 & 237)
                            if (false && false) {
                                v4 -= 201
                            } else {
                                v2 = fi36
                            }
                            v38 := GenS1{F0: "hello", F1: 1}
                            _ = v38
                        }
                        if len(v5) > 0 {
                            v5[0] = v1
                        }
                    case 221:
                        v39 := GenS1{F0: "hello", F1: 740}
                        _ = v39
                    }
                    var v40 bool = true
                    _ = fi36
                    _ = v37
                    _ = v40
                }
            }
        default:
            var v41 float32 = v28
            _ = v41
        }
        _ = fi32
        _ = v33
    }
    switch "" {
    case "hello":
        var v42 float32 = v28
        _ = v42
    }
    if len(v5) > 0 {
        v5[0] = v1
    }
    var v43 string = "a"
    v3 -= (v3 & (v3 & (v3 + genFn1(42))))
    var v44 int8 = 12
    v44 = v44
    fmt.Println(v44, v44)
    var v45 string = v43
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
