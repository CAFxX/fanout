package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint64 }

func main() {
    var v1 string = "\n\t"
    var v2 string = "\n\t"
    var v3 int32 = -387
    var v4 int64 = 0
    v5 := GenS1{F0: 482, F1: 4294967295}
    v6 := []int{-873}
    var v7 float64 = 3.14159
    var v8 uint8 = 42
    var v9 int16 = 42
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(v1, fi10)
        var v11 int8 = 63
        switch ("a" + "\n\t") {
        case "foo bar":
            var v12 int16 = -955
            _ = v12
        case "a":
            if (false && true) {
                var v13 bool = true
                switch "hello" {
                case "foo bar":
                    var v14 uint64 = 1
                    _ = v14
                case "":
                    if len(v6) > 0 {
                        v6[0] = (fi10 + 1)
                    }
                case "\n\t":
                    fmt.Println(v9)
                }
                _ = v13
            } else {
                if ((-(v3)) < v3) {
                    var v15 float64 = (float64((0.1 - 0.0)) / ((v7 / -2.25) - v7))
                    v6 = append(v6, (fi10 % ((127) | 1)))
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v17 := GenS1{F0: 65535, F1: 2}
                        _ = fi16
                        _ = v17
                    }
                    _ = v15
                } else {
                    for fi18 := 0; fi18 < 10; fi18++ {
                        switch fi10 {
                        case -2147483648:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        default:
                            v19 := GenS1{F0: 255, F1: 66}
                            _ = v19
                        }
                        _ = fi18
                    }
                    switch fi10 {
                    case 4294967295:
                        for fi20 := 0; fi20 < 13; fi20++ {
                            if ((3.14159 / 1.0) <= 1e10) {
                                var v21 uint32 = 1
                                var v22 int32 = (v3 >> ((42) & 63))
                                _ = v21
                                _ = v22
                            }
                            v6 = append(v6, fi20)
                            v23 := GenS2{F0: (^(v3))}
                            _ = fi20
                            _ = v23
                        }
                    default:
                        v2 = ("\n\t" + v2)
                    }
                }
                switch ("hello" + "") {
                case "a":
                    var v24 int8 = v11
                    _ = v24
                case "":
                    _ = len(v6)
                case "hello":
                    switch "" {
                    case "\n\t":
                        switch v2 {
                        case "foo bar":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case "\n\t":
                            var v25 uint32 = 2
                            _ = v25
                        default:
                            var v26 int64 = -294
                            _ = v26
                        }
                    case "hello":
                        var v27 float32 = 3.14159
                        _ = v27
                    case "foo bar":
                        v7 -= 0.0
                    }
                }
            }
        case "hello":
            v7 += v7
        }
        _ = fi10
        _ = v11
    }
    switch "" {
    case "\n\t":
        v28 := GenS2{F0: (v3 << ((42) & 63))}
        _ = v28
    case "foo bar":
        switch ((v2 + ("hello" + v1)) + v1) {
        case "\n\t":
            v9 = v9
        case "foo bar":
            fmt.Println(v4)
        case "":
            v29 := GenS3{F0: 730, F1: false, F2: 334}
            _ = v29
        default:
            var v30 int8 = 1
            _ = v30
        }
    case "hello":
        for fi31 := 0; fi31 < 9; fi31++ {
            v32 := GenS2{F0: (v3 & v3)}
            var v33 bool = false
            v6 = append(v6, (fi31 / (((fi31 << ((924) & 63))) | 1)))
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    var v34 uint16 = 254
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
