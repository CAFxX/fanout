package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint8 }
func genFn0() int32 {
    v4 := GenS1{F0: -9223372036854775808, F1: 32767}
    var v5 int32 = -891
    var v6 uint32 = 65535
    _ = v4
    _ = v5
    _ = v6
    return v5
}

func main() {
    var v1 string = ""
    v2 := make(map[string]string)
    var v3 int = 127
    var v7 float64 = (float64((-2.25 - (0.1 - ((0.0 / (float32(3.14159) - (-2.25 - (((1e10 + float32(0.1)) + -2.25) + -2.25)))) / 1.5)))) * 1.5)
    switch ("hello" + v1) {
    case "":
        v1 = ("foo bar" + "foo bar")
    case "foo bar":
        v7 -= v7
    default:
        v8 := []int{}
        _ = v8
    }
    v7 += 0.0
    var v9 float32 = 1.5
    fmt.Println(v1)
    delete(v2, "a")
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := v2[v1]
        _ = fi10
        _ = v11
    }
    switch (v3 / ((v3) | 1)) {
    case 0:
        fmt.Println(v3)
    case 32768:
        v1 = v1
    }
    v12 := GenS1{F0: -988, F1: 409}
    var v13 int64 = -198
    if true {
        v14 := GenS2{F0: 2}
        if (v3 == v3) {
            switch v1 {
            case "foo bar":
                v1 = "\n\t"
            }
            if (false && true) {
                var v15 uint = 819
                fmt.Println(v1, v1)
                _ = v15
            }
        } else {
            if ("a" != v1) {
                fmt.Println(v1)
                v7 = -2.25
            }
            v7 -= v7
        }
        _ = v14
    }
    var v16 int8 = 1
    v17 := GenS3{F0: v13, F1: 183, F2: 127}
    fmt.Println(v16, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
