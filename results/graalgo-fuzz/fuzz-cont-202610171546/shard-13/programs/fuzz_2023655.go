package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint64, p1 string, p2 int8) bool {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 float32 = float32(3.14159)
        v6 := GenS1{F0: 1}
        v7 := []int{fi4, 127}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    if (42 < 255) {
        var v8 int64 = 9223372036854775807
        v9 := []int{128}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return true
}
func genFn1(p0 string) uint {
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int8 = 100
        v11 -= v11
        _ = fi10
        _ = v11
    }
    v12 := []int{-2147483648, 9223372036854775807}
    p0 = p0
    _ = p0
    _ = v12
    return 272
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    var v3 string = "hello"
    if false {
        v3 = v3
        v3 = v3
    }
    v3 = v3
    v13 := GenS1{F0: 0}
    if false {
        for fi14 := 0; fi14 < 4; fi14++ {
            v3 = ((v3 + v3) + v3)
            _ = fi14
        }
        v15 := GenS1{F0: -834}
        if (!false) {
            if false {
                v3 = v3
            } else {
                var v16 float64 = 0.0
                v3 = (v3 + v3)
                fmt.Println(v3, v16)
                _ = v16
            }
        }
        _ = v15
    } else {
        v3 = v3
    }
    v3 = v3
    if false {
        switch ("\n\t" + v3) {
        case "foo bar":
            v17 := []int{497, 2147483647}
            _ = v17
        default:
            v3 = ("hello" + "foo bar")
        }
        fmt.Println(v3)
    } else {
        v18 := GenS1{F0: 0}
        var v19 uint32 = 65535
        _ = v18
        _ = v19
    }
    fmt.Println(v3)
    v20 := GenS1{F0: 259}
    v21 := GenS1{F0: 42}
    if ((!((((float32(3.14159) - 1.5) - (1.5 - 0.0)) * 1e10) <= 0.1)) || true) {
        v3 = (("" + (("foo bar" + ((v3 + v3) + v3)) + (v3 + v3))) + ((v3 + "\n\t") + v3))
        if (false || ((127 > 32768) && (float32(0.0) == 1e10))) {
            v22 := GenS1{F0: 0}
            switch "a" {
            case "":
                v3 = (v3 + v3)
            case "\n\t":
                fmt.Println(v3)
            case "hello":
                v3 = v3
            default:
                if (255 != 134) {
                    var v23 int32 = -1
                    for fi24 := 0; fi24 < 9; fi24++ {
                        var v25 uint64 = 2
                        v26 := v1[""]
                        v27 := GenS1{F0: (-(v23))}
                        _ = fi24
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    v3 = v3
                    _ = v23
                } else {
                    v28 := GenS2{F0: -753, F1: 184}
                    v29 := []int{1, 127, 9223372036854775807, 42}
                    _ = v28
                    _ = v29
                }
            }
            v3 = ("hello" + ("a" + (v3 + ("foo bar" + "foo bar"))))
            _ = v22
        }
        for fi30 := 0; fi30 < 3; fi30++ {
            for fi31 := 0; fi31 < 4; fi31++ {
                if true {
                    fmt.Println(fi30, fi31)
                    var v32 int32 = 42
                    _ = v32
                } else {
                    v33 := []int{fi30, 128, fi30}
                    v3 = "a"
                    v34 := GenS2{F0: -183, F1: 0}
                    _ = v33
                    _ = v34
                }
                v3 = v3
                switch (fi30 - 517) {
                case -2147483648:
                    var v35 uint32 = 2
                    _ = v35
                case -9223372036854775808:
                    v3 = (((v3 + "foo bar") + ("a" + v3)) + "")
                case 128:
                    var v36 int16 = -1
                    _ = v36
                default:
                    var v37 uint8 = 0
                    _ = v37
                }
                _ = fi31
            }
            _ = fi30
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
