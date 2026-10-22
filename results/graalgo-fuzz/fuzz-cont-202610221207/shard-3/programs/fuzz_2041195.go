package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 string; F1 uint32 }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[uint32]string)
    var v3 uint64 = 4294967295
    v4 := make(map[string]bool)
    var v5 uint32 = 0
    var v6 float64 = 3.14159
    if ("\n\t" != ("a" + "a")) {
        v7 := []int{32767, -145, 4294967295}
        v5 -= 0
        if (!true) {
            fmt.Println(v3, v3)
            v1 += 0
            var v8 string = "foo bar"
            _ = v8
        } else {
            v9 := GenS1{F0: 32767}
            v5 += v5
            _ = v9
        }
        _ = v7
    } else {
        if true {
            v5 = (^(v5))
        } else {
            if true {
                v6 = float64((0.0 - float32((0.0 * 1e10))))
            } else {
                v3 -= 4294967295
            }
            if (1 != -1) {
                var v10 int32 = 0
                var v11 bool = (!true)
                for fi12 := 0; fi12 < 7; fi12++ {
                    v2[v5] = "foo bar"
                    switch int(v3) {
                    case -753:
                        if (!true) {
                            v13 := GenS2{F0: 42, F1: 1}
                            v14 := GenS2{F0: 128, F1: 1}
                            v15 := GenS1{F0: uint16(v5)}
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                    case 984:
                        if ((-20 > 0) || true) {
                            var v16 uint16 = 607
                            _ = v16
                        } else {
                            v11 = false
                        }
                    }
                    _ = fi12
                }
                _ = v10
                _ = v11
            } else {
                var v17 uint32 = v5
                _ = v17
            }
            var v18 int32 = -2147483648
            _ = v18
        }
    }
    _ = len(v4)
    fmt.Println(v1, v5)
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := GenS2{F0: 2, F1: 42}
        var v21 float32 = 1e10
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := []int{1}
        _ = fi22
        _ = v23
    }
    v1 += v1
    v24 := []int{4294967295}
    v6 = float64(0.0)
    v3 += (v3 + (v3 * 1))
    var v25 float32 = -2.25
    for fi26 := 0; fi26 < 16; fi26++ {
        v27 := v4[("hello" + ("\n\t" + (("\n\t" + "a") + "\n\t")))]
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        v29 := GenS1{F0: 0}
        if (748 != 42) {
            if false {
                v30 := GenS3{F0: 252, F1: -9223372036854775808}
                v5 = v1
                switch fi28 {
                case 4294967295:
                    switch (fi28 * fi28) {
                    case 9223372036854775807:
                        var v31 uint32 = (^(v1))
                        _ = v31
                    case 42:
                        v5 = v5
                    case -2147483648:
                        var v32 int16 = 42
                        _ = v32
                    default:
                        if (!((1 < 839) || true)) {
                            v25 += (v25 + v25)
                            var v33 int8 = 0
                            _ = v33
                        }
                    }
                case -1:
                    var v34 float32 = (v25 - v25)
                    _ = v34
                case 1:
                    if true {
                        if (32768 < -2147483648) {
                            if true {
                                var v35 string = "a"
                                _ = v35
                            }
                            var v36 int8 = 100
                            _ = v36
                        }
                        var v37 bool = true
                        v37 = (v37 || (!v37))
                        _ = v37
                    }
                default:
                    if true {
                        fmt.Println(fi28, v3, v25)
                    }
                }
                _ = v30
            }
        } else {
            v38 := v4[(("\n\t" + "\n\t") + "hello")]
            v39 := GenS4{F0: "hello", F1: v5}
            _ = v38
            _ = v39
        }
        var v40 bool = (((fi28 / ((fi28) | 1)) > (fi28 * 4294967295)) && (!((("" + "hello") + "a") != "\n\t")))
        _ = fi28
        _ = v29
        _ = v40
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        switch "hello" {
        case "\n\t":
            for fi42 := 0; fi42 < 11; fi42++ {
                fmt.Println(fi41)
                _ = fi42
            }
        case "hello":
            v43 := GenS3{F0: 499, F1: 179}
            _ = v43
        }
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
