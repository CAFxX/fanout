package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int }
func genFn0(p0 int16, p1 int8, p2 int16) uint32 {
    var v7 int16 = -1
    var v8 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{}
    var v3 string = ((((((("hello" + ("foo bar" + "a")) + ("\n\t" + ((("" + "") + "hello") + "a"))) + "\n\t") + ("" + "foo bar")) + "a") + ("" + "")) + ("" + "foo bar"))
    var v4 uint32 = 149
    var v5 int = 960
    v6 := make(map[uint32]bool)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint16 = 255
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v2 = append(v2, fi11)
        var v12 uint16 = 32767
        _ = fi11
        _ = v12
    }
    v3 = v3
    v3 = v3
    v13 := GenS1{F0: 64, F1: 255, F2: 128}
    v14 := GenS1{F0: 42, F1: 42, F2: (v5 ^ 4294967295)}
    if false {
        var v15 uint16 = 289
        fmt.Println(v3)
        v16 := GenS1{F0: 9223372036854775807, F1: 2, F2: 32768}
        _ = v15
        _ = v16
    }
    if (-32768 < 84) {
        if (4294967295 >= 42) {
            var v17 uint16 = 618
            _ = v17
        } else {
            v3 = "a"
            var v18 int8 = 127
            if (-9223372036854775808 > 42) {
                var v19 float64 = 3.14159
                _ = v19
            }
            _ = v18
        }
    } else {
        v3 = (v3 + ("foo bar" + v3))
    }
    v20 := v6[205]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
