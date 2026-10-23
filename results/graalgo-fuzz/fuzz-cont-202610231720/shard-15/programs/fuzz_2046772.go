package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int64, p1 uint8, p2 uint) int {
    if false {
        var v7 int32 = -682
        _ = v7
    } else {
        var v8 int16 = -1
        _ = v8
    }
    v9 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -789
}

func main() {
    v1 := []uint32{}
    var v2 float64 = 1.5
    v3 := make(map[uint32]string)
    var v4 uint = 18446744073709551615
    var v5 uint = 437
    var v6 float64 = (3.14159 - 0.1)
    fmt.Println(v6)
    v3[373] = "foo bar"
    var v10 string = "hello"
    v4 = (v5 | 341)
    v5 = 1
    v10 = "\n\t"
    v11 := GenS1{F0: 2}
    var v12 bool = false
    var v13 uint16 = 255
    v5 += (v5 | (v4 ^ uint(v13)))
    v1 = append(v1, 2147483648)
    var v14 float64 = 0.1
    var v15 string = v10
    switch v10 {
    case "\n\t":
        v16 := v3[42]
        _ = v16
    }
    v17 := GenS1{F0: 255}
    switch -101 {
    case 2147483647:
        for fi18 := 0; fi18 < 10; fi18++ {
            if ((1 < 45) || (42 >= 9223372036854775807)) {
                fmt.Println(v14)
                var v19 float64 = (1.5 * v14)
                _ = v19
            }
            _ = fi18
        }
    case 32768:
        for fi20 := 0; fi20 < 8; fi20++ {
            fmt.Println(v12, v4, v5)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v12)
            _ = fi20
        }
    default:
        switch int(v4) {
        case -9223372036854775808:
            var v21 float32 = 0.0
            _ = v21
        case 2147483647:
            for fi22 := 0; fi22 < 11; fi22++ {
                for fi23 := 0; fi23 < 1; fi23++ {
                    if (false || false) {
                        switch (v15 + "\n\t") {
                        case "foo bar":
                            var v24 int8 = 2
                            _ = v24
                        }
                        v5 = 1
                        v10 = "foo bar"
                    } else {
                        var v25 uint = (v4 | 2)
                        switch (fi23 >> (((v5 / ((v4) | 1))) & 63)) {
                        case 127:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 32767:
                            switch genFn0(-9223372036854775808, 255, 157) {
                            case -9223372036854775808:
                                v12 = (!v12)
                            }
                        default:
                            v26 := GenS2{F0: (!false)}
                            _ = v26
                        }
                        _ = v25
                    }
                    var v27 uint = v4
                    var v28 uint32 = 4294967295
                    _ = fi23
                    _ = v27
                    _ = v28
                }
                _ = fi22
            }
        case 32767:
            var v29 bool = (-334 > 1)
            _ = v29
        default:
            v30 := GenS1{F0: (v13 | v13)}
            _ = v30
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 9223372036854775807 {
    case 42:
        for fi31 := 0; fi31 < 19; fi31++ {
            v3[526] = (v10 + ("" + v10))
            _ = fi31
        }
    case 32767:
        var v32 uint16 = (^(v13))
        _ = v32
    default:
        if v12 {
            delete(v3, 4294967295)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
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
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
