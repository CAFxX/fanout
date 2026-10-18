package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int; F2 int16 }
type GenS3 struct { F0 int16; F1 int8; F2 bool }
func genFn0(p0 string) uint8 {
    v7 := []int{-2147483648}
    switch ("hello" + ("a" + "\n\t")) {
    case "a":
        p0 = ((p0 + p0) + p0)
    case "":
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 int32 = 2147483647
            for fi10 := 0; fi10 < 15; fi10++ {
                v9 -= v9
                v9 = (v9 - (v9 | v9))
                _ = fi10
            }
            var v11 int64 = -9223372036854775808
            _ = fi8
            _ = v11
            _ = v9
        }
    case "foo bar":
        v7 = append(v7, 32767)
    }
    v7 = append(v7, 974)
    p0 = p0
    _ = p0
    _ = v7
    return 35
}
func genFn1(p0 uint16, p1 uint32) uint32 {
    var v12 int8 = 1
    switch "a" {
    case "foo bar":
        for fi13 := 0; fi13 < 14; fi13++ {
            if (!false) {
                if (400 == 42) {
                    for fi14 := 0; fi14 < 19; fi14++ {
                        p1 -= (p1 * p1)
                        switch (("\n\t" + "foo bar") + ("" + "\n\t")) {
                        case "foo bar":
                            var v15 int16 = 32767
                            _ = v15
                        }
                        v12 += (v12 >> ((4294967295) & 63))
                        _ = fi14
                    }
                    v12 = (v12 & 30)
                } else {
                    fmt.Println(fi13, v12)
                    v16 := []int{}
                    _ = v16
                }
            }
            var v17 string = ("hello" + ("\n\t" + ("foo bar" + "a")))
            _ = fi13
            _ = v17
        }
    case "\n\t":
        for fi18 := 0; fi18 < 7; fi18++ {
            v12 = v12
            var v19 uint64 = 20
            fmt.Println(p0, p0, v12)
            _ = fi18
            _ = v19
        }
    case "":
        var v20 int64 = -9223372036854775808
        _ = v20
    }
    fmt.Println(p0)
    var v21 int32 = 1
    _ = p0
    _ = p1
    _ = v12
    _ = v21
    return 2
}

func main() {
    var v1 string = "foo bar"
    v2 := []int{0, -898}
    v3 := make(map[string]uint32)
    var v4 uint32 = 0
    v5 := []uint32{v4}
    var v6 float32 = 0.1
    v5 = append(v5, v4)
    v22 := GenS1{F0: -9223372036854775808}
    if (-32768 < 215) {
        var v23 int16 = 1000
        if false {
            if (648 > 32767) {
                if (false && (!(false && true))) {
                    v6 -= v6
                    if ((!true) && false) {
                        for fi24 := 0; fi24 < 16; fi24++ {
                            fmt.Println(fi24, v6, v6)
                            var v25 int64 = 1
                            var v26 int = fi24
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        if true {
                            v27 := GenS1{F0: 141}
                            _ = v27
                        }
                    } else {
                        v28 := GenS2{F0: 42, F1: 4294967295, F2: (v23 >> ((65535) & 63))}
                        if (true || (v6 == v6)) {
                            v6 = v6
                        }
                        v29 := GenS2{F0: 211, F1: -909, F2: (v23 << ((18446744073709551615) & 63))}
                        _ = v28
                        _ = v29
                    }
                    _ = len(v3)
                }
            } else {
                var v30 int64 = 1
                _ = v30
            }
        }
        _ = v23
    } else {
        var v31 float64 = (((1e10 - (-2.25 + float64(1.5))) + 0.1) - 0.0)
        for fi32 := 0; fi32 < 8; fi32++ {
            var v33 float32 = (v6 * float32(v31))
            var v34 int8 = -42
            if (!(false || (v31 <= v31))) {
                fmt.Println(v31, v4)
                if (!((v4 >= genFn1(42, 145)) || false)) {
                    v35 := GenS3{F0: 1, F1: 1, F2: (false && (!((-74 > v34) || (42 >= 65535))))}
                    fmt.Println(v33)
                    _ = v35
                }
            } else {
                v6 = v33
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v31
    }
    var v36 uint32 = v4
    v1 = ("foo bar" + "a")
    v6 = -2.25
    var v37 uint = 65535
    fmt.Println(v37)
    var v38 int = 651
    if (1 <= v4) {
        if true {
            v39 := GenS1{F0: 42}
            v5 = append(v5, v4)
            _ = v39
        } else {
            if false {
                for fi40 := 0; fi40 < 2; fi40++ {
                    for fi41 := 0; fi41 < 7; fi41++ {
                        var v42 uint8 = 42
                        _ = fi41
                        _ = v42
                    }
                    _ = fi40
                }
            } else {
                for fi43 := 0; fi43 < 2; fi43++ {
                    fmt.Println(v1, v37)
                    _ = fi43
                }
            }
        }
        fmt.Println(v1)
    }
    if (uint(v4) != v37) {
        v38 = (v38 / (((v38 >> ((v37) & 63))) | 1))
        if true {
            if false {
                var v44 uint16 = 945
                _ = v44
            }
            fmt.Println(v1)
        }
    } else {
        var v45 int64 = 42
        switch ("" + (v1 + (v1 + v1))) {
        case "":
            var v46 int32 = 935
            _ = v46
        case "a":
            var v47 int = (v38 / ((v38) | 1))
            _ = v47
        case "foo bar":
            fmt.Println(v6, v1, v37)
        default:
            var v48 uint = (v37 >> ((4294967295) & 63))
            _ = v48
        }
        v36 += (^(v36))
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
