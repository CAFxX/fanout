package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint32; F1 string }
func genFn0(p0 int64, p1 uint16, p2 uint16) uint16 {
    p1 += (p1 + 65535)
    fmt.Println(p1)
    if false {
        v4 := GenS1{F0: -1, F1: 65535}
        _ = v4
    }
    v5 := GenS2{F0: (("\n\t" + "hello") + ("a" + ("hello" + "hello"))), F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 65535
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 987
    v3 := make(map[int]string)
    var v6 uint8 = 42
    v1[9223372036854775807] = 0
    fmt.Println(v2, v6)
    var v7 int8 = -32
    v7 = v7
    v6 -= v6
    fmt.Println(v7, v2)
    fmt.Println(v7)
    var v8 bool = false
    if v8 {
        v7 += v7
        if false {
            var v9 uint16 = 423
            _ = v9
        } else {
            switch "a" {
            case "foo bar":
                switch -9223372036854775808 {
                case 9223372036854775807:
                    v2 -= v2
                case 224:
                    var v10 uint64 = 0
                    _ = v10
                case 2147483647:
                    delete(v1, 127)
                default:
                    fmt.Println(v7, v8, v6)
                }
            }
            v11 := GenS3{F0: 32768, F1: (v2 & v2)}
            v12 := []int{}
            _ = v11
            _ = v12
        }
        fmt.Println(v6)
    } else {
        fmt.Println(v8, v7)
        for fi13 := 0; fi13 < 1; fi13++ {
            for fi14 := 0; fi14 < 17; fi14++ {
                if v8 {
                    v15 := GenS4{F0: v2, F1: ("\n\t" + "a")}
                    v7 -= v7
                    fmt.Println(fi13, fi13)
                    _ = v15
                }
                v16 := GenS4{F0: v2, F1: ("a" + ("foo bar" + "a"))}
                _ = fi14
                _ = v16
            }
            fmt.Println(v2, v6, fi13)
            _ = fi13
        }
    }
    fmt.Println(v8, v8)
    var v17 uint = 4294967295
    fmt.Println(v6)
    var v18 bool = v8
    v7 += v7
    for fi19 := 0; fi19 < 17; fi19++ {
        if (!v18) {
            v20 := GenS3{F0: -738, F1: v2}
            v21 := v1[-2147483648]
            var v22 uint64 = 42
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    fmt.Println(v2, v18)
    var v23 float64 = ((-2.25 / 1e10) - 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
