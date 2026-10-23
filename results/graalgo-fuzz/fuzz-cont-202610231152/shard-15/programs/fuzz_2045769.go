package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 string; F1 uint16 }
func genFn0(p0 int8, p1 uint32, p2 bool) bool {
    var v5 float32 = 3.14159
    if (127 <= 551) {
        v6 := GenS1{F0: 100}
        p1 = p1
        _ = v6
    } else {
        v7 := GenS2{F0: -1}
        var v8 int64 = -581
        var v9 int8 = int8(v8)
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ((!p2) && p2)
}
func genFn1(p0 uint8, p1 int16, p2 int) int64 {
    fmt.Println(p2)
    var v10 string = ""
    for fi11 := 0; fi11 < 2; fi11++ {
        fmt.Println(p0, p1)
        if (true || true) {
            v12 := []int{32768, fi11}
            if true {
                var v13 int = fi11
                var v14 int8 = -128
                _ = v13
                _ = v14
            }
            _ = v12
        } else {
            var v15 uint64 = uint64(p0)
            v15 = v15
            for fi16 := 0; fi16 < 2; fi16++ {
                p0 = 74
                v17 := GenS2{F0: 42}
                if true {
                    var v18 uint = 42
                    v19 := GenS1{F0: 74}
                    var v20 int = fi16
                    _ = v18
                    _ = v19
                    _ = v20
                }
                p2 += (fi16 + (fi11 + -9223372036854775808))
                var v21 int = (p2 - fi11)
                v22 := GenS3{F0: 42, F1: 42}
                var v23 bool = true
                _ = fi16
                _ = v17
                _ = v21
                _ = v22
                _ = v23
            }
            _ = v15
        }
        _ = fi11
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := GenS3{F0: 0, F1: 1}
        v26 := []int{}
        fmt.Println(p1, p0)
        _ = fi24
        _ = v25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -9223372036854775808
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[string]string)
    var v3 uint = uint(v1)
    var v4 uint8 = 127
    var v27 int16 = 1
    fmt.Println(v27)
    fmt.Println(v3, v1)
    v4 += 133
    for fi28 := 0; fi28 < 17; fi28++ {
        _ = len(v2)
        _ = fi28
    }
    v4 = v4
    v29 := []int{127, -2147483648}
    v30 := GenS4{F0: "hello", F1: 32767}
    var v31 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
