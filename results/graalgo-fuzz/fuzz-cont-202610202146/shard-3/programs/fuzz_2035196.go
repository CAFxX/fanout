package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 string }

func main() {
    var v1 string = ""
    var v2 string = v1
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    v4[578] = "foo bar"
    fmt.Println(v1, v2)
    switch -756 {
    case -619:
        v5 := []int{9223372036854775807, 284, 128}
        _ = v5
    }
    var v6 float32 = 1e10
    v1 = v1
    fmt.Println(v6, v6)
    switch v1 {
    case "foo bar":
        switch "" {
        case "foo bar":
            v1 = ("" + ("foo bar" + ""))
        }
    case "":
        v7 := GenS1{F0: 128, F1: 2}
        _ = v7
    case "\n\t":
        var v8 bool = true
        _ = v8
    default:
        var v9 int = -275
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if (!true) {
            v6 += v6
        } else {
            v11 := []int{fi10, (fi10 ^ 1)}
            _ = v11
        }
        _ = fi10
    }
    if ((0 == 128) || true) {
        v2 = (v2 + ("a" + "foo bar"))
    } else {
        fmt.Println(v6)
    }
    var v12 uint64 = 0
    var v13 int8 = 1
    fmt.Println(v2, v2, v2)
    var v14 uint8 = 128
    v6 = -2.25
    v15 := []int{}
    v6 -= 1.5
    v16 := GenS1{F0: 32768, F1: v12}
    if false {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 bool = false
            v19 := GenS1{F0: fi17, F1: (v12 / (((v12 + v12)) | 1))}
            var v20 int8 = (-(v13))
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
    } else {
        if false {
            fmt.Println(v6, v13)
            v12 -= (v12 % ((v12) | 1))
            var v21 uint = 2
            _ = v21
        } else {
            if ((v2 != ((v1 + ("a" + "\n\t")) + v2)) || ((false || true) && true)) {
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 uint = 65535
                    if true {
                        switch fi22 {
                        case 783:
                            v13 = (v13 % ((33) | 1))
                        case 964:
                            if false {
                                switch (fi22 + 4294967295) {
                                case -300:
                                    if false {
                                        var v24 int = fi22
                                        _ = v24
                                    }
                                }
                                delete(v4, 65535)
                            }
                        default:
                            var v25 float32 = float32(((0.1 * (3.14159 - -2.25)) + -2.25))
                            _ = v25
                        }
                        var v26 uint16 = 2
                        _ = v26
                    }
                    _ = fi22
                    _ = v23
                }
                fmt.Println(v6, v13)
            }
            fmt.Println(v2)
            for fi27 := 0; fi27 < 18; fi27++ {
                v28 := GenS1{F0: fi27, F1: v12}
                v29 := GenS2{F0: v1}
                _ = fi27
                _ = v28
                _ = v29
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
