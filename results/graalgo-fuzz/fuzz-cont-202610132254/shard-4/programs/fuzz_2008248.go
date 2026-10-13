package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 bool; F2 uint }
type GenS7 struct { F0 bool }
func genFn0(p0 bool, p1 uint, p2 string) int16 {
    v4 := GenS1{F0: (p1 - (p1 & (p1 / (((p1 ^ 65535)) | 1)))), F1: -1}
    switch 128 {
    case 0:
        p1 += p1
    }
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 uint8 = 161
        _ = fi5
        _ = v6
    }
    p2 = ("foo bar" + (p2 + (p2 + p2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}
func genFn1(p0 uint16) int32 {
    var v7 int = 1
    if (!true) {
        if (42 < 941) {
            var v8 int64 = 1
            for fi9 := 0; fi9 < 8; fi9++ {
                switch ("\n\t" + "\n\t") {
                case "hello":
                    if (false && false) {
                        var v10 uint = 65535
                        v11 := GenS2{F0: -32768, F1: (v8 ^ v8)}
                        _ = v10
                        _ = v11
                    }
                }
                _ = fi9
            }
            v12 := GenS1{F0: 852, F1: (v8 ^ v8)}
            _ = v12
            _ = v8
        }
    } else {
        v7 = 130
        v7 = -2147483648
    }
    var v13 float32 = (3.14159 * (float32(float64(3.14159)) + 0.1))
    for fi14 := 0; fi14 < 5; fi14++ {
        p0 -= p0
        v15 := []int{v7}
        for fi16 := 0; fi16 < 4; fi16++ {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            if false {
                if len(v15) > 0 {
                    v15[0] = 42
                }
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    var v17 int8 = -1
    _ = p0
    _ = v13
    _ = v17
    _ = v7
    return 252
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    switch "foo bar" {
    case "foo bar":
        for fi18 := 0; fi18 < 20; fi18++ {
            v19 := GenS3{F0: 215, F1: 42}
            fmt.Println(fi18, fi18)
            _ = fi18
            _ = v19
        }
    case "":
        var v20 float64 = 3.14159
        _ = v20
    }
    var v21 float32 = 1e10
    v21 = 1.5
    if (0 == 87) {
        var v22 float64 = 1e10
        _ = v22
    } else {
        if (true || true) {
            var v23 int64 = 9223372036854775807
            var v24 int32 = -1
            v25 := []int{32767, 0, 9223372036854775807, 1}
            _ = v23
            _ = v24
            _ = v25
        }
    }
    switch 2147483647 {
    case 127:
        for fi26 := 0; fi26 < 1; fi26++ {
            v27 := GenS3{F0: 42, F1: 18446744073709551615}
            _ = fi26
            _ = v27
        }
    default:
        if (!true) {
            v28 := GenS4{F0: 51}
            if (-57 == -105) {
                for fi29 := 0; fi29 < 12; fi29++ {
                    v21 -= v21
                    v30 := GenS5{F0: 1}
                    _ = fi29
                    _ = v30
                }
                var v31 float64 = float64(v21)
                var v32 uint8 = 1
                _ = v31
                _ = v32
            }
            _ = v28
        } else {
            if (1 != -1) {
                if false {
                    for fi33 := 0; fi33 < 4; fi33++ {
                        v34 := GenS6{F0: 32767, F1: true, F2: 765}
                        var v35 uint32 = 2
                        if false {
                            v1 += v1
                            v1 = -2147483648
                        } else {
                            v36 := []int{(fi33 * (fi33 ^ fi33))}
                            _ = v36
                        }
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    v37 := GenS5{F0: 42}
                    _ = v37
                }
                for fi38 := 0; fi38 < 4; fi38++ {
                    fmt.Println(v1)
                    v1 = v1
                    v39 := GenS1{F0: 48, F1: 1}
                    _ = fi38
                    _ = v39
                }
                v21 -= v21
            }
            v40 := GenS1{F0: 18446744073709551615, F1: 231}
            for fi41 := 0; fi41 < 9; fi41++ {
                v42 := GenS7{F0: false}
                _ = fi41
                _ = v42
            }
            _ = v40
        }
    }
    var v43 uint64 = 4294967295
    v21 += ((v21 * -2.25) * v21)
    var v44 int16 = 897
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v43
    _ = v44
}
