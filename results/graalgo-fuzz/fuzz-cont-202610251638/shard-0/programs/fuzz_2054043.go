package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int64 = -456
    var v2 bool = false
    var v3 uint = 4294967295
    var v4 int64 = 1
    var v5 float32 = float32(3.14159)
    var v6 uint = (v3 + v3)
    var v7 uint16 = 1
    var v8 bool = (-2147483648 != -2147483648)
    if (!(!v2)) {
        v9 := GenS1{F0: 0}
        var v10 string = "a"
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v1 = v1
        fmt.Println(v2, v7)
        _ = fi11
    }
    if true {
        fmt.Println(v3)
        v12 := GenS1{F0: (^(v3))}
        v1 += v4
        _ = v12
    }
    if (v8 && v2) {
        if v2 {
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 uint32 = 300
                for fi15 := 0; fi15 < 14; fi15++ {
                    var v16 float32 = ((1e10 * 0.0) + -2.25)
                    if (!(!(100 >= 100))) {
                        switch fi15 {
                        case 2147483647:
                            fmt.Println(v8, fi13)
                        case 0:
                            if (42 >= 2) {
                                v17 := []int{32767, fi13}
                                v8 = v2
                                _ = v17
                            }
                        case 42:
                            switch int(v6) {
                            case -1:
                                v8 = (1 == 1)
                            }
                        default:
                            v18 := []int{-1, (^(fi13)), (fi13 & 9223372036854775807)}
                            _ = v18
                        }
                    }
                    var v19 int8 = 127
                    _ = fi15
                    _ = v16
                    _ = v19
                }
                var v20 int64 = (v4 | (v4 >> (((v3 / (((^(v3))) | 1))) & 63)))
                _ = fi13
                _ = v14
                _ = v20
            }
            var v21 uint32 = 65535
            _ = v21
        }
    }
    v22 := GenS2{F0: -706}
    for fi23 := 0; fi23 < 6; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            v6 = v3
            _ = fi24
        }
        _ = fi23
    }
    v3 = 616
    switch 4294967295 {
    case -9223372036854775808:
        v1 -= 520
    }
    switch -1 {
    case 4294967295:
        fmt.Println(v2, v2, v7)
    case 2147483647:
        var v25 float32 = 3.14159
        _ = v25
    default:
        v1 = (^(v4))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
