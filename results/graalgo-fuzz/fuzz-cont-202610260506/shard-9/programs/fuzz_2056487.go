package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int }
type GenS2 struct { F0 int32 }

func main() {
    var v1 string = ("\n\t" + ("foo bar" + ("hello" + ("" + "hello"))))
    v2 := make(map[uint32]string)
    var v3 string = v1
    v4 := make(map[uint32]bool)
    var v5 uint64 = 2
    v6 := GenS1{F0: 0, F1: v1, F2: 128}
    v7 := GenS2{F0: -390}
    if (1 > -1) {
        if (-60 >= 1) {
            v3 = ("\n\t" + ((v3 + "a") + (((("hello" + "foo bar") + (v3 + "foo bar")) + "") + v3)))
        }
    } else {
        v3 = v3
    }
    v8 := GenS1{F0: -937, F1: v3, F2: 128}
    fmt.Println(v5, v5)
    v9 := []int{42, -1, 127}
    v10 := GenS2{F0: 0}
    switch v3 {
    case "":
        v9 = append(v9, 128)
    case "foo bar":
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    if (!(!false)) {
        for fi11 := 0; fi11 < 11; fi11++ {
            fmt.Println(v5, v1)
            if (-9223372036854775808 == 669) {
                if len(v9) > 0 {
                    v9[0] = fi11
                }
            }
            var v12 uint = 110
            _ = fi11
            _ = v12
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
