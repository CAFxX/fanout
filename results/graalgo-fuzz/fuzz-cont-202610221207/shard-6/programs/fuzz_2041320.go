package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int }

func main() {
    v1 := []uint32{}
    var v2 uint32 = 426
    var v3 int64 = 808
    var v4 int64 = (v3 ^ 882)
    v5 := make(map[string]uint32)
    var v6 bool = (255 <= 65535)
    switch (("hello" + ("foo bar" + "")) + (("foo bar" + ("" + "a")) + "a")) {
    case "a":
        var v7 int16 = -942
        _ = v7
    default:
        v8 := GenS1{F0: -422, F1: -2147483648, F2: (v2 >> ((uint(v2)) & 63))}
        _ = v8
    }
    var v9 uint16 = 42
    switch 0 {
    case -101:
        fmt.Println(v4, v4)
    case -597:
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := GenS1{F0: (fi11 | 337), F1: 0, F2: 2147483647}
                var v13 uint = 0
                _ = fi11
                _ = v12
                _ = v13
            }
            if v6 {
                v6 = (true || ((1e10 + 3.14159) < (0.1 - 1.5)))
            } else {
                v9 += 0
            }
            if (!v6) {
                if v6 {
                    v4 += (v4 / (((v4 - (v4 % ((v4) | 1)))) | 1))
                } else {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        if v6 {
                            if false {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                fmt.Println(v6)
                            } else {
                                var v15 string = ("a" + "foo bar")
                                v9 -= v9
                                _ = v15
                            }
                        } else {
                            fmt.Println(v4)
                            v16 := GenS1{F0: fi10, F1: 486, F2: (v2 % (((^(v2))) | 1))}
                            v2 = 404
                            _ = v16
                        }
                        var v17 int16 = 32767
                        _ = fi14
                        _ = v17
                    }
                    v18 := GenS1{F0: (-(fi10)), F1: 1, F2: (v2 * v2)}
                    _ = v18
                }
                v5[("hello" + ("foo bar" + "hello"))] = v2
                v19 := GenS2{F0: 1}
                _ = v19
            } else {
                fmt.Println(v3, v2)
            }
            _ = fi10
        }
    default:
        fmt.Println(v3)
    }
    v20 := v5["foo bar"]
    v9 = (v9 << ((681) & 63))
    var v21 int16 = -1
    v6 = (!v6)
    for fi22 := 0; fi22 < 13; fi22++ {
        fmt.Println(v4)
        v23 := GenS1{F0: fi22, F1: 42, F2: (v20 >> ((255) & 63))}
        _ = fi22
        _ = v23
    }
    if (v20 > (^(v2))) {
        if false {
            fmt.Println(v20, v9)
            var v24 int32 = -2147483648
            _ = v24
        } else {
            switch "" {
            case "hello":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "":
                var v25 float64 = 0.1
                _ = v25
            case "a":
                if len(v1) > 0 {
                    v1[0] = (v2 * (v2 * 2))
                }
            }
            v6 = ("" == (((("" + "hello") + "a") + "\n\t") + "hello"))
            var v26 int32 = 2147483647
            _ = v26
        }
        for fi27 := 0; fi27 < 8; fi27++ {
            v28 := GenS1{F0: (-(fi27)), F1: 2147483647, F2: v2}
            _ = fi27
            _ = v28
        }
    }
    switch 32768 {
    case 1:
        fmt.Println(v3)
    default:
        var v29 uint16 = (v9 - 843)
        _ = v29
    }
    v30 := GenS1{F0: 32767, F1: 1, F2: (^(v2))}
    for fi31 := 0; fi31 < 15; fi31++ {
        v32 := GenS3{F0: fi31, F1: fi31}
        v6 = false
        fmt.Println(v21, v4, v2)
        _ = fi31
        _ = v32
    }
    fmt.Println(v21)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
