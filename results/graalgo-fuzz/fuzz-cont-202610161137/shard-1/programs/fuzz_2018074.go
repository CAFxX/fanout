package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
func genFn0() int {
    var v4 uint64 = 18446744073709551615
    v5 := GenS1{F0: 100, F1: 1}
    _ = v4
    _ = v5
    return 128
}

func main() {
    v1 := []int{}
    var v2 bool = true
    v3 := []int{}
    if v2 {
        v2 = false
        fmt.Println(v2)
    } else {
        v2 = v2
        if v2 {
            _ = len(v3)
        }
        v6 := GenS2{F0: 135}
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            v2 = (int16(fi8) <= -1)
            _ = fi8
        }
        _ = fi7
    }
    v1 = append(v1, -53)
    v9 := GenS1{F0: -106, F1: 681}
    switch ("hello" + "\n\t") {
    case "foo bar":
        if (v2 && (!false)) {
            switch -9223372036854775808 {
            case 32767:
                v2 = ("" != "foo bar")
            case -2147483648:
                var v10 int16 = 1
                _ = v10
            default:
                v11 := GenS3{F0: 122}
                _ = v11
            }
            var v12 float32 = 0.0
            _ = v12
        }
    case "":
        var v13 float32 = float32(1.5)
        _ = v13
    case "hello":
        for fi14 := 0; fi14 < 7; fi14++ {
            _ = len(v3)
            var v15 int8 = -95
            _ = fi14
            _ = v15
        }
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    v2 = (!true)
    v2 = v2
    v16 := GenS3{F0: 100}
    var v17 int32 = 1
    if v2 {
        v2 = v2
        v2 = v2
    }
    switch 1 {
    case 4294967295:
        fmt.Println(v2, v17)
    case 762:
        v2 = v2
    case 1:
        for fi18 := 0; fi18 < 1; fi18++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v17 = v17
            _ = len(v3)
            _ = fi18
        }
    default:
        fmt.Println(v2, v2)
    }
    v19 := GenS4{F0: 69}
    fmt.Println(v17, v17)
    v20 := GenS1{F0: -1, F1: 1000}
    if true {
        var v21 uint64 = 2
        for fi22 := 0; fi22 < 18; fi22++ {
            for fi23 := 0; fi23 < 10; fi23++ {
                v2 = true
                _ = fi23
            }
            if ((0.1 >= 0.1) || v2) {
                var v24 uint = 671
                _ = v24
            } else {
                if true {
                    switch ("hello" + "\n\t") {
                    case "":
                        var v25 int64 = 42
                        _ = v25
                    case "foo bar":
                        var v26 uint64 = v21
                        _ = v26
                    case "a":
                        v27 := GenS4{F0: (v21 - v21)}
                        _ = v27
                    }
                    v28 := GenS1{F0: 127, F1: -32768}
                    _ = v28
                }
                var v29 uint16 = 65535
                v30 := GenS1{F0: 0, F1: 883}
                _ = v29
                _ = v30
            }
            switch (fi22 - (fi22 & int(v17))) {
            case -9223372036854775808:
                v17 = (^(v17))
            case 42:
                v17 += v17
            }
            _ = fi22
        }
        var v31 uint8 = 0
        _ = v21
        _ = v31
    }
    v32 := GenS3{F0: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v9
}
