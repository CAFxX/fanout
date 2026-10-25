package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int64 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[int]string)
    v2 := []int{1, 2147483647, -9223372036854775808, 32767}
    var v3 int = 635
    var v4 int32 = 2147483647
    var v5 uint = 2
    switch "\n\t" {
    case "foo bar":
        var v6 uint32 = 571
        _ = v6
    case "a":
        v4 += v4
    case "hello":
        v7 := GenS1{F0: false, F1: -313, F2: 285}
        _ = v7
    default:
        switch (("foo bar" + "\n\t") + ("hello" + "\n\t")) {
        case "a":
            if false {
                fmt.Println(v3, v4, v3)
                for fi8 := 0; fi8 < 8; fi8++ {
                    switch v3 {
                    case -531:
                        if (128 < 215) {
                            if false {
                                var v9 uint8 = uint8(v5)
                                _ = v9
                            } else {
                                var v10 int = v3
                                _ = len(v2)
                                _ = v10
                            }
                        } else {
                            switch "foo bar" {
                            case "":
                                var v11 uint64 = 1
                                _ = v11
                            case "foo bar":
                                var v12 uint32 = 419
                                _ = v12
                            }
                        }
                    case -719:
                        var v13 float32 = (float32((0.0 - 3.14159)) / 1.0)
                        _ = v13
                    case 4294967295:
                        switch "" {
                        case "\n\t":
                            for fi14 := 0; fi14 < 12; fi14++ {
                                v2 = append(v2, fi8)
                                var v15 uint64 = 4294967295
                                fmt.Println(v3)
                                _ = fi14
                                _ = v15
                            }
                        case "foo bar":
                            v3 = (v3 - v3)
                        default:
                            v4 = (-(v4))
                        }
                    default:
                        var v16 int32 = (^(v4))
                        _ = v16
                    }
                    _ = fi8
                }
            }
        case "foo bar":
            if false {
                for fi17 := 0; fi17 < 18; fi17++ {
                    v4 += 2147483647
                    v18 := GenS2{F0: 379}
                    if (!false) {
                        v4 = v4
                        for fi19 := 0; fi19 < 7; fi19++ {
                            var v20 uint64 = 4294967295
                            _ = fi19
                            _ = v20
                        }
                        fmt.Println(v5)
                    }
                    _ = fi17
                    _ = v18
                }
                var v21 uint32 = 65535
                _ = v21
            } else {
                for fi22 := 0; fi22 < 13; fi22++ {
                    fmt.Println(v3, v3)
                    if true {
                        v1[-2147483648] = "hello"
                        switch (fi22 & 32768) {
                        case -394:
                            for fi23 := 0; fi23 < 15; fi23++ {
                                fmt.Println(v5, v5, v4)
                                v4 = (v4 >> ((v5) & 63))
                                _ = fi23
                            }
                        case -1:
                            if (18446744073709551615 < 472) {
                                fmt.Println(fi22, v3)
                                fmt.Println(v3, v5, v4)
                                var v24 int = 127
                                _ = v24
                            } else {
                                v25 := v1[(^(v3))]
                                _ = v25
                            }
                        }
                    }
                    _ = fi22
                }
            }
        case "hello":
            delete(v1, 978)
        }
    }
    var v26 int64 = -1
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 uint16 = 2
        fmt.Println(v28, v5)
        v26 += (^(v26))
        _ = fi27
        _ = v28
    }
    delete(v1, -243)
    fmt.Println(v4, v26)
    var v29 uint8 = 42
    var v30 int8 = 36
    fmt.Println(v30, v3, v30)
    v29 = 63
    var v31 uint32 = 42
    var v32 int32 = v4
    v2 = append(v2, v3)
    if true {
        var v33 uint32 = v31
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
