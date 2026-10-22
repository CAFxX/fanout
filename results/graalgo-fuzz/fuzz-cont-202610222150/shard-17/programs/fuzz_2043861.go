package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int32; F1 bool; F2 int16 }
type GenS3 struct { F0 uint32; F1 uint8; F2 bool }
type GenS4 struct { F0 string; F1 uint16; F2 uint32 }
func genFn0(p0 uint32, p1 string, p2 uint16) uint16 {
    fmt.Println(p2, p2)
    v7 := []int{32768, 0}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 + 32767)
}
func genFn1(p0 int32, p1 uint16, p2 int16) int {
    p0 = p0
    switch 4294967295 {
    case 32767:
        v8 := []int{-9223372036854775808, 569}
        _ = v8
    case 1:
        fmt.Println(p1, p2)
    default:
        v9 := GenS1{F0: 1000, F1: 1}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 uint = 572
    var v2 uint16 = 255
    v3 := make(map[int]int)
    var v4 int8 = 42
    var v5 uint8 = 18
    var v6 int16 = 1
    v6 = (^(v6))
    if (!false) {
        for fi10 := 0; fi10 < 1; fi10++ {
            v2 = genFn0(0, "foo bar", 42)
            _ = fi10
        }
        v11 := []int{32767, int(v5), 0, 128}
        _ = v11
    }
    var v12 uint8 = 2
    var v13 int = genFn1(2147483647, v2, v6)
    var v14 float32 = (((3.14159 + 0.0) - ((1.5 * (3.14159 - 3.14159)) + 0.1)) * ((0.0 + float32(1.5)) / (-2.25 * (0.1 + -2.25))))
    v4 += 100
    var v15 uint64 = 42
    if true {
        var v16 int8 = (v4 ^ v4)
        v2 -= genFn0(0, "", v2)
        _ = v16
    } else {
        switch "foo bar" {
        case "foo bar":
            var v17 string = "\n\t"
            _ = v17
        default:
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                v5 += 235
            case "hello":
                if (2 != (v1 << (((v1 - (v1 & 65535))) & 63))) {
                    v14 += (float32(3.14159) * (v14 * float32((3.14159 - (float64(0.1) * 0.0)))))
                } else {
                    var v18 uint16 = v2
                    v19 := v3[(v13 % (((v13 * int(v5))) | 1))]
                    switch -534 {
                    case 128:
                        fmt.Println(v13)
                    case 32768:
                        v19 = (-(v13))
                    }
                    _ = v18
                    _ = v19
                }
            case "":
                var v20 uint16 = 187
                _ = v20
            default:
                var v21 int32 = -754
                _ = v21
            }
        }
    }
    v22 := []int{}
    switch "\n\t" {
    case "hello":
        fmt.Println(v12, v15, v6)
    case "":
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            v4 = v4
        case "foo bar":
            fmt.Println(v15)
        default:
            v13 += (v13 + v13)
        }
    case "\n\t":
        fmt.Println(v12, v1, v14)
    }
    v2 -= uint16(v15)
    v12 -= (v12 + (v12 % ((42) | 1)))
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 uint64 = v15
        if ((0 > (^(v15))) || (!((float64(v14) - 3.14159) >= (0.0 * 0.1)))) {
            for fi25 := 0; fi25 < 15; fi25++ {
                switch ("hello" + ("foo bar" + ("\n\t" + ""))) {
                case "a":
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                default:
                    if true {
                        var v26 uint64 = (v24 << (((v1 / ((1) | 1))) & 63))
                        _ = v26
                    } else {
                        fmt.Println(v5, v5)
                        var v27 uint8 = v12
                        v5 = v12
                        v1 = v1
                        if true {
                            v14 = 0.0
                            var v28 int32 = 0
                            v27 = uint8(v1)
                            _ = v28
                        }
                        v29 := GenS2{F0: 2147483647, F1: true, F2: (v6 >> ((1) & 63))}
                        _ = v27
                        _ = v29
                    }
                }
                _ = fi25
            }
            v14 += (1e10 * v14)
        }
        _ = fi23
        _ = v24
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        switch genFn1(-642, 32767, (v6 & v6)) {
        case 1:
            for fi31 := 0; fi31 < 2; fi31++ {
                switch (fi31 - 128) {
                case 42:
                    v22 = append(v22, fi31)
                case 92:
                    v3[0] = fi31
                case 32767:
                    v32 := GenS1{F0: 31, F1: -226}
                    _ = v32
                default:
                    var v33 bool = false
                    _ = v33
                }
                var v34 float64 = -2.25
                _ = len(v3)
                var v35 uint8 = (^(v12))
                var v36 bool = (false || (float64(-2.25) <= (v34 + ((-2.25 / 1.0) * v34))))
                _ = fi31
                _ = v34
                _ = v35
                _ = v36
            }
        }
        _ = fi30
    }
    v37 := GenS3{F0: 1, F1: v12, F2: (false || false)}
    fmt.Println(v14)
    v38 := GenS4{F0: ("a" + "foo bar"), F1: v2, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
