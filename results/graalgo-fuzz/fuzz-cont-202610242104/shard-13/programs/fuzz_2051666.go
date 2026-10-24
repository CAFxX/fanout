package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 uint, p1 int16, p2 int64) int64 {
    v7 := []int{195, 9223372036854775807, 128, 1}
    switch "a" {
    case "a":
        fmt.Println(p1, p1)
    case "foo bar":
        p2 = (p2 + 42)
    case "":
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 uint8 = 1
            var v10 bool = (1 < 9223372036854775807)
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        var v11 int16 = -1
        _ = v11
    }
    if true {
        var v12 uint64 = 898
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 bool = false
            if v14 {
                var v15 int8 = -1
                var v16 uint = p0
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 string = "hello"
                    _ = fi17
                    _ = v18
                }
                _ = v15
                _ = v16
            } else {
                var v19 uint16 = 42
                if (p1 == (p1 | p1)) {
                    v20 := GenS1{F0: 2}
                    if v14 {
                        fmt.Println(v19)
                        v21 := GenS1{F0: 255}
                        _ = v21
                    } else {
                        v19 += v19
                        fmt.Println(fi13, p1, v14)
                    }
                    _ = v20
                }
                fmt.Println(p0, fi13, v12)
                _ = v19
            }
            var v22 string = "foo bar"
            _ = fi13
            _ = v14
            _ = v22
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            v24 := GenS2{F0: "", F1: 127, F2: 2}
            _ = fi23
            _ = v24
        }
        _ = v12
    } else {
        p1 -= 1
        for fi25 := 0; fi25 < 13; fi25++ {
            p1 = p1
            _ = fi25
        }
        var v26 uint = 224
        _ = v26
    }
    if false {
        var v27 int8 = -31
        _ = v27
    }
    v28 := GenS2{F0: "\n\t", F1: 155, F2: 50}
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    _ = v7
    return 0
}

func main() {
    var v1 int64 = 42
    var v2 uint64 = 1
    v3 := make(map[string]uint32)
    var v4 string = "hello"
    var v5 uint64 = 18446744073709551615
    v6 := make(map[int]string)
    if (v5 == v5) {
        v1 = v1
    } else {
        var v29 int = 127
        _ = v29
    }
    var v30 uint64 = v5
    var v31 uint16 = 42
    v5 -= (v5 * (v5 * 31))
    switch ((((v4 + (((v4 + ("\n\t" + ((v4 + v4) + "\n\t"))) + v4) + "hello")) + "a") + "a") + (("a" + (v4 + v4)) + "a")) {
    case "hello":
        fmt.Println(v31, v4)
    case "foo bar":
        switch v4 {
        case "hello":
            delete(v3, (("a" + v4) + (v4 + (v4 + "foo bar"))))
        case "a":
            for fi32 := 0; fi32 < 1; fi32++ {
                for fi33 := 0; fi33 < 3; fi33++ {
                    var v34 uint16 = (v31 * 333)
                    v31 -= v31
                    _ = fi33
                    _ = v34
                }
                switch (fi32 & fi32) {
                case 42:
                    var v35 uint = 0
                    _ = v35
                case 138:
                    if (404 >= v5) {
                        if (1e10 >= float32(3.14159)) {
                            v36 := GenS3{F0: 1, F1: (^(v31)), F2: 2147483647}
                            v37 := GenS4{F0: 138, F1: 42}
                            for fi38 := 0; fi38 < 11; fi38++ {
                                v30 = 18446744073709551615
                                v31 -= (v31 >> ((2) & 63))
                                _ = fi38
                            }
                            _ = v36
                            _ = v37
                        } else {
                            v39 := GenS5{F0: (v1 >> ((1) & 63)), F1: (fi32 * fi32)}
                            if ((v1 | v1) != 42) {
                                v40 := GenS5{F0: (v1 * 1), F1: (fi32 + (fi32 * fi32))}
                                var v41 int8 = 100
                                _ = v40
                                _ = v41
                            }
                            fmt.Println(fi32, fi32, v4)
                            _ = v39
                        }
                        var v42 int8 = -1
                        for fi43 := 0; fi43 < 7; fi43++ {
                            v44 := GenS1{F0: 128}
                            var v45 int8 = (-(v42))
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                        _ = v42
                    }
                default:
                    if (2 >= 208) {
                        if false {
                            v46 := GenS4{F0: v31, F1: 640}
                            var v47 uint16 = 42
                            _ = v46
                            _ = v47
                        }
                        var v48 int8 = -55
                        for fi49 := 0; fi49 < 13; fi49++ {
                            fmt.Println(v31, v2, v5)
                            _ = fi49
                        }
                        _ = v48
                    }
                }
                _ = fi32
            }
        case "foo bar":
            if false {
                v50 := GenS6{F0: 42, F1: 37}
                _ = v50
            }
        }
    }
    fmt.Println(v30, v1, v5)
    var v51 uint8 = 2
    var v52 float64 = float64(1e10)
    v2 = 0
    v2 = (v30 & 444)
    var v53 uint64 = (v2 ^ (v2 & (v2 >> ((uint(v51)) & 63))))
    v52 += v52
    if false {
        v54 := GenS7{F0: v51}
        fmt.Println(v31, v51)
        v55 := GenS2{F0: (((v4 + v4) + "") + v4), F1: uint8(v31), F2: 1}
        _ = v54
        _ = v55
    }
    for fi56 := 0; fi56 < 14; fi56++ {
        v57 := GenS6{F0: fi56, F1: 100}
        fmt.Println(v2, v52, v30)
        _ = fi56
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v6
}
