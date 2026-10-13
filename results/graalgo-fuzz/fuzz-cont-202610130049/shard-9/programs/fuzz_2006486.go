package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }
type GenS2 struct { F0 uint32; F1 bool }
func genFn0(p0 int8) int32 {
    var v4 uint = 255
    if ((!true) && true) {
        var v5 uint64 = 2
        _ = v5
    } else {
        p0 = p0
        switch -2147483648 {
        case 2147483647:
            var v6 int8 = p0
            _ = v6
        case -2147483648:
            p0 -= (p0 & p0)
        }
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := []int{-1, (fi7 << ((v4) & 63)), 4294967295}
            var v9 int64 = 9223372036854775807
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    v4 += v4
    var v11 uint = v4
    _ = p0
    _ = v11
    _ = v4
    return -2147483648
}

func main() {
    var v1 int64 = -1
    var v2 uint8 = 42
    v3 := make(map[int]bool)
    var v12 uint8 = 127
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 uint8 = (^(v12))
        _ = fi13
        _ = v14
    }
    switch ((("foo bar" + "foo bar") + "") + "a") {
    case "a":
        v15 := []int{4294967295, 127}
        _ = v15
    case "":
        if (true && false) {
            v16 := GenS1{F0: -911, F1: 32767, F2: 504}
            _ = v16
        } else {
            if false {
                var v17 int8 = 127
                for fi18 := 0; fi18 < 2; fi18++ {
                    for fi19 := 0; fi19 < 18; fi19++ {
                        fmt.Println(v12)
                        var v20 float64 = (0.1 / 1e10)
                        _ = fi19
                        _ = v20
                    }
                    _ = fi18
                }
                for fi21 := 0; fi21 < 20; fi21++ {
                    for fi22 := 0; fi22 < 7; fi22++ {
                        var v23 float64 = 3.14159
                        var v24 uint32 = 379
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS1{F0: 862, F1: 0, F2: 255}
                    _ = fi21
                    _ = v25
                }
                _ = v17
            } else {
                v2 = v12
            }
            v12 = (^(v2))
        }
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 int64 = (v1 << ((65535) & 63))
        _ = fi26
        _ = v27
    }
    var v28 uint32 = 911
    fmt.Println(v12, v1, v1)
    v29 := v3[-978]
    v12 = 127
    v30 := GenS1{F0: -194, F1: 770, F2: 42}
    v28 -= 2
    if (true || (127 >= v12)) {
        v28 = uint32(v12)
        fmt.Println(v29, v29, v28)
        switch (("a" + "") + "foo bar") {
        case "":
            for fi31 := 0; fi31 < 2; fi31++ {
                var v32 bool = (18446744073709551615 >= 42)
                _ = fi31
                _ = v32
            }
        case "a":
            var v33 float32 = 1e10
            _ = v33
        default:
            for fi34 := 0; fi34 < 2; fi34++ {
                v12 -= (v12 ^ v12)
                v35 := GenS2{F0: (v28 >> ((1) & 63)), F1: v29}
                _ = fi34
                _ = v35
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
