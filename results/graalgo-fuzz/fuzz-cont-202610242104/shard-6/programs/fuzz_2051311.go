package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int32; F1 int; F2 int64 }
type GenS3 struct { F0 uint; F1 uint8; F2 string }

func main() {
    var v1 float64 = ((0.0 + ((0.0 / -2.25) - (1.5 + (3.14159 * 1.5)))) / (float64((0.1 + (-2.25 * 3.14159))) / (1.5 / 3.14159)))
    var v2 uint64 = 42
    var v3 uint8 = 127
    if true {
        if true {
            v4 := []int{}
            _ = v4
        }
        switch ("hello" + "\n\t") {
        case "hello":
            v1 -= v1
        }
    }
    switch -2147483648 {
    case 347:
        for fi5 := 0; fi5 < 3; fi5++ {
            v3 = 1
            _ = fi5
        }
    case 128:
        var v6 float64 = v1
        _ = v6
    case 127:
        var v7 uint8 = v3
        _ = v7
    }
    fmt.Println(v1)
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: (fi8 % (((fi8 * (fi8 + 4294967295))) | 1)), F1: 855}
        v10 := GenS2{F0: -496, F1: fi8, F2: -9223372036854775808}
        v11 := GenS3{F0: 1, F1: 2, F2: ""}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v3 = (^(v3))
    var v12 int = -9223372036854775808
    var v13 int16 = -32768
    if true {
        v12 += v12
        if true {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 bool = false
                for fi16 := 0; fi16 < 12; fi16++ {
                    fmt.Println(v3)
                    fmt.Println(fi14)
                    _ = fi16
                }
                _ = fi14
                _ = v15
            }
            var v17 int64 = 9223372036854775807
            _ = v17
        }
        fmt.Println(v13)
    } else {
        var v18 uint64 = 18446744073709551615
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := []int{(v12 % (((v12 ^ fi19)) | 1)), (v12 << ((741) & 63)), fi19, v12}
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 5; fi21++ {
            fmt.Println(v18)
            _ = fi21
        }
        _ = v18
    }
    var v22 uint = 42
    var v23 uint64 = v2
    var v24 int64 = 0
    var v25 string = "foo bar"
    if false {
        v22 += v22
        var v26 int = (v12 * v12)
        v27 := GenS2{F0: 2147483647, F1: (^(v12)), F2: (v24 & (^(v24)))}
        _ = v26
        _ = v27
    }
    v28 := GenS3{F0: v22, F1: v3, F2: (v25 + v25)}
    v29 := []int{-58, v12}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
}
