package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int; F2 string }
type GenS3 struct { F0 bool; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint8) uint {
    if true {
        fmt.Println(p0)
        v7 := GenS1{F0: -128}
        switch (("foo bar" + "hello") + "a") {
        case "foo bar":
            v8 := []int{-1, -2147483648}
            _ = v8
        case "a":
            var v9 float32 = 1e10
            _ = v9
        case "":
            p0 = (^(p0))
        }
        _ = v7
    } else {
        if (((1e10 - (1.5 * 0.1)) - 0.0) <= ((1.5 * float64(0.0)) + 1.5)) {
            fmt.Println(p0)
            v10 := []int{0}
            for fi11 := 0; fi11 < 16; fi11++ {
                var v12 bool = false
                if len(v10) > 0 {
                    v10[0] = fi11
                }
                _ = fi11
                _ = v12
            }
            _ = v10
        } else {
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 int64 = -1
                _ = fi13
                _ = v14
            }
        }
        if false {
            var v15 uint64 = 0
            if (3 < 42) {
                fmt.Println(v15, v15)
            } else {
                var v16 int = -166
                v17 := []int{32767, 9223372036854775807, (^(v16)), 4294967295}
                _ = v16
                _ = v17
            }
            switch (("foo bar" + (("\n\t" + "") + "hello")) + "hello") {
            case "a":
                fmt.Println(p0, p0)
            default:
                var v18 int16 = 1000
                _ = v18
            }
            _ = v15
        }
        p0 = 253
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        if ((1000 > 0) || (1 > 0)) {
            if true {
                v20 := GenS1{F0: 0}
                if ((1.5 * 0.0) == -2.25) {
                    if (((fi19 % ((fi19) | 1)) < fi19) && true) {
                        if false {
                            var v21 uint16 = 0
                            for fi22 := 0; fi22 < 2; fi22++ {
                                if (false && false) {
                                    v23 := []int{fi19}
                                    var v24 int32 = 0
                                    _ = v23
                                    _ = v24
                                } else {
                                    if true {
                                        var v25 int = -624
                                        _ = v25
                                    }
                                    switch "" {
                                    case "\n\t":
                                        if (0 > (fi19 << ((851) & 63))) {
                                            if (true || (1 >= 305)) {
                                                p0 -= (p0 >> ((65535) & 63))
                                                if ((v21 >> ((954) & 63)) > (v21 + (v21 | 369))) {
                                                    var v26 uint8 = (^(p0))
                                                    fmt.Println(fi22, v26)
                                                    _ = v26
                                                }
                                            }
                                            v21 += v21
                                        } else {
                                            switch "hello" {
                                            case "hello":
                                                v27 := []int{-18, fi22, (fi22 - 2147483647)}
                                                _ = v27
                                            }
                                        }
                                    case "hello":
                                        v28 := GenS2{F0: 138, F1: (fi22 << ((2) & 63)), F2: "a"}
                                        _ = v28
                                    default:
                                        v29 := GenS2{F0: 18446744073709551615, F1: fi19, F2: ("a" + ((("\n\t" + "") + "a") + "a"))}
                                        _ = v29
                                    }
                                    var v30 int16 = 32767
                                    _ = v30
                                }
                                var v31 bool = ((((true || true) && (255 > (v21 | (^(v21))))) || false) && true)
                                fmt.Println(p0, fi19)
                                _ = fi22
                                _ = v31
                            }
                            _ = v21
                        }
                        for fi32 := 0; fi32 < 15; fi32++ {
                            v33 := GenS1{F0: 100}
                            var v34 uint = 18446744073709551615
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        if ((!true) && (-990 == -540)) {
                            v35 := []int{(fi19 ^ (fi19 - fi19)), (-(fi19))}
                            _ = v35
                        }
                    }
                    p0 += p0
                } else {
                    v36 := GenS3{F0: ((p0 / ((p0) | 1)) >= 128), F1: 244, F2: 73}
                    _ = v36
                }
                _ = v20
            } else {
                for fi37 := 0; fi37 < 3; fi37++ {
                    v38 := GenS4{F0: 1, F1: (p0 & p0)}
                    var v39 bool = (!((3.14159 - 1.5) >= ((3.14159 * 3.14159) + -2.25)))
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                for fi40 := 0; fi40 < 9; fi40++ {
                    fmt.Println(fi40, p0)
                    if false {
                        v41 := []int{-787}
                        _ = v41
                    } else {
                        p0 = 1
                    }
                    _ = fi40
                }
                v42 := GenS3{F0: (1 != 1), F1: 18446744073709551615, F2: -30}
                _ = v42
            }
        } else {
            var v43 int8 = -35
            fmt.Println(fi19, p0, v43)
            _ = v43
        }
        if (((!true) || true) && true) {
            v44 := []int{0, (fi19 % (((fi19 & fi19)) | 1)), (fi19 % (((fi19 - (fi19 & -1))) | 1)), fi19}
            var v45 int64 = 0
            _ = v44
            _ = v45
        }
        v46 := []int{(fi19 >> ((18446744073709551615) & 63)), fi19}
        _ = fi19
        _ = v46
    }
    _ = p0
    return 2
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 uint32 = 2
    var v4 uint32 = v3
    var v5 bool = true
    var v6 int16 = -32768
    v6 = v6
    for fi47 := 0; fi47 < 6; fi47++ {
        var v48 int8 = -128
        _ = fi47
        _ = v48
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        v6 = int16(v4)
        var v50 float32 = -2.25
        _ = fi49
        _ = v50
    }
    if v5 {
        var v51 float32 = 1e10
        switch "a" {
        case "\n\t":
            var v52 int32 = 2147483647
            _ = v52
        case "a":
            v53 := GenS1{F0: 42}
            _ = v53
        case "foo bar":
            v54 := GenS4{F0: 4294967295, F1: uint8(v4)}
            _ = v54
        }
        v2[(("hello" + "foo bar") + "\n\t")] = "foo bar"
        _ = v51
    }
    v5 = false
    v55 := GenS2{F0: 414, F1: -2147483648, F2: ""}
    switch 128 {
    case -410:
        v56 := v1[-572]
        _ = v56
    }
    v5 = ((-920 <= 0) || (false || v5))
    var v57 uint16 = 65535
    v57 = v57
    var v58 int = 698
    v58 += (v58 << ((65535) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v57
    _ = v58
    _ = v6
}
