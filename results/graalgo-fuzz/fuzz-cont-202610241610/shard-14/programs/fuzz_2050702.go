package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int16; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int64 }
func genFn0(p0 uint16, p1 int8, p2 bool) uint16 {
    v4 := []int{42, 9223372036854775807, -1, 128}
    if (!p2) {
        p1 = (p1 >> ((0) & 63))
        var v5 int32 = 42
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 801
}

func main() {
    var v1 uint32 = 2
    v2 := []string{"hello", ("foo bar" + "\n\t"), ("" + "foo bar"), "\n\t"}
    v3 := []int{}
    v1 = v1
    v6 := GenS1{F0: true}
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(v1)
        if (42 == 42) {
            if (!(1 <= 0)) {
                for fi8 := 0; fi8 < 14; fi8++ {
                    v9 := GenS2{F0: false}
                    _ = fi8
                    _ = v9
                }
                var v10 string = "a"
                if false {
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v10 = (v10 + ("a" + ""))
                        _ = fi11
                    }
                    for fi12 := 0; fi12 < 19; fi12++ {
                        var v13 uint8 = 128
                        _ = fi12
                        _ = v13
                    }
                    fmt.Println(v1, v1, fi7)
                } else {
                    v10 = "\n\t"
                    var v14 bool = true
                    _ = v14
                }
                _ = v10
            }
            var v15 string = "a"
            if false {
                if true {
                    var v16 int8 = -80
                    _ = v16
                }
            } else {
                if (-1 < 42) {
                    fmt.Println(v1, fi7, v1)
                    for fi17 := 0; fi17 < 8; fi17++ {
                        v18 := GenS3{F0: 479, F1: 2147483647}
                        _ = fi17
                        _ = v18
                    }
                    var v19 float32 = (1e10 * 0.1)
                    _ = v19
                } else {
                    if (true || (0.1 < (0.0 / float64(float32(1e10))))) {
                        fmt.Println(v15, v15, fi7)
                        fmt.Println(v15, v15)
                    } else {
                        if true {
                            v15 = (v15 + v15)
                        }
                        if true {
                            var v20 int32 = -1
                            _ = v20
                        }
                        switch (^(fi7)) {
                        case 80:
                            if true {
                                for fi21 := 0; fi21 < 12; fi21++ {
                                    var v22 bool = true
                                    _ = fi21
                                    _ = v22
                                }
                                var v23 int16 = 32767
                                _ = v23
                            } else {
                                fmt.Println(v15, v1, v1)
                            }
                        case -924:
                            var v24 int64 = 1
                            _ = v24
                        default:
                            switch (^(fi7)) {
                            case 127:
                                var v25 uint = 962
                                _ = v25
                            default:
                                var v26 int8 = 127
                                _ = v26
                            }
                        }
                    }
                }
                switch v15 {
                case "":
                    for fi27 := 0; fi27 < 10; fi27++ {
                        var v28 int16 = -32768
                        fmt.Println(fi27)
                        _ = fi27
                        _ = v28
                    }
                default:
                    v15 = ""
                }
                v15 = v15
            }
            _ = v15
        } else {
            switch "\n\t" {
            case "a":
                for fi29 := 0; fi29 < 10; fi29++ {
                    var v30 bool = (fi7 >= (fi7 ^ (fi29 << ((420) & 63))))
                    v30 = true
                    var v31 string = "foo bar"
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            case "foo bar":
                v32 := GenS2{F0: (0 <= 2147483647)}
                _ = v32
            case "hello":
                _ = len(v2)
            default:
                if len(v3) > 0 {
                    v3[0] = -817
                }
            }
        }
        _ = fi7
    }
    v33 := GenS3{F0: 0, F1: 2147483647}
    var v34 int16 = 1
    var v35 uint64 = 2
    fmt.Println(v34, v35)
    v34 = v34
    for fi36 := 0; fi36 < 2; fi36++ {
        v35 += (v35 + v35)
        v1 = v1
        _ = fi36
    }
    v37 := GenS1{F0: (-2.25 >= 1.5)}
    for fi38 := 0; fi38 < 18; fi38++ {
        v39 := GenS3{F0: v35, F1: -632}
        var v40 uint8 = 42
        v41 := GenS4{F0: (-(v34)), F1: (v34 & v34), F2: v40}
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    fmt.Println(v1, v34, v35)
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    if (199 != 2) {
        v42 := GenS5{F0: v1, F1: 398}
        var v43 uint64 = v35
        _ = v42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v37
    _ = v6
}
