package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0() uint {
    v6 := []int{}
    if (-32768 <= -1) {
        var v7 float32 = 3.14159
        _ = v7
    } else {
        var v8 int16 = -403
        v8 -= v8
        _ = v8
    }
    var v9 uint16 = 892
    var v10 int8 = -1
    _ = v10
    _ = v6
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 uint32) int16 {
    if (!false) {
        fmt.Println(p0)
        var v11 string = "hello"
        _ = v11
    } else {
        if (!((!true) && (1 > -31))) {
            fmt.Println(p0)
        } else {
            v12 := GenS1{F0: 237}
            v13 := []int{4294967295, -434, 2147483647}
            p0 = 2
            _ = v12
            _ = v13
        }
        v14 := []int{}
        _ = v14
    }
    var v15 string = "\n\t"
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 int8 = 100
        var v18 uint16 = 427
        var v19 string = "a"
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = v15
    return -677
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 int16 = 0
    v4 := []string{"a", ""}
    var v5 uint32 = 166
    var v20 int32 = 912
    var v21 uint32 = (v5 * 1)
    v20 = (v20 ^ -2147483648)
    if false {
        v21 += (v21 / ((v5) | 1))
    }
    fmt.Println(v20, v20)
    if false {
        v5 += (^(v5))
        var v22 bool = (0.1 > 1.5)
        _ = v22
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        fmt.Println(v5, v5, v21)
        var v24 uint64 = 18446744073709551615
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            v21 -= v21
            switch "a" {
            case "a":
                switch "\n\t" {
                case "\n\t":
                    v5 += v5
                }
            }
            if (540 < 29) {
                var v27 uint16 = 0
                if (true || ((42 >= 2) || ((false && (127 >= fi26)) || (false || false)))) {
                    v28 := GenS1{F0: (^(v27))}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v28
                }
                v29 := GenS2{F0: v5}
                v30 := GenS2{F0: 2147483648}
                _ = v27
                _ = v29
                _ = v30
            } else {
                fmt.Println(fi26, v3)
                _ = len(v4)
                var v31 uint = 906
                _ = len(v2)
                var v32 int32 = (v20 ^ v20)
                _ = v31
                _ = v32
            }
            _ = fi26
        }
        _ = fi25
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        v3 -= v3
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        var v34 int32 = v20
        _ = fi33
        _ = v34
    }
    var v35 int8 = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        fmt.Println(v5, fi36, v35)
        var v37 int64 = 0
        _ = fi36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
