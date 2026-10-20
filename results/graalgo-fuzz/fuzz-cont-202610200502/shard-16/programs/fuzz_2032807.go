package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint64; F1 int64; F2 bool }

func main() {
    var v1 uint64 = 887
    var v2 float64 = 0.0
    var v3 bool = (2 != 217)
    switch -66 {
    case -2147483648:
        for fi4 := 0; fi4 < 5; fi4++ {
            v1 += (v1 >> ((2) & 63))
            _ = fi4
        }
    case 127:
        v2 -= v2
    case -283:
        var v5 float64 = 0.0
        _ = v5
    }
    fmt.Println(v2, v1, v1)
    v3 = (true && (1e10 >= 3.14159))
    v3 = true
    v6 := []int{}
    if v3 {
        v3 = false
        var v7 string = (("hello" + "\n\t") + ("\n\t" + (("foo bar" + "") + "foo bar")))
        _ = v7
    } else {
        v1 -= (^(v1))
    }
    var v8 uint64 = v1
    var v9 string = "foo bar"
    var v10 int = 2147483647
    v11 := GenS1{F0: v8, F1: v1, F2: -476}
    v8 -= v1
    v12 := GenS2{F0: v1, F1: -1, F2: (!v3)}
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        var v13 uint8 = 1
        _ = v13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 uint8 = 154
        var v16 float64 = v2
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        if v3 {
            switch (fi17 & -1) {
            case -130:
                if v3 {
                    for fi18 := 0; fi18 < 6; fi18++ {
                        var v19 int32 = -2147483648
                        switch v9 {
                        case "":
                            var v20 uint = 42
                            _ = v20
                        case "\n\t":
                            var v21 uint32 = 65535
                            _ = v21
                        }
                        var v22 int64 = 243
                        _ = fi18
                        _ = v19
                        _ = v22
                    }
                }
            case -2147483648:
                v23 := GenS1{F0: 0, F1: v8, F2: 99}
                _ = v23
            }
            v24 := GenS2{F0: (v8 / ((v8) | 1)), F1: 971, F2: v3}
            _ = v24
        } else {
            var v25 uint32 = 235
            _ = v25
        }
        var v26 int16 = -32768
        fmt.Println(v9, v10)
        _ = fi17
        _ = v26
    }
    var v27 int64 = 42
    var v28 bool = v3
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v6
    _ = v8
    _ = v9
}
