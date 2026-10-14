package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int8; F1 int64; F2 int16 }
func genFn0() uint32 {
    var v7 uint16 = 65535
    fmt.Println(v7)
    if false {
        if (42 != 18446744073709551615) {
            var v8 bool = true
            v9 := GenS1{F0: 1000}
            _ = v8
            _ = v9
        } else {
            var v10 uint16 = 42
            var v11 float32 = 3.14159
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint = 912
            _ = fi12
            _ = v13
        }
        var v14 int16 = 0
        _ = v14
    }
    var v15 uint64 = 1
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 uint = 255
        var v18 int64 = -1
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = v15
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 int64, p2 uint) int32 {
    var v19 bool = (!(!(-308 <= p1)))
    v20 := GenS2{F0: 47, F1: (-(p0))}
    switch ((("a" + ("\n\t" + ("foo bar" + "hello"))) + "a") + "") {
    case "a":
        if v19 {
            for fi21 := 0; fi21 < 17; fi21++ {
                v22 := GenS3{F0: (-128 <= -128), F1: 65535}
                v23 := []int{}
                var v24 float64 = 0.1
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            fmt.Println(p1, p0, v19)
        } else {
            var v25 float64 = ((float64((-2.25 + 0.0)) * (-2.25 + 3.14159)) * 0.0)
            _ = v25
        }
    case "hello":
        var v26 int16 = p0
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 0
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[int]int)
    var v3 uint = 98
    var v4 uint64 = 0
    v5 := []uint32{257}
    var v6 string = "hello"
    v1 -= v1
    v1 += 1e10
    switch "a" {
    case "a":
        v1 += ((float32(1.5) + 0.0) * ((v1 - (v1 / v1)) - 3.14159))
    case "":
        for fi27 := 0; fi27 < 9; fi27++ {
            if (!(-1 != 42)) {
                delete(v2, (fi27 % (((fi27 + int(v4))) | 1)))
            }
            fmt.Println(v1, v6, v1)
            _ = fi27
        }
    case "\n\t":
        for fi28 := 0; fi28 < 20; fi28++ {
            v4 += v4
            var v29 uint16 = 32767
            var v30 bool = ((("" + (v6 + v6)) != v6) || (false || false))
            _ = fi28
            _ = v29
            _ = v30
        }
    default:
        if true {
            for fi31 := 0; fi31 < 17; fi31++ {
                v6 = (v6 + "\n\t")
                _ = fi31
            }
            v6 = ("foo bar" + (v6 + (v6 + "\n\t")))
            if true {
                v1 += 3.14159
                var v32 uint8 = 213
                for fi33 := 0; fi33 < 14; fi33++ {
                    if false {
                        switch v6 {
                        case "\n\t":
                            v34 := GenS4{F0: fi33, F1: -1}
                            _ = v34
                        case "a":
                            fmt.Println(v6, fi33, fi33)
                        case "":
                            v4 = v4
                        }
                    } else {
                        var v35 int64 = -9223372036854775808
                        _ = v35
                    }
                    v36 := GenS1{F0: 188}
                    _ = fi33
                    _ = v36
                }
                _ = v32
            }
        }
    }
    var v37 int16 = -32768
    fmt.Println(v37, v6)
    for fi38 := 0; fi38 < 15; fi38++ {
        var v39 bool = true
        _ = fi38
        _ = v39
    }
    v1 = float32(-2.25)
    if false {
        for fi40 := 0; fi40 < 13; fi40++ {
            if (!true) {
                var v41 int32 = -242
                var v42 int64 = -9223372036854775808
                fmt.Println(v1, v3, v6)
                _ = v41
                _ = v42
            } else {
                _ = len(v5)
                delete(v2, (fi40 * (fi40 + (fi40 & (-(fi40))))))
            }
            fmt.Println(v3, v6, v3)
            _ = fi40
        }
        v43 := GenS5{F0: 1, F1: 0, F2: (v37 * 32767)}
        _ = v43
    }
    var v44 bool = false
    v45 := GenS4{F0: 2147483647, F1: -9}
    var v46 float64 = 1e10
    v47 := GenS3{F0: (true || ((v6 + v6) == "hello")), F1: genFn0()}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v4, v44)
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
