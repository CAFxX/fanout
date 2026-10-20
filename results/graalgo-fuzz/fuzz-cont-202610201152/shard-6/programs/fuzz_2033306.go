package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int; F1 int; F2 bool }
func genFn0() uint8 {
    var v4 float64 = 1e10
    switch 0 {
    case -2147483648:
        fmt.Println(v4)
    case -879:
        v4 += 0.1
    default:
        v5 := GenS1{F0: 2147483647, F1: 0}
        _ = v5
    }
    v6 := GenS2{F0: 2147483648, F1: 0}
    var v7 uint8 = 255
    _ = v4
    _ = v6
    _ = v7
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    v8 := []int{}
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int16 = 1
            switch fi9 {
            case -411:
                v11 += v11
            case 32767:
                v12 := GenS2{F0: 65535, F1: -126}
                _ = v12
            case 42:
                var v13 uint32 = 0
                _ = v13
            }
            fmt.Println(v11)
            fmt.Println(fi10)
            var v14 int16 = -32768
            v15 := GenS3{F0: 127, F1: 2}
            if (0.1 > 0.0) {
                switch "hello" {
                case "a":
                    if (true || (599 <= fi10)) {
                        v11 -= v11
                        var v16 float64 = (-2.25 + (-2.25 * -2.25))
                        if true {
                            fmt.Println(v11, v16)
                        }
                        _ = v16
                    }
                case "\n\t":
                    if false {
                        var v17 uint64 = 42
                        fmt.Println(fi10, fi9)
                        v14 -= (v11 >> ((4294967295) & 63))
                        if (((("" + "a") + "\n\t") + "foo bar") >= "a") {
                            fmt.Println(fi10)
                            v14 = v11
                        }
                        var v18 float32 = 0.0
                        _ = v17
                        _ = v18
                    } else {
                        v14 = (v11 * v14)
                        v11 -= 0
                    }
                case "":
                    switch (fi9 ^ fi10) {
                    case -2147483648:
                        v11 = (v11 >> ((65535) & 63))
                    default:
                        if ((649 <= 65535) || (!((-1 == 1) || (("" + "a") <= "\n\t")))) {
                            v11 -= v14
                            fmt.Println(fi9)
                        }
                        v19 := GenS3{F0: 255, F1: 18446744073709551615}
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        _ = v19
                    }
                }
                var v20 int8 = 127
                _ = v20
            }
            v14 = (v14 | v11)
            _ = fi10
            _ = v11
            _ = v14
            _ = v15
        }
        if false {
            for fi21 := 0; fi21 < 5; fi21++ {
                v22 := GenS2{F0: 0, F1: 127}
                var v23 uint32 = 0
                _ = fi21
                _ = v22
                _ = v23
            }
            v24 := GenS1{F0: 42, F1: 551}
            _ = v24
        } else {
            for fi25 := 0; fi25 < 10; fi25++ {
                v26 := GenS1{F0: 4294967295, F1: 9223372036854775807}
                _ = fi25
                _ = v26
            }
            var v27 float64 = 0.0
            v27 = (0.1 - v27)
            _ = v27
        }
        _ = fi9
    }
    var v28 int16 = 42
    var v29 uint32 = 148
    v30 := GenS3{F0: 153, F1: 1}
    if true {
        for fi31 := 0; fi31 < 11; fi31++ {
            v32 := GenS2{F0: 42, F1: 100}
            _ = fi31
            _ = v32
        }
        v33 := GenS4{F0: 4294967295, F1: 2147483647, F2: ((0.1 * ((((((3.14159 + 3.14159) / ((0.1 + 0.0) + float64(float32(0.1)))) + 1e10) + 0.0) / (3.14159 * (0.1 / 1e10))) + float64(((float32(0.1) - 1e10) / 0.1)))) > 0.1)}
        v34 := GenS4{F0: 9223372036854775807, F1: 42, F2: true}
        _ = v33
        _ = v34
    }
    var v35 uint64 = 0
    var v36 uint32 = v29
    for fi37 := 0; fi37 < 12; fi37++ {
        v28 += v28
        _ = fi37
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        if (128 < 0) {
            v28 += (v28 & 1000)
            for fi39 := 0; fi39 < 9; fi39++ {
                v36 = v29
                _ = fi39
            }
            v40 := GenS3{F0: 0, F1: (^(v35))}
            _ = v40
        } else {
            fmt.Println(v29, v35, fi38)
        }
        _ = fi38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v8
}
