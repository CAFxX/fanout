package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0() uint32 {
    v4 := []int{764, 1, 42, 32767}
    var v5 uint32 = 2
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int16 = 1
        var v8 uint32 = v5
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = v4
    _ = v5
    return 2147483648
}
func genFn1(p0 uint16) int16 {
    v9 := GenS1{F0: 954, F1: 107, F2: 1}
    v10 := []int{}
    _ = p0
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := []string{("hello" + ("\n\t" + "hello")), "", "foo bar"}
    var v2 uint16 = 2
    var v3 bool = (167 == 804)
    v3 = (3.14159 > 0.1)
    if v3 {
        v11 := GenS1{F0: 42, F1: 0, F2: -2147483648}
        _ = len(v1)
        v3 = true
        _ = v11
    }
    var v12 int64 = 42
    var v13 int8 = 75
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = (83 <= 2)
    v13 += v13
    var v14 int16 = 42
    v15 := GenS1{F0: 300, F1: v13, F2: 9223372036854775807}
    fmt.Println(v2, v12)
    var v16 int = int(v12)
    switch v16 {
    case -368:
        var v17 float64 = 0.1
        _ = v17
    case 1:
        v18 := GenS2{F0: v14}
        _ = v18
    }
    v14 = genFn1(0)
    switch ("" + (((("foo bar" + "") + (("\n\t" + "\n\t") + "hello")) + ("\n\t" + "")) + ("hello" + "a"))) {
    case "hello":
        for fi19 := 0; fi19 < 3; fi19++ {
            if false {
                for fi20 := 0; fi20 < 12; fi20++ {
                    if (42 >= 4294967295) {
                        if (!v3) {
                            var v21 int16 = (v14 + (v14 ^ v14))
                            _ = v21
                        }
                        v13 = -128
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                    }
                    _ = fi20
                }
                v3 = v3
            } else {
                var v22 uint = 482
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi23 := 0; fi23 < 18; fi23++ {
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    var v24 bool = v3
                    _ = len(v1)
                    var v25 int8 = 0
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = v22
            }
            _ = fi19
        }
    default:
        v26 := GenS3{F0: 2, F1: int32(v2)}
        _ = v26
    }
    v27 := GenS2{F0: v14}
    var v28 int64 = -9223372036854775808
    for fi29 := 0; fi29 < 16; fi29++ {
        fmt.Println(v12)
        for fi30 := 0; fi30 < 13; fi30++ {
            v31 := GenS3{F0: 4294967295, F1: 42}
            _ = fi30
            _ = v31
        }
        if (v12 > (v28 * 42)) {
            var v32 int16 = v14
            _ = v32
        } else {
            for fi33 := 0; fi33 < 6; fi33++ {
                v34 := GenS2{F0: genFn1(2)}
                var v35 int64 = (^(v12))
                var v36 int16 = (-(v14))
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
            v3 = (!v3)
            for fi37 := 0; fi37 < 5; fi37++ {
                if (!false) {
                    var v38 int8 = (v13 & v13)
                    switch "hello" {
                    case "":
                        v14 = v14
                    case "\n\t":
                        v13 = v13
                    }
                    switch (fi37 << ((675) & 63)) {
                    case -2147483648:
                        _ = len(v1)
                    }
                    _ = v38
                } else {
                    v13 += (v13 & v13)
                }
                var v39 int32 = 42
                _ = fi37
                _ = v39
            }
        }
        _ = fi29
    }
    v40 := GenS1{F0: 2, F1: v13, F2: v16}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v40
}
