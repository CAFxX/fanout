package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int64; F2 int16 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS4 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint32 = 53
    var v2 int = -469
    v3 := make(map[int]int)
    var v4 uint8 = 0
    v5 := make(map[string]bool)
    if ((!(703 <= 0)) && ((!(!(false && (65535 >= (v1 % (((v1 + 2147483647)) | 1)))))) && false)) {
        v6 := GenS1{F0: 42}
        v5["hello"] = false
        _ = v6
    } else {
        v1 = v1
        v7 := GenS1{F0: 0}
        _ = v7
    }
    v4 = (v4 * v4)
    var v8 float32 = 1e10
    fmt.Println(v4)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint16 = 789
        if ((false && false) || true) {
            v2 += (fi9 - 215)
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 string = "hello"
                _ = fi11
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 11; fi13++ {
                v8 = ((((1.5 / (-2.25 - (v8 - v8))) + -2.25) - (v8 / (v8 - v8))) / 1e10)
                v14 := GenS2{F0: (("foo bar" + (("a" + ("foo bar" + "a")) + "\n\t")) + "\n\t"), F1: -1, F2: 42}
                _ = fi13
                _ = v14
            }
            var v15 uint = 2
            _ = v15
        }
        _ = fi9
        _ = v10
    }
    switch "\n\t" {
    case "a":
        v3[v2] = -360
    case "foo bar":
        fmt.Println(v4, v2, v2)
    case "":
        for fi16 := 0; fi16 < 11; fi16++ {
            v17 := []int{}
            v18 := GenS1{F0: -248}
            switch ("\n\t" + ("\n\t" + "\n\t")) {
            case "":
                v19 := GenS3{F0: 18446744073709551615, F1: (v4 / ((v4) | 1)), F2: (v4 ^ (v4 / (((v4 ^ v4)) | 1)))}
                _ = v19
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    default:
        var v20 uint8 = v4
        _ = v20
    }
    v1 = (v1 & v1)
    for fi21 := 0; fi21 < 1; fi21++ {
        v1 = (^(v1))
        _ = fi21
    }
    _ = len(v5)
    v4 = v4
    v4 += (v4 | v4)
    switch v2 {
    case 32768:
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 uint32 = v1
            if ("\n\t" > (("hello" + "\n\t") + (("\n\t" + ("\n\t" + (("a" + ("foo bar" + "hello")) + ("hello" + (("" + "a") + ("foo bar" + "\n\t")))))) + "foo bar"))) {
                v24 := []int{(fi22 + v2), 9223372036854775807, fi22, fi22}
                for fi25 := 0; fi25 < 20; fi25++ {
                    var v26 float32 = ((v8 + v8) / ((v8 / (v8 - v8)) + (1e10 * 1e10)))
                    var v27 float32 = v26
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                v4 -= (v4 / ((v4) | 1))
                _ = v24
            } else {
                v28 := GenS2{F0: "\n\t", F1: int64(v4), F2: -827}
                v4 -= (^(v4))
                _ = v28
            }
            if false {
                var v29 uint8 = 0
                fmt.Println(fi22, fi22)
                _ = v29
            }
            _ = fi22
            _ = v23
        }
    case -780:
        var v30 uint32 = v1
        _ = v30
    case 680:
        v8 = (0.1 * (0.1 - 1.5))
    }
    v1 += (^(v1))
    if (-2147483648 <= -1) {
        v31 := []int{158, (v2 >> ((0) & 63)), v2, (v2 | 133)}
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 uint8 = (^(v4))
            var v34 float64 = 3.14159
            _ = fi32
            _ = v33
            _ = v34
        }
        if (true && true) {
            v1 = 891
        }
        _ = v31
    }
    v35 := GenS4{F0: -874, F1: 20}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v8
}
