package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint8; F1 int; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 uint32; F2 int8 }
type GenS6 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS7 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS8 struct { F0 int16; F1 int16; F2 int8 }
func genFn0() int8 {
    v5 := []int{}
    if len(v5) > 0 {
        v5[0] = -340
    }
    var v6 float64 = 0.1
    if len(v5) > 0 {
        v5[0] = -739
    }
    v6 -= v6
    _ = v5
    _ = v6
    return 127
}

func main() {
    var v1 uint = 65535
    var v2 int16 = 32767
    var v3 int16 = v2
    var v4 uint64 = 1
    var v7 bool = false
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint16 = 383
        _ = fi8
        _ = v9
    }
    switch -2147483648 {
    case 2147483647:
        var v10 uint = (v1 | v1)
        _ = v10
    case 489:
        fmt.Println(v1, v3, v3)
    default:
        var v11 float64 = float64(float32(3.14159))
        _ = v11
    }
    switch ("\n\t" + (("hello" + "hello") + "")) {
    case "a":
        v1 -= 42
    case "hello":
        var v12 int = 32768
        _ = v12
    case "":
        v13 := []int{907}
        _ = v13
    default:
        v7 = ((v1 <= uint(v4)) || false)
    }
    fmt.Println(v4)
    v14 := GenS1{F0: 491, F1: true}
    if (42 == 2) {
        v15 := []int{128, 32767}
        if (v7 && v7) {
            v7 = (3.14159 >= 1e10)
            for fi16 := 0; fi16 < 16; fi16++ {
                fmt.Println(v1, v7)
                v7 = v7
                _ = fi16
            }
            fmt.Println(v2, v3)
        }
        _ = v15
    } else {
        fmt.Println(v1)
        v17 := GenS2{F0: 229, F1: 127}
        _ = v17
    }
    if v7 {
        for fi18 := 0; fi18 < 20; fi18++ {
            for fi19 := 0; fi19 < 12; fi19++ {
                v20 := GenS2{F0: 127, F1: -1}
                _ = fi19
                _ = v20
            }
            v4 += 1
            _ = fi18
        }
        if v7 {
            var v21 bool = (32768 != -228)
            fmt.Println(v3, v2)
            v22 := []int{0}
            _ = v21
            _ = v22
        } else {
            v23 := GenS2{F0: uint8(v4), F1: -1}
            _ = v23
        }
        var v24 uint16 = 829
        _ = v24
    } else {
        switch 127 {
        case -1:
            v7 = ((false || true) && true)
        default:
            switch ("\n\t" + "a") {
            case "\n\t":
                fmt.Println(v4)
            case "hello":
                for fi25 := 0; fi25 < 18; fi25++ {
                    v26 := GenS3{F0: 2, F1: fi25, F2: 0}
                    var v27 int8 = genFn0()
                    v28 := GenS4{F0: v4}
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
            }
        }
    }
    v29 := GenS5{F0: "", F1: 2147483648, F2: -25}
    var v30 bool = (1.5 >= 0.0)
    for fi31 := 0; fi31 < 9; fi31++ {
        v30 = v30
        v4 -= (v4 >> (((v1 + (v1 - (^(v1))))) & 63))
        _ = fi31
    }
    if (128 <= 1) {
        fmt.Println(v3)
        switch "" {
        case "":
            var v32 float64 = 1.5
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        switch "a" {
        case "a":
            v2 = (-(v2))
        case "\n\t":
            for fi34 := 0; fi34 < 6; fi34++ {
                var v35 uint = (v1 / ((v1) | 1))
                v36 := GenS6{F0: 255, F1: v7, F2: 477}
                if (!(1e10 != (0.0 * (3.14159 * float64(-2.25))))) {
                    v37 := GenS7{F0: 780, F1: v3, F2: 255}
                    fmt.Println(v35, fi33, v30)
                    v38 := GenS5{F0: "a", F1: 950, F2: genFn0()}
                    _ = v37
                    _ = v38
                }
                _ = fi34
                _ = v35
                _ = v36
            }
        case "foo bar":
            var v39 int32 = 1
            _ = v39
        default:
            v3 -= (v2 / ((v3) | 1))
        }
        switch (fi33 ^ (fi33 + (fi33 % ((-9223372036854775808) | 1)))) {
        case 697:
            v40 := []int{fi33}
            _ = v40
        case -9223372036854775808:
            fmt.Println(v3)
        case 42:
            fmt.Println(v3, v3)
        }
        _ = fi33
    }
    v41 := GenS8{F0: -719, F1: v3, F2: 0}
    v30 = (true && (v7 || (9223372036854775807 > 9223372036854775807)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v7
}
