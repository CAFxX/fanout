package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{2147483647, 858}
    var v3 uint8 = 128
    v4 := make(map[string]int)
    var v5 int16 = -1
    var v6 uint16 = 65535
    var v7 float32 = 0.0
    var v8 int32 = 1
    if true {
        var v9 uint8 = (^(v3))
        _ = v9
    } else {
        v10 := GenS1{F0: -66, F1: 9}
        v11 := GenS1{F0: 1, F1: -1}
        var v12 int8 = -64
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 float64 = 3.14159
    var v14 uint8 = v3
    v6 -= v6
    switch -989 {
    case 350:
        switch "\n\t" {
        case "hello":
            var v15 int32 = v8
            _ = v15
        case "":
            v16 := GenS1{F0: -1, F1: 1}
            _ = v16
        case "foo bar":
            if false {
                for fi17 := 0; fi17 < 5; fi17++ {
                    switch "" {
                    case "\n\t":
                        for fi18 := 0; fi18 < 20; fi18++ {
                            fmt.Println(fi18, v5, fi18)
                            v13 += 1.5
                            v14 += v14
                            _ = fi18
                        }
                    case "a":
                        fmt.Println(v8)
                    default:
                        fmt.Println(v13, v5, fi17)
                    }
                    for fi19 := 0; fi19 < 9; fi19++ {
                        if (v7 <= v7) {
                            v3 += 2
                            var v20 uint = 2
                            _ = v20
                        } else {
                            switch fi17 {
                            case 128:
                                v21 := GenS1{F0: 127, F1: -31}
                                _ = v21
                            case 2147483647:
                                v22 := GenS1{F0: 100, F1: 987}
                                _ = v22
                            }
                        }
                        _ = fi19
                    }
                    if (!false) {
                        switch fi17 {
                        case -171:
                            v3 = v3
                        case 9223372036854775807:
                            delete(v4, ("hello" + "a"))
                        case 130:
                            var v23 int8 = 127
                            _ = v23
                        }
                    }
                    _ = fi17
                }
                if true {
                    var v24 uint64 = 18446744073709551615
                    _ = v24
                }
                v13 += v13
            }
        }
    case -762:
        if (-13 <= 0) {
            v25 := GenS1{F0: 100, F1: 1}
            _ = v25
        }
    case 42:
        if (0 < -1) {
            v3 = 128
        } else {
            v26 := GenS1{F0: -106, F1: 1}
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 16; fi27++ {
        fmt.Println(v14, v5, v8)
        _ = len(v2)
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
