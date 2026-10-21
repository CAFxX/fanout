package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool }
func genFn0(p0 int8) int64 {
    p0 = 127
    v7 := GenS1{F0: 718, F1: 127}
    v8 := GenS1{F0: 2147483647, F1: 1}
    _ = p0
    _ = v7
    _ = v8
    return -999
}
func genFn1(p0 int32, p1 int16, p2 int16) int {
    switch "a" {
    case "a":
        v9 := GenS1{F0: -9223372036854775808, F1: 255}
        _ = v9
    case "":
        var v10 uint32 = 65535
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := []int{}
            if false {
                if false {
                    var v13 float32 = (0.0 / 1.0)
                    fmt.Println(p2, p1, p2)
                    _ = v13
                }
                for fi14 := 0; fi14 < 7; fi14++ {
                    p1 += p1
                    p2 -= p2
                    _ = fi14
                }
            } else {
                p0 += (p0 >> ((395) & 63))
                _ = len(v12)
                v12 = append(v12, (fi11 | (^(fi11))))
            }
            fmt.Println(p2)
            _ = fi11
            _ = v12
        }
    default:
        v15 := GenS1{F0: 127, F1: 42}
        _ = v15
    }
    var v16 int64 = -9223372036854775808
    var v17 uint = 504
    for fi18 := 0; fi18 < 20; fi18++ {
        v17 -= v17
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 26
}

func main() {
    v1 := []string{}
    v2 := []string{}
    v3 := make(map[string]bool)
    var v4 int8 = 127
    var v5 uint8 = 127
    var v6 uint64 = 99
    var v19 int64 = 0
    var v20 string = "hello"
    v21 := GenS1{F0: 32768, F1: 1}
    v2 = append(v2, (v20 + ("" + "hello")))
    var v22 uint64 = v6
    v23 := v3[v20]
    var v24 int8 = 0
    if (42 != 0) {
        _ = len(v1)
        var v25 bool = v23
        _ = v25
    }
    if true {
        if false {
            if v23 {
                v5 = v5
                switch v20 {
                case "":
                    v20 = "foo bar"
                }
                _ = len(v1)
            } else {
                for fi26 := 0; fi26 < 13; fi26++ {
                    v20 = "hello"
                    for fi27 := 0; fi27 < 2; fi27++ {
                        _ = len(v2)
                        _ = fi27
                    }
                    switch (fi26 % ((fi26) | 1)) {
                    case 128:
                        switch fi26 {
                        case -9223372036854775808:
                            v20 = (v20 + "a")
                        case 32768:
                            _ = len(v1)
                        case 276:
                            for fi28 := 0; fi28 < 16; fi28++ {
                                fmt.Println(v22)
                                var v29 uint8 = v5
                                _ = fi28
                                _ = v29
                            }
                        default:
                            if (2 <= 65535) {
                                var v30 int16 = 0
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v31 float64 = (3.14159 * 0.1)
                                _ = v30
                                _ = v31
                            }
                        }
                    }
                    _ = fi26
                }
                fmt.Println(v5, v4, v6)
                fmt.Println(v6)
            }
        } else {
            fmt.Println(v5, v23, v20)
        }
        if ((v23 && (!v23)) && true) {
            var v32 float32 = 1.5
            _ = v32
        }
    } else {
        v22 = (v6 + (v22 + v22))
        for fi33 := 0; fi33 < 13; fi33++ {
            v4 += 127
            _ = fi33
        }
        if (!((243 < 33) || true)) {
            var v34 uint32 = 2147483647
            v20 = (("foo bar" + (v20 + v20)) + (v20 + (v20 + v20)))
            _ = v34
        }
    }
    var v35 uint64 = (v6 % ((v6) | 1))
    v36 := GenS2{F0: v19}
    v22 -= (^(v22))
    for fi37 := 0; fi37 < 9; fi37++ {
        v38 := GenS3{F0: ((0.0 * 1e10) <= (1.5 - -2.25))}
        v2 = append(v2, (((v20 + (v20 + v20)) + v20) + (v20 + v20)))
        _ = fi37
        _ = v38
    }
    fmt.Println(v5, v22)
    var v39 string = v20
    var v40 uint16 = 255
    fmt.Println(v40, v24)
    v3[v20] = (((v19 >> ((18446744073709551615) & 63)) >= 9223372036854775807) && v23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
