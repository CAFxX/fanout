package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int }
func genFn0() int64 {
    switch 2147483647 {
    case 0:
        var v6 int64 = 780
        _ = v6
    case 128:
        var v7 int16 = -1
        _ = v7
    }
    if ((127 <= -127) || true) {
        var v8 string = ""
        fmt.Println(v8)
        _ = v8
    }
    var v9 string = ("a" + "\n\t")
    _ = v9
    return 1
}
func genFn1(p0 uint) uint16 {
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 string = "a"
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 1
    _ = p0
    _ = v12
    return (v12 ^ v12)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 42
    v3 := []int{-2147483648, 4294967295, 127, -704}
    var v4 uint32 = 967
    var v5 uint16 = v2
    var v13 int16 = 42
    v14 := GenS1{F0: (v13 - (v13 * v13)), F1: 131}
    delete(v1, v4)
    if true {
        _ = len(v1)
    } else {
        for fi15 := 0; fi15 < 2; fi15++ {
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 uint16 = (v5 & (v2 >> ((42) & 63)))
                v2 -= v17
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        _ = len(v1)
        _ = fi18
    }
    v5 = (v2 - (v5 >> ((42) & 63)))
    var v19 float64 = 0.0
    v20 := GenS1{F0: (v13 % ((-32768) | 1)), F1: 42}
    v1[0] = false
    v2 += v2
    var v21 int16 = (v13 * v13)
    var v22 uint8 = 115
    for fi23 := 0; fi23 < 5; fi23++ {
        if (false && (0 != -9223372036854775808)) {
            if (v21 != 1000) {
                fmt.Println(v2, v5, v4)
            }
            var v24 int8 = 127
            switch (fi23 | (fi23 << ((171) & 63))) {
            case -9223372036854775808:
                v22 += v22
            case 32768:
                fmt.Println(v19, v5)
            case 127:
                var v25 uint32 = v4
                _ = v25
            default:
                v26 := GenS1{F0: v13, F1: (v22 ^ 255)}
                _ = v26
            }
            _ = v24
        } else {
            v4 += 2147483647
        }
        switch fi23 {
        case -886:
            v27 := GenS2{F0: (fi23 | fi23)}
            _ = v27
        case 42:
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                delete(v1, (v4 | (v4 >> ((42) & 63))))
            }
        }
        _ = len(v3)
        _ = fi23
    }
    v21 = (v13 + -190)
    var v28 bool = (128 > 128)
    if v28 {
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS2{F0: fi29}
            _ = fi29
            _ = v30
        }
    } else {
        switch ("" + ("foo bar" + "a")) {
        case "foo bar":
            v22 = v22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
