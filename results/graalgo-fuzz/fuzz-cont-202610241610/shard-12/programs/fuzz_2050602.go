package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint8, p1 int8) int {
    v4 := GenS1{F0: uint16(p0), F1: (!true)}
    if true {
        v5 := []int{}
        switch "" {
        case "a":
            v6 := GenS2{F0: -381}
            _ = v6
        case "foo bar":
            if false {
                p1 -= -1
            }
        default:
            var v7 string = ""
            _ = v7
        }
        _ = v5
    }
    fmt.Println(p0)
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]int)
    var v3 uint32 = 118
    switch 1 {
    case -2147483648:
        if false {
            v8 := v1["hello"]
            _ = v8
        }
    case 9223372036854775807:
        switch -1 {
        case 9223372036854775807:
            v3 = (^(v3))
        case 0:
            var v9 int64 = 1
            _ = v9
        }
    case 42:
        for fi10 := 0; fi10 < 17; fi10++ {
            fmt.Println(fi10)
            _ = fi10
        }
    }
    switch 188 {
    case -132:
        if (true && false) {
            var v11 uint8 = 65
            var v12 float64 = 0.0
            _ = v11
            _ = v12
        }
    }
    v3 -= v3
    fmt.Println(v3)
    switch (("foo bar" + "hello") + ("\n\t" + "hello")) {
    case "hello":
        var v13 bool = (true && (-2.25 >= ((3.14159 + -2.25) / 3.14159)))
        _ = v13
    case "\n\t":
        switch 884 {
        case 2147483647:
            var v14 int32 = -2147483648
            _ = v14
        case -937:
            switch "a" {
            case "\n\t":
                fmt.Println(v3)
            default:
                var v15 int = genFn0(2, -128)
                _ = v15
            }
        case 32767:
            fmt.Println(v3)
        }
    default:
        v16 := GenS2{F0: genFn0(170, -96)}
        _ = v16
    }
    var v17 int = 1
    v18 := GenS3{F0: 1, F1: 42, F2: 255}
    if (false || false) {
        v3 -= v3
        for fi19 := 0; fi19 < 11; fi19++ {
            fmt.Println(v3)
            _ = fi19
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            if (0 > 468) {
                v21 := v1[("hello" + "")]
                if true {
                    fmt.Println(v17, fi20, fi20)
                    fmt.Println(v21)
                    v2[34] = fi20
                } else {
                    delete(v2, 2)
                    var v22 uint = 4294967295
                    v1["\n\t"] = (v21 << (((v22 + 691)) & 63))
                    _ = v22
                }
                var v23 uint = 708
                _ = v21
                _ = v23
            }
            _ = fi20
        }
    } else {
        if ((0.0 + (0.0 - (-2.25 * (1.5 * (1.5 - 3.14159))))) > 1e10) {
            for fi24 := 0; fi24 < 16; fi24++ {
                v25 := GenS2{F0: genFn0(215, -122)}
                if true {
                    var v26 uint64 = 4294967295
                    var v27 uint64 = (v26 << ((42) & 63))
                    _ = v26
                    _ = v27
                }
                _ = fi24
                _ = v25
            }
        }
    }
    if (false && false) {
        fmt.Println(v3)
        switch "" {
        case "":
            var v28 float32 = 1.5
            _ = v28
        case "foo bar":
            if false {
                v29 := v2[v3]
                var v30 uint32 = (v3 | 1)
                _ = v29
                _ = v30
            }
        case "\n\t":
            if (!false) {
                var v31 uint8 = 2
                _ = v31
            } else {
                v32 := GenS2{F0: 32767}
                var v33 int16 = 0
                v3 -= (v3 - v3)
                _ = v32
                _ = v33
            }
        }
        var v34 int32 = 42
        _ = v34
    } else {
        v3 = (v3 & v3)
        if (true && true) {
            v35 := GenS3{F0: 391, F1: -1, F2: 72}
            for fi36 := 0; fi36 < 1; fi36++ {
                if false {
                    if true {
                        var v37 uint = 4294967295
                        _ = v37
                    } else {
                        v38 := GenS4{F0: 1}
                        _ = len(v1)
                        _ = v38
                    }
                    switch genFn0(58, -128) {
                    case 877:
                        v2[v3] = fi36
                    case 119:
                        for fi39 := 0; fi39 < 15; fi39++ {
                            var v40 uint8 = 255
                            fmt.Println(v40, fi39)
                            fmt.Println(fi36, fi36, v17)
                            _ = fi39
                            _ = v40
                        }
                    case 722:
                        v17 -= 32767
                    default:
                        for fi41 := 0; fi41 < 5; fi41++ {
                            var v42 uint32 = (v3 ^ 4294967295)
                            v42 = 42
                            v43 := GenS1{F0: 1, F1: (!true)}
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                    }
                    v17 = (fi36 % ((fi36) | 1))
                } else {
                    delete(v1, ("hello" + "a"))
                }
                _ = fi36
            }
            _ = v35
        }
    }
    var v44 string = "a"
    v3 += (v3 - v3)
    v45 := GenS2{F0: (v17 & genFn0(80, -128))}
    v46 := GenS1{F0: 65535, F1: (2 >= 0)}
    fmt.Println(v44, v44, v17)
    if false {
        v3 = (v3 % (((v3 + v3)) | 1))
        v47 := []int{(v17 << ((22) & 63))}
        _ = v47
    } else {
        for fi48 := 0; fi48 < 1; fi48++ {
            v49 := []int{61}
            _ = fi48
            _ = v49
        }
        var v50 uint8 = 247
        switch v17 {
        case 0:
            v44 = "\n\t"
        case 42:
            fmt.Println(v50, v50, v50)
        case 127:
            var v51 int = genFn0(v50, 42)
            _ = v51
        }
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
}
