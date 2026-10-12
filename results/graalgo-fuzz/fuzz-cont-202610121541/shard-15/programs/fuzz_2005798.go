package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 uint16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"foo bar", "hello", "\n\t", "hello"}
    var v3 int8 = 0
    switch "" {
    case "foo bar":
        v4 := GenS1{F0: true, F1: 1}
        _ = v4
    case "\n\t":
        var v5 int64 = 0
        _ = v5
    }
    var v6 uint64 = 42
    var v7 uint16 = 2
    if (false || (2147483647 < -2147483648)) {
        v3 = 1
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := v1["hello"]
            v10 := GenS2{F0: 189}
            if false {
                if ((-(fi8)) >= fi8) {
                    v11 := GenS1{F0: (int64(fi8) >= 9223372036854775807), F1: 1}
                    if false {
                        for fi12 := 0; fi12 < 20; fi12++ {
                            if false {
                                _ = len(v1)
                            } else {
                                var v13 int8 = 127
                                fmt.Println(v7)
                                var v14 float32 = 1e10
                                _ = v13
                                _ = v14
                            }
                            _ = fi12
                        }
                        _ = len(v2)
                    }
                    _ = v11
                } else {
                    var v15 uint32 = 790
                    switch fi8 {
                    case 9223372036854775807:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = v15
                }
                v1["hello"] = (v9 + 0)
            } else {
                v7 -= uint16(v6)
                v6 += (^(v6))
            }
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        var v16 int16 = -2
        _ = v16
    }
    if (217 >= 42) {
        for fi17 := 0; fi17 < 18; fi17++ {
            v3 = (v3 ^ 127)
            _ = fi17
        }
        delete(v1, "foo bar")
        v2 = append(v2, "hello")
    } else {
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 uint8 = 255
            _ = fi18
            _ = v19
        }
    }
    if false {
        switch "\n\t" {
        case "a":
            fmt.Println(v6)
        }
    } else {
        v20 := GenS2{F0: 12}
        var v21 uint = 42
        fmt.Println(v21)
        _ = v20
        _ = v21
    }
    switch "a" {
    case "hello":
        for fi22 := 0; fi22 < 2; fi22++ {
            if (704 >= -1) {
                for fi23 := 0; fi23 < 9; fi23++ {
                    v6 -= (v6 & 42)
                    v7 = 32767
                    _ = fi23
                }
            }
            switch (^(fi22)) {
            case 1:
                var v24 int64 = -778
                _ = v24
            default:
                var v25 int16 = -645
                _ = v25
            }
            _ = fi22
        }
    default:
        if true {
            v26 := GenS2{F0: 127}
            v3 = -128
            var v27 float64 = 3.14159
            _ = v26
            _ = v27
        } else {
            var v28 uint = 65535
            switch -713 {
            case 127:
                fmt.Println(v28, v7, v28)
            case -328:
                v29 := GenS1{F0: false, F1: -1}
                _ = v29
            case 4294967295:
                var v30 int8 = v3
                _ = v30
            default:
                v2 = append(v2, "a")
            }
            _ = v28
        }
    }
    if true {
        var v31 int16 = 1000
        v32 := GenS2{F0: 156}
        if (-32768 != (v31 - v31)) {
            var v33 uint16 = v7
            _ = v33
        }
        _ = v31
        _ = v32
    } else {
        v34 := GenS3{F0: 2, F1: 0}
        for fi35 := 0; fi35 < 13; fi35++ {
            v6 = v6
            _ = fi35
        }
        v6 = v6
        _ = v34
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        fmt.Println(v7)
        v6 = v6
        if true {
            for fi37 := 0; fi37 < 14; fi37++ {
                switch "\n\t" {
                case "a":
                    switch "foo bar" {
                    case "\n\t":
                        var v38 uint = 0
                        _ = len(v2)
                        _ = v38
                    case "a":
                        var v39 float64 = -2.25
                        _ = v39
                    case "foo bar":
                        v7 += v7
                    }
                case "\n\t":
                    v6 -= v6
                case "foo bar":
                    if (0 < 280) {
                        var v40 string = ("" + "foo bar")
                        _ = v40
                    } else {
                        v2 = append(v2, "foo bar")
                        v41 := GenS4{F0: false}
                        _ = v41
                    }
                default:
                    if true {
                        var v42 uint8 = 127
                        var v43 bool = (!false)
                        _ = v42
                        _ = v43
                    } else {
                        var v44 int16 = -142
                        var v45 int32 = 42
                        _ = v44
                        _ = v45
                    }
                }
                v7 = v7
                if ("a" == "foo bar") {
                    if true {
                        v46 := GenS5{F0: 42, F1: (v7 | v7)}
                        if len(v2) > 0 {
                            v2[0] = ("a" + "\n\t")
                        }
                        var v47 int8 = v3
                        _ = v46
                        _ = v47
                    }
                }
                _ = fi37
            }
            var v48 float32 = 0.0
            _ = v48
        } else {
            _ = len(v1)
            var v49 uint16 = (v7 & v7)
            var v50 uint8 = 128
            _ = v49
            _ = v50
        }
        _ = fi36
    }
    if true {
        _ = len(v2)
    } else {
        var v51 int32 = 647
        var v52 uint64 = v6
        _ = v51
        _ = v52
    }
    switch 32768 {
    case 9223372036854775807:
        for fi53 := 0; fi53 < 14; fi53++ {
            for fi54 := 0; fi54 < 10; fi54++ {
                _ = len(v2)
                v55 := GenS5{F0: 4294967295, F1: 0}
                var v56 string = ((((("foo bar" + "hello") + "foo bar") + "hello") + "a") + "hello")
                _ = fi54
                _ = v55
                _ = v56
            }
            fmt.Println(v7)
            v57 := GenS3{F0: 255, F1: 888}
            _ = fi53
            _ = v57
        }
    default:
        var v58 bool = true
        _ = v58
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
