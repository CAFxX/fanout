package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint; F1 uint16 }

func main() {
    var v1 uint = 583
    v2 := []uint32{4294967295}
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    var v5 int = 32768
    var v6 uint = 254
    var v7 uint64 = 222
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: 42, F1: -257}
        fmt.Println(v1, v7, v7)
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 float32 = (float32(1.5) / 1.0)
            v1 = v6
            if false {
                v12 := GenS1{F0: 0, F1: -1}
                v4[245] = (("a" + "foo bar") + ((("a" + "hello") + "") + "a"))
                v7 += v7
                var v13 uint8 = 128
                fmt.Println(v6, fi10, v5)
                var v14 string = "foo bar"
                v15 := GenS2{F0: (v14 + ((v14 + v14) + v14))}
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            } else {
                v5 -= (^(fi8))
            }
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v6 -= 1
    for fi16 := 0; fi16 < 19; fi16++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v1)
        if true {
            v17 := v3[("hello" + "\n\t")]
            if (((false || true) || (42 > -9223372036854775808)) || false) {
                for fi18 := 0; fi18 < 8; fi18++ {
                    fmt.Println(v1, v7)
                    var v19 uint = 779
                    if (!true) {
                        if (false || (32767 <= 42)) {
                            v7 += 1
                            v20 := GenS2{F0: (((("" + (("" + ("" + ("hello" + ("hello" + "hello")))) + ("" + "foo bar"))) + (("hello" + "\n\t") + (("\n\t" + "foo bar") + ""))) + "foo bar") + "hello")}
                            _ = v20
                        }
                    }
                    _ = fi18
                    _ = v19
                }
                for fi21 := 0; fi21 < 12; fi21++ {
                    fmt.Println(v17, v17, fi21)
                    var v22 int = (fi16 * fi21)
                    _ = fi21
                    _ = v22
                }
            }
            _ = len(v2)
            _ = v17
        } else {
            v7 -= v7
            var v23 float64 = 0.0
            v6 += uint(v7)
            _ = v23
        }
        _ = fi16
    }
    v24 := GenS1{F0: -1, F1: 9223372036854775807}
    v1 -= (v1 ^ (v1 | v6))
    for fi25 := 0; fi25 < 14; fi25++ {
        if (v5 != fi25) {
            var v26 int8 = -107
            _ = v26
        } else {
            for fi27 := 0; fi27 < 17; fi27++ {
                _ = len(v2)
                _ = fi27
            }
            v28 := GenS2{F0: ""}
            _ = v28
        }
        v29 := GenS1{F0: -1, F1: -1}
        _ = fi25
        _ = v29
    }
    v4[482] = ("a" + ("a" + "hello"))
    for fi30 := 0; fi30 < 7; fi30++ {
        if false {
            v6 = v6
        }
        switch "foo bar" {
        case "hello":
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v1, v5)
            case "":
                if (1 <= 42) {
                    v31 := GenS3{F0: v6, F1: 42}
                    _ = v31
                } else {
                    v1 += (v6 >> ((v1) & 63))
                    if ((fi30 > fi30) || (4294967295 < 555)) {
                        var v32 uint32 = 1
                        var v33 uint = v1
                        _ = v32
                        _ = v33
                    } else {
                        var v34 int16 = 0
                        var v35 int64 = 866
                        _ = v34
                        _ = v35
                    }
                }
            }
        default:
            switch v5 {
            case -2147483648:
                v1 = (v1 + 556)
            case 32768:
                var v36 uint64 = 1
                _ = v36
            }
        }
        _ = fi30
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        if ((!(!(217 < 0))) || true) {
            if true {
                v38 := GenS4{F0: 294, F1: 65535}
                var v39 float64 = float64((0.0 - float32(((0.0 - 0.1) + 0.1))))
                _ = v38
                _ = v39
            } else {
                v6 = (v1 % ((v6) | 1))
            }
        } else {
            var v40 int16 = 32767
            fmt.Println(v7, v6)
            _ = v40
        }
        var v41 bool = true
        _ = fi37
        _ = v41
    }
    if false {
        var v42 string = ("hello" + "\n\t")
        _ = v42
    }
    if (22 == 490) {
        for fi43 := 0; fi43 < 8; fi43++ {
            var v44 uint = v6
            for fi45 := 0; fi45 < 12; fi45++ {
                v7 += 669
                _ = fi45
            }
            var v46 int32 = 1
            _ = fi43
            _ = v44
            _ = v46
        }
    } else {
        v2 = append(v2, 65535)
        var v47 int32 = 1
        var v48 int16 = 32767
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
