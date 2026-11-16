package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint32 }
func genFn0(p0 int16) bool {
    var v6 string = "\n\t"
    fmt.Println(v6, v6)
    v6 = v6
    _ = p0
    _ = v6
    return (v6 < v6)
}

func main() {
    var v1 string = ("a" + (("foo bar" + ("\n\t" + "a")) + ((("hello" + "foo bar") + "foo bar") + "")))
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 uint = 18446744073709551615
    var v5 bool = (false && (!false))
    v7 := GenS1{F0: 42, F1: 4294967295, F2: 766}
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint8 = 128
        _ = fi8
        _ = v9
    }
    switch v1 {
    case "foo bar":
        if (0 != -597) {
            if false {
                fmt.Println(v1)
                if (!true) {
                    if v5 {
                        if true {
                            var v10 string = "hello"
                            _ = v10
                        } else {
                            for fi11 := 0; fi11 < 13; fi11++ {
                                v5 = true
                                switch v1 {
                                case "\n\t":
                                    var v12 uint16 = 255
                                    _ = v12
                                case "":
                                    v13 := GenS1{F0: 100, F1: 32768, F2: fi11}
                                    _ = v13
                                case "hello":
                                    v5 = false
                                }
                                for fi14 := 0; fi14 < 15; fi14++ {
                                    if false {
                                        switch "a" {
                                        case "":
                                            v5 = v5
                                        }
                                    }
                                    var v15 int16 = -1
                                    fmt.Println(fi11, v4, fi11)
                                    _ = fi14
                                    _ = v15
                                }
                                _ = fi11
                            }
                        }
                        for fi16 := 0; fi16 < 14; fi16++ {
                            fmt.Println(v5, fi16)
                            _ = fi16
                        }
                    }
                } else {
                    v17 := GenS1{F0: 42, F1: -2147483648, F2: -9223372036854775808}
                    var v18 float32 = 0.0
                    _ = v17
                    _ = v18
                }
            } else {
                v19 := []int{42, 67, 4294967295}
                for fi20 := 0; fi20 < 20; fi20++ {
                    delete(v2, (^(fi20)))
                    var v21 float64 = 1.5
                    _ = len(v3)
                    _ = fi20
                    _ = v21
                }
                _ = len(v19)
                _ = v19
            }
            v5 = true
            var v22 float64 = 0.0
            _ = v22
        }
    }
    if v5 {
        var v23 int8 = 100
        _ = v23
    } else {
        switch (((v1 + (v1 + "\n\t")) + ("" + v1)) + "hello") {
        case "":
            switch -533 {
            case 32767:
                if v5 {
                    for fi24 := 0; fi24 < 19; fi24++ {
                        for fi25 := 0; fi25 < 9; fi25++ {
                            var v26 uint8 = 128
                            _ = fi25
                            _ = v26
                        }
                        var v27 int8 = 0
                        _ = fi24
                        _ = v27
                    }
                    v28 := v2[2147483647]
                    for fi29 := 0; fi29 < 11; fi29++ {
                        if true {
                            fmt.Println(v5, fi29, v28)
                            var v30 int = (v28 % (((fi29 | (fi29 / ((4294967295) | 1)))) | 1))
                            var v31 float32 = (((float32(-2.25) + -2.25) + (float32(((1.5 / 0.1) * (1.5 - 1e10))) + -2.25)) * (1e10 - float32(0.0)))
                            _ = v30
                            _ = v31
                        }
                        _ = fi29
                    }
                    _ = v28
                }
            case 127:
                v1 = (v1 + "")
            case 124:
                var v32 float64 = 1.5
                _ = v32
            }
        case "foo bar":
            if v5 {
                v33 := GenS2{F0: 0, F1: 125}
                if ((float64(0.0) != -2.25) || v5) {
                    var v34 uint8 = 255
                    _ = v34
                } else {
                    var v35 int64 = -1
                    _ = v35
                }
                fmt.Println(v5, v1)
                _ = v33
            } else {
                var v36 uint = (v4 << ((0) & 63))
                _ = v36
            }
        }
        v37 := GenS3{F0: 0, F1: 2, F2: 0}
        _ = v37
    }
    v38 := []int{}
    var v39 int8 = 55
    var v40 uint16 = 491
    v1 = ((v1 + (("" + v1) + v1)) + v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v7
}
