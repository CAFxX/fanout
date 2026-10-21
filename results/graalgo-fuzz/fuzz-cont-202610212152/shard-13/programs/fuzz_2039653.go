package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint16 }
func genFn0(p0 int32) int8 {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return -128
}

func main() {
    var v1 int8 = 0
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    var v5 bool = true
    var v6 int8 = v1
    if (916 > -793) {
        var v7 bool = ((v5 && ("" != "hello")) && (1 <= 2))
        fmt.Println(v7, v2)
        var v8 uint64 = 42
        _ = v7
        _ = v8
    }
    v1 -= v6
    for fi9 := 0; fi9 < 12; fi9++ {
        fmt.Println(fi9, v1, v5)
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint8 = 113
            v2 = -1
            v2 += v2
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            if ("foo bar" >= "foo bar") {
                v13 := GenS1{F0: 42}
                v14 := GenS1{F0: 28}
                _ = v13
                _ = v14
            } else {
                switch "hello" {
                case "":
                    v15 := []int{32768, fi12, (fi9 & fi12)}
                    _ = v15
                case "hello":
                    v16 := GenS2{F0: 1, F1: 436}
                    _ = v16
                case "\n\t":
                    if (v5 && (v5 && (!v5))) {
                        v17 := v3["foo bar"]
                        _ = v17
                    } else {
                        v3["a"] = fi12
                        if (980 != 42) {
                            var v18 int = (fi12 << ((4294967295) & 63))
                            var v19 uint = 404
                            _ = v18
                            _ = v19
                        }
                        if v5 {
                            if false {
                                v6 -= -128
                                delete(v4, 4294967295)
                                delete(v4, 950)
                            }
                            v6 += 127
                            v20 := GenS2{F0: 96, F1: 566}
                            _ = v20
                        } else {
                            v21 := []int{fi9}
                            _ = v21
                        }
                    }
                }
                fmt.Println(v1, fi12, v5)
                switch "a" {
                case "\n\t":
                    fmt.Println(fi12, v5, fi9)
                case "foo bar":
                    var v22 uint32 = 2147483647
                    _ = v22
                case "hello":
                    if ((!false) && true) {
                        v5 = false
                        v23 := GenS1{F0: 128}
                        var v24 int = 92
                        _ = v23
                        _ = v24
                    }
                }
            }
            _ = fi12
        }
        _ = fi9
    }
    v25 := GenS1{F0: 128}
    v1 -= -128
    v26 := GenS1{F0: 239}
    v27 := GenS1{F0: 0}
    v1 = (v1 & (v1 / ((v6) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
