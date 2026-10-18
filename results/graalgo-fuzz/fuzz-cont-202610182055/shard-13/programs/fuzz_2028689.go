package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 string; F1 uint8; F2 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "foo bar"
    var v3 float64 = (((float64(1.5) / 1.0) / -2.25) / 1.0)
    var v4 int = 899
    v5 := GenS1{F0: 0, F1: 65535}
    v2 = "\n\t"
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int8 = 0
        var v8 int32 = -1
        if (0 <= -32768) {
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 uint = 483
                fmt.Println(v10, v8, v10)
                _ = fi9
                _ = v10
            }
        } else {
            for fi11 := 0; fi11 < 9; fi11++ {
                v7 += (v7 * 1)
                switch (v2 + v2) {
                case "hello":
                    v12 := GenS1{F0: 2, F1: 164}
                    _ = v12
                case "":
                    v2 = "hello"
                }
                _ = fi11
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v4 -= (v4 - (v4 >> ((764) & 63)))
    v13 := GenS1{F0: 255, F1: 42}
    v1[0] = 966
    var v14 string = v2
    fmt.Println(v2, v3)
    fmt.Println(v3, v14, v4)
    var v15 uint = 2
    delete(v1, 2147483648)
    switch ((v2 + v2) + v14) {
    case "foo bar":
        v2 = v2
    case "a":
        switch ("foo bar" + ((v2 + (v2 + "\n\t")) + ("" + "hello"))) {
        case "\n\t":
            v16 := GenS1{F0: 127, F1: 2147483647}
            _ = v16
        case "foo bar":
            var v17 uint = (v15 ^ 65535)
            _ = v17
        case "":
            v3 = v3
        default:
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 uint64 = 18446744073709551615
                fmt.Println(v14)
                for fi20 := 0; fi20 < 12; fi20++ {
                    var v21 uint64 = v19
                    _ = fi20
                    _ = v21
                }
                _ = fi18
                _ = v19
            }
        }
    case "\n\t":
        if (false && (-9223372036854775808 == 9223372036854775807)) {
            switch v4 {
            case 32768:
                fmt.Println(v14, v15)
            case 42:
                var v22 int64 = 9223372036854775807
                _ = v22
            default:
                v2 = "hello"
            }
            v4 = v4
        }
    default:
        var v23 float64 = 1e10
        _ = v23
    }
    if (993 >= 399) {
        switch (v14 + ((v14 + v2) + "hello")) {
        case "a":
            if false {
                var v24 float32 = (float32(((float64(0.0) - ((v3 / (v3 - v3)) + v3)) - ((v3 + v3) / v3))) + -2.25)
                _ = v24
            }
        case "\n\t":
            switch "a" {
            case "":
                var v25 int = 128
                _ = v25
            case "hello":
                var v26 uint = 65535
                _ = v26
            case "foo bar":
                if true {
                    var v27 bool = (-917 <= -1)
                    fmt.Println(v4)
                    _ = v27
                } else {
                    v3 -= 1e10
                    for fi28 := 0; fi28 < 17; fi28++ {
                        fmt.Println(fi28, v2, v4)
                        _ = fi28
                    }
                    if (!(128 < 1)) {
                        for fi29 := 0; fi29 < 19; fi29++ {
                            v30 := GenS2{F0: "foo bar", F1: 114, F2: int32(v15)}
                            fmt.Println(v14, v3, fi29)
                            var v31 int64 = int64(v4)
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        for fi32 := 0; fi32 < 5; fi32++ {
                            v2 = ("hello" + "foo bar")
                            v15 += 65535
                            _ = fi32
                        }
                        v2 = v14
                    } else {
                        var v33 uint16 = 65535
                        fmt.Println(v2, v33, v14)
                        var v34 int16 = 42
                        _ = v33
                        _ = v34
                    }
                }
            default:
                var v35 uint = 65535
                _ = v35
            }
        default:
            if ((v4 * (v4 * (v4 & v4))) < (^(v4))) {
                fmt.Println(v3, v3)
                if (255 < uint8(v15)) {
                    v36 := GenS1{F0: 1, F1: uint32(v15)}
                    _ = v36
                }
            }
        }
        v37 := GenS2{F0: v14, F1: 42, F2: -2147483648}
        _ = v37
    }
    v2 = ("" + (v2 + "hello"))
    var v38 uint32 = 194
    v39 := GenS1{F0: 195, F1: (^(v38))}
    v40 := []int{(v4 & -1)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
