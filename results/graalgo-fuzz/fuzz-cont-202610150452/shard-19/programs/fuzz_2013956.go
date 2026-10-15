package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 bool = false
    var v2 int8 = 100
    var v3 int64 = -1
    var v4 float32 = -2.25
    v5 := make(map[int]uint32)
    var v6 int64 = v3
    var v7 float32 = float32((((((1e10 * (-2.25 * ((1e10 * 1e10) / 1.5))) + (3.14159 - 0.0)) / 1e10) * (0.0 * (1e10 - (3.14159 + (0.0 - 1e10))))) * float64(float32(0.0))))
    var v8 float64 = 1.5
    v8 -= 1e10
    switch 127 {
    case 59:
        switch -1 {
        case 9223372036854775807:
            v9 := GenS1{F0: ("" + ("" + ((("hello" + "") + "a") + ""))), F1: 63}
            _ = v9
        case 32768:
            var v10 bool = false
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        if v1 {
            var v12 int8 = v2
            switch (fi11 ^ (fi11 << ((70) & 63))) {
            case -2147483648:
                var v13 string = "foo bar"
                _ = v13
            case 2147483647:
                var v14 uint16 = 255
                _ = v14
            case -100:
                fmt.Println(v3, v7)
            }
            for fi15 := 0; fi15 < 11; fi15++ {
                if v1 {
                    v16 := GenS2{F0: -272}
                    var v17 float32 = 1.5
                    _ = v16
                    _ = v17
                }
                fmt.Println(fi15, v6)
                _ = fi15
            }
            _ = v12
        }
        v18 := GenS1{F0: "hello", F1: 255}
        v19 := []int{}
        _ = fi11
        _ = v18
        _ = v19
    }
    v20 := []int{}
    fmt.Println(v4)
    v4 += v7
    v21 := v5[-9223372036854775808]
    _ = len(v20)
    v22 := GenS1{F0: "hello", F1: 98}
    if true {
        v7 = v7
    }
    v3 = v6
    if false {
        if (("\n\t" + "a") > "foo bar") {
            fmt.Println(v2, v1)
        }
        var v23 uint64 = 784
        if (int(v21) == 128) {
            var v24 int8 = (v2 % ((v2) | 1))
            _ = v24
        }
        _ = v23
    } else {
        var v25 float32 = v4
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
