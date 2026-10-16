package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32; F1 int64 }

func main() {
    v1 := []string{"", "hello"}
    var v2 uint32 = 2147483647
    var v3 uint8 = 62
    v4 := make(map[uint32]bool)
    var v5 int32 = 2147483647
    v2 -= 406
    switch ("foo bar" + "") {
    case "\n\t":
        v6 := GenS1{F0: int8(v5)}
        _ = v6
    case "hello":
        var v7 int = 9223372036854775807
        _ = v7
    case "foo bar":
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint8 = 0
            var v10 int16 = 0
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        if false {
            _ = len(v4)
            delete(v4, v2)
            v2 = v2
        }
    }
    switch ("a" + ("a" + "")) {
    case "\n\t":
        var v11 int64 = -183
        _ = v11
    }
    v12 := GenS1{F0: -94}
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 uint64 = 1
            _ = fi14
            _ = v15
        }
        v16 := GenS2{F0: v5, F1: 42}
        _ = fi13
        _ = v16
    }
    v3 = 255
    v2 -= v2
    switch "a" {
    case "\n\t":
        for fi17 := 0; fi17 < 19; fi17++ {
            fmt.Println(v3, v5, v2)
            v3 -= (v3 + 255)
            _ = fi17
        }
    default:
        switch "" {
        case "foo bar":
            var v18 uint8 = v3
            _ = v18
        case "\n\t":
            if (255 != 1) {
                v19 := GenS3{F0: -1}
                var v20 uint64 = 0
                _ = v19
                _ = v20
            }
        }
    }
    var v21 string = ""
    for fi22 := 0; fi22 < 14; fi22++ {
        if false {
            switch fi22 {
            case 0:
                switch 4294967295 {
                case 32767:
                    fmt.Println(v2)
                case 692:
                    v2 += v2
                default:
                    if (!true) {
                        var v23 int32 = (v5 + (v5 & (v5 << ((65535) & 63))))
                        _ = v23
                    }
                }
            case -123:
                if false {
                    var v24 uint64 = 687
                    _ = v24
                } else {
                    _ = len(v4)
                    v3 -= (v3 * v3)
                }
            default:
                if (!(!true)) {
                    var v25 int8 = 53
                    _ = v25
                } else {
                    var v26 int8 = -57
                    var v27 uint16 = 42
                    var v28 float32 = float32(0.1)
                    _ = v26
                    _ = v27
                    _ = v28
                }
            }
            switch (v21 + v21) {
            case "":
                _ = len(v4)
            case "\n\t":
                fmt.Println(fi22, fi22, v3)
            case "hello":
                if true {
                    var v29 string = v21
                    if false {
                        v1 = append(v1, "foo bar")
                    }
                    _ = v29
                }
            default:
                v21 = (v21 + v21)
            }
            fmt.Println(v3, v3)
        } else {
            v21 = ("foo bar" + (((("foo bar" + ((("a" + v21) + v21) + (v21 + v21))) + v21) + "hello") + (v21 + "hello")))
            _ = len(v1)
        }
        _ = fi22
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        _ = len(v1)
        var v31 uint = 512
        v2 = (^(v2))
        _ = fi30
        _ = v31
    }
    v32 := GenS4{F0: 2}
    _ = len(v1)
    if (!(25 != 255)) {
        switch 32767 {
        case 1:
            v5 += int32(v3)
        }
        v1 = append(v1, v21)
    }
    v21 = (v21 + v21)
    if (true && false) {
        if true {
            v3 = (v3 / (((v3 / ((9) | 1))) | 1))
            if true {
                v2 = (v2 | (v2 & (v2 / ((2147483648) | 1))))
            }
        } else {
            v33 := GenS4{F0: 286}
            _ = v33
        }
        var v34 int = -118
        v3 = (v3 >> ((291) & 63))
        _ = v34
    }
    v35 := GenS5{F0: (v2 / ((2) | 1)), F1: 1}
    if (((-(v5)) <= (v5 | v5)) && ((v21 + (v21 + "hello")) == v21)) {
        for fi36 := 0; fi36 < 17; fi36++ {
            v4[v2] = false
            var v37 int = (fi36 - 2147483647)
            switch v21 {
            case "a":
                for fi38 := 0; fi38 < 1; fi38++ {
                    v4[(v2 + v2)] = (true && (-837 > 9223372036854775807))
                    v39 := GenS2{F0: 42, F1: 42}
                    v5 -= (v5 * (v5 * (v5 ^ v5)))
                    _ = fi38
                    _ = v39
                }
            }
            _ = fi36
            _ = v37
        }
        v2 += 0
        v3 += 0
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v5
}
