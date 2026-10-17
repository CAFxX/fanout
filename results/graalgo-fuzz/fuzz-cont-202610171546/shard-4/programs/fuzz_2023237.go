package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int8; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]uint32)
    v3 := []int{}
    v4 := make(map[int]int)
    delete(v2, 42)
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int8 = 0
        if (3 <= 1) {
            var v7 float64 = (1.5 / 1e10)
            for fi8 := 0; fi8 < 19; fi8++ {
                if true {
                    delete(v4, -1)
                } else {
                    v9 := GenS1{F0: fi8}
                    v6 += v6
                    _ = v9
                }
                var v10 uint64 = 2
                _ = fi8
                _ = v10
            }
            _ = v7
        } else {
            v1 -= 18446744073709551615
        }
        var v11 float32 = float32((1e10 - (-2.25 * 0.0)))
        _ = fi5
        _ = v11
        _ = v6
    }
    var v12 float32 = float32(1.5)
    v1 = (^(v1))
    var v13 string = "foo bar"
    for fi14 := 0; fi14 < 17; fi14++ {
        v1 -= v1
        if (!(v13 > "hello")) {
            for fi15 := 0; fi15 < 11; fi15++ {
                v1 -= v1
                v13 = (v13 + ("\n\t" + v13))
                if true {
                    fmt.Println(v12, v1, fi15)
                }
                fmt.Println(fi14, fi14)
                _ = fi15
            }
        }
        _ = fi14
    }
    v16 := v2[128]
    if (!(42 <= 32767)) {
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 uint32 = v16
            _ = len(v4)
            _ = fi17
            _ = v18
        }
        var v19 int32 = 368
        for fi20 := 0; fi20 < 19; fi20++ {
            v16 -= v16
            if (fi20 != (fi20 ^ (fi20 ^ 1))) {
                if true {
                    v21 := GenS1{F0: (fi20 * fi20)}
                    v22 := GenS2{F0: v13, F1: 100, F2: -920}
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v24 := GenS3{F0: v13}
                        v25 := GenS4{F0: 4294967295, F1: fi20}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = v21
                    _ = v22
                } else {
                    v16 = v16
                    switch int(v16) {
                    case 2147483647:
                        v26 := GenS2{F0: v13, F1: 100, F2: 32767}
                        _ = v26
                    default:
                        if false {
                            v4[-9223372036854775808] = fi20
                            var v27 int16 = -32768
                            v12 = v12
                            _ = v27
                        } else {
                            var v28 float64 = 1.5
                            _ = v28
                        }
                    }
                    for fi29 := 0; fi29 < 7; fi29++ {
                        v30 := GenS2{F0: v13, F1: 127, F2: -721}
                        var v31 uint8 = 2
                        fmt.Println(v19)
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                }
                switch fi20 {
                case 2147483647:
                    var v32 float32 = 1e10
                    _ = v32
                case 42:
                    var v33 int16 = -1
                    _ = v33
                default:
                    if false {
                        var v34 uint = 2
                        var v35 int8 = 0
                        _ = v34
                        _ = v35
                    } else {
                        if ((0.1 - v12) < (3.14159 * (v12 + (v12 - (1e10 * v12))))) {
                            var v36 int16 = 1000
                            _ = v36
                        }
                    }
                }
                for fi37 := 0; fi37 < 14; fi37++ {
                    if false {
                        v38 := GenS3{F0: v13}
                        _ = v38
                    }
                    _ = fi37
                }
            }
            if ((!false) && true) {
                if true {
                    var v39 int8 = -69
                    switch v13 {
                    case "":
                        fmt.Println(v13, v13, fi20)
                    case "\n\t":
                        var v40 int64 = 0
                        _ = v40
                    case "foo bar":
                        if false {
                            var v41 uint8 = 1
                            _ = v41
                        } else {
                            var v42 string = (("" + "foo bar") + "a")
                            _ = v42
                        }
                    default:
                        v16 = v16
                    }
                    _ = v39
                } else {
                    v16 -= 439
                    switch fi20 {
                    case -2147483648:
                        fmt.Println(v12, v1, v1)
                    default:
                        v19 += v19
                    }
                    v16 = 1
                }
                fmt.Println(v1, v12)
            }
            _ = fi20
        }
        _ = v19
    } else {
        var v43 int8 = -107
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v44 := GenS4{F0: 4294967295, F1: 4294967295}
        _ = v43
        _ = v44
    }
    v13 = "\n\t"
    var v45 bool = false
    var v46 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
}
