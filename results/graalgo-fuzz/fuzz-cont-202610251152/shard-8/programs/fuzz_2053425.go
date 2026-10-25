package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 uint; F2 int64 }
func genFn0(p0 uint8, p1 int, p2 string) int16 {
    var v5 uint8 = 0
    var v6 int32 = 42
    fmt.Println(p1, v6, p1)
    var v7 string = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 1000
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]int)
    var v3 float64 = -2.25
    var v4 bool = false
    v8 := []int{42, 127, 1, 9223372036854775807}
    if ((-(v1)) == v1) {
        v4 = (true || v4)
        v4 = (0 < 65535)
    } else {
        if false {
            v1 -= v1
            fmt.Println(v4)
        }
        if true {
            if false {
                v2[0] = 32767
            } else {
                v3 -= 0.1
                for fi9 := 0; fi9 < 10; fi9++ {
                    v4 = (v4 || true)
                    fmt.Println(fi9)
                    for fi10 := 0; fi10 < 16; fi10++ {
                        v11 := GenS1{F0: int8(v1)}
                        v12 := GenS1{F0: 100}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = fi9
                }
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            if (2 <= 953) {
                fmt.Println(v3, v1)
                v4 = v4
                v3 += 0.1
            }
        } else {
            var v13 int32 = -1
            v1 = v1
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 int64 = (v1 | v1)
            var v17 uint = 18446744073709551615
            fmt.Println(fi15, v17, v3)
            _ = len(v8)
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    v18 := GenS1{F0: 0}
    for fi19 := 0; fi19 < 19; fi19++ {
        if true {
            switch "hello" {
            case "":
                if (!v4) {
                    _ = len(v8)
                    v1 += 259
                }
            }
        }
        switch "foo bar" {
        case "hello":
            v3 = v3
        case "a":
            fmt.Println(v4, fi19, v4)
        }
        _ = fi19
    }
    var v20 int64 = -600
    v21 := GenS1{F0: 1}
    var v22 int8 = 42
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := GenS1{F0: 1}
        var v25 int8 = v22
        v26 := GenS2{F0: 901, F1: 649, F2: -972}
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    v1 = v20
    var v27 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v8
}
