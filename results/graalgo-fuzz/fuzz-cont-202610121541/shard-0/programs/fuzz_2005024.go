package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 130
    var v3 int = -248
    var v4 bool = true
    v5 := GenS1{F0: 0, F1: 42}
    var v6 int8 = 127
    if v4 {
        for fi7 := 0; fi7 < 2; fi7++ {
            v4 = false
            var v8 int64 = 9223372036854775807
            _ = fi7
            _ = v8
        }
        v6 = (v6 & 100)
        var v9 uint32 = 2147483647
        _ = v9
    } else {
        v10 := GenS1{F0: 218, F1: -1}
        _ = v10
    }
    var v11 float64 = ((-2.25 - 1.5) + 0.0)
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 float64 = v11
            switch fi13 {
            case 4294967295:
                v15 := GenS1{F0: 127, F1: 874}
                v3 -= (^(fi13))
                _ = v15
            case 549:
                _ = len(v1)
            default:
                if (v4 && v4) {
                    var v16 int8 = v6
                    v17 := GenS1{F0: 128, F1: -736}
                    v18 := v1[fi12]
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
            v11 += (v14 / v11)
            _ = fi13
            _ = v14
        }
        var v19 int64 = -9223372036854775808
        _ = fi12
        _ = v19
    }
    delete(v1, 4294967295)
    if (1 < 2) {
        v1[v3] = v3
        v1[v3] = 9223372036854775807
        fmt.Println(v4)
    } else {
        var v20 uint16 = 2
        v21 := GenS1{F0: 99, F1: -32768}
        v22 := []int{(v3 + int(v20)), (v3 >> ((1) & 63)), (v3 % (((v3 - (v3 & (v3 % ((v3) | 1))))) | 1))}
        _ = v20
        _ = v21
        _ = v22
    }
    switch (("\n\t" + "") + ("a" + "\n\t")) {
    case "hello":
        fmt.Println(v3)
    default:
        if (!(0 == 407)) {
            var v23 string = "hello"
            var v24 uint16 = 65535
            _ = v23
            _ = v24
        } else {
            var v25 int = 592
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        switch "" {
        case "":
            for fi27 := 0; fi27 < 15; fi27++ {
                if (!(42 <= 1)) {
                    v6 = 42
                } else {
                    v28 := v1[1]
                    var v29 int16 = -920
                    var v30 int64 = 433
                    var v31 int8 = 0
                    _ = v28
                    _ = v29
                    _ = v30
                    _ = v31
                }
                _ = fi27
            }
        case "\n\t":
            switch fi26 {
            case -500:
                var v32 uint8 = 2
                _ = v32
            case 32768:
                for fi33 := 0; fi33 < 20; fi33++ {
                    v34 := []int{(fi33 >> ((382) & 63))}
                    v4 = v4
                    _ = fi33
                    _ = v34
                }
            default:
                v35 := v1[128]
                _ = v35
            }
        case "a":
            v36 := v1[v3]
            _ = v36
        default:
            v4 = (18446744073709551615 < 18446744073709551615)
        }
        if (0 < 2) {
            v4 = v4
            if (0 <= 42) {
                v6 = v6
            }
            var v37 float32 = 1e10
            _ = v37
        }
        switch (-(v3)) {
        case 127:
            v4 = ((v2 % (((v2 / ((65535) | 1))) | 1)) <= v2)
        case 32767:
            v38 := []int{v3, (v3 ^ (v3 | (v3 + fi26))), v3}
            _ = v38
        case -1:
            fmt.Println(v3, v4)
        }
        _ = fi26
    }
    v3 = (-(v3))
    var v39 uint16 = 2
    v40 := []int{(v3 | v3), 4294967295}
    v1[(v3 & (v3 >> ((317) & 63)))] = 2147483647
    var v41 float64 = v11
    switch (v3 << ((172) & 63)) {
    case 0:
        fmt.Println(v41, v2, v6)
    case 1:
        v40 = append(v40, (v3 * 9223372036854775807))
    }
    _ = len(v40)
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
