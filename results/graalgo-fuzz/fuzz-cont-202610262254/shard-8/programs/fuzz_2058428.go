package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]bool)
    var v3 bool = true
    var v4 uint64 = 1
    fmt.Println(v4, v1)
    var v5 uint64 = v4
    v6 := []int{144, 42}
    v7 := GenS1{F0: 2}
    var v8 uint8 = v1
    switch 156 {
    case 348:
        if true {
            v9 := GenS1{F0: 18446744073709551615}
            var v10 uint8 = (^(v1))
            if (255 > 869) {
                var v11 uint32 = 42
                var v12 int8 = 127
                for fi13 := 0; fi13 < 1; fi13++ {
                    v3 = v3
                    fmt.Println(v1)
                    _ = fi13
                }
                _ = v11
                _ = v12
            }
            _ = v10
            _ = v9
        } else {
            v14 := GenS1{F0: 18446744073709551615}
            _ = v14
        }
    }
    v15 := GenS1{F0: 0}
    for fi16 := 0; fi16 < 15; fi16++ {
        switch ("\n\t" + "foo bar") {
        case "hello":
            if v3 {
                v1 -= (v8 & v1)
                var v17 int = (fi16 % ((fi16) | 1))
                v5 -= (v4 / ((v4) | 1))
                _ = v17
            } else {
                var v18 uint8 = (^(v8))
                _ = v18
            }
        }
        if (387 > 732) {
            fmt.Println(fi16)
            var v19 int16 = -455
            _ = v19
        } else {
            var v20 uint8 = v1
            var v21 float64 = 0.0
            _ = v20
            _ = v21
        }
        _ = fi16
    }
    v22 := GenS1{F0: 4294967295}
    var v23 uint64 = 1
    _ = len(v6)
    switch int(v23) {
    case -9223372036854775808:
        var v24 int32 = 1
        _ = v24
    default:
        v5 = v4
    }
    v25 := GenS1{F0: 345}
    for fi26 := 0; fi26 < 13; fi26++ {
        fmt.Println(v5)
        _ = fi26
    }
    if v3 {
        var v27 uint8 = 2
        if (!(100 >= 0)) {
            v28 := GenS1{F0: 42}
            _ = v28
        }
        _ = v27
    } else {
        if (!(v3 && (((1.5 + float32(0.0)) - -2.25) >= 0.0))) {
            v4 -= (v5 & (v23 % ((118) | 1)))
        }
        _ = len(v6)
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 string = "hello"
        _ = fi29
        _ = v30
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
