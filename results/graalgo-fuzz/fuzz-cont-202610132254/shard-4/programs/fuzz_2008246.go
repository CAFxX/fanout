package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint32 = 4294967295
    var v2 uint8 = 1
    var v3 int32 = 0
    var v4 string = "a"
    v3 = v3
    switch v4 {
    case "foo bar":
        v5 := []int{-9223372036854775808, 1, -1, -2147483648}
        _ = v5
    default:
        v4 = (((v4 + v4) + v4) + "foo bar")
    }
    if (2 < 464) {
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: (^(fi6))}
            if (32767 <= -32768) {
                v8 := GenS1{F0: 32768}
                v9 := []int{-9223372036854775808, fi6, (^(fi6))}
                _ = v8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            v2 = v2
            var v11 uint64 = 535
            fmt.Println(v3, v11)
            _ = fi10
            _ = v11
        }
        v1 += v1
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        v3 += v3
        v13 := GenS1{F0: 1}
        _ = fi12
        _ = v13
    }
    if true {
        switch 32767 {
        case 4294967295:
            if true {
                if false {
                    var v14 float32 = (0.1 + (0.0 / 1e10))
                    _ = v14
                } else {
                    var v15 bool = false
                    _ = v15
                }
            }
        }
        v2 += v2
        v16 := GenS1{F0: 4294967295}
        _ = v16
    } else {
        if (0 > -876) {
            for fi17 := 0; fi17 < 10; fi17++ {
                if ((v1 + (v1 & v1)) < v1) {
                    v18 := GenS1{F0: fi17}
                    var v19 int8 = 0
                    var v20 float32 = 0.1
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(v4, v3, fi17)
                }
                _ = fi17
            }
        }
        var v21 int32 = (v3 >> ((2) & 63))
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v2 += 0
        v2 = 128
        _ = fi22
    }
    var v23 int = 9223372036854775807
    for fi24 := 0; fi24 < 17; fi24++ {
        fmt.Println(v1)
        var v25 uint = 42
        fmt.Println(fi24, v1)
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
