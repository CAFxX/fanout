package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 string }
func genFn0() uint {
    v5 := []int{-2147483648, 690, 32767}
    var v6 uint8 = 255
    v6 = v6
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 uint32, p1 string) int8 {
    var v7 int16 = 32767
    var v8 int32 = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 33
}

func main() {
    var v1 uint16 = 145
    var v2 bool = false
    var v3 int = 42
    v4 := make(map[string]string)
    var v9 int = -724
    _ = len(v4)
    v4[""] = "\n\t"
    var v10 int16 = 1
    var v11 uint32 = 1
    for fi12 := 0; fi12 < 7; fi12++ {
        _ = len(v4)
        var v13 int = int(v11)
        _ = fi12
        _ = v13
    }
    v10 -= v10
    v11 += (v11 << ((uint(v11)) & 63))
    v14 := []int{(v9 | (-(v9))), (v3 + v3), (v3 ^ int(v11))}
    for fi15 := 0; fi15 < 11; fi15++ {
        v9 = v9
        fmt.Println(fi15)
        _ = fi15
    }
    delete(v4, (("\n\t" + ("a" + "foo bar")) + ""))
    switch (("a" + "foo bar") + "hello") {
    case "":
        v16 := GenS1{F0: 0}
        _ = v16
    default:
        if (42 > -9223372036854775808) {
            v11 = v11
            if true {
                v9 -= (v3 << ((1) & 63))
            }
            v3 -= (v3 % (((v9 >> ((482) & 63))) | 1))
        } else {
            switch 4294967295 {
            case 127:
                v17 := GenS2{F0: (v3 % (((-(v9))) | 1))}
                _ = v17
            case -187:
                v18 := GenS2{F0: v9}
                _ = v18
            default:
                var v19 uint64 = 677
                _ = v19
            }
            fmt.Println(v9, v3, v9)
            var v20 uint16 = v1
            _ = v20
        }
    }
    if (1e10 < float64(3.14159)) {
        if v2 {
            var v21 int32 = 1
            fmt.Println(v11, v9, v9)
            switch "a" {
            case "foo bar":
                if ((3.14159 - float64(0.0)) != 0.0) {
                    var v22 float32 = 1e10
                    _ = v22
                }
            case "hello":
                var v23 int32 = v21
                _ = v23
            case "\n\t":
                if (!(((1.5 / 1.5) * 0.1) == -2.25)) {
                    var v24 uint8 = 228
                    _ = v24
                } else {
                    v21 -= -2147483648
                    switch "" {
                    case "hello":
                        var v25 uint16 = v1
                        _ = v25
                    case "":
                        for fi26 := 0; fi26 < 10; fi26++ {
                            v9 -= (v9 / ((fi26) | 1))
                            v27 := GenS2{F0: v9}
                            _ = fi26
                            _ = v27
                        }
                    case "foo bar":
                        if ((v2 && v2) || v2) {
                            v2 = ((^(v10)) >= (v10 * v10))
                            v28 := GenS1{F0: 222}
                            _ = v28
                        } else {
                            fmt.Println(v21)
                        }
                    }
                    switch "foo bar" {
                    case "a":
                        switch v9 {
                        case 2147483647:
                            switch (v9 << ((696) & 63)) {
                            case 32768:
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    v3 += (^(v9))
                                    var v30 uint8 = 42
                                    if false {
                                        if (0.0 > float32((float64(0.0) - float64(0.1)))) {
                                            var v31 uint16 = v1
                                            _ = v31
                                        } else {
                                            delete(v4, "foo bar")
                                        }
                                        v2 = ((^(v21)) > v21)
                                        if (115 <= 42) {
                                            for fi32 := 0; fi32 < 6; fi32++ {
                                                v33 := GenS1{F0: v30}
                                                _ = fi32
                                                _ = v33
                                            }
                                            v21 = v21
                                        }
                                    } else {
                                        v34 := v4["foo bar"]
                                        v34 = "a"
                                        _ = v34
                                    }
                                    _ = fi29
                                    _ = v30
                                }
                            case -2147483648:
                                if len(v14) > 0 {
                                    v14[0] = -719
                                }
                            }
                        default:
                            fmt.Println(v1, v21)
                        }
                    case "\n\t":
                        var v35 int = int(v1)
                        _ = v35
                    case "foo bar":
                        var v36 uint = 2
                        _ = v36
                    default:
                        v9 = 42
                    }
                }
            }
            _ = v21
        }
        for fi37 := 0; fi37 < 19; fi37++ {
            v38 := GenS3{F0: 1, F1: "foo bar"}
            v3 -= (fi37 ^ -260)
            _ = fi37
            _ = v38
        }
    }
    var v39 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v9
}
