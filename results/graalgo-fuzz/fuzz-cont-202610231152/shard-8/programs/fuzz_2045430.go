package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int8 = -36
    var v2 bool = (false || (!true))
    var v3 bool = v2
    var v4 uint16 = 676
    v5 := GenS1{F0: "", F1: 0}
    v1 = (v1 << ((860) & 63))
    v6 := []int{-382, 978}
    var v7 float32 = 1e10
    fmt.Println(v1)
    var v8 float32 = ((0.0 + v7) * v7)
    fmt.Println(v3)
    v7 = (v8 + v7)
    var v9 uint32 = 0
    switch 42 {
    case 32767:
        switch -2147483648 {
        case 763:
            fmt.Println(v8, v8, v4)
        case 2147483647:
            var v10 int64 = 9223372036854775807
            _ = v10
        }
    default:
        if (-1 != -2147483648) {
            v2 = ((1 == 42) && false)
            fmt.Println(v3, v2)
            var v11 string = ("foo bar" + "foo bar")
            _ = v11
        } else {
            switch ("foo bar" + (("hello" + "") + "foo bar")) {
            case "a":
                var v12 bool = false
                _ = v12
            default:
                for fi13 := 0; fi13 < 5; fi13++ {
                    v7 += (-2.25 - 0.1)
                    v7 -= (v8 / v7)
                    _ = fi13
                }
            }
            v1 = v1
            for fi14 := 0; fi14 < 19; fi14++ {
                fmt.Println(v4)
                var v15 uint64 = uint64(v4)
                _ = fi14
                _ = v15
            }
        }
    }
    switch 17 {
    case 128:
        var v16 uint16 = (v4 | v4)
        _ = v16
    }
    switch "\n\t" {
    case "hello":
        var v17 uint8 = 127
        _ = v17
    case "\n\t":
        var v18 float64 = 0.1
        _ = v18
    }
    fmt.Println(v8, v3)
    v8 -= v7
    var v19 uint32 = v9
    var v20 int32 = 1
    for fi21 := 0; fi21 < 19; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            if ((3.14159 == 0.0) || true) {
                if true {
                    fmt.Println(v9, v8)
                    v1 += (-(v1))
                    var v23 string = "a"
                    _ = v23
                } else {
                    var v24 uint64 = 331
                    _ = v24
                }
                if true {
                    switch fi21 {
                    case -1:
                        v25 := GenS1{F0: "a", F1: 42}
                        _ = v25
                    default:
                        fmt.Println(v2)
                        v7 = v8
                        if (!false) {
                            v26 := GenS2{F0: (^(v4)), F1: -2147483648}
                            fmt.Println(fi21, v20, v7)
                            _ = len(v6)
                            _ = v26
                        } else {
                            switch "hello" {
                            case "hello":
                                if v3 {
                                    fmt.Println(v20, v3)
                                    v2 = v3
                                    v27 := GenS3{F0: 962}
                                    _ = v27
                                }
                            case "\n\t":
                                var v28 uint8 = 212
                                _ = v28
                            }
                            var v29 uint16 = v4
                            _ = len(v6)
                            _ = v29
                        }
                    }
                    var v30 int = fi22
                    var v31 string = ("hello" + "\n\t")
                    _ = v30
                    _ = v31
                }
            }
            _ = fi22
        }
        switch "foo bar" {
        case "foo bar":
            var v32 uint64 = 4294967295
            _ = v32
        case "\n\t":
            var v33 float64 = 3.14159
            _ = v33
        }
        _ = fi21
    }
    var v34 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
