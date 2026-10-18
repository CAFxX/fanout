package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint }
func genFn0(p0 uint16, p1 uint, p2 int32) uint16 {
    fmt.Println(p2, p2)
    for fi6 := 0; fi6 < 9; fi6++ {
        switch -9223372036854775808 {
        case 32767:
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 uint = p1
                _ = fi7
                _ = v8
            }
        default:
            fmt.Println(fi6, fi6, fi6)
        }
        switch "hello" {
        case "foo bar":
            p0 = (p0 << (((p1 / ((884) | 1))) & 63))
        default:
            v9 := GenS1{F0: 0}
            _ = v9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            p0 += p0
            _ = fi10
        }
        _ = fi6
    }
    fmt.Println(p0)
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return (^(p0))
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{42, 0, 0}
    v3 := []int{-64, 818}
    v4 := make(map[int]bool)
    var v5 int = -2147483648
    if false {
        switch "" {
        case "":
            fmt.Println(v5)
        case "foo bar":
            var v11 float64 = (0.1 / 3.14159)
            _ = v11
        case "\n\t":
            var v12 int16 = 0
            _ = v12
        default:
            v5 = 32767
        }
        var v13 string = "foo bar"
        _ = v13
    }
    v14 := GenS1{F0: 444}
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 uint = 0
        _ = fi15
        _ = v16
    }
    var v17 int8 = -1
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 int64 = 42
            if false {
                v5 -= (fi19 ^ (v5 & -2147483648))
                var v21 uint16 = 255
                _ = v21
            }
            switch ("hello" + "") {
            case "foo bar":
                fmt.Println(v20)
            case "\n\t":
                v5 += (fi18 / ((-888) | 1))
            case "hello":
                if true {
                    fmt.Println(fi19)
                    fmt.Println(fi18)
                    v22 := GenS2{F0: 529, F1: 1}
                    _ = v22
                }
            }
            _ = fi19
            _ = v20
        }
        var v23 int16 = -1
        _ = fi18
        _ = v23
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    var v24 uint8 = 1
    switch (v5 - (-(v5))) {
    case -1:
        fmt.Println(v5, v5)
    case 1:
        v17 = (v17 << ((1) & 63))
    default:
        for fi25 := 0; fi25 < 9; fi25++ {
            if (!false) {
                v26 := GenS3{F0: 9223372036854775807, F1: int8(v24)}
                _ = v26
            } else {
                switch v5 {
                case 32768:
                    var v27 uint = 1
                    _ = v27
                case 9223372036854775807:
                    var v28 int32 = -1
                    _ = v28
                case 464:
                    for fi29 := 0; fi29 < 13; fi29++ {
                        var v30 uint64 = 0
                        switch -9223372036854775808 {
                        case 9223372036854775807:
                            v31 := GenS4{F0: true, F1: 456, F2: 4294967295}
                            _ = v31
                        }
                        var v32 uint8 = (^(v24))
                        _ = fi29
                        _ = v30
                        _ = v32
                    }
                }
                _ = len(v3)
            }
            _ = fi25
        }
    }
    v2 = append(v2, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
