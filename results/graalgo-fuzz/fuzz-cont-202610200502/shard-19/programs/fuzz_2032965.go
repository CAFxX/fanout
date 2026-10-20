package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint64; F1 int16; F2 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 1
    var v3 string = "hello"
    var v4 int64 = 127
    var v5 uint16 = v2
    if true {
        fmt.Println(v4, v4)
        v6 := GenS1{F0: -123}
        _ = v6
    }
    switch v3 {
    case "\n\t":
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 bool = true
            _ = fi7
            _ = v8
        }
    case "":
        var v9 int16 = -6
        _ = v9
    }
    v10 := GenS1{F0: 126}
    fmt.Println(v4)
    v11 := []int{}
    var v12 int32 = -1
    if true {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 uint64 = 0
            var v15 int32 = v12
            _ = fi13
            _ = v14
            _ = v15
        }
    } else {
        var v16 float32 = float32((0.0 / (float64(0.0) + float64(1e10))))
        _ = v16
    }
    fmt.Println(v4, v3)
    switch -440 {
    case 295:
        fmt.Println(v5, v5, v12)
    default:
        v11 = append(v11, -1)
    }
    var v17 uint = 65535
    if false {
        if false {
            v5 = 340
            v18 := v1[337]
            _ = v18
        }
        var v19 float64 = ((1.5 * 3.14159) - 0.0)
        v3 = (v3 + ((v3 + v3) + v3))
        _ = v19
    } else {
        var v20 int8 = -128
        if false {
            switch "" {
            case "\n\t":
                _ = len(v11)
            }
            fmt.Println(v2, v2)
            if (v4 > (v4 % (((-(v4))) | 1))) {
                if false {
                    var v21 uint64 = 199
                    _ = v21
                } else {
                    if true {
                        var v22 string = v3
                        var v23 int16 = 42
                        _ = v22
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 19; fi24++ {
                        for fi25 := 0; fi25 < 7; fi25++ {
                            var v26 int16 = 32767
                            _ = fi25
                            _ = v26
                        }
                        _ = fi24
                    }
                    switch 32768 {
                    case 127:
                        switch v3 {
                        case "\n\t":
                            v2 -= (v2 & (v2 >> ((4294967295) & 63)))
                        case "a":
                            v11 = append(v11, 9223372036854775807)
                        case "":
                            v27 := GenS2{F0: ""}
                            _ = v27
                        }
                    case 4294967295:
                        var v28 uint8 = 42
                        _ = v28
                    case 128:
                        switch 9223372036854775807 {
                        case 42:
                            v2 += 0
                        case 280:
                            switch 923 {
                            case 32767:
                                v29 := GenS2{F0: ("\n\t" + v3)}
                                _ = v29
                            default:
                                v12 += v12
                            }
                        case 351:
                            var v30 uint = (v17 ^ v17)
                            _ = v30
                        }
                    }
                }
            } else {
                fmt.Println(v17)
            }
        } else {
            v31 := GenS3{F0: v4, F1: 42}
            _ = v31
        }
        _ = v20
    }
    var v32 uint16 = (^(v2))
    v11 = append(v11, -252)
    var v33 uint16 = (^(v32))
    if false {
        v34 := GenS2{F0: ("a" + ((v3 + v3) + (v3 + ("\n\t" + "\n\t"))))}
        _ = v34
    } else {
        fmt.Println(v33, v3, v33)
        if ((v12 & v12) != v12) {
            switch v3 {
            case "\n\t":
                v17 = (^(v17))
            case "hello":
                if false {
                    if true {
                        if false {
                            var v35 uint8 = 1
                            v36 := GenS4{F0: 4294967295, F1: -437, F2: 2}
                            _ = v35
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 11; fi37++ {
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v12)
                            case "a":
                                v3 = "\n\t"
                            case "hello":
                                v5 = (v2 / ((65535) | 1))
                            default:
                                v38 := v1[2]
                                _ = v38
                            }
                            var v39 uint = v17
                            _ = fi37
                            _ = v39
                        }
                        v32 = v2
                    }
                } else {
                    fmt.Println(v5)
                }
            case "a":
                if false {
                    if ((-1 <= 0) && false) {
                        fmt.Println(v17, v17)
                        switch "a" {
                        case "foo bar":
                            fmt.Println(v33, v33)
                        case "":
                            fmt.Println(v32)
                        default:
                            v11 = append(v11, 61)
                        }
                    } else {
                        for fi40 := 0; fi40 < 13; fi40++ {
                            if (v12 < v12) {
                                v41 := GenS3{F0: (v4 % (((^(v4))) | 1)), F1: -375}
                                _ = v41
                            }
                            var v42 uint = v17
                            var v43 int8 = 100
                            _ = fi40
                            _ = v42
                            _ = v43
                        }
                    }
                    v44 := GenS1{F0: 22}
                    _ = v44
                }
            }
            for fi45 := 0; fi45 < 6; fi45++ {
                v46 := GenS1{F0: 19}
                _ = fi45
                _ = v46
            }
            v4 -= v4
        }
    }
    fmt.Println(v12, v4, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
