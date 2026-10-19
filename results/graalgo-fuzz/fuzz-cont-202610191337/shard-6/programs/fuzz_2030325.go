package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint64; F2 bool }
type GenS3 struct { F0 uint; F1 uint64 }

func main() {
    v1 := []uint32{65535, 0, 1, 42}
    var v2 string = "a"
    var v3 uint = 2
    v3 = v3
    fmt.Println(v3, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v4 float64 = float64(1.5)
    var v5 bool = true
    v6 := GenS1{F0: 127}
    var v7 string = v2
    switch 339 {
    case -958:
        var v8 float32 = (3.14159 * float32(v4))
        _ = v8
    default:
        switch "" {
        case "foo bar":
            if (!true) {
                fmt.Println(v3, v2, v7)
                if v5 {
                    var v9 uint32 = 2147483648
                    var v10 float32 = -2.25
                    _ = v10
                    _ = v9
                } else {
                    if v5 {
                        v2 = "foo bar"
                        if false {
                            fmt.Println(v2, v2)
                            fmt.Println(v5, v3, v3)
                            fmt.Println(v7, v7, v3)
                        } else {
                            v4 += ((v4 + (v4 + v4)) - 0.1)
                            v3 = 0
                        }
                        for fi11 := 0; fi11 < 6; fi11++ {
                            if v5 {
                                if v5 {
                                    v3 = 689
                                } else {
                                    v2 = v2
                                }
                                var v12 bool = (!(1 != 109))
                                _ = v12
                            } else {
                                for fi13 := 0; fi13 < 6; fi13++ {
                                    if (!(127 != 255)) {
                                        var v14 uint32 = 2147483647
                                        _ = len(v1)
                                        if v5 {
                                            if false {
                                                fmt.Println(v4)
                                                var v15 int16 = 332
                                                v14 -= v14
                                                _ = v15
                                            } else {
                                                v16 := GenS1{F0: 255}
                                                var v17 uint8 = 0
                                                switch (fi13 >> ((v3) & 63)) {
                                                case 128:
                                                    if (((3.14159 * (0.0 * -2.25)) / 1.0) >= 0.0) {
                                                        v5 = ((!(v5 && (v5 || v5))) || ((v5 || (!(v5 || (-128 <= -58)))) && (579 != 577)))
                                                    }
                                                }
                                                var v18 int64 = -1
                                                _ = v16
                                                _ = v17
                                                _ = v18
                                            }
                                            v14 = 583
                                        } else {
                                            v19 := GenS2{F0: -9223372036854775808, F1: 2, F2: true}
                                            _ = v19
                                        }
                                        var v20 int = (fi13 >> (((^(v3))) & 63))
                                        switch v20 {
                                        case -327:
                                            var v21 bool = v5
                                            v5 = true
                                            _ = v21
                                        case 0:
                                            v14 += (^(v14))
                                        case 9223372036854775807:
                                            if ((-2.25 * 0.1) <= float32(v4)) {
                                                v22 := GenS2{F0: 42, F1: 4294967295, F2: v5}
                                                var v23 string = v2
                                                _ = v22
                                                _ = v23
                                            }
                                        }
                                        if v5 {
                                            v24 := GenS3{F0: v3, F1: 48}
                                            _ = v24
                                        }
                                        if v5 {
                                            var v25 int = fi11
                                            _ = v25
                                        }
                                        _ = v14
                                        _ = v20
                                    }
                                    _ = fi13
                                }
                                var v26 uint8 = 42
                                for fi27 := 0; fi27 < 18; fi27++ {
                                    v5 = false
                                    _ = fi27
                                }
                                _ = v26
                            }
                            fmt.Println(v5, v7)
                            _ = fi11
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v5 = (v5 || v5)
            } else {
                v28 := GenS2{F0: 946, F1: 0, F2: true}
                _ = v28
            }
        default:
            var v29 int64 = -9223372036854775808
            _ = v29
        }
    }
    if v5 {
        for fi30 := 0; fi30 < 5; fi30++ {
            if len(v1) > 0 {
                v1[0] = uint32(v3)
            }
            for fi31 := 0; fi31 < 16; fi31++ {
                v32 := GenS1{F0: 42}
                _ = fi31
                _ = v32
            }
            _ = fi30
        }
        var v33 int32 = -1
        _ = v33
    }
    fmt.Println(v7, v2)
    if ("hello" <= v7) {
        var v34 uint64 = 4294967295
        _ = v34
    }
    var v35 float64 = 0.0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
