package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 string; F2 uint }
func genFn0() int16 {
    var v6 uint64 = 0
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 1}
            _ = fi8
            _ = v9
        }
        if (!false) {
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := []int{2147483647, 4294967295, (fi10 & fi10)}
                _ = fi10
                _ = v11
            }
            v12 := []int{int(v6), fi7}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 string = ""
                v6 = 174
                _ = fi13
                _ = v14
            }
            switch "foo bar" {
            case "":
                v15 := []int{fi7, (fi7 & 9223372036854775807), -2147483648}
                _ = v15
            case "foo bar":
                var v16 int8 = -1
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            v6 = 0
            _ = fi17
        }
        _ = fi7
    }
    fmt.Println(v6)
    for fi18 := 0; fi18 < 16; fi18++ {
        if true {
            if true {
                fmt.Println(v6, v6)
                v6 += (v6 << ((706) & 63))
                var v19 uint64 = (^(v6))
                _ = v19
            }
        }
        _ = fi18
    }
    _ = v6
    return 42
}
func genFn1(p0 uint64, p1 int32) uint {
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            v22 := GenS2{F0: 1}
            p0 += p0
            p0 -= p0
            _ = fi21
            _ = v22
        }
        p1 = p1
        _ = fi20
    }
    v23 := []int{4294967295, 128, -96}
    var v24 float32 = float32((0.1 + (0.1 / (1.5 / -2.25))))
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 42
    var v2 int32 = 2147483647
    v3 := make(map[string]string)
    var v4 int16 = 1000
    var v5 int32 = 1
    fmt.Println(v1, v5)
    var v25 uint64 = 1
    var v26 uint = 18446744073709551615
    fmt.Println(v26, v26, v1)
    fmt.Println(v26, v5, v26)
    v27 := GenS2{F0: 2}
    v28 := []int{9223372036854775807, 32767, -558, -1}
    _ = len(v3)
    var v29 int16 = (v4 >> ((v26) & 63))
    var v30 int = 798
    var v31 uint8 = 2
    fmt.Println(v2)
    v32 := GenS3{F0: -1, F1: "hello", F2: (v26 << (((^(v26))) & 63))}
    if len(v28) > 0 {
        v28[0] = v30
    }
    fmt.Println(v26, v26, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
