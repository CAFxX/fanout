package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 bool; F2 int16 }
func genFn0() int {
    v5 := GenS1{F0: 127}
    v6 := GenS1{F0: 59}
    _ = v5
    _ = v6
    return 32768
}

func main() {
    var v1 uint8 = 60
    var v2 uint8 = v1
    var v3 int64 = 42
    var v4 float64 = 3.14159
    var v7 bool = false
    for fi8 := 0; fi8 < 20; fi8++ {
        v3 = v3
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 float32 = (1.5 / float32(v4))
            v11 := []int{-9223372036854775808, fi9, fi9}
            switch (fi9 | fi8) {
            case 32768:
                var v12 int16 = 0
                _ = v12
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        var v13 float64 = (v4 * v4)
        _ = fi8
        _ = v13
    }
    v3 = (v3 | v3)
    v14 := []int{9223372036854775807, 4294967295}
    if v7 {
        for fi15 := 0; fi15 < 14; fi15++ {
            v7 = (v7 && v7)
            var v16 int = (^(fi15))
            _ = fi15
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 int8 = 1
            switch (fi17 % ((-1) | 1)) {
            case 32767:
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            case 21:
                if ((true && (!false)) || v7) {
                    if v7 {
                        if v7 {
                            for fi19 := 0; fi19 < 3; fi19++ {
                                switch fi19 {
                                case -9223372036854775808:
                                    switch (fi19 & fi19) {
                                    case 301:
                                        v20 := GenS1{F0: v18}
                                        _ = v20
                                    }
                                case -1:
                                    var v21 int8 = v18
                                    _ = v21
                                default:
                                    var v22 string = "hello"
                                    _ = v22
                                }
                                _ = fi19
                            }
                            v23 := GenS1{F0: 0}
                            var v24 int8 = (v18 & v18)
                            _ = v23
                            _ = v24
                        } else {
                            v25 := GenS1{F0: v18}
                            _ = v25
                        }
                        var v26 int64 = v3
                        _ = len(v14)
                        _ = v26
                    } else {
                        switch fi17 {
                        case 1:
                            fmt.Println(v18)
                        case -2147483648:
                            v2 += 215
                        }
                    }
                    var v27 string = ("foo bar" + "\n\t")
                    _ = v27
                } else {
                    var v28 int32 = 530
                    var v29 string = "foo bar"
                    if ((!(!((fi17 >> ((4294967295) & 63)) == (fi17 << ((0) & 63))))) || v7) {
                        _ = len(v14)
                        var v30 int16 = 1
                        _ = v30
                    } else {
                        var v31 uint = 849
                        fmt.Println(v18)
                        _ = v31
                    }
                    _ = v28
                    _ = v29
                }
            default:
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
            _ = fi17
            _ = v18
        }
        v32 := GenS1{F0: 100}
        var v33 int64 = -1
        _ = v32
        _ = v33
    }
    v1 += v2
    for fi34 := 0; fi34 < 20; fi34++ {
        var v35 float32 = ((1e10 / float32(-2.25)) / float32((v4 + 1e10)))
        var v36 int8 = 31
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 float32 = (1e10 + float32(((v4 + v4) * float64((((((0.1 / (1.5 * -2.25)) + ((3.14159 / (1e10 * (0.1 / 3.14159))) / (-2.25 * 1.5))) * float32((1e10 / (((0.1 + float64(1e10)) + -2.25) + v4)))) - 0.0) / (1e10 / (((0.1 * 3.14159) * 3.14159) / 3.14159)))))))
    switch "hello" {
    case "hello":
        switch 1 {
        case 4294967295:
            var v38 uint64 = 42
            _ = v38
        default:
            v7 = false
        }
    case "foo bar":
        for fi39 := 0; fi39 < 18; fi39++ {
            v14 = append(v14, 1)
            v40 := GenS1{F0: -128}
            _ = fi39
            _ = v40
        }
    case "":
        v4 += v4
    }
    var v41 int = 42
    v42 := GenS1{F0: 47}
    for fi43 := 0; fi43 < 2; fi43++ {
        fmt.Println(v1, v1, v2)
        fmt.Println(v2, v41)
        _ = fi43
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        for fi45 := 0; fi45 < 1; fi45++ {
            if v7 {
                v46 := GenS2{F0: 642, F1: true, F2: -708}
                _ = v46
            }
            _ = fi45
        }
        fmt.Println(v41, v1, fi44)
        _ = fi44
    }
    v7 = (0 < -128)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v7
}
