package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint32 }
func genFn0(p0 int, p1 bool) uint8 {
    if true {
        fmt.Println(p0, p0)
        var v7 float64 = 1.5
        v7 -= 0.1
        _ = v7
    } else {
        var v8 uint = 18446744073709551615
        if p1 {
            v9 := []int{p0}
            _ = len(v9)
            _ = v9
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := []int{396, (-(fi10)), 32767, fi10}
            v12 := GenS1{F0: 0, F1: 18446744073709551615, F2: 334}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v8
    }
    var v13 string = "hello"
    switch (v13 + "") {
    case "foo bar":
        v13 = ("hello" + v13)
    case "hello":
        fmt.Println(p0, v13, p0)
    }
    v13 = ("a" + v13)
    _ = p0
    _ = p1
    _ = v13
    return 42
}
func genFn1(p0 int16, p1 int64, p2 int16) uint {
    fmt.Println(p2, p0)
    p1 -= (p1 / (((p1 | p1)) | 1))
    p0 = p2
    _ = p0
    _ = p1
    _ = p2
    return 4
}

func main() {
    var v1 uint32 = 1
    var v2 int8 = 100
    var v3 uint64 = 1
    var v4 string = "hello"
    v5 := make(map[string]string)
    v6 := make(map[string]int)
    var v14 int = 630
    v15 := GenS2{F0: 748, F1: 1, F2: (^(v1))}
    fmt.Println(v1, v14)
    v4 = ((v4 + ("\n\t" + v4)) + v4)
    v14 += (v14 & v14)
    var v16 int64 = 9223372036854775807
    for fi17 := 0; fi17 < 11; fi17++ {
        switch (^(fi17)) {
        case -1:
            var v18 uint64 = (^(v3))
            _ = v18
        case 32768:
            v4 = ((v4 + "hello") + (v4 + v4))
        case 216:
            v3 = v3
        }
        fmt.Println(v3, v1, v16)
        v14 = (v14 % ((-998) | 1))
        _ = fi17
    }
    v19 := GenS2{F0: 18446744073709551615, F1: uint8(v3), F2: (v1 >> ((60) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
