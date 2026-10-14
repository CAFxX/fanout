package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint64) int16 {
    p0 = p0
    v7 := GenS1{F0: 127}
    if true {
        v8 := []int{}
        _ = v8
    } else {
        var v9 uint8 = 0
        p0 += (^(p0))
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 bool = (!(true && (!true)))
            var v12 bool = v11
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []string{("hello" + "foo bar"), "", (("\n\t" + "hello") + ""), "foo bar"}
    var v3 bool = false
    var v4 uint64 = 329
    var v5 uint8 = 253
    var v6 uint32 = v1
    var v13 int8 = 42
    if v3 {
        v14 := GenS1{F0: -103}
        if ((v6 | v6) != (v6 ^ (v6 << ((42) & 63)))) {
            if (v6 < (v6 >> ((2) & 63))) {
                fmt.Println(v5)
                fmt.Println(v1, v1, v5)
            } else {
                v13 = v13
                v15 := GenS1{F0: v13}
                v13 -= v13
                _ = v15
            }
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            _ = len(v2)
        }
        v4 = v4
        _ = v14
    } else {
        if v3 {
            v3 = v3
            v4 += v4
        }
    }
    fmt.Println(v13, v6)
    v1 = 42
    v2 = append(v2, "a")
    v16 := GenS2{F0: v13}
    switch "" {
    case "a":
        v3 = (42 >= -9223372036854775808)
    case "foo bar":
        for fi17 := 0; fi17 < 9; fi17++ {
            fmt.Println(v1)
            v13 += v13
            if v3 {
                if false {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        var v19 float32 = (-2.25 / 1.0)
                        _ = fi18
                        _ = v19
                    }
                    var v20 uint64 = (^(v4))
                    v6 = (^(v1))
                    _ = v20
                }
                for fi21 := 0; fi21 < 19; fi21++ {
                    if (((^(v6)) > 2147483647) || (("a" != ("hello" + "")) || v3)) {
                        switch fi17 {
                        case 128:
                            v6 = 2147483647
                        case 4294967295:
                            var v22 uint32 = (v6 << ((0) & 63))
                            _ = v22
                        case 127:
                            var v23 int = (-(fi17))
                            _ = v23
                        default:
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                        }
                        v24 := GenS1{F0: v13}
                        _ = v24
                    } else {
                        fmt.Println(fi17)
                        switch fi21 {
                        case 1:
                            if (!(v3 && false)) {
                                v13 = v13
                                var v25 int64 = 996
                                _ = v25
                            } else {
                                var v26 uint32 = 2147483648
                                _ = v26
                            }
                        case -2147483648:
                            v13 += 127
                        case -529:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            var v27 uint8 = v5
                            v28 := GenS2{F0: (v13 % ((v13) | 1))}
                            _ = v27
                            _ = v28
                        }
                        v1 -= uint32(v4)
                        var v29 float64 = ((3.14159 * 0.0) - float64(((0.0 / (-2.25 / -2.25)) + 1e10)))
                        var v30 int16 = -311
                        _ = v29
                        _ = v30
                    }
                    v4 = 0
                    v31 := GenS2{F0: v13}
                    _ = fi21
                    _ = v31
                }
            }
            _ = fi17
        }
    default:
        var v32 string = "\n\t"
        _ = v32
    }
    v5 = (v5 ^ (v5 - 42))
    for fi33 := 0; fi33 < 20; fi33++ {
        var v34 int = (fi33 + (fi33 / ((-2147483648) | 1)))
        var v35 int16 = 228
        if v3 {
            v1 -= v1
            for fi36 := 0; fi36 < 17; fi36++ {
                _ = len(v2)
                fmt.Println(v3, fi33)
                _ = fi36
            }
            for fi37 := 0; fi37 < 5; fi37++ {
                v38 := GenS2{F0: (v13 & -128)}
                v39 := GenS1{F0: 0}
                var v40 float64 = -2.25
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    switch "\n\t" {
    case "hello":
        switch 9223372036854775807 {
        case 2147483647:
            v41 := GenS1{F0: v13}
            _ = v41
        case -1:
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
        case 0:
            var v42 int64 = 1
            _ = v42
        default:
            v43 := GenS3{F0: (^(v1))}
            _ = v43
        }
    case "":
        var v44 uint64 = (v4 * 177)
        _ = v44
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
