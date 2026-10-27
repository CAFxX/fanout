package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 int32; F2 uint }
func genFn0(p0 int, p1 uint8, p2 int) int {
    p0 = p2
    v4 := GenS1{F0: 127}
    v5 := GenS1{F0: 42}
    v6 := []int{p0, (p0 & p0), (^(p0))}
    var v7 string = ("\n\t" + ("hello" + "\n\t"))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -651
}

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{28, 854, 635, 2}
    v3 := make(map[string]string)
    if ((42 < 1) || (((((("hello" + v1) + "foo bar") + v1) + v1) != ((((v1 + (v1 + (v1 + v1))) + ((v1 + "a") + v1)) + (v1 + ((v1 + "foo bar") + ("foo bar" + v1)))) + ("a" + (v1 + ("a" + (v1 + (v1 + v1))))))) || false)) {
        v1 = "hello"
        v8 := GenS1{F0: 127}
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS2{F0: -9223372036854775808}
            var v11 string = v1
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    }
    if true {
        var v12 uint16 = 32767
        v13 := GenS3{F0: -723, F1: 997, F2: 0}
        v12 -= (v12 + v12)
        _ = v12
        _ = v13
    } else {
        var v14 int64 = -279
        v14 = 367
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 13; fi15++ {
            if (-1 >= 1) {
                delete(v3, v1)
                _ = len(v2)
                if (false || (!false)) {
                    if (712 >= -2147483648) {
                        v3["hello"] = "hello"
                        v2 = append(v2, 42)
                    }
                    var v16 int8 = 100
                    for fi17 := 0; fi17 < 10; fi17++ {
                        switch ((v1 + "") + (v1 + v1)) {
                        case "foo bar":
                            fmt.Println(fi15, fi15)
                        case "hello":
                            switch fi17 {
                            case 32768:
                                if (false && true) {
                                    v3[v1] = v1
                                    delete(v3, v1)
                                } else {
                                    v18 := GenS2{F0: 42}
                                    switch -2147483648 {
                                    case -187:
                                        v16 = -128
                                    case 918:
                                        v16 = v16
                                    case 32768:
                                        fmt.Println(v16, fi17)
                                    }
                                    _ = v18
                                }
                            case 0:
                                v3["hello"] = "\n\t"
                            case 2147483647:
                                v1 = ("foo bar" + ("a" + ("\n\t" + v1)))
                            }
                        default:
                            v3[("foo bar" + v1)] = "hello"
                        }
                        var v19 uint16 = 875
                        _ = fi17
                        _ = v19
                    }
                    _ = v16
                }
            } else {
                var v20 uint16 = 869
                fmt.Println(fi15, v1)
                _ = v20
            }
            _ = fi15
        }
        var v21 uint = 65535
        _ = v21
    }
    var v22 uint8 = 143
    fmt.Println(v1)
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := GenS3{F0: -1, F1: -835, F2: 65535}
        if (!false) {
            delete(v3, (v1 + v1))
        }
        _ = fi23
        _ = v24
    }
    v25 := GenS2{F0: 9223372036854775807}
    var v26 float32 = (3.14159 + 1e10)
    fmt.Println(v1)
    v22 = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
}
