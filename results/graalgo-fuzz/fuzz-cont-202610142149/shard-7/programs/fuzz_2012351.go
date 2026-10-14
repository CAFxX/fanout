package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -1
    var v3 string = (("foo bar" + ("foo bar" + ("" + (("foo bar" + "a") + "")))) + ((("\n\t" + (("a" + ((("foo bar" + ("" + "hello")) + "\n\t") + "a")) + "hello")) + "") + "foo bar"))
    v4 := []uint32{65535}
    var v5 int = 4294967295
    switch "" {
    case "hello":
        v2 += v2
    case "\n\t":
        var v6 float64 = 1e10
        _ = v6
    default:
        if ((542 <= 18446744073709551615) || true) {
            var v7 int8 = (v2 % ((100) | 1))
            v8 := GenS1{F0: 930, F1: 984, F2: v5}
            _ = v7
            _ = v8
        } else {
            var v9 int16 = 749
            _ = v9
        }
    }
    var v10 int8 = v2
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS1{F0: 0, F1: 255, F2: fi11}
        _ = fi11
        _ = v12
    }
    switch (v5 | (v5 & v5)) {
    case 42:
        var v13 uint8 = 255
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v10 = (-(v2))
        _ = fi14
    }
    var v15 int64 = 9223372036854775807
    var v16 int = v5
    var v17 uint32 = 4294967295
    switch int(v17) {
    case -407:
        if (997 == 42) {
            switch -932 {
            case 127:
                if (!false) {
                    v2 = (v2 / ((v10) | 1))
                }
            case -842:
                for fi18 := 0; fi18 < 9; fi18++ {
                    for fi19 := 0; fi19 < 3; fi19++ {
                        v10 = (v2 + (v2 % (((v2 % ((v2) | 1))) | 1)))
                        _ = fi19
                    }
                    v20 := v1[(^(v17))]
                    _ = fi18
                    _ = v20
                }
            }
        }
    case 539:
        var v21 int32 = 1
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := GenS2{F0: true, F1: (^(fi22))}
        _ = fi22
        _ = v23
    }
    v24 := GenS3{F0: v17}
    switch 128 {
    case 154:
        v25 := GenS3{F0: 4294967295}
        _ = v25
    case -203:
        var v26 uint8 = 0
        _ = v26
    default:
        if false {
            var v27 float32 = (1.5 - 3.14159)
            _ = v27
        } else {
            for fi28 := 0; fi28 < 14; fi28++ {
                v5 = -9223372036854775808
                _ = len(v1)
                _ = fi28
            }
            switch (v5 - v16) {
            case 1:
                if true {
                    for fi29 := 0; fi29 < 5; fi29++ {
                        delete(v1, (^(v17)))
                        _ = fi29
                    }
                    for fi30 := 0; fi30 < 10; fi30++ {
                        v31 := GenS4{F0: (v15 >> ((700) & 63)), F1: 18446744073709551615, F2: -1}
                        _ = fi30
                        _ = v31
                    }
                }
            case 127:
                for fi32 := 0; fi32 < 17; fi32++ {
                    if false {
                        if len(v4) > 0 {
                            v4[0] = (v17 / ((v17) | 1))
                        }
                        v2 = 127
                    }
                    v33 := v1[(v17 | v17)]
                    var v34 uint16 = 977
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
        }
    }
    switch (v5 % ((32767) | 1)) {
    case 2147483647:
        v35 := GenS4{F0: (v15 + v15), F1: 2, F2: (v15 & -432)}
        _ = v35
    case -9223372036854775808:
        fmt.Println(v15, v2, v15)
    case -2147483648:
        v5 -= v5
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
