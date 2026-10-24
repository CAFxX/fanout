package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int16; F1 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = -992
    v3 := make(map[uint32]int)
    v4 := []uint32{65535, 38, 712}
    var v5 string = "foo bar"
    v6 := make(map[string]bool)
    v7 := GenS1{F0: 430}
    var v8 int64 = -1
    switch 128 {
    case 2147483647:
        v8 -= v8
    case -2147483648:
        v9 := GenS1{F0: 314}
        _ = v9
    case 4294967295:
        v2 = (-(v2))
    default:
        var v10 int = -328
        _ = v10
    }
    v8 = v8
    _ = len(v4)
    v5 = (v5 + (v5 + v5))
    var v11 float64 = (0.0 + (0.1 * 0.1))
    switch v5 {
    case "foo bar":
        v12 := GenS2{F0: 1, F1: 42}
        _ = v12
    default:
        switch "hello" {
        case "foo bar":
            if false {
                v2 += -32768
            } else {
                var v13 string = "foo bar"
                if false {
                    var v14 int16 = (^(v2))
                    v15 := v3[2147483647]
                    _ = v14
                    _ = v15
                }
                _ = v13
            }
        }
    }
    if false {
        var v16 float64 = v11
        _ = v16
    } else {
        if true {
            if (v11 > (((float64(1e10) * 0.1) / (1.5 * 1.5)) + (v11 + 1.5))) {
                v2 -= 1000
                v8 -= (v8 % ((v8) | 1))
                for fi17 := 0; fi17 < 5; fi17++ {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        _ = len(v4)
                        if true {
                            if true {
                                var v19 uint = 0
                                _ = v19
                            }
                            if false {
                                v2 += -1
                                if false {
                                    var v20 uint = 887
                                    _ = v20
                                }
                            }
                            v11 = (v11 - 1e10)
                        }
                        var v21 uint16 = 45
                        _ = len(v4)
                        _ = fi18
                        _ = v21
                    }
                    var v22 string = v5
                    _ = fi17
                    _ = v22
                }
            } else {
                if true {
                    fmt.Println(v11)
                } else {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v24 := v6[("foo bar" + (((v5 + v5) + (v5 + (v5 + (v5 + "foo bar")))) + "a"))]
                        _ = fi23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 2; fi25++ {
                        v5 = ("hello" + "foo bar")
                        for fi26 := 0; fi26 < 15; fi26++ {
                            delete(v6, v5)
                            fmt.Println(fi25, v5)
                            _ = fi26
                        }
                        _ = fi25
                    }
                }
                v5 = v5
            }
            v5 = ("foo bar" + "a")
        }
    }
    _ = len(v3)
    var v27 uint8 = 2
    if false {
        var v28 float32 = 0.0
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS2{F0: 2147483647, F1: 2}
            var v31 uint16 = 65535
            for fi32 := 0; fi32 < 12; fi32++ {
                switch 0 {
                case 42:
                    var v33 uint64 = 2
                    switch (fi29 - 32768) {
                    case -1:
                        switch v5 {
                        case "a":
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                if false {
                                    if (0 < 65535) {
                                        v4 = append(v4, uint32(v31))
                                    } else {
                                        delete(v3, 2147483647)
                                    }
                                } else {
                                    var v34 float32 = float32(v11)
                                    if len(v4) > 0 {
                                        v4[0] = 346
                                    }
                                    var v35 bool = (false && ((!(!false)) && ((!true) && true)))
                                    var v36 int8 = 0
                                    v37 := GenS3{F0: (v2 >> ((255) & 63)), F1: fi32}
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                            }
                        case "\n\t":
                            var v38 uint32 = 0
                            _ = v38
                        case "foo bar":
                            v2 = 42
                        }
                    case 128:
                        if true {
                            var v39 int32 = -93
                            var v40 uint64 = 42
                            _ = v39
                            _ = v40
                        }
                    }
                    v41 := GenS3{F0: (-(v2)), F1: 2147483647}
                    _ = v33
                    _ = v41
                }
                _ = fi32
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        for fi42 := 0; fi42 < 15; fi42++ {
            var v43 int8 = 100
            _ = fi42
            _ = v43
        }
        _ = v28
    } else {
        if len(v4) > 0 {
            v4[0] = 42
        }
        var v44 float64 = (0.1 - ((v11 / 3.14159) * v11))
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
