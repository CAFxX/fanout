package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 uint8; F1 int8; F2 int32 }
func genFn0(p0 int16, p1 uint8, p2 uint64) uint32 {
    var v7 int64 = -1
    if true {
        p0 += p0
    }
    var v8 int16 = -32768
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 812}
        p1 = 176
        p1 = p1
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        v8 -= (p0 << ((4294967295) & 63))
        v7 -= v7
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 uint32, p1 bool) int8 {
    v12 := GenS2{F0: 1}
    if (-1 == 962) {
        v13 := GenS3{F0: p0}
        p1 = (p1 && p1)
        fmt.Println(p0)
        _ = v13
    } else {
        var v14 int = -2147483648
        var v15 uint16 = 42
        _ = v14
        _ = v15
    }
    switch "foo bar" {
    case "hello":
        v16 := []int{-843, 128, 1, 42}
        _ = v16
    case "":
        p0 = p0
    default:
        v17 := []int{-637, -8, 2147483647, 4294967295}
        _ = v17
    }
    var v18 uint = 0
    _ = p0
    _ = p1
    _ = v12
    _ = v18
    return -1
}

func main() {
    var v1 int8 = 65
    var v2 int64 = -166
    var v3 uint8 = 0
    var v4 uint = 1
    var v5 int = 32767
    var v6 uint = v4
    for fi19 := 0; fi19 < 5; fi19++ {
        if (false || false) {
            switch "\n\t" {
            case "a":
                v20 := GenS4{F0: -215}
                _ = v20
            case "hello":
                v5 = 0
            }
            if false {
                v21 := []int{-546, 881, (v5 | v5)}
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 float64 = (3.14159 - ((float64(-2.25) + (-2.25 - (((1.5 * 3.14159) - 1e10) / (0.1 * 3.14159)))) * (-2.25 + (-2.25 - 1.5))))
            v2 = v2
            _ = fi22
            _ = v23
        }
        v2 = v2
        _ = fi19
    }
    switch "a" {
    case "":
        v6 -= 42
    case "a":
        v24 := []int{v5, -1}
        _ = v24
    }
    v25 := GenS1{F0: (v6 + (^(v4)))}
    switch "" {
    case "hello":
        var v26 float64 = -2.25
        _ = v26
    }
    fmt.Println(v1, v1, v4)
    if false {
        switch (v5 | 532) {
        case 42:
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 float64 = 3.14159
                _ = fi27
                _ = v28
            }
        }
    }
    v2 -= v2
    for fi29 := 0; fi29 < 2; fi29++ {
        for fi30 := 0; fi30 < 13; fi30++ {
            switch ("" + "a") {
            case "":
                fmt.Println(v2)
            }
            fmt.Println(v1, v4, v2)
            _ = fi30
        }
        _ = fi29
    }
    if (false || (-70 > 1)) {
        var v31 int = (v5 + (^(v5)))
        _ = v31
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        if true {
            for fi33 := 0; fi33 < 17; fi33++ {
                if (!false) {
                    v34 := GenS1{F0: 2}
                    v6 = v4
                    var v35 int8 = v1
                    _ = v34
                    _ = v35
                }
                if true {
                    v2 -= (v2 - 9223372036854775807)
                }
                _ = fi33
            }
            v36 := GenS5{F0: v5}
            v37 := []int{(^(v5))}
            _ = v36
            _ = v37
        } else {
            var v38 int32 = -964
            var v39 string = ("a" + ("a" + ""))
            _ = v38
            _ = v39
        }
        _ = fi32
    }
    v6 = 255
    v4 -= v4
    for fi40 := 0; fi40 < 4; fi40++ {
        for fi41 := 0; fi41 < 18; fi41++ {
            v3 += (v3 - v3)
            _ = fi41
        }
        v5 += fi40
        v42 := GenS2{F0: v3}
        _ = fi40
        _ = v42
    }
    if true {
        switch v5 {
        case 42:
            v6 = v4
        case 56:
            switch ("hello" + ("a" + ("" + (("\n\t" + "") + "foo bar")))) {
            case "":
                var v43 float64 = 0.0
                _ = v43
            case "foo bar":
                v44 := []int{}
                _ = v44
            case "hello":
                v45 := GenS6{F0: -989, F1: false}
                _ = v45
            }
        default:
            var v46 int32 = -1
            _ = v46
        }
        v4 = v6
    } else {
        v47 := GenS7{F0: v3, F1: genFn1(4294967295, false), F2: 607}
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
