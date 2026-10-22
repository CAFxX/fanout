package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int32; F2 int64 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint64) int32 {
    p0 = (p0 * p0)
    if (!false) {
        var v4 uint = 42
        _ = v4
    }
    if true {
        p0 += (p0 / ((p0) | 1))
    } else {
        if (-1 != 1) {
            v5 := []int{}
            p0 = (p0 / ((p0) | 1))
            _ = v5
        } else {
            var v6 string = "foo bar"
            v7 := GenS1{F0: 1}
            _ = v6
            _ = v7
        }
        var v8 uint16 = 42
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    return -459
}

func main() {
    var v1 uint8 = 255
    var v2 uint64 = 179
    var v3 uint16 = 65535
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS2{F0: -1, F1: 1, F2: int64(v2)}
        _ = fi9
        _ = v10
    }
    var v11 int32 = 1
    if (-32768 >= -211) {
        if (!("hello" >= "a")) {
            var v12 uint16 = 255
            v13 := []int{-9223372036854775808}
            _ = v12
            _ = v13
        }
        v2 = 4294967295
        var v14 bool = false
        _ = v14
    } else {
        if true {
            fmt.Println(v3, v11, v2)
            for fi15 := 0; fi15 < 19; fi15++ {
                if (1 == 577) {
                    var v16 int64 = 797
                    _ = v16
                }
                v3 = 65535
                _ = fi15
            }
            v2 += v2
        } else {
            if true {
                v1 = 0
                for fi17 := 0; fi17 < 5; fi17++ {
                    v18 := GenS3{F0: 1}
                    _ = fi17
                    _ = v18
                }
                v2 = 4294967295
            }
            for fi19 := 0; fi19 < 17; fi19++ {
                for fi20 := 0; fi20 < 19; fi20++ {
                    switch "hello" {
                    case "hello":
                        var v21 int64 = 1
                        _ = v21
                    case "\n\t":
                        v2 = (v2 & 4294967295)
                    }
                    v11 = v11
                    _ = fi20
                }
                _ = fi19
            }
        }
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 int8 = -128
            _ = fi22
            _ = v23
        }
    }
    var v24 uint64 = v2
    var v25 float32 = 1e10
    fmt.Println(v24)
    v26 := []int{32767, 9223372036854775807}
    if true {
        if len(v26) > 0 {
            v26[0] = 128
        }
        if false {
            v25 -= (float32(3.14159) / v25)
            if (!true) {
                fmt.Println(v1, v24, v25)
                v1 += (v1 ^ v1)
                for fi27 := 0; fi27 < 12; fi27++ {
                    var v28 uint32 = 0
                    switch (fi27 & fi27) {
                    case -9223372036854775808:
                        v29 := GenS3{F0: (v11 / ((v11) | 1))}
                        _ = v29
                    case 4294967295:
                        v30 := GenS2{F0: -9223372036854775808, F1: -281, F2: -553}
                        _ = v30
                    case 2147483647:
                        fmt.Println(v11, v28, v1)
                    }
                    var v31 int32 = v11
                    _ = fi27
                    _ = v28
                    _ = v31
                }
            }
        } else {
            var v32 uint64 = 42
            var v33 int = 32767
            _ = v32
            _ = v33
        }
    } else {
        v34 := GenS1{F0: (v11 | (-(v11)))}
        _ = v34
    }
    var v35 float32 = v25
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v35
}
