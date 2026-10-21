package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
type GenS3 struct { F0 bool; F1 int64 }
func genFn0(p0 uint64, p1 int) int8 {
    var v7 int32 = 42
    var v8 uint64 = p0
    v9 := []int{p1, (p1 | p1), (p1 * (p1 | -2147483648)), (p1 / ((p1) | 1))}
    var v10 string = ("foo bar" + "foo bar")
    if true {
        fmt.Println(p0, v7, p0)
        var v11 uint32 = 990
        var v12 uint32 = (v11 | v11)
        _ = v11
        _ = v12
    } else {
        v13 := GenS1{F0: 2}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() bool {
    var v14 uint64 = 217
    for fi15 := 0; fi15 < 19; fi15++ {
        switch fi15 {
        case 817:
            switch fi15 {
            case -660:
                var v16 int = (fi15 << ((255) & 63))
                _ = v16
            }
        }
        fmt.Println(v14, v14)
        var v17 bool = true
        _ = fi15
        _ = v17
    }
    var v18 bool = (!true)
    v18 = true
    var v19 string = "foo bar"
    _ = v14
    _ = v18
    _ = v19
    return true
}

func main() {
    var v1 float64 = 3.14159
    v2 := []int{690, 32767, -558, 127}
    var v3 uint16 = 42
    var v4 uint = 1
    var v5 string = ""
    v6 := make(map[uint32]uint32)
    var v20 int64 = int64(v3)
    var v21 int = -215
    v20 = v20
    var v22 float32 = ((0.0 + (1.5 + 3.14159)) * 3.14159)
    for fi23 := 0; fi23 < 8; fi23++ {
        for fi24 := 0; fi24 < 11; fi24++ {
            v25 := GenS1{F0: (v4 * (v4 << ((v4) & 63)))}
            if true {
                if false {
                    _ = len(v2)
                } else {
                    v20 = -1
                    switch v5 {
                    case "foo bar":
                        v26 := v6[1]
                        if false {
                            var v27 bool = true
                            _ = v27
                        } else {
                            if false {
                                fmt.Println(v4)
                            }
                        }
                        v28 := GenS1{F0: (v4 | 18446744073709551615)}
                        _ = len(v2)
                        if (!(-559 == 1000)) {
                            v4 += v4
                            v26 -= 2
                        } else {
                            v26 = v26
                            if (-594 > v20) {
                                fmt.Println(v26, v21, fi23)
                                v4 = (^(v4))
                            } else {
                                v26 = (^(v26))
                                fmt.Println(v26, fi23)
                            }
                            v29 := GenS1{F0: 116}
                            var v30 int = 9223372036854775807
                            _ = v29
                            _ = v30
                        }
                        v4 = (^(v4))
                        _ = v26
                        _ = v28
                    default:
                        fmt.Println(v4)
                    }
                }
            } else {
                v31 := GenS2{F0: (v4 - v4), F1: 127, F2: 1000}
                var v32 string = v5
                _ = v31
                _ = v32
            }
            v6[1] = 65535
            v21 = 127
            _ = fi24
            _ = v25
        }
        var v33 uint8 = 2
        for fi34 := 0; fi34 < 2; fi34++ {
            _ = len(v2)
            v35 := GenS2{F0: (v4 | v4), F1: 127, F2: 1000}
            var v36 uint = v4
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi23
        _ = v33
    }
    v4 += (^(v4))
    v37 := GenS3{F0: (!(v4 != (v4 >> ((v4) & 63)))), F1: (-(v20))}
    switch v21 {
    case 475:
        v38 := GenS2{F0: 964, F1: 100, F2: -1}
        _ = v38
    case 127:
        var v39 uint8 = 128
        _ = v39
    case 32767:
        switch "" {
        case "a":
            v5 = v5
        }
    }
    switch v21 {
    case 266:
        fmt.Println(v4, v22)
    case 4294967295:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 2147483647:
        var v40 int64 = (v20 | (^(v20)))
        _ = v40
    }
    v21 -= v21
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
