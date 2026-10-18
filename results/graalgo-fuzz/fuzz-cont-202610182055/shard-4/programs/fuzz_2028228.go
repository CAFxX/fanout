package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint64 }
func genFn0() uint64 {
    v7 := []int{32767, 1}
    for fi8 := 0; fi8 < 5; fi8++ {
        switch (fi8 | (fi8 + (fi8 >> ((1) & 63)))) {
        case -1:
            if true {
                fmt.Println(fi8)
            } else {
                var v9 uint16 = 255
                var v10 bool = (true || false)
                _ = v10
                _ = v9
            }
        default:
            for fi11 := 0; fi11 < 17; fi11++ {
                switch (-(fi11)) {
                case 55:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                case 2147483647:
                    v12 := GenS1{F0: (-(fi11))}
                    _ = v12
                case -1:
                    v13 := GenS2{F0: 2, F1: 819, F2: 2}
                    fmt.Println(fi8, fi11)
                    var v14 bool = (382 > 255)
                    _ = v13
                    _ = v14
                }
                _ = fi11
            }
        }
        var v15 int8 = -47
        _ = fi8
        _ = v15
    }
    var v16 float32 = 1e10
    if true {
        v17 := GenS3{F0: 255, F1: 183, F2: 18446744073709551615}
        v16 = v16
        _ = v17
    }
    _ = v16
    _ = v7
    return 18446744073709551615
}
func genFn1() int8 {
    var v18 int64 = 42
    var v19 uint8 = 0
    v18 -= 0
    var v20 int16 = 881
    _ = v18
    _ = v19
    _ = v20
    return 1
}

func main() {
    var v1 string = (("foo bar" + "\n\t") + "a")
    v2 := []uint32{42}
    v3 := make(map[string]uint32)
    var v4 string = "a"
    var v5 uint32 = 65535
    var v6 int8 = 42
    var v21 uint64 = 18446744073709551615
    for fi22 := 0; fi22 < 17; fi22++ {
        var v23 bool = false
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        v25 := GenS2{F0: v5, F1: 28, F2: (v21 - (^(v21)))}
        v6 = v6
        _ = fi24
        _ = v25
    }
    var v26 int = 859
    v4 = "a"
    var v27 int = (^(v26))
    fmt.Println(v27)
    for fi28 := 0; fi28 < 1; fi28++ {
        if true {
            v29 := v3[v1]
            _ = v29
        }
        _ = fi28
    }
    v21 += (v21 | genFn0())
    v30 := GenS1{F0: (v27 << ((65535) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
