package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 string; F1 uint32; F2 bool }
type GenS6 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS7 struct { F0 int; F1 int32 }
func genFn0() int {
    v7 := GenS1{F0: 116}
    if (1 < 1) {
        v8 := GenS1{F0: -1}
        var v9 int64 = -1
        v9 = v9
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS1{F0: 127}
            _ = fi10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS1{F0: -128}
        switch fi12 {
        case -9223372036854775808:
            switch "\n\t" {
            case "":
                var v14 float32 = (1e10 - 1.5)
                _ = v14
            case "\n\t":
                for fi15 := 0; fi15 < 17; fi15++ {
                    v16 := GenS2{F0: -128, F1: 127, F2: 2}
                    var v17 int8 = -74
                    v18 := GenS1{F0: (v17 ^ 0)}
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            default:
                var v19 bool = false
                _ = v19
            }
        }
        fmt.Println(fi12)
        _ = fi12
        _ = v13
    }
    var v20 int = 4294967295
    _ = v20
    _ = v7
    return (v20 | 258)
}

func main() {
    var v1 uint = 42
    v2 := []int{32767, 32768, 42, 127}
    v3 := make(map[uint32]string)
    var v4 string = (((("\n\t" + "a") + "hello") + "a") + "a")
    var v5 uint16 = 847
    var v6 float64 = float64(1e10)
    v21 := GenS2{F0: int8(v5), F1: 2147483647, F2: (v1 & (^(v1)))}
    var v22 uint = v1
    v4 = (v4 + "foo bar")
    switch (v4 + "a") {
    case "\n\t":
        for fi23 := 0; fi23 < 14; fi23++ {
            switch (fi23 + genFn0()) {
            case -1:
                if ((42 >= -1) && false) {
                    switch v4 {
                    case "a":
                        v6 = v6
                    case "foo bar":
                        for fi24 := 0; fi24 < 20; fi24++ {
                            if true {
                                v5 += (v5 << (((v1 / ((v22) | 1))) & 63))
                                if (!false) {
                                    v25 := GenS3{F0: -1, F1: 32767}
                                    switch (fi24 & fi24) {
                                    case 0:
                                        var v26 bool = (1 <= -1)
                                        _ = v26
                                    default:
                                        var v27 bool = (true && true)
                                        _ = v27
                                    }
                                    switch fi24 {
                                    case 64:
                                        if (false && true) {
                                            var v28 float64 = (v6 * (((0.1 * ((float64(1e10) + 0.0) + v6)) - float64(3.14159)) - (float64(1e10) / v6)))
                                            _ = v28
                                        }
                                    }
                                    _ = v25
                                }
                                switch (("hello" + v4) + v4) {
                                case "":
                                    v22 = (v1 / (((v22 % ((v1) | 1))) | 1))
                                case "foo bar":
                                    v29 := v3[729]
                                    _ = v29
                                default:
                                    if true {
                                        delete(v3, uint32(v22))
                                    }
                                }
                            }
                            switch ((v4 + ((v4 + "a") + (v4 + v4))) + v4) {
                            case "\n\t":
                                v30 := GenS4{F0: 2, F1: (uint16(v1) > 0)}
                                _ = v30
                            case "foo bar":
                                var v31 int32 = 1
                                _ = v31
                            case "a":
                                var v32 bool = (true || (v5 == 42))
                                _ = v32
                            }
                            if true {
                                switch (fi23 << ((42) & 63)) {
                                case -76:
                                    var v33 string = v4
                                    _ = v33
                                default:
                                    v1 += 668
                                }
                                var v34 int = (fi24 >> ((v1) & 63))
                                var v35 bool = (!true)
                                _ = v34
                                _ = v35
                            }
                            _ = fi24
                        }
                    default:
                        _ = len(v3)
                    }
                    if (!false) {
                        var v36 int64 = 42
                        v37 := GenS5{F0: "\n\t", F1: 362, F2: false}
                        for fi38 := 0; fi38 < 19; fi38++ {
                            var v39 bool = (true || (-32768 <= -174))
                            var v40 uint16 = uint16(v1)
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                        _ = v36
                        _ = v37
                    }
                    var v41 int32 = 0
                    _ = v41
                } else {
                    var v42 string = v4
                    _ = v42
                }
            case 927:
                v6 = (v6 * 1.5)
            }
            _ = fi23
        }
    case "a":
        v5 = 1
    }
    if true {
        var v43 float32 = (-2.25 / (((-2.25 * ((0.0 + 0.1) * -2.25)) / (-2.25 + 1e10)) * float32(((0.1 / (v6 / v6)) + v6))))
        fmt.Println(v43)
        _ = v43
    }
    _ = len(v3)
    v4 = v4
    fmt.Println(v22, v22, v1)
    var v44 int16 = 1000
    for fi45 := 0; fi45 < 5; fi45++ {
        v4 = v4
        _ = fi45
    }
    var v46 string = ""
    v47 := GenS2{F0: -90, F1: 149, F2: v22}
    delete(v3, 400)
    v48 := GenS6{F0: 550, F1: -2147483648, F2: 97}
    v6 = (v6 + (float64(0.0) / ((((0.0 + ((1e10 + v6) * v6)) * v6) + v6) - 0.1)))
    v49 := GenS7{F0: 42, F1: 42}
    var v50 float64 = v6
    fmt.Println(v46, v22)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
