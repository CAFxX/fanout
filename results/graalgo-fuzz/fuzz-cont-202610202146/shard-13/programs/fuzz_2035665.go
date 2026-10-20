package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int32; F1 string; F2 int8 }
type GenS4 struct { F0 uint32; F1 int32; F2 int }
type GenS5 struct { F0 uint; F1 int; F2 uint32 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint8, p2 int64) uint8 {
    v4 := []int{}
    v5 := GenS1{F0: "hello", F1: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p1 - 42)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    v3 := []int{2147483647}
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            if true {
                if true {
                    var v8 string = "hello"
                    _ = v8
                }
            } else {
                var v9 float32 = 0.0
                switch ("\n\t" + "\n\t") {
                case "a":
                    v10 := GenS2{F0: 0, F1: fi6}
                    _ = v10
                case "":
                    fmt.Println(fi7, fi7)
                }
                var v11 int16 = -827
                _ = v11
                _ = v9
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (176 < -748) {
                if false {
                    v12 := GenS3{F0: 2147483647, F1: ("hello" + ""), F2: 100}
                    fmt.Println(fi6)
                    _ = v12
                } else {
                    if (127 == (fi7 | fi7)) {
                        v13 := GenS2{F0: 2147483647, F1: fi7}
                        _ = v13
                    }
                    var v14 int32 = 1
                    _ = v14
                }
                if false {
                    v15 := GenS4{F0: 2, F1: 0, F2: fi7}
                    if true {
                        switch "a" {
                        case "hello":
                            if (true || false) {
                                if (!false) {
                                    var v16 int8 = 25
                                    var v17 uint8 = 1
                                    _ = v16
                                    _ = v17
                                } else {
                                    if true {
                                        var v18 uint64 = 42
                                        v19 := GenS4{F0: 2147483648, F1: 503, F2: 42}
                                        if false {
                                            fmt.Println(v18, fi7, v18)
                                            v18 += (v18 >> ((255) & 63))
                                        } else {
                                            v20 := GenS5{F0: 117, F1: (fi7 * fi6), F2: 949}
                                            if false {
                                                switch "" {
                                                case "":
                                                    var v21 uint64 = (^(v18))
                                                    _ = v21
                                                default:
                                                    var v22 int64 = 284
                                                    _ = v22
                                                }
                                                v18 = (v18 % ((684) | 1))
                                                var v23 bool = (!true)
                                                _ = v23
                                            } else {
                                                v18 = v18
                                                fmt.Println(fi7)
                                            }
                                            _ = v20
                                        }
                                        _ = v18
                                        _ = v19
                                    }
                                    var v24 uint32 = 611
                                    _ = v24
                                }
                                fmt.Println(fi6, fi7)
                            }
                        case "foo bar":
                            v25 := GenS2{F0: -980, F1: 0}
                            _ = v25
                        }
                        v3 = append(v3, -2147483648)
                    }
                    _ = v15
                }
                var v26 bool = true
                _ = v26
            }
            v27 := GenS6{F0: 1}
            _ = fi7
            _ = v27
        }
        _ = fi6
    }
    var v28 uint64 = 915
    v29 := GenS1{F0: "", F1: (-1 != 122)}
    for fi30 := 0; fi30 < 2; fi30++ {
        var v31 int32 = 1
        _ = fi30
        _ = v31
    }
    v32 := GenS5{F0: 947, F1: 2147483647, F2: 0}
    v28 = 1
    switch "foo bar" {
    case "a":
        v28 = (v28 >> ((2) & 63))
    }
    var v33 int32 = 1
    var v34 int16 = -642
    if false {
        v34 += v34
        v28 -= v28
    } else {
        v35 := v1[2147483647]
        _ = v35
    }
    v2[-2147483648] = -635
    v36 := GenS6{F0: 2147483648}
    v33 = (v33 % (((v33 | (v33 & v33))) | 1))
    v37 := v2[32767]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v37
}
