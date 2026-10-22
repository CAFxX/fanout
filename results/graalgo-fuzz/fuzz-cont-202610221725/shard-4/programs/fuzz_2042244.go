package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 bool }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int16; F1 int64; F2 int }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int32, p1 int) uint {
    fmt.Println(p0, p1)
    var v4 string = "hello"
    p1 = p1
    _ = p0
    _ = p1
    _ = v4
    return 302
}

func main() {
    v1 := []int{-187, 1, -2147483648}
    v2 := make(map[string]bool)
    v3 := []int{1}
    var v5 int = 0
    v5 += (v5 >> ((255) & 63))
    for fi6 := 0; fi6 < 15; fi6++ {
        switch "\n\t" {
        case "\n\t":
            switch (("\n\t" + "foo bar") + "a") {
            case "hello":
                switch "hello" {
                case "\n\t":
                    switch (v5 % ((fi6) | 1)) {
                    case 32768:
                        v5 = (fi6 + 856)
                    default:
                        v7 := GenS1{F0: -152, F1: -787, F2: true}
                        _ = v7
                    }
                }
            case "a":
                if true {
                    var v8 uint = 4294967295
                    _ = v8
                } else {
                    var v9 int16 = -1
                    if true {
                        for fi10 := 0; fi10 < 4; fi10++ {
                            fmt.Println(v9)
                            v2["hello"] = false
                            v9 += v9
                            _ = fi10
                        }
                        v1 = append(v1, (v5 | -2147483648))
                        _ = len(v2)
                    } else {
                        for fi11 := 0; fi11 < 1; fi11++ {
                            v9 = (v9 / (((^(v9))) | 1))
                            _ = fi11
                        }
                        v9 -= v9
                        for fi12 := 0; fi12 < 18; fi12++ {
                            switch fi12 {
                            case 128:
                                v13 := GenS1{F0: 42, F1: (v5 >> ((1) & 63)), F2: true}
                                _ = v13
                            }
                            v14 := GenS2{F0: genFn0(-2147483648, (fi12 - v5)), F1: 2147483648}
                            if true {
                                v9 = (v9 & 1000)
                                v5 = v5
                                var v15 int64 = -1
                                v16 := GenS2{F0: 18446744073709551615, F1: 48}
                                v17 := GenS3{F0: (v9 & -335), F1: v15, F2: fi6}
                                _ = v15
                                _ = v16
                                _ = v17
                            } else {
                                fmt.Println(v9, v9)
                                var v18 bool = false
                                v1 = append(v1, fi6)
                                _ = v18
                            }
                            var v19 uint = 255
                            _ = fi12
                            _ = v14
                            _ = v19
                        }
                    }
                    v5 = (v5 % ((fi6) | 1))
                    _ = v9
                }
            }
        }
        if (!(!true)) {
            v5 += (v5 >> ((254) & 63))
            var v20 uint32 = 2147483647
            v5 += v5
            _ = v20
        }
        _ = fi6
    }
    v5 = v5
    v21 := GenS3{F0: 915, F1: -445, F2: (v5 >> ((65535) & 63))}
    delete(v2, "\n\t")
    var v22 string = "hello"
    var v23 int8 = -1
    var v24 bool = true
    v25 := GenS4{F0: 1}
    if true {
        v22 = "foo bar"
        var v26 float32 = (-2.25 + 0.1)
        _ = v26
    }
    _ = len(v3)
    v24 = true
    v22 = v22
    var v27 float32 = ((1.5 + -2.25) - ((-2.25 + (1e10 - float32(float64(float32(1e10))))) / 1.0))
    v3 = append(v3, (v5 ^ (-(v5))))
    v24 = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v5
}
