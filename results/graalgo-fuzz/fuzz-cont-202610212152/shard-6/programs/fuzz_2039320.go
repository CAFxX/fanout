package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int64 = 42
    v2 := []int{127, -372}
    var v3 uint16 = 2
    var v4 uint16 = (^(v3))
    var v5 int8 = 1
    if ((0 < 100) && true) {
        switch -2147483648 {
        case 128:
            v2 = append(v2, int(v5))
        case -2147483648:
            v6 := GenS1{F0: v5, F1: -1}
            _ = v6
        case -1:
            if len(v2) > 0 {
                v2[0] = -1
            }
        }
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        switch "" {
        case "":
            switch "\n\t" {
            case "\n\t":
                if len(v2) > 0 {
                    v2[0] = fi7
                }
            default:
                if (("\n\t" + "hello") == ("hello" + "\n\t")) {
                    for fi8 := 0; fi8 < 7; fi8++ {
                        _ = len(v2)
                        _ = fi8
                    }
                    _ = len(v2)
                    var v9 uint32 = 4294967295
                    _ = v9
                }
            }
        }
        _ = fi7
    }
    var v10 uint = 42
    for fi11 := 0; fi11 < 13; fi11++ {
        v2 = append(v2, (fi11 & 214))
        _ = fi11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint = 705
        _ = fi12
        _ = v13
    }
    var v14 int = -1
    var v15 int16 = 0
    v2 = append(v2, v14)
    for fi16 := 0; fi16 < 12; fi16++ {
        fmt.Println(fi16)
        if true {
            var v17 uint16 = (v3 * v4)
            fmt.Println(v3)
            var v18 uint = v10
            _ = v17
            _ = v18
        }
        if len(v2) > 0 {
            v2[0] = (v14 % ((fi16) | 1))
        }
        _ = fi16
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        _ = len(v2)
        if (((0 > 0) || false) || true) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v15, v5, v4)
            switch v14 {
            case -577:
                if (2147483647 < -62) {
                    _ = len(v2)
                    v15 = int16(v1)
                    if false {
                        var v20 uint16 = v4
                        fmt.Println(v14, v15)
                        var v21 int32 = 2147483647
                        _ = v20
                        _ = v21
                    }
                }
            case 0:
                for fi22 := 0; fi22 < 5; fi22++ {
                    if len(v2) > 0 {
                        v2[0] = (v14 % ((-755) | 1))
                    }
                    var v23 float32 = 1e10
                    _ = fi22
                    _ = v23
                }
            case -2147483648:
                for fi24 := 0; fi24 < 1; fi24++ {
                    fmt.Println(fi19, v4, v5)
                    if true {
                        fmt.Println(v5, v15)
                    }
                    v25 := GenS2{F0: (v10 & v10), F1: v3}
                    v15 = v15
                    v15 = (v15 + v15)
                    if false {
                        v5 -= (v5 + (v5 / (((v5 / ((v5) | 1))) | 1)))
                    }
                    var v26 uint8 = 112
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
        }
        _ = fi19
    }
    v14 = v14
    var v27 uint32 = 65535
    for fi28 := 0; fi28 < 11; fi28++ {
        v29 := GenS3{F0: v27}
        for fi30 := 0; fi30 < 1; fi30++ {
            v31 := GenS3{F0: (v27 / ((v27) | 1))}
            v10 -= (v10 + v10)
            _ = fi30
            _ = v31
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi28
        _ = v29
    }
    var v32 uint = (v10 >> ((uint(v27)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
