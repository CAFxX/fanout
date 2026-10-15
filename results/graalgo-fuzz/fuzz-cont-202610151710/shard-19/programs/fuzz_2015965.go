package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 string, p1 int64, p2 bool) int8 {
    v7 := GenS1{F0: (p1 * (p1 - (p1 << ((783) & 63)))), F1: ("a" + (p0 + "")), F2: (p1 != 1)}
    var v8 float32 = 0.1
    v9 := GenS1{F0: (p1 & p1), F1: p0, F2: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 100
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = (v1 | (v1 << ((18446744073709551615) & 63)))
    var v3 uint = 4294967295
    v4 := make(map[uint32]uint32)
    var v5 int = -9223372036854775808
    var v6 string = ""
    if (("\n\t" != (v6 + "")) && true) {
        v3 += 7
        if false {
            for fi10 := 0; fi10 < 12; fi10++ {
                v1 += -9223372036854775808
                switch (fi10 >> ((v3) & 63)) {
                case 865:
                    for fi11 := 0; fi11 < 18; fi11++ {
                        switch v6 {
                        case "":
                            v12 := []int{(-(fi11))}
                            _ = v12
                        case "hello":
                            v3 -= v3
                        case "a":
                            var v13 string = (((v6 + v6) + v6) + v6)
                            _ = v13
                        default:
                            v3 = v3
                        }
                        var v14 uint16 = 715
                        _ = fi11
                        _ = v14
                    }
                case 127:
                    v6 = "\n\t"
                default:
                    if (2147483647 != fi10) {
                        var v15 float32 = (0.0 / (-2.25 + 0.0))
                        fmt.Println(fi10, v1)
                        v16 := GenS2{F0: v6}
                        _ = v15
                        _ = v16
                    } else {
                        v5 -= 0
                        v17 := []int{(-(v5))}
                        _ = v17
                    }
                }
                v6 = "foo bar"
                _ = fi10
            }
        } else {
            v18 := []int{}
            var v19 int16 = -998
            _ = v18
            _ = v19
        }
    }
    v20 := v4[2]
    v21 := GenS3{F0: -480, F1: (^(v3))}
    v6 = v6
    v2 -= (v2 % ((v1) | 1))
    v22 := []int{int(v3)}
    _ = len(v22)
    var v23 uint = v3
    switch 32767 {
    case 0:
        var v24 bool = ((v2 >> ((v3) & 63)) == v2)
        _ = v24
    case 9223372036854775807:
        for fi25 := 0; fi25 < 7; fi25++ {
            v5 = 958
            _ = fi25
        }
    }
    var v26 uint = (v23 / ((876) | 1))
    for fi27 := 0; fi27 < 7; fi27++ {
        _ = len(v22)
        var v28 float32 = ((((((0.0 + 1.5) - 3.14159) * float32(1.5)) - (0.1 * 0.0)) - -2.25) / 1.0)
        var v29 uint16 = 1
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int8 = 100
    v5 = (^(v5))
    v31 := GenS4{F0: (v2 & v1)}
    if true {
        if true {
            v3 = v26
            if true {
                fmt.Println(v20, v1, v20)
            } else {
                var v32 int8 = -1
                v30 = genFn0("\n\t", (v2 % (((v2 & v1)) | 1)), true)
                _ = v32
            }
        }
        var v33 int8 = (v30 - v30)
        v33 = v33
        _ = v33
    }
    if false {
        if false {
            v2 = v2
        } else {
            v34 := GenS2{F0: (v6 + v6)}
            var v35 uint16 = 32767
            v2 = (v2 + v1)
            _ = v34
            _ = v35
        }
        v3 = v23
        if (!false) {
            v36 := GenS5{F0: v30}
            _ = v36
        } else {
            fmt.Println(v26)
            for fi37 := 0; fi37 < 3; fi37++ {
                for fi38 := 0; fi38 < 11; fi38++ {
                    v39 := GenS3{F0: int64(v20), F1: v26}
                    _ = fi38
                    _ = v39
                }
                _ = fi37
            }
            switch v6 {
            case "a":
                if false {
                    switch v5 {
                    case 4294967295:
                        switch v6 {
                        case "hello":
                            var v40 uint32 = v20
                            _ = v40
                        case "\n\t":
                            for fi41 := 0; fi41 < 18; fi41++ {
                                var v42 uint32 = 71
                                _ = fi41
                                _ = v42
                            }
                        case "":
                            var v43 uint32 = (^(v20))
                            _ = v43
                        }
                    default:
                        v44 := GenS4{F0: (v1 - v2)}
                        _ = v44
                    }
                }
            case "foo bar":
                fmt.Println(v2)
            }
        }
    }
    if (!false) {
        v45 := GenS3{F0: v1, F1: v3}
        fmt.Println(v1)
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
