package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() uint64 {
    var v5 int32 = 684
    v5 = v5
    v6 := GenS1{F0: 652}
    switch ("hello" + "") {
    case "\n\t":
        v7 := GenS1{F0: 32767}
        _ = v7
    case "a":
        if true {
            v8 := []int{122, -2147483648, -454, 351}
            fmt.Println(v5)
            v9 := GenS1{F0: 0}
            _ = v8
            _ = v9
        }
    case "hello":
        v5 = 1
    }
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 float32 = (1.5 / 3.14159)
    var v2 uint = 0
    var v3 float64 = 3.14159
    var v4 uint32 = 42
    var v10 uint32 = (v4 ^ (v4 ^ (v4 / ((v4) | 1))))
    if false {
        v10 -= 0
        var v11 uint16 = 65535
        if ((("a" + "") + "foo bar") <= "") {
            for fi12 := 0; fi12 < 3; fi12++ {
                v13 := []int{fi12}
                fmt.Println(v10, fi12, v3)
                _ = fi12
                _ = v13
            }
        }
        _ = v11
    } else {
        if false {
            v4 -= v10
        }
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := []int{(fi14 % (((-(fi14))) | 1)), fi14, -152, int(v2)}
        _ = fi14
        _ = v15
    }
    fmt.Println(v2, v2, v4)
    v2 = 327
    if true {
        v16 := []int{-9223372036854775808, -9223372036854775808, 128}
        for fi17 := 0; fi17 < 14; fi17++ {
            v10 -= v10
            _ = fi17
        }
        _ = v16
    }
    var v18 uint16 = 2
    var v19 float64 = v3
    v20 := []int{-692}
    var v21 int16 = 42
    var v22 string = ""
    for fi23 := 0; fi23 < 20; fi23++ {
        if len(v20) > 0 {
            v20[0] = fi23
        }
        if len(v20) > 0 {
            v20[0] = 2147483647
        }
        _ = fi23
    }
    v24 := GenS1{F0: (v21 ^ (v21 << (((v2 - (v2 - (v2 % (((v2 >> (((v2 & v2)) & 63))) | 1))))) & 63)))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}
