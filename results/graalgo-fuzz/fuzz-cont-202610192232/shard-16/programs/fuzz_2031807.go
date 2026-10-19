package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 bool; F1 int8 }
func genFn0(p0 int32) int16 {
    if (!(2 < 649)) {
        p0 = p0
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        switch (("foo bar" + "a") + "foo bar") {
        case "":
            switch fi7 {
            case 32767:
                v8 := []int{(fi7 ^ 9223372036854775807), fi7, (^(fi7)), (^(fi7))}
                _ = v8
            case 32768:
                switch "hello" {
                case "":
                    fmt.Println(fi7)
                }
            case -753:
                for fi9 := 0; fi9 < 11; fi9++ {
                    switch (("" + "\n\t") + ("\n\t" + "foo bar")) {
                    case "":
                        if false {
                            p0 += (p0 & (p0 / (((^(p0))) | 1)))
                            v10 := GenS1{F0: "a", F1: true}
                            if true {
                                p0 = 1
                                v11 := []int{}
                                switch (fi7 % ((-2147483648) | 1)) {
                                case -9223372036854775808:
                                    if false {
                                        if len(v11) > 0 {
                                            v11[0] = fi9
                                        }
                                    }
                                case 1:
                                    p0 = p0
                                default:
                                    v11 = append(v11, 42)
                                    fmt.Println(p0, fi7)
                                    v12 := GenS1{F0: "foo bar", F1: true}
                                    _ = v12
                                }
                                _ = v11
                            }
                            _ = v10
                        }
                    case "a":
                        var v13 float32 = ((-2.25 * (0.0 + (1.5 + float32((0.0 / -2.25))))) / (1.5 + 3.14159))
                        p0 = p0
                        var v14 uint16 = 65535
                        _ = v13
                        _ = v14
                    default:
                        var v15 uint64 = 469
                        _ = v15
                    }
                    var v16 bool = (18446744073709551615 >= 798)
                    if v16 {
                        p0 += -2147483648
                        p0 = (p0 & (p0 >> ((4294967295) & 63)))
                    } else {
                        if (18446744073709551615 > 287) {
                            var v17 string = "hello"
                            v18 := GenS2{F0: 32}
                            _ = v17
                            _ = v18
                        }
                        p0 = 42
                        v16 = ("hello" != "hello")
                        var v19 int = 1
                        _ = v19
                    }
                    _ = fi9
                    _ = v16
                }
            default:
                fmt.Println(fi7)
            }
        default:
            p0 = p0
        }
        switch "a" {
        case "foo bar":
            switch (("a" + (("a" + (("" + "\n\t") + "")) + ((("" + "foo bar") + "hello") + "hello"))) + "hello") {
            case "foo bar":
                v20 := GenS1{F0: "hello", F1: (p0 >= (^(p0)))}
                _ = v20
            case "\n\t":
                var v21 int64 = -879
                _ = v21
            case "":
                var v22 int64 = -1
                _ = v22
            default:
                var v23 int16 = 140
                _ = v23
            }
        }
        var v24 uint64 = 18446744073709551615
        _ = fi7
        _ = v24
    }
    _ = p0
    return 0
}

func main() {
    var v1 int16 = 107
    var v2 int64 = 0
    var v3 uint64 = 620
    v4 := make(map[uint32]string)
    var v5 float64 = 1.5
    var v6 string = "a"
    v25 := []int{-304, 42, 440}
    fmt.Println(v5, v5)
    v26 := GenS2{F0: 127}
    if true {
        switch -2147483648 {
        case 861:
            v27 := GenS2{F0: -1}
            _ = v27
        }
    }
    v28 := GenS2{F0: 119}
    v4[2] = v6
    fmt.Println(v1, v6)
    fmt.Println(v6, v6, v2)
    var v29 int = -9223372036854775808
    v30 := GenS3{F0: v29, F1: -2147483648}
    v1 = (^(v1))
    v6 = v6
    v29 -= 128
    if ((v6 <= "") && ((0.0 - ((3.14159 / (v5 * (v5 + v5))) + v5)) > (1.5 + ((v5 + ((float64(3.14159) / v5) / (float64(1e10) - 1.5))) + float64(1e10))))) {
        v5 = float64(0.1)
        v31 := GenS4{F0: v29, F1: 127}
        _ = v31
    }
    if (53 < (v3 << ((65535) & 63))) {
        v32 := GenS5{F0: true, F1: 42}
        _ = v32
    }
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    v33 := GenS5{F0: false, F1: 80}
    if true {
        var v34 int16 = genFn0(-2147483648)
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
