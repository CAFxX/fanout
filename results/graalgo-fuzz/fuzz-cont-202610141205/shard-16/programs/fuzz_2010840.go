package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint }
type GenS4 struct { F0 int32; F1 int16; F2 uint64 }
func genFn0() string {
    if true {
        var v4 int16 = 1000
        var v5 uint16 = 677
        _ = v4
        _ = v5
    }
    var v6 uint = 65535
    _ = v6
    return ((("" + "a") + "\n\t") + "foo bar")
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{102, -2147483648}
    var v3 int64 = -937
    v7 := v1[2147483648]
    _ = len(v2)
    v8 := GenS1{F0: -1}
    if (1 <= v3) {
        v7 = v7
        fmt.Println(v3, v3)
    } else {
        if ((!(!(!v7))) || (1 > 42)) {
            v9 := v1[42]
            _ = v9
        } else {
            var v10 uint64 = 1
            v3 = (^(v3))
            _ = v10
        }
    }
    if (2 > 384) {
        var v11 int16 = 617
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        _ = len(v1)
        _ = fi12
    }
    switch "\n\t" {
    case "hello":
        if v7 {
            v7 = v7
        } else {
            if v7 {
                fmt.Println(v3, v7)
            }
            for fi13 := 0; fi13 < 10; fi13++ {
                if true {
                    switch fi13 {
                    case 0:
                        v14 := GenS2{F0: v3, F1: 2}
                        _ = v14
                    default:
                        fmt.Println(v7, v3)
                    }
                }
                _ = fi13
            }
        }
    }
    fmt.Println(v7, v3)
    v3 += v3
    v15 := GenS3{F0: 349, F1: 100, F2: 255}
    if v7 {
        v7 = v7
    }
    if (2 != 2) {
        var v16 int = 487
        for fi17 := 0; fi17 < 5; fi17++ {
            fmt.Println(v7)
            v3 = v3
            var v18 uint64 = 1
            _ = fi17
            _ = v18
        }
        if false {
            v2 = append(v2, -300)
            v19 := v1[1]
            switch (((("a" + "\n\t") + "\n\t") + ("hello" + ("foo bar" + ""))) + ("a" + ("" + "hello"))) {
            case "a":
                _ = len(v2)
            case "foo bar":
                v20 := GenS4{F0: -2147483648, F1: -32768, F2: 18446744073709551615}
                _ = v20
            case "hello":
                if v7 {
                    v7 = v7
                    for fi21 := 0; fi21 < 20; fi21++ {
                        switch ("" + "") {
                        case "\n\t":
                            fmt.Println(fi21)
                        case "":
                            if (!(!v7)) {
                                fmt.Println(v19, v7)
                            } else {
                                if (1000 == 768) {
                                    var v22 uint = 780
                                    _ = v22
                                }
                            }
                        default:
                            var v23 int32 = 2147483647
                            _ = v23
                        }
                        v19 = (1 == 0)
                        for fi24 := 0; fi24 < 6; fi24++ {
                            var v25 uint = 18446744073709551615
                            v7 = (v16 > (fi24 + 127))
                            switch "a" {
                            case "\n\t":
                                v19 = ((int16(v25) == -1) || ((-1 >= 0) && (v19 || v19)))
                            case "hello":
                                var v26 int16 = -32768
                                _ = v26
                            case "a":
                                v27 := v1[65535]
                                _ = v27
                            default:
                                v19 = (-128 > 1)
                            }
                            _ = fi24
                            _ = v25
                        }
                        _ = fi21
                    }
                    var v28 bool = v19
                    _ = v28
                }
            }
            _ = v19
        }
        _ = v16
    } else {
        v7 = (true || v7)
    }
    var v29 int = 127
    if len(v2) > 0 {
        v2[0] = (v29 >> ((0) & 63))
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        delete(v1, 0)
        for fi31 := 0; fi31 < 20; fi31++ {
            v3 = (v3 & -1)
            _ = fi31
        }
        v3 = v3
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v7
    _ = v8
}
