package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int; F1 uint64; F2 bool }
type GenS5 struct { F0 int8; F1 uint64; F2 uint8 }

func main() {
    var v1 bool = (false && (((73 >= 1) && (!false)) && true))
    v2 := make(map[uint32]int)
    var v3 uint32 = 294
    var v4 int64 = 42
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    var v7 string = "a"
    fmt.Println(v1, v4, v4)
    if v1 {
        v4 -= (v4 & v4)
    }
    var v8 uint64 = 0
    if (4294967295 < 9223372036854775807) {
        for fi9 := 0; fi9 < 15; fi9++ {
            v4 += 42
            var v10 uint32 = v3
            _ = fi9
            _ = v10
        }
    } else {
        v7 = "a"
        for fi11 := 0; fi11 < 7; fi11++ {
            if false {
                switch (v7 + v7) {
                case "hello":
                    v1 = (((v7 + "a") + "foo bar") > "\n\t")
                case "":
                    switch (fi11 >> ((313) & 63)) {
                    case 4294967295:
                        switch v7 {
                        case "foo bar":
                            v12 := []int{-655, fi11, (fi11 % ((0) | 1)), fi11}
                            _ = v12
                        case "":
                            for fi13 := 0; fi13 < 9; fi13++ {
                                v3 -= v3
                                v3 = 752
                                var v14 float64 = 1.5
                                if v1 {
                                    v7 = "foo bar"
                                } else {
                                    if ("foo bar" != v7) {
                                        v15 := GenS1{F0: int32(v4), F1: 23, F2: (v4 / ((int64(v8)) | 1))}
                                        v7 = v7
                                        var v16 int = int(v4)
                                        _ = v15
                                        _ = v16
                                    }
                                }
                                _ = fi13
                                _ = v14
                            }
                        case "a":
                            v17 := GenS2{F0: fi11}
                            _ = v17
                        }
                    case -1:
                        fmt.Println(fi11, v8)
                    default:
                        switch (fi11 / ((-827) | 1)) {
                        case -295:
                            if v1 {
                                if (-716 >= -1) {
                                    v18 := []int{(fi11 / ((fi11) | 1)), (^(fi11)), fi11, (fi11 ^ fi11)}
                                    v7 = v7
                                    var v19 int64 = v4
                                    _ = v18
                                    _ = v19
                                }
                            }
                        case 42:
                            v20 := []int{}
                            _ = v20
                        default:
                            v5[""] = v7
                        }
                    }
                }
                v21 := GenS3{F0: 0, F1: 2}
                _ = v21
            }
            _ = fi11
        }
        var v22 uint16 = 255
        _ = v22
    }
    v23 := v6[v3]
    fmt.Println(v7, v23, v4)
    var v24 float32 = -2.25
    var v25 uint = 42
    if (255 != 219) {
        switch 305 {
        case 32768:
            var v26 float64 = 1e10
            _ = v26
        case -371:
            var v27 uint16 = 1
            _ = v27
        }
    } else {
        fmt.Println(v7, v4, v8)
        for fi28 := 0; fi28 < 17; fi28++ {
            v7 = (v23 + "a")
            _ = fi28
        }
        var v29 uint32 = v3
        _ = v29
    }
    switch 400 {
    case 32768:
        var v30 int32 = 0
        _ = v30
    case -5:
        v7 = v7
    case 2147483647:
        fmt.Println(v4)
    default:
        for fi31 := 0; fi31 < 18; fi31++ {
            v32 := []int{-9223372036854775808, 32767, fi31}
            if v1 {
                if (true || (true && v1)) {
                    _ = len(v32)
                    var v33 int32 = -1
                    if ((v1 || (1.5 != 1e10)) || (42 <= 63)) {
                        fmt.Println(v7, v4)
                        switch 0 {
                        case 2147483647:
                            for fi34 := 0; fi34 < 12; fi34++ {
                                v35 := GenS2{F0: 127}
                                fmt.Println(v4)
                                _ = fi34
                                _ = v35
                            }
                        case 398:
                            v36 := GenS3{F0: -29, F1: (v25 - uint(v3))}
                            _ = v36
                        case -9223372036854775808:
                            if v1 {
                                fmt.Println(v33)
                                v5["hello"] = (v7 + ("foo bar" + "\n\t"))
                            } else {
                                v37 := GenS4{F0: fi31, F1: v8, F2: (v25 <= (v25 >> ((255) & 63)))}
                                _ = v37
                            }
                        default:
                            for fi38 := 0; fi38 < 16; fi38++ {
                                var v39 uint = 18446744073709551615
                                fmt.Println(v8)
                                v40 := GenS5{F0: 0, F1: (v8 % ((uint64(v3)) | 1)), F2: 1}
                                var v41 int = fi38
                                _ = fi38
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                        }
                    }
                    _ = v33
                } else {
                    v6[(v3 >> ((2) & 63))] = v23
                }
            }
            _ = fi31
            _ = v32
        }
    }
    delete(v6, v3)
    var v42 bool = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
