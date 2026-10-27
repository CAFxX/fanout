package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int64; F2 int32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int8; F1 uint }
type GenS5 struct { F0 int8; F1 bool; F2 string }
type GenS6 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{871}
    var v3 string = "hello"
    v4 := make(map[int]uint32)
    v5 := GenS1{F0: 127}
    fmt.Println(v3)
    v3 = v3
    v6 := GenS2{F0: 209, F1: 9223372036854775807, F2: 629}
    v3 = v3
    var v7 uint16 = 196
    v7 = v7
    v8 := GenS3{F0: 2, F1: 32768}
    switch v3 {
    case "foo bar":
        v9 := GenS3{F0: 0, F1: -426}
        _ = v9
    case "":
        v7 = v7
    default:
        if (159 > 346) {
            var v10 float32 = (1.5 * (-2.25 + 0.0))
            var v11 uint16 = (v7 + (^(v7)))
            v7 += 2
            _ = v10
            _ = v11
        } else {
            v3 = "a"
            for fi12 := 0; fi12 < 11; fi12++ {
                fmt.Println(fi12, v7, v7)
                _ = fi12
            }
        }
    }
    if false {
        if ((v7 & (^(v7))) > v7) {
            var v13 uint = 18446744073709551615
            v3 = "a"
            _ = v13
        }
        fmt.Println(v3, v7)
        var v14 bool = (false && (1 >= 2))
        _ = v14
    }
    v4[128] = 1
    if false {
        v7 = (v7 * v7)
        v3 = v3
    } else {
        v1[1] = 9223372036854775807
        v3 = ("\n\t" + (("hello" + "foo bar") + v3))
        var v15 uint = 701
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        _ = len(v2)
        for fi17 := 0; fi17 < 1; fi17++ {
            if true {
                fmt.Println(fi17, v3, v7)
            } else {
                var v18 uint16 = (v7 & (^(v7)))
                fmt.Println(v18, v3, v3)
                v3 = v3
                var v19 uint16 = (^(v18))
                v3 = (v3 + ("\n\t" + (v3 + v3)))
                var v20 uint = 4294967295
                _ = v18
                _ = v19
                _ = v20
            }
            v7 = v7
            fmt.Println(v3, v7)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v21 float32 = -2.25
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi17
            _ = v21
        }
        _ = fi16
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if false {
            var v23 uint8 = 134
            if false {
                var v24 float64 = 1e10
                _ = v24
            } else {
                var v25 int = int(v7)
                _ = v25
            }
            _ = v23
        } else {
            fmt.Println(v7, v3)
            for fi26 := 0; fi26 < 17; fi26++ {
                switch v3 {
                case "foo bar":
                    v7 -= v7
                case "a":
                    var v27 int64 = 606
                    _ = v27
                case "hello":
                    var v28 uint16 = v7
                    _ = v28
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi26
            }
            v29 := GenS4{F0: -128, F1: 2}
            _ = v29
        }
        _ = fi22
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        switch v3 {
        case "":
            v7 += v7
        }
        _ = fi30
    }
    if false {
        v3 = v3
        if true {
            if true {
                var v31 uint16 = v7
                if true {
                    v7 -= v31
                    switch ((v3 + v3) + ((v3 + "foo bar") + (v3 + v3))) {
                    case "foo bar":
                        var v32 int8 = 1
                        _ = v32
                    default:
                        v33 := GenS2{F0: 2, F1: -9223372036854775808, F2: -1}
                        _ = v33
                    }
                }
                _ = v31
            } else {
                switch 803 {
                case 42:
                    fmt.Println(v7, v3)
                case -569:
                    for fi34 := 0; fi34 < 11; fi34++ {
                        fmt.Println(v7, v3)
                        _ = fi34
                    }
                case 0:
                    v35 := v4[-9223372036854775808]
                    _ = v35
                default:
                    var v36 uint32 = 2147483648
                    _ = v36
                }
                for fi37 := 0; fi37 < 1; fi37++ {
                    switch (((v3 + "foo bar") + (v3 + "\n\t")) + v3) {
                    case "foo bar":
                        var v38 uint8 = 128
                        _ = v38
                    }
                    v3 = v3
                    _ = fi37
                }
                var v39 uint32 = 0
                _ = v39
            }
            for fi40 := 0; fi40 < 17; fi40++ {
                for fi41 := 0; fi41 < 18; fi41++ {
                    fmt.Println(fi40, v3)
                    _ = fi41
                }
                fmt.Println(fi40)
                _ = fi40
            }
        } else {
            if len(v2) > 0 {
                v2[0] = -9223372036854775808
            }
            fmt.Println(v7)
        }
        v42 := GenS5{F0: -66, F1: (128 < -147), F2: "hello"}
        _ = v42
    }
    v7 = (^(v7))
    v43 := GenS6{F0: (v7 ^ v7), F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
