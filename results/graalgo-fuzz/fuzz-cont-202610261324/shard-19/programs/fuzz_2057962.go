package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int32 }
type GenS2 struct { F0 uint; F1 bool; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 int16; F1 int32; F2 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = (1.5 - 1e10)
    var v3 int64 = 0
    var v4 uint64 = 938
    v3 = -276
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := []int{fi5}
        _ = len(v6)
        _ = len(v1)
        _ = fi5
        _ = v6
    }
    if ("\n\t" < ("a" + "a")) {
        switch 128 {
        case 0:
            if (false && false) {
                v7 := GenS1{F0: ("hello" + ""), F1: -9223372036854775808, F2: 1}
                _ = v7
            }
        case 836:
            v8 := GenS2{F0: 280, F1: ((4294967295 <= 733) || ((1.5 + -2.25) >= -2.25)), F2: ("\n\t" + "a")}
            _ = v8
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: "", F1: (v3 >> ((2) & 63)), F2: -2147483648}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 7; fi11++ {
            v4 = (^(v4))
            _ = fi11
        }
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 float64 = -2.25
        v14 := GenS3{F0: v4}
        v15 := GenS4{F0: 52, F1: -1}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    v16 := GenS5{F0: ("" < "a"), F1: -128}
    switch 277 {
    case 9223372036854775807:
        if false {
            v4 = v4
            v3 = (^(v3))
        } else {
            if true {
                v2 += (v2 - -2.25)
                switch "foo bar" {
                case "foo bar":
                    v4 += 627
                case "a":
                    var v17 uint32 = 65535
                    _ = v17
                default:
                    var v18 int16 = -1
                    _ = v18
                }
            } else {
                var v19 bool = false
                var v20 uint8 = 0
                v21 := GenS4{F0: -128, F1: 32767}
                _ = v19
                _ = v20
                _ = v21
            }
            v22 := v1[(("\n\t" + "") + "")]
            if ((!(!v22)) || (59 < 42)) {
                switch 128 {
                case 4294967295:
                    var v23 string = "a"
                    _ = v23
                }
                var v24 uint = 1
                v25 := v1[(("\n\t" + "hello") + ("hello" + "a"))]
                _ = v24
                _ = v25
            } else {
                v26 := []int{-424, 32768, 32768, 1}
                _ = v26
            }
            _ = v22
        }
    case 4294967295:
        var v27 uint = 220
        _ = v27
    case -319:
        var v28 bool = (!false)
        _ = v28
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        if false {
            var v30 int16 = 339
            v3 = (v3 % ((v3) | 1))
            _ = v30
        }
        fmt.Println(v3)
        _ = fi29
    }
    v31 := GenS6{F0: "\n\t", F1: 2147483647}
    v4 = (v4 >> ((65535) & 63))
    if true {
        var v32 int8 = 100
        v33 := GenS6{F0: "\n\t", F1: 32768}
        _ = v32
        _ = v33
    } else {
        v4 = v4
        for fi34 := 0; fi34 < 16; fi34++ {
            var v35 float32 = (v2 - (float32(0.1) - 3.14159))
            v36 := GenS7{F0: 0, F1: -2147483648, F2: false}
            v35 = v2
            _ = fi34
            _ = v35
            _ = v36
        }
        v2 = v2
    }
    if (!false) {
        switch "hello" {
        case "foo bar":
            v37 := []int{-1, 2147483647, 748}
            _ = v37
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
