package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint8, p1 uint64, p2 string) bool {
    var v4 float64 = -2.25
    v5 := GenS1{F0: 838, F1: (-1 < -173), F2: -1}
    v4 += (v4 + 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return true
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"\n\t", "hello", ("" + "foo bar"), ""}
    v3 := []string{}
    var v6 int32 = 1
    if false {
        v6 -= (v6 + (v6 / ((-2147483648) | 1)))
        if true {
            var v7 uint32 = 550
            switch ("a" + ("" + ("hello" + "\n\t"))) {
            case "":
                v6 -= v6
            }
            _ = v7
        } else {
            if true {
                fmt.Println(v6)
                if (true || (v6 >= (v6 >> ((42) & 63)))) {
                    v8 := GenS2{F0: 645}
                    _ = v8
                }
                v6 -= -370
            }
        }
    }
    switch -150 {
    case 32768:
        v9 := GenS3{F0: 958, F1: 561, F2: 18446744073709551615}
        _ = v9
    case 42:
        v6 = (v6 - 1)
    case 32767:
        if (!(-748 > 9223372036854775807)) {
            v2 = append(v2, "")
        } else {
            fmt.Println(v6)
        }
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v6 += v6
        _ = fi10
    }
    fmt.Println(v6)
    var v11 int16 = 1000
    if false {
        v11 = v11
        v6 -= 42
        v11 = -600
    } else {
        switch (("hello" + "foo bar") + "") {
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "a":
            v12 := GenS4{F0: 32767}
            _ = v12
        }
        v13 := GenS5{F0: "", F1: 960}
        _ = v13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        if ("hello" < (("hello" + "\n\t") + "\n\t")) {
            v11 = v11
            var v15 uint = 732
            v16 := GenS3{F0: 42, F1: 255, F2: (^(v15))}
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 uint16 = 255
            _ = fi17
            _ = v18
        }
        if (false && (!true)) {
            var v19 int64 = 9223372036854775807
            _ = len(v3)
            _ = v19
        } else {
            if false {
                v20 := GenS6{F0: 42}
                v6 -= 2147483647
                _ = v20
            }
            v6 = (v6 / ((0) | 1))
        }
        _ = fi14
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v6)
    v6 -= (^(v6))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v6
}
