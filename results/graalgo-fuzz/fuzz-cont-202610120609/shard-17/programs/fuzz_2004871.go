package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int8; F2 int }
type GenS4 struct { F0 bool; F1 int16; F2 int8 }
type GenS5 struct { F0 int32; F1 uint8; F2 int64 }
type GenS6 struct { F0 int; F1 int32; F2 int64 }
type GenS7 struct { F0 uint8; F1 int; F2 uint32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 int32 }
func genFn0() int {
    var v6 int = 42
    v6 += (v6 % (((v6 % (((v6 - v6)) | 1))) | 1))
    var v7 uint16 = 255
    v8 := []int{}
    _ = v6
    _ = v7
    _ = v8
    return (v6 % (((v6 - (v6 % (((v6 ^ v6)) | 1)))) | 1))
}
func genFn1(p0 int16, p1 int32) uint64 {
    v9 := GenS1{F0: "hello", F1: 1}
    if true {
        p1 = p1
        if (false || (((779 < 42) && false) && ("a" < "hello"))) {
            switch 0 {
            case 524:
                p0 = p0
            case 759:
                fmt.Println(p0)
            case -963:
                v10 := GenS2{F0: 1}
                _ = v10
            }
        } else {
            fmt.Println(p1)
            v11 := GenS3{F0: 1, F1: 127, F2: genFn0()}
            _ = v11
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    var v1 uint = 4294967295
    v2 := []string{("hello" + "\n\t"), "\n\t", "hello", (("foo bar" + "") + "a")}
    var v3 int16 = 42
    v4 := []string{}
    var v5 uint64 = 317
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    if true {
        if false {
            if true {
                var v12 uint64 = (v5 ^ 1)
                v13 := GenS3{F0: 604, F1: -111, F2: 795}
                v14 := GenS4{F0: ((v12 << (((v1 >> ((v1) & 63))) & 63)) < v5), F1: (v3 & -52), F2: -33}
                _ = v12
                _ = v13
                _ = v14
            } else {
                if (false || false) {
                    for fi15 := 0; fi15 < 5; fi15++ {
                        fmt.Println(v3, fi15, v5)
                        _ = fi15
                    }
                } else {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v5 = 4294967295
                        _ = fi16
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v17 := GenS5{F0: 2147483647, F1: 1, F2: 9}
                    _ = v17
                }
                if true {
                    v18 := GenS6{F0: 487, F1: -1, F2: -9223372036854775808}
                    _ = v18
                } else {
                    var v19 uint16 = 65535
                    if false {
                        var v20 uint32 = 0
                        fmt.Println(v5)
                        var v21 string = "a"
                        _ = v20
                        _ = v21
                    } else {
                        var v22 uint8 = 0
                        v5 -= v5
                        _ = v22
                    }
                    _ = v19
                }
                for fi23 := 0; fi23 < 11; fi23++ {
                    switch "" {
                    case "a":
                        if true {
                            for fi24 := 0; fi24 < 14; fi24++ {
                                v1 = (^(v1))
                                var v25 int = (fi24 * (fi24 << ((v1) & 63)))
                                _ = fi24
                                _ = v25
                            }
                        }
                    case "hello":
                        for fi26 := 0; fi26 < 14; fi26++ {
                            switch int(v5) {
                            case 128:
                                switch fi26 {
                                case 552:
                                    var v27 int16 = (^(v3))
                                    _ = v27
                                case 0:
                                    v5 = v5
                                }
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            var v28 int8 = 1
                            _ = fi26
                            _ = v28
                        }
                    }
                    _ = len(v2)
                    _ = fi23
                }
            }
        }
        v5 = genFn1(v3, 0)
        var v29 int8 = -81
        _ = v29
    }
    if len(v4) > 0 {
        v4[0] = "a"
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        fmt.Println(v1, fi30, fi30)
        if (false || (-2147483648 <= -2147483648)) {
            var v31 uint64 = genFn1(v3, 2147483647)
            var v32 int16 = v3
            v33 := GenS7{F0: 89, F1: fi30, F2: 194}
            _ = v31
            _ = v32
            _ = v33
        } else {
            var v34 uint = 2
            if false {
                fmt.Println(v1, v3, v1)
                var v35 int16 = (v3 >> ((v34) & 63))
                _ = v35
            }
            _ = v34
        }
        _ = fi30
    }
    v3 = 0
    if false {
        switch "a" {
        case "":
            var v36 bool = false
            _ = v36
        case "\n\t":
            var v37 uint32 = 2147483647
            _ = v37
        case "hello":
            var v38 string = ((((("\n\t" + "foo bar") + ("a" + "a")) + "") + ("foo bar" + "\n\t")) + "\n\t")
            _ = v38
        }
        v39 := GenS4{F0: true, F1: -32768, F2: -128}
        _ = v39
    } else {
        switch ("\n\t" + "hello") {
        case "a":
            for fi40 := 0; fi40 < 6; fi40++ {
                if ((fi40 + genFn0()) == (fi40 / ((fi40) | 1))) {
                    v41 := GenS8{F0: 433}
                    _ = v41
                } else {
                    var v42 float64 = 0.0
                    v42 = 0.0
                    _ = v42
                }
                v5 += genFn1(-546, 2147483647)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi40
            }
        case "hello":
            if true {
                var v43 float32 = -2.25
                var v44 uint8 = 0
                _ = v43
                _ = v44
            }
        case "foo bar":
            v1 = uint(v5)
        default:
            var v45 uint64 = v5
            _ = v45
        }
    }
    if false {
        switch "a" {
        case "foo bar":
            v46 := GenS9{F0: (v5 - 4294967295), F1: -641}
            _ = v46
        case "a":
            var v47 int64 = -9223372036854775808
            _ = v47
        case "\n\t":
            if true {
                v48 := GenS9{F0: v5, F1: -2147483648}
                _ = v48
            } else {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
            }
        default:
            var v49 uint8 = 255
            _ = v49
        }
    }
    fmt.Println(v1, v3, v1)
    var v50 uint64 = genFn1(v3, -1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
}
