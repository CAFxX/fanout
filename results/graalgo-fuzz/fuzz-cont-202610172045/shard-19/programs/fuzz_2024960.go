package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }

func main() {
    var v1 bool = true
    var v2 int = -43
    var v3 uint32 = 470
    v4 := make(map[string]bool)
    v5 := []uint32{}
    v3 = 2
    if (3.14159 > 1.5) {
        switch (v2 - 4294967295) {
        case 948:
            v6 := GenS1{F0: -656, F1: 574, F2: "a"}
            _ = v6
        case 993:
            if false {
                v5 = append(v5, v3)
                switch v2 {
                case -9223372036854775808:
                    for fi7 := 0; fi7 < 9; fi7++ {
                        for fi8 := 0; fi8 < 18; fi8++ {
                            v3 += v3
                            _ = fi8
                        }
                        fmt.Println(v1, v3, v2)
                        _ = fi7
                    }
                case 4294967295:
                    var v9 uint = 971
                    _ = v9
                }
            } else {
                var v10 int64 = -986
                var v11 uint64 = 980
                _ = v10
                _ = v11
            }
        case -2147483648:
            v3 = (^(v3))
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if true {
        v12 := GenS2{F0: 0}
        _ = v12
    } else {
        fmt.Println(v1, v1)
        var v13 uint64 = 1
        for fi14 := 0; fi14 < 16; fi14++ {
            switch "a" {
            case "a":
                var v15 float64 = (0.1 * 1.5)
                _ = v15
            }
            _ = fi14
        }
        _ = v13
    }
    var v16 int64 = -481
    var v17 int8 = 87
    fmt.Println(v16)
    var v18 uint32 = v3
    for fi19 := 0; fi19 < 3; fi19++ {
        if v1 {
            v2 = v2
        } else {
            v20 := GenS3{F0: "foo bar"}
            _ = v20
        }
        _ = fi19
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 int8 = 100
        _ = fi21
        _ = v22
    }
    v23 := GenS2{F0: v17}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
