package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 uint16; F2 int32 }
type GenS7 struct { F0 uint32; F1 uint; F2 uint8 }
func genFn0(p0 int, p1 bool, p2 uint32) int32 {
    var v7 float64 = 3.14159
    var v8 uint = 1
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 float64 = -2.25
            var v12 int64 = 1
            v13 := []int{-1, 1, p0, fi9}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi9
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v8 += (v8 | 65535)
        p0 += -9223372036854775808
        _ = fi14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS1{F0: 426, F1: 1, F2: 127}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 bool, p1 int8) string {
    v17 := []int{-9223372036854775808, 9223372036854775807, -1, 42}
    if p0 {
        for fi18 := 0; fi18 < 9; fi18++ {
            p0 = (42 <= 0)
            if len(v17) > 0 {
                v17[0] = -1
            }
            v19 := GenS2{F0: (p1 >> ((255) & 63)), F1: 2147483648}
            _ = fi18
            _ = v19
        }
        var v20 uint16 = 65535
        if (3.14159 < ((((0.0 - float32(-2.25)) / (0.0 + (float32((1e10 + float64(float32(1.5)))) * 0.1))) * 1e10) - 3.14159)) {
            if true {
                v21 := GenS2{F0: -120, F1: 42}
                _ = v21
            }
            fmt.Println(v20, p1, p1)
        }
        _ = v20
    } else {
        var v22 uint64 = 42
        _ = v22
    }
    if (-807 < 0) {
        var v23 bool = p0
        _ = v23
    } else {
        switch ("\n\t" + ("foo bar" + "a")) {
        case "":
            fmt.Println(p0, p0)
        }
        v17 = append(v17, 128)
    }
    v24 := GenS3{F0: p0}
    switch "\n\t" {
    case "hello":
        v25 := GenS4{F0: 223}
        _ = v25
    default:
        p1 = (p1 - p1)
    }
    _ = p0
    _ = p1
    _ = v17
    _ = v24
    return (("hello" + "hello") + ("\n\t" + "hello"))
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]bool)
    v3 := []string{"\n\t", "\n\t"}
    v4 := []string{("\n\t" + "hello"), "foo bar", ""}
    var v5 uint8 = 127
    v6 := make(map[int]string)
    v26 := GenS5{F0: 0}
    v27 := GenS6{F0: -2147483648, F1: 0, F2: 883}
    var v28 float64 = -2.25
    if (42 == 0) {
        _ = len(v4)
        var v29 float64 = (v28 + ((-2.25 * v28) * 3.14159))
        _ = v29
    } else {
        if false {
            delete(v6, -120)
            if ((v28 / ((v28 / 0.1) + v28)) != 1e10) {
                if (65535 >= uint32(v5)) {
                    fmt.Println(v5)
                }
                var v30 int16 = 1000
                v28 = (((v28 - 3.14159) * v28) - ((v28 / v28) * v28))
                _ = v30
            } else {
                v31 := GenS7{F0: 511, F1: 2, F2: 0}
                fmt.Println(v28, v1)
                _ = v31
            }
        } else {
            if true {
                for fi32 := 0; fi32 < 19; fi32++ {
                    switch -144 {
                    case 9223372036854775807:
                        v4 = append(v4, ("hello" + "foo bar"))
                    }
                    v28 = float64((v1 + ((v1 - 1e10) - ((3.14159 + (((3.14159 + 1e10) * float32(v28)) - v1)) * (v1 * 0.1)))))
                    delete(v6, fi32)
                    _ = fi32
                }
                var v33 float32 = 1.5
                var v34 float32 = v33
                _ = v33
                _ = v34
            } else {
                v28 = (-2.25 - 1.5)
                if (4294967295 < 1) {
                    var v35 uint8 = (v5 & (v5 | 255))
                    v1 += v1
                    _ = v35
                } else {
                    v28 = v28
                }
            }
            var v36 uint16 = 1
            _ = v36
        }
    }
    v37 := GenS2{F0: -1, F1: 2}
    fmt.Println(v5)
    v1 = 1.5
    delete(v2, 4294967295)
    fmt.Println(v28, v5)
    switch -166 {
    case 1:
        v1 = (1e10 / (float32(0.1) + 0.1))
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
