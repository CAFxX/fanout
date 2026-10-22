package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 string }

func main() {
    var v1 uint64 = 2
    var v2 int8 = -1
    var v3 float64 = 1.5
    var v4 int64 = -1
    var v5 uint = 65535
    switch -2147483648 {
    case 4294967295:
        v6 := []int{-591}
        _ = v6
    case 128:
        v1 = v1
    case -1:
        v1 += (v1 | v1)
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(v5, v4, v5)
        switch (fi7 ^ (-(fi7))) {
        case 538:
            v8 := []int{-9223372036854775808, (-(fi7))}
            _ = v8
        case -557:
            v3 = v3
        }
        _ = fi7
    }
    v9 := []int{42}
    for fi10 := 0; fi10 < 15; fi10++ {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi10
    }
    var v11 int8 = (^(v2))
    if false {
        var v12 int8 = v2
        _ = v12
    } else {
        var v13 int16 = 3
        var v14 int32 = -1
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint8 = 42
        var v17 float32 = ((1.5 - 3.14159) - (1e10 - ((0.1 / (1.5 - 1e10)) + 1.5)))
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 string = ((("hello" + "\n\t") + ("foo bar" + "foo bar")) + ("\n\t" + "\n\t"))
    fmt.Println(v1, v18)
    if true {
        if true {
            v19 := GenS1{F0: 128, F1: 183, F2: (((v18 + v18) + ("foo bar" + v18)) + "foo bar")}
            _ = v19
        } else {
            v9 = append(v9, -185)
        }
        v5 = (v5 % ((v5) | 1))
        var v20 uint8 = 42
        _ = v20
    }
    var v21 bool = (v4 <= -1)
    v5 -= 255
    if true {
        var v22 uint16 = 32767
        var v23 bool = (int16(v22) > 1000)
        v18 = (("foo bar" + v18) + v18)
        _ = v22
        _ = v23
    }
    if v21 {
        v4 -= v4
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 float32 = 0.0
            _ = fi24
            _ = v25
        }
    } else {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    var v26 int = 135
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
