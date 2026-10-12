package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 string }
func genFn0() int {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{fi6, fi6, -231}
        var v8 uint = 255
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 bool = false
    var v10 uint32 = 1
    _ = v10
    _ = v9
    return -9223372036854775808
}
func genFn1(p0 uint64, p1 int64, p2 bool) uint64 {
    p1 = (^(p1))
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := []int{}
        var v13 string = "foo bar"
        v14 := GenS1{F0: (fi11 ^ (fi11 >> ((1) & 63)))}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 bool = (!(1 == 1))
    if v15 {
        p1 += -480
    }
    fmt.Println(p2, v15)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return (p0 % ((p0) | 1))
}

func main() {
    var v1 string = "hello"
    v2 := []uint32{82, 2}
    v3 := make(map[int]int)
    var v4 int = -266
    var v5 uint16 = 51
    if (42 > (v4 / ((v4) | 1))) {
        var v16 int32 = -927
        _ = v16
    }
    var v17 float64 = 3.14159
    v5 = (v5 % (((v5 ^ v5)) | 1))
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 float32 = 1e10
        v5 = v5
        _ = len(v3)
        _ = fi18
        _ = v19
    }
    var v20 uint64 = genFn1(2, -9223372036854775808, false)
    if true {
        v5 = (v5 + v5)
        v5 = (v5 - (v5 >> ((38) & 63)))
    } else {
        switch "hello" {
        case "\n\t":
            switch v4 {
            case 32767:
                var v21 int64 = 42
                _ = v21
            case 2147483647:
                if true {
                    v2 = append(v2, 1)
                    v2 = append(v2, uint32(v5))
                }
            }
        case "hello":
            var v22 int8 = -1
            _ = v22
        case "":
            var v23 uint64 = genFn1(v20, int64(v4), (true || (154 != 4294967295)))
            _ = v23
        }
        var v24 int32 = 1
        v17 += 1e10
        _ = v24
    }
    v25 := v3[(v4 & (v4 + (v4 << ((0) & 63))))]
    if false {
        v17 += v17
        var v26 uint = 42
        _ = v26
    } else {
        switch (v4 - v25) {
        case -497:
            v1 = "hello"
        case 685:
            fmt.Println(v25)
        case -2147483648:
            if false {
                fmt.Println(v17, v4, v4)
                if false {
                    v27 := GenS2{F0: v20, F1: "hello"}
                    _ = v27
                }
                var v28 uint32 = 4294967295
                _ = v28
            }
        default:
            for fi29 := 0; fi29 < 6; fi29++ {
                var v30 string = v1
                _ = fi29
                _ = v30
            }
        }
        if (!false) {
            v4 -= genFn0()
            v31 := GenS3{F0: v1}
            _ = v31
        }
    }
    var v32 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
