package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint64 }
func genFn0(p0 uint, p1 int, p2 int32) int {
    fmt.Println(p1, p0, p1)
    var v6 float32 = 0.1
    switch p1 {
    case 42:
        var v7 int16 = -32768
        _ = v7
    case -2147483648:
        var v8 uint = (p0 / ((p0) | 1))
        _ = v8
    default:
        var v9 float64 = 1.5
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}
func genFn1() bool {
    var v10 uint = 2
    var v11 uint64 = 4294967295
    var v12 int16 = 32767
    for fi13 := 0; fi13 < 20; fi13++ {
        v10 = (^(v10))
        v14 := GenS1{F0: 113, F1: 40, F2: 18446744073709551615}
        _ = fi13
        _ = v14
    }
    _ = v10
    _ = v11
    _ = v12
    return true
}

func main() {
    v1 := []string{"", ((((("" + "\n\t") + ("" + "hello")) + "\n\t") + "foo bar") + (("hello" + "hello") + "foo bar")), "hello"}
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    var v4 int64 = -1
    v5 := []int{}
    if false {
        v4 += (v4 & (v4 + (v4 % ((-9223372036854775808) | 1))))
        v2["a"] = false
    } else {
        if true {
            v4 = v4
            delete(v3, "\n\t")
            if false {
                var v15 uint8 = 255
                _ = v15
            }
        } else {
            v4 += v4
            v4 -= v4
            for fi16 := 0; fi16 < 2; fi16++ {
                for fi17 := 0; fi17 < 15; fi17++ {
                    if (false && true) {
                        var v18 int64 = v4
                        _ = v18
                    } else {
                        if (!false) {
                            fmt.Println(v4, v4)
                            v4 = (^(v4))
                            var v19 string = "a"
                            _ = v19
                        } else {
                            v4 += v4
                            var v20 float32 = -2.25
                            _ = v20
                        }
                        var v21 float32 = -2.25
                        _ = len(v2)
                        _ = v21
                    }
                    v4 -= (v4 / (((v4 >> ((983) & 63))) | 1))
                    v4 = v4
                    _ = fi17
                }
                _ = fi16
            }
        }
    }
    var v22 float32 = (1e10 / 1e10)
    if len(v5) > 0 {
        v5[0] = 568
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        if true {
            v4 += v4
            if (!true) {
                v4 -= 9223372036854775807
                for fi24 := 0; fi24 < 11; fi24++ {
                    switch "hello" {
                    case "":
                        var v25 uint32 = 2
                        _ = v25
                    }
                    _ = len(v5)
                    _ = fi24
                }
            }
            fmt.Println(fi23)
        }
        var v26 int64 = (v4 >> ((65535) & 63))
        _ = fi23
        _ = v26
    }
    fmt.Println(v22, v4)
    var v27 string = "\n\t"
    var v28 int16 = 0
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 int8 = -1
        var v31 int16 = (v28 >> ((18446744073709551615) & 63))
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int = 32768
    v2[v27] = true
    for fi33 := 0; fi33 < 11; fi33++ {
        switch (((v27 + "a") + ("\n\t" + (v27 + v27))) + ("" + ((v27 + "hello") + ("hello" + v27)))) {
        case "a":
            delete(v3, (("\n\t" + v27) + v27))
        case "":
            switch (v32 ^ v32) {
            case 1:
                var v34 string = v27
                _ = v34
            case -440:
                v35 := v2[""]
                _ = v35
            case 965:
                var v36 int8 = 42
                _ = v36
            }
        default:
            for fi37 := 0; fi37 < 10; fi37++ {
                if (0.1 >= 0.1) {
                    fmt.Println(fi33, v4, v4)
                    var v38 uint64 = 739
                    v39 := GenS1{F0: 42, F1: 230, F2: v38}
                    _ = v38
                    _ = v39
                }
                _ = fi37
            }
        }
        _ = len(v5)
        fmt.Println(v28, v28)
        _ = fi33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
