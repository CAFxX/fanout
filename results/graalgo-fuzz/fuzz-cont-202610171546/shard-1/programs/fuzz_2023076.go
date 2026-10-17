package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int16; F1 int64; F2 int64 }
type GenS4 struct { F0 uint64; F1 bool }

func main() {
    var v1 int32 = 0
    var v2 uint64 = 42
    var v3 uint = 255
    v4 := make(map[string]string)
    v1 -= (v1 + v1)
    switch -114 {
    case -923:
        if true {
            for fi5 := 0; fi5 < 6; fi5++ {
                v2 += (v2 ^ v2)
                var v6 uint32 = 65535
                v2 = v2
                _ = fi5
                _ = v6
            }
            if true {
                if (!false) {
                    switch "hello" {
                    case "a":
                        v3 = 196
                    }
                    var v7 bool = (-32768 >= 1000)
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 11; fi8++ {
                        v9 := GenS1{F0: -1, F1: uint8(v3)}
                        _ = fi8
                        _ = v9
                    }
                }
            }
        }
    case 4294967295:
        v1 = (v1 & 42)
    default:
        if true {
            var v10 string = "\n\t"
            _ = v10
        }
    }
    var v11 int32 = v1
    fmt.Println(v3, v1)
    if false {
        delete(v4, "\n\t")
    } else {
        v1 += v1
    }
    v12 := []int{9223372036854775807, 32767}
    fmt.Println(v11, v3, v1)
    v4[("hello" + "foo bar")] = "a"
    fmt.Println(v11)
    if (3.14159 <= 1e10) {
        if (0 < 0) {
            for fi13 := 0; fi13 < 12; fi13++ {
                for fi14 := 0; fi14 < 18; fi14++ {
                    v15 := GenS2{F0: 0, F1: -607}
                    if true {
                        v16 := GenS1{F0: fi13, F1: 111}
                        var v17 uint = (^(v3))
                        delete(v4, ("foo bar" + ""))
                        _ = v16
                        _ = v17
                    } else {
                        if false {
                            var v18 uint64 = 820
                            _ = v18
                        } else {
                            v19 := GenS1{F0: fi14, F1: 42}
                            v1 -= v1
                            var v20 string = "foo bar"
                            _ = v19
                            _ = v20
                        }
                        if len(v12) > 0 {
                            v12[0] = (fi14 + fi13)
                        }
                        var v21 int64 = -720
                        _ = v21
                    }
                    var v22 uint32 = 899
                    switch fi13 {
                    case -246:
                        v3 = v3
                    }
                    _ = fi14
                    _ = v15
                    _ = v22
                }
                var v23 int8 = 100
                var v24 int16 = 1
                _ = fi13
                _ = v23
                _ = v24
            }
        } else {
            var v25 int64 = -890
            switch -600 {
            case 32767:
                var v26 int16 = -32768
                _ = v26
            case 2147483647:
                var v27 int32 = v11
                _ = v27
            }
            _ = v25
        }
    }
    var v28 float64 = (float64(float32(-2.25)) / (float64(((((((((1.5 / 3.14159) - 1.5) * 1e10) - -2.25) + 0.1) + -2.25) / 1.5) * 1e10)) + float64(1e10)))
    for fi29 := 0; fi29 < 10; fi29++ {
        if true {
            v11 = v11
            var v30 uint16 = 1
            _ = v30
        }
        _ = fi29
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        v32 := GenS3{F0: -32768, F1: -605, F2: 1}
        _ = fi31
        _ = v32
    }
    var v33 int = -9223372036854775808
    var v34 uint64 = (^(v2))
    if (true && false) {
        switch "foo bar" {
        case "foo bar":
            v28 = v28
        default:
            for fi35 := 0; fi35 < 18; fi35++ {
                if (0 != 42) {
                    v36 := GenS4{F0: 490, F1: (!(true || false))}
                    var v37 int64 = 42
                    var v38 uint8 = 0
                    _ = v36
                    _ = v37
                    _ = v38
                }
                if false {
                    fmt.Println(v1, v1, v3)
                }
                _ = fi35
            }
        }
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
