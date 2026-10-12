package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int32; F1 string }
func genFn0(p0 int16, p1 uint8, p2 int64) int16 {
    v7 := []int{-2147483648, 128, 2147483647, -917}
    var v8 int64 = p2
    p0 = p0
    var v9 string = ("" + "")
    p2 = -984
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[uint32]bool)
    var v3 string = ""
    v4 := make(map[int]int)
    v5 := make(map[string]uint32)
    var v6 int16 = 0
    v10 := GenS1{F0: 4294967295, F1: 1}
    v1 = (-(v1))
    if ((!(true || (genFn0(-428, 1, -9223372036854775808) < v6))) && ((0.1 <= float32((1e10 / 1e10))) && false)) {
        if true {
            if true {
                var v11 uint16 = 696
                v12 := GenS2{F0: 2}
                _ = v11
                _ = v12
            } else {
                var v13 uint32 = 1
                _ = v13
            }
        } else {
            v3 = (v3 + (v3 + ("\n\t" + v3)))
            v14 := GenS3{F0: 160, F1: (v3 + (v3 + v3))}
            v3 = ((("hello" + "foo bar") + v3) + v3)
            _ = v14
        }
        if false {
            var v15 uint32 = 950
            if false {
                v6 = v6
            }
            _ = v15
        }
    } else {
        v6 = (v6 >> ((2) & 63))
        for fi16 := 0; fi16 < 14; fi16++ {
            delete(v2, 2147483647)
            fmt.Println(fi16, v3)
            for fi17 := 0; fi17 < 1; fi17++ {
                fmt.Println(fi17)
                _ = fi17
            }
            _ = fi16
        }
    }
    v18 := GenS4{F0: v1, F1: ("a" + (("hello" + ((("foo bar" + ("hello" + v3)) + v3) + "\n\t")) + ((v3 + (v3 + "")) + (v3 + v3))))}
    var v19 float32 = 0.1
    for fi20 := 0; fi20 < 10; fi20++ {
        v2[65535] = (1 < -9223372036854775808)
        fmt.Println(fi20, v6)
        _ = fi20
    }
    v3 = "a"
    var v21 uint64 = 443
    switch "hello" {
    case "a":
        if true {
            for fi22 := 0; fi22 < 18; fi22++ {
                fmt.Println(fi22)
                _ = fi22
            }
        }
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        v21 -= 42
        var v24 int16 = (v6 >> ((530) & 63))
        var v25 string = "foo bar"
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 uint8 = 10
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
