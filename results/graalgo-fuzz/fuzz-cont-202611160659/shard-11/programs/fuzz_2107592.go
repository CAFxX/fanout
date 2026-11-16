package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string }

func main() {
    v1 := []int{127, 128}
    var v2 int8 = 1
    var v3 int = 16
    v4 := make(map[uint32]string)
    v5 := []uint32{198, 2, 210, 2147483647}
    v2 += v2
    if len(v1) > 0 {
        v1[0] = (^(v3))
    }
    fmt.Println(v2, v3)
    v6 := GenS1{F0: 42}
    var v7 uint16 = 0
    switch 348 {
    case 42:
        v7 = v7
    default:
        delete(v4, 2147483647)
    }
    var v8 int64 = 42
    v9 := v4[2]
    switch int(v2) {
    case 536:
        switch v9 {
        case "a":
            fmt.Println(v8, v7, v3)
        case "\n\t":
            delete(v4, 65535)
        default:
            for fi10 := 0; fi10 < 3; fi10++ {
                if (false || (1 != 18446744073709551615)) {
                    var v11 float64 = 0.1
                    _ = len(v5)
                    for fi12 := 0; fi12 < 19; fi12++ {
                        if true {
                            var v13 int = (v3 - v3)
                            _ = len(v4)
                            delete(v4, 42)
                            _ = v13
                        }
                        v7 = (^(v7))
                        _ = len(v4)
                        fmt.Println(fi12, v2, v11)
                        _ = fi12
                    }
                    _ = v11
                }
                if (412 > 1) {
                    var v14 int8 = (v2 % ((-128) | 1))
                    v4[2147483647] = "foo bar"
                    _ = v14
                } else {
                    delete(v4, 2147483648)
                    var v15 int = 597
                    v16 := GenS1{F0: 166}
                    _ = v15
                    _ = v16
                }
                if (!false) {
                    fmt.Println(v2, v2, v7)
                } else {
                    v17 := GenS1{F0: 801}
                    v18 := GenS1{F0: 42}
                    switch (v3 | v3) {
                    case 965:
                        var v19 bool = true
                        _ = v19
                    case 42:
                        for fi20 := 0; fi20 < 15; fi20++ {
                            var v21 string = v9
                            _ = fi20
                            _ = v21
                        }
                    case -9223372036854775808:
                        for fi22 := 0; fi22 < 10; fi22++ {
                            v3 = -195
                            if (!true) {
                                switch "foo bar" {
                                case "\n\t":
                                    var v23 int = (^(v3))
                                    _ = v23
                                case "":
                                    _ = len(v5)
                                case "foo bar":
                                    switch fi22 {
                                    case 1:
                                        v24 := GenS1{F0: 42}
                                        _ = v24
                                    case 32767:
                                        v9 = "hello"
                                    }
                                    if (0 != v8) {
                                        v25 := GenS2{F0: 0}
                                        fmt.Println(fi22)
                                        v2 = v2
                                        _ = v25
                                    } else {
                                        var v26 int64 = (^(v8))
                                        v7 = 0
                                        v27 := v4[2147483648]
                                        _ = v26
                                        _ = v27
                                    }
                                default:
                                    var v28 float64 = ((-2.25 - 0.0) - 1e10)
                                    _ = v28
                                }
                                v9 = "foo bar"
                                fmt.Println(v2, v8)
                            } else {
                                var v29 bool = (true && false)
                                _ = v29
                            }
                            if true {
                                v30 := GenS3{F0: (v3 | (fi10 & (fi10 + -1)))}
                                v2 -= (v2 - v2)
                                _ = v30
                            } else {
                                if ((true || (-933 <= int32(v2))) || (-1 <= -155)) {
                                    switch ("hello" + "foo bar") {
                                    case "a":
                                        v9 = v9
                                    }
                                }
                                v31 := GenS4{F0: "foo bar"}
                                var v32 uint16 = 42
                                _ = v31
                                _ = v32
                            }
                            _ = fi22
                        }
                    default:
                        v33 := GenS4{F0: v9}
                        _ = v33
                    }
                    _ = v17
                    _ = v18
                }
                _ = fi10
            }
        }
    case -1:
        for fi34 := 0; fi34 < 12; fi34++ {
            delete(v4, 370)
            v8 += (v8 * (v8 / ((-1) | 1)))
            for fi35 := 0; fi35 < 18; fi35++ {
                v7 = (v7 & 255)
                v36 := v4[42]
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
    }
    var v37 int8 = (v2 / (((v2 % ((v2) | 1))) | 1))
    for fi38 := 0; fi38 < 8; fi38++ {
        v39 := GenS3{F0: v3}
        v40 := GenS4{F0: ((v9 + v9) + (v9 + v9))}
        var v41 int64 = v8
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    var v42 bool = (!((v7 | v7) == (^(v7))))
    if v42 {
        var v43 bool = v42
        _ = v43
    }
    fmt.Println(v37, v42)
    var v44 uint64 = 4294967295
    v45 := GenS3{F0: v3}
    v3 -= (v3 | v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v42
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
