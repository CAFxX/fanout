package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS4 struct { F0 int16; F1 int32; F2 bool }
type GenS5 struct { F0 int8; F1 int16; F2 int32 }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 bool; F1 uint64; F2 uint32 }
func genFn0(p0 int64, p1 int8, p2 int64) bool {
    p1 = p1
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: "a"}
        var v8 uint = 129
        v9 := GenS2{F0: 0}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        if true {
            if (42 < 18446744073709551615) {
                v11 := GenS3{F0: 73, F1: 2147483647, F2: (-(p1))}
                var v12 bool = false
                var v13 int = 128
                _ = v11
                _ = v12
                _ = v13
            }
            v14 := []int{}
            _ = v14
        }
        _ = fi10
    }
    var v15 int = -9223372036854775808
    v16 := GenS3{F0: 0, F1: 1, F2: -128}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return ((0.1 - 0.1) < 1e10)
}

func main() {
    var v1 uint8 = 18
    var v2 uint16 = 1
    v3 := []string{"", ("\n\t" + "foo bar"), ("a" + ("" + "foo bar")), ("" + "a")}
    v4 := make(map[uint32]int)
    v5 := make(map[int]string)
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            v2 -= v2
            _ = fi18
        }
        if true {
            v19 := GenS1{F0: "foo bar"}
            _ = v19
        }
        v1 = v1
        _ = fi17
    }
    v5[32768] = "\n\t"
    v20 := GenS4{F0: 0, F1: -2147483648, F2: true}
    fmt.Println(v2, v1)
    v21 := GenS5{F0: 0, F1: 1, F2: -1}
    for fi22 := 0; fi22 < 13; fi22++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v2 = (v2 * v2)
        _ = fi22
    }
    v1 += v1
    if ("hello" == "") {
        for fi23 := 0; fi23 < 14; fi23++ {
            if false {
                fmt.Println(fi23)
                for fi24 := 0; fi24 < 1; fi24++ {
                    v3 = append(v3, "foo bar")
                    var v25 int8 = 42
                    _ = fi24
                    _ = v25
                }
            }
            v2 += 793
            for fi26 := 0; fi26 < 19; fi26++ {
                v27 := GenS4{F0: 32767, F1: -887, F2: (((816 > 1) || (true && false)) && false)}
                _ = fi26
                _ = v27
            }
            _ = fi23
        }
        for fi28 := 0; fi28 < 5; fi28++ {
            if (((!false) || (true && true)) && true) {
                switch (fi28 % ((fi28) | 1)) {
                case 877:
                    var v29 int16 = -687
                    _ = v29
                default:
                    v30 := GenS5{F0: 0, F1: 42, F2: -1}
                    _ = v30
                }
            } else {
                v31 := GenS5{F0: -67, F1: -1, F2: 713}
                _ = v31
            }
            _ = fi28
        }
        v2 = 488
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        var v33 int64 = -9223372036854775808
        if false {
            v34 := GenS6{F0: 42, F1: 42}
            _ = v34
        } else {
            v33 = v33
        }
        _ = fi32
        _ = v33
    }
    if (false && false) {
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 uint32 = 2147483648
            _ = fi35
            _ = v36
        }
        var v37 int64 = 1
        _ = v37
    } else {
        var v38 string = ("" + "hello")
        v39 := GenS5{F0: 100, F1: 42, F2: 0}
        _ = v38
        _ = v39
    }
    var v40 string = "foo bar"
    for fi41 := 0; fi41 < 9; fi41++ {
        var v42 int8 = int8(fi41)
        fmt.Println(fi41, fi41, fi41)
        _ = fi41
        _ = v42
    }
    fmt.Println(v40, v1)
    switch -2147483648 {
    case -167:
        v2 = (v2 + 42)
    case 1:
        if ("a" <= v40) {
            if (18446744073709551615 >= 1) {
                v43 := GenS7{F0: false, F1: 2, F2: 675}
                for fi44 := 0; fi44 < 13; fi44++ {
                    v2 -= 868
                    v45 := v4[2147483648]
                    _ = len(v4)
                    _ = fi44
                    _ = v45
                }
                for fi46 := 0; fi46 < 18; fi46++ {
                    delete(v4, 106)
                    var v47 float64 = (float64(3.14159) / (((3.14159 + 3.14159) * -2.25) - 1e10))
                    if (1 > v1) {
                        fmt.Println(v2)
                        var v48 float64 = float64(0.0)
                        _ = v48
                    } else {
                        var v49 uint8 = 156
                        var v50 int16 = 32767
                        v2 = v2
                        _ = v49
                        _ = v50
                    }
                    _ = fi46
                    _ = v47
                }
                _ = v43
            }
        } else {
            switch "foo bar" {
            case "":
                if (18446744073709551615 > 918) {
                    _ = len(v5)
                    v51 := v4[4294967295]
                    v51 = v51
                    _ = v51
                } else {
                    var v52 uint32 = 619
                    _ = v52
                }
            }
            fmt.Println(v2, v1)
            var v53 int8 = 42
            _ = v53
        }
    case 9223372036854775807:
        var v54 uint32 = 2147483647
        _ = v54
    }
    v3 = append(v3, v40)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
