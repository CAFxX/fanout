package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }

func main() {
    var v1 int16 = 914
    var v2 int = -118
    var v3 uint32 = 2147483647
    var v4 uint32 = v3
    v5 := make(map[int]int)
    v6 := GenS1{F0: 255, F1: 622}
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 float64 = 0.0
        _ = fi7
        _ = v8
    }
    var v9 uint = 42
    var v10 int16 = (v1 | -32768)
    var v11 string = ("foo bar" + ("hello" + ("a" + "foo bar")))
    for fi12 := 0; fi12 < 7; fi12++ {
        if (100 >= -1) {
            v13 := []int{v2, (v2 << ((v9) & 63))}
            delete(v5, (^(fi12)))
            _ = v13
        }
        v14 := GenS1{F0: 0, F1: -9223372036854775808}
        v9 -= (^(v9))
        _ = fi12
        _ = v14
    }
    fmt.Println(v2, v9, v1)
    switch v2 {
    case 401:
        var v15 float32 = ((0.0 * ((3.14159 + ((0.1 * (0.0 / ((1e10 * 1.5) - 0.0))) * (0.0 - 3.14159))) - -2.25)) + float32((0.0 - 0.1)))
        _ = v15
    case 42:
        v16 := GenS1{F0: 32767, F1: 1}
        _ = v16
    default:
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 float32 = 1e10
            v18 += v18
            _ = fi17
            _ = v18
        }
    }
    var v19 string = (v11 + (v11 + v11))
    switch -2147483648 {
    case 0:
        v20 := GenS1{F0: 42, F1: 9223372036854775807}
        _ = v20
    case 1:
        var v21 uint32 = 65535
        _ = v21
    case 42:
        fmt.Println(v11, v3)
    }
    fmt.Println(v3, v11, v9)
    v4 -= v4
    switch v11 {
    case "a":
        fmt.Println(v2)
    }
    v2 = (v2 | v2)
    fmt.Println(v1, v4)
    if false {
        if true {
            v19 = (v11 + (v19 + (v19 + v19)))
            for fi22 := 0; fi22 < 17; fi22++ {
                for fi23 := 0; fi23 < 1; fi23++ {
                    _ = len(v5)
                    switch (v19 + "foo bar") {
                    case "hello":
                        var v24 uint32 = v3
                        _ = v24
                    case "a":
                        switch "" {
                        case "foo bar":
                            var v25 uint = (v9 + v9)
                            _ = v25
                        default:
                            v11 = (v19 + "foo bar")
                        }
                    case "\n\t":
                        var v26 int64 = -746
                        _ = v26
                    }
                    v19 = (((v11 + ("foo bar" + v19)) + v11) + ((v19 + (v11 + ((v19 + "hello") + v11))) + ("\n\t" + v19)))
                    _ = fi23
                }
                if true {
                    var v27 uint64 = 42
                    v19 = v19
                    v4 = v3
                    _ = v27
                }
                var v28 bool = true
                _ = fi22
                _ = v28
            }
        } else {
            for fi29 := 0; fi29 < 7; fi29++ {
                var v30 uint = (v9 + (^(v9)))
                _ = fi29
                _ = v30
            }
            fmt.Println(v1)
            for fi31 := 0; fi31 < 5; fi31++ {
                switch v19 {
                case "foo bar":
                    if (false || false) {
                        for fi32 := 0; fi32 < 6; fi32++ {
                            fmt.Println(fi32, v9, fi32)
                            v9 -= (v9 - uint(v4))
                            _ = fi32
                        }
                    } else {
                        switch v11 {
                        case "\n\t":
                            var v33 int16 = int16(v3)
                            _ = v33
                        default:
                            switch (-(v2)) {
                            case -9223372036854775808:
                                var v34 uint = 4294967295
                                _ = v34
                            case 289:
                                switch v11 {
                                case "":
                                    var v35 int64 = -51
                                    _ = v35
                                default:
                                    if true {
                                        v36 := GenS1{F0: 879, F1: 9223372036854775807}
                                        _ = v36
                                    }
                                }
                            default:
                                v9 += (v9 % (((v9 | 0)) | 1))
                            }
                        }
                    }
                case "\n\t":
                    var v37 float32 = (3.14159 + 0.1)
                    _ = v37
                case "":
                    v11 = v11
                }
                v38 := []int{}
                _ = fi31
                _ = v38
            }
        }
        switch v2 {
        case 515:
            var v39 uint16 = 65535
            _ = v39
        case 214:
            v4 += v3
        }
    }
    v19 = "hello"
    v11 = ((((v19 + ("\n\t" + v19)) + v11) + "") + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
