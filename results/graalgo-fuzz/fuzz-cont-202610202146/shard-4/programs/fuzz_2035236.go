package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint; F1 int32; F2 int32 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 int8; F1 int16 }
func genFn0(p0 uint16) uint8 {
    var v6 uint64 = 42
    v7 := GenS1{F0: p0, F1: 255, F2: (^(p0))}
    v6 += v6
    var v8 float32 = (0.1 - 1e10)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint, p1 int16, p2 int) bool {
    var v9 uint16 = 42
    v9 -= v9
    v10 := []int{p2}
    fmt.Println(v9, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return false
}

func main() {
    var v1 uint16 = 42
    var v2 float64 = 0.0
    v3 := make(map[string]bool)
    var v4 int16 = -1
    v5 := []int{}
    _ = len(v5)
    for fi11 := 0; fi11 < 15; fi11++ {
        fmt.Println(v4, v4)
        if false {
            for fi12 := 0; fi12 < 5; fi12++ {
                if (0 != -2147483648) {
                    var v13 string = "a"
                    _ = v13
                } else {
                    fmt.Println(fi11)
                    v14 := GenS2{F0: 653, F1: 165, F2: 1}
                    _ = v14
                }
                _ = fi12
            }
            if (0 == 1) {
                var v15 int32 = 2147483647
                var v16 int16 = (v4 - (v4 / ((v4) | 1)))
                _ = v15
                _ = v16
            } else {
                var v17 int = (fi11 >> ((242) & 63))
                _ = v17
            }
            switch fi11 {
            case -614:
                v4 -= v4
            }
        } else {
            var v18 uint64 = 1
            if false {
                if false {
                    for fi19 := 0; fi19 < 5; fi19++ {
                        if (-1 != 42) {
                            _ = len(v5)
                            switch 4294967295 {
                            case 127:
                                delete(v3, "foo bar")
                            case 176:
                                var v20 string = "foo bar"
                                _ = v20
                            case 4294967295:
                                v18 -= v18
                            }
                            var v21 int32 = -717
                            var v22 string = ("\n\t" + "")
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v21
                            _ = v22
                        }
                        v18 = v18
                        _ = fi19
                    }
                    if (false && (125 < 65535)) {
                        v2 -= float64((-2.25 + -2.25))
                        for fi23 := 0; fi23 < 16; fi23++ {
                            v5 = append(v5, fi11)
                            v4 = v4
                            var v24 int16 = 42
                            _ = fi23
                            _ = v24
                        }
                        fmt.Println(v18, v18, fi11)
                    }
                } else {
                    var v25 int = int(v1)
                    var v26 int = (^(fi11))
                    v3[("a" + "hello")] = false
                    _ = v25
                    _ = v26
                }
                v27 := GenS3{F0: v4, F1: 341}
                _ = v27
            }
            _ = v18
        }
        v28 := GenS3{F0: 32767, F1: 1}
        _ = fi11
        _ = v28
    }
    v1 -= 32767
    var v29 float32 = (0.0 / 1.0)
    v4 -= v4
    var v30 bool = true
    var v31 uint = 42
    switch "a" {
    case "\n\t":
        v31 -= 780
    case "foo bar":
        var v32 int32 = 42
        _ = v32
    case "":
        for fi33 := 0; fi33 < 15; fi33++ {
            v34 := GenS4{F0: (v31 & v31), F1: fi33}
            _ = fi33
            _ = v34
        }
    }
    var v35 int = -713
    v35 -= v35
    if true {
        v4 -= int16(v31)
    } else {
        switch "foo bar" {
        case "":
            fmt.Println(v4, v29, v31)
        case "foo bar":
            for fi36 := 0; fi36 < 19; fi36++ {
                for fi37 := 0; fi37 < 12; fi37++ {
                    var v38 int32 = -2147483648
                    _ = fi37
                    _ = v38
                }
                _ = len(v3)
                _ = fi36
            }
        case "a":
            var v39 string = ("" + "\n\t")
            _ = v39
        }
        v2 += ((-2.25 + v2) / 3.14159)
        for fi40 := 0; fi40 < 15; fi40++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi40
        }
    }
    switch int(v4) {
    case 4294967295:
        switch "hello" {
        case "a":
            v4 = (v4 << ((v31) & 63))
        case "":
            _ = len(v3)
        case "hello":
            v41 := GenS1{F0: (v1 >> (((v31 >> ((v31) & 63))) & 63)), F1: 118, F2: 216}
            _ = v41
        }
    case 9223372036854775807:
        var v42 uint64 = 960
        _ = v42
    default:
        v43 := GenS5{F0: 127, F1: v4}
        _ = v43
    }
    v44 := GenS1{F0: 351, F1: 42, F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v44
    _ = v5
}
