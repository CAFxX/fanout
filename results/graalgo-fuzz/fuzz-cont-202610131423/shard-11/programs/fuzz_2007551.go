package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }
type GenS2 struct { F0 int64; F1 bool }
func genFn0() int64 {
    var v7 uint16 = 255
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint8 = 255
        fmt.Println(v7, v7)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 801, F1: (v7 < (^(v7))), F2: 0}
    fmt.Println(v7)
    _ = v10
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 bool, p1 uint, p2 uint) int {
    switch "" {
    case "hello":
        switch ("a" + "\n\t") {
        case "foo bar":
            if (false && true) {
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 uint32 = uint32(p1)
                    v12 -= (v12 & 42)
                    _ = fi11
                    _ = v12
                }
            } else {
                p2 = p2
                v13 := GenS1{F0: -819, F1: (2147483647 > 1), F2: 0}
                p2 -= 65535
                _ = v13
            }
        default:
            v14 := []int{743, 1}
            _ = v14
        }
    }
    var v15 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 1
}

func main() {
    var v1 int64 = 620
    var v2 float32 = 1.5
    v3 := []uint32{}
    v4 := make(map[int]bool)
    var v5 uint8 = 0
    var v6 float64 = 3.14159
    v5 = (v5 / (((^(v5))) | 1))
    var v16 int16 = 1000
    var v17 int8 = -105
    var v18 uint64 = 0
    switch "" {
    case "":
        v18 += v18
    case "hello":
        fmt.Println(v6, v5)
    default:
        switch "hello" {
        case "\n\t":
            for fi19 := 0; fi19 < 17; fi19++ {
                v6 -= v6
                if (264 == v18) {
                    if (!true) {
                        switch (("" + "foo bar") + ("hello" + "a")) {
                        case "a":
                            var v20 uint8 = v5
                            _ = v20
                        default:
                            if (!true) {
                                v2 = float32(v6)
                                switch fi19 {
                                case 127:
                                    var v21 bool = (363 < -1)
                                    _ = v21
                                case -2147483648:
                                    v22 := GenS1{F0: (fi19 & (-(fi19))), F1: true, F2: (v16 / (((v16 ^ (v16 - (v16 & v16)))) | 1))}
                                    _ = v22
                                case 0:
                                    v1 = (v1 * (v1 & v1))
                                }
                                delete(v4, -1)
                            }
                        }
                    }
                }
                fmt.Println(v6)
                _ = fi19
            }
        case "foo bar":
            v23 := GenS2{F0: 408, F1: true}
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        v16 = 1000
        v6 += v6
        v25 := GenS2{F0: genFn0(), F1: (v2 != float32(v6))}
        _ = fi24
        _ = v25
    }
    var v26 bool = ((42 > 18446744073709551615) && (("" + ("\n\t" + "hello")) <= "hello"))
    v3 = append(v3, 4294967295)
    if true {
        var v27 uint = 65535
        _ = v27
    } else {
        for fi28 := 0; fi28 < 20; fi28++ {
            if (4294967295 <= 42) {
                var v29 float32 = v2
                for fi30 := 0; fi30 < 9; fi30++ {
                    var v31 int = -963
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
            v32 := GenS2{F0: (-(v1)), F1: v26}
            _ = fi28
            _ = v32
        }
    }
    if len(v3) > 0 {
        v3[0] = uint32(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
