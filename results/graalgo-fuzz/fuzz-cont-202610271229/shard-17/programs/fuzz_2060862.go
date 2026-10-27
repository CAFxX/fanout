package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int32; F1 uint8; F2 uint32 }

func main() {
    var v1 int16 = 0
    var v2 float32 = 1.5
    var v3 int8 = 100
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(v3, fi5, v1)
        var v6 uint64 = 1
        _ = fi5
        _ = v6
    }
    v2 = v2
    v2 = v2
    fmt.Println(v1)
    fmt.Println(v3)
    fmt.Println(v1)
    v3 -= (-(v3))
    _ = len(v4)
    v7 := GenS1{F0: (2147483647 >= 38), F1: 119}
    switch (("\n\t" + "a") + "") {
    case "a":
        var v8 int32 = -2147483648
        _ = v8
    default:
        var v9 float32 = v2
        _ = v9
    }
    if (530 <= -9223372036854775808) {
        fmt.Println(v2, v2)
    } else {
        v10 := GenS2{F0: 18446744073709551615, F1: 36}
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                _ = len(v4)
                _ = fi12
            }
            switch (fi11 / ((fi11) | 1)) {
            case 128:
                var v13 bool = false
                _ = v13
            case 2147483647:
                if true {
                    v14 := []int{(-(fi11)), -573}
                    _ = v14
                } else {
                    var v15 string = "foo bar"
                    _ = v15
                }
            default:
                var v16 float64 = 1.5
                _ = v16
            }
            if (false && ((!false) && true)) {
                fmt.Println(fi11, v1, v2)
                v1 = 1
            }
            _ = fi11
        }
        _ = v10
    }
    var v17 uint = 686
    v18 := GenS3{F0: (v17 >> (((v17 * v17)) & 63))}
    _ = len(v4)
    switch (((("a" + "a") + "foo bar") + "\n\t") + (("a" + "hello") + "hello")) {
    case "\n\t":
        switch 2147483647 {
        case 170:
            for fi19 := 0; fi19 < 5; fi19++ {
                v3 = (^(v3))
                v1 -= (-(v1))
                _ = fi19
            }
        case 4294967295:
            var v20 int32 = 157
            _ = v20
        case 0:
            switch "" {
            case "a":
                var v21 string = ""
                _ = v21
            case "":
                var v22 uint16 = 0
                _ = v22
            case "hello":
                v3 -= (v3 << (((v17 >> ((v17) & 63))) & 63))
            default:
                var v23 uint8 = 163
                _ = v23
            }
        default:
            v2 = (v2 - (v2 * v2))
        }
    case "a":
        for fi24 := 0; fi24 < 5; fi24++ {
            if false {
                if (0.1 == float64((v2 - v2))) {
                    v4[(fi24 * fi24)] = ""
                    switch "a" {
                    case "\n\t":
                        if true {
                            var v25 float32 = (v2 - float32(3.14159))
                            _ = v25
                        }
                    case "a":
                        var v26 uint8 = 42
                        _ = v26
                    case "":
                        var v27 int = fi24
                        _ = v27
                    }
                    var v28 uint8 = 1
                    _ = v28
                }
                for fi29 := 0; fi29 < 3; fi29++ {
                    v30 := GenS4{F0: 2147483648, F1: (fi29 | 9223372036854775807)}
                    v31 := GenS4{F0: 2147483648, F1: (fi29 * (fi24 + fi24))}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            } else {
                v32 := []int{fi24, (fi24 >> ((v17) & 63))}
                _ = v32
            }
            if false {
                switch (("foo bar" + "") + ("foo bar" + "a")) {
                case "":
                    var v33 int32 = 2147483647
                    _ = v33
                }
                for fi34 := 0; fi34 < 15; fi34++ {
                    fmt.Println(v17, v3, fi34)
                    v35 := GenS2{F0: 328, F1: 1}
                    _ = fi34
                    _ = v35
                }
                var v36 uint16 = 32767
                _ = v36
            }
            _ = fi24
        }
    case "foo bar":
        var v37 uint = v17
        _ = v37
    }
    v38 := GenS5{F0: 1, F1: uint8(v17), F2: 2}
    switch 0 {
    case 0:
        var v39 uint = 42
        _ = v39
    case -333:
        var v40 int = 1
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v7
}
