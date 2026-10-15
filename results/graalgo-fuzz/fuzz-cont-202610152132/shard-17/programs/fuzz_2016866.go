package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint; F2 uint }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint8; F1 bool }
func genFn0() int {
    v5 := GenS1{F0: "foo bar", F1: false, F2: 83}
    var v6 int = -2147483648
    _ = v5
    _ = v6
    return v6
}

func main() {
    var v1 bool = true
    var v2 int8 = -1
    var v3 float32 = 1.5
    v4 := []int{}
    _ = len(v4)
    var v7 uint16 = 42
    switch ((((("" + "") + "\n\t") + "") + "foo bar") + "") {
    case "a":
        if v1 {
            v4 = append(v4, -29)
        } else {
            _ = len(v4)
            var v8 int8 = v2
            _ = v8
        }
    case "hello":
        var v9 uint64 = 589
        _ = v9
    default:
        if v1 {
            if (110 == 255) {
                v10 := GenS1{F0: ("hello" + "\n\t"), F1: (1 != 704), F2: (v2 >> ((637) & 63))}
                fmt.Println(v2, v3)
                _ = v10
            } else {
                fmt.Println(v1)
            }
            fmt.Println(v2, v7)
            v11 := GenS1{F0: "\n\t", F1: (255 < 42), F2: (v2 + 87)}
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v1 = (-593 >= 128)
        v7 -= (v7 >> ((65535) & 63))
        var v13 int16 = -32768
        _ = fi12
        _ = v13
    }
    var v14 string = ""
    for fi15 := 0; fi15 < 3; fi15++ {
        switch v14 {
        case "a":
            v7 = (v7 | (v7 - (v7 << ((0) & 63))))
        case "foo bar":
            switch (fi15 / (((-(fi15))) | 1)) {
            case 127:
                v7 -= (v7 % (((v7 / ((32767) | 1))) | 1))
            case 32767:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 0:
                for fi16 := 0; fi16 < 14; fi16++ {
                    if (true && (2 <= 219)) {
                        switch (fi16 & fi16) {
                        case 742:
                            if false {
                                v17 := GenS1{F0: v14, F1: v1, F2: (v2 * v2)}
                                var v18 int32 = 1
                                _ = v17
                                _ = v18
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(fi15)
                                case "foo bar":
                                    v19 := GenS2{F0: 42}
                                    _ = v19
                                }
                                if (237 == 1) {
                                    switch "a" {
                                    case "":
                                        v20 := GenS3{F0: 0, F1: 65535, F2: 4294967295}
                                        _ = v20
                                    default:
                                        switch (-(fi15)) {
                                        case -9223372036854775808:
                                            fmt.Println(v7, v14)
                                        case 2147483647:
                                            fmt.Println(v3, v2)
                                        case 0:
                                            if len(v4) > 0 {
                                                v4[0] = fi16
                                            }
                                        default:
                                            var v21 uint16 = v7
                                            _ = v21
                                        }
                                    }
                                } else {
                                    v14 = (v14 + v14)
                                }
                            }
                        case 0:
                            var v22 uint = 4294967295
                            _ = v22
                        default:
                            v14 = v14
                        }
                        var v23 bool = v1
                        _ = v23
                    } else {
                        v14 = ("\n\t" + "")
                        v24 := GenS1{F0: "foo bar", F1: (!false), F2: v2}
                        _ = v24
                    }
                    _ = fi16
                }
            default:
                v25 := GenS4{F0: 2, F1: v1}
                _ = v25
            }
        }
        v26 := GenS2{F0: 1}
        for fi27 := 0; fi27 < 13; fi27++ {
            v28 := GenS4{F0: 0, F1: true}
            var v29 string = (v14 + v14)
            v1 = v1
            v30 := GenS5{F0: 1, F1: false}
            var v31 int = fi27
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
            _ = v31
        }
        _ = fi15
        _ = v26
    }
    var v32 uint8 = 0
    v7 -= (v7 % ((v7) | 1))
    var v33 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v7
}
