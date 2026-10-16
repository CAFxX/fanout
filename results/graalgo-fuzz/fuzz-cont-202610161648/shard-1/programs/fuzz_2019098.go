package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint8; F2 string }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 string; F1 string; F2 int }
func genFn0() string {
    var v4 uint32 = 1
    v5 := GenS1{F0: -2147483648}
    for fi6 := 0; fi6 < 16; fi6++ {
        v4 += (^(v4))
        _ = fi6
    }
    _ = v4
    _ = v5
    return "a"
}
func genFn1(p0 string, p1 uint, p2 uint16) uint {
    p1 -= p1
    p0 = "foo bar"
    v7 := []int{-1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return uint(p2)
}

func main() {
    var v1 int = 417
    v2 := make(map[string]uint32)
    var v3 uint32 = 2147483648
    _ = len(v2)
    v1 -= v1
    v8 := GenS1{F0: v1}
    fmt.Println(v1, v3)
    v2[("" + ("a" + "\n\t"))] = 4294967295
    v1 += v1
    v1 += 0
    v3 = v3
    v9 := GenS2{F0: -2147483648, F1: 2, F2: (("" + "a") + ("hello" + ("foo bar" + "foo bar")))}
    if (false && true) {
        var v10 bool = (true || true)
        for fi11 := 0; fi11 < 10; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                fmt.Println(fi12, v10, v3)
                _ = fi12
            }
            _ = fi11
        }
        v13 := GenS3{F0: 42, F1: (v3 - v3)}
        _ = v10
        _ = v13
    }
    if ("" <= "hello") {
        var v14 uint32 = (v3 << ((genFn1("\n\t", 61, 482)) & 63))
        v1 = v1
        fmt.Println(v3, v1)
        _ = v14
    }
    fmt.Println(v3, v3)
    v1 = (v1 & int(v3))
    fmt.Println(v1)
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := []int{795, (v1 & v1)}
        var v17 int32 = 420
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: v1, F1: 172, F2: (("foo bar" + "") + ("a" + "hello"))}
    if true {
        switch v1 {
        case 42:
            switch -2147483648 {
            case -131:
                for fi19 := 0; fi19 < 5; fi19++ {
                    if true {
                        for fi20 := 0; fi20 < 6; fi20++ {
                            v2[""] = (^(v3))
                            _ = fi20
                        }
                    } else {
                        var v21 float64 = 0.0
                        v22 := v2["\n\t"]
                        _ = v21
                        _ = v22
                    }
                    if false {
                        var v23 float32 = ((0.0 / 1.0) - (1e10 / float32((0.1 - 0.0))))
                        v24 := GenS1{F0: 2147483647}
                        _ = v23
                        _ = v24
                    } else {
                        delete(v2, (("hello" + "\n\t") + "a"))
                    }
                    switch "\n\t" {
                    case "":
                        var v25 uint16 = 255
                        _ = v25
                    case "hello":
                        for fi26 := 0; fi26 < 6; fi26++ {
                            switch (fi19 - fi19) {
                            case 9223372036854775807:
                                if false {
                                    v27 := GenS1{F0: fi19}
                                    v1 -= 1
                                    switch "\n\t" {
                                    case "":
                                        v1 = (v1 ^ (fi19 + -2147483648))
                                    case "a":
                                        v28 := GenS4{F0: -492, F1: 558}
                                        _ = v28
                                    case "\n\t":
                                        v29 := GenS5{F0: 1, F1: -699}
                                        _ = v29
                                    default:
                                        _ = len(v2)
                                    }
                                    _ = v27
                                } else {
                                    switch v1 {
                                    case 2147483647:
                                        if (-9223372036854775808 != 388) {
                                            fmt.Println(fi19, v1, v3)
                                            var v30 string = "foo bar"
                                            v31 := []int{-159, fi19, (^(fi26))}
                                            _ = v30
                                            _ = v31
                                        }
                                    case -9223372036854775808:
                                        v3 = (v3 << ((625) & 63))
                                    case 53:
                                        v32 := GenS6{F0: (("a" + "foo bar") + "\n\t"), F1: "", F2: (fi19 * -1)}
                                        _ = v32
                                    default:
                                        fmt.Println(v1)
                                    }
                                    v1 += -972
                                    v33 := GenS6{F0: "foo bar", F1: "", F2: 32767}
                                    _ = v33
                                }
                            case 181:
                                var v34 uint8 = 1
                                _ = v34
                            default:
                                if true {
                                    fmt.Println(fi26, fi19, fi26)
                                } else {
                                    var v35 uint16 = 1
                                    _ = v35
                                }
                            }
                            v1 += (fi26 >> ((42) & 63))
                            _ = fi26
                        }
                    }
                    _ = fi19
                }
            case 1:
                v1 += -2147483648
            case 0:
                fmt.Println(v3, v3)
            default:
                fmt.Println(v1)
            }
        case 0:
            v2["a"] = v3
        }
        v1 += (v1 + (v1 | v1))
        v1 -= (-(v1))
    } else {
        v36 := []int{}
        for fi37 := 0; fi37 < 1; fi37++ {
            var v38 string = ""
            v3 += v3
            for fi39 := 0; fi39 < 5; fi39++ {
                var v40 string = "foo bar"
                v41 := GenS3{F0: 2, F1: 298}
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = fi37
            _ = v38
        }
        switch ("" + "") {
        case "":
            for fi42 := 0; fi42 < 7; fi42++ {
                fmt.Println(fi42, fi42)
                _ = fi42
            }
        case "hello":
            fmt.Println(v1, v3)
        case "foo bar":
            for fi43 := 0; fi43 < 15; fi43++ {
                var v44 uint16 = 0
                _ = fi43
                _ = v44
            }
        }
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
