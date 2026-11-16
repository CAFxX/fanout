package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint8; F2 string }
func genFn0(p0 uint32) int {
    p0 -= p0
    for fi5 := 0; fi5 < 5; fi5++ {
        fmt.Println(fi5, p0)
        v6 := GenS1{F0: 4294967295, F1: 9223372036854775807}
        _ = fi5
        _ = v6
    }
    _ = p0
    return 664
}

func main() {
    var v1 uint = 4294967295
    v2 := []string{"foo bar", "\n\t", ("a" + "\n\t"), "hello"}
    var v3 int32 = 732
    v4 := []string{}
    v1 = (v1 << ((v1) & 63))
    v7 := GenS2{F0: -1, F1: 2}
    _ = len(v4)
    if (-2147483648 < (-(v3))) {
        var v8 float32 = 0.1
        _ = v8
    } else {
        v9 := GenS3{F0: 2147483648, F1: 2, F2: "foo bar"}
        v10 := GenS2{F0: -501, F1: 2}
        _ = v10
        _ = v9
    }
    var v11 int = 563
    for fi12 := 0; fi12 < 1; fi12++ {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        _ = len(v2)
        fmt.Println(fi12, v3)
        _ = fi12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v1 = 65535
        _ = fi13
    }
    fmt.Println(v1)
    fmt.Println(v1, v11)
    v3 = v3
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
