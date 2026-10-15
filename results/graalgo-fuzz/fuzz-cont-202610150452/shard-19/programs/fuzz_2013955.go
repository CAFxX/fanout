package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
func genFn0(p0 int64, p1 uint) int {
    var v4 uint32 = 899
    if ((p0 + (p0 ^ p0)) <= 9223372036854775807) {
        v5 := []int{32767, 32767}
        _ = v5
    }
    switch "" {
    case "":
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 uint = p1
            _ = fi6
            _ = v7
        }
    }
    v8 := []int{-303, 128, 2147483647, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return -2147483648
}

func main() {
    var v1 uint8 = 176
    var v2 float32 = ((-2.25 - (float32(1e10) - 0.0)) * 1.5)
    v3 := make(map[int]uint32)
    v9 := GenS1{F0: 0}
    var v10 uint = 65535
    var v11 uint16 = 912
    var v12 uint64 = 4294967295
    if false {
        _ = len(v3)
        v12 = uint64(v11)
        v11 += v11
    } else {
        if (v12 < v12) {
            v13 := GenS1{F0: 914}
            var v14 bool = true
            switch genFn0(int64(v12), 226) {
            case 32767:
                v15 := GenS2{F0: v14}
                _ = v15
            case -161:
                fmt.Println(v12)
            }
            _ = v13
            _ = v14
        }
        v16 := v3[32768]
        _ = v16
    }
    delete(v3, 32768)
    for fi17 := 0; fi17 < 10; fi17++ {
        if (1000 == -297) {
            if (42 == 750) {
                v3[genFn0(298, v10)] = 41
            }
            v11 += uint16(v10)
        } else {
            v18 := []int{genFn0(1, (^(v10)))}
            v19 := GenS1{F0: 42}
            _ = v18
            _ = v19
        }
        v1 -= uint8(v11)
        v10 -= (v10 & 4294967295)
        _ = fi17
    }
    v20 := []int{1}
    v10 += 65535
    var v21 uint64 = uint64(v11)
    v22 := GenS1{F0: 2147483648}
    v11 += v11
    v23 := GenS2{F0: false}
    _ = len(v3)
    v21 -= v12
    for fi24 := 0; fi24 < 9; fi24++ {
        v21 = 1
        _ = fi24
    }
    var v25 uint64 = (^(v21))
    switch (("foo bar" + ("foo bar" + (("a" + ("foo bar" + (("hello" + ("foo bar" + "foo bar")) + "foo bar"))) + "a"))) + ("\n\t" + ("" + ""))) {
    case "\n\t":
        fmt.Println(v25, v21, v25)
    default:
        for fi26 := 0; fi26 < 20; fi26++ {
            for fi27 := 0; fi27 < 19; fi27++ {
                fmt.Println(v10)
                var v28 int64 = 1
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v9
}
