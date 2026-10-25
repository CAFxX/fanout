package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 string }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint16; F1 uint16; F2 string }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 bool; F1 uint; F2 uint64 }
func genFn0(p0 int32) int32 {
    v4 := GenS1{F0: p0, F1: 255, F2: ("hello" + ("hello" + "foo bar"))}
    var v5 float32 = 0.0
    fmt.Println(p0, v5)
    switch "\n\t" {
    case "hello":
        var v6 float32 = 3.14159
        _ = v6
    }
    _ = p0
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 string = ""
    v2 := make(map[int]uint32)
    var v3 uint64 = 0
    v7 := []int{-538, 9223372036854775807, 127, 846}
    v2[4294967295] = 4294967295
    if true {
        var v8 float32 = float32(float64((1e10 / float32((-2.25 - (0.1 + 0.0))))))
        var v9 uint = uint(v3)
        _ = v8
        _ = v9
    } else {
        if false {
            v3 = v3
            switch v1 {
            case "a":
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case "":
                _ = len(v7)
            case "foo bar":
                v10 := GenS1{F0: 0, F1: 4294967295, F2: v1}
                _ = v10
            }
        }
        if false {
            v11 := GenS1{F0: -2147483648, F1: 4294967295, F2: "foo bar"}
            _ = v11
        }
    }
    _ = len(v7)
    var v12 int = 9223372036854775807
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint8 = 1
        _ = fi13
        _ = v14
    }
    v3 += (v3 + v3)
    var v15 uint64 = 2
    switch v1 {
    case "foo bar":
        v16 := GenS2{F0: v15, F1: v1}
        _ = v16
    }
    _ = len(v2)
    var v17 uint8 = 2
    if true {
        if true {
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 string = v1
                switch (v12 * v12) {
                case -971:
                    v20 := GenS2{F0: v3, F1: v19}
                    _ = v20
                case -79:
                    delete(v2, fi18)
                }
                if ((-2.25 != 0.0) || (true || (4294967295 >= 4294967295))) {
                    var v21 float64 = 1e10
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 10; fi22++ {
                        var v23 uint16 = 0
                        v24 := GenS3{F0: (v23 * (v23 * v23)), F1: v23, F2: "foo bar"}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS4{F0: 763, F1: (v15 + 42)}
                    _ = v25
                }
                _ = fi18
                _ = v19
            }
            if false {
                var v26 uint32 = 2147483647
                v3 -= (v15 * v3)
                _ = len(v2)
                _ = v26
            } else {
                v17 += v17
            }
            switch v12 {
            case 32768:
                if false {
                    var v27 uint16 = 2
                    v28 := GenS5{F0: true, F1: 4294967295, F2: 1}
                    _ = v27
                    _ = v28
                } else {
                    var v29 int8 = -1
                    v12 = v12
                    _ = v29
                }
            default:
                for fi30 := 0; fi30 < 12; fi30++ {
                    switch v1 {
                    case "\n\t":
                        var v31 uint = 2
                        _ = v31
                    case "":
                        if false {
                            var v32 float64 = 0.1
                            if (int64(v17) < -9223372036854775808) {
                                var v33 float32 = 1e10
                                var v34 uint32 = 488
                                v7 = append(v7, v12)
                                _ = v33
                                _ = v34
                            } else {
                                var v35 int64 = -9223372036854775808
                                var v36 uint16 = 0
                                v37 := GenS2{F0: v3, F1: ("\n\t" + v1)}
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                            _ = v32
                        } else {
                            for fi38 := 0; fi38 < 13; fi38++ {
                                v39 := GenS1{F0: 1, F1: 892, F2: ("" + v1)}
                                fmt.Println(v15, fi30, fi30)
                                _ = fi38
                                _ = v39
                            }
                            if false {
                                fmt.Println(fi30, v1, v3)
                                var v40 string = (v1 + v1)
                                for fi41 := 0; fi41 < 10; fi41++ {
                                    var v42 float32 = ((((float32((0.0 - 1e10)) * 3.14159) / ((0.1 / 1.0) - 1e10)) * 1e10) + 1.5)
                                    fmt.Println(fi30, v15, v15)
                                    _ = fi41
                                    _ = v42
                                }
                                _ = v40
                            } else {
                                v17 = v17
                            }
                        }
                    default:
                        switch v1 {
                        case "\n\t":
                            switch v1 {
                            case "hello":
                                v12 = (v12 / ((32767) | 1))
                            case "foo bar":
                                if false {
                                    switch v1 {
                                    case "a":
                                        if true {
                                            switch (v12 * (fi30 & (v12 | (fi30 + fi30)))) {
                                            case 32768:
                                                var v43 uint16 = 0
                                                _ = v43
                                            case -438:
                                                v12 = fi30
                                            case 2147483647:
                                                if true {
                                                    _ = len(v2)
                                                } else {
                                                    var v44 string = v1
                                                    v12 = (fi30 >> ((42) & 63))
                                                    _ = v44
                                                }
                                            }
                                        }
                                    case "":
                                        var v45 float64 = 3.14159
                                        _ = v45
                                    }
                                }
                            case "":
                                v15 = 2
                            }
                        case "a":
                            var v46 int64 = 9223372036854775807
                            _ = v46
                        case "hello":
                            var v47 int16 = 870
                            _ = v47
                        default:
                            _ = len(v7)
                        }
                    }
                    _ = fi30
                }
            }
        } else {
            var v48 float64 = (((-2.25 - float64(float32(3.14159))) * 0.0) * 1e10)
            var v49 float64 = (v48 - 0.0)
            v48 = v48
            _ = v48
            _ = v49
        }
        v17 -= v17
        v15 += 42
    }
    if false {
        var v50 uint16 = 2
        var v51 uint64 = 1
        var v52 int64 = 1
        _ = v50
        _ = v51
        _ = v52
    }
    v1 = (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v7
}
