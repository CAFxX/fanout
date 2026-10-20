package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS4 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint32 = 65535
    var v2 uint32 = (^(v1))
    var v3 string = "hello"
    var v4 int16 = 1000
    var v5 uint = 2
    v6 := GenS1{F0: -826}
    v7 := GenS1{F0: 1}
    v8 := []int{32767, -1, 32767}
    v1 += (v1 | (v2 / ((v2) | 1)))
    if len(v8) > 0 {
        v8[0] = 0
    }
    var v9 string = v3
    var v10 uint32 = v1
    switch "" {
    case "foo bar":
        for fi11 := 0; fi11 < 9; fi11++ {
            _ = len(v8)
            switch (^(fi11)) {
            case 32767:
                if false {
                    var v12 float64 = float64((0.1 - 3.14159))
                    _ = v12
                }
            case -305:
                for fi13 := 0; fi13 < 17; fi13++ {
                    _ = len(v8)
                    _ = fi13
                }
            }
            _ = fi11
        }
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int32 = 42
        var v16 uint = (v5 & 920)
        _ = fi14
        _ = v15
        _ = v16
    }
    switch 2147483647 {
    case 42:
        var v17 int64 = 9223372036854775807
        _ = v17
    default:
        if false {
            switch (v3 + "") {
            case "foo bar":
                fmt.Println(v2, v1, v1)
            case "":
                v8 = append(v8, 4294967295)
            }
            v3 = v9
        } else {
            if ((float64(3.14159) * (((((1.5 / (1e10 * (((-2.25 * 1e10) * ((((3.14159 + 1e10) / 1.0) + 0.1) + ((-2.25 * 1e10) * -2.25))) * float64(3.14159)))) - 0.0) - 1e10) * -2.25) + (((0.1 - -2.25) + 0.1) + (0.0 - -2.25)))) <= 0.0) {
                var v18 float64 = float64(-2.25)
                _ = v18
            }
        }
    }
    switch v9 {
    case "foo bar":
        v1 = 65535
    default:
        switch 127 {
        case 4294967295:
            fmt.Println(v10, v3)
        case -1:
            fmt.Println(v4, v2, v10)
        }
    }
    v19 := GenS2{F0: (((v9 + ("\n\t" + v3)) + "hello") + "hello"), F1: 42, F2: 42}
    v20 := GenS3{F0: 0, F1: 255, F2: 379}
    switch 4294967295 {
    case -2147483648:
        for fi21 := 0; fi21 < 8; fi21++ {
            switch "foo bar" {
            case "":
                var v22 float64 = 1e10
                _ = v22
            case "\n\t":
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            case "foo bar":
                if ((65535 <= 1) || (100 == 0)) {
                    v2 = v2
                    v23 := GenS4{F0: true, F1: v1}
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        var v25 int = fi21
                        _ = fi24
                        _ = v25
                    }
                }
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi21
        }
    case -736:
        fmt.Println(v2, v2)
    case 42:
        switch 127 {
        case 9223372036854775807:
            var v26 float64 = (0.0 / (((((1.5 + -2.25) * 0.0) * 1.5) - -2.25) - ((0.0 - 0.0) / 1.0)))
            _ = v26
        case -2147483648:
            v4 = (v4 >> (((v5 >> (((v5 % ((v5) | 1))) & 63))) & 63))
        }
    default:
        for fi27 := 0; fi27 < 2; fi27++ {
            v10 = 4294967295
            if true {
                var v28 string = v3
                for fi29 := 0; fi29 < 3; fi29++ {
                    var v30 uint = v5
                    if true {
                        fmt.Println(fi27, v2)
                    } else {
                        v31 := GenS3{F0: 0, F1: 128, F2: 1}
                        var v32 uint8 = 128
                        var v33 int = fi27
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    v9 = ("foo bar" + ("hello" + v3))
                    var v34 uint = v5
                    v34 = (^(v30))
                    _ = fi29
                    _ = v30
                    _ = v34
                }
                var v35 string = v28
                _ = v28
                _ = v35
            }
            var v36 int64 = -1
            _ = fi27
            _ = v36
        }
    }
    var v37 uint16 = 0
    for fi38 := 0; fi38 < 5; fi38++ {
        v3 = "hello"
        _ = len(v8)
        v39 := GenS4{F0: true, F1: (v2 ^ v1)}
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
