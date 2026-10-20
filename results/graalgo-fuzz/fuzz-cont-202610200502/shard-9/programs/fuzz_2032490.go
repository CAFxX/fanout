package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 string }
func genFn0(p0 uint64) int {
    fmt.Println(p0)
    if (!(false && true)) {
        for fi5 := 0; fi5 < 13; fi5++ {
            fmt.Println(p0, fi5)
            _ = fi5
        }
        p0 = (p0 / (((p0 >> ((65535) & 63))) | 1))
    }
    p0 = (^(p0))
    _ = p0
    return 486
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    var v4 uint32 = 973
    var v6 int8 = -1
    fmt.Println(v4, v1)
    if ((3.14159 == 0.1) && (42 >= (v6 >> ((286) & 63)))) {
        for fi7 := 0; fi7 < 19; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                switch ("\n\t" + "\n\t") {
                case "hello":
                    v9 := GenS1{F0: ("a" != (("a" + "foo bar") + ("foo bar" + "hello"))), F1: 729, F2: (fi7 | (fi8 % (((fi8 | (fi7 + -410))) | 1)))}
                    _ = v9
                }
                fmt.Println(v6, v6, v6)
                v4 += (v4 & v4)
                _ = fi8
            }
            switch "hello" {
            case "a":
                var v10 uint32 = (v4 - v4)
                _ = v10
            case "foo bar":
                v11 := v3[(v4 * (v4 + v4))]
                _ = v11
            case "\n\t":
                fmt.Println(v4)
            }
            v4 = v4
            _ = fi7
        }
        v12 := v2[(v4 ^ v4)]
        _ = v12
    }
    v1 += (v1 + v1)
    v6 -= v6
    switch ("\n\t" + "foo bar") {
    case "a":
        v13 := []int{32767, genFn0(42), 32768, 127}
        _ = v13
    case "":
        fmt.Println(v6, v6)
    }
    v14 := GenS2{F0: 530, F1: (v4 / ((2147483647) | 1))}
    for fi15 := 0; fi15 < 18; fi15++ {
        fmt.Println(v6, v4, v1)
        fmt.Println(v6)
        _ = fi15
    }
    var v16 float64 = 3.14159
    v17 := GenS3{F0: "hello"}
    v18 := []int{}
    for fi19 := 0; fi19 < 9; fi19++ {
        v4 += 42
        fmt.Println(v6, v6)
        _ = fi19
    }
    var v20 uint16 = 65535
    for fi21 := 0; fi21 < 7; fi21++ {
        v1 = int16(v4)
        fmt.Println(v1, fi21)
        _ = fi21
    }
    v6 += v6
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
}
