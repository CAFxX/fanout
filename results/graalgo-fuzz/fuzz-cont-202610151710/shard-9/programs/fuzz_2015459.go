package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint64; F2 uint64 }
func genFn0(p0 int) int {
    fmt.Println(p0)
    v6 := GenS1{F0: -1}
    if (false && false) {
        for fi7 := 0; fi7 < 18; fi7++ {
            if false {
                fmt.Println(fi7, fi7)
                if false {
                    fmt.Println(fi7, fi7)
                    for fi8 := 0; fi8 < 2; fi8++ {
                        if ("" != ("a" + "")) {
                            var v9 uint16 = 1
                            fmt.Println(p0)
                            _ = v9
                        }
                        _ = fi8
                    }
                }
            }
            v10 := GenS1{F0: p0}
            switch "" {
            case "a":
                var v11 float64 = (1e10 - 3.14159)
                _ = v11
            }
            _ = fi7
            _ = v10
        }
        var v12 bool = (!(568 < 320))
        _ = v12
    }
    var v13 uint32 = 2147483648
    p0 += 42
    _ = p0
    _ = v13
    _ = v6
    return p0
}
func genFn1(p0 int16, p1 bool) int8 {
    var v14 uint64 = 0
    fmt.Println(p1)
    p0 += p0
    v14 += 42
    _ = p0
    _ = p1
    _ = v14
    return 50
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]int)
    var v3 int = 0
    var v4 int32 = -992
    v5 := []string{("foo bar" + ""), ("" + "foo bar")}
    var v15 uint32 = 42
    v15 = (v15 ^ 482)
    v15 -= (v15 ^ 2147483647)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v16 bool = true
    fmt.Println(v4, v1)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, "foo bar")
    if false {
        v17 := GenS2{F0: (2147483647 < (^(v3))), F1: 42, F2: uint64(v1)}
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(v4, v16)
            var v19 uint64 = 42
            _ = fi18
            _ = v19
        }
        v2[-518] = 42
        _ = v17
    }
    var v20 int8 = -31
    v16 = (v16 && (v16 && v16))
    v21 := v2[v3]
    var v22 uint32 = (v15 >> ((65535) & 63))
    v3 = -203
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
