package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int8; F2 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int16) uint64 {
    for fi4 := 0; fi4 < 9; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    p0 = 0
    p0 -= 1000
    var v5 int16 = (p0 - p0)
    _ = p0
    _ = v5
    return 4294967295
}
func genFn1(p0 uint, p1 string, p2 uint) uint16 {
    p1 = "hello"
    switch p1 {
    case "a":
        p2 -= (^(p2))
    default:
        p1 = (p1 + "\n\t")
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 float32 = float32(0.1)
    v2 := []int{1, -2147483648, 127}
    v3 := make(map[uint32]bool)
    if true {
        var v6 bool = (-1 > 0)
        switch 9223372036854775807 {
        case 506:
            if v6 {
                var v7 float32 = 1.5
                fmt.Println(v7)
                _ = v7
            }
        }
        _ = v6
    } else {
        var v8 int16 = 638
        if (4294967295 <= 0) {
            v9 := GenS1{F0: 1}
            v8 -= (v8 / ((v8) | 1))
            _ = v9
        } else {
            v8 = v8
            if ((true && (!((v8 & 32767) != -32768))) || ((714 >= 0) || false)) {
                v1 += float32(3.14159)
            }
            v8 = v8
        }
        _ = v8
    }
    switch 2147483647 {
    case 42:
        v10 := GenS2{F0: 42, F1: 127, F2: 42}
        _ = v10
    case -731:
        switch "\n\t" {
        case "hello":
            switch 32768 {
            case -9223372036854775808:
                v1 = -2.25
            default:
                for fi11 := 0; fi11 < 8; fi11++ {
                    if true {
                        for fi12 := 0; fi12 < 13; fi12++ {
                            fmt.Println(v1, v1, v1)
                            var v13 uint64 = genFn0(-551)
                            v1 = v1
                            _ = fi12
                            _ = v13
                        }
                        v14 := GenS1{F0: 1}
                        for fi15 := 0; fi15 < 10; fi15++ {
                            v1 += v1
                            v1 += (0.0 - v1)
                            _ = fi15
                        }
                        _ = v14
                    } else {
                        v1 -= -2.25
                        var v16 float32 = v1
                        if false {
                            v16 += v16
                        } else {
                            if (1.5 <= (v16 + v1)) {
                                v16 = v1
                                switch (fi11 << ((751) & 63)) {
                                case 32767:
                                    switch (fi11 / ((0) | 1)) {
                                    case 128:
                                        var v17 int = fi11
                                        _ = v17
                                    case -2147483648:
                                        var v18 int8 = -98
                                        _ = v18
                                    }
                                case 9223372036854775807:
                                    if (!(true && false)) {
                                        fmt.Println(v16)
                                        fmt.Println(v1, v16, v16)
                                        v19 := GenS1{F0: 42}
                                        _ = v19
                                    } else {
                                        delete(v3, 65535)
                                        fmt.Println(fi11, v1, v16)
                                        for fi20 := 0; fi20 < 19; fi20++ {
                                            if false {
                                                fmt.Println(v1)
                                            }
                                            var v21 uint64 = 18446744073709551615
                                            v16 -= 3.14159
                                            _ = fi20
                                            _ = v21
                                        }
                                    }
                                default:
                                    var v22 string = "a"
                                    _ = v22
                                }
                                var v23 int8 = -8
                                _ = v23
                            }
                        }
                        _ = v16
                    }
                    v24 := v3[42]
                    _ = fi11
                    _ = v24
                }
            }
        case "\n\t":
            for fi25 := 0; fi25 < 10; fi25++ {
                switch "a" {
                case "":
                    fmt.Println(fi25)
                default:
                    for fi26 := 0; fi26 < 6; fi26++ {
                        switch ((("\n\t" + ("hello" + "a")) + ("a" + "foo bar")) + "") {
                        case "foo bar":
                            var v27 bool = (-32768 < 1000)
                            _ = v27
                        case "a":
                            v1 += (float32(3.14159) / (v1 * 1.5))
                        default:
                            var v28 float64 = 0.0
                            _ = v28
                        }
                        delete(v3, 690)
                        _ = fi26
                    }
                }
                _ = fi25
            }
        default:
            fmt.Println(v1)
        }
    case -745:
        var v29 uint64 = 18446744073709551615
        _ = v29
    default:
        var v30 string = ("\n\t" + "a")
        _ = v30
    }
    v1 += (v1 * (v1 / (v1 - v1)))
    v1 = v1
    v31 := GenS3{F0: 0}
    v1 = v1
    var v32 uint = 18446744073709551615
    v1 -= (0.1 - float32(3.14159))
    v32 = v32
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
}
