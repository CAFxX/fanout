package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 bool, p1 int) uint16 {
    if (((!(-199 == -489)) || true) || p0) {
        var v5 uint = 2
        _ = v5
    } else {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 float32 = (1.5 / -2.25)
            v8 := []int{-385}
            for fi9 := 0; fi9 < 4; fi9++ {
                fmt.Println(p0, p1)
                if (((0.1 * (0.0 + float64(v7))) / 1e10) <= 3.14159) {
                    p1 += (-(fi9))
                }
                _ = fi9
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 string = "\n\t"
            var v12 float32 = -2.25
            v13 := []int{(-(p1)), fi10, p1, 0}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        p1 = p1
    }
    switch (p1 + -1) {
    case 1:
        var v14 bool = false
        _ = v14
    case 2147483647:
        v15 := []int{20, (p1 << ((817) & 63)), (-(p1))}
        _ = v15
    }
    _ = p0
    _ = p1
    return 255
}
func genFn1(p0 string, p1 bool) bool {
    v16 := []int{-2147483648}
    p1 = p1
    p0 = p0
    var v17 uint = 2
    p0 = ("hello" + p0)
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return p1
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int32 = -2147483648
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v18 int32 = (^(v2))
    switch "foo bar" {
    case "a":
        v19 := []int{2147483647, -2147483648, 32767}
        _ = v19
    case "":
        v1 = 1
    case "hello":
        v20 := GenS1{F0: 0}
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        fmt.Println(v18)
        v2 = v18
        fmt.Println(v2, v1)
        _ = fi21
    }
    v22 := GenS1{F0: 1}
    var v23 float32 = 0.0
    fmt.Println(v1, v23, v1)
    var v24 bool = (!false)
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 uint8 = 132
        v23 = v23
        var v27 int8 = 100
        _ = fi25
        _ = v26
        _ = v27
    }
    v4[""] = 2147483647
    for fi28 := 0; fi28 < 17; fi28++ {
        if (1 <= 67) {
            var v29 int32 = 672
            v30 := GenS1{F0: 0}
            v18 += v29
            _ = v29
            _ = v30
        } else {
            v31 := []int{fi28, 32768, -215}
            _ = v31
        }
        if v24 {
            for fi32 := 0; fi32 < 11; fi32++ {
                delete(v3, (fi28 / ((fi28) | 1)))
                _ = fi32
            }
        }
        fmt.Println(v1)
        _ = fi28
    }
    switch "a" {
    case "":
        for fi33 := 0; fi33 < 3; fi33++ {
            v34 := GenS2{F0: 4294967295}
            _ = fi33
            _ = v34
        }
    default:
        for fi35 := 0; fi35 < 15; fi35++ {
            for fi36 := 0; fi36 < 19; fi36++ {
                v2 = (v18 << ((18446744073709551615) & 63))
                var v37 int64 = -382
                fmt.Println(v37, v18)
                _ = fi36
                _ = v37
            }
            v38 := GenS1{F0: 1}
            v23 = v23
            _ = fi35
            _ = v38
        }
    }
    var v39 int = 9223372036854775807
    var v40 int32 = (v18 & (^(v18)))
    if true {
        switch 4294967295 {
        case -1:
            v41 := GenS2{F0: 927}
            _ = v41
        }
        v3[v39] = 838
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}
