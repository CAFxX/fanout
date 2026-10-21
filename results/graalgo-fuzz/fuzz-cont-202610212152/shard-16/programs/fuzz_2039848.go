package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 string = ("foo bar" + "")
    var v2 uint8 = 42
    var v3 int16 = -1
    var v4 bool = (0 == -9223372036854775808)
    if v4 {
        v5 := GenS1{F0: 1}
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 int = fi6
            v8 := []int{}
            var v9 int32 = 42
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        switch 2147483647 {
        case -9223372036854775808:
            if (v1 == (v1 + ("a" + (((v1 + v1) + "\n\t") + "hello")))) {
                var v10 uint32 = 160
                _ = v10
            } else {
                var v11 uint = 409
                v12 := GenS2{F0: int64(v11), F1: (v3 << (((v11 - v11)) & 63))}
                switch 9223372036854775807 {
                case 153:
                    fmt.Println(v2, v1, v4)
                case 127:
                    v13 := []int{-650}
                    _ = v13
                case -1:
                    v3 = (v3 ^ (-(v3)))
                }
                _ = v11
                _ = v12
            }
        case 2147483647:
            var v14 int16 = v3
            _ = v14
        case 1:
            var v15 uint64 = 79
            _ = v15
        }
        _ = v5
    } else {
        switch 32767 {
        case 2147483647:
            v1 = ""
        default:
            for fi16 := 0; fi16 < 16; fi16++ {
                if (32767 != 0) {
                    for fi17 := 0; fi17 < 11; fi17++ {
                        v18 := GenS2{F0: -1, F1: 377}
                        var v19 int8 = 42
                        var v20 string = v1
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                }
                _ = fi16
            }
        }
        v1 = "\n\t"
    }
    if v4 {
        if ("foo bar" <= v1) {
            for fi21 := 0; fi21 < 2; fi21++ {
                fmt.Println(v2, v4)
                _ = fi21
            }
            var v22 int64 = -9223372036854775808
            v23 := []int{877, -507, 9223372036854775807, 1}
            _ = v22
            _ = v23
        } else {
            for fi24 := 0; fi24 < 12; fi24++ {
                v25 := []int{}
                _ = fi24
                _ = v25
            }
            switch v1 {
            case "a":
                var v26 int64 = 42
                _ = v26
            }
        }
        var v27 float64 = 3.14159
        _ = v27
    } else {
        v4 = (1 != 0)
    }
    var v28 uint16 = 32767
    var v29 int = -9223372036854775808
    for fi30 := 0; fi30 < 14; fi30++ {
        v4 = (-363 == (v3 / ((v3) | 1)))
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 uint16 = (v28 - (v28 / (((^(v28))) | 1)))
            v33 := []int{fi30}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    switch (v1 + v1) {
    case "a":
        v34 := GenS3{F0: 517}
        _ = v34
    default:
        var v35 int64 = 1
        _ = v35
    }
    v36 := []int{128, 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
}
