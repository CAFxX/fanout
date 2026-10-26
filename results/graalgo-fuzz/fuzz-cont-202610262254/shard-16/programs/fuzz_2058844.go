package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint64; F2 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
func genFn0(p0 bool) uint32 {
    if (-2.25 >= 0.0) {
        fmt.Println(p0)
    }
    v4 := []int{}
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: fi5, F1: 18446744073709551615}
        var v7 uint32 = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 = (!(-128 <= 81))
    for fi8 := 0; fi8 < 3; fi8++ {
        if len(v4) > 0 {
            v4[0] = fi8
        }
        p0 = (782 > 652)
        var v9 uint8 = 255
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v4
    return 65
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[uint32]uint32)
    var v3 string = "\n\t"
    v3 = (v3 + ("\n\t" + (v3 + v3)))
    v10 := GenS1{F0: v1, F1: 258}
    var v11 uint32 = 2147483648
    fmt.Println(v11, v1, v11)
    if true {
        v12 := GenS1{F0: v1, F1: 0}
        switch "" {
        case "foo bar":
            v13 := v2[(^(v11))]
            _ = v13
        case "\n\t":
            v14 := GenS2{F0: true, F1: 864, F2: -1}
            _ = v14
        case "a":
            delete(v2, genFn0(((true && true) || true)))
        }
        v15 := GenS2{F0: ((!((2 <= 886) || false)) && false), F1: 0, F2: -796}
        _ = v12
        _ = v15
    } else {
        v16 := GenS1{F0: (v1 % ((v1) | 1)), F1: 0}
        for fi17 := 0; fi17 < 19; fi17++ {
            v18 := []int{}
            var v19 float32 = ((0.0 * -2.25) - -2.25)
            var v20 uint16 = 588
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        _ = v16
    }
    _ = len(v2)
    if true {
        v1 = v1
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 uint16 = 32767
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 int8 = -128
            _ = fi23
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    if true {
        if ((1 != 959) || (32767 == 746)) {
            fmt.Println(v11, v3)
            v25 := GenS2{F0: true, F1: 336, F2: 42}
            _ = v25
        } else {
            var v26 int8 = 22
            for fi27 := 0; fi27 < 20; fi27++ {
                v2[v11] = (v11 * (v11 & v11))
                var v28 int64 = 265
                _ = fi27
                _ = v28
            }
            _ = v26
        }
    } else {
        fmt.Println(v11, v11)
    }
    var v29 int = -1
    v30 := []int{-2147483648, (v1 >> ((255) & 63))}
    if true {
        v29 += -1
        var v31 string = v3
        var v32 float32 = float32(0.1)
        _ = v31
        _ = v32
    } else {
        v33 := GenS3{F0: 131}
        if (true || true) {
            var v34 bool = false
            v34 = (-32768 >= 32767)
            for fi35 := 0; fi35 < 7; fi35++ {
                v36 := v2[genFn0(false)]
                _ = fi35
                _ = v36
            }
            _ = v34
        } else {
            var v37 float64 = ((-2.25 + 0.0) - 0.1)
            if (4 < 18446744073709551615) {
                v3 = ("\n\t" + ((v3 + (((("a" + ((v3 + v3) + "\n\t")) + (v3 + v3)) + (v3 + v3)) + "foo bar")) + ((((((v3 + v3) + v3) + v3) + ("" + v3)) + v3) + ((v3 + (("foo bar" + "a") + v3)) + v3))))
            }
            if true {
                v38 := GenS4{F0: 621}
                _ = v38
            } else {
                var v39 int32 = -2147483648
                _ = v39
            }
            _ = v37
        }
        _ = v33
    }
    v11 += (v11 * (^(v11)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
