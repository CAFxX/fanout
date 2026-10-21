package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int16; F1 int8; F2 string }
type GenS3 struct { F0 int32 }
func genFn0(p0 int32, p1 int, p2 int8) int32 {
    switch (p1 ^ p1) {
    case 2147483647:
        v5 := []int{-574, p1, p1, (^(p1))}
        _ = v5
    case 700:
        v6 := []int{(-(p1)), p1}
        _ = v6
    default:
        v7 := GenS1{F0: 39, F1: 4294967295}
        _ = v7
    }
    p0 = -378
    _ = p0
    _ = p1
    _ = p2
    return int32(p2)
}
func genFn1(p0 int16, p1 int8) uint {
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: 128, F1: 1}
        p0 += 389
        _ = fi8
        _ = v9
    }
    switch ("hello" + ("" + "a")) {
    case "foo bar":
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    var v4 uint32 = 2147483648
    fmt.Println(v4)
    fmt.Println(v4)
    v10 := GenS2{F0: 1, F1: -109, F2: "a"}
    fmt.Println(v4)
    if (0 < 1) {
        v11 := []int{42}
        for fi12 := 0; fi12 < 11; fi12++ {
            if (!true) {
                v4 = v4
            }
            _ = fi12
        }
        _ = v11
    }
    v1[4294967295] = v4
    if (true && (127 >= -125)) {
        var v13 uint32 = (v4 & v4)
        v14 := []int{42}
        _ = v13
        _ = v14
    } else {
        fmt.Println(v4)
    }
    switch -2147483648 {
    case 29:
        for fi15 := 0; fi15 < 1; fi15++ {
            switch fi15 {
            case 514:
                var v16 int = (fi15 / ((fi15) | 1))
                _ = v16
            case 698:
                for fi17 := 0; fi17 < 20; fi17++ {
                    var v18 int16 = 1000
                    _ = fi17
                    _ = v18
                }
            case 2147483647:
                var v19 string = "a"
                _ = v19
            }
            v4 += v4
            if ((v4 >> ((18446744073709551615) & 63)) != 8) {
                switch -279 {
                case 2147483647:
                    fmt.Println(v4, v4)
                case 42:
                    v4 = 147
                default:
                    if ((false || ((fi15 - (fi15 & fi15)) < (-(fi15)))) || false) {
                        fmt.Println(v4, v4)
                    }
                }
                for fi20 := 0; fi20 < 2; fi20++ {
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        v21 := []int{fi20}
                        var v22 string = "a"
                        _ = v21
                        _ = v22
                    case "hello":
                        v23 := []int{127, 722}
                        _ = v23
                    }
                    var v24 float64 = 1.5
                    _ = fi20
                    _ = v24
                }
                v4 += 0
            }
            _ = fi15
        }
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        _ = len(v1)
        _ = fi25
    }
    var v26 uint16 = 1
    for fi27 := 0; fi27 < 16; fi27++ {
        if true {
            v28 := GenS3{F0: 0}
            _ = v28
        }
        v26 = (v26 % ((v26) | 1))
        _ = fi27
    }
    var v29 int32 = -765
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
