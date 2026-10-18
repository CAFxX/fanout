package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int8; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{("foo bar" + ((("a" + ("" + "hello")) + "a") + "\n\t")), "foo bar", ("foo bar" + ("" + "\n\t")), ("\n\t" + "a")}
    var v3 int16 = 32767
    var v4 int8 = -1
    v5 := []uint32{}
    v6 := GenS1{F0: 1}
    var v7 string = ""
    switch "hello" {
    case "a":
        _ = len(v1)
    case "":
        if true {
            for fi8 := 0; fi8 < 3; fi8++ {
                for fi9 := 0; fi9 < 5; fi9++ {
                    delete(v1, 0)
                    _ = fi9
                }
                var v10 uint8 = 42
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi8
                _ = v10
            }
            fmt.Println(v7, v3, v4)
        }
    default:
        v4 += v4
    }
    if len(v2) > 0 {
        v2[0] = v7
    }
    v4 = v4
    if ((((0.1 / 1.0) / 1.0) + 0.0) <= -2.25) {
        var v11 int16 = v3
        _ = v11
    } else {
        fmt.Println(v7)
    }
    switch 127 {
    case 127:
        if true {
            var v12 int64 = 0
            if true {
                var v13 int8 = (-(v4))
                var v14 uint32 = 0
                _ = v13
                _ = v14
            }
            _ = v12
        }
    case 0:
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        fmt.Println(v4, v7, v7)
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := GenS1{F0: 0}
            if (!(4294967295 == 947)) {
                v4 -= (-(v4))
                v18 := GenS1{F0: 0}
                var v19 string = "foo bar"
                _ = v18
                _ = v19
            }
            if true {
                v3 = v3
                var v20 float64 = 0.1
                _ = v20
            } else {
                v2 = append(v2, v7)
            }
            _ = fi16
            _ = v17
        }
        fmt.Println(v4)
        _ = fi15
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := GenS1{F0: 1}
        _ = fi21
        _ = v22
    }
    v3 -= 42
    var v23 int32 = -2147483648
    fmt.Println(v4)
    var v24 int8 = -4
    if true {
        fmt.Println(v4, v7)
        v25 := GenS2{F0: (v3 % (((v3 % ((v3) | 1))) | 1))}
        if (int32(v3) > -855) {
            var v26 uint = 253
            _ = v26
        }
        _ = v25
    } else {
        fmt.Println(v4, v3)
        fmt.Println(v7, v4, v24)
        v27 := GenS3{F0: -1, F1: v4, F2: 280}
        _ = v27
    }
    var v28 uint = 50
    for fi29 := 0; fi29 < 12; fi29++ {
        fmt.Println(v23)
        var v30 string = (v7 + (v7 + v7))
        v5 = append(v5, 65535)
        _ = fi29
        _ = v30
    }
    v31 := v1[42]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
