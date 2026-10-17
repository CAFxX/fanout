package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int; F1 int64; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint }
func genFn0(p0 int16) string {
    switch -2147483648 {
    case 42:
        var v4 uint16 = 255
        _ = v4
    case 2147483647:
        p0 -= (^(p0))
    }
    p0 = p0
    v5 := []int{}
    v6 := GenS1{F0: 0, F1: 0}
    switch "hello" {
    case "hello":
        if (230 > 51) {
            _ = len(v5)
            _ = len(v5)
            for fi7 := 0; fi7 < 19; fi7++ {
                v8 := GenS1{F0: 32767, F1: 446}
                var v9 int16 = p0
                _ = fi7
                _ = v8
                _ = v9
            }
        } else {
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 int16 = (^(p0))
                p0 = (^(p0))
                fmt.Println(p0)
                _ = fi10
                _ = v11
            }
            p0 = p0
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int = -785
    v3 := make(map[uint32]int)
    if ("\n\t" > (((((("" + "foo bar") + "") + "foo bar") + "") + (("hello" + "hello") + "foo bar")) + "hello")) {
        if (-1 == -909) {
            var v12 int16 = 0
            _ = v12
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            v2 += fi13
            v1 = v1
            _ = fi13
        }
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint16 = 461
            if (fi14 == fi14) {
                var v16 uint64 = 8
                switch "foo bar" {
                case "a":
                    for fi17 := 0; fi17 < 2; fi17++ {
                        var v18 int8 = -108
                        v19 := GenS1{F0: 0, F1: v15}
                        fmt.Println(v2)
                        v20 := []int{v2, fi17, v2, int(v15)}
                        v15 -= (v15 ^ (v15 * 42))
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                case "hello":
                    v1 = (v1 | (v1 % ((2147483647) | 1)))
                default:
                    switch (-(fi14)) {
                    case 156:
                        var v21 uint16 = 1
                        _ = v21
                    }
                }
                _ = v16
            }
            v22 := GenS2{F0: fi14, F1: -1, F2: false}
            _ = fi14
            _ = v15
            _ = v22
        }
    }
    var v23 int16 = -1
    var v24 bool = false
    v23 += 1
    var v25 int8 = 42
    switch v2 {
    case 9223372036854775807:
        v25 += (v25 % ((v25) | 1))
    }
    var v26 uint32 = (v1 % ((986) | 1))
    if ((v25 * v25) != (v25 & v25)) {
        var v27 uint64 = 42
        var v28 int32 = -947
        for fi29 := 0; fi29 < 13; fi29++ {
            if (!(v2 >= (v2 + fi29))) {
                delete(v3, (^(v1)))
                if ((v26 % (((v26 * v26)) | 1)) < v26) {
                    var v30 bool = v24
                    _ = v30
                } else {
                    var v31 float32 = 3.14159
                    fmt.Println(v31, v24, v1)
                    _ = v31
                }
                v32 := GenS3{F0: (fi29 & fi29)}
                _ = v32
            }
            _ = fi29
        }
        _ = v27
        _ = v28
    }
    v26 += v26
    var v33 int8 = (v25 ^ (v25 + v25))
    v26 = v26
    v23 -= v23
    v34 := GenS4{F0: 0, F1: 18446744073709551615}
    var v35 string = "a"
    var v36 float32 = -2.25
    fmt.Println(v26, v2, v36)
    for fi37 := 0; fi37 < 9; fi37++ {
        v24 = ((!v24) || (v1 == v26))
        var v38 int64 = 9223372036854775807
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
