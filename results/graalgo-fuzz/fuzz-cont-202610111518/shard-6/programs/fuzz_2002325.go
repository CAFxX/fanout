package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int32; F2 uint64 }
type GenS3 struct { F0 int; F1 string }
func genFn0(p0 uint32, p1 bool, p2 bool) int16 {
    v7 := GenS1{F0: 42}
    var v8 float32 = (((1.5 * (-2.25 / -2.25)) * -2.25) + float32(0.1))
    var v9 uint32 = 1
    var v10 int64 = 9223372036854775807
    v11 := []int{42, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return -942
}
func genFn1() uint64 {
    var v12 int8 = 1
    v13 := []int{127, 1, 639, -547}
    var v14 int8 = (v12 + v12)
    var v15 int16 = 1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 2
}

func main() {
    v1 := []int{}
    v2 := []uint32{0, 2147483648}
    v3 := make(map[uint32]int)
    v4 := []int{1, 9223372036854775807}
    var v5 uint = 2
    var v6 uint16 = 98
    _ = len(v2)
    if (!true) {
        switch 42 {
        case 981:
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(v6, fi16, v5)
                _ = fi16
            }
        case -9223372036854775808:
            var v17 uint16 = (^(v6))
            _ = v17
        case -690:
            var v18 uint8 = 128
            _ = v18
        default:
            v6 = 42
        }
        var v19 int64 = 536
        _ = v19
    } else {
        var v20 int64 = 9223372036854775807
        _ = v20
    }
    var v21 int64 = -1
    v6 -= 65535
    if true {
        for fi22 := 0; fi22 < 8; fi22++ {
            if false {
                var v23 float32 = 0.1
                switch (("\n\t" + ((("foo bar" + "") + "a") + (("hello" + "\n\t") + "foo bar"))) + "\n\t") {
                case "\n\t":
                    fmt.Println(v6, fi22)
                }
                _ = v23
            } else {
                fmt.Println(fi22)
            }
            for fi24 := 0; fi24 < 9; fi24++ {
                v25 := GenS2{F0: -1, F1: 0, F2: 278}
                _ = fi24
                _ = v25
            }
            _ = fi22
        }
        v21 += v21
    } else {
        if true {
            v26 := GenS3{F0: 42, F1: ""}
            fmt.Println(v6, v5, v6)
            v21 = v21
            _ = v26
        } else {
            if (2 <= 1) {
                var v27 int8 = 1
                _ = v27
            }
            var v28 int16 = -204
            if true {
                v29 := GenS2{F0: 42, F1: -2147483648, F2: genFn1()}
                v30 := GenS2{F0: 38, F1: 1, F2: 2}
                var v31 uint32 = 347
                _ = v29
                _ = v30
                _ = v31
            }
            _ = v28
        }
        v21 = (^(v21))
        v32 := GenS1{F0: 2}
        _ = v32
    }
    var v33 uint = (v5 << (((v5 | v5)) & 63))
    v3[42] = -170
    var v34 int32 = -151
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
