package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 string = "foo bar"
    var v3 uint8 = 176
    var v4 float64 = ((-2.25 * (0.0 * -2.25)) + (0.0 + 0.1))
    fmt.Println(v3, v3, v4)
    var v5 uint = 664
    v4 = v4
    v6 := GenS1{F0: -457}
    if true {
        var v7 int64 = -9223372036854775808
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 11; fi8++ {
            switch fi8 {
            case 1:
                v2 = (v2 + v2)
            }
            switch 32768 {
            case 128:
                v9 := []int{fi8}
                _ = v9
            case -143:
                v10 := GenS2{F0: -1}
                _ = v10
            }
            _ = fi8
        }
    }
    fmt.Println(v3, v3, v2)
    var v11 float32 = ((float32((0.0 * v4)) * (3.14159 * ((float32(float64(((0.1 / 1.5) - 1.5))) + float32((float64(0.0) - v4))) / 1.0))) * 3.14159)
    v4 -= (v4 - v4)
    var v12 int32 = -257
    v12 = (v12 % ((0) | 1))
    var v13 uint32 = 0
    v14 := GenS2{F0: 9223372036854775807}
    if true {
        v15 := GenS2{F0: 0}
        _ = v15
    } else {
        v11 -= v11
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := v1[v2]
            v18 := GenS3{F0: 1, F1: 42}
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 int8 = 0
        _ = v19
    }
    var v20 uint8 = (v3 << (((v5 >> (((^(v5))) & 63))) & 63))
    if (false || (v4 >= v4)) {
        for fi21 := 0; fi21 < 9; fi21++ {
            switch fi21 {
            case 128:
                switch (v2 + v2) {
                case "a":
                    if false {
                        v22 := GenS3{F0: 42, F1: -1}
                        _ = v22
                    } else {
                        v11 -= (1e10 - (1.5 + -2.25))
                        for fi23 := 0; fi23 < 12; fi23++ {
                            fmt.Println(fi23, v11, fi23)
                            _ = fi23
                        }
                        if (!(false || true)) {
                            var v24 uint = v5
                            _ = v24
                        } else {
                            switch "\n\t" {
                            case "hello":
                                if (!(128 <= v20)) {
                                    var v25 int = fi21
                                    v11 = v11
                                    var v26 uint = 113
                                    _ = v25
                                    _ = v26
                                } else {
                                    switch (fi21 & (fi21 | (fi21 & fi21))) {
                                    case 32767:
                                        for fi27 := 0; fi27 < 5; fi27++ {
                                            v4 += float64(float32(v4))
                                            var v28 bool = ((("" >= ("" + "a")) && true) || (v12 >= v12))
                                            _ = fi27
                                            _ = v28
                                        }
                                    case 127:
                                        var v29 int32 = (-(v12))
                                        _ = v29
                                    }
                                    v30 := v1["hello"]
                                    v31 := []int{32767, (fi21 & (^(fi21))), fi21, 127}
                                    _ = v30
                                    _ = v31
                                }
                            case "foo bar":
                                v32 := GenS3{F0: 64, F1: -32768}
                                _ = v32
                            }
                            v5 = (v5 / ((v5) | 1))
                        }
                    }
                case "\n\t":
                    switch (v2 + v2) {
                    case "a":
                        fmt.Println(v2, v2, v11)
                    }
                }
            case 4294967295:
                var v33 uint32 = 897
                _ = v33
            }
            _ = fi21
        }
        v34 := []int{794, 9223372036854775807, 9223372036854775807, 877}
        var v35 int = -669
        _ = v34
        _ = v35
    } else {
        var v36 string = ("" + v2)
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
