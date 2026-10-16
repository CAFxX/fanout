package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int32 }
func genFn0() uint {
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int64 = 9223372036854775807
        v8 := []int{}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 128, F1: 4294967295, F2: 1}
    v10 := GenS1{F0: 128, F1: 4294967295, F2: 2}
    var v11 uint = 18446744073709551615
    _ = v10
    _ = v11
    _ = v9
    return (^(v11))
}
func genFn1(p0 uint32, p1 uint16, p2 string) uint8 {
    fmt.Println(p2, p1, p0)
    p2 = ""
    if true {
        p2 = p2
        v12 := GenS1{F0: 255, F1: p0, F2: 65535}
        _ = v12
    } else {
        v13 := GenS1{F0: 127, F1: p0, F2: 1}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return uint8(p0)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 18446744073709551615
    var v3 bool = false
    v4 := make(map[int]string)
    var v5 int16 = 32767
    _ = len(v1)
    v14 := []int{}
    switch ("\n\t" + "hello") {
    case "\n\t":
        fmt.Println(v2, v2)
    case "":
        for fi15 := 0; fi15 < 17; fi15++ {
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            v2 = 194
            _ = fi15
        }
    }
    v4[9223372036854775807] = "\n\t"
    var v16 uint16 = 32767
    v17 := GenS2{F0: 664}
    if true {
        for fi18 := 0; fi18 < 3; fi18++ {
            for fi19 := 0; fi19 < 1; fi19++ {
                v3 = (!true)
                var v20 uint64 = 588
                _ = fi19
                _ = v20
            }
            v2 += v2
            v21 := GenS3{F0: 292, F1: 2147483647}
            _ = fi18
            _ = v21
        }
    }
    v16 = (v16 & v16)
    if true {
        var v22 int64 = 316
        _ = v22
    }
    fmt.Println(v16, v5)
    for fi23 := 0; fi23 < 8; fi23++ {
        if (1 <= -570) {
            if len(v14) > 0 {
                v14[0] = -1
            }
            v14 = append(v14, int(v2))
            var v24 float32 = 1.5
            _ = v24
        }
        fmt.Println(v2)
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
