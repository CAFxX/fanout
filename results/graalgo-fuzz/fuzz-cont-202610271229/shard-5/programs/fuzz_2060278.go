package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int }

func main() {
    var v1 uint32 = 100
    var v2 int64 = -407
    var v3 int16 = 1000
    if false {
        v4 := GenS1{F0: 65535, F1: "foo bar", F2: 1}
        _ = v4
    } else {
        var v5 string = ("\n\t" + "hello")
        for fi6 := 0; fi6 < 13; fi6++ {
            fmt.Println(fi6, v3)
            var v7 bool = false
            _ = fi6
            _ = v7
        }
        v8 := []int{-2147483648, -2147483648, -315, -1}
        _ = v5
        _ = v8
    }
    v9 := []int{-132}
    fmt.Println(v3, v3, v1)
    var v10 int8 = 76
    v3 += v3
    if true {
        if len(v9) > 0 {
            v9[0] = -1
        }
        var v11 uint32 = 65535
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 int = 883
            var v14 float32 = -2.25
            for fi15 := 0; fi15 < 13; fi15++ {
                switch -9223372036854775808 {
                case -2147483648:
                    _ = len(v9)
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v16 uint8 = 2
                    _ = v16
                case -944:
                    v17 := GenS1{F0: 255, F1: ("a" + "\n\t"), F2: 9223372036854775807}
                    _ = v17
                }
                var v18 int32 = 541
                _ = fi15
                _ = v18
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    }
    if len(v9) > 0 {
        v9[0] = 32768
    }
    v9 = append(v9, int(v10))
    v10 = v10
    var v19 bool = ((-2.25 + float64(0.0)) <= 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
