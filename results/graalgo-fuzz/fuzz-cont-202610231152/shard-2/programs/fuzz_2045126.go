package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int, p1 int8, p2 bool) uint32 {
    if (p2 && p2) {
        for fi4 := 0; fi4 < 11; fi4++ {
            var v5 int64 = -140
            fmt.Println(p1)
            switch fi4 {
            case 1:
                fmt.Println(p0)
            case 4294967295:
                var v6 float32 = 3.14159
                _ = v6
            case 2147483647:
                fmt.Println(p0)
            default:
                v7 := GenS1{F0: 1000, F1: 42}
                _ = v7
            }
            _ = fi4
            _ = v5
        }
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    fmt.Println(v1)
    var v8 int32 = -1
    v1 = ("hello" + v1)
    if true {
        var v9 int = 32768
        _ = v9
    }
    var v10 bool = true
    v11 := []int{32768}
    v10 = true
    if v10 {
        var v12 uint8 = 127
        if v10 {
            for fi13 := 0; fi13 < 8; fi13++ {
                v1 = v1
                _ = fi13
            }
            v2[v1] = 9223372036854775807
        } else {
            v14 := GenS1{F0: 1, F1: 0}
            if len(v11) > 0 {
                v11[0] = -285
            }
            _ = v14
        }
        _ = v12
    }
    var v15 uint64 = 496
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            switch (fi17 << ((65535) & 63)) {
            case -2147483648:
                switch fi16 {
                case 32768:
                    if len(v11) > 0 {
                        v11[0] = (fi17 % ((fi16) | 1))
                    }
                case 826:
                    _ = len(v11)
                default:
                    v15 += (v15 * 0)
                }
            }
            fmt.Println(v8)
            _ = fi17
        }
        v18 := GenS2{F0: (v1 + v1)}
        v1 = "foo bar"
        _ = fi16
        _ = v18
    }
    if v10 {
        if v10 {
            fmt.Println(v8)
            v15 += (v15 * (v15 ^ (v15 + (^(v15)))))
        }
        v19 := GenS2{F0: "hello"}
        var v20 uint32 = genFn0(int(v8), -1, ((-2.25 / 1.0) <= (1e10 + 1.5)))
        _ = v19
        _ = v20
    } else {
        fmt.Println(v10)
    }
    fmt.Println(v8, v10)
    v21 := GenS3{F0: 42, F1: v15}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v8
}
