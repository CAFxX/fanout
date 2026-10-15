package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int8; F1 bool }

func main() {
    var v1 int32 = -1
    var v2 bool = (true && false)
    var v3 bool = (499 == 42)
    v4 := []int{}
    var v5 int8 = -59
    var v6 string = "a"
    var v7 float32 = 1.5
    fmt.Println(v7)
    var v8 uint = 42
    fmt.Println(v5)
    v1 = int32(v8)
    if false {
        _ = len(v4)
        v4 = append(v4, 1)
        var v9 float32 = ((v7 + ((v7 * (v7 / float32(1e10))) * (v7 / 3.14159))) + ((v7 + v7) * (1.5 + (((float32(0.0) / (v7 / v7)) * 0.0) * v7))))
        _ = v9
    }
    v10 := GenS1{F0: 195, F1: 538}
    v4 = append(v4, 2147483647)
    _ = len(v4)
    v11 := GenS2{F0: 0, F1: false}
    v6 = v6
    v4 = append(v4, 2147483647)
    switch v6 {
    case "":
        for fi12 := 0; fi12 < 2; fi12++ {
            if len(v4) > 0 {
                v4[0] = -9223372036854775808
            }
            if (-944 == -529) {
                v13 := GenS2{F0: (v5 / (((^(v5))) | 1)), F1: (-1 > -760)}
                _ = v13
            } else {
                if true {
                    if (false || (!((1 >= 64) && (2147483648 <= 810)))) {
                        _ = len(v4)
                        fmt.Println(v1, v7)
                        if len(v4) > 0 {
                            v4[0] = (fi12 * fi12)
                        }
                    } else {
                        if len(v4) > 0 {
                            v4[0] = (^(fi12))
                        }
                        switch (^(fi12)) {
                        case 4294967295:
                            v14 := GenS1{F0: 2, F1: -1}
                            _ = v14
                        case 903:
                            var v15 uint16 = 65535
                            _ = v15
                        case -239:
                            _ = len(v4)
                        }
                        _ = len(v4)
                    }
                }
            }
            _ = fi12
        }
    }
    var v16 int16 = -32768
    var v17 int8 = (v5 | (v5 - (^(v5))))
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
