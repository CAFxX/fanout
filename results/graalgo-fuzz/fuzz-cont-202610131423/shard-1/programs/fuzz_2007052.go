package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }

func main() {
    var v1 uint32 = 2
    var v2 float64 = 1e10
    var v3 int64 = 275
    var v4 int16 = -32768
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            if false {
                var v7 int = -2147483648
                _ = v7
            } else {
                switch "a" {
                case "\n\t":
                    if true {
                        var v8 int32 = 2147483647
                        var v9 int = (^(fi5))
                        _ = v8
                        _ = v9
                    }
                case "":
                    if (-31 == 0) {
                        var v10 uint = 18446744073709551615
                        v11 := GenS1{F0: 842, F1: -429}
                        _ = v10
                        _ = v11
                    }
                default:
                    fmt.Println(v1, fi6)
                }
                if (-9223372036854775808 > (v3 << ((26) & 63))) {
                    v4 -= v4
                    fmt.Println(fi5, v4)
                    v12 := []int{(fi6 & (fi5 + (fi5 << ((255) & 63)))), -408}
                    _ = v12
                } else {
                    var v13 int = (-(fi6))
                    _ = v13
                }
            }
            if false {
                v14 := GenS1{F0: 2, F1: 69}
                var v15 int = fi6
                v16 := []int{(-(v15))}
                var v17 uint64 = 854
                switch ((("" + "foo bar") + "\n\t") + "foo bar") {
                case "":
                    v3 = (v3 & v3)
                }
                _ = v14
                _ = v15
                _ = v16
                _ = v17
            } else {
                if false {
                    var v18 int64 = v3
                    var v19 int16 = v4
                    v1 -= v1
                    v1 -= v1
                    _ = v18
                    _ = v19
                }
                v4 = 32767
                var v20 int = 9223372036854775807
                _ = v20
            }
            if (2147483647 == 2147483647) {
                var v21 bool = false
                if false {
                    v22 := GenS1{F0: 232, F1: 1}
                    fmt.Println(v21)
                    _ = v22
                } else {
                    var v23 uint = 768
                    _ = v23
                }
                _ = v21
            } else {
                v24 := GenS1{F0: 4294967295, F1: -512}
                if (!true) {
                    v25 := []int{-529, (-(fi6)), (^(fi6)), (^(fi5))}
                    var v26 string = ""
                    _ = v25
                    _ = v26
                }
                _ = v24
            }
            _ = fi6
        }
        v3 = (v3 % ((v3) | 1))
        _ = fi5
    }
    if true {
        v27 := GenS1{F0: 4294967295, F1: 1}
        var v28 int32 = 942
        _ = v27
        _ = v28
    }
    v3 = v3
    if false {
        if true {
            v2 = v2
            for fi29 := 0; fi29 < 13; fi29++ {
                var v30 uint16 = 603
                _ = fi29
                _ = v30
            }
        }
        var v31 int32 = -1
        _ = v31
    }
    var v32 uint32 = (v1 | 42)
    var v33 float32 = (1e10 + 0.1)
    v34 := GenS1{F0: 42, F1: 0}
    var v35 float64 = v2
    v36 := []int{}
    v37 := GenS2{F0: 65535, F1: (v32 + v32), F2: uint(v32)}
    switch "" {
    case "foo bar":
        for fi38 := 0; fi38 < 4; fi38++ {
            for fi39 := 0; fi39 < 14; fi39++ {
                fmt.Println(v2, v3, v3)
                v35 += v2
                var v40 uint = 1
                _ = fi39
                _ = v40
            }
            _ = fi38
        }
    case "":
        _ = len(v36)
    }
    v41 := GenS2{F0: 514, F1: 855, F2: 1}
    for fi42 := 0; fi42 < 19; fi42++ {
        if len(v36) > 0 {
            v36[0] = (fi42 | fi42)
        }
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
}
