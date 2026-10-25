package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int; F1 int8 }
func genFn0(p0 string, p1 string, p2 int8) int8 {
    v7 := GenS1{F0: 0, F1: p2}
    var v8 float32 = ((0.1 / 0.1) * 0.0)
    p0 = "foo bar"
    var v9 int64 = -1
    if false {
        var v10 int32 = -2147483648
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (p2 - 42)
}
func genFn1(p0 bool) bool {
    v11 := []int{127, 286, 32767, 4294967295}
    var v12 float32 = 0.1
    var v13 float32 = (0.1 / 1.0)
    var v14 uint32 = 4294967295
    v12 += 1e10
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return false
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 int64 = 0
    var v4 string = "hello"
    var v5 string = ""
    v6 := []string{(v4 + ""), ((((v5 + v5) + "\n\t") + v4) + (v5 + ((v5 + v5) + v5)))}
    if true {
        if ((v3 / ((v3) | 1)) < v3) {
            for fi15 := 0; fi15 < 10; fi15++ {
                switch -9223372036854775808 {
                case 548:
                    if (!(-1 > -351)) {
                        var v16 int64 = 1
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 15; fi17++ {
                            v4 = ("a" + "hello")
                            fmt.Println(v3, v5)
                            _ = fi17
                        }
                    }
                default:
                    _ = len(v6)
                }
                var v18 float32 = (float32(3.14159) / 3.14159)
                v3 = v3
                _ = fi15
                _ = v18
            }
            v5 = (v5 + (v4 + ((v5 + v4) + v4)))
            var v19 uint64 = 1
            _ = v19
        } else {
            v20 := GenS2{F0: 9223372036854775807, F1: (v3 ^ (-(v3)))}
            fmt.Println(v5, v5)
            _ = v20
        }
    } else {
        var v21 float64 = 0.0
        v22 := v1[2147483647]
        fmt.Println(v3)
        _ = v21
        _ = v22
    }
    v2[-119] = 127
    fmt.Println(v3, v4)
    var v23 bool = true
    var v24 uint8 = 0
    if false {
        v23 = v23
        if v23 {
            switch (v4 + "\n\t") {
            case "foo bar":
                switch 1 {
                case 9223372036854775807:
                    if (4294967295 != 727) {
                        fmt.Println(v4, v4)
                    } else {
                        v25 := GenS1{F0: v3, F1: -19}
                        _ = v25
                    }
                case 128:
                    var v26 int = 0
                    _ = v26
                default:
                    v27 := GenS2{F0: -9223372036854775808, F1: 9223372036854775807}
                    _ = v27
                }
            default:
                for fi28 := 0; fi28 < 7; fi28++ {
                    v4 = v4
                    var v29 int16 = 0
                    for fi30 := 0; fi30 < 13; fi30++ {
                        var v31 uint16 = 0
                        _ = fi30
                        _ = v31
                    }
                    _ = fi28
                    _ = v29
                }
            }
        }
    }
    switch "\n\t" {
    case "foo bar":
        v4 = v4
    case "":
        if true {
            var v32 string = (v5 + v5)
            _ = v32
        } else {
            for fi33 := 0; fi33 < 18; fi33++ {
                for fi34 := 0; fi34 < 16; fi34++ {
                    var v35 int64 = -635
                    var v36 uint = 0
                    v37 := GenS1{F0: (-(v3)), F1: 127}
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                _ = fi33
            }
            v4 = v5
            for fi38 := 0; fi38 < 10; fi38++ {
                fmt.Println(v5, v5, v4)
                _ = fi38
            }
        }
    case "hello":
        v39 := v1[100]
        _ = v39
    }
    if len(v6) > 0 {
        v6[0] = v4
    }
    var v40 float64 = ((1e10 / (((float64(0.1) + (0.1 - 0.1)) - float64(-2.25)) / float64(1.5))) / -2.25)
    var v41 int16 = -1
    v42 := GenS2{F0: -1, F1: (^(v3))}
    v43 := GenS3{F0: -314, F1: -1}
    var v44 uint32 = 858
    fmt.Println(v24, v44, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
