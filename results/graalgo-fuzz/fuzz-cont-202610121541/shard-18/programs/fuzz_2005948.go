package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 string; F2 int }
type GenS4 struct { F0 int; F1 int16; F2 uint64 }
type GenS5 struct { F0 int16 }
func genFn0(p0 int) int64 {
    p0 = p0
    switch "foo bar" {
    case "hello":
        v7 := GenS1{F0: 0, F1: 1, F2: 995}
        _ = v7
    case "a":
        var v8 int64 = -9223372036854775808
        _ = v8
    case "foo bar":
        v9 := []int{p0, -1, (^(p0))}
        _ = v9
    default:
        v10 := []int{}
        _ = v10
    }
    v11 := GenS2{F0: (true || false)}
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v11
    return 0
}
func genFn1(p0 bool, p1 uint8) bool {
    v12 := GenS2{F0: p0}
    fmt.Println(p0, p0)
    var v13 float32 = 0.0
    var v14 float64 = (0.0 * (-2.25 - ((0.0 / 0.1) + float64(((v13 / v13) / v13)))))
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return (714 > 18446744073709551615)
}

func main() {
    var v1 uint16 = 42
    v2 := []string{"a"}
    v3 := []uint32{639}
    var v4 int8 = 0
    v5 := make(map[uint32]uint32)
    var v6 int8 = v4
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS1{F0: 243, F1: 1, F2: 656}
        for fi17 := 0; fi17 < 7; fi17++ {
            v1 = (^(v1))
            v18 := GenS3{F0: (v1 ^ (v1 ^ (v1 + v1))), F1: "hello", F2: (fi17 / ((423) | 1))}
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        v20 := GenS4{F0: (fi19 | (-(fi19))), F1: -607, F2: 4294967295}
        v3 = append(v3, 2147483647)
        v1 -= (v1 & (v1 & (v1 - (v1 / ((v1) | 1)))))
        _ = fi19
        _ = v20
    }
    v6 -= v4
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = 32767
    v21 := GenS4{F0: 32768, F1: -32768, F2: 1}
    v22 := GenS5{F0: 1}
    v3 = append(v3, 2147483647)
    var v23 uint16 = v1
    switch -2147483648 {
    case 128:
        v24 := v5[375]
        _ = v24
    case 365:
        v25 := GenS1{F0: 65535, F1: 2, F2: -9223372036854775808}
        _ = v25
    default:
        if (1 == 2147483647) {
            if len(v3) > 0 {
                v3[0] = 2
            }
        }
    }
    v1 = 2
    fmt.Println(v6, v23, v1)
    v1 = (^(v23))
    switch "a" {
    case "a":
        var v26 uint8 = 42
        _ = v26
    case "hello":
        v27 := GenS1{F0: 2, F1: 18446744073709551615, F2: 1}
        _ = v27
    case "\n\t":
        var v28 float32 = 0.1
        _ = v28
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        _ = len(v5)
        switch (("foo bar" + "\n\t") + ("\n\t" + (("a" + (("hello" + "foo bar") + "foo bar")) + "hello"))) {
        case "\n\t":
            if (1e10 > float64(1.5)) {
                fmt.Println(v6, v23, v1)
                var v30 float64 = -2.25
                _ = v30
            }
        case "hello":
            v1 = v1
        }
        fmt.Println(v4, v23, v4)
        _ = fi29
    }
    v3 = append(v3, 466)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
