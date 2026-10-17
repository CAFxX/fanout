package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint, p1 int, p2 bool) int {
    fmt.Println(p0, p1, p2)
    if (!p2) {
        var v4 uint = p0
        _ = v4
    } else {
        p2 = p2
    }
    var v5 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 int = 0
    v6 := []int{104}
    var v7 uint32 = 2
    v7 = v7
    v3 = v3
    v3 += v3
    if len(v6) > 0 {
        v6[0] = (v3 % ((genFn0(530, v3, ((3.14159 / 0.1) == 1.5))) | 1))
    }
    _ = len(v6)
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 string = "hello"
            var v11 uint16 = 2
            _ = fi9
            _ = v10
            _ = v11
        }
        if (!(!true)) {
            v12 := GenS1{F0: false, F1: -519}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 8; fi13++ {
                switch ((("\n\t" + ("hello" + "a")) + ("foo bar" + "a")) + "hello") {
                case "a":
                    if (true && true) {
                        if len(v6) > 0 {
                            v6[0] = 42
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    } else {
                        if (v7 == (v7 ^ v7)) {
                            v14 := GenS1{F0: false, F1: -1}
                            switch v3 {
                            case 127:
                                fmt.Println(v3, v7, v3)
                            default:
                                switch (-(fi8)) {
                                case 128:
                                    fmt.Println(v7, v3, v7)
                                case -2147483648:
                                    var v15 float64 = 0.1
                                    _ = v15
                                case 471:
                                    v7 = v7
                                default:
                                    var v16 string = ""
                                    _ = v16
                                }
                            }
                            v6 = append(v6, (^(fi13)))
                            _ = v14
                        } else {
                            _ = len(v6)
                            var v17 uint8 = 157
                            if false {
                                var v18 bool = (2 != 526)
                                if false {
                                    switch v3 {
                                    case 32768:
                                        if v18 {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        }
                                    case 127:
                                        var v19 uint64 = 18446744073709551615
                                        _ = v19
                                    default:
                                        var v20 float64 = 0.0
                                        _ = v20
                                    }
                                }
                                fmt.Println(v18, v7)
                                _ = v18
                            }
                            _ = v17
                        }
                    }
                case "hello":
                    switch "\n\t" {
                    case "\n\t":
                        if ((753 <= 1) && false) {
                            if true {
                                v21 := GenS2{F0: -1}
                                var v22 float64 = 3.14159
                                _ = v21
                                _ = v22
                            }
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    case "hello":
                        delete(v2, (fi8 << ((2) & 63)))
                    default:
                        v3 = fi8
                    }
                }
                _ = fi13
            }
            v7 += (v7 + 740)
        }
        _ = fi8
    }
    var v23 float64 = (((3.14159 * ((0.0 / -2.25) - float64(3.14159))) - 1.5) - 0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v6
    _ = v7
}
