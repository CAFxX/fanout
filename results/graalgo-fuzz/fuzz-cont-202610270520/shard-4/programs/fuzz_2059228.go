package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 uint16 = 811
    var v2 uint8 = 128
    var v3 float64 = (1.5 / -2.25)
    switch ("hello" + "a") {
    case "a":
        v4 := []int{128, -629, 32767}
        _ = v4
    case "foo bar":
        v5 := []int{-2147483648, -548, -1, -247}
        _ = v5
    case "":
        fmt.Println(v3, v1)
    default:
        v1 = 319
    }
    v1 += (v1 | v1)
    switch ("" + ((("" + "foo bar") + (("" + "a") + ("\n\t" + "hello"))) + "")) {
    case "foo bar":
        v2 = 235
    case "\n\t":
        for fi6 := 0; fi6 < 12; fi6++ {
            if false {
                v7 := GenS1{F0: 1, F1: (^(v1))}
                _ = v7
            } else {
                for fi8 := 0; fi8 < 13; fi8++ {
                    var v9 int = 0
                    var v10 int64 = 0
                    switch "\n\t" {
                    case "":
                        v11 := []int{fi8, (fi8 ^ fi8), v9, fi6}
                        _ = v11
                    case "hello":
                        v12 := []int{}
                        _ = v12
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                if (1 < -112) {
                    v13 := GenS1{F0: -1, F1: (^(v1))}
                    _ = v13
                }
                fmt.Println(v1)
            }
            _ = fi6
        }
    }
    v3 = v3
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 string = "foo bar"
        _ = fi14
        _ = v15
    }
    v1 = (v1 & v1)
    if true {
        var v16 string = "foo bar"
        _ = v16
    } else {
        v1 -= (^(v1))
        v17 := GenS1{F0: 1000, F1: 666}
        if (false || false) {
            for fi18 := 0; fi18 < 14; fi18++ {
                switch (-(fi18)) {
                case 193:
                    fmt.Println(v2)
                case 32768:
                    var v19 float64 = ((0.1 - (1.5 * 1.5)) - (v3 + 0.1))
                    _ = v19
                }
                if false {
                    v20 := []int{(fi18 + 32768)}
                    v1 = (v1 << ((4294967295) & 63))
                    v21 := GenS1{F0: 226, F1: (v1 * 1)}
                    _ = v20
                    _ = v21
                }
                if (!(43 >= 1)) {
                    var v22 int32 = 1
                    _ = v22
                } else {
                    var v23 float32 = -2.25
                    var v24 uint64 = 1
                    _ = v23
                    _ = v24
                }
                _ = fi18
            }
        }
        _ = v17
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        if true {
            var v26 int8 = 37
            var v27 uint8 = 2
            _ = v26
            _ = v27
        }
        switch ("hello" + ("foo bar" + "")) {
        case "hello":
            if true {
                v28 := []int{(fi25 * 42), (fi25 & -841), fi25}
                fmt.Println(v2, v1, v2)
                _ = v28
            }
        }
        _ = fi25
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        v2 += v2
        v30 := GenS1{F0: 32767, F1: v1}
        _ = fi29
        _ = v30
    }
    v2 = (v2 & 0)
    v31 := []int{42, 90, 128}
    if (9223372036854775807 != -9223372036854775808) {
        v3 += (0.1 + (0.0 / -2.25))
        var v32 int = 0
        _ = v32
    } else {
        if false {
            if len(v31) > 0 {
                v31[0] = int(v1)
            }
        }
    }
    v33 := GenS2{F0: -355, F1: 1}
    v34 := GenS2{F0: 32768, F1: -1}
    var v35 float32 = 0.1
    var v36 uint16 = (v1 ^ (v1 * (v1 % ((v1) | 1))))
    v37 := GenS3{F0: 18446744073709551615}
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
