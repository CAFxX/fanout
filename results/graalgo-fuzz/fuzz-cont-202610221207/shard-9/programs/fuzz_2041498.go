package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 bool = (true && (!false))
    var v2 uint64 = 0
    var v3 int8 = 1
    v4 := make(map[string]uint32)
    var v5 uint32 = 16
    for fi6 := 0; fi6 < 10; fi6++ {
        v2 = (^(v2))
        _ = fi6
    }
    v7 := GenS1{F0: 1}
    v8 := GenS2{F0: 65535, F1: 127}
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: 255}
        var v11 int = (fi9 * fi9)
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int = 128
    v13 := GenS3{F0: -88}
    v5 = 602
    var v14 uint64 = (^(v2))
    switch "a" {
    case "a":
        v15 := GenS3{F0: v3}
        _ = v15
    }
    fmt.Println(v2)
    if (!(true && (int16(v12) != -879))) {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 float64 = (0.0 - (((1.5 + (3.14159 / (3.14159 * (0.1 / (((1.5 + ((0.0 / float64(3.14159)) / 1.0)) - 0.0) - 0.1))))) - float64(1e10)) - 0.1))
            switch (v12 % (((^(fi16))) | 1)) {
            case 0:
                if (int32(v5) >= int32(fi16)) {
                    var v18 bool = true
                    delete(v4, "a")
                    fmt.Println(v2, v17)
                    _ = v18
                }
            case -9223372036854775808:
                var v19 uint8 = 68
                _ = v19
            default:
                v20 := []int{-440, (v12 >> ((18446744073709551615) & 63)), (fi16 / ((-2147483648) | 1))}
                _ = v20
            }
            switch "foo bar" {
            case "hello":
                if v1 {
                    v21 := []int{}
                    fmt.Println(v5, v3)
                    _ = v21
                }
            case "":
                var v22 string = ("a" + "")
                _ = v22
            default:
                v14 += 18446744073709551615
            }
            _ = fi16
            _ = v17
        }
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        if v1 {
            v24 := GenS1{F0: 32767}
            for fi25 := 0; fi25 < 2; fi25++ {
                fmt.Println(v14, fi23)
                _ = fi25
            }
            _ = v24
        }
        _ = fi23
    }
    if (v1 && false) {
        v14 = v14
        if (0 != 42) {
            v26 := GenS2{F0: 18446744073709551615, F1: 238}
            _ = v26
        } else {
            for fi27 := 0; fi27 < 7; fi27++ {
                switch 127 {
                case -109:
                    var v28 float64 = (0.1 / 0.1)
                    _ = v28
                case -280:
                    for fi29 := 0; fi29 < 14; fi29++ {
                        v14 += 18446744073709551615
                        v30 := []int{(fi27 ^ fi27)}
                        _ = fi29
                        _ = v30
                    }
                }
                _ = fi27
            }
        }
        v31 := []int{v12, (^(v12)), (v12 + v12), v12}
        _ = v31
    } else {
        if true {
            for fi32 := 0; fi32 < 19; fi32++ {
                if ((v5 % (((^(v5))) | 1)) > (v5 / ((1) | 1))) {
                    var v33 int8 = (v3 - (-(v3)))
                    fmt.Println(v1, v14, v12)
                    _ = v33
                }
                var v34 int = fi32
                _ = fi32
                _ = v34
            }
            var v35 float32 = ((1.5 * (1e10 + (0.1 / 1e10))) - 3.14159)
            _ = v35
        } else {
            fmt.Println(v1, v2, v5)
            v12 = v12
        }
        fmt.Println(v1, v1)
    }
    v14 += 2
    fmt.Println(v5, v3, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
