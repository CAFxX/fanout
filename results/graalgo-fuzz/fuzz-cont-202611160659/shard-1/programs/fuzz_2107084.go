package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint; F1 int64; F2 string }
type GenS3 struct { F0 string; F1 int8; F2 string }
type GenS4 struct { F0 int8; F1 int8; F2 uint32 }
func genFn0(p0 bool, p1 bool) uint32 {
    if p0 {
        if (!(4294967295 > 0)) {
            v5 := GenS1{F0: -1, F1: 32767}
            v6 := GenS1{F0: -456, F1: 255}
            _ = v5
            _ = v6
        } else {
            for fi7 := 0; fi7 < 14; fi7++ {
                switch fi7 {
                case 0:
                    for fi8 := 0; fi8 < 19; fi8++ {
                        p0 = false
                        if p0 {
                            if p0 {
                                if ((255 == 127) || (1 > -1)) {
                                    fmt.Println(p1, fi8, p1)
                                    p1 = (!false)
                                    v9 := []int{fi8}
                                    fmt.Println(fi8, fi7)
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(fi7)
                                    default:
                                        switch (fi8 % ((fi8) | 1)) {
                                        case 4294967295:
                                            if len(v9) > 0 {
                                                v9[0] = fi7
                                            }
                                        default:
                                            fmt.Println(fi7, fi7, fi8)
                                        }
                                    }
                                    v9 = append(v9, (-(fi8)))
                                    fmt.Println(fi7, p0, p1)
                                    _ = v9
                                }
                                p1 = (217 == 0)
                            } else {
                                fmt.Println(fi7)
                            }
                            v10 := GenS1{F0: 0, F1: 332}
                            _ = v10
                        }
                        _ = fi8
                    }
                case 960:
                    switch ("" + "hello") {
                    case "foo bar":
                        var v11 bool = ((2 < 65535) || (-1 >= -741))
                        _ = v11
                    default:
                        v12 := GenS2{F0: 255, F1: 1, F2: "a"}
                        _ = v12
                    }
                default:
                    p0 = (212 != 572)
                }
                _ = fi7
            }
        }
        var v13 string = ""
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS3{F0: ((v13 + (v13 + v13)) + (v13 + (v13 + v13))), F1: 104, F2: ("\n\t" + v13)}
            fmt.Println(v13, fi14, fi14)
            if true {
                var v16 uint = 18446744073709551615
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 int16 = 0
                _ = fi18
                _ = v19
            }
            v20 := GenS3{F0: "", F1: -1, F2: "foo bar"}
            _ = fi17
            _ = v20
        }
    }
    if false {
        p1 = (!(1000 < 676))
        v21 := GenS1{F0: 1, F1: 639}
        _ = v21
    }
    fmt.Println(p1)
    p0 = p0
    _ = p0
    _ = p1
    return 65535
}
func genFn1(p0 int8) bool {
    p0 = p0
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 16; fi23++ {
            p0 += p0
            _ = fi23
        }
        _ = fi22
    }
    _ = p0
    return true
}

func main() {
    var v1 int64 = 1
    var v2 uint16 = 42
    var v3 uint = 18446744073709551615
    var v4 string = ("a" + "")
    switch -1 {
    case 128:
        v24 := []int{}
        _ = v24
    case 557:
        v25 := GenS4{F0: 52, F1: 50, F2: 193}
        _ = v25
    default:
        v26 := []int{974}
        _ = v26
    }
    var v27 float32 = 1.5
    for fi28 := 0; fi28 < 13; fi28++ {
        v29 := []int{-1, (fi28 | fi28), (-(fi28))}
        _ = fi28
        _ = v29
    }
    v3 = 641
    v2 = v2
    v3 -= (v3 - (v3 >> ((18446744073709551615) & 63)))
    var v30 bool = (v4 <= "foo bar")
    var v31 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
