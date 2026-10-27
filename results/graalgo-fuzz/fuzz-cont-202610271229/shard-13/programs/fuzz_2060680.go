package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint64 = 0
    v2 := []string{("hello" + "a")}
    v3 := []int{4294967295, -184}
    v4 := make(map[int]uint32)
    var v5 float64 = 3.14159
    if false {
        v5 += v5
        var v6 int8 = 100
        for fi7 := 0; fi7 < 17; fi7++ {
            fmt.Println(v5, v6)
            if true {
                var v8 uint64 = 2
                _ = v8
            }
            _ = fi7
        }
        _ = v6
    } else {
        var v9 uint = 255
        _ = v9
    }
    var v10 int16 = 925
    var v11 int32 = 0
    if true {
        switch 212 {
        case -9223372036854775808:
            var v12 string = "foo bar"
            _ = v12
        default:
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v11)
            case "":
                var v13 uint = 978
                _ = v13
            case "foo bar":
                if true {
                    var v14 float32 = 1.5
                    for fi15 := 0; fi15 < 9; fi15++ {
                        fmt.Println(v1, fi15, v14)
                        var v16 bool = (uint16(v1) > 0)
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                } else {
                    if true {
                        var v17 uint64 = (v1 | (v1 >> ((2) & 63)))
                        v4[32768] = 63
                        delete(v4, 1)
                        _ = v17
                    }
                    var v18 bool = false
                    _ = v18
                }
            default:
                if ((0 > 2147483648) || false) {
                    fmt.Println(v5)
                    var v19 int8 = 100
                    _ = v19
                }
            }
        }
        v20 := GenS1{F0: 1, F1: 1}
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 float32 = -2.25
            v5 -= (float64(v22) * (0.1 - v5))
            var v23 bool = false
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = v20
    } else {
        if false {
            delete(v4, 270)
            for fi24 := 0; fi24 < 19; fi24++ {
                switch "hello" {
                case "":
                    var v25 uint = 0
                    _ = v25
                case "\n\t":
                    v26 := GenS2{F0: 1}
                    _ = v26
                case "a":
                    switch fi24 {
                    case 428:
                        switch 755 {
                        case 4294967295:
                            fmt.Println(v11, v5, v1)
                        case -621:
                            var v27 uint16 = 255
                            _ = v27
                        case 84:
                            v28 := GenS3{F0: -20, F1: 0}
                            _ = v28
                        }
                    }
                default:
                    v1 += 470
                }
                _ = fi24
            }
        } else {
            v1 -= (v1 * v1)
            if ((v5 - 3.14159) == ((((v5 * ((v5 - (v5 + 1e10)) + v5)) / float64(0.1)) / v5) * v5)) {
                if true {
                    v5 = (v5 + v5)
                    v29 := v4[-586]
                    var v30 float64 = (v5 / 1.0)
                    _ = v29
                    _ = v30
                }
            } else {
                v31 := v4[32767]
                if true {
                    v11 = v11
                    var v32 float32 = -2.25
                    v10 = 33
                    _ = v32
                }
                if false {
                    for fi33 := 0; fi33 < 1; fi33++ {
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                        _ = fi33
                    }
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                    var v34 uint = 4294967295
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 14; fi35++ {
                        var v36 int16 = v10
                        _ = len(v3)
                        v2 = append(v2, "hello")
                        _ = fi35
                        _ = v36
                    }
                }
                _ = v31
            }
        }
    }
    var v37 string = "\n\t"
    v38 := GenS2{F0: 0}
    switch "hello" {
    case "":
        switch 0 {
        case 1:
            for fi39 := 0; fi39 < 13; fi39++ {
                v40 := GenS4{F0: (^(v11))}
                _ = fi39
                _ = v40
            }
        case 2147483647:
            var v41 uint32 = 4294967295
            _ = v41
        case -348:
            v42 := GenS4{F0: v11}
            _ = v42
        }
    }
    fmt.Println(v10, v5, v5)
    fmt.Println(v11, v5, v10)
    var v43 uint8 = 128
    v44 := GenS4{F0: (v11 & v11)}
    var v45 float32 = 0.0
    v1 -= v1
    v11 = v11
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
