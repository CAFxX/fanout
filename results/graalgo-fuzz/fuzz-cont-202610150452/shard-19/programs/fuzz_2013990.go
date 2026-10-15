package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint; F1 string; F2 uint }
type GenS4 struct { F0 uint64; F1 int; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int32; F2 bool }
type GenS6 struct { F0 int8; F1 uint32; F2 int32 }

func main() {
    var v1 uint = 0
    var v2 uint8 = 127
    v3 := []uint32{}
    var v4 int64 = -9223372036854775808
    var v5 float32 = 0.1
    if len(v3) > 0 {
        v3[0] = 779
    }
    var v6 int64 = (-(v4))
    fmt.Println(v6, v5, v1)
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            if ("hello" >= "a") {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v5 = (float32(1e10) - (v5 - v5))
            }
            v5 -= 0.0
            var v9 float64 = 1.5
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 2, F1: 4294967295, F2: 42}
        var v11 string = "\n\t"
        _ = fi7
        _ = v10
        _ = v11
    }
    var v12 uint64 = 2
    _ = len(v3)
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 bool = (!(9223372036854775807 >= v4))
        _ = fi13
        _ = v14
    }
    if false {
        v15 := GenS2{F0: 42, F1: -71}
        v16 := GenS3{F0: (v1 | v1), F1: ("hello" + ("foo bar" + "")), F2: (v1 >> ((v1) & 63))}
        for fi17 := 0; fi17 < 10; fi17++ {
            fmt.Println(v12, fi17, v1)
            _ = fi17
        }
        _ = v15
        _ = v16
    } else {
        var v18 float32 = (v5 / v5)
        v5 -= (v18 + -2.25)
        _ = v18
    }
    fmt.Println(v1, v1)
    v3 = append(v3, 1)
    v6 += (v6 & (v4 ^ int64(v12)))
    var v19 uint = v1
    if (true || false) {
        v20 := GenS3{F0: uint(v2), F1: "foo bar", F2: (v1 / ((v1) | 1))}
        if false {
            v3 = append(v3, 0)
            v6 = 531
            for fi21 := 0; fi21 < 3; fi21++ {
                var v22 uint32 = 65535
                fmt.Println(v12)
                _ = fi21
                _ = v22
            }
        } else {
            fmt.Println(v5, v4)
        }
        _ = v20
    } else {
        v23 := GenS4{F0: 796, F1: 42, F2: 65535}
        if true {
            v6 = -395
        } else {
            var v24 uint = (v19 | (v19 + 1))
            var v25 uint = (v1 ^ (v24 << ((650) & 63)))
            var v26 uint16 = 560
            _ = v24
            _ = v25
            _ = v26
        }
        if (false || (0 > 585)) {
            switch ("\n\t" + "a") {
            case "a":
                if ("a" != "foo bar") {
                    var v27 string = "\n\t"
                    _ = v27
                }
            case "":
                v28 := GenS3{F0: v19, F1: "\n\t", F2: v1}
                _ = v28
            default:
                for fi29 := 0; fi29 < 1; fi29++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi29
                }
            }
        }
        _ = v23
    }
    var v30 int64 = (v4 ^ v6)
    for fi31 := 0; fi31 < 12; fi31++ {
        v32 := GenS5{F0: 2, F1: int32(v12), F2: (v2 >= v2)}
        if ((v4 < 46) && (true && false)) {
            v33 := GenS6{F0: 42, F1: 1, F2: -1}
            _ = v33
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi31
        _ = v32
    }
    v1 = 57
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
