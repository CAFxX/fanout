package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    v1 := []int{42, 127, 24, 32768}
    var v2 uint64 = 566
    var v3 uint64 = (v2 % (((v2 << ((75) & 63))) | 1))
    var v4 int8 = 0
    var v5 uint8 = 0
    switch "" {
    case "hello":
        var v6 uint32 = 42
        _ = v6
    case "\n\t":
        v3 -= (v3 ^ v3)
    }
    var v7 uint = 886
    v3 = v3
    var v8 uint8 = 160
    if true {
        switch ("\n\t" + ("" + ("" + ("" + ("\n\t" + ("a" + ("a" + "\n\t"))))))) {
        case "foo bar":
            var v9 int64 = 487
            _ = v9
        case "hello":
            fmt.Println(v8)
        }
    } else {
        switch 32768 {
        case -740:
            if (!(v7 >= v7)) {
                v7 = 925
                v3 -= 4294967295
                v10 := GenS1{F0: (^(v2))}
                _ = v10
            }
        case 2147483647:
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 bool = true
                fmt.Println(v5)
                for fi13 := 0; fi13 < 20; fi13++ {
                    v14 := GenS2{F0: ""}
                    var v15 int16 = 1000
                    v16 := GenS1{F0: v2}
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = fi11
                _ = v12
            }
        default:
            var v17 int16 = 1000
            _ = v17
        }
    }
    _ = len(v1)
    var v18 int64 = -327
    if true {
        v8 += (v5 * 172)
        fmt.Println(v5)
    } else {
        fmt.Println(v5)
        for fi19 := 0; fi19 < 7; fi19++ {
            v4 += v4
            if true {
                _ = len(v1)
                if ((!true) && ((879 > 0) && true)) {
                    v18 += v18
                }
            }
            _ = fi19
        }
    }
    if false {
        v20 := GenS3{F0: v18, F1: 65535}
        var v21 int8 = (^(v4))
        _ = v20
        _ = v21
    } else {
        switch -396 {
        case -9223372036854775808:
            var v22 int32 = -2147483648
            _ = v22
        case 4294967295:
            var v23 bool = false
            _ = v23
        case -987:
            var v24 uint8 = (v5 & 127)
            _ = v24
        default:
            v18 += (v18 | 925)
        }
    }
    var v25 float64 = 0.0
    if len(v1) > 0 {
        v1[0] = 991
    }
    if len(v1) > 0 {
        v1[0] = -995
    }
    var v26 string = "a"
    if true {
        var v27 uint64 = 2
        _ = v27
    } else {
        switch 32768 {
        case -738:
            var v28 uint64 = (^(v3))
            _ = v28
        default:
            v1 = append(v1, 42)
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
