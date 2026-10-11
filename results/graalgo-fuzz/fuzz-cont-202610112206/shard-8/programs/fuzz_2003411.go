package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -2147483648
    var v3 uint16 = 73
    var v4 string = ("\n\t" + "foo bar")
    fmt.Println(v4, v2, v3)
    v3 = v3
    v5 := GenS1{F0: 654}
    switch 128 {
    case 2147483647:
        if (0 != 0) {
            v6 := []int{-9223372036854775808, 42, -1}
            _ = v6
        }
    case 128:
        fmt.Println(v2, v3)
    case 32767:
        v7 := GenS1{F0: 1}
        _ = v7
    default:
        var v8 int64 = 9223372036854775807
        _ = v8
    }
    var v9 int8 = 1
    var v10 float64 = float64(0.1)
    if true {
        delete(v1, 269)
    }
    v11 := []int{}
    for fi12 := 0; fi12 < 3; fi12++ {
        _ = len(v11)
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 int16 = 42
            _ = fi13
            _ = v14
        }
        v3 -= v3
        _ = fi12
    }
    fmt.Println(v10, v2, v4)
    if true {
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 7; fi16++ {
                if false {
                    fmt.Println(v2, fi15)
                } else {
                    v17 := GenS1{F0: 4294967295}
                    switch v4 {
                    case "":
                        if false {
                            v18 := GenS1{F0: (fi16 ^ (fi15 ^ 0))}
                            _ = v18
                        } else {
                            v19 := GenS1{F0: (-(fi15))}
                            v10 = v10
                            _ = v19
                        }
                    }
                    var v20 uint64 = 42
                    var v21 bool = false
                    var v22 uint32 = 2147483648
                    if v21 {
                        v4 = v4
                        var v23 uint = 619
                        _ = v23
                    }
                    v21 = false
                    _ = v17
                    _ = v20
                    _ = v21
                    _ = v22
                }
                if false {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    fmt.Println(v4, v10)
                    if false {
                        v2 -= v2
                        v4 = ""
                        if (3.14159 < v10) {
                            if (!true) {
                                var v24 float64 = ((v10 / (v10 - 1.5)) / v10)
                                var v25 uint64 = 42
                                _ = v24
                                _ = v25
                            }
                            v2 = -2147483648
                            var v26 int16 = 356
                            _ = v26
                        } else {
                            var v27 int = 621
                            v28 := GenS1{F0: (^(v27))}
                            v29 := GenS2{F0: v3, F1: v3}
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    }
                }
                if false {
                    var v30 uint32 = 1
                    switch "a" {
                    case "foo bar":
                        v1[1] = false
                    case "":
                        var v31 int8 = (v9 ^ v9)
                        _ = v31
                    }
                    v30 = (^(v30))
                    _ = v30
                }
                _ = fi16
            }
            if (true || (65535 == 226)) {
                if len(v11) > 0 {
                    v11[0] = (fi15 - (fi15 + 128))
                }
                v2 = -1
            }
            for fi32 := 0; fi32 < 12; fi32++ {
                v33 := GenS1{F0: (fi32 * fi15)}
                fmt.Println(fi32, v9, fi15)
                v2 += v2
                _ = fi32
                _ = v33
            }
            _ = fi15
        }
        v34 := GenS2{F0: 42, F1: (v3 % (((v3 - v3)) | 1))}
        _ = v34
    } else {
        for fi35 := 0; fi35 < 2; fi35++ {
            var v36 float32 = 0.1
            _ = fi35
            _ = v36
        }
        v9 = v9
        switch 9223372036854775807 {
        case -9223372036854775808:
            var v37 uint32 = 65535
            _ = v37
        }
    }
    var v38 uint16 = 65535
    _ = len(v1)
    var v39 uint32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v9
}
