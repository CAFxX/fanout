package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 string; F2 string }
type GenS4 struct { F0 int16; F1 int8; F2 int8 }
type GenS5 struct { F0 int8; F1 int64 }
func genFn0() int {
    switch -9223372036854775808 {
    case 32767:
        if (0 > 42) {
            if true {
                v5 := GenS1{F0: 42}
                var v6 uint64 = 277
                _ = v5
                _ = v6
            }
        } else {
            if true {
                switch (("\n\t" + "foo bar") + "") {
                case "hello":
                    switch 0 {
                    case -9223372036854775808:
                        var v7 uint32 = 768
                        _ = v7
                    }
                case "foo bar":
                    var v8 int = -318
                    _ = v8
                case "":
                    var v9 int32 = 1
                    _ = v9
                default:
                    v10 := GenS1{F0: -40}
                    _ = v10
                }
                if false {
                    v11 := []int{-221, 32768}
                    _ = v11
                } else {
                    v12 := []int{}
                    _ = v12
                }
                var v13 uint8 = 127
                _ = v13
            }
            var v14 uint32 = 0
            switch 912 {
            case 127:
                v14 -= 65535
            case 128:
                if false {
                    fmt.Println(v14)
                } else {
                    v15 := GenS1{F0: 0}
                    _ = v15
                }
            default:
                fmt.Println(v14)
            }
            _ = v14
        }
    default:
        var v16 int = 0
        _ = v16
    }
    var v17 bool = false
    var v18 uint16 = 987
    _ = v17
    _ = v18
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    var v4 uint32 = 2147483647
    v19 := GenS1{F0: 108}
    if true {
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := []int{(fi20 | fi20)}
            fmt.Println(v1, v4, fi20)
            _ = fi20
            _ = v21
        }
        v22 := GenS2{F0: (v4 + 65535)}
        _ = v22
    } else {
        var v23 uint8 = 72
        v4 = (v4 << ((0) & 63))
        _ = v23
    }
    if false {
        for fi24 := 0; fi24 < 11; fi24++ {
            delete(v3, "a")
            var v25 uint32 = v4
            v25 -= v4
            _ = fi24
            _ = v25
        }
    } else {
        if (true && false) {
            var v26 int64 = -1
            fmt.Println(v4, v26, v1)
            _ = v26
        } else {
            fmt.Println(v1)
            v27 := v3["\n\t"]
            switch (v27 & (-(v27))) {
            case -9223372036854775808:
                var v28 int64 = 0
                _ = v28
            default:
                var v29 uint8 = 27
                _ = v29
            }
            _ = v27
        }
        var v30 uint64 = 18446744073709551615
        _ = v30
    }
    fmt.Println(v1, v4)
    var v31 int64 = 9223372036854775807
    v32 := GenS2{F0: v4}
    v33 := v3[""]
    if ((^(v4)) == (v4 * (v4 - (v4 & (^(v4)))))) {
        if (((v33 % ((-97) | 1)) > 32768) || (100 <= 1)) {
            fmt.Println(v31, v31)
        }
        var v34 uint16 = 65535
        v35 := GenS1{F0: 71}
        _ = v34
        _ = v35
    }
    v36 := GenS3{F0: ("\n\t" + "hello"), F1: (("a" + ("" + "")) + "hello"), F2: "foo bar"}
    if (0.1 <= (0.0 + 0.1)) {
        v37 := GenS4{F0: 42, F1: 31, F2: -1}
        if ((((!true) || false) && false) || true) {
            if (42 <= v4) {
                fmt.Println(v33)
            } else {
                v38 := []int{v33}
                _ = v38
            }
        } else {
            delete(v3, ("a" + "hello"))
        }
        for fi39 := 0; fi39 < 6; fi39++ {
            if ((v31 - (v31 & 1)) < v31) {
                if false {
                    v40 := v3["hello"]
                    _ = v40
                }
                switch v33 {
                case 536:
                    for fi41 := 0; fi41 < 2; fi41++ {
                        var v42 int32 = 0
                        var v43 int32 = v42
                        fmt.Println(v43, fi39, v33)
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                case -986:
                    v33 = (fi39 - (-(fi39)))
                default:
                    v44 := []int{}
                    _ = v44
                }
                var v45 uint16 = 2
                _ = v45
            }
            v31 = -9223372036854775808
            _ = fi39
        }
        _ = v37
    }
    var v46 uint32 = (v4 >> ((2) & 63))
    _ = len(v3)
    for fi47 := 0; fi47 < 20; fi47++ {
        v33 -= -1
        v46 = v46
        v48 := GenS1{F0: 106}
        _ = fi47
        _ = v48
    }
    v49 := GenS5{F0: 127, F1: -9223372036854775808}
    v50 := []int{127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v46
    _ = v49
    _ = v50
}
