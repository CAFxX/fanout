package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int64 }
type GenS2 struct { F0 int16; F1 bool; F2 bool }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint }
func genFn0() uint32 {
    v5 := []int{32768, 128, -884, -9223372036854775808}
    var v6 float32 = 0.1
    _ = v5
    _ = v6
    return 397
}
func genFn1(p0 uint32, p1 uint) int16 {
    var v7 int = int(p1)
    v8 := GenS1{F0: 0, F1: -32768, F2: 42}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -32768
}

func main() {
    var v1 uint8 = 127
    var v2 uint8 = (^(v1))
    var v3 string = "\n\t"
    var v4 uint = 255
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS2{F0: int16(v4), F1: false, F2: (-454 > -1)}
            if (!(!true)) {
                if false {
                    var v12 int8 = 1
                    v12 = v12
                    if ((fi9 | fi10) >= fi9) {
                        v3 = v3
                        v13 := []int{(fi9 - fi10), (-(fi10)), 1, (fi9 & fi10)}
                        _ = v13
                    } else {
                        switch 0 {
                        case 771:
                            var v14 uint32 = uint32(v2)
                            _ = v14
                        case -499:
                            if true {
                                v3 = "\n\t"
                                v15 := GenS1{F0: 277, F1: -32768, F2: 9223372036854775807}
                                _ = v15
                            }
                        }
                    }
                    _ = v12
                }
                var v16 string = ("foo bar" + v3)
                _ = v16
            } else {
                var v17 uint32 = genFn0()
                var v18 uint32 = (v17 >> (((v4 % (((v4 + (^(v4)))) | 1))) & 63))
                switch v3 {
                case "hello":
                    if (true && true) {
                        var v19 string = v3
                        v18 = genFn0()
                        _ = v19
                    } else {
                        if true {
                            var v20 uint = v4
                            if true {
                                switch "a" {
                                case "hello":
                                    v21 := []int{(-(fi10)), fi9, (^(fi9))}
                                    _ = v21
                                case "":
                                    switch 2147483647 {
                                    case 385:
                                        v17 = (v18 | (^(v17)))
                                    }
                                    v22 := []int{(-(fi10)), fi9, -9223372036854775808}
                                    _ = v22
                                }
                                v23 := GenS2{F0: -32768, F1: true, F2: false}
                                switch ("a" + "") {
                                case "foo bar":
                                    var v24 uint16 = 65535
                                    _ = v24
                                case "\n\t":
                                    v18 += (v18 ^ 656)
                                case "hello":
                                    switch (fi10 + 4294967295) {
                                    case 32767:
                                        var v25 string = v3
                                        _ = v25
                                    case 2147483647:
                                        if true {
                                            fmt.Println(v18, v3, v20)
                                            if ((0.1 - (1.5 - (1e10 * ((3.14159 / 3.14159) / -2.25)))) <= (((0.1 / -2.25) * ((0.0 - 0.0) + 0.1)) + 0.1)) {
                                                if (!true) {
                                                    v26 := []int{fi9}
                                                    _ = v26
                                                } else {
                                                    v27 := []int{(^(fi9)), 453, (fi10 % ((0) | 1))}
                                                    var v28 float64 = 0.0
                                                    _ = v27
                                                    _ = v28
                                                }
                                            } else {
                                                fmt.Println(fi9, v17, v2)
                                                v20 += 18446744073709551615
                                                if false {
                                                    v29 := GenS2{F0: 0, F1: (false || (-1 < 656)), F2: true}
                                                    _ = v29
                                                }
                                            }
                                        } else {
                                            var v30 uint16 = 65535
                                            v31 := GenS2{F0: 42, F1: (4294967295 == 831), F2: (true || true)}
                                            _ = v30
                                            _ = v31
                                        }
                                    case -2147483648:
                                        var v32 int = (-(fi10))
                                        _ = v32
                                    }
                                }
                                _ = v23
                            }
                            _ = v20
                        }
                        fmt.Println(fi9, v17)
                    }
                case "a":
                    v33 := GenS1{F0: v18, F1: 1000, F2: int64(v4)}
                    _ = v33
                default:
                    v34 := []int{fi10, -2147483648}
                    _ = v34
                }
                _ = v17
                _ = v18
            }
            var v35 uint16 = 2
            _ = fi10
            _ = v11
            _ = v35
        }
        v36 := []int{fi9, (fi9 - fi9)}
        _ = fi9
        _ = v36
    }
    if true {
        fmt.Println(v1)
        v2 -= 1
        v1 -= 52
    }
    v37 := []int{-2147483648}
    fmt.Println(v2, v1, v2)
    var v38 uint8 = v2
    var v39 uint = 934
    _ = len(v37)
    fmt.Println(v39, v1)
    v40 := GenS1{F0: 1, F1: 32767, F2: -816}
    var v41 int16 = -824
    if false {
        v42 := GenS1{F0: genFn0(), F1: v41, F2: 1}
        var v43 float64 = -2.25
        _ = v42
        _ = v43
    } else {
        var v44 uint32 = 2147483647
        for fi45 := 0; fi45 < 9; fi45++ {
            var v46 uint16 = 0
            fmt.Println(v4)
            if true {
                fmt.Println(v2, v41)
                switch (v3 + ((((v3 + (v3 + v3)) + v3) + ("foo bar" + ((v3 + v3) + (v3 + "")))) + "")) {
                case "\n\t":
                    v37 = append(v37, (^(fi45)))
                case "":
                    v47 := GenS3{F0: (v44 * 1), F1: 424}
                    _ = v47
                default:
                    var v48 uint8 = v38
                    _ = v48
                }
            } else {
                if true {
                    if true {
                        fmt.Println(v44)
                        var v49 int64 = -1
                        if (!true) {
                            v50 := GenS4{F0: v39}
                            v51 := GenS1{F0: v44, F1: v41, F2: v49}
                            v3 = (v3 + v3)
                            _ = v50
                            _ = v51
                        } else {
                            v52 := GenS3{F0: v44, F1: v44}
                            _ = v52
                        }
                        _ = v49
                    }
                    v37 = append(v37, (fi45 + (fi45 + (^(fi45)))))
                } else {
                    if false {
                        v53 := GenS4{F0: (v4 % (((v39 & uint(v2))) | 1))}
                        if true {
                            v54 := GenS2{F0: 1, F1: false, F2: (((v4 >> (((v4 ^ (v39 << ((v39) & 63)))) & 63)) == 255) && (-9223372036854775808 == 42))}
                            v46 = 65535
                            _ = v54
                        } else {
                            var v55 uint32 = (v44 % (((v44 | v44)) | 1))
                            _ = v55
                        }
                        var v56 int32 = 42
                        _ = v53
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 16; fi57++ {
                        v44 -= v44
                        _ = fi57
                    }
                    for fi58 := 0; fi58 < 18; fi58++ {
                        v41 -= (^(v41))
                        _ = len(v37)
                        var v59 float32 = 0.1
                        _ = fi58
                        _ = v59
                    }
                }
                v44 -= (v44 * (v44 | 2))
            }
            _ = fi45
            _ = v46
        }
        _ = v44
    }
    switch -876 {
    case 32768:
        var v60 string = (("\n\t" + ("" + "a")) + (v3 + ("a" + v3)))
        _ = v60
    case 32767:
        if len(v37) > 0 {
            v37[0] = -2147483648
        }
    }
    v2 += (v1 & v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
