package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 uint64 }
func genFn0() int64 {
    v4 := GenS1{F0: 42, F1: 65535, F2: 233}
    var v5 int64 = -1
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint32 = 1
        v8 := GenS2{F0: 1}
        _ = fi6
        _ = v7
        _ = v8
    }
    v5 += (v5 | 1)
    v5 -= v5
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    var v3 float32 = 0.0
    var v9 int32 = 0
    fmt.Println(v3, v9)
    _ = len(v1)
    for fi10 := 0; fi10 < 17; fi10++ {
        if len(v2) > 0 {
            v2[0] = fi10
        }
        _ = fi10
    }
    if (!true) {
        if true {
            fmt.Println(v3, v9)
            if false {
                _ = len(v2)
                _ = len(v1)
                for fi11 := 0; fi11 < 19; fi11++ {
                    var v12 uint = 2
                    _ = fi11
                    _ = v12
                }
            } else {
                delete(v1, "foo bar")
                fmt.Println(v9, v9)
                switch "foo bar" {
                case "foo bar":
                    for fi13 := 0; fi13 < 5; fi13++ {
                        var v14 bool = false
                        _ = fi13
                        _ = v14
                    }
                case "":
                    v15 := GenS3{F0: 2147483648}
                    _ = v15
                case "a":
                    switch "a" {
                    case "\n\t":
                        var v16 uint64 = 42
                        _ = v16
                    }
                }
            }
        } else {
            if len(v2) > 0 {
                v2[0] = 13
            }
        }
    } else {
        if true {
            var v17 bool = (9223372036854775807 < -1)
            if (!(!v17)) {
                var v18 int16 = -514
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 uint64 = 858
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 2; fi21++ {
                    for fi22 := 0; fi22 < 9; fi22++ {
                        if ((4294967295 == 2) && v17) {
                            v9 -= v9
                            v23 := GenS1{F0: 1, F1: 510, F2: (fi22 + (fi21 & (fi22 & fi21)))}
                            if v17 {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                var v24 float64 = float64((v3 * v3))
                                _ = v24
                            }
                            v17 = v17
                            _ = v23
                        }
                        fmt.Println(v9, v18)
                        _ = fi22
                    }
                    var v25 float64 = 1e10
                    fmt.Println(v3)
                    _ = fi21
                    _ = v25
                }
                _ = v18
            }
            _ = v17
        }
        v9 -= (-(v9))
        fmt.Println(v9, v3)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v26 := GenS4{F0: (v9 % ((-1) | 1))}
    if (false && true) {
        fmt.Println(v3)
    } else {
        fmt.Println(v3)
        v2 = append(v2, -990)
        fmt.Println(v3, v3)
    }
    if true {
        var v27 int = 0
        var v28 string = "hello"
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        v30 := GenS3{F0: 65535}
        var v31 uint64 = 2
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 uint = 2
    var v33 uint = v32
    v34 := GenS5{F0: 42, F1: 961}
    v9 += -717
    v9 = (v9 ^ 0)
    fmt.Println(v3, v33, v33)
    fmt.Println(v9, v32)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v9
}
