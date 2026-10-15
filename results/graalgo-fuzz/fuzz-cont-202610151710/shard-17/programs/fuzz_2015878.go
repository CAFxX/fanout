package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int; F1 uint8; F2 int8 }
type GenS5 struct { F0 string; F1 uint32 }

func main() {
    var v1 int32 = -523
    var v2 uint32 = 4294967295
    v3 := make(map[string]int)
    var v4 int8 = -114
    v5 := make(map[string]string)
    v6 := []string{("hello" + "a"), ("" + "a"), (("hello" + ("" + "a")) + ""), "foo bar"}
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int8 = v4
        var v9 bool = true
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := v5["\n\t"]
    var v11 int64 = 1
    v11 += v11
    switch "foo bar" {
    case "":
        if (-2.25 == (1e10 * 0.1)) {
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 float64 = float64((1.5 + (-2.25 / float32(1.5))))
                _ = fi12
                _ = v13
            }
            v14 := GenS1{F0: 42, F1: -1}
            _ = v14
        }
    }
    v15 := GenS1{F0: 9223372036854775807, F1: -32768}
    v16 := GenS1{F0: 1, F1: -32768}
    if false {
        v17 := GenS2{F0: 65535}
        v4 = 100
        _ = v17
    }
    v10 = (("hello" + v10) + "\n\t")
    if (-964 != 598) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v18 string = v10
        v19 := GenS1{F0: 42, F1: 372}
        _ = v18
        _ = v19
    } else {
        var v20 int8 = v4
        v5[v10] = v10
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        if len(v6) > 0 {
            v6[0] = v10
        }
        for fi22 := 0; fi22 < 17; fi22++ {
            if true {
                v23 := GenS2{F0: 326}
                delete(v3, ("foo bar" + v10))
                if (v10 <= (v10 + "")) {
                    v6 = append(v6, v10)
                } else {
                    fmt.Println(v1, fi21, fi21)
                    v24 := GenS3{F0: (-(v11)), F1: (v2 / ((v2) | 1)), F2: 65535}
                    _ = v24
                }
                _ = v23
            } else {
                _ = len(v6)
            }
            _ = fi22
        }
        _ = len(v6)
        _ = fi21
    }
    if (v4 > 113) {
        fmt.Println(v10)
        v6 = append(v6, (v10 + ("a" + v10)))
    }
    fmt.Println(v1, v11, v10)
    if len(v6) > 0 {
        v6[0] = v10
    }
    _ = len(v6)
    v25 := GenS4{F0: 127, F1: 128, F2: v4}
    for fi26 := 0; fi26 < 20; fi26++ {
        switch 128 {
        case -9223372036854775808:
            for fi27 := 0; fi27 < 18; fi27++ {
                v4 = int8(fi26)
                _ = fi27
            }
        default:
            for fi28 := 0; fi28 < 17; fi28++ {
                switch (v10 + "hello") {
                case "a":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    v29 := GenS5{F0: v10, F1: 384}
                    _ = v29
                }
                v1 -= v1
                v3["\n\t"] = (fi26 - fi28)
                fmt.Println(fi28)
                _ = fi28
            }
        }
        _ = fi26
    }
    delete(v3, "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
