package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 bool }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int; F1 string; F2 int8 }
func genFn0() int32 {
    v5 := GenS1{F0: 32767, F1: 128, F2: (true || (!true))}
    var v6 uint = 18446744073709551615
    v7 := []int{-9223372036854775808, 128}
    _ = v5
    _ = v6
    _ = v7
    return 598
}

func main() {
    v1 := []string{"a", "", ("" + ("" + ((((("" + "") + "hello") + "a") + "") + (("a" + ("hello" + ("\n\t" + "foo bar"))) + "")))), "a"}
    v2 := make(map[string]string)
    v3 := []uint32{}
    v4 := []uint32{2147483647, 1}
    if (!(!true)) {
        var v8 string = "\n\t"
        var v9 int64 = 9223372036854775807
        for fi10 := 0; fi10 < 7; fi10++ {
            switch (^(fi10)) {
            case -552:
                v9 -= (v9 << ((4294967295) & 63))
            case 128:
                v11 := GenS1{F0: fi10, F1: 255, F2: false}
                _ = v11
            case 127:
                switch ((((v8 + v8) + v8) + (v8 + "a")) + "") {
                case "a":
                    var v12 int = fi10
                    _ = v12
                case "hello":
                    v9 = 1
                case "":
                    fmt.Println(v8, v9)
                default:
                    var v13 uint32 = 766
                    _ = v13
                }
            default:
                var v14 int8 = -1
                _ = v14
            }
            v15 := GenS1{F0: 32768, F1: 255, F2: true}
            _ = fi10
            _ = v15
        }
        _ = v8
        _ = v9
    }
    var v16 int32 = -248
    switch -472 {
    case -1:
        v16 -= (-(v16))
    case 683:
        var v17 uint32 = 549
        _ = v17
    default:
        fmt.Println(v16)
    }
    switch "a" {
    case "foo bar":
        if (v16 != (v16 ^ genFn0())) {
            v18 := GenS2{F0: (-(v16)), F1: 1}
            if (true && (true && true)) {
                for fi19 := 0; fi19 < 5; fi19++ {
                    fmt.Println(v16, v16)
                    for fi20 := 0; fi20 < 15; fi20++ {
                        var v21 uint64 = 1
                        var v22 int64 = 0
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    _ = fi19
                }
                for fi23 := 0; fi23 < 13; fi23++ {
                    if ((22 >= 45) && ("foo bar" == "foo bar")) {
                        if ((!false) || true) {
                            v16 += (v16 << ((834) & 63))
                        }
                        if false {
                            fmt.Println(fi23, v16)
                            for fi24 := 0; fi24 < 17; fi24++ {
                                v16 = v16
                                if false {
                                    if false {
                                        if (!true) {
                                            _ = len(v2)
                                            fmt.Println(v16, fi23, v16)
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v25 := GenS3{F0: true, F1: -32768}
                                            v16 = 2147483647
                                            _ = v25
                                        }
                                        v16 = genFn0()
                                    }
                                    v26 := GenS2{F0: genFn0(), F1: 42}
                                    v16 = v16
                                    _ = v26
                                } else {
                                    v27 := GenS2{F0: v16, F1: 9223372036854775807}
                                    _ = v27
                                }
                                v28 := GenS3{F0: (261 >= 255), F1: 42}
                                _ = fi24
                                _ = v28
                            }
                        }
                    }
                    var v29 float64 = ((0.0 / float64((1.5 * (((0.1 / ((float32(float64(1e10)) - (3.14159 * 0.0)) + 3.14159)) - float32(float64(-2.25))) / (float32(1.5) + -2.25))))) - -2.25)
                    for fi30 := 0; fi30 < 14; fi30++ {
                        if true {
                            v29 -= v29
                            switch fi30 {
                            case -2147483648:
                                var v31 int32 = (v16 - (v16 ^ genFn0()))
                                _ = v31
                            case 1:
                                v16 += (v16 + genFn0())
                            }
                        }
                        _ = fi30
                    }
                    _ = fi23
                    _ = v29
                }
            } else {
                switch "a" {
                case "":
                    v16 = (v16 & genFn0())
                }
                fmt.Println(v16)
            }
            _ = v18
        } else {
            if (-1 > 0) {
                v32 := v2[(((("foo bar" + "a") + ("a" + "foo bar")) + ("" + "\n\t")) + "hello")]
                switch (v32 + v32) {
                case "\n\t":
                    if (!(((!(1e10 >= (0.1 - 0.1))) || true) || (true || false))) {
                        fmt.Println(v32)
                        v16 = (v16 | 1)
                        var v33 int16 = -32768
                        _ = v33
                    }
                case "a":
                    v1 = append(v1, ((("a" + "a") + (v32 + (((v32 + v32) + v32) + "hello"))) + ""))
                default:
                    if (315 <= 4294967295) {
                        v34 := GenS3{F0: true, F1: -1}
                        var v35 int32 = v16
                        _ = v34
                        _ = v35
                    }
                }
                _ = v32
            } else {
                var v36 float32 = 3.14159
                _ = v36
            }
        }
    }
    fmt.Println(v16)
    for fi37 := 0; fi37 < 15; fi37++ {
        var v38 uint8 = 127
        v2["foo bar"] = "foo bar"
        _ = len(v1)
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 9; fi39++ {
        if true {
            var v40 string = (("hello" + "\n\t") + "a")
            v41 := GenS4{F0: fi39, F1: (v40 + ("hello" + "\n\t")), F2: 93}
            _ = v40
            _ = v41
        }
        v16 = (v16 >> ((2) & 63))
        _ = fi39
    }
    switch 0 {
    case -9223372036854775808:
        for fi42 := 0; fi42 < 20; fi42++ {
            v16 -= genFn0()
            _ = fi42
        }
    case -2147483648:
        if ((-2.25 + 1.5) != (1.5 + 1.5)) {
            fmt.Println(v16)
            if true {
                v16 = (v16 % (((v16 & v16)) | 1))
                if (false || true) {
                    v16 -= v16
                } else {
                    delete(v2, "a")
                }
                var v43 float64 = float64(0.1)
                _ = v43
            } else {
                var v44 uint16 = 1
                var v45 int8 = 100
                _ = v44
                _ = v45
            }
            if (1000 != 1000) {
                _ = len(v4)
                var v46 uint64 = 31
                _ = v46
            } else {
                var v47 string = "hello"
                switch 1 {
                case -395:
                    _ = len(v4)
                default:
                    v48 := GenS4{F0: 97, F1: "", F2: 127}
                    _ = v48
                }
                var v49 uint16 = 326
                _ = v47
                _ = v49
            }
        }
    default:
        if len(v1) > 0 {
            v1[0] = ("foo bar" + ((("hello" + ("" + "foo bar")) + "foo bar") + "hello"))
        }
    }
    var v50 bool = false
    v16 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v50
}
