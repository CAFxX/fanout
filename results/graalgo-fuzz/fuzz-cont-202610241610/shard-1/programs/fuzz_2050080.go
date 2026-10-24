package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int8; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint, p1 uint8, p2 uint16) int {
    var v7 uint8 = p1
    p1 = 0
    var v8 bool = false
    var v9 int16 = -26
    v10 := GenS1{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 32768
}
func genFn1(p0 uint32, p1 int16) int16 {
    v11 := []int{0}
    if (127 != 128) {
        var v12 uint16 = 0
        _ = v12
    } else {
        var v13 bool = (true || ((p0 / ((752) | 1)) <= p0))
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = v13
    }
    v14 := GenS1{F0: 168}
    _ = p0
    _ = p1
    _ = v11
    _ = v14
    return (p1 - -32768)
}

func main() {
    var v1 int16 = 1000
    var v2 bool = (0.1 >= ((0.1 - float32((float64(float32(-2.25)) - (float64(1.5) + -2.25)))) + 0.1))
    v3 := []string{("\n\t" + "foo bar"), ("\n\t" + ""), "a", "a"}
    var v4 int16 = v1
    v5 := make(map[int]bool)
    var v6 int = 32768
    fmt.Println(v2, v1, v4)
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint64 = 2
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: genFn0(2, 42, 1), F1: 1, F2: v4}
    var v18 string = ((("\n\t" + "a") + "a") + ("a" + ""))
    v19 := GenS1{F0: 32767}
    switch genFn0(255, 127, 2) {
    case 2147483647:
        v6 -= -624
    }
    switch ("hello" + v18) {
    case "":
        v20 := GenS3{F0: 983, F1: 42}
        _ = v20
    case "hello":
        v21 := v5[-2147483648]
        _ = v21
    case "a":
        v4 = v1
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            fmt.Println(v2)
            if true {
                v24 := GenS4{F0: 1}
                _ = v24
            }
            delete(v5, (fi22 % (((^(fi23))) | 1)))
            _ = fi23
        }
        delete(v5, (fi22 + 310))
        v25 := GenS2{F0: fi22, F1: 100, F2: v4}
        _ = fi22
        _ = v25
    }
    v18 = "a"
    v3 = append(v3, v18)
    v26 := GenS5{F0: -1}
    fmt.Println(v1, v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v27 := GenS2{F0: (v6 + genFn0(418, 0, 42)), F1: 42, F2: v1}
    switch genFn0(18446744073709551615, 42, 804) {
    case 9223372036854775807:
        v18 = "foo bar"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
