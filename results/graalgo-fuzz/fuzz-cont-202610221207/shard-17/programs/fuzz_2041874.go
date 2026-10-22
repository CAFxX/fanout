package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int; F2 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 int32; F2 string }
func genFn0() uint {
    v7 := GenS1{F0: "a", F1: 42, F2: false}
    v8 := GenS2{F0: 1}
    _ = v7
    _ = v8
    return 923
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 2147483647
    var v3 int32 = 1
    var v4 float32 = 0.0
    var v5 uint16 = 255
    v6 := make(map[string]int)
    v3 += v3
    switch "a" {
    case "":
        v5 = v5
    default:
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS2{F0: genFn0()}
            var v11 int = (fi9 + fi9)
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    if false {
        var v12 string = "hello"
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v13 int = 1
            _ = v13
        }
        _ = v12
    }
    fmt.Println(v2, v4)
    v4 -= (v4 - 0.1)
    var v14 int64 = 1
    v4 -= ((0.0 + (v4 + float32(0.1))) / v4)
    v15 := GenS3{F0: v5}
    for fi16 := 0; fi16 < 17; fi16++ {
        fmt.Println(v4, v5, v2)
        if (!(!(!(-283 < -805)))) {
            v6[("\n\t" + ("\n\t" + "foo bar"))] = (fi16 & (fi16 | (^(fi16))))
        } else {
            var v17 int = (fi16 << ((289) & 63))
            if (-1 >= 1) {
                if true {
                    v5 = (v5 / ((849) | 1))
                    v18 := GenS4{F0: (^(fi16)), F1: (fi16 + fi16), F2: 2147483648}
                    v19 := v1[(fi16 - fi16)]
                    _ = v18
                    _ = v19
                }
            }
            for fi20 := 0; fi20 < 2; fi20++ {
                if (103 <= 1) {
                    v3 -= (v3 / (((v2 + -176)) | 1))
                } else {
                    if true {
                        switch "\n\t" {
                        case "hello":
                            v3 = (v3 / ((v2) | 1))
                        case "":
                            v4 += v4
                        case "a":
                            var v21 int64 = (v14 >> ((4294967295) & 63))
                            var v22 int8 = -1
                            _ = v21
                            _ = v22
                        }
                        switch "" {
                        case "\n\t":
                            var v23 int32 = (v2 ^ (v3 ^ 1))
                            _ = v23
                        case "foo bar":
                            var v24 string = ("hello" + "")
                            _ = v24
                        }
                        if true {
                            v25 := GenS1{F0: "", F1: 42, F2: true}
                            v14 += v14
                            v26 := GenS5{F0: v5}
                            _ = v25
                            _ = v26
                        } else {
                            var v27 int = (fi20 | (fi16 ^ 32767))
                            _ = v27
                        }
                    }
                    fmt.Println(v4)
                    var v28 uint = 65535
                    _ = v28
                }
                fmt.Println(v5)
                if (genFn0() > 435) {
                    fmt.Println(v14, fi20, v3)
                    v29 := GenS6{F0: 255, F1: v3, F2: "a"}
                    v6[("foo bar" + ("\n\t" + "\n\t"))] = (fi16 << ((42) & 63))
                    _ = v29
                }
                _ = fi20
            }
            _ = v17
        }
        for fi30 := 0; fi30 < 18; fi30++ {
            if (42 == 42) {
                var v31 int8 = 0
                _ = v31
            }
            var v32 float32 = (-2.25 / v4)
            var v33 float64 = ((((0.0 * -2.25) / 3.14159) / 1.0) / 0.1)
            _ = fi30
            _ = v32
            _ = v33
        }
        _ = fi16
    }
    v34 := []int{-466, -28}
    var v35 int64 = (v14 & (v14 + 42))
    v2 += (v2 / (((v3 - (v3 << ((1) & 63)))) | 1))
    fmt.Println(v4, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
