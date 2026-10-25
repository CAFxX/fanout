package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int32; F1 bool; F2 int }
func genFn0() int {
    for fi5 := 0; fi5 < 15; fi5++ {
        fmt.Println(fi5)
        if ((("foo bar" + ("" + ("" + ""))) + "hello") < "foo bar") {
            v6 := GenS1{F0: 42, F1: fi5, F2: 65535}
            switch (fi5 + fi5) {
            case -942:
                var v7 float32 = ((1.5 * ((1e10 - 1e10) * 1e10)) + 0.1)
                _ = v7
            case 42:
                v8 := GenS2{F0: -33, F1: 0, F2: 2147483647}
                _ = v8
            case -353:
                switch "foo bar" {
                case "foo bar":
                    var v9 uint8 = 240
                    _ = v9
                case "a":
                    var v10 float64 = 1e10
                    _ = v10
                case "":
                    fmt.Println(fi5)
                }
            }
            _ = v6
        }
        var v11 int16 = 1
        _ = fi5
        _ = v11
    }
    var v12 uint32 = 445
    var v13 uint16 = 42
    var v14 uint8 = 1
    _ = v12
    _ = v13
    _ = v14
    return 653
}
func genFn1(p0 int16, p1 int8) int16 {
    p0 += p0
    v15 := GenS3{F0: p0, F1: (false && (0.1 <= 0.0))}
    switch 252 {
    case 687:
        v16 := []int{}
        _ = v16
    case -752:
        fmt.Println(p1, p0)
    case -9223372036854775808:
        for fi17 := 0; fi17 < 12; fi17++ {
            v18 := GenS3{F0: (p0 & (p0 & (p0 / ((p0) | 1)))), F1: false}
            p1 = (p1 & 3)
            for fi19 := 0; fi19 < 8; fi19++ {
                v20 := []int{fi17, 127, (fi19 >> ((42) & 63)), fi19}
                fmt.Println(p1, fi19)
                var v21 int8 = (p1 * p1)
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi17
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = v15
    return 0
}

func main() {
    v1 := []uint32{601}
    var v2 uint = 2
    var v3 uint8 = 0
    var v4 uint16 = 32767
    for fi22 := 0; fi22 < 6; fi22++ {
        _ = len(v1)
        var v23 int = genFn0()
        var v24 string = ""
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS4{F0: -471, F1: true, F2: 127}
    fmt.Println(v4, v4)
    fmt.Println(v4)
    if true {
        v26 := GenS3{F0: 1, F1: true}
        _ = v26
    } else {
        v1 = append(v1, 851)
    }
    var v27 uint = 232
    fmt.Println(v3, v3)
    v28 := GenS1{F0: v4, F1: 9223372036854775807, F2: 601}
    fmt.Println(v4, v3, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
