package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 string }
type GenS2 struct { F0 string; F1 int16; F2 string }
type GenS3 struct { F0 int8; F1 int64 }
func genFn0(p0 int32, p1 int64) int64 {
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 bool = true
        p1 = (p1 & (p1 / ((-516) | 1)))
        v8 := GenS1{F0: -62, F1: v7, F2: ("hello" + "hello")}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{-65}
    _ = p0
    _ = p1
    _ = v9
    return -652
}
func genFn1(p0 int, p1 string) string {
    var v10 int = (p0 - p0)
    p1 = "a"
    _ = p0
    _ = p1
    _ = v10
    return (((p1 + (p1 + ("a" + ("\n\t" + "")))) + ((p1 + (p1 + p1)) + "foo bar")) + "foo bar")
}

func main() {
    var v1 uint8 = 2
    v2 := []int{4294967295, 32768, 42, -805}
    v3 := make(map[int]bool)
    var v4 int = 2147483647
    var v5 uint32 = 2147483648
    delete(v3, 412)
    v5 += (v5 ^ (v5 ^ v5))
    if false {
        v2 = append(v2, v4)
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        if len(v2) > 0 {
            v2[0] = v4
        }
        _ = fi11
    }
    v1 = (v1 - 42)
    var v12 float64 = 1e10
    var v13 uint64 = 4294967295
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v14 float64 = 0.0
    for fi15 := 0; fi15 < 10; fi15++ {
        v1 -= 127
        fmt.Println(fi15, v1, v4)
        _ = fi15
    }
    if true {
        v13 = 672
    } else {
        v5 -= v5
    }
    v16 := GenS2{F0: ("" + ""), F1: -724, F2: "hello"}
    if false {
        if false {
            if true {
                v17 := GenS3{F0: 3, F1: 299}
                if true {
                    fmt.Println(v4, v13, v4)
                }
                _ = v17
            }
        }
    }
    v18 := GenS3{F0: -1, F1: -9223372036854775808}
    v1 += (v1 * (v1 / ((255) | 1)))
    fmt.Println(v1, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
