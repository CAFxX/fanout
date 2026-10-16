package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int8, p1 uint8) int8 {
    p0 = p0
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 bool = (42 != 232)
        var v7 float32 = 1e10
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint8 = p1
            v6 = v6
            v10 := GenS1{F0: v9}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v11 := GenS2{F0: 4294967295, F1: p1}
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            p0 = p0
            _ = fi13
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v11
    return p0
}
func genFn1(p0 string) bool {
    if (float64(-2.25) <= float64(1.5)) {
        p0 = p0
        v14 := []int{2147483647, -2147483648, -215, -9223372036854775808}
        _ = v14
    } else {
        if true {
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 int = (fi15 / (((fi15 << ((861) & 63))) | 1))
                var v17 int8 = genFn0(-1, 42)
                fmt.Println(p0)
                _ = fi15
                _ = v16
                _ = v17
            }
            v18 := GenS3{F0: 218, F1: 1}
            var v19 float64 = 0.0
            _ = v18
            _ = v19
        }
        var v20 float64 = 1.5
        _ = v20
    }
    p0 = p0
    p0 = ((p0 + p0) + "a")
    v21 := []int{}
    _ = p0
    _ = v21
    return true
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 0
    var v3 uint32 = 366
    var v4 uint8 = 127
    v22 := GenS1{F0: (v4 >> ((18446744073709551615) & 63))}
    var v23 int8 = 100
    fmt.Println(v3)
    var v24 uint = 960
    fmt.Println(v4, v4, v2)
    switch ("a" + v1) {
    case "hello":
        var v25 float32 = (1.5 - 3.14159)
        _ = v25
    }
    var v26 int64 = 0
    switch ("" + ("\n\t" + "\n\t")) {
    case "hello":
        var v27 int32 = 2147483647
        _ = v27
    case "foo bar":
        var v28 uint16 = 0
        _ = v28
    case "":
        fmt.Println(v1, v26)
    }
    v29 := []int{(v2 ^ v2), v2}
    v2 = v2
    v26 -= (v26 / ((int64(v4)) | 1))
    if true {
        if len(v29) > 0 {
            v29[0] = -280
        }
        if false {
            var v30 int32 = 0
            _ = v30
        }
        if false {
            v26 = (v26 ^ (^(v26)))
            v31 := GenS3{F0: (v4 + uint8(v3)), F1: v4}
            if false {
                for fi32 := 0; fi32 < 4; fi32++ {
                    if (true || false) {
                        var v33 float32 = float32(float64(1e10))
                        v3 -= v3
                        _ = v33
                    } else {
                        fmt.Println(v4)
                        fmt.Println(v23)
                        switch (v2 * (-(v2))) {
                        case 4294967295:
                            var v34 uint8 = (^(v4))
                            _ = v34
                        }
                    }
                    v35 := GenS1{F0: v4}
                    _ = fi32
                    _ = v35
                }
            }
            _ = v31
        }
    } else {
        v26 = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
