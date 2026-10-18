package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 string) uint64 {
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{(^(fi7)), (fi7 / ((fi7) | 1)), fi7}
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    if true {
        v9 := GenS1{F0: ("" + ("\n\t" + "a")), F1: 42}
        _ = v9
    } else {
        v10 := []int{}
        var v11 int64 = 9223372036854775807
        _ = v10
        _ = v11
    }
    _ = p0
    return 530
}
func genFn1(p0 int64) int16 {
    p0 += -15
    v12 := []int{-488, 1, 2147483647}
    var v13 int = 4294967295
    if false {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        v13 = (v13 << ((65535) & 63))
    }
    _ = p0
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 float64 = 0.0
    var v2 string = "a"
    var v3 float64 = v1
    v4 := make(map[uint32]int)
    var v5 int32 = 42
    var v6 uint8 = 1
    var v14 int = -8
    var v15 int = -1
    fmt.Println(v3, v1, v3)
    switch v14 {
    case -1:
        var v16 uint = 4294967295
        _ = v16
    case 32768:
        fmt.Println(v5, v1, v3)
    default:
        var v17 uint64 = 42
        _ = v17
    }
    fmt.Println(v15)
    var v18 int = 102
    v19 := GenS2{F0: -128}
    v20 := []int{(v18 + 0), (v14 - 405), 12, v15}
    var v21 bool = true
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 int8 = -1
        fmt.Println(v2, fi22)
        _ = fi22
        _ = v23
    }
    delete(v4, 710)
    for fi24 := 0; fi24 < 13; fi24++ {
        if v21 {
            var v25 uint64 = 1
            if (!false) {
                v14 = v14
                var v26 int16 = -361
                v6 = v6
                _ = v26
            }
            _ = len(v4)
            _ = v25
        } else {
            var v27 string = "hello"
            _ = v27
        }
        v28 := GenS2{F0: 42}
        switch v18 {
        case 128:
            for fi29 := 0; fi29 < 19; fi29++ {
                if (v21 && true) {
                    if false {
                        v30 := GenS3{F0: uint16(v6)}
                        fmt.Println(v3, v14)
                        _ = v30
                    } else {
                        v2 = v2
                        if v21 {
                            v31 := GenS4{F0: genFn0(v2)}
                            if (765 != 65535) {
                                v15 = 9223372036854775807
                            } else {
                                var v32 string = ("hello" + v2)
                                v14 = (v15 + (fi24 >> ((65535) & 63)))
                                fmt.Println(v2, v14)
                                _ = v32
                            }
                            _ = v31
                        }
                    }
                    var v33 uint32 = 4294967295
                    _ = v33
                } else {
                    var v34 uint32 = 42
                    var v35 string = ((v2 + "foo bar") + (v2 + ("\n\t" + ("hello" + v2))))
                    v18 = v14
                    if false {
                        v35 = v35
                        v36 := GenS4{F0: 839}
                        var v37 uint16 = 0
                        _ = v36
                        _ = v37
                    }
                    if v21 {
                        var v38 string = v35
                        var v39 string = "\n\t"
                        delete(v4, uint32(v6))
                        _ = v38
                        _ = v39
                    }
                    _ = v34
                    _ = v35
                }
                v15 = (fi29 ^ v14)
                var v40 bool = v21
                _ = fi29
                _ = v40
            }
        default:
            var v41 float32 = 1.5
            _ = v41
        }
        _ = fi24
        _ = v28
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        var v43 string = "hello"
        var v44 uint = 63
        _ = fi42
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
