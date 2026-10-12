package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 uint32; F2 int32 }
func genFn0(p0 string) int {
    p0 = (p0 + p0)
    var v7 uint8 = 29
    var v8 int16 = 1000
    fmt.Println(v8, v8)
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint32) uint {
    fmt.Println(p0)
    var v9 uint32 = (p0 ^ p0)
    if ((-288 < 1) || true) {
        fmt.Println(p0, v9)
    }
    _ = p0
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 1
    var v3 string = "\n\t"
    var v4 uint32 = 65535
    var v5 string = ((((v3 + (v3 + v3)) + (("foo bar" + "foo bar") + v3)) + v3) + "hello")
    v6 := make(map[uint32]bool)
    v10 := []int{-1, 4294967295}
    delete(v1, v5)
    v11 := v6[1]
    v6[(v4 | 2147483648)] = (v11 || (733 <= 292))
    var v12 float32 = (3.14159 / 1e10)
    switch 261 {
    case -255:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    case 127:
        for fi13 := 0; fi13 < 8; fi13++ {
            v5 = v3
            var v14 string = (v3 + "")
            _ = fi13
            _ = v14
        }
    }
    v5 = "foo bar"
    switch -299 {
    case -505:
        switch 127 {
        case 32768:
            var v15 bool = v11
            _ = v15
        case 128:
            switch v5 {
            case "":
                if (!true) {
                    v12 -= v12
                    v2 = v2
                } else {
                    switch -9223372036854775808 {
                    case -2147483648:
                        v12 += v12
                    default:
                        if v11 {
                            switch -1 {
                            case -892:
                                switch ("foo bar" + v5) {
                                case "\n\t":
                                    v5 = "foo bar"
                                default:
                                    var v16 int = 1
                                    _ = v16
                                }
                            case 42:
                                v12 = ((float32((1e10 * -2.25)) / (v12 - (((v12 - v12) / v12) * v12))) * (v12 + ((-2.25 / v12) * v12)))
                            }
                            for fi17 := 0; fi17 < 5; fi17++ {
                                v10 = append(v10, 127)
                                _ = fi17
                            }
                        }
                    }
                    var v18 uint16 = 2
                    _ = v18
                }
            case "a":
                v19 := GenS1{F0: 8}
                _ = v19
            default:
                var v20 bool = false
                _ = v20
            }
        }
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(v11, v3, v5)
        if (!(!false)) {
            switch fi21 {
            case 127:
                var v22 uint8 = uint8(v4)
                _ = v22
            case -293:
                v3 = (v5 + v3)
            default:
                if ("\n\t" > ("" + ("a" + v3))) {
                    var v23 float32 = (0.1 / v12)
                    switch 32767 {
                    case 4294967295:
                        if false {
                            v5 = "a"
                        } else {
                            for fi24 := 0; fi24 < 7; fi24++ {
                                if false {
                                    switch "\n\t" {
                                    case "a":
                                        v3 = v5
                                    }
                                    v4 += 2147483647
                                    var v25 float64 = ((-2.25 * 3.14159) / (3.14159 + (3.14159 / 0.1)))
                                    if (0 >= 4294967295) {
                                        v26 := GenS2{F0: v2, F1: 0}
                                        delete(v6, v4)
                                        v3 = "a"
                                        _ = v26
                                    } else {
                                        switch ("" + v3) {
                                        case "":
                                            var v27 string = ""
                                            var v28 uint8 = (v2 + v2)
                                            _ = v27
                                            _ = v28
                                        case "hello":
                                            var v29 uint32 = v4
                                            _ = len(v10)
                                            _ = v29
                                        case "foo bar":
                                            v6[(v4 & (v4 % ((v4) | 1)))] = ((!(!(v11 || v11))) && (v11 && (-449 <= 1)))
                                        }
                                    }
                                    _ = v25
                                } else {
                                    v11 = v11
                                }
                                var v30 uint32 = 2147483647
                                v30 -= 2147483648
                                _ = fi24
                                _ = v30
                            }
                            for fi31 := 0; fi31 < 17; fi31++ {
                                v23 = float32(((-2.25 - 0.1) + 3.14159))
                                _ = fi31
                            }
                        }
                    case 127:
                        var v32 uint8 = 1
                        _ = v32
                    case 32768:
                        fmt.Println(v12, fi21, v5)
                    }
                    _ = v23
                } else {
                    if (1000 >= -1) {
                        fmt.Println(v5)
                        v3 = "hello"
                        var v33 uint32 = v4
                        _ = v33
                    }
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
            }
            var v34 int16 = 1
            _ = v34
        } else {
            v35 := GenS3{F0: 1}
            v12 += float32(1e10)
            for fi36 := 0; fi36 < 9; fi36++ {
                v12 += -2.25
                _ = fi36
            }
            _ = v35
        }
        _ = fi21
    }
    v37 := GenS4{F0: v2}
    v38 := GenS5{F0: -1, F1: (v4 / ((1) | 1)), F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
