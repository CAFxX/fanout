package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
func genFn0(p0 int, p1 string) int {
    p1 = (p1 + (p1 + p1))
    for fi5 := 0; fi5 < 14; fi5++ {
        switch 1 {
        case 117:
            v6 := GenS1{F0: 1}
            _ = v6
        case -1:
            p0 -= (p0 * (fi5 ^ (fi5 >> ((1) & 63))))
        }
        var v7 uint32 = 2
        v8 := []int{p0, (fi5 << ((18446744073709551615) & 63))}
        _ = fi5
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: -2}
    fmt.Println(p1, p1)
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(fi10, p1, fi10)
        if (float64(((0.0 + (((1.5 / -2.25) / 1e10) + float32(1e10))) - (0.1 * 3.14159))) >= 0.0) {
            var v11 int16 = -461
            for fi12 := 0; fi12 < 15; fi12++ {
                p0 -= (fi10 >> ((65535) & 63))
                if true {
                    var v13 uint32 = 190
                    v14 := []int{(fi10 >> ((1) & 63)), (fi10 | -1), fi12}
                    var v15 bool = true
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = fi12
            }
            _ = v11
        } else {
            for fi16 := 0; fi16 < 7; fi16++ {
                switch p0 {
                case 259:
                    if false {
                        p1 = p1
                    } else {
                        v17 := GenS1{F0: -13}
                        _ = v17
                    }
                case 860:
                    v18 := GenS1{F0: -128}
                    _ = v18
                case 0:
                    v19 := GenS1{F0: 127}
                    _ = v19
                }
                _ = fi16
            }
            p0 = (p0 * p0)
            if false {
                var v20 uint64 = 18446744073709551615
                _ = v20
            }
        }
        p0 = p0
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return (p0 & p0)
}
func genFn1(p0 uint32, p1 string) int16 {
    p0 = 993
    if false {
        if true {
            if true {
                v21 := GenS1{F0: -1}
                _ = v21
            } else {
                var v22 string = ((((p1 + (((p1 + "") + "hello") + "a")) + (p1 + ("foo bar" + p1))) + "foo bar") + (p1 + ((((p1 + ("foo bar" + p1)) + "a") + (p1 + ((p1 + p1) + ""))) + (((p1 + (("" + "a") + ("\n\t" + "foo bar"))) + ((p1 + (p1 + (("hello" + "") + "foo bar"))) + p1)) + p1))))
                for fi23 := 0; fi23 < 17; fi23++ {
                    if true {
                        v24 := []int{(fi23 >> ((42) & 63))}
                        for fi25 := 0; fi25 < 14; fi25++ {
                            var v26 int16 = 296
                            var v27 int16 = v26
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        if (667 >= 355) {
                            v22 = v22
                            v28 := GenS1{F0: -14}
                            for fi29 := 0; fi29 < 17; fi29++ {
                                p0 -= 1
                                _ = fi29
                            }
                            _ = v28
                        }
                        _ = v24
                    }
                    var v30 int32 = 42
                    v31 := GenS1{F0: -128}
                    _ = fi23
                    _ = v30
                    _ = v31
                }
                _ = v22
            }
            p1 = "a"
            fmt.Println(p0, p1)
        }
        if true {
            switch ("a" + p1) {
            case "a":
                p0 += 656
            case "hello":
                p0 = (p0 / ((p0) | 1))
            case "foo bar":
                fmt.Println(p0)
            }
            v32 := GenS1{F0: 42}
            _ = v32
        }
        var v33 uint64 = 914
        _ = v33
    }
    v34 := GenS1{F0: 1}
    p0 = p0
    _ = p0
    _ = p1
    _ = v34
    return 926
}

func main() {
    var v1 int = 42
    var v2 int16 = 42
    var v3 string = ((("" + (("\n\t" + "foo bar") + "foo bar")) + ("" + ("foo bar" + ""))) + "foo bar")
    var v4 uint16 = 1
    switch v3 {
    case "hello":
        for fi35 := 0; fi35 < 20; fi35++ {
            var v36 float32 = 3.14159
            _ = fi35
            _ = v36
        }
    case "foo bar":
        v37 := GenS2{F0: 178, F1: v3}
        _ = v37
    case "\n\t":
        var v38 int16 = (v2 / ((v2) | 1))
        _ = v38
    default:
        v2 += -1
    }
    v2 = 113
    fmt.Println(v4, v3)
    switch v3 {
    case "":
        v39 := GenS3{F0: uint32(v4), F1: 2}
        _ = v39
    }
    v1 = genFn0((v1 * 205), "foo bar")
    v2 -= -1
    v40 := []int{(v1 | v1), (^(v1))}
    for fi41 := 0; fi41 < 18; fi41++ {
        if true {
            fmt.Println(v1, v1, v4)
        }
        for fi42 := 0; fi42 < 18; fi42++ {
            fmt.Println(v1, v2, v4)
            fmt.Println(fi41)
            _ = fi42
        }
        _ = fi41
    }
    if true {
        if true {
            var v43 int8 = -128
            _ = v43
        } else {
            v44 := GenS4{F0: -400, F1: 409, F2: true}
            switch (v1 ^ (^(v1))) {
            case -9223372036854775808:
                var v45 int = -651
                _ = v45
            case 32768:
                _ = len(v40)
            case 128:
                var v46 int64 = 1
                _ = v46
            }
            _ = v44
        }
        for fi47 := 0; fi47 < 3; fi47++ {
            var v48 float64 = float64(1.5)
            _ = fi47
            _ = v48
        }
        v49 := GenS4{F0: 0, F1: 1, F2: true}
        _ = v49
    }
    v4 += (v4 & v4)
    v2 = (v2 / (((v2 ^ (v2 / ((genFn1(2147483647, "\n\t")) | 1)))) | 1))
    v50 := GenS2{F0: 128, F1: "hello"}
    for fi51 := 0; fi51 < 20; fi51++ {
        v52 := GenS3{F0: 2147483648, F1: 0}
        _ = fi51
        _ = v52
    }
    if (65535 <= 255) {
        var v53 uint8 = 172
        var v54 uint64 = 0
        _ = v53
        _ = v54
    }
    v4 -= (v4 & 255)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v50
}
