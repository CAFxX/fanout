package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 string; F1 int; F2 uint64 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64; F1 int32; F2 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 380
    v3 := []string{"hello", ("" + "a"), "a"}
    v1[0] = "foo bar"
    if (-9223372036854775808 >= 9223372036854775807) {
        v2 = 350
        v4 := GenS1{F0: ""}
        for fi5 := 0; fi5 < 8; fi5++ {
            v2 = 255
            _ = fi5
        }
        _ = v4
    } else {
        v2 += (v2 ^ 32767)
        v6 := GenS2{F0: 42, F1: 1}
        var v7 int64 = 9223372036854775807
        _ = v6
        _ = v7
    }
    if (true && false) {
        if ((!false) && false) {
            var v8 uint = 0
            for fi9 := 0; fi9 < 6; fi9++ {
                for fi10 := 0; fi10 < 6; fi10++ {
                    var v11 float32 = float32(1.5)
                    v2 += 65535
                    var v12 int64 = 0
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v13 uint64 = 301
                var v14 int8 = 100
                _ = fi9
                _ = v13
                _ = v14
            }
            _ = v8
        } else {
            v2 -= v2
            var v15 uint16 = 462
            _ = v15
        }
    } else {
        var v16 float64 = ((1.5 + 1e10) / float64(0.1))
        _ = v16
    }
    v17 := GenS3{F0: "\n\t", F1: 32768, F2: 1}
    if ("" > "hello") {
        v3 = append(v3, ("" + ("hello" + ("foo bar" + "\n\t"))))
        v2 = v2
        var v18 string = ("a" + "\n\t")
        _ = v18
    } else {
        var v19 uint64 = 1
        _ = v19
    }
    var v20 uint32 = 4294967295
    if false {
        if (1e10 <= float32(0.0)) {
            fmt.Println(v2)
            if true {
                var v21 int = -674
                if false {
                    var v22 int8 = 1
                    switch -9223372036854775808 {
                    case 32768:
                        var v23 int64 = 0
                        _ = v23
                    case -91:
                        v20 = v20
                    case 937:
                        delete(v1, (v20 ^ v20))
                    }
                    switch "a" {
                    case "hello":
                        for fi24 := 0; fi24 < 20; fi24++ {
                            v3 = append(v3, (("hello" + "foo bar") + "a"))
                            switch ("foo bar" + ("hello" + "\n\t")) {
                            case "":
                                var v25 int16 = 584
                                _ = v25
                            default:
                                for fi26 := 0; fi26 < 9; fi26++ {
                                    v22 = 0
                                    v27 := GenS4{F0: -1, F1: (v2 == (^(v2)))}
                                    _ = fi26
                                    _ = v27
                                }
                            }
                            fmt.Println(v2, v20, v20)
                            _ = fi24
                        }
                    }
                    _ = v22
                }
                _ = v21
            } else {
                var v28 int8 = -9
                if (!false) {
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 float32 = 0.0
                        _ = fi29
                        _ = v30
                    }
                    v28 = (-(v28))
                }
                _ = v28
            }
            var v31 uint16 = v2
            _ = v31
        }
        v32 := GenS2{F0: -9, F1: 2147483647}
        _ = v32
    } else {
        var v33 float32 = 0.1
        v34 := GenS1{F0: "\n\t"}
        _ = v33
        _ = v34
    }
    var v35 int16 = 1
    _ = len(v1)
    v35 -= (-(v35))
    if true {
        fmt.Println(v2, v20, v2)
        v35 = (v35 * v35)
    } else {
        if false {
            var v36 uint16 = (v2 & (v2 & (v2 * v2)))
            _ = v36
        }
        var v37 int16 = 1000
        switch 127 {
        case 9223372036854775807:
            fmt.Println(v35, v37)
        case 32768:
            v38 := GenS5{F0: v20}
            _ = v38
        case 4294967295:
            v39 := GenS2{F0: -1, F1: 42}
            _ = v39
        default:
            var v40 int8 = -128
            _ = v40
        }
        _ = v37
    }
    switch 700 {
    case 1:
        switch -296 {
        case 32768:
            switch -1 {
            case -2147483648:
                switch 1 {
                case 42:
                    v2 += 65535
                }
            case -9223372036854775808:
                v41 := GenS1{F0: "hello"}
                _ = v41
            case 4294967295:
                var v42 int16 = v35
                _ = v42
            }
        case 42:
            var v43 float64 = 1e10
            _ = v43
        default:
            for fi44 := 0; fi44 < 2; fi44++ {
                var v45 int16 = (v35 ^ (v35 + (-(v35))))
                for fi46 := 0; fi46 < 12; fi46++ {
                    v47 := GenS1{F0: "hello"}
                    fmt.Println(fi44, fi46, fi44)
                    _ = fi46
                    _ = v47
                }
                _ = fi44
                _ = v45
            }
        }
    case 32767:
        v2 = v2
    }
    if false {
        if (18446744073709551615 >= 255) {
            var v48 int64 = -427
            for fi49 := 0; fi49 < 7; fi49++ {
                for fi50 := 0; fi50 < 4; fi50++ {
                    v20 += 0
                    _ = fi50
                }
                var v51 float64 = -2.25
                v48 = v48
                _ = fi49
                _ = v51
            }
            _ = v48
        } else {
            var v52 int64 = 1
            _ = v52
        }
        v1[2] = ""
    } else {
        fmt.Println(v20, v2)
        var v53 int8 = 42
        var v54 uint = 18446744073709551615
        _ = v53
        _ = v54
    }
    v2 += v2
    var v55 uint = 255
    v56 := GenS6{F0: -1, F1: 42, F2: 32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v55
    _ = v56
}
