package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int; F1 int8; F2 uint16 }
func genFn0(p0 int, p1 int16, p2 bool) int16 {
    p0 -= p0
    v7 := []int{-9223372036854775808, p0, (p0 / ((9223372036854775807) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint32 = 0
    var v3 uint64 = 2
    v4 := []string{"\n\t", ("foo bar" + "hello")}
    var v5 int32 = -2147483648
    var v6 int8 = -31
    v8 := GenS1{F0: "", F1: v6}
    if true {
        var v9 bool = (42 <= v1)
        if ((414 > 0) || false) {
            v3 -= (v3 + 855)
            v10 := GenS1{F0: "hello", F1: 42}
            _ = v10
        } else {
            if (int(v2) == 9223372036854775807) {
                switch 9223372036854775807 {
                case 32767:
                    for fi11 := 0; fi11 < 20; fi11++ {
                        v3 = 4294967295
                        v2 = v2
                        for fi12 := 0; fi12 < 6; fi12++ {
                            var v13 int = fi11
                            v9 = true
                            _ = fi12
                            _ = v13
                        }
                        _ = fi11
                    }
                default:
                    if false {
                        switch "foo bar" {
                        case "foo bar":
                            var v14 bool = true
                            _ = v14
                        case "\n\t":
                            v15 := GenS1{F0: "\n\t", F1: -1}
                            _ = v15
                        case "a":
                            var v16 uint = v1
                            _ = v16
                        }
                        var v17 uint32 = 4294967295
                        var v18 uint64 = (v3 ^ v3)
                        _ = v17
                        _ = v18
                    } else {
                        var v19 int32 = (v5 >> ((v1) & 63))
                        v20 := GenS1{F0: "a", F1: (v6 | v6)}
                        _ = v19
                        _ = v20
                    }
                }
                v21 := GenS1{F0: (((("" + (((("hello" + ((((("hello" + "a") + ((("foo bar" + (("hello" + "a") + "")) + (("a" + "foo bar") + "foo bar")) + "a")) + "foo bar") + ("foo bar" + "")) + ("\n\t" + "foo bar"))) + "a") + (("" + ("" + "\n\t")) + ("foo bar" + "foo bar"))) + "")) + "\n\t") + "foo bar") + "hello"), F1: 32}
                _ = v21
            }
        }
        _ = v9
    } else {
        var v22 string = ("hello" + "foo bar")
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        v6 += v6
        _ = fi23
    }
    fmt.Println(v5)
    fmt.Println(v5, v6, v5)
    fmt.Println(v3, v6)
    switch "foo bar" {
    case "":
        v1 = (^(v1))
    case "foo bar":
        v24 := GenS2{F0: (^(v1)), F1: 65535}
        _ = v24
    default:
        v5 = (v5 - v5)
    }
    v4 = append(v4, "a")
    switch 4294967295 {
    case 4294967295:
        v3 += uint64(v2)
    case -880:
        v2 += v2
    }
    var v25 uint = 255
    fmt.Println(v5, v1, v1)
    v26 := GenS3{F0: 32768, F1: v6, F2: 160}
    v27 := GenS2{F0: (v1 & 482), F1: 1}
    var v28 float32 = (1.5 - 0.1)
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 uint8 = 127
        var v31 float64 = -2.25
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
