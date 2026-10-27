package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint; F2 int64 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 uint = 2
    v2 := []uint32{}
    var v3 string = "hello"
    v4 := make(map[int]uint32)
    var v5 string = (v3 + (("\n\t" + (v3 + v3)) + v3))
    switch 520 {
    case -9:
        for fi6 := 0; fi6 < 20; fi6++ {
            switch (fi6 / ((fi6) | 1)) {
            case 0:
                if (false || false) {
                    v7 := GenS1{F0: 42}
                    _ = v7
                }
            case 2147483647:
                for fi8 := 0; fi8 < 16; fi8++ {
                    var v9 uint = (^(v1))
                    _ = fi8
                    _ = v9
                }
            }
            for fi10 := 0; fi10 < 14; fi10++ {
                v11 := GenS2{F0: 541, F1: false}
                v12 := GenS3{F0: 32767, F1: 729}
                _ = len(v4)
                fmt.Println(v5)
                switch fi10 {
                case -464:
                    var v13 uint8 = 20
                    _ = v13
                case -2147483648:
                    v1 -= v1
                }
                var v14 uint16 = 1
                var v15 bool = false
                _ = len(v4)
                _ = fi10
                _ = v11
                _ = v12
                _ = v14
                _ = v15
            }
            var v16 uint32 = 0
            _ = fi6
            _ = v16
        }
    case 32767:
        v3 = ("\n\t" + ((v3 + "hello") + (v3 + (("" + v5) + v5))))
    case -9223372036854775808:
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 bool = (0 < -9223372036854775808)
            v19 := GenS1{F0: 2}
            v20 := GenS1{F0: 0}
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
    default:
        delete(v4, -2147483648)
    }
    switch "hello" {
    case "a":
        delete(v4, 887)
    case "foo bar":
        if false {
            var v21 int16 = 1
            var v22 float32 = (3.14159 / 1.5)
            _ = v21
            _ = v22
        }
    }
    var v23 int = -487
    var v24 float64 = 0.0
    for fi25 := 0; fi25 < 3; fi25++ {
        switch v23 {
        case 32767:
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 string = ("" + ("hello" + "foo bar"))
                var v28 int32 = -2147483648
                _ = fi26
                _ = v27
                _ = v28
            }
        default:
            var v29 uint = (v1 + v1)
            _ = v29
        }
        var v30 float32 = 0.1
        var v31 uint8 = 128
        _ = fi25
        _ = v30
        _ = v31
    }
    switch (v23 | -1) {
    case -1:
        var v32 float64 = ((v24 / (v24 * (1e10 + 0.1))) - v24)
        _ = v32
    case 32768:
        var v33 bool = true
        _ = v33
    case 32767:
        v34 := GenS2{F0: 97, F1: (1.5 != float32((((v24 - 1e10) / v24) + v24)))}
        _ = v34
    }
    switch 4294967295 {
    case 127:
        for fi35 := 0; fi35 < 1; fi35++ {
            if (!false) {
                fmt.Println(v1, v23)
                fmt.Println(v23)
                switch "a" {
                case "foo bar":
                    var v36 int = v23
                    _ = v36
                case "":
                    v24 -= float64(((float32((0.0 - v24)) * 0.0) + -2.25))
                case "hello":
                    v24 += v24
                }
            }
            switch -766 {
            case 4294967295:
                fmt.Println(v3, v24, v24)
            case 2147483647:
                v5 = v3
            case 32767:
                v37 := v4[v23]
                _ = v37
            }
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            _ = fi35
        }
    case 128:
        fmt.Println(v23, v5, v24)
    case 276:
        if false {
            v24 -= (1.5 * v24)
            for fi38 := 0; fi38 < 17; fi38++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi38
            }
            for fi39 := 0; fi39 < 16; fi39++ {
                if len(v2) > 0 {
                    v2[0] = 0
                }
                v40 := GenS1{F0: 1}
                switch (^(v23)) {
                case -459:
                    switch (v23 & fi39) {
                    case 9223372036854775807:
                        fmt.Println(v24, fi39, v3)
                    case 42:
                        for fi41 := 0; fi41 < 10; fi41++ {
                            switch (^(fi41)) {
                            case 42:
                                fmt.Println(v3, v1, v23)
                                var v42 bool = true
                                _ = v42
                            case -296:
                                v43 := GenS4{F0: 613, F1: (v1 / (((^(v1))) | 1)), F2: 0}
                                _ = v43
                            case -2147483648:
                                v44 := GenS3{F0: 65535, F1: -1}
                                _ = v44
                            }
                            _ = fi41
                        }
                    default:
                        v24 += v24
                    }
                case 2147483647:
                    v2 = append(v2, 393)
                case 743:
                    var v45 uint8 = 2
                    _ = v45
                default:
                    v46 := GenS2{F0: 900, F1: false}
                    _ = v46
                }
                _ = fi39
                _ = v40
            }
        } else {
            v47 := GenS4{F0: 0, F1: (^(v1)), F2: 474}
            if (v24 == v24) {
                v1 -= v1
                if (true && (-1 < 1)) {
                    for fi48 := 0; fi48 < 8; fi48++ {
                        var v49 bool = (550 < 52)
                        v50 := GenS5{F0: int64(v1), F1: 42}
                        switch v5 {
                        case "foo bar":
                            v5 = "a"
                        case "":
                            switch v3 {
                            case "":
                                if v49 {
                                    var v51 uint32 = 2147483647
                                    _ = v51
                                }
                            case "foo bar":
                                v3 = ("" + "foo bar")
                            }
                        }
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                } else {
                    var v52 uint64 = 4294967295
                    var v53 uint64 = v52
                    v5 = v3
                    _ = v52
                    _ = v53
                }
            } else {
                for fi54 := 0; fi54 < 9; fi54++ {
                    var v55 float32 = float32(v24)
                    _ = fi54
                    _ = v55
                }
            }
            v56 := GenS6{F0: 2}
            _ = v47
            _ = v56
        }
    }
    for fi57 := 0; fi57 < 9; fi57++ {
        v3 = v3
        var v58 string = ("" + "hello")
        v59 := GenS4{F0: 804, F1: (^(v1)), F2: 1}
        _ = fi57
        _ = v58
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
