package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 bool = (true || (0 <= 42))
    var v2 int = 517
    var v3 float64 = 0.1
    v4 := make(map[string]string)
    v5 := GenS1{F0: (v2 % ((v2) | 1)), F1: 127, F2: 18446744073709551615}
    switch ("" + "foo bar") {
    case "":
        var v6 bool = true
        _ = v6
    case "hello":
        v7 := GenS2{F0: -1}
        _ = v7
    }
    if (177 >= 708) {
        fmt.Println(v2, v3)
        v8 := []int{v2}
        _ = v8
    } else {
        if (!v1) {
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 float32 = ((3.14159 - 0.0) / 0.1)
                if true {
                    if ((!(65535 > 42)) && v1) {
                        var v11 uint = 61
                        v12 := []int{(fi9 + 32768)}
                        _ = v11
                        _ = v12
                    } else {
                        v2 = v2
                    }
                    fmt.Println(v2)
                } else {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        var v14 uint = 42
                        switch "hello" {
                        case "foo bar":
                            var v15 bool = false
                            _ = v15
                        }
                        _ = fi13
                        _ = v14
                    }
                    var v16 uint16 = 42
                    v1 = ((v10 > v10) && (!(!v1)))
                    _ = v16
                }
                _ = fi9
                _ = v10
            }
            v2 = v2
        }
    }
    v1 = (v1 && v1)
    if v1 {
        delete(v4, "hello")
    } else {
        v17 := GenS2{F0: 118}
        _ = v17
    }
    var v18 uint16 = 65535
    v3 -= v3
    fmt.Println(v18, v1, v1)
    v18 += v18
    if (!(v1 && (250 > 4294967295))) {
        var v19 int32 = -373
        fmt.Println(v19, v2)
        _ = v19
    }
    switch 852 {
    case 791:
        for fi20 := 0; fi20 < 5; fi20++ {
            fmt.Println(v3, v18, v18)
            v21 := GenS3{F0: 42}
            var v22 uint64 = 2
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    v2 = (^(v2))
    v23 := GenS1{F0: v2, F1: (v2 + v2), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
