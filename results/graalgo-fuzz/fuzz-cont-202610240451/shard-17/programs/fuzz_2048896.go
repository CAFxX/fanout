package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint }
type GenS4 struct { F0 uint32; F1 int32 }

func main() {
    var v1 string = ((("" + "a") + "foo bar") + ("a" + (("\n\t" + "") + (((("\n\t" + ("foo bar" + ("a" + ""))) + "foo bar") + "a") + ("\n\t" + "a")))))
    var v2 uint8 = 2
    v3 := []int{}
    var v4 int8 = -1
    var v5 int32 = 42
    var v6 uint = 42
    var v7 int16 = -32768
    v8 := GenS1{F0: 4294967295}
    var v9 int64 = -213
    for fi10 := 0; fi10 < 16; fi10++ {
        if false {
            var v11 uint = 4294967295
            _ = v11
        } else {
            v12 := GenS1{F0: 1}
            for fi13 := 0; fi13 < 11; fi13++ {
                v6 += (^(v6))
                switch (fi13 | fi10) {
                case 0:
                    var v14 int32 = (v5 + (^(v5)))
                    v3 = append(v3, fi13)
                    if true {
                        fmt.Println(v5, v6)
                    }
                    _ = v14
                }
                _ = fi13
            }
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 uint64 = 0
                var v17 int = (-(fi15))
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v12
        }
        for fi18 := 0; fi18 < 6; fi18++ {
            if true {
                v3 = append(v3, fi10)
            } else {
                v19 := GenS2{F0: -2147483648, F1: 18446744073709551615}
                var v20 uint8 = (v2 ^ v2)
                if false {
                    var v21 uint = (^(v6))
                    fmt.Println(fi18, v6, fi10)
                    _ = len(v3)
                    if len(v3) > 0 {
                        v3[0] = (fi10 + (fi10 / ((9223372036854775807) | 1)))
                    }
                    _ = v21
                } else {
                    v22 := GenS3{F0: 2147483647, F1: 1, F2: (v6 & v6)}
                    var v23 float64 = -2.25
                    _ = v22
                    _ = v23
                }
                if len(v3) > 0 {
                    v3[0] = -800
                }
                v3 = append(v3, -9223372036854775808)
                v24 := GenS3{F0: 42, F1: (v9 % ((v9) | 1)), F2: 1}
                var v25 int = fi10
                _ = v19
                _ = v20
                _ = v24
                _ = v25
            }
            switch v1 {
            case "foo bar":
                var v26 int16 = v7
                _ = v26
            case "a":
                _ = len(v3)
                var v27 int64 = (v9 * v9)
                var v28 bool = (100 <= 1)
                v3 = append(v3, fi18)
                fmt.Println(fi10, fi18, v7)
                _ = v27
                _ = v28
            }
            _ = fi18
        }
        _ = fi10
    }
    if true {
        fmt.Println(v9)
        v29 := GenS2{F0: 4294967295, F1: 697}
        v30 := GenS4{F0: 4294967295, F1: v5}
        _ = v29
        _ = v30
    }
    v31 := GenS2{F0: -2147483648, F1: 4294967295}
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v32 int64 = -783
    if ((4294967295 > 984) || (535 != 600)) {
        switch 977 {
        case -493:
            fmt.Println(v5, v1, v7)
        }
    }
    var v33 int32 = v5
    v5 = v33
    var v34 uint8 = v2
    if ((-(v4)) <= v4) {
        if (1.5 <= (1e10 - (0.0 / 1.0))) {
            var v35 uint16 = 783
            _ = v35
        }
        fmt.Println(v7, v32, v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
