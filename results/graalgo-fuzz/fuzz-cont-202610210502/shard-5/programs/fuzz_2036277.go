package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int32; F2 int8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint) bool {
    p0 = p0
    v6 := GenS1{F0: 2}
    var v7 uint64 = 320
    if (817 == 910) {
        var v8 float64 = (-2.25 + ((((1e10 * 1e10) + 3.14159) / 1.0) + ((1e10 - 0.0) + 0.0)))
        switch "" {
        case "":
            var v9 int64 = 311
            _ = v9
        case "hello":
            for fi10 := 0; fi10 < 8; fi10++ {
                fmt.Println(v8)
                _ = fi10
            }
        }
        _ = v8
    } else {
        for fi11 := 0; fi11 < 6; fi11++ {
            if ((!false) && false) {
                v12 := []int{}
                _ = v12
            }
            if (!true) {
                p0 += 2
            }
            v13 := []int{fi11, fi11, -546}
            _ = fi11
            _ = v13
        }
        var v14 uint16 = 0
        _ = v14
    }
    _ = p0
    _ = v6
    _ = v7
    return (!true)
}
func genFn1() int {
    if (0 > 129) {
        var v15 float64 = (0.0 - 0.0)
        var v16 uint32 = 352
        _ = v15
        _ = v16
    }
    if (-32 == -1) {
        var v17 float64 = -2.25
        v18 := GenS1{F0: 127}
        _ = v17
        _ = v18
    }
    switch "foo bar" {
    case "a":
        var v19 uint8 = 42
        _ = v19
    default:
        v20 := []int{4294967295, -532}
        _ = v20
    }
    return -901
}

func main() {
    var v1 int = 32768
    var v2 int16 = -219
    v3 := make(map[string]int)
    v4 := make(map[string]int)
    var v5 float64 = -2.25
    var v21 int16 = v2
    fmt.Println(v21, v5, v21)
    var v22 int32 = 2147483647
    v5 -= (3.14159 + float64(-2.25))
    if true {
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS2{F0: 1, F1: (v22 ^ (v22 ^ v22)), F2: -128}
            v1 = genFn1()
            v25 := GenS2{F0: (v22 >> ((155) & 63)), F1: v22, F2: -1}
            _ = fi23
            _ = v24
            _ = v25
        }
        v5 = float64(3.14159)
        for fi26 := 0; fi26 < 9; fi26++ {
            _ = len(v4)
            _ = fi26
        }
    } else {
        var v27 uint16 = 1
        var v28 int8 = 100
        v29 := []int{(v1 / ((v1) | 1)), 9223372036854775807, v1}
        _ = v27
        _ = v28
        _ = v29
    }
    if (!false) {
        v21 -= (v2 * (v21 ^ v21))
    } else {
        delete(v4, "a")
        v30 := GenS1{F0: 2}
        _ = v30
    }
    if false {
        if (!(v21 < 1000)) {
            if true {
                v21 = (^(v2))
                fmt.Println(v5, v21, v21)
                if (596 > 42) {
                    switch ("foo bar" + "foo bar") {
                    case "\n\t":
                        var v31 uint8 = 128
                        _ = v31
                    case "hello":
                        fmt.Println(v1)
                    case "foo bar":
                        fmt.Println(v22, v21, v5)
                    default:
                        _ = len(v3)
                    }
                    v5 += v5
                    fmt.Println(v22)
                }
            }
        } else {
            v32 := []int{(v1 + (v1 << ((42) & 63))), 2147483647}
            v21 = -1
            var v33 int32 = v22
            _ = v32
            _ = v33
        }
        var v34 float32 = ((1.5 * (1e10 - (((((((0.1 * -2.25) * ((((1e10 / (-2.25 + 1e10)) - -2.25) * (0.0 / 1.0)) / ((1e10 * 1e10) + 0.0))) - 1.5) * 3.14159) + float32(v5)) + 3.14159) / 1.0))) - (((1.5 + 3.14159) - -2.25) + float32(0.1)))
        _ = v34
    }
    if false {
        fmt.Println(v22)
    } else {
        v1 = (-(v1))
        for fi35 := 0; fi35 < 8; fi35++ {
            var v36 int = (v1 / ((genFn1()) | 1))
            v37 := GenS3{F0: 55}
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    if true {
        for fi38 := 0; fi38 < 16; fi38++ {
            var v39 int64 = -1
            switch (v1 + (fi38 >> ((18446744073709551615) & 63))) {
            case 127:
                v1 = -1
            }
            v40 := []int{128, 2147483647, genFn1()}
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
