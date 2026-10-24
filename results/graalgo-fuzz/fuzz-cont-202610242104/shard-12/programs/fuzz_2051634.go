package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
func genFn0(p0 int32, p1 int8, p2 uint64) int8 {
    p0 -= p0
    p1 -= p1
    var v6 int8 = -128
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -115
}
func genFn1() int8 {
    var v7 int = -1
    v7 -= 4294967295
    _ = v7
    return -1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]bool)
    switch "hello" {
    case "":
        v8 := GenS1{F0: 884, F1: 127}
        _ = v8
    }
    var v9 int64 = 798
    v10 := GenS1{F0: (v9 << ((69) & 63)), F1: genFn0(4, -1, 0)}
    v11 := []int{}
    switch (("hello" + "a") + ((("\n\t" + "") + "") + "foo bar")) {
    case "":
        v9 -= -748
    case "foo bar":
        if (-922 != 815) {
            var v12 uint64 = 2
            v9 = v9
            _ = v12
        }
    }
    var v13 int16 = 547
    v9 -= (^(v9))
    if (1e10 != (-2.25 * 1e10)) {
        if (17 > 1) {
            _ = len(v2)
            v14 := GenS1{F0: -1, F1: 124}
            _ = v14
        } else {
            for fi15 := 0; fi15 < 15; fi15++ {
                v9 -= (^(v9))
                if true {
                    if (18446744073709551615 >= 2) {
                        _ = len(v5)
                        if true {
                            if (false || (false || (42 == 255))) {
                                if false {
                                    var v16 uint = 1
                                    _ = v16
                                } else {
                                    switch fi15 {
                                    case 9223372036854775807:
                                        v2[913] = (42 >= 584)
                                    case 127:
                                        if false {
                                            if len(v11) >= 2 {
                                                v11 = v11[1:len(v11)]
                                            }
                                        } else {
                                            fmt.Println(v9)
                                        }
                                    case -1:
                                        var v17 string = "foo bar"
                                        _ = v17
                                    }
                                }
                            }
                            _ = len(v1)
                        } else {
                            v13 -= 0
                            v13 += v13
                        }
                        var v18 string = "a"
                        _ = v18
                    }
                }
                _ = fi15
            }
            v9 = v9
            for fi19 := 0; fi19 < 5; fi19++ {
                if true {
                    v9 = v9
                }
                fmt.Println(fi19, v9)
                switch fi19 {
                case 0:
                    switch fi19 {
                    case 42:
                        v20 := GenS2{F0: true}
                        _ = v20
                    case 1:
                        var v21 float32 = 1e10
                        _ = v21
                    case 740:
                        fmt.Println(fi19, v9)
                    }
                case 4294967295:
                    fmt.Println(v9, v9, v9)
                default:
                    var v22 bool = false
                    _ = v22
                }
                _ = fi19
            }
        }
        if false {
            var v23 int8 = -128
            _ = len(v1)
            for fi24 := 0; fi24 < 8; fi24++ {
                v11 = append(v11, fi24)
                v13 = v13
                if len(v11) > 0 {
                    v11[0] = (-(fi24))
                }
                _ = fi24
            }
            _ = v23
        }
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        var v26 int = (fi25 & -244)
        v26 -= -618
        _ = fi25
        _ = v26
    }
    _ = len(v11)
    v27 := GenS1{F0: v9, F1: -128}
    var v28 uint8 = 128
    switch "a" {
    case "":
        var v29 uint = 67
        _ = v29
    case "a":
        var v30 uint32 = 400
        _ = v30
    default:
        v13 = v13
    }
    switch "foo bar" {
    case "foo bar":
        for fi31 := 0; fi31 < 3; fi31++ {
            var v32 uint16 = 2
            delete(v4, 642)
            v33 := GenS3{F0: 113}
            _ = fi31
            _ = v32
            _ = v33
        }
    default:
        var v34 uint = 0
        _ = v34
    }
    var v35 uint16 = 32767
    fmt.Println(v28)
    var v36 uint64 = 886
    if (1 >= v28) {
        for fi37 := 0; fi37 < 19; fi37++ {
            _ = len(v1)
            _ = fi37
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v9
}
