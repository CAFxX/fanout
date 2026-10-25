package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int8 {
    switch "\n\t" {
    case "foo bar":
        var v4 int64 = -9223372036854775808
        _ = v4
    }
    if false {
        v5 := []int{0}
        v6 := GenS1{F0: 42}
        var v7 int64 = -1
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 int16 = 42
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: -2147483648}
        _ = fi9
        _ = v10
    }
    _ = v8
    return 100
}
func genFn1() uint8 {
    v11 := []int{}
    var v12 int = -385
    if false {
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 bool = (32767 < 42)
            if len(v11) > 0 {
                v11[0] = fi13
            }
            fmt.Println(v14)
            _ = fi13
            _ = v14
        }
    } else {
        v12 = (v12 ^ (-(v12)))
    }
    _ = v11
    _ = v12
    return 2
}

func main() {
    var v1 float32 = 3.14159
    var v2 float32 = (((v1 / v1) - float32(1.5)) + 0.1)
    var v3 int32 = -2147483648
    if false {
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := []int{(fi15 + fi15), -939, fi15}
            fmt.Println(v3, v2, v1)
            v17 := GenS1{F0: (v3 - (v3 & v3))}
            _ = fi15
            _ = v16
            _ = v17
        }
    } else {
        switch "\n\t" {
        case "":
            if (float32(0.0) >= ((v1 * (v1 / ((v1 * -2.25) + 0.1))) * 1.5)) {
                if false {
                    v18 := []int{654}
                    if true {
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                    } else {
                        if len(v18) > 0 {
                            v18[0] = 2147483647
                        }
                        v1 -= v1
                        var v19 int64 = -9223372036854775808
                        _ = v19
                    }
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    _ = v18
                } else {
                    v3 = (v3 | v3)
                    v20 := GenS1{F0: v3}
                    var v21 string = "a"
                    _ = v20
                    _ = v21
                }
                fmt.Println(v2)
            } else {
                for fi22 := 0; fi22 < 8; fi22++ {
                    v23 := []int{}
                    v1 = 3.14159
                    v1 += v2
                    _ = fi22
                    _ = v23
                }
            }
        case "hello":
            switch ("hello" + "\n\t") {
            case "\n\t":
                v24 := []int{32767, 32767, 32767}
                _ = v24
            }
        default:
            v25 := []int{-9223372036854775808}
            _ = v25
        }
        switch 212 {
        case 481:
            v26 := GenS1{F0: (v3 & (v3 ^ (v3 * (v3 & v3))))}
            _ = v26
        default:
            var v27 float64 = (float64(1.5) * (1e10 / 1.0))
            _ = v27
        }
        v1 = v2
    }
    v2 = ((0.0 * v2) / float32(1e10))
    if true {
        v28 := GenS1{F0: 0}
        _ = v28
    }
    v2 = (1e10 * float32(3.14159))
    v29 := []int{42, -735, -2147483648}
    for fi30 := 0; fi30 < 16; fi30++ {
        v3 = v3
        for fi31 := 0; fi31 < 16; fi31++ {
            v2 = float32(3.14159)
            v32 := GenS1{F0: (v3 >> ((447) & 63))}
            _ = fi31
            _ = v32
        }
        var v33 int16 = -32768
        _ = fi30
        _ = v33
    }
    v2 = 0.1
    var v34 int8 = 1
    var v35 uint16 = 817
    var v36 uint8 = 171
    var v37 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
