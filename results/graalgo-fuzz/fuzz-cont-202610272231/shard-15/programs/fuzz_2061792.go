package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int64 }
func genFn0() int32 {
    var v4 float64 = (1.5 - 0.1)
    v5 := []int{9223372036854775807}
    _ = v4
    _ = v5
    return -417
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = true
    var v3 uint16 = 32
    v6 := GenS1{F0: 775}
    if (877 < 255) {
        v7 := GenS1{F0: 2147483647}
        v8 := GenS1{F0: 32768}
        _ = v7
        _ = v8
    } else {
        var v9 bool = (42 == 0)
        _ = v9
    }
    fmt.Println(v3)
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := GenS2{F0: "a"}
            var v13 uint16 = v3
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    if v2 {
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS1{F0: (-(fi14))}
            _ = fi14
            _ = v15
        }
        v16 := GenS3{F0: 214, F1: -9223372036854775808}
        _ = v16
    } else {
        if ((uint8(v3) != 42) && (0 > 31)) {
            v17 := []int{1, -882, -9223372036854775808}
            v3 -= (v3 | (v3 - v3))
            if true {
                var v18 int8 = 4
                v19 := GenS1{F0: 42}
                v3 += (v3 & v3)
                _ = v18
                _ = v19
            }
            _ = v17
        } else {
            if true {
                v1[("hello" + "hello")] = "\n\t"
                v2 = (79 == -128)
            } else {
                v2 = v2
                for fi20 := 0; fi20 < 19; fi20++ {
                    for fi21 := 0; fi21 < 15; fi21++ {
                        fmt.Println(fi21, v2, v2)
                        v22 := []int{(^(fi20))}
                        _ = fi21
                        _ = v22
                    }
                    v3 -= (^(v3))
                    var v23 uint32 = 2147483647
                    _ = fi20
                    _ = v23
                }
            }
        }
        v2 = (-584 > 274)
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        if false {
            for fi25 := 0; fi25 < 12; fi25++ {
                var v26 uint32 = 4294967295
                _ = fi25
                _ = v26
            }
            var v27 uint16 = 65535
            var v28 int16 = 0
            _ = v27
            _ = v28
        }
        _ = fi24
    }
    var v29 uint64 = 1
    var v30 uint = 0
    v29 = uint64(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v6
}
