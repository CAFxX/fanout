package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint8; F1 int64; F2 uint64 }
func genFn0(p0 uint, p1 int16) int64 {
    switch ("hello" + ("hello" + "")) {
    case "":
        var v4 uint = p0
        _ = v4
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            var v5 int16 = p1
            _ = v5
        case "foo bar":
            var v6 uint32 = 65535
            _ = v6
        }
    case "hello":
        v7 := GenS1{F0: 1000}
        _ = v7
    }
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v8
    return 340
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (0 >= 369)
    var v3 float64 = 3.14159
    var v9 uint64 = 4294967295
    v10 := []int{}
    v3 = float64(1.5)
    var v11 int32 = -2147483648
    var v12 int = 128
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 string = ("\n\t" + "hello")
        var v15 uint16 = 0
        if true {
            fmt.Println(v14, v12, fi13)
            if (42 != 2) {
                fmt.Println(v14)
                v16 := GenS1{F0: 1000}
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 float32 = 0.1
                    _ = fi17
                    _ = v18
                }
                _ = v16
            } else {
                v1["\n\t"] = (v9 > 502)
                if v2 {
                    v11 = (v11 ^ 1)
                    for fi19 := 0; fi19 < 12; fi19++ {
                        fmt.Println(v11)
                        fmt.Println(fi19)
                        switch ((v14 + "foo bar") + v14) {
                        case "":
                            var v20 int32 = v11
                            _ = v20
                        case "hello":
                            var v21 string = v14
                            _ = v21
                        case "a":
                            v22 := GenS2{F0: v15, F1: v14}
                            _ = v22
                        }
                        _ = fi19
                    }
                }
                if len(v10) > 0 {
                    v10[0] = (fi13 | v12)
                }
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                if ((v12 & (-(fi13))) < fi13) {
                    v1["hello"] = v2
                    _ = len(v1)
                    fmt.Println(v14, v2, v15)
                }
                _ = fi23
            }
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    if v2 {
        v24 := GenS1{F0: int16(v11)}
        _ = v24
    }
    v25 := GenS3{F0: 128, F1: int64(v11), F2: (v9 & v9)}
    for fi26 := 0; fi26 < 3; fi26++ {
        fmt.Println(fi26, v11, fi26)
        _ = fi26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        if ((float64((1.5 * 0.1)) - ((float64(-2.25) / float64((((1.5 + 3.14159) * ((1.5 - ((float32(((((0.0 - v3) / 0.1) - (1.5 - (v3 - (v3 - 1.5)))) + float64((-2.25 - (1e10 - 0.0))))) - (0.1 / 3.14159)) / 0.1)) + (-2.25 / (-2.25 / 1e10)))) + 0.0))) * 3.14159)) >= v3) {
            v9 += 0
        } else {
            v2 = true
            for fi28 := 0; fi28 < 20; fi28++ {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v3 = 0.0
                _ = fi28
            }
        }
        if len(v10) > 0 {
            v10[0] = fi27
        }
        _ = fi27
    }
    v2 = v2
    var v29 uint32 = 1
    v10 = append(v10, v12)
    for fi30 := 0; fi30 < 8; fi30++ {
        if v2 {
            v3 = (float64(0.0) / (v3 * (1.5 + (0.0 + 1e10))))
        }
        _ = fi30
    }
    var v31 int16 = 32767
    fmt.Println(v2, v12, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v31
    _ = v9
}
