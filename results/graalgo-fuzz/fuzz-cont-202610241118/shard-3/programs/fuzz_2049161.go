package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 uint8 = 42
    v2 := []uint32{0, 986}
    var v3 float64 = 0.0
    var v4 int16 = 42
    v5 := make(map[int]string)
    v6 := make(map[string]string)
    v7 := GenS1{F0: -2147483648}
    switch 1 {
    case -597:
        for fi8 := 0; fi8 < 10; fi8++ {
            switch (fi8 - fi8) {
            case 239:
                v1 -= 55
            case 1:
                if false {
                    v9 := GenS1{F0: 4294967295}
                    if true {
                        v1 = v1
                        switch fi8 {
                        case 42:
                            v4 -= v4
                        default:
                            v4 = -32768
                        }
                        v10 := GenS2{F0: 0, F1: 255, F2: 2147483647}
                        _ = v10
                    } else {
                        for fi11 := 0; fi11 < 8; fi11++ {
                            var v12 uint8 = (v1 + v1)
                            _ = fi11
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 3; fi13++ {
                            v14 := v6[("foo bar" + "")]
                            _ = fi13
                            _ = v14
                        }
                        fmt.Println(v4)
                    }
                    _ = v9
                }
            case 42:
                v4 -= v4
            }
            _ = fi8
        }
    case -476:
        v1 -= (v1 / (((v1 | v1)) | 1))
    case -9223372036854775808:
        if true {
            var v15 bool = false
            v16 := GenS3{F0: 1, F1: 0}
            switch -559 {
            case 747:
                v3 += -2.25
            case 0:
                var v17 string = ("a" + "hello")
                _ = v17
            default:
                v3 = v3
            }
            _ = v15
            _ = v16
        } else {
            var v18 bool = true
            if (32 != 610) {
                v1 += v1
                if v18 {
                    for fi19 := 0; fi19 < 18; fi19++ {
                        fmt.Println(v4)
                        fmt.Println(v4, v1, fi19)
                        _ = fi19
                    }
                    if (v3 <= 1.5) {
                        v1 = (v1 | v1)
                        fmt.Println(v1)
                        if (42 != -128) {
                            v4 -= -584
                            fmt.Println(v18, v4, v1)
                            for fi20 := 0; fi20 < 16; fi20++ {
                                v2 = append(v2, 42)
                                var v21 string = ""
                                _ = fi20
                                _ = v21
                            }
                        } else {
                            v18 = (4294967295 <= 18446744073709551615)
                            v3 += 0.0
                            fmt.Println(v3, v3)
                        }
                    } else {
                        fmt.Println(v18, v3)
                        v22 := GenS4{F0: 519}
                        _ = v22
                    }
                    var v23 string = ""
                    _ = v23
                }
                for fi24 := 0; fi24 < 20; fi24++ {
                    v25 := GenS4{F0: 2}
                    for fi26 := 0; fi26 < 18; fi26++ {
                        fmt.Println(v18, v3)
                        var v27 int16 = (-(v4))
                        _ = fi26
                        _ = v27
                    }
                    _ = fi24
                    _ = v25
                }
            } else {
                fmt.Println(v4)
                v3 -= v3
                switch "a" {
                case "hello":
                    for fi28 := 0; fi28 < 2; fi28++ {
                        if (v18 || (fi28 >= 32767)) {
                            var v29 int16 = 1
                            fmt.Println(v1, v18, v3)
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 9; fi30++ {
                            delete(v6, "a")
                            _ = fi30
                        }
                        _ = fi28
                    }
                case "a":
                    fmt.Println(v4, v4, v18)
                }
            }
            _ = v18
        }
    default:
        fmt.Println(v3, v4)
    }
    var v31 uint64 = 4294967295
    if false {
        var v32 uint8 = v1
        _ = v32
    } else {
        delete(v5, -2147483648)
        v3 += -2.25
    }
    fmt.Println(v4, v31, v1)
    var v33 int32 = 718
    v3 -= (v3 - (1e10 / (v3 / (float64(-2.25) - v3))))
    v31 += (^(v31))
    v34 := GenS2{F0: -128, F1: 32767, F2: v33}
    var v35 int64 = -789
    var v36 int8 = 100
    var v37 int64 = v35
    if true {
        var v38 bool = (v3 < ((v3 - float64((float32(v3) / float32((float64((((-2.25 / 3.14159) + 3.14159) + (3.14159 / 1.0))) - ((-2.25 - float64(-2.25)) * -2.25)))))) - v3))
        v39 := GenS2{F0: (v36 & 127), F1: 0, F2: 2147483647}
        _ = v38
        _ = v39
    } else {
        for fi40 := 0; fi40 < 6; fi40++ {
            var v41 uint64 = v31
            v42 := GenS3{F0: 0, F1: 1}
            _ = fi40
            _ = v41
            _ = v42
        }
        v43 := GenS1{F0: 32767}
        fmt.Println(v31, v36)
        _ = v43
    }
    v37 = 5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
