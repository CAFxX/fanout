package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint }
type GenS4 struct { F0 int16; F1 string; F2 int64 }
type GenS5 struct { F0 uint32; F1 int16 }
func genFn0(p0 int32) int32 {
    if (true || true) {
        for fi5 := 0; fi5 < 20; fi5++ {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 uint32 = 874
                _ = fi6
                _ = v7
            }
            switch "\n\t" {
            case "a":
                switch -1 {
                case 32767:
                    p0 += -2147483648
                case 127:
                    p0 += p0
                case -156:
                    for fi8 := 0; fi8 < 9; fi8++ {
                        v9 := GenS1{F0: -32768, F1: 0}
                        p0 = p0
                        _ = fi8
                        _ = v9
                    }
                default:
                    v10 := []int{(^(fi5)), fi5, int(p0)}
                    _ = v10
                }
            case "":
                v11 := GenS2{F0: ""}
                _ = v11
            }
            _ = fi5
        }
    }
    var v12 uint = 42
    p0 = (p0 / ((p0) | 1))
    var v13 float32 = -2.25
    _ = p0
    _ = v12
    _ = v13
    return 327
}

func main() {
    v1 := []int{32767}
    var v2 uint8 = 0
    var v3 bool = (-2147483648 > 42)
    var v4 float32 = 0.1
    var v14 float32 = 0.1
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 int8 = -1
        switch 733 {
        case 42:
            for fi17 := 0; fi17 < 1; fi17++ {
                fmt.Println(v16, v16, v14)
                v14 -= (((((v4 * -2.25) - (1.5 + 0.0)) * v4) + v14) / 1e10)
                _ = fi17
            }
        case 127:
            var v18 uint = 2
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        _ = len(v1)
        if (v3 || v3) {
            v14 -= float32(0.0)
            var v20 uint64 = 664
            _ = v20
        }
        _ = fi19
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        fmt.Println(v14, fi21, v4)
        _ = fi21
    }
    v22 := GenS3{F0: v2, F1: 334, F2: 0}
    fmt.Println(v4)
    if v3 {
        var v23 int = 1
        for fi24 := 0; fi24 < 18; fi24++ {
            if (!(!(int64(fi24) < 9223372036854775807))) {
                _ = len(v1)
                for fi25 := 0; fi25 < 4; fi25++ {
                    fmt.Println(v23)
                    var v26 uint8 = v2
                    _ = fi25
                    _ = v26
                }
            }
            _ = fi24
        }
        v27 := GenS1{F0: 1, F1: 2147483647}
        _ = v23
        _ = v27
    }
    if ((127 == 127) || (!false)) {
        fmt.Println(v3, v14, v14)
        v4 = ((float32(1e10) / (v14 + 0.1)) - float32((((1e10 / 1.0) * -2.25) - (0.0 - 0.1))))
    } else {
        for fi28 := 0; fi28 < 3; fi28++ {
            if v3 {
                for fi29 := 0; fi29 < 6; fi29++ {
                    if (735 == 915) {
                        v30 := GenS4{F0: -122, F1: "foo bar", F2: -1}
                        _ = v30
                    }
                    v4 = 1e10
                    v1 = append(v1, (fi29 + fi29))
                    _ = fi29
                }
                switch "a" {
                case "":
                    v31 := GenS4{F0: 32767, F1: "hello", F2: -296}
                    _ = v31
                }
            } else {
                v14 = -2.25
            }
            _ = fi28
        }
        v32 := GenS3{F0: v2, F1: 2, F2: 18446744073709551615}
        v33 := GenS4{F0: 1000, F1: "hello", F2: 819}
        _ = v32
        _ = v33
    }
    v34 := GenS5{F0: 2147483647, F1: -32768}
    var v35 uint8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
