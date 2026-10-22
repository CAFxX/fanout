package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 float32 = 0.0
    v2 := []uint32{}
    var v3 float64 = 0.0
    v4 := make(map[uint32]int)
    var v5 string = ("hello" + "\n\t")
    var v6 int = 127
    var v7 uint = 2
    v1 = (0.1 - (float32(v3) - ((v1 - (v1 / v1)) - ((0.0 + 3.14159) + v1))))
    if (!(42 <= 88)) {
        switch v6 {
        case 896:
            if (2147483647 != 0) {
                fmt.Println(v1, v3, v6)
                for fi8 := 0; fi8 < 18; fi8++ {
                    fmt.Println(v7, fi8, v3)
                    for fi9 := 0; fi9 < 3; fi9++ {
                        var v10 int16 = 0
                        var v11 int32 = 132
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    if true {
                        var v12 uint8 = 1
                        _ = v12
                    } else {
                        v5 = v5
                    }
                    _ = fi8
                }
            } else {
                var v13 int = (-(v6))
                delete(v4, 1)
                v14 := v4[42]
                _ = v13
                _ = v14
            }
        default:
            var v15 int32 = 42
            _ = v15
        }
    } else {
        v16 := GenS1{F0: 0}
        _ = v16
    }
    if false {
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 uint32 = 2147483648
            var v19 int64 = int64(v7)
            _ = fi17
            _ = v18
            _ = v19
        }
        v7 = (v7 << ((v7) & 63))
        fmt.Println(v1)
    }
    _ = len(v2)
    v3 += (3.14159 + 1.5)
    var v20 int16 = 294
    if false {
        for fi21 := 0; fi21 < 10; fi21++ {
            v20 = (-(v20))
            if true {
                for fi22 := 0; fi22 < 2; fi22++ {
                    if ((!(-1 != -9223372036854775808)) || (100 == 127)) {
                        if false {
                            fmt.Println(v7, v6)
                        }
                        if true {
                            v3 -= -2.25
                            var v23 uint = v7
                            switch ((v5 + v5) + ((v5 + (v5 + (("foo bar" + (v5 + v5)) + (("\n\t" + "") + (("" + (("hello" + (v5 + (v5 + (v5 + ("hello" + ""))))) + v5)) + ((v5 + "foo bar") + v5)))))) + v5)) {
                            case "":
                                var v24 uint32 = 671
                                _ = v24
                            case "a":
                                fmt.Println(v20)
                            default:
                                var v25 int = 1
                                _ = v25
                            }
                            _ = v23
                        }
                    } else {
                        v20 = (^(v20))
                    }
                    var v26 uint32 = 1
                    _ = fi22
                    _ = v26
                }
                switch v5 {
                case "\n\t":
                    if false {
                        v3 -= 1e10
                    } else {
                        v7 += v7
                        if false {
                            v4[0] = (v6 << ((255) & 63))
                            if (false || (true || true)) {
                                fmt.Println(v6, v6, v20)
                            }
                            for fi27 := 0; fi27 < 13; fi27++ {
                                fmt.Println(v20, v5)
                                v3 += ((1.5 - 3.14159) - v3)
                                v28 := GenS2{F0: 9223372036854775807}
                                var v29 float64 = float64(v1)
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                        }
                        v5 = v5
                    }
                default:
                    v7 = v7
                }
                var v30 bool = (v3 < 1.5)
                _ = v30
            }
            _ = fi21
        }
    } else {
        _ = len(v2)
    }
    var v31 uint64 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
