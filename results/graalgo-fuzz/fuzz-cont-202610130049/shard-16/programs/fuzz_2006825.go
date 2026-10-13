package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint16; F2 uint8 }
func genFn0(p0 int64) int16 {
    v7 := []int{9223372036854775807}
    v8 := GenS1{F0: "foo bar", F1: 42, F2: -32768}
    p0 += p0
    p0 = p0
    _ = p0
    _ = v7
    _ = v8
    return 362
}
func genFn1(p0 uint8) string {
    var v9 uint32 = 4294967295
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 14; fi11++ {
            if ("a" < "hello") {
                v9 = v9
            }
            var v12 int32 = 2147483647
            v13 := []int{fi10, fi10}
            v14 := GenS1{F0: ("hello" + (("\n\t" + "") + ("\n\t" + "hello"))), F1: 42, F2: 0}
            p0 += uint8(v9)
            fmt.Println(v12, v9, fi11)
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        var v15 uint = 42
        var v16 uint = (v15 - v15)
        _ = fi10
        _ = v15
        _ = v16
    }
    fmt.Println(v9)
    var v17 uint32 = (v9 - 250)
    for fi18 := 0; fi18 < 1; fi18++ {
        p0 = 42
        _ = fi18
    }
    _ = p0
    _ = v17
    _ = v9
    return "a"
}

func main() {
    var v1 bool = false
    var v2 uint32 = 31
    var v3 int8 = -48
    v4 := []uint32{v2, v2, (^(v2)), v2}
    v5 := make(map[string]bool)
    v6 := make(map[string]int)
    var v19 int32 = 1
    v1 = true
    v4 = append(v4, 2147483648)
    v20 := GenS2{F0: v19}
    v21 := GenS3{F0: "", F1: 419, F2: 199}
    v22 := v6["\n\t"]
    v1 = true
    v23 := GenS2{F0: (v19 * (-(v19)))}
    v24 := GenS3{F0: "foo bar", F1: 65535, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
