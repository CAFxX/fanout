package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int32; F2 bool }
func genFn0(p0 int64, p1 int32, p2 uint8) uint {
    switch "" {
    case "":
        v6 := GenS1{F0: 4294967295}
        _ = v6
    case "hello":
        v7 := GenS1{F0: 66}
        _ = v7
    }
    v8 := []int{9223372036854775807}
    if len(v8) > 0 {
        v8[0] = 9223372036854775807
    }
    p1 = (p1 | int32(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 355
}
func genFn1(p0 int) string {
    var v9 int = (-(p0))
    switch "a" {
    case "foo bar":
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS1{F0: 6}
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = v9
    return "a"
}

func main() {
    v1 := []uint32{1, 2}
    v2 := []string{"a"}
    var v3 int8 = -128
    var v4 int = 127
    v5 := []string{("\n\t" + "a")}
    for fi12 := 0; fi12 < 15; fi12++ {
        v1 = append(v1, 800)
        _ = fi12
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    fmt.Println(v3, v4)
    if (-2.25 <= (((0.1 - 0.0) / (-2.25 - (3.14159 - ((3.14159 / 1e10) * (3.14159 / -2.25))))) / 3.14159)) {
        if true {
            if true {
                v1 = append(v1, 2147483648)
            }
        } else {
            if false {
                if (!true) {
                    v3 += int8(v4)
                    var v13 int = 128
                    if true {
                        if true {
                            switch "foo bar" {
                            case "":
                                v13 = -9223372036854775808
                            }
                        }
                        if false {
                            v3 = (v3 & v3)
                            v3 = v3
                        } else {
                            v4 = v13
                            var v14 bool = true
                            var v15 uint16 = 1
                            _ = v14
                            _ = v15
                        }
                    }
                    _ = v13
                } else {
                    v3 -= v3
                    var v16 int64 = 42
                    _ = len(v5)
                    _ = v16
                }
                v4 = (v4 ^ (v4 - (-(v4))))
            }
            fmt.Println(v3, v3)
            switch "\n\t" {
            case "hello":
                if ((("" + "foo bar") + ("hello" + (("foo bar" + ("foo bar" + "foo bar")) + ""))) > "foo bar") {
                    switch "a" {
                    case "":
                        for fi17 := 0; fi17 < 7; fi17++ {
                            for fi18 := 0; fi18 < 12; fi18++ {
                                _ = len(v5)
                                _ = len(v5)
                                _ = fi18
                            }
                            var v19 int32 = 42
                            v20 := GenS2{F0: 2147483648}
                            _ = fi17
                            _ = v19
                            _ = v20
                        }
                    default:
                        for fi21 := 0; fi21 < 13; fi21++ {
                            fmt.Println(v4)
                            _ = fi21
                        }
                    }
                    if false {
                        if true {
                            v22 := GenS1{F0: 0}
                            var v23 uint64 = 4294967295
                            _ = v22
                            _ = v23
                        }
                        v3 -= -47
                    } else {
                        _ = len(v2)
                    }
                    v24 := GenS3{F0: 1, F1: -211, F2: (true || false)}
                    _ = v24
                } else {
                    var v25 float64 = 3.14159
                    _ = v25
                }
            case "":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "foo bar":
                v3 += -128
            default:
                fmt.Println(v3, v3)
            }
        }
        var v26 uint64 = 42
        _ = v26
    }
    v3 = (^(v3))
    var v27 float64 = float64(3.14159)
    switch (v4 << ((4294967295) & 63)) {
    case 9223372036854775807:
        v3 -= 100
    }
    v1 = append(v1, 857)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
