package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int, p1 uint8, p2 bool) bool {
    for fi4 := 0; fi4 < 17; fi4++ {
        switch ("a" + "hello") {
        case "foo bar":
            var v5 uint32 = 1
            _ = v5
        case "hello":
            p2 = p2
        case "a":
            var v6 int32 = 0
            _ = v6
        }
        p1 -= p1
        _ = fi4
    }
    if (0 == 0) {
        switch p0 {
        case -9223372036854775808:
            if (65535 == 1) {
                fmt.Println(p1)
            }
        }
        switch "hello" {
        case "hello":
            var v7 uint = 65535
            _ = v7
        case "":
            v8 := GenS1{F0: 128}
            _ = v8
        }
    } else {
        p1 -= 10
        v9 := GenS2{F0: 2, F1: 1}
        _ = v9
    }
    var v10 int32 = 42
    fmt.Println(p2, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p2
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int32 = 2147483647
    v3 := make(map[int]string)
    var v11 float32 = 0.1
    v1 = v1
    switch "a" {
    case "\n\t":
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 15; fi13++ {
                fmt.Println(v11, v2, fi12)
                var v14 float32 = (v11 - v11)
                _ = fi13
                _ = v14
            }
            v15 := v3[56]
            _ = fi12
            _ = v15
        }
    case "a":
        fmt.Println(v11)
    case "":
        var v16 int8 = 100
        _ = v16
    }
    delete(v3, 2147483647)
    if (0.0 >= 0.1) {
        switch "" {
        case "\n\t":
            v1 += v1
        case "foo bar":
            switch 2147483647 {
            case -1:
                for fi17 := 0; fi17 < 14; fi17++ {
                    v18 := []int{fi17, fi17, fi17, -2147483648}
                    var v19 string = ("" + "")
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            case 128:
                var v20 uint = 392
                _ = v20
            case 9223372036854775807:
                v11 = (1e10 / (v11 - 1.5))
            }
        }
        v21 := GenS1{F0: 128}
        _ = v21
    } else {
        v2 += v2
    }
    var v22 string = "\n\t"
    fmt.Println(v22, v22)
    for fi23 := 0; fi23 < 15; fi23++ {
        v24 := []int{408, fi23, (fi23 % (((fi23 % (((fi23 ^ (fi23 + 2147483647))) | 1))) | 1))}
        fmt.Println(v2)
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        if (!(!(1 != 0))) {
            for fi26 := 0; fi26 < 7; fi26++ {
                switch "hello" {
                case "":
                    v27 := GenS2{F0: 4294967295, F1: 608}
                    _ = v27
                case "hello":
                    fmt.Println(v11)
                default:
                    fmt.Println(fi25, v22, v1)
                }
                _ = fi26
            }
        }
        v28 := GenS1{F0: 2}
        _ = fi25
        _ = v28
    }
    fmt.Println(v1, v1)
    fmt.Println(v22, v1, v22)
    for fi29 := 0; fi29 < 11; fi29++ {
        v1 = 0
        v3[127] = v22
        _ = fi29
    }
    var v30 float64 = (1.5 * float64((float32((0.1 - ((0.0 + 1.5) + 0.1))) / (v11 / (3.14159 / -2.25)))))
    var v31 int64 = -861
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v31
}
