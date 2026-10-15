package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint }

func main() {
    var v1 uint8 = 92
    v2 := make(map[string]string)
    var v3 uint16 = 755
    v4 := []uint32{2, 4294967295, 1, 1}
    v5 := []uint32{4294967295, 65535, 1, 88}
    v3 -= (^(v3))
    if (2 != 65535) {
        if false {
            switch "\n\t" {
            case "\n\t":
                v6 := GenS1{F0: 127, F1: 23, F2: 4294967295}
                _ = v6
            default:
                v1 = v1
            }
            for fi7 := 0; fi7 < 11; fi7++ {
                for fi8 := 0; fi8 < 11; fi8++ {
                    var v9 uint16 = 727
                    _ = fi8
                    _ = v9
                }
                var v10 uint16 = v3
                _ = fi7
                _ = v10
            }
        } else {
            fmt.Println(v3, v3)
            fmt.Println(v3, v1)
            var v11 int32 = 0
            _ = v11
        }
        var v12 uint32 = 2147483647
        _ = v12
    }
    v1 = uint8(v3)
    v1 += 250
    var v13 float32 = ((-2.25 + float32(1e10)) * 0.1)
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int16 = 32767
        _ = fi14
        _ = v15
    }
    switch "a" {
    case "a":
        switch 2147483647 {
        case 42:
            var v16 int32 = 42
            _ = v16
        case 1:
            v3 = (^(v3))
        case 2147483647:
            switch -451 {
            case 127:
                for fi17 := 0; fi17 < 20; fi17++ {
                    fmt.Println(fi17, v3, fi17)
                    switch "a" {
                    case "":
                        var v18 float64 = 0.0
                        _ = v18
                    case "foo bar":
                        for fi19 := 0; fi19 < 16; fi19++ {
                            if ((!true) || false) {
                                fmt.Println(v13)
                                var v20 int64 = 0
                                _ = v20
                            } else {
                                var v21 float32 = 0.1
                                _ = v21
                            }
                            v22 := GenS1{F0: 0, F1: 127, F2: 552}
                            _ = fi19
                            _ = v22
                        }
                    case "hello":
                        var v23 uint64 = 42
                        _ = v23
                    }
                    _ = fi17
                }
            case 1:
                var v24 uint32 = 337
                _ = v24
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        default:
            var v25 uint16 = 160
            _ = v25
        }
    case "foo bar":
        var v26 int8 = 1
        _ = v26
    }
    var v27 uint = 1
    var v28 float32 = v13
    v29 := GenS1{F0: 100, F1: -116, F2: 18446744073709551615}
    var v30 uint8 = v1
    v31 := GenS2{F0: 32767}
    v30 = (^(v1))
    for fi32 := 0; fi32 < 15; fi32++ {
        v33 := GenS1{F0: -1, F1: int8(v27), F2: 42}
        switch "" {
        case "foo bar":
            v34 := v2[("" + (("a" + ("" + "")) + "\n\t"))]
            _ = v34
        case "a":
            switch fi32 {
            case -626:
                var v35 uint64 = 240
                _ = v35
            }
        case "\n\t":
            if (-361 >= 1000) {
                _ = len(v2)
                if true {
                    for fi36 := 0; fi36 < 14; fi36++ {
                        if (-762 > -1) {
                            if ((("hello" + "foo bar") <= ("a" + "a")) && false) {
                                var v37 uint64 = 42
                                _ = v37
                            } else {
                                v28 = (v13 * v13)
                                if (false && (661 < 42)) {
                                    v38 := v2[""]
                                    _ = v38
                                }
                                switch (("a" + ("a" + "\n\t")) + "hello") {
                                case "foo bar":
                                    v13 -= v13
                                case "a":
                                    var v39 int32 = -793
                                    _ = v39
                                default:
                                    if false {
                                        v40 := GenS1{F0: 85, F1: 42, F2: 566}
                                        _ = v40
                                    }
                                }
                            }
                            v41 := GenS3{F0: 1}
                            delete(v2, ("a" + (("foo bar" + "hello") + "a")))
                            _ = v41
                        } else {
                            if false {
                                v42 := GenS4{F0: 42, F1: (^(v27))}
                                v27 += (v27 / ((v27) | 1))
                                var v43 float64 = 1e10
                                _ = v42
                                _ = v43
                            }
                            var v44 int = (fi32 >> (((v27 ^ (v27 & (v27 << (((v27 % (((v27 | (v27 - (v27 + (v27 & v27))))) | 1))) & 63))))) & 63))
                            _ = v44
                        }
                        v45 := GenS3{F0: 1}
                        _ = fi36
                        _ = v45
                    }
                    v27 = (v27 * (v27 ^ v27))
                    for fi46 := 0; fi46 < 6; fi46++ {
                        var v47 int = -2147483648
                        _ = fi46
                        _ = v47
                    }
                }
            }
        default:
            var v48 int64 = 126
            _ = v48
        }
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
