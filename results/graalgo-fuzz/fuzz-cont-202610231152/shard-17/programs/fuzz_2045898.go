package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint32 }
func genFn0(p0 int32) uint {
    v7 := []int{1}
    var v8 float32 = 3.14159
    v9 := GenS1{F0: 1, F1: 364, F2: ((("foo bar" + ("a" + "hello")) + "a") + "hello")}
    v8 = (v8 / 1e10)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("\n\t" + "a")
    var v3 int64 = 477
    var v4 uint16 = 721
    v5 := []string{("hello" + "\n\t"), v2}
    var v6 uint8 = 42
    v10 := GenS1{F0: -1, F1: 18446744073709551615, F2: (v2 + "")}
    fmt.Println(v2, v4)
    _ = len(v5)
    for fi11 := 0; fi11 < 1; fi11++ {
        v1[fi11] = true
        fmt.Println(fi11)
        if (!false) {
            v12 := GenS2{F0: v4}
            _ = v12
        }
        _ = fi11
    }
    fmt.Println(v6)
    v4 = v4
    for fi13 := 0; fi13 < 13; fi13++ {
        fmt.Println(v2, v4, v3)
        if false {
            for fi14 := 0; fi14 < 5; fi14++ {
                switch "\n\t" {
                case "\n\t":
                    var v15 int = (fi14 & (fi13 | 9223372036854775807))
                    _ = v15
                }
                var v16 uint8 = (v6 | v6)
                _ = fi14
                _ = v16
            }
        }
        _ = fi13
    }
    if (128 >= 32767) {
        v17 := GenS3{F0: 0, F1: 42}
        var v18 uint64 = 944
        _ = v17
        _ = v18
    }
    v6 -= v6
    var v19 uint = 255
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := GenS3{F0: (v3 % ((42) | 1)), F1: 0}
        _ = fi20
        _ = v21
    }
    var v22 int16 = 1
    if false {
        v23 := GenS4{F0: v3, F1: 1}
        _ = v23
    }
    fmt.Println(v3)
    v5 = append(v5, ((v2 + (v2 + v2)) + "a"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
