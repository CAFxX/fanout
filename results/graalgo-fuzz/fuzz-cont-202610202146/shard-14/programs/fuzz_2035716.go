package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint; F1 int16; F2 uint64 }
func genFn0(p0 int16, p1 int, p2 int32) int32 {
    switch "a" {
    case "foo bar":
        for fi4 := 0; fi4 < 9; fi4++ {
            v5 := GenS1{F0: ("hello" + "foo bar")}
            for fi6 := 0; fi6 < 4; fi6++ {
                fmt.Println(p2, p0)
                _ = fi6
            }
            _ = fi4
            _ = v5
        }
    }
    var v7 float64 = (((float64(0.0) - 3.14159) + ((-2.25 + 0.1) / 1e10)) * 1e10)
    switch ("" + "") {
    case "":
        switch p1 {
        case -9223372036854775808:
            var v8 bool = (2 < 314)
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 ^ -171)
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    var v9 int8 = 21
    v9 = v9
    for fi10 := 0; fi10 < 6; fi10++ {
        if (-2147483648 >= -128) {
            if true {
                if ((v9 / ((v9) | 1)) < (v9 ^ (^(v9)))) {
                    v11 := GenS1{F0: ("\n\t" + "a")}
                    fmt.Println(v9, fi10)
                    _ = v11
                }
                if (true || true) {
                    if (true && (420 > 886)) {
                        for fi12 := 0; fi12 < 20; fi12++ {
                            switch 630 {
                            case 42:
                                fmt.Println(fi10)
                            case 127:
                                if true {
                                    v9 = v9
                                    fmt.Println(fi10)
                                    v13 := []int{fi12, (fi10 - fi10), (^(fi10)), fi12}
                                    _ = v13
                                }
                            case 128:
                                fmt.Println(fi10, fi10)
                            default:
                                v14 := GenS1{F0: (("" + "a") + "foo bar")}
                                _ = v14
                            }
                            _ = fi12
                        }
                    }
                    v15 := GenS1{F0: "foo bar"}
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v9 = (^(v9))
                        fmt.Println(v9)
                        v17 := GenS1{F0: "foo bar"}
                        if (("hello" + ("" + "a")) > ("a" + "a")) {
                            fmt.Println(fi16, v9)
                        } else {
                            if (378 > 1) {
                                if (v9 >= v9) {
                                    if (!((344 != 18446744073709551615) && (true || false))) {
                                        v18 := GenS2{F0: (v9 / ((1) | 1)), F1: -32768}
                                        _ = v18
                                    }
                                    var v19 int32 = 1
                                    v19 = 42
                                    _ = v19
                                }
                            }
                            fmt.Println(fi16)
                        }
                        var v20 uint16 = 0
                        v9 = (^(v9))
                        var v21 float64 = 0.0
                        _ = fi16
                        _ = v17
                        _ = v20
                        _ = v21
                    }
                    _ = v15
                } else {
                    v22 := GenS3{F0: 877, F1: 537, F2: 1}
                    v23 := GenS3{F0: 255, F1: -32768, F2: 42}
                    fmt.Println(v9, v9)
                    _ = v22
                    _ = v23
                }
            }
        } else {
            v9 += (v9 & (v9 & v9))
        }
        _ = fi10
    }
    v24 := v3[1]
    var v25 uint64 = 58
    for fi26 := 0; fi26 < 20; fi26++ {
        fmt.Println(v24, v24, v24)
        for fi27 := 0; fi27 < 1; fi27++ {
            v9 -= 0
            var v28 uint = 306
            v9 = 68
            v28 = (^(v28))
            _ = fi27
            _ = v28
        }
        v9 -= -1
        _ = fi26
    }
    switch v24 {
    case "hello":
        if (((false || false) && false) && (3.14159 > 0.0)) {
            v29 := v2[4294967295]
            _ = v29
        }
    case "\n\t":
        var v30 uint16 = 42
        _ = v30
    default:
        v31 := []int{-9223372036854775808}
        _ = v31
    }
    v25 += 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v9
}
