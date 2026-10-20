package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS2 struct { F0 string; F1 int64 }

func main() {
    var v1 float64 = (3.14159 + ((1e10 + 0.0) * 0.0))
    v2 := make(map[int]uint32)
    var v3 int8 = 58
    var v4 uint8 = 42
    v5 := make(map[uint32]int)
    switch 127 {
    case 128:
        for fi6 := 0; fi6 < 8; fi6++ {
            switch (("\n\t" + "") + "") {
            case "":
                if (!false) {
                    v3 = -1
                } else {
                    v1 += v1
                }
            case "foo bar":
                var v7 float64 = v1
                _ = v7
            case "\n\t":
                fmt.Println(v4, fi6, v1)
            }
            var v8 int = (fi6 | -9223372036854775808)
            v9 := []int{(v8 + fi6), v8, fi6}
            _ = fi6
            _ = v8
            _ = v9
        }
    case -381:
        if (42 >= 565) {
            var v10 int8 = (v3 - v3)
            _ = v10
        }
    }
    switch 0 {
    case -2147483648:
        v11 := GenS1{F0: (v4 ^ (v4 / (((v4 << ((42) & 63))) | 1))), F1: 65535, F2: 1}
        _ = v11
    case -9223372036854775808:
        fmt.Println(v4)
    default:
        v4 = (v4 & v4)
    }
    v12 := GenS1{F0: 128, F1: 1, F2: -1}
    var v13 uint64 = 857
    v3 = v3
    switch "" {
    case "":
        v1 -= float64(-2.25)
    case "a":
        for fi14 := 0; fi14 < 17; fi14++ {
            v5[2147483647] = (fi14 & (fi14 + fi14))
            fmt.Println(v13)
            for fi15 := 0; fi15 < 7; fi15++ {
                v16 := GenS1{F0: 128, F1: 286, F2: 403}
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    case "hello":
        v5[2] = 32767
    default:
        v17 := []int{}
        _ = v17
    }
    var v18 int = 9223372036854775807
    v13 = v13
    fmt.Println(v1, v18)
    if true {
        v18 += v18
        if (!(!((v3 + v3) == v3))) {
            var v19 uint = 4294967295
            _ = v19
        }
    } else {
        switch v18 {
        case 128:
            v20 := GenS2{F0: "foo bar", F1: -1}
            _ = v20
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            for fi22 := 0; fi22 < 13; fi22++ {
                v4 = (v4 << ((42) & 63))
                delete(v2, v18)
                switch (fi22 - (v18 * (v18 | fi22))) {
                case 127:
                    if false {
                        if ((!(true && ((!(0 < 65535)) && (!true)))) || true) {
                            delete(v5, 2147483647)
                            switch (fi22 | fi21) {
                            case -930:
                                v23 := v5[749]
                                _ = v23
                            }
                            fmt.Println(fi21, fi21, fi21)
                        }
                        var v24 int32 = -1
                        switch (fi21 / (((^(fi21))) | 1)) {
                        case 42:
                            if false {
                                switch "" {
                                case "a":
                                    fmt.Println(v18, v24, v24)
                                case "":
                                    v1 -= v1
                                default:
                                    var v25 string = "a"
                                    _ = v25
                                }
                                fmt.Println(v4, v18)
                                var v26 uint16 = 1
                                _ = v26
                            } else {
                                switch ("hello" + "a") {
                                case "foo bar":
                                    v4 += (v4 * (^(v4)))
                                default:
                                    v27 := v2[fi22]
                                    _ = v27
                                }
                                fmt.Println(fi22, v3)
                                var v28 float64 = v1
                                _ = v28
                            }
                        }
                        _ = v24
                    } else {
                        if true {
                            if false {
                                if (1 != 47) {
                                    v29 := GenS2{F0: "a", F1: -9223372036854775808}
                                    v3 = (v3 + (v3 % ((v3) | 1)))
                                    var v30 int16 = -280
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v31 uint64 = v13
                                    if true {
                                        delete(v5, 1)
                                        fmt.Println(v4, v13, fi21)
                                        v31 += (v13 | v13)
                                    }
                                    _ = v31
                                }
                                switch "" {
                                case "hello":
                                    v32 := []int{fi21, 1, (-(fi22))}
                                    _ = v32
                                }
                                v33 := []int{127}
                                _ = v33
                            }
                        }
                        v3 = (v3 * v3)
                    }
                case -1:
                    v34 := []int{1}
                    _ = v34
                }
                _ = fi22
            }
            _ = fi21
        }
        var v35 float32 = 0.0
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
