package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int8 }
type GenS2 struct { F0 uint8; F1 string; F2 bool }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int16; F1 string; F2 uint32 }
type GenS5 struct { F0 uint64; F1 string; F2 string }
func genFn0(p0 int8, p1 int, p2 int) int64 {
    v4 := []int{p1, 2147483647}
    p1 = (p1 + 2147483647)
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 float64 = ((float64((-2.25 - 1e10)) * (float64(0.1) + 1.5)) / (float64((float32(0.1) / 1.5)) + float64(3.14159)))
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    var v1 uint8 = 242
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    v7 := GenS1{F0: 786, F1: -128, F2: 42}
    var v8 string = "hello"
    if false {
        v1 -= 128
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int = -950
            var v11 uint64 = 42
            _ = fi9
            _ = v10
            _ = v11
        }
        v8 = ""
    } else {
        v12 := GenS1{F0: 189, F1: 100, F2: 0}
        v8 = ((v8 + ("\n\t" + (v8 + (v8 + (v8 + (v8 + "foo bar")))))) + v8)
        v13 := []int{0, 32768, -1}
        _ = v12
        _ = v13
    }
    var v14 uint64 = 666
    v8 = v8
    v15 := v3[65535]
    v15 = (-(v15))
    switch v8 {
    case "a":
        var v16 bool = (!true)
        _ = v16
    case "":
        var v17 uint8 = (v1 / (((v1 << ((65535) & 63))) | 1))
        _ = v17
    }
    var v18 float32 = (0.1 + 1e10)
    if false {
        if ((true && false) && false) {
            if true {
                fmt.Println(v15, v1)
            }
            var v19 uint16 = 105
            v18 = (v18 - (float32(1.5) * 3.14159))
            _ = v19
        } else {
            fmt.Println(v8)
            v14 += (v14 & v14)
            if (v14 <= 42) {
                var v20 int = (-(v15))
                if ((-2147483648 == -141) || false) {
                    v21 := GenS2{F0: 0, F1: (v8 + "foo bar"), F2: false}
                    var v22 uint8 = 1
                    _ = v21
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 17; fi23++ {
                        v20 -= 32767
                        _ = fi23
                    }
                    for fi24 := 0; fi24 < 20; fi24++ {
                        var v25 bool = false
                        v26 := GenS3{F0: v20, F1: ((v8 + "foo bar") + "foo bar")}
                        fmt.Println(v20, v15, v18)
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    v27 := []int{32768, (v20 % ((v20) | 1)), v20}
                    _ = v27
                }
                v20 = -1
                _ = v20
            } else {
                if false {
                    var v28 int32 = -2147483648
                    var v29 int16 = 1000
                    _ = v28
                    _ = v29
                } else {
                    if false {
                        switch -9223372036854775808 {
                        case 4294967295:
                            v30 := GenS2{F0: (^(v1)), F1: (v8 + "a"), F2: true}
                            _ = v30
                        case -9223372036854775808:
                            if true {
                                fmt.Println(v18)
                                switch v15 {
                                case 32768:
                                    fmt.Println(v15, v8, v1)
                                case 9223372036854775807:
                                    for fi31 := 0; fi31 < 17; fi31++ {
                                        var v32 uint32 = uint32(v1)
                                        _ = fi31
                                        _ = v32
                                    }
                                case -21:
                                    v1 = (v1 ^ (^(v1)))
                                }
                            }
                        default:
                            if false {
                                v33 := []int{}
                                fmt.Println(v15)
                                _ = v33
                            }
                        }
                        fmt.Println(v18, v1, v1)
                        fmt.Println(v1)
                    }
                }
                v34 := []int{(v15 % ((int(v1)) | 1)), v15}
                switch "a" {
                case "hello":
                    if false {
                        var v35 bool = ((v1 >> ((4294967295) & 63)) > v1)
                        if (v8 >= "") {
                            switch (v15 + (v15 << ((429) & 63))) {
                            case 381:
                                v36 := GenS4{F0: -166, F1: v8, F2: 2147483648}
                                _ = v36
                            }
                            for fi37 := 0; fi37 < 17; fi37++ {
                                var v38 uint = 0
                                var v39 string = "foo bar"
                                if false {
                                    for fi40 := 0; fi40 < 5; fi40++ {
                                        var v41 uint = (^(v38))
                                        _ = fi40
                                        _ = v41
                                    }
                                    _ = len(v3)
                                } else {
                                    for fi42 := 0; fi42 < 16; fi42++ {
                                        var v43 uint64 = 4294967295
                                        v15 = v15
                                        var v44 string = ""
                                        _ = fi42
                                        _ = v43
                                        _ = v44
                                    }
                                    v35 = true
                                }
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        }
                        _ = v35
                    }
                }
                _ = v34
            }
        }
        v45 := GenS4{F0: -517, F1: (v8 + ""), F2: 380}
        fmt.Println(v14)
        _ = v45
    }
    fmt.Println(v15, v18, v14)
    if (2147483647 == 524) {
        for fi46 := 0; fi46 < 11; fi46++ {
            switch (((v8 + ("" + "foo bar")) + "foo bar") + (((v8 + "") + ("a" + (v8 + "\n\t"))) + (v8 + v8))) {
            case "hello":
                fmt.Println(v18)
            default:
                var v47 bool = (false || (!(541 <= uint32(v1))))
                _ = v47
            }
            v15 += fi46
            _ = fi46
        }
        v48 := GenS5{F0: 42, F1: v8, F2: v8}
        v1 = v1
        _ = v48
    } else {
        var v49 int16 = -617
        v50 := GenS1{F0: 0, F1: -1, F2: 42}
        _ = v49
        _ = v50
    }
    v3[2] = (v15 + v15)
    for fi51 := 0; fi51 < 18; fi51++ {
        fmt.Println(v15, fi51)
        _ = fi51
    }
    for fi52 := 0; fi52 < 18; fi52++ {
        for fi53 := 0; fi53 < 1; fi53++ {
            v15 += 2147483647
            var v54 float32 = (0.1 * (v18 - 0.1))
            var v55 uint = 64
            _ = fi53
            _ = v54
            _ = v55
        }
        _ = fi52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
