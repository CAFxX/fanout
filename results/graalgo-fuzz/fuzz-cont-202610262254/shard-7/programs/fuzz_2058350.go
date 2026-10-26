package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 42
    v3 := make(map[string]int)
    var v4 uint = 4294967295
    var v5 int32 = -305
    v6 := make(map[int]string)
    switch (("" + "\n\t") + "") {
    case "a":
        _ = len(v6)
    case "foo bar":
        fmt.Println(v2, v4)
    default:
        if (true && true) {
            var v7 bool = (true && false)
            v4 += v4
            v6[-296] = ("\n\t" + "hello")
            _ = v7
        }
    }
    var v8 uint32 = 2
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint = (v4 | (v4 & v4))
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: 1, F1: 65535}
        _ = fi9
        _ = v12
    }
    v13 := GenS1{F0: 32767, F1: 177}
    var v14 uint = v4
    fmt.Println(v8, v5, v4)
    var v15 uint8 = 2
    if false {
        if true {
            if true {
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := []int{2147483647, fi16, 555, (fi16 & fi16)}
                    v6[(fi16 - (fi16 * 4294967295))] = "foo bar"
                    switch "a" {
                    case "a":
                        var v18 float32 = -2.25
                        _ = v18
                    }
                    _ = fi16
                    _ = v17
                }
                v4 += v14
                v2 = int64(v4)
            }
            if false {
                v19 := GenS1{F0: uint16(v8), F1: 32767}
                v20 := GenS2{F0: 0}
                var v21 float32 = 0.1
                _ = v19
                _ = v20
                _ = v21
            }
        } else {
            v22 := []int{15, -519, -2147483648}
            _ = v22
        }
        v23 := []int{42, int(v4), 32767, 4294967295}
        if (255 != (v15 ^ v15)) {
            v2 = v2
        }
        _ = v23
    } else {
        v24 := GenS1{F0: 42, F1: 900}
        v25 := []int{-2147483648, -240}
        _ = v24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
