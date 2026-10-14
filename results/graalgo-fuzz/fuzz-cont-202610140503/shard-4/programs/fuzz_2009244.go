package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 bool; F1 string; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 uint32 }
func genFn0() uint16 {
    var v7 int = 32768
    for fi8 := 0; fi8 < 14; fi8++ {
        v7 = fi8
        if true {
            for fi9 := 0; fi9 < 1; fi9++ {
                v10 := GenS1{F0: 9223372036854775807, F1: 143, F2: 20}
                var v11 int8 = 127
                var v12 float32 = float32(1e10)
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            switch (("hello" + "foo bar") + ("foo bar" + "hello")) {
            case "\n\t":
                var v13 int16 = -1
                _ = v13
            }
        }
        _ = fi8
    }
    _ = v7
    return 787
}
func genFn1(p0 uint16) uint8 {
    v14 := []int{1, 1, 32768}
    fmt.Println(p0)
    _ = p0
    _ = v14
    return 225
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    var v5 int32 = 156
    v6 := []string{("foo bar" + "foo bar"), (("a" + "foo bar") + ""), (("" + "hello") + "hello")}
    for fi15 := 0; fi15 < 16; fi15++ {
        if (3.14159 != -2.25) {
            if (1 <= 141) {
                fmt.Println(fi15)
                switch "" {
                case "a":
                    fmt.Println(v5, v5)
                }
            } else {
                var v16 string = "\n\t"
                _ = v16
            }
            var v17 uint8 = 255
            _ = v17
        } else {
            v18 := GenS2{F0: 89, F1: 12}
            _ = v18
        }
        _ = fi15
    }
    v5 = (v5 + 2147483647)
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := GenS3{F0: (false || false), F1: "hello", F2: 0}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        fmt.Println(fi21, fi21)
        _ = fi21
    }
    switch "hello" {
    case "foo bar":
        v5 += v5
    case "":
        switch "a" {
        case "":
            delete(v3, -534)
        case "\n\t":
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 float32 = 1.5
                _ = fi22
                _ = v23
            }
        default:
            v24 := GenS4{F0: v5, F1: 705}
            _ = v24
        }
    case "hello":
        var v25 uint64 = 18446744073709551615
        _ = v25
    }
    v5 -= v5
    v26 := GenS3{F0: (((v5 ^ -1) < v5) || true), F1: "", F2: 255}
    if (true || false) {
        _ = len(v6)
        switch "foo bar" {
        case "\n\t":
            var v27 int32 = (v5 & (v5 | (v5 << ((259) & 63))))
            _ = v27
        case "foo bar":
            if len(v6) > 0 {
                v6[0] = "hello"
            }
        }
    } else {
        if true {
            var v28 uint8 = 127
            if true {
                v2[uint32(v28)] = (!(744 != 0))
                for fi29 := 0; fi29 < 5; fi29++ {
                    fmt.Println(fi29, v5, v5)
                    var v30 uint8 = v28
                    v31 := GenS5{F0: true}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                v28 = 165
            } else {
                v28 -= (^(v28))
                _ = len(v1)
                if true {
                    var v32 int32 = v5
                    _ = v32
                }
            }
            v28 += v28
            _ = v28
        }
        for fi33 := 0; fi33 < 12; fi33++ {
            var v34 float64 = 0.0
            var v35 float32 = (3.14159 + ((1e10 - 0.0) + (-2.25 / ((float32(v34) + 1.5) - 1e10))))
            v36 := GenS1{F0: 0, F1: 950, F2: 155}
            _ = fi33
            _ = v34
            _ = v35
            _ = v36
        }
        if false {
            v37 := GenS3{F0: false, F1: "\n\t", F2: 0}
            for fi38 := 0; fi38 < 11; fi38++ {
                var v39 uint16 = 668
                _ = len(v1)
                _ = fi38
                _ = v39
            }
            var v40 uint8 = 127
            _ = v37
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 2; fi41++ {
        for fi42 := 0; fi42 < 14; fi42++ {
            switch "" {
            case "":
                v43 := GenS2{F0: 198, F1: 75}
                _ = v43
            case "a":
                fmt.Println(fi41, fi41)
            default:
                v5 = (v5 / ((v5) | 1))
            }
            _ = fi42
        }
        delete(v4, 65535)
        _ = fi41
    }
    var v44 uint16 = 65535
    var v45 uint = 18446744073709551615
    v44 -= v44
    var v46 float32 = 0.1
    v6 = append(v6, ("hello" + ("a" + (("foo bar" + "") + "hello"))))
    fmt.Println(v5, v5)
    v47 := GenS6{F0: 16, F1: 17}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
