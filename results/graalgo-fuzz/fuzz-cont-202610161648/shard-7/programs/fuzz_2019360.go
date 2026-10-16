package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
func genFn0() bool {
    v7 := GenS1{F0: -1, F1: "a", F2: 0}
    var v8 uint8 = 0
    if (0 <= -18) {
        var v9 uint8 = 2
        v8 -= v9
        _ = v9
    } else {
        switch 1 {
        case 9223372036854775807:
            if false {
                if ((v8 / ((v8) | 1)) < (^(v8))) {
                    fmt.Println(v8)
                }
                switch "a" {
                case "hello":
                    v10 := []int{}
                    _ = v10
                }
                var v11 int8 = -128
                _ = v11
            } else {
                var v12 bool = true
                _ = v12
            }
        case 2147483647:
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 uint64 = 4294967295
                _ = fi13
                _ = v14
            }
        }
        v8 -= (^(v8))
    }
    v8 = v8
    _ = v7
    _ = v8
    return (true && true)
}
func genFn1(p0 int32) int {
    v15 := []int{}
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 string = ""
        v18 := GenS1{F0: 9223372036854775807, F1: v17, F2: 80}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: -985, F1: "", F2: 179}
    if len(v15) > 0 {
        v15[0] = 32768
    }
    _ = p0
    _ = v15
    _ = v19
    return 0
}

func main() {
    var v1 int = 32767
    var v2 float64 = (-2.25 - -2.25)
    v3 := make(map[uint32]bool)
    var v4 bool = (4294967295 != 255)
    var v5 float32 = (0.1 / 1e10)
    var v6 int16 = 1000
    fmt.Println(v1)
    var v20 int8 = 84
    var v21 int = (-(v1))
    for fi22 := 0; fi22 < 1; fi22++ {
        v4 = true
        _ = fi22
    }
    if (false || ((!v4) || (2147483647 != 2147483647))) {
        v23 := GenS1{F0: -9223372036854775808, F1: "hello", F2: 42}
        var v24 int64 = 336
        _ = v23
        _ = v24
    }
    var v25 int32 = 42
    switch genFn1(1) {
    case 127:
        if ((4294967295 > 498) && (42 == 18446744073709551615)) {
            v26 := GenS1{F0: 667, F1: "\n\t", F2: 32767}
            _ = v26
        } else {
            if v4 {
                if (1 != v20) {
                    v2 = -2.25
                    v27 := GenS2{F0: 2147483647}
                    fmt.Println(v1)
                    _ = v27
                }
                for fi28 := 0; fi28 < 4; fi28++ {
                    v29 := GenS3{F0: 252}
                    _ = fi28
                    _ = v29
                }
            } else {
                for fi30 := 0; fi30 < 5; fi30++ {
                    var v31 float64 = v2
                    for fi32 := 0; fi32 < 1; fi32++ {
                        fmt.Println(v1)
                        v4 = v4
                        _ = fi32
                    }
                    _ = fi30
                    _ = v31
                }
                switch ("a" + "") {
                case "hello":
                    var v33 uint8 = 127
                    _ = v33
                case "foo bar":
                    if (!((42 != 482) && (65535 <= 362))) {
                        var v34 int16 = 32767
                        _ = v34
                    } else {
                        _ = len(v3)
                        v2 = 0.1
                    }
                }
            }
            v5 = float32(-2.25)
            if ((42 <= 661) && (!((0.1 > (v2 + ((v2 - v2) + v2))) || v4))) {
                v35 := v3[541]
                switch "hello" {
                case "hello":
                    v2 -= (v2 + ((v2 - v2) + (1.5 / v2)))
                default:
                    for fi36 := 0; fi36 < 2; fi36++ {
                        var v37 uint64 = 0
                        v4 = v35
                        _ = fi36
                        _ = v37
                    }
                }
                v6 -= v6
                _ = v35
            }
        }
    case -9223372036854775808:
        v1 = genFn1((v25 / ((-1) | 1)))
    }
    var v38 int64 = 107
    var v39 int8 = v20
    v40 := GenS3{F0: 255}
    if v4 {
        if true {
            fmt.Println(v6)
            var v41 uint = 65535
            v2 = v2
            _ = v41
        }
        v42 := GenS2{F0: v25}
        _ = v42
    } else {
        v5 = float32(v2)
    }
    switch (v1 >> ((255) & 63)) {
    case 200:
        var v43 uint64 = 537
        _ = v43
    }
    switch -2147483648 {
    case 4294967295:
        var v44 uint = 1
        _ = v44
    default:
        v21 += v21
    }
    var v45 uint16 = 990
    var v46 uint8 = 1
    if ((uint64(v45) > 2) || v4) {
        if (605 > genFn1((v25 << ((453) & 63)))) {
            if v4 {
                v47 := GenS1{F0: int64(v25), F1: "\n\t", F2: v45}
                _ = v47
            }
        }
        fmt.Println(v45, v21)
        v3[2147483648] = v4
    }
    v20 = v20
    var v48 int32 = (v25 + -116)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v45
    _ = v46
    _ = v48
    _ = v5
    _ = v6
}
