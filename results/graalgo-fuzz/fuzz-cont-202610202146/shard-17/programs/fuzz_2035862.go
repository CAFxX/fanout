package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int32; F2 int }
type GenS3 struct { F0 int; F1 uint16; F2 int8 }
func genFn0(p0 int32) int64 {
    var v6 uint32 = 2147483648
    var v7 uint64 = 0
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int16 = 1000
        var v10 uint32 = 2147483648
        switch "" {
        case "":
            v6 = (v6 | (v6 - v6))
        default:
            var v11 string = "foo bar"
            _ = v11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v7 = 18446744073709551615
    _ = p0
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint8) int64 {
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := []int{}
            _ = fi13
            _ = v14
        }
        fmt.Println(p0, p0)
        var v15 uint32 = 0
        _ = fi12
        _ = v15
    }
    v16 := GenS1{F0: -2147483648}
    fmt.Println(p0)
    _ = p0
    _ = v16
    return 1
}

func main() {
    v1 := []int{}
    v2 := []uint32{678, 65535}
    var v3 uint16 = 255
    var v4 uint64 = 0
    var v5 uint32 = 1
    v3 = (v3 - 880)
    v4 += v4
    for fi17 := 0; fi17 < 13; fi17++ {
        if true {
            var v18 uint16 = 673
            v19 := GenS1{F0: -2147483648}
            v20 := GenS1{F0: 42}
            _ = v18
            _ = v19
            _ = v20
        } else {
            v5 = 308
            for fi21 := 0; fi21 < 20; fi21++ {
                var v22 float64 = 1.5
                v5 -= 2
                _ = fi21
                _ = v22
            }
            _ = len(v2)
        }
        var v23 float64 = (float64(0.0) * 0.0)
        _ = fi17
        _ = v23
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 int8 = 100
        _ = fi24
        _ = v25
    }
    v3 = 255
    var v26 uint8 = 46
    if (0.1 < 3.14159) {
        v26 = 2
        v27 := GenS1{F0: -1}
        _ = v27
    } else {
        for fi28 := 0; fi28 < 13; fi28++ {
            var v29 int8 = 0
            var v30 uint32 = (v5 - (v5 | (v5 ^ v5)))
            if len(v1) > 0 {
                v1[0] = fi28
            }
            _ = fi28
            _ = v29
            _ = v30
        }
        if len(v1) > 0 {
            v1[0] = 42
        }
        var v31 uint = 156
        _ = v31
    }
    if (0 != 0) {
        fmt.Println(v26, v5)
        fmt.Println(v4, v4, v4)
        if true {
            var v32 uint32 = 42
            var v33 int16 = 42
            for fi34 := 0; fi34 < 16; fi34++ {
                v33 += v33
                var v35 uint64 = v4
                _ = fi34
                _ = v35
            }
            _ = v32
            _ = v33
        } else {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v4, v5)
            case "foo bar":
                var v36 uint = 715
                _ = v36
            default:
                if true {
                    var v37 int = -401
                    var v38 uint = 680
                    _ = v37
                    _ = v38
                }
            }
            v39 := GenS1{F0: 202}
            if false {
                v3 = uint16(v4)
                switch 9223372036854775807 {
                case 304:
                    fmt.Println(v5)
                case 385:
                    switch -768 {
                    case 4294967295:
                        var v40 uint8 = (v26 | (v26 / ((v26) | 1)))
                        _ = v40
                    case 2147483647:
                        v41 := GenS1{F0: 853}
                        _ = v41
                    case -332:
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                    }
                default:
                    _ = len(v2)
                }
            }
            _ = v39
        }
    }
    for fi42 := 0; fi42 < 13; fi42++ {
        for fi43 := 0; fi43 < 17; fi43++ {
            switch ((("a" + ("hello" + (("hello" + ("" + "")) + "foo bar"))) + ((("foo bar" + ((("a" + "hello") + "a") + ("" + "a"))) + "hello") + ("" + (("\n\t" + "") + (("a" + ("" + "\n\t")) + ("" + (("hello" + "") + ("a" + "\n\t")))))))) + (("hello" + ("\n\t" + "foo bar")) + "\n\t")) {
            case "hello":
                v26 -= 1
            case "foo bar":
                var v44 string = ("hello" + "a")
                fmt.Println(v44, v44, v3)
                v45 := GenS1{F0: 2147483647}
                if len(v2) > 0 {
                    v2[0] = v5
                }
                v46 := GenS2{F0: 647, F1: 451, F2: (^(fi42))}
                fmt.Println(v3, v44)
                fmt.Println(v4)
                _ = v44
                _ = v45
                _ = v46
            default:
                switch (("\n\t" + "foo bar") + ("hello" + "foo bar")) {
                case "\n\t":
                    var v47 int64 = genFn0(-1)
                    _ = v47
                default:
                    v3 = 2
                }
            }
            v4 = (^(v4))
            if false {
                _ = len(v1)
            } else {
                v26 = uint8(v3)
                v5 -= (v5 ^ v5)
                if false {
                    v2 = append(v2, 536)
                    v3 = (^(v3))
                } else {
                    _ = len(v2)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            _ = fi43
        }
        _ = fi42
    }
    switch "\n\t" {
    case "foo bar":
        for fi48 := 0; fi48 < 2; fi48++ {
            v5 += (v5 & (v5 - (v5 - (v5 / (((v5 / ((v5) | 1))) | 1)))))
            v49 := GenS3{F0: 4294967295, F1: 436, F2: 1}
            _ = fi48
            _ = v49
        }
    case "hello":
        _ = len(v2)
    default:
        fmt.Println(v4, v5)
    }
    if false {
        v5 -= (v5 ^ (v5 - (v5 % (((v5 & (v5 << ((65535) & 63)))) | 1))))
        v4 -= (^(v4))
    } else {
        v5 -= (v5 * (v5 << ((2) & 63)))
    }
    switch 127 {
    case 128:
        var v50 float32 = (1.5 - (float32(3.14159) * 3.14159))
        _ = v50
    case 9223372036854775807:
        if (42 >= 42) {
            for fi51 := 0; fi51 < 10; fi51++ {
                v3 -= v3
                _ = fi51
            }
        } else {
            var v52 int = 32767
            _ = v52
        }
    default:
        for fi53 := 0; fi53 < 4; fi53++ {
            var v54 float64 = float64(3.14159)
            _ = fi53
            _ = v54
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
