package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int; F1 uint8; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 int32; F1 uint16; F2 uint64 }
func genFn0() int64 {
    var v4 uint8 = 2
    if true {
        var v5 int64 = 1
        v6 := GenS1{F0: (v5 << ((398) & 63)), F1: 1}
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: -466, F1: 0}
        var v8 uint32 = 2147483647
        v9 := GenS1{F0: 315, F1: 42}
        _ = v7
        _ = v8
        _ = v9
    }
    _ = v4
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"foo bar", "a", "hello", "\n\t"}
    var v3 string = ""
    var v10 float32 = -2.25
    v3 = ("\n\t" + "foo bar")
    v11 := v1[540]
    v11 = v11
    v12 := GenS1{F0: 168, F1: 147}
    v3 = "\n\t"
    if (v11 || true) {
        delete(v1, 4294967295)
    }
    var v13 uint8 = 42
    if false {
        for fi14 := 0; fi14 < 15; fi14++ {
            v13 = (^(v13))
            var v15 uint8 = 128
            v11 = (1e10 <= (3.14159 * 1.5))
            _ = fi14
            _ = v15
        }
        v3 = "a"
        v2 = append(v2, v3)
    } else {
        v10 -= (v10 * 1e10)
        if len(v2) > 0 {
            v2[0] = v3
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 int = fi16
            for fi18 := 0; fi18 < 16; fi18++ {
                switch fi16 {
                case 4294967295:
                    if (!((0 <= 4294967295) && ((1.5 * float64(1.5)) < -2.25))) {
                        var v19 float32 = (v10 * v10)
                        switch v3 {
                        case "hello":
                            switch 32768 {
                            case -2147483648:
                                v20 := v1[630]
                                _ = v20
                            case -1:
                                v21 := GenS1{F0: 0, F1: 599}
                                _ = v21
                            case 32767:
                                v22 := GenS2{F0: fi16, F1: v13, F2: ((((v3 + ((v3 + ((v3 + "a") + v3)) + "")) + v3) + v3) + "\n\t")}
                                _ = v22
                            default:
                                var v23 uint8 = v13
                                _ = v23
                            }
                        case "\n\t":
                            if true {
                                var v24 int16 = 42
                                _ = v24
                            } else {
                                var v25 uint16 = 65535
                                fmt.Println(v19)
                                if (!true) {
                                    var v26 int64 = 782
                                    v27 := GenS3{F0: 0}
                                    if ((v13 <= 0) && false) {
                                        delete(v1, 2147483647)
                                        var v28 int = fi16
                                        v28 = int(v13)
                                        _ = v28
                                    } else {
                                        delete(v1, 142)
                                    }
                                    _ = v26
                                    _ = v27
                                } else {
                                    if true {
                                        fmt.Println(v17, v10)
                                    } else {
                                        var v29 int64 = 9223372036854775807
                                        _ = v29
                                    }
                                }
                                _ = v25
                            }
                        }
                        var v30 float64 = (0.1 - 3.14159)
                        _ = v19
                        _ = v30
                    }
                case 9223372036854775807:
                    v17 += (-(v17))
                default:
                    if (1.5 > v10) {
                        fmt.Println(v11)
                        if v11 {
                            v13 -= v13
                            fmt.Println(v11, v3)
                        } else {
                            v13 = v13
                            v2 = append(v2, "a")
                            v31 := v1[1]
                            v10 -= v10
                            if v31 {
                                v11 = false
                            } else {
                                switch v3 {
                                case "":
                                    switch fi16 {
                                    case -767:
                                        v32 := GenS4{F0: 32767, F1: (fi16 - fi16)}
                                        if ((v17 - (v17 | 935)) <= fi18) {
                                            var v33 string = "a"
                                            _ = v33
                                        }
                                        _ = len(v1)
                                        var v34 uint64 = 18446744073709551615
                                        var v35 bool = v11
                                        _ = v32
                                        _ = v34
                                        _ = v35
                                    }
                                case "a":
                                    switch "foo bar" {
                                    case "hello":
                                        var v36 string = "a"
                                        v37 := GenS5{F0: v17, F1: 0}
                                        _ = v36
                                        _ = v37
                                    default:
                                        var v38 float64 = (-2.25 - 0.0)
                                        _ = v38
                                    }
                                default:
                                    var v39 int64 = -990
                                    _ = v39
                                }
                                if (-1 < -2147483648) {
                                    v40 := GenS6{F0: -2147483648, F1: 1, F2: 762}
                                    _ = v40
                                }
                            }
                            _ = v31
                        }
                        var v41 int64 = genFn0()
                        _ = v41
                    } else {
                        v13 = (v13 + v13)
                    }
                }
                var v42 int32 = 231
                _ = fi18
                _ = v42
            }
            _ = fi16
            _ = v17
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
