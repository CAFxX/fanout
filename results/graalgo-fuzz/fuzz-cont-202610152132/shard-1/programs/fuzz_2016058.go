package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint64; F2 uint16 }
type GenS3 struct { F0 string; F1 int64; F2 int }
func genFn0(p0 int, p1 int32, p2 int32) int {
    v5 := []int{p0, p0, 128}
    switch ("hello" + (("hello" + ("hello" + "hello")) + "\n\t")) {
    case "":
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 float64 = 0.0
            _ = fi6
            _ = v7
        }
    case "\n\t":
        fmt.Println(p2)
    }
    _ = len(v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}
func genFn1(p0 uint32, p1 uint64, p2 uint16) uint64 {
    if (!false) {
        fmt.Println(p0, p0, p1)
    }
    v8 := GenS1{F0: false}
    var v9 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (p1 >> ((18446744073709551615) & 63))
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int8 = 100
    v3 := []uint32{}
    var v4 int = -598
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS2{F0: "a", F1: 0, F2: 814}
        _ = fi10
        _ = v11
    }
    if ((!true) || false) {
        v12 := GenS1{F0: true}
        switch -9223372036854775808 {
        case 32767:
            var v13 float32 = -2.25
            _ = v13
        case 42:
            fmt.Println(v2, v4)
        }
        var v14 uint32 = 4294967295
        _ = v12
        _ = v14
    } else {
        v15 := GenS1{F0: false}
        fmt.Println(v4, v4)
        fmt.Println(v1)
        _ = v15
    }
    var v16 bool = (3.14159 >= float64((0.0 + 1e10)))
    v4 += (v4 | v4)
    var v17 uint8 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (1000 <= 42) {
        var v18 float32 = ((-2.25 / 1.5) * (0.0 + (float32(float64(-2.25)) + float32(0.1))))
        v16 = (42 == 655)
        _ = len(v3)
        _ = v18
    } else {
        v1 += (^(v1))
    }
    var v19 uint64 = 746
    v3 = append(v3, (v1 >> ((2) & 63)))
    switch (("\n\t" + "foo bar") + "a") {
    case "":
        switch "a" {
        case "\n\t":
            fmt.Println(v16, v2, v19)
        }
    case "hello":
        for fi20 := 0; fi20 < 4; fi20++ {
            if ((359 >= 1) && v16) {
                fmt.Println(v1)
            } else {
                v21 := GenS1{F0: (!true)}
                switch "foo bar" {
                case "":
                    var v22 uint16 = 0
                    _ = v22
                case "hello":
                    v1 = v1
                default:
                    switch ("" + "foo bar") {
                    case "a":
                        for fi23 := 0; fi23 < 19; fi23++ {
                            var v24 uint8 = v17
                            if (42 < 430) {
                                v3 = append(v3, v1)
                                v17 -= v17
                            } else {
                                v25 := GenS1{F0: (359 == -2147483648)}
                                v3 = append(v3, 2147483647)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v25
                            }
                            _ = fi23
                            _ = v24
                        }
                    case "\n\t":
                        v2 -= v2
                    case "foo bar":
                        v26 := GenS3{F0: "\n\t", F1: 42, F2: fi20}
                        _ = v26
                    }
                }
                _ = v21
            }
            var v27 bool = v16
            var v28 uint32 = 2
            _ = fi20
            _ = v27
            _ = v28
        }
    case "a":
        v4 = v4
    default:
        v4 -= (v4 >> ((255) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
