package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 string) int16 {
    v7 := []int{127, 4294967295, 554}
    for fi8 := 0; fi8 < 16; fi8++ {
        switch (("foo bar" + p0) + "hello") {
        case "hello":
            p0 = ""
        case "\n\t":
            var v9 int8 = 42
            _ = v9
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 1000
}

func main() {
    var v1 int16 = 1
    v2 := []uint32{454, 569, 0}
    var v3 uint16 = 2
    var v4 int16 = (v1 % (((v1 >> ((4294967295) & 63))) | 1))
    var v5 uint32 = 368
    var v6 uint32 = v5
    for fi10 := 0; fi10 < 14; fi10++ {
        if true {
            if false {
                v2 = append(v2, v5)
                for fi11 := 0; fi11 < 9; fi11++ {
                    if false {
                        if false {
                            var v12 uint16 = (v3 ^ (v3 & 905))
                            v13 := GenS1{F0: v6, F1: "hello"}
                            v12 -= (^(v12))
                            v3 += (v12 % (((v3 * (v12 + 2))) | 1))
                            v12 = 0
                            _ = v12
                            _ = v13
                        }
                        var v14 int64 = 0
                        _ = v14
                    } else {
                        fmt.Println(v3, fi10)
                        if false {
                            if len(v2) > 0 {
                                v2[0] = (^(v6))
                            }
                            var v15 uint8 = 1
                            _ = v15
                        }
                        switch "" {
                        case "":
                            switch 9223372036854775807 {
                            case -17:
                                switch fi10 {
                                case 32768:
                                    fmt.Println(fi10, v1, v4)
                                case 32767:
                                    v16 := GenS1{F0: 65535, F1: "a"}
                                    _ = v16
                                case 420:
                                    v4 = (v4 + -324)
                                }
                            default:
                                v2 = append(v2, 2)
                            }
                        }
                    }
                    fmt.Println(v3, v6)
                    _ = fi11
                }
                var v17 int32 = 1
                _ = v17
            }
            for fi18 := 0; fi18 < 11; fi18++ {
                switch 0 {
                case 32767:
                    v4 += v4
                }
                v19 := GenS2{F0: 1}
                _ = fi18
                _ = v19
            }
        } else {
            var v20 string = "\n\t"
            var v21 float32 = 1.5
            _ = v20
            _ = v21
        }
        _ = fi10
    }
    _ = len(v2)
    for fi22 := 0; fi22 < 5; fi22++ {
        fmt.Println(v4, fi22, v5)
        fmt.Println(v3)
        _ = fi22
    }
    var v23 float64 = float64(1.5)
    _ = len(v2)
    v24 := GenS3{F0: 939}
    var v25 uint32 = (v6 % ((v5) | 1))
    if ((v25 / ((v25) | 1)) > (v25 + v5)) {
        for fi26 := 0; fi26 < 7; fi26++ {
            fmt.Println(v3, v6)
            var v27 uint32 = (v25 >> ((357) & 63))
            _ = fi26
            _ = v27
        }
        switch "" {
        case "a":
            if true {
                fmt.Println(v3, v3)
                if (true || false) {
                    var v28 int64 = 1
                    v28 -= v28
                    _ = v28
                } else {
                    var v29 int16 = -221
                    v30 := GenS1{F0: (v6 | 65535), F1: "a"}
                    _ = v29
                    _ = v30
                }
                v3 = v3
            } else {
                if false {
                    fmt.Println(v1, v6, v25)
                }
                v3 = (v3 << ((42) & 63))
                switch "" {
                case "a":
                    v4 += (v4 & (v1 % ((v4) | 1)))
                case "":
                    var v31 int64 = 42
                    _ = v31
                }
            }
        default:
            v25 -= (v25 - v6)
        }
        v1 -= (v4 | v1)
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        for fi33 := 0; fi33 < 9; fi33++ {
            _ = len(v2)
            v25 = v5
            _ = fi33
        }
        for fi34 := 0; fi34 < 15; fi34++ {
            v25 = 65535
            if (-9223372036854775808 != 9223372036854775807) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v25)
            }
            _ = fi34
        }
        if len(v2) > 0 {
            v2[0] = v6
        }
        _ = fi32
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        var v36 uint32 = 1
        v37 := GenS2{F0: 255}
        for fi38 := 0; fi38 < 13; fi38++ {
            v39 := GenS4{F0: 42, F1: (v4 ^ (v1 / (((v1 | v1)) | 1))), F2: (v36 >> ((368) & 63))}
            if (35 != 42) {
                switch fi35 {
                case 2147483647:
                    v4 -= genFn0(("hello" + ("hello" + "foo bar")))
                default:
                    v2 = append(v2, v36)
                    v36 -= (v6 ^ v5)
                }
                fmt.Println(fi35)
            } else {
                var v40 int64 = 0
                _ = v40
            }
            _ = fi38
            _ = v39
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    v2 = append(v2, (v25 << ((65535) & 63)))
    var v41 int64 = 1
    v23 -= (((((v23 / 1.5) * v23) / v23) / v23) * 1.5)
    for fi42 := 0; fi42 < 12; fi42++ {
        for fi43 := 0; fi43 < 1; fi43++ {
            if true {
                var v44 float32 = (3.14159 * ((float32(0.0) / (float32(((1.5 + 1e10) / v23)) / (float32((v23 / (v23 + v23))) / 3.14159))) / 0.1))
                v6 -= v25
                _ = v44
            }
            _ = fi43
        }
        for fi45 := 0; fi45 < 1; fi45++ {
            v46 := GenS5{F0: v6}
            _ = fi45
            _ = v46
        }
        _ = fi42
    }
    v1 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
