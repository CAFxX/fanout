package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    var v1 uint8 = 41
    var v2 int16 = -1
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    var v5 string = "a"
    var v6 uint16 = 763
    fmt.Println(v1, v1, v2)
    for fi7 := 0; fi7 < 12; fi7++ {
        if (1 <= 972) {
            v8 := GenS1{F0: v1}
            var v9 uint64 = 4294967295
            switch 0 {
            case 128:
                var v10 int32 = int32(v2)
                _ = v10
            case 32767:
                v6 -= (v6 - (v6 & 255))
            case 652:
                switch 4294967295 {
                case 32768:
                    v5 = (v5 + ("hello" + v5))
                case 1:
                    if ((!(((128 != v1) || (((0 <= 2147483648) && (!false)) && false)) || (701 != 2147483648))) && false) {
                        if (!false) {
                            if false {
                                v5 = v5
                                fmt.Println(v2, v2, v1)
                                var v11 int32 = 0
                                _ = v11
                            } else {
                                if (!false) {
                                    var v12 int32 = -1
                                    var v13 float64 = 1e10
                                    v14 := GenS1{F0: (^(v1))}
                                    _ = v12
                                    _ = v13
                                    _ = v14
                                } else {
                                    var v15 int64 = -1
                                    _ = v15
                                }
                                var v16 int8 = 0
                                if (!((((v5 + v5) + "a") >= (("hello" + v5) + ("hello" + ("hello" + v5)))) && (int16(v9) < v2))) {
                                    var v17 int = fi7
                                    _ = v17
                                } else {
                                    var v18 uint64 = v9
                                    v16 = (^(v16))
                                    var v19 float32 = (0.1 + 3.14159)
                                    _ = v18
                                    _ = v19
                                }
                                _ = v16
                            }
                            _ = len(v3)
                            var v20 int64 = 42
                            _ = v20
                        }
                    }
                }
            }
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    var v21 bool = true
    if (0.1 >= (((0.1 * 0.0) - 0.1) * 0.1)) {
        for fi22 := 0; fi22 < 5; fi22++ {
            for fi23 := 0; fi23 < 6; fi23++ {
                switch (fi22 % ((fi23) | 1)) {
                case 32768:
                    v1 = v1
                case 2147483647:
                    v24 := GenS2{F0: -2147483648, F1: false}
                    _ = v24
                default:
                    v25 := GenS2{F0: -990, F1: (uint64(v1) > 42)}
                    _ = v25
                }
                var v26 uint64 = 701
                switch v5 {
                case "":
                    v27 := GenS1{F0: (^(v1))}
                    _ = v27
                case "foo bar":
                    v1 += (v1 / ((15) | 1))
                }
                _ = fi23
                _ = v26
            }
            v28 := []int{fi22, 32768}
            _ = fi22
            _ = v28
        }
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 uint8 = 2
            _ = fi29
            _ = v30
        }
        switch v5 {
        case "a":
            var v31 float64 = 1e10
            _ = v31
        case "":
            v32 := []int{749}
            _ = v32
        default:
            var v33 int = -659
            _ = v33
        }
    }
    _ = len(v4)
    v34 := []int{32767, 127}
    var v35 float64 = 3.14159
    v36 := v3[v5]
    if len(v34) > 0 {
        v34[0] = -9223372036854775808
    }
    var v37 int64 = 710
    v37 = 584
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
