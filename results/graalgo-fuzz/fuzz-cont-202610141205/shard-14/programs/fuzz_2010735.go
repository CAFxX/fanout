package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 bool }

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 1e10
    var v3 string = ("a" + ((("foo bar" + "a") + "a") + "foo bar"))
    var v4 bool = (!(0 < 18446744073709551615))
    v5 := make(map[string]string)
    var v6 uint16 = 488
    for fi7 := 0; fi7 < 12; fi7++ {
        if false {
            for fi8 := 0; fi8 < 5; fi8++ {
                v3 = ((v3 + v3) + "a")
                switch (v3 + v3) {
                case "foo bar":
                    var v9 int32 = 2147483647
                    fmt.Println(fi7)
                    v3 = "a"
                    _ = v9
                case "hello":
                    fmt.Println(fi7, v3)
                case "\n\t":
                    v1[v3] = v3
                }
                v10 := []int{}
                v11 := GenS1{F0: 0, F1: 142, F2: (false || v4)}
                _ = fi8
                _ = v10
                _ = v11
            }
            var v12 string = v3
            _ = v12
        }
        _ = fi7
    }
    v13 := GenS1{F0: 656, F1: 2, F2: false}
    v14 := GenS1{F0: 0, F1: 2147483647, F2: (32767 == 1)}
    v4 = false
    if (603 < 2) {
        v15 := []int{128, 421, -230, 128}
        for fi16 := 0; fi16 < 11; fi16++ {
            v3 = "hello"
            for fi17 := 0; fi17 < 10; fi17++ {
                if true {
                    var v18 float64 = 0.0
                    if true {
                        fmt.Println(v2, v4)
                        v4 = true
                    }
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v4, v4)
                    case "foo bar":
                        var v19 int64 = 1
                        _ = v19
                    }
                    v20 := v1[v3]
                    var v21 int32 = -217
                    _ = v18
                    _ = v20
                    _ = v21
                }
                _ = len(v5)
                fmt.Println(v4, fi17, v4)
                _ = fi17
            }
            _ = fi16
        }
        var v22 int64 = -9223372036854775808
        _ = v15
        _ = v22
    } else {
        switch "\n\t" {
        case "\n\t":
            v23 := GenS1{F0: int32(v6), F1: 1, F2: false}
            _ = v23
        case "hello":
            var v24 bool = v4
            _ = v24
        case "":
            fmt.Println(v6, v2)
        }
        fmt.Println(v4, v3)
    }
    var v25 bool = (!((1 >= 983) || false))
    fmt.Println(v2, v25)
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := []int{}
        _ = fi26
        _ = v27
    }
    if (v2 <= 0.1) {
        v1[v3] = ((v3 + "hello") + "foo bar")
        switch v3 {
        case "":
            v25 = v4
        default:
            var v28 bool = (("hello" + v3) <= v3)
            _ = v28
        }
    }
    fmt.Println(v6, v6, v25)
    var v29 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
