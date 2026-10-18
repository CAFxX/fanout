package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int8; F2 int16 }
type GenS3 struct { F0 int }
func genFn0(p0 uint8, p1 int) int64 {
    v7 := GenS1{F0: (("hello" + "hello") + "hello")}
    p1 = p1
    if false {
        if (18446744073709551615 <= 4294967295) {
            p0 = p0
        }
        for fi8 := 0; fi8 < 13; fi8++ {
            if false {
                v9 := []int{}
                if (!true) {
                    fmt.Println(p1)
                    for fi10 := 0; fi10 < 9; fi10++ {
                        if (false || ((p0 / (((p0 << ((1) & 63))) | 1)) == 255)) {
                            v11 := GenS1{F0: "a"}
                            _ = v11
                        }
                        _ = fi10
                    }
                    v12 := GenS2{F0: 100, F1: 1, F2: 42}
                    _ = v12
                }
                _ = v9
            } else {
                fmt.Println(p1, p1)
                p0 -= p0
            }
            _ = fi8
        }
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int8 = 100
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v7
    return -311
}
func genFn1(p0 int32, p1 int, p2 string) int {
    fmt.Println(p2)
    switch p2 {
    case "":
        var v15 uint64 = 4294967295
        _ = v15
    case "foo bar":
        p1 -= p1
    default:
        for fi16 := 0; fi16 < 2; fi16++ {
            p1 -= (fi16 << ((4294967295) & 63))
            _ = fi16
        }
    }
    p1 -= p1
    v17 := []int{p1}
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return (-(p1))
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]string)
    var v3 uint8 = 42
    var v4 string = ""
    var v5 int64 = -1
    v6 := []string{}
    v4 = ((v4 + v4) + v4)
    switch (v4 + v4) {
    case "\n\t":
        var v18 int32 = int32(v5)
        _ = v18
    case "foo bar":
        switch 726 {
        case 42:
            var v19 uint64 = 4294967295
            _ = v19
        default:
            var v20 uint = 2
            _ = v20
        }
    default:
        var v21 uint32 = 0
        _ = v21
    }
    v3 = (v3 + (v3 | (v3 | (v3 * (^(v3))))))
    if (2 == 65535) {
        for fi22 := 0; fi22 < 15; fi22++ {
            v23 := GenS3{F0: genFn1(670, fi22, (("hello" + "") + "a"))}
            _ = fi22
            _ = v23
        }
        var v24 int8 = -128
        _ = v24
    }
    switch "foo bar" {
    case "":
        v4 = ""
    case "\n\t":
        fmt.Println(v5, v4, v4)
    case "a":
        if (0.1 > (3.14159 - 1e10)) {
            v2[42] = v4
        }
    default:
        v25 := GenS2{F0: 72, F1: -128, F2: v1}
        _ = v25
    }
    var v26 bool = true
    v2[101] = (v4 + v4)
    v1 -= (v1 * (v1 + v1))
    var v27 int32 = 0
    v28 := GenS3{F0: genFn1((v27 >> ((65535) & 63)), -9223372036854775808, ("\n\t" + (("" + v4) + ("foo bar" + ("\n\t" + v4)))))}
    for fi29 := 0; fi29 < 20; fi29++ {
        if v26 {
            var v30 string = ("a" + ("hello" + ("foo bar" + (v4 + v4))))
            var v31 uint = 0
            _ = v30
            _ = v31
        } else {
            var v32 int32 = -787
            _ = v32
        }
        delete(v2, 4294967295)
        _ = fi29
    }
    v33 := GenS1{F0: "a"}
    v5 += -552
    var v34 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
