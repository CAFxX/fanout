package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
func genFn0(p0 uint32) uint16 {
    fmt.Println(p0)
    var v5 int = 127
    switch v5 {
    case 2147483647:
        p0 += (p0 + (p0 * p0))
    case 32767:
        v5 = (^(v5))
    }
    v5 = (v5 % (((-(v5))) | 1))
    switch (v5 * v5) {
    case 0:
        v5 -= v5
    case 549:
        switch v5 {
        case 1:
            switch "foo bar" {
            case "\n\t":
                if (v5 < (v5 ^ (^(v5)))) {
                    var v6 float32 = 1.5
                    _ = v6
                }
            case "foo bar":
                v7 := GenS1{F0: v5, F1: 65535}
                _ = v7
            case "hello":
                if (false || (!true)) {
                    switch ("hello" + "a") {
                    case "hello":
                        v8 := []int{}
                        _ = v8
                    case "foo bar":
                        for fi9 := 0; fi9 < 5; fi9++ {
                            v5 = (fi9 << ((962) & 63))
                            var v10 uint16 = 295
                            _ = fi9
                            _ = v10
                        }
                    case "":
                        for fi11 := 0; fi11 < 3; fi11++ {
                            fmt.Println(v5, v5)
                            _ = fi11
                        }
                    }
                }
            }
        }
    case -9223372036854775808:
        v5 += -265
    }
    _ = p0
    _ = v5
    return 378
}

func main() {
    var v1 uint16 = 1
    var v2 int64 = 0
    v3 := make(map[string]int)
    var v4 uint16 = 32767
    v1 = (v4 + v1)
    var v12 int16 = -996
    v13 := []int{-520}
    fmt.Println(v12, v12, v4)
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v1, v4)
        var v15 int = fi14
        v16 := GenS1{F0: v15, F1: 255}
        _ = fi14
        _ = v15
        _ = v16
    }
    if false {
        if false {
            var v17 float32 = 3.14159
            _ = v17
        }
        v18 := GenS1{F0: 32768, F1: 18446744073709551615}
        fmt.Println(v1, v4)
        _ = v18
    } else {
        v4 -= (v4 % (((v1 | 743)) | 1))
    }
    v19 := GenS1{F0: 440, F1: 18446744073709551615}
    var v20 float32 = -2.25
    switch 128 {
    case -851:
        fmt.Println(v12, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
