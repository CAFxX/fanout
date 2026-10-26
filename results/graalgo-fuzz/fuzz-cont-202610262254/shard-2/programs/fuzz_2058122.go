package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int64; F1 uint64; F2 uint16 }
func genFn0(p0 string, p1 int, p2 int16) bool {
    v4 := []int{p1, p1, (p1 & (p1 << ((428) & 63))), -805}
    switch (-(p1)) {
    case 2147483647:
        p0 = ""
    }
    if false {
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 int8 = -1
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: 42, F1: 0, F2: 54}
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := GenS2{F0: 128, F1: (((p0 + "hello") + ("hello" + "a")) + "")}
            var v10 uint64 = 0
            p1 = p1
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ((p0 + p0) < ("foo bar" + (p0 + (((p0 + (("foo bar" + "\n\t") + p0)) + (p0 + (p0 + (p0 + p0)))) + "\n\t"))))
}
func genFn1(p0 int32) uint {
    switch 2147483647 {
    case 4294967295:
        var v11 uint16 = 589
        _ = v11
    case -29:
        p0 = (-(p0))
    }
    var v12 string = ("foo bar" + (("hello" + "a") + ""))
    v13 := []int{4294967295}
    _ = p0
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]bool)
    var v3 uint16 = 212
    var v14 float64 = 0.0
    v15 := []int{-1}
    v3 = (v3 ^ (^(v3)))
    for fi16 := 0; fi16 < 5; fi16++ {
        v3 -= 0
        var v17 uint = 65535
        _ = fi16
        _ = v17
    }
    var v18 uint16 = 1
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 uint64 = 42
            v3 += (v3 << ((42) & 63))
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    if true {
        v22 := GenS2{F0: 275, F1: v1}
        var v23 int = -189
        var v24 int64 = 9223372036854775807
        _ = v22
        _ = v23
        _ = v24
    }
    var v25 bool = true
    var v26 uint64 = 2
    v27 := GenS3{F0: 9223372036854775807, F1: 766, F2: (v3 * 65535)}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
