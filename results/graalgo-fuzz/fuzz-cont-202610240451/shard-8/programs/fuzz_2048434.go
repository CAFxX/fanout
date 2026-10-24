package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int32 }
func genFn0() uint {
    v7 := GenS1{F0: 305}
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int8 = -70
        _ = fi8
        _ = v9
    }
    var v10 uint = 0
    _ = v10
    _ = v7
    return v10
}
func genFn1() int64 {
    var v11 int = 750
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            if false {
                var v14 bool = (!true)
                v14 = v14
                if (418 < fi13) {
                    if (164 > 128) {
                        var v15 float32 = 0.0
                        if true {
                            v14 = ((v11 % ((v11) | 1)) >= v11)
                            fmt.Println(v15, v11)
                        } else {
                            v11 = 9223372036854775807
                        }
                        _ = v15
                    } else {
                        var v16 int32 = 175
                        _ = v16
                    }
                } else {
                    v17 := []int{(fi13 % ((42) | 1)), 206}
                    v18 := GenS1{F0: 1}
                    v17 = append(v17, v11)
                    var v19 int16 = 32767
                    switch (fi12 & 128) {
                    case 198:
                        v20 := GenS1{F0: 739}
                        _ = v20
                    case -1:
                        v14 = (((18446744073709551615 < 2) || (-2.25 <= 1e10)) && (1e10 <= -2.25))
                    case -198:
                        v19 = -558
                        v21 := GenS2{F0: (v11 | (v11 % (((v11 << ((42) & 63))) | 1))), F1: v19}
                        _ = v21
                    default:
                        if len(v17) > 0 {
                            v17[0] = (v11 % ((int(v19)) | 1))
                        }
                    }
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = v14
            } else {
                v11 = (fi13 / (((-(fi12))) | 1))
            }
            _ = fi13
        }
        _ = fi12
    }
    _ = v11
    return 876
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 1e10
    var v3 uint8 = 42
    var v4 string = "hello"
    v5 := []string{("a" + ((("\n\t" + (v4 + v4)) + v4) + ((v4 + ("\n\t" + ("\n\t" + v4))) + (v4 + (v4 + (v4 + "a")))))), (("hello" + "") + "a")}
    var v6 float64 = 3.14159
    switch "a" {
    case "\n\t":
        v4 = ("" + "hello")
    case "foo bar":
        v2 = v2
    case "hello":
        _ = len(v5)
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        if (568 >= 0) {
            fmt.Println(v4, v3)
            var v23 uint = 830
            _ = v23
        } else {
            fmt.Println(v4, fi22)
        }
        _ = fi22
    }
    if (float64((v2 * (v2 - (v2 + ((((0.0 + v2) - v2) / 1e10) / float32((((3.14159 - v6) / v6) * ((v6 * v6) * v6)))))))) < v6) {
        var v24 int32 = 42
        fmt.Println(v24, v3, v6)
        _ = len(v1)
        _ = v24
    }
    switch "" {
    case "hello":
        switch 128 {
        case 4294967295:
            fmt.Println(v3, v3, v4)
        }
    default:
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 int16 = 155
            _ = fi25
            _ = v26
        }
    }
    var v27 int16 = -32768
    if false {
        switch 82 {
        case 191:
            for fi28 := 0; fi28 < 12; fi28++ {
                fmt.Println(v6)
                v29 := GenS3{F0: 0, F1: 100}
                var v30 int = fi28
                _ = fi28
                _ = v29
                _ = v30
            }
        case 42:
            v4 = v4
        case 9223372036854775807:
            var v31 uint = 4294967295
            _ = v31
        }
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if len(v5) > 0 {
                v5[0] = (v4 + "")
            }
        }
        for fi32 := 0; fi32 < 4; fi32++ {
            v33 := v1[((v4 + ("" + ((v4 + v4) + (v4 + v4)))) + "foo bar")]
            var v34 int8 = 97
            _ = fi32
            _ = v33
            _ = v34
        }
    } else {
        v2 = v2
        if len(v5) > 0 {
            v5[0] = (v4 + "hello")
        }
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        _ = len(v5)
        v2 += (0.1 * v2)
        _ = fi35
    }
    v36 := GenS4{F0: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
