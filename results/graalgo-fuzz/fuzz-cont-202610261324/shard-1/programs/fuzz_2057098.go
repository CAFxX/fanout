package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint16, p1 int32) int8 {
    p1 = p1
    switch "hello" {
    case "":
        v7 := GenS1{F0: 643}
        _ = v7
    }
    p0 += p0
    switch "foo bar" {
    case "hello":
        v8 := []int{385, 127, 1, 32768}
        _ = v8
    case "foo bar":
        var v9 float32 = 0.1
        _ = v9
    case "a":
        v10 := []int{}
        _ = v10
    default:
        p1 = (p1 ^ p1)
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]bool)
    var v3 int16 = 32767
    var v4 int64 = -1
    v5 := make(map[uint32]string)
    var v6 int8 = 42
    fmt.Println(v4, v4)
    v4 = v4
    v4 = (v4 % ((v4) | 1))
    v11 := GenS1{F0: 198}
    v6 += int8(v3)
    v6 += v6
    if true {
        v6 = (-(v6))
        _ = len(v5)
    } else {
        var v12 int8 = v6
        v4 -= -1
        _ = v12
    }
    fmt.Println(v6, v4, v3)
    v13 := []int{int(v4), 692, -307}
    var v14 uint64 = 4294967295
    switch -9223372036854775808 {
    case -232:
        var v15 int8 = v6
        _ = v15
    case 9223372036854775807:
        var v16 uint16 = 42
        _ = v16
    }
    switch 128 {
    case 4294967295:
        v17 := GenS1{F0: 2}
        _ = v17
    case 32767:
        for fi18 := 0; fi18 < 13; fi18++ {
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 uint32 = uint32(v14)
                v21 := GenS1{F0: 65535}
                _ = fi19
                _ = v20
                _ = v21
            }
            if false {
                if (!true) {
                    v4 += v4
                }
            }
            _ = fi18
        }
    default:
        for fi22 := 0; fi22 < 17; fi22++ {
            var v23 int32 = 0
            v4 += -9223372036854775808
            _ = fi22
            _ = v23
        }
    }
    fmt.Println(v14)
    fmt.Println(v3, v1)
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
