package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int; F1 int8; F2 int32 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint }
type GenS4 struct { F0 int16; F1 int64 }
func genFn0(p0 int32) uint8 {
    if true {
        v4 := []int{-914, -9223372036854775808}
        _ = v4
    }
    v5 := GenS1{F0: "", F1: 128}
    if ((true && (-1 <= 100)) || (0 < 1)) {
        p0 += -538
        fmt.Println(p0)
    }
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 uint32) int16 {
    v6 := GenS1{F0: ("" + ((("foo bar" + "\n\t") + ("" + "a")) + "\n\t")), F1: 32768}
    p0 = p0
    v7 := []int{}
    fmt.Println(p0)
    switch "a" {
    case "hello":
        p0 = (p0 * (p0 | p0))
    default:
        v8 := GenS1{F0: "hello", F1: 0}
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("foo bar" + "foo bar")
    v3 := make(map[int]int)
    if true {
        if true {
            fmt.Println(v2)
            fmt.Println(v2)
        }
        if false {
            var v9 int8 = -1
            _ = v9
        } else {
            var v10 int8 = 1
            _ = v10
        }
    }
    fmt.Println(v2)
    if false {
        var v11 int32 = 58
        var v12 int = 2147483647
        var v13 string = "a"
        _ = v11
        _ = v12
        _ = v13
    } else {
        var v14 int32 = 1
        if (!false) {
            v15 := GenS2{F0: -1, F1: 101, F2: 0}
            _ = v15
        }
        v16 := v1[558]
        _ = v14
        _ = v16
    }
    v2 = "hello"
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS3{F0: 32767, F1: int64(fi17), F2: 1}
        _ = fi17
        _ = v18
    }
    delete(v3, -762)
    v2 = v2
    v19 := GenS4{F0: 0, F1: 1}
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 bool = (2147483647 != 4294967295)
        if v21 {
            v21 = v21
        } else {
            for fi22 := 0; fi22 < 18; fi22++ {
                switch ("\n\t" + "\n\t") {
                case "a":
                    v23 := []int{(fi20 << ((255) & 63)), 32767, -2147483648}
                    _ = v23
                case "\n\t":
                    fmt.Println(v2, fi22, fi22)
                case "hello":
                    fmt.Println(fi22, fi20, v21)
                default:
                    var v24 float32 = (1.5 + (((0.0 - 1e10) + 0.0) * (1e10 * (0.1 - (1e10 - (0.1 - (0.0 - 0.1)))))))
                    _ = v24
                }
                _ = fi22
            }
            fmt.Println(v2)
        }
        v2 = ((v2 + (v2 + v2)) + v2)
        _ = fi20
        _ = v21
    }
    var v25 uint16 = 1
    var v26 uint64 = 4294967295
    for fi27 := 0; fi27 < 6; fi27++ {
        switch -558 {
        case -645:
            switch (v2 + "") {
            case "foo bar":
                for fi28 := 0; fi28 < 15; fi28++ {
                    _ = len(v3)
                    _ = fi28
                }
            }
        }
        _ = fi27
    }
    v3[-9223372036854775808] = 32768
    if true {
        v29 := []int{-2147483648}
        v30 := v3[-508]
        var v31 int8 = 42
        _ = v29
        _ = v30
        _ = v31
    } else {
        v32 := GenS1{F0: "foo bar", F1: 9223372036854775807}
        fmt.Println(v26)
        v33 := []int{32768, 127}
        _ = v32
        _ = v33
    }
    var v34 float64 = (3.14159 + ((float64((((1.5 / 0.1) / 1.0) / float32((1.5 - (3.14159 * 1e10))))) * float64((3.14159 - 1e10))) * 0.1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v34
}
