package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint32 }
type GenS2 struct { F0 uint; F1 int64; F2 bool }

func main() {
    var v1 int = 9223372036854775807
    var v2 int8 = -1
    var v3 float32 = (3.14159 - 1.5)
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 uint32 = 4294967295
        var v6 int16 = 725
        _ = fi4
        _ = v5
        _ = v6
    }
    switch "" {
    case "a":
        v7 := GenS1{F0: -1, F1: (("a" + "\n\t") + "a"), F2: 2147483647}
        _ = v7
    case "foo bar":
        var v8 float64 = 3.14159
        _ = v8
    default:
        var v9 int = (v1 ^ (-(v1)))
        _ = v9
    }
    var v10 float32 = v3
    switch v1 {
    case 32768:
        var v11 uint8 = 2
        _ = v11
    case -1:
        if (446 != 2) {
            v12 := GenS1{F0: 9223372036854775807, F1: "foo bar", F2: 347}
            var v13 int16 = 42
            v2 = (v2 - (v2 % (((v2 >> ((65535) & 63))) | 1)))
            _ = v12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := GenS2{F0: 1, F1: 42, F2: (430 < 1)}
            _ = fi15
            _ = v16
        }
        var v17 uint = 669
        _ = fi14
        _ = v17
    }
    v2 = (v2 << ((0) & 63))
    for fi18 := 0; fi18 < 8; fi18++ {
        if false {
            if (((0 < -9223372036854775808) && true) && (2 <= 743)) {
                v19 := []int{}
                _ = v19
            }
            for fi20 := 0; fi20 < 14; fi20++ {
                switch ("" + "a") {
                case "hello":
                    var v21 int64 = 0
                    _ = v21
                default:
                    fmt.Println(fi18, v1)
                }
                var v22 uint16 = 255
                var v23 uint32 = 150
                _ = fi20
                _ = v22
                _ = v23
            }
        } else {
            fmt.Println(v2, v2, v3)
        }
        v24 := GenS1{F0: -1, F1: "\n\t", F2: 4294967295}
        for fi25 := 0; fi25 < 4; fi25++ {
            v2 += v2
            v26 := []int{(fi18 & -2147483648)}
            _ = fi25
            _ = v26
        }
        _ = fi18
        _ = v24
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        switch "a" {
        case "foo bar":
            var v28 uint32 = 962
            _ = v28
        case "a":
            var v29 int64 = 0
            _ = v29
        default:
            v30 := []int{(-(v1))}
            _ = v30
        }
        _ = fi27
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        var v32 uint8 = 128
        var v33 uint = 0
        var v34 uint = (^(v33))
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    var v35 int32 = 42
    var v36 int = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v35
    _ = v36
}
