package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint8 = 255
    v2 := []int{9223372036854775807, 4294967295, -9223372036854775808}
    v3 := make(map[int]bool)
    v4 := []string{"a", "a", "hello", "hello"}
    var v5 float32 = 0.1
    v6 := GenS1{F0: -1, F1: 32767, F2: 65535}
    var v7 bool = true
    v8 := GenS2{F0: -496, F1: 444}
    switch -9223372036854775808 {
    case -9223372036854775808:
        if len(v2) > 0 {
            v2[0] = -1
        }
    case 32767:
        fmt.Println(v1, v5)
    }
    if (!true) {
        v7 = false
        v9 := GenS2{F0: 385, F1: 2147483648}
        _ = v9
    } else {
        v3[-497] = (2147483647 > -1)
        var v10 uint8 = v1
        _ = v10
    }
    fmt.Println(v5, v1)
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 bool = (0 < 1)
        v13 := GenS3{F0: 1}
        var v14 int16 = 0
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    if v7 {
        if v7 {
            v15 := GenS3{F0: -324}
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            fmt.Println(fi16, v5)
            for fi17 := 0; fi17 < 15; fi17++ {
                v3[fi16] = (-9223372036854775808 >= 9223372036854775807)
                _ = fi17
            }
            _ = fi16
        }
    }
    v4 = append(v4, "")
    for fi18 := 0; fi18 < 11; fi18++ {
        v2 = append(v2, (^(fi18)))
        v19 := v3[fi18]
        v5 -= 0.0
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 92, F1: 32767, F2: 255}
    for fi21 := 0; fi21 < 2; fi21++ {
        if ((!v7) && (!(9223372036854775807 != -1))) {
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
        } else {
            v3[(-(fi21))] = v7
        }
        _ = fi21
    }
    delete(v3, 9223372036854775807)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
