package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }

func main() {
    var v1 int8 = -128
    v2 := make(map[int]uint32)
    var v3 uint8 = 42
    v4 := []uint32{2147483647, 0}
    v5 := make(map[string]uint32)
    if len(v4) > 0 {
        v4[0] = 2
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 float64 = ((-2.25 - 0.0) - -2.25)
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        if (1 >= 2) {
            switch "" {
            case "foo bar":
                if true {
                    v9 := GenS1{F0: 2147483648, F1: -517, F2: ""}
                    var v10 string = "foo bar"
                    _ = v10
                    _ = v9
                } else {
                    v2[(fi8 - (-(fi8)))] = 0
                    v11 := GenS1{F0: 877, F1: -929, F2: "foo bar"}
                    switch "" {
                    case "foo bar":
                        v3 -= v3
                    case "\n\t":
                        switch "\n\t" {
                        case "":
                            for fi12 := 0; fi12 < 3; fi12++ {
                                if true {
                                    fmt.Println(fi12, v3, v1)
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                fmt.Println(fi8, fi8, fi8)
                                _ = fi12
                            }
                        case "hello":
                            var v13 int64 = -540
                            _ = v13
                        }
                    default:
                        v14 := GenS2{F0: 1}
                        _ = v14
                    }
                    _ = v11
                }
            case "":
                var v15 int = fi8
                _ = v15
            }
        } else {
            var v16 int64 = 1
            for fi17 := 0; fi17 < 10; fi17++ {
                v1 = (v1 + (v1 | v1))
                _ = fi17
            }
            _ = v16
        }
        _ = fi8
    }
    v3 = v3
    v3 -= v3
    v3 -= (v3 * (^(v3)))
    for fi18 := 0; fi18 < 2; fi18++ {
        delete(v5, ("hello" + "\n\t"))
        var v19 uint64 = 1
        v20 := GenS1{F0: 4294967295, F1: 332, F2: ("a" + "a")}
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v3 = (v3 | 42)
        switch (-(fi21)) {
        case 32768:
            fmt.Println(fi21, fi21)
        default:
            if len(v4) > 0 {
                v4[0] = 873
            }
        }
        _ = fi21
    }
    v3 = v3
    if true {
        var v22 float32 = 0.0
        fmt.Println(v3, v22)
        for fi23 := 0; fi23 < 1; fi23++ {
            for fi24 := 0; fi24 < 20; fi24++ {
                fmt.Println(fi23, v1, v3)
                v25 := GenS1{F0: 2, F1: -458, F2: "a"}
                v26 := GenS3{F0: false}
                _ = fi24
                _ = v25
                _ = v26
            }
            fmt.Println(v3)
            var v27 int16 = 42
            _ = fi23
            _ = v27
        }
        _ = v22
    } else {
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 bool = (2 < 127)
            if ("foo bar" < (((("a" + "foo bar") + ("a" + "foo bar")) + "\n\t") + ("foo bar" + ""))) {
                v3 = 0
                v29 = (0.1 > 3.14159)
            } else {
                fmt.Println(v3)
            }
            _ = fi28
            _ = v29
        }
        for fi30 := 0; fi30 < 2; fi30++ {
            for fi31 := 0; fi31 < 11; fi31++ {
                fmt.Println(v1, fi30)
                v3 += 2
                var v32 float64 = float64(0.1)
                fmt.Println(v3, v1, fi30)
                _ = fi31
                _ = v32
            }
            _ = fi30
        }
        for fi33 := 0; fi33 < 15; fi33++ {
            fmt.Println(fi33)
            var v34 uint8 = (v3 << ((18446744073709551615) & 63))
            var v35 uint32 = 1
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
