package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint64; F1 int64; F2 string }
func genFn0(p0 uint16) bool {
    var v4 int16 = -1
    var v5 int64 = 0
    if (!(((^(p0)) != 0) && false)) {
        var v6 int8 = -1
        if (!(p0 > 1)) {
            var v7 int32 = -2147483648
            _ = v7
        } else {
            v8 := []int{0, 4294967295}
            var v9 uint = 42
            _ = v8
            _ = v9
        }
        _ = v6
    }
    _ = p0
    _ = v4
    _ = v5
    return true
}
func genFn1(p0 string) uint8 {
    v10 := GenS1{F0: 1, F1: 32768}
    for fi11 := 0; fi11 < 19; fi11++ {
        if false {
            var v12 int16 = 1
            _ = v12
        } else {
            switch p0 {
            case "foo bar":
                v13 := []int{fi11, (-(fi11)), (^(fi11))}
                _ = v13
            }
            var v14 int32 = 1
            _ = v14
        }
        _ = fi11
    }
    _ = p0
    _ = v10
    return 1
}

func main() {
    v1 := []string{"hello", ""}
    var v2 int16 = 993
    var v3 int64 = -627
    _ = len(v1)
    for fi15 := 0; fi15 < 10; fi15++ {
        v2 += (v2 << ((1) & 63))
        v2 = (v2 + int16(v3))
        _ = fi15
    }
    v16 := GenS1{F0: 255, F1: 32768}
    v17 := GenS2{F0: 65535, F1: 140}
    v3 = (^(v3))
    v3 = (v3 << ((374) & 63))
    var v18 float32 = -2.25
    v19 := GenS3{F0: 1, F1: -699, F2: "hello"}
    fmt.Println(v18, v18)
    v18 = -2.25
    var v20 uint16 = 935
    var v21 int8 = -122
    var v22 uint32 = 0
    fmt.Println(v21, v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
