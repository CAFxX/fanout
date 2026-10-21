package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 string; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 uint16; F2 uint }

func main() {
    var v1 float64 = float64(1e10)
    v2 := make(map[int]bool)
    v3 := []int{-35}
    v1 -= (0.0 + 0.0)
    v1 = v1
    v2[640] = (-768 != 0)
    v1 -= v1
    if (32767 <= -32768) {
        switch ("a" + "a") {
        case "a":
            if false {
                switch ("\n\t" + ("foo bar" + ("a" + ("" + "hello")))) {
                case "foo bar":
                    switch ((((("foo bar" + ("a" + "foo bar")) + ("a" + ((("\n\t" + ("\n\t" + "foo bar")) + ("hello" + ("\n\t" + ("a" + "a")))) + "\n\t"))) + ("\n\t" + "hello")) + "\n\t") + "\n\t") {
                    case "foo bar":
                        v1 = v1
                    case "hello":
                        if (1 <= 120) {
                            for fi4 := 0; fi4 < 19; fi4++ {
                                if (!true) {
                                    v5 := GenS1{F0: -47}
                                    _ = v5
                                }
                                var v6 uint32 = 114
                                _ = fi4
                                _ = v6
                            }
                            var v7 float32 = ((float32((v1 - ((-2.25 - 0.0) + ((v1 * v1) / v1)))) + 1.5) + 1.5)
                            _ = v7
                        } else {
                            var v8 float64 = (v1 + 0.0)
                            _ = v8
                        }
                    case "\n\t":
                        v9 := GenS2{F0: false, F1: "hello", F2: 128}
                        _ = v9
                    }
                case "":
                    switch -1 {
                    case 761:
                        var v10 uint16 = 374
                        _ = v10
                    }
                case "a":
                    v1 -= float64(0.0)
                default:
                    fmt.Println(v1)
                }
            }
        default:
            v11 := GenS1{F0: -2147483648}
            _ = v11
        }
    } else {
        v1 = 3.14159
        v1 -= float64((float32((v1 * v1)) / 1.5))
        var v12 int = 0
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v1 -= float64(3.14159)
        _ = fi13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 int32 = -1
        if len(v3) > 0 {
            v3[0] = 9223372036854775807
        }
        _ = fi14
        _ = v15
    }
    var v16 uint32 = 4294967295
    var v17 int32 = 560
    switch ((("hello" + ("" + "hello")) + "\n\t") + ("foo bar" + "")) {
    case "":
        var v18 float32 = 0.0
        _ = v18
    case "a":
        switch "hello" {
        case "\n\t":
            v19 := GenS3{F0: 42}
            _ = v19
        case "hello":
            fmt.Println(v1, v16)
        case "a":
            fmt.Println(v1, v16, v1)
        }
    case "foo bar":
        if true {
            v1 += (v1 + (float64(float32((v1 + 1.5))) * (float64(float32((-2.25 * -2.25))) - v1)))
            var v20 uint = 42
            _ = v20
        }
    default:
        var v21 int = 1
        _ = v21
    }
    v3 = append(v3, 128)
    var v22 int32 = -933
    if (true && (32767 > -32768)) {
        var v23 int32 = 1
        v22 += v17
        v1 -= float64((3.14159 / 3.14159))
        _ = v23
    } else {
        v1 = v1
    }
    fmt.Println(v22, v1)
    switch "hello" {
    case "foo bar":
        v16 -= v16
    case "":
        var v24 uint64 = 389
        _ = v24
    case "hello":
        var v25 int16 = 32767
        _ = v25
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        if len(v3) > 0 {
            v3[0] = fi26
        }
        v27 := GenS4{F0: "foo bar", F1: 42, F2: 0}
        for fi28 := 0; fi28 < 13; fi28++ {
            var v29 string = "foo bar"
            var v30 int16 = -1
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
}
