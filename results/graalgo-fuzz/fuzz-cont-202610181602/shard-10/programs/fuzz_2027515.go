package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int16; F2 uint32 }
type GenS3 struct { F0 int16; F1 int32; F2 bool }
type GenS4 struct { F0 int64; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int, p1 uint16) uint8 {
    fmt.Println(p1, p1)
    v5 := GenS1{F0: p0}
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS2{F0: 9223372036854775807, F1: 1000, F2: 2147483648}
        v8 := GenS1{F0: p0}
        v9 := GenS2{F0: -9223372036854775808, F1: 42, F2: 0}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    return 97
}

func main() {
    var v1 uint = 947
    var v2 uint16 = 1
    v3 := make(map[int]string)
    var v4 uint8 = 117
    switch 113 {
    case -2147483648:
        if true {
            v10 := []int{1, -2147483648}
            var v11 int64 = 228
            for fi12 := 0; fi12 < 18; fi12++ {
                if true {
                    v11 = (v11 & -332)
                }
                _ = fi12
            }
            _ = v10
            _ = v11
        }
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        if (true && (!false)) {
            if ((v1 % ((0) | 1)) >= v1) {
                for fi14 := 0; fi14 < 14; fi14++ {
                    var v15 uint16 = (^(v2))
                    v15 = 0
                    _ = fi14
                    _ = v15
                }
                v4 -= (^(v4))
            }
            if false {
                v16 := []int{}
                var v17 uint8 = genFn0(fi13, uint16(v4))
                v18 := GenS2{F0: 0, F1: 620, F2: 65535}
                _ = v16
                _ = v17
                _ = v18
            } else {
                v2 -= (v2 - (^(v2)))
                _ = len(v3)
                for fi19 := 0; fi19 < 9; fi19++ {
                    v4 = (v4 * (v4 + 2))
                    fmt.Println(fi19)
                    v20 := []int{fi13}
                    _ = fi19
                    _ = v20
                }
            }
        } else {
            v21 := GenS3{F0: 42, F1: 1, F2: (((1.5 / 1.0) - 3.14159) > 0.0)}
            var v22 uint16 = (v2 - v2)
            _ = v21
            _ = v22
        }
        _ = fi13
    }
    v4 += genFn0(32767, v2)
    v23 := GenS3{F0: 1000, F1: 2147483647, F2: true}
    v1 = (v1 + v1)
    var v24 string = ("foo bar" + "a")
    v25 := []int{-984, 128}
    var v26 float32 = 1.5
    v27 := GenS2{F0: -157, F1: 254, F2: 65535}
    for fi28 := 0; fi28 < 13; fi28++ {
        v3[505] = v24
        v29 := GenS3{F0: -32768, F1: -2147483648, F2: true}
        _ = fi28
        _ = v29
    }
    if (!false) {
        v26 = (v26 / v26)
    } else {
        v26 = 3.14159
    }
    v30 := GenS4{F0: -1, F1: 65535, F2: -352}
    fmt.Println(v26, v26, v4)
    if false {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v1)
        case "hello":
            fmt.Println(v24)
        case "\n\t":
            for fi31 := 0; fi31 < 16; fi31++ {
                v24 = ""
                for fi32 := 0; fi32 < 19; fi32++ {
                    var v33 float32 = float32((0.1 - 0.1))
                    _ = fi32
                    _ = v33
                }
                _ = fi31
            }
        default:
            var v34 uint64 = 4294967295
            _ = v34
        }
        v1 = v1
    } else {
        for fi35 := 0; fi35 < 15; fi35++ {
            _ = len(v25)
            if ((v26 / v26) < (v26 * ((((float32(-2.25) + v26) - 0.1) + v26) + v26))) {
                v26 = (v26 - (float32(-2.25) * v26))
                var v36 float64 = (float64(v26) + 0.1)
                _ = v36
            } else {
                var v37 int64 = -1
                _ = v37
            }
            _ = fi35
        }
        for fi38 := 0; fi38 < 7; fi38++ {
            v39 := GenS5{F0: (v2 ^ (v2 >> ((65535) & 63)))}
            _ = fi38
            _ = v39
        }
        for fi40 := 0; fi40 < 10; fi40++ {
            var v41 int = -1
            for fi42 := 0; fi42 < 19; fi42++ {
                v24 = ((((v24 + v24) + (("foo bar" + (v24 + (v24 + v24))) + ("\n\t" + "a"))) + v24) + v24)
                _ = len(v3)
                v24 = ((("a" + (v24 + v24)) + (((v24 + "a") + ("hello" + v24)) + v24)) + "\n\t")
                _ = fi42
            }
            fmt.Println(v4, v4)
            _ = fi40
            _ = v41
        }
    }
    var v43 uint = v1
    for fi44 := 0; fi44 < 14; fi44++ {
        if (!false) {
            v43 += v43
            fmt.Println(v26, fi44, v4)
            if len(v25) > 0 {
                v25[0] = fi44
            }
        } else {
            fmt.Println(v4, v1, fi44)
            v45 := GenS3{F0: 661, F1: 42, F2: (fi44 <= fi44)}
            for fi46 := 0; fi46 < 14; fi46++ {
                if false {
                    var v47 int64 = -9223372036854775808
                    delete(v3, 32768)
                    var v48 int64 = v47
                    _ = v47
                    _ = v48
                }
                _ = len(v3)
                var v49 string = v24
                _ = fi46
                _ = v49
            }
            _ = v45
        }
        if ((v1 % (((^(v1))) | 1)) != 1) {
            if false {
                var v50 uint32 = 725
                for fi51 := 0; fi51 < 20; fi51++ {
                    fmt.Println(fi44, fi44)
                    v4 += 1
                    fmt.Println(v50, fi51)
                    v26 += float32(-2.25)
                    _ = fi51
                }
                v52 := GenS5{F0: v2}
                _ = v50
                _ = v52
            } else {
                v2 = (^(v2))
            }
            var v53 int16 = 1
            _ = v53
        }
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v43
}
