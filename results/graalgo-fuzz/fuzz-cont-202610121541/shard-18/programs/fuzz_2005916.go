package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    var v3 string = "\n\t"
    v4 := []string{}
    var v5 float32 = 3.14159
    delete(v2, 58)
    v3 = "a"
    delete(v2, 719)
    var v6 int64 = -9223372036854775808
    v7 := GenS1{F0: -545}
    v2[0] = (315 >= 2)
    fmt.Println(v5, v6, v3)
    fmt.Println(v5, v3, v3)
    var v8 bool = true
    v8 = v8
    var v9 float32 = (0.1 + (float32((0.0 * 1e10)) * v5))
    if (!(!(!v8))) {
        fmt.Println(v6, v6)
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                fmt.Println(fi10, v5)
                if len(v4) > 0 {
                    v4[0] = v3
                }
                _ = fi11
            }
            var v12 uint8 = 0
            _ = fi10
            _ = v12
        }
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 uint16 = 42
            fmt.Println(v8, v14, v6)
            _ = fi13
            _ = v14
        }
    } else {
        v5 = (v9 * v9)
        switch 42 {
        case 98:
            for fi15 := 0; fi15 < 6; fi15++ {
                _ = len(v1)
                v3 = (v3 + "a")
                if ((-2.25 * (1.5 - 0.0)) >= 1e10) {
                    _ = len(v4)
                } else {
                    var v16 uint = 65535
                    v17 := GenS1{F0: -32768}
                    _ = v16
                    _ = v17
                }
                _ = fi15
            }
        case -911:
            v18 := GenS1{F0: 32767}
            _ = v18
        case 42:
            fmt.Println(v5, v6, v5)
        default:
            var v19 int = 1
            _ = v19
        }
        var v20 bool = (!false)
        _ = v20
    }
    if v8 {
        v8 = false
        fmt.Println(v8)
    }
    if v8 {
        v21 := v2[523]
        if v8 {
            for fi22 := 0; fi22 < 17; fi22++ {
                if len(v4) > 0 {
                    v4[0] = (v3 + (v3 + "foo bar"))
                }
                v23 := v2[fi22]
                delete(v1, fi22)
                _ = fi22
                _ = v23
            }
        } else {
            if ((v21 && (2147483647 > -256)) && false) {
                for fi24 := 0; fi24 < 8; fi24++ {
                    v3 = v3
                    v6 = v6
                    var v25 int = fi24
                    _ = fi24
                    _ = v25
                }
                for fi26 := 0; fi26 < 2; fi26++ {
                    v27 := GenS1{F0: 42}
                    _ = fi26
                    _ = v27
                }
            } else {
                for fi28 := 0; fi28 < 16; fi28++ {
                    delete(v2, 186)
                    v29 := GenS2{F0: fi28, F1: -1}
                    for fi30 := 0; fi30 < 2; fi30++ {
                        if len(v4) > 0 {
                            v4[0] = (((("\n\t" + "") + (v3 + v3)) + "foo bar") + v3)
                        }
                        fmt.Println(v5, v9)
                        _ = fi30
                    }
                    _ = fi28
                    _ = v29
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        fmt.Println(v6, v3, v5)
        _ = v21
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        v32 := GenS2{F0: (^(fi31)), F1: -1}
        _ = fi31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
