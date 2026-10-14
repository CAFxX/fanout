package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 uint; F1 int64; F2 int64 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 uint; F1 int64 }

func main() {
    var v1 int = 117
    var v2 int = v1
    v3 := make(map[string]bool)
    var v4 int8 = -108
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint16 = 1
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    var v8 uint32 = 42
    v9 := GenS1{F0: 413, F1: int8(v8), F2: 42}
    v10 := GenS1{F0: 65535, F1: 127, F2: 65535}
    fmt.Println(v2, v1)
    var v11 uint = 2
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS1{F0: v11, F1: (v4 + v4), F2: (v11 | (v11 << ((362) & 63)))}
        v1 += (v2 % ((v2) | 1))
        v14 := GenS2{F0: (v11 >> ((65535) & 63)), F1: 364, F2: 1}
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int64 = 0
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS3{F0: 848, F1: v11}
        _ = fi16
        _ = v17
    }
    v18 := v3["foo bar"]
    for fi19 := 0; fi19 < 4; fi19++ {
        if true {
            v1 = fi19
        }
        if v18 {
            var v20 int32 = 2147483647
            fmt.Println(v1)
            for fi21 := 0; fi21 < 20; fi21++ {
                switch ("a" + "") {
                case "foo bar":
                    v22 := GenS3{F0: 0, F1: 2}
                    _ = v22
                case "hello":
                    v23 := GenS2{F0: v11, F1: 1, F2: (v15 / (((v15 ^ -1)) | 1))}
                    _ = v23
                case "a":
                    var v24 float64 = (1e10 + -2.25)
                    _ = v24
                }
                v25 := v3["a"]
                v20 -= v20
                if (v18 && (!false)) {
                    switch "foo bar" {
                    case "hello":
                        var v26 float64 = float64(((3.14159 + ((0.1 - (-2.25 / -2.25)) + (1.5 / 1.0))) - 0.1))
                        _ = v26
                    case "":
                        v27 := []int{-9223372036854775808, -989, (v2 + 42), (v1 + 0)}
                        _ = v27
                    default:
                        v28 := GenS2{F0: (v11 + v11), F1: 623, F2: (v15 & (v15 / ((v15) | 1)))}
                        _ = v28
                    }
                } else {
                    var v29 float64 = (((-2.25 / 1.5) / (1e10 - -2.25)) / -2.25)
                    if false {
                        fmt.Println(v4, fi21, v2)
                        v30 := []int{(^(v1)), (^(v1)), v2, (v2 & fi19)}
                        _ = v30
                    } else {
                        v31 := GenS1{F0: 475, F1: 0, F2: (v11 ^ (v11 ^ (v11 % ((v11) | 1))))}
                        _ = v31
                    }
                    _ = v29
                }
                if v25 {
                    var v32 bool = (!(true && ((1e10 * float32((0.0 * 1.5))) >= -2.25)))
                    _ = v32
                } else {
                    v2 = 4294967295
                    v11 += (v11 << ((42) & 63))
                    fmt.Println(fi21, fi19)
                }
                _ = fi21
                _ = v25
            }
            _ = v20
        } else {
            for fi33 := 0; fi33 < 19; fi33++ {
                if v18 {
                    if v18 {
                        v34 := GenS4{F0: 127, F1: (((("\n\t" + "\n\t") + "a") + "") + (("hello" + "hello") + ""))}
                        _ = v34
                    }
                }
                v35 := []int{fi19}
                _ = fi33
                _ = v35
            }
            for fi36 := 0; fi36 < 17; fi36++ {
                var v37 uint16 = 72
                v18 = (!v18)
                _ = fi36
                _ = v37
            }
            switch (v2 - v1) {
            case 1:
                fmt.Println(v2, v4)
            case -969:
                v38 := GenS5{F0: (v11 / ((0) | 1)), F1: -682}
                _ = v38
            }
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
