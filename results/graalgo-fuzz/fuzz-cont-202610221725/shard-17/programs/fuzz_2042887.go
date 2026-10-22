package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 100
    var v3 string = "\n\t"
    v4 := make(map[uint32]int)
    v5 := []int{32767, 42, -9223372036854775808, 482}
    var v6 int = -150
    v6 = v6
    var v7 int64 = -637
    if false {
        var v8 string = ("" + (v3 + "a"))
        _ = v8
    }
    if (!false) {
        v6 = -1
        for fi9 := 0; fi9 < 15; fi9++ {
            v3 = v3
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := GenS1{F0: 255, F1: v7, F2: (-(fi10))}
                fmt.Println(fi9)
                _ = fi10
                _ = v11
            }
            _ = len(v5)
            _ = fi9
        }
    }
    v12 := GenS1{F0: 2, F1: (-(v7)), F2: v6}
    var v13 uint = 57
    fmt.Println(v7, v3, v2)
    v7 = -939
    switch v3 {
    case "a":
        var v14 string = v3
        _ = v14
    case "":
        var v15 int32 = 1
        _ = v15
    default:
        v16 := GenS2{F0: ((42 > 1000) || false)}
        _ = v16
    }
    var v17 string = (((v3 + ("a" + v3)) + (((v3 + "foo bar") + v3) + v3)) + (("a" + "a") + v3))
    switch (v6 % (((v6 * v6)) | 1)) {
    case 128:
        switch (-(v6)) {
        case -206:
            switch (v6 % ((v6) | 1)) {
            case 377:
                if (3.14159 != (3.14159 * ((float32((float64(3.14159) * 1e10)) / 1e10) - 3.14159))) {
                    for fi18 := 0; fi18 < 6; fi18++ {
                        v19 := GenS1{F0: 32767, F1: (v7 ^ (v7 * (-(v7)))), F2: v6}
                        _ = fi18
                        _ = v19
                    }
                } else {
                    v6 = (v6 + (v6 | v6))
                }
            case -2147483648:
                var v20 uint = 238
                _ = v20
            case -798:
                v6 -= v6
            }
        case 9223372036854775807:
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 uint32 = 252
                _ = fi21
                _ = v22
            }
        default:
            v3 = ("" + v17)
        }
    case 32767:
        var v23 float32 = 1e10
        _ = v23
    default:
        switch -4 {
        case 1:
            if false {
                switch v6 {
                case 720:
                    v3 = v3
                default:
                    for fi24 := 0; fi24 < 6; fi24++ {
                        v6 = fi24
                        var v25 uint64 = 42
                        v26 := GenS3{F0: 1, F1: v2}
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                }
            }
        default:
            v6 = (v6 + (v6 - v6))
        }
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        if true {
            v2 = -128
        } else {
            var v28 int16 = 943
            fmt.Println(v7)
            _ = v28
        }
        v6 -= fi27
        _ = fi27
    }
    if (!(42 <= 42)) {
        v7 = (v7 & (^(v7)))
        v29 := GenS2{F0: true}
        v30 := GenS3{F0: (-(v2)), F1: (v2 & (v2 / ((v2) | 1)))}
        _ = v29
        _ = v30
    } else {
        switch v6 {
        case 9223372036854775807:
            for fi31 := 0; fi31 < 9; fi31++ {
                v32 := GenS3{F0: (-(v2)), F1: v2}
                v6 -= (fi31 & v6)
                _ = fi31
                _ = v32
            }
        case -1:
            switch v17 {
            case "hello":
                var v33 int16 = -304
                _ = v33
            }
        case -47:
            delete(v4, 1)
        }
        for fi34 := 0; fi34 < 1; fi34++ {
            var v35 int8 = v2
            v36 := GenS4{F0: 65535, F1: 109, F2: -32768}
            for fi37 := 0; fi37 < 19; fi37++ {
                var v38 int = fi34
                v38 += fi34
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi37
                _ = v38
            }
            _ = fi34
            _ = v35
            _ = v36
        }
        v39 := GenS4{F0: 42, F1: 601, F2: -211}
        _ = v39
    }
    v40 := GenS2{F0: (((0.0 + (float64(0.0) + (0.1 * 1.5))) / 1.5) == float64((0.1 + 1.5)))}
    fmt.Println(v3, v2)
    v17 = ""
    v6 = int(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
