package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint64 = 89
    var v2 string = "hello"
    v3 := make(map[int]uint32)
    v2 = (v2 + (v2 + v2))
    var v4 uint32 = 42
    v4 += v4
    var v5 uint64 = (^(v1))
    delete(v3, 32768)
    if (127 >= 255) {
        if true {
            for fi6 := 0; fi6 < 13; fi6++ {
                v5 -= v5
                switch (fi6 >> ((0) & 63)) {
                case 127:
                    v5 = (v5 ^ v5)
                default:
                    v7 := []int{fi6, fi6, -388}
                    _ = v7
                }
                _ = fi6
            }
            fmt.Println(v2, v2)
        }
        if true {
            var v8 int64 = 251
            if (true && (!false)) {
                for fi9 := 0; fi9 < 6; fi9++ {
                    for fi10 := 0; fi10 < 11; fi10++ {
                        v5 += v5
                        v4 -= v4
                        _ = fi10
                    }
                    var v11 uint16 = 2
                    _ = fi9
                    _ = v11
                }
            } else {
                var v12 float64 = 1e10
                _ = v12
            }
            _ = v8
        } else {
            v4 = v4
        }
        v13 := GenS1{F0: 0, F1: 459}
        _ = v13
    }
    var v14 bool = true
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 int16 = -1
        var v17 bool = v14
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 int8 = 17
    switch 396 {
    case -2147483648:
        if v14 {
            fmt.Println(v18, v14, v1)
            if ((4294967295 == 0) || v14) {
                if v14 {
                    switch (v2 + ("hello" + v2)) {
                    case "\n\t":
                        for fi19 := 0; fi19 < 20; fi19++ {
                            if v14 {
                                for fi20 := 0; fi20 < 8; fi20++ {
                                    fmt.Println(fi19, fi20)
                                    _ = fi20
                                }
                            } else {
                                v2 = ((v2 + v2) + "a")
                            }
                            _ = fi19
                        }
                    case "":
                        var v21 uint8 = 129
                        _ = v21
                    default:
                        var v22 uint8 = 42
                        _ = v22
                    }
                } else {
                    for fi23 := 0; fi23 < 7; fi23++ {
                        var v24 uint = 18446744073709551615
                        var v25 uint = v24
                        v25 += v24
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    v26 := GenS2{F0: 128, F1: 1}
                    var v27 uint = 0
                    _ = v26
                    _ = v27
                }
                if v14 {
                    for fi28 := 0; fi28 < 19; fi28++ {
                        for fi29 := 0; fi29 < 19; fi29++ {
                            fmt.Println(v18, v2, v4)
                            var v30 uint32 = (v4 ^ 119)
                            _ = fi29
                            _ = v30
                        }
                        switch fi28 {
                        case -210:
                            v1 = v1
                        default:
                            for fi31 := 0; fi31 < 8; fi31++ {
                                switch (fi31 * 32767) {
                                case -2147483648:
                                    switch (((v2 + "hello") + (("hello" + v2) + ("foo bar" + v2))) + "") {
                                    case "foo bar":
                                        v32 := GenS2{F0: 1, F1: -32768}
                                        _ = v32
                                    default:
                                        switch (("foo bar" + "") + ("\n\t" + (v2 + v2))) {
                                        case "a":
                                            delete(v3, (fi28 + fi31))
                                        }
                                    }
                                }
                                _ = fi31
                            }
                        }
                        _ = fi28
                    }
                    fmt.Println(v14, v4, v4)
                }
            }
        } else {
            for fi33 := 0; fi33 < 6; fi33++ {
                var v34 int64 = 9223372036854775807
                var v35 int8 = v18
                _ = fi33
                _ = v34
                _ = v35
            }
        }
    default:
        fmt.Println(v2)
    }
    fmt.Println(v14, v5)
    switch (v2 + v2) {
    case "foo bar":
        v36 := GenS1{F0: 1, F1: 42}
        _ = v36
    case "hello":
        v2 = ""
    default:
        switch 811 {
        case 598:
            v37 := GenS3{F0: 1000}
            _ = v37
        }
    }
    var v38 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
