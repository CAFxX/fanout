package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint32, p1 uint32, p2 int8) bool {
    v5 := GenS1{F0: 42, F1: 42}
    var v6 uint8 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return true
}

func main() {
    var v1 int16 = 42
    var v2 int8 = -20
    v3 := make(map[int]bool)
    v4 := []int{128, -75, -1}
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: 255, F1: (v7 - 9223372036854775807)}
    var v9 int8 = (^(v2))
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint8 = 97
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint8 = 0
        _ = fi12
        _ = v13
    }
    v1 -= (-(v1))
    var v14 string = "\n\t"
    if true {
        _ = len(v3)
        v1 = v1
        for fi15 := 0; fi15 < 17; fi15++ {
            v1 = v1
            if true {
                for fi16 := 0; fi16 < 10; fi16++ {
                    var v17 uint8 = 255
                    _ = fi16
                    _ = v17
                }
                v1 += v1
            } else {
                if (!(true || true)) {
                    var v18 uint64 = 18446744073709551615
                    _ = v18
                } else {
                    var v19 uint16 = 2
                    var v20 uint = 527
                    _ = v19
                    _ = v20
                }
            }
            if (true || ((v7 - v7) == v7)) {
                var v21 int64 = v7
                _ = v21
            } else {
                fmt.Println(v7, v1, v1)
                v7 = int64(v9)
            }
            _ = fi15
        }
    }
    v1 -= 32767
    if ((false || (!true)) && false) {
        v22 := GenS1{F0: 128, F1: v7}
        var v23 int = -927
        if false {
            for fi24 := 0; fi24 < 20; fi24++ {
                v1 = -142
                v14 = (v14 + v14)
                v14 = v14
                _ = fi24
            }
            switch v23 {
            case -1:
                for fi25 := 0; fi25 < 19; fi25++ {
                    fmt.Println(v9, v9)
                    _ = fi25
                }
            case 42:
                var v26 int8 = (v9 % ((v9) | 1))
                _ = v26
            }
        }
        _ = v22
        _ = v23
    } else {
        if false {
            fmt.Println(v14, v14)
            fmt.Println(v7, v2, v14)
        } else {
            var v27 uint64 = 42
            _ = v27
        }
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    v3[1] = true
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
