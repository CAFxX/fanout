package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint64; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string }

func main() {
    var v1 uint16 = 312
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    var v4 uint64 = 2
    v5 := make(map[string]uint32)
    var v6 int16 = 247
    var v7 uint64 = (^(v4))
    fmt.Println(v6, v1)
    if false {
        v6 -= v6
        switch "foo bar" {
        case "hello":
            switch 127 {
            case 0:
                v3["a"] = false
            case 230:
                var v8 uint64 = (v4 << ((475) & 63))
                _ = v8
            }
        }
    } else {
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            fmt.Println(v7, v6, v1)
        default:
            var v9 uint64 = v7
            _ = v9
        }
        v10 := []int{int(v1)}
        v6 -= 1000
        _ = v10
    }
    v1 -= 43
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(fi12, fi11, v1)
            v6 -= (^(v6))
            v6 = v6
            v13 := GenS1{F0: ("\n\t" + "a")}
            if (255 <= 42) {
                if (253 <= 1) {
                    fmt.Println(v1)
                    v7 -= v7
                    var v14 uint8 = 0
                    _ = v14
                }
                var v15 bool = (235 != v4)
                fmt.Println(v7, v15)
                var v16 uint32 = 2147483648
                _ = v15
                _ = v16
            }
            v3[""] = false
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v6 += v6
    switch 128 {
    case -1:
        if true {
            for fi17 := 0; fi17 < 7; fi17++ {
                v1 += 62
                if (int(v7) != (^(fi17))) {
                    v1 += (v1 * (v1 % ((v1) | 1)))
                    if true {
                        v3[(((("a" + "foo bar") + "hello") + (("hello" + ((("hello" + "a") + "\n\t") + "foo bar")) + "")) + ("hello" + "\n\t"))] = true
                        v18 := GenS1{F0: "hello"}
                        _ = v18
                    } else {
                        v4 = v7
                    }
                } else {
                    var v19 uint8 = 42
                    v1 = v1
                    if (434 >= v6) {
                        if true {
                            if (-2.25 <= 1e10) {
                                for fi20 := 0; fi20 < 13; fi20++ {
                                    var v21 uint16 = 1
                                    var v22 uint16 = (v1 & (^(v1)))
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                                if (65535 >= 18446744073709551615) {
                                    v1 += v1
                                    for fi23 := 0; fi23 < 17; fi23++ {
                                        v6 += (v6 & v6)
                                        v4 -= (^(v4))
                                        _ = fi23
                                    }
                                }
                            } else {
                                var v24 bool = (true || (!(!false)))
                                var v25 uint32 = 4294967295
                                v26 := []int{fi17}
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                            v27 := GenS2{F0: 0, F1: 42}
                            _ = v27
                        }
                        fmt.Println(fi17, v1, v6)
                        v28 := GenS3{F0: 1, F1: (^(v4)), F2: true}
                        _ = v28
                    }
                    _ = v19
                }
                if true {
                    for fi29 := 0; fi29 < 3; fi29++ {
                        var v30 float64 = 3.14159
                        _ = fi29
                        _ = v30
                    }
                    if (true && (674 != 2147483647)) {
                        for fi31 := 0; fi31 < 12; fi31++ {
                            v32 := v2[(("\n\t" + ("\n\t" + (("\n\t" + "") + "hello"))) + ("a" + "\n\t"))]
                            _ = fi31
                            _ = v32
                        }
                    } else {
                        if (true || true) {
                            switch "\n\t" {
                            case "hello":
                                v33 := GenS1{F0: ((("foo bar" + ("hello" + "a")) + "\n\t") + "foo bar")}
                                _ = v33
                            }
                        }
                        v7 = (v7 / ((18446744073709551615) | 1))
                    }
                } else {
                    v6 = (v6 | v6)
                    v7 -= v7
                }
                _ = fi17
            }
            v7 += (v7 * (v4 & v4))
        }
    }
    if false {
        v34 := GenS3{F0: int32(v7), F1: v4, F2: (-314 < 336)}
        if true {
            switch "a" {
            case "":
                var v35 uint32 = 2147483647
                _ = v35
            }
            v36 := GenS4{F0: uint(v7)}
            _ = v36
        }
        _ = v34
    }
    v37 := GenS5{F0: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
