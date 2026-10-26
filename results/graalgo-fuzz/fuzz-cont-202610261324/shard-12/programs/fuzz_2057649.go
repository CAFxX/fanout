package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 string; F2 int16 }
type GenS4 struct { F0 int32; F1 uint64 }
func genFn0(p0 string) uint16 {
    var v7 float32 = (0.1 * -2.25)
    if (4294967295 != 384) {
        p0 = ((p0 + p0) + "foo bar")
        v8 := []int{4294967295}
        _ = v8
    }
    var v9 int16 = -32768
    v10 := GenS1{F0: 628, F1: 562, F2: -1}
    if false {
        var v11 string = (p0 + "foo bar")
        _ = v11
    }
    _ = p0
    _ = v10
    _ = v7
    _ = v9
    return 375
}
func genFn1(p0 int32, p1 string, p2 int16) int {
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int = fi12
        v14 := GenS1{F0: 42, F1: 854, F2: 0}
        var v15 float32 = 1.5
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    switch p1 {
    case "\n\t":
        v16 := []int{128}
        _ = v16
    }
    fmt.Println(p2, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return -492
}

func main() {
    var v1 int16 = -1
    v2 := []uint32{317, 42, 42, 2}
    var v3 uint32 = 138
    var v4 float32 = 1e10
    var v5 uint = 86
    var v6 int = 9223372036854775807
    var v17 int32 = -2147483648
    fmt.Println(v6, v1)
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := GenS2{F0: 244}
        if ((float64((0.1 / (0.0 + -2.25))) + ((0.1 - (((0.1 + (float64(v4) * 0.1)) - 1.5) - 0.1)) / float64(v4))) >= 0.1) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi20 := 0; fi20 < 4; fi20++ {
                v3 -= 42
                v21 := GenS2{F0: uint8(v3)}
                var v22 uint8 = 2
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        if true {
            v3 = v3
            var v23 string = "foo bar"
            if ((!(v5 > v5)) && true) {
                for fi24 := 0; fi24 < 13; fi24++ {
                    v6 = 32768
                    v3 -= v3
                    v1 += 0
                    _ = fi24
                }
            }
            _ = v23
        }
        _ = fi18
        _ = v19
    }
    var v25 int32 = (v17 | (v17 | -910))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = v1
    fmt.Println(v5)
    switch ("hello" + ("\n\t" + ("hello" + "hello"))) {
    case "foo bar":
        var v26 bool = ((-1 >= 1) || true)
        _ = v26
    default:
        _ = len(v2)
    }
    switch v6 {
    case 0:
        if (42 < 44) {
            fmt.Println(v1)
        } else {
            fmt.Println(v25, v25)
            if (255 == genFn0("a")) {
                var v27 uint32 = (^(v3))
                fmt.Println(v3, v17)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v27
            }
            fmt.Println(v25)
        }
    case -1:
        for fi28 := 0; fi28 < 5; fi28++ {
            var v29 float64 = 3.14159
            _ = fi28
            _ = v29
        }
    default:
        if (!false) {
            v30 := GenS3{F0: 558, F1: "a", F2: 32767}
            var v31 uint8 = 222
            v31 = 1
            _ = v30
            _ = v31
        } else {
            v5 = (v5 / (((v5 << (((^(v5))) & 63))) | 1))
            v32 := GenS4{F0: int32(v5), F1: 2}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v32
        }
    }
    var v33 uint64 = 42
    if (v17 < (v17 % ((v25) | 1))) {
        var v34 uint64 = (v33 % ((1) | 1))
        _ = v34
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        if ((!true) && ((false || false) || (true && false))) {
            v5 = v5
            var v36 uint = 1
            _ = v36
        } else {
            if false {
                v37 := GenS2{F0: 210}
                for fi38 := 0; fi38 < 15; fi38++ {
                    var v39 bool = (!true)
                    if v39 {
                        v40 := GenS4{F0: (v25 ^ int32(v33)), F1: (v33 % ((v33) | 1))}
                        var v41 int64 = 42
                        _ = v40
                        _ = v41
                    }
                    _ = fi38
                    _ = v39
                }
                for fi42 := 0; fi42 < 9; fi42++ {
                    v25 -= (v25 + v17)
                    if false {
                        v3 += (v3 + 2147483648)
                        fmt.Println(v25, v6)
                    }
                    _ = fi42
                }
                _ = v37
            } else {
                var v43 uint = 255
                fmt.Println(v6)
                for fi44 := 0; fi44 < 10; fi44++ {
                    v45 := GenS2{F0: 26}
                    _ = fi44
                    _ = v45
                }
                _ = v43
            }
            v46 := GenS2{F0: 0}
            _ = v46
        }
        if (true || false) {
            if (v25 > (v25 / (((v25 ^ -1)) | 1))) {
                switch (^(fi35)) {
                case 32768:
                    if len(v2) > 0 {
                        v2[0] = (^(v3))
                    }
                }
                v25 -= (^(v25))
                v3 = 112
            } else {
                var v47 float32 = v4
                if false {
                    if (false && ((v25 << ((v5) & 63)) < (-(v17)))) {
                        var v48 int = (^(fi35))
                        v4 -= (v4 * 1.5)
                        _ = v48
                    }
                    switch (fi35 * (-(v6))) {
                    case 32768:
                        _ = len(v2)
                    }
                }
                _ = v47
            }
            v25 += v17
        } else {
            if len(v2) > 0 {
                v2[0] = v3
            }
        }
        v17 += v25
        _ = fi35
    }
    var v49 int32 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
