package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint; F2 int32 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8; F1 uint32; F2 int64 }
func genFn0(p0 uint16, p1 int16) string {
    p0 = 2
    p1 -= p1
    _ = p0
    _ = p1
    return "a"
}
func genFn1(p0 int64, p1 string) bool {
    var v4 uint64 = 1
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: p1, F1: 42, F2: 42}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := []int{4294967295, (fi7 | fi7), 128}
        p0 = 845
        fmt.Println(p0, fi7, p1)
        _ = fi7
        _ = v8
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = v4
    return (2 > 2147483648)
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int = 9223372036854775807
    var v3 float64 = float64(1e10)
    switch "" {
    case "\n\t":
        switch "\n\t" {
        case "foo bar":
            var v9 int64 = 0
            _ = v9
        case "hello":
            if (0 == 42) {
                for fi10 := 0; fi10 < 16; fi10++ {
                    fmt.Println(fi10, v2)
                    _ = fi10
                }
            } else {
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 uint8 = 155
                    v13 := []int{(v2 >> ((2) & 63)), v2}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
        case "a":
            var v14 bool = (((1 == int32(v2)) || false) || true)
            _ = v14
        }
    default:
        v15 := GenS1{F0: "\n\t", F1: v1, F2: 0}
        _ = v15
    }
    if true {
        fmt.Println(v1, v3)
        v16 := GenS2{F0: 127}
        v17 := GenS1{F0: ("" + ("hello" + ("foo bar" + (("foo bar" + "hello") + "a")))), F1: (v1 << ((255) & 63)), F2: 18446744073709551615}
        _ = v16
        _ = v17
    }
    if ("" < ("hello" + "\n\t")) {
        var v18 uint32 = 0
        _ = v18
    } else {
        v1 = v1
    }
    var v19 int8 = 127
    v19 = (v19 >> ((65535) & 63))
    v20 := GenS3{F0: 42}
    v21 := GenS4{F0: 4294967295, F1: 1, F2: -578}
    v22 := GenS5{F0: 530, F1: 255}
    switch ("foo bar" + "hello") {
    case "foo bar":
        var v23 int8 = v19
        _ = v23
    }
    v3 = v3
    var v24 uint8 = 42
    v24 += v24
    v1 += (v1 * v1)
    v25 := []int{(v2 | -9223372036854775808), v2, 42, -1}
    v26 := GenS6{F0: (v1 % (((v1 << ((1) & 63))) | 1))}
    if (!(!true)) {
        for fi27 := 0; fi27 < 9; fi27++ {
            for fi28 := 0; fi28 < 14; fi28++ {
                if (int16(v24) > 780) {
                    if true {
                        var v29 float32 = (((0.1 * -2.25) - 1e10) / -2.25)
                        fmt.Println(fi28, v2, v1)
                        _ = v29
                    }
                } else {
                    var v30 string = ((("\n\t" + "a") + ("foo bar" + (((((("" + "\n\t") + "foo bar") + "hello") + "\n\t") + (((("hello" + "") + ("hello" + "")) + "hello") + "foo bar")) + "a"))) + ((("foo bar" + "\n\t") + "foo bar") + ("hello" + ("\n\t" + "foo bar"))))
                    var v31 uint = 2
                    if len(v25) > 0 {
                        v25[0] = 128
                    }
                    _ = v30
                    _ = v31
                }
                _ = fi28
            }
            var v32 bool = false
            v2 = (fi27 & fi27)
            _ = fi27
            _ = v32
        }
    }
    v33 := GenS3{F0: -2147483648}
    v34 := GenS7{F0: v24, F1: 42, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v34
}
