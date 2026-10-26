package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int; F1 int8 }

func main() {
    var v1 uint64 = 935
    var v2 int = 0
    var v3 int64 = 0
    var v4 int16 = 32767
    for fi5 := 0; fi5 < 10; fi5++ {
        switch (v2 * fi5) {
        case -9223372036854775808:
            v6 := []int{fi5, fi5, 32767, (fi5 | fi5)}
            _ = v6
        case 23:
            v7 := []int{v2}
            _ = v7
        case -1:
            var v8 uint64 = 42
            _ = v8
        default:
            var v9 int64 = v3
            _ = v9
        }
        _ = fi5
    }
    fmt.Println(v1)
    v10 := []int{(^(v2)), (v2 % ((216) | 1)), v2, v2}
    v3 = (^(v3))
    _ = len(v10)
    var v11 uint64 = v1
    fmt.Println(v11, v11)
    var v12 uint32 = 1
    v1 -= v11
    if true {
        var v13 int64 = int64(v12)
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = v13
    } else {
        v14 := GenS1{F0: 2, F1: 1}
        var v15 float32 = 0.1
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            if len(v10) > 0 {
                v10[0] = (fi17 / ((-1) | 1))
            }
            if true {
                switch (fi16 ^ (fi16 % ((32767) | 1))) {
                case -1:
                    v11 = 18446744073709551615
                }
                var v18 float32 = 0.0
                _ = v18
            }
            v1 += v11
            _ = fi17
        }
        var v19 int16 = (v4 - 42)
        _ = fi16
        _ = v19
    }
    v3 -= -9223372036854775808
    if ((127 > 212) || false) {
        if ("\n\t" <= "") {
            v2 = (v2 ^ 4294967295)
            if (false && (-128 < -1)) {
                var v20 bool = true
                fmt.Println(v12, v20)
                _ = v20
            } else {
                fmt.Println(v12, v4)
                if false {
                    if ((-(v3)) > (v3 << ((2) & 63))) {
                        for fi21 := 0; fi21 < 10; fi21++ {
                            fmt.Println(fi21)
                            switch 0 {
                            case 523:
                                if (1 < 0) {
                                    v3 += 897
                                } else {
                                    switch "hello" {
                                    case "":
                                        fmt.Println(v1, v3, v2)
                                    }
                                }
                            default:
                                v12 = v12
                            }
                            _ = fi21
                        }
                        var v22 uint8 = 54
                        v10 = append(v10, v2)
                        _ = v22
                    } else {
                        v12 += v12
                    }
                    var v23 int64 = v3
                    var v24 string = "foo bar"
                    _ = v23
                    _ = v24
                }
            }
            v12 -= (^(v12))
        }
    }
    var v25 uint64 = (v1 >> ((4294967295) & 63))
    if (!(!true)) {
        if (-2147483648 >= 470) {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        } else {
            for fi26 := 0; fi26 < 10; fi26++ {
                for fi27 := 0; fi27 < 13; fi27++ {
                    v11 += v25
                    _ = fi27
                }
                _ = fi26
            }
            if false {
                if ((v25 ^ (v25 * 1)) >= 42) {
                    for fi28 := 0; fi28 < 19; fi28++ {
                        for fi29 := 0; fi29 < 9; fi29++ {
                            fmt.Println(v2)
                            v1 = (v11 << ((644) & 63))
                            if (520 <= v4) {
                                fmt.Println(v11, v3)
                                var v30 uint64 = 0
                                _ = v30
                            } else {
                                if (float32(1.5) >= 0.0) {
                                    v11 += (v11 & 604)
                                    var v31 uint16 = 255
                                    _ = v31
                                }
                            }
                            _ = fi29
                        }
                        _ = fi28
                    }
                    var v32 int32 = 342
                    _ = v32
                }
            } else {
                v4 = v4
                v33 := GenS2{F0: v2, F1: -128}
                _ = v33
            }
            fmt.Println(v4, v25)
        }
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        v11 = (v1 >> ((834) & 63))
        _ = fi34
    }
    _ = len(v10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
