package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int32; F2 bool }
type GenS3 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint; F2 int16 }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS6 struct { F0 uint8; F1 uint8; F2 int16 }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 int8 = 127
    v5 := make(map[int]int)
    var v6 float64 = ((float64((3.14159 - 1.5)) / 0.1) - 1.5)
    fmt.Println(v1, v6)
    v7 := GenS1{F0: 246}
    v8 := GenS2{F0: 255, F1: v1, F2: (!((2 > 127) || false))}
    var v9 string = ("a" + "a")
    _ = len(v2)
    var v10 float64 = v6
    v4 += v4
    if true {
        v10 = (v10 * float64(-2.25))
        v10 = (0.1 + ((3.14159 - (1.5 / v6)) / v10))
    } else {
        v1 = v1
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int8 = 1
        _ = fi11
        _ = v12
    }
    if false {
        var v13 uint8 = 56
        _ = v13
    } else {
        v6 = 1e10
        if (true && true) {
            v10 -= v6
        }
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        switch fi14 {
        case -747:
            v4 -= (-(v4))
        case -303:
            for fi15 := 0; fi15 < 1; fi15++ {
                if ((fi14 != (fi14 << ((65535) & 63))) && (!false)) {
                    fmt.Println(v6)
                    switch "a" {
                    case "hello":
                        if (3 != 2) {
                            var v16 string = (v9 + v9)
                            if true {
                                if true {
                                    fmt.Println(v6)
                                    if true {
                                        v17 := GenS3{F0: 0, F1: (v4 * v4), F2: 2147483648}
                                        switch v9 {
                                        case "hello":
                                            fmt.Println(fi15, v4, v4)
                                        case "":
                                            fmt.Println(v16, v1)
                                        default:
                                            if false {
                                                v6 = float64(1e10)
                                                v4 -= (-(v4))
                                                var v18 bool = true
                                                _ = v18
                                            } else {
                                                v19 := []int{(fi15 | (fi14 >> ((42) & 63)))}
                                                var v20 int16 = -378
                                                v21 := GenS4{F0: true, F1: 1, F2: (v20 / ((92) | 1))}
                                                _ = v19
                                                _ = v20
                                                _ = v21
                                            }
                                        }
                                        var v22 uint64 = 175
                                        var v23 string = ""
                                        v24 := []int{fi14, fi14, fi14}
                                        _ = v17
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    } else {
                                        var v25 uint = 1
                                        switch (fi14 / ((int(v4)) | 1)) {
                                        case 912:
                                            switch (v9 + v9) {
                                            case "":
                                                fmt.Println(v9)
                                            default:
                                                v26 := []int{fi14}
                                                _ = v26
                                            }
                                        }
                                        _ = v25
                                    }
                                } else {
                                    if (((0.0 * 0.1) > -2.25) || true) {
                                        if ((((v10 * v6) - (v10 + v10)) / v10) != float64(-2.25)) {
                                            v27 := GenS5{F0: 112, F1: 59, F2: 2}
                                            var v28 int16 = 0
                                            _ = v27
                                            _ = v28
                                        }
                                        v6 = -2.25
                                        var v29 int = (fi14 << ((42) & 63))
                                        _ = v29
                                    } else {
                                        if false {
                                            v10 -= v10
                                            fmt.Println(fi14)
                                            var v30 int8 = int8(v1)
                                            _ = v30
                                        } else {
                                            if true {
                                                var v31 int64 = int64(fi15)
                                                var v32 int64 = (v31 | 9223372036854775807)
                                                fmt.Println(v16)
                                                v33 := GenS3{F0: 1, F1: v4, F2: 2147483647}
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                            }
                                            v16 = ""
                                        }
                                        var v34 float64 = float64((0.0 - 0.0))
                                        var v35 int64 = 983
                                        _ = v34
                                        _ = v35
                                    }
                                    v36 := []int{}
                                    _ = v36
                                }
                                _ = len(v5)
                                fmt.Println(v1, v10)
                            } else {
                                if true {
                                    var v37 uint8 = 2
                                    _ = v37
                                }
                            }
                            var v38 float64 = (1e10 * 3.14159)
                            _ = v16
                            _ = v38
                        } else {
                            var v39 float32 = 1.5
                            v6 -= float64(v39)
                            v40 := GenS5{F0: 206, F1: 2, F2: 0}
                            _ = v39
                            _ = v40
                        }
                    case "foo bar":
                        switch fi14 {
                        case -9223372036854775808:
                            var v41 int8 = -36
                            _ = v41
                        }
                    case "a":
                        var v42 float32 = 0.0
                        _ = v42
                    default:
                        v1 = -653
                    }
                    if true {
                        fmt.Println(v6)
                        var v43 float64 = (3.14159 - v6)
                        _ = v43
                    }
                    v44 := GenS6{F0: 160, F1: 2, F2: 1000}
                    fmt.Println(fi14)
                    _ = v44
                }
                if true {
                    var v45 uint32 = 183
                    _ = v45
                } else {
                    var v46 int32 = -1
                    fmt.Println(fi15, v4, v4)
                    v47 := GenS4{F0: (false && (true || false)), F1: 255, F2: 0}
                    _ = v46
                    _ = v47
                }
                _ = fi15
            }
        default:
            switch 32767 {
            case 697:
                v48 := GenS3{F0: 67, F1: (v4 & v4), F2: 42}
                _ = v48
            }
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
