package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint16 }
func genFn0(p0 int, p1 uint32, p2 uint64) int16 {
    var v6 uint16 = 0
    if (-1 > 1) {
        for fi7 := 0; fi7 < 8; fi7++ {
            v6 += 42
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := GenS1{F0: 1, F1: fi8, F2: (v6 + (v6 ^ 444))}
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        var v10 uint32 = (p1 & (p1 | p1))
        _ = v10
    } else {
        v11 := GenS1{F0: (p1 - 52), F1: p0, F2: (^(v6))}
        p2 = p2
        _ = v11
    }
    var v12 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    return 1000
}
func genFn1() uint64 {
    var v13 uint8 = 42
    v14 := []int{9223372036854775807}
    _ = len(v14)
    _ = v13
    _ = v14
    return 2
}

func main() {
    var v1 int16 = 0
    var v2 uint8 = 128
    v3 := make(map[string]int)
    var v4 int16 = 629
    v5 := make(map[string]int)
    v4 = 1000
    v2 = 2
    v2 += v2
    v5[("foo bar" + "a")] = -1
    delete(v3, "foo bar")
    var v15 uint64 = 462
    var v16 uint64 = (v15 / ((4294967295) | 1))
    var v17 int8 = 0
    if true {
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 float32 = float32((1.5 / ((-2.25 * float64((1e10 + (0.1 * ((-2.25 - (1e10 / 1e10)) / (((((1e10 - (float32(-2.25) * -2.25)) * 1e10) + 3.14159) * 1e10) + 3.14159)))))) / 3.14159)))
            if true {
                var v20 uint = 1
                for fi21 := 0; fi21 < 12; fi21++ {
                    v16 += v15
                    if (v2 > v2) {
                        v22 := GenS1{F0: 94, F1: fi18, F2: 323}
                        _ = len(v3)
                        var v23 int8 = int8(fi21)
                        _ = v22
                        _ = v23
                    } else {
                        fmt.Println(v2, v15, v20)
                        fmt.Println(v17)
                        v3["foo bar"] = (^(fi21))
                    }
                    _ = fi21
                }
                _ = v20
            }
            if true {
                switch "\n\t" {
                case "a":
                    fmt.Println(v2, v16, v1)
                }
                var v24 uint = 906
                _ = v24
            } else {
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 float32 = 1e10
                    switch fi25 {
                    case 4294967295:
                        if true {
                            v27 := GenS1{F0: 0, F1: int(v16), F2: 42}
                            _ = v27
                        }
                    default:
                        v28 := []int{fi25, (-(fi25)), fi25}
                        _ = v28
                    }
                    v19 = (float32(float64(((v26 - -2.25) * (v19 * (3.14159 - ((((float32((1e10 - 0.0)) * float32((-2.25 - 0.0))) / float32(1e10)) / (0.1 + float32(0.0))) / (v26 - ((v26 / -2.25) * v26)))))))) * v19)
                    _ = fi25
                    _ = v26
                }
            }
            _ = fi18
            _ = v19
        }
        switch -540 {
        case -642:
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 int = fi29
                for fi31 := 0; fi31 < 2; fi31++ {
                    var v32 uint8 = (v2 * 143)
                    v16 = (^(v16))
                    _ = fi31
                    _ = v32
                }
                var v33 bool = (!true)
                _ = fi29
                _ = v30
                _ = v33
            }
        case 221:
            for fi34 := 0; fi34 < 12; fi34++ {
                var v35 uint32 = 0
                _ = fi34
                _ = v35
            }
        case 0:
            var v36 int = 32768
            _ = v36
        default:
            if true {
                v17 += v17
            } else {
                var v37 int32 = 651
                if (247 > 4294967295) {
                    switch 127 {
                    case 736:
                        v38 := []int{300, -847}
                        _ = v38
                    case -1:
                        v16 = genFn1()
                    }
                }
                v17 = v17
                _ = v37
            }
        }
    } else {
        v15 += genFn1()
        v15 -= (v15 & genFn1())
    }
    v39 := GenS1{F0: uint32(v16), F1: -2147483648, F2: 22}
    v5[("\n\t" + "\n\t")] = 0
    var v40 float32 = (3.14159 * (3.14159 + 0.1))
    var v41 uint = 18446744073709551615
    v41 = (v41 >> (((v41 - v41)) & 63))
    v16 -= genFn1()
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
