package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint32) uint64 {
    v6 := GenS1{F0: p0}
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 = 229
        p0 = p0
        _ = fi7
    }
    p0 -= 2147483647
    fmt.Println(p0)
    if true {
        if (0 > 11) {
            var v8 uint16 = 167
            _ = v8
        }
        var v9 int = 128
        _ = v9
    } else {
        if false {
            v10 := GenS1{F0: 856}
            if (!true) {
                fmt.Println(p0)
                fmt.Println(p0)
                if (127 < 1) {
                    var v11 int16 = int16(p0)
                    _ = v11
                }
            } else {
                v12 := GenS1{F0: p0}
                _ = v12
            }
            _ = v10
        } else {
            p0 -= p0
            v13 := GenS1{F0: (p0 >> ((867) & 63))}
            var v14 float64 = ((3.14159 - (-2.25 * (0.1 * 3.14159))) + (1.5 + 3.14159))
            _ = v13
            _ = v14
        }
        var v15 uint32 = p0
        _ = v15
    }
    _ = p0
    _ = v6
    return 4294967295
}
func genFn1(p0 int64) int16 {
    v16 := GenS1{F0: 42}
    fmt.Println(p0)
    for fi17 := 0; fi17 < 20; fi17++ {
        p0 -= p0
        p0 = p0
        _ = fi17
    }
    if true {
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 string = "hello"
            for fi20 := 0; fi20 < 7; fi20++ {
                var v21 string = v19
                v21 = v21
                switch fi20 {
                case 32768:
                    v19 = (v19 + v19)
                default:
                    v21 = v19
                }
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
    } else {
        p0 = (p0 ^ p0)
        fmt.Println(p0)
        switch 583 {
        case -9223372036854775808:
            switch -236 {
            case 9223372036854775807:
                for fi22 := 0; fi22 < 6; fi22++ {
                    if (true && (!(32 >= 42))) {
                        v23 := []int{fi22, (fi22 % ((fi22) | 1)), 2147483647}
                        var v24 int16 = 1
                        _ = v23
                        _ = v24
                    } else {
                        var v25 uint8 = 128
                        p0 = 488
                        switch fi22 {
                        case -1:
                            v26 := []int{}
                            _ = v26
                        default:
                            var v27 uint16 = 255
                            _ = v27
                        }
                        _ = v25
                    }
                    var v28 int = (fi22 * 2147483647)
                    fmt.Println(p0)
                    _ = fi22
                    _ = v28
                }
            case 127:
                v29 := []int{0}
                _ = v29
            case 128:
                if true {
                    var v30 int = 1
                    _ = v30
                }
            default:
                var v31 uint8 = 1
                _ = v31
            }
        }
    }
    _ = p0
    _ = v16
    return -32768
}

func main() {
    var v1 int32 = 2147483647
    var v2 float32 = 3.14159
    var v3 int = 1
    var v4 string = "\n\t"
    var v5 uint8 = 42
    v3 = (v3 | (v3 / ((v3) | 1)))
    v32 := GenS1{F0: 402}
    var v33 uint32 = 4294967295
    v1 = (v1 << ((255) & 63))
    var v34 uint16 = 357
    v33 = (^(v33))
    var v35 uint64 = 143
    v36 := GenS2{F0: 88}
    fmt.Println(v2)
    v34 -= (v34 << ((565) & 63))
    var v37 uint8 = (v5 + (v5 % (((v5 - (^(v5)))) | 1)))
    for fi38 := 0; fi38 < 19; fi38++ {
        for fi39 := 0; fi39 < 19; fi39++ {
            v40 := []int{fi39, v3, (fi39 >> ((65535) & 63)), (fi39 + (fi38 >> ((18446744073709551615) & 63)))}
            _ = len(v40)
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    fmt.Println(v4)
    var v41 uint32 = (v33 / (((v33 % (((v33 % ((893) | 1))) | 1))) | 1))
    var v42 uint32 = (v41 ^ v33)
    fmt.Println(v4, v33)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
