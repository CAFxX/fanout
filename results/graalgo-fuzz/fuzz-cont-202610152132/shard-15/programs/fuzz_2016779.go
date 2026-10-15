package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint) string {
    if ((false && (-1 <= 2147483647)) || false) {
        var v7 int16 = -337
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 float64 = 1.5
        v10 := []int{(^(fi8)), 4294967295}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint8 = 1
    _ = p0
    _ = v11
    return "hello"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"\n\t", "\n\t", "hello", ("foo bar" + "a")}
    v3 := []int{}
    v4 := make(map[string]bool)
    var v5 int = -720
    v6 := make(map[string]uint32)
    v5 = (^(v5))
    switch ("a" + "hello") {
    case "\n\t":
        if (-115 <= -856) {
            fmt.Println(v5)
        }
    }
    v12 := GenS1{F0: 1, F1: 65535, F2: 678}
    v13 := GenS2{F0: -1}
    v5 = (-(v5))
    var v14 uint8 = 80
    var v15 float32 = 3.14159
    v14 = v14
    var v16 string = ("" + "foo bar")
    if true {
        for fi17 := 0; fi17 < 1; fi17++ {
            for fi18 := 0; fi18 < 11; fi18++ {
                fmt.Println(fi18, fi18)
                if false {
                    if true {
                        v14 += (v14 ^ 2)
                        if (fi18 <= fi17) {
                            if (523 <= 2) {
                                v19 := GenS3{F0: -26, F1: 32767}
                                fmt.Println(v14)
                                var v20 int64 = 585
                                _ = v19
                                _ = v20
                            } else {
                                v14 -= (v14 - v14)
                            }
                            if false {
                                fmt.Println(v14)
                                fmt.Println(fi18, fi18)
                                var v21 uint8 = v14
                                v14 += (v21 / ((v14) | 1))
                                delete(v4, v16)
                                _ = v21
                            } else {
                                var v22 int = fi17
                                v23 := GenS1{F0: -1, F1: 711, F2: 65535}
                                var v24 uint16 = 2
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            var v25 uint8 = v14
                            if (0.0 <= v15) {
                                if true {
                                    if (42 >= -603) {
                                        _ = len(v1)
                                    }
                                    if true {
                                        v16 = (((("hello" + "") + v16) + "") + "a")
                                    }
                                    var v26 int = v5
                                    switch (fi18 & v26) {
                                    case 127:
                                        if (-9223372036854775808 >= 9223372036854775807) {
                                            switch fi17 {
                                            case -28:
                                                if len(v3) > 0 {
                                                    v3[0] = (-(fi18))
                                                }
                                            default:
                                                v16 = v16
                                            }
                                        }
                                    case -2147483648:
                                        fmt.Println(v25, fi18)
                                    }
                                    _ = v26
                                }
                            }
                            v27 := GenS1{F0: -176, F1: 458, F2: 65535}
                            _ = v25
                            _ = v27
                        }
                    }
                    v6[v16] = 2
                } else {
                    v28 := GenS3{F0: -1, F1: 0}
                    var v29 string = ("a" + ("foo bar" + ((((((("a" + (v16 + v16)) + "foo bar") + "\n\t") + v16) + (v16 + "a")) + v16) + ((v16 + "a") + (v16 + (v16 + ""))))))
                    _ = v28
                    _ = v29
                }
                if (v14 <= 0) {
                    v5 = 81
                    v15 = (v15 + v15)
                    fmt.Println(v15, fi18, v16)
                    v16 = v16
                    v15 = (-2.25 * 1.5)
                }
                _ = fi18
            }
            v30 := GenS2{F0: -2147483648}
            _ = fi17
            _ = v30
        }
    }
    if (true && (42 > 42)) {
        var v31 int16 = 1000
        v14 = (v14 - v14)
        switch (v16 + "a") {
        case "\n\t":
            v15 += -2.25
        }
        _ = v31
    } else {
        if ((v14 <= v14) || (!false)) {
            v32 := v6[v16]
            for fi33 := 0; fi33 < 7; fi33++ {
                fmt.Println(v5)
                for fi34 := 0; fi34 < 5; fi34++ {
                    v35 := GenS3{F0: 0, F1: 0}
                    _ = fi34
                    _ = v35
                }
                _ = fi33
            }
            for fi36 := 0; fi36 < 20; fi36++ {
                var v37 int16 = -1
                _ = fi36
                _ = v37
            }
            _ = v32
        }
    }
    if (731 >= 705) {
        v38 := v4["a"]
        var v39 int32 = 0
        _ = v38
        _ = v39
    }
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
