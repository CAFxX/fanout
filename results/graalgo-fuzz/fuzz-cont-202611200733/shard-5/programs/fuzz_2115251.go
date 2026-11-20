package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int; F1 uint8 }

func main() {
    var v1 int16 = 0
    var v2 uint64 = 2
    var v3 uint = 1
    v4 := make(map[uint32]bool)
    var v5 uint32 = uint32(v3)
    v6 := GenS1{F0: -1}
    v3 = (v3 << ((42) & 63))
    fmt.Println(v1, v5, v2)
    v5 = v5
    var v7 uint64 = 487
    var v8 int = 0
    v2 = (v7 / ((v2) | 1))
    var v9 string = "hello"
    var v10 int64 = 202
    v11 := []int{(-(v8)), 12}
    switch 0 {
    case 32768:
        v9 = "a"
    case -836:
        v12 := GenS1{F0: 0}
        _ = v12
    default:
        fmt.Println(v10, v7, v5)
    }
    if (false || true) {
        v10 += (-(v10))
        v13 := GenS1{F0: 591}
        v14 := GenS2{F0: false}
        _ = v13
        _ = v14
    }
    if true {
        if (247 > 0) {
            if (-1 >= 0) {
                if (true && true) {
                    _ = len(v4)
                    var v15 uint = uint(v2)
                    _ = v15
                }
                if ((!true) && true) {
                    delete(v4, v5)
                    v16 := GenS1{F0: int32(v5)}
                    _ = v16
                }
            } else {
                var v17 float64 = 0.1
                switch v9 {
                case "foo bar":
                    delete(v4, v5)
                default:
                    for fi18 := 0; fi18 < 7; fi18++ {
                        switch "foo bar" {
                        case "\n\t":
                            for fi19 := 0; fi19 < 10; fi19++ {
                                _ = len(v11)
                                var v20 uint = 692
                                if (187 >= 33) {
                                    if (!true) {
                                        var v21 uint = v20
                                        _ = v21
                                    } else {
                                        v3 = v3
                                        v5 -= (^(v5))
                                        v17 += v17
                                    }
                                    var v22 float64 = v17
                                    var v23 int16 = v1
                                    _ = v22
                                    _ = v23
                                }
                                v10 = 9223372036854775807
                                _ = fi19
                                _ = v20
                            }
                        case "foo bar":
                            v24 := v4[(v5 - (v5 << ((65535) & 63)))]
                            _ = v24
                        case "hello":
                            if len(v11) > 0 {
                                v11[0] = (^(fi18))
                            }
                        }
                        _ = fi18
                    }
                }
                _ = v17
            }
            var v25 uint16 = 2
            for fi26 := 0; fi26 < 6; fi26++ {
                if true {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        if len(v11) > 0 {
                            v11[0] = v8
                        }
                        _ = fi27
                    }
                    switch -89 {
                    case -1:
                        for fi28 := 0; fi28 < 20; fi28++ {
                            switch 717 {
                            case 9223372036854775807:
                                v29 := v4[(^(v5))]
                                v30 := GenS2{F0: (!((0 < 0) || ((v2 - (v2 * (v7 ^ 0))) >= v2)))}
                                _ = v29
                                _ = v30
                            case -172:
                                if (true || false) {
                                    var v31 uint8 = 42
                                    fmt.Println(fi28, v9)
                                    if (!(v10 < -9223372036854775808)) {
                                        v9 = "foo bar"
                                        var v32 uint32 = v5
                                        var v33 int64 = (v10 | v10)
                                        _ = v32
                                        _ = v33
                                    } else {
                                        v5 -= v5
                                    }
                                    _ = v31
                                } else {
                                    if (v9 < "hello") {
                                        v8 = v8
                                        if ((^(v3)) > uint(v5)) {
                                            v34 := GenS3{F0: fi28, F1: v1}
                                            var v35 int8 = 127
                                            v1 += v1
                                            _ = v34
                                            _ = v35
                                        }
                                        fmt.Println(v3, v2, v2)
                                    } else {
                                        var v36 uint64 = 120
                                        _ = v36
                                    }
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                }
                            default:
                                var v37 int = (v8 ^ v8)
                                _ = v37
                            }
                            _ = fi28
                        }
                    case 9223372036854775807:
                        var v38 int = v8
                        _ = v38
                    }
                    _ = len(v11)
                }
                if false {
                    if true {
                        if len(v11) > 0 {
                            v11[0] = (-(v8))
                        }
                    }
                    v11 = append(v11, fi26)
                } else {
                    fmt.Println(v9)
                    v3 = 236
                }
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi26
            }
            _ = v25
        } else {
            v39 := GenS4{F0: 2, F1: 1}
            _ = v39
        }
        switch (^(v8)) {
        case 128:
            var v40 uint8 = 1
            _ = v40
        }
    }
    var v41 int16 = v1
    v42 := GenS5{F0: v8, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
