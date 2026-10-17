package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 string }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 1e10
    var v3 uint64 = 893
    var v4 string = (("hello" + (("" + "foo bar") + "hello")) + "")
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(v4, v2)
        for fi6 := 0; fi6 < 5; fi6++ {
            if true {
                v2 -= 0.0
                v1[-1] = (778 < 688)
            }
            _ = fi6
        }
        _ = fi5
    }
    v7 := GenS1{F0: 1}
    v1[0] = (2147483647 <= 1)
    v2 = 0.1
    v4 = "a"
    var v8 int = 9223372036854775807
    v9 := v1[(v8 % (((v8 << ((4294967295) & 63))) | 1))]
    v9 = v9
    switch v8 {
    case -1:
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: -78}
            fmt.Println(fi10, fi10, v2)
            _ = fi10
            _ = v11
        }
    case 4294967295:
        var v12 uint16 = 759
        _ = v12
    }
    v13 := GenS2{F0: (v3 << ((749) & 63)), F1: 18446744073709551615, F2: (v4 + ((v4 + v4) + (((v4 + ("hello" + ("" + (v4 + v4)))) + v4) + "hello")))}
    var v14 uint32 = 65535
    v15 := GenS2{F0: (v3 & 4294967295), F1: 435, F2: v4}
    fmt.Println(v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
