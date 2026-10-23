package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int; F1 int16; F2 uint32 }
type GenS4 struct { F0 string }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[int]uint32)
    var v3 float64 = (((0.1 + 0.1) + ((float64(0.1) / 0.1) + 1.5)) - (3.14159 * 1e10))
    v4 := make(map[uint32]string)
    var v5 string = (("" + ("hello" + "\n\t")) + ("a" + ""))
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: 65535}
    for fi8 := 0; fi8 < 7; fi8++ {
        if true {
            for fi9 := 0; fi9 < 18; fi9++ {
                if true {
                    if (fi9 >= fi9) {
                        v10 := GenS2{F0: v5, F1: 255}
                        _ = v10
                    }
                    var v11 uint64 = 522
                    v5 = ("hello" + "a")
                    _ = v11
                } else {
                    if true {
                        var v12 uint8 = 128
                        if (2 < 42) {
                            var v13 uint16 = 0
                            fmt.Println(v3, v5, v12)
                            var v14 int64 = int64(v13)
                            v13 += v13
                            v1 = (^(v1))
                            _ = v13
                            _ = v14
                        } else {
                            v15 := GenS2{F0: ("" + v5), F1: 255}
                            var v16 uint = 255
                            v17 := GenS2{F0: "", F1: 757}
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        _ = v12
                    }
                    var v18 float32 = 1e10
                    _ = v18
                }
                v19 := GenS3{F0: fi9, F1: 42, F2: 2147483647}
                _ = fi9
                _ = v19
            }
        }
        fmt.Println(v5, v5)
        var v20 uint = 255
        _ = fi8
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        switch "\n\t" {
        case "a":
            v3 -= float64((0.0 / (0.1 + 0.0)))
        case "":
            for fi22 := 0; fi22 < 18; fi22++ {
                if ("hello" != "foo bar") {
                    v23 := GenS3{F0: (fi21 ^ (fi21 ^ fi22)), F1: 1, F2: 2147483647}
                    var v24 uint16 = 899
                    var v25 uint8 = 128
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    switch (-(fi21)) {
                    case 32768:
                        v26 := GenS1{F0: 181}
                        _ = v26
                    }
                    var v27 uint = 4294967295
                    _ = v27
                }
                switch (fi21 << ((378) & 63)) {
                case 127:
                    switch 32768 {
                    case 142:
                        var v28 uint32 = 2
                        _ = v28
                    }
                case 9223372036854775807:
                    v29 := []int{}
                    _ = v29
                case 42:
                    v30 := GenS3{F0: (fi22 << ((255) & 63)), F1: 965, F2: 959}
                    _ = v30
                default:
                    fmt.Println(v5, v5, v5)
                }
                _ = fi22
            }
        default:
            v5 = ""
        }
        if ((v1 <= v1) || (!((1.5 <= (0.1 * 0.0)) || (fi21 != (fi21 / (((fi21 - (fi21 >> ((65535) & 63)))) | 1)))))) {
            fmt.Println(v1)
            switch "hello" {
            case "foo bar":
                v1 += v1
            case "hello":
                fmt.Println(v1)
            default:
                for fi31 := 0; fi31 < 8; fi31++ {
                    var v32 uint32 = 2
                    var v33 int = 607
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
        }
        for fi34 := 0; fi34 < 8; fi34++ {
            var v35 int32 = -1
            _ = fi34
            _ = v35
        }
        _ = fi21
    }
    var v36 bool = false
    for fi37 := 0; fi37 < 1; fi37++ {
        v36 = false
        v5 = ("foo bar" + "hello")
        _ = fi37
    }
    v36 = true
    v38 := GenS3{F0: -9223372036854775808, F1: -32768, F2: 42}
    var v39 int8 = 23
    var v40 float64 = (1.5 / (float64(1e10) + float64(1e10)))
    var v41 int64 = -910
    if true {
        var v42 int = -825
        v43 := []int{42, 2147483647, v42}
        _ = v42
        _ = v43
    } else {
        fmt.Println(v41, v39, v3)
        for fi44 := 0; fi44 < 3; fi44++ {
            switch 0 {
            case 0:
                if (4294967295 <= 721) {
                    v45 := GenS2{F0: "a", F1: 0}
                    switch fi44 {
                    case 2147483647:
                        v46 := GenS4{F0: "\n\t"}
                        _ = v46
                    }
                    _ = v45
                }
            case 958:
                if false {
                    v47 := []int{(fi44 / ((-2147483648) | 1)), fi44, -9223372036854775808, 0}
                    v48 := GenS4{F0: v5}
                    v49 := GenS2{F0: "a", F1: 0}
                    _ = v47
                    _ = v48
                    _ = v49
                } else {
                    var v50 int8 = (v39 / ((v39) | 1))
                    _ = v50
                }
            default:
                v51 := []int{}
                _ = v51
            }
            _ = fi44
        }
        var v52 uint32 = 1
        _ = v52
    }
    v53 := GenS4{F0: "a"}
    var v54 uint16 = 264
    for fi55 := 0; fi55 < 9; fi55++ {
        v54 = (v54 / ((v54) | 1))
        var v56 float64 = (float64(-2.25) * 0.1)
        fmt.Println(v39)
        _ = fi55
        _ = v56
    }
    var v57 uint16 = (v54 ^ (v54 & 65535))
    v58 := GenS3{F0: 128, F1: -1, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v53
    _ = v54
    _ = v57
    _ = v58
    _ = v6
    _ = v7
}
