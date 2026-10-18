package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int32 }

func main() {
    var v1 uint8 = 128
    var v2 int8 = 127
    var v3 uint64 = 4294967295
    if (!((3.14159 / 1.0) >= float32((((1.5 + 0.0) * (0.0 + ((3.14159 / 1e10) * 0.0))) / 1.0)))) {
        for fi4 := 0; fi4 < 4; fi4++ {
            v5 := GenS1{F0: "foo bar"}
            _ = fi4
            _ = v5
        }
    } else {
        v6 := GenS1{F0: "\n\t"}
        v1 -= (v1 / ((v1) | 1))
        v7 := []int{int(v3), 950}
        _ = v6
        _ = v7
    }
    v8 := []int{-2147483648}
    if true {
        v3 = v3
    } else {
        fmt.Println(v2)
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS1{F0: ((("hello" + "\n\t") + ("a" + "hello")) + "foo bar")}
        _ = fi9
        _ = v10
    }
    var v11 uint32 = 65535
    for fi12 := 0; fi12 < 4; fi12++ {
        switch (fi12 | fi12) {
        case 32767:
            fmt.Println(fi12)
        case 418:
            switch "" {
            case "a":
                switch "a" {
                case "\n\t":
                    v2 -= v2
                case "":
                    v13 := GenS1{F0: ""}
                    _ = v13
                case "a":
                    v1 = (v1 + (v1 - v1))
                }
            default:
                if (!(!false)) {
                    v2 += 100
                } else {
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v1 -= 1
                        v2 = (-(v2))
                        _ = fi14
                    }
                    if len(v8) > 0 {
                        v8[0] = fi12
                    }
                    switch "\n\t" {
                    case "hello":
                        switch ((("hello" + "a") + "a") + "") {
                        case "\n\t":
                            var v15 int8 = (v2 + (v2 - v2))
                            _ = v15
                        case "hello":
                            _ = len(v8)
                        case "foo bar":
                            if false {
                                var v16 float64 = (0.1 - float64((-2.25 - (1.5 + float32(1e10)))))
                                _ = v16
                            } else {
                                fmt.Println(fi12, v11)
                            }
                        }
                    case "foo bar":
                        var v17 int16 = -43
                        _ = v17
                    case "":
                        var v18 uint8 = 161
                        _ = v18
                    }
                }
            }
        }
        _ = fi12
    }
    var v19 string = ("foo bar" + "\n\t")
    var v20 int16 = -32768
    v3 = 2
    var v21 int64 = -9223372036854775808
    var v22 float64 = 0.1
    for fi23 := 0; fi23 < 10; fi23++ {
        v21 += v21
        _ = len(v8)
        _ = fi23
    }
    if (1.5 == 0.0) {
        var v24 int8 = (v2 ^ (v2 % (((v2 >> ((1) & 63))) | 1)))
        v25 := GenS2{F0: 0, F1: 2147483647}
        _ = v24
        _ = v25
    } else {
        v1 = v1
    }
    v20 -= v20
    switch "a" {
    case "foo bar":
        if ((v2 << ((0) & 63)) <= v2) {
            if len(v8) > 0 {
                v8[0] = 4294967295
            }
        } else {
            v3 += v3
            v26 := GenS1{F0: "hello"}
            _ = v26
        }
    case "hello":
        if false {
            if false {
                for fi27 := 0; fi27 < 11; fi27++ {
                    v21 -= v21
                    _ = fi27
                }
                var v28 bool = (!(2147483647 >= 9223372036854775807))
                if ((v1 >= v1) && v28) {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v29 uint32 = v11
                    _ = v29
                }
                _ = v28
            }
            if (1 != v20) {
                if len(v8) > 0 {
                    v8[0] = -2147483648
                }
            } else {
                var v30 uint8 = (v1 - v1)
                var v31 int8 = v2
                if len(v8) > 0 {
                    v8[0] = 128
                }
                _ = v30
                _ = v31
            }
        }
    case "a":
        fmt.Println(v11, v19, v1)
    default:
        v22 = v22
    }
    v19 = (v19 + ((v19 + v19) + v19))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v8
}
