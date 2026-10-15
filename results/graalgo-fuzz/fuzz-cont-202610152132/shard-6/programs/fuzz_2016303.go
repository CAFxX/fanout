package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 bool, p1 uint64, p2 int16) int64 {
    switch "" {
    case "\n\t":
        switch "hello" {
        case "hello":
            v4 := []int{42}
            _ = v4
        }
    case "hello":
        var v5 int32 = -1
        _ = v5
    default:
        v6 := []int{-9223372036854775808, 0}
        _ = v6
    }
    var v7 int8 = -128
    p0 = ((p0 || ((!false) || ((1 == 979) && false))) && ((p2 >> ((1) & 63)) == 8))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 int64, p1 int64) int16 {
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 float32 = 3.14159
            _ = fi9
            _ = v10
        }
        fmt.Println(p1, p1)
        p1 -= (p0 % (((^(p0))) | 1))
        _ = fi8
    }
    if false {
        var v11 uint16 = 667
        _ = v11
    } else {
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := GenS1{F0: -960}
            var v14 uint = 0
            v15 := []int{(fi12 / (((^(fi12))) | 1)), fi12, fi12, (fi12 / ((fi12) | 1))}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        var v16 uint16 = 32767
        v17 := []int{128, 9223372036854775807}
        _ = v16
        _ = v17
    }
    var v18 int8 = 77
    var v19 float32 = 0.1
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return -529
}

func main() {
    var v1 uint32 = 65535
    var v2 int64 = 1
    var v3 int32 = 850
    v20 := GenS1{F0: (v3 >> ((uint(v1)) & 63))}
    var v21 string = ("a" + "foo bar")
    v22 := GenS1{F0: 756}
    v2 -= (v2 | genFn0(true, 0, 246))
    var v23 uint = 2
    var v24 uint32 = (v1 >> ((v23) & 63))
    var v25 uint64 = 1
    for fi26 := 0; fi26 < 6; fi26++ {
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 int8 = 42
            switch (v21 + "foo bar") {
            case "hello":
                v29 := GenS2{F0: (v28 | -1)}
                _ = v29
            case "foo bar":
                switch 128 {
                case -253:
                    v30 := []int{(^(fi26)), fi27, (fi26 >> (((^(v23))) & 63)), 4294967295}
                    _ = v30
                default:
                    v31 := []int{(fi27 * 1), (fi26 ^ (fi27 ^ (^(fi26))))}
                    _ = v31
                }
            default:
                var v32 float64 = 3.14159
                _ = v32
            }
            v33 := []int{fi26, fi27}
            _ = fi27
            _ = v28
            _ = v33
        }
        v34 := []int{fi26, (fi26 | (fi26 << ((uint(v1)) & 63))), (fi26 | (fi26 & (-(fi26)))), fi26}
        for fi35 := 0; fi35 < 2; fi35++ {
            var v36 uint16 = 466
            var v37 uint32 = v1
            fmt.Println(fi26, v3)
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi26
        _ = v34
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        v2 -= genFn0((!(2147483648 >= uint32(v25))), (^(v25)), 0)
        var v39 uint16 = 2
        _ = fi38
        _ = v39
    }
    if false {
        for fi40 := 0; fi40 < 14; fi40++ {
            v1 = v24
            _ = fi40
        }
        v41 := []int{198}
        _ = v41
    } else {
        if true {
            var v42 int8 = 112
            v43 := GenS2{F0: v42}
            fmt.Println(v24, v3)
            _ = v42
            _ = v43
        }
    }
    var v44 bool = true
    var v45 uint64 = (v25 % ((v25) | 1))
    var v46 bool = (v21 > "")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v44
    _ = v45
    _ = v46
}
