package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS2 struct { F0 int16; F1 bool; F2 string }
type GenS3 struct { F0 uint; F1 bool; F2 int64 }
type GenS4 struct { F0 int32 }
func genFn0(p0 int32, p1 uint16) uint64 {
    v5 := GenS1{F0: 1, F1: 65535, F2: 42}
    p1 -= (p1 ^ (p1 >> ((653) & 63)))
    switch "foo bar" {
    case "hello":
        p0 = (p0 ^ (p0 * p0))
    case "":
        p1 += p1
    default:
        var v6 int = 471
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return 892
}

func main() {
    var v1 uint64 = 1
    v2 := []string{("hello" + "hello")}
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    for fi7 := 0; fi7 < 15; fi7++ {
        v1 = 453
        _ = fi7
    }
    var v8 uint = 2
    v2 = append(v2, "foo bar")
    fmt.Println(v1, v8)
    if (1 < 504) {
        fmt.Println(v8, v1)
        var v9 int32 = 1
        _ = v9
    } else {
        fmt.Println(v1, v8)
        if (true || true) {
            var v10 uint16 = 0
            v2 = append(v2, "hello")
            _ = v10
        }
        v1 = genFn0(1, 32767)
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint32 = 0
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int16 = -32768
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 2147483648, F1: v8, F2: 2147483648}
    fmt.Println(v8)
    _ = len(v2)
    for fi16 := 0; fi16 < 7; fi16++ {
        if true {
            v8 = 18446744073709551615
            v8 = v8
        } else {
            v17 := GenS1{F0: 133, F1: v8, F2: uint32(v1)}
            if false {
                if ((v8 % ((v8) | 1)) != (v8 * v8)) {
                    if false {
                        var v18 int64 = -1
                        for fi19 := 0; fi19 < 13; fi19++ {
                            var v20 int = 2147483647
                            _ = fi19
                            _ = v20
                        }
                        _ = v18
                    } else {
                        delete(v4, "a")
                        for fi21 := 0; fi21 < 14; fi21++ {
                            v1 -= genFn0(-1, 65535)
                            v22 := GenS1{F0: 2, F1: v8, F2: 42}
                            _ = fi21
                            _ = v22
                        }
                    }
                }
            }
            for fi23 := 0; fi23 < 7; fi23++ {
                switch "a" {
                case "a":
                    fmt.Println(v8)
                case "foo bar":
                    v8 = (v8 & 182)
                }
                v1 += 2
                if len(v2) > 0 {
                    v2[0] = ("a" + "\n\t")
                }
                _ = fi23
            }
            _ = v17
        }
        if (fi16 != (fi16 / (((fi16 >> (((v8 + (v8 / ((v8) | 1)))) & 63))) | 1))) {
            for fi24 := 0; fi24 < 5; fi24++ {
                if false {
                    switch "foo bar" {
                    case "hello":
                        var v25 int32 = 0
                        _ = v25
                    case "\n\t":
                        v2 = append(v2, "foo bar")
                    case "foo bar":
                        var v26 uint32 = 0
                        _ = v26
                    default:
                        v27 := GenS2{F0: -32768, F1: false, F2: "foo bar"}
                        _ = v27
                    }
                    fmt.Println(fi16)
                    v1 -= (v1 / ((genFn0(2147483647, 2)) | 1))
                } else {
                    v28 := GenS3{F0: (v8 * v8), F1: true, F2: 42}
                    switch 32767 {
                    case -342:
                        fmt.Println(fi24, fi24)
                    }
                    _ = v28
                }
                if true {
                    var v29 uint64 = (v1 + v1)
                    var v30 int8 = 100
                    switch (fi24 * fi24) {
                    case 9223372036854775807:
                        var v31 uint32 = 2147483647
                        _ = v31
                    case -9223372036854775808:
                        v32 := GenS2{F0: -83, F1: ((v30 / (((v30 % ((127) | 1))) | 1)) == v30), F2: ("foo bar" + "")}
                        _ = v32
                    case 42:
                        var v33 bool = false
                        _ = v33
                    }
                    _ = v29
                    _ = v30
                }
                v34 := GenS1{F0: 2147483648, F1: v8, F2: 2147483647}
                _ = fi24
                _ = v34
            }
        }
        var v35 float64 = 0.1
        _ = fi16
        _ = v35
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        v2 = append(v2, "")
        _ = fi36
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        v38 := GenS3{F0: v8, F1: false, F2: 0}
        var v39 uint32 = 42
        _ = fi37
        _ = v38
        _ = v39
    }
    v8 = v8
    switch 32767 {
    case 1:
        v2 = append(v2, ("a" + "\n\t"))
    case 139:
        for fi40 := 0; fi40 < 7; fi40++ {
            v41 := GenS4{F0: 42}
            _ = fi40
            _ = v41
        }
    case -251:
        v8 = v8
    default:
        v42 := GenS2{F0: 42, F1: true, F2: ("" + "")}
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
