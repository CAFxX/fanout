package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int }
func genFn0() int8 {
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: 4294967295, F1: true, F2: 596}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(fi6)
        if false {
            var v7 uint16 = 311
            var v8 uint8 = 255
            _ = v7
            _ = v8
        } else {
            v9 := GenS1{F0: 416, F1: true, F2: (fi6 & fi6)}
            v10 := []int{32767}
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 uint16 = 65535
                _ = fi11
                _ = v12
            }
            _ = v10
            _ = v9
        }
        _ = fi6
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        if (-2147483648 <= 905) {
            v14 := []int{(-(fi13)), 1, fi13}
            _ = v14
        }
        _ = fi13
    }
    return -34
}
func genFn1() int64 {
    if true {
        v15 := GenS1{F0: 42, F1: false, F2: 1}
        v16 := []int{-703, 32768, 237}
        _ = v15
        _ = v16
    }
    var v17 int64 = 498
    var v18 uint16 = 255
    var v19 int8 = -128
    _ = v17
    _ = v18
    _ = v19
    return (v17 / (((v17 + v17)) | 1))
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{65535, 8}
    var v3 uint16 = 378
    if (1 == 1) {
        v3 -= (v3 | v3)
    } else {
        v20 := GenS1{F0: 18446744073709551615, F1: (!false), F2: -1}
        _ = v20
    }
    var v21 uint32 = 2
    v3 -= 255
    v21 = (v21 % ((v21) | 1))
    fmt.Println(v3, v3)
    if (32767 > 350) {
        v22 := GenS1{F0: 42, F1: true, F2: -9223372036854775808}
        if (65535 < v3) {
            switch "a" {
            case "":
                var v23 int16 = 358
                _ = v23
            case "hello":
                fmt.Println(v21, v21)
            default:
                fmt.Println(v3, v3)
            }
        }
        _ = v22
    } else {
        var v24 bool = true
        v25 := GenS1{F0: 0, F1: v24, F2: -9223372036854775808}
        for fi26 := 0; fi26 < 1; fi26++ {
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 string = "\n\t"
                _ = fi27
                _ = v28
            }
            if v24 {
                if v24 {
                    for fi29 := 0; fi29 < 5; fi29++ {
                        v24 = (1e10 != (0.0 / ((0.1 * (1e10 - 3.14159)) + float32(float64(0.1)))))
                        var v30 bool = false
                        _ = fi29
                        _ = v30
                    }
                }
                _ = len(v2)
            }
            fmt.Println(fi26, fi26)
            _ = fi26
        }
        _ = v24
        _ = v25
    }
    if true {
        var v31 string = (("foo bar" + "\n\t") + ((((((("" + "foo bar") + (((("hello" + "hello") + "a") + "foo bar") + "foo bar")) + ("hello" + "a")) + "a") + ("\n\t" + "foo bar")) + ("foo bar" + "")) + ("\n\t" + "")))
        var v32 bool = (true || false)
        var v33 int64 = 0
        _ = v31
        _ = v32
        _ = v33
    }
    v21 = v21
    v3 = v3
    var v34 int8 = genFn0()
    if true {
        var v35 uint8 = 42
        _ = v35
    } else {
        v36 := GenS1{F0: 18446744073709551615, F1: (!false), F2: 128}
        for fi37 := 0; fi37 < 6; fi37++ {
            fmt.Println(v21, fi37)
            var v38 int8 = v34
            v39 := GenS1{F0: 1, F1: (4294967295 >= 823), F2: (fi37 * -2147483648)}
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = v36
    }
    var v40 uint64 = 4294967295
    fmt.Println(v34, v3)
    for fi41 := 0; fi41 < 5; fi41++ {
        for fi42 := 0; fi42 < 10; fi42++ {
            var v43 int32 = 1
            var v44 float64 = 1.5
            _ = fi42
            _ = v43
            _ = v44
        }
        fmt.Println(v40, v40)
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v40
}
