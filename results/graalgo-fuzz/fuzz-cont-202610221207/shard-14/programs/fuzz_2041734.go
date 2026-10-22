package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint8, p1 int64) int {
    var v5 uint = 42
    fmt.Println(v5, v5, p0)
    v6 := GenS1{F0: true, F1: false}
    var v7 float32 = (1.5 * (3.14159 + (1e10 + (0.1 * (0.1 * ((-2.25 / (1.5 * 3.14159)) - (3.14159 / 1.0)))))))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1() int8 {
    if (!false) {
        var v8 uint64 = 0
        _ = v8
    } else {
        v9 := GenS1{F0: (0 <= 81), F1: (65535 > 2)}
        if (-820 > -783) {
            var v10 bool = (!true)
            v10 = v10
            _ = v10
        }
        _ = v9
    }
    switch "a" {
    case "hello":
        v11 := []int{32768, -2147483648, 1, 0}
        _ = v11
    case "\n\t":
        switch "\n\t" {
        case "a":
            var v12 string = "hello"
            _ = v12
        case "\n\t":
            for fi13 := 0; fi13 < 13; fi13++ {
                for fi14 := 0; fi14 < 10; fi14++ {
                    v15 := []int{}
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
        case "hello":
            var v16 int64 = 0
            _ = v16
        }
    case "a":
        if (true && (!(852 == 1))) {
            var v17 int16 = 0
            _ = v17
        } else {
            v18 := []int{128}
            var v19 int32 = -2147483648
            _ = v18
            _ = v19
        }
    }
    v20 := GenS1{F0: false, F1: (!false)}
    _ = v20
    return 100
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = ("" + "\n\t")
    var v3 uint32 = 2147483648
    var v4 bool = (false && false)
    var v21 string = (v2 + (v2 + v2))
    var v22 uint16 = 255
    var v23 int8 = 42
    v24 := GenS2{F0: 6}
    if false {
        if v4 {
            switch -2147483648 {
            case 128:
                if true {
                    v25 := GenS3{F0: 127, F1: (v21 + ("hello" + ("hello" + (("foo bar" + (v2 + v21)) + v21))))}
                    _ = v25
                }
            case 14:
                var v26 int64 = 42
                _ = v26
            case -9223372036854775808:
                for fi27 := 0; fi27 < 2; fi27++ {
                    for fi28 := 0; fi28 < 2; fi28++ {
                        fmt.Println(v21, v2, v22)
                        v22 += uint16(v3)
                        var v29 uint = 0
                        _ = fi28
                        _ = v29
                    }
                    v23 -= (v23 ^ (^(v23)))
                    v2 = ((v21 + (v21 + "foo bar")) + "")
                    _ = fi27
                }
            }
            v4 = v4
        } else {
            for fi30 := 0; fi30 < 1; fi30++ {
                fmt.Println(v21)
                _ = fi30
            }
            fmt.Println(v21)
            v31 := GenS1{F0: v4, F1: v4}
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        for fi33 := 0; fi33 < 12; fi33++ {
            fmt.Println(v23)
            _ = fi33
        }
        fmt.Println(fi32, fi32)
        fmt.Println(fi32, v21, v23)
        _ = fi32
    }
    v34 := []int{}
    switch (v21 + "a") {
    case "":
        for fi35 := 0; fi35 < 9; fi35++ {
            var v36 uint32 = v3
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
            _ = fi35
            _ = v36
        }
    }
    var v37 int8 = v23
    var v38 uint16 = v22
    v1["\n\t"] = v3
    v34 = append(v34, int(v3))
    var v39 int32 = 42
    if v4 {
        fmt.Println(v37)
    }
    for fi40 := 0; fi40 < 6; fi40++ {
        v41 := GenS4{F0: 610, F1: v39}
        for fi42 := 0; fi42 < 5; fi42++ {
            v43 := GenS5{F0: 0, F1: 535}
            v34 = append(v34, -573)
            v44 := GenS6{F0: -32768}
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi40
        _ = v41
    }
    fmt.Println(v3, v37, v21)
    if (v4 && (9223372036854775807 > 486)) {
        fmt.Println(v2)
        var v45 uint32 = (v3 % ((v3) | 1))
        var v46 int16 = 32767
        _ = v45
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
