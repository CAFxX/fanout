package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int16; F1 int32; F2 uint64 }
func genFn0(p0 int8, p1 bool) int8 {
    v7 := GenS1{F0: 32767, F1: (("" + ("a" + "")) + ("hello" + ("hello" + "\n\t")))}
    v8 := []int{128, 238, 42, 278}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v9 := GenS2{F0: (-(p0))}
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 uint16 = 803
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 int16 = 32767
    v2 := []int{-9223372036854775808}
    v3 := make(map[string]int)
    var v4 uint8 = 128
    var v5 int32 = 221
    v6 := make(map[string]int)
    fmt.Println(v5, v5)
    if (!false) {
        v13 := GenS2{F0: 76}
        for fi14 := 0; fi14 < 7; fi14++ {
            if true {
                v15 := GenS3{F0: fi14, F1: -712}
                v16 := GenS2{F0: -1}
                v17 := GenS4{F0: int16(v5), F1: v5, F2: 18446744073709551615}
                _ = v15
                _ = v16
                _ = v17
            }
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 uint = 42
                var v20 int32 = (v5 ^ 790)
                var v21 int32 = v5
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            var v22 uint = 42
            _ = fi14
            _ = v22
        }
        v1 = (v1 ^ 0)
        _ = v13
    }
    var v23 int32 = (v5 & v5)
    var v24 int16 = (v1 * -156)
    delete(v3, ("hello" + "foo bar"))
    var v25 uint64 = 0
    var v26 bool = true
    var v27 float64 = (1.5 / ((1.5 + -2.25) - -2.25))
    var v28 uint16 = 878
    v6[("hello" + "\n\t")] = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
