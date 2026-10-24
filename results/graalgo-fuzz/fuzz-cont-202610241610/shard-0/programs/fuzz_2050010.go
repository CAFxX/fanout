package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }
func genFn0(p0 string, p1 int64) uint16 {
    v4 := GenS1{F0: 255, F1: 127, F2: 978}
    p1 = (p1 >> ((0) & 63))
    var v5 int16 = -1
    if (true || (42 >= 42)) {
        v5 -= (v5 + (^(v5)))
        fmt.Println(p1, p0)
    }
    var v6 int32 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 233
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 2
    var v3 int8 = 100
    var v7 bool = (-93 >= -9223372036854775808)
    v3 += -128
    v8 := GenS2{F0: -9223372036854775808, F1: 42}
    v9 := GenS3{F0: "a"}
    if (18446744073709551615 >= v2) {
        fmt.Println(v2, v2)
    } else {
        v10 := v1["\n\t"]
        v11 := []int{}
        _ = v10
        _ = v11
    }
    var v12 string = ("" + (("\n\t" + "\n\t") + "hello"))
    v13 := GenS4{F0: 804}
    v2 += v2
    v14 := v1[(("hello" + ((((v12 + "hello") + v12) + v12) + v12)) + v12)]
    for fi15 := 0; fi15 < 3; fi15++ {
        if v7 {
            if false {
                fmt.Println(v14, v12, fi15)
                if (!true) {
                    var v16 int16 = 1000
                    _ = v16
                } else {
                    var v17 int16 = 32767
                    _ = v17
                }
            }
            fmt.Println(v3, fi15, v2)
            v3 -= v3
        } else {
            var v18 int16 = -568
            _ = v18
        }
        var v19 int32 = -2147483648
        v1["hello"] = (fi15 | 2147483647)
        _ = fi15
        _ = v19
    }
    var v20 float32 = -2.25
    var v21 int32 = 1
    v12 = v12
    for fi22 := 0; fi22 < 14; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            v21 -= 0
            _ = fi23
        }
        _ = fi22
    }
    if false {
        if v7 {
            if (true || (42 > 367)) {
                if (v14 != v14) {
                    var v24 uint = (v2 & (v2 << ((42) & 63)))
                    _ = v24
                }
            }
            if v7 {
                var v25 int32 = 42
                _ = v25
            }
        } else {
            v26 := GenS4{F0: 65535}
            _ = v26
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 float32 = -2.25
            _ = fi27
            _ = v28
        }
        v29 := GenS2{F0: -1, F1: v21}
        _ = v29
    } else {
        var v30 int8 = v3
        v31 := GenS2{F0: -1, F1: v21}
        _ = v30
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
