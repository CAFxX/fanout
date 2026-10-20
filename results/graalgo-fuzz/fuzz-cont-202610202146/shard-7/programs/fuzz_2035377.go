package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint; F1 int; F2 uint32 }
type GenS3 struct { F0 string; F1 uint8; F2 uint8 }
func genFn0(p0 int64) bool {
    var v7 bool = ((-1 < (^(p0))) || false)
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint8 = 1
        for fi10 := 0; fi10 < 20; fi10++ {
            v7 = false
            _ = fi10
        }
        v11 := GenS1{F0: 0, F1: (^(v9))}
        _ = fi8
        _ = v11
        _ = v9
    }
    var v12 int = 1
    var v13 int8 = 19
    v14 := []int{-632, v12, v12}
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    _ = v7
    return (v12 != v12)
}
func genFn1(p0 uint) int32 {
    for fi15 := 0; fi15 < 14; fi15++ {
        if true {
            p0 = p0
        }
        var v16 uint16 = 255
        var v17 bool = true
        _ = fi15
        _ = v16
        _ = v17
    }
    if (42 == (p0 | 835)) {
        p0 -= (p0 / (((^(p0))) | 1))
    } else {
        fmt.Println(p0)
        var v18 int64 = 1
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        p0 += p0
        v20 := []int{fi19, 32767, 2147483647}
        var v21 float32 = -2.25
        _ = fi19
        _ = v20
        _ = v21
    }
    p0 = p0
    _ = p0
    return -2147483648
}

func main() {
    v1 := []string{"a"}
    v2 := []int{}
    var v3 int32 = 1
    v4 := make(map[uint32]uint32)
    var v5 int = int(v3)
    v6 := make(map[uint32]string)
    v3 = genFn1(4294967295)
    if (!(true && true)) {
        var v22 bool = false
        _ = v22
    } else {
        v23 := v4[42]
        _ = len(v1)
        for fi24 := 0; fi24 < 14; fi24++ {
            v25 := GenS1{F0: 42, F1: 0}
            for fi26 := 0; fi26 < 7; fi26++ {
                v3 = int32(v23)
                switch "a" {
                case "hello":
                    v23 -= (^(v23))
                case "\n\t":
                    switch (("a" + "a") + "hello") {
                    case "a":
                        v27 := GenS2{F0: 45, F1: (^(fi24)), F2: 42}
                        _ = v27
                    }
                    var v28 uint16 = 2
                    v3 -= genFn1(42)
                    _ = v28
                case "a":
                    v23 += v23
                default:
                    v5 = 128
                }
                _ = fi26
            }
            for fi29 := 0; fi29 < 16; fi29++ {
                var v30 bool = (false || false)
                v23 = (v23 << ((4294967295) & 63))
                _ = fi29
                _ = v30
            }
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    v3 += v3
    v6[2147483648] = ""
    v3 = (-(v3))
    v5 = (v5 ^ (v5 % ((v5) | 1)))
    delete(v4, 860)
    fmt.Println(v3, v3)
    v3 = (-(v3))
    v2 = append(v2, 0)
    if true {
        for fi31 := 0; fi31 < 11; fi31++ {
            for fi32 := 0; fi32 < 5; fi32++ {
                switch fi32 {
                case 828:
                    v5 = 265
                case 4294967295:
                    var v33 float32 = 0.0
                    _ = v33
                default:
                    if true {
                        switch (v5 % ((fi31) | 1)) {
                        case -9223372036854775808:
                            switch fi32 {
                            case -452:
                                var v34 uint8 = 111
                                _ = v34
                            default:
                                v5 += (^(fi31))
                            }
                            switch 609 {
                            case 914:
                                if true {
                                    _ = len(v2)
                                }
                            case 73:
                                var v35 string = "foo bar"
                                _ = v35
                            case 507:
                                v3 -= 0
                            default:
                                v36 := GenS3{F0: "", F1: 150, F2: 166}
                                _ = v36
                            }
                            if false {
                                var v37 int16 = 752
                                if false {
                                    switch "" {
                                    case "":
                                        var v38 int8 = 100
                                        _ = v38
                                    case "\n\t":
                                        if true {
                                            if true {
                                                var v39 float64 = (0.1 * ((0.1 * 3.14159) + -2.25))
                                                fmt.Println(v37, v37)
                                                _ = v39
                                            }
                                            v5 -= (v5 - (fi31 / (((fi31 >> ((4294967295) & 63))) | 1)))
                                            v3 -= v3
                                        }
                                    case "a":
                                        if false {
                                            switch (v5 + (fi31 ^ 0)) {
                                            case 42:
                                                var v40 bool = ((v3 % ((374) | 1)) <= int32(v37))
                                                _ = v40
                                            case -452:
                                                delete(v4, 42)
                                            default:
                                                var v41 uint8 = 25
                                                _ = v41
                                            }
                                            v5 = (-(fi32))
                                        } else {
                                            if true {
                                                fmt.Println(v3)
                                                var v42 uint64 = 18446744073709551615
                                                if (!(true || false)) {
                                                    if false {
                                                        fmt.Println(v5, v42, v3)
                                                        var v43 string = (("hello" + "") + ("hello" + "\n\t"))
                                                        fmt.Println(v42, fi32, fi32)
                                                        var v44 int64 = int64(v42)
                                                        v45 := GenS1{F0: 0, F1: 1}
                                                        _ = len(v1)
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    var v46 int64 = -10
                                                    v46 = (v46 - -218)
                                                    _ = v46
                                                } else {
                                                    var v47 bool = ("\n\t" != "\n\t")
                                                    v42 = 4294967295
                                                    var v48 uint8 = 1
                                                    _ = v47
                                                    _ = v48
                                                }
                                                _ = v42
                                            } else {
                                                var v49 bool = (((false || ((v37 > v37) || true)) && false) || (65535 < 1))
                                                var v50 string = "hello"
                                                _ = v49
                                                _ = v50
                                            }
                                        }
                                    }
                                }
                                var v51 int = (fi31 >> ((4294967295) & 63))
                                _ = v37
                                _ = v51
                            } else {
                                var v52 int8 = 127
                                _ = v52
                            }
                        }
                        _ = len(v2)
                        v5 += (fi31 >> ((255) & 63))
                    } else {
                        fmt.Println(fi31)
                    }
                }
                _ = fi32
            }
            v5 += 1
            _ = fi31
        }
        for fi53 := 0; fi53 < 13; fi53++ {
            switch 794 {
            case 825:
                fmt.Println(v5, fi53)
            case 127:
                v3 = (v3 + genFn1(0))
            case 42:
                switch fi53 {
                case 42:
                    v5 = 504
                case 32768:
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v3)
                    case "":
                        v5 -= 4294967295
                    case "hello":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                default:
                    var v54 uint32 = 1
                    _ = v54
                }
            }
            _ = fi53
        }
    } else {
        v3 -= (v3 >> ((2) & 63))
        v5 = (v5 + (v5 * 32768))
    }
    fmt.Println(v3, v3)
    v3 -= v3
    switch ("" + ((((("hello" + "hello") + ("foo bar" + (("a" + "foo bar") + "a"))) + "foo bar") + "hello") + "\n\t")) {
    case "\n\t":
        fmt.Println(v5)
    }
    var v55 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
