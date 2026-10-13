package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int32; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 string; F2 uint64 }
type GenS5 struct { F0 int8; F1 bool; F2 uint16 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 uint }

func main() {
    v1 := []string{}
    v2 := []uint32{4294967295, 4294967295}
    var v3 uint32 = 2
    if (true || (false || true)) {
        var v4 uint = 42
        _ = v4
    } else {
        v5 := GenS1{F0: 0}
        fmt.Println(v3)
        if ((false && true) || false) {
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 uint = 907
                v8 := GenS1{F0: -1}
                for fi9 := 0; fi9 < 11; fi9++ {
                    switch fi6 {
                    case 4294967295:
                        if ((v7 & (^(v7))) >= v7) {
                            var v10 int64 = 0
                            _ = v10
                        } else {
                            var v11 float32 = (float32((float64((1e10 + 1e10)) / 1.0)) * 0.0)
                            _ = v11
                        }
                    case -954:
                        var v12 uint32 = (v3 >> ((uint(v3)) & 63))
                        _ = v12
                    default:
                        if false {
                            fmt.Println(v7)
                            v3 = (^(v3))
                        }
                    }
                    v7 += v7
                    v13 := GenS2{F0: 0, F1: -445, F2: uint8(v3)}
                    _ = fi9
                    _ = v13
                }
                _ = fi6
                _ = v7
                _ = v8
            }
            var v14 int32 = -185
            _ = v14
        } else {
            switch 9223372036854775807 {
            case 202:
                var v15 uint = 1
                _ = v15
            case 32767:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        _ = v5
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v3 = (v3 % (((^(v3))) | 1))
        _ = fi16
    }
    v17 := GenS3{F0: 283}
    v18 := GenS4{F0: uint16(v3), F1: "hello", F2: 2}
    for fi19 := 0; fi19 < 7; fi19++ {
        switch (-(fi19)) {
        case -2147483648:
            if true {
                var v20 uint = 445
                var v21 int32 = 846
                var v22 uint8 = 127
                _ = v20
                _ = v21
                _ = v22
            }
        case 32768:
            var v23 float32 = (0.1 + 0.0)
            _ = v23
        default:
            v24 := GenS1{F0: 42}
            _ = v24
        }
        v25 := GenS1{F0: 2147483647}
        if true {
            v1 = append(v1, "a")
            var v26 float64 = 1.5
            for fi27 := 0; fi27 < 15; fi27++ {
                v1 = append(v1, (("hello" + (("a" + "a") + (("hello" + (("hello" + ("foo bar" + "hello")) + "\n\t")) + "foo bar"))) + ("a" + "\n\t")))
                var v28 int8 = -1
                switch fi19 {
                case 42:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 0:
                    v29 := GenS5{F0: -1, F1: ((fi19 << ((907) & 63)) > -550), F2: 488}
                    _ = v29
                case 32768:
                    v3 = v3
                default:
                    v26 = ((v26 * (v26 + v26)) / v26)
                }
                _ = fi27
                _ = v28
            }
            _ = v26
        }
        _ = fi19
        _ = v25
    }
    switch "a" {
    case "hello":
        var v30 int16 = 0
        _ = v30
    case "a":
        switch 1 {
        case 9223372036854775807:
            fmt.Println(v3)
        case 32767:
            v2 = append(v2, (^(v3)))
        case 2147483647:
            var v31 float32 = (-2.25 / -2.25)
            _ = v31
        }
    case "":
        if (("foo bar" + "foo bar") == ("foo bar" + "hello")) {
            var v32 uint32 = 4294967295
            v32 += v3
            _ = v32
        } else {
            v33 := GenS6{F0: 0, F1: 186}
            if len(v1) > 0 {
                v1[0] = ("a" + ((("\n\t" + "a") + ("a" + "\n\t")) + ""))
            }
            v34 := GenS4{F0: 65535, F1: "foo bar", F2: 762}
            _ = v33
            _ = v34
        }
    default:
        v3 = 65535
    }
    var v35 int8 = 0
    for fi36 := 0; fi36 < 12; fi36++ {
        v3 -= (v3 ^ 2147483648)
        var v37 uint16 = 979
        v38 := GenS7{F0: 65535}
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 uint64 = 191
    var v40 uint16 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v40
}
