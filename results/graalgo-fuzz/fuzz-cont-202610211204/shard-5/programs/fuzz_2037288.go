package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0() uint8 {
    var v7 uint8 = 128
    var v8 uint8 = v7
    var v9 uint64 = 2
    v7 += v8
    _ = v7
    _ = v8
    _ = v9
    return v8
}
func genFn1() int {
    var v10 int32 = 42
    var v11 bool = (4294967295 != 1)
    v12 := []int{-1}
    fmt.Println(v10, v11)
    _ = v10
    _ = v11
    _ = v12
    return 128
}

func main() {
    var v1 int64 = 1
    v2 := make(map[string]uint32)
    var v3 int16 = 42
    v4 := []string{"hello", "hello", ("foo bar" + "a")}
    var v5 int8 = 100
    v6 := make(map[string]int)
    fmt.Println(v1)
    switch "a" {
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 int64 = v1
        var v15 int16 = (v3 | (v3 - (-(v3))))
        switch genFn1() {
        case 32767:
            var v16 int32 = 1
            _ = v16
        case -894:
            if (v5 > int8(v1)) {
                var v17 uint64 = 18446744073709551615
                _ = v17
            }
        case 127:
            if true {
                v18 := GenS1{F0: 71, F1: fi13, F2: 25}
                _ = v18
            } else {
                for fi19 := 0; fi19 < 8; fi19++ {
                    v3 = v3
                    if (1.5 == -2.25) {
                        fmt.Println(v1)
                    }
                    v3 += (^(v15))
                    v20 := GenS2{F0: 2147483647}
                    v1 -= -973
                    _ = fi19
                    _ = v20
                }
            }
        default:
            var v21 float64 = (3.14159 / 3.14159)
            _ = v21
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v5, v1)
    v22 := GenS2{F0: 32768}
    if false {
        fmt.Println(v5, v3, v5)
        if true {
            if true {
                for fi23 := 0; fi23 < 20; fi23++ {
                    v1 += 1
                    _ = fi23
                }
                if true {
                    var v24 bool = ((0.1 + 1e10) < 0.0)
                    v1 -= v1
                    _ = v24
                }
                v5 = v5
            }
            var v25 uint = 404
            v3 = v3
            _ = v25
        } else {
            v5 += (v5 - 87)
            delete(v6, "foo bar")
        }
    } else {
        fmt.Println(v1)
    }
    if false {
        v26 := v2["a"]
        _ = v26
    } else {
        if (255 >= 42) {
            v27 := v6["hello"]
            fmt.Println(v1)
            for fi28 := 0; fi28 < 3; fi28++ {
                if (!((0 == 2) && (false && false))) {
                    if (!(1 > 152)) {
                        if (false || true) {
                            if len(v4) > 0 {
                                v4[0] = "\n\t"
                            }
                            for fi29 := 0; fi29 < 15; fi29++ {
                                v5 += (v5 & v5)
                                var v30 int64 = (v1 % ((9223372036854775807) | 1))
                                _ = len(v4)
                                _ = fi29
                                _ = v30
                            }
                        }
                    }
                    var v31 uint = 18446744073709551615
                    v6["foo bar"] = (fi28 * genFn1())
                    _ = v31
                }
                _ = fi28
            }
            _ = v27
        }
        fmt.Println(v5, v3)
    }
    var v32 uint = 735
    v33 := GenS2{F0: 228}
    switch -894 {
    case -237:
        v3 = 1
    case -1:
        var v34 int8 = (^(v5))
        _ = v34
    case -14:
        var v35 int32 = 0
        _ = v35
    default:
        for fi36 := 0; fi36 < 16; fi36++ {
            v4 = append(v4, ("foo bar" + ""))
            switch int(v32) {
            case 9223372036854775807:
                var v37 uint8 = 251
                _ = v37
            case 511:
                var v38 uint8 = 208
                _ = v38
            }
            fmt.Println(v1)
            _ = fi36
        }
    }
    if true {
        switch -146 {
        case -74:
            v3 = (v3 ^ v3)
        case 658:
            v1 = -1
        default:
            v39 := GenS3{F0: 18446744073709551615}
            _ = v39
        }
        v3 -= (v3 - v3)
        fmt.Println(v1)
    }
    v3 += v3
    var v40 int = 397
    fmt.Println(v1, v3)
    v6[""] = (v40 - (-(v40)))
    fmt.Println(v5, v32)
    if false {
        v3 += (v3 | v3)
        var v41 uint = (^(v32))
        _ = v41
    } else {
        if (v1 < int64(v3)) {
            for fi42 := 0; fi42 < 13; fi42++ {
                var v43 int32 = -1
                switch 0 {
                case -835:
                    switch ("a" + ("foo bar" + ("a" + "hello"))) {
                    case "":
                        for fi44 := 0; fi44 < 6; fi44++ {
                            if false {
                                v43 = v43
                                v3 = (v3 ^ 0)
                                v40 = 4294967295
                            }
                            v3 += v3
                            var v45 uint32 = 377
                            _ = fi44
                            _ = v45
                        }
                    case "\n\t":
                        v46 := GenS4{F0: 455}
                        _ = v46
                    default:
                        var v47 int32 = -850
                        _ = v47
                    }
                case 1:
                    if (fi42 >= fi42) {
                        v1 -= 378
                    }
                case 769:
                    v40 += 581
                default:
                    v4 = append(v4, ("" + "foo bar"))
                }
                _ = fi42
                _ = v43
            }
        }
    }
    v40 -= v40
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
