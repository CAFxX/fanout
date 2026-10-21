package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 int8 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 float32 = (1.5 / float32((0.0 + (0.1 / 0.1))))
    var v2 int32 = 2147483647
    var v3 bool = (-1 > 1)
    var v4 int32 = v2
    v5 := make(map[string]int)
    var v6 float64 = (float64(1.5) + 1e10)
    fmt.Println(v6, v6, v6)
    var v7 string = ("hello" + "")
    switch -9223372036854775808 {
    case -383:
        var v8 uint16 = 65535
        _ = v8
    default:
        if (!true) {
            fmt.Println(v1, v6)
        } else {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 int64 = -1
                var v11 uint64 = 2
                _ = fi9
                _ = v10
                _ = v11
            }
        }
    }
    if (4294967295 > 4294967295) {
        var v12 string = (v7 + "\n\t")
        _ = v12
    } else {
        var v13 uint8 = 255
        _ = v13
    }
    v14 := GenS1{F0: 114}
    var v15 int16 = 42
    fmt.Println(v6, v7)
    if v3 {
        v16 := []int{298, -804, -613}
        _ = v16
    }
    _ = len(v5)
    v17 := []int{-2147483648, 32767}
    for fi18 := 0; fi18 < 15; fi18++ {
        v3 = (true || (!((v3 && (((v7 + (v7 + ("\n\t" + ""))) + v7) == v7)) && ((!false) || (((v1 / 1e10) == 1.5) && (!false))))))
        for fi19 := 0; fi19 < 8; fi19++ {
            fmt.Println(v1)
            _ = fi19
        }
        _ = fi18
    }
    var v20 uint16 = 255
    if ((v15 + v15) == (v15 / (((^(v15))) | 1))) {
        v15 = (v15 * (v15 | (v15 / ((v15) | 1))))
        var v21 bool = false
        _ = v21
    } else {
        switch 4294967295 {
        case 818:
            for fi22 := 0; fi22 < 2; fi22++ {
                for fi23 := 0; fi23 < 18; fi23++ {
                    if false {
                        if v3 {
                            var v24 uint64 = 2
                            var v25 int16 = (v15 & v15)
                            var v26 int = -231
                            _ = v24
                            _ = v25
                            _ = v26
                        } else {
                            v15 -= v15
                            v27 := v5[v7]
                            var v28 int8 = 100
                            _ = v27
                            _ = v28
                        }
                        v5[(((v7 + v7) + ((("" + ((v7 + (v7 + v7)) + v7)) + "hello") + (((v7 + (v7 + v7)) + v7) + v7))) + "foo bar")] = fi22
                    } else {
                        var v29 int32 = v2
                        _ = v29
                    }
                    var v30 bool = (-9223372036854775808 < -227)
                    var v31 uint32 = 4294967295
                    _ = fi23
                    _ = v30
                    _ = v31
                }
                _ = fi22
            }
        case -245:
            v5[(v7 + (v7 + v7))] = 128
        case 635:
            v32 := GenS1{F0: -128}
            _ = v32
        }
        v33 := GenS2{F0: 0, F1: 2147483647, F2: -61}
        _ = v33
    }
    switch v7 {
    case "":
        v4 = (v4 - v4)
    }
    v4 -= v2
    v34 := GenS3{F0: 165}
    v35 := GenS2{F0: 0, F1: 1, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
