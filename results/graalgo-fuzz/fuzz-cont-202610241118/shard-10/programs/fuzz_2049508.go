package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 int16; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint; F1 bool }

func main() {
    var v1 int64 = -710
    v2 := make(map[string]uint32)
    var v3 int = -234
    var v4 int8 = 1
    var v5 string = ((("foo bar" + "") + ("\n\t" + "")) + (("\n\t" + "\n\t") + ("foo bar" + "")))
    var v6 float32 = 0.0
    fmt.Println(v6, v6, v6)
    for fi7 := 0; fi7 < 10; fi7++ {
        v6 -= v6
        v5 = v5
        v8 := []int{(fi7 | 128), -635}
        _ = fi7
        _ = v8
    }
    fmt.Println(v6, v6, v5)
    v4 = 1
    v4 += 100
    var v9 int8 = 1
    v10 := []int{(v3 % ((v3) | 1))}
    var v11 int = (v3 * (^(v3)))
    v9 += v9
    var v12 int = -1
    v4 = -128
    var v13 int8 = (v9 / (((v4 ^ 49)) | 1))
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int8 = (v9 % ((0) | 1))
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int64 = (v1 % (((^(v1))) | 1))
            delete(v2, (((v5 + (v5 + (v5 + v5))) + (v5 + ((("foo bar" + "\n\t") + (v5 + v5)) + "a"))) + v5))
            v10 = append(v10, -658)
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    v5 = (("" + "foo bar") + "")
    if ((1 >= -1) || ((v6 - float32(-2.25)) == -2.25)) {
        var v18 float64 = float64(v6)
        if ((v1 >= v1) && true) {
            fmt.Println(v11, v11, v5)
            switch v5 {
            case "a":
                if (163 < 4294967295) {
                    for fi19 := 0; fi19 < 3; fi19++ {
                        v20 := GenS1{F0: 0, F1: 2147483647, F2: 0}
                        for fi21 := 0; fi21 < 18; fi21++ {
                            v22 := GenS2{F0: "", F1: -1, F2: 0}
                            fmt.Println(fi21, v12)
                            _ = fi21
                            _ = v22
                        }
                        _ = fi19
                        _ = v20
                    }
                    for fi23 := 0; fi23 < 13; fi23++ {
                        var v24 uint = 18446744073709551615
                        fmt.Println(v5)
                        for fi25 := 0; fi25 < 16; fi25++ {
                            switch (fi25 | 4294967295) {
                            case 42:
                                v4 -= (v4 + v13)
                            }
                            _ = fi25
                        }
                        _ = fi23
                        _ = v24
                    }
                    switch v3 {
                    case 32767:
                        v18 -= (v18 - 0.0)
                    case -624:
                        var v26 int8 = (v4 ^ (v9 ^ v9))
                        _ = v26
                    }
                } else {
                    var v27 float32 = (v6 - ((v6 + v6) + v6))
                    v28 := GenS3{F0: 1, F1: 42, F2: 929}
                    _ = v27
                    _ = v28
                }
            default:
                fmt.Println(v1, v12, v6)
            }
        }
        fmt.Println(v9, v13, v6)
        _ = v18
    } else {
        v29 := GenS4{F0: 396, F1: false}
        v9 = 127
        _ = len(v2)
        _ = v29
    }
    var v30 bool = true
    var v31 uint16 = 368
    v32 := GenS3{F0: (^(v31)), F1: 4294967295, F2: v31}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
