package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint64; F1 int; F2 uint64 }
func genFn0(p0 uint16) bool {
    var v7 uint8 = 2
    var v8 int = -2147483648
    v9 := []int{int(p0), (v8 + (v8 - 152)), (v8 >> ((858) & 63)), v8}
    fmt.Println(v7, v7)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 int, p1 int64, p2 int16) int16 {
    if true {
        v10 := GenS1{F0: 42, F1: "hello"}
        var v11 uint16 = 2
        _ = v10
        _ = v11
    }
    var v12 string = "hello"
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p2
}

func main() {
    var v1 float32 = 1e10
    var v2 int8 = -128
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    var v5 uint8 = 127
    var v6 uint16 = 65535
    v2 = 42
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := []int{fi13, (fi13 ^ (fi13 - -738)), fi13}
        fmt.Println(v2, v1)
        var v15 uint16 = 718
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 bool = (!true)
    switch -958 {
    case 32767:
        fmt.Println(v1, v1)
    case -802:
        switch "\n\t" {
        case "\n\t":
            fmt.Println(v16, v6)
        case "a":
            var v17 uint32 = 42
            _ = v17
        }
    case -9223372036854775808:
        v6 = (^(v6))
    default:
        if true {
            if false {
                v4[2147483647] = "foo bar"
            }
            v18 := []int{-321, 184, -18, 128}
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            _ = v18
        } else {
            _ = len(v3)
            var v19 uint32 = 490
            var v20 uint16 = 0
            _ = v19
            _ = v20
        }
    }
    v16 = (512 < 4294967295)
    var v21 string = "foo bar"
    var v22 uint = 1
    if v16 {
        for fi23 := 0; fi23 < 8; fi23++ {
            for fi24 := 0; fi24 < 9; fi24++ {
                v25 := []int{}
                _ = fi24
                _ = v25
            }
            if ((true && false) && (127 > (v5 & (v5 + v5)))) {
                v21 = v21
            }
            _ = fi23
        }
    } else {
        if v16 {
            for fi26 := 0; fi26 < 18; fi26++ {
                v27 := GenS1{F0: 65535, F1: ""}
                _ = fi26
                _ = v27
            }
            if true {
                var v28 bool = true
                _ = v28
            }
            v21 = (v21 + "hello")
        }
        for fi29 := 0; fi29 < 18; fi29++ {
            for fi30 := 0; fi30 < 10; fi30++ {
                if v16 {
                    v31 := GenS2{F0: (v5 * v5)}
                    v32 := []int{fi30, fi29, fi29}
                    switch v21 {
                    case "a":
                        v16 = v16
                    case "hello":
                        var v33 float32 = (v1 + (v1 * (v1 + v1)))
                        _ = v33
                    default:
                        v5 = 42
                    }
                    _ = v31
                    _ = v32
                }
                _ = fi30
            }
            if (!(!v16)) {
                v34 := GenS3{F0: (v6 | (v6 << ((255) & 63))), F1: 381, F2: 0}
                v35 := GenS1{F0: 4294967295, F1: ((("hello" + (v21 + "a")) + v21) + v21)}
                v16 = (v16 && ((^(v6)) != (v6 * (v6 * 0))))
                _ = v34
                _ = v35
            }
            _ = fi29
        }
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        fmt.Println(fi36, fi36)
        v16 = (!v16)
        switch fi36 {
        case 42:
            fmt.Println(v6)
        case 127:
            var v37 uint64 = 810
            _ = v37
        case -878:
            v6 += 1
        }
        _ = fi36
    }
    if ((515 <= 2) && ((v21 + (v21 + (v21 + "\n\t"))) == "foo bar")) {
        for fi38 := 0; fi38 < 2; fi38++ {
            switch (v21 + v21) {
            case "\n\t":
                v2 = int8(fi38)
            case "a":
                if v16 {
                    v5 = v5
                } else {
                    v39 := []int{}
                    if v16 {
                        v22 = (^(v22))
                    }
                    for fi40 := 0; fi40 < 14; fi40++ {
                        var v41 uint = 271
                        _ = fi40
                        _ = v41
                    }
                    _ = v39
                }
            case "foo bar":
                fmt.Println(v16, v1)
            default:
                if v16 {
                    var v42 int64 = 765
                    if false {
                        if (!true) {
                            v16 = v16
                        } else {
                            for fi43 := 0; fi43 < 3; fi43++ {
                                switch (fi38 * (fi43 | (fi43 ^ fi43))) {
                                case 42:
                                    if (genFn1(1, (v42 % (((-(v42))) | 1)), genFn1(fi43, (^(v42)), 0)) >= -1) {
                                        var v44 uint = (^(v22))
                                        var v45 string = (v21 + v21)
                                        fmt.Println(v5, v45)
                                        _ = v44
                                        _ = v45
                                    } else {
                                        switch v21 {
                                        case "\n\t":
                                            switch fi43 {
                                            case 2147483647:
                                                v46 := GenS2{F0: 137}
                                                _ = v46
                                            default:
                                                v47 := GenS1{F0: uint32(v22), F1: ("a" + v21)}
                                                _ = v47
                                            }
                                        case "a":
                                            var v48 int = fi43
                                            _ = v48
                                        case "":
                                            v49 := GenS4{F0: -283, F1: 255, F2: 42}
                                            _ = v49
                                        }
                                        var v50 uint16 = v6
                                        _ = v50
                                    }
                                case -277:
                                    switch fi38 {
                                    case 42:
                                        fmt.Println(v21, v6)
                                    }
                                case -187:
                                    v51 := []int{int(v42), (fi43 >> (((v22 - (v22 >> (((v22 << ((v22) & 63))) & 63)))) & 63)), (fi43 << (((v22 + v22)) & 63)), 176}
                                    _ = v51
                                default:
                                    switch v21 {
                                    case "a":
                                        var v52 uint16 = uint16(v22)
                                        _ = v52
                                    case "":
                                        v53 := GenS3{F0: 63, F1: v42, F2: 1}
                                        _ = v53
                                    default:
                                        fmt.Println(v6, fi43, fi43)
                                        v54 := GenS5{F0: 284, F1: fi43, F2: 4294967295}
                                        _ = v54
                                    }
                                }
                                v42 -= -1
                                v55 := GenS4{F0: 1, F1: v6, F2: 644}
                                _ = fi43
                                _ = v55
                            }
                        }
                    } else {
                        v56 := []int{fi38, fi38}
                        var v57 uint16 = (v6 & (v6 / (((v6 << ((59) & 63))) | 1)))
                        fmt.Println(v16)
                        _ = v56
                        _ = v57
                    }
                    _ = v42
                }
            }
            var v58 bool = (v16 || (v16 || v16))
            _ = fi38
            _ = v58
        }
        fmt.Println(v2)
        fmt.Println(v22, v5)
    }
    v59 := GenS1{F0: 109, F1: v21}
    var v60 int64 = -873
    v4[4294967295] = ((v21 + (v21 + (v21 + v21))) + (v21 + ((((v21 + v21) + "") + "") + "foo bar")))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
}
