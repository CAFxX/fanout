package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS4 struct { F0 int32; F1 bool; F2 int8 }
type GenS5 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]int)
    var v3 float64 = 0.1
    var v4 int64 = -1
    var v5 string = (("" + "foo bar") + "\n\t")
    fmt.Println(v3, v3)
    if false {
        if true {
            var v6 uint16 = 0
            _ = v6
        } else {
            v7 := GenS1{F0: 1, F1: 274}
            v8 := GenS2{F0: int(v4), F1: 338}
            _ = v7
            _ = v8
        }
        if false {
            switch (v5 + v5) {
            case "":
                switch v5 {
                case "\n\t":
                    for fi9 := 0; fi9 < 19; fi9++ {
                        var v10 int8 = 127
                        switch v5 {
                        case "\n\t":
                            switch ("" + ((v5 + v5) + v5)) {
                            case "a":
                                var v11 int32 = 2147483647
                                _ = v11
                            case "\n\t":
                                var v12 float64 = 0.0
                                _ = v12
                            }
                        case "foo bar":
                            if true {
                                if (!((0.0 != ((((1e10 + float32((float64(((0.0 + float32(v3)) - 1e10)) + v3))) + -2.25) * 0.1) - (float32(0.1) * 1e10))) && true)) {
                                    fmt.Println(fi9, v3, v4)
                                }
                                v13 := GenS1{F0: -1, F1: -1}
                                _ = v13
                            } else {
                                var v14 int8 = v10
                                var v15 int16 = -32768
                                fmt.Println(v15, v14, v4)
                                _ = v14
                                _ = v15
                            }
                        case "hello":
                            v2[2] = (^(fi9))
                        default:
                            if (2 <= 979) {
                                fmt.Println(v10)
                                fmt.Println(v4)
                                if true {
                                    delete(v2, 42)
                                }
                            }
                        }
                        fmt.Println(v10)
                        _ = fi9
                        _ = v10
                    }
                case "foo bar":
                    v2[995] = 32767
                case "hello":
                    for fi16 := 0; fi16 < 20; fi16++ {
                        var v17 float64 = (v3 * (float64(0.0) + 1e10))
                        _ = fi16
                        _ = v17
                    }
                }
            }
            v18 := GenS3{F0: 12, F1: 65535, F2: 1}
            switch v5 {
            case "foo bar":
                fmt.Println(v5)
            case "":
                var v19 string = v5
                _ = v19
            default:
                var v20 int8 = 93
                _ = v20
            }
            _ = v18
        } else {
            var v21 int8 = 127
            v5 = (v5 + v5)
            _ = v21
        }
    }
    fmt.Println(v5, v5)
    switch v5 {
    case "a":
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := GenS3{F0: 127, F1: 65535, F2: 2}
            if true {
                v4 = -445
                fmt.Println(v5)
                var v24 string = v5
                _ = v24
            } else {
                if false {
                    v3 += 0.1
                    if (42 == 612) {
                        if false {
                            var v25 int64 = (v4 << ((237) & 63))
                            _ = v25
                        }
                    }
                }
                if true {
                    for fi26 := 0; fi26 < 18; fi26++ {
                        v1 = append(v1, v5)
                        _ = fi26
                    }
                    v27 := v2[1]
                    _ = v27
                } else {
                    var v28 int16 = -1
                    v3 -= (((v3 / 1.5) + v3) * v3)
                    var v29 uint16 = 735
                    _ = v28
                    _ = v29
                }
                v30 := GenS4{F0: -1, F1: (false && false), F2: 4}
                _ = v30
            }
            _ = fi22
            _ = v23
        }
    default:
        fmt.Println(v3)
    }
    v3 = float64(1.5)
    if true {
        switch v5 {
        case "hello":
            v31 := GenS4{F0: -1, F1: (((!(false || false)) || (!(false && false))) || (true && false)), F2: 1}
            _ = v31
        case "":
            v32 := GenS1{F0: 0, F1: 1000}
            _ = v32
        case "a":
            var v33 int32 = -1
            _ = v33
        default:
            var v34 float64 = v3
            _ = v34
        }
        v35 := v2[705]
        _ = v35
    }
    v36 := GenS5{F0: 255, F1: 0}
    for fi37 := 0; fi37 < 20; fi37++ {
        var v38 int = 2147483647
        v1 = append(v1, v5)
        _ = fi37
        _ = v38
    }
    v39 := GenS3{F0: 255, F1: 80, F2: 42}
    var v40 int64 = (v4 - (v4 - v4))
    var v41 float64 = (v3 - v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
