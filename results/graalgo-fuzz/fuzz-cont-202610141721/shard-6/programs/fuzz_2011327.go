package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 string; F2 int32 }
func genFn0(p0 int32, p1 int16, p2 uint) int {
    fmt.Println(p2, p2, p1)
    switch 2147483647 {
    case 161:
        p0 = p0
    case 721:
        switch -471 {
        case -145:
            p1 = (p1 % (((^(p1))) | 1))
        default:
            fmt.Println(p2)
        }
    }
    if false {
        v5 := []int{42, 32768, 32767}
        var v6 uint16 = 65535
        _ = v5
        _ = v6
    }
    v7 := []int{}
    var v8 int32 = (-(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32768
}
func genFn1(p0 uint32, p1 uint8, p2 bool) int64 {
    var v9 bool = (p2 || p2)
    switch 112 {
    case 2147483647:
        p2 = (32767 != 333)
    default:
        p2 = v9
    }
    p2 = (1 != 1000)
    p1 = (p1 >> ((424) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 uint32 = 172
    v4 := make(map[string]int)
    var v10 uint8 = 0
    v11 := []int{32767, -285, -543, 1}
    v11 = append(v11, 356)
    var v12 int32 = 317
    v13 := GenS1{F0: "\n\t"}
    switch 9223372036854775807 {
    case 127:
        if false {
            fmt.Println(v3, v12, v3)
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := GenS2{F0: 42, F1: "foo bar", F2: (v12 >> ((647) & 63))}
                v3 = 2147483648
                var v16 uint16 = 190
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            var v17 uint64 = 42
            if true {
                v18 := GenS1{F0: "hello"}
                for fi19 := 0; fi19 < 17; fi19++ {
                    v11 = append(v11, fi19)
                    var v20 uint32 = v3
                    _ = fi19
                    _ = v20
                }
                fmt.Println(v3, v12)
                _ = v18
            } else {
                fmt.Println(v10)
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 string = "a"
                    v23 := GenS2{F0: 255, F1: (v22 + "hello"), F2: v12}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = len(v2)
            }
            _ = v17
        }
    case -2147483648:
        fmt.Println(v3, v10, v10)
    }
    v10 = 104
    v12 += v12
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
