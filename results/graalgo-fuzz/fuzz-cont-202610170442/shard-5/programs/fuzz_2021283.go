package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    v4 := make(map[int]uint32)
    v5 := make(map[int]uint32)
    if false {
        var v6 int16 = 1000
        _ = v6
    }
    if true {
        var v7 float64 = (float64((((float32((1.5 - -2.25)) - 1.5) * 3.14159) * 1.5)) + -2.25)
        v8 := GenS1{F0: 100, F1: 61, F2: 128}
        if false {
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS1{F0: 0, F1: 43, F2: 149}
                _ = fi9
                _ = v10
            }
            switch "" {
            case "foo bar":
                for fi11 := 0; fi11 < 6; fi11++ {
                    if false {
                        for fi12 := 0; fi12 < 14; fi12++ {
                            delete(v4, fi11)
                            fmt.Println(fi12, fi12)
                            _ = fi12
                        }
                        var v13 uint64 = 42
                        _ = v13
                    } else {
                        if len(v2) > 0 {
                            v2[0] = 437
                        }
                        v7 = v7
                        if true {
                            for fi14 := 0; fi14 < 11; fi14++ {
                                var v15 float32 = 0.1
                                v16 := GenS1{F0: -103, F1: 127, F2: 77}
                                var v17 uint = 65535
                                _ = fi14
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            v7 = (1.5 * (v7 * (1.5 * (1.5 + v7))))
                        } else {
                            var v18 string = "hello"
                            for fi19 := 0; fi19 < 1; fi19++ {
                                var v20 string = (v18 + ("foo bar" + (v18 + "foo bar")))
                                v21 := GenS2{F0: false}
                                v22 := GenS3{F0: 1, F1: 710}
                                _ = fi19
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            _ = len(v2)
                            _ = v18
                        }
                    }
                    _ = fi11
                }
            case "\n\t":
                v23 := GenS1{F0: -128, F1: 127, F2: 122}
                _ = v23
            case "":
                v24 := GenS4{F0: 4294967295}
                _ = v24
            }
            v7 -= v7
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 uint = 255
                _ = fi25
                _ = v26
            }
            var v27 uint32 = 106
            _ = v27
        }
        _ = v7
        _ = v8
    }
    switch -9223372036854775808 {
    case 2147483647:
        var v28 int8 = 1
        _ = v28
    case -1:
        var v29 uint64 = 42
        _ = v29
    }
    var v30 int32 = -822
    if true {
        if false {
            _ = len(v1)
            v30 += (^(v30))
        }
        fmt.Println(v30)
    } else {
        if len(v2) > 0 {
            v2[0] = 838
        }
        v31 := GenS5{F0: 1, F1: 2}
        v30 = (v30 - 42)
        _ = v31
    }
    if false {
        fmt.Println(v30)
    } else {
        _ = len(v2)
    }
    var v32 int16 = 32767
    var v33 int16 = (v32 + 149)
    for fi34 := 0; fi34 < 8; fi34++ {
        v35 := v1[65535]
        _ = fi34
        _ = v35
    }
    v36 := v5[-1]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
}
