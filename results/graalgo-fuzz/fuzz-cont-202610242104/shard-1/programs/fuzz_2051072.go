package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int32 }
func genFn0(p0 uint16, p1 int16) uint32 {
    p1 += p1
    switch -704 {
    case 128:
        p0 += p0
    }
    v7 := []int{128, 0, -104}
    _ = p0
    _ = p1
    _ = v7
    return 598
}

func main() {
    var v1 bool = (!true)
    var v2 int8 = -4
    v3 := make(map[string]string)
    var v4 bool = true
    var v5 uint8 = 42
    v6 := make(map[string]string)
    var v8 int64 = 1
    fmt.Println(v1, v4, v5)
    if (!v1) {
        v2 = 127
        v9 := GenS1{F0: v2, F1: 0, F2: -802}
        _ = v9
    }
    if (!v1) {
        v8 += v8
        v10 := []int{2147483647, -9223372036854775808, -9223372036854775808, 502}
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v6[(("a" + ("a" + "\n\t")) + "\n\t")] = ""
        v1 = v1
        _ = fi11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint = 65535
        v2 -= (v2 % (((v2 | -97)) | 1))
        _ = fi12
        _ = v13
    }
    _ = len(v6)
    var v14 int = 1
    var v15 bool = v4
    var v16 uint16 = 255
    if true {
        v17 := []int{32768}
        v5 = v5
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
