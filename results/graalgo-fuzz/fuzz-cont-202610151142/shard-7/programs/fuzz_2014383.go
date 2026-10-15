package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int32 }
type GenS2 struct { F0 int8; F1 int8; F2 uint }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int32; F1 uint }
func genFn0(p0 int16, p1 int32, p2 uint) uint16 {
    var v7 float32 = (0.0 + (0.1 + 1e10))
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            if ((p1 % ((p1) | 1)) > (-(p1))) {
                var v10 int64 = 0
                _ = v10
            } else {
                var v11 int = -48
                var v12 bool = (true || ((p1 | (p1 - p1)) != (p1 << (((p2 - p2)) & 63))))
                _ = v11
                _ = v12
            }
            v13 := []int{}
            v14 := GenS1{F0: -905, F1: (p0 & p0), F2: (-(p1))}
            p1 -= (p1 / ((p1) | 1))
            _ = fi9
            _ = v13
            _ = v14
        }
        _ = fi8
    }
    v15 := []int{1, 1, 128}
    var v16 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v7
    return uint16(p2)
}

func main() {
    var v1 bool = false
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    v4 := make(map[string]string)
    var v5 uint16 = 2
    var v6 uint16 = v5
    v5 -= v6
    if (!false) {
        v1 = false
    } else {
        var v17 int64 = -1
        _ = v17
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v1 = (127 <= 127)
        _ = fi18
    }
    fmt.Println(v5, v6)
    for fi19 := 0; fi19 < 19; fi19++ {
        v1 = v1
        switch "\n\t" {
        case "hello":
            fmt.Println(v6)
        case "\n\t":
            delete(v2, -2147483648)
        }
        var v20 string = ""
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: 0, F1: -45, F2: 4294967295}
    v22 := GenS3{F0: 996, F1: 633}
    var v23 uint64 = 18446744073709551615
    switch ("" + "foo bar") {
    case "a":
        if (v1 && ((0.1 == (float32(((3.14159 - 0.0) * 1e10)) / (0.0 + 0.1))) && true)) {
            var v24 string = "hello"
            _ = v24
        } else {
            var v25 uint64 = v23
            for fi26 := 0; fi26 < 4; fi26++ {
                if true {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v23 += v23
                        var v28 int16 = 32767
                        if v1 {
                            if (true || (uint32(v5) == 2147483648)) {
                                fmt.Println(v28, v5)
                                v25 = v23
                                fmt.Println(v23)
                            }
                            var v29 float64 = -2.25
                            v6 += 200
                            _ = v29
                        } else {
                            v30 := GenS4{F0: 293, F1: 0}
                            v23 = v23
                            _ = v30
                        }
                        _ = fi27
                        _ = v28
                    }
                    v5 += genFn0(42, 0, 1)
                    var v31 uint64 = (^(v25))
                    _ = v31
                } else {
                    v23 += (^(v23))
                    var v32 uint8 = 128
                    fmt.Println(fi26)
                    _ = v32
                }
                _ = fi26
            }
            _ = v25
        }
    case "hello":
        var v33 int = -1
        _ = v33
    case "\n\t":
        if (42 < 157) {
            for fi34 := 0; fi34 < 19; fi34++ {
                switch ("" + "\n\t") {
                case "hello":
                    v5 = 2
                case "":
                    if v1 {
                        v35 := v4["\n\t"]
                        var v36 float32 = (3.14159 / 3.14159)
                        for fi37 := 0; fi37 < 11; fi37++ {
                            var v38 int64 = -1
                            var v39 uint16 = v5
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        _ = v35
                        _ = v36
                    }
                default:
                    var v40 float32 = 3.14159
                    _ = v40
                }
                var v41 uint16 = 42
                _ = fi34
                _ = v41
            }
            v5 -= (v5 ^ 0)
        } else {
            v6 = (^(v5))
            var v42 uint8 = 0
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        if true {
            v44 := GenS4{F0: -84, F1: 895}
            v6 = (v6 - 177)
            _ = v44
        }
        for fi45 := 0; fi45 < 14; fi45++ {
            var v46 float64 = float64(1.5)
            _ = fi45
            _ = v46
        }
        _ = fi43
    }
    var v47 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
