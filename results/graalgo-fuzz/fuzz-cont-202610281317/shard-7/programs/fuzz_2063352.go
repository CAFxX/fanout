package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    fmt.Println(v1)
    switch "foo bar" {
    case "foo bar":
        v4 := []int{-9223372036854775808, 9223372036854775807, -2147483648}
        _ = v4
    default:
        switch 960 {
        case 9223372036854775807:
            switch ("\n\t" + "a") {
            case "":
                var v5 float32 = 1.5
                _ = v5
            case "\n\t":
                v1 = (((v1 / (v1 - 0.1)) + float64(0.0)) / 0.1)
            }
        case -2147483648:
            if true {
                switch "foo bar" {
                case "a":
                    v6 := GenS1{F0: 32768, F1: -1, F2: 65535}
                    _ = v6
                }
            } else {
                var v7 bool = (32768 != -1)
                switch 127 {
                case 344:
                    fmt.Println(v7)
                case 9223372036854775807:
                    v8 := GenS1{F0: 42, F1: -77, F2: 60}
                    _ = v8
                case -2147483648:
                    var v9 int32 = -914
                    _ = v9
                }
                v10 := []int{-1, 9223372036854775807, -166}
                _ = v10
                _ = v7
            }
        case -59:
            if false {
                v11 := []int{32767, 32767, 2147483647, 1}
                v12 := GenS1{F0: 0, F1: 127, F2: 892}
                _ = v11
                _ = v12
            }
        default:
            v1 = v1
        }
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS1{F0: 127, F1: 1, F2: 1}
        _ = fi13
        _ = v14
    }
    v1 = v1
    var v15 int = -966
    v16 := GenS2{F0: -8}
    v17 := GenS3{F0: 65535, F1: 65535}
    v18 := GenS2{F0: 1}
    fmt.Println(v1, v1)
    switch (v15 * 548) {
    case -1:
        var v19 uint16 = 65535
        _ = v19
    case 915:
        var v20 uint32 = 65535
        _ = v20
    case 9223372036854775807:
        v1 += 1e10
    default:
        for fi21 := 0; fi21 < 1; fi21++ {
            if ((!(float64((1e10 / 0.1)) > v1)) && true) {
                if false {
                    var v22 float64 = v1
                    if ((98 <= 585) || (42 > 4294967295)) {
                        v15 = (fi21 % ((v15) | 1))
                    } else {
                        v23 := GenS4{F0: 589, F1: 42}
                        _ = v23
                    }
                    _ = v22
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    if ((42 < 0) || true) {
                        var v25 int = fi21
                        var v26 float32 = 1e10
                        _ = v25
                        _ = v26
                    } else {
                        v27 := []int{}
                        _ = v27
                    }
                    var v28 bool = false
                    v15 -= 0
                    _ = fi24
                    _ = v28
                }
            } else {
                v15 = v15
                for fi29 := 0; fi29 < 17; fi29++ {
                    v30 := GenS5{F0: 32767}
                    switch ("foo bar" + "\n\t") {
                    case "a":
                        v1 = v1
                    case "foo bar":
                        var v31 string = "foo bar"
                        _ = v31
                    }
                    if (!false) {
                        v32 := []int{(v15 | (fi21 + (v15 + 128))), fi29, 467}
                        if (489 == 292) {
                            delete(v3, 0)
                            if true {
                                var v33 int = v15
                                _ = v33
                            }
                        }
                        var v34 int64 = 378
                        _ = v32
                        _ = v34
                    }
                    _ = fi29
                    _ = v30
                }
            }
            var v35 int = fi21
            v15 = v15
            _ = fi21
            _ = v35
        }
    }
    var v36 int = (v15 & v15)
    var v37 uint16 = 904
    for fi38 := 0; fi38 < 7; fi38++ {
        fmt.Println(v1)
        _ = fi38
    }
    v37 -= (^(v37))
    var v39 uint32 = 713
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v39
}
