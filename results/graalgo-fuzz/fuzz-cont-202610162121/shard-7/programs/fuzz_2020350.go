package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint8 }
func genFn0(p0 int, p1 bool, p2 int8) string {
    p0 += 1
    var v4 float32 = 0.0
    if p1 {
        v5 := GenS1{F0: (p0 ^ p0)}
        _ = v5
    }
    if true {
        for fi6 := 0; fi6 < 1; fi6++ {
            fmt.Println(fi6, p2, p2)
            v7 := []int{fi6, -1, fi6, -9223372036854775808}
            for fi8 := 0; fi8 < 18; fi8++ {
                switch (^(fi6)) {
                case 32767:
                    p2 += (p2 >> ((940) & 63))
                case 2147483647:
                    v4 = v4
                }
                _ = fi8
            }
            _ = fi6
            _ = v7
        }
        if (917 >= 9223372036854775807) {
            var v9 int8 = (^(p2))
            for fi10 := 0; fi10 < 4; fi10++ {
                fmt.Println(fi10, p0, p2)
                _ = fi10
            }
            fmt.Println(p1, p0, p2)
            _ = v9
        }
        p0 = p0
    } else {
        switch "a" {
        case "a":
            if (42 <= 822) {
                if (("" + ("foo bar" + "foo bar")) < "\n\t") {
                    if true {
                        var v11 bool = true
                        v12 := GenS1{F0: p0}
                        switch (^(p0)) {
                        case -992:
                            var v13 float32 = (((1e10 / float32((1e10 / 1.0))) * v4) + (float32(3.14159) * float32(1.5)))
                            _ = v13
                        case 1:
                            var v14 uint8 = 0
                            _ = v14
                        case 438:
                            v15 := []int{(p0 * (p0 >> ((65535) & 63))), p0, p0, -926}
                            _ = v15
                        default:
                            v16 := []int{p0, -1, p0}
                            _ = v16
                        }
                        _ = v11
                        _ = v12
                    }
                }
            }
        }
        v17 := GenS2{F0: 1, F1: 1}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "hello"
}

func main() {
    var v1 uint32 = 645
    var v2 int32 = -2147483648
    var v3 uint8 = 253
    var v18 uint32 = v1
    v19 := []int{4294967295, 128, -1}
    for fi20 := 0; fi20 < 2; fi20++ {
        switch (fi20 * (fi20 * (fi20 << ((0) & 63)))) {
        case 128:
            var v21 string = ("hello" + "foo bar")
            _ = v21
        case 9223372036854775807:
            v3 -= v3
        }
        _ = fi20
    }
    _ = len(v19)
    v22 := GenS2{F0: (v2 % ((v2) | 1)), F1: 4294967295}
    v2 = (v2 % ((v2) | 1))
    var v23 uint32 = (^(v18))
    switch "a" {
    case "foo bar":
        v24 := GenS3{F0: v3, F1: 100, F2: (v3 << ((255) & 63))}
        _ = v24
    case "":
        v1 -= 2147483647
    default:
        var v25 uint64 = 42
        _ = v25
    }
    _ = len(v19)
    var v26 int64 = -965
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
}
