package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint32; F2 bool }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 1
    var v3 int32 = 42
    v4 := []int{}
    if len(v4) > 0 {
        v4[0] = 127
    }
    var v5 uint64 = v2
    v6 := GenS1{F0: 97}
    v7 := GenS1{F0: 238}
    var v8 uint16 = 645
    v8 = v8
    for fi9 := 0; fi9 < 12; fi9++ {
        if (255 > 42) {
            for fi10 := 0; fi10 < 4; fi10++ {
                if false {
                    fmt.Println(fi9)
                    var v11 uint64 = (v2 >> ((324) & 63))
                    var v12 string = "a"
                    _ = v11
                    _ = v12
                } else {
                    var v13 uint = 104
                    if (true && false) {
                        var v14 float64 = 1e10
                        var v15 int32 = (v3 + 42)
                        v16 := GenS1{F0: 255}
                        fmt.Println(v2, v14)
                        if false {
                            var v17 int16 = 1000
                            _ = len(v4)
                            _ = v17
                        } else {
                            fmt.Println(fi10, v8)
                            var v18 int32 = v15
                            if true {
                                var v19 int8 = 0
                                v20 := GenS1{F0: 1}
                                v14 = (v14 + 1.5)
                                v3 = v15
                                switch "hello" {
                                case "a":
                                    var v21 int = 127
                                    _ = v21
                                case "\n\t":
                                    var v22 float32 = 0.1
                                    _ = v22
                                default:
                                    v23 := GenS2{F0: 42, F1: uint32(v8), F2: (192 > 401)}
                                    _ = v23
                                }
                                _ = v19
                                _ = v20
                            }
                            _ = v18
                        }
                        var v24 uint16 = v8
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v24
                    }
                    if false {
                        v13 += (v13 ^ 4294967295)
                        var v25 int8 = -116
                        switch 32768 {
                        case 733:
                            _ = len(v1)
                        case 127:
                            if (int16(v25) > -907) {
                                var v26 uint = v13
                                _ = v26
                            } else {
                                v27 := GenS3{F0: ("hello" + "")}
                                v28 := GenS2{F0: -1, F1: 2147483647, F2: false}
                                _ = v27
                                _ = v28
                            }
                        default:
                            var v29 bool = (2 != 711)
                            _ = v29
                        }
                        _ = v25
                    }
                    _ = v13
                }
                if false {
                    var v30 uint8 = 0
                    if len(v4) > 0 {
                        v4[0] = 1
                    }
                    _ = v30
                } else {
                    v2 = 42
                }
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = len(v4)
                }
                _ = fi10
            }
        }
        _ = fi9
    }
    fmt.Println(v3)
    for fi31 := 0; fi31 < 8; fi31++ {
        if ("hello" != (("" + ("a" + "foo bar")) + "a")) {
            var v32 uint32 = 143
            var v33 bool = false
            _ = v32
            _ = v33
        } else {
            var v34 uint8 = 110
            var v35 uint8 = 2
            _ = v34
            _ = v35
        }
        _ = fi31
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        var v37 bool = true
        _ = fi36
        _ = v37
    }
    var v38 string = ((("a" + (("hello" + "hello") + ("" + "a"))) + "") + "\n\t")
    var v39 int16 = 575
    v8 += 255
    switch v38 {
    case "foo bar":
        switch -2147483648 {
        case 563:
            v2 += v5
        default:
            for fi40 := 0; fi40 < 8; fi40++ {
                fmt.Println(v2, v5)
                _ = fi40
            }
        }
    case "hello":
        var v41 int64 = -56
        _ = v41
    case "a":
        for fi42 := 0; fi42 < 4; fi42++ {
            fmt.Println(v38, v39, v5)
            _ = fi42
        }
    default:
        v8 = 2
    }
    var v43 uint8 = 255
    v38 = "\n\t"
    if true {
        var v44 float32 = 1e10
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
