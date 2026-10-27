package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int32; F1 int8; F2 uint32 }

func main() {
    var v1 int32 = -830
    var v2 int64 = 42
    v3 := make(map[int]string)
    var v4 uint16 = 2
    for fi5 := 0; fi5 < 11; fi5++ {
        switch "" {
        case "hello":
            v6 := []int{954, 42, fi5}
            _ = v6
        case "":
            var v7 float32 = 3.14159
            _ = v7
        case "\n\t":
            for fi8 := 0; fi8 < 6; fi8++ {
                if true {
                    var v9 bool = (false || false)
                    _ = v9
                } else {
                    v10 := GenS1{F0: 530, F1: 0, F2: 271}
                    _ = v10
                }
                v11 := []int{(fi5 / (((fi5 - (fi5 << ((944) & 63)))) | 1))}
                _ = fi8
                _ = v11
            }
        }
        var v12 uint64 = 372
        fmt.Println(v4, fi5, v1)
        _ = fi5
        _ = v12
    }
    fmt.Println(v4, v2)
    switch -9223372036854775808 {
    case 42:
        v13 := GenS2{F0: (!(v2 == v2)), F1: 0}
        _ = v13
    default:
        if true {
            fmt.Println(v2, v1, v4)
            for fi14 := 0; fi14 < 20; fi14++ {
                v2 = v2
                for fi15 := 0; fi15 < 15; fi15++ {
                    v1 += v1
                    v1 = v1
                    _ = fi15
                }
                _ = fi14
            }
            fmt.Println(v1, v1)
        }
    }
    var v16 uint64 = 0
    var v17 float64 = (1e10 * 1.5)
    v4 -= (^(v4))
    fmt.Println(v16, v2, v1)
    switch ("foo bar" + "hello") {
    case "\n\t":
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 uint64 = (v16 / ((v16) | 1))
            _ = fi18
            _ = v19
        }
    case "foo bar":
        if true {
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 string = ""
                for fi22 := 0; fi22 < 15; fi22++ {
                    if (false && true) {
                        v1 = (v1 - -2147483648)
                        v23 := GenS1{F0: 42, F1: v16, F2: 734}
                        v24 := v3[(fi20 >> ((171) & 63))]
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(v16)
                    if ((true && ((1.5 - 1e10) == 0.1)) && (uint8(v4) < 204)) {
                        v2 += (v2 >> ((2) & 63))
                    } else {
                        fmt.Println(v2, v16)
                        var v25 int32 = 0
                        _ = v25
                    }
                    v26 := GenS2{F0: (!(!(v2 != v2))), F1: 1}
                    _ = fi22
                    _ = v26
                }
                _ = fi20
                _ = v21
            }
            switch ("a" + "") {
            case "hello":
                switch ("" + "foo bar") {
                case "foo bar":
                    var v27 int32 = 1
                    _ = v27
                case "hello":
                    fmt.Println(v17, v4)
                case "a":
                    v28 := GenS1{F0: 2147483648, F1: 1, F2: 2}
                    _ = v28
                }
            case "foo bar":
                v29 := GenS1{F0: 974, F1: v16, F2: uint(v4)}
                _ = v29
            default:
                v3[42] = ("foo bar" + "")
            }
        } else {
            var v30 int = -9223372036854775808
            v31 := []int{v30, (-(v30)), int(v4), (v30 / (((v30 / ((v30) | 1))) | 1))}
            _ = v30
            _ = v31
        }
    }
    v4 -= v4
    v2 = v2
    v16 = v16
    fmt.Println(v4)
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 bool = true
        var v34 int64 = v2
        _ = fi32
        _ = v33
        _ = v34
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        if true {
            var v35 int8 = 1
            _ = v35
        }
    case "":
        v17 = float64(1.5)
    default:
        v36 := GenS3{F0: v1, F1: 0, F2: 2}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
