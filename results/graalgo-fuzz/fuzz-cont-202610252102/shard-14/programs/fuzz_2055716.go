package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int, p1 uint32, p2 int32) uint64 {
    v5 := []int{p0, 128, 1, p0}
    for fi6 := 0; fi6 < 4; fi6++ {
        switch "a" {
        case "":
            if true {
                p0 = (p0 * -1)
                switch "hello" {
                case "\n\t":
                    v7 := GenS1{F0: 100}
                    _ = v7
                case "":
                    var v8 int64 = -686
                    _ = v8
                case "hello":
                    v9 := GenS1{F0: -1}
                    _ = v9
                default:
                    var v10 uint64 = 18446744073709551615
                    _ = v10
                }
            }
        case "hello":
            switch "a" {
            case "\n\t":
                for fi11 := 0; fi11 < 4; fi11++ {
                    switch fi6 {
                    case -2147483648:
                        var v12 string = ("a" + ("a" + ""))
                        _ = v12
                    }
                    if (-931 >= -1) {
                        p1 += p1
                        p0 = fi6
                        fmt.Println(p0)
                    }
                    switch ("hello" + "") {
                    case "hello":
                        p2 = (^(p2))
                    case "foo bar":
                        if false {
                            p1 = 0
                            switch (fi11 | fi6) {
                            case 42:
                                p0 = (p0 / ((p0) | 1))
                            }
                        }
                    default:
                        v5 = append(v5, 42)
                    }
                    _ = fi11
                }
            }
        }
        v13 := GenS2{F0: -1}
        if true {
            var v14 int64 = 554
            _ = v14
        } else {
            _ = len(v5)
            p0 -= 42
            for fi15 := 0; fi15 < 13; fi15++ {
                if false {
                    p2 += 42
                    p1 += 2147483647
                    switch fi6 {
                    case 198:
                        switch "" {
                        case "\n\t":
                            var v16 float32 = (0.1 * 1e10)
                            _ = v16
                        case "hello":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    case 0:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case 9223372036854775807:
                        p1 -= p1
                    default:
                        fmt.Println(p2, p2, p1)
                    }
                }
                _ = fi15
            }
        }
        _ = fi6
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 uint32, p1 uint64, p2 uint8) uint16 {
    p0 = (p0 | uint32(p1))
    var v17 uint64 = (p1 >> ((260) & 63))
    var v18 uint32 = (p0 | p0)
    for fi19 := 0; fi19 < 18; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            v21 := GenS3{F0: 65535}
            var v22 int = fi19
            v22 = (fi20 & fi19)
            _ = fi20
            _ = v21
            _ = v22
        }
        v18 += (p0 | (^(p0)))
        _ = fi19
    }
    v23 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v23
    return 32767
}

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 uint = 42
    v24 := []int{-746, 9223372036854775807, 873}
    delete(v3, 128)
    v25 := GenS3{F0: 2}
    v4 = 1
    var v26 float32 = 1.5
    if (!false) {
        v27 := GenS1{F0: v1}
        _ = v27
    }
    var v28 uint64 = genFn0(128, uint32(v4), 370)
    fmt.Println(v28, v1, v28)
    var v29 float64 = (-2.25 / (3.14159 - (1.5 * 1e10)))
    delete(v2, 717)
    v4 = v4
    _ = len(v24)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
