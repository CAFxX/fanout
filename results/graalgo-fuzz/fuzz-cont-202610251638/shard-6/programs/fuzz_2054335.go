package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS6 struct { F0 int32; F1 uint8; F2 uint8 }
func genFn0(p0 int32, p1 int) uint32 {
    if false {
        v5 := GenS1{F0: 4294967295, F1: "hello"}
        p1 = p1
        _ = v5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{fi6, fi6, (p1 | fi6), fi6}
        if true {
            var v8 uint8 = 230
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v8
        }
        p1 = (p1 | (fi6 | (^(p1))))
        _ = fi6
        _ = v7
    }
    var v9 uint64 = 1
    fmt.Println(v9, v9, p0)
    var v10 uint32 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (^(v10))
}
func genFn1(p0 int, p1 uint8) bool {
    switch (p0 << ((255) & 63)) {
    case 42:
        p0 -= (p0 ^ 128)
    case 128:
        fmt.Println(p0, p0)
    default:
        if true {
            v11 := []int{42, p0, (p0 + (p0 % ((p0) | 1))), -2147483648}
            _ = v11
        }
    }
    fmt.Println(p1)
    switch ("a" + ("hello" + ("\n\t" + ("" + "\n\t")))) {
    case "":
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int = p0
            _ = fi12
            _ = v13
        }
    case "hello":
        var v14 int64 = 786
        _ = v14
    case "foo bar":
        var v15 string = "foo bar"
        _ = v15
    default:
        if (true || (42 > 42)) {
            var v16 uint8 = (^(p1))
            if true {
                switch (^(p0)) {
                case 32767:
                    v17 := []int{p0, (p0 / ((p0) | 1))}
                    _ = v17
                case 128:
                    v18 := []int{int(v16), -9223372036854775808, p0}
                    _ = v18
                case 9223372036854775807:
                    v19 := GenS1{F0: 18, F1: "foo bar"}
                    _ = v19
                }
            }
            v20 := GenS2{F0: 255, F1: ("\n\t" < "\n\t")}
            _ = v16
            _ = v20
        }
    }
    v21 := []int{}
    p0 = (p0 << ((65535) & 63))
    _ = p0
    _ = p1
    _ = v21
    return false
}

func main() {
    var v1 int32 = 0
    v2 := []uint32{2147483647}
    var v3 int32 = 681
    v4 := make(map[int]int)
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := GenS3{F0: 4294967295, F1: 32767}
        _ = fi22
        _ = v23
    }
    switch -1 {
    case 634:
        var v24 int8 = 127
        _ = v24
    default:
        v1 -= (v1 | (v3 | -2147483648))
    }
    switch "a" {
    case "hello":
        v1 += (v3 << ((65535) & 63))
    case "foo bar":
        v3 = v1
    case "":
        _ = len(v4)
    default:
        _ = len(v2)
    }
    _ = len(v4)
    v3 -= v3
    if true {
        v25 := GenS4{F0: -1, F1: (v1 - (v1 / ((v3) | 1)))}
        _ = v25
    } else {
        var v26 float64 = 1e10
        switch -9223372036854775808 {
        case -396:
            var v27 string = "a"
            _ = v27
        }
        var v28 uint16 = 301
        _ = v26
        _ = v28
    }
    fmt.Println(v3, v3)
    for fi29 := 0; fi29 < 17; fi29++ {
        v30 := GenS4{F0: 9223372036854775807, F1: (^(v1))}
        if len(v2) > 0 {
            v2[0] = 483
        }
        _ = fi29
        _ = v30
    }
    if true {
        for fi31 := 0; fi31 < 12; fi31++ {
            switch (fi31 & (fi31 | fi31)) {
            case 32767:
                var v32 uint32 = 51
                _ = v32
            }
            v33 := GenS5{F0: 858, F1: 167, F2: 32767}
            var v34 uint = 2
            _ = fi31
            _ = v33
            _ = v34
        }
        var v35 uint64 = 4294967295
        _ = v35
    }
    switch 0 {
    case 9223372036854775807:
        v36 := GenS5{F0: 1, F1: 1, F2: 1}
        _ = v36
    case 2147483647:
        v37 := GenS6{F0: (^(v3)), F1: 1, F2: 223}
        _ = v37
    case 4294967295:
        v3 = v1
    default:
        v3 += (^(v3))
    }
    v3 = (v3 - 42)
    var v38 float64 = 3.14159
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
