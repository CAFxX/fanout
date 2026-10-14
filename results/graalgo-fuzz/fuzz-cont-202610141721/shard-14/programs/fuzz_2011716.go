package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint32; F2 string }
func genFn0() bool {
    var v7 uint = 1
    var v8 uint8 = 255
    fmt.Println(v8, v7)
    var v9 uint32 = 0
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]string)
    var v3 uint16 = 42
    var v4 int64 = 1
    v5 := []int{}
    var v6 uint64 = 2
    v6 -= (v6 << ((131) & 63))
    if ((!true) || true) {
        if true {
            v10 := v2[0]
            _ = v10
        } else {
            v11 := GenS1{F0: -1}
            for fi12 := 0; fi12 < 3; fi12++ {
                for fi13 := 0; fi13 < 2; fi13++ {
                    v3 -= v3
                    v14 := GenS1{F0: -1}
                    if (!false) {
                        if true {
                            _ = len(v5)
                        } else {
                            v15 := GenS2{F0: 32767}
                            _ = v15
                        }
                        fmt.Println(v3, v1, fi12)
                        if false {
                            v6 = (v6 + v6)
                            fmt.Println(fi12, v3, fi13)
                            fmt.Println(v6, v6)
                        }
                    } else {
                        _ = len(v5)
                        var v16 bool = (!false)
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                }
                if true {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        v18 := GenS2{F0: -824}
                        fmt.Println(fi17, v1, v4)
                        v19 := GenS3{F0: (!(1 == 888)), F1: uint32(v6), F2: "a"}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    fmt.Println(v3)
                } else {
                    v4 -= 9223372036854775807
                }
                _ = fi12
            }
            _ = v11
        }
    }
    v6 = v6
    v20 := GenS2{F0: 1000}
    fmt.Println(v6, v3)
    for fi21 := 0; fi21 < 19; fi21++ {
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 float32 = float32(v1)
            var v24 uint64 = (v6 - v6)
            _ = len(v2)
            v24 += (v6 / (((v6 << ((18446744073709551615) & 63))) | 1))
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 uint8 = 42
        var v26 uint64 = v6
        _ = fi21
        _ = v25
        _ = v26
    }
    if len(v5) > 0 {
        v5[0] = 0
    }
    v27 := GenS2{F0: 42}
    for fi28 := 0; fi28 < 1; fi28++ {
        v29 := v2[763]
        var v30 int64 = v4
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
