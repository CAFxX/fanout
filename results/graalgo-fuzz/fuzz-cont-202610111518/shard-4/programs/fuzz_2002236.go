package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
func genFn0(p0 string, p1 uint) int32 {
    var v7 string = (p0 + "\n\t")
    if false {
        if (-2.25 == (((0.0 - 1.5) * 0.1) - 0.1)) {
            v8 := GenS1{F0: 2, F1: 2, F2: 42}
            var v9 int32 = -1
            _ = v8
            _ = v9
        }
        var v10 int16 = -938
        fmt.Println(p0)
        _ = v10
    } else {
        if (343 <= 358) {
            var v11 uint = p1
            v12 := []int{641, 455, 127}
            v13 := GenS1{F0: 2, F1: 0, F2: 42}
            _ = v11
            _ = v12
            _ = v13
        }
        v14 := []int{}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v7
    return 2147483647
}

func main() {
    var v1 float32 = (-2.25 + ((float32((0.1 / (3.14159 / 0.1))) * (1e10 / 1e10)) - 1e10))
    var v2 uint8 = 176
    var v3 uint8 = (^(v2))
    var v4 uint16 = 1
    var v5 uint64 = 768
    var v6 int32 = 2147483647
    var v15 float32 = v1
    fmt.Println(v3)
    if (true || true) {
        v16 := []int{2147483647, 42, -1}
        if true {
            v16 = append(v16, 4294967295)
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
        if true {
            var v17 uint64 = 628
            v18 := GenS1{F0: v5, F1: (^(v17)), F2: 127}
            _ = v17
            _ = v18
        }
        _ = v16
    } else {
        var v19 int32 = -2147483648
        _ = v19
    }
    fmt.Println(v4, v15)
    v5 += v5
    v20 := GenS2{F0: v4}
    v21 := GenS3{F0: "a"}
    var v22 int64 = 9223372036854775807
    if false {
        v23 := GenS1{F0: 947, F1: v5, F2: 0}
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 bool = (!false)
            fmt.Println(v4, v25)
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 bool = v25
                _ = fi26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
