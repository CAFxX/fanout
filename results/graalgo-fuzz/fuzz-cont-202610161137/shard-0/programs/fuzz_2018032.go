package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint8, p1 int8) uint32 {
    p1 += p1
    var v7 uint64 = 42
    var v8 float32 = (0.1 - float32(0.1))
    v9 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 437
}
func genFn1() int32 {
    var v10 bool = true
    var v11 uint64 = 4294967295
    v11 += v11
    v12 := GenS1{F0: 804}
    _ = v10
    _ = v11
    _ = v12
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 int64 = 646
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    v6 := []string{"hello", ("a" + "hello")}
    v3 -= v3
    var v13 int64 = (v3 & 1)
    v14 := v5[""]
    fmt.Println(v13, v3, v13)
    v15 := GenS1{F0: 114}
    var v16 int32 = 1
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 uint64 = 2
            var v20 uint16 = 32767
            if (-128 <= 100) {
                var v21 int16 = -361
                _ = v21
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        v14 = v14
        _ = fi17
    }
    v3 -= -1
    v14 = v14
    if true {
        v22 := GenS1{F0: 2}
        var v23 int16 = 42
        _ = v22
        _ = v23
    }
    var v24 uint32 = 42
    fmt.Println(v14, v24, v24)
    fmt.Println(v14, v16)
    v14 = (("a" >= "a") || (127 > 147))
    fmt.Println(v13)
    var v25 uint32 = v24
    v26 := GenS1{F0: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
