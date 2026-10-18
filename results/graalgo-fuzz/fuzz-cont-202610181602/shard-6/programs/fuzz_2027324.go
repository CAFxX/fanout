package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint32 }
func genFn0(p0 int, p1 string) uint64 {
    fmt.Println(p0)
    p0 = (^(p0))
    v4 := GenS1{F0: 0, F1: p1, F2: 200}
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    var v1 int = 324
    var v2 uint8 = 0
    v3 := make(map[uint32]int)
    v5 := []int{128, (^(v1)), -448}
    if false {
        fmt.Println(v1)
    } else {
        var v6 int16 = -1
        v5 = append(v5, v1)
        var v7 string = "foo bar"
        _ = v6
        _ = v7
    }
    v1 -= 159
    var v8 uint64 = 1
    v2 -= (v2 << ((622) & 63))
    if false {
        if false {
            var v9 int32 = -866
            for fi10 := 0; fi10 < 1; fi10++ {
                switch "" {
                case "foo bar":
                    if true {
                        var v11 bool = true
                        _ = v11
                    } else {
                        for fi12 := 0; fi12 < 17; fi12++ {
                            var v13 float64 = 0.1
                            var v14 uint = 65535
                            delete(v3, 1)
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                    }
                case "hello":
                    v2 -= v2
                }
                for fi15 := 0; fi15 < 18; fi15++ {
                    if ((false || false) && false) {
                        v8 = 324
                        v3[669] = fi15
                    } else {
                        fmt.Println(v8)
                        var v16 int8 = 42
                        _ = v16
                    }
                    _ = fi15
                }
                v8 = (v8 << ((4294967295) & 63))
                _ = fi10
            }
            var v17 float64 = 1e10
            _ = v17
            _ = v9
        } else {
            var v18 string = ""
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        switch "hello" {
        case "":
            var v20 uint8 = 204
            _ = v20
        case "a":
            if true {
                delete(v3, 837)
                fmt.Println(fi19)
            }
        case "\n\t":
            if (!true) {
                var v21 int64 = 794
                var v22 int32 = 1
                v23 := GenS1{F0: (v22 << ((4294967295) & 63)), F1: "hello", F2: 4294967295}
                _ = v21
                _ = v22
                _ = v23
            }
        }
        _ = fi19
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        switch "" {
        case "a":
            v8 = v8
        }
        var v25 float64 = float64(1e10)
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v8
}
