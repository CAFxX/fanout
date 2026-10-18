package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint32; F2 int8 }
type GenS4 struct { F0 int32; F1 int8 }

func main() {
    var v1 float64 = 0.1
    var v2 bool = false
    v3 := []string{}
    var v4 bool = false
    if (!(-290 >= 127)) {
        _ = len(v3)
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 bool = v4
            var v7 uint = 42
            v8 := GenS1{F0: 0}
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    } else {
        fmt.Println(v2)
        switch 808 {
        case -803:
            v4 = v4
        }
        if (((0.1 + -2.25) < (0.0 - (0.0 + ((0.0 / (-2.25 + 0.1)) - -2.25)))) && (42 < 255)) {
            var v9 uint = 255
            _ = v9
        } else {
            fmt.Println(v2, v1, v4)
            switch "foo bar" {
            case "":
                if (!true) {
                    _ = len(v3)
                    if (42 != 42) {
                        fmt.Println(v4, v4)
                    } else {
                        v1 = 0.1
                        v2 = v4
                        var v10 float32 = float32(((v1 - ((1e10 / (1e10 / (((0.0 - 3.14159) + (((v1 * (((-2.25 * (((((float64((1.5 - 1.5)) * (v1 - (0.0 + (v1 + (v1 + v1))))) * 1e10) + (v1 * float64((3.14159 * (-2.25 / -2.25))))) / v1) + v1)) - (3.14159 * v1)) / (1.5 * 1e10))) - 3.14159) - v1)) / (float64(3.14159) - ((v1 * -2.25) / 1.0))))) * (v1 * (3.14159 / v1)))) * float64(-2.25)))
                        _ = v10
                    }
                    var v11 float32 = ((0.1 + 0.0) / 0.1)
                    _ = v11
                }
            case "a":
                switch -2147483648 {
                case 127:
                    var v12 int = 4294967295
                    _ = v12
                case -1:
                    v13 := GenS2{F0: 42}
                    _ = v13
                default:
                    fmt.Println(v1)
                }
            case "hello":
                v2 = (!(v2 || (0 <= 1)))
            }
        }
    }
    fmt.Println(v4, v4, v2)
    if v2 {
        if len(v3) > 0 {
            v3[0] = ""
        }
        var v14 uint = 65535
        _ = v14
    }
    var v15 int32 = 42
    if true {
        v1 -= v1
        for fi16 := 0; fi16 < 15; fi16++ {
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 bool = v2
                fmt.Println(v4)
                v19 := GenS3{F0: v15, F1: 65535, F2: int8(v15)}
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 uint64 = 397
            if (42 == 0) {
                v20 -= v20
                if len(v3) > 0 {
                    v3[0] = ("foo bar" + "foo bar")
                }
                var v21 int32 = (v15 | (v15 >> ((689) & 63)))
                _ = v21
            } else {
                var v22 float64 = v1
                v4 = true
                _ = v22
            }
            _ = fi16
            _ = v20
        }
    } else {
        var v23 int = 128
        if len(v3) > 0 {
            v3[0] = ("foo bar" + "foo bar")
        }
        fmt.Println(v2)
        _ = v23
    }
    var v24 int8 = 100
    v1 = (v1 + v1)
    switch 32767 {
    case 128:
        var v25 float32 = 3.14159
        _ = v25
    case -658:
        for fi26 := 0; fi26 < 18; fi26++ {
            switch fi26 {
            case 615:
                if v2 {
                    var v27 int64 = 42
                    for fi28 := 0; fi28 < 10; fi28++ {
                        if (v1 < float64(0.0)) {
                            v29 := GenS4{F0: 789, F1: (v24 / (((-(v24))) | 1))}
                            fmt.Println(v2, fi26)
                            fmt.Println(v15, v24)
                            _ = v29
                        } else {
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                        }
                        v4 = true
                        var v30 uint = 905
                        _ = fi28
                        _ = v30
                    }
                    _ = v27
                }
            }
            _ = fi26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
