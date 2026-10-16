package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 string }
type GenS2 struct { F0 uint; F1 int64; F2 int16 }
type GenS3 struct { F0 int; F1 int32; F2 uint8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int64, p1 string, p2 int16) int8 {
    p1 = ("a" + p1)
    for fi7 := 0; fi7 < 3; fi7++ {
        if false {
            var v8 uint = 636
            _ = v8
        } else {
            var v9 int32 = 2147483647
            var v10 uint = 4294967295
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: true, F1: 2, F2: "foo bar"}
        fmt.Println(p2, p0, p2)
        _ = fi7
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := []string{("\n\t" + "foo bar"), "foo bar", ("foo bar" + ("hello" + ("a" + ("\n\t" + (("hello" + "foo bar") + "hello"))))), "foo bar"}
    var v2 uint16 = 2
    var v3 uint16 = v2
    var v4 int8 = 3
    v5 := []int{-1, -9223372036854775808}
    v6 := make(map[int]uint32)
    var v12 uint32 = 2
    fmt.Println(v12, v3)
    delete(v6, 519)
    switch ("foo bar" + "\n\t") {
    case "a":
        switch ("hello" + "") {
        case "":
            v5 = append(v5, 9223372036854775807)
        case "\n\t":
            v12 = v12
        case "a":
            for fi13 := 0; fi13 < 19; fi13++ {
                v12 -= 2
                v14 := GenS1{F0: (65535 <= 2), F1: 686, F2: "foo bar"}
                var v15 int32 = -319
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            fmt.Println(v12, v2)
        }
    }
    var v16 uint64 = 1
    v17 := GenS2{F0: 0, F1: 9223372036854775807, F2: 312}
    for fi18 := 0; fi18 < 3; fi18++ {
        if false {
            var v19 float64 = (float64((1e10 / 1.0)) / -2.25)
            var v20 string = "a"
            if true {
                var v21 float32 = 0.1
                v22 := GenS3{F0: fi18, F1: -381, F2: 2}
                fmt.Println(v19)
                _ = v21
                _ = v22
            }
            _ = v19
            _ = v20
        }
        v23 := GenS4{F0: v4}
        _ = fi18
        _ = v23
    }
    switch 142 {
    case 485:
        v3 = (v2 % (((v3 << ((109) & 63))) | 1))
    }
    switch "" {
    case "hello":
        v12 += (v12 >> ((604) & 63))
    }
    v4 = v4
    _ = len(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
