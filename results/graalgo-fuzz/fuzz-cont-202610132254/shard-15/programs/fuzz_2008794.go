package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int32; F2 int64 }
func genFn0(p0 uint16, p1 int32) int16 {
    for fi7 := 0; fi7 < 7; fi7++ {
        p0 = p0
        _ = fi7
    }
    p0 = p0
    _ = p0
    _ = p1
    return 103
}

func main() {
    var v1 string = "hello"
    var v2 string = "hello"
    var v3 uint32 = 0
    var v4 int64 = 0
    var v5 float64 = -2.25
    var v6 uint = 1
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := []int{fi8, fi8, fi8, (fi8 - fi8)}
        _ = fi8
        _ = v9
    }
    var v10 float64 = (v5 - (1e10 + 3.14159))
    var v11 int32 = 42
    var v12 string = v1
    v13 := []int{67, 2147483647}
    v14 := GenS1{F0: 0}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v15 bool = (!(3.14159 < 3.14159))
    if v15 {
        var v16 uint = v6
        _ = v16
    } else {
        if true {
            if v15 {
                v13 = append(v13, 128)
                if (v15 && (4294967295 == v6)) {
                    for fi17 := 0; fi17 < 3; fi17++ {
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        fmt.Println(v12, v2)
                        var v18 float32 = 1e10
                        _ = fi17
                        _ = v18
                    }
                    switch v1 {
                    case "a":
                        for fi19 := 0; fi19 < 5; fi19++ {
                            v20 := GenS1{F0: -68}
                            v4 = (v4 / ((int64(v6)) | 1))
                            _ = fi19
                            _ = v20
                        }
                    }
                }
            }
            for fi21 := 0; fi21 < 3; fi21++ {
                v22 := GenS2{F0: (v11 >> ((2) & 63)), F1: v11, F2: (v4 + v4)}
                _ = fi21
                _ = v22
            }
            var v23 float32 = 0.0
            _ = v23
        }
    }
    var v24 int64 = (^(v4))
    var v25 float32 = (3.14159 / 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
