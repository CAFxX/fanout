package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 int; F2 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint64; F2 int }

func main() {
    var v1 uint32 = 568
    var v2 uint16 = 32767
    var v3 int8 = -128
    v4 := make(map[uint32]string)
    v5 := make(map[string]string)
    v6 := make(map[string]int)
    var v7 int32 = -2147483648
    if (false || false) {
        var v8 bool = false
        _ = v8
    }
    v1 = (^(v1))
    if true {
        v9 := GenS1{F0: v1}
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 uint = 0
            v7 = v7
            var v12 int64 = -917
            _ = fi10
            _ = v11
            _ = v12
        }
        fmt.Println(v2, v3)
        _ = v9
    } else {
        v1 += v1
    }
    v7 -= v7
    for fi13 := 0; fi13 < 11; fi13++ {
        if true {
            if false {
                v14 := GenS2{F0: -448, F1: v1}
                switch (fi13 % (((-(fi13))) | 1)) {
                case 137:
                    switch ("\n\t" + "a") {
                    case "hello":
                        fmt.Println(v3, v1, v3)
                    default:
                        v15 := []int{}
                        _ = v15
                    }
                case 4294967295:
                    v6["hello"] = fi13
                case 527:
                    v16 := GenS1{F0: (v1 * (v1 & v1))}
                    _ = v16
                default:
                    v1 -= v1
                }
                var v17 int = fi13
                _ = v14
                _ = v17
            } else {
                var v18 uint32 = v1
                for fi19 := 0; fi19 < 18; fi19++ {
                    fmt.Println(v2, v7, v2)
                    _ = fi19
                }
                switch "hello" {
                case "\n\t":
                    switch (fi13 * fi13) {
                    case 1:
                        v7 = (v7 / ((-1) | 1))
                    case -296:
                        v2 -= (^(v2))
                    }
                case "":
                    v18 -= (v18 & 2147483648)
                case "a":
                    v20 := []int{-9223372036854775808}
                    _ = v20
                default:
                    v3 = (v3 << ((534) & 63))
                }
                _ = v18
            }
        } else {
            v21 := []int{fi13, (fi13 - 230), 32767, fi13}
            v2 -= 32767
            var v22 string = "foo bar"
            _ = v21
            _ = v22
        }
        v23 := GenS3{F0: false}
        _ = fi13
        _ = v23
    }
    _ = len(v5)
    fmt.Println(v7, v7, v1)
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := GenS4{F0: 0, F1: 0, F2: 1}
        v7 = 42
        _ = fi24
        _ = v25
    }
    fmt.Println(v7, v2)
    v26 := []int{9223372036854775807, -1, -493}
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS5{F0: true}
        switch ("hello" + "") {
        case "":
            var v29 float32 = ((((float32(float64(0.1)) + 0.0) - 1.5) / 1e10) / float32((float64(-2.25) * 1e10)))
            _ = v29
        case "foo bar":
            if (!false) {
                fmt.Println(fi27)
            }
        case "hello":
            v5[""] = ("foo bar" + "a")
        default:
            var v30 uint64 = 42
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    var v31 string = "a"
    for fi32 := 0; fi32 < 14; fi32++ {
        v33 := GenS6{F0: v31, F1: 42, F2: (fi32 % ((fi32) | 1))}
        var v34 string = v31
        _ = fi32
        _ = v33
        _ = v34
    }
    if (2 <= 741) {
        fmt.Println(v31)
    }
    fmt.Println(v31, v31, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
