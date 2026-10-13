package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint; F1 int16 }

func main() {
    var v1 uint32 = 2
    var v2 bool = (((210 >= -587) && ((("hello" + "") + ("foo bar" + "hello")) != "")) || ("a" > "a"))
    var v3 bool = (!v2)
    var v4 uint64 = 4294967295
    var v5 int16 = 42
    var v6 int64 = -9223372036854775808
    v1 = v1
    fmt.Println(v3)
    var v7 int = 9223372036854775807
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: 185, F1: -1, F2: 209}
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int16 = v5
    var v12 string = "a"
    if (!false) {
        fmt.Println(v11)
        switch int(v1) {
        case 139:
            fmt.Println(v4)
        case 363:
            v13 := GenS1{F0: v4, F1: -2147483648, F2: (^(v4))}
            _ = v13
        case 32767:
            v11 = (v11 * v5)
        }
        v4 -= v4
    } else {
        v3 = ((((0.1 - (float64(1.5) + 1.5)) / 0.1) == (((0.0 - (-2.25 * (1e10 - 1.5))) - -2.25) - (((1.5 + 0.0) / 0.1) * 1.5))) || v2)
        for fi14 := 0; fi14 < 5; fi14++ {
            if v2 {
                v15 := GenS2{F0: "", F1: -884, F2: 199}
                if v3 {
                    v11 = (v11 ^ (v11 - (v11 << ((190) & 63))))
                } else {
                    var v16 uint8 = 0
                    if v3 {
                        for fi17 := 0; fi17 < 8; fi17++ {
                            var v18 int64 = (-(v6))
                            v19 := GenS3{F0: 0, F1: v5}
                            switch fi17 {
                            case 32767:
                                fmt.Println(v2, fi17)
                            }
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        v16 = 179
                        v20 := []int{(v7 + (^(fi14)))}
                        _ = v20
                    }
                    var v21 float64 = -2.25
                    _ = v16
                    _ = v21
                }
                _ = v15
            } else {
                if false {
                    var v22 int8 = 1
                    _ = v22
                }
                for fi23 := 0; fi23 < 13; fi23++ {
                    switch v12 {
                    case "a":
                        var v24 string = v12
                        _ = v24
                    }
                    var v25 int16 = -32768
                    v3 = true
                    var v26 int64 = -9223372036854775808
                    _ = fi23
                    _ = v25
                    _ = v26
                }
            }
            if (2 == uint8(v1)) {
                v6 = int64(v5)
                v2 = v3
            } else {
                var v27 uint8 = 0
                v28 := []int{(^(v7)), (v7 >> ((0) & 63)), 4294967295}
                _ = v27
                _ = v28
            }
            _ = fi14
        }
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        fmt.Println(v6, v7, v3)
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
