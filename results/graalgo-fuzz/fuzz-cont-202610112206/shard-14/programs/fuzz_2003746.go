package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint16; F1 string }
func genFn0(p0 int, p1 uint32, p2 int16) int8 {
    p1 -= (^(p1))
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := []int{fi6, -9223372036854775808, fi6, (p0 / (((fi6 % (((fi6 / ((fi6) | 1))) | 1))) | 1))}
        var v8 int = p0
        fmt.Println(p1, fi6)
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = (255 <= 278)
    v3 := []string{"a", ""}
    var v4 uint32 = 2147483647
    var v5 int16 = 1000
    var v9 int = 1
    switch v9 {
    case 32768:
        v10 := GenS1{F0: 42, F1: 751, F2: 9223372036854775807}
        _ = v10
    case -474:
        v4 = (v4 % (((^(v4))) | 1))
    }
    var v11 uint8 = 220
    v11 -= (v11 + (^(v11)))
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS1{F0: 1, F1: 4294967295, F2: 42}
        _ = fi12
        _ = v13
    }
    var v14 uint = 42
    var v15 uint8 = v11
    if (!((v15 & v15) > (v15 >> ((v14) & 63)))) {
        v16 := GenS2{F0: 1, F1: "hello"}
        v9 = v9
        v9 += 128
        _ = v16
    } else {
        for fi17 := 0; fi17 < 7; fi17++ {
            if true {
                v2 = false
            } else {
                v14 -= 65535
            }
            _ = fi17
        }
        v15 = (^(v11))
    }
    if (!v2) {
        v14 = (v14 & v14)
        if true {
            v3 = append(v3, "a")
            for fi18 := 0; fi18 < 2; fi18++ {
                for fi19 := 0; fi19 < 13; fi19++ {
                    v9 -= fi18
                    var v20 uint32 = (v4 + (v4 | 791))
                    v1[v20] = (!v2)
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
            v5 -= (v5 | (v5 % ((-32768) | 1)))
        }
    }
    var v21 string = "hello"
    var v22 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
