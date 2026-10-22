package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint16, p2 uint32) uint64 {
    v5 := GenS1{F0: 927}
    v6 := []int{1}
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 int8, p1 uint8, p2 uint) int {
    if (p2 > p2) {
        p1 = 42
        v7 := GenS1{F0: 42}
        v8 := []int{-254}
        _ = v7
        _ = v8
    }
    switch "" {
    case "":
        fmt.Println(p1, p0)
    case "foo bar":
        switch "hello" {
        case "a":
            fmt.Println(p1)
        default:
            p0 = p0
        }
    default:
        v9 := GenS1{F0: -1}
        _ = v9
    }
    var v10 int = 2147483647
    fmt.Println(v10, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 128
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{("" + ""), "hello", "", "\n\t"}
    v3 := []uint32{}
    var v4 int8 = -50
    fmt.Println(v4)
    _ = len(v2)
    fmt.Println(v4)
    switch ("\n\t" + "") {
    case "hello":
        var v11 uint8 = 1
        _ = v11
    case "\n\t":
        for fi12 := 0; fi12 < 11; fi12++ {
            if ((65535 > 1) && false) {
                v4 = (v4 >> ((464) & 63))
            }
            _ = fi12
        }
    }
    if false {
        v4 -= v4
        switch 2147483647 {
        case 2147483647:
            var v13 uint32 = 2
            _ = v13
        case -1:
            switch 2147483647 {
            case -2147483648:
                v14 := v1[2147483648]
                _ = v14
            default:
                if (false && false) {
                    v4 = (v4 >> ((232) & 63))
                    var v15 int = 504
                    _ = v15
                } else {
                    _ = len(v3)
                    var v16 int8 = v4
                    var v17 uint64 = genFn0(32767, 2, 65535)
                    _ = v16
                    _ = v17
                }
            }
        case 598:
            for fi18 := 0; fi18 < 10; fi18++ {
                switch (^(fi18)) {
                case 1:
                    v19 := GenS1{F0: 42}
                    _ = v19
                case -1:
                    v1[65535] = 65535
                }
                _ = fi18
            }
        }
        var v20 uint = 439
        _ = v20
    } else {
        v4 = -56
    }
    var v21 int8 = v4
    var v22 uint64 = 296
    fmt.Println(v21)
    var v23 int8 = v4
    var v24 int16 = 317
    v1[2] = 436
    v25 := GenS2{F0: v24}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
