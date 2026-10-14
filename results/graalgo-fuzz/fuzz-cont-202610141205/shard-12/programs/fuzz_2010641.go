package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 string }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 bool }
func genFn0(p0 int64) int64 {
    var v4 uint64 = 4294967295
    switch "a" {
    case "":
        p0 = p0
    default:
        p0 = p0
    }
    v5 := []int{-262, 808}
    _ = p0
    _ = v4
    _ = v5
    return (p0 - (p0 & (-(p0))))
}

func main() {
    var v1 uint32 = 210
    var v2 int8 = 42
    var v3 int32 = -861
    v6 := []int{}
    if (!(0 <= 42)) {
        for fi7 := 0; fi7 < 15; fi7++ {
            v1 += v1
            v3 += (v3 * (v3 * (v3 ^ v3)))
            _ = fi7
        }
        v3 += (-(v3))
    } else {
        v3 = v3
    }
    v8 := GenS1{F0: (v2 | (v2 ^ 42)), F1: 42, F2: 2}
    var v9 uint32 = (v1 * v1)
    for fi10 := 0; fi10 < 1; fi10++ {
        v9 += (v9 & 65535)
        if ((!false) || ((fi10 << ((42) & 63)) < 474)) {
            fmt.Println(v3, v2, v1)
            var v11 uint = 278
            _ = v11
        }
        if false {
            for fi12 := 0; fi12 < 19; fi12++ {
                if true {
                    var v13 uint16 = 0
                    fmt.Println(v1, v13, v2)
                    v2 = v2
                    _ = v13
                }
                _ = fi12
            }
        } else {
            var v14 float32 = 0.1
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 string = "a"
                v17 := GenS2{F0: 945, F1: -143, F2: (v16 + v16)}
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        }
        _ = fi10
    }
    v1 = v9
    v18 := GenS1{F0: (v2 >> ((81) & 63)), F1: 946, F2: 0}
    fmt.Println(v2, v3)
    _ = len(v6)
    v3 -= (^(v3))
    for fi19 := 0; fi19 < 15; fi19++ {
        if ((42 != 752) && (v9 <= (v9 >> ((255) & 63)))) {
            var v20 float64 = 1.5
            _ = v20
        } else {
            v1 -= (v1 >> ((42) & 63))
            var v21 uint = 4294967295
            if len(v6) > 0 {
                v6[0] = 128
            }
            _ = v21
        }
        v3 = v3
        _ = fi19
    }
    fmt.Println(v9, v2)
    v22 := GenS3{F0: 952, F1: int16(v2)}
    v23 := GenS4{F0: 65535, F1: 1}
    v24 := GenS2{F0: 1, F1: (v3 & v3), F2: "a"}
    if true {
        v1 = 42
        switch -1 {
        case 841:
            for fi25 := 0; fi25 < 8; fi25++ {
                for fi26 := 0; fi26 < 11; fi26++ {
                    v1 -= (v1 & v1)
                    var v27 float32 = 0.0
                    _ = fi26
                    _ = v27
                }
                var v28 uint8 = 108
                _ = len(v6)
                _ = fi25
                _ = v28
            }
        case 32767:
            if true {
                if (true || ((((0 == 0) && ((0 > (v3 >> ((255) & 63))) && (0.1 == (1e10 + 1.5)))) || false) || (255 >= 605))) {
                    var v29 uint32 = v1
                    _ = v29
                } else {
                    v30 := GenS5{F0: 127}
                    if false {
                        var v31 float64 = 3.14159
                        _ = v31
                    }
                    _ = v30
                }
                switch "" {
                case "foo bar":
                    v32 := GenS2{F0: 224, F1: v3, F2: "hello"}
                    _ = v32
                }
                v33 := GenS6{F0: ("a" + ("a" + "hello")), F1: ("foo bar" == "")}
                _ = v33
            }
        case -9223372036854775808:
            var v34 uint64 = 4294967295
            _ = v34
        default:
            fmt.Println(v1, v1, v1)
        }
        fmt.Println(v9, v1)
    }
    if true {
        if false {
            var v35 float64 = float64(float32(-2.25))
            if false {
                v36 := GenS7{F0: (true && false)}
                v6 = append(v6, 32768)
                _ = v36
            } else {
                if false {
                    if len(v6) > 0 {
                        v6[0] = 32767
                    }
                    fmt.Println(v35)
                    v6 = append(v6, 32768)
                }
                if true {
                    v9 = v9
                }
                fmt.Println(v3)
            }
            v1 = v9
            _ = v35
        }
    } else {
        v6 = append(v6, 32768)
        v37 := GenS1{F0: -30, F1: 291, F2: 42}
        v6 = append(v6, 127)
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v6
    _ = v8
    _ = v9
}
