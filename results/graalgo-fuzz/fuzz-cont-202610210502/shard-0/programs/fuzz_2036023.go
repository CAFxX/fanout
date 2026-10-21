package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 uint; F2 int }
func genFn0(p0 int, p1 bool, p2 uint8) int {
    var v6 uint = 397
    v7 := GenS1{F0: 398, F1: (^(p0))}
    var v8 uint8 = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (^(p0))
}
func genFn1() uint8 {
    var v9 uint32 = 1
    var v10 int8 = -128
    v11 := GenS2{F0: 2}
    _ = v10
    _ = v11
    _ = v9
    return 87
}

func main() {
    var v1 bool = (!true)
    var v2 uint32 = 2
    var v3 string = ("" + "\n\t")
    var v4 int8 = 15
    v5 := make(map[string]int)
    fmt.Println(v4)
    for fi12 := 0; fi12 < 3; fi12++ {
        v1 = false
        var v13 int32 = 42
        var v14 int16 = -32768
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := []int{127}
    fmt.Println(v4)
    switch "\n\t" {
    case "":
        if v1 {
            _ = len(v5)
            var v16 int8 = v4
            var v17 uint8 = 2
            _ = v16
            _ = v17
        } else {
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 bool = (("\n\t" + v3) <= v3)
                _ = fi18
                _ = v19
            }
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            v4 += (v4 + v4)
        }
    case "a":
        v20 := GenS2{F0: 436}
        _ = v20
    case "foo bar":
        if (!(!(631 >= 265))) {
            for fi21 := 0; fi21 < 15; fi21++ {
                v1 = ((0.0 + 1e10) <= 0.0)
                v22 := GenS2{F0: 329}
                _ = fi21
                _ = v22
            }
            v15 = append(v15, 1)
        } else {
            var v23 int = genFn0(9223372036854775807, true, 234)
            _ = v23
        }
    }
    if true {
        v5["a"] = 32767
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 int = 2147483647
            for fi26 := 0; fi26 < 18; fi26++ {
                v4 += -61
                _ = fi26
            }
            _ = fi24
            _ = v25
        }
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        _ = len(v15)
        if (0 < -9223372036854775808) {
            var v28 uint16 = 2
            _ = v28
        } else {
            v1 = false
            v5[v3] = (fi27 / ((fi27) | 1))
            v2 = v2
        }
        _ = fi27
    }
    switch 128 {
    case 2147483647:
        v2 = v2
    }
    var v29 float32 = (3.14159 * 0.0)
    var v30 int8 = (v4 * 127)
    if true {
        v31 := GenS2{F0: 436}
        for fi32 := 0; fi32 < 4; fi32++ {
            v5["foo bar"] = -1
            v33 := GenS3{F0: (v4 * 116), F1: v3}
            _ = fi32
            _ = v33
        }
        fmt.Println(v4)
        _ = v31
    }
    if v1 {
        v30 += v4
        for fi34 := 0; fi34 < 17; fi34++ {
            v29 -= v29
            _ = fi34
        }
    }
    switch v3 {
    case "hello":
        v35 := GenS4{F0: v2}
        _ = v35
    case "foo bar":
        if (v1 && v1) {
            var v36 uint = 65535
            v4 = -128
            _ = v36
        } else {
            switch "hello" {
            case "hello":
                var v37 int8 = 58
                _ = v37
            }
        }
    }
    var v38 string = (((v3 + (((v3 + v3) + "\n\t") + (v3 + v3))) + v3) + v3)
    v2 = v2
    v39 := GenS5{F0: 756, F1: 4294967295, F2: -639}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
