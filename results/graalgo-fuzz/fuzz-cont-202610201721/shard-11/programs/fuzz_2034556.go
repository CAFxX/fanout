package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint64; F1 uint; F2 uint32 }
func genFn0() int64 {
    v4 := GenS1{F0: 0}
    var v5 int16 = 1000
    _ = v4
    _ = v5
    return 970
}
func genFn1(p0 int) uint32 {
    v6 := []int{1, -9223372036854775808, (p0 + p0)}
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            p0 = -1
            if (651 > 42) {
                switch "hello" {
                case "":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case "foo bar":
                    if len(v6) > 0 {
                        v6[0] = -1
                    }
                case "hello":
                    fmt.Println(fi7, fi8)
                }
                switch (("\n\t" + "") + "hello") {
                case "hello":
                    v9 := GenS1{F0: 18446744073709551615}
                    _ = v9
                case "a":
                    v10 := GenS1{F0: 65535}
                    _ = v10
                }
                _ = len(v6)
                _ = len(v6)
                var v11 int = (-(p0))
                _ = v11
            }
            _ = fi8
        }
        _ = fi7
    }
    var v12 uint32 = 2
    _ = p0
    _ = v12
    _ = v6
    return v12
}

func main() {
    var v1 int8 = -1
    var v2 uint = 65535
    v3 := make(map[string]bool)
    v13 := GenS2{F0: 342}
    v14 := GenS2{F0: 2}
    for fi15 := 0; fi15 < 3; fi15++ {
        v2 = v2
        var v16 bool = (fi15 > (^(fi15)))
        _ = fi15
        _ = v16
    }
    var v17 int64 = int64(v1)
    v1 = 42
    var v18 float64 = 0.0
    v19 := GenS3{F0: (^(v17)), F1: false}
    var v20 uint16 = 0
    v18 = v18
    v20 -= v20
    v21 := v3["foo bar"]
    v21 = ((!(0.1 > 0.0)) || (!v21))
    switch -1 {
    case 486:
        fmt.Println(v2, v20)
    case 32767:
        for fi22 := 0; fi22 < 6; fi22++ {
            v23 := GenS2{F0: 2147483648}
            _ = fi22
            _ = v23
        }
    case 0:
        v3["hello"] = v21
    }
    v24 := GenS4{F0: 18446744073709551615, F1: v2, F2: 296}
    v25 := []int{9223372036854775807}
    v17 -= genFn0()
    var v26 float32 = (1.5 - (-2.25 + 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
