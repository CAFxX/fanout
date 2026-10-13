package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int64; F1 string; F2 bool }
type GenS3 struct { F0 uint }
func genFn0() int64 {
    v6 := GenS1{F0: 764, F1: 341}
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint16 = 42
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 int8 = 127
    if (!(false || false)) {
        v11 := []int{-9223372036854775808}
        _ = len(v11)
        _ = v11
    }
    _ = v10
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 bool = false
    var v2 int16 = 1000
    var v3 uint64 = 4294967295
    v4 := []int{-274}
    var v5 int = 4294967295
    if v1 {
        var v12 int64 = 678
        _ = v12
    }
    var v13 uint32 = 2147483648
    fmt.Println(v2, v1, v5)
    var v14 uint16 = 65535
    v15 := GenS1{F0: int(v3), F1: 697}
    v2 += (^(v2))
    switch "\n\t" {
    case "hello":
        for fi16 := 0; fi16 < 18; fi16++ {
            fmt.Println(v1, v2, v13)
            for fi17 := 0; fi17 < 4; fi17++ {
                fmt.Println(v14, fi17, fi17)
                _ = fi17
            }
            for fi18 := 0; fi18 < 15; fi18++ {
                fmt.Println(v3)
                v13 = v13
                var v19 uint64 = (^(v3))
                _ = fi18
                _ = v19
            }
            _ = fi16
        }
    case "foo bar":
        var v20 uint8 = 255
        _ = v20
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v21 string = ("\n\t" + "a")
    for fi22 := 0; fi22 < 5; fi22++ {
        fmt.Println(fi22, v2, fi22)
        _ = fi22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        _ = len(v4)
        for fi24 := 0; fi24 < 20; fi24++ {
            switch v21 {
            case "hello":
                fmt.Println(fi24, fi23)
            case "foo bar":
                v25 := GenS2{F0: 1, F1: v21, F2: (128 <= 2)}
                _ = v25
            case "":
                fmt.Println(v13)
            default:
                fmt.Println(fi24, v13, v13)
            }
            _ = fi24
        }
        _ = fi23
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        if true {
            var v27 uint32 = (v13 % ((0) | 1))
            _ = v27
        }
        v28 := GenS3{F0: 0}
        _ = fi26
        _ = v28
    }
    v1 = true
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
