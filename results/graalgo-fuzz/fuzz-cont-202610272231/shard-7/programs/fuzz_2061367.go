package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int64 }
func genFn0() uint32 {
    v7 := GenS1{F0: 803}
    for fi8 := 0; fi8 < 5; fi8++ {
        switch 1 {
        case 32767:
            var v9 string = "\n\t"
            _ = v9
        case -584:
            switch (fi8 << ((1) & 63)) {
            case 4294967295:
                switch "a" {
                case "":
                    if false {
                        v10 := []int{127, -459, (fi8 + 42), (fi8 % (((fi8 & fi8)) | 1))}
                        var v11 float64 = 3.14159
                        switch "hello" {
                        case "":
                            var v12 int = 2147483647
                            _ = v12
                        case "foo bar":
                            if (-121 <= -1) {
                                var v13 float64 = (v11 / (-2.25 + float64(((0.1 + float32(v11)) - (1e10 + 3.14159)))))
                                v14 := GenS2{F0: "", F1: -1}
                                v13 += v13
                                _ = v13
                                _ = v14
                            }
                        }
                        _ = v10
                        _ = v11
                    }
                default:
                    switch fi8 {
                    case 32767:
                        var v15 uint16 = 809
                        _ = v15
                    default:
                        fmt.Println(fi8)
                    }
                }
            case -1:
                var v16 uint16 = 149
                _ = v16
            }
        default:
            var v17 int32 = 0
            _ = v17
        }
        _ = fi8
    }
    _ = v7
    return 0
}
func genFn1(p0 uint, p1 int16, p2 uint) uint8 {
    v18 := GenS2{F0: "foo bar", F1: 218}
    if ("foo bar" >= "hello") {
        var v19 int64 = 0
        _ = v19
    }
    p1 -= (p1 << (((p2 ^ p2)) & 63))
    v20 := []int{int(p1), 664}
    var v21 int16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v20
    _ = v21
    return 2
}

func main() {
    var v1 int = 32767
    v2 := []int{v1, -863, (v1 * 4294967295)}
    var v3 int64 = 42
    v4 := make(map[string]uint32)
    v5 := []string{(("\n\t" + "") + "\n\t"), "foo bar", "foo bar", ("\n\t" + "hello")}
    var v6 uint16 = 42
    v6 = v6
    switch (v1 ^ v1) {
    case 1:
        v3 = 741
    case -9223372036854775808:
        var v22 float64 = (1e10 / 1.5)
        _ = v22
    case 0:
        for fi23 := 0; fi23 < 19; fi23++ {
            for fi24 := 0; fi24 < 19; fi24++ {
                if len(v2) > 0 {
                    v2[0] = fi24
                }
                if (!(true && false)) {
                    v25 := GenS3{F0: ((!false) || true)}
                    v6 = 9
                    _ = v25
                }
                _ = fi24
            }
            _ = fi23
        }
    }
    v26 := GenS2{F0: "foo bar", F1: (v3 / (((v3 | (v3 / ((v3) | 1)))) | 1))}
    var v27 float32 = (float32(1e10) / (1e10 - ((1e10 * 1.5) * 0.0)))
    v1 = (v1 << ((1) & 63))
    v3 -= v3
    var v28 int16 = 1
    v3 = 1
    switch v1 {
    case -422:
        v29 := GenS3{F0: true}
        _ = v29
    default:
        v28 = (v28 >> ((0) & 63))
    }
    var v30 bool = false
    v31 := GenS4{F0: 182, F1: (v3 >> ((42) & 63))}
    var v32 int16 = v28
    if (true || v30) {
        for fi33 := 0; fi33 < 10; fi33++ {
            var v34 float64 = -2.25
            _ = fi33
            _ = v34
        }
        v6 = v6
    } else {
        v1 = (v1 & (v1 >> ((65535) & 63)))
        var v35 uint16 = 2
        if (42 == 42) {
            var v36 int16 = v28
            var v37 int = (-(v1))
            _ = v36
            _ = v37
        } else {
            if ("" < (("\n\t" + "foo bar") + "")) {
                var v38 uint16 = (^(v35))
                _ = v38
            } else {
                switch "\n\t" {
                case "\n\t":
                    if (false && (!v30)) {
                        var v39 int32 = 947
                        v6 = v6
                        v32 = v32
                        _ = v39
                    }
                }
            }
            v6 -= v6
            switch (^(v1)) {
            case -210:
                for fi40 := 0; fi40 < 10; fi40++ {
                    if (false && (!v30)) {
                        switch 928 {
                        case 1:
                            for fi41 := 0; fi41 < 12; fi41++ {
                                var v42 uint8 = 127
                                _ = fi41
                                _ = v42
                            }
                        }
                    }
                    if ((-(v28)) != 32767) {
                        var v43 int8 = 127
                        _ = v43
                    }
                    _ = fi40
                }
            case 9223372036854775807:
                v6 -= v6
            case -779:
                v3 -= v3
            default:
                fmt.Println(v27)
            }
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
