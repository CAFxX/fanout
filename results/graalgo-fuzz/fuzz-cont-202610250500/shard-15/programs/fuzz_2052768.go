package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 uint32; F2 int16 }

func main() {
    var v1 int64 = 543
    var v2 uint64 = 0
    var v3 string = "a"
    var v4 uint16 = 0
    v1 += v1
    if true {
        v5 := GenS1{F0: v3, F1: -241, F2: 128}
        v1 -= v1
        v6 := GenS2{F0: (v4 * (^(v4)))}
        _ = v5
        _ = v6
    } else {
        if (255 == 1) {
            v7 := GenS3{F0: 80, F1: v3}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := []int{fi8, fi8, 9223372036854775807, fi8}
                var v10 int64 = v1
                v3 = ("\n\t" + "\n\t")
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        switch -916 {
        case -178:
            if true {
                v11 := []int{9223372036854775807, 127, 128, 9223372036854775807}
                for fi12 := 0; fi12 < 10; fi12++ {
                    if len(v11) > 0 {
                        v11[0] = fi12
                    }
                    _ = fi12
                }
                if (!(0 < -9223372036854775808)) {
                    for fi13 := 0; fi13 < 4; fi13++ {
                        _ = len(v11)
                        for fi14 := 0; fi14 < 4; fi14++ {
                            var v15 uint32 = 424
                            var v16 uint8 = 128
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        var v17 bool = true
                        _ = fi13
                        _ = v17
                    }
                }
                _ = v11
            } else {
                switch v3 {
                case "foo bar":
                    v18 := GenS4{F0: 203}
                    _ = v18
                case "a":
                    switch "a" {
                    case "a":
                        v3 = v3
                    }
                case "hello":
                    v19 := []int{-9223372036854775808}
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 12; fi20++ {
                        v21 := GenS4{F0: 0}
                        _ = fi20
                        _ = v21
                    }
                }
                fmt.Println(v1)
            }
        case 1:
            var v22 uint64 = v2
            _ = v22
        case 2147483647:
            if (!false) {
                fmt.Println(v1, v1, v3)
            }
        }
        fmt.Println(v3, v3, v4)
    }
    v1 -= (v1 ^ (v1 & v1))
    var v23 string = (v3 + v3)
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 uint8 = 255
        _ = fi24
        _ = v25
    }
    fmt.Println(v1)
    if false {
        if false {
            v26 := []int{128, 0, 42}
            _ = v26
        } else {
            if (-32768 < 32767) {
                fmt.Println(v3)
            }
        }
        var v27 float32 = 3.14159
        _ = v27
    }
    v4 += (v4 * v4)
    v28 := []int{581, 0, 342}
    switch "" {
    case "a":
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 int = fi29
            fmt.Println(v30, v23)
            if (42 < 59) {
                v31 := GenS5{F0: 95, F1: 628, F2: 275}
                _ = v31
            }
            _ = fi29
            _ = v30
        }
    case "\n\t":
        v23 = v3
    default:
        var v32 int8 = 100
        _ = v32
    }
    var v33 bool = (229 >= 0)
    switch ((v3 + "hello") + (v3 + (v23 + v3))) {
    case "":
        var v34 bool = true
        _ = v34
    case "foo bar":
        v23 = v23
    }
    v3 = "\n\t"
    v35 := GenS4{F0: 127}
    v36 := GenS3{F0: 1, F1: (v23 + v3)}
    fmt.Println(v33)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v4
}
