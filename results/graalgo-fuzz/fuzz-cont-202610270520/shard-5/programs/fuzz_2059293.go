package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint32 }
func genFn0(p0 uint) uint16 {
    if (!true) {
        if (32767 < 441) {
            switch "" {
            case "foo bar":
                v7 := GenS1{F0: 723, F1: 0, F2: 45}
                _ = v7
            }
            v8 := []int{-221, 127}
            fmt.Println(p0)
            _ = v8
        } else {
            var v9 int = -2147483648
            _ = v9
        }
        v10 := []int{1, 290, 2147483647}
        _ = v10
    }
    fmt.Println(p0)
    for fi11 := 0; fi11 < 20; fi11++ {
        if false {
            var v12 int32 = 293
            _ = v12
        } else {
            var v13 uint16 = 0
            if true {
                var v14 bool = false
                var v15 uint8 = 128
                _ = v14
                _ = v15
            } else {
                v16 := GenS1{F0: 42, F1: -1, F2: 65535}
                var v17 uint8 = 209
                v18 := []int{(fi11 / ((625) | 1)), int(v13), (fi11 & fi11), fi11}
                _ = v16
                _ = v17
                _ = v18
            }
            _ = v13
        }
        var v19 uint32 = 4294967295
        _ = fi11
        _ = v19
    }
    _ = p0
    return 65535
}
func genFn1(p0 int16, p1 bool, p2 int64) uint64 {
    var v20 uint32 = 2147483647
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := []int{4294967295, fi21, fi21}
        var v23 uint64 = 4294967295
        fmt.Println(p2, v23)
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 2
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int32 = 0
    v3 := make(map[uint32]bool)
    v4 := make(map[string]string)
    v5 := []int{42, 127, 4294967295}
    var v6 bool = false
    var v24 uint16 = 928
    v24 = v24
    var v25 string = ("\n\t" + "a")
    v2 = v2
    if (!v6) {
        switch ((v25 + "a") + "hello") {
        case "":
            v25 = (v25 + v25)
        }
    }
    var v26 int = -1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch v25 {
    case "\n\t":
        var v27 int64 = int64(v24)
        _ = v27
    default:
        if false {
            var v28 string = v25
            var v29 int8 = 100
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v28
            _ = v29
        }
    }
    v3[4294967295] = true
    v2 += (v2 << ((4294967295) & 63))
    fmt.Println(v1, v1, v26)
    v5 = append(v5, (v26 >> ((255) & 63)))
    var v30 string = (v25 + "")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
