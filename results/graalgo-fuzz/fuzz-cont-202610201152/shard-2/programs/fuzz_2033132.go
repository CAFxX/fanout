package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int; F2 int16 }
type GenS3 struct { F0 uint64; F1 bool; F2 int16 }
func genFn0(p0 uint16) uint8 {
    v4 := GenS1{F0: -9223372036854775808}
    for fi5 := 0; fi5 < 7; fi5++ {
        p0 -= p0
        var v6 int8 = int8(p0)
        var v7 string = ("" + ("foo bar" + ("\n\t" + ("\n\t" + ("foo bar" + "\n\t")))))
        _ = fi5
        _ = v6
        _ = v7
    }
    switch 1 {
    case -1:
        p0 = 42
    }
    _ = p0
    _ = v4
    return 128
}
func genFn1() uint32 {
    var v8 int8 = 1
    switch "a" {
    case "foo bar":
        v9 := []int{127, 32768}
        _ = v9
    case "a":
        switch 951 {
        case 0:
            var v10 uint64 = 18446744073709551615
            _ = v10
        }
    case "":
        v11 := GenS1{F0: 9223372036854775807}
        _ = v11
    default:
        fmt.Println(v8)
    }
    v12 := []int{}
    fmt.Println(v8)
    v13 := GenS1{F0: -822}
    _ = v12
    _ = v13
    _ = v8
    return 2
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    v14 := []int{9223372036854775807}
    switch -503 {
    case 0:
        _ = len(v14)
    case -234:
        if (16 < 65) {
            if true {
                switch 32767 {
                case -9223372036854775808:
                    v15 := GenS2{F0: (("foo bar" + (((("hello" + "a") + ("hello" + ("foo bar" + (((((("\n\t" + ("\n\t" + "hello")) + "") + "") + "hello") + "") + "foo bar")))) + "a") + "foo bar")) + "a"), F1: 0, F2: 171}
                    _ = v15
                case 1:
                    if (2147483647 <= 0) {
                        var v16 uint8 = (v1 + (v1 % (((v1 >> ((42) & 63))) | 1)))
                        _ = v16
                    }
                default:
                    var v17 float32 = -2.25
                    _ = v17
                }
                for fi18 := 0; fi18 < 3; fi18++ {
                    _ = len(v3)
                    _ = fi18
                }
                if (0 == 206) {
                    switch -575 {
                    case 128:
                        v1 -= (v1 + v1)
                    default:
                        switch -1 {
                        case 128:
                            v1 = 1
                        }
                    }
                } else {
                    _ = len(v3)
                    v1 = 11
                }
            } else {
                for fi19 := 0; fi19 < 4; fi19++ {
                    v1 = 93
                    _ = fi19
                }
                switch 32768 {
                case 42:
                    switch "\n\t" {
                    case "hello":
                        v20 := v2[2]
                        _ = v20
                    case "foo bar":
                        var v21 int64 = 1
                        _ = v21
                    default:
                        var v22 uint16 = 0
                        _ = v22
                    }
                case 2147483647:
                    switch ("\n\t" + ("\n\t" + ("foo bar" + (("hello" + ("a" + "")) + "\n\t")))) {
                    case "":
                        for fi23 := 0; fi23 < 8; fi23++ {
                            v14 = append(v14, (fi23 % ((fi23) | 1)))
                            for fi24 := 0; fi24 < 7; fi24++ {
                                if (1 <= 823) {
                                    fmt.Println(v1, fi24)
                                    var v25 uint16 = 615
                                    var v26 int16 = 42
                                    _ = v25
                                    _ = v26
                                }
                                v1 += (v1 * v1)
                                _ = fi24
                            }
                            _ = fi23
                        }
                    }
                default:
                    fmt.Println(v1)
                }
            }
        } else {
            v1 -= v1
        }
    default:
        v1 -= (v1 & (v1 % ((genFn0(42)) | 1)))
    }
    v1 = v1
    fmt.Println(v1)
    for fi27 := 0; fi27 < 13; fi27++ {
        v28 := GenS3{F0: 206, F1: (3.14159 >= 1.5), F2: 42}
        v29 := GenS2{F0: "a", F1: fi27, F2: 1}
        _ = fi27
        _ = v28
        _ = v29
    }
    _ = len(v2)
    var v30 int = -273
    v1 += v1
    v31 := GenS2{F0: "foo bar", F1: v30, F2: 0}
    if false {
        fmt.Println(v30, v30)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
