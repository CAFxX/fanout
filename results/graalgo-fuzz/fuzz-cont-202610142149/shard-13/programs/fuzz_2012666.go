package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
type GenS3 struct { F0 int32; F1 uint }
func genFn0(p0 int8, p1 string, p2 uint16) bool {
    var v6 int32 = 42
    for fi7 := 0; fi7 < 12; fi7++ {
        if false {
            v8 := []int{fi7, -871, -119, fi7}
            _ = v8
        }
        switch (p1 + p1) {
        case "a":
            p2 = (p2 >> ((65535) & 63))
        }
        _ = fi7
    }
    v9 := []int{128, 127, 32767, -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return true
}
func genFn1(p0 uint16, p1 uint16, p2 uint16) bool {
    v10 := GenS1{F0: 128}
    v11 := []int{4294967295, -284, -1}
    v12 := GenS2{F0: -1, F1: 0, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return true
}

func main() {
    v1 := []string{"hello", "\n\t", ("" + ((("\n\t" + "") + "a") + ("a" + "foo bar"))), "a"}
    v2 := make(map[string]int)
    var v3 float32 = 3.14159
    v4 := make(map[uint32]bool)
    v5 := make(map[int]int)
    v13 := GenS3{F0: 0, F1: 18446744073709551615}
    v14 := GenS3{F0: -81, F1: 4294967295}
    for fi15 := 0; fi15 < 9; fi15++ {
        switch (fi15 + -533) {
        case 32767:
            for fi16 := 0; fi16 < 6; fi16++ {
                fmt.Println(v3, fi16, fi16)
                _ = fi16
            }
        case 128:
            v3 = ((v3 - 1e10) - (v3 + (-2.25 + float32(0.1))))
        }
        if (-150 <= 440) {
            v17 := GenS1{F0: 0}
            v18 := v5[-664]
            _ = v17
            _ = v18
        } else {
            fmt.Println(fi15, fi15)
            if false {
                v19 := GenS1{F0: fi15}
                _ = v19
            }
            v3 = (v3 * v3)
        }
        _ = fi15
    }
    fmt.Println(v3)
    delete(v2, ("\n\t" + "hello"))
    var v20 int8 = -102
    var v21 uint = 18446744073709551615
    switch ("hello" + "a") {
    case "a":
        v2[("" + "")] = 9223372036854775807
    }
    var v22 int = -485
    v5[v22] = 9223372036854775807
    if false {
        var v23 int = 0
        var v24 int16 = -32768
        _ = v23
        _ = v24
    }
    fmt.Println(v3, v3)
    v3 += v3
    var v25 uint64 = 2
    fmt.Println(v3)
    v25 += (v25 >> ((729) & 63))
    var v26 float64 = (-2.25 + 1.5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
