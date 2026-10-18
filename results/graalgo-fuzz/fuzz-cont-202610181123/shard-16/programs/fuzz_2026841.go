package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint64 }

func main() {
    var v1 bool = (!true)
    v2 := make(map[uint32]string)
    var v3 int = 2147483647
    v4 := make(map[uint32]bool)
    v5 := make(map[int]uint32)
    v6 := make(map[string]uint32)
    v3 = (-(v3))
    v4[0] = false
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 249, F1: 4294967295, F2: 959}
        if v1 {
            var v9 float64 = float64(3.14159)
            _ = v9
        } else {
            var v10 uint64 = 753
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    if false {
        var v11 bool = v1
        v12 := []int{}
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 int8 = 42
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            _ = fi13
            _ = v14
        }
        _ = v11
        _ = v12
    }
    fmt.Println(v1)
    fmt.Println(v1, v3)
    if true {
        for fi15 := 0; fi15 < 11; fi15++ {
            for fi16 := 0; fi16 < 20; fi16++ {
                if v1 {
                    v1 = (!v1)
                    v17 := []int{fi15, 32768, v3}
                    if true {
                        var v18 int = fi15
                        var v19 uint64 = 208
                        _ = v18
                        _ = v19
                    }
                    _ = v17
                }
                fmt.Println(v1)
                v20 := []int{}
                _ = fi16
                _ = v20
            }
            _ = fi15
        }
        fmt.Println(v1)
        v3 -= -9223372036854775808
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        switch "hello" {
        case "":
            var v22 float64 = 0.1
            _ = v22
        }
        v3 -= (^(fi21))
        v23 := GenS1{F0: 399, F1: 579, F2: 266}
        _ = fi21
        _ = v23
    }
    switch "a" {
    case "hello":
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            v3 = v3
        case "hello":
            var v24 float32 = (3.14159 * ((3.14159 * (float32((1e10 / 1.0)) * 0.0)) + 0.0))
            _ = v24
        case "\n\t":
            v1 = v1
        }
    case "":
        for fi25 := 0; fi25 < 1; fi25++ {
            switch (fi25 + (v3 + 32768)) {
            case 4294967295:
                var v26 float32 = (3.14159 * 3.14159)
                _ = v26
            case 2147483647:
                switch (("hello" + "hello") + "a") {
                case "\n\t":
                    fmt.Println(v1)
                case "foo bar":
                    delete(v4, 42)
                case "":
                    fmt.Println(fi25, fi25)
                }
            case -1:
                fmt.Println(fi25, v1)
            default:
                v1 = v1
            }
            var v27 uint64 = 42
            _ = fi25
            _ = v27
        }
    case "\n\t":
        fmt.Println(v1, v3)
    default:
        v1 = v1
    }
    v28 := GenS1{F0: 4294967295, F1: 1, F2: 169}
    v29 := GenS1{F0: 1, F1: 2, F2: 379}
    var v30 uint = 255
    v31 := []int{127, 42, v3, 32767}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
