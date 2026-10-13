package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS6 struct { F0 bool }

func main() {
    v1 := []int{32767, -860, 2147483647}
    v2 := []string{"foo bar"}
    v3 := make(map[uint32]bool)
    switch "foo bar" {
    case "":
        switch 166 {
        case -734:
            var v4 int32 = -1
            _ = v4
        }
    }
    for fi5 := 0; fi5 < 14; fi5++ {
        if ((((("hello" + ("hello" + ("a" + "hello"))) + ("a" + "")) + "foo bar") + "foo bar") != "a") {
            var v6 int32 = 1
            var v7 int = (fi5 >> ((1) & 63))
            _ = v6
            _ = v7
        } else {
            if true {
                var v8 uint16 = 32767
                var v9 uint64 = 18446744073709551615
                if false {
                    for fi10 := 0; fi10 < 4; fi10++ {
                        v9 = v9
                        v8 = v8
                        _ = fi10
                    }
                    var v11 uint64 = 18446744073709551615
                    _ = v11
                } else {
                    var v12 uint = 216
                    v8 = 32767
                    _ = v12
                }
                _ = v8
                _ = v9
            } else {
                var v13 int32 = 24
                v13 += 976
                fmt.Println(v13)
                _ = v13
            }
            v14 := GenS1{F0: 255}
            var v15 float64 = -2.25
            _ = v14
            _ = v15
        }
        _ = fi5
    }
    v16 := GenS2{F0: ((!(!true)) && true)}
    var v17 float64 = 0.1
    if true {
        var v18 uint8 = 1
        var v19 int64 = 42
        var v20 int16 = 1
        _ = v18
        _ = v19
        _ = v20
    }
    v17 = (-2.25 - -2.25)
    _ = len(v3)
    fmt.Println(v17)
    if (("a" != "a") || true) {
        v21 := GenS3{F0: 2, F1: true}
        var v22 float32 = 1e10
        v17 += (0.0 * 0.1)
        _ = v21
        _ = v22
    } else {
        for fi23 := 0; fi23 < 1; fi23++ {
            v17 = float64(-2.25)
            v24 := GenS4{F0: 834, F1: -888}
            var v25 int32 = 0
            _ = fi23
            _ = v24
            _ = v25
        }
        switch 9223372036854775807 {
        case 0:
            v26 := GenS3{F0: 108, F1: ((true && true) || ((("a" + "") + "hello") != "hello"))}
            _ = v26
        default:
            if true {
                for fi27 := 0; fi27 < 17; fi27++ {
                    if false {
                        for fi28 := 0; fi28 < 15; fi28++ {
                            var v29 int16 = 1
                            v30 := GenS5{F0: 610, F1: 2147483648, F2: 296}
                            v29 = v29
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                        var v31 int32 = -383
                        _ = v31
                    } else {
                        switch fi27 {
                        case 127:
                            if (690 >= 42) {
                                for fi32 := 0; fi32 < 20; fi32++ {
                                    if (1 >= 947) {
                                        if len(v2) > 0 {
                                            v2[0] = ("hello" + "a")
                                        }
                                        fmt.Println(fi27, fi32, v17)
                                        var v33 uint = 65535
                                        _ = v33
                                    }
                                    if (false && true) {
                                        var v34 int32 = 0
                                        _ = v34
                                    }
                                    fmt.Println(fi27, fi27, v17)
                                    _ = fi32
                                }
                                v17 = (((v17 + 3.14159) / 1e10) / 1.0)
                            }
                        default:
                            for fi35 := 0; fi35 < 6; fi35++ {
                                v17 = 0.1
                                switch fi35 {
                                case -1:
                                    v17 = (v17 - v17)
                                }
                                var v36 int32 = 1
                                _ = fi35
                                _ = v36
                            }
                        }
                        var v37 float64 = v17
                        v17 += ((v37 - v17) + 0.1)
                        _ = v37
                    }
                    var v38 float64 = (float64(3.14159) * v17)
                    _ = fi27
                    _ = v38
                }
                var v39 int = 127
                v40 := GenS6{F0: (1 <= 255)}
                _ = v39
                _ = v40
            }
        }
    }
    v41 := GenS5{F0: 375, F1: 357, F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v41
}
