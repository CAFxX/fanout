package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int8; F2 int8 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint32, p1 uint64) uint8 {
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := GenS1{F0: -54, F1: ((fi4 / ((109) | 1)) != 908)}
        if (true || true) {
            v6 := []int{127, (fi4 >> ((18446744073709551615) & 63)), fi4}
            p1 -= 0
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    v7 := GenS2{F0: 113, F1: p0}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 2
    v3 := make(map[int]uint32)
    fmt.Println(v2)
    v8 := []int{1, -9223372036854775808, -598, 9223372036854775807}
    var v9 int64 = -1
    for fi10 := 0; fi10 < 8; fi10++ {
        v2 = v2
        _ = fi10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 uint32 = 65535
            v13 -= (v13 >> ((42) & 63))
            fmt.Println(v2, fi11)
            _ = fi12
            _ = v13
        }
        switch "hello" {
        case "a":
            v9 = (v9 + v9)
        default:
            if true {
                v14 := GenS3{F0: -706, F1: 0}
                _ = v14
            } else {
                if (42 == 0) {
                    v9 -= v9
                    v2 = (v2 & v2)
                }
                v15 := GenS3{F0: 1, F1: 4294967295}
                _ = v15
            }
        }
        switch (fi11 % (((fi11 + fi11)) | 1)) {
        case 2147483647:
            var v16 int16 = 42
            _ = v16
        }
        _ = fi11
    }
    fmt.Println(v2)
    var v17 uint = 432
    if (false && (((!(!true)) && false) || (127 != int(v2)))) {
        var v18 string = ""
        fmt.Println(v18, v2, v9)
        var v19 uint32 = 568
        _ = v18
        _ = v19
    }
    if (false || true) {
        v9 = 408
        var v20 uint64 = uint64(v2)
        var v21 uint = (v17 | v17)
        _ = v20
        _ = v21
    }
    v2 += v2
    v22 := GenS4{F0: 32767, F1: -1, F2: 127}
    v3[9223372036854775807] = (^(v2))
    var v23 int64 = 538
    switch ("" + "a") {
    case "foo bar":
        v24 := GenS5{F0: v2}
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v8
    _ = v9
}
