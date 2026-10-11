package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 bool, p1 int8) int32 {
    p1 = p1
    v6 := []int{526, 2147483647, -646}
    _ = p0
    _ = p1
    _ = v6
    return -1
}

func main() {
    var v1 int64 = -630
    var v2 uint16 = 42
    v3 := []string{"\n\t"}
    v4 := []int{9223372036854775807, 4294967295, -190}
    var v5 uint = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 = append(v3, "foo bar")
    v7 := GenS1{F0: v5, F1: 0}
    var v8 bool = false
    var v9 string = "\n\t"
    if true {
        v1 += 0
        for fi10 := 0; fi10 < 15; fi10++ {
            v1 -= 0
            fmt.Println(v9)
            v4 = append(v4, fi10)
            _ = fi10
        }
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        if v8 {
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := GenS1{F0: (v5 - v5), F1: 0}
                _ = fi12
                _ = v13
            }
            fmt.Println(v9, v9)
        } else {
            if (!(!v8)) {
                v8 = ((v8 || v8) || (0 <= 604))
                switch v9 {
                case "hello":
                    var v14 int8 = -128
                    _ = v14
                }
                var v15 string = "hello"
                _ = v15
            } else {
                fmt.Println(v8, v5)
            }
            if len(v4) > 0 {
                v4[0] = -2147483648
            }
        }
        _ = len(v3)
        _ = fi11
    }
    switch 209 {
    case -810:
        _ = len(v3)
    }
    _ = len(v3)
    if (32767 <= 1000) {
        fmt.Println(v1, v8)
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        _ = len(v3)
        if len(v4) > 0 {
            v4[0] = -444
        }
        _ = len(v3)
        _ = fi16
    }
    v17 := GenS2{F0: 255}
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
