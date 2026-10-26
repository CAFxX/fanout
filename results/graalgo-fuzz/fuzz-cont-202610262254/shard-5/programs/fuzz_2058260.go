package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint64; F2 bool }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 string, p1 bool) uint32 {
    v7 := []int{9223372036854775807, 42}
    for fi8 := 0; fi8 < 18; fi8++ {
        p0 = (p0 + p0)
        _ = fi8
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v7
    return 2147483648
}
func genFn1(p0 uint, p1 string, p2 uint8) int64 {
    p1 = ("foo bar" + ((p1 + ((p1 + "hello") + ((p1 + (p1 + (p1 + (p1 + "a")))) + (p1 + ("hello" + "foo bar"))))) + "a"))
    var v9 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 uint8 = 2
    var v2 uint32 = 869
    v3 := make(map[uint32]bool)
    var v4 string = ""
    v5 := []string{"foo bar"}
    var v6 uint32 = v2
    v6 = 879
    if false {
        fmt.Println(v6, v6, v4)
        v3[v6] = true
        if false {
            fmt.Println(v1, v4)
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := GenS1{F0: fi10}
                for fi12 := 0; fi12 < 17; fi12++ {
                    if (true || false) {
                        v13 := GenS1{F0: 701}
                        var v14 int64 = 408
                        _ = v13
                        _ = v14
                    }
                    switch ("foo bar" + (v4 + v4)) {
                    case "":
                        var v15 uint64 = 42
                        _ = v15
                    case "\n\t":
                        v2 -= (v6 % ((610) | 1))
                        v16 := v3[(v6 - v2)]
                        v6 += 970
                        _ = v16
                    case "a":
                        var v17 int64 = 1
                        _ = v17
                    }
                    v18 := v3[(^(v6))]
                    _ = fi12
                    _ = v18
                }
                _ = fi10
                _ = v11
            }
        }
    }
    var v19 uint16 = 42
    fmt.Println(v19, v6, v19)
    var v20 bool = true
    _ = len(v5)
    fmt.Println(v19, v20, v6)
    v19 += v19
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := GenS2{F0: 998, F1: 18446744073709551615, F2: false}
        v4 = ((v4 + "a") + "hello")
        _ = fi21
        _ = v22
    }
    switch 32767 {
    case 32767:
        switch (v4 + (v4 + "foo bar")) {
        case "a":
            var v23 bool = v20
            _ = v23
        case "hello":
            v2 -= v6
        default:
            for fi24 := 0; fi24 < 10; fi24++ {
                v4 = ("hello" + (v4 + (v4 + (v4 + "hello"))))
                _ = fi24
            }
        }
    case -9223372036854775808:
        v25 := GenS3{F0: (4294967295 <= 423), F1: (v2 | (v6 + v2))}
        _ = v25
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        fmt.Println(v6)
        _ = fi26
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        fmt.Println(v2)
        var v28 bool = v20
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
