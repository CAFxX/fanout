package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint }
type GenS2 struct { F0 uint16; F1 int32; F2 uint }
type GenS3 struct { F0 int32; F1 uint16; F2 string }
type GenS4 struct { F0 uint; F1 int }
func genFn0(p0 string) int {
    if false {
        fmt.Println(p0)
    } else {
        v7 := GenS1{F0: 2, F1: 42, F2: 65535}
        if (125 > 388) {
            p0 = (("\n\t" + p0) + ("a" + (p0 + "")))
            v8 := GenS1{F0: 255, F1: 852, F2: 0}
            v9 := GenS1{F0: 255, F1: 4294967295, F2: 255}
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: 1, F1: 518, F2: 243}
            _ = v10
        }
        _ = v7
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(fi11)
        _ = fi11
    }
    var v12 int32 = 1
    _ = p0
    _ = v12
    return 2147483647
}
func genFn1() int64 {
    if true {
        v13 := []int{2147483647, 127}
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        v14 := GenS1{F0: 255, F1: 65535, F2: 691}
        _ = v13
        _ = v14
    }
    if (255 > 157) {
        var v15 uint32 = 42
        v16 := GenS2{F0: 42, F1: 0, F2: 147}
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := []int{(fi17 - fi17)}
            var v19 int16 = 32767
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 18; fi20++ {
            switch (fi20 % ((-9223372036854775808) | 1)) {
            case 41:
                var v21 int16 = 1
                _ = v21
            default:
                if false {
                    fmt.Println(fi20)
                    var v22 uint8 = 219
                    _ = v22
                }
            }
            v23 := []int{(-(fi20))}
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            _ = fi20
            _ = v23
        }
    }
    var v24 bool = true
    v25 := GenS3{F0: -2147483648, F1: 65535, F2: "foo bar"}
    _ = v24
    _ = v25
    return -9223372036854775808
}

func main() {
    v1 := []string{"hello", ("\n\t" + "foo bar"), "foo bar"}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    v4 := []int{}
    var v5 int64 = 748
    v6 := []int{}
    for fi26 := 0; fi26 < 10; fi26++ {
        v5 += 42
        for fi27 := 0; fi27 < 9; fi27++ {
            v5 -= genFn1()
            fmt.Println(v5)
            _ = fi27
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            fmt.Println(fi28, fi28, fi26)
            _ = fi28
        }
        _ = fi26
    }
    switch -866 {
    case 127:
        var v29 uint64 = 242
        _ = v29
    case 9223372036854775807:
        v3[0] = "foo bar"
    default:
        for fi30 := 0; fi30 < 4; fi30++ {
            if true {
                var v31 string = ("\n\t" + "a")
                if false {
                    for fi32 := 0; fi32 < 6; fi32++ {
                        fmt.Println(fi30)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi32
                    }
                } else {
                    _ = len(v3)
                    for fi33 := 0; fi33 < 18; fi33++ {
                        fmt.Println(fi30)
                        v34 := GenS1{F0: 2, F1: 65535, F2: 65535}
                        _ = fi33
                        _ = v34
                    }
                }
                _ = v31
            } else {
                v5 += (v5 % ((genFn1()) | 1))
                v35 := GenS3{F0: -2147483648, F1: 32767, F2: "\n\t"}
                fmt.Println(v5)
                _ = v35
            }
            v2["a"] = 0
            _ = fi30
        }
    }
    if true {
        v5 += (^(v5))
        v5 = v5
        for fi36 := 0; fi36 < 15; fi36++ {
            fmt.Println(v5, v5)
            _ = fi36
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (1 < 42) {
        v5 += v5
        var v37 float64 = (((3.14159 * 0.1) - 0.0) * 0.1)
        v37 = v37
        _ = v37
    } else {
        if false {
            for fi38 := 0; fi38 < 10; fi38++ {
                for fi39 := 0; fi39 < 8; fi39++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v40 uint = 4294967295
                    _ = fi39
                    _ = v40
                }
                for fi41 := 0; fi41 < 14; fi41++ {
                    v5 = v5
                    v42 := v2["hello"]
                    v5 = (-(v5))
                    _ = fi41
                    _ = v42
                }
                _ = fi38
            }
            if ((0 == 1) && true) {
                var v43 float32 = (float32(3.14159) + 0.0)
                for fi44 := 0; fi44 < 7; fi44++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch "hello" {
                    case "":
                        var v45 float32 = v43
                        _ = v45
                    case "foo bar":
                        var v46 int32 = -1
                        _ = v46
                    case "\n\t":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    switch "a" {
                    case "a":
                        delete(v3, 1)
                    case "foo bar":
                        for fi47 := 0; fi47 < 16; fi47++ {
                            var v48 int = -670
                            v49 := GenS2{F0: 0, F1: -2147483648, F2: 4294967295}
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                    case "hello":
                        v50 := GenS2{F0: 42, F1: 0, F2: 1}
                        _ = v50
                    }
                    _ = fi44
                }
                var v51 float64 = 1e10
                _ = v43
                _ = v51
            }
        }
        if false {
            var v52 int16 = 1000
            _ = v52
        }
        switch (("a" + "a") + "\n\t") {
        case "\n\t":
            var v53 float32 = 0.0
            _ = v53
        }
    }
    var v54 float32 = -2.25
    var v55 int64 = (v5 << ((165) & 63))
    v5 = v55
    v56 := GenS4{F0: 4294967295, F1: genFn0("\n\t")}
    var v57 int = genFn0("foo bar")
    v55 += (^(v5))
    fmt.Println(v54, v5)
    var v58 uint32 = 42
    fmt.Println(v57, v58, v5)
    v59 := GenS4{F0: 18446744073709551615, F1: v57}
    v55 += (v5 & 551)
    v4 = append(v4, v57)
    v54 += (1.5 / 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
}
