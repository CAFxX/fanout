package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int32; F2 string }
type GenS3 struct { F0 int32; F1 bool; F2 uint64 }
type GenS4 struct { F0 string }
func genFn0(p0 int32, p1 uint, p2 bool) uint {
    v7 := GenS1{F0: -918}
    var v8 int8 = -1
    var v9 string = "\n\t"
    v10 := GenS2{F0: p0, F1: -2147483648, F2: (v9 + ("a" + ("foo bar" + ((((v9 + "a") + (v9 + v9)) + "foo bar") + ("" + v9)))))}
    var v11 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return (^(p1))
}

func main() {
    var v1 int = 127
    v2 := []int{v1, (^(v1)), (v1 + 42), (v1 >> ((2) & 63))}
    var v3 int8 = 1
    var v4 uint32 = 793
    v5 := make(map[string]int)
    var v6 int8 = -22
    var v12 int16 = -1
    for fi13 := 0; fi13 < 7; fi13++ {
        if true {
            var v14 bool = false
            _ = v14
        } else {
            v15 := GenS1{F0: v12}
            for fi16 := 0; fi16 < 20; fi16++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v5)
                _ = fi16
            }
            _ = v15
        }
        switch v1 {
        case 32767:
            v12 -= v12
        case -916:
            v4 -= (^(v4))
        }
        _ = fi13
    }
    var v17 float64 = float64(1e10)
    v18 := GenS3{F0: -2147483648, F1: true, F2: 631}
    v4 = (v4 >> ((0) & 63))
    var v19 int64 = 1
    v20 := v5["foo bar"]
    fmt.Println(v17)
    if len(v2) > 0 {
        v2[0] = (-(v20))
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 int32 = -2147483648
        fmt.Println(fi21, v3)
        _ = fi21
        _ = v22
    }
    if false {
        for fi23 := 0; fi23 < 19; fi23++ {
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 float32 = 0.0
                switch "hello" {
                case "foo bar":
                    _ = len(v2)
                case "":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "hello":
                    fmt.Println(v17, v1, fi24)
                default:
                    fmt.Println(v17, v6, fi24)
                }
                _ = fi24
                _ = v25
            }
            v12 += v12
            _ = fi23
        }
        if (!((-(v19)) >= 42)) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            fmt.Println(v17, v19, v4)
            v2 = append(v2, (-(v20)))
        }
        v3 -= 127
    }
    fmt.Println(v4, v19, v19)
    v12 -= (v12 % ((v12) | 1))
    v26 := GenS2{F0: 42, F1: 706, F2: (("a" + "") + (("foo bar" + "a") + (((("\n\t" + (("" + "foo bar") + "")) + ((("\n\t" + "foo bar") + "foo bar") + "foo bar")) + "hello") + (("\n\t" + "hello") + ""))))}
    for fi27 := 0; fi27 < 2; fi27++ {
        if true {
            v28 := GenS1{F0: v12}
            _ = v28
        } else {
            switch "\n\t" {
            case "a":
                v29 := GenS3{F0: -98, F1: (v19 == v19), F2: 18446744073709551615}
                _ = v29
            case "hello":
                if (-793 != 1) {
                    v12 -= v12
                    switch 9223372036854775807 {
                    case -2147483648:
                        fmt.Println(v6, fi27)
                    case 42:
                        v19 -= (v19 % (((v19 >> ((685) & 63))) | 1))
                    case 4294967295:
                        fmt.Println(v20, v17)
                    default:
                        switch 127 {
                        case -9223372036854775808:
                            switch ("foo bar" + ("hello" + "a")) {
                            case "\n\t":
                                v1 = int(v4)
                            case "":
                                v12 = -1
                            case "hello":
                                v6 += v6
                            }
                        case 4294967295:
                            v3 = v3
                        case 32767:
                            v30 := GenS2{F0: -606, F1: 521, F2: "a"}
                            _ = v30
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                }
            }
        }
        var v31 uint64 = 42
        var v32 uint32 = 2147483647
        _ = fi27
        _ = v31
        _ = v32
    }
    v33 := GenS4{F0: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
