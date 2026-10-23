package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 string; F1 string; F2 uint64 }
func genFn0() uint16 {
    switch "hello" {
    case "\n\t":
        v7 := GenS1{F0: 121, F1: 2}
        _ = v7
    case "hello":
        var v8 bool = true
        _ = v8
    case "a":
        switch 42 {
        case -691:
            if true {
                v9 := []int{}
                var v10 uint16 = 514
                _ = v10
                _ = v9
            } else {
                var v11 uint16 = 669
                v11 = 42
                v12 := GenS1{F0: -1, F1: 1}
                _ = v11
                _ = v12
            }
        case -322:
            if ((1e10 + (1e10 + (-2.25 + (3.14159 / (3.14159 + float64(float32((0.0 + 0.1)))))))) == float64((0.1 - -2.25))) {
                v13 := []int{-9223372036854775808}
                for fi14 := 0; fi14 < 19; fi14++ {
                    var v15 uint8 = 128
                    _ = fi14
                    _ = v15
                }
                var v16 int32 = -351
                _ = v13
                _ = v16
            } else {
                for fi17 := 0; fi17 < 11; fi17++ {
                    v18 := GenS1{F0: -119, F1: 109}
                    _ = fi17
                    _ = v18
                }
                var v19 float64 = (-2.25 / -2.25)
                _ = v19
            }
        case 4294967295:
            var v20 int16 = 32767
            _ = v20
        default:
            switch 127 {
            case 32768:
                v21 := []int{-442, 127, -267, 0}
                _ = v21
            case 188:
                for fi22 := 0; fi22 < 3; fi22++ {
                    switch ("" + ("" + "a")) {
                    case "a":
                        var v23 float32 = 1e10
                        _ = v23
                    case "":
                        switch fi22 {
                        case 32767:
                            v24 := []int{(fi22 + 32768), (fi22 << ((2) & 63))}
                            _ = v24
                        case 1:
                            var v25 bool = true
                            _ = v25
                        }
                    default:
                        var v26 int64 = 9223372036854775807
                        _ = v26
                    }
                    var v27 int16 = -923
                    _ = fi22
                    _ = v27
                }
            }
        }
    default:
        var v28 uint16 = 65535
        _ = v28
    }
    var v29 int16 = 32767
    _ = v29
    return 42
}

func main() {
    var v1 uint16 = 1
    v2 := []uint32{65535}
    v3 := make(map[string]string)
    var v4 float32 = 0.0
    v5 := make(map[int]uint32)
    v6 := make(map[string]bool)
    v4 = v4
    v30 := GenS1{F0: 42, F1: 222}
    var v31 uint64 = 2
    fmt.Println(v31, v1)
    fmt.Println(v4, v1, v1)
    for fi32 := 0; fi32 < 13; fi32++ {
        switch "foo bar" {
        case "\n\t":
            var v33 float32 = 0.1
            _ = v33
        default:
            if true {
                v1 -= 42
            }
        }
        switch "hello" {
        case "\n\t":
            v34 := GenS2{F0: fi32, F1: 100}
            _ = v34
        case "hello":
            fmt.Println(fi32, v31)
        case "a":
            switch "foo bar" {
            case "foo bar":
                v31 = 188
            }
        }
        var v35 uint = 18446744073709551615
        _ = fi32
        _ = v35
    }
    switch (("a" + "a") + ("a" + "a")) {
    case "":
        for fi36 := 0; fi36 < 9; fi36++ {
            v37 := GenS3{F0: "hello", F1: ("a" + "a"), F2: 2}
            fmt.Println(fi36)
            _ = fi36
            _ = v37
        }
    case "hello":
        for fi38 := 0; fi38 < 5; fi38++ {
            fmt.Println(v1, v31)
            _ = fi38
        }
    }
    v31 -= (v31 ^ v31)
    v39 := GenS1{F0: 1, F1: 128}
    switch "\n\t" {
    case "\n\t":
        delete(v5, 2147483647)
    default:
        var v40 uint16 = 65535
        _ = v40
    }
    var v41 int16 = 800
    if ((0.1 / (1e10 - 1.5)) > 0.1) {
        v4 = 1e10
        delete(v3, "a")
        v1 = (^(v1))
    }
    v41 += 42
    v31 += (v31 + v31)
    var v42 uint = 0
    switch 388 {
    case -108:
        for fi43 := 0; fi43 < 10; fi43++ {
            v44 := v6[(("" + "foo bar") + "")]
            for fi45 := 0; fi45 < 20; fi45++ {
                fmt.Println(fi45)
                _ = fi45
            }
            if len(v2) > 0 {
                v2[0] = 461
            }
            _ = fi43
            _ = v44
        }
    default:
        v1 += 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
