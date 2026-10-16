package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint }
func genFn0(p0 uint) bool {
    p0 = p0
    fmt.Println(p0)
    if true {
        for fi4 := 0; fi4 < 11; fi4++ {
            p0 += (p0 | p0)
            for fi5 := 0; fi5 < 6; fi5++ {
                p0 = (p0 + (p0 % (((p0 & p0)) | 1)))
                _ = fi5
            }
            for fi6 := 0; fi6 < 3; fi6++ {
                if false {
                    v7 := []int{(fi6 | (^(fi4))), (fi6 & (fi6 / (((fi4 + (-(fi4)))) | 1))), fi4, fi6}
                    v8 := GenS1{F0: 384, F1: 100, F2: p0}
                    _ = v7
                    _ = v8
                }
                _ = fi6
            }
            _ = fi4
        }
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint = 4294967295
            _ = fi9
            _ = v10
        }
    } else {
        var v11 uint16 = 396
        _ = v11
    }
    if (150 <= 42) {
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 float32 = 3.14159
            fmt.Println(fi12)
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := []int{}
            p0 += 65535
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = fi14
            _ = v15
        }
    }
    _ = p0
    return (1 > 183)
}

func main() {
    var v1 uint32 = 4294967295
    var v2 string = ("" + "")
    var v3 bool = (true || false)
    fmt.Println(v3, v2)
    v2 = v2
    v16 := GenS1{F0: 42, F1: 1, F2: 232}
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 uint8 = 128
        _ = fi17
        _ = v18
    }
    v3 = v3
    v2 = ("\n\t" + ("hello" + "\n\t"))
    v19 := []int{-2147483648, 127, 472}
    for fi20 := 0; fi20 < 7; fi20++ {
        for fi21 := 0; fi21 < 16; fi21++ {
            v22 := GenS1{F0: 548, F1: 100, F2: 0}
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            if false {
                var v23 uint16 = 138
                var v24 int16 = int16(v1)
                _ = v23
                _ = v24
            }
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    fmt.Println(v1, v2, v2)
    if (v3 || v3) {
        if (!(true || (!false))) {
            if v3 {
                v1 += v1
                var v25 bool = (!v3)
                _ = v25
            }
        } else {
            v26 := GenS1{F0: 1, F1: 127, F2: 2}
            if v3 {
                v2 = v2
            }
            _ = v26
        }
    } else {
        v27 := GenS1{F0: -1, F1: 42, F2: 326}
        v19 = append(v19, 32767)
        v28 := GenS1{F0: -1, F1: 65, F2: 65535}
        _ = v27
        _ = v28
    }
    v2 = "a"
    v19 = append(v19, -2147483648)
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 string = v2
        _ = fi29
        _ = v30
    }
    if v3 {
        fmt.Println(v3)
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
