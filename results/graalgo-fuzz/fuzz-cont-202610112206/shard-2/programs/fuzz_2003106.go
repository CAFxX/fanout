package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint16 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        if false {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 int8 = 1
                var v10 string = ("foo bar" + "")
                v11 := GenS1{F0: int64(fi7), F1: 621}
                fmt.Println(v10, fi7, fi8)
                v12 := GenS1{F0: -9223372036854775808, F1: 32767}
                _ = fi8
                _ = v10
                _ = v11
                _ = v12
                _ = v9
            }
            var v13 float32 = 1e10
            var v14 int8 = 116
            _ = v13
            _ = v14
        } else {
            switch "\n\t" {
            case "\n\t":
                switch "a" {
                case "":
                    var v15 uint32 = 35
                    _ = v15
                case "\n\t":
                    v16 := GenS1{F0: -580, F1: -1}
                    _ = v16
                }
            }
            v17 := []int{42, (fi7 | fi7), (^(fi7)), 128}
            _ = v17
        }
        if (18446744073709551615 >= 768) {
            v18 := GenS2{F0: 519}
            var v19 uint32 = 345
            for fi20 := 0; fi20 < 1; fi20++ {
                var v21 int16 = 1
                var v22 int32 = 49
                if (((-(fi7)) != fi20) || ((v21 / ((268) | 1)) == 322)) {
                    v23 := []int{}
                    var v24 uint = uint(v19)
                    var v25 int = (^(fi20))
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    var v26 int = (fi20 & (fi20 >> ((255) & 63)))
                    v19 = v19
                    _ = v26
                }
                if (4294967295 <= 1) {
                    v27 := GenS2{F0: 1}
                    fmt.Println(v21)
                    _ = v27
                } else {
                    var v28 int32 = v22
                    v21 = v21
                    v29 := []int{fi20, (fi7 + fi20)}
                    fmt.Println(fi7, v28, v28)
                    _ = v28
                    _ = v29
                }
                v30 := []int{32767, fi20}
                _ = fi20
                _ = v21
                _ = v22
                _ = v30
            }
            _ = v18
            _ = v19
        }
        _ = fi7
    }
    var v31 int = -2147483648
    if (!(true && ((0.1 * 0.0) <= ((-2.25 * 1.5) + (1.5 - -2.25))))) {
        for fi32 := 0; fi32 < 5; fi32++ {
            v31 += v31
            _ = fi32
        }
        for fi33 := 0; fi33 < 18; fi33++ {
            v31 -= 42
            var v34 bool = true
            _ = fi33
            _ = v34
        }
    }
    fmt.Println(v31)
    fmt.Println(v31)
    _ = v31
    return -383
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    var v3 uint64 = 2
    v4 := make(map[int]uint32)
    var v5 uint16 = 32767
    var v6 float64 = (1.5 * (0.1 * float64(1.5)))
    if true {
        for fi35 := 0; fi35 < 4; fi35++ {
            fmt.Println(v3)
            _ = fi35
        }
        v36 := GenS1{F0: 42, F1: 1000}
        _ = v36
    } else {
        var v37 int8 = -125
        var v38 float64 = v6
        var v39 string = "\n\t"
        _ = v37
        _ = v38
        _ = v39
    }
    v5 -= (^(v5))
    v40 := []int{4294967295, 32768}
    if (true || true) {
        v5 = (v5 & v5)
        var v41 uint16 = 65535
        _ = v41
    }
    if len(v40) > 0 {
        v40[0] = -1
    }
    v6 += float64(0.1)
    var v42 int16 = 552
    switch 128 {
    case 0:
        v40 = append(v40, 2147483647)
    case 1:
        v5 = (v5 & (v5 % ((v5) | 1)))
    case 127:
        fmt.Println(v6)
    }
    v5 = v5
    var v43 uint8 = 128
    var v44 float64 = (0.0 / (-2.25 * -2.25))
    for fi45 := 0; fi45 < 3; fi45++ {
        if (!true) {
            fmt.Println(v42, v5)
            for fi46 := 0; fi46 < 19; fi46++ {
                switch (fi45 - (^(fi46))) {
                case -404:
                    var v47 uint16 = v5
                    switch ("foo bar" + ("\n\t" + ("hello" + (("a" + "a") + "hello")))) {
                    case "\n\t":
                        v47 += 32767
                    case "":
                        var v48 int32 = -657
                        _ = v48
                    }
                    v1[-2147483648] = 2147483648
                    _ = v47
                case 128:
                    var v49 int8 = -23
                    _ = v49
                default:
                    var v50 uint = 42
                    _ = v50
                }
                _ = len(v1)
                _ = fi46
            }
        }
        _ = fi45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
