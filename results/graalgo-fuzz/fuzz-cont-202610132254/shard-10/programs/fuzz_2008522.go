package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint16; F1 int32 }

func main() {
    var v1 int = 0
    v2 := make(map[int]bool)
    var v3 bool = false
    v4 := v2[v1]
    v5 := GenS1{F0: -1}
    v6 := v2[(^(v1))]
    _ = len(v2)
    var v7 string = "foo bar"
    if (1e10 >= 3.14159) {
        v8 := GenS2{F0: 956, F1: v7}
        if true {
            v9 := []int{42, v1, (v1 ^ (v1 / ((v1) | 1)))}
            var v10 bool = (((((float64(((1e10 * float32(((0.1 + 1e10) + (0.1 / 1e10)))) - (((-2.25 * float32(-2.25)) - 0.1) - 0.1))) * 3.14159) != 1.5) || (78 == 42)) || (!v3)) || v6)
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 10; fi11++ {
                if v6 {
                    v12 := GenS2{F0: 9223372036854775807, F1: "\n\t"}
                    _ = v12
                }
                v13 := []int{}
                _ = fi11
                _ = v13
            }
        }
        fmt.Println(v7)
        _ = v8
    }
    v14 := GenS1{F0: 114}
    if (!(!(v4 || (!false)))) {
        var v15 uint64 = 42
        if true {
            if true {
                var v16 uint32 = 4294967295
                v16 += (^(v16))
                _ = v16
            }
            for fi17 := 0; fi17 < 4; fi17++ {
                for fi18 := 0; fi18 < 15; fi18++ {
                    var v19 int64 = 1
                    var v20 uint16 = 565
                    v4 = (127 > 1)
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                var v21 bool = v3
                _ = fi17
                _ = v21
            }
        }
        fmt.Println(v4, v3)
        _ = v15
    } else {
        v4 = v3
    }
    if v4 {
        v22 := GenS3{F0: 0, F1: 2147483647}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 2; fi23++ {
            v7 = v7
            v6 = v3
            _ = fi23
        }
        v3 = ((1e10 / (0.1 + float32(3.14159))) < (1.5 - float32(0.1)))
    }
    _ = len(v2)
    var v24 uint8 = 1
    var v25 int64 = 1
    v2[(v1 + v1)] = v3
    for fi26 := 0; fi26 < 8; fi26++ {
        switch fi26 {
        case 9223372036854775807:
            switch (v1 + (fi26 - fi26)) {
            case -964:
                var v27 uint32 = 17
                _ = v27
            case -2147483648:
                if v6 {
                    fmt.Println(v24)
                    v7 = v7
                    var v28 float64 = 3.14159
                    _ = v28
                } else {
                    v1 = v1
                    v29 := []int{(fi26 | -979), v1, 211, 4294967295}
                    for fi30 := 0; fi30 < 18; fi30++ {
                        var v31 int8 = 42
                        var v32 int64 = (v25 - int64(v24))
                        if len(v29) >= 2 {
                            v29 = v29[1:len(v29)]
                        }
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    _ = v29
                }
            case 32768:
                v33 := GenS2{F0: (v25 | -850), F1: v7}
                _ = v33
            }
        default:
            if (1 == -61) {
                fmt.Println(v7)
                switch v1 {
                case 4294967295:
                    v7 = v7
                case -9223372036854775808:
                    v34 := []int{}
                    _ = v34
                }
            }
        }
        var v35 float32 = -2.25
        var v36 float32 = v35
        _ = fi26
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
