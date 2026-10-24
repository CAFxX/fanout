package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
func genFn0() int64 {
    var v4 uint64 = 372
    fmt.Println(v4)
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := GenS1{F0: "a", F1: 860, F2: 2}
        if true {
            v7 := GenS1{F0: "hello", F1: 65535, F2: 849}
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 int32 = 823
                v10 := GenS2{F0: 32767}
                fmt.Println(v4)
                _ = fi8
                _ = v10
                _ = v9
            }
            v11 := []int{(^(fi5)), 9223372036854775807, 42, (fi5 & fi5)}
            _ = v11
            _ = v7
        } else {
            v4 += (^(v4))
            var v12 int64 = -1
            _ = v12
        }
        var v13 bool = false
        _ = fi5
        _ = v13
        _ = v6
    }
    _ = v4
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 2
    var v3 int64 = 9223372036854775807
    v14 := GenS1{F0: ("a" + ("foo bar" + "foo bar")), F1: 794, F2: 65535}
    v3 += (v3 + (v3 * (-(v3))))
    if false {
        v15 := GenS2{F0: -415}
        _ = v15
    } else {
        var v16 int8 = -1
        if true {
            v16 -= (v16 + v16)
        }
        _ = v16
    }
    v3 -= (v3 / (((v3 - v3)) | 1))
    v17 := GenS2{F0: -283}
    var v18 int = 127
    v3 -= (-(v3))
    var v19 int8 = 42
    v20 := GenS3{F0: 42}
    v21 := GenS2{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
