package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8 }
func genFn0(p0 string) uint8 {
    if (958 != 984) {
        var v6 float64 = float64(3.14159)
        v7 := []int{0, 127}
        switch p0 {
        case "a":
            v6 -= (1e10 / v6)
        case "\n\t":
            var v8 uint64 = 742
            _ = v8
        }
        _ = v6
        _ = v7
    }
    if (false || (true || false)) {
        var v9 int64 = 42
        fmt.Println(p0, p0)
        _ = v9
    }
    var v10 uint64 = 42
    v11 := GenS1{F0: (-643 > -2147483648)}
    v12 := GenS1{F0: false}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return 165
}
func genFn1(p0 uint64) bool {
    p0 += p0
    v13 := GenS1{F0: (false || ((float32(-2.25) - 0.1) >= float32(1e10)))}
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint8 = 1
        _ = fi14
        _ = v15
    }
    if true {
        v16 := []int{2147483647}
        switch "" {
        case "hello":
            for fi17 := 0; fi17 < 17; fi17++ {
                v18 := GenS1{F0: (!(((453 == 42) && (false && true)) && (!(!true))))}
                _ = fi17
                _ = v18
            }
        case "\n\t":
            p0 -= (p0 >> ((1) & 63))
        case "a":
            var v19 bool = false
            _ = v19
        default:
            var v20 uint64 = (p0 % (((p0 << ((0) & 63))) | 1))
            _ = v20
        }
        _ = v16
    }
    switch -1 {
    case 32768:
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 uint = 26
            _ = fi21
            _ = v22
        }
    }
    _ = p0
    _ = v13
    return true
}

func main() {
    var v1 int64 = 1
    v2 := []uint32{206, 2147483648, 1}
    var v3 uint = 0
    v4 := make(map[uint32]bool)
    v5 := []uint32{2147483647, 42, 2}
    v23 := GenS2{F0: 127, F1: -1}
    var v24 int16 = 1
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 int64 = v1
        switch (("a" + "\n\t") + ("hello" + ("a" + ("foo bar" + "foo bar")))) {
        case "hello":
            fmt.Println(v24, fi25)
        case "\n\t":
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 uint = v3
                v29 := GenS1{F0: (false || (!(-1 != -840)))}
                var v30 bool = true
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        case "foo bar":
            var v31 uint = (v3 ^ v3)
            _ = v31
        default:
            var v32 int8 = 84
            _ = v32
        }
        v3 = v3
        _ = fi25
        _ = v26
    }
    var v33 int64 = 1
    switch "hello" {
    case "a":
        switch ("" + "") {
        case "\n\t":
            var v34 uint8 = 2
            _ = v34
        case "hello":
            fmt.Println(v3, v33)
        }
    }
    var v35 int16 = (v24 - v24)
    var v36 int8 = -128
    v36 = -44
    var v37 int32 = -2147483648
    v36 = v36
    if len(v2) > 0 {
        v2[0] = 1
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        v24 = (-(v35))
        fmt.Println(v33, v3, v37)
        _ = fi38
    }
    if false {
        for fi39 := 0; fi39 < 13; fi39++ {
            var v40 int16 = -552
            _ = fi39
            _ = v40
        }
        fmt.Println(v36, v33, v36)
    } else {
        v41 := GenS3{F0: 39}
        for fi42 := 0; fi42 < 19; fi42++ {
            v35 += v35
            v24 = 1000
            _ = fi42
        }
        v3 -= v3
        _ = v41
    }
    if (v3 >= v3) {
        v1 = (v33 | (-(v1)))
    } else {
        v43 := GenS1{F0: false}
        fmt.Println(v36, v36, v36)
        fmt.Println(v1)
        _ = v43
    }
    v44 := GenS4{F0: (v36 - 100)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v44
    _ = v5
}
