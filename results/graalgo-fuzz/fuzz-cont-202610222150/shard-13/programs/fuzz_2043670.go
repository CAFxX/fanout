package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int16, p1 uint64) uint32 {
    for fi7 := 0; fi7 < 14; fi7++ {
        p1 = p1
        for fi8 := 0; fi8 < 20; fi8++ {
            switch "a" {
            case "":
                v9 := []int{fi7, 42, (fi8 - fi7)}
                _ = v9
            }
            _ = fi8
        }
        var v10 uint16 = 541
        _ = fi7
        _ = v10
    }
    v11 := []int{}
    if false {
        var v12 uint16 = 65535
        _ = v12
    } else {
        if (((("a" + (("foo bar" + "hello") + ("\n\t" + "\n\t"))) <= "foo bar") && true) || true) {
            _ = len(v11)
            v11 = append(v11, 32767)
        } else {
            if len(v11) > 0 {
                v11[0] = 32768
            }
        }
        p1 -= (^(p1))
    }
    var v13 bool = true
    _ = p0
    _ = p1
    _ = v11
    _ = v13
    return 65535
}
func genFn1(p0 int16, p1 uint16, p2 uint) bool {
    p0 -= (p0 << ((p2) & 63))
    v14 := GenS1{F0: -39}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (684 <= 774)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "a"
    var v3 uint64 = 42
    v4 := make(map[int]int)
    v5 := make(map[string]string)
    var v6 int8 = 100
    v3 += v3
    if false {
        v2 = v2
    }
    v15 := GenS2{F0: 128}
    v6 = (v6 & 0)
    var v16 int32 = 1
    switch v2 {
    case "a":
        var v17 bool = (v16 != v16)
        _ = v17
    case "":
        if (-1 != 9223372036854775807) {
            var v18 uint64 = 18446744073709551615
            _ = v18
        }
    default:
        var v19 uint32 = 4294967295
        _ = v19
    }
    v6 += v6
    if true {
        fmt.Println(v3, v16)
        v1[101] = 765
        v1[1] = 42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
