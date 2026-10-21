package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint }
func genFn0(p0 string, p1 int8, p2 uint8) uint8 {
    p1 = (^(p1))
    var v7 float32 = 0.1
    switch -2147483648 {
    case 9223372036854775807:
        var v8 string = "a"
        _ = v8
    }
    v9 := GenS1{F0: 306}
    var v10 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v9
    return p2
}

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    v3 := []string{("hello" + "foo bar")}
    v4 := make(map[int]bool)
    var v5 int16 = 1
    var v6 uint16 = 32767
    switch "foo bar" {
    case "\n\t":
        v11 := GenS1{F0: 0}
        _ = v11
    default:
        var v12 int = -779
        _ = v12
    }
    var v13 uint = 18446744073709551615
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int64 = 9223372036854775807
        for fi16 := 0; fi16 < 14; fi16++ {
            if (!(v1 && v1)) {
                fmt.Println(v6)
                _ = len(v3)
            } else {
                switch "\n\t" {
                case "":
                    if v1 {
                        v17 := v2[(("hello" + "a") + "foo bar")]
                        _ = v17
                    }
                }
                v1 = v1
                var v18 uint = v13
                _ = v18
            }
            v5 = v5
            switch 42 {
            case -1:
                v6 = (v6 % ((0) | 1))
            case 0:
                v19 := GenS2{F0: 2, F1: 255, F2: -1}
                _ = v19
            case 42:
                if (v1 && v1) {
                    if v1 {
                        if (1 >= 299) {
                            var v20 int = 4294967295
                            _ = v20
                        } else {
                            fmt.Println(fi16)
                        }
                        v5 -= -1
                        var v21 float64 = 3.14159
                        _ = v21
                    } else {
                        var v22 int = fi14
                        var v23 uint16 = v6
                        fmt.Println(fi16, v6, v22)
                        v24 := GenS1{F0: (v13 / ((v13) | 1))}
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    if false {
                        if v1 {
                            fmt.Println(fi16, v13)
                        } else {
                            v25 := GenS2{F0: 255, F1: v13, F2: 1}
                            var v26 uint64 = 18446744073709551615
                            if v1 {
                                switch fi14 {
                                case -557:
                                    if ("hello" >= "\n\t") {
                                        v6 += (v6 & 42)
                                        v27 := GenS1{F0: v13}
                                        if len(v3) > 0 {
                                            v3[0] = ("hello" + ("\n\t" + "hello"))
                                        }
                                        _ = v27
                                    }
                                case -841:
                                    fmt.Println(v15, v1, fi14)
                                case 0:
                                    var v28 uint64 = (^(v26))
                                    _ = v28
                                default:
                                    v13 = 42
                                }
                                var v29 float64 = float64((0.1 * 0.1))
                                _ = v29
                            }
                            _ = v25
                            _ = v26
                        }
                        switch "" {
                        case "a":
                            v30 := GenS1{F0: (v13 % ((v13) | 1))}
                            fmt.Println(fi14, v5, v5)
                            _ = v30
                        case "":
                            v1 = true
                        case "\n\t":
                            v13 -= 867
                        default:
                            switch 4294967295 {
                            case 9223372036854775807:
                                v31 := GenS3{F0: 128, F1: v13, F2: v13}
                                _ = v31
                            case 946:
                                fmt.Println(v15, v6, v6)
                            }
                        }
                        fmt.Println(v6, v1)
                    }
                }
            default:
                delete(v4, fi14)
            }
            var v32 float64 = 0.0
            _ = fi16
            _ = v32
        }
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 int = 4294967295
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi33
            _ = v34
        }
        _ = fi14
        _ = v15
    }
    var v35 uint = 127
    var v36 int = 9223372036854775807
    v35 = v35
    delete(v2, "\n\t")
    for fi37 := 0; fi37 < 7; fi37++ {
        if len(v3) > 0 {
            v3[0] = (("\n\t" + "\n\t") + "a")
        }
        var v38 int8 = 1
        if (((!((true && (1.5 == -2.25)) || (1e10 == -2.25))) || true) && v1) {
            var v39 int16 = -582
            _ = v39
        }
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
