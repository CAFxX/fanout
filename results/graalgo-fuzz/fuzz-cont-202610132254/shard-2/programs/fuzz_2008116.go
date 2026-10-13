package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int16 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint16 = 32767
    var v2 int8 = 0
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    v5 := make(map[int]bool)
    if (2147483647 > 4294967295) {
        fmt.Println(v1)
        v1 = (v1 ^ (^(v1)))
    }
    var v6 bool = false
    delete(v5, 403)
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(v2, v6, v2)
        _ = fi7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 string = "\n\t"
            _ = fi9
            _ = v10
        }
        var v11 int8 = v2
        _ = fi8
        _ = v11
    }
    var v12 float64 = float64(1.5)
    if (0 <= 42) {
        var v13 int = 42
        var v14 float32 = ((float32(0.1) / ((0.0 - 1.5) - 1.5)) - 0.1)
        _ = v13
        _ = v14
    } else {
        fmt.Println(v6, v1)
        fmt.Println(v2, v2, v1)
    }
    v15 := GenS1{F0: 726, F1: -1, F2: -333}
    v16 := GenS1{F0: 2147483647, F1: 2147483647, F2: -329}
    _ = len(v5)
    v17 := []int{-9223372036854775808, 4294967295, 297}
    if false {
        if (!(v6 && (v6 || false))) {
            v12 -= float64(1.5)
            v18 := GenS1{F0: 42, F1: 0, F2: 466}
            _ = len(v4)
            _ = v18
        } else {
            fmt.Println(v1, v12, v1)
            var v19 bool = v6
            v1 = v1
            _ = v19
        }
        var v20 uint32 = 2147483648
        _ = v20
    }
    var v21 int16 = 32767
    if true {
        var v22 float32 = -2.25
        var v23 uint32 = 0
        _ = v22
        _ = v23
    } else {
        var v24 uint = 42
        v1 -= v1
        v24 = 255
        _ = v24
    }
    v25 := GenS2{F0: 359}
    for fi26 := 0; fi26 < 6; fi26++ {
        v1 = 526
        for fi27 := 0; fi27 < 16; fi27++ {
            if (fi27 >= fi27) {
                switch fi26 {
                case 42:
                    fmt.Println(fi27, v1, v1)
                case -256:
                    var v28 int64 = -9223372036854775808
                    _ = v28
                }
                v29 := GenS2{F0: fi26}
                _ = v29
            } else {
                if v6 {
                    fmt.Println(fi27, v21, v2)
                } else {
                    delete(v4, ("a" + ("hello" + "\n\t")))
                }
                if len(v17) > 0 {
                    v17[0] = (fi27 | (fi26 ^ (fi26 + (fi26 / (((fi26 % ((fi26) | 1))) | 1)))))
                }
            }
            fmt.Println(fi26)
            switch (("" + (((("foo bar" + "a") + "a") + ("a" + "")) + "a")) + "foo bar") {
            case "hello":
                var v30 string = "a"
                _ = v30
            }
            _ = fi27
        }
        var v31 uint64 = 0
        _ = fi26
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
