package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 bool }
type GenS2 struct { F0 string; F1 int32; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint; F2 int32 }
type GenS4 struct { F0 uint64; F1 int8; F2 string }
func genFn0(p0 int8, p1 int32, p2 int16) int32 {
    p0 += (p0 << ((991) & 63))
    for fi4 := 0; fi4 < 5; fi4++ {
        p1 -= (p1 * 2147483647)
        _ = fi4
    }
    var v5 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}
func genFn1(p0 uint16, p1 uint) int8 {
    var v6 int32 = 1
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint8 = 255
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 int32 = v6
            p0 = p0
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    p0 -= p0
    switch "a" {
    case "\n\t":
        v13 := GenS1{F0: 2147483647, F1: (^(v6)), F2: true}
        _ = v13
    case "foo bar":
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 int32 = v6
            v6 -= genFn0(100, v15, 42)
            var v16 int = (-(fi14))
            _ = fi14
            _ = v15
            _ = v16
        }
    case "a":
        fmt.Println(v6, p1, v6)
    }
    _ = p0
    _ = p1
    _ = v6
    return -1
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int8 = -96
    var v3 int = 128
    v17 := GenS2{F0: "", F1: 1, F2: 288}
    v3 = v3
    var v18 int32 = 0
    var v19 int16 = 0
    v20 := []int{v3, 9223372036854775807}
    for fi21 := 0; fi21 < 20; fi21++ {
        switch "a" {
        case "hello":
            v19 -= (^(v19))
        case "foo bar":
            v2 = (^(v2))
        case "\n\t":
            if true {
                for fi22 := 0; fi22 < 4; fi22++ {
                    v3 = (fi22 % (((-(v3))) | 1))
                    _ = fi22
                }
                switch (v3 ^ v3) {
                case 85:
                    for fi23 := 0; fi23 < 11; fi23++ {
                        var v24 uint = 638
                        var v25 float64 = 1.5
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                case 553:
                    for fi26 := 0; fi26 < 2; fi26++ {
                        var v27 uint16 = 587
                        _ = fi26
                        _ = v27
                    }
                case 127:
                    if (v19 != (v19 + (-(v19)))) {
                        var v28 bool = (42 <= 653)
                        _ = v28
                    }
                default:
                    v29 := GenS3{F0: (-(v1)), F1: 65535, F2: (v18 ^ (v18 + -179))}
                    _ = v29
                }
                var v30 uint32 = 24
                _ = v30
            }
        }
        _ = fi21
    }
    if (true && (!true)) {
        v20 = append(v20, -327)
        if true {
            var v31 int = (v3 << ((65535) & 63))
            if len(v20) > 0 {
                v20[0] = v31
            }
            fmt.Println(v19, v3)
            _ = v31
        } else {
            v3 -= (v3 % (((v3 + -647)) | 1))
        }
    } else {
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
    }
    switch 32768 {
    case 42:
        if (v19 >= v19) {
            if false {
                v32 := GenS2{F0: "hello", F1: (v18 * (v18 << ((821) & 63))), F2: 1}
                _ = v32
            } else {
                v1 = -1
            }
            fmt.Println(v1)
            v33 := GenS4{F0: 2, F1: (v2 & v2), F2: ("" + "foo bar")}
            _ = v33
        }
    case -652:
        v1 = -9223372036854775808
    }
    fmt.Println(v3, v18, v19)
    v34 := GenS4{F0: 392, F1: v2, F2: "a"}
    v3 -= (v3 + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
}
