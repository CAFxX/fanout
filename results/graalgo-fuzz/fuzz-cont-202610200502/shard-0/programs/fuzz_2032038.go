package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }

func main() {
    var v1 int = 780
    var v2 float64 = (((float64(1e10) - -2.25) - (((3.14159 + 0.0) * 1.5) + ((1e10 / 1.5) + (1e10 - 1.5)))) / (3.14159 / ((0.1 / 0.1) + float64((1.5 + float32(1.5))))))
    var v3 string = "foo bar"
    var v4 uint8 = 42
    var v5 int = (v1 - (v1 >> ((1) & 63)))
    v6 := []string{v3, ("hello" + v3), v3, v3}
    v6 = append(v6, ((v3 + "a") + "foo bar"))
    v5 = v1
    fmt.Println(v5, v2, v3)
    if true {
        if (589 >= v5) {
            v7 := GenS1{F0: (!((1e10 + float32(0.1)) >= 1e10)), F1: 255, F2: (32767 > 0)}
            switch v3 {
            case "hello":
                fmt.Println(v3, v2)
            case "\n\t":
                v8 := GenS2{F0: (v3 + "foo bar")}
                _ = v8
            case "foo bar":
                v9 := GenS1{F0: false, F1: 4294967295, F2: true}
                _ = v9
            }
            _ = v7
        }
        v2 += (v2 * v2)
    } else {
        for fi10 := 0; fi10 < 19; fi10++ {
            for fi11 := 0; fi11 < 8; fi11++ {
                v6 = append(v6, "\n\t")
                _ = fi11
            }
            _ = fi10
        }
        fmt.Println(v5, v4, v3)
        v2 += (v2 - 3.14159)
    }
    if false {
        v12 := GenS2{F0: v3}
        var v13 int64 = -1
        var v14 uint32 = 2147483647
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 int32 = 1
    v16 := GenS1{F0: (((v3 > v3) && true) && ((490 >= 255) || ((1 != 1) && (true && true)))), F1: uint(v4), F2: false}
    var v17 int32 = v15
    var v18 uint32 = 2
    var v19 uint8 = 0
    var v20 string = ((v3 + "") + v3)
    switch v3 {
    case "":
        var v21 int8 = int8(v18)
        _ = v21
    case "\n\t":
        v3 = ""
    default:
        var v22 uint32 = (v18 | 4294967295)
        _ = v22
    }
    fmt.Println(v17, v15, v5)
    v2 += 0.0
    if (false || ((v18 % (((v18 >> ((151) & 63))) | 1)) >= v18)) {
        v6 = append(v6, v3)
        if false {
            v23 := GenS2{F0: v20}
            var v24 int16 = 38
            for fi25 := 0; fi25 < 11; fi25++ {
                if false {
                    if ((!(0 > -9223372036854775808)) && (!(-1 != 9223372036854775807))) {
                        for fi26 := 0; fi26 < 7; fi26++ {
                            var v27 int32 = v15
                            var v28 float64 = 1e10
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        v29 := GenS3{F0: 202}
                        _ = v29
                    } else {
                        fmt.Println(v4)
                        for fi30 := 0; fi30 < 4; fi30++ {
                            v2 -= v2
                            var v31 uint16 = 40
                            _ = fi30
                            _ = v31
                        }
                        if true {
                            if true {
                                var v32 uint8 = 0
                                _ = v32
                            }
                            v19 -= v19
                        }
                    }
                } else {
                    var v33 int64 = -124
                    var v34 uint16 = 482
                    _ = v33
                    _ = v34
                }
                fmt.Println(v4, v15, v15)
                for fi35 := 0; fi35 < 17; fi35++ {
                    v18 = (^(v18))
                    _ = fi35
                }
                _ = fi25
            }
            _ = v23
            _ = v24
        }
        var v36 int64 = -1
        _ = v36
    }
    v1 = v1
    v17 -= v15
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
