package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint32; F2 int32 }
type GenS4 struct { F0 bool }
func genFn0(p0 int, p1 int32) uint64 {
    for fi5 := 0; fi5 < 3; fi5++ {
        p1 += (-(p1))
        _ = fi5
    }
    p1 -= (p1 * p1)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 601
}
func genFn1(p0 uint64, p1 string) string {
    p0 -= (p0 ^ 2)
    if (false || ((1 > 127) && (554 != uint16(p0)))) {
        switch 32767 {
        case 32768:
            var v6 uint16 = 0
            _ = v6
        case 4294967295:
            p1 = "foo bar"
        case 42:
            var v7 uint64 = 42
            _ = v7
        }
        fmt.Println(p0, p0)
    } else {
        v8 := []int{-26, 128}
        v9 := GenS1{F0: -2147483648}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return (p1 + p1)
}

func main() {
    var v1 uint16 = 1
    var v2 float32 = 3.14159
    v3 := []string{}
    var v4 string = "a"
    v10 := GenS2{F0: 224}
    v1 -= v1
    fmt.Println(v4, v1)
    if true {
        v4 = ("\n\t" + v4)
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        if (true && false) {
            v12 := GenS3{F0: 42, F1: 268, F2: 2147483647}
            if (!false) {
                if (-2147483648 > 1) {
                    var v13 uint = 814
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v3 = append(v3, ((v4 + v4) + v4))
                        _ = fi14
                    }
                    _ = v13
                }
                fmt.Println(v4, v4, fi11)
            }
            _ = v12
        }
        if true {
            if (!(!true)) {
                var v15 int8 = -10
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 bool = false
                    var v18 bool = false
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                if false {
                    if false {
                        switch (fi11 << ((42) & 63)) {
                        case 128:
                            if (!(-1 < 0)) {
                                v3 = append(v3, (v4 + v4))
                            }
                        case 42:
                            v2 = v2
                        }
                    }
                    var v19 int16 = 424
                    _ = v19
                }
                _ = v15
            } else {
                v20 := GenS4{F0: false}
                for fi21 := 0; fi21 < 2; fi21++ {
                    if (1 < 0) {
                        if false {
                            v4 = v4
                            fmt.Println(fi21, v2, v4)
                            v1 -= v1
                            if (18446744073709551615 < 930) {
                                v2 = v2
                                if (367 > v1) {
                                    _ = len(v3)
                                    v22 := GenS2{F0: 255}
                                    _ = v22
                                }
                            } else {
                                var v23 int64 = -1
                                _ = v23
                            }
                        } else {
                            v2 = v2
                            if (!(65535 < 255)) {
                                v24 := GenS1{F0: -1}
                                _ = v24
                            }
                            var v25 uint32 = 42
                            _ = v25
                        }
                        switch v4 {
                        case "":
                            var v26 uint = 42
                            _ = v26
                        case "hello":
                            v3 = append(v3, "hello")
                        case "\n\t":
                            v1 = (v1 >> ((65535) & 63))
                        }
                        v27 := GenS1{F0: -1}
                        _ = v27
                    }
                    _ = fi21
                }
                v4 = "hello"
                _ = v20
            }
            fmt.Println(v2, v4, v2)
        } else {
            fmt.Println(v4, v4)
        }
        _ = fi11
    }
    if true {
        v28 := GenS1{F0: 524}
        _ = v28
    }
    fmt.Println(v2, v4)
    for fi29 := 0; fi29 < 7; fi29++ {
        if len(v3) > 0 {
            v3[0] = ("a" + (v4 + ("\n\t" + "")))
        }
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 bool = (!true)
            var v32 uint16 = (v1 ^ (v1 / ((v1) | 1)))
            _ = fi30
            _ = v31
            _ = v32
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
