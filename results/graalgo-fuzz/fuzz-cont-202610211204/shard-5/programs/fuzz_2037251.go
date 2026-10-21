package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int; F2 int32 }
type GenS4 struct { F0 int64 }
func genFn0(p0 int8) int8 {
    v5 := []int{}
    var v6 uint = 65535
    _ = p0
    _ = v5
    _ = v6
    return (p0 - 42)
}
func genFn1(p0 uint64, p1 uint8, p2 uint16) int {
    v7 := []int{42, -240, 1, 0}
    if ("\n\t" > "a") {
        v8 := GenS1{F0: 0}
        _ = v8
    }
    v7 = append(v7, 127)
    switch 32768 {
    case 32768:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    default:
        if len(v7) > 0 {
            v7[0] = 2147483647
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 127
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 1
    var v3 uint32 = 4294967295
    v4 := make(map[string]uint32)
    for fi9 := 0; fi9 < 5; fi9++ {
        fmt.Println(v2)
        _ = fi9
    }
    var v10 uint16 = v2
    v2 += (^(v10))
    v3 = (^(v3))
    if false {
        fmt.Println(v10)
        var v11 string = "a"
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v3 += v3
        switch (fi12 - (fi12 | fi12)) {
        case 42:
            for fi13 := 0; fi13 < 2; fi13++ {
                v14 := GenS1{F0: -239}
                switch fi12 {
                case 724:
                    if false {
                        switch ("foo bar" + "foo bar") {
                        case "\n\t":
                            var v15 uint64 = 42
                            _ = v15
                        }
                    }
                }
                _ = fi13
                _ = v14
            }
        }
        v2 += (v10 >> ((776) & 63))
        _ = fi12
    }
    if ((v2 ^ v2) == 0) {
        var v16 uint16 = v10
        v10 = 414
        _ = v16
    } else {
        if (-1 <= 42) {
            fmt.Println(v2)
        }
    }
    v2 = 1
    switch (("a" + "") + "") {
    case "a":
        var v17 uint16 = v10
        _ = v17
    case "hello":
        v3 -= 2147483648
    }
    if (0 > 127) {
        v2 = 237
        fmt.Println(v2, v3, v2)
        v18 := []int{127, -843, 2147483647}
        _ = v18
    } else {
        switch -314 {
        case 32767:
            v19 := GenS2{F0: "hello"}
            _ = v19
        case 794:
            if (!false) {
                v20 := GenS3{F0: ("foo bar" + "foo bar"), F1: 1, F2: -2147483648}
                _ = v20
            } else {
                var v21 int32 = -2147483648
                _ = v21
            }
        case -480:
            var v22 string = "hello"
            _ = v22
        }
        switch -2147483648 {
        case 115:
            fmt.Println(v3)
        }
        for fi23 := 0; fi23 < 6; fi23++ {
            v24 := GenS4{F0: 883}
            if (435 < -980) {
                if (v3 != (^(v3))) {
                    switch genFn1(4294967295, 1, 2) {
                    case -9223372036854775808:
                        fmt.Println(fi23, v3)
                    case -1:
                        switch -624 {
                        case 128:
                            v25 := []int{(fi23 << ((42) & 63)), fi23, (^(fi23)), 0}
                            _ = v25
                        case 42:
                            if (false || (1 >= 64)) {
                                v10 += uint16(v3)
                            }
                        }
                    }
                    v2 += (v10 % (((v2 >> ((255) & 63))) | 1))
                } else {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        fmt.Println(v3)
                        _ = fi26
                    }
                    delete(v4, "hello")
                    var v27 uint64 = 0
                    _ = v27
                }
                v2 = (v2 % ((v10) | 1))
                fmt.Println(v2, v10)
            }
            _ = fi23
            _ = v24
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
