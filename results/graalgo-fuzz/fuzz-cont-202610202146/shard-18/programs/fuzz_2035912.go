package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 int32; F1 uint64 }

func main() {
    var v1 int8 = -1
    var v2 string = "\n\t"
    v3 := make(map[int]bool)
    var v4 bool = true
    v5 := GenS1{F0: "", F1: 9223372036854775807}
    var v6 uint64 = 4294967295
    switch -1 {
    case 42:
        v7 := GenS1{F0: "hello", F1: 1}
        _ = v7
    default:
        var v8 uint32 = 1
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS2{F0: (!(!(1 >= 298)))}
        v6 = 635
        switch fi9 {
        case -473:
            if false {
                v11 := []int{(fi9 - 42)}
                _ = v11
            } else {
                fmt.Println(v4, v2)
                v12 := []int{(fi9 * 610)}
                _ = v12
            }
        case -9223372036854775808:
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 int32 = 177
                fmt.Println(fi13)
                _ = fi13
                _ = v14
            }
        default:
            var v15 int = -441
            _ = v15
        }
        _ = fi9
        _ = v10
    }
    switch v2 {
    case "hello":
        if false {
            for fi16 := 0; fi16 < 9; fi16++ {
                fmt.Println(v2, v6, v4)
                var v17 string = v2
                var v18 uint32 = 300
                _ = fi16
                _ = v17
                _ = v18
            }
        } else {
            v2 = "\n\t"
            v19 := []int{32768}
            var v20 int8 = -1
            _ = v19
            _ = v20
        }
    case "\n\t":
        for fi21 := 0; fi21 < 17; fi21++ {
            if (1 > 129) {
                v3[(fi21 << ((18446744073709551615) & 63))] = false
                var v22 int8 = v1
                if v4 {
                    v1 = (v1 * (v22 % ((1) | 1)))
                    if false {
                        v23 := GenS3{F0: 42}
                        _ = v23
                    }
                    var v24 int32 = -2147483648
                    _ = v24
                } else {
                    v25 := GenS3{F0: 740}
                    v26 := []int{32767, (fi21 * fi21), fi21}
                    v4 = false
                    _ = v25
                    _ = v26
                }
                _ = v22
            } else {
                v27 := GenS4{F0: (v1 & 127), F1: v6, F2: (v6 & v6)}
                _ = v27
            }
            _ = fi21
        }
    }
    var v28 int8 = 114
    var v29 float32 = -2.25
    if (v4 && true) {
        v28 = (-(v28))
    }
    var v30 uint16 = 392
    var v31 uint16 = (^(v30))
    var v32 uint = 347
    switch 870 {
    case -827:
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 uint64 = (^(v6))
            var v35 int = 1
            if true {
                for fi36 := 0; fi36 < 7; fi36++ {
                    switch v2 {
                    case "foo bar":
                        switch v35 {
                        case 1:
                            if v4 {
                                if (true && false) {
                                    switch (fi33 - v35) {
                                    case 32768:
                                        v28 = (v28 * v28)
                                    default:
                                        v37 := []int{}
                                        _ = v37
                                    }
                                }
                                var v38 bool = v4
                                _ = v38
                            }
                        default:
                            if v4 {
                                if (2147483648 <= 2) {
                                    v31 = (v30 + (v30 % (((v31 << ((0) & 63))) | 1)))
                                } else {
                                    v39 := GenS5{F0: fi33, F1: -353}
                                    if true {
                                        var v40 int8 = 42
                                        var v41 float64 = (0.0 / (1e10 + (-2.25 - ((0.1 / 0.1) - (1.5 / 0.1)))))
                                        _ = v40
                                        _ = v41
                                    } else {
                                        var v42 string = "hello"
                                        if v4 {
                                            v43 := []int{v35, v35, -610}
                                            _ = v43
                                        }
                                        v44 := GenS4{F0: v1, F1: (v34 + (v6 - 1)), F2: v6}
                                        _ = v42
                                        _ = v44
                                    }
                                    v4 = (2147483647 != 2147483647)
                                    _ = v39
                                }
                                v45 := GenS6{F0: 1, F1: -411}
                                var v46 int8 = -128
                                if (!(1000 != 1)) {
                                    switch ("hello" + v2) {
                                    case "foo bar":
                                        if v4 {
                                            var v47 uint8 = 114
                                            v34 += (v34 >> ((18446744073709551615) & 63))
                                            _ = v47
                                        } else {
                                            if v4 {
                                                var v48 uint32 = 129
                                                v49 := GenS6{F0: 127, F1: 9223372036854775807}
                                                var v50 int32 = 1
                                                _ = v48
                                                _ = v49
                                                _ = v50
                                            }
                                            v31 -= (v31 * (v31 % ((v31) | 1)))
                                        }
                                    }
                                }
                                _ = v45
                                _ = v46
                            }
                        }
                    default:
                        v2 = v2
                        v29 -= v29
                        v30 = (^(v31))
                    }
                    fmt.Println(v35, v34, v35)
                    var v51 uint16 = v30
                    _ = fi36
                    _ = v51
                }
                var v52 float64 = 0.0
                if ((v32 | (v32 * (v32 << ((v32) & 63)))) != (v32 & 2)) {
                    var v53 int32 = -1
                    _ = v53
                } else {
                    for fi54 := 0; fi54 < 6; fi54++ {
                        v28 = v28
                        _ = fi54
                    }
                }
                _ = v52
            }
            _ = fi33
            _ = v34
            _ = v35
        }
    default:
        var v55 int16 = 1
        _ = v55
    }
    for fi56 := 0; fi56 < 3; fi56++ {
        v31 = (v31 + 2)
        for fi57 := 0; fi57 < 15; fi57++ {
            var v58 string = "\n\t"
            v59 := GenS1{F0: (v58 + v2), F1: fi56}
            _ = fi57
            _ = v58
            _ = v59
        }
        _ = fi56
    }
    var v60 uint = 4294967295
    switch ((("\n\t" + "hello") + v2) + "\n\t") {
    case "foo bar":
        var v61 float64 = (1e10 + float64(3.14159))
        _ = v61
    case "\n\t":
        v62 := GenS7{F0: 1, F1: v6}
        _ = v62
    }
    var v63 int16 = -831
    var v64 uint16 = (v31 / ((255) | 1))
    var v65 uint32 = 639
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v63
    _ = v64
    _ = v65
}
