package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }

func main() {
    v1 := []int{32768}
    var v2 uint8 = 25
    var v3 int64 = -9223372036854775808
    switch ("a" + "") {
    case "hello":
        var v4 uint32 = 926
        _ = v4
    case "":
        v5 := GenS1{F0: 1, F1: 1}
        _ = v5
    case "a":
        var v6 bool = (401 > 216)
        _ = v6
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        switch ("hello" + ("\n\t" + ("a" + "a"))) {
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            switch ("\n\t" + "a") {
            case "hello":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    } else {
        fmt.Println(v3, v3)
        if ((1.5 > 1e10) && (1 <= -1)) {
            v2 += 226
        } else {
            if false {
                v2 -= v2
                var v7 string = ("hello" + "a")
                _ = v7
            }
        }
        var v8 int32 = 2147483647
        _ = v8
    }
    var v9 int32 = -2147483648
    _ = len(v1)
    var v10 uint = 2
    var v11 int32 = (v9 / (((v9 >> (((v10 + v10)) & 63))) | 1))
    v3 -= v3
    var v12 float64 = ((3.14159 + float64((float32(float64(1.5)) + 3.14159))) - -2.25)
    fmt.Println(v12, v11, v12)
    v13 := GenS1{F0: (v10 - v10), F1: 1}
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 string = ("foo bar" + "\n\t")
        fmt.Println(v12, v2, v9)
        var v16 float64 = ((1e10 + v12) + (-2.25 + v12))
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
