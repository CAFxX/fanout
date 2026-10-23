package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 float64 = 1e10
    var v2 int = 445
    var v3 int = v2
    var v4 uint32 = 1
    switch v2 {
    case 981:
        if (false || false) {
            for fi5 := 0; fi5 < 5; fi5++ {
                var v6 int16 = 54
                v7 := []int{1, 32768, 4294967295, 32767}
                v8 := GenS1{F0: 42, F1: (^(fi5))}
                _ = fi5
                _ = v6
                _ = v7
                _ = v8
            }
            v9 := []int{}
            v9 = append(v9, v3)
            _ = v9
        } else {
            var v10 uint64 = 0
            v11 := []int{-2147483648}
            _ = v10
            _ = v11
        }
    case -819:
        var v12 float64 = v1
        _ = v12
    case 128:
        var v13 uint = 966
        _ = v13
    default:
        var v14 uint = 677
        _ = v14
    }
    var v15 int64 = -1
    var v16 int64 = v15
    v16 -= v16
    v3 += v3
    for fi17 := 0; fi17 < 8; fi17++ {
        if (!(1 >= 54)) {
            v15 += v15
            for fi18 := 0; fi18 < 17; fi18++ {
                v19 := GenS2{F0: 2}
                var v20 int8 = -128
                v21 := []int{(fi17 - 2147483647), fi17, (fi18 ^ fi17)}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        } else {
            v16 = 1
        }
        v16 -= (v16 & (v15 >> ((763) & 63)))
        v22 := []int{fi17, 907, fi17}
        _ = fi17
        _ = v22
    }
    var v23 int16 = 0
    var v24 uint16 = 620
    var v25 int = v2
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 int16 = v23
        var v28 string = (("hello" + "foo bar") + ("hello" + ("foo bar" + "\n\t")))
        for fi29 := 0; fi29 < 8; fi29++ {
            v16 = (v16 - (v15 & v16))
            _ = fi29
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    var v30 float32 = 3.14159
    switch v25 {
    case 131:
        if (!true) {
            var v31 int64 = (v15 ^ v15)
            v16 -= (v16 << ((42) & 63))
            v32 := GenS3{F0: 499}
            _ = v31
            _ = v32
        } else {
            v33 := []int{(^(v2)), v25, (-(v25)), -2147483648}
            switch "foo bar" {
            case "a":
                if len(v33) > 0 {
                    v33[0] = (v2 | v25)
                }
            case "":
                v15 -= -1
            }
            _ = v33
        }
    case -2147483648:
        v34 := GenS1{F0: 42, F1: v25}
        _ = v34
    default:
        v35 := GenS2{F0: (v24 ^ (v24 ^ 255))}
        _ = v35
    }
    v36 := []int{(v3 * v2)}
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    if true {
        fmt.Println(v24, v2, v2)
        v16 -= -9223372036854775808
        var v37 int64 = -1
        _ = v37
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        var v39 float32 = 0.1
        fmt.Println(v1, v23)
        var v40 uint64 = 520
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v41
}
