package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int; F1 uint16; F2 bool }
type GenS5 struct { F0 int16; F1 uint8 }
func genFn0() int64 {
    v6 := []int{0, -9223372036854775808}
    if len(v6) > 0 {
        v6[0] = 42
    }
    var v7 float32 = (float32(1.5) / (1e10 + ((0.1 * 0.0) + (((((0.1 * 1e10) - ((0.0 - 0.0) + 0.1)) * 0.1) - 3.14159) - 3.14159))))
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: 2}
        _ = fi8
        _ = v9
    }
    v7 += float32(-2.25)
    _ = v6
    _ = v7
    return 588
}

func main() {
    var v1 int8 = -71
    v2 := []string{("a" + ""), "", "\n\t", ("foo bar" + "hello")}
    v3 := []int{1, 4294967295}
    var v4 int8 = 0
    var v5 uint64 = 18446744073709551615
    v3 = append(v3, -646)
    switch (("" + "hello") + "foo bar") {
    case "":
        if (65535 <= 833) {
            v1 -= (v4 - 127)
            for fi10 := 0; fi10 < 2; fi10++ {
                v5 = (v5 % (((v5 % ((4294967295) | 1))) | 1))
                if true {
                    if (v1 <= int8(v5)) {
                        if true {
                            var v11 string = "hello"
                            var v12 int = 357
                            if false {
                                v4 = 80
                                v11 = "foo bar"
                            } else {
                                if (!(2147483647 <= -570)) {
                                    var v13 uint16 = 1
                                    if len(v2) > 0 {
                                        v2[0] = v11
                                    }
                                    var v14 int8 = (v4 + 42)
                                    _ = v13
                                    _ = v14
                                } else {
                                    v12 -= fi10
                                    v15 := GenS1{F0: 0}
                                    v12 -= fi10
                                    _ = v15
                                }
                            }
                            _ = v11
                            _ = v12
                        }
                    } else {
                        var v16 uint16 = 32767
                        _ = v16
                    }
                    if (905 <= 1) {
                        for fi17 := 0; fi17 < 5; fi17++ {
                            fmt.Println(v5, fi17)
                            v18 := GenS2{F0: 4294967295, F1: int8(fi10), F2: 275}
                            _ = fi17
                            _ = v18
                        }
                    }
                } else {
                    var v19 uint32 = 65535
                    v4 -= v4
                    _ = len(v2)
                    _ = v19
                }
                if true {
                    fmt.Println(v5)
                    v20 := GenS3{F0: 42, F1: (fi10 | (fi10 * fi10))}
                    _ = v20
                }
                _ = fi10
            }
        } else {
            for fi21 := 0; fi21 < 7; fi21++ {
                v22 := GenS4{F0: fi21, F1: 65535, F2: (255 == uint16(v5))}
                var v23 int32 = 0
                if len(v3) > 0 {
                    v3[0] = fi21
                }
                _ = fi21
                _ = v22
                _ = v23
            }
        }
    case "foo bar":
        fmt.Println(v4, v1, v4)
    }
    if (!true) {
        v24 := GenS1{F0: 42}
        for fi25 := 0; fi25 < 5; fi25++ {
            var v26 uint64 = (v5 % (((v5 % ((v5) | 1))) | 1))
            v26 += (v5 / ((v5) | 1))
            for fi27 := 0; fi27 < 9; fi27++ {
                _ = len(v2)
                if true {
                    var v28 uint64 = v26
                    if false {
                        v26 = 1
                        var v29 int16 = -675
                        fmt.Println(v26, fi27, v29)
                        _ = v29
                    } else {
                        var v30 uint8 = 255
                        var v31 string = "\n\t"
                        if true {
                            fmt.Println(fi27, v5)
                            v31 = (v31 + ((v31 + v31) + v31))
                        }
                        _ = v30
                        _ = v31
                    }
                    v5 = 159
                    _ = v28
                } else {
                    var v32 int = (fi27 & fi27)
                    v5 = (v26 - 4294967295)
                    _ = v32
                }
                _ = fi27
            }
            _ = fi25
            _ = v26
        }
        _ = v24
    } else {
        fmt.Println(v1, v1, v5)
    }
    _ = len(v2)
    v5 += v5
    switch ("\n\t" + "") {
    case "\n\t":
        v1 = (v1 / ((v1) | 1))
    case "a":
        var v33 uint16 = 678
        _ = v33
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = 1
        }
    }
    if false {
        v34 := GenS2{F0: 120, F1: (v1 % ((42) | 1)), F2: 65535}
        _ = v34
    } else {
        v35 := GenS5{F0: 42, F1: 165}
        var v36 uint64 = v5
        _ = v35
        _ = v36
    }
    if false {
        v5 += 796
        for fi37 := 0; fi37 < 17; fi37++ {
            switch "hello" {
            case "hello":
                v38 := GenS5{F0: 32767, F1: 130}
                _ = v38
            case "foo bar":
                var v39 uint16 = 644
                _ = v39
            case "a":
                v4 += 1
            }
            for fi40 := 0; fi40 < 4; fi40++ {
                v4 = v1
                _ = fi40
            }
            _ = fi37
        }
    }
    v41 := GenS1{F0: 166}
    v4 = (v1 ^ v4)
    var v42 int8 = (v4 + (v1 / ((-1) | 1)))
    if len(v2) > 0 {
        v2[0] = ""
    }
    var v43 int16 = -933
    v4 = (v1 + 0)
    _ = len(v2)
    v1 = (v1 >> ((674) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
