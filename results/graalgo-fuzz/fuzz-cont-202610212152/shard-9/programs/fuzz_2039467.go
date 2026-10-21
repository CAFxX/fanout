package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int64 = -550
    var v2 string = ("hello" + "")
    var v3 string = v2
    var v4 bool = ((1.5 + 1.5) < -2.25)
    v2 = ("hello" + ("hello" + v3))
    var v5 uint8 = 128
    v6 := []int{}
    var v7 uint8 = v5
    v7 += v7
    if (v4 || v4) {
        var v8 int = 128
        _ = v8
    }
    if (!(255 < 255)) {
        var v9 int64 = 1
        _ = v9
    } else {
        switch -2147483648 {
        case 32767:
            v1 = v1
        case 9223372036854775807:
            var v10 float32 = 1.5
            _ = v10
        default:
            fmt.Println(v7, v2)
        }
        v11 := GenS1{F0: ((v3 + v2) + v3), F1: 0}
        _ = v11
    }
    fmt.Println(v1, v4)
    fmt.Println(v5, v3, v4)
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 string = v3
        _ = fi12
        _ = v13
    }
    fmt.Println(v3)
    v6 = append(v6, -9223372036854775808)
    var v14 float64 = -2.25
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            v6 = append(v6, fi16)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi16
        }
        _ = fi15
    }
    v6 = append(v6, 0)
    if (35 < 0) {
        for fi17 := 0; fi17 < 12; fi17++ {
            if ((!(3.14159 >= -2.25)) || true) {
                for fi18 := 0; fi18 < 7; fi18++ {
                    var v19 string = v2
                    if (1 >= 2147483647) {
                        v20 := GenS1{F0: "", F1: 42}
                        var v21 int16 = 32767
                        _ = v20
                        _ = v21
                    }
                    _ = fi18
                    _ = v19
                }
                var v22 int32 = -583
                v1 -= v1
                _ = v22
            } else {
                switch fi17 {
                case 814:
                    for fi23 := 0; fi23 < 1; fi23++ {
                        switch fi23 {
                        case 32768:
                            v24 := GenS1{F0: v2, F1: 1000}
                            _ = v24
                        default:
                            var v25 int = fi23
                            v2 = v2
                            v26 := GenS1{F0: v3, F1: 1}
                            _ = v25
                            _ = v26
                        }
                        var v27 float64 = (v14 * v14)
                        if (42 >= 4294967295) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if true {
                                if (1e10 > 1e10) {
                                    v28 := GenS2{F0: 42}
                                    if len(v6) > 0 {
                                        v6[0] = fi23
                                    }
                                    _ = v28
                                }
                            }
                        }
                        _ = fi23
                        _ = v27
                    }
                }
            }
            _ = fi17
        }
        v1 = (^(v1))
    }
    fmt.Println(v14)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
