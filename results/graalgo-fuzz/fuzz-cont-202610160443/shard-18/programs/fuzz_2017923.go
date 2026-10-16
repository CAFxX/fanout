package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int; F2 int32 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 uint8; F1 uint64; F2 int }

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = 1000
    var v3 string = "\n\t"
    v4 := make(map[int]int)
    var v5 uint = 18446744073709551615
    v6 := GenS1{F0: (!false), F1: 153, F2: 42}
    v3 = ""
    var v7 int64 = 1
    switch 9223372036854775807 {
    case -598:
        var v8 bool = true
        _ = v8
    case -9223372036854775808:
        v9 := GenS2{F0: (v3 + v3)}
        _ = v9
    default:
        for fi10 := 0; fi10 < 2; fi10++ {
            v2 += 653
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 string = "hello"
                if false {
                    v13 := GenS1{F0: ((true || true) || true), F1: (fi11 - fi10), F2: 2147483647}
                    if true {
                        fmt.Println(v12, v12)
                        if true {
                            if true {
                                v12 = ""
                            }
                        } else {
                            var v14 int8 = 42
                            if (255 > 62) {
                                fmt.Println(v2)
                            }
                            _ = v14
                        }
                    } else {
                        var v15 bool = false
                        v16 := GenS3{F0: 528, F1: fi10, F2: 973}
                        fmt.Println(fi11, v7)
                        _ = v15
                        _ = v16
                    }
                    var v17 uint64 = 42
                    _ = v13
                    _ = v17
                }
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    }
    var v18 uint16 = 255
    switch ((v3 + "a") + "hello") {
    case "hello":
        fmt.Println(v2, v2)
    case "\n\t":
        var v19 uint8 = 0
        _ = v19
    case "a":
        switch 865 {
        case 824:
            var v20 int = 873
            _ = v20
        case -2147483648:
            for fi21 := 0; fi21 < 19; fi21++ {
                switch fi21 {
                case 408:
                    v2 = (v2 >> ((v5) & 63))
                case -9223372036854775808:
                    var v22 uint64 = 1
                    _ = v22
                case -942:
                    if false {
                        for fi23 := 0; fi23 < 12; fi23++ {
                            var v24 string = (v3 + "hello")
                            _ = fi23
                            _ = v24
                        }
                        if false {
                            v7 -= 42
                            v25 := GenS4{F0: (0.0 >= (1e10 - 1.5)), F1: (v18 & 0)}
                            v26 := GenS5{F0: 128, F1: 42}
                            _ = v25
                            _ = v26
                        } else {
                            if true {
                                fmt.Println(v18, v18)
                                var v27 uint64 = 18446744073709551615
                                v18 += v18
                                _ = v27
                            } else {
                                for fi28 := 0; fi28 < 19; fi28++ {
                                    var v29 uint = (^(v5))
                                    _ = fi28
                                    _ = v29
                                }
                                fmt.Println(v18, v18)
                                v30 := v4[(fi21 >> (((v5 ^ 1)) & 63))]
                                _ = v30
                            }
                            var v31 uint64 = 18446744073709551615
                            fmt.Println(fi21, v3)
                            _ = v31
                        }
                        for fi32 := 0; fi32 < 5; fi32++ {
                            v33 := []int{fi32, fi32, fi32}
                            _ = fi32
                            _ = v33
                        }
                    } else {
                        fmt.Println(fi21, v3, v18)
                    }
                default:
                    switch v3 {
                    case "\n\t":
                        switch 893 {
                        case 551:
                            for fi34 := 0; fi34 < 6; fi34++ {
                                if ((-(fi34)) < fi34) {
                                    var v35 uint16 = v18
                                    v7 = (v7 + 1)
                                    _ = v35
                                } else {
                                    if false {
                                        if (-1 > 1) {
                                            v36 := []int{32767}
                                            fmt.Println(fi21, v5, fi34)
                                            v2 -= (-(v2))
                                            _ = v36
                                        }
                                        delete(v1, (fi34 << (((v5 + uint(v18))) & 63)))
                                        v18 = 32767
                                    }
                                }
                                var v37 uint16 = 725
                                v37 += (v18 - (v37 << ((0) & 63)))
                                _ = fi34
                                _ = v37
                            }
                        case 0:
                            v38 := GenS4{F0: false, F1: (v18 % ((65535) | 1))}
                            _ = v38
                        case 32767:
                            v1[1] = (1 >= -69)
                        }
                    case "foo bar":
                        v39 := GenS6{F0: 0, F1: 4294967295, F2: 32768}
                        _ = v39
                    case "hello":
                        v40 := []int{}
                        _ = v40
                    }
                }
                _ = fi21
            }
        default:
            fmt.Println(v2, v2)
        }
    }
    v18 = v18
    if true {
        v3 = ("\n\t" + (((v3 + ("hello" + v3)) + (v3 + ((("" + "hello") + ("\n\t" + (v3 + "foo bar"))) + "\n\t"))) + v3))
        for fi41 := 0; fi41 < 20; fi41++ {
            fmt.Println(fi41, v18)
            switch v3 {
            case "\n\t":
                v4[(fi41 ^ fi41)] = 2147483647
            }
            for fi42 := 0; fi42 < 15; fi42++ {
                var v43 bool = (false || (!false))
                _ = fi42
                _ = v43
            }
            _ = fi41
        }
        var v44 uint64 = uint64(v18)
        _ = v44
    }
    var v45 uint = v5
    var v46 int32 = -765
    v47 := []int{32767, 9223372036854775807, 142}
    v48 := GenS4{F0: false, F1: v18}
    v49 := GenS3{F0: 0, F1: 9223372036854775807, F2: (v46 << ((731) & 63))}
    v50 := GenS1{F0: true, F1: 4294967295, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
