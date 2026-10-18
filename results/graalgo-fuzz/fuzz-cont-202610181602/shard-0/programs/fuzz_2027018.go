package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint32; F2 bool }

func main() {
    var v1 uint64 = 547
    v2 := make(map[int]int)
    v3 := []string{"", "hello"}
    var v4 uint64 = v1
    if false {
        for fi5 := 0; fi5 < 5; fi5++ {
            delete(v2, fi5)
            _ = fi5
        }
    } else {
        v3 = append(v3, ("" + "a"))
        v1 = v4
        if (v4 > v4) {
            _ = len(v2)
        }
    }
    if len(v3) > 0 {
        v3[0] = ""
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = v1
    if false {
        var v6 int32 = 2147483647
        if true {
            if true {
                v6 = (v6 >> ((2) & 63))
            }
            for fi7 := 0; fi7 < 10; fi7++ {
                for fi8 := 0; fi8 < 5; fi8++ {
                    switch "" {
                    case "a":
                        v9 := GenS1{F0: 2, F1: 2}
                        _ = v9
                    case "\n\t":
                        if (2 < 654) {
                            var v10 float64 = (0.0 + 1.5)
                            if false {
                                var v11 uint16 = 1
                                var v12 int8 = 0
                                switch "\n\t" {
                                case "hello":
                                    var v13 int8 = v12
                                    _ = v13
                                }
                                _ = v11
                                _ = v12
                            } else {
                                v10 = float64((((3.14159 + 0.0) / 1e10) + float32(1.5)))
                            }
                            v14 := GenS2{F0: v6}
                            _ = v10
                            _ = v14
                        } else {
                            if ((^(v4)) != (v1 | 834)) {
                                v1 = (v4 + 0)
                                v6 = (v6 | 42)
                            }
                            if false {
                                if (65535 < 65535) {
                                    fmt.Println(fi7, v6)
                                    fmt.Println(v1, v6)
                                    _ = len(v2)
                                }
                                v4 = v4
                                v1 += v1
                            }
                            v6 -= (v6 << ((0) & 63))
                            var v15 uint = 255
                            _ = v15
                        }
                    }
                    _ = fi8
                }
                if false {
                    if true {
                        if (false && true) {
                            var v16 int16 = -745
                            var v17 int32 = (v6 >> ((65535) & 63))
                            _ = v16
                            _ = v17
                        }
                    } else {
                        var v18 int8 = 42
                        v1 = (^(v1))
                        _ = v18
                    }
                }
                _ = fi7
            }
            v4 = (v1 >> ((1) & 63))
        }
        _ = v6
    } else {
        fmt.Println(v4)
        var v19 uint = 2
        _ = v19
    }
    fmt.Println(v4, v4)
    v20 := GenS3{F0: 315, F1: 185, F2: false}
    var v21 bool = true
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            _ = len(v3)
            _ = fi23
        }
        var v24 uint16 = uint16(v4)
        _ = fi22
        _ = v24
    }
    var v25 bool = true
    v25 = v21
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
}
