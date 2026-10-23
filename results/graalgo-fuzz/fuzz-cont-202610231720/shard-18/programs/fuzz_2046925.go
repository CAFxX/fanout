package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int64 }
func genFn0(p0 int16, p1 int64, p2 uint64) int8 {
    fmt.Println(p0)
    fmt.Println(p2, p0, p2)
    p0 = p0
    v7 := []int{-2147483648, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -90
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -69
    var v3 float64 = (-2.25 - 3.14159)
    v4 := make(map[int]bool)
    var v5 uint32 = 1
    var v6 int = 42
    var v8 uint16 = 32767
    if true {
        if false {
            for fi9 := 0; fi9 < 14; fi9++ {
                if false {
                    v5 += v5
                } else {
                    switch ("hello" + "a") {
                    case "\n\t":
                        if false {
                            for fi10 := 0; fi10 < 11; fi10++ {
                                v3 = (v3 + (float64(-2.25) / -2.25))
                                if ((v2 % ((genFn0(-517, 9223372036854775807, 684)) | 1)) <= (v2 ^ (v2 % (((v2 ^ (v2 >> ((525) & 63)))) | 1)))) {
                                    v11 := GenS1{F0: true, F1: v6}
                                    var v12 float64 = float64(-2.25)
                                    var v13 uint = 0
                                    _ = v11
                                    _ = v12
                                    _ = v13
                                }
                                _ = fi10
                            }
                            switch "a" {
                            case "":
                                v14 := GenS1{F0: true, F1: (^(fi9))}
                                _ = v14
                            }
                            v15 := GenS1{F0: (!false), F1: fi9}
                            _ = v15
                        }
                    case "a":
                        v16 := v4[fi9]
                        _ = v16
                    case "":
                        switch "a" {
                        case "hello":
                            for fi17 := 0; fi17 < 11; fi17++ {
                                if (-1 >= 9223372036854775807) {
                                    v3 += float64((float32((0.0 + v3)) / 1.5))
                                    v4[0] = (!(true || (true && false)))
                                } else {
                                    var v18 uint16 = (^(v8))
                                    fmt.Println(v6)
                                    _ = v18
                                }
                                var v19 int8 = 42
                                _ = fi17
                                _ = v19
                            }
                        case "foo bar":
                            var v20 int64 = -1
                            _ = v20
                        case "":
                            v21 := GenS1{F0: true, F1: -9223372036854775808}
                            _ = v21
                        default:
                            v22 := GenS2{F0: 42}
                            _ = v22
                        }
                    }
                }
                v5 -= (v5 * 687)
                for fi23 := 0; fi23 < 5; fi23++ {
                    v24 := []int{(fi23 % (((fi23 - fi23)) | 1))}
                    v25 := GenS1{F0: (255 == uint8(v5)), F1: 42}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi9
            }
            var v26 uint = 2
            _ = v26
        }
    }
    var v27 int32 = 1
    for fi28 := 0; fi28 < 2; fi28++ {
        var v29 float32 = ((((float32((v3 - 0.1)) / (0.0 + 3.14159)) / (0.1 - (1e10 * 1.5))) / 0.1) + -2.25)
        switch "hello" {
        case "\n\t":
            var v30 uint64 = 1
            _ = v30
        case "a":
            v31 := GenS2{F0: 681}
            _ = v31
        default:
            var v32 int32 = 1
            _ = v32
        }
        v27 += (v27 - v27)
        _ = fi28
        _ = v29
    }
    var v33 float32 = -2.25
    v6 = 32768
    var v34 int16 = 42
    v35 := []int{(^(v6)), 9223372036854775807}
    fmt.Println(v3)
    if false {
        v33 += 1.5
        v27 = (^(v27))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
