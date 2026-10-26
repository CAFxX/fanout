package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int64 }
type GenS2 struct { F0 int64; F1 uint32; F2 int }
func genFn0(p0 int64, p1 bool) uint {
    for fi4 := 0; fi4 < 20; fi4++ {
        switch fi4 {
        case 2147483647:
            var v5 int = (fi4 / (((fi4 - (fi4 / (((fi4 % ((-1) | 1))) | 1)))) | 1))
            _ = v5
        case 484:
            if false {
                switch fi4 {
                case 128:
                    var v6 int64 = p0
                    _ = v6
                case 32767:
                    p1 = (42 != 745)
                case -1:
                    var v7 uint64 = 42
                    _ = v7
                }
                for fi8 := 0; fi8 < 3; fi8++ {
                    fmt.Println(fi8, fi8, fi8)
                    var v9 int64 = (p0 - (p0 / (((p0 * p0)) | 1)))
                    _ = fi8
                    _ = v9
                }
                v10 := []int{(fi4 | fi4), 1, (fi4 - (fi4 & (-(fi4))))}
                _ = v10
            } else {
                v11 := GenS1{F0: 483, F1: -1000, F2: p0}
                var v12 uint64 = 117
                _ = v11
                _ = v12
            }
        }
        var v13 string = ""
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 string = "hello"
            if p1 {
                fmt.Println(v15, fi4, v13)
                var v16 uint8 = 0
                v17 := []int{(fi14 ^ (fi4 % ((189) | 1))), -870, fi4, (^(fi14))}
                _ = v16
                _ = v17
            } else {
                var v18 uint32 = 1
                v15 = "hello"
                var v19 uint32 = (v18 - v18)
                _ = v18
                _ = v19
            }
            var v20 bool = (!(853 > 179))
            _ = fi14
            _ = v15
            _ = v20
        }
        _ = fi4
        _ = v13
    }
    p1 = false
    switch "foo bar" {
    case "foo bar":
        switch 181 {
        case 2147483647:
            v21 := []int{}
            _ = v21
        }
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint16 = 32767
    var v3 string = "a"
    v3 = ("" + ("hello" + v3))
    v22 := GenS1{F0: 42, F1: -32768, F2: 154}
    for fi23 := 0; fi23 < 1; fi23++ {
        v3 = v3
        v24 := GenS2{F0: -1, F1: 765, F2: (fi23 | (fi23 * (fi23 ^ (fi23 % ((fi23) | 1)))))}
        v2 = v2
        _ = fi23
        _ = v24
    }
    var v25 float32 = 1.5
    var v26 uint8 = 42
    v27 := []int{}
    v28 := GenS2{F0: 42, F1: 2147483648, F2: 4294967295}
    v25 = v25
    var v29 float32 = 0.0
    var v30 int = 32767
    v3 = (("foo bar" + (((("\n\t" + v3) + (v3 + v3)) + v3) + "\n\t")) + "foo bar")
    if (false && true) {
        switch (v30 % ((v30) | 1)) {
        case 1:
            var v31 int8 = -128
            _ = v31
        case -816:
            var v32 uint16 = (v2 ^ v2)
            _ = v32
        case 436:
            if (!(0 <= 42)) {
                for fi33 := 0; fi33 < 2; fi33++ {
                    v26 += (v26 & 42)
                    switch ("" + v3) {
                    case "foo bar":
                        v34 := GenS1{F0: 0, F1: 32767, F2: -825}
                        _ = v34
                    }
                    fmt.Println(v25, fi33)
                    _ = fi33
                }
                v1 += v1
            } else {
                var v35 int = int(v2)
                var v36 uint32 = 1
                _ = v35
                _ = v36
            }
        default:
            v25 = (((v29 + 1.5) / ((((((float32(0.0) + v29) / v29) / v29) / v29) * 1.5) / v29)) + v25)
        }
    } else {
        if false {
            var v37 int16 = -785
            var v38 uint32 = 935
            _ = v37
            _ = v38
        }
    }
    fmt.Println(v2, v29, v29)
    var v39 int64 = 356
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v39
}
