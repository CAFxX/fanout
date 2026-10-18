package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int8; F1 int; F2 uint8 }

func main() {
    v1 := []string{("" + ""), ("a" + ("foo bar" + (("\n\t" + "a") + "a")))}
    var v2 int32 = -1
    v3 := make(map[uint32]int)
    var v4 uint64 = 1
    fmt.Println(v4, v2)
    var v5 string = "a"
    var v6 bool = (163 < 127)
    if (2 <= 23) {
        switch 0 {
        case 0:
            for fi7 := 0; fi7 < 15; fi7++ {
                v6 = v6
                for fi8 := 0; fi8 < 18; fi8++ {
                    fmt.Println(v4, fi8, v5)
                    v4 -= (v4 % (((v4 - (v4 ^ (v4 % (((v4 >> ((255) & 63))) | 1))))) | 1))
                    if v6 {
                        switch 32768 {
                        case 32767:
                            var v9 string = v5
                            _ = v9
                        case 42:
                            var v10 int32 = v2
                            _ = v10
                        case 128:
                            var v11 uint32 = 65535
                            var v12 bool = true
                            _ = v11
                            _ = v12
                        }
                    } else {
                        v4 -= 42
                        var v13 int32 = (v2 * (-(v2)))
                        switch v5 {
                        case "foo bar":
                            v14 := v3[0]
                            _ = v14
                        case "a":
                            _ = len(v1)
                        }
                        v15 := GenS1{F0: 784}
                        if ((127 != -128) || (!(true && ((!v6) || false)))) {
                            v16 := GenS1{F0: 2}
                            v4 = (v4 ^ (v4 % ((42) | 1)))
                            _ = v16
                        } else {
                            if true {
                                var v17 float64 = (-2.25 + 1.5)
                                _ = v17
                            } else {
                                if v6 {
                                    v2 = v13
                                }
                                if (v6 && v6) {
                                    v5 = ""
                                    fmt.Println(v13)
                                } else {
                                    v6 = v6
                                    v1 = append(v1, (("\n\t" + v5) + ("hello" + v5)))
                                    if false {
                                        switch fi7 {
                                        case 0:
                                            fmt.Println(fi8, fi7, v5)
                                        case 32768:
                                            v2 = 891
                                        default:
                                            v2 = v13
                                        }
                                        switch -308 {
                                        case 9223372036854775807:
                                            v18 := GenS2{F0: fi7, F1: fi7}
                                            _ = v18
                                        case -2147483648:
                                            if (93 < 1) {
                                                var v19 int8 = -82
                                                _ = v19
                                            }
                                        }
                                        if false {
                                            v20 := GenS1{F0: 222}
                                            _ = v20
                                        } else {
                                            if ((float32((1e10 / 3.14159)) + 3.14159) <= 1e10) {
                                                v21 := v3[2]
                                                _ = v21
                                            } else {
                                                fmt.Println(v13, v13)
                                                _ = len(v3)
                                            }
                                            v13 += v2
                                        }
                                        var v22 int64 = 42
                                        _ = v22
                                    } else {
                                        v5 = v5
                                        v6 = (((v5 + v5) + v5) <= "a")
                                        if (v6 && v6) {
                                            if (((!(!(!v6))) || (-2.25 == -2.25)) || true) {
                                                v6 = (42 < 1)
                                                v6 = (32767 != 1000)
                                                var v23 int32 = 1
                                                _ = v23
                                            }
                                            var v24 bool = v6
                                            if (v24 || (v6 || (0 > -128))) {
                                                if v24 {
                                                    v3[65535] = fi7
                                                    v4 += v4
                                                    var v25 uint8 = 2
                                                    _ = v25
                                                }
                                                v26 := GenS1{F0: 42}
                                                v27 := GenS3{F0: 1, F1: v24}
                                                _ = v26
                                                _ = v27
                                            }
                                            delete(v3, 1)
                                            v13 = (v2 * (^(v2)))
                                            _ = v24
                                        } else {
                                            v6 = v6
                                        }
                                    }
                                }
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        _ = v13
                        _ = v15
                    }
                    _ = fi8
                }
                _ = fi7
            }
        case 128:
            if (v6 || v6) {
                for fi28 := 0; fi28 < 4; fi28++ {
                    v29 := GenS2{F0: 507, F1: (fi28 | fi28)}
                    v30 := GenS4{F0: -1, F1: (fi28 | fi28), F2: 34}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
        }
        fmt.Println(v2, v2)
    }
    var v31 bool = (876 >= 2)
    fmt.Println(v2, v5, v4)
    v5 = v5
    v2 = v2
    var v32 uint16 = 255
    v33 := GenS1{F0: v32}
    v32 -= v32
    var v34 int32 = v2
    if true {
        v32 -= v32
    } else {
        v2 = -697
        v6 = (42 <= 42)
        fmt.Println(v31)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
