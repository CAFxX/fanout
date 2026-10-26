package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int64, p2 int32) int16 {
    v6 := []int{0, 128, -1}
    var v7 int8 = 0
    var v8 uint16 = (p0 << ((364) & 63))
    switch 2147483647 {
    case -1:
        var v9 int64 = p1
        _ = v9
    default:
        if (2 == 42) {
            p1 -= -221
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 673
}
func genFn1(p0 int, p1 uint32, p2 uint8) uint32 {
    v10 := GenS1{F0: 42}
    fmt.Println(p1, p1)
    if true {
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint8 = 128
            fmt.Println(v12, p0)
            _ = fi11
            _ = v12
        }
        var v13 int64 = 523
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := GenS1{F0: 42}
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        v16 := []int{}
        _ = v16
    }
    if true {
        var v17 float32 = 3.14159
        p2 += p2
        if (!((true || false) || false)) {
            fmt.Println(p0, p1, p0)
        }
        _ = v17
    }
    var v18 int8 = int8(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v18
    return 1
}

func main() {
    var v1 int64 = -951
    var v2 float32 = 3.14159
    var v3 int16 = 1000
    v4 := []int{}
    var v5 int8 = 0
    var v19 uint8 = 2
    var v20 int8 = v5
    var v21 uint16 = 454
    v20 = v20
    if (!(802 <= 2)) {
        v22 := GenS1{F0: 1}
        v4 = append(v4, -9223372036854775808)
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 int8 = (v5 & int8(fi23))
            v2 = v2
            v25 := GenS1{F0: 1}
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v22
    }
    v2 += v2
    switch 1 {
    case 32768:
        for fi26 := 0; fi26 < 15; fi26++ {
            for fi27 := 0; fi27 < 1; fi27++ {
                v5 += int8(v1)
                v4 = append(v4, fi26)
                switch "\n\t" {
                case "":
                    var v28 uint64 = 42
                    _ = v28
                case "hello":
                    v4 = append(v4, -420)
                }
                _ = fi27
            }
            v29 := GenS1{F0: 1}
            _ = fi26
            _ = v29
        }
    case 9223372036854775807:
        var v30 int16 = v3
        _ = v30
    case 147:
        v3 = -897
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        if (!true) {
            v21 += (v21 & 119)
        } else {
            switch -767 {
            case -838:
                for fi32 := 0; fi32 < 20; fi32++ {
                    var v33 bool = (!((("hello" + "foo bar") + "foo bar") >= "hello"))
                    _ = len(v4)
                    v1 += (v1 >> ((42) & 63))
                    _ = fi32
                    _ = v33
                }
            case 128:
                switch "foo bar" {
                case "a":
                    v1 -= (v1 >> ((364) & 63))
                }
            default:
                v20 = v5
            }
            switch fi31 {
            case 32768:
                v21 += v21
            }
        }
        fmt.Println(fi31, v5)
        var v34 string = "a"
        _ = fi31
        _ = v34
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        var v36 int32 = 2147483647
        v37 := GenS2{F0: 4294967295}
        var v38 uint = uint(v21)
        _ = fi35
        _ = v36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 19; fi39++ {
        v21 -= (v21 / ((65535) | 1))
        var v40 float32 = (v2 + v2)
        v41 := GenS2{F0: genFn1(-2147483648, 284, (^(v19)))}
        _ = fi39
        _ = v40
        _ = v41
    }
    if (591 < 42) {
        if ((^(v21)) == 42) {
            v42 := GenS1{F0: 255}
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
