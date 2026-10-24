package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int64 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint64 = 2
    var v2 int = -9223372036854775808
    var v3 int64 = -1
    var v4 int64 = -9223372036854775808
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 float32 = 1e10
        v7 := []int{(fi5 - (v2 ^ (fi5 & 9223372036854775807)))}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 1000, F1: (v2 & v2), F2: -1}
    fmt.Println(v4, v4)
    v9 := GenS1{F0: -32768, F1: -1, F2: (v3 / ((v4) | 1))}
    var v10 float64 = (0.1 / (3.14159 - 0.1))
    v11 := []int{(^(v2)), 127, v2}
    switch ("" + "a") {
    case "":
        fmt.Println(v2, v3)
    case "foo bar":
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int64 = v4
            v1 -= v1
            if true {
                fmt.Println(v2)
                var v14 int = fi12
                if true {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    v3 += (v4 << ((727) & 63))
                } else {
                    if true {
                        _ = len(v11)
                        v2 -= (-(v14))
                        if false {
                            switch "" {
                            case "a":
                                if (-364 < 1) {
                                    v15 := GenS2{F0: v14}
                                    v16 := GenS2{F0: v14}
                                    var v17 uint16 = 1
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                } else {
                                    switch (fi12 % ((1) | 1)) {
                                    case -1:
                                        if (false && true) {
                                            fmt.Println(v3)
                                            for fi18 := 0; fi18 < 8; fi18++ {
                                                var v19 int8 = -123
                                                v20 := GenS2{F0: (-(v2))}
                                                _ = fi18
                                                _ = v19
                                                _ = v20
                                            }
                                            var v21 string = ("" + "\n\t")
                                            _ = v21
                                        } else {
                                            v22 := GenS2{F0: (v14 ^ (v14 >> ((936) & 63)))}
                                            var v23 uint = 18446744073709551615
                                            v23 = (v23 * (v23 / ((v23) | 1)))
                                            _ = v22
                                            _ = v23
                                        }
                                    case 32768:
                                        fmt.Println(v1, v14)
                                    case 32767:
                                        fmt.Println(fi12, fi12, v10)
                                    }
                                    if (-127 <= 22) {
                                        var v24 uint64 = v1
                                        var v25 float32 = float32(((1.5 + ((v10 + v10) + float64(0.0))) + 3.14159))
                                        if false {
                                            var v26 int = 32768
                                            fmt.Println(v26)
                                            _ = v26
                                        } else {
                                            if (true || ((-2.25 - 0.0) <= (0.0 + (v10 + (0.0 - 1e10))))) {
                                                v11 = append(v11, (v14 ^ -9223372036854775808))
                                                var v27 uint32 = 4294967295
                                                var v28 uint64 = (v1 % ((uint64(v27)) | 1))
                                                _ = v27
                                                _ = v28
                                            }
                                            v29 := GenS2{F0: v2}
                                            _ = v29
                                        }
                                        _ = v24
                                        _ = v25
                                    } else {
                                        v3 -= v3
                                    }
                                }
                            case "\n\t":
                                var v30 bool = ((59 > 0) && (v14 == (v14 >> ((255) & 63))))
                                _ = v30
                            default:
                                var v31 bool = false
                                _ = v31
                            }
                            v13 += (v4 / ((-1) | 1))
                            switch "foo bar" {
                            case "hello":
                                _ = len(v11)
                            case "":
                                var v32 int8 = 1
                                _ = v32
                            default:
                                v13 += (-(v13))
                            }
                        }
                    }
                }
                _ = v14
            } else {
                if (true || (!(!false))) {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        var v34 string = ""
                        _ = fi33
                        _ = v34
                    }
                }
            }
            _ = fi12
            _ = v13
        }
    }
    v35 := GenS2{F0: v2}
    var v36 int32 = 0
    switch "foo bar" {
    case "a":
        for fi37 := 0; fi37 < 13; fi37++ {
            var v38 uint64 = (v1 * 0)
            v4 -= v4
            fmt.Println(v10, v36)
            _ = fi37
            _ = v38
        }
    case "\n\t":
        var v39 int8 = -128
        _ = v39
    }
    switch -9223372036854775808 {
    case 822:
        if len(v11) > 0 {
            v11[0] = v2
        }
    case 32767:
        var v40 int = (^(v2))
        _ = v40
    }
    v41 := GenS1{F0: 1000, F1: (^(v2)), F2: v4}
    if true {
        var v42 int = (^(v2))
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v8
    _ = v9
}
