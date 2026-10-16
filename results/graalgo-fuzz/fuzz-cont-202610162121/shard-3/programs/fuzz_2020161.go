package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS2 struct { F0 string }
func genFn0() int32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(fi7)
        var v8 string = ("a" + "")
        _ = fi7
        _ = v8
    }
    var v9 string = "hello"
    if true {
        v10 := []int{42}
        if ((-1 == 0) || (!(0 <= 0))) {
            v11 := GenS1{F0: 467, F1: 1, F2: 32767}
            _ = v11
        }
        _ = v10
    }
    v12 := []int{2147483647, -983, 32767, 32767}
    var v13 string = "\n\t"
    _ = v12
    _ = v13
    _ = v9
    return 1
}

func main() {
    v1 := []int{}
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]bool)
    var v6 string = ""
    for fi14 := 0; fi14 < 6; fi14++ {
        switch (fi14 & -9223372036854775808) {
        case 9223372036854775807:
            v6 = "\n\t"
        case -82:
            v6 = ("" + v6)
        default:
            var v15 uint64 = 18446744073709551615
            _ = v15
        }
        _ = fi14
    }
    if ((2 > 2147483648) && (!true)) {
        _ = len(v1)
    } else {
        if (true || true) {
            if true {
                switch v6 {
                case "hello":
                    v6 = v6
                }
                if (!(!(false || true))) {
                    fmt.Println(v6)
                }
            }
            var v16 bool = (1e10 != 1e10)
            v17 := GenS2{F0: v6}
            _ = v16
            _ = v17
        } else {
            v2 = -9223372036854775808
            var v18 uint = 124
            var v19 string = (((v6 + "foo bar") + v6) + "foo bar")
            _ = v18
            _ = v19
        }
        if ((v2 % (((^(v2))) | 1)) == v2) {
            switch 2147483647 {
            case 2147483647:
                var v20 int64 = -9223372036854775808
                _ = v20
            }
            v21 := v5[2147483648]
            _ = v21
        }
    }
    var v22 float64 = 3.14159
    var v23 float32 = 3.14159
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 float64 = 3.14159
        _ = fi24
        _ = v25
    }
    var v26 uint32 = 1
    var v27 int8 = -128
    var v28 int16 = 32767
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
