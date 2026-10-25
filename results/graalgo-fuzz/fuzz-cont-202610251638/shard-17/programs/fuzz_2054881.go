package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint32 }
type GenS2 struct { F0 uint8; F1 string }
func genFn0(p0 int32, p1 int32) string {
    v5 := []int{4294967295, -2147483648, 1}
    v6 := GenS1{F0: (p1 % (((p0 + (p0 >> ((180) & 63)))) | 1)), F1: -2147483648, F2: 1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[string]int)
    v3 := []uint32{4294967295, 65}
    var v4 string = "foo bar"
    var v7 int = 471
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int8 = 114
        v1 = v1
        v7 += fi8
        _ = fi8
        _ = v9
    }
    v4 = (("foo bar" + ("hello" + "a")) + "")
    for fi10 := 0; fi10 < 8; fi10++ {
        switch (fi10 >> ((255) & 63)) {
        case -691:
            for fi11 := 0; fi11 < 10; fi11++ {
                if (-1 != 1) {
                    if true {
                        if false {
                            if true {
                                if (!false) {
                                    v7 += fi11
                                }
                            } else {
                                fmt.Println(v7, v7)
                                fmt.Println(fi11, fi10, v7)
                                var v12 uint16 = 65535
                                _ = v12
                            }
                        } else {
                            delete(v2, "foo bar")
                            _ = len(v3)
                            v13 := GenS1{F0: 1, F1: fi11, F2: 79}
                            _ = v13
                        }
                        switch (v7 & -247) {
                        case -724:
                            if true {
                                switch v4 {
                                case "":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                v4 = v4
                                var v14 float32 = 0.0
                                var v15 float64 = v1
                                _ = v14
                                _ = v15
                            } else {
                                v16 := v2[v4]
                                v4 = "\n\t"
                                _ = v16
                            }
                        }
                        var v17 uint64 = 4294967295
                        _ = v17
                    } else {
                        v18 := GenS1{F0: -833, F1: (fi11 | fi11), F2: 4294967295}
                        v1 -= v1
                        v1 = v1
                        _ = v18
                    }
                    var v19 int8 = 12
                    v1 += (0.0 + (v1 * v1))
                    _ = v19
                }
                var v20 float64 = 1e10
                var v21 bool = (0 > 2147483648)
                _ = fi11
                _ = v20
                _ = v21
            }
        }
        _ = fi10
    }
    v7 -= v7
    var v22 uint8 = 42
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 float32 = 1.5
        var v25 uint32 = 332
        v1 = 1.5
        _ = fi23
        _ = v24
        _ = v25
    }
    delete(v2, v4)
    var v26 uint64 = 760
    switch v4 {
    case "\n\t":
        v1 = v1
    }
    v27 := GenS2{F0: (v22 * (v22 * (v22 % ((127) | 1)))), F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v7
}
