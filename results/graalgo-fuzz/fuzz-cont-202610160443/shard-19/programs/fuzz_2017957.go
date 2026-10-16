package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int; F2 int8 }
type GenS5 struct { F0 uint8; F1 bool; F2 string }

func main() {
    var v1 int16 = 84
    var v2 uint = 4294967295
    v3 := make(map[int]bool)
    var v4 uint16 = 65535
    v5 := make(map[int]uint32)
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 52}
        _ = fi6
        _ = v7
    }
    if true {
        v8 := GenS2{F0: 0}
        _ = v8
    }
    v9 := GenS2{F0: 0}
    v2 -= v2
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v10 uint64 = 18446744073709551615
        _ = v10
    case 4294967295:
        v1 -= v1
    default:
        var v11 int16 = (-(v1))
        _ = v11
    }
    v12 := GenS3{F0: 7}
    var v13 int64 = -9223372036854775808
    switch "a" {
    case "a":
        if false {
            v13 -= v13
        } else {
            if true {
                if (-1 < -996) {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            var v16 uint32 = 4294967295
                            var v17 string = ("a" + "\n\t")
                            delete(v3, 0)
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        v2 += v2
                        v18 := GenS1{F0: 127}
                        _ = fi14
                        _ = v18
                    }
                    switch "foo bar" {
                    case "a":
                        var v19 uint32 = 42
                        _ = v19
                    case "\n\t":
                        for fi20 := 0; fi20 < 1; fi20++ {
                            v4 += 647
                            _ = fi20
                        }
                    default:
                        v21 := GenS1{F0: -1}
                        _ = v21
                    }
                    v13 = 9223372036854775807
                }
            }
        }
    case "":
        v22 := GenS4{F0: 128, F1: 1, F2: -83}
        _ = v22
    case "hello":
        if true {
            v23 := GenS4{F0: 2, F1: -752, F2: 29}
            v24 := GenS2{F0: v4}
            _ = v23
            _ = v24
        } else {
            v25 := GenS5{F0: 127, F1: true, F2: (("foo bar" + ("hello" + "\n\t")) + "")}
            if ((v4 | v4) >= 42) {
                fmt.Println(v2, v1)
                switch "foo bar" {
                case "foo bar":
                    v5[-2147483648] = 2147483648
                }
                v4 -= (v4 - v4)
            } else {
                var v26 bool = false
                _ = v26
            }
            _ = v25
        }
    }
    var v27 int64 = v13
    v2 -= (v2 >> ((2) & 63))
    v13 = 0
    for fi28 := 0; fi28 < 15; fi28++ {
        v29 := GenS4{F0: 86, F1: int(v2), F2: -83}
        var v30 bool = (!(255 >= v2))
        v1 -= -882
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
