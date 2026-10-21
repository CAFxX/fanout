package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint; F1 int8; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 uint32 }
func genFn0() int64 {
    var v4 int16 = 1000
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 uint32 = 744
        v7 := []int{fi5, fi5}
        var v8 uint16 = 391
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    _ = v4
    return -9223372036854775808
}
func genFn1(p0 uint) int64 {
    v9 := GenS1{F0: 834}
    fmt.Println(p0)
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(fi10)
        v11 := GenS2{F0: 42, F1: 0}
        p0 = p0
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v9
    return 42
}

func main() {
    v1 := []int{}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    var v12 int32 = -1
    var v13 uint = 11
    var v14 uint64 = 2
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS1{F0: 984}
        _ = fi15
        _ = v16
    }
    if true {
        delete(v3, 2147483647)
        switch "\n\t" {
        case "":
            v14 = v14
        case "a":
            for fi17 := 0; fi17 < 4; fi17++ {
                for fi18 := 0; fi18 < 14; fi18++ {
                    switch (("a" + "") + "\n\t") {
                    case "a":
                        if (!false) {
                            v14 -= (v14 / ((v14) | 1))
                        }
                    case "foo bar":
                        var v19 uint16 = 2
                        _ = v19
                    default:
                        v20 := GenS3{F0: 1, F1: 42}
                        _ = v20
                    }
                    _ = fi18
                }
                var v21 int32 = (v12 / (((^(v12))) | 1))
                switch "a" {
                case "a":
                    v13 = v13
                case "hello":
                    v14 -= 0
                }
                _ = fi17
                _ = v21
            }
        default:
            v13 = (v13 ^ 451)
        }
        switch 42 {
        case 789:
            fmt.Println(v13, v12, v13)
        case 4294967295:
            if (-9223372036854775808 > -19) {
                var v22 uint16 = 65535
                _ = v22
            }
        default:
            v14 -= v14
        }
    } else {
        if ((0 < uint32(v13)) && false) {
            var v23 float64 = (-2.25 - 0.0)
            _ = v23
        }
        switch -248 {
        case 315:
            var v24 int = 4294967295
            _ = v24
        }
    }
    v12 += (v12 * 695)
    v25 := GenS4{F0: (v13 ^ v13), F1: 42, F2: ("\n\t" + ("\n\t" + ((("a" + (("foo bar" + "hello") + "")) + "\n\t") + "")))}
    for fi26 := 0; fi26 < 13; fi26++ {
        if (2 == 60) {
            v12 = (^(v12))
        }
        _ = fi26
    }
    var v27 uint = v13
    var v28 int64 = genFn1(v27)
    if false {
        for fi29 := 0; fi29 < 11; fi29++ {
            v3[42] = "hello"
            v30 := GenS5{F0: -1}
            var v31 uint16 = 0
            _ = fi29
            _ = v30
            _ = v31
        }
        v3[uint32(v14)] = "a"
    }
    v32 := GenS4{F0: (v27 ^ 65535), F1: 42, F2: "\n\t"}
    v33 := GenS6{F0: "foo bar", F1: 2147483647}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
}
