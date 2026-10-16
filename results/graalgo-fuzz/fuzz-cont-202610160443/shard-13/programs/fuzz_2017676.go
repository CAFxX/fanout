package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 uint; F2 uint }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 int64; F1 uint }
func genFn0(p0 uint8) uint32 {
    var v6 int16 = -1
    v6 += v6
    _ = p0
    _ = v6
    return 1
}
func genFn1(p0 bool, p1 uint32, p2 uint32) uint32 {
    var v7 float32 = 3.14159
    v8 := []int{32767, -881, 4294967295, 1}
    v7 = v7
    for fi9 := 0; fi9 < 12; fi9++ {
        if p0 {
            for fi10 := 0; fi10 < 7; fi10++ {
                if p0 {
                    switch ("hello" + ("a" + "\n\t")) {
                    case "\n\t":
                        fmt.Println(p2)
                    case "foo bar":
                        if (!p0) {
                            if true {
                                switch fi10 {
                                case 128:
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    switch (^(fi10)) {
                                    case 2147483647:
                                        v11 := GenS1{F0: genFn0(127), F1: "hello"}
                                        p1 = genFn0(128)
                                        _ = v11
                                    case 128:
                                        v12 := GenS1{F0: p1, F1: "foo bar"}
                                        _ = v12
                                    }
                                    fmt.Println(fi9)
                                case 127:
                                    switch (fi10 & fi9) {
                                    case 128:
                                        var v13 uint16 = 255
                                        var v14 int16 = 42
                                        _ = v13
                                        _ = v14
                                    case 32768:
                                        var v15 uint = 378
                                        _ = v15
                                    default:
                                        switch 1 {
                                        case -901:
                                            v7 = 0.1
                                            _ = len(v8)
                                            if len(v8) > 0 {
                                                v8[0] = fi10
                                            }
                                        }
                                    }
                                case 32767:
                                    if ((fi9 ^ (fi9 - fi9)) > (fi10 | fi10)) {
                                        var v16 bool = false
                                        _ = v16
                                    } else {
                                        p0 = (!true)
                                        fmt.Println(v7, v7, v7)
                                        var v17 uint32 = (p2 | genFn0(1))
                                        _ = v17
                                    }
                                default:
                                    switch fi9 {
                                    case 128:
                                        if len(v8) > 0 {
                                            v8[0] = 42
                                        }
                                    }
                                }
                                p1 -= genFn0(2)
                                fmt.Println(p1)
                            } else {
                                v18 := GenS1{F0: p1, F1: ("foo bar" + (("foo bar" + "\n\t") + "\n\t"))}
                                var v19 float32 = v7
                                switch -127 {
                                case 1:
                                    var v20 uint32 = (p1 >> ((65535) & 63))
                                    _ = v20
                                default:
                                    var v21 uint8 = 130
                                    _ = v21
                                }
                                var v22 uint64 = 18446744073709551615
                                if len(v8) > 0 {
                                    v8[0] = (-(fi10))
                                }
                                var v23 float64 = (-2.25 / 0.1)
                                _ = v18
                                _ = v19
                                _ = v22
                                _ = v23
                            }
                        }
                        _ = len(v8)
                        v7 = float32(0.0)
                    }
                    if p0 {
                        v7 = (((v7 * (float32(1e10) - (float32(1.5) + 1e10))) * 0.0) * 0.0)
                    } else {
                        var v24 uint32 = p2
                        p1 = v24
                        _ = v24
                    }
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
                v25 := GenS1{F0: (p1 & genFn0(87)), F1: "foo bar"}
                var v26 uint16 = 45
                _ = fi10
                _ = v25
                _ = v26
            }
            v27 := GenS2{F0: 1000, F1: 42, F2: 0}
            _ = v27
        }
        var v28 int8 = -1
        _ = fi9
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p2 >> ((2) & 63))
}

func main() {
    v1 := make(map[string]int)
    var v2 int = 32767
    var v3 uint32 = 732
    var v4 uint64 = 371
    var v5 int16 = -32768
    fmt.Println(v3, v4)
    if (false || (1 < 28)) {
        v29 := v1[""]
        _ = v29
    } else {
        v4 -= v4
        var v30 uint8 = 255
        _ = v30
    }
    var v31 bool = true
    v32 := GenS3{F0: 140, F1: 42}
    fmt.Println(v2)
    fmt.Println(v3, v3)
    v31 = v31
    for fi33 := 0; fi33 < 8; fi33++ {
        _ = len(v1)
        if true {
            var v34 bool = false
            v5 += (-(v5))
            v4 -= 232
            _ = v34
        } else {
            switch (("a" + "") + "foo bar") {
            case "a":
                v2 = (fi33 % (((fi33 | fi33)) | 1))
            case "\n\t":
                v4 += (v4 + v4)
            }
            var v35 string = "foo bar"
            _ = v35
        }
        var v36 int32 = -2147483648
        _ = fi33
        _ = v36
    }
    _ = len(v1)
    if v31 {
        v31 = (-169 > 15)
        if v31 {
            var v37 int8 = -1
            for fi38 := 0; fi38 < 5; fi38++ {
                for fi39 := 0; fi39 < 4; fi39++ {
                    var v40 uint16 = 1
                    _ = fi39
                    _ = v40
                }
                v41 := GenS4{F0: (!((v2 | v2) > v2)), F1: true}
                _ = fi38
                _ = v41
            }
            _ = v37
        }
    } else {
        v42 := []int{(-(v2)), 814}
        if (-1 > 42) {
            fmt.Println(v2, v5, v2)
            fmt.Println(v2, v2)
            v4 += uint64(v3)
        } else {
            v43 := GenS5{F0: 66}
            var v44 uint64 = v4
            _ = len(v1)
            _ = v43
            _ = v44
        }
        _ = v42
    }
    v45 := GenS6{F0: v31, F1: 296, F2: 1}
    if v31 {
        v46 := GenS7{F0: (v2 * v2), F1: 1}
        for fi47 := 0; fi47 < 13; fi47++ {
            if (-111 < -128) {
                v5 += 1000
            }
            if (!v31) {
                switch 395 {
                case -9223372036854775808:
                    var v48 int16 = int16(fi47)
                    _ = v48
                }
            }
            v49 := GenS8{F0: 9223372036854775807, F1: 42}
            _ = fi47
            _ = v49
        }
        fmt.Println(v2, v5)
        _ = v46
    } else {
        v50 := v1["foo bar"]
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v45
    _ = v5
}
