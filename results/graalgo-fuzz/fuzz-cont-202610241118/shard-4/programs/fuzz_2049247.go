package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int8, p1 uint) uint8 {
    v6 := GenS1{F0: -109}
    fmt.Println(p0)
    var v7 string = ""
    var v8 int64 = 9223372036854775807
    if ((p0 == (p0 ^ p0)) && false) {
        fmt.Println(v7, p1, p0)
        p1 = (p1 + p1)
        var v9 int16 = 1
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 133
}
func genFn1() int16 {
    var v10 uint = 638
    v11 := []int{894, -407, 127}
    v12 := GenS2{F0: 707, F1: 0}
    fmt.Println(v10)
    if (759 > v10) {
        v13 := GenS3{F0: 0}
        _ = v13
    } else {
        var v14 string = "hello"
        var v15 uint32 = 65535
        _ = v14
        _ = v15
    }
    _ = v10
    _ = v11
    _ = v12
    return 662
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 int8 = -1
    v5 := make(map[int]string)
    v16 := GenS2{F0: -32768, F1: 42}
    v1 -= v1
    var v17 uint8 = genFn0(-128, 848)
    v17 -= v17
    var v18 uint32 = 2
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 string = (("foo bar" + "a") + ("a" + ""))
        _ = fi19
        _ = v20
    }
    fmt.Println(v1)
    for fi21 := 0; fi21 < 13; fi21++ {
        if ((v18 & 2147483648) != (v18 >> ((197) & 63))) {
            v4 = -1
            v4 = (v4 >> ((4294967295) & 63))
        } else {
            var v22 float32 = (((v1 - float32(float64(float32(3.14159)))) * float32(0.1)) - v1)
            v3[""] = true
            _ = v22
        }
        if (false || true) {
            v4 += (v4 ^ 127)
        } else {
            v17 += (^(v17))
        }
        _ = fi21
    }
    var v23 bool = (true || false)
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := GenS4{F0: v17}
        if v23 {
            for fi26 := 0; fi26 < 8; fi26++ {
                v27 := []int{42, (fi26 - fi26), 128, fi24}
                if false {
                    v23 = v23
                    var v28 int8 = 42
                    _ = v28
                } else {
                    var v29 int = (fi26 ^ -9223372036854775808)
                    _ = v29
                }
                _ = fi26
                _ = v27
            }
        }
        _ = fi24
        _ = v25
    }
    var v30 int16 = -262
    var v31 string = (("a" + "\n\t") + ("a" + ("\n\t" + "hello")))
    fmt.Println(v1, v4)
    var v32 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
