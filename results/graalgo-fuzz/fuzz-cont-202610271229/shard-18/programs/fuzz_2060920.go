package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint32; F2 int8 }

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]string)
    var v3 uint32 = 2147483648
    var v4 uint8 = 128
    v1 = v1
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int64 = v1
        switch fi5 {
        case -9223372036854775808:
            delete(v2, v3)
        case -1:
            if (fi5 != fi5) {
                for fi7 := 0; fi7 < 5; fi7++ {
                    v4 = v4
                    _ = fi7
                }
            }
        default:
            v8 := []int{}
            _ = v8
        }
        var v9 int = fi5
        _ = fi5
        _ = v6
        _ = v9
    }
    v1 = (v1 - v1)
    if (("hello" + "") != "foo bar") {
        var v10 uint64 = 0
        var v11 uint64 = v10
        switch "hello" {
        case "foo bar":
            fmt.Println(v10, v10)
        case "":
            for fi12 := 0; fi12 < 19; fi12++ {
                switch "foo bar" {
                case "a":
                    fmt.Println(v10, v10)
                case "foo bar":
                    v11 = v10
                case "":
                    for fi13 := 0; fi13 < 15; fi13++ {
                        switch ("hello" + "hello") {
                        case "a":
                            var v14 int32 = 2147483647
                            _ = v14
                        case "":
                            var v15 int64 = (v1 * v1)
                            _ = v15
                        case "\n\t":
                            v16 := []int{(fi13 << ((0) & 63)), fi13, 1}
                            _ = v16
                        default:
                            var v17 string = ("" + ("" + "hello"))
                            _ = v17
                        }
                        _ = fi13
                    }
                default:
                    var v18 float64 = 1e10
                    _ = v18
                }
                _ = fi12
            }
        default:
            v19 := []int{-2147483648, 42, 4294967295, 32767}
            _ = v19
        }
        _ = v10
        _ = v11
    } else {
        v4 -= v4
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := v2[v3]
            _ = fi20
            _ = v21
        }
        var v22 int32 = 42
        _ = v22
    }
    fmt.Println(v1, v3)
    for fi23 := 0; fi23 < 19; fi23++ {
        delete(v2, 2)
        v24 := GenS1{F0: 1}
        _ = fi23
        _ = v24
    }
    _ = len(v2)
    if true {
        _ = len(v2)
        if (42 == 482) {
            var v25 uint = 42
            switch 32767 {
            case 2147483647:
                fmt.Println(v3)
            default:
                v4 -= 1
            }
            var v26 int32 = 1
            _ = v25
            _ = v26
        } else {
            for fi27 := 0; fi27 < 20; fi27++ {
                switch fi27 {
                case 2147483647:
                    var v28 uint32 = v3
                    _ = v28
                }
                var v29 string = "a"
                if true {
                    for fi30 := 0; fi30 < 15; fi30++ {
                        v4 = 128
                        _ = len(v2)
                        _ = fi30
                    }
                    var v31 uint32 = v3
                    var v32 uint16 = 0
                    _ = v31
                    _ = v32
                } else {
                    switch v29 {
                    case "a":
                        v29 = ((v29 + ("" + (v29 + ("" + v29)))) + ((v29 + v29) + v29))
                    }
                    switch v29 {
                    case "foo bar":
                        var v33 int32 = -2147483648
                        _ = v33
                    }
                }
                _ = fi27
                _ = v29
            }
            var v34 uint64 = uint64(v4)
            if false {
                v34 += v34
            }
            _ = v34
        }
    } else {
        fmt.Println(v4, v1)
    }
    v35 := GenS2{F0: (v3 ^ (v3 | 875)), F1: -2147483648, F2: 1}
    v36 := GenS3{F0: 548, F1: 65535, F2: int8(v4)}
    var v37 int = 0
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
