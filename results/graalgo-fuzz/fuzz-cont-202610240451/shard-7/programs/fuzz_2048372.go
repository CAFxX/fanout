package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int16) int32 {
    var v7 float64 = 1e10
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint16 = 2
        _ = fi8
        _ = v9
    }
    if false {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 float64 = float64((3.14159 / 1e10))
            for fi12 := 0; fi12 < 4; fi12++ {
                fmt.Println(fi12, v7, v11)
                var v13 int = (fi10 | (fi12 * (-(fi10))))
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        if false {
            for fi14 := 0; fi14 < 17; fi14++ {
                v15 := []int{(fi14 + fi14)}
                var v16 bool = (323 >= 42)
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 uint8 = 42
                v19 := GenS1{F0: ("a" + "foo bar")}
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 uint8 = 2
            fmt.Println(v20, v20)
            _ = v20
        }
        fmt.Println(p0)
    }
    v21 := []int{336}
    var v22 int64 = -9223372036854775808
    _ = p0
    _ = v21
    _ = v22
    _ = v7
    return -228
}
func genFn1(p0 uint, p1 int64) int {
    v23 := []int{32768, 42}
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 string = "hello"
        p0 -= (p0 ^ (p0 * (p0 << ((145) & 63))))
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        fmt.Println(p0, p1, p1)
        _ = fi26
    }
    _ = len(v23)
    for fi27 := 0; fi27 < 15; fi27++ {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        var v28 int8 = 0
        _ = fi27
        _ = v28
    }
    _ = p0
    _ = p1
    _ = v23
    return 1
}

func main() {
    var v1 int32 = -2147483648
    var v2 bool = ((32768 >= -9223372036854775808) || (false && (!(!false))))
    v3 := []int{-187, -9223372036854775808, 1, -1}
    var v4 bool = true
    v5 := make(map[uint32]uint32)
    var v6 int32 = v1
    var v29 int8 = -78
    var v30 uint32 = 479
    v4 = (v4 && (!v2))
    v4 = v4
    v30 -= (v30 - (v30 - 757))
    for fi31 := 0; fi31 < 20; fi31++ {
        v4 = ((v6 - genFn0(-1)) >= genFn0(1000))
        switch fi31 {
        case 9223372036854775807:
            v6 += 0
        case 127:
            _ = len(v5)
        default:
            _ = len(v3)
        }
        _ = fi31
    }
    fmt.Println(v29, v6, v6)
    if (!v4) {
        v30 = (v30 >> ((954) & 63))
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 int8 = (v29 * 1)
            _ = fi32
            _ = v33
        }
    } else {
        v30 = (v30 - (v30 / ((1) | 1)))
        if (true || (!false)) {
            if v4 {
                _ = len(v5)
                fmt.Println(v1)
                v6 = v6
            }
            switch int(v1) {
            case -422:
                _ = len(v5)
            default:
                v34 := GenS1{F0: "a"}
                _ = v34
            }
            for fi35 := 0; fi35 < 19; fi35++ {
                var v36 bool = true
                _ = fi35
                _ = v36
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
