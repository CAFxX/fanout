package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int; F2 int16 }
type GenS3 struct { F0 uint; F1 string; F2 int }
type GenS4 struct { F0 uint; F1 int8; F2 string }
type GenS5 struct { F0 bool; F1 bool }
func genFn0(p0 uint8, p1 int, p2 uint) string {
    v7 := GenS1{F0: 2147483647}
    v8 := GenS2{F0: (p2 << ((p2) & 63)), F1: (p1 / (((p1 ^ (p1 - -448))) | 1)), F2: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    var v1 uint16 = 671
    v2 := make(map[uint32]string)
    var v3 int = -280
    v4 := make(map[uint32]int)
    var v5 uint32 = 847
    v6 := make(map[string]int)
    switch "hello" {
    case "hello":
        var v9 uint8 = 0
        _ = v9
    default:
        v5 += (v5 % ((2147483647) | 1))
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v1 = (v1 | (v1 / ((v1) | 1)))
        switch "" {
        case "foo bar":
            v11 := GenS2{F0: 18446744073709551615, F1: fi10, F2: 1000}
            _ = v11
        case "a":
            fmt.Println(v1, v3)
        }
        if (255 >= 245) {
            v12 := GenS3{F0: 1, F1: "", F2: (fi10 * 70)}
            for fi13 := 0; fi13 < 11; fi13++ {
                if false {
                    v14 := GenS3{F0: 255, F1: (("\n\t" + "") + ""), F2: (fi10 & (fi13 | fi10))}
                    _ = v14
                } else {
                    switch ("hello" + "\n\t") {
                    case "":
                        var v15 int32 = 380
                        _ = v15
                    case "hello":
                        switch (("a" + ("\n\t" + "foo bar")) + "foo bar") {
                        case "":
                            v1 += (v1 - v1)
                        case "hello":
                            delete(v4, v5)
                        case "foo bar":
                            var v16 uint32 = v5
                            _ = v16
                        }
                    default:
                        fmt.Println(v3, v3, fi13)
                    }
                }
                _ = fi13
            }
            fmt.Println(fi10, v3)
            _ = v12
        }
        _ = fi10
    }
    var v17 bool = true
    v5 = (^(v5))
    v17 = false
    switch 32767 {
    case 127:
        v5 -= 431
    case 0:
        if true {
            fmt.Println(v5)
            v18 := []int{v3}
            _ = v18
        } else {
            v17 = (2147483647 == v3)
            switch (("hello" + ("hello" + "\n\t")) + "a") {
            case "foo bar":
                if (1.5 > (0.1 * 3.14159)) {
                    switch (-(v3)) {
                    case 0:
                        v19 := []int{}
                        _ = v19
                    case -9223372036854775808:
                        var v20 uint = 4294967295
                        _ = v20
                    case 2147483647:
                        v21 := []int{v3, v3, -9223372036854775808, 128}
                        _ = v21
                    default:
                        v22 := GenS3{F0: 255, F1: ((("a" + ("foo bar" + "hello")) + "foo bar") + ("a" + ("a" + ("foo bar" + "hello")))), F2: (v3 >> ((0) & 63))}
                        _ = v22
                    }
                    if (v17 || (!v17)) {
                        v1 = (v1 & v1)
                        v3 -= v3
                    }
                }
            case "hello":
                v5 -= (v5 ^ (v5 | (v5 * 0)))
            case "\n\t":
                v23 := GenS1{F0: 2147483647}
                _ = v23
            }
        }
    case 9223372036854775807:
        v24 := GenS4{F0: 255, F1: -128, F2: ""}
        _ = v24
    }
    var v25 uint16 = 498
    v6[("a" + ((("foo bar" + "") + "a") + ("\n\t" + "foo bar")))] = v3
    var v26 int16 = -700
    if ((!(((-1 <= -2147483648) || (int32(v5) < 1)) && false)) && (-23 < -1)) {
        switch ("a" + ("" + ("foo bar" + "a"))) {
        case "\n\t":
            var v27 uint32 = (v5 >> ((261) & 63))
            _ = v27
        }
    } else {
        v28 := GenS5{F0: v17, F1: false}
        delete(v4, v5)
        _ = v28
    }
    if (-81 == 0) {
        v29 := []int{v3}
        fmt.Println(v3, v26, v3)
        _ = v29
    }
    var v30 uint16 = 65535
    v1 += (v1 | v25)
    var v31 uint = 2
    v32 := []int{(v3 | 1), -182, int(v5)}
    fmt.Println(v3, v31, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
