package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint; F1 uint; F2 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 int8; F2 uint }
func genFn0(p0 int, p1 int16) uint64 {
    v7 := GenS1{F0: p1, F1: 4294967295, F2: 42}
    p1 = (p1 & p1)
    _ = p0
    _ = p1
    _ = v7
    return 2
}

func main() {
    v1 := []uint32{42, 531, 0}
    var v2 string = "hello"
    var v3 float32 = (-2.25 + 3.14159)
    var v4 float32 = 1.5
    v5 := make(map[uint32]string)
    v6 := make(map[int]string)
    var v8 int8 = -1
    v2 = (v2 + (("\n\t" + v2) + ("" + (((v2 + "\n\t") + v2) + "a"))))
    var v9 int8 = v8
    var v10 int = 0
    if (float64(0.1) > 0.0) {
        _ = len(v6)
        for fi11 := 0; fi11 < 10; fi11++ {
            _ = len(v6)
            v12 := GenS1{F0: 202, F1: 42, F2: 32767}
            _ = fi11
            _ = v12
        }
    }
    var v13 int32 = -497
    switch "a" {
    case "foo bar":
        switch ("hello" + v2) {
        case "a":
            switch (("a" + ("hello" + v2)) + v2) {
            case "\n\t":
                if (1000 <= 412) {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        fmt.Println(fi14, v3, v9)
                        _ = fi14
                    }
                } else {
                    v13 += v13
                }
            case "foo bar":
                var v15 uint64 = 18446744073709551615
                _ = v15
            case "":
                if true {
                    v9 = -17
                    switch ((v2 + v2) + ("foo bar" + "a")) {
                    case "hello":
                        for fi16 := 0; fi16 < 13; fi16++ {
                            v17 := GenS1{F0: 32767, F1: 0, F2: 65535}
                            _ = fi16
                            _ = v17
                        }
                    case "":
                        delete(v6, v10)
                    default:
                        switch v10 {
                        case 42:
                            var v18 int64 = 0
                            _ = v18
                        case 32767:
                            if true {
                                switch v2 {
                                case "\n\t":
                                    fmt.Println(v10)
                                case "foo bar":
                                    v9 += 127
                                default:
                                    v19 := GenS2{F0: 1, F1: 313}
                                    _ = v19
                                }
                                for fi20 := 0; fi20 < 10; fi20++ {
                                    if true {
                                        for fi21 := 0; fi21 < 11; fi21++ {
                                            v9 -= -96
                                            var v22 uint = 559
                                            fmt.Println(v13, v22, v22)
                                            _ = fi21
                                            _ = v22
                                        }
                                    }
                                    _ = fi20
                                }
                                fmt.Println(v9)
                            }
                        case 127:
                            var v23 uint = 42
                            _ = v23
                        }
                    }
                }
            }
        case "hello":
            v9 = (v9 % (((-(v8))) | 1))
        case "\n\t":
            switch v10 {
            case 32767:
                _ = len(v5)
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        v24 := GenS3{F0: 70, F1: 475, F2: 9223372036854775807}
        _ = v24
    }
    if false {
        var v25 bool = true
        _ = v25
    } else {
        for fi26 := 0; fi26 < 7; fi26++ {
            if (776 > 410) {
                fmt.Println(v2)
            } else {
                v27 := GenS4{F0: v2}
                _ = len(v6)
                _ = v27
            }
            if false {
                _ = len(v6)
                var v28 bool = false
                _ = v28
            }
            _ = fi26
        }
        for fi29 := 0; fi29 < 6; fi29++ {
            switch (v2 + (("foo bar" + v2) + (v2 + (v2 + (v2 + v2))))) {
            case "hello":
                v4 = (v4 / v3)
            default:
                v30 := GenS1{F0: -252, F1: 4294967295, F2: 883}
                _ = v30
            }
            if true {
                v10 -= v10
                v10 = fi29
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            switch (v10 + (-(fi29))) {
            case 0:
                delete(v5, 65535)
            }
            _ = fi29
        }
        v31 := GenS4{F0: "a"}
        _ = v31
    }
    fmt.Println(v3, v2, v2)
    for fi32 := 0; fi32 < 14; fi32++ {
        var v33 uint8 = 128
        _ = fi32
        _ = v33
    }
    v34 := GenS5{F0: 0, F1: 42, F2: 482}
    var v35 uint = 65535
    var v36 uint16 = 46
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
