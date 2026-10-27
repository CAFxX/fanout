package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 int64; F2 string }
func genFn0(p0 bool, p1 int64) int8 {
    fmt.Println(p1)
    fmt.Println(p1)
    _ = p0
    _ = p1
    return -124
}
func genFn1(p0 uint) int16 {
    var v4 int64 = 1
    var v5 uint64 = 392
    fmt.Println(v4, v4, v5)
    for fi6 := 0; fi6 < 4; fi6++ {
        if true {
            v7 := GenS1{F0: 421}
            _ = v7
        } else {
            if (1.5 <= 1e10) {
                switch (fi6 * -997) {
                case 950:
                    switch ("a" + ("hello" + "foo bar")) {
                    case "a":
                        if false {
                            fmt.Println(v4, p0)
                            p0 = (^(p0))
                            v8 := GenS2{F0: int32(v5)}
                            _ = v8
                        } else {
                            switch "hello" {
                            case "hello":
                                if false {
                                    v9 := GenS1{F0: 1}
                                    v10 := GenS1{F0: -1}
                                    var v11 float64 = float64(float32(1.5))
                                    _ = v10
                                    _ = v11
                                    _ = v9
                                }
                            default:
                                var v12 float32 = -2.25
                                _ = v12
                            }
                        }
                    case "":
                        fmt.Println(fi6)
                    case "foo bar":
                        if false {
                            var v13 uint64 = 42
                            _ = v13
                        } else {
                            v14 := []int{}
                            if true {
                                v15 := GenS2{F0: -2147483648}
                                _ = v15
                            }
                            _ = v14
                        }
                    default:
                        var v16 bool = (-2.25 > float64(1e10))
                        _ = v16
                    }
                }
                if (!false) {
                    var v17 string = ""
                    _ = v17
                } else {
                    var v18 uint = (p0 % ((p0) | 1))
                    v19 := GenS2{F0: -2147483648}
                    _ = v18
                    _ = v19
                }
                if ((v5 >> (((p0 ^ p0)) & 63)) <= v5) {
                    var v20 uint64 = v5
                    _ = v20
                }
            } else {
                p0 += p0
            }
            v4 = v4
        }
        v5 = v5
        var v21 uint8 = 2
        _ = fi6
        _ = v21
    }
    v4 -= (v4 & -9223372036854775808)
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 int = 127
    var v2 bool = false
    v3 := make(map[string]string)
    var v22 int64 = -351
    var v23 uint = 42
    if v2 {
        v22 = (v22 & int64(v1))
    }
    if v2 {
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 float64 = -2.25
            v3["a"] = "hello"
            for fi26 := 0; fi26 < 6; fi26++ {
                v27 := GenS3{F0: 32768}
                _ = fi26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
        v28 := GenS2{F0: -1}
        for fi29 := 0; fi29 < 16; fi29++ {
            if (165 > 255) {
                v1 = fi29
            }
            _ = fi29
        }
        _ = v28
    }
    if (!(v2 || ((v23 % ((v23) | 1)) <= (^(v23))))) {
        v30 := []int{1, (v1 - -1), v1}
        for fi31 := 0; fi31 < 7; fi31++ {
            v2 = v2
            _ = fi31
        }
        v3["a"] = ("a" + ("hello" + "foo bar"))
        _ = v30
    }
    v32 := v3["foo bar"]
    v33 := GenS3{F0: (v1 & v1)}
    v22 = (v22 % ((v22) | 1))
    var v34 int = v1
    for fi35 := 0; fi35 < 11; fi35++ {
        fmt.Println(v23, v32)
        _ = fi35
    }
    if (-301 > 42) {
        v2 = (v23 == v23)
    }
    for fi36 := 0; fi36 < 16; fi36++ {
        if false {
            if false {
                fmt.Println(v2)
                var v37 int16 = genFn1(v23)
                _ = v37
            }
            switch v34 {
            case 127:
                v38 := GenS3{F0: (-(v1))}
                _ = v38
            case -2147483648:
                for fi39 := 0; fi39 < 10; fi39++ {
                    fmt.Println(v2, fi39, v1)
                    v2 = v2
                    _ = fi39
                }
            }
        }
        _ = fi36
    }
    v23 -= 2
    v1 = (v1 / (((^(v1))) | 1))
    switch v34 {
    case 374:
        var v40 float32 = 0.1
        _ = v40
    case 124:
        fmt.Println(v32)
    default:
        v41 := GenS4{F0: 2, F1: v22, F2: "foo bar"}
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
