package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint32) uint16 {
    v7 := []int{-152}
    for fi8 := 0; fi8 < 14; fi8++ {
        if (32767 != -32768) {
            if len(v7) > 0 {
                v7[0] = -9223372036854775808
            }
            _ = len(v7)
            var v9 float64 = 3.14159
            _ = v9
        }
        for fi10 := 0; fi10 < 16; fi10++ {
            p0 += (p0 * (p0 + 4294967295))
            _ = fi10
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 32767
}
func genFn1(p0 int64, p1 int, p2 int) uint16 {
    switch p2 {
    case 32768:
        v11 := GenS1{F0: 2}
        _ = v11
    case 128:
        p2 = (^(p2))
    default:
        fmt.Println(p2, p0, p2)
    }
    var v12 uint64 = 665
    var v13 uint = 6
    for fi14 := 0; fi14 < 15; fi14++ {
        switch 2147483647 {
        case 4294967295:
            switch "\n\t" {
            case "a":
                switch 1 {
                case -1:
                    var v15 float32 = 1e10
                    _ = v15
                default:
                    fmt.Println(v12, v13, v13)
                }
            case "\n\t":
                if (!(!false)) {
                    var v16 string = "foo bar"
                    p2 = (p2 * 0)
                    switch (p1 + (^(fi14))) {
                    case -9223372036854775808:
                        if ((p2 << ((v13) & 63)) == (p1 * (-(p2)))) {
                            v17 := []int{(-(fi14)), (p2 - 2147483647), 113}
                            _ = v17
                        }
                    }
                    _ = v16
                }
            }
        }
        fmt.Println(p1, p0, v12)
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 int64 = -1
    v4 := []uint32{1}
    var v5 uint8 = 255
    var v6 int64 = 9223372036854775807
    var v18 uint8 = (^(v5))
    var v19 int32 = 937
    if (!false) {
        v5 = 150
        v6 += v3
    } else {
        switch ("a" + "\n\t") {
        case "a":
            v6 = (v3 << ((1) & 63))
        case "foo bar":
            switch "" {
            case "foo bar":
                delete(v1, 32768)
            }
        case "hello":
            v3 += (v3 - v3)
        default:
            _ = len(v4)
        }
    }
    delete(v1, -9223372036854775808)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v2)
    var v20 uint64 = 933
    v5 -= v5
    v19 = int32(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
