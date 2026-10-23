package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint16; F1 int32; F2 int }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int8; F1 int8; F2 int16 }
func genFn0() int32 {
    if ((true || (true && false)) && false) {
        var v4 int16 = -1
        _ = v4
    } else {
        if true {
            if false {
                v5 := []int{}
                var v6 string = "foo bar"
                _ = v5
                _ = v6
            }
        } else {
            v7 := []int{}
            var v8 int8 = 127
            fmt.Println(v8)
            _ = v7
            _ = v8
        }
    }
    var v9 uint16 = 0
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint64 = 1
    var v2 string = ""
    var v3 int = -2147483648
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int8 = 101
        _ = fi10
        _ = v11
    }
    if (!true) {
        switch "hello" {
        case "foo bar":
            v1 += (v1 & (v1 | (v1 << ((0) & 63))))
        case "hello":
            switch "hello" {
            case "":
                switch "a" {
                case "foo bar":
                    if false {
                        v3 = 32768
                        var v12 uint8 = 127
                        for fi13 := 0; fi13 < 3; fi13++ {
                            if false {
                                v3 += 829
                                v3 = (fi13 ^ fi13)
                                var v14 uint64 = (v1 << ((255) & 63))
                                _ = v14
                            }
                            v15 := []int{fi13, v3}
                            fmt.Println(v1)
                            _ = fi13
                            _ = v15
                        }
                        _ = v12
                    } else {
                        v16 := GenS1{F0: 16}
                        var v17 uint16 = 65535
                        _ = v16
                        _ = v17
                    }
                }
            default:
                for fi18 := 0; fi18 < 16; fi18++ {
                    v2 = (v2 + v2)
                    _ = fi18
                }
            }
        case "a":
            fmt.Println(v2, v1, v3)
        }
        var v19 uint16 = 2
        _ = v19
    } else {
        fmt.Println(v3)
    }
    v20 := []int{(v3 & 127)}
    v2 = (v2 + (("" + (v2 + "\n\t")) + "\n\t"))
    _ = len(v20)
    var v21 float32 = 0.1
    if false {
        if true {
            v22 := GenS2{F0: (v1 + v1), F1: 9223372036854775807}
            fmt.Println(v21)
            _ = v22
        } else {
            fmt.Println(v2, v1, v21)
            _ = len(v20)
            v3 = 32767
        }
    } else {
        fmt.Println(v3, v21, v21)
        for fi23 := 0; fi23 < 10; fi23++ {
            switch v2 {
            case "hello":
                var v24 uint16 = 65535
                _ = v24
            case "":
                var v25 int = 0
                _ = v25
            }
            v26 := GenS1{F0: int64(v1)}
            if (2 < 42) {
                switch "" {
                case "\n\t":
                    if (4294967295 >= (v1 / ((18446744073709551615) | 1))) {
                        for fi27 := 0; fi27 < 7; fi27++ {
                            var v28 uint16 = 42
                            var v29 uint8 = 86
                            v30 := GenS3{F0: 2, F1: 42, F2: v3}
                            _ = fi27
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    } else {
                        v21 = v21
                        v31 := GenS4{F0: 233, F1: 459}
                        switch v2 {
                        case "a":
                            _ = len(v20)
                        case "\n\t":
                            v2 = "foo bar"
                        default:
                            if true {
                                v1 += (v1 + v1)
                                var v32 uint = 635
                                _ = v32
                            }
                        }
                        _ = v31
                    }
                case "a":
                    var v33 uint32 = 309
                    _ = v33
                default:
                    if ((false || false) || true) {
                        v34 := GenS2{F0: (v1 * v1), F1: 1}
                        v35 := GenS3{F0: 255, F1: 1, F2: (-(v3))}
                        fmt.Println(v3, v2, v3)
                        _ = v34
                        _ = v35
                    }
                }
            }
            _ = fi23
            _ = v26
        }
    }
    switch v2 {
    case "\n\t":
        v2 = (v2 + "a")
    case "foo bar":
        v21 += float32(0.1)
    }
    var v36 bool = (42 < 9223372036854775807)
    v37 := GenS5{F0: -128, F1: 87, F2: -517}
    var v38 int64 = 9223372036854775807
    v20 = append(v20, v3)
    if v36 {
        v39 := GenS2{F0: v1, F1: v38}
        if (1 != 635) {
            var v40 uint16 = 32767
            if (246 < 2) {
                switch v2 {
                case "foo bar":
                    v20 = append(v20, v3)
                default:
                    for fi41 := 0; fi41 < 14; fi41++ {
                        var v42 uint16 = v40
                        _ = fi41
                        _ = v42
                    }
                }
                if (true || ("\n\t" != "a")) {
                    for fi43 := 0; fi43 < 11; fi43++ {
                        _ = len(v20)
                        _ = fi43
                    }
                } else {
                    v21 = 0.0
                }
            }
            _ = v40
        }
        var v44 uint8 = 251
        _ = v39
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}
