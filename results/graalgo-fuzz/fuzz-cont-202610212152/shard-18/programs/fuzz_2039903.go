package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int64; F2 uint64 }
type GenS4 struct { F0 int8 }
func genFn0() uint {
    if (128 != 628) {
        var v4 float32 = (0.1 / (1e10 * 3.14159))
        if true {
            var v5 int64 = -1
            if false {
                switch "\n\t" {
                case "foo bar":
                    v6 := []int{118, 4294967295, 883}
                    _ = v6
                }
            } else {
                var v7 uint64 = 42
                _ = v7
            }
            _ = v5
        }
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(v4)
            var v9 uint16 = 0
            switch "a" {
            case "a":
                var v10 string = ("" + "a")
                _ = v10
            case "hello":
                v11 := []int{}
                _ = v11
            case "foo bar":
                var v12 bool = ((v9 - (v9 * (^(v9)))) < v9)
                _ = v12
            }
            _ = fi8
            _ = v9
        }
        _ = v4
    }
    var v13 int32 = 42
    fmt.Println(v13)
    v14 := GenS1{F0: 230}
    switch -9223372036854775808 {
    case 1:
        v15 := []int{}
        _ = v15
    case 2147483647:
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 int16 = -62
            _ = fi16
            _ = v17
        }
    case 9223372036854775807:
        for fi18 := 0; fi18 < 8; fi18++ {
            for fi19 := 0; fi19 < 5; fi19++ {
                if false {
                    fmt.Println(fi19, fi18, fi19)
                    v13 -= 148
                }
                _ = fi19
            }
            _ = fi18
        }
    }
    _ = v13
    _ = v14
    return 1
}
func genFn1() int16 {
    var v20 float64 = ((-2.25 + (-2.25 * (0.0 * 1e10))) + (1e10 * 1.5))
    for fi21 := 0; fi21 < 2; fi21++ {
        v20 = (v20 - float64(1.5))
        _ = fi21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            switch (-(fi22)) {
            case 2147483647:
                v20 = float64(1e10)
            default:
                if (float64((3.14159 - ((1e10 - (-2.25 * (((1.5 / 3.14159) + (1.5 + 1e10)) * ((1e10 * (0.0 + -2.25)) * (((-2.25 + 1e10) / 1.0) + ((-2.25 * 0.0) - 3.14159)))))) + 1.5))) >= v20) {
                    var v24 int = fi22
                    _ = v24
                }
                fmt.Println(fi22)
                if (42 == 9223372036854775807) {
                    v20 = (0.1 * (1.5 * v20))
                    v25 := GenS1{F0: fi23}
                    _ = v25
                } else {
                    var v26 string = "a"
                    var v27 uint = 1
                    _ = v26
                    _ = v27
                }
            }
            var v28 int8 = 39
            v29 := []int{}
            _ = fi23
            _ = v28
            _ = v29
        }
        _ = fi22
    }
    _ = v20
    return -753
}

func main() {
    var v1 int16 = -8
    v2 := []int{42, 128, 763}
    var v3 string = "a"
    v1 = v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v30 := GenS2{F0: (v1 % ((v1) | 1))}
    if (("foo bar" + v3) > "foo bar") {
        v3 = "\n\t"
        v31 := GenS3{F0: ((((v3 + ("hello" + v3)) + ("" + "foo bar")) + ("" + "foo bar")) + "a"), F1: -572, F2: 18446744073709551615}
        var v32 bool = ((false && (1 < 42)) && (42 >= 187))
        _ = v31
        _ = v32
    }
    v1 += (v1 % ((1000) | 1))
    var v33 uint32 = 304
    v1 -= 685
    fmt.Println(v3, v1)
    _ = len(v2)
    if len(v2) > 0 {
        v2[0] = -2147483648
    }
    v33 = v33
    if (true || (genFn1() == v1)) {
        var v34 uint8 = 1
        _ = v34
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!(1 == 126)) {
        fmt.Println(v1)
        fmt.Println(v1)
        v33 = 4294967295
    } else {
        v3 = v3
        var v35 bool = (1 > 1)
        v36 := GenS4{F0: 42}
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        v1 = v1
        var v38 string = v3
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
}
