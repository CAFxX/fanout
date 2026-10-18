package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
func genFn0() int8 {
    var v7 int8 = 1
    v8 := []int{0}
    var v9 uint16 = 42
    _ = v7
    _ = v8
    _ = v9
    return v7
}

func main() {
    v1 := []int{4294967295, 32767, 32768, 0}
    v2 := []string{}
    v3 := make(map[uint32]bool)
    var v4 int64 = 1
    var v5 uint16 = 208
    var v6 float32 = 1e10
    switch 278 {
    case 4294967295:
        v6 -= v6
    default:
        fmt.Println(v6)
    }
    var v10 uint8 = 71
    v11 := GenS1{F0: 42, F1: (v5 + 65535)}
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 int64 = (v4 ^ 0)
        var v14 int32 = 42
        _ = fi12
        _ = v13
        _ = v14
    }
    v10 = uint8(v5)
    var v15 uint = 18446744073709551615
    fmt.Println(v5, v10)
    v15 = (^(v15))
    v5 -= v5
    if false {
        var v16 int = 32767
        _ = v16
    }
    switch ("hello" + "\n\t") {
    case "":
        if false {
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 uint16 = v5
                for fi19 := 0; fi19 < 16; fi19++ {
                    if (!true) {
                        if (-1 < 557) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            v15 = v15
                        }
                        switch (("a" + ("hello" + "\n\t")) + ((((("hello" + ("hello" + ("hello" + ""))) + "\n\t") + "foo bar") + "hello") + "hello")) {
                        case "foo bar":
                            var v20 string = ""
                            _ = v20
                        case "hello":
                            v18 += 255
                        case "":
                            fmt.Println(v18)
                        default:
                            switch "a" {
                            case "a":
                                v15 -= v15
                            case "\n\t":
                                v6 -= (v6 * v6)
                            }
                            switch "a" {
                            case "foo bar":
                                switch fi17 {
                                case 9223372036854775807:
                                    if false {
                                        v4 -= (v4 & (v4 & (v4 >> ((v15) & 63))))
                                        fmt.Println(v10)
                                        var v21 uint64 = uint64(v15)
                                        var v22 bool = (0 == 2)
                                        v23 := v3[14]
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    }
                                }
                            case "":
                                fmt.Println(v5)
                            }
                            v18 = v5
                            fmt.Println(v15, v15)
                        }
                        v24 := GenS1{F0: int32(v18), F1: v5}
                        v4 -= 9223372036854775807
                        if (19 == 65535) {
                            v6 = (0.0 - v6)
                            fmt.Println(v15)
                        }
                        _ = v24
                    }
                    _ = fi19
                }
                _ = fi17
                _ = v18
            }
        } else {
            for fi25 := 0; fi25 < 6; fi25++ {
                for fi26 := 0; fi26 < 19; fi26++ {
                    if false {
                        v10 = 255
                    }
                    var v27 uint8 = (^(v10))
                    var v28 bool = true
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = fi25
            }
            v29 := GenS1{F0: int32(v5), F1: 32767}
            _ = v29
        }
    case "hello":
        var v30 float32 = 1.5
        _ = v30
    default:
        v10 = v10
    }
    var v31 int64 = v4
    v5 = v5
    var v32 int32 = -1
    var v33 int32 = 0
    for fi34 := 0; fi34 < 3; fi34++ {
        v4 += (v31 % ((v31) | 1))
        _ = fi34
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        v36 := GenS1{F0: 42, F1: (v5 * v5)}
        fmt.Println(v15)
        v4 -= v4
        _ = fi35
        _ = v36
    }
    v3[1] = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
