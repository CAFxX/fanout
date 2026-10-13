package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int8) string {
    v6 := []int{}
    if ((0 >= 66) || true) {
        v7 := GenS1{F0: 65535}
        _ = v7
    }
    _ = p0
    _ = v6
    return "hello"
}

func main() {
    var v1 bool = (!(225 == 42))
    var v2 uint8 = 42
    var v3 int8 = -128
    v4 := make(map[string]uint32)
    v5 := make(map[int]int)
    v8 := []int{}
    var v9 int = 127
    v8 = append(v8, 1)
    if (1 == 42) {
        v10 := GenS1{F0: 65535}
        var v11 uint8 = v2
        _ = len(v8)
        _ = v10
        _ = v11
    } else {
        var v12 int8 = (v3 / (((-(v3))) | 1))
        v9 -= 0
        var v13 string = "hello"
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 1}
    var v15 int64 = 1
    switch ("" + ("hello" + "hello")) {
    case "foo bar":
        if true {
            var v16 int = v9
            if (true && v1) {
                v5[v16] = -9223372036854775808
                var v17 bool = false
                _ = v17
            } else {
                if (!(v1 && (v1 && (4294967295 < v9)))) {
                    switch (("hello" + ("a" + "")) + "a") {
                    case "hello":
                        var v18 uint = 65535
                        _ = v18
                    case "foo bar":
                        switch 1 {
                        case -998:
                            v19 := GenS2{F0: 9223372036854775807, F1: -1}
                            _ = v19
                        case 504:
                            for fi20 := 0; fi20 < 2; fi20++ {
                                var v21 uint = 205
                                _ = fi20
                                _ = v21
                            }
                        case -2147483648:
                            for fi22 := 0; fi22 < 16; fi22++ {
                                v16 = fi22
                                var v23 bool = false
                                _ = fi22
                                _ = v23
                            }
                        }
                    case "a":
                        v16 += (-(v16))
                    }
                }
                switch -2147483648 {
                case 4294967295:
                    var v24 float32 = 1e10
                    _ = v24
                case -502:
                    v16 += (-(v16))
                case -2147483648:
                    switch v9 {
                    case 32767:
                        fmt.Println(v1, v2, v15)
                    case -2147483648:
                        fmt.Println(v2, v3)
                    default:
                        for fi25 := 0; fi25 < 8; fi25++ {
                            v26 := GenS3{F0: 18446744073709551615, F1: 42}
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            _ = fi25
                            _ = v26
                        }
                    }
                default:
                    v16 = (v16 & v16)
                }
                v3 -= -106
            }
            _ = v16
        } else {
            var v27 uint16 = 853
            var v28 int8 = (v3 + (v3 * 46))
            _ = v27
            _ = v28
        }
    case "":
        if (!(v1 && v1)) {
            _ = len(v4)
        }
    case "hello":
        var v29 uint = 2
        _ = v29
    }
    var v30 int = (-(v9))
    v31 := GenS4{F0: 770}
    v30 += v30
    fmt.Println(v9, v1, v9)
    v32 := GenS2{F0: -139, F1: 1}
    var v33 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
