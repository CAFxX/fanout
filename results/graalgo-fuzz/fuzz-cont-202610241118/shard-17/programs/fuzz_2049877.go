package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint32; F2 int8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int32, p1 uint32, p2 bool) int64 {
    fmt.Println(p1, p1, p2)
    p1 = p1
    switch 2147483647 {
    case 240:
        switch "hello" {
        case "":
            switch "a" {
            case "\n\t":
                v6 := GenS1{F0: (55 == 42)}
                _ = v6
            case "hello":
                if (!(!(p2 || (p2 || ((307 >= 42) || (2 >= 24)))))) {
                    var v7 bool = (1e10 != 0.1)
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 3; fi8++ {
                        p1 -= (^(p1))
                        _ = fi8
                    }
                }
            }
        }
    case -9223372036854775808:
        v9 := GenS2{F0: 291, F1: p1, F2: -128}
        _ = v9
    case 127:
        var v10 int = 0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := []string{"", ("" + "\n\t")}
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    var v4 float32 = 3.14159
    var v5 int32 = 595
    v4 = v4
    v11 := GenS1{F0: true}
    if true {
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 bool = true
            v13 = false
            for fi14 := 0; fi14 < 16; fi14++ {
                if (v13 && (-106 >= 127)) {
                    var v15 int = 1
                    _ = v15
                }
                if v13 {
                    fmt.Println(v13, v13)
                }
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
    }
    v5 += v5
    v4 = (v4 / v4)
    v16 := GenS3{F0: 100}
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 string = "\n\t"
        _ = fi17
        _ = v18
    }
    if ((!(1 <= 175)) || false) {
        switch ("hello" + ("hello" + "")) {
        case "\n\t":
            var v19 bool = false
            _ = v19
        }
    }
    var v20 int = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
