package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int64 }
type GenS2 struct { F0 int; F1 string }

func main() {
    var v1 uint = 367
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    v4 := []int{}
    var v5 uint8 = 252
    var v6 uint8 = 0
    var v7 float64 = 3.14159
    v8 := GenS1{F0: 87, F1: 1, F2: -643}
    switch "foo bar" {
    case "":
        var v9 int32 = 2147483647
        _ = v9
    case "\n\t":
        var v10 uint8 = (v6 + (v5 % (((v5 << (((^(v1))) & 63))) | 1)))
        _ = v10
    }
    delete(v2, 0)
    v5 = 38
    _ = len(v4)
    v1 -= v1
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            if false {
                switch fi12 {
                case 1:
                    if false {
                        var v13 uint64 = 4294967295
                        var v14 int64 = 9223372036854775807
                        _ = v13
                        _ = v14
                    }
                default:
                    v3[("a" + ("foo bar" + "a"))] = ""
                }
                v7 = ((v7 / (float64(3.14159) - v7)) - float64(0.0))
                var v15 uint32 = 2147483648
                var v16 float64 = v7
                _ = v15
                _ = v16
            } else {
                _ = len(v3)
                fmt.Println(v5, v5)
                var v17 float64 = ((0.1 - -2.25) - float64(0.0))
                _ = v17
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v7 += 1e10
            _ = fi12
        }
        if true {
            v5 += v6
            switch 32768 {
            case -9223372036854775808:
                v18 := GenS1{F0: 127, F1: 1, F2: 1}
                _ = v18
            case 127:
                for fi19 := 0; fi19 < 12; fi19++ {
                    v7 = (v7 * float64(float32(1e10)))
                    v7 += v7
                    var v20 float32 = -2.25
                    _ = fi19
                    _ = v20
                }
            }
        } else {
            var v21 uint8 = (v6 * (v5 / ((v5) | 1)))
            var v22 int16 = 1000
            for fi23 := 0; fi23 < 8; fi23++ {
                var v24 int32 = -2147483648
                var v25 int8 = 126
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v21 = v5
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = v21
            _ = v22
        }
        _ = fi11
    }
    if (v7 != v7) {
        fmt.Println(v5, v7, v7)
        v7 = v7
        v4 = append(v4, 0)
    } else {
        if (!((527 >= 65535) && (!(2147483647 < 907)))) {
            fmt.Println(v5, v5, v5)
        }
        switch 482 {
        case 9223372036854775807:
            v1 = uint(v5)
        case -1:
            v1 += 567
        }
    }
    var v26 uint8 = (v5 | (v6 / (((v6 << (((^(v1))) & 63))) | 1)))
    for fi27 := 0; fi27 < 16; fi27++ {
        if (3.14159 != 0.0) {
            v28 := v2[927]
            if v28 {
                v29 := GenS1{F0: 100, F1: 9223372036854775807, F2: 42}
                _ = v29
            } else {
                v30 := GenS1{F0: 100, F1: 552, F2: 461}
                for fi31 := 0; fi31 < 10; fi31++ {
                    switch (("a" + "hello") + "") {
                    case "\n\t":
                        var v32 string = "hello"
                        _ = v32
                    case "":
                        v33 := v2[290]
                        _ = v33
                    }
                    _ = fi31
                }
                v6 = (^(v6))
                _ = v30
            }
            for fi34 := 0; fi34 < 10; fi34++ {
                v26 = (v26 << ((42) & 63))
                v4 = append(v4, (^(fi34)))
                _ = fi34
            }
            _ = v28
        }
        var v35 float64 = v7
        if len(v4) > 0 {
            v4[0] = (fi27 | (fi27 << (((v1 >> (((v1 >> (((v1 - (v1 * 1))) & 63))) & 63))) & 63)))
        }
        _ = fi27
        _ = v35
    }
    v7 = v7
    v36 := GenS1{F0: -39, F1: 1, F2: 459}
    v1 += (v1 | v1)
    switch "foo bar" {
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            switch 127 {
            case -9223372036854775808:
                v37 := GenS1{F0: 1, F1: 988, F2: -125}
                _ = v37
            case 32767:
                var v38 int32 = 2147483647
                _ = v38
            }
        case "a":
            if true {
                fmt.Println(v6, v5, v5)
                v1 += v1
                v26 = (v5 << (((v1 % ((v1) | 1))) & 63))
            }
        case "":
            v4 = append(v4, 4294967295)
        }
    case "a":
        _ = len(v4)
    case "hello":
        var v39 float64 = 0.1
        _ = v39
    }
    v40 := GenS2{F0: 0, F1: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
