package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int; F2 int }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int; F1 uint32; F2 uint8 }
type GenS6 struct { F0 string; F1 uint16; F2 uint16 }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 uint32; F1 uint16; F2 int32 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]int)
    var v3 float32 = (0.1 + ((0.0 + 3.14159) - float32(1e10)))
    var v4 int = 2147483647
    _ = len(v2)
    if true {
        var v5 float64 = float64(0.0)
        fmt.Println(v3, v4, v3)
        fmt.Println(v1)
        _ = v5
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 uint8 = 0
        v7 += (v7 >> ((0) & 63))
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 107
    v3 = float32((float64(0.0) * -2.25))
    switch v4 {
    case -578:
        var v9 uint16 = 42
        _ = v9
    case 4294967295:
        switch v4 {
        case -1:
            switch ((("a" + (("hello" + ("foo bar" + "hello")) + "\n\t")) + (("hello" + "") + "hello")) + "hello") {
            case "":
                for fi10 := 0; fi10 < 17; fi10++ {
                    for fi11 := 0; fi11 < 15; fi11++ {
                        if (v1 || (((("" != ("hello" + "a")) && (2147483647 < -538)) && (65535 > 466)) && v1)) {
                            fmt.Println(v1)
                        } else {
                            v8 = v8
                            if false {
                                v3 -= float32((((float64(v3) + 3.14159) - 0.1) - (float64(v3) * ((1.5 * float64(v3)) / (float64((((float32(0.0) / v3) * (1.5 - ((v3 + 3.14159) + v3))) * v3)) * float64((v3 / 1.0)))))))
                                var v12 uint64 = 732
                                v13 := GenS1{F0: 42, F1: (("\n\t" + "\n\t") + ("" + "a"))}
                                switch "hello" {
                                case "hello":
                                    var v14 int32 = 2147483647
                                    _ = v14
                                case "":
                                    if true {
                                        v15 := GenS2{F0: v12}
                                        switch fi10 {
                                        case 0:
                                            var v16 uint32 = 581
                                            _ = v16
                                        case 1:
                                            v1 = (!false)
                                        }
                                        _ = v15
                                    }
                                case "foo bar":
                                    switch fi11 {
                                    case -2147483648:
                                        v4 -= (v4 << ((145) & 63))
                                    case 0:
                                        v3 -= float32(-2.25)
                                    case 2147483647:
                                        var v17 uint8 = 28
                                        _ = v17
                                    }
                                default:
                                    fmt.Println(v1)
                                }
                                v8 -= v8
                                _ = v12
                                _ = v13
                            }
                        }
                        v18 := []int{2147483647}
                        var v19 int64 = 9223372036854775807
                        _ = fi11
                        _ = v18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 6; fi20++ {
                        if ((!(0 <= 0)) || v1) {
                            if (334 < 147) {
                                var v21 bool = (0 == 9223372036854775807)
                                if v21 {
                                    v22 := GenS3{F0: "hello", F1: (fi10 << ((4294967295) & 63)), F2: (fi10 ^ -739)}
                                    v23 := GenS1{F0: 1, F1: ("\n\t" + "")}
                                    _ = v22
                                    _ = v23
                                } else {
                                    _ = len(v2)
                                    fmt.Println(fi20, v3)
                                    var v24 int64 = -9223372036854775808
                                    _ = v24
                                }
                                _ = v21
                            } else {
                                var v25 string = ((("hello" + "") + "\n\t") + "hello")
                                _ = v25
                            }
                        } else {
                            if (400 > -32768) {
                                v3 += (v3 * (((v3 + float32(3.14159)) * 0.0) * (1e10 - v3)))
                                if v1 {
                                    switch -1 {
                                    case -406:
                                        var v26 uint16 = 65535
                                        _ = v26
                                    }
                                } else {
                                    var v27 uint16 = 32767
                                    fmt.Println(v4, v4)
                                    var v28 uint32 = 65535
                                    v28 = 197
                                    v4 = fi10
                                    _ = v27
                                    _ = v28
                                }
                                v8 = 128
                            } else {
                                switch (("" + "a") + "a") {
                                case "":
                                    var v29 uint = 42
                                    _ = v29
                                case "hello":
                                    v1 = ((4294967295 >= 4294967295) && false)
                                }
                                v30 := GenS4{F0: v8, F1: (fi10 + v4)}
                                v31 := GenS1{F0: -43, F1: "a"}
                                _ = v30
                                _ = v31
                            }
                            v32 := GenS5{F0: fi20, F1: 567, F2: v8}
                            _ = v32
                        }
                        var v33 uint8 = 255
                        var v34 string = "foo bar"
                        _ = fi20
                        _ = v33
                        _ = v34
                    }
                    _ = fi10
                }
            case "\n\t":
                var v35 uint64 = 1
                _ = v35
            case "a":
                var v36 int64 = 803
                _ = v36
            }
        case -615:
            v37 := GenS3{F0: "\n\t", F1: 32767, F2: v4}
            _ = v37
        case -9223372036854775808:
            switch (v4 | (^(v4))) {
            case 32768:
                for fi38 := 0; fi38 < 3; fi38++ {
                    if v1 {
                        v39 := GenS4{F0: v8, F1: (v4 << ((1) & 63))}
                        if (18446744073709551615 < 42) {
                            v40 := GenS2{F0: uint64(v8)}
                            var v41 uint = 123
                            v42 := v2[1]
                            _ = v40
                            _ = v41
                            _ = v42
                        } else {
                            for fi43 := 0; fi43 < 8; fi43++ {
                                if (42 >= -1) {
                                    var v44 uint64 = 2
                                    v45 := GenS2{F0: (v44 >> ((42) & 63))}
                                    _ = v44
                                    _ = v45
                                }
                                var v46 int = 32768
                                var v47 int32 = 305
                                _ = fi43
                                _ = v46
                                _ = v47
                            }
                            if v1 {
                                var v48 uint = 4294967295
                                if ((1 < 813) && false) {
                                    v1 = false
                                    v49 := GenS6{F0: "", F1: 42, F2: 430}
                                    _ = v49
                                }
                                _ = v48
                            } else {
                                v50 := []int{}
                                v51 := GenS7{F0: (42 <= 4294967295), F1: ("a" + "foo bar")}
                                var v52 int16 = 32767
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                        }
                        _ = v39
                    } else {
                        v53 := v2[0]
                        v3 = v3
                        var v54 bool = v1
                        _ = v53
                        _ = v54
                    }
                    var v55 int16 = 42
                    _ = fi38
                    _ = v55
                }
            case 127:
                if v1 {
                    for fi56 := 0; fi56 < 11; fi56++ {
                        var v57 int16 = -39
                        if v1 {
                            var v58 uint64 = 0
                            v1 = v1
                            v59 := v2[65535]
                            _ = v58
                            _ = v59
                        }
                        _ = fi56
                        _ = v57
                    }
                } else {
                    v60 := []int{}
                    _ = v60
                }
            case -480:
                for fi61 := 0; fi61 < 1; fi61++ {
                    if v1 {
                        fmt.Println(fi61, v4, v3)
                    }
                    for fi62 := 0; fi62 < 20; fi62++ {
                        var v63 uint = 1
                        switch 128 {
                        case 9223372036854775807:
                            v64 := GenS1{F0: -1, F1: "a"}
                            _ = v64
                        }
                        v65 := v2[1]
                        _ = fi62
                        _ = v63
                        _ = v65
                    }
                    _ = fi61
                }
            default:
                switch v4 {
                case -1:
                    if (((0.1 / 0.1) + (-2.25 + ((3.14159 * 0.1) + 1.5))) >= 1e10) {
                        v8 = v8
                    }
                case 42:
                    v66 := []int{9223372036854775807, v4, (^(v4))}
                    _ = v66
                }
            }
        }
    }
    if true {
        if false {
            v8 += (v8 / (((v8 * 0)) | 1))
        } else {
            if v1 {
                var v67 int8 = -1
                _ = v67
            }
        }
        var v68 uint8 = (^(v8))
        var v69 uint16 = 42
        _ = v68
        _ = v69
    } else {
        v1 = v1
    }
    v70 := v2[881]
    var v71 uint64 = 42
    if (v8 == v8) {
        fmt.Println(v4)
        var v72 int16 = 498
        fmt.Println(v3, v4, v70)
        _ = v72
    }
    v73 := []int{int(v71), v70, (v4 & (v4 | v70))}
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        fmt.Println(v4)
    }
    v71 += 66
    for fi74 := 0; fi74 < 15; fi74++ {
        v75 := GenS8{F0: 0, F1: 32767, F2: -1}
        if len(v73) >= 2 {
            v73 = v73[1:len(v73)]
        }
        _ = fi74
        _ = v75
    }
    v70 = v70
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("slicelen v73:", len(v73), v73)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v70
    _ = v71
    _ = v73
    _ = v8
}
