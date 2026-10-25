package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int, p1 bool, p2 uint8) uint {
    p1 = (0 > 2)
    var v4 uint32 = 98
    var v5 string = (("\n\t" + "foo bar") + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 255
}

func main() {
    var v1 uint = 273
    var v2 int8 = -45
    v3 := make(map[string]string)
    v2 = v2
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int = 952
        if (false || (false || (2 > 2))) {
            if false {
                switch "" {
                case "hello":
                    var v8 int8 = v2
                    _ = v8
                case "":
                    v9 := []int{9223372036854775807}
                    _ = v9
                case "foo bar":
                    v10 := GenS1{F0: 318}
                    _ = v10
                default:
                    var v11 uint64 = 1
                    _ = v11
                }
            } else {
                delete(v3, "foo bar")
            }
            fmt.Println(v7)
            if false {
                switch v7 {
                case 42:
                    fmt.Println(v1, v2, v2)
                case 1:
                    v12 := GenS1{F0: 624}
                    _ = v12
                default:
                    fmt.Println(fi6, fi6, v2)
                }
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 float64 = 1.5
                    v15 := v3["\n\t"]
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            } else {
                if ((675 >= 1) && (v1 != (v1 ^ (v1 * (v1 ^ v1))))) {
                    if ((float64(((0.0 - (-2.25 + (-2.25 + float32(3.14159)))) - 0.1)) - 1.5) == float64(1.5)) {
                        v1 += (v1 + v1)
                    } else {
                        v2 -= (v2 * v2)
                        v7 -= (fi6 - v7)
                    }
                    var v16 uint16 = 163
                    _ = v16
                }
                switch "a" {
                case "a":
                    v17 := GenS1{F0: 0}
                    _ = v17
                case "foo bar":
                    if (true || (true && false)) {
                        v18 := []int{32768}
                        for fi19 := 0; fi19 < 8; fi19++ {
                            var v20 uint = (v1 * genFn0(v7, (0.0 != (0.0 + 1.5)), 128))
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    }
                }
                fmt.Println(v2, fi6, v1)
            }
        } else {
            var v21 int16 = 0
            v22 := GenS1{F0: 712}
            if false {
                for fi23 := 0; fi23 < 8; fi23++ {
                    delete(v3, "\n\t")
                    var v24 int64 = 9223372036854775807
                    _ = fi23
                    _ = v24
                }
                switch "" {
                case "foo bar":
                    var v25 uint32 = 42
                    _ = v25
                case "":
                    var v26 int64 = int64(v7)
                    _ = v26
                case "hello":
                    var v27 string = (("a" + "\n\t") + "")
                    _ = v27
                }
            }
            _ = v21
            _ = v22
        }
        _ = fi6
        _ = v7
    }
    delete(v3, ("\n\t" + ("\n\t" + "hello")))
    v1 += 65535
    v28 := GenS2{F0: 133}
    v2 = (v2 / ((v2) | 1))
    for fi29 := 0; fi29 < 17; fi29++ {
        v2 -= (v2 - v2)
        var v30 int32 = 1
        _ = fi29
        _ = v30
    }
    switch 1 {
    case -878:
        v3["\n\t"] = ""
    case 95:
        v31 := []int{-9223372036854775808}
        _ = v31
    case 9223372036854775807:
        v32 := GenS3{F0: 885, F1: 42, F2: 2}
        _ = v32
    default:
        switch "a" {
        case "\n\t":
            fmt.Println(v2)
        case "hello":
            v2 = (v2 / (((v2 << ((v1) & 63))) | 1))
        }
    }
    v1 = (^(v1))
    v33 := []int{-9223372036854775808}
    var v34 int64 = -1
    var v35 float32 = ((1.5 * -2.25) + (3.14159 * (0.1 - ((3.14159 + (((-2.25 / -2.25) * -2.25) * -2.25)) / -2.25))))
    if true {
        if ((!true) || true) {
            for fi36 := 0; fi36 < 18; fi36++ {
                v37 := GenS4{F0: 2}
                if (false && true) {
                    if (false && true) {
                        for fi38 := 0; fi38 < 12; fi38++ {
                            fmt.Println(v2)
                            if false {
                                var v39 bool = true
                                if (v35 != float32(0.0)) {
                                    var v40 int8 = v2
                                    v2 = v40
                                    v1 += (v1 + (^(v1)))
                                    _ = v40
                                }
                                v41 := GenS5{F0: 702, F1: 1}
                                _ = v39
                                _ = v41
                            }
                            _ = fi38
                        }
                        v42 := GenS4{F0: 2147483647}
                        fmt.Println(v2, v2, fi36)
                        _ = v42
                    } else {
                        if true {
                            v43 := GenS2{F0: 133}
                            v44 := GenS2{F0: 128}
                            fmt.Println(v34, v35)
                            _ = v43
                            _ = v44
                        }
                    }
                }
                _ = fi36
                _ = v37
            }
            var v45 float32 = v35
            _ = v45
        } else {
            v2 = v2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
}
