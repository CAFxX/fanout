package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint32; F1 int32; F2 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 string }

func main() {
    var v1 bool = false
    v2 := make(map[int]string)
    var v3 uint32 = 472
    v4 := make(map[int]int)
    v5 := make(map[uint32]bool)
    v6 := make(map[int]int)
    v7 := GenS1{F0: 1}
    fmt.Println(v1, v3)
    switch "hello" {
    case "a":
        switch 2147483647 {
        case 127:
            var v8 string = "hello"
            _ = v8
        case 9223372036854775807:
            switch ("hello" + ("hello" + ("a" + "hello"))) {
            case "hello":
                v1 = v1
            default:
                for fi9 := 0; fi9 < 12; fi9++ {
                    v10 := GenS1{F0: 111}
                    fmt.Println(v3, fi9)
                    v11 := GenS1{F0: -1}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
        case -790:
            for fi12 := 0; fi12 < 7; fi12++ {
                if v1 {
                    for fi13 := 0; fi13 < 14; fi13++ {
                        if (42 == 0) {
                            v1 = v1
                        } else {
                            v14 := GenS2{F0: 0}
                            v1 = ((("a" + ("\n\t" + "")) + ("\n\t" + "foo bar")) >= "")
                            _ = v14
                        }
                        _ = fi13
                    }
                } else {
                    switch ("" + "foo bar") {
                    case "":
                        if v1 {
                            if (!((0.1 * 1e10) == ((0.0 + (3.14159 * (-2.25 / 1e10))) - float32(float64(0.1))))) {
                                switch "hello" {
                                case "a":
                                    v15 := []int{fi12, fi12, -501}
                                    _ = v15
                                case "hello":
                                    if (42 != -977) {
                                        v16 := GenS3{F0: -958}
                                        v17 := GenS4{F0: 4294967295, F1: "foo bar"}
                                        _ = v16
                                        _ = v17
                                    }
                                }
                                delete(v6, (fi12 * (fi12 + fi12)))
                                v2[262] = "\n\t"
                            }
                            switch "foo bar" {
                            case "":
                                if false {
                                    var v18 float64 = 3.14159
                                    v3 -= v3
                                    v4[(fi12 + -653)] = 9223372036854775807
                                    _ = v18
                                }
                            case "foo bar":
                                fmt.Println(v1)
                            case "\n\t":
                                v3 = v3
                            }
                            var v19 bool = false
                            _ = v19
                        }
                    case "hello":
                        if (fi12 > (fi12 + (fi12 & fi12))) {
                            v3 -= 65535
                            delete(v6, (fi12 | (^(fi12))))
                        }
                    }
                    var v20 float32 = 1e10
                    for fi21 := 0; fi21 < 20; fi21++ {
                        var v22 float64 = float64(v20)
                        v23 := []int{9223372036854775807, fi21}
                        var v24 uint = 14
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = v20
                }
                if (128 <= 255) {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        if v1 {
                            fmt.Println(fi12)
                        } else {
                            v1 = ((!(!(0 > 42))) && (0 <= int8(fi12)))
                            var v26 bool = v1
                            switch ("" + "hello") {
                            case "\n\t":
                                var v27 float32 = 0.1
                                _ = v27
                            }
                            _ = v26
                        }
                        _ = fi25
                    }
                    v28 := []int{fi12, (fi12 & fi12), 127}
                    _ = v28
                } else {
                    v29 := []int{}
                    if v1 {
                        for fi30 := 0; fi30 < 6; fi30++ {
                            v3 -= v3
                            v3 = v3
                            _ = fi30
                        }
                        v1 = false
                        v3 = 610
                    } else {
                        var v31 uint16 = 113
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 19; fi32++ {
                        fmt.Println(v1, v1, fi32)
                        fmt.Println(v3, v1, v1)
                        _ = fi32
                    }
                    _ = v29
                }
                _ = fi12
            }
        }
    }
    if v1 {
        var v33 string = "hello"
        v3 += (^(v3))
        _ = v33
    } else {
        v34 := GenS1{F0: -91}
        for fi35 := 0; fi35 < 10; fi35++ {
            if v1 {
                var v36 int = (fi35 << ((2) & 63))
                var v37 int = (^(fi35))
                var v38 float32 = -2.25
                _ = v36
                _ = v37
                _ = v38
            } else {
                v1 = false
            }
            fmt.Println(v3)
            _ = fi35
        }
        _ = v34
    }
    switch 32768 {
    case 32768:
        v39 := GenS1{F0: 103}
        _ = v39
    case 4294967295:
        v3 = (^(v3))
    }
    v1 = (v1 && (!v1))
    if (!v1) {
        v40 := GenS5{F0: (v3 ^ (^(v3))), F1: 2147483647, F2: 741}
        _ = v40
    } else {
        for fi41 := 0; fi41 < 15; fi41++ {
            if (!(0.0 == 0.0)) {
                for fi42 := 0; fi42 < 17; fi42++ {
                    fmt.Println(fi41, fi42)
                    _ = fi42
                }
            }
            var v43 uint = 421
            _ = fi41
            _ = v43
        }
        v44 := GenS6{F0: -128}
        _ = v44
    }
    v1 = (-973 <= 0)
    switch -276 {
    case 332:
        v45 := GenS6{F0: 42}
        _ = v45
    case -462:
        var v46 int32 = 1
        _ = v46
    default:
        var v47 int32 = int32(v3)
        _ = v47
    }
    switch "" {
    case "foo bar":
        var v48 int8 = -128
        _ = v48
    case "hello":
        switch -722 {
        case 2147483647:
            for fi49 := 0; fi49 < 17; fi49++ {
                v50 := GenS7{F0: 124, F1: 65535}
                _ = fi49
                _ = v50
            }
        case 266:
            switch "" {
            case "a":
                if v1 {
                    fmt.Println(v1, v1)
                    for fi51 := 0; fi51 < 7; fi51++ {
                        switch "\n\t" {
                        case "\n\t":
                            _ = len(v5)
                        case "":
                            v3 = v3
                        }
                        _ = fi51
                    }
                } else {
                    if true {
                        switch 32767 {
                        case 178:
                            fmt.Println(v3, v1)
                        case -489:
                            if false {
                                fmt.Println(v3)
                                v52 := []int{}
                                _ = v52
                            }
                        default:
                            var v53 uint32 = (v3 >> ((249) & 63))
                            _ = v53
                        }
                    }
                    var v54 uint32 = (v3 ^ 2)
                    _ = v54
                }
            case "hello":
                v55 := GenS1{F0: -128}
                _ = v55
            default:
                for fi56 := 0; fi56 < 8; fi56++ {
                    v57 := GenS3{F0: 32767}
                    _ = fi56
                    _ = v57
                }
            }
        default:
            v58 := GenS8{F0: (("\n\t" + ("foo bar" + ("" + ((("" + (("" + "foo bar") + "a")) + (("hello" + "hello") + (("" + "foo bar") + "foo bar"))) + "a")))) + ("" + "\n\t"))}
            _ = v58
        }
    case "\n\t":
        for fi59 := 0; fi59 < 18; fi59++ {
            if true {
                var v60 int16 = 32767
                _ = v60
            }
            v61 := GenS2{F0: 886}
            for fi62 := 0; fi62 < 2; fi62++ {
                fmt.Println(fi59, fi59, fi59)
                v6[fi59] = (fi59 >> ((2) & 63))
                _ = fi62
            }
            _ = fi59
            _ = v61
        }
    }
    var v63 uint64 = 657
    var v64 bool = (42 < 1)
    switch ("" + "\n\t") {
    case "foo bar":
        var v65 int8 = 0
        _ = v65
    default:
        v66 := GenS1{F0: 100}
        _ = v66
    }
    var v67 int64 = 111
    var v68 float32 = 1.5
    if ((0.0 / (0.1 + 3.14159)) < (3.14159 - 1e10)) {
        var v69 string = ((("foo bar" + ((("\n\t" + "hello") + "") + "foo bar")) + "") + (("hello" + "foo bar") + "\n\t"))
        v70 := GenS2{F0: 878}
        _ = v69
        _ = v70
    }
    var v71 uint8 = 178
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v67
    _ = v68
    _ = v7
    _ = v71
}
