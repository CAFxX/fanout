package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    v2 := []int{-213, 2147483647, -900, 0}
    v3 := []int{128, 127}
    var v4 float64 = float64(1.5)
    var v5 int64 = 0
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            v5 += v5
            v8 := v1[""]
            delete(v1, "\n\t")
            if true {
                v9 := GenS1{F0: 1, F1: 65535}
                _ = v9
            } else {
                v10 := GenS1{F0: 65535, F1: 833}
                _ = v10
            }
            switch 9223372036854775807 {
            case 39:
                v4 += v4
            }
            _ = fi7
            _ = v8
        }
        v11 := GenS2{F0: 100, F1: 42}
        var v12 int16 = -32768
        _ = fi6
        _ = v11
        _ = v12
    }
    v4 = v4
    v1[""] = ("hello" + "a")
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 int64 = -1
        switch "\n\t" {
        case "a":
            v15 := GenS1{F0: 1, F1: 18446744073709551615}
            _ = v15
        default:
            v16 := GenS1{F0: 4294967295, F1: 171}
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    v5 = (^(v5))
    v17 := GenS1{F0: 303, F1: 65535}
    v4 += ((3.14159 / v4) - v4)
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 uint = 255
        fmt.Println(v5)
        for fi20 := 0; fi20 < 14; fi20++ {
            v4 = 1e10
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    v4 -= (v4 / (-2.25 / 0.1))
    v5 = (v5 + 9223372036854775807)
    if true {
        var v21 uint16 = 0
        if false {
            fmt.Println(v5)
        } else {
            for fi22 := 0; fi22 < 6; fi22++ {
                if false {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        v24 := GenS3{F0: v5}
                        v1["a"] = ""
                        if (!(!true)) {
                            fmt.Println(v5, v21)
                            if false {
                                switch "" {
                                case "foo bar":
                                    _ = len(v3)
                                default:
                                    switch "foo bar" {
                                    case "a":
                                        v21 += v21
                                    case "":
                                        v25 := GenS3{F0: (v5 + (v5 << ((65535) & 63)))}
                                        _ = v25
                                    default:
                                        v26 := GenS4{F0: 100, F1: 0}
                                        _ = v26
                                    }
                                }
                                if false {
                                    v27 := GenS4{F0: 1, F1: 4294967295}
                                    _ = v27
                                } else {
                                    v28 := GenS5{F0: 301}
                                    if (-2.25 > (v4 * 0.0)) {
                                        var v29 uint32 = 0
                                        _ = v29
                                    }
                                    _ = v28
                                }
                            } else {
                                var v30 float32 = 0.1
                                switch 1 {
                                case 152:
                                    _ = len(v3)
                                case 1:
                                    switch fi23 {
                                    case 42:
                                        v31 := GenS1{F0: 358, F1: 456}
                                        _ = v31
                                    case 127:
                                        var v32 string = "hello"
                                        _ = v32
                                    case 772:
                                        v33 := GenS1{F0: 42, F1: 0}
                                        _ = v33
                                    }
                                }
                                switch "hello" {
                                case "a":
                                    v30 += (float32((v4 + (((float64(((v30 / v30) - v30)) / v4) - v4) - (0.0 * (v4 / v4))))) + (v30 + 1e10))
                                default:
                                    v4 = (float64(3.14159) / v4)
                                }
                                _ = v30
                            }
                        }
                        _ = fi23
                        _ = v24
                    }
                }
                _ = len(v1)
                _ = fi22
            }
            v34 := GenS1{F0: 0, F1: uint(v21)}
            _ = v34
        }
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 int16 = -1
            v37 := GenS4{F0: 127, F1: 2147483648}
            _ = len(v3)
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = v21
    }
    var v38 string = "hello"
    for fi39 := 0; fi39 < 12; fi39++ {
        if len(v2) > 0 {
            v2[0] = fi39
        }
        v38 = v38
        _ = fi39
    }
    var v40 string = "foo bar"
    v40 = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v5
}
