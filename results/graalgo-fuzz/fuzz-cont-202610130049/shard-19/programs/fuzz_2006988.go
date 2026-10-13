package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 string }

func main() {
    var v1 float64 = 0.1
    v2 := []int{}
    v3 := make(map[uint32]string)
    var v4 float32 = 0.1
    var v5 uint64 = 785
    var v6 string = ("foo bar" + "a")
    switch v6 {
    case "foo bar":
        if (42 > 255) {
            if false {
                switch 2147483647 {
                case 32768:
                    for fi7 := 0; fi7 < 13; fi7++ {
                        v8 := GenS1{F0: (-562 > 42), F1: 1}
                        _ = len(v3)
                        _ = fi7
                        _ = v8
                    }
                case 59:
                    var v9 uint16 = 410
                    _ = v9
                case 608:
                    var v10 uint32 = 148
                    _ = v10
                default:
                    v6 = "\n\t"
                }
                var v11 int64 = 42
                _ = v11
            } else {
                v12 := GenS1{F0: true, F1: 42}
                _ = v12
            }
            delete(v3, 635)
            for fi13 := 0; fi13 < 19; fi13++ {
                _ = len(v3)
                _ = len(v3)
                for fi14 := 0; fi14 < 6; fi14++ {
                    _ = len(v2)
                    _ = fi14
                }
                _ = fi13
            }
        } else {
            v15 := GenS2{F0: 2147483647}
            _ = v15
        }
    case "":
        v16 := GenS2{F0: 2147483647}
        _ = v16
    case "hello":
        v17 := GenS1{F0: (("" + "a") >= "a"), F1: 568}
        _ = v17
    }
    v18 := GenS2{F0: 0}
    v4 = ((((((1e10 - float32(v1)) - float32(0.1)) / 3.14159) * 0.0) - (0.0 + (v4 / v4))) / -2.25)
    var v19 int16 = 0
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    var v20 bool = (1 != 0)
    switch 270 {
    case 9223372036854775807:
        if v20 {
            switch ("a" + "") {
            case "hello":
                var v21 uint64 = v5
                _ = v21
            case "foo bar":
                fmt.Println(v6, v6, v6)
            default:
                if v20 {
                    if v20 {
                        var v22 float64 = v1
                        for fi23 := 0; fi23 < 1; fi23++ {
                            for fi24 := 0; fi24 < 6; fi24++ {
                                if true {
                                    v19 -= (v19 >> ((783) & 63))
                                    if (v20 || true) {
                                        var v25 uint16 = 682
                                        var v26 float64 = v22
                                        var v27 int = (-(fi24))
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    } else {
                                        fmt.Println(v4, fi23)
                                        var v28 float32 = float32(v1)
                                        _ = v28
                                    }
                                    if v20 {
                                        var v29 float64 = v1
                                        var v30 uint = 255
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v31 := GenS2{F0: -2147483648}
                                        _ = v31
                                    }
                                }
                                _ = fi24
                            }
                            fmt.Println(v4, v5, v6)
                            _ = fi23
                        }
                        _ = v22
                    } else {
                        v4 -= (float32(v1) * v4)
                        if false {
                            v32 := GenS3{F0: v20}
                            _ = v32
                        }
                    }
                    v33 := GenS4{F0: ("" + v6), F1: 0}
                    _ = v33
                }
            }
            var v34 int16 = v19
            _ = v34
        } else {
            switch "hello" {
            case "hello":
                if len(v2) > 0 {
                    v2[0] = 302
                }
            case "a":
                var v35 int32 = 583
                _ = v35
            }
        }
    case 2147483647:
        for fi36 := 0; fi36 < 15; fi36++ {
            var v37 uint64 = v5
            _ = fi36
            _ = v37
        }
    }
    switch v6 {
    case "hello":
        fmt.Println(v5, v4)
    }
    for fi38 := 0; fi38 < 3; fi38++ {
        v6 = v6
        _ = fi38
    }
    if (42 > 512) {
        if true {
            for fi39 := 0; fi39 < 4; fi39++ {
                switch v6 {
                case "a":
                    v40 := GenS5{F0: v6}
                    _ = v40
                case "foo bar":
                    for fi41 := 0; fi41 < 10; fi41++ {
                        v20 = (!((v4 - 3.14159) < v4))
                        v42 := GenS2{F0: 0}
                        _ = fi41
                        _ = v42
                    }
                default:
                    v43 := GenS5{F0: "\n\t"}
                    _ = v43
                }
                _ = fi39
            }
        } else {
            var v44 uint8 = 109
            _ = len(v3)
            _ = v44
        }
    } else {
        v4 -= -2.25
        v45 := GenS1{F0: (v20 || (!false)), F1: 4294967295}
        _ = v45
    }
    delete(v3, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
