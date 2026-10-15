package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint8; F2 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint32; F2 uint16 }
func genFn0() bool {
    if (0 < 338) {
        for fi4 := 0; fi4 < 13; fi4++ {
            v5 := GenS1{F0: 255, F1: -1, F2: false}
            _ = fi4
            _ = v5
        }
        var v6 int = -1
        _ = v6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        if false {
            var v8 uint64 = 2
            v9 := []int{fi7, (fi7 >> ((938) & 63)), fi7, (fi7 >> ((2) & 63))}
            fmt.Println(v8, fi7)
            _ = v8
            _ = v9
        } else {
            switch 127 {
            case -478:
                v10 := GenS1{F0: 65535, F1: -9223372036854775808, F2: false}
                _ = v10
            }
            v11 := GenS2{F0: -481}
            v12 := GenS3{F0: 42, F1: 163, F2: 1}
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    return false
}
func genFn1() string {
    var v13 int8 = -1
    v14 := []int{9223372036854775807, -492, -191}
    v15 := GenS2{F0: 0}
    _ = v13
    _ = v14
    _ = v15
    return (("hello" + "foo bar") + (("a" + "a") + "hello"))
}

func main() {
    var v1 uint = 42
    v2 := make(map[uint32]int)
    var v3 int64 = 1
    switch 32768 {
    case 535:
        v1 = 255
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        v3 = v3
        var v17 uint16 = 1
        _ = fi16
        _ = v17
    }
    v18 := []int{}
    if false {
        v19 := GenS4{F0: true}
        _ = v19
    }
    var v20 uint16 = 10
    v2[808] = -805
    fmt.Println(v20)
    _ = len(v2)
    v20 = 1
    v1 -= (^(v1))
    if ((v1 % ((4294967295) | 1)) != 489) {
        if false {
            switch "foo bar" {
            case "":
                switch "foo bar" {
                case "":
                    v21 := GenS5{F0: -9223372036854775808, F1: 997, F2: v20}
                    _ = v21
                case "\n\t":
                    var v22 uint32 = 2147483648
                    _ = v22
                case "foo bar":
                    var v23 float64 = float64(1.5)
                    _ = v23
                }
            case "foo bar":
                v1 += 0
            case "\n\t":
                fmt.Println(v1, v1, v3)
            default:
                if (((v3 ^ (v3 | 0)) <= (v3 / (((v3 + v3)) | 1))) || (v3 == (-(v3)))) {
                    v20 = v20
                    v20 = v20
                    if true {
                        var v24 uint32 = 2147483647
                        _ = v24
                    }
                }
            }
            var v25 float64 = float64(float32(3.14159))
            v3 += (v3 >> ((v1) & 63))
            _ = v25
        } else {
            if false {
                var v26 uint64 = 123
                if true {
                    v1 = (v1 % (((v1 / ((v1) | 1))) | 1))
                    fmt.Println(v26, v3, v26)
                } else {
                    _ = len(v2)
                    v20 -= 32767
                }
                fmt.Println(v20)
                _ = v26
            } else {
                var v27 float32 = 1e10
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 2; fi28++ {
            v29 := GenS5{F0: (fi28 << ((1) & 63)), F1: 1, F2: (v20 * v20)}
            for fi30 := 0; fi30 < 18; fi30++ {
                fmt.Println(v1)
                _ = fi30
            }
            _ = fi28
            _ = v29
        }
        if false {
            v3 += (^(v3))
        }
    } else {
        fmt.Println(v1)
        var v31 float32 = 3.14159
        _ = v31
    }
    if (!false) {
        fmt.Println(v20, v1)
    } else {
        v32 := GenS3{F0: 42, F1: 2, F2: -714}
        v1 = (v1 * 1)
        for fi33 := 0; fi33 < 16; fi33++ {
            if false {
                if (!false) {
                    var v34 int = 32768
                    v20 -= 42
                    fmt.Println(v20, v3, v3)
                    _ = v34
                }
            } else {
                v35 := GenS4{F0: false}
                v18 = append(v18, (fi33 / (((-(fi33))) | 1)))
                _ = v35
            }
            _ = fi33
        }
        _ = v32
    }
    v20 -= (^(v20))
    switch ("hello" + ("foo bar" + "\n\t")) {
    case "":
        v3 += 42
    case "a":
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}
