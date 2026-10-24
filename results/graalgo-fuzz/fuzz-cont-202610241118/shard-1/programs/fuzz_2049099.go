package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32 }
func genFn0() bool {
    var v4 float32 = 0.0
    for fi5 := 0; fi5 < 6; fi5++ {
        if false {
            v6 := []int{}
            v4 = ((v4 - (v4 + v4)) * (float32(1e10) / float32(1.5)))
            var v7 int = fi5
            _ = v6
            _ = v7
        }
        var v8 uint32 = 424
        _ = fi5
        _ = v8
    }
    var v9 bool = true
    _ = v4
    _ = v9
    return (v9 || v9)
}
func genFn1() uint16 {
    v10 := []int{624, -219, 32767, 588}
    var v11 float32 = (0.1 / 1.0)
    switch 42 {
    case -2147483648:
        if false {
            if true {
                v11 += v11
                if (42 > 42) {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                } else {
                    switch ("a" + "hello") {
                    case "\n\t":
                        if (32767 >= 665) {
                            _ = len(v10)
                            var v12 int64 = 42
                            _ = v12
                        } else {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                        }
                    case "foo bar":
                        if (!(!("foo bar" <= "foo bar"))) {
                            v11 = ((1e10 * -2.25) - v11)
                            fmt.Println(v11)
                            v13 := GenS1{F0: 1000, F1: 4294967295, F2: 2147483647}
                            _ = v13
                        } else {
                            v14 := GenS1{F0: -32768, F1: 1, F2: 2}
                            for fi15 := 0; fi15 < 1; fi15++ {
                                switch (fi15 ^ fi15) {
                                case 32768:
                                    fmt.Println(v11, fi15)
                                case 42:
                                    switch (fi15 - (^(fi15))) {
                                    case 9223372036854775807:
                                        var v16 uint64 = 42
                                        _ = v16
                                    }
                                case 9223372036854775807:
                                    _ = len(v10)
                                default:
                                    var v17 uint64 = 4294967295
                                    _ = v17
                                }
                                _ = fi15
                            }
                            var v18 uint8 = 160
                            _ = v14
                            _ = v18
                        }
                    }
                    v19 := GenS1{F0: 32767, F1: 491, F2: 0}
                    _ = v19
                }
            } else {
                var v20 uint = 18446744073709551615
                _ = len(v10)
                var v21 float64 = 0.1
                _ = v20
                _ = v21
            }
            fmt.Println(v11)
        }
    default:
        if ((true || false) && true) {
            switch 4294967295 {
            case 32768:
                v11 = (-2.25 / v11)
            case -9223372036854775808:
                fmt.Println(v11)
            default:
                v22 := GenS2{F0: 1}
                _ = v22
            }
            if true {
                v23 := GenS2{F0: 1}
                _ = v23
            }
            if true {
                var v24 string = ((("foo bar" + "a") + (("hello" + "foo bar") + ("foo bar" + ""))) + "foo bar")
                fmt.Println(v24, v11)
                if len(v10) > 0 {
                    v10[0] = 128
                }
                _ = v24
            }
        }
    }
    _ = v10
    _ = v11
    return 1
}

func main() {
    v1 := []string{"\n\t", (("foo bar" + "hello") + ("" + ("" + ("foo bar" + "hello")))), (("hello" + "foo bar") + ("\n\t" + "foo bar"))}
    v2 := []string{"foo bar", ("a" + "hello"), "foo bar", (("\n\t" + "") + "\n\t")}
    var v3 float64 = (1e10 / (3.14159 * (0.1 - 1e10)))
    var v25 int16 = -379
    if false {
        v25 = 1000
    } else {
        v26 := GenS2{F0: 106}
        _ = v26
    }
    switch 814 {
    case -326:
        fmt.Println(v25, v25)
    case -87:
        v27 := GenS2{F0: 339}
        _ = v27
    case -9223372036854775808:
        v25 = v25
    default:
        if (((1e10 + 0.1) * 1e10) != 0.1) {
            v3 = 3.14159
            var v28 int8 = -1
            _ = len(v2)
            _ = v28
        } else {
            v3 = 0.0
            v1 = append(v1, "a")
            v25 = v25
        }
    }
    fmt.Println(v3)
    var v29 float64 = (0.0 * v3)
    if (86 >= -79) {
        v2 = append(v2, "hello")
        switch ((("\n\t" + ("foo bar" + "hello")) + "foo bar") + "\n\t") {
        case "":
            if ((v25 ^ (v25 / ((32767) | 1))) <= (v25 + v25)) {
                fmt.Println(v25, v25, v25)
                v30 := GenS2{F0: 42}
                var v31 float64 = (v3 * float64(0.0))
                _ = v30
                _ = v31
            } else {
                var v32 uint64 = 0
                _ = v32
            }
        }
        var v33 float64 = 1e10
        _ = v33
    }
    var v34 int64 = 9223372036854775807
    for fi35 := 0; fi35 < 5; fi35++ {
        v34 = v34
        _ = fi35
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        v37 := GenS3{F0: -2147483648}
        _ = fi36
        _ = v37
    }
    var v38 uint64 = 409
    var v39 uint16 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v34
    _ = v38
    _ = v39
}
