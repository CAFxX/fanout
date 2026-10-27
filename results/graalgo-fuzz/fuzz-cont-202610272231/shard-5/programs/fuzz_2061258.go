package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int, p1 uint64, p2 uint8) uint64 {
    if false {
        var v5 float32 = 1e10
        var v6 float64 = (((1.5 + ((-2.25 / 1.5) * 1e10)) + -2.25) / ((0.0 - (-2.25 + (1.5 * 1.5))) + 1.5))
        if true {
            switch "hello" {
            case "a":
                p0 += (p0 % ((9223372036854775807) | 1))
            }
        }
        _ = v5
        _ = v6
    }
    var v7 int32 = 1
    var v8 int = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 int64, p1 int8) uint8 {
    v9 := []int{128}
    fmt.Println(p1)
    fmt.Println(p1, p1)
    for fi10 := 0; fi10 < 18; fi10++ {
        p0 -= (p0 - p0)
        var v11 string = "a"
        p0 += p0
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return 255
}

func main() {
    v1 := []uint32{2147483647, 764, 4294967295}
    var v2 float64 = 1.5
    var v3 string = "hello"
    v4 := []int{-241, 1, 9223372036854775807}
    if false {
        if (32767 != 1) {
            var v12 string = (v3 + v3)
            _ = len(v1)
            _ = v12
        }
    } else {
        v2 += ((v2 - 0.0) * v2)
    }
    if true {
        if ((float32(1.5) / -2.25) < 1.5) {
            var v13 uint16 = 1
            _ = v13
        } else {
            switch v3 {
            case "hello":
                v2 = 3.14159
            }
        }
        _ = len(v4)
        v14 := GenS1{F0: 1}
        _ = v14
    } else {
        var v15 int8 = -128
        _ = v15
    }
    v4 = append(v4, -193)
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int32 = 2147483647
        v18 := GenS1{F0: 938}
        _ = fi16
        _ = v17
        _ = v18
    }
    switch (v3 + "a") {
    case "foo bar":
        var v19 int64 = 332
        _ = v19
    default:
        var v20 int = 1
        _ = v20
    }
    v3 = (v3 + (v3 + ("" + v3)))
    v2 += v2
    fmt.Println(v3, v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v4)
    fmt.Println(v3)
    var v21 int8 = 1
    switch v3 {
    case "\n\t":
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := GenS1{F0: 2}
            v24 := GenS1{F0: 1}
            _ = fi22
            _ = v23
            _ = v24
        }
    case "hello":
        fmt.Println(v3, v3, v3)
    case "":
        if true {
            v2 -= v2
            v21 += v21
            var v25 int64 = -1
            _ = v25
        } else {
            switch -221 {
            case 4294967295:
                if false {
                    v26 := GenS1{F0: 71}
                    _ = v26
                }
            case 9223372036854775807:
                if true {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v2 += v2
                        var v28 int8 = (v21 % ((127) | 1))
                        _ = fi27
                        _ = v28
                    }
                } else {
                    v21 += v21
                }
            case -1:
                v3 = v3
            }
        }
    }
    v29 := GenS1{F0: 2147483647}
    for fi30 := 0; fi30 < 6; fi30++ {
        v3 = (v3 + v3)
        v3 = (v3 + ("\n\t" + ("foo bar" + v3)))
        switch fi30 {
        case 42:
            if true {
                var v31 int16 = 1
                _ = v31
            } else {
                if (-1 != 0) {
                    if false {
                        v2 += (v2 + (v2 / v2))
                    }
                } else {
                    v2 -= ((1e10 + (v2 * v2)) + float64(float32(v2)))
                }
                for fi32 := 0; fi32 < 20; fi32++ {
                    v2 = (((3.14159 + float64((0.1 - float32(v2)))) + v2) / ((-2.25 + v2) + v2))
                    fmt.Println(fi32, v2)
                    _ = fi32
                }
                v33 := GenS1{F0: 2147483647}
                _ = v33
            }
        }
        _ = fi30
    }
    v4 = append(v4, 32768)
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v4
}
