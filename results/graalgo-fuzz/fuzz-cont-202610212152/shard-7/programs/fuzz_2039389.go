package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint32, p1 uint16, p2 uint64) uint {
    switch "\n\t" {
    case "hello":
        switch 4294967295 {
        case -968:
            v7 := []int{127, -848, 490}
            _ = v7
        }
    }
    switch 2147483647 {
    case 32768:
        if (1.5 < -2.25) {
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := []int{(fi8 << ((2) & 63))}
                if true {
                    fmt.Println(p1, p0)
                    for fi10 := 0; fi10 < 6; fi10++ {
                        var v11 float64 = 0.1
                        v11 += -2.25
                        _ = fi10
                        _ = v11
                    }
                }
                if true {
                    _ = len(v9)
                    var v12 int = fi8
                    _ = v12
                }
                _ = fi8
                _ = v9
            }
            var v13 int64 = 849
            _ = v13
        }
    case 42:
        if (!true) {
            for fi14 := 0; fi14 < 6; fi14++ {
                v15 := GenS1{F0: p2, F1: 24, F2: 236}
                p2 += 2
                fmt.Println(p2, p2, p0)
                _ = fi14
                _ = v15
            }
            switch "foo bar" {
            case "a":
                p1 = 720
            case "\n\t":
                p0 -= (p0 + 1)
            }
        }
    }
    var v16 int16 = 42
    if false {
        switch -1 {
        case 0:
            var v17 uint = 1
            _ = v17
        case 826:
            var v18 int16 = (v16 >> ((707) & 63))
            _ = v18
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            p2 = p2
            p1 = (p1 + (p1 % ((p1) | 1)))
            _ = fi19
        }
        p2 -= 1
    } else {
        v20 := []int{}
        for fi21 := 0; fi21 < 19; fi21++ {
            for fi22 := 0; fi22 < 6; fi22++ {
                v23 := GenS2{F0: 32768, F1: (("a" + "\n\t") + "a")}
                switch (("foo bar" + "\n\t") + "a") {
                case "a":
                    fmt.Println(p1, p2, p1)
                    p0 = (^(p0))
                    p2 -= (p2 - p2)
                    v16 -= (^(v16))
                case "hello":
                    v24 := GenS3{F0: int64(p1), F1: p1}
                    v25 := GenS1{F0: p2, F1: -1, F2: -1}
                    _ = v24
                    _ = v25
                }
                switch 127 {
                case 0:
                    _ = len(v20)
                }
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 18446744073709551615
}

func main() {
    var v1 uint = 912
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]int)
    var v4 uint16 = 1
    var v5 uint = (v1 >> (((v1 << ((18446744073709551615) & 63))) & 63))
    var v6 int8 = 42
    v5 = 0
    fmt.Println(v2, v1, v2)
    for fi26 := 0; fi26 < 6; fi26++ {
        fmt.Println(v1, fi26, v2)
        var v27 bool = (2 < (v4 << ((v1) & 63)))
        v28 := GenS4{F0: -374, F1: (^(v4)), F2: -32768}
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 string = ("foo bar" + "hello")
    var v30 bool = false
    v31 := GenS5{F0: 1, F1: 65535}
    var v32 int32 = -1
    v4 = (v4 & v4)
    var v33 uint16 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
