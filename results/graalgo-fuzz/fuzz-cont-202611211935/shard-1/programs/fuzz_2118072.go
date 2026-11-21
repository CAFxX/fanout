package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int64 }
type GenS2 struct { F0 int16 }
func genFn0() uint16 {
    var v6 uint16 = 405
    for fi7 := 0; fi7 < 15; fi7++ {
        switch "hello" {
        case "\n\t":
            if false {
                var v8 uint64 = 727
                if true {
                    var v9 int = (fi7 >> ((42) & 63))
                    _ = v9
                } else {
                    v10 := []int{(fi7 / (((-(fi7))) | 1)), fi7, fi7}
                    _ = v10
                }
                _ = v8
            }
        case "hello":
            v11 := []int{}
            _ = v11
        default:
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := []int{-675, fi7, fi12}
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                v6 -= v6
                _ = fi12
                _ = v13
            }
        }
        _ = fi7
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        switch "a" {
        case "":
            v15 := []int{(fi14 % ((fi14) | 1))}
            _ = v15
        case "\n\t":
            v16 := GenS1{F0: -2, F1: 0, F2: -9223372036854775808}
            _ = v16
        default:
            v17 := GenS1{F0: 116, F1: -2147483648, F2: 42}
            _ = v17
        }
        v6 -= 65535
        v18 := GenS1{F0: 100, F1: -1, F2: 0}
        _ = fi14
        _ = v18
    }
    _ = v6
    return 0
}
func genFn1(p0 uint16, p1 bool, p2 uint32) int64 {
    p0 = (p0 | p0)
    var v19 bool = (565 == 1)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return -9223372036854775808
}

func main() {
    var v1 int8 = 1
    var v2 float32 = 3.14159
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    v5 := []uint32{2147483647, 885, 65535}
    if false {
        for fi20 := 0; fi20 < 7; fi20++ {
            fmt.Println(v2)
            _ = fi20
        }
    }
    var v21 int64 = -9223372036854775808
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 uint8 = 127
        _ = fi22
        _ = v23
    }
    fmt.Println(v1, v1, v1)
    v24 := GenS2{F0: 1}
    v21 = v21
    switch 32767 {
    case -2147483648:
        v2 = 1.5
    case 42:
        var v25 int64 = (v21 % ((v21) | 1))
        _ = v25
    default:
        for fi26 := 0; fi26 < 6; fi26++ {
            v27 := GenS1{F0: (v1 >> ((65535) & 63)), F1: 42, F2: (v21 ^ v21)}
            var v28 int32 = -2147483648
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    v1 = -128
    var v29 float32 = v2
    var v30 int16 = -32768
    switch 833 {
    case -1:
        v31 := GenS2{F0: v30}
        _ = v31
    case 32768:
        var v32 uint16 = 0
        _ = v32
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        if true {
            v34 := GenS2{F0: -370}
            _ = v34
        }
        switch ("a" + "") {
        case "hello":
            switch fi33 {
            case 1:
                var v35 bool = false
                _ = v35
            case -381:
                fmt.Println(v2)
            case -341:
                var v36 uint32 = 4294967295
                _ = v36
            default:
                if (v2 == (0.0 - v29)) {
                    v3[2147483647] = -519
                    var v37 string = "foo bar"
                    fmt.Println(v37, fi33, v21)
                    _ = v37
                } else {
                    var v38 uint = 177
                    _ = v38
                }
            }
        case "foo bar":
            if (true && (!false)) {
                v39 := GenS1{F0: v1, F1: -2147483648, F2: (v21 >> ((65535) & 63))}
                _ = v39
            }
        case "\n\t":
            v40 := v3[810]
            _ = v40
        default:
            fmt.Println(v2, v30, fi33)
        }
        _ = fi33
    }
    v5 = append(v5, 2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
