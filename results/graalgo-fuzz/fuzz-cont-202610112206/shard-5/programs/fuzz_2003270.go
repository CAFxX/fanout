package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int64 = 1
    var v2 uint32 = 2
    var v3 uint32 = 691
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 bool = true
        v1 = (v1 * v1)
        v6 := []int{(fi4 % (((fi4 * 0)) | 1)), -597, 9223372036854775807, fi4}
        _ = fi4
        _ = v5
        _ = v6
    }
    switch ("foo bar" + "a") {
    case "a":
        switch ((("" + "a") + ((("a" + ((("" + "foo bar") + "hello") + "foo bar")) + "") + "foo bar")) + ("\n\t" + ("" + "hello"))) {
        case "":
            v7 := []int{}
            _ = v7
        }
    default:
        var v8 float32 = (-2.25 + (0.1 * ((1.5 + float32(0.1)) + 3.14159)))
        _ = v8
    }
    var v9 float32 = 3.14159
    v10 := GenS1{F0: 2}
    v11 := []int{}
    var v12 uint16 = 65535
    v3 += v2
    var v13 uint64 = 441
    var v14 uint = 528
    switch ("" + "hello") {
    case "a":
        for fi15 := 0; fi15 < 8; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                switch "\n\t" {
                case "a":
                    var v17 uint8 = uint8(v13)
                    _ = v17
                case "foo bar":
                    var v18 uint = v14
                    _ = v18
                case "":
                    v1 = -1
                }
                var v19 int64 = v1
                _ = fi16
                _ = v19
            }
            var v20 int = fi15
            _ = fi15
            _ = v20
        }
    case "foo bar":
        if false {
            var v21 bool = (1.5 > float32(-2.25))
            _ = v21
        }
    case "hello":
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    default:
        v3 = (v3 & (v2 >> ((v14) & 63)))
    }
    fmt.Println(v14)
    fmt.Println(v1, v9, v14)
    var v22 uint64 = uint64(v12)
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 float32 = ((-2.25 - v9) * 0.1)
        var v25 int = 42
        _ = fi23
        _ = v24
        _ = v25
    }
    if false {
        fmt.Println(v2, v3, v1)
    }
    v14 -= (v14 | 42)
    for fi26 := 0; fi26 < 14; fi26++ {
        v9 = (v9 / 1.5)
        v3 = v2
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v9
}
