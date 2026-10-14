package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint; F1 bool; F2 uint64 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 32767
    v3 := make(map[int]string)
    v4 := []int{}
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(v2, v2)
        _ = fi5
    }
    v6 := GenS1{F0: 88}
    v7 := GenS1{F0: 128}
    v2 += v2
    fmt.Println(v2)
    var v8 uint8 = 237
    v9 := GenS1{F0: 205}
    switch 127 {
    case 32767:
        v10 := GenS1{F0: 7}
        _ = v10
    case 651:
        for fi11 := 0; fi11 < 2; fi11++ {
            v2 = 901
            _ = fi11
        }
    default:
        if (0 > -194) {
            delete(v1, ("a" + "\n\t"))
            if (42 > -1) {
                if (127 > 88) {
                    switch "a" {
                    case "":
                        for fi12 := 0; fi12 < 15; fi12++ {
                            switch -2147483648 {
                            case -975:
                                if true {
                                    for fi13 := 0; fi13 < 11; fi13++ {
                                        var v14 uint8 = v8
                                        var v15 float32 = (float32((0.0 + 0.1)) + 0.1)
                                        v16 := GenS1{F0: 0}
                                        var v17 int8 = 25
                                        _ = fi13
                                        _ = v14
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                    }
                                } else {
                                    _ = len(v3)
                                    v2 -= -377
                                }
                            }
                            var v18 uint = 65535
                            v19 := GenS1{F0: v8}
                            _ = fi12
                            _ = v18
                            _ = v19
                        }
                    case "hello":
                        switch 0 {
                        case -708:
                            v8 += v8
                        case 97:
                            if len(v4) > 0 {
                                v4[0] = 155
                            }
                        }
                    case "foo bar":
                        var v20 int16 = v2
                        _ = v20
                    default:
                        var v21 uint8 = v8
                        _ = v21
                    }
                    v8 -= (^(v8))
                    var v22 string = "hello"
                    _ = v22
                }
                var v23 float64 = ((3.14159 - 1e10) * (1e10 + float64(0.1)))
                for fi24 := 0; fi24 < 6; fi24++ {
                    if false {
                        v8 -= v8
                        v23 = v23
                        if false {
                            v2 = (-(v2))
                        }
                    } else {
                        v25 := GenS2{F0: -122, F1: 65535}
                        _ = v25
                    }
                    switch -2147483648 {
                    case 2147483647:
                        switch -141 {
                        case 32767:
                            v2 += v2
                        case 128:
                            switch -548 {
                            case 233:
                                var v26 uint16 = 1
                                _ = v26
                            case 903:
                                var v27 int = (fi24 / ((fi24) | 1))
                                _ = v27
                            case -699:
                                v28 := GenS3{F0: 2, F1: (733 < 60), F2: 513}
                                _ = v28
                            }
                        }
                    }
                    delete(v3, 42)
                    _ = fi24
                }
                _ = v23
            } else {
                var v29 uint16 = 255
                _ = v29
            }
        } else {
            if (true || false) {
                v30 := GenS4{F0: 42}
                _ = v30
            }
        }
    }
    fmt.Println(v8, v2)
    var v31 float64 = 0.0
    v4 = append(v4, -2147483648)
    if true {
        var v32 uint8 = 255
        _ = v32
    } else {
        var v33 int64 = 646
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
