package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 int64; F2 int }
func genFn0(p0 uint64, p1 bool, p2 int32) int16 {
    v5 := GenS1{F0: (-(p2))}
    switch ("hello" + "a") {
    case "a":
        var v6 string = "a"
        _ = v6
    case "foo bar":
        if (1 > 2) {
            var v7 uint16 = 42
            var v8 string = (("foo bar" + "hello") + "\n\t")
            v9 := []int{-933, 924}
            _ = v7
            _ = v8
            _ = v9
        }
    case "":
        var v10 float64 = (float64(0.1) * (0.0 - (1e10 / 1.5)))
        _ = v10
    }
    switch "" {
    case "a":
        fmt.Println(p1, p2)
    }
    v11 := []int{32768, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v5
    return -754
}
func genFn1(p0 bool, p1 bool) int32 {
    var v12 bool = p0
    p1 = p0
    switch "" {
    case "\n\t":
        p1 = (!(p0 && (true || (897 >= 42))))
    default:
        var v13 int32 = 1
        _ = v13
    }
    var v14 int = 42
    _ = p0
    _ = p1
    _ = v12
    _ = v14
    return 2147483647
}

func main() {
    v1 := []string{("hello" + ("\n\t" + "foo bar")), "hello", (("" + ("\n\t" + ("hello" + "hello"))) + "hello")}
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    var v4 uint32 = 245
    _ = len(v1)
    fmt.Println(v4)
    if (1 > 23) {
        if (32767 >= genFn0(1, false, 42)) {
            if true {
                fmt.Println(v4)
                for fi15 := 0; fi15 < 10; fi15++ {
                    v16 := GenS2{F0: 0}
                    _ = fi15
                    _ = v16
                }
            } else {
                fmt.Println(v4)
            }
        }
    }
    v17 := GenS2{F0: 42}
    v18 := GenS3{F0: (1 >= 42), F1: v4}
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 int32 = genFn1((!((!(!true)) && true)), true)
        _ = fi19
        _ = v20
    }
    v4 = 4294967295
    v4 = v4
    fmt.Println(v4)
    v21 := GenS4{F0: "hello"}
    var v22 string = "a"
    switch (v22 + ("hello" + (((v22 + (v22 + v22)) + "a") + v22))) {
    case "a":
        v23 := GenS5{F0: 1, F1: 0, F2: 2147483647}
        _ = v23
    }
    v24 := GenS5{F0: 2, F1: -835, F2: -495}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}
