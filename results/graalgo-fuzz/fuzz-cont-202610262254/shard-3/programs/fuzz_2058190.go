package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
func genFn0() int16 {
    if (!false) {
        var v5 uint8 = 2
        v5 -= (v5 / ((2) | 1))
        _ = v5
    }
    var v6 float64 = 3.14159
    var v7 float32 = 3.14159
    v6 -= (((v6 * v6) + v6) / float64((-2.25 / 0.1)))
    _ = v6
    _ = v7
    return -592
}

func main() {
    var v1 uint64 = 733
    var v2 bool = (849 != -1)
    var v3 int8 = 127
    v4 := []int{1, 32767, 127}
    fmt.Println(v3, v1, v3)
    if false {
        v2 = false
    } else {
        if (!(766 >= v1)) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v1 -= (^(v1))
        }
    }
    v4 = append(v4, 607)
    var v8 string = "\n\t"
    fmt.Println(v3)
    var v9 int32 = 1
    fmt.Println(v3, v2, v3)
    v3 = (v3 / ((v3) | 1))
    v9 += v9
    if false {
        if len(v4) > 0 {
            v4[0] = 9223372036854775807
        }
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v2 = true
        if len(v4) > 0 {
            v4[0] = (fi10 + (fi10 + fi10))
        }
        _ = fi10
    }
    v4 = append(v4, 128)
    v11 := GenS1{F0: 267, F1: 439}
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS1{F0: 1, F1: fi12}
        fmt.Println(v9)
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 42, F1: 32767}
    switch (v8 + v8) {
    case "a":
        switch v8 {
        case "hello":
            _ = len(v4)
        default:
            v2 = (1 >= (v1 / (((v1 & v1)) | 1)))
        }
    case "\n\t":
        for fi15 := 0; fi15 < 12; fi15++ {
            fmt.Println(v8, fi15)
            for fi16 := 0; fi16 < 20; fi16++ {
                if ((float64(0.1) <= (1e10 * 1e10)) && (32767 < 42)) {
                    v9 -= v9
                }
                var v17 int16 = 1
                _ = fi16
                _ = v17
            }
            v18 := GenS1{F0: 158, F1: fi15}
            _ = fi15
            _ = v18
        }
    case "foo bar":
        for fi19 := 0; fi19 < 12; fi19++ {
            if (v2 || ((v3 ^ v3) >= (v3 ^ 67))) {
                v8 = "hello"
            }
            var v20 uint = 226
            _ = fi19
            _ = v20
        }
    }
    v4 = append(v4, 2147483647)
    for fi21 := 0; fi21 < 4; fi21++ {
        v1 += v1
        var v22 string = (v8 + "")
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
