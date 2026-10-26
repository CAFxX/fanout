package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 bool; F1 bool }
func genFn0() uint64 {
    v7 := GenS1{F0: -519, F1: 878, F2: 2}
    v8 := []int{}
    if (!(127 <= 91)) {
        var v9 int = 128
        var v10 float64 = float64((3.14159 * float32(3.14159)))
        _ = v10
        _ = v9
    }
    var v11 int = -2147483648
    _ = v11
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int8, p1 string) uint32 {
    p0 = 41
    v12 := []int{}
    p1 = p1
    if len(v12) > 0 {
        v12[0] = 707
    }
    _ = p0
    _ = p1
    _ = v12
    return 905
}

func main() {
    var v1 int8 = 0
    var v2 uint64 = 1
    var v3 bool = (42 >= 0)
    var v4 uint = 326
    v5 := make(map[int]int)
    var v6 int64 = -150
    var v13 int32 = 2147483647
    switch "\n\t" {
    case "\n\t":
        var v14 uint64 = v2
        _ = v14
    case "":
        fmt.Println(v13, v2, v4)
    default:
        fmt.Println(v1, v4, v3)
    }
    v4 = (^(v4))
    for fi15 := 0; fi15 < 5; fi15++ {
        switch (("a" + "hello") + ("foo bar" + "hello")) {
        case "foo bar":
            for fi16 := 0; fi16 < 19; fi16++ {
                switch "a" {
                case "a":
                    var v17 float64 = 1e10
                    _ = v17
                case "hello":
                    v13 += (v13 << ((v4) & 63))
                }
                var v18 uint32 = genFn1((v1 >> (((v4 * 0)) & 63)), ("" + "foo bar"))
                _ = fi16
                _ = v18
            }
        case "":
            if (v1 < 100) {
                var v19 int = 32768
                var v20 string = ("\n\t" + (("a" + ("\n\t" + "\n\t")) + ("hello" + "\n\t")))
                v21 := []int{(fi15 << ((v4) & 63)), 127}
                _ = v19
                _ = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 17; fi22++ {
                    var v23 int = 501
                    _ = fi22
                    _ = v23
                }
                for fi24 := 0; fi24 < 6; fi24++ {
                    v25 := GenS2{F0: v4, F1: v2}
                    switch fi24 {
                    case 2147483647:
                        var v26 int = fi15
                        _ = v26
                    default:
                        switch (fi15 + (fi24 * fi24)) {
                        case 4294967295:
                            var v27 uint8 = 169
                            _ = v27
                        case 42:
                            v6 -= v6
                        case 32768:
                            if (v3 && v3) {
                                fmt.Println(v6)
                            } else {
                                fmt.Println(v1)
                                v6 -= int64(v4)
                                if v3 {
                                    switch (("hello" + "a") + "") {
                                    case "":
                                        fmt.Println(fi24)
                                        if false {
                                            delete(v5, fi24)
                                        }
                                    }
                                    v2 = (v2 ^ (v2 ^ 0))
                                    v28 := GenS1{F0: (fi15 - (^(fi24))), F1: uint(v2), F2: (v2 + (v2 / (((v2 / (((v2 - v2)) | 1))) | 1)))}
                                    var v29 uint8 = 0
                                    v3 = v3
                                    _ = v28
                                    _ = v29
                                }
                            }
                        default:
                            v6 = 0
                            var v30 uint32 = 65535
                            _ = v30
                        }
                    }
                    _ = fi24
                    _ = v25
                }
                if v3 {
                    if v3 {
                        switch "" {
                        case "\n\t":
                            var v31 string = ""
                            _ = v31
                        case "hello":
                            v32 := GenS3{F0: true, F1: (!(-1 == v1))}
                            _ = v32
                        case "a":
                            v5[(fi15 ^ fi15)] = fi15
                        }
                    }
                    v2 += genFn0()
                }
            }
        }
        v6 = -762
        _ = fi15
    }
    fmt.Println(v4, v4, v6)
    for fi33 := 0; fi33 < 12; fi33++ {
        v1 += 0
        _ = fi33
    }
    v34 := []int{-135, -214}
    if false {
        v3 = false
    }
    var v35 string = ("foo bar" + "hello")
    if v3 {
        v36 := GenS1{F0: 128, F1: 2, F2: 1}
        _ = len(v5)
        _ = v36
    } else {
        var v37 float64 = float64(1e10)
        v3 = false
        _ = v37
    }
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
