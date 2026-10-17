package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 string; F1 string }
func genFn0() uint8 {
    var v4 uint8 = 1
    for fi5 := 0; fi5 < 5; fi5++ {
        if true {
            v6 := GenS1{F0: fi5, F1: false, F2: v4}
            _ = v6
        } else {
            if (float64(3.14159) < (1.5 * (0.0 / -2.25))) {
                v4 -= 2
                v4 = 0
                fmt.Println(fi5, v4)
            }
        }
        var v7 uint32 = 1
        _ = fi5
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{1, fi8, fi8, (fi8 ^ (fi8 << ((2) & 63)))}
        for fi10 := 0; fi10 < 10; fi10++ {
            v4 = (v4 | (v4 * (v4 >> ((620) & 63))))
            var v11 int32 = 847
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v4 -= v4
    _ = v4
    return 0
}

func main() {
    v1 := []string{(("a" + "hello") + ""), ("foo bar" + "hello"), ("" + "\n\t"), ("" + "a")}
    v2 := []string{}
    var v3 int = -392
    v3 = -1
    fmt.Println(v3)
    var v12 int32 = 0
    var v13 int = (v3 >> ((60) & 63))
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (false && false) {
                v12 = -55
                var v16 uint64 = 18446744073709551615
                _ = v16
            }
            _ = fi15
        }
        v3 += v13
        _ = fi14
    }
    if false {
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 float64 = -2.25
                var v20 float32 = (3.14159 / (0.0 * float32((v19 + v19))))
                _ = fi18
                _ = v19
                _ = v20
            }
            var v21 bool = (66 > 42)
            var v22 int8 = 42
            _ = fi17
            _ = v21
            _ = v22
        }
        var v23 uint16 = 179
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v23
    } else {
        if len(v1) > 0 {
            v1[0] = ""
        }
        fmt.Println(v13, v12, v13)
    }
    fmt.Println(v13, v12)
    if true {
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 int = -2147483648
            var v26 uint16 = 277
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 19; fi27++ {
            v28 := GenS2{F0: 128}
            fmt.Println(v13, v12)
            if ((-1000 < -722) || (!false)) {
                if (263 >= 32767) {
                    if true {
                        v29 := GenS3{F0: (!false)}
                        var v30 uint = 95
                        v31 := GenS4{F0: 32767}
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        switch v3 {
                        case 42:
                            if true {
                                for fi32 := 0; fi32 < 20; fi32++ {
                                    v12 = (v12 | v12)
                                    v33 := GenS5{F0: 271, F1: (("hello" + "a") + "\n\t")}
                                    if (!(false || false)) {
                                        if false {
                                            if ((79 < 2) || false) {
                                                fmt.Println(fi27, v3, fi27)
                                                fmt.Println(v3)
                                            } else {
                                                var v34 bool = true
                                                if len(v2) > 0 {
                                                    v2[0] = ("" + ("\n\t" + (("hello" + "") + ("hello" + "hello"))))
                                                }
                                                if v34 {
                                                    fmt.Println(v34, v13, v12)
                                                }
                                                _ = v34
                                            }
                                            fmt.Println(v12)
                                            fmt.Println(fi32)
                                        }
                                        v35 := GenS6{F0: 255, F1: 10}
                                        _ = v35
                                    } else {
                                        v3 += 4294967295
                                        v3 -= (fi27 ^ (-(fi32)))
                                        var v36 bool = true
                                        _ = v36
                                    }
                                    v37 := GenS1{F0: 286, F1: (fi27 <= (-(v13))), F2: 87}
                                    _ = fi32
                                    _ = v33
                                    _ = v37
                                }
                                var v38 int16 = 901
                                v38 -= -32768
                                _ = v38
                            }
                        case -9223372036854775808:
                            v1 = append(v1, "\n\t")
                        case 806:
                            var v39 int = 2147483647
                            _ = v39
                        }
                        fmt.Println(fi27, v3, fi27)
                    }
                } else {
                    v40 := GenS7{F0: ((!true) || ((!(1 > 113)) && false))}
                    _ = v40
                }
            } else {
                switch (-(fi27)) {
                case 654:
                    v41 := GenS8{F0: 100}
                    _ = v41
                }
                var v42 float64 = 0.1
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v42
            }
            _ = fi27
            _ = v28
        }
        v43 := GenS9{F0: ("hello" + "hello"), F1: ("a" + "\n\t")}
        _ = v43
    } else {
        for fi44 := 0; fi44 < 16; fi44++ {
            v3 += fi44
            var v45 bool = true
            _ = fi44
            _ = v45
        }
    }
    v1 = append(v1, "hello")
    var v46 int64 = -9223372036854775808
    fmt.Println(v3, v12, v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v46
}
