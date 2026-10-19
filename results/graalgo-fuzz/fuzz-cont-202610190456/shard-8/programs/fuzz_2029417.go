package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint64; F1 bool; F2 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 uint32; F2 int8 }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: -1, F1: false, F2: 42}
        switch fi7 {
        case -2147483648:
            fmt.Println(fi7)
        case 0:
            if false {
                for fi9 := 0; fi9 < 9; fi9++ {
                    v10 := GenS1{F0: -128, F1: true, F2: 42}
                    _ = fi9
                    _ = v10
                }
                switch fi7 {
                case 1:
                    v11 := GenS2{F0: 2147483648, F1: 187}
                    _ = v11
                case 0:
                    var v12 uint16 = 255
                    _ = v12
                case 32768:
                    var v13 int64 = 1
                    _ = v13
                default:
                    var v14 int32 = -1
                    _ = v14
                }
                var v15 string = "\n\t"
                _ = v15
            } else {
                var v16 int32 = -1
                var v17 uint32 = 4294967295
                v16 += (v16 + (v16 & (v16 | v16)))
                _ = v16
                _ = v17
            }
        case -919:
            var v18 uint32 = 4294967295
            _ = v18
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            if false {
                var v20 int32 = -741
                _ = v20
            } else {
                var v21 float32 = 3.14159
                _ = v21
            }
            switch "\n\t" {
            case "a":
                fmt.Println(fi19, fi19)
            case "":
                switch (^(fi7)) {
                case 4294967295:
                    var v22 int64 = 1
                    _ = v22
                default:
                    switch (-(fi19)) {
                    case 42:
                        var v23 int32 = 664
                        _ = v23
                    case -2147483648:
                        var v24 bool = false
                        switch "" {
                        case "hello":
                            if true {
                                var v25 string = "foo bar"
                                v25 = "\n\t"
                                _ = v25
                            }
                        }
                        switch fi19 {
                        case 127:
                            v26 := GenS1{F0: 20, F1: v24, F2: 1}
                            _ = v26
                        case 32768:
                            fmt.Println(v24, fi7)
                        case -843:
                            switch 0 {
                            case 92:
                                v27 := []int{0}
                                _ = v27
                            case 32768:
                                var v28 int64 = -9223372036854775808
                                _ = v28
                            default:
                                var v29 bool = (0.0 < 0.1)
                                _ = v29
                            }
                        default:
                            v30 := GenS3{F0: 65535, F1: ("" + ((("\n\t" + ("a" + (("" + "foo bar") + ""))) + "hello") + "\n\t"))}
                            _ = v30
                        }
                        _ = v24
                    default:
                        if (false || (65535 < 65535)) {
                            var v31 string = ""
                            var v32 int = fi19
                            v33 := GenS3{F0: 772, F1: ""}
                            _ = v31
                            _ = v32
                            _ = v33
                        } else {
                            var v34 bool = (false || true)
                            var v35 uint16 = 255
                            _ = v34
                            _ = v35
                        }
                    }
                }
            }
            _ = fi19
        }
        _ = fi7
        _ = v8
    }
    var v36 uint = 4294967295
    if true {
        if (((!true) || ((284 <= 1) || true)) || (!false)) {
            var v37 int32 = 0
            _ = v37
        } else {
            if true {
                var v38 uint8 = 127
                v39 := []int{-582, 344, 128}
                _ = v38
                _ = v39
            }
            v36 += v36
            for fi40 := 0; fi40 < 1; fi40++ {
                for fi41 := 0; fi41 < 9; fi41++ {
                    var v42 float32 = 3.14159
                    _ = fi41
                    _ = v42
                }
                _ = fi40
            }
        }
    } else {
        v36 = 65535
    }
    _ = v36
    return 376
}

func main() {
    var v1 uint32 = 2
    v2 := []string{("\n\t" + "")}
    v3 := make(map[int]bool)
    v4 := []uint32{53, 0, v1}
    var v5 int8 = -35
    var v6 uint32 = v1
    var v43 float64 = (float64(((0.1 + (3.14159 * -2.25)) + 3.14159)) * (3.14159 * 1.5))
    fmt.Println(v43, v43)
    for fi44 := 0; fi44 < 8; fi44++ {
        fmt.Println(fi44, v6)
        if true {
            for fi45 := 0; fi45 < 20; fi45++ {
                var v46 string = "foo bar"
                var v47 float32 = 1.5
                var v48 int32 = 2147483647
                _ = fi45
                _ = v46
                _ = v47
                _ = v48
            }
            v1 = (v6 << ((366) & 63))
        } else {
            v3[(fi44 / ((-611) | 1))] = (("" + (("a" + ("foo bar" + "a")) + "foo bar")) > "foo bar")
        }
        _ = fi44
    }
    v5 -= (v5 - v5)
    v49 := GenS4{F0: 788, F1: true, F2: 2147483647}
    var v50 int16 = 32767
    switch "hello" {
    case "\n\t":
        var v51 int16 = v50
        _ = v51
    case "":
        v52 := GenS1{F0: -56, F1: false, F2: v50}
        _ = v52
    default:
        var v53 float64 = v43
        _ = v53
    }
    var v54 int64 = 42
    v55 := GenS5{F0: v1}
    for fi56 := 0; fi56 < 10; fi56++ {
        var v57 uint = 18446744073709551615
        _ = fi56
        _ = v57
    }
    v58 := GenS6{F0: v1, F1: 2, F2: v5}
    fmt.Println(v6, v43)
    var v59 uint = 65535
    var v60 int64 = v54
    fmt.Println(v6, v54)
    v3[32767] = false
    v1 += v6
    fmt.Println(v50)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v54:", v54)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v49
    _ = v5
    _ = v50
    _ = v54
    _ = v55
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
