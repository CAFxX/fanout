package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint; F2 uint64 }
func genFn0() bool {
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: 92, F1: -1}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS2{F0: 1}
        var v10 int16 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    return false
}

func main() {
    var v1 uint32 = 2
    var v2 string = ""
    var v3 float64 = 3.14159
    var v4 uint32 = 4294967295
    var v5 int64 = 1
    for fi11 := 0; fi11 < 10; fi11++ {
        if true {
            v3 = 1e10
        } else {
            fmt.Println(v4)
        }
        switch fi11 {
        case 128:
            fmt.Println(fi11, v2)
        default:
            if (0 < 255) {
                if true {
                    v12 := []int{23, fi11, fi11}
                    _ = len(v12)
                    _ = v12
                }
                var v13 int = fi11
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 3; fi14++ {
            switch v2 {
            case "\n\t":
                var v15 uint32 = 1
                _ = v15
            case "hello":
                v16 := GenS3{F0: (fi14 % (((fi14 * fi11)) | 1))}
                _ = v16
            }
            if (true && (106 > v5)) {
                v17 := []int{-2147483648, fi14, (fi14 ^ (^(fi11))), (fi14 ^ fi14)}
                v2 = v2
                switch v2 {
                case "foo bar":
                    var v18 uint32 = v4
                    _ = v18
                default:
                    var v19 uint16 = 65535
                    _ = v19
                }
                _ = v17
            } else {
                v20 := []int{fi14, (fi14 % (((fi11 / (((fi14 / ((-478) | 1))) | 1))) | 1)), fi11, (fi11 | (fi11 + int(v5)))}
                var v21 int = fi11
                var v22 int32 = 2147483647
                fmt.Println(v3)
                var v23 int16 = -32768
                _ = v20
                _ = v21
                _ = v22
                _ = v23
            }
            _ = fi14
        }
        _ = fi11
    }
    v3 = v3
    var v24 int16 = 456
    v25 := GenS3{F0: 128}
    if true {
        v26 := []int{}
        _ = v26
    } else {
        v1 += (^(v4))
        var v27 uint64 = 353
        v28 := GenS1{F0: 42, F1: int32(v27)}
        _ = v27
        _ = v28
    }
    v2 = "a"
    v4 += (v4 * v4)
    var v29 uint = 548
    var v30 int16 = 206
    v31 := GenS4{F0: false, F1: (v29 - v29), F2: 0}
    var v32 uint = 1
    v33 := []int{183, -1, -153, 127}
    var v34 string = (("" + "hello") + "a")
    var v35 bool = (!true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
