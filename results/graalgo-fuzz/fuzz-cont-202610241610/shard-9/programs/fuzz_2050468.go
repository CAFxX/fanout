package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint64, p1 string) uint8 {
    switch p1 {
    case "a":
        v4 := GenS1{F0: "hello"}
        _ = v4
    case "":
        p0 = (^(p0))
    default:
        p1 = "hello"
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"foo bar", "a", "hello"}
    var v3 bool = false
    v5 := GenS2{F0: 0, F1: 85}
    if v3 {
        switch ("foo bar" + ("foo bar" + "\n\t")) {
        case "\n\t":
            var v6 int8 = 127
            _ = v6
        case "":
            v3 = (255 < 2)
        case "hello":
            if (!((0.1 >= ((1.5 / -2.25) - 1.5)) && v3)) {
                if (-98 != 1) {
                    var v7 string = "\n\t"
                    _ = v7
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if v3 {
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                } else {
                    switch 0 {
                    case 4294967295:
                        v3 = (v3 && ((v3 && v3) || (!(1e10 < -2.25))))
                    default:
                        v3 = (false && (!(1 > 42)))
                    }
                    fmt.Println(v3)
                    if false {
                        for fi8 := 0; fi8 < 9; fi8++ {
                            if (!(float64(3.14159) == ((float64(float32((0.0 / ((-2.25 + (0.1 / 1.0)) / 1e10)))) - 1.5) - 0.1))) {
                                v3 = (127 <= -1)
                                if (2147483647 <= 2147483647) {
                                    var v9 int32 = -70
                                    _ = v9
                                }
                                switch fi8 {
                                case 9223372036854775807:
                                    if (!v3) {
                                        for fi10 := 0; fi10 < 1; fi10++ {
                                            var v11 uint64 = 0
                                            if true {
                                                if len(v2) > 0 {
                                                    v2[0] = "\n\t"
                                                }
                                                v12 := v1[fi10]
                                                v11 = v11
                                                v11 += v11
                                                switch "hello" {
                                                case "foo bar":
                                                    var v13 uint16 = 110
                                                    _ = v13
                                                }
                                                fmt.Println(fi10, v11, v12)
                                                _ = v12
                                            } else {
                                                v14 := GenS3{F0: 100, F1: -717}
                                                var v15 int16 = -375
                                                var v16 bool = (genFn0(684, "foo bar") >= genFn0(42, ("" + "\n\t")))
                                                var v17 uint = 218
                                                _ = v14
                                                _ = v15
                                                _ = v16
                                                _ = v17
                                            }
                                            _ = fi10
                                            _ = v11
                                        }
                                        if (255 > 866) {
                                            for fi18 := 0; fi18 < 9; fi18++ {
                                                switch "" {
                                                case "a":
                                                    v19 := GenS4{F0: 840, F1: 1}
                                                    _ = v19
                                                default:
                                                    var v20 int8 = 100
                                                    _ = v20
                                                }
                                                var v21 int32 = 491
                                                var v22 float64 = (0.1 * ((float64(((0.1 - ((3.14159 * 1e10) / 1.5)) - 0.1)) * ((-2.25 / -2.25) + float64(0.1))) + 3.14159))
                                                _ = fi18
                                                _ = v21
                                                _ = v22
                                            }
                                        }
                                    }
                                case -2147483648:
                                    fmt.Println(fi8, fi8)
                                }
                            } else {
                                v23 := GenS2{F0: 2, F1: 127}
                                _ = v23
                            }
                            v3 = (true && false)
                            _ = fi8
                        }
                        if (false && (v3 && v3)) {
                            var v24 bool = false
                            _ = v24
                        } else {
                            v3 = v3
                            fmt.Println(v3)
                        }
                        v3 = false
                    } else {
                        var v25 uint16 = 255
                        var v26 uint16 = v25
                        v25 = 255
                        _ = v25
                        _ = v26
                    }
                }
            }
        }
    }
    var v27 uint32 = 2
    var v28 int = -1
    v28 = (^(v28))
    var v29 int32 = 608
    for fi30 := 0; fi30 < 5; fi30++ {
        v31 := GenS3{F0: -58, F1: v29}
        _ = fi30
        _ = v31
    }
    switch "hello" {
    case "a":
        var v32 uint64 = 4294967295
        _ = v32
    case "hello":
        v3 = (("\n\t" + "foo bar") == "")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v5
}
