package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int16) bool {
    v4 := []int{2147483647}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    var v5 uint16 = 65535
    var v6 uint32 = 597
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (2147483647 <= 974)
}

func main() {
    v1 := []string{}
    var v2 uint16 = 0
    v3 := make(map[int]uint32)
    var v7 bool = (!true)
    var v8 float64 = (0.1 - -2.25)
    switch 32767 {
    case 32767:
        var v9 int32 = 508
        _ = v9
    case 433:
        if false {
            delete(v3, 127)
            if (-1 >= 127) {
                if v7 {
                    var v10 string = ""
                    switch -1 {
                    case 42:
                        v10 = "hello"
                    case -9223372036854775808:
                        switch v10 {
                        case "\n\t":
                            var v11 uint16 = 0
                            _ = v11
                        }
                    case 128:
                        if v7 {
                            if v7 {
                                _ = len(v3)
                                v12 := v3[4294967295]
                                _ = v12
                            } else {
                                v13 := GenS1{F0: "\n\t"}
                                var v14 int = 128
                                _ = v13
                                _ = v14
                            }
                        }
                    default:
                        var v15 uint64 = 2
                        _ = v15
                    }
                    _ = v10
                } else {
                    v16 := GenS1{F0: "\n\t"}
                    v17 := v3[-2147483648]
                    _ = v16
                    _ = v17
                }
                fmt.Println(v7, v7, v8)
                v2 -= 419
            } else {
                v2 += 255
            }
        } else {
            fmt.Println(v7, v7, v7)
            var v18 uint8 = 0
            _ = v18
        }
    }
    switch "foo bar" {
    case "a":
        v3[-1] = 609
    default:
        fmt.Println(v2, v2, v2)
    }
    switch "\n\t" {
    case "\n\t":
        v19 := GenS1{F0: ("a" + "")}
        _ = v19
    }
    switch "" {
    case "":
        for fi20 := 0; fi20 < 9; fi20++ {
            switch 4294967295 {
            case -9223372036854775808:
                for fi21 := 0; fi21 < 8; fi21++ {
                    if v7 {
                        fmt.Println(v8, fi20, v8)
                        var v22 bool = v7
                        _ = v22
                    } else {
                        fmt.Println(fi20)
                        var v23 float64 = ((float64((0.1 + 1.5)) / (((((float64((1.5 / float32(v8))) / v8) / v8) / v8) * v8) - v8)) * (v8 * 0.1))
                        v24 := GenS1{F0: "\n\t"}
                        _ = v23
                        _ = v24
                    }
                    _ = fi21
                }
            case 128:
                switch fi20 {
                case 42:
                    if len(v1) > 0 {
                        v1[0] = ("" + (("a" + "foo bar") + ""))
                    }
                default:
                    var v25 uint32 = 2147483648
                    _ = v25
                }
            case -1:
                v7 = false
            }
            for fi26 := 0; fi26 < 16; fi26++ {
                v27 := GenS1{F0: "foo bar"}
                switch "hello" {
                case "":
                    delete(v3, (fi26 & 128))
                default:
                    var v28 uint64 = 0
                    _ = v28
                }
                _ = fi26
                _ = v27
            }
            var v29 int32 = 697
            _ = fi20
            _ = v29
        }
    case "a":
        v30 := GenS1{F0: ("foo bar" + ("a" + "hello"))}
        _ = v30
    default:
        var v31 int = 32768
        _ = v31
    }
    if (v7 && ((float64(1.5) - 1e10) < (v8 + (v8 - v8)))) {
        for fi32 := 0; fi32 < 4; fi32++ {
            v33 := GenS1{F0: ("a" + "\n\t")}
            v34 := v3[fi32]
            for fi35 := 0; fi35 < 14; fi35++ {
                v36 := GenS1{F0: ("" + "foo bar")}
                if (("a" < "foo bar") && v7) {
                    v2 += v2
                }
                fmt.Println(v8)
                _ = fi35
                _ = v36
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        for fi37 := 0; fi37 < 20; fi37++ {
            v1 = append(v1, "a")
            _ = fi37
        }
    }
    v38 := v3[9223372036854775807]
    var v39 float64 = ((0.0 - (0.1 / (((((0.1 + (v8 * (v8 / 0.1))) / (1.5 * float64((-2.25 * 3.14159)))) + (float64((float32(v8) + 1.5)) * (float64((float32(0.0) / 3.14159)) - 1.5))) - -2.25) * (0.0 / v8)))) * float64((1e10 - 1.5)))
    var v40 int32 = 672
    var v41 float64 = v39
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    _ = v7
    _ = v8
}
