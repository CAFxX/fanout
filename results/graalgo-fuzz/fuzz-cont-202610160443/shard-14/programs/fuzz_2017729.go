package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int64; F1 string; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 uint8 }
func genFn0(p0 int16) int64 {
    v7 := GenS1{F0: 128, F1: 63}
    switch "\n\t" {
    case "a":
        for fi8 := 0; fi8 < 11; fi8++ {
            if true {
                var v9 uint64 = 2
                fmt.Println(v9, fi8)
                v10 := []int{fi8, (fi8 << ((255) & 63))}
                _ = v10
                _ = v9
            }
            v11 := []int{(^(fi8)), 1, (fi8 - fi8), fi8}
            _ = fi8
            _ = v11
        }
    }
    _ = p0
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint32, p1 int, p2 uint32) string {
    v12 := GenS1{F0: 986, F1: p1}
    v13 := GenS2{F0: 42, F1: ("a" + "hello"), F2: 816}
    v14 := []int{(p1 * (p1 & p1)), (^(p1)), -958}
    v14 = append(v14, -936)
    if ("a" != "\n\t") {
        if (0.1 >= (1e10 / 1e10)) {
            v14 = append(v14, p1)
            var v15 int32 = -57
            switch (p1 * p1) {
            case -559:
                fmt.Println(p1)
            case -458:
                _ = len(v14)
            case 0:
                var v16 bool = false
                _ = v16
            default:
                v15 = (v15 % (((v15 >> ((867) & 63))) | 1))
            }
            _ = v15
        }
        switch p1 {
        case 248:
            p2 = p0
        default:
            p2 = p0
        }
        for fi17 := 0; fi17 < 7; fi17++ {
            p0 += p2
            _ = fi17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return (("a" + "\n\t") + "\n\t")
}

func main() {
    v1 := make(map[string]int)
    var v2 string = ("hello" + ("\n\t" + "a"))
    v3 := make(map[int]string)
    var v4 uint32 = 0
    v5 := []uint32{(v4 >> ((1) & 63)), 65535, v4}
    v6 := []uint32{(v4 % ((v4) | 1)), (v4 * v4)}
    v4 -= v4
    var v18 float32 = -2.25
    for fi19 := 0; fi19 < 10; fi19++ {
        v20 := GenS3{F0: 255, F1: 666}
        _ = fi19
        _ = v20
    }
    var v21 string = v2
    for fi22 := 0; fi22 < 5; fi22++ {
        v2 = "foo bar"
        var v23 float64 = (float64(((v18 / (1e10 - v18)) * (v18 - 3.14159))) - ((0.1 / -2.25) * 0.0))
        _ = fi22
        _ = v23
    }
    v18 -= v18
    var v24 uint64 = 18446744073709551615
    switch -2147483648 {
    case 0:
        if (!(true && false)) {
            v25 := GenS4{F0: v21}
            var v26 int8 = 100
            _ = v25
            _ = v26
        }
    case -1:
        fmt.Println(v4, v24)
    case 2147483647:
        v27 := GenS3{F0: 42, F1: v24}
        _ = v27
    default:
        switch 32768 {
        case 148:
            if (v18 < 1e10) {
                if false {
                    v24 -= (v24 - v24)
                } else {
                    v6 = append(v6, v4)
                }
                switch 634 {
                case 4294967295:
                    v24 -= (v24 - (v24 + v24))
                }
            }
        }
    }
    v28 := GenS5{F0: 0, F1: 41}
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 int32 = -2147483648
        v30 += (v30 >> ((196) & 63))
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        for fi32 := 0; fi32 < 1; fi32++ {
            switch 2147483647 {
            case 0:
                var v33 int64 = -896
                _ = v33
            }
            _ = fi32
        }
        _ = fi31
    }
    v34 := v3[-9223372036854775808]
    var v35 float32 = v18
    v24 -= 796
    var v36 float32 = 0.0
    if (66 != 0) {
        delete(v3, 437)
    } else {
        for fi37 := 0; fi37 < 15; fi37++ {
            for fi38 := 0; fi38 < 14; fi38++ {
                fmt.Println(v21, v34)
                _ = fi38
            }
            v18 = v35
            v24 = 953
            _ = fi37
        }
        switch 4294967295 {
        case -2147483648:
            var v39 int64 = -496
            _ = v39
        case 0:
            var v40 uint16 = 547
            _ = v40
        }
        v34 = ("" + v2)
    }
    var v41 uint8 = 2
    var v42 int64 = genFn0(46)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v24
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
