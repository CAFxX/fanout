package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint32, p1 bool, p2 int8) uint8 {
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            switch "\n\t" {
            case "":
                var v9 uint32 = p0
                var v10 bool = (!p1)
                if (v10 || p1) {
                    p0 = 4294967295
                    var v11 float32 = (float32((3.14159 / 1e10)) * 1e10)
                    var v12 float32 = 3.14159
                    _ = v11
                    _ = v12
                }
                _ = v10
                _ = v9
            case "hello":
                if p1 {
                    v13 := GenS1{F0: 386, F1: p2}
                    _ = v13
                } else {
                    var v14 string = ""
                    var v15 int64 = 0
                    v16 := []int{(fi8 | 161), fi7, fi7, fi7}
                    _ = v14
                    _ = v15
                    _ = v16
                }
            }
            _ = fi8
        }
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS2{F0: p0, F1: 0}
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 8; fi19++ {
            fmt.Println(p1)
            fmt.Println(fi19, fi7)
            var v20 uint8 = 42
            _ = fi19
            _ = v20
        }
        _ = fi7
    }
    p0 = (p0 * 42)
    var v21 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 0
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 uint8 = (v1 / ((v1) | 1))
    v5 := make(map[string]string)
    var v6 uint64 = 1
    var v22 string = (((("hello" + "") + "foo bar") + "\n\t") + "")
    v23 := []int{98, -2147483648, 42, 42}
    fmt.Println(v1, v1)
    v4 += genFn0(2147483648, ((v6 >> ((65535) & 63)) > v6), 76)
    if true {
        v4 += 127
        v6 = (v6 - (v6 & (v6 >> ((521) & 63))))
        for fi24 := 0; fi24 < 6; fi24++ {
            fmt.Println(fi24)
            _ = fi24
        }
    } else {
        switch v22 {
        case "":
            v25 := GenS1{F0: 0, F1: 0}
            _ = v25
        case "a":
            v26 := GenS2{F0: 2147483647, F1: v6}
            _ = v26
        }
    }
    if true {
        if (-1 > 1000) {
            fmt.Println(v22)
            v4 += (v4 / ((genFn0(1, false, 0)) | 1))
        }
    }
    v1 -= (v1 ^ v1)
    v27 := GenS1{F0: 0, F1: 0}
    _ = len(v3)
    for fi28 := 0; fi28 < 18; fi28++ {
        v29 := GenS2{F0: 2, F1: v6}
        for fi30 := 0; fi30 < 6; fi30++ {
            if false {
                if len(v23) > 0 {
                    v23[0] = (^(fi28))
                }
                v2[0] = 42
                v1 -= 189
            }
            fmt.Println(fi30, v22)
            _ = fi30
        }
        v4 = v4
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
