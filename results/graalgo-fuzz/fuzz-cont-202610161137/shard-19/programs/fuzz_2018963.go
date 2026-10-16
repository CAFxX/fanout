package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint16; F1 uint64; F2 int32 }
func genFn0() uint32 {
    var v5 uint8 = 212
    v5 += 32
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int64 = -278
        _ = fi6
        _ = v7
    }
    var v8 int = 2147483647
    _ = v5
    _ = v8
    return 42
}
func genFn1(p0 int, p1 int, p2 string) int64 {
    if true {
        fmt.Println(p1)
        var v9 string = p2
        switch (v9 + (((((v9 + (v9 + v9)) + v9) + v9) + "a") + p2)) {
        case "":
            fmt.Println(p2, v9, p0)
        default:
            switch 0 {
            case 4294967295:
                for fi10 := 0; fi10 < 11; fi10++ {
                    if (!(1000 != 0)) {
                        if false {
                            for fi11 := 0; fi11 < 16; fi11++ {
                                v12 := GenS1{F0: 42}
                                p1 = p1
                                v13 := []int{(p1 >> ((18446744073709551615) & 63))}
                                var v14 int8 = 100
                                var v15 string = v9
                                _ = fi11
                                _ = v12
                                _ = v13
                                _ = v14
                                _ = v15
                            }
                        }
                        var v16 int16 = -1
                        v17 := GenS1{F0: 42}
                        _ = v16
                        _ = v17
                    } else {
                        v18 := []int{9223372036854775807, (p0 ^ (^(fi10)))}
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 19; fi19++ {
                        p0 = (p0 * (p0 - (-(p0))))
                        v20 := GenS1{F0: 32767}
                        p1 = p0
                        _ = fi19
                        _ = v20
                    }
                    var v21 int64 = -9223372036854775808
                    _ = fi10
                    _ = v21
                }
            case 127:
                fmt.Println(p0, v9)
            }
        }
        _ = v9
    }
    v22 := []int{(p0 ^ p0), 42}
    var v23 uint32 = 2
    p0 += (p0 & 545)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return 1
}

func main() {
    var v1 int32 = -1
    var v2 int64 = -581
    var v3 uint32 = 120
    var v4 uint64 = 447
    var v24 int = -8
    if false {
        var v25 uint32 = (v3 & v3)
        switch ("\n\t" + "") {
        case "a":
            if (!false) {
                var v26 uint64 = v4
                _ = v26
            } else {
                var v27 uint = 255
                for fi28 := 0; fi28 < 19; fi28++ {
                    v1 = v1
                    v29 := GenS1{F0: 2}
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
        case "":
            v30 := GenS2{F0: 32767, F1: ("foo bar" + ("" + "\n\t"))}
            _ = v30
        case "foo bar":
            var v31 int32 = 0
            _ = v31
        default:
            if true {
                v24 -= -955
                for fi32 := 0; fi32 < 20; fi32++ {
                    v33 := GenS2{F0: -22, F1: "a"}
                    v4 -= v4
                    v1 -= (v1 / ((-1) | 1))
                    _ = fi32
                    _ = v33
                }
                v4 = v4
            } else {
                var v34 uint = 103
                v1 -= 1
                _ = v34
            }
        }
        _ = v25
    } else {
        switch ("hello" + "") {
        case "":
            v35 := GenS3{F0: 65535, F1: (^(v4)), F2: (v1 % ((-1) | 1))}
            _ = v35
        }
    }
    fmt.Println(v2)
    v2 = v2
    var v36 int = (v24 + (v24 % ((2147483647) | 1)))
    v1 = (^(v1))
    fmt.Println(v2)
    fmt.Println(v1, v1, v1)
    for fi37 := 0; fi37 < 10; fi37++ {
        v4 = (v4 >> ((42) & 63))
        for fi38 := 0; fi38 < 15; fi38++ {
            v39 := []int{(fi37 * v36), (v24 - (fi37 ^ (^(fi38)))), 69, v24}
            var v40 float64 = float64(((0.0 + 3.14159) * 0.1))
            switch "hello" {
            case "\n\t":
                var v41 float64 = v40
                _ = v41
            }
            _ = fi38
            _ = v39
            _ = v40
        }
        for fi42 := 0; fi42 < 19; fi42++ {
            switch (fi37 & (v24 * fi42)) {
            case 128:
                v1 = -1
            default:
                v3 = 0
            }
            switch "\n\t" {
            case "\n\t":
                if (2 != 4294967295) {
                    v3 -= 0
                }
            case "foo bar":
                v43 := []int{v24, (fi37 ^ v24), (fi37 + (v24 * (fi42 * 973)))}
                _ = v43
            case "hello":
                v1 -= (v1 / (((^(v1))) | 1))
            default:
                v44 := GenS1{F0: 573}
                _ = v44
            }
            _ = fi42
        }
        _ = fi37
    }
    var v45 int16 = -403
    var v46 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v4
    _ = v45
    _ = v46
}
