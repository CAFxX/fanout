package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint32 }
func genFn0(p0 int64) int16 {
    if true {
        var v6 int32 = 42
        v7 := GenS1{F0: 1000, F1: -745, F2: (1000 != 686)}
        v6 = v6
        _ = v6
        _ = v7
    } else {
        p0 = (p0 / ((p0) | 1))
        var v8 uint8 = 127
        for fi9 := 0; fi9 < 19; fi9++ {
            if ((105 >= 91) && false) {
                if false {
                    fmt.Println(fi9, p0, v8)
                }
                if ("foo bar" >= "hello") {
                    if ("foo bar" < "\n\t") {
                        var v10 uint64 = 42
                        fmt.Println(v10, v8, p0)
                        v11 := GenS2{F0: -500}
                        _ = v10
                        _ = v11
                    } else {
                        if ((v8 + 210) <= 0) {
                            for fi12 := 0; fi12 < 2; fi12++ {
                                if true {
                                    v8 = v8
                                }
                                p0 = p0
                                _ = fi12
                            }
                        } else {
                            var v13 uint32 = 2147483648
                            var v14 int = fi9
                            _ = v13
                            _ = v14
                        }
                        v15 := GenS2{F0: 128}
                        _ = v15
                    }
                }
            } else {
                var v16 int = 2147483647
                var v17 uint32 = 2
                _ = v16
                _ = v17
            }
            var v18 int32 = 246
            _ = fi9
            _ = v18
        }
        _ = v8
    }
    var v19 int16 = -32768
    fmt.Println(p0, v19)
    v20 := []int{-239, -978}
    v19 = (v19 << ((1) & 63))
    _ = p0
    _ = v19
    _ = v20
    return v19
}
func genFn1() int {
    var v21 bool = false
    switch "a" {
    case "\n\t":
        if false {
            var v22 int8 = -128
            v23 := GenS2{F0: 127}
            v24 := []int{}
            _ = v22
            _ = v23
            _ = v24
        }
    case "":
        var v25 int32 = -752
        _ = v25
    default:
        for fi26 := 0; fi26 < 2; fi26++ {
            v21 = (!v21)
            v21 = true
            for fi27 := 0; fi27 < 3; fi27++ {
                var v28 uint64 = 7
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    }
    _ = v21
    return 127
}

func main() {
    v1 := []int{}
    var v2 float64 = 0.1
    var v3 float64 = 1.5
    var v4 uint64 = 0
    var v5 string = ("hello" + "")
    v2 = -2.25
    v3 = float64(float32(v3))
    fmt.Println(v3, v3)
    v2 += v3
    v3 = v3
    var v29 uint16 = 255
    var v30 uint32 = 2147483648
    v1 = append(v1, 0)
    v2 = v2
    for fi31 := 0; fi31 < 14; fi31++ {
        v32 := GenS1{F0: -1, F1: genFn0(9223372036854775807), F2: true}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi31
        _ = v32
    }
    if len(v1) > 0 {
        v1[0] = int(v30)
    }
    fmt.Println(v30, v5)
    for fi33 := 0; fi33 < 1; fi33++ {
        if false {
            for fi34 := 0; fi34 < 11; fi34++ {
                v35 := GenS3{F0: -163, F1: 0}
                var v36 uint = 42
                _ = fi34
                _ = v35
                _ = v36
            }
            fmt.Println(v5)
        } else {
            var v37 int8 = -72
            var v38 float32 = 1e10
            var v39 int32 = -738
            _ = v37
            _ = v38
            _ = v39
        }
        fmt.Println(v29)
        _ = fi33
    }
    for fi40 := 0; fi40 < 2; fi40++ {
        fmt.Println(v3, fi40)
        _ = fi40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
