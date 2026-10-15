package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int8 = 127
    var v2 int8 = (v1 * (v1 | (v1 % ((v1) | 1))))
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    var v5 int16 = 0
    var v6 uint = 118
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{}
        if false {
            var v9 int16 = 32767
            var v10 uint64 = 135
            _ = v10
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    if ((false && true) || true) {
        v11 := GenS1{F0: -2147483648}
        fmt.Println(v1)
        _ = v11
    } else {
        v12 := []int{1, 974}
        var v13 uint16 = 0
        _ = v12
        _ = v13
    }
    switch "" {
    case "":
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := GenS1{F0: (fi14 | (fi14 & fi14))}
            switch (fi14 * fi14) {
            case 0:
                v2 += v2
            }
            _ = fi14
            _ = v15
        }
    case "\n\t":
        v2 = -63
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := GenS1{F0: fi16}
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v19 := GenS1{F0: -1}
    fmt.Println(v5)
    v2 -= (v1 & (v1 << ((255) & 63)))
    var v20 int8 = v2
    if (2 == 257) {
        fmt.Println(v20)
        for fi21 := 0; fi21 < 16; fi21++ {
            fmt.Println(v6, v1)
            var v22 string = ""
            _ = fi21
            _ = v22
        }
    }
    if (((true || (42 < 9223372036854775807)) && false) && ((255 < 208) && (65535 > 2))) {
        var v23 int8 = 127
        var v24 float32 = (0.0 - 1e10)
        _ = v23
        _ = v24
    }
    fmt.Println(v20, v20)
    v5 = (v5 * v5)
    v1 += v20
    v25 := []int{}
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        v5 += (-(v5))
    case "":
        for fi26 := 0; fi26 < 7; fi26++ {
            fmt.Println(v1)
            v5 -= 42
            v1 = 1
            _ = fi26
        }
    default:
        _ = len(v3)
    }
    var v27 int = 0
    for fi28 := 0; fi28 < 14; fi28++ {
        for fi29 := 0; fi29 < 10; fi29++ {
            if (3.14159 <= 3.14159) {
                v6 = v6
                v25 = append(v25, (fi29 * -100))
                v27 = (fi28 * (v27 - fi29))
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                var v30 int16 = 793
                _ = v30
            }
            v31 := GenS1{F0: (v27 % ((int(v2)) | 1))}
            _ = fi29
            _ = v31
        }
        fmt.Println(v20, fi28)
        _ = fi28
    }
    var v32 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
