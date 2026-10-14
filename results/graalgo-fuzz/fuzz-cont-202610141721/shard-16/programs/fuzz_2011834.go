package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 int16, p1 int16, p2 int) string {
    p1 = (p0 % (((p1 | p0)) | 1))
    var v6 uint64 = 685
    var v7 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "hello"
}
func genFn1(p0 int64) bool {
    p0 -= p0
    if (127 != 127) {
        for fi8 := 0; fi8 < 4; fi8++ {
            if true {
                v9 := GenS1{F0: 0, F1: "hello", F2: int(p0)}
                v10 := []int{}
                _ = v10
                _ = v9
            }
            v11 := GenS1{F0: 32767, F1: "foo bar", F2: 2147483647}
            if (0 >= 127) {
                var v12 uint64 = 326
                _ = v12
            } else {
                switch "\n\t" {
                case "a":
                    for fi13 := 0; fi13 < 3; fi13++ {
                        var v14 uint = 1
                        v14 = (^(v14))
                        _ = fi13
                        _ = v14
                    }
                }
            }
            _ = fi8
            _ = v11
        }
    }
    if (!("foo bar" != "\n\t")) {
        v15 := GenS2{F0: (!(!(42 <= 50))), F1: 0}
        v16 := GenS1{F0: -1, F1: "a", F2: 378}
        _ = v15
        _ = v16
    } else {
        p0 = -9223372036854775808
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 int16 = int16(fi17)
        fmt.Println(fi17, fi17, p0)
        _ = fi17
        _ = v18
    }
    if true {
        v19 := []int{-820, 4294967295, 0, 128}
        _ = v19
    }
    _ = p0
    return (2147483648 != 2147483648)
}

func main() {
    var v1 uint = 0
    v2 := []int{}
    var v3 int8 = 0
    v4 := make(map[string]uint32)
    var v5 int16 = 32767
    var v20 int64 = -1
    fmt.Println(v20)
    if (2147483647 == -1) {
        v21 := GenS2{F0: false, F1: 2147483647}
        _ = v21
    } else {
        v20 = v20
    }
    var v22 uint64 = 18446744073709551615
    v3 += 127
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 uint32 = 930
        v3 += int8(v22)
        _ = fi23
        _ = v24
    }
    if len(v2) > 0 {
        v2[0] = 32768
    }
    v25 := v4["hello"]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
