package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int32, p1 int8, p2 int32) int16 {
    p1 = 74
    switch (((("a" + "a") + "hello") + ("hello" + "foo bar")) + "foo bar") {
    case "\n\t":
        fmt.Println(p1, p1, p0)
    case "hello":
        for fi5 := 0; fi5 < 5; fi5++ {
            v6 := []int{(fi5 - 0)}
            _ = fi5
            _ = v6
        }
    case "a":
        p1 -= (p1 >> ((18446744073709551615) & 63))
    }
    v7 := []int{9223372036854775807, 128}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}
func genFn1(p0 uint8, p1 int16, p2 string) uint64 {
    p2 = p2
    var v8 bool = (113 != 0)
    var v9 float32 = (float32(float64(0.0)) / -2.25)
    var v10 bool = (!((v9 / (((3.14159 + float32(3.14159)) + v9) - (1.5 * v9))) != v9))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 int32 = -594
    var v4 uint16 = 66
    v11 := GenS1{F0: 2, F1: true, F2: 983}
    fmt.Println(v3)
    v12 := GenS1{F0: 255, F1: (0 != 193), F2: 42}
    switch "foo bar" {
    case "hello":
        _ = len(v1)
    case "foo bar":
        fmt.Println(v3, v3)
    }
    v3 += v3
    v4 += 255
    switch -872 {
    case 619:
        v13 := GenS1{F0: 255, F1: true, F2: 982}
        _ = v13
    case 0:
        v3 += v3
    default:
        v2["hello"] = true
    }
    v4 = v4
    switch "foo bar" {
    case "\n\t":
        v4 = (v4 | (v4 / ((v4) | 1)))
    case "":
        v4 = v4
    default:
        var v14 int8 = 127
        _ = v14
    }
    if (42 <= 255) {
        switch 1 {
        case 339:
            v4 = v4
        default:
            v15 := []int{57, 4294967295, -9223372036854775808}
            _ = v15
        }
    } else {
        fmt.Println(v4, v4)
    }
    v3 = -1
    if (!(-32768 == -1)) {
        if (1 > 0) {
            var v16 int64 = 887
            v16 += v16
            v17 := GenS2{F0: 0}
            _ = v16
            _ = v17
        }
        if true {
            v3 = (v3 << ((598) & 63))
        }
        v3 -= (v3 << ((255) & 63))
    } else {
        v4 += 550
        switch ("foo bar" + "a") {
        case "a":
            switch ("" + "\n\t") {
            case "\n\t":
                for fi18 := 0; fi18 < 4; fi18++ {
                    var v19 uint32 = 4294967295
                    var v20 int64 = -363
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
        case "\n\t":
            var v21 int8 = 127
            _ = v21
        case "foo bar":
            fmt.Println(v4, v3)
        }
        switch 0 {
        case -2147483648:
            fmt.Println(v4, v3)
        case 32767:
            var v22 int8 = 100
            _ = v22
        default:
            v23 := GenS3{F0: genFn1(2, 42, "hello")}
            _ = v23
        }
    }
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
