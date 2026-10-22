package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = ((0.0 - 3.14159) + 0.1)
    var v3 int = 9223372036854775807
    var v4 float32 = (3.14159 - float32((3.14159 - (0.1 * v2))))
    var v5 float32 = (v4 - (0.1 - float32(v2)))
    if (-498 >= 75) {
        if false {
            if true {
                var v6 uint = 18446744073709551615
                v7 := GenS1{F0: -1, F1: 42}
                _ = v6
                _ = v7
            } else {
                v2 = 3.14159
            }
            v8 := GenS2{F0: 18446744073709551615}
            _ = v8
        }
    } else {
        fmt.Println(v2, v2, v3)
        var v9 int16 = 184
        _ = v9
    }
    var v10 string = "foo bar"
    v11 := GenS2{F0: 18446744073709551615}
    v12 := v1[1]
    var v13 int64 = 0
    var v14 float32 = v4
    v2 = 3.14159
    var v15 float64 = v2
    if (v4 != 1.5) {
        var v16 uint16 = 42
        var v17 uint16 = 622
        for fi18 := 0; fi18 < 3; fi18++ {
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 uint = 977
                v20 -= v20
                switch v10 {
                case "\n\t":
                    if false {
                        if false {
                            switch fi18 {
                            case -2147483648:
                                fmt.Println(fi19)
                            case 57:
                                var v21 bool = (v12 && v12)
                                _ = v21
                            }
                            fmt.Println(v10)
                            v22 := []int{fi19}
                            v23 := v1[42]
                            var v24 uint8 = 0
                            switch ((v10 + v10) + ("hello" + ("" + "\n\t"))) {
                            case "foo bar":
                                v17 = v17
                            case "\n\t":
                                v3 = (fi18 % ((v3) | 1))
                            case "":
                                switch "foo bar" {
                                case "\n\t":
                                    v25 := GenS3{F0: uint32(v20), F1: -5}
                                    _ = v25
                                case "":
                                    var v26 uint16 = v17
                                    _ = v26
                                }
                            }
                            v3 += fi19
                            v17 -= (v17 >> (((^(v20))) & 63))
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        if (!(!v12)) {
                            if ((4294967295 > 4294967295) && (v17 < 142)) {
                                var v27 int = v3
                                _ = v27
                            }
                        }
                        v28 := GenS1{F0: 100, F1: -63}
                        _ = v28
                    }
                case "foo bar":
                    switch (-(fi18)) {
                    case 4294967295:
                        var v29 bool = v12
                        _ = v29
                    case 553:
                        v4 -= float32(v15)
                    case -2147483648:
                        var v30 uint8 = 89
                        _ = v30
                    }
                case "hello":
                    if (!false) {
                        v20 += (v20 / (((v20 ^ (v20 & (v20 % ((v20) | 1))))) | 1))
                        switch 128 {
                        case 4294967295:
                            v15 -= (0.0 * 0.1)
                        }
                        switch (fi19 + v3) {
                        case 0:
                            if v12 {
                                switch (fi18 * 32768) {
                                case -9223372036854775808:
                                    var v31 int32 = int32(v13)
                                    _ = v31
                                case 9223372036854775807:
                                    fmt.Println(v16)
                                default:
                                    v32 := []int{-2147483648, -9223372036854775808, 673}
                                    _ = v32
                                }
                                v16 += (^(v16))
                            } else {
                                var v33 float32 = v5
                                if v12 {
                                    v5 += v14
                                } else {
                                    var v34 bool = v12
                                    v35 := []int{fi19, 127}
                                    v36 := GenS1{F0: -1, F1: 100}
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                                _ = v33
                            }
                        case 2147483647:
                            v16 = v17
                            v15 += -2.25
                        }
                    }
                }
                var v37 uint = v20
                _ = fi19
                _ = v20
                _ = v37
            }
            fmt.Println(v4, v17, fi18)
            _ = fi18
        }
        _ = v16
        _ = v17
    } else {
        fmt.Println(v3, v10)
        v13 = (v13 * v13)
    }
    if false {
        fmt.Println(v4, v3, v10)
        v14 = v4
    } else {
        for fi38 := 0; fi38 < 11; fi38++ {
            var v39 int = (^(fi38))
            fmt.Println(fi38, v5, v12)
            v40 := []int{9223372036854775807}
            _ = fi38
            _ = v39
            _ = v40
        }
        fmt.Println(v2, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
