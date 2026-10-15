package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int = 1
    var v2 uint16 = 2
    var v3 int16 = 32767
    v4 := make(map[string]int)
    v5 := []int{489, (v1 << ((4294967295) & 63))}
    var v6 float64 = 1e10
    fmt.Println(v6, v1)
    fmt.Println(v6, v3, v2)
    var v7 uint = 0
    var v8 uint = (^(v7))
    delete(v4, ("" + ""))
    v8 += 287
    switch (v1 - -9223372036854775808) {
    case -958:
        v9 := GenS1{F0: 0, F1: (v7 | 670), F2: false}
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v8 += v8
        _ = fi10
    }
    v11 := GenS1{F0: 1, F1: v7, F2: (((v6 > v6) || true) && (42 != 1))}
    v12 := GenS1{F0: 9223372036854775807, F1: (v7 / (((^(v7))) | 1)), F2: (v2 > (v2 % ((v2) | 1)))}
    if (true || true) {
        v13 := GenS1{F0: -9223372036854775808, F1: (^(v7)), F2: true}
        _ = v13
    } else {
        if (!((^(v8)) <= v8)) {
            var v14 float64 = (float64((0.0 + -2.25)) + ((v6 / v6) / 1.5))
            switch "foo bar" {
            case "hello":
                var v15 uint64 = 0
                _ = v15
            case "\n\t":
                for fi16 := 0; fi16 < 5; fi16++ {
                    if true {
                        switch "hello" {
                        case "hello":
                            var v17 float64 = ((v14 + 1.5) * v6)
                            _ = v17
                        case "":
                            fmt.Println(v14, v2)
                        case "a":
                            if len(v5) > 0 {
                                v5[0] = (fi16 + (fi16 << ((270) & 63)))
                            }
                        default:
                            for fi18 := 0; fi18 < 13; fi18++ {
                                v6 = 0.1
                                v3 += v3
                                _ = fi18
                            }
                        }
                        var v19 uint8 = 128
                        v7 = (v8 << (((v8 >> ((v7) & 63))) & 63))
                        _ = v19
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            var v20 int64 = 42
                            _ = v20
                        case "a":
                            v21 := GenS2{F0: 32767}
                            _ = v21
                        case "":
                            switch (v1 / (((v1 ^ v1)) | 1)) {
                            case 9223372036854775807:
                                switch ("" + "a") {
                                case "":
                                    v4["a"] = fi16
                                case "\n\t":
                                    switch "hello" {
                                    case "a":
                                        for fi22 := 0; fi22 < 14; fi22++ {
                                            var v23 float32 = (3.14159 / 1e10)
                                            var v24 int = fi22
                                            var v25 int64 = -864
                                            _ = fi22
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                    case "foo bar":
                                        switch ("hello" + (("" + "a") + (((("foo bar" + "\n\t") + "") + "hello") + ("hello" + "a")))) {
                                        case "foo bar":
                                            v26 := GenS3{F0: 42}
                                            _ = v26
                                        }
                                    case "hello":
                                        for fi27 := 0; fi27 < 16; fi27++ {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = fi27
                                        }
                                    default:
                                        v28 := GenS2{F0: v2}
                                        _ = v28
                                    }
                                default:
                                    _ = len(v5)
                                }
                            case 32767:
                                if (127 > 2) {
                                    fmt.Println(v3, fi16, v3)
                                }
                            default:
                                var v29 uint64 = 1
                                _ = v29
                            }
                        }
                        var v30 uint32 = 831
                        _ = v30
                    }
                    _ = fi16
                }
            default:
                v8 = v7
            }
            _ = v14
        }
        fmt.Println(v6, v2, v6)
        var v31 uint64 = 42
        _ = v31
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        var v33 string = "foo bar"
        _ = fi32
        _ = v33
    }
    var v34 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
