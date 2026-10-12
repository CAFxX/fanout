package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 string; F1 string }

func main() {
    var v1 int32 = -2147483648
    var v2 float64 = 1e10
    v3 := []uint32{311}
    for fi4 := 0; fi4 < 2; fi4++ {
        if false {
            v5 := GenS1{F0: 905}
            v1 -= 182
            _ = v5
        }
        fmt.Println(fi4)
        _ = fi4
    }
    switch "hello" {
    case "foo bar":
        v2 -= (-2.25 - ((v2 * 1.5) - v2))
    }
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = 569
    }
    v1 += v1
    v2 = (v2 * v2)
    if true {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 string = "\n\t"
            _ = fi6
            _ = v7
        }
        _ = len(v3)
    }
    v8 := GenS1{F0: 42}
    if (!false) {
        if false {
            _ = len(v3)
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := GenS2{F0: 40}
                v11 := GenS3{F0: 65535, F1: false}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = len(v3)
        }
        fmt.Println(v1, v1)
        if (322 != 0) {
            _ = len(v3)
        } else {
            if (!false) {
                if true {
                    v2 = -2.25
                    var v12 string = (("a" + "hello") + "hello")
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        if true {
                            var v14 string = ("" + ("a" + "hello"))
                            if true {
                                if false {
                                    v15 := GenS2{F0: 435}
                                    _ = v15
                                } else {
                                    var v16 uint32 = 2147483648
                                    _ = v16
                                }
                                fmt.Println(v14)
                            }
                            _ = v14
                        }
                        v1 = (v1 ^ v1)
                        switch (-(fi13)) {
                        case 9223372036854775807:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            for fi17 := 0; fi17 < 4; fi17++ {
                                v1 -= v1
                                _ = fi17
                            }
                        }
                        _ = fi13
                    }
                }
            }
            v1 = 1
            v2 += (v2 / ((((v2 + v2) * v2) - v2) * v2))
        }
    } else {
        v18 := GenS1{F0: 0}
        _ = v18
    }
    switch 128 {
    case 42:
        v2 = v2
    case 127:
        var v19 int16 = 1000
        _ = v19
    case 2147483647:
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := GenS4{F0: "foo bar", F1: "a"}
            if false {
                var v22 uint8 = 127
                _ = v22
            } else {
                fmt.Println(v2)
                var v23 int32 = (v1 << ((2) & 63))
                _ = v23
            }
            _ = fi20
            _ = v21
        }
    default:
        var v24 int64 = -211
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
