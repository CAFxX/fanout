package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int; F2 int64 }
type GenS3 struct { F0 int; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int32 }
func genFn0(p0 bool, p1 int32, p2 int) bool {
    var v7 bool = (p2 > 4294967295)
    for fi8 := 0; fi8 < 9; fi8++ {
        if (p0 || ("foo bar" == ((("\n\t" + ("hello" + "")) + "hello") + "hello"))) {
            p2 = (p2 >> ((1) & 63))
            var v9 uint = 461
            var v10 uint32 = 614
            _ = v10
            _ = v9
        }
        fmt.Println(p1, p2)
        var v11 int = p2
        _ = fi8
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (!((1 > 1) && ("\n\t" != "foo bar")))
}
func genFn1(p0 uint64, p1 uint64) int64 {
    p1 -= 1
    if (!true) {
        var v12 uint16 = 255
        for fi13 := 0; fi13 < 3; fi13++ {
            v12 = (^(v12))
            if true {
                v14 := GenS1{F0: -806}
                _ = v14
            }
            v15 := []int{(fi13 % ((fi13) | 1))}
            _ = fi13
            _ = v15
        }
        _ = v12
    } else {
        v16 := GenS1{F0: 1}
        v17 := []int{32767, 32768, 127}
        _ = v16
        _ = v17
    }
    if true {
        p0 = 0
    } else {
        if true {
            var v18 float32 = (3.14159 / 0.1)
            _ = v18
        }
    }
    v19 := GenS1{F0: -1}
    _ = p0
    _ = p1
    _ = v19
    return -89
}

func main() {
    v1 := []int{4294967295, -1, 56}
    var v2 uint = 0
    var v3 uint16 = 255
    var v4 uint = v2
    var v5 bool = (v3 == (^(v3)))
    v6 := []string{}
    var v20 uint = (v2 >> ((2) & 63))
    v20 -= 678
    switch int(v4) {
    case 9223372036854775807:
        var v21 int32 = -2147483648
        _ = v21
    }
    v22 := GenS2{F0: "hello", F1: -2147483648, F2: 1}
    if ((!false) || ((-32768 < 32767) && (2 < 35))) {
        for fi23 := 0; fi23 < 16; fi23++ {
            var v24 int16 = 562
            fmt.Println(v2)
            _ = fi23
            _ = v24
        }
        var v25 int8 = 127
        _ = v25
    } else {
        for fi26 := 0; fi26 < 4; fi26++ {
            var v27 int16 = 1000
            var v28 uint = (v4 ^ (v4 | 42))
            for fi29 := 0; fi29 < 8; fi29++ {
                v30 := GenS2{F0: (("" + "a") + "foo bar"), F1: -1, F2: 693}
                v20 = v4
                _ = fi29
                _ = v30
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        if (!v5) {
            var v31 float32 = 3.14159
            _ = v31
        }
    }
    switch ("" + ("a" + "a")) {
    case "foo bar":
        var v32 string = "\n\t"
        _ = v32
    case "":
        v33 := GenS3{F0: -143, F1: 42, F2: v3}
        _ = v33
    case "\n\t":
        fmt.Println(v2, v2)
    default:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    fmt.Println(v4, v20, v5)
    fmt.Println(v20, v2, v5)
    fmt.Println(v4)
    fmt.Println(v3, v20)
    var v34 int16 = 1000
    v3 = 370
    var v35 int64 = 42
    v36 := GenS2{F0: "\n\t", F1: 1, F2: (v35 * (^(v35)))}
    for fi37 := 0; fi37 < 7; fi37++ {
        var v38 float64 = (3.14159 + ((0.0 / 1.0) - ((1e10 + 1e10) - 0.0)))
        _ = fi37
        _ = v38
    }
    if len(v6) > 0 {
        v6[0] = ("a" + "")
    }
    if (v3 >= (^(v3))) {
        v39 := GenS4{F0: -1}
        for fi40 := 0; fi40 < 11; fi40++ {
            var v41 uint = v4
            _ = fi40
            _ = v41
        }
        v1 = append(v1, 2147483647)
        _ = v39
    } else {
        var v42 bool = false
        _ = v42
    }
    fmt.Println(v34)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
