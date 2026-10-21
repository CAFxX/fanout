package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int8; F1 uint8; F2 uint }

func main() {
    v1 := []uint32{227}
    var v2 int8 = 89
    var v3 uint = 201
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    switch ("hello" + ("" + "\n\t")) {
    case "":
        v3 -= 18446744073709551615
    case "\n\t":
        switch -9223372036854775808 {
        case -9223372036854775808:
            if (0 < 4294967295) {
                fmt.Println(v3)
                fmt.Println(v3, v2)
                if false {
                    v6 := GenS1{F0: 65535}
                    var v7 uint64 = 18446744073709551615
                    var v8 uint16 = 1
                    _ = v6
                    _ = v7
                    _ = v8
                } else {
                    v4[9223372036854775807] = (("hello" + "foo bar") + "")
                }
            } else {
                if (!(int16(v3) > 1000)) {
                    v9 := GenS2{F0: 823, F1: 65535}
                    v3 += 255
                    _ = v9
                }
                switch "" {
                case "foo bar":
                    if true {
                        v2 = 20
                    }
                case "":
                    for fi10 := 0; fi10 < 17; fi10++ {
                        v3 -= 226
                        _ = fi10
                    }
                }
                fmt.Println(v2)
            }
        case 1:
            if (!true) {
                v11 := GenS2{F0: 255, F1: 65535}
                _ = v11
            }
        case 4294967295:
            for fi12 := 0; fi12 < 18; fi12++ {
                if (("\n\t" > "foo bar") && (true && (4294967295 <= 4294967295))) {
                    v3 = v3
                } else {
                    var v13 int8 = -128
                    if (1.5 >= float64(1.5)) {
                        v14 := GenS3{F0: -9223372036854775808, F1: 42}
                        var v15 int8 = 42
                        _ = v14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 18; fi16++ {
                        v2 -= (v13 ^ (v2 & (v13 + v2)))
                        var v17 uint16 = 65535
                        var v18 float64 = 1e10
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    _ = v13
                }
                _ = fi12
            }
        default:
            fmt.Println(v3, v2)
        }
    }
    v2 += v2
    fmt.Println(v2, v3)
    var v19 bool = (32767 > 314)
    v20 := GenS3{F0: 42, F1: (v3 - v3)}
    v2 += (^(v2))
    switch 9223372036854775807 {
    case -349:
        var v21 int = 0
        _ = v21
    case 128:
        for fi22 := 0; fi22 < 15; fi22++ {
            fmt.Println(v3, v19, v3)
            for fi23 := 0; fi23 < 8; fi23++ {
                v24 := GenS4{F0: 0}
                var v25 uint32 = 65535
                fmt.Println(fi23)
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi22
        }
    case -800:
        v26 := GenS5{F0: 2, F1: v3}
        _ = v26
    default:
        if v19 {
            var v27 string = ("" + (("" + "a") + "hello"))
            _ = v27
        } else {
            var v28 uint32 = 2147483647
            v2 += 56
            for fi29 := 0; fi29 < 11; fi29++ {
                switch fi29 {
                case 42:
                    var v30 float32 = 1.5
                    _ = v30
                default:
                    v3 -= v3
                }
                v31 := GenS6{F0: (v2 * v2), F1: 1, F2: v3}
                for fi32 := 0; fi32 < 11; fi32++ {
                    switch (("a" + ("foo bar" + ("foo bar" + ("" + ("" + "a"))))) + "a") {
                    case "hello":
                        var v33 int = (fi29 / ((fi32) | 1))
                        _ = v33
                    case "foo bar":
                        fmt.Println(v2, v2)
                    }
                    v3 += 1
                    var v34 float32 = 1e10
                    var v35 int64 = -9223372036854775808
                    _ = fi32
                    _ = v34
                    _ = v35
                }
                _ = fi29
                _ = v31
            }
            _ = v28
        }
    }
    if false {
        v36 := GenS6{F0: 127, F1: 42, F2: 260}
        _ = v36
    }
    v37 := GenS2{F0: 65535, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
