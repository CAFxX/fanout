package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint; F1 int8; F2 uint8 }
type GenS3 struct { F0 int32; F1 string; F2 int64 }
func genFn0(p0 uint64, p1 int16, p2 string) int {
    p1 -= (^(p1))
    p2 = p2
    p1 = p1
    p2 = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 int, p1 uint) int16 {
    p0 -= genFn0(4294967295, 32767, ((("" + "\n\t") + "a") + ""))
    fmt.Println(p1)
    var v4 int16 = -683
    _ = p0
    _ = p1
    _ = v4
    return v4
}

func main() {
    var v1 float32 = 1e10
    var v2 bool = (!true)
    var v3 string = ((("" + ("foo bar" + "")) + ("" + "foo bar")) + ("foo bar" + "a"))
    fmt.Println(v1)
    for fi5 := 0; fi5 < 8; fi5++ {
        if (v2 && (v2 || true)) {
            var v6 uint16 = 32767
            switch ((((v3 + (("" + "hello") + "foo bar")) + v3) + v3) + "foo bar") {
            case "a":
                switch genFn0(18446744073709551615, 1, v3) {
                case -245:
                    v7 := GenS1{F0: 255, F1: 182}
                    _ = v7
                }
            }
            fmt.Println(fi5, v2)
            _ = v6
        }
        _ = fi5
    }
    v8 := []int{-1, -2147483648, -1, 2147483647}
    v9 := GenS1{F0: 128, F1: 4294967295}
    v3 = "hello"
    var v10 int = 4294967295
    var v11 int = (^(v10))
    if (v2 && true) {
        var v12 int16 = 0
        v11 += (v11 + (v10 - (v11 / ((1) | 1))))
        var v13 int = genFn0(132, (v12 * 460), (((v3 + "hello") + "\n\t") + v3))
        _ = v12
        _ = v13
    } else {
        v2 = true
    }
    fmt.Println(v2, v10, v1)
    v11 = -2147483648
    v1 -= float32(1e10)
    switch (v10 * genFn0(4294967295, 0, "\n\t")) {
    case -2147483648:
        v14 := GenS1{F0: 128, F1: 0}
        _ = v14
    case 1:
        for fi15 := 0; fi15 < 11; fi15++ {
            for fi16 := 0; fi16 < 17; fi16++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                v17 := GenS2{F0: 2, F1: 100, F2: 144}
                var v18 uint = 42
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 int8 = -1
            if (v2 && (!(3.14159 <= float64((v1 * (0.0 * 0.1)))))) {
                if len(v8) > 0 {
                    v8[0] = -398
                }
            } else {
                var v20 uint16 = 65535
                fmt.Println(v19)
                v21 := GenS3{F0: -1, F1: v3, F2: 829}
                _ = v20
                _ = v21
            }
            _ = fi15
            _ = v19
        }
    case 128:
        var v22 int = v11
        _ = v22
    }
    switch v11 {
    case 0:
        if (!v2) {
            v3 = (("\n\t" + (v3 + v3)) + v3)
        }
    case 153:
        fmt.Println(v2)
    case -9223372036854775808:
        v11 -= 32767
    }
    if false {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        switch v3 {
        case "\n\t":
            v8 = append(v8, v11)
        case "foo bar":
            v10 = v11
        case "hello":
            v23 := GenS2{F0: 255, F1: -49, F2: 1}
            _ = v23
        }
        v2 = v2
    } else {
        v11 = (v10 | v11)
        switch v3 {
        case "":
            for fi24 := 0; fi24 < 9; fi24++ {
                v2 = false
                _ = fi24
            }
        case "hello":
            v2 = (2 != 632)
        }
        var v25 int16 = genFn1((v10 | v10), 0)
        _ = v25
    }
    if len(v8) > 0 {
        v8[0] = 2147483647
    }
    if (391 == 2147483648) {
        _ = len(v8)
        fmt.Println(v3, v10)
        v26 := GenS2{F0: 12, F1: -128, F2: 212}
        _ = v26
    } else {
        v10 += v11
    }
    fmt.Println(v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
