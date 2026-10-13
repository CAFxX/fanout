package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 979
    var v3 uint = 131
    var v4 int16 = -1
    v5 := make(map[int]int)
    v6 := GenS1{F0: 224, F1: 2147483647}
    if (0 != 53) {
        for fi7 := 0; fi7 < 13; fi7++ {
            if (9223372036854775807 > 42) {
                switch "a" {
                case "":
                    var v8 int = fi7
                    _ = v8
                case "foo bar":
                    v1 = 203
                default:
                    var v9 int32 = int32(fi7)
                    _ = v9
                }
                for fi10 := 0; fi10 < 6; fi10++ {
                    delete(v5, (fi7 + (^(fi7))))
                    if ("\n\t" <= "foo bar") {
                        v1 -= (^(v1))
                    } else {
                        v1 -= 42
                        v11 := GenS1{F0: 211, F1: 2147483647}
                        _ = v11
                    }
                    switch (fi7 << (((v3 << (((v3 << (((v3 & v3)) & 63))) & 63))) & 63)) {
                    case 9223372036854775807:
                        fmt.Println(v2)
                    default:
                        var v12 uint = v3
                        _ = v12
                    }
                    _ = fi10
                }
            }
            v13 := GenS2{F0: -54, F1: v4}
            delete(v5, (fi7 % ((42) | 1)))
            _ = fi7
            _ = v13
        }
        if true {
            v4 = v4
        }
        v2 += 2
    } else {
        if false {
            if true {
                switch 0 {
                case 32767:
                    switch (("foo bar" + ((("" + ("a" + ("hello" + "foo bar"))) + "hello") + (("foo bar" + ("foo bar" + ("a" + "\n\t"))) + "a"))) + "") {
                    case "\n\t":
                        var v14 uint8 = 224
                        _ = v14
                    case "a":
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v1, v2)
                        case "hello":
                            var v15 string = "a"
                            _ = v15
                        default:
                            var v16 float64 = 1e10
                            _ = v16
                        }
                    }
                }
                var v17 uint32 = v2
                var v18 float64 = (3.14159 / -2.25)
                _ = v17
                _ = v18
            }
            switch -747 {
            case 1:
                v19 := GenS1{F0: 193, F1: 0}
                _ = v19
            case 127:
                var v20 bool = (v1 <= (v1 | (^(v1))))
                _ = v20
            default:
                v21 := []int{480, 128, 127, 4294967295}
                _ = v21
            }
            var v22 uint = v3
            _ = v22
        }
        v23 := GenS3{F0: 839, F1: -1, F2: 0}
        v24 := []int{288, -905}
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 bool = ((^(v2)) >= (^(v2)))
        for fi27 := 0; fi27 < 20; fi27++ {
            delete(v5, int(v1))
            v28 := GenS4{F0: 32767}
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    fmt.Println(v1, v4, v4)
    v29 := GenS3{F0: 128, F1: -951, F2: 4294967295}
    switch "" {
    case "hello":
        v2 -= v2
    case "":
        v30 := GenS3{F0: -1, F1: 0, F2: v1}
        _ = v30
    default:
        v31 := GenS2{F0: 127, F1: -276}
        _ = v31
    }
    switch "hello" {
    case "hello":
        v1 -= v1
    case "":
        if (!(563 > v2)) {
            var v32 uint16 = 0
            v33 := GenS5{F0: 213}
            _ = v32
            _ = v33
        }
    case "\n\t":
        var v34 uint = v3
        _ = v34
    default:
        switch 1 {
        case -2147483648:
            var v35 uint32 = (v2 & v2)
            _ = v35
        case 32768:
            v36 := []int{1, -9223372036854775808}
            _ = v36
        case 128:
            v4 = v4
        }
    }
    v2 += v2
    v37 := []int{-149}
    var v38 int = 9223372036854775807
    v2 += v2
    _ = len(v37)
    v38 += v38
    for fi39 := 0; fi39 < 13; fi39++ {
        v4 = 42
        var v40 int = (^(v38))
        var v41 float64 = ((0.1 * (-2.25 * 1.5)) * (0.0 * (3.14159 - (((-2.25 - float64(3.14159)) * 1.5) * (float64(1e10) / float64(1e10))))))
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
