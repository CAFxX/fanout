package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int16) string {
    p0 = -32768
    v7 := GenS1{F0: 42}
    if (988 != 882) {
        var v8 int16 = (p0 << ((808) & 63))
        v9 := []int{32767, 2147483647, -927, 1}
        v10 := GenS1{F0: 0}
        _ = v10
        _ = v8
        _ = v9
    }
    p0 -= (p0 & p0)
    if true {
        if (1 > 42) {
            var v11 int8 = 127
            fmt.Println(v11, p0)
            var v12 float32 = (1.5 * 0.0)
            _ = v11
            _ = v12
        } else {
            v13 := GenS2{F0: -9223372036854775808, F1: 42}
            _ = v13
        }
    }
    _ = p0
    _ = v7
    return ("a" + ((("\n\t" + "\n\t") + "foo bar") + (("" + "") + ((((("a" + "a") + "") + "a") + "\n\t") + (("" + "\n\t") + "foo bar")))))
}
func genFn1(p0 string, p1 int64, p2 bool) int32 {
    var v14 bool = (false && ("\n\t" >= p0))
    v15 := []int{128, 4294967295}
    if true {
        v16 := GenS3{F0: 992}
        fmt.Println(v14, v14, p2)
        _ = v16
    }
    fmt.Println(v14)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 942
}

func main() {
    var v1 float64 = (1.5 - 1.5)
    v2 := []string{"a"}
    v3 := make(map[uint32]bool)
    var v4 int8 = 127
    var v5 float64 = v1
    var v6 uint8 = 127
    for fi17 := 0; fi17 < 2; fi17++ {
        v4 = v4
        for fi18 := 0; fi18 < 9; fi18++ {
            if (!true) {
                var v19 int64 = 42
                v4 += (v4 ^ (v4 & (^(v4))))
                _ = v19
            } else {
                var v20 int = fi17
                var v21 int64 = -767
                _ = v20
                _ = v21
            }
            _ = fi18
        }
        _ = fi17
    }
    v22 := GenS4{F0: 324, F1: 1}
    v23 := v3[1]
    var v24 float64 = v1
    for fi25 := 0; fi25 < 4; fi25++ {
        switch "" {
        case "foo bar":
            var v26 uint8 = v6
            _ = v26
        case "a":
            for fi27 := 0; fi27 < 3; fi27++ {
                if v23 {
                    if false {
                        v6 += v6
                        switch (("" + "\n\t") + "foo bar") {
                        case "foo bar":
                            if ((true && true) && ((fi25 % ((fi25) | 1)) < fi25)) {
                                if v23 {
                                    switch (fi27 | fi27) {
                                    case -365:
                                        v24 -= (v1 * v24)
                                    default:
                                        v2 = append(v2, ("foo bar" + "\n\t"))
                                    }
                                } else {
                                    _ = len(v2)
                                    v28 := GenS1{F0: (v4 % (((v4 & (^(v4)))) | 1))}
                                    _ = v28
                                }
                                if true {
                                    var v29 bool = false
                                    _ = v29
                                } else {
                                    v6 -= (v6 - v6)
                                }
                                fmt.Println(fi27)
                            }
                        }
                        v5 = v5
                        v5 -= v5
                        v6 -= (v6 * (v6 & (v6 % ((1) | 1))))
                    }
                }
                _ = len(v3)
                fmt.Println(v23)
                _ = fi27
            }
        case "":
            var v30 int32 = 2147483647
            _ = v30
        }
        _ = fi25
    }
    v3[4294967295] = (163 != 907)
    var v31 string = "a"
    var v32 float32 = (1.5 / ((-2.25 + (3.14159 / 0.1)) / (3.14159 * (1e10 * 3.14159))))
    if false {
        var v33 float64 = v5
        if v23 {
            v23 = (1 > int32(v4))
            if v23 {
                delete(v3, 0)
            } else {
                var v34 int32 = 2147483647
                fmt.Println(v1, v34)
                _ = v34
            }
        }
        _ = v33
    } else {
        var v35 int32 = 2147483647
        _ = v35
    }
    var v36 int16 = 42
    switch 2147483647 {
    case -9223372036854775808:
        var v37 uint8 = v6
        _ = v37
    }
    if v23 {
        if v23 {
            if (-85 <= 0) {
                fmt.Println(v32)
                for fi38 := 0; fi38 < 16; fi38++ {
                    v4 = v4
                    var v39 uint = 172
                    _ = fi38
                    _ = v39
                }
            } else {
                v40 := GenS5{F0: 0}
                _ = v40
            }
        }
    }
    v24 += (((v24 + v24) * v24) - v24)
    switch 32768 {
    case 42:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 2147483647:
        for fi41 := 0; fi41 < 6; fi41++ {
            var v42 int8 = v4
            v32 += float32((float64((v32 * (v32 + 3.14159))) + float64(float32(v5))))
            v43 := GenS3{F0: 103}
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
