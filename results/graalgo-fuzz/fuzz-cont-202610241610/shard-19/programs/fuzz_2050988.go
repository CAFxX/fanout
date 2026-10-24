package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0(p0 int8, p1 uint16, p2 int) int16 {
    p0 -= (p0 >> ((18446744073709551615) & 63))
    switch (p2 + p2) {
    case -690:
        for fi5 := 0; fi5 < 5; fi5++ {
            switch 9223372036854775807 {
            case 128:
                if (542 < 2) {
                    v6 := GenS1{F0: ""}
                    p1 = p1
                    var v7 int64 = -381
                    _ = v6
                    _ = v7
                } else {
                    switch "" {
                    case "hello":
                        var v8 uint64 = 386
                        _ = v8
                    case "foo bar":
                        if ((1 < 42) || false) {
                            var v9 string = ""
                            _ = v9
                        }
                    case "a":
                        p1 += 32767
                    default:
                        var v10 string = "\n\t"
                        _ = v10
                    }
                    p0 -= p0
                }
            case 4294967295:
                var v11 uint16 = p1
                _ = v11
            }
            for fi12 := 0; fi12 < 18; fi12++ {
                switch "a" {
                case "\n\t":
                    if (0.0 != 1e10) {
                        v13 := []int{}
                        _ = v13
                    }
                case "hello":
                    p1 -= (p1 - (p1 << ((442) & 63)))
                }
                v14 := []int{p2, (fi5 % (((fi12 / ((p2) | 1))) | 1)), -2147483648, 127}
                _ = fi12
                _ = v14
            }
            _ = fi5
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}
func genFn1(p0 bool, p1 int) int {
    var v15 uint32 = 2147483648
    v16 := []int{p1}
    switch (-(p1)) {
    case 732:
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return p1
}

func main() {
    var v1 int32 = -722
    var v2 uint16 = 2
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    v2 += 4
    v2 = v2
    v17 := []int{32768}
    if (32767 < 42) {
        var v18 int32 = v1
        var v19 int16 = genFn0(42, (v2 / ((42) | 1)), 127)
        _ = v18
        _ = v19
    } else {
        if (true || false) {
            var v20 uint8 = 2
            _ = v20
        } else {
            v1 -= 42
            _ = len(v3)
        }
    }
    var v21 float64 = 1e10
    var v22 int64 = 0
    switch "hello" {
    case "\n\t":
        for fi23 := 0; fi23 < 19; fi23++ {
            if (fi23 != 32768) {
                _ = len(v17)
            }
            v22 -= (v22 + (v22 % ((v22) | 1)))
            v24 := GenS2{F0: (v22 % ((v22) | 1)), F1: v22}
            _ = fi23
            _ = v24
        }
    case "":
        fmt.Println(v22, v21)
    default:
        v21 = v21
    }
    var v25 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
