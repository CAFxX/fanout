package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int32; F2 int }
func genFn0(p0 bool, p1 uint32) uint16 {
    v5 := []int{1}
    v6 := GenS1{F0: ((-1 <= 32767) && false)}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 string, p1 uint32, p2 int16) uint64 {
    p1 = p1
    if true {
        for fi7 := 0; fi7 < 12; fi7++ {
            p1 = p1
            fmt.Println(p2, p0)
            p0 = (p0 + ((p0 + (p0 + p0)) + ""))
            _ = fi7
        }
    } else {
        if true {
            if true {
                var v8 int = 4294967295
                _ = v8
            }
            p1 -= 2147483647
            if false {
                fmt.Println(p0, p2)
            } else {
                switch p0 {
                case "hello":
                    var v9 float64 = (1.5 + 0.0)
                    _ = v9
                default:
                    fmt.Println(p2, p0, p1)
                }
            }
        } else {
            fmt.Println(p1, p2)
            p2 = p2
        }
        var v10 float64 = (3.14159 * float64((0.1 - 1e10)))
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            p1 = 2147483647
            fmt.Println(p2)
            _ = fi12
        }
        p0 = ("" + p0)
        _ = fi11
    }
    var v13 int = -837
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 767
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = (658 <= 1000)
    var v3 bool = v2
    var v4 float64 = float64((0.0 + 3.14159))
    fmt.Println(v3, v4)
    v14 := GenS1{F0: (!v3)}
    v15 := GenS2{F0: 661, F1: 1}
    v3 = true
    switch 1 {
    case 32767:
        var v16 int8 = 42
        _ = v16
    default:
        switch -136 {
        case -1:
            v3 = (-9223372036854775808 != -161)
        default:
            if ((1 >= 1) && (255 > 102)) {
                v17 := GenS3{F0: 1, F1: 46, F2: 537}
                for fi18 := 0; fi18 < 16; fi18++ {
                    var v19 int32 = 42
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 uint8 = 25
                    v22 := []int{153, fi20, (fi20 % ((-339) | 1))}
                    for fi23 := 0; fi23 < 9; fi23++ {
                        switch (fi23 / (((-(fi23))) | 1)) {
                        case 32768:
                            v2 = (v3 && v2)
                        case 42:
                            var v24 string = ("\n\t" + (("\n\t" + ("foo bar" + "\n\t")) + ("" + "a")))
                            _ = v24
                        case -616:
                            fmt.Println(v2)
                        }
                        v25 := GenS3{F0: (v21 % (((v21 ^ (v21 >> ((18446744073709551615) & 63)))) | 1)), F1: -128, F2: 42}
                        var v26 float32 = (0.1 - 0.0)
                        _ = fi23
                        _ = v25
                        _ = v26
                    }
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = v17
            }
        }
    }
    var v27 uint32 = 367
    v28 := GenS2{F0: 1, F1: v27}
    v29 := GenS1{F0: (true || v3)}
    var v30 int = -821
    if v2 {
        if v2 {
            v31 := []int{}
            if v2 {
                if (v4 >= (v4 * 1.5)) {
                    if (!(v2 && (476 > 562))) {
                        for fi32 := 0; fi32 < 14; fi32++ {
                            for fi33 := 0; fi33 < 2; fi33++ {
                                fmt.Println(v3, v3, v27)
                                var v34 uint64 = genFn1("hello", v27, 83)
                                _ = fi33
                                _ = v34
                            }
                            _ = fi32
                        }
                        v35 := GenS2{F0: 0, F1: v27}
                        _ = v35
                    }
                }
                if (!(!true)) {
                    for fi36 := 0; fi36 < 20; fi36++ {
                        var v37 string = ""
                        v2 = v3
                        for fi38 := 0; fi38 < 20; fi38++ {
                            v39 := GenS2{F0: 42, F1: v27}
                            v30 = (-(fi38))
                            switch ("hello" + (v37 + v37)) {
                            case "a":
                                var v40 int16 = -32768
                                _ = v40
                            case "foo bar":
                                var v41 int = (v30 ^ 177)
                                _ = v41
                            case "\n\t":
                                v30 = 128
                            default:
                                var v42 int16 = 0
                                _ = v42
                            }
                            _ = fi38
                            _ = v39
                        }
                        _ = fi36
                        _ = v37
                    }
                } else {
                    if v2 {
                        fmt.Println(v30, v27, v2)
                    }
                    var v43 bool = ((v27 % (((v27 ^ (v27 ^ 698))) | 1)) >= v27)
                    _ = v43
                }
                var v44 int32 = -2147483648
                _ = v44
            }
            _ = v31
        }
    } else {
        v45 := GenS4{F0: (^(v27)), F1: 1, F2: v30}
        fmt.Println(v30, v30)
        if v2 {
            v27 -= v27
        } else {
            var v46 int8 = 0
            v2 = (65535 != 2)
            for fi47 := 0; fi47 < 12; fi47++ {
                if true {
                    switch (fi47 + 42) {
                    case 2147483647:
                        _ = len(v1)
                    default:
                        v48 := []int{(fi47 * (v30 | fi47)), (fi47 | v30)}
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 10; fi49++ {
                        v27 = (^(v27))
                        _ = fi49
                    }
                }
                var v50 uint32 = 0
                _ = fi47
                _ = v50
            }
            _ = v46
        }
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
