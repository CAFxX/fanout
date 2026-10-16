package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int, p1 uint, p2 string) uint16 {
    p2 = ""
    var v4 uint16 = 202
    v5 := GenS1{F0: p2, F1: 4294967295}
    p2 = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return v4
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]int)
    var v3 uint8 = 48
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            delete(v2, 2147483647)
            v1 -= 0
            v3 += (^(v3))
            _ = fi7
        }
        _ = fi6
    }
    var v8 uint8 = v3
    v9 := GenS2{F0: 1, F1: (v1 * 4294967295)}
    var v10 int64 = -1
    var v11 string = "foo bar"
    if true {
        var v12 bool = true
        var v13 int16 = 1000
        v10 = (v10 * (v10 + int64(v3)))
        _ = v12
        _ = v13
    } else {
        if true {
            fmt.Println(v11, v1)
        } else {
            fmt.Println(v8, v3)
            var v14 uint32 = 124
            v15 := GenS2{F0: 42, F1: 53}
            _ = v14
            _ = v15
        }
        fmt.Println(v3)
        if true {
            for fi16 := 0; fi16 < 11; fi16++ {
                v17 := GenS1{F0: ("" + ("" + (v11 + v11))), F1: v1}
                v18 := GenS1{F0: (v11 + ((v11 + v11) + (v11 + ((("hello" + "hello") + (("a" + v11) + ("\n\t" + v11))) + "\n\t")))), F1: v1}
                fmt.Println(v3, fi16, fi16)
                _ = fi16
                _ = v17
                _ = v18
            }
        } else {
            if ((!(-2.25 <= 0.1)) && (1 < -32768)) {
                var v19 int16 = 759
                _ = v19
            } else {
                var v20 uint64 = v1
                if (3.14159 != 0.1) {
                    var v21 int32 = -999
                    for fi22 := 0; fi22 < 17; fi22++ {
                        fmt.Println(fi22)
                        v23 := GenS1{F0: (v11 + (v11 + "")), F1: 1}
                        _ = fi22
                        _ = v23
                    }
                    fmt.Println(v1, v20)
                    _ = v21
                } else {
                    if ((v3 | 255) <= (v8 & v3)) {
                        v8 = 44
                        if (2 < 42) {
                            var v24 string = "\n\t"
                            _ = v24
                        }
                    }
                }
                _ = v20
            }
        }
    }
    v25 := GenS3{F0: -2147483648}
    if true {
        fmt.Println(v11)
        var v26 uint32 = 2147483648
        v26 = (v26 * v26)
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        v1 = (v1 >> ((807) & 63))
        fmt.Println(v11, v11, v10)
        _ = fi27
    }
    var v28 float32 = (1.5 + float32(0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v8
    _ = v9
}
