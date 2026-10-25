package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 int = 4294967295
    var v2 bool = (2147483647 != -1)
    var v3 uint64 = 1
    var v4 uint32 = 322
    var v5 uint64 = (^(v3))
    var v6 uint64 = (v3 & (v5 + 18446744073709551615))
    fmt.Println(v3)
    v7 := GenS1{F0: 161}
    for fi8 := 0; fi8 < 3; fi8++ {
        if false {
            var v9 uint = 734
            _ = v9
        } else {
            var v10 uint32 = 42
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS1{F0: 2}
                if true {
                    v13 := []int{fi8, -1}
                    v4 += (v10 << ((2) & 63))
                    if (!(1 > 65535)) {
                        var v14 int64 = 42
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = v14
                    }
                    _ = v13
                } else {
                    var v15 float64 = -2.25
                    if (false || v2) {
                        var v16 float32 = (0.1 * (float32(0.1) / ((1.5 / (-2.25 + 1.5)) / float32(1.5))))
                        v17 := []int{}
                        if v2 {
                            fmt.Println(v2, v1, v2)
                            switch 0 {
                            case 128:
                                switch (v1 | fi8) {
                                case 877:
                                    v18 := GenS1{F0: 255}
                                    _ = v18
                                case 1:
                                    var v19 int64 = 0
                                    _ = v19
                                case -832:
                                    if len(v17) > 0 {
                                        v17[0] = v1
                                    }
                                default:
                                    if (1 < 903) {
                                        switch ("hello" + "") {
                                        case "foo bar":
                                            _ = len(v17)
                                        }
                                    }
                                }
                            }
                            if v2 {
                                var v20 uint8 = 174
                                _ = v20
                            }
                        }
                        _ = v16
                        _ = v17
                    } else {
                        if v2 {
                            v21 := []int{}
                            v21 = append(v21, 9223372036854775807)
                            _ = v21
                        }
                        var v22 uint64 = v6
                        _ = v22
                    }
                    v23 := []int{v1, 4294967295, (v1 << ((255) & 63)), (^(fi11))}
                    _ = v15
                    _ = v23
                }
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        _ = fi8
    }
    var v24 float64 = 3.14159
    for fi25 := 0; fi25 < 9; fi25++ {
        v26 := GenS1{F0: 558}
        for fi27 := 0; fi27 < 13; fi27++ {
            v28 := []int{32768, fi27}
            v28 = append(v28, -568)
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        v2 = (255 != 1)
        v2 = (1 >= 0)
        _ = fi29
    }
    v30 := []int{}
    var v31 int32 = 472
    var v32 int16 = 42
    for fi33 := 0; fi33 < 1; fi33++ {
        v34 := GenS2{F0: -304, F1: 255}
        _ = fi33
        _ = v34
    }
    _ = len(v30)
    v6 = v3
    var v35 int32 = v31
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
