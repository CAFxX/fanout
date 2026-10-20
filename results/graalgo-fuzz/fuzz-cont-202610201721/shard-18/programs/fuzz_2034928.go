package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int64; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 string }
func genFn0() bool {
    var v5 int = -217
    fmt.Println(v5)
    fmt.Println(v5)
    fmt.Println(v5)
    if false {
        v5 += -527
        v5 = v5
    } else {
        var v6 uint8 = 255
        var v7 int = (v5 + v5)
        fmt.Println(v5, v5)
        _ = v6
        _ = v7
    }
    _ = v5
    return false
}
func genFn1(p0 int16) bool {
    var v8 int = 755
    v8 += v8
    switch v8 {
    case 32768:
        switch 32767 {
        case -2147483648:
            v9 := GenS1{F0: 42, F1: (!true), F2: 18446744073709551615}
            _ = v9
        default:
            v10 := []int{(-(v8))}
            _ = v10
        }
    case -2147483648:
        switch v8 {
        case 9223372036854775807:
            v11 := GenS2{F0: 42, F1: 127}
            _ = v11
        case -791:
            if false {
                for fi12 := 0; fi12 < 19; fi12++ {
                    if (2147483648 >= 2) {
                        for fi13 := 0; fi13 < 15; fi13++ {
                            var v14 int16 = -32768
                            v14 = -32768
                            v8 += (fi13 + (-(fi13)))
                            _ = fi13
                            _ = v14
                        }
                    }
                    for fi15 := 0; fi15 < 3; fi15++ {
                        p0 -= (p0 << ((835) & 63))
                        _ = fi15
                    }
                    if (!("foo bar" > "hello")) {
                        var v16 uint64 = 1
                        v8 = v8
                        fmt.Println(v8, v8)
                        _ = v16
                    }
                    _ = fi12
                }
            } else {
                fmt.Println(p0, p0)
            }
        case 0:
            for fi17 := 0; fi17 < 12; fi17++ {
                v18 := []int{9223372036854775807, (v8 - (v8 + fi17)), v8, fi17}
                fmt.Println(v8)
                _ = fi17
                _ = v18
            }
        }
    }
    v19 := GenS3{F0: 9223372036854775807, F1: 948, F2: 142}
    _ = p0
    _ = v19
    _ = v8
    return (!(false || false))
}

func main() {
    var v1 bool = false
    v2 := make(map[string]uint32)
    var v3 int = 127
    var v4 uint = 65535
    var v20 int16 = 42
    v21 := GenS1{F0: 791, F1: true, F2: 4294967295}
    var v22 string = "hello"
    var v23 int8 = -1
    fmt.Println(v4, v22)
    switch (-(v3)) {
    case 127:
        for fi24 := 0; fi24 < 11; fi24++ {
            fmt.Println(v3)
            for fi25 := 0; fi25 < 12; fi25++ {
                if (v4 != (v4 + v4)) {
                    if (4294967295 <= v4) {
                        if (!true) {
                            v4 -= 255
                            if (v20 != v20) {
                                var v26 int16 = 1000
                                _ = v26
                            }
                            if (!v1) {
                                v23 -= (v23 / ((v23) | 1))
                            } else {
                                v22 = v22
                                fmt.Println(v23, v4, v20)
                            }
                            fmt.Println(v22, v1)
                        }
                    } else {
                        var v27 uint = (v4 - v4)
                        var v28 int64 = 1
                        var v29 int16 = (v20 & -1)
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    v1 = true
                }
                _ = len(v2)
                if (v1 && v1) {
                    if false {
                        var v30 uint32 = 2147483647
                        v31 := GenS3{F0: 42, F1: -236, F2: 1}
                        if false {
                            var v32 int8 = (^(v23))
                            fmt.Println(v32)
                            _ = v32
                        } else {
                            switch 329 {
                            case 0:
                                if (1.5 > float64(0.1)) {
                                    v33 := GenS4{F0: (v30 | v30)}
                                    v23 -= v23
                                    _ = v33
                                } else {
                                    if (false && true) {
                                        v30 += 0
                                        v22 = ("\n\t" + v22)
                                    }
                                    v34 := []int{-775}
                                    v22 = "foo bar"
                                    _ = v34
                                }
                            }
                        }
                        if true {
                            v35 := v2[v22]
                            _ = v35
                        }
                        v36 := v2[v22]
                        _ = v30
                        _ = v31
                        _ = v36
                    }
                }
                _ = fi25
            }
            _ = fi24
        }
    }
    if (!v1) {
        v37 := GenS5{F0: -1, F1: ""}
        var v38 uint = (^(v4))
        _ = v37
        _ = v38
    } else {
        v39 := v2["hello"]
        var v40 uint = 1
        _ = v39
        _ = v40
    }
    if false {
        delete(v2, "\n\t")
        for fi41 := 0; fi41 < 8; fi41++ {
            v22 = v22
            for fi42 := 0; fi42 < 15; fi42++ {
                v43 := GenS2{F0: (v4 | (^(v4))), F1: 128}
                var v44 int = (-(fi41))
                _ = fi42
                _ = v43
                _ = v44
            }
            _ = fi41
        }
        var v45 uint32 = 0
        _ = v45
    }
    var v46 uint8 = 227
    v20 = v20
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v46
}
