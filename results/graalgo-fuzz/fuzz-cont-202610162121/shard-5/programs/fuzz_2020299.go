package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32 }
func genFn0(p0 int) uint64 {
    p0 = p0
    for fi6 := 0; fi6 < 14; fi6++ {
        p0 = fi6
        _ = fi6
    }
    fmt.Println(p0)
    v7 := []int{p0}
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 bool) bool {
    if (-1 < 576) {
        p0 = p0
    } else {
        p0 = (1 == 42)
    }
    p0 = (17 != 0)
    switch 42 {
    case 9223372036854775807:
        p0 = (0 < -684)
    }
    v8 := []int{2147483647, 42, 1}
    if len(v8) > 0 {
        v8[0] = -913
    }
    _ = p0
    _ = v8
    return p0
}

func main() {
    var v1 float32 = 0.0
    v2 := []uint32{42}
    v3 := []string{((("\n\t" + "hello") + "\n\t") + "a"), (("\n\t" + "") + "a")}
    v4 := make(map[string]uint32)
    var v5 uint8 = 1
    v9 := GenS1{F0: true}
    if false {
        var v10 int = 4294967295
        switch (v10 >> ((52) & 63)) {
        case 2147483647:
            if (false && true) {
                v2 = append(v2, 2)
                v2 = append(v2, 2147483647)
                v5 -= v5
            }
        default:
            if false {
                v11 := GenS1{F0: false}
                fmt.Println(v10)
                var v12 int32 = 1
                _ = v11
                _ = v12
            } else {
                v13 := v4["a"]
                _ = v13
            }
        }
        _ = v10
    } else {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        if true {
            fmt.Println(v1, v1)
            if len(v3) > 0 {
                v3[0] = "a"
            }
            v14 := GenS2{F0: 664, F1: 1, F2: v5}
            _ = v14
        }
    }
    var v15 uint8 = v5
    v16 := GenS3{F0: 42}
    v15 = v5
    v1 = (((v1 + v1) * float32((1e10 * 3.14159))) * v1)
    fmt.Println(v1)
    for fi17 := 0; fi17 < 18; fi17++ {
        fmt.Println(v1, fi17)
        switch "a" {
        case "foo bar":
            switch 127 {
            case -114:
                v15 -= v5
            case 32768:
                v15 -= v5
            case 0:
                if true {
                    if false {
                        v18 := GenS3{F0: 32767}
                        var v19 int32 = 2147483647
                        v15 = (v5 % ((v15) | 1))
                        _ = v18
                        _ = v19
                    }
                    if false {
                        switch fi17 {
                        case 9223372036854775807:
                            var v20 uint32 = 0
                            _ = v20
                        case -9223372036854775808:
                            for fi21 := 0; fi21 < 20; fi21++ {
                                fmt.Println(fi21, v1, fi17)
                                _ = fi21
                            }
                        }
                        v22 := GenS4{F0: 42}
                        if (true && (((!(1e10 > 0.1)) && true) && (-1 == 889))) {
                            v23 := GenS4{F0: -60}
                            for fi24 := 0; fi24 < 7; fi24++ {
                                if (0 >= 944) {
                                    var v25 float64 = -2.25
                                    v5 -= v5
                                    fmt.Println(fi24)
                                    _ = v25
                                } else {
                                    v26 := GenS5{F0: (fi17 << ((274) & 63))}
                                    fmt.Println(v5, v1)
                                    _ = v26
                                }
                                v27 := GenS3{F0: 32767}
                                v28 := v4[("\n\t" + "")]
                                _ = fi24
                                _ = v27
                                _ = v28
                            }
                            var v29 bool = (42 <= 554)
                            _ = v23
                            _ = v29
                        }
                        _ = v22
                    }
                } else {
                    v5 += 0
                    var v30 int64 = 0
                    switch fi17 {
                    case 32768:
                        fmt.Println(v1)
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = v30
                }
            default:
                v5 += (v15 + (v5 & v5))
            }
        case "":
            v31 := v4["a"]
            _ = v31
        case "\n\t":
            var v32 uint16 = 32767
            _ = v32
        }
        switch (("" + ("\n\t" + (("hello" + "foo bar") + "hello"))) + ("\n\t" + ("" + "a"))) {
        case "a":
            fmt.Println(fi17, v15, v1)
        }
        _ = fi17
    }
    v33 := GenS2{F0: 374, F1: -32768, F2: v15}
    for fi34 := 0; fi34 < 2; fi34++ {
        v5 -= v5
        _ = len(v2)
        _ = fi34
    }
    var v35 string = "\n\t"
    switch 357 {
    case -1:
        for fi36 := 0; fi36 < 6; fi36++ {
            var v37 float32 = (3.14159 / -2.25)
            _ = fi36
            _ = v37
        }
    case 0:
        _ = len(v2)
    }
    var v38 bool = true
    for fi39 := 0; fi39 < 14; fi39++ {
        v38 = (377 != 828)
        v40 := GenS6{F0: 42}
        v1 = v1
        _ = fi39
        _ = v40
    }
    if (127 < -1) {
        v5 -= (v5 << ((1) & 63))
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        if (717 >= -9223372036854775808) {
            var v42 uint16 = 2
            _ = v42
        }
        _ = fi41
    }
    if len(v2) > 0 {
        v2[0] = 374
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v9
}
