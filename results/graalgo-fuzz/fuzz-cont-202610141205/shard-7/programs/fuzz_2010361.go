package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int8; F1 int16; F2 uint16 }
func genFn0() int16 {
    v6 := GenS1{F0: "", F1: 42, F2: 748}
    switch ("a" + "foo bar") {
    case "":
        var v7 uint8 = 2
        _ = v7
    }
    var v8 int64 = -1
    switch "" {
    case "":
        var v9 int64 = (v8 - v8)
        _ = v9
    }
    _ = v6
    _ = v8
    return -957
}

func main() {
    var v1 uint = 318
    var v2 int64 = 9223372036854775807
    var v3 uint = 383
    var v4 int8 = 0
    v5 := []int{-495, 127, -2147483648, -197}
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "foo bar":
        var v10 uint16 = 233
        _ = v10
    }
    v2 -= (v2 * 1)
    var v11 bool = (!(!(100 == (-(v4)))))
    if (!((v11 && v11) || (!(false && v11)))) {
        if v11 {
            var v12 string = ""
            v4 = 100
            _ = v12
        } else {
            v13 := GenS2{F0: v3, F1: 0, F2: (v2 + -123)}
            v2 = (v2 - v2)
            v14 := GenS1{F0: "\n\t", F1: 2, F2: 1}
            _ = v13
            _ = v14
        }
    }
    var v15 int16 = 1
    fmt.Println(v2, v3, v15)
    fmt.Println(v3, v4)
    if v11 {
        if (!v11) {
            _ = len(v5)
            if (!(v11 && v11)) {
                var v16 uint16 = 255
                for fi17 := 0; fi17 < 6; fi17++ {
                    for fi18 := 0; fi18 < 14; fi18++ {
                        if len(v5) > 0 {
                            v5[0] = (fi17 * (fi17 * -9223372036854775808))
                        }
                        if v11 {
                            if len(v5) > 0 {
                                v5[0] = fi18
                            }
                            v19 := GenS3{F0: 1, F1: (v15 | (v15 - (v15 - (v15 / ((int16(v16)) | 1)))))}
                            if ((v16 >> (((v1 - (v1 ^ v1))) & 63)) > (v16 ^ v16)) {
                                v16 = v16
                                switch "\n\t" {
                                case "":
                                    switch "a" {
                                    case "a":
                                        var v20 uint16 = v16
                                        _ = v20
                                    default:
                                        var v21 uint64 = 493
                                        if false {
                                            v4 = (v4 + v4)
                                            var v22 uint = (v1 / ((v1) | 1))
                                            _ = len(v5)
                                            v1 -= v3
                                            _ = v22
                                        }
                                        _ = v21
                                    }
                                case "hello":
                                    var v23 uint32 = 0
                                    _ = v23
                                case "a":
                                    v1 -= (^(v1))
                                }
                            }
                            _ = v19
                        } else {
                            if v11 {
                                v24 := GenS1{F0: "hello", F1: 42, F2: 42}
                                _ = v24
                            }
                            var v25 int16 = v15
                            _ = v25
                        }
                        _ = fi18
                    }
                    _ = fi17
                }
                _ = v16
            }
            v26 := GenS3{F0: 26, F1: v15}
            _ = v26
        }
        if (true && ((v4 ^ 49) < (v4 & v4))) {
            fmt.Println(v3, v1)
            v5 = append(v5, -9223372036854775808)
        }
        for fi27 := 0; fi27 < 3; fi27++ {
            if len(v5) > 0 {
                v5[0] = fi27
            }
            var v28 int16 = (v15 & v15)
            _ = fi27
            _ = v28
        }
    } else {
        var v29 int16 = 42
        v3 = (v3 << ((v1) & 63))
        _ = v29
    }
    var v30 string = "hello"
    switch 128 {
    case 42:
        switch -899 {
        case 32767:
            v31 := GenS4{F0: (v4 - (v4 << ((v1) & 63))), F1: v15, F2: 32767}
            _ = v31
        case 0:
            v32 := GenS4{F0: (-(v4)), F1: v15, F2: 255}
            _ = v32
        }
    }
    v30 = (v30 + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
