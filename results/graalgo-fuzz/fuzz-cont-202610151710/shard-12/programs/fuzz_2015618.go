package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int; F1 int32; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int64; F2 uint8 }
func genFn0() string {
    var v7 int64 = -839
    var v8 int = 406
    _ = v7
    _ = v8
    return ("" + "a")
}

func main() {
    var v1 int8 = 0
    var v2 uint = 4294967295
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    var v5 uint32 = 556
    v6 := make(map[string]bool)
    _ = len(v4)
    for fi9 := 0; fi9 < 2; fi9++ {
        v2 = v2
        v10 := []int{(-(fi9)), 4294967295, fi9, (fi9 - 42)}
        for fi11 := 0; fi11 < 11; fi11++ {
            v2 = (v2 | v2)
            switch fi9 {
            case 128:
                var v12 int64 = -247
                _ = v12
            default:
                v13 := GenS1{F0: "a", F1: -2147483648}
                _ = v13
            }
            var v14 bool = true
            _ = fi11
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    var v15 int = 32768
    switch (v15 | (v15 / ((v15) | 1))) {
    case 9223372036854775807:
        for fi16 := 0; fi16 < 15; fi16++ {
            if true {
                switch -1 {
                case -210:
                    if true {
                        fmt.Println(v5)
                        for fi17 := 0; fi17 < 12; fi17++ {
                            v18 := []int{-220, (fi16 >> (((^(v2))) & 63)), (fi17 & (fi16 % (((fi17 & v15)) | 1))), (fi17 % ((-1) | 1))}
                            v19 := GenS1{F0: ((("" + "a") + "foo bar") + "\n\t"), F1: -1}
                            if (((0.0 * float32((3.14159 / ((0.1 * -2.25) * -2.25)))) * float32((1e10 * (1.5 * 3.14159)))) == -2.25) {
                                v2 += (^(v2))
                            }
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        v20 := GenS2{F0: fi16, F1: 0, F2: 1}
                        _ = v20
                    }
                case -9223372036854775808:
                    v1 = 1
                default:
                    var v21 uint32 = 65535
                    _ = v21
                }
            } else {
                v22 := GenS3{F0: 127}
                switch ("foo bar" + "a") {
                case "foo bar":
                    var v23 int8 = v1
                    _ = v23
                case "hello":
                    v24 := v4["hello"]
                    _ = v24
                default:
                    var v25 uint32 = v5
                    _ = v25
                }
                _ = v22
            }
            _ = fi16
        }
    case 4294967295:
        v26 := []int{v15}
        _ = v26
    default:
        var v27 int = v15
        _ = v27
    }
    fmt.Println(v2, v1, v2)
    var v28 uint32 = (v5 + (^(v5)))
    fmt.Println(v2, v15)
    fmt.Println(v5, v1)
    for fi29 := 0; fi29 < 5; fi29++ {
        for fi30 := 0; fi30 < 4; fi30++ {
            v28 = (v28 / (((v28 | v5)) | 1))
            _ = fi30
        }
        _ = fi29
    }
    if (-98 >= 32767) {
        _ = len(v4)
        if ("a" != "foo bar") {
            v31 := GenS4{F0: true, F1: -9223372036854775808, F2: 128}
            _ = v31
        } else {
            v28 -= 265
            for fi32 := 0; fi32 < 7; fi32++ {
                v28 += (v28 * 42)
                if ((v5 * (v28 / ((v28) | 1))) > v28) {
                    if (-9223372036854775808 >= 1) {
                        v2 -= 255
                    }
                } else {
                    v33 := v6["a"]
                    var v34 int16 = 943
                    _ = v33
                    _ = v34
                }
                _ = fi32
            }
            var v35 bool = false
            _ = v35
        }
        var v36 int16 = -1
        _ = v36
    } else {
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
