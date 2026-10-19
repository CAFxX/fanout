package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint; F1 int32; F2 int32 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[int]string)
    v3 := []string{((("\n\t" + "\n\t") + "a") + "foo bar")}
    var v4 int32 = -866
    var v5 uint32 = 785
    var v6 int8 = 97
    var v7 string = "a"
    var v8 uint64 = 4294967295
    switch 12 {
    case 0:
        var v9 uint8 = 255
        _ = v9
    case 42:
        switch "\n\t" {
        case "hello":
            v4 -= v4
        case "foo bar":
            var v10 uint64 = (v8 * (^(v8)))
            _ = v10
        }
    case 32767:
        switch v7 {
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = (v7 + "foo bar")
            }
        default:
            fmt.Println(v4)
        }
    }
    fmt.Println(v6, v5, v8)
    v6 += v6
    delete(v2, -9223372036854775808)
    delete(v2, 32768)
    switch v7 {
    case "\n\t":
        switch "a" {
        case "":
            switch (((v7 + (("a" + v7) + v7)) + "a") + ("hello" + ((v7 + v7) + v7))) {
            case "\n\t":
                var v11 float64 = (0.1 / -2.25)
                _ = v11
            }
        default:
            if true {
                switch 2147483647 {
                case 9223372036854775807:
                    for fi12 := 0; fi12 < 5; fi12++ {
                        if (!((-2.25 * (3.14159 / 3.14159)) == (1e10 / (0.0 + float64((0.1 - (1.5 * 1.5))))))) {
                            _ = len(v2)
                        }
                        _ = fi12
                    }
                default:
                    v5 += (v5 - v5)
                }
                _ = len(v2)
            }
        }
    case "hello":
        v13 := GenS1{F0: 4294967295, F1: v8, F2: 2}
        _ = v13
    case "foo bar":
        v8 = 968
    default:
        v5 = v5
    }
    var v14 float32 = 1e10
    v15 := GenS2{F0: 0, F1: (v4 % ((v4) | 1)), F2: (v4 - -2147483648)}
    if false {
        v14 = v14
        for fi16 := 0; fi16 < 18; fi16++ {
            v1 = v1
            _ = fi16
        }
        var v17 uint = 42
        _ = v17
    } else {
        var v18 bool = false
        _ = v18
    }
    v19 := GenS2{F0: 4294967295, F1: -1, F2: v4}
    switch (v7 + v7) {
    case "a":
        var v20 int64 = 1
        _ = v20
    case "":
        var v21 string = v7
        _ = v21
    case "\n\t":
        v1 += v1
    }
    var v22 int = 2147483647
    if (0 != v4) {
        for fi23 := 0; fi23 < 6; fi23++ {
            switch fi23 {
            case 266:
                fmt.Println(v8, v22)
            }
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 string = (v7 + v7)
                v26 := v2[fi23]
                switch (v7 + ("foo bar" + v7)) {
                case "hello":
                    var v27 int16 = (v1 | v1)
                    _ = v27
                case "a":
                    var v28 uint8 = 1
                    _ = v28
                case "\n\t":
                    v5 = (v5 % ((1) | 1))
                default:
                    var v29 float32 = 0.0
                    _ = v29
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = fi23
        }
        v2[(^(v22))] = v7
    } else {
        v2[v22] = v7
        fmt.Println(v6, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
