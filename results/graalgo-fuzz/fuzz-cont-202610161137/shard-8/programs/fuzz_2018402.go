package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS4 struct { F0 bool; F1 int16 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    var v4 uint = v1
    var v5 uint32 = 445
    var v6 string = "hello"
    if false {
        var v7 int8 = 127
        var v8 int = 32768
        _ = v7
        _ = v8
    }
    if true {
        v4 += (v1 + v1)
        for fi9 := 0; fi9 < 15; fi9++ {
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := []int{}
                v12 := GenS1{F0: 2, F1: 42}
                v11 = append(v11, fi9)
                _ = fi10
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 17; fi13++ {
                if true {
                    v6 = "a"
                    v14 := GenS2{F0: ((v5 / (((v5 - (v5 | (v5 & v5)))) | 1)) < v5), F1: 100}
                    _ = v14
                } else {
                    var v15 bool = true
                    if v15 {
                        v5 += (v5 - v5)
                    }
                    fmt.Println(fi13, fi13, v15)
                    delete(v3, (^(v5)))
                    if v15 {
                        v15 = v15
                        v16 := GenS1{F0: 1, F1: 950}
                        _ = v16
                    }
                    v15 = ((!v15) && false)
                    _ = v15
                }
                fmt.Println(fi13, fi9, fi13)
                switch v6 {
                case "":
                    v17 := GenS3{F0: (v1 % ((18446744073709551615) | 1)), F1: 68, F2: 32767}
                    _ = v17
                default:
                    v18 := GenS4{F0: false, F1: -415}
                    _ = v18
                }
                var v19 float64 = 3.14159
                v6 = v6
                v20 := []int{9223372036854775807, (fi13 | (fi9 << ((18446744073709551615) & 63)))}
                _ = fi13
                _ = v19
                _ = v20
            }
            v3[(v5 >> (((v4 | 18446744073709551615)) & 63))] = (fi9 + (fi9 ^ (fi9 & (fi9 - fi9))))
            _ = fi9
        }
    } else {
        for fi21 := 0; fi21 < 4; fi21++ {
            for fi22 := 0; fi22 < 3; fi22++ {
                v1 += 586
                v2[v6] = false
                _ = fi22
            }
            fmt.Println(v1, v4)
            _ = fi21
        }
    }
    v6 = (v6 + (((v6 + (("" + v6) + ((("a" + v6) + (v6 + "foo bar")) + v6))) + (("\n\t" + (v6 + ("\n\t" + v6))) + v6)) + "hello"))
    var v23 uint16 = 255
    v24 := GenS3{F0: 255, F1: 255, F2: (^(v23))}
    var v25 bool = false
    v25 = v25
    var v26 uint32 = v5
    v23 = (v23 << (((v1 + v4)) & 63))
    fmt.Println(v25)
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 uint = (v1 >> (((v4 / ((v1) | 1))) & 63))
        for fi29 := 0; fi29 < 7; fi29++ {
            v28 -= (v4 << ((uint(v26)) & 63))
            _ = fi29
        }
        _ = fi27
        _ = v28
    }
    if (0 == 42) {
        var v30 bool = v25
        var v31 uint = (v4 % ((v1) | 1))
        _ = v30
        _ = v31
    } else {
        var v32 int16 = 42
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
