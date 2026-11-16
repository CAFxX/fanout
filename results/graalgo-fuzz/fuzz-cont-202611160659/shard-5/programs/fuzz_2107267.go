package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int16; F1 string }
type GenS9 struct { F0 uint32; F1 int }
func genFn0() int {
    var v5 uint8 = 128
    switch 1 {
    case 1:
        if (!false) {
            v5 -= (v5 + v5)
            v5 += (v5 - 16)
        } else {
            var v6 bool = (42 < -1)
            v6 = false
            _ = v6
        }
    case 756:
        v7 := GenS1{F0: 1}
        _ = v7
    case 620:
        v5 = v5
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS2{F0: 39}
        _ = fi8
        _ = v9
    }
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := []int{}
    var v2 uint32 = 2147483647
    v3 := make(map[string]string)
    var v4 float64 = (1.5 - (3.14159 / (-2.25 - 1.5)))
    for fi10 := 0; fi10 < 16; fi10++ {
        switch "a" {
        case "foo bar":
            switch (fi10 / ((1) | 1)) {
            case -2147483648:
                fmt.Println(v4, v4, v2)
            default:
                var v11 uint8 = 255
                _ = v11
            }
        default:
            for fi12 := 0; fi12 < 8; fi12++ {
                v4 -= 1.5
                v2 = (v2 | (v2 & v2))
                _ = fi12
            }
        }
        _ = len(v1)
        _ = fi10
    }
    v4 = 3.14159
    for fi13 := 0; fi13 < 8; fi13++ {
        if true {
            _ = len(v1)
            var v14 uint16 = 1
            _ = v14
        }
        var v15 uint16 = 893
        for fi16 := 0; fi16 < 2; fi16++ {
            switch (("" + ("a" + "hello")) + ("\n\t" + "")) {
            case "hello":
                v17 := GenS3{F0: -753, F1: 848}
                _ = v17
            default:
                var v18 float32 = 0.1
                _ = v18
            }
            v4 -= (v4 + (v4 * (v4 + float64((1e10 - 1e10)))))
            if (false && false) {
                if len(v1) > 0 {
                    v1[0] = 127
                }
                v19 := GenS4{F0: -32768, F1: 0}
                _ = v19
            }
            v20 := GenS3{F0: 928, F1: 1}
            var v21 int16 = 654
            _ = fi16
            _ = v20
            _ = v21
        }
        _ = fi13
        _ = v15
    }
    v4 -= float64(1.5)
    switch -1 {
    case -748:
        for fi22 := 0; fi22 < 18; fi22++ {
            for fi23 := 0; fi23 < 7; fi23++ {
                switch fi22 {
                case 1:
                    v24 := GenS5{F0: 0, F1: -32768}
                    _ = v24
                case 128:
                    if (-128 == 42) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        v25 := GenS6{F0: 1, F1: 543, F2: 920}
                        fmt.Println(fi22)
                        var v26 uint = 4294967295
                        v27 := GenS7{F0: "hello"}
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                case 0:
                    _ = len(v1)
                default:
                    v28 := GenS8{F0: 1000, F1: ("hello" + "foo bar")}
                    _ = v28
                }
                var v29 int64 = 1
                var v30 int64 = 0
                _ = fi23
                _ = v29
                _ = v30
            }
            var v31 uint8 = 42
            _ = fi22
            _ = v31
        }
    default:
        var v32 string = "hello"
        _ = v32
    }
    switch genFn0() {
    case -2147483648:
        if (1 <= -32768) {
            _ = len(v1)
        }
    case 5:
        v33 := GenS7{F0: ""}
        _ = v33
    case -9223372036854775808:
        v3["hello"] = ""
    default:
        v34 := v3[""]
        _ = v34
    }
    var v35 uint = 860
    switch "\n\t" {
    case "hello":
        var v36 int16 = -32768
        _ = v36
    default:
        v1 = append(v1, 117)
    }
    switch "hello" {
    case "":
        for fi37 := 0; fi37 < 18; fi37++ {
            v38 := GenS4{F0: -1, F1: 2}
            _ = fi37
            _ = v38
        }
    }
    if true {
        fmt.Println(v2, v4, v2)
        v39 := GenS9{F0: (v2 / ((v2) | 1)), F1: genFn0()}
        _ = len(v1)
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
