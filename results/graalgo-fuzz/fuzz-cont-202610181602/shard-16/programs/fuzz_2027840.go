package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int }
func genFn0(p0 int16, p1 uint16, p2 int8) int {
    var v4 uint16 = (p1 ^ (^(p1)))
    switch "hello" {
    case "\n\t":
        if (0.1 == 1.5) {
            var v5 int8 = (^(p2))
            _ = v5
        }
    }
    var v6 uint8 = 255
    for fi7 := 0; fi7 < 6; fi7++ {
        if false {
            v8 := []int{}
            if (!false) {
                var v9 int = 127
                fmt.Println(p2, v6)
                _ = v9
            }
            _ = v8
        } else {
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 bool = (699 > 875)
                _ = fi10
                _ = v11
            }
            switch (fi7 % ((-546) | 1)) {
            case 2147483647:
                if (0 >= 42) {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        v13 := GenS1{F0: 2}
                        var v14 float64 = 0.1
                        v15 := GenS2{F0: (p2 | 1)}
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    p1 = p1
                    p2 = (p2 / ((0) | 1))
                }
            }
            v4 = p1
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return 2147483647
}
func genFn1() bool {
    if false {
        var v16 uint8 = 2
        var v17 uint32 = 2147483647
        _ = v16
        _ = v17
    }
    if (false && (2 > 158)) {
        var v18 uint64 = 112
        _ = v18
    }
    var v19 int64 = 9223372036854775807
    _ = v19
    return false
}

func main() {
    var v1 uint = 129
    v2 := []int{898}
    v3 := make(map[uint32]uint32)
    var v20 uint16 = 1
    var v21 string = "hello"
    v2 = append(v2, 42)
    var v22 uint32 = 2
    for fi23 := 0; fi23 < 18; fi23++ {
        for fi24 := 0; fi24 < 7; fi24++ {
            v25 := GenS1{F0: v22}
            var v26 uint = (v1 >> ((v1) & 63))
            v26 -= 1
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi23
    }
    var v27 int32 = 1
    var v28 uint32 = (v22 << ((899) & 63))
    if (true && false) {
        v28 = (v22 | (v22 + (v22 + uint32(v1))))
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 bool = false
            v21 = ("foo bar" + ("\n\t" + (v21 + v21)))
            _ = fi29
            _ = v30
        }
    }
    fmt.Println(v28, v28)
    delete(v3, 0)
    v20 = v20
    var v31 int64 = -9223372036854775808
    v32 := GenS3{F0: 46, F1: 2147483647}
    var v33 int = 341
    v3[(v28 >> ((616) & 63))] = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
