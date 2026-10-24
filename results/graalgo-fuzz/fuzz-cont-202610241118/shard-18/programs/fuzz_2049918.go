package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int64; F1 string; F2 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 int; F2 int64 }
type GenS6 struct { F0 int }

func main() {
    var v1 string = ""
    var v2 uint = 18446744073709551615
    var v3 float64 = (float64(1e10) / 1e10)
    switch 883 {
    case -643:
        for fi4 := 0; fi4 < 4; fi4++ {
            var v5 uint8 = 191
            switch ("a" + v1) {
            case "\n\t":
                fmt.Println(v1, v5, fi4)
            }
            _ = fi4
            _ = v5
        }
    case 42:
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 uint = 255
            _ = fi6
            _ = v7
        }
    }
    v2 = (v2 / ((2) | 1))
    v3 = v3
    v1 = v1
    var v8 uint8 = 235
    v1 = "a"
    v9 := GenS1{F0: 9223372036854775807, F1: 2147483648, F2: (9223372036854775807 < 42)}
    var v10 uint64 = 4294967295
    var v11 uint = 0
    for fi12 := 0; fi12 < 2; fi12++ {
        v11 -= (v11 << ((0) & 63))
        for fi13 := 0; fi13 < 20; fi13++ {
            if (1 > -10) {
                var v14 uint32 = 2147483647
                _ = v14
            }
            var v15 uint32 = 0
            _ = fi13
            _ = v15
        }
        _ = fi12
    }
    fmt.Println(v3, v1)
    v16 := GenS1{F0: 9223372036854775807, F1: 492, F2: true}
    if false {
        for fi17 := 0; fi17 < 13; fi17++ {
            v11 += v11
            v18 := GenS1{F0: 42, F1: 232, F2: false}
            _ = fi17
            _ = v18
        }
    } else {
        v10 = v10
    }
    var v19 uint64 = 4294967295
    if (false && true) {
        v20 := []int{}
        for fi21 := 0; fi21 < 20; fi21++ {
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 float64 = (1e10 * v3)
                _ = fi22
                _ = v23
            }
            switch ("a" + v1) {
            case "":
                if false {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v25 := GenS1{F0: -1, F1: 1, F2: (float32((v3 + ((3.14159 / 0.1) + 0.0))) >= 0.1)}
                        v26 := GenS1{F0: 0, F1: 0, F2: false}
                        var v27 uint64 = (v19 ^ v10)
                        switch (^(fi21)) {
                        case 9223372036854775807:
                            switch -9223372036854775808 {
                            case -1:
                                v1 = "\n\t"
                            case 2147483647:
                                switch int(v10) {
                                case 9223372036854775807:
                                    fmt.Println(v10, fi24)
                                }
                            case 32768:
                                switch v1 {
                                case "\n\t":
                                    switch "a" {
                                    case "":
                                        switch (v1 + v1) {
                                        case "\n\t":
                                            switch v1 {
                                            case "hello":
                                                v28 := GenS1{F0: 1, F1: 42, F2: false}
                                                _ = v28
                                            case "foo bar":
                                                v2 = (v11 * v11)
                                            case "a":
                                                v1 = v1
                                            }
                                        case "foo bar":
                                            var v29 int64 = 0
                                            _ = v29
                                        }
                                    case "hello":
                                        v8 -= v8
                                    }
                                case "a":
                                    v10 -= 388
                                case "foo bar":
                                    v30 := GenS2{F0: false, F1: 214}
                                    _ = v30
                                }
                            }
                        }
                        if false {
                            if false {
                                v31 := GenS3{F0: 42, F1: v1, F2: 2}
                                v32 := GenS4{F0: v1}
                                _ = v31
                                _ = v32
                            } else {
                                v10 -= (v19 + 18446744073709551615)
                            }
                            if true {
                                fmt.Println(v19, v19, v8)
                                v33 := GenS5{F0: 259, F1: (fi21 % ((fi21) | 1)), F2: 590}
                                _ = v33
                            }
                            fmt.Println(fi24, v3, v27)
                            v2 -= (v2 ^ (v2 << (((^(v2))) & 63)))
                            v3 += ((v3 / v3) * 0.1)
                            v1 = v1
                        } else {
                            v27 = 0
                        }
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                } else {
                    v19 = 528
                }
            default:
                if (-768 > -53) {
                    if true {
                        var v34 float32 = ((0.1 / 1.0) + -2.25)
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 12; fi35++ {
                            v36 := GenS2{F0: false, F1: 100}
                            v37 := GenS6{F0: fi21}
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                    }
                }
            }
            _ = fi21
        }
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = v20
    } else {
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
