package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int; F2 uint }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int16; F1 string }
func genFn0(p0 uint) int64 {
    p0 = p0
    switch 1 {
    case 127:
        var v4 int = 0
        _ = v4
    }
    _ = p0
    return -351
}
func genFn1(p0 uint64, p1 int8, p2 string) int64 {
    v5 := []int{1, 127, 502}
    p1 = -56
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 951
    var v3 uint32 = 705
    fmt.Println(v2)
    if ((true && true) && true) {
        fmt.Println(v3, v2)
    } else {
        v2 -= (^(v2))
        v2 = (v2 ^ v2)
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int32 = (v2 - (v2 >> ((18446744073709551615) & 63)))
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int16 = -643
            var v10 int64 = 42
            v11 := GenS1{F0: 0}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        fmt.Println(v3)
        _ = fi6
        _ = v7
    }
    if (255 != 2) {
        if false {
            v12 := GenS1{F0: 2}
            if ((((1.5 - 1.5) + 0.0) + 1.5) >= -2.25) {
                v13 := GenS1{F0: 47}
                _ = v13
            } else {
                switch 9223372036854775807 {
                case 1:
                    v14 := GenS2{F0: 2, F1: 454, F2: 252}
                    _ = v14
                }
            }
            _ = v12
        }
    }
    fmt.Println(v2)
    if (!("\n\t" >= (("\n\t" + "hello") + ""))) {
        v15 := []int{4294967295, 128, -9223372036854775808}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 uint32 = v3
            _ = fi16
            _ = v17
        }
    }
    if (114 <= 2) {
        if ((-1 < -9223372036854775808) || (false && true)) {
            var v18 uint = 255
            _ = v18
        } else {
            switch int(v2) {
            case -9223372036854775808:
                switch "" {
                case "a":
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v20 := GenS2{F0: 0, F1: fi19, F2: 0}
                        v2 = -1
                        _ = fi19
                        _ = v20
                    }
                case "foo bar":
                    var v21 bool = (0 > 305)
                    _ = v21
                default:
                    v2 += int32(v3)
                }
            case 9223372036854775807:
                var v22 uint64 = 1
                _ = v22
            case 42:
                if (!false) {
                    var v23 float32 = (1.5 - 1e10)
                    if false {
                        v24 := []int{32767}
                        _ = v24
                    } else {
                        fmt.Println(v3, v23)
                    }
                    v1["\n\t"] = (("hello" + "\n\t") + "\n\t")
                    _ = v23
                }
            }
            delete(v1, ("foo bar" + "\n\t"))
            switch 443 {
            case 0:
                v2 -= v2
            case 32768:
                fmt.Println(v2)
            case 784:
                v25 := GenS1{F0: 128}
                _ = v25
            }
        }
        fmt.Println(v2, v3)
        v26 := GenS3{F0: 2, F1: "foo bar"}
        _ = v26
    }
    v2 += v2
    v27 := GenS1{F0: 1}
    for fi28 := 0; fi28 < 9; fi28++ {
        v29 := GenS2{F0: 712, F1: fi28, F2: 42}
        _ = fi28
        _ = v29
    }
    v30 := []int{9223372036854775807, -567}
    if (!(genFn1(uint64(v3), 42, "") >= -352)) {
        fmt.Println(v2, v2)
        switch "foo bar" {
        case "foo bar":
            v31 := GenS4{F0: -944, F1: ("a" + "")}
            _ = v31
        case "hello":
            if (42 > 633) {
                v3 -= (^(v3))
                var v32 string = "hello"
                fmt.Println(v32, v2)
                _ = v32
            } else {
                if true {
                    switch "hello" {
                    case "hello":
                        for fi33 := 0; fi33 < 11; fi33++ {
                            fmt.Println(fi33)
                            _ = fi33
                        }
                    case "a":
                        v3 = v3
                    default:
                        if len(v30) > 0 {
                            v30[0] = 32767
                        }
                    }
                } else {
                    v3 = (v3 / ((517) | 1))
                    v1["a"] = ""
                    v2 += v2
                }
                var v34 float64 = (0.1 / 1e10)
                _ = v34
            }
        case "a":
            v35 := v1["hello"]
            _ = v35
        default:
            v36 := GenS3{F0: 0, F1: ("" + "\n\t")}
            _ = v36
        }
    }
    if true {
        fmt.Println(v2, v3)
    }
    v3 = (v3 << ((18446744073709551615) & 63))
    switch 127 {
    case -542:
        fmt.Println(v2, v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
}
