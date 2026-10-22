package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint16 = 0
    v3 := make(map[uint32]string)
    v4 := []int{9223372036854775807, 500}
    var v5 bool = false
    v1 -= (v1 - v1)
    v5 = v5
    var v6 uint32 = 2147483647
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v7 string = "foo bar"
    var v8 uint = 42
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: 1}
        if ((v8 >> ((v8) & 63)) == v8) {
            v11 := GenS1{F0: 1}
            if v5 {
                if v5 {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        var v13 uint32 = v6
                        v5 = v5
                        _ = fi12
                        _ = v13
                    }
                    if ((v2 + uint16(v8)) == v2) {
                        for fi14 := 0; fi14 < 11; fi14++ {
                            v8 += v8
                            v7 = "foo bar"
                            _ = fi14
                        }
                    }
                    var v15 uint8 = 222
                    _ = v15
                } else {
                    _ = len(v4)
                    for fi16 := 0; fi16 < 6; fi16++ {
                        var v17 int = fi9
                        _ = fi16
                        _ = v17
                    }
                }
            } else {
                var v18 int16 = 42
                _ = v18
            }
            _ = v11
        } else {
            v19 := GenS1{F0: 4294967295}
            _ = v19
        }
        v20 := GenS1{F0: 42}
        _ = fi9
        _ = v10
        _ = v20
    }
    fmt.Println(v5, v5, v7)
    v7 = "a"
    var v21 uint32 = (v6 >> ((v8) & 63))
    var v22 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
