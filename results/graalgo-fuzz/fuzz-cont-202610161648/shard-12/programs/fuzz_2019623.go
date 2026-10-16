package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint64; F2 int }
type GenS5 struct { F0 string; F1 uint8; F2 uint64 }
type GenS6 struct { F0 uint64; F1 string }
func genFn0(p0 int, p1 uint, p2 uint8) uint16 {
    if ((p2 << (((p1 % ((p1) | 1))) & 63)) >= p2) {
        var v5 uint8 = p2
        _ = v5
    }
    var v6 uint8 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1(p0 uint32, p1 uint, p2 string) bool {
    switch p2 {
    case "\n\t":
        switch int(p1) {
        case -109:
            p0 = p0
        case 32768:
            if (true && true) {
                p2 = p2
                var v7 uint16 = 291
                _ = v7
            } else {
                var v8 uint8 = 255
                var v9 int32 = 1
                var v10 int64 = 42
                _ = v10
                _ = v8
                _ = v9
            }
        case -1:
            var v11 int = -738
            _ = v11
        }
    }
    if false {
        var v12 int = int(p0)
        if false {
            var v13 uint = 42
            v14 := []int{v12, 1, (v12 << (((v13 << (((^(v13))) & 63))) & 63)), -69}
            _ = v13
            _ = v14
        }
        switch p2 {
        case "hello":
            if true {
                for fi15 := 0; fi15 < 1; fi15++ {
                    v12 = v12
                    _ = fi15
                }
            } else {
                if true {
                    var v16 int32 = -2147483648
                    fmt.Println(p0)
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v18 := GenS1{F0: 18446744073709551615, F1: (42 == 0)}
                        v19 := []int{v12, v12, -761, (fi17 & fi17)}
                        for fi20 := 0; fi20 < 2; fi20++ {
                            p0 = (p0 >> (((p1 >> (((p1 & 421)) & 63))) & 63))
                            _ = fi20
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
                var v21 uint8 = 2
                if ((0.0 / 0.1) < float64(-2.25)) {
                    if true {
                        v22 := GenS2{F0: 1, F1: v21}
                        _ = v22
                    } else {
                        v23 := []int{(v12 * v12), 451, v12, (v12 % (((v12 << ((4294967295) & 63))) | 1))}
                        switch p2 {
                        case "hello":
                            v24 := GenS1{F0: 551, F1: (p1 == 2)}
                            _ = v24
                        default:
                            v25 := GenS1{F0: 2, F1: (uint32(v21) > 65535)}
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 2; fi26++ {
                            v27 := GenS1{F0: 211, F1: (false || (((p2 + ((p2 + p2) + p2)) + p2) == "a"))}
                            if true {
                                var v28 int32 = 2147483647
                                if (0 <= 42) {
                                    var v29 int16 = 1000
                                    var v30 int = (v12 ^ (v12 * fi26))
                                    v31 := GenS3{F0: false}
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                                _ = v28
                            } else {
                                if (!true) {
                                    v32 := GenS4{F0: (p2 + ("foo bar" + (p2 + ((p2 + "\n\t") + (((("\n\t" + p2) + ("\n\t" + (p2 + "hello"))) + ((p2 + p2) + p2)) + (p2 + p2)))))), F1: 1, F2: (fi26 / (((^(v12))) | 1))}
                                    _ = v32
                                }
                            }
                            var v33 int16 = -1
                            _ = fi26
                            _ = v27
                            _ = v33
                        }
                        _ = v23
                    }
                    if false {
                        v34 := []int{(v12 * int(p1)), v12}
                        p2 = ("a" + p2)
                        if len(v34) > 0 {
                            v34[0] = (v12 << ((p1) & 63))
                        }
                        _ = v34
                    } else {
                        if (p1 != p1) {
                            v12 = (v12 & v12)
                        }
                        v35 := GenS5{F0: "", F1: (v21 * v21), F2: 182}
                        _ = v35
                    }
                }
                _ = v21
            }
        default:
            p2 = "hello"
        }
        _ = v12
    } else {
        switch p2 {
        case "hello":
            var v36 int16 = 32767
            _ = v36
        }
        for fi37 := 0; fi37 < 14; fi37++ {
            var v38 int32 = -480
            _ = fi37
            _ = v38
        }
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return (105 > 255)
}

func main() {
    var v1 uint8 = 164
    var v2 uint64 = 18446744073709551615
    var v3 int16 = 0
    var v4 bool = true
    for fi39 := 0; fi39 < 5; fi39++ {
        for fi40 := 0; fi40 < 5; fi40++ {
            var v41 uint8 = v1
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    v4 = ((v2 | v2) < (^(v2)))
    var v42 float32 = 3.14159
    v1 = v1
    if (!v4) {
        var v43 uint16 = 768
        fmt.Println(v43, v42, v4)
        for fi44 := 0; fi44 < 14; fi44++ {
            v3 = v3
            _ = fi44
        }
        _ = v43
    }
    if v4 {
        v1 = v1
        v3 = v3
        v45 := []int{1, 1}
        _ = v45
    } else {
        var v46 int64 = 722
        for fi47 := 0; fi47 < 13; fi47++ {
            switch (^(fi47)) {
            case 42:
                switch 2147483647 {
                case 1:
                    var v48 string = "a"
                    _ = v48
                default:
                    v4 = (!v4)
                }
            case 32767:
                v49 := []int{(fi47 * (fi47 | fi47)), (fi47 | fi47)}
                _ = v49
            }
            var v50 uint64 = 18446744073709551615
            for fi51 := 0; fi51 < 7; fi51++ {
                var v52 uint16 = 1
                var v53 float32 = 3.14159
                _ = fi51
                _ = v52
                _ = v53
            }
            _ = fi47
            _ = v50
        }
        _ = v46
    }
    v4 = (true || (!(42 <= 2)))
    if ((!((v1 | (v1 / ((v1) | 1))) > 42)) || v4) {
        var v54 float32 = (v42 / 1.0)
        _ = v54
    }
    switch "\n\t" {
    case "":
        v55 := GenS6{F0: v2, F1: "foo bar"}
        _ = v55
    default:
        var v56 uint = 42
        _ = v56
    }
    v57 := []int{}
    for fi58 := 0; fi58 < 9; fi58++ {
        v57 = append(v57, fi58)
        _ = fi58
    }
    v3 = (^(v3))
    var v59 bool = v4
    fmt.Println(v3, v1)
    v42 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v57
    _ = v59
}
