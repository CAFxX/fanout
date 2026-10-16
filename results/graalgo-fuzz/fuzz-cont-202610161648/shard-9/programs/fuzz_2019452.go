package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 4294967295
    var v3 float32 = float32((float64(((1.5 / 1.0) * 1e10)) / 3.14159))
    var v4 float32 = (v3 * float32(3.14159))
    v5 := GenS1{F0: -9223372036854775808, F1: false}
    fmt.Println(v4, v2, v2)
    if false {
        var v6 int = 865
        if true {
            v7 := GenS1{F0: 42, F1: true}
            fmt.Println(v3, v6, v3)
            _ = v7
        }
        v8 := GenS1{F0: 1, F1: (false || (42 > 255))}
        _ = v6
        _ = v8
    }
    var v9 float64 = float64((float32((float64(v4) - (0.0 * 1e10))) * ((v3 - 1e10) - (v4 + (float32(0.0) - ((1.5 / v3) * (v3 - v3)))))))
    v10 := []int{-27, 127}
    fmt.Println(v2, v2)
    if ("a" != ("\n\t" + "foo bar")) {
        for fi11 := 0; fi11 < 2; fi11++ {
            v9 = float64(v4)
            switch "hello" {
            case "":
                var v12 string = ("" + ("foo bar" + ("a" + "")))
                _ = v12
            case "hello":
                v13 := GenS1{F0: 0, F1: false}
                _ = v13
            case "\n\t":
                var v14 uint16 = 141
                _ = v14
            }
            _ = fi11
        }
        v4 -= 0.0
    }
    v2 -= (v2 % ((v2) | 1))
    v2 += (v2 - v2)
    var v15 bool = (false || false)
    switch 1 {
    case -96:
        v16 := GenS1{F0: -1, F1: ((("" + ((("" + "\n\t") + "a") + "\n\t")) + "\n\t") != "a")}
        _ = v16
    default:
        for fi17 := 0; fi17 < 17; fi17++ {
            if false {
                delete(v1, "hello")
            } else {
                if (!(520 <= fi17)) {
                    v4 += (v4 - (v3 + v4))
                } else {
                    v3 = v3
                    v18 := v1[""]
                    var v19 uint32 = 0
                    _ = v18
                    _ = v19
                }
                v20 := GenS1{F0: -1, F1: false}
                switch (fi17 | (-(fi17))) {
                case 9223372036854775807:
                    if false {
                        v1["foo bar"] = 1
                        var v21 uint8 = 1
                        var v22 uint8 = (v21 ^ (v21 - v21))
                        _ = v21
                        _ = v22
                    } else {
                        var v23 float64 = 0.1
                        if true {
                            for fi24 := 0; fi24 < 7; fi24++ {
                                v25 := GenS2{F0: 127}
                                v26 := GenS2{F0: 145}
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            if v15 {
                                if (300 >= 2) {
                                    var v27 uint32 = 2
                                    _ = v27
                                } else {
                                    fmt.Println(v3, v2, v23)
                                    switch fi17 {
                                    case 919:
                                        fmt.Println(v23, v4)
                                    default:
                                        for fi28 := 0; fi28 < 8; fi28++ {
                                            v9 = -2.25
                                            switch 9223372036854775807 {
                                            case 127:
                                                if len(v10) > 0 {
                                                    v10[0] = (fi17 * fi28)
                                                }
                                            case -1:
                                                v4 = v4
                                            case 0:
                                                if (v15 || true) {
                                                    if true {
                                                        v1["hello"] = fi17
                                                    }
                                                    v3 -= float32((v23 / (3.14159 + 1e10)))
                                                } else {
                                                    switch 198 {
                                                    case 563:
                                                        v29 := GenS1{F0: -9223372036854775808, F1: v15}
                                                        _ = v29
                                                    case 747:
                                                        fmt.Println(v2, v4)
                                                    default:
                                                        _ = len(v10)
                                                    }
                                                    _ = len(v10)
                                                }
                                            default:
                                                v30 := GenS2{F0: 127}
                                                _ = v30
                                            }
                                            _ = fi28
                                        }
                                    }
                                    v2 += 497
                                }
                                for fi31 := 0; fi31 < 3; fi31++ {
                                    fmt.Println(v3, fi31)
                                    v23 = v9
                                    _ = fi31
                                }
                                var v32 uint8 = 42
                                _ = v32
                            }
                            var v33 uint64 = (v2 << ((819) & 63))
                            _ = v33
                        } else {
                            if v15 {
                                var v34 bool = (2 != 65535)
                                _ = v34
                            }
                            v1["hello"] = (fi17 % ((fi17) | 1))
                        }
                        _ = v23
                    }
                }
                _ = v20
            }
            v35 := GenS1{F0: -1, F1: v15}
            if false {
                v36 := GenS3{F0: 1}
                fmt.Println(v3)
                _ = v36
            }
            _ = fi17
            _ = v35
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
