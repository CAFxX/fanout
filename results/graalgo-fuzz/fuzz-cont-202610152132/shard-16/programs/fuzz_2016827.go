package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint; F2 int }

func main() {
    v1 := make(map[int]int)
    var v2 uint = 178
    v3 := []string{"foo bar", "foo bar"}
    var v4 uint32 = 0
    v5 := []int{-9223372036854775808, 32767, 748}
    v6 := GenS1{F0: 681, F1: 4294967295}
    v7 := GenS1{F0: 1, F1: 1}
    switch 128 {
    case -452:
        if false {
            switch "a" {
            case "":
                v2 -= (^(v2))
            default:
                fmt.Println(v2, v4)
            }
            switch 920 {
            case 42:
                fmt.Println(v4, v2)
            default:
                _ = len(v1)
            }
            switch 1 {
            case -9223372036854775808:
                v1[273] = 32767
            default:
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 int8 = 0
                    var v10 uint64 = 1
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
        } else {
            switch ("a" + ("" + ("foo bar" + "foo bar"))) {
            case "a":
                fmt.Println(v4, v4)
            case "":
                v4 += (v4 - (v4 * v4))
            case "\n\t":
                var v11 uint16 = 65535
                _ = v11
            }
            if false {
                v4 += 65535
            } else {
                v2 = uint(v4)
                for fi12 := 0; fi12 < 3; fi12++ {
                    v4 += v4
                    _ = fi12
                }
            }
        }
    case 42:
        if false {
            v13 := GenS1{F0: 551, F1: v4}
            _ = v13
        }
    default:
        var v14 uint64 = 4294967295
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        switch ("" + "hello") {
        case "hello":
            if true {
                for fi16 := 0; fi16 < 11; fi16++ {
                    v17 := GenS1{F0: 279, F1: v4}
                    _ = fi16
                    _ = v17
                }
                v18 := GenS1{F0: 690, F1: (v4 + v4)}
                _ = v18
            }
        case "a":
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 float64 = 0.1
                _ = fi19
                _ = v20
            }
        case "":
            _ = len(v1)
        }
        v2 = 4294967295
        _ = fi15
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        switch "foo bar" {
        case "\n\t":
            v4 = 2
        case "hello":
            for fi22 := 0; fi22 < 7; fi22++ {
                switch 32767 {
                case 4294967295:
                    var v23 int32 = 1
                    _ = v23
                default:
                    _ = len(v1)
                }
                v5 = append(v5, 127)
                _ = fi22
            }
        case "":
            v24 := GenS2{F0: -40, F1: v2, F2: fi21}
            _ = v24
        }
        _ = len(v5)
        _ = fi21
    }
    var v25 uint8 = 43
    if (v4 > (v4 ^ v4)) {
        var v26 uint16 = 255
        v2 -= uint(v25)
        v26 = 93
        _ = v26
    }
    fmt.Println(v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
