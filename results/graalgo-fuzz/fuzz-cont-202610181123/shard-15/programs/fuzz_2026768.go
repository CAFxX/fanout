package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
func genFn0(p0 string, p1 int, p2 int64) int8 {
    switch int(p2) {
    case -2147483648:
        var v4 uint64 = 126
        _ = v4
    case 32768:
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 uint64 = 0
            var v7 uint16 = uint16(v6)
            _ = fi5
            _ = v6
            _ = v7
        }
    case -9223372036854775808:
        v8 := []int{}
        _ = v8
    }
    var v9 int8 = 41
    p2 = (-(p2))
    v9 += v9
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return v9
}
func genFn1() int32 {
    var v10 int8 = 42
    if ((true || (!(135 != 200))) || (1 > 899)) {
        if (!true) {
            for fi11 := 0; fi11 < 12; fi11++ {
                v10 -= (v10 % ((100) | 1))
                if false {
                    v10 -= v10
                    v12 := []int{(^(fi11)), fi11}
                    _ = v12
                }
                v13 := GenS1{F0: false}
                _ = fi11
                _ = v13
            }
        }
        switch 9223372036854775807 {
        case 839:
            v14 := []int{}
            _ = v14
        case 4294967295:
            var v15 int32 = 42
            _ = v15
        }
        v10 = genFn0("hello", 485, 0)
    } else {
        v10 += 42
        var v16 int8 = 78
        for fi17 := 0; fi17 < 18; fi17++ {
            fmt.Println(fi17, v16, v10)
            _ = fi17
        }
        _ = v16
    }
    _ = v10
    return -2147483648
}

func main() {
    v1 := []uint32{702}
    var v2 bool = ((float64(((float32(float64(0.1)) + 0.0) + 1.5)) != 3.14159) || false)
    v3 := make(map[string]bool)
    var v18 uint8 = 114
    v2 = v2
    v19 := GenS2{F0: ""}
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int16 = 32767
            v18 = (v18 - (v18 ^ v18))
            switch "\n\t" {
            case "a":
                var v23 int = (fi20 / ((fi21) | 1))
                _ = v23
            case "\n\t":
                v18 = 128
            }
            var v24 int16 = (^(v22))
            v18 = 239
            _ = fi21
            _ = v22
            _ = v24
        }
        _ = fi20
    }
    v1 = append(v1, 610)
    var v25 uint8 = v18
    switch 9223372036854775807 {
    case 42:
        if v2 {
            _ = len(v1)
            switch "foo bar" {
            case "foo bar":
                v26 := GenS1{F0: v2}
                _ = v26
            case "a":
                v18 = (v25 << ((1) & 63))
            case "\n\t":
                v25 = (v18 | (^(v18)))
            }
        }
    case -2147483648:
        v18 = (v18 * (v18 * v25))
    case 2147483647:
        switch ("a" + "foo bar") {
        case "foo bar":
            switch 128 {
            case 213:
                if len(v1) > 0 {
                    v1[0] = 0
                }
            case 1:
                v25 = v25
            case -9223372036854775808:
                var v27 uint32 = 4294967295
                _ = v27
            }
        }
    }
    _ = len(v1)
    var v28 uint16 = 1
    for fi29 := 0; fi29 < 7; fi29++ {
        if false {
            for fi30 := 0; fi30 < 19; fi30++ {
                fmt.Println(fi30)
                _ = fi30
            }
        }
        v2 = true
        _ = fi29
    }
    v3["hello"] = (v2 || (!v2))
    if v2 {
        v2 = v2
        var v31 bool = (!(1 != -2147483648))
        _ = v31
    } else {
        v18 = (v18 % (((v18 << ((42) & 63))) | 1))
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v28
    _ = v3
}
