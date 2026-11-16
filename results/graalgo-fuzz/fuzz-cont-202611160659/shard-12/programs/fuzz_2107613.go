package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint32; F2 int32 }
func genFn0() string {
    var v5 uint32 = 0
    v5 += (v5 & 1)
    _ = v5
    return "hello"
}
func genFn1(p0 string, p1 uint8) string {
    var v6 uint16 = 187
    fmt.Println(p1)
    var v7 bool = (-47 < 127)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (("foo bar" + "") + p0)
}

func main() {
    var v1 bool = false
    v2 := []string{"foo bar", (("\n\t" + "") + (("foo bar" + "hello") + ("" + "a")))}
    v3 := make(map[int]int)
    var v4 uint64 = 0
    v1 = (false || (127 != 128))
    v8 := v3[128]
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v9 := GenS1{F0: 9223372036854775807, F1: (v4 ^ 2)}
    if false {
        switch "\n\t" {
        case "a":
            v2 = append(v2, "a")
        case "":
            v4 = 4294967295
        case "\n\t":
            v1 = v1
        default:
            v10 := GenS1{F0: -1, F1: 1}
            _ = v10
        }
    }
    v3[(^(v8))] = v8
    var v11 uint32 = 42
    switch "foo bar" {
    case "foo bar":
        var v12 uint16 = 1
        _ = v12
    case "\n\t":
        var v13 float64 = (0.1 * 1e10)
        _ = v13
    default:
        v8 += (v8 + (v8 << ((302) & 63)))
    }
    if ((v8 - (-(v8))) <= v8) {
        v14 := GenS1{F0: 42, F1: (v4 ^ (v4 / (((v4 | (v4 / ((4294967295) | 1)))) | 1)))}
        v15 := GenS2{F0: -332, F1: (v11 << ((65535) & 63)), F2: 42}
        v16 := GenS2{F0: 9223372036854775807, F1: v11, F2: 361}
        _ = v14
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 13; fi18++ {
                var v19 uint32 = (v11 + (v11 % ((59) | 1)))
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
        if ((v11 / ((0) | 1)) != v11) {
            v11 -= (v11 & v11)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    var v20 float32 = 0.0
    switch (^(v8)) {
    case 2147483647:
        delete(v3, 2147483647)
    case 128:
        if false {
            v8 -= (v8 & -9223372036854775808)
            v1 = true
            var v21 string = "a"
            _ = v21
        }
    case -9223372036854775808:
        v8 = (^(v8))
    }
    var v22 uint32 = 1
    v4 = 1
    v23 := v3[v8]
    if true {
        if v1 {
            if (float64(0.0) > ((((0.1 * 1.5) - 1e10) + 1.5) * 1.5)) {
                if v1 {
                    var v24 int64 = 1
                    switch 32767 {
                    case 704:
                        for fi25 := 0; fi25 < 6; fi25++ {
                            var v26 int16 = 0
                            _ = fi25
                            _ = v26
                        }
                    case -943:
                        switch "" {
                        case "hello":
                            for fi27 := 0; fi27 < 11; fi27++ {
                                var v28 int32 = 1
                                var v29 int16 = 1
                                var v30 uint64 = 18446744073709551615
                                _ = fi27
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                        case "foo bar":
                            var v31 uint32 = 1
                            _ = v31
                        default:
                            for fi32 := 0; fi32 < 6; fi32++ {
                                v33 := GenS1{F0: v24, F1: (v4 % ((2) | 1))}
                                _ = fi32
                                _ = v33
                            }
                        }
                    case 401:
                        var v34 bool = true
                        _ = v34
                    default:
                        var v35 int16 = 165
                        _ = v35
                    }
                    delete(v3, (^(v8)))
                    _ = v24
                }
                fmt.Println(v22, v11)
                var v36 uint8 = 13
                _ = v36
            }
            fmt.Println(v8, v22, v11)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
