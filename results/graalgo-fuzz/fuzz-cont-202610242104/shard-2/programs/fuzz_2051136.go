package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 bool; F1 bool; F2 int }

func main() {
    var v1 uint32 = 0
    var v2 uint32 = (^(v1))
    v3 := []int{-489}
    var v4 uint16 = 1
    var v5 uint16 = 2
    var v6 int = -580
    var v7 float64 = 3.14159
    switch (v6 >> ((901) & 63)) {
    case -970:
        fmt.Println(v4, v7)
    default:
        var v8 uint8 = 255
        _ = v8
    }
    var v9 float32 = (1e10 * 0.1)
    fmt.Println(v5, v1)
    v1 -= (v2 << ((614) & 63))
    var v10 int64 = 42
    if false {
        switch v6 {
        case -532:
            v11 := GenS1{F0: (^(v4)), F1: 128, F2: v2}
            _ = v11
        case 42:
            switch -624 {
            case 32767:
                for fi12 := 0; fi12 < 6; fi12++ {
                    v3 = append(v3, fi12)
                    var v13 float64 = 1.5
                    fmt.Println(v1, v7, v5)
                    _ = fi12
                    _ = v13
                }
            case -1:
                switch "\n\t" {
                case "\n\t":
                    switch "foo bar" {
                    case "\n\t":
                        var v14 string = "\n\t"
                        _ = v14
                    case "hello":
                        fmt.Println(v6, v7, v9)
                    default:
                        var v15 string = "hello"
                        _ = v15
                    }
                case "":
                    var v16 int32 = -1
                    _ = v16
                }
            case 0:
                switch "" {
                case "hello":
                    var v17 int64 = v10
                    _ = v17
                default:
                    v1 += v2
                }
            }
        }
        if true {
            switch (v6 | (v6 >> ((2) & 63))) {
            case -9223372036854775808:
                v2 = (^(v1))
            default:
                for fi18 := 0; fi18 < 4; fi18++ {
                    if ((!((v7 + v7) == float64(0.0))) && (4294967295 <= 0)) {
                        for fi19 := 0; fi19 < 5; fi19++ {
                            var v20 float32 = v9
                            var v21 float32 = v20
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                    for fi22 := 0; fi22 < 2; fi22++ {
                        var v23 int = fi18
                        _ = fi22
                        _ = v23
                    }
                    _ = fi18
                }
            }
            v24 := GenS2{F0: (((true && ("hello" >= ("hello" + "hello"))) && (!false)) || (17 < 90)), F1: (3.14159 != ((v7 + v7) * (0.1 / ((v7 / v7) + float64(v9))))), F2: v6}
            _ = v24
        } else {
            var v25 int32 = -1
            var v26 int8 = int8(v10)
            v6 -= v6
            _ = v25
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 int32 = -2147483648
        var v29 bool = (-2147483648 <= 96)
        for fi30 := 0; fi30 < 20; fi30++ {
            v2 -= 447
            switch "a" {
            case "":
                fmt.Println(v9)
            case "hello":
                v28 += (-(v28))
            case "a":
                if (false || ((!v29) || (((v28 >> ((693) & 63)) < v28) || (0.0 > v7)))) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    v7 = float64((0.1 * (1e10 - v9)))
                }
            default:
                v29 = (false || (!true))
            }
            fmt.Println(v2, v2, v6)
            _ = fi30
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
