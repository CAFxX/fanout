package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    v2 := []int{4294967295, -804}
    var v3 uint64 = 399
    v4 := []int{}
    var v5 int64 = 42
    var v6 int = -1
    switch ("" + "\n\t") {
    case "":
        if len(v4) > 0 {
            v4[0] = v6
        }
    case "hello":
        for fi7 := 0; fi7 < 3; fi7++ {
            _ = len(v1)
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := GenS1{F0: 175, F1: -2147483648, F2: 42}
                if false {
                    v6 -= 1
                    switch "" {
                    case "foo bar":
                        var v10 bool = false
                        _ = v10
                    }
                    delete(v1, "foo bar")
                    _ = len(v2)
                    var v11 int64 = (-(v5))
                    var v12 int8 = 87
                    var v13 int64 = (v11 + (^(v11)))
                    var v14 int16 = int16(v11)
                    _ = v11
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS1{F0: 1, F1: -1, F2: 42}
                    fmt.Println(v6)
                    _ = v15
                }
                v16 := GenS1{F0: 90, F1: -2147483648, F2: 452}
                _ = fi8
                _ = v16
                _ = v9
            }
            _ = fi7
        }
    }
    if (1.5 >= 0.1) {
        for fi17 := 0; fi17 < 2; fi17++ {
            v6 += fi17
            _ = fi17
        }
    }
    var v18 uint = 255
    v18 -= 18446744073709551615
    v3 = 42
    var v19 uint64 = (v3 & (^(v3)))
    for fi20 := 0; fi20 < 20; fi20++ {
        switch (fi20 | fi20) {
        case -118:
            v21 := GenS2{F0: 128}
            _ = v21
        }
        fmt.Println(v19, v18)
        _ = fi20
    }
    var v22 int64 = (-(v5))
    if true {
        v23 := GenS3{F0: v18}
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
