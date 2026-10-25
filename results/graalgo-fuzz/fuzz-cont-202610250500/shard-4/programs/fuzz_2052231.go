package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint16, p1 int8, p2 string) int64 {
    if (0 >= 2147483647) {
        if false {
            var v6 uint64 = 117
            var v7 string = "a"
            _ = v6
            _ = v7
        } else {
            v8 := []int{int(p1)}
            var v9 float32 = (-2.25 + ((-2.25 * (3.14159 * -2.25)) - -2.25))
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 16; fi10++ {
            if true {
                var v11 uint32 = 337
                _ = v11
            }
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 bool = false
                v14 := GenS1{F0: 128}
                _ = fi12
                _ = v13
                _ = v14
            }
            p2 = "a"
            _ = fi10
        }
    }
    switch "hello" {
    case "a":
        var v15 uint = 255
        _ = v15
    }
    fmt.Println(p2)
    switch "foo bar" {
    case "hello":
        v16 := GenS2{F0: 801}
        _ = v16
    default:
        v17 := GenS1{F0: 127}
        _ = v17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        fmt.Println(p1, p2, p0)
        if ((0 > 4294967295) || (!true)) {
            var v19 string = p2
            _ = v19
        } else {
            if (!((!true) || true)) {
                fmt.Println(p2, fi18, p2)
            } else {
                for fi20 := 0; fi20 < 14; fi20++ {
                    p2 = p2
                    fmt.Println(p1, fi18, p0)
                    if false {
                        v21 := GenS1{F0: 127}
                        _ = v21
                    } else {
                        v22 := GenS2{F0: 4294967295}
                        _ = v22
                    }
                    var v23 uint64 = 4294967295
                    _ = fi20
                    _ = v23
                }
            }
            v24 := []int{fi18, (-(fi18))}
            _ = v24
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}
func genFn1(p0 uint16, p1 uint8, p2 uint64) uint8 {
    v25 := []int{}
    p1 = (p1 << ((1) & 63))
    v26 := GenS1{F0: (^(p1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    return (p1 / (((p1 & p1)) | 1))
}

func main() {
    var v1 uint = 377
    var v2 float32 = -2.25
    var v3 float64 = 0.0
    v4 := make(map[string]bool)
    var v5 uint32 = 1
    v27 := GenS1{F0: 42}
    for fi28 := 0; fi28 < 10; fi28++ {
        v2 += 0.0
        _ = fi28
    }
    v3 += v3
    var v29 uint8 = 128
    v30 := []int{0, 2147483647, 127, 9223372036854775807}
    if len(v30) > 0 {
        v30[0] = -2147483648
    }
    v31 := GenS3{F0: 0, F1: 42}
    for fi32 := 0; fi32 < 11; fi32++ {
        v33 := GenS1{F0: v29}
        if (42 <= 0) {
            for fi34 := 0; fi34 < 8; fi34++ {
                fmt.Println(v5)
                v30 = append(v30, -762)
                v5 += (v5 >> ((v1) & 63))
                _ = fi34
            }
            for fi35 := 0; fi35 < 11; fi35++ {
                var v36 int = (fi32 % (((fi32 / (((fi35 >> ((4294967295) & 63))) | 1))) | 1))
                var v37 string = "hello"
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        _ = fi32
        _ = v33
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 string = "\n\t"
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
