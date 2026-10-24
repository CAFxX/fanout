package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 bool }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int64 }

func main() {
    var v1 string = ("foo bar" + "foo bar")
    var v2 bool = (true && (0.1 != float32(-2.25)))
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    fmt.Println(v1, v2)
    var v5 uint32 = 2147483647
    var v6 uint16 = 599
    switch 1 {
    case 593:
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint64 = 373
            v6 += (v6 >> ((255) & 63))
            if true {
                v1 = ((v1 + "foo bar") + v1)
                switch v1 {
                case "":
                    var v9 float32 = 1e10
                    _ = v9
                case "a":
                    v10 := GenS1{F0: (v8 ^ (v8 % (((^(v8))) | 1))), F1: (fi7 ^ (fi7 * 128)), F2: false}
                    _ = v10
                case "hello":
                    v6 += (v6 | v6)
                }
            } else {
                if (1 >= 987) {
                    v2 = (v8 < (v8 % ((v8) | 1)))
                    switch 4294967295 {
                    case 9223372036854775807:
                        for fi11 := 0; fi11 < 14; fi11++ {
                            var v12 uint64 = (v8 + (v8 % (((v8 << ((766) & 63))) | 1)))
                            switch "a" {
                            case "\n\t":
                                var v13 int64 = 1
                                _ = v13
                            case "":
                                if (-1 == 0) {
                                    switch "hello" {
                                    case "a":
                                        var v14 bool = ((!true) || (!(v2 || false)))
                                        _ = v14
                                    case "hello":
                                        fmt.Println(fi7)
                                    case "":
                                        if (false && (v2 || (0.0 <= (0.0 / (3.14159 + 3.14159))))) {
                                            v15 := GenS1{F0: v12, F1: (fi7 - fi7), F2: (("\n\t" + "a") >= "a")}
                                            _ = v15
                                        }
                                        v6 = v6
                                        switch fi11 {
                                        case 32767:
                                            v4[("foo bar" + (v1 + v1))] = ((v8 & (v8 >> ((330) & 63))) > v12)
                                        case 0:
                                            v6 = v6
                                        }
                                    }
                                    var v16 float64 = float64(0.1)
                                    _ = v16
                                }
                            default:
                                v17 := GenS1{F0: 0, F1: fi7, F2: v2}
                                _ = v17
                            }
                            _ = fi11
                            _ = v12
                        }
                    }
                    var v18 bool = true
                    _ = v18
                }
                for fi19 := 0; fi19 < 19; fi19++ {
                    if v2 {
                        switch v1 {
                        case "foo bar":
                            fmt.Println(fi19)
                        }
                    }
                    if v2 {
                        v20 := GenS2{F0: (v8 ^ (v8 / ((v8) | 1))), F1: 42}
                        _ = v20
                    } else {
                        v3[65535] = "hello"
                        v1 = v1
                        if (v5 > (v5 & (^(v5)))) {
                            v1 = (v1 + (v1 + v1))
                            if v2 {
                                v5 -= uint32(v6)
                            }
                        }
                    }
                    _ = fi19
                }
                v21 := v4[v1]
                _ = v21
            }
            _ = fi7
            _ = v8
        }
    case 501:
        v22 := []int{1, -9223372036854775808, -15}
        _ = v22
    }
    if (v2 || (v6 <= v6)) {
        var v23 bool = false
        _ = v23
    } else {
        v1 = "hello"
    }
    var v24 float64 = float64(1.5)
    v25 := GenS3{F0: -165}
    switch 128 {
    case 6:
        _ = len(v3)
    case 127:
        switch ("hello" + ("foo bar" + "")) {
        case "a":
            var v26 uint8 = 127
            _ = v26
        case "hello":
            v1 = v1
        }
    }
    var v27 int64 = 9223372036854775807
    var v28 float32 = (float32((v24 / v24)) + -2.25)
    v6 = v6
    switch v1 {
    case "":
        var v29 int64 = v27
        _ = v29
    case "a":
        switch v1 {
        case "\n\t":
            v4[v1] = (889 >= 255)
        case "":
            fmt.Println(v5, v6, v5)
        }
    }
    if (2 >= 702) {
        if ((v6 / ((v6) | 1)) != 800) {
            fmt.Println(v6, v2)
            v1 = v1
        }
        switch v1 {
        case "\n\t":
            fmt.Println(v6)
        default:
            switch -142 {
            case 4294967295:
                v30 := GenS4{F0: 1, F1: 42}
                _ = v30
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
