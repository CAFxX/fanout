package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 string; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 uint32; F1 string; F2 uint }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string; F1 uint8; F2 int }
func genFn0(p0 int32) uint64 {
    v5 := []int{-1}
    var v6 bool = (false && (((1.5 <= 1e10) && true) && false))
    v6 = (v6 || v6)
    v7 := GenS1{F0: 0}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 42
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    if false {
        var v8 uint = v2
        var v9 bool = true
        v10 := v4[-169]
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        switch fi11 {
        case 42:
            v12 := GenS2{F0: 98, F1: "\n\t", F2: -1}
            _ = v12
        case 941:
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 uint64 = 42
                v15 := GenS3{F0: 763}
                if (("\n\t" >= ("" + ("hello" + "hello"))) && (!true)) {
                    v16 := []int{4294967295, (fi13 & fi13), 9223372036854775807, fi13}
                    if (true || false) {
                        v2 = v2
                        fmt.Println(fi11, v2)
                    } else {
                        var v17 float32 = 1e10
                        v2 = 619
                        delete(v4, (fi13 * (fi11 >> ((v2) & 63))))
                        v17 = ((v17 / v17) * (-2.25 - v17))
                        _ = v17
                    }
                    v14 -= 128
                    _ = v16
                }
                _ = fi13
                _ = v14
                _ = v15
            }
        case -2147483648:
            switch fi11 {
            case 9223372036854775807:
                switch (^(fi11)) {
                case 868:
                    v2 = (v2 * (v2 << ((v2) & 63)))
                case 32767:
                    var v18 int8 = 17
                    _ = v18
                case 2147483647:
                    switch ("\n\t" + ("" + (("a" + "foo bar") + (("\n\t" + "foo bar") + ((("\n\t" + "") + ("" + ((((("\n\t" + "foo bar") + "\n\t") + "\n\t") + "") + "hello"))) + ""))))) {
                    case "\n\t":
                        switch "hello" {
                        case "a":
                            var v19 uint = v2
                            _ = v19
                        case "\n\t":
                            v2 = 689
                        case "":
                            for fi20 := 0; fi20 < 15; fi20++ {
                                var v21 int64 = -746
                                v2 = (v2 + 65535)
                                _ = fi20
                                _ = v21
                            }
                        }
                    case "hello":
                        var v22 bool = (false && false)
                        _ = v22
                    }
                default:
                    v23 := GenS4{F0: genFn0(-1), F1: (false && (((true && false) && (!((fi11 % ((fi11) | 1)) > fi11))) && ((0.0 == (1.5 - (float32(3.14159) * 3.14159))) || false)))}
                    _ = v23
                }
            case 641:
                if (true && (true || ((0.1 / 1.0) > 1e10))) {
                    v2 -= v2
                    v24 := GenS4{F0: genFn0(-2147483648), F1: (!true)}
                    _ = v24
                }
            case 20:
                for fi25 := 0; fi25 < 6; fi25++ {
                    v26 := []int{}
                    _ = fi25
                    _ = v26
                }
            }
        default:
            v27 := GenS5{F0: 1, F1: 307}
            _ = v27
        }
        v2 = v2
        fmt.Println(v2, fi11)
        _ = fi11
    }
    fmt.Println(v2)
    var v28 uint = 18446744073709551615
    v2 += v28
    delete(v3, 2147483647)
    v28 += (v28 / (((v2 >> ((v2) & 63))) | 1))
    v29 := GenS6{F0: 2, F1: "foo bar", F2: v2}
    var v30 int32 = 42
    v31 := GenS7{F0: (v28 + 1), F1: 0}
    v32 := GenS8{F0: ((genFn0((v30 << (((^(v28))) & 63))) < 1) && (4294967295 != uint32(v28)))}
    v33 := GenS9{F0: "foo bar", F1: 127, F2: 9223372036854775807}
    fmt.Println(v28, v30, v30)
    v28 = v2
    v34 := []int{32768, 32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
