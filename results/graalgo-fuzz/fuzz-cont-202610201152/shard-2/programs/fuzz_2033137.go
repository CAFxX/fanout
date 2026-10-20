package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint32 }

func main() {
    v1 := []string{"a", "a"}
    v2 := make(map[int]int)
    var v3 string = "a"
    v4 := make(map[int]int)
    var v5 string = v3
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    if true {
        var v9 string = v5
        var v10 uint = 2
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int8 = 0
            _ = len(v1)
            if ((false || (false || (true && ((float32(-2.25) < float32(0.0)) && false)))) || ((v10 >> ((v10) & 63)) >= (v10 ^ (v10 - v10)))) {
                if ((fi11 & 878) < (fi11 >> (((v10 % (((v10 >> ((v10) & 63))) | 1))) & 63))) {
                    _ = len(v1)
                    v13 := GenS1{F0: false, F1: 2147483648, F2: 42}
                    fmt.Println(v12, v5)
                    _ = v13
                } else {
                    switch fi11 {
                    case 1:
                        v12 = (v12 - v12)
                    case 32767:
                        for fi14 := 0; fi14 < 15; fi14++ {
                            var v15 uint16 = 2
                            fmt.Println(v12, fi14, v12)
                            v16 := GenS1{F0: false, F1: 4294967295, F2: v15}
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    case -186:
                        v12 += v12
                    default:
                        v17 := GenS1{F0: ((962 <= uint32(v10)) && ((v9 + "a") > (v5 + ("foo bar" + ("foo bar" + v3))))), F1: 321, F2: 1}
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 8; fi18++ {
                        v10 -= (^(v10))
                        v19 := GenS1{F0: true, F1: 4294967295, F2: 267}
                        _ = fi18
                        _ = v19
                    }
                    if true {
                        fmt.Println(v12)
                        v3 = v5
                    } else {
                        v3 = (v9 + ("foo bar" + (v3 + ("\n\t" + "\n\t"))))
                    }
                }
                fmt.Println(v3, v3)
                v20 := GenS1{F0: true, F1: 2147483648, F2: 1}
                _ = v20
            } else {
                v21 := GenS1{F0: true, F1: 42, F2: 332}
                if false {
                    var v22 float64 = 3.14159
                    _ = v22
                } else {
                    fmt.Println(v10)
                }
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 int8 = (v12 * 63)
                    _ = fi23
                    _ = v24
                }
                _ = v21
            }
            _ = fi11
            _ = v12
        }
        _ = v10
        _ = v9
    }
    v3 = (v3 + v5)
    var v25 float32 = -2.25
    v3 = ("a" + v5)
    if false {
        _ = len(v1)
        var v26 uint64 = 0
        _ = v26
    }
    v5 = v3
    v1 = append(v1, ("" + ("" + "foo bar")))
    if (32768 <= 2147483647) {
        if (2 == 217) {
            v5 = v5
        }
    } else {
        for fi27 := 0; fi27 < 10; fi27++ {
            if (!false) {
                if (100 < -54) {
                    var v28 uint8 = 42
                    _ = v28
                } else {
                    if (127 == 219) {
                        var v29 int16 = -272
                        var v30 uint = 736
                        _ = v29
                        _ = v30
                    } else {
                        v3 = v5
                        v31 := GenS1{F0: false, F1: 234, F2: 32767}
                        v32 := GenS2{F0: 2147483647, F1: 128, F2: 65535}
                        _ = v31
                        _ = v32
                    }
                }
                fmt.Println(v25, v25)
            }
            _ = fi27
        }
        v33 := GenS3{F0: 766, F1: 897, F2: 2147483648}
        if true {
            fmt.Println(v5)
            _ = len(v1)
        }
        _ = v33
    }
    fmt.Println(v3)
    var v34 int = 4294967295
    if (!(v34 == 587)) {
        var v35 float32 = v25
        var v36 uint16 = 32767
        _ = v35
        _ = v36
    }
    v5 = (v5 + v3)
    var v37 float64 = 0.1
    var v38 string = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
