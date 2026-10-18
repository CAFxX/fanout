package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 0.0
    v3 := make(map[string]uint32)
    v4 := []int{32768, 2147483647}
    var v5 int64 = -589
    if false {
        fmt.Println(v2)
        delete(v3, "foo bar")
        for fi6 := 0; fi6 < 8; fi6++ {
            v5 -= v5
            var v7 int32 = 2147483647
            switch int(v7) {
            case -2147483648:
                var v8 int16 = 1
                _ = v8
            default:
                if true {
                    for fi9 := 0; fi9 < 9; fi9++ {
                        switch "foo bar" {
                        case "a":
                            v4 = append(v4, 32767)
                        default:
                            switch 9223372036854775807 {
                            case 287:
                                var v10 uint64 = 4294967295
                                _ = v10
                            case 127:
                                v4 = append(v4, (fi9 - (^(fi9))))
                            case 32768:
                                var v11 uint = 2
                                _ = v11
                            }
                        }
                        _ = fi9
                    }
                } else {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        var v13 int16 = -1
                        _ = fi12
                        _ = v13
                    }
                }
            }
            _ = fi6
            _ = v7
        }
    }
    var v14 int8 = -88
    v14 -= v14
    var v15 uint32 = 0
    for fi16 := 0; fi16 < 10; fi16++ {
        if false {
            if true {
                fmt.Println(v14, v14, v5)
                v17 := v1[2]
                _ = v17
            } else {
                v18 := GenS1{F0: "a"}
                switch "\n\t" {
                case "":
                    if true {
                        v19 := GenS2{F0: v5, F1: 703}
                        var v20 uint16 = 2
                        _ = v19
                        _ = v20
                    } else {
                        v2 = v2
                    }
                default:
                    v5 = (v5 % (((v5 + (v5 - (-(v5))))) | 1))
                }
                _ = v18
            }
        } else {
            if (0 >= 2) {
                v14 = (v14 & -128)
            }
            switch "\n\t" {
            case "\n\t":
                v21 := GenS1{F0: ""}
                _ = v21
            case "a":
                fmt.Println(v2, v15, v2)
            default:
                for fi22 := 0; fi22 < 17; fi22++ {
                    v15 -= v15
                    if (!(true && ((!(255 >= 65535)) || false))) {
                        v23 := GenS1{F0: ""}
                        var v24 int = 4294967295
                        var v25 int16 = -1
                        _ = v23
                        _ = v24
                        _ = v25
                    } else {
                        _ = len(v1)
                        v26 := GenS1{F0: ((("hello" + "\n\t") + "foo bar") + ("hello" + ("hello" + "")))}
                        var v27 int16 = 32767
                        v15 += v15
                        var v28 uint16 = 0
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    v2 -= v2
                    _ = fi22
                }
            }
        }
        v15 = 2147483647
        var v29 uint = 0
        _ = fi16
        _ = v29
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        switch fi30 {
        case 501:
            v5 = v5
        case 32767:
            fmt.Println(fi30, fi30, v2)
        }
        _ = fi30
    }
    if (!(1e10 == 0.0)) {
        var v31 uint32 = v15
        if true {
            v32 := v1[2147483648]
            var v33 int8 = v14
            var v34 int64 = -127
            _ = v32
            _ = v33
            _ = v34
        }
        var v35 int8 = 0
        _ = v31
        _ = v35
    } else {
        v2 += float64(-2.25)
        v2 = float64(1e10)
        if true {
            v4 = append(v4, 4294967295)
        }
    }
    if len(v4) > 0 {
        v4[0] = 32767
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v15 = 2147483647
    v15 += (^(v15))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
