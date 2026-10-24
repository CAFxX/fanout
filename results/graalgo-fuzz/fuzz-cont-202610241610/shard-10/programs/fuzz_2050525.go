package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint32; F1 int32; F2 uint16 }

func main() {
    var v1 int8 = 127
    var v2 string = "hello"
    v3 := make(map[uint32]bool)
    var v4 int64 = 399
    v5 := []int{-166}
    var v6 int = -1
    var v7 int16 = -443
    v8 := GenS1{F0: "", F1: 624, F2: 18446744073709551615}
    for fi9 := 0; fi9 < 17; fi9++ {
        switch ("" + "hello") {
        case "hello":
            if true {
                if false {
                    v10 := GenS2{F0: 42, F1: (fi9 << ((255) & 63))}
                    switch -1 {
                    case 633:
                        var v11 int8 = v1
                        _ = v11
                    case -673:
                        var v12 bool = ((!true) && (1 >= 372))
                        _ = v12
                    }
                    _ = v10
                }
                var v13 int8 = (v1 / (((v1 << ((535) & 63))) | 1))
                _ = v13
            }
        case "a":
            v4 += v4
        case "foo bar":
            v7 = (-(v7))
        }
        _ = fi9
    }
    var v14 uint32 = 712
    v3[v14] = (false && (v7 > (v7 % ((v7) | 1))))
    var v15 int64 = v4
    var v16 int = (v6 % ((-544) | 1))
    _ = len(v3)
    v15 -= (-(v15))
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 string = ""
        _ = fi17
        _ = v18
    }
    fmt.Println(v2, v1)
    fmt.Println(v14, v15)
    v6 += v16
    if true {
        v19 := GenS3{F0: (v14 & v14), F1: -2147483648, F2: 917}
        var v20 uint64 = 0
        switch (v2 + "") {
        case "foo bar":
            var v21 uint32 = 19
            _ = v21
        }
        _ = v19
        _ = v20
    }
    v16 += -2147483648
    fmt.Println(v7)
    v15 -= v4
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
