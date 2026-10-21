package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 string }
func genFn0(p0 uint64, p1 int64, p2 uint16) int8 {
    fmt.Println(p2, p2)
    v5 := GenS1{F0: 2147483647}
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 53
}
func genFn1(p0 int64) uint32 {
    for fi6 := 0; fi6 < 5; fi6++ {
        switch "hello" {
        case "a":
            v7 := GenS1{F0: 1}
            _ = v7
        case "foo bar":
            v8 := GenS1{F0: -1}
            _ = v8
        case "\n\t":
            if (255 == 42) {
                fmt.Println(fi6, p0)
                p0 -= (p0 | (p0 & p0))
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 uint32 = 545
                    _ = fi9
                    _ = v10
                }
            }
        }
        _ = fi6
    }
    fmt.Println(p0)
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint16 = 619
        if (false || true) {
            fmt.Println(v12, v12, fi11)
            v13 := []int{739, (fi11 ^ (fi11 ^ fi11))}
            p0 = -237
            _ = v13
        } else {
            var v14 float64 = 0.0
            v15 := []int{(fi11 * fi11), fi11, (fi11 & 128)}
            v16 := GenS1{F0: -1}
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi11
        _ = v12
    }
    var v17 string = (("\n\t" + ("foo bar" + "hello")) + ((("\n\t" + "a") + "") + "hello"))
    var v18 float64 = (0.0 + float64(3.14159))
    _ = p0
    _ = v17
    _ = v18
    return 4294967295
}

func main() {
    var v1 uint8 = 128
    var v2 uint16 = 0
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    switch "foo bar" {
    case "":
        var v19 uint = 0
        _ = v19
    case "a":
        fmt.Println(v2)
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        delete(v4, "a")
        fmt.Println(v2, v2)
        _ = fi20
    }
    if (2147483647 == -2147483648) {
        v21 := []int{}
        _ = v21
    } else {
        v2 += (v2 << ((42) & 63))
        var v22 bool = (21 != -128)
        _ = v22
    }
    v1 = 255
    var v23 uint16 = v2
    fmt.Println(v2, v1, v23)
    if true {
        for fi24 := 0; fi24 < 19; fi24++ {
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 float64 = ((1.5 * 1.5) * float64(1e10))
                v23 += 12
                v27 := GenS1{F0: 1}
                _ = fi25
                _ = v26
                _ = v27
            }
            if true {
                switch ("" + (("\n\t" + ("a" + "\n\t")) + ("a" + ""))) {
                case "\n\t":
                    switch "\n\t" {
                    case "hello":
                        for fi28 := 0; fi28 < 10; fi28++ {
                            var v29 bool = true
                            if (v1 <= 255) {
                                var v30 uint64 = 42
                                v31 := []int{333, (fi24 - -595), (fi24 << ((829) & 63))}
                                _ = v30
                                _ = v31
                            }
                            v32 := []int{fi28, 26, fi28, (-(fi28))}
                            _ = fi28
                            _ = v29
                            _ = v32
                        }
                    case "\n\t":
                        v33 := GenS2{F0: 18446744073709551615}
                        _ = v33
                    case "a":
                        v34 := []int{(fi24 | (fi24 << ((65535) & 63))), (fi24 << ((480) & 63)), fi24}
                        _ = v34
                    }
                }
            } else {
                v23 = v23
            }
            v2 += (v23 / ((65535) | 1))
            _ = fi24
        }
    }
    switch -9223372036854775808 {
    case 32767:
        switch 1 {
        case 2147483647:
            switch ("\n\t" + ("a" + "\n\t")) {
            case "":
                for fi35 := 0; fi35 < 16; fi35++ {
                    var v36 uint16 = (v2 ^ v23)
                    v37 := v3[(fi35 - -178)]
                    v38 := GenS3{F0: 0, F1: ("foo bar" + "")}
                    _ = fi35
                    _ = v36
                    _ = v37
                    _ = v38
                }
            default:
                fmt.Println(v1, v23)
            }
        default:
            var v39 int64 = -731
            _ = v39
        }
    case 0:
        var v40 int = 32767
        _ = v40
    case 128:
        v2 -= 732
    }
    if false {
        if true {
            v41 := []int{}
            _ = v41
        }
        if false {
            v23 -= v23
            for fi42 := 0; fi42 < 1; fi42++ {
                switch "a" {
                case "a":
                    var v43 uint16 = v23
                    _ = v43
                case "hello":
                    v44 := []int{(fi42 % (((fi42 << ((95) & 63))) | 1)), fi42, fi42}
                    _ = v44
                case "":
                    fmt.Println(v23, v2, fi42)
                default:
                    fmt.Println(v2)
                }
                _ = fi42
            }
            var v45 int32 = -1
            _ = v45
        }
    } else {
        v23 += 0
    }
    v46 := GenS3{F0: 18446744073709551615, F1: ("hello" + "foo bar")}
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v46
}
