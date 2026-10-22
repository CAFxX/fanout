package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 int64; F2 int64 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int64; F1 string }
func genFn0() int16 {
    var v7 uint16 = 912
    fmt.Println(v7)
    v8 := []int{-552, 2147483647, 32767, 42}
    if true {
        _ = len(v8)
    } else {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    var v9 float64 = -2.25
    _ = v7
    _ = v8
    _ = v9
    return -1
}
func genFn1(p0 uint16, p1 uint, p2 bool) int8 {
    var v10 string = ("a" + "hello")
    for fi11 := 0; fi11 < 18; fi11++ {
        v12 := []int{fi11, (fi11 >> ((p1) & 63)), (fi11 + fi11)}
        if p2 {
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 uint32 = 32
                v15 := GenS1{F0: 32767}
                _ = fi13
                _ = v14
                _ = v15
            }
            v16 := GenS2{F0: 1, F1: 1, F2: -1}
            switch v10 {
            case "a":
                p0 -= (p0 ^ (p0 * (p0 >> ((p1) & 63))))
            case "foo bar":
                p2 = p2
            }
            _ = v16
        } else {
            for fi17 := 0; fi17 < 8; fi17++ {
                if len(v12) > 0 {
                    v12[0] = (fi17 / (((-(fi11))) | 1))
                }
                _ = fi17
            }
            switch fi11 {
            case 4294967295:
                switch "foo bar" {
                case "hello":
                    v18 := GenS3{F0: 32768, F1: fi11}
                    _ = v18
                case "\n\t":
                    v12 = append(v12, 1)
                case "a":
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
            case 42:
                var v19 int = fi11
                _ = v19
            default:
                var v20 uint = 580
                _ = v20
            }
            for fi21 := 0; fi21 < 9; fi21++ {
                if false {
                    p0 += 255
                } else {
                    _ = len(v12)
                    p0 = (p0 << ((p1) & 63))
                }
                p2 = p2
                var v22 uint8 = 2
                _ = fi21
                _ = v22
            }
        }
        switch fi11 {
        case 2147483647:
            fmt.Println(p1, fi11, p2)
        default:
            switch fi11 {
            case -93:
                var v23 int32 = 42
                _ = v23
            case 906:
                var v24 int16 = -393
                _ = v24
            }
        }
        _ = fi11
        _ = v12
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        for fi26 := 0; fi26 < 3; fi26++ {
            v27 := GenS4{F0: 1, F1: (v10 + v10)}
            var v28 int8 = 0
            _ = fi26
            _ = v27
            _ = v28
        }
        switch ("" + (v10 + v10)) {
        case "a":
            v29 := []int{(fi25 >> ((p1) & 63))}
            _ = v29
        case "foo bar":
            p1 = (p1 << (((^(p1))) & 63))
        case "":
            if (3.14159 == ((1.5 * 1e10) * (1e10 / (((1e10 / 1.5) - (-2.25 / (3.14159 / 1e10))) / -2.25)))) {
                p1 += p1
                for fi30 := 0; fi30 < 13; fi30++ {
                    p0 += p0
                    fmt.Println(v10, p1, fi30)
                    v31 := GenS1{F0: -1}
                    _ = fi30
                    _ = v31
                }
                fmt.Println(p0, v10)
            }
        default:
            fmt.Println(v10, p1)
        }
        _ = fi25
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 127
}

func main() {
    var v1 uint64 = 107
    var v2 uint8 = 128
    var v3 int8 = 42
    var v4 int = 0
    var v5 int32 = -61
    v6 := []uint32{1, 2}
    var v32 int32 = 42
    fmt.Println(v1, v4, v3)
    v2 = (v2 << ((2) & 63))
    switch "foo bar" {
    case "foo bar":
        var v33 uint16 = 1
        _ = v33
    case "\n\t":
        if len(v6) > 0 {
            v6[0] = 2147483648
        }
    }
    var v34 int64 = 806
    var v35 int16 = 1000
    var v36 uint16 = 0
    v32 = 1
    var v37 float64 = 3.14159
    var v38 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
