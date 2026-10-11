package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 string }
func genFn0() int64 {
    if true {
        var v5 string = "\n\t"
        _ = v5
    } else {
        var v6 bool = (false && (2 == 42))
        var v7 string = "hello"
        _ = v6
        _ = v7
    }
    var v8 uint64 = 503
    v8 = 234
    _ = v8
    return -452
}
func genFn1() int64 {
    if (!(!false)) {
        if (65535 >= 151) {
            var v9 float32 = 1.5
            v10 := []int{}
            var v11 float32 = 0.0
            _ = v10
            _ = v11
            _ = v9
        } else {
            var v12 uint = 2
            if true {
                switch 4294967295 {
                case -2147483648:
                    var v13 int = -9223372036854775808
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 16; fi14++ {
                        if (!true) {
                            v12 -= (v12 + 65535)
                        } else {
                            var v15 int64 = -1
                            _ = v15
                        }
                        _ = fi14
                    }
                }
            } else {
                var v16 int32 = 2147483647
                switch "" {
                case "hello":
                    v17 := GenS1{F0: -9223372036854775808}
                    _ = v17
                }
                _ = v16
            }
            _ = v12
        }
        var v18 uint64 = 463
        v19 := GenS1{F0: -1}
        _ = v18
        _ = v19
    }
    switch "hello" {
    case "":
        var v20 uint8 = 103
        _ = v20
    case "\n\t":
        for fi21 := 0; fi21 < 1; fi21++ {
            v22 := GenS1{F0: 42}
            fmt.Println(fi21)
            _ = fi21
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 uint = 435
            v25 := GenS1{F0: genFn0()}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    if (612 >= 1) {
        var v26 float64 = 1e10
        v26 += (1e10 / 1.0)
        if (!false) {
            var v27 uint16 = 0
            v26 = (1.5 - v26)
            _ = v27
        }
        _ = v26
    } else {
        var v28 string = ""
        v29 := GenS1{F0: genFn0()}
        _ = v28
        _ = v29
    }
    v30 := []int{0, 0, 9223372036854775807}
    if (768 == 18446744073709551615) {
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
        v31 := GenS1{F0: genFn0()}
        var v32 int8 = 100
        _ = v31
        _ = v32
    }
    _ = v30
    return 42
}

func main() {
    var v1 int = 32767
    var v2 uint = 42
    v3 := make(map[uint32]uint32)
    v4 := []int{(-(v1)), 0}
    if len(v4) > 0 {
        v4[0] = v1
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        for fi34 := 0; fi34 < 4; fi34++ {
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v1 += (^(v1))
                v35 := GenS1{F0: genFn1()}
                _ = v35
            } else {
                fmt.Println(v1)
            }
            _ = len(v3)
            switch "foo bar" {
            case "\n\t":
                v2 -= v2
            case "a":
                var v36 uint32 = 42
                _ = v36
            }
            _ = fi34
        }
        _ = fi33
    }
    v2 = (v2 & 370)
    v37 := GenS2{F0: 0, F1: (("a" + ("" + "a")) + "hello")}
    v2 = v2
    v1 = (v1 << (((^(v2))) & 63))
    v2 -= 629
    var v38 uint16 = 32767
    var v39 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
