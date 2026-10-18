package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int64 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 string }

func main() {
    var v1 string = ""
    var v2 int8 = 127
    var v3 int32 = -1
    v4 := []int{}
    var v5 bool = true
    fmt.Println(v3)
    fmt.Println(v2, v5)
    var v6 int32 = (v3 | (v3 | 77))
    if len(v4) > 0 {
        v4[0] = 128
    }
    if (2 == 779) {
        fmt.Println(v1, v5, v6)
        var v7 int64 = 42
        v8 := GenS1{F0: (v6 - (v6 % ((v3) | 1))), F1: 811, F2: (-(v7))}
        _ = v7
        _ = v8
    }
    v4 = append(v4, 1)
    v5 = ((v1 + v1) > (v1 + "foo bar"))
    if ((!false) || (709 > 1)) {
        for fi9 := 0; fi9 < 4; fi9++ {
            v5 = v5
            _ = fi9
        }
        if true {
            v1 = (v1 + ((((v1 + (v1 + "a")) + ((("\n\t" + (v1 + "")) + v1) + "foo bar")) + "") + v1))
            if ((-(v3)) == 42) {
                fmt.Println(v3, v1, v6)
                var v10 int64 = 9223372036854775807
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v10
            } else {
                v11 := GenS2{F0: 984, F1: 42}
                for fi12 := 0; fi12 < 6; fi12++ {
                    fmt.Println(v3, v2, v5)
                    _ = fi12
                }
                _ = v11
            }
        } else {
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := GenS3{F0: ((v1 + v1) + ((v1 + v1) + v1))}
                for fi15 := 0; fi15 < 14; fi15++ {
                    if ((3.14159 < 0.1) || (v5 || ((69 >= 331) && ((v6 - 1) < v3)))) {
                        var v16 bool = true
                        _ = v16
                    } else {
                        var v17 uint64 = 42
                        _ = v17
                    }
                    v3 = v3
                    v1 = v1
                    _ = fi15
                }
                _ = fi13
                _ = v14
            }
            var v18 uint = 764
            _ = v18
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            fmt.Println(v5, v3)
        } else {
            v2 = (v2 ^ 100)
            var v19 float64 = (0.1 + (1.5 / (float64((0.0 - (((float32(0.1) + 0.0) * 3.14159) / 1e10))) / 1.5)))
            v1 = "a"
            _ = v19
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            fmt.Println(fi20, v6)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi20
        }
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        fmt.Println(v2, v2)
        switch v1 {
        case "hello":
            var v22 int8 = -60
            _ = v22
        case "":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi21
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v24 int16 = 878
        _ = fi23
        _ = v24
    }
    v1 = "a"
    fmt.Println(v6, v5)
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 9; fi26++ {
            v27 := GenS3{F0: "a"}
            _ = fi26
            _ = v27
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi25
    }
    v1 = v1
    fmt.Println(v3, v6, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
