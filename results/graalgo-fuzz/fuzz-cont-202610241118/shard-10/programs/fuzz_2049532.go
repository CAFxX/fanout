package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int) uint8 {
    if false {
        fmt.Println(p0)
        var v5 float32 = float32(-2.25)
        _ = v5
    } else {
        fmt.Println(p0)
        p0 += (^(p0))
        v6 := GenS1{F0: 1}
        _ = v6
    }
    switch (-(p0)) {
    case 0:
        if true {
            var v7 float32 = 0.0
            v8 := GenS1{F0: -1}
            _ = v7
            _ = v8
        } else {
            var v9 float64 = 1.5
            _ = v9
        }
    }
    var v10 int8 = -111
    _ = p0
    _ = v10
    return 42
}
func genFn1(p0 uint64, p1 int32, p2 int64) uint16 {
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(p1, p0)
        var v12 float64 = (1e10 / 0.1)
        _ = fi11
        _ = v12
    }
    var v13 float64 = 1e10
    switch "\n\t" {
    case "\n\t":
        var v14 uint16 = 2
        _ = v14
    case "hello":
        if (32768 >= 9223372036854775807) {
            v15 := GenS1{F0: p1}
            v16 := []int{-2147483648, 135, int(p2), 9223372036854775807}
            _ = v15
            _ = v16
        } else {
            if true {
                v17 := []int{9223372036854775807, -737, 2147483647}
                var v18 float64 = ((0.1 + 1e10) / float64((float32(float64(float32((v13 * ((1e10 - v13) / 1.0))))) - 0.0)))
                fmt.Println(v13, v13)
                _ = v17
                _ = v18
            }
        }
    case "":
        fmt.Println(p2)
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        if true {
            p0 += (p0 & p0)
        }
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 593
}

func main() {
    v1 := []int{128, 996}
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    v4 := []uint32{2, 438, 580, 491}
    switch "" {
    case "foo bar":
        var v20 uint = 255
        _ = v20
    }
    v21 := GenS1{F0: 2147483647}
    v22 := GenS2{F0: -727}
    v23 := GenS3{F0: 4294967295}
    var v24 uint16 = 552
    var v25 uint16 = (v24 ^ v24)
    for fi26 := 0; fi26 < 15; fi26++ {
        fmt.Println(v25)
        if true {
            var v27 int = 1
            fmt.Println(v24, fi26)
            if true {
                for fi28 := 0; fi28 < 20; fi28++ {
                    if len(v4) > 0 {
                        v4[0] = 2147483648
                    }
                    var v29 string = (("a" + ("\n\t" + "")) + "\n\t")
                    _ = fi28
                    _ = v29
                }
                var v30 uint32 = 25
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v30
            }
            _ = v27
        }
        _ = fi26
    }
    var v31 float32 = -2.25
    for fi32 := 0; fi32 < 10; fi32++ {
        v25 = v25
        var v33 string = "a"
        _ = fi32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v4
}
