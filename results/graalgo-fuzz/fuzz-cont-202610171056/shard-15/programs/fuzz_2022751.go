package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 string; F2 uint }
type GenS5 struct { F0 int; F1 bool }
func genFn0(p0 uint8) uint16 {
    v6 := GenS1{F0: 0}
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(fi8, fi8)
        v9 := GenS1{F0: 0}
        _ = fi8
        _ = v9
    }
    v10 := []int{32768, 0, 1}
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 int8 = 100
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 string, p1 int32, p2 bool) string {
    p0 = (p0 + p0)
    fmt.Println(p2, p0, p2)
    if (32767 > -1) {
        if (p2 && (!p2)) {
            var v13 int8 = 42
            v13 = (v13 - -1)
            var v14 uint = 18446744073709551615
            _ = v13
            _ = v14
        }
    } else {
        var v15 int16 = 1000
        p1 = (-(p1))
        for fi16 := 0; fi16 < 4; fi16++ {
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 float64 = (3.14159 / 1e10)
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        _ = v15
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 uint8 = 255
        fmt.Println(fi19, p1, p1)
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := GenS1{F0: 2147483647}
            p0 = (p0 + (p0 + "hello"))
            if p2 {
                if true {
                    v23 := []int{127, fi21}
                    _ = v23
                } else {
                    v24 := GenS1{F0: (p1 & int32(v20))}
                    if true {
                        var v25 int16 = 1
                        if false {
                            if p2 {
                                if (0 != 23) {
                                    if p2 {
                                        p1 -= 0
                                        if p2 {
                                            if false {
                                                switch "a" {
                                                case "":
                                                    v26 := GenS2{F0: 1, F1: 857}
                                                    _ = v26
                                                }
                                            } else {
                                                v27 := []int{}
                                                switch "" {
                                                case "":
                                                    p1 -= (-(p1))
                                                default:
                                                    fmt.Println(p1, v20, v25)
                                                }
                                                v28 := GenS3{F0: v20}
                                                v25 -= 32767
                                                if ((!false) || (!true)) {
                                                    _ = len(v27)
                                                } else {
                                                    v20 += v20
                                                }
                                                p1 = 1
                                                var v29 uint16 = 918
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                            }
                                            v30 := []int{fi19}
                                            fmt.Println(v20, fi19, v25)
                                            if (v20 >= 127) {
                                                p1 -= (p1 | p1)
                                                var v31 int = -2147483648
                                                if p2 {
                                                    if p2 {
                                                        var v32 int = (v31 | fi21)
                                                        _ = len(v30)
                                                        _ = v32
                                                    } else {
                                                        if p2 {
                                                            if len(v30) > 0 {
                                                                v30[0] = (fi21 + fi21)
                                                            }
                                                            switch (p0 + p0) {
                                                            case "":
                                                                v30 = append(v30, (fi21 & v31))
                                                            case "a":
                                                                switch (fi21 | (fi21 ^ 962)) {
                                                                case 1:
                                                                    p0 = p0
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if len(v30) >= 2 {
                                                        v30 = v30[1:len(v30)]
                                                    }
                                                }
                                                _ = v31
                                            } else {
                                                var v33 int64 = 1
                                                if (1 != (v20 % (((^(v20))) | 1))) {
                                                    var v34 uint32 = 65535
                                                    fmt.Println(v34)
                                                    _ = v34
                                                } else {
                                                    var v35 uint16 = 0
                                                    switch p0 {
                                                    case "foo bar":
                                                        switch p0 {
                                                        case "a":
                                                            v33 -= 1
                                                        }
                                                    }
                                                    p2 = true
                                                    var v36 int16 = (v25 >> ((2) & 63))
                                                    fmt.Println(v33, v25)
                                                    _ = v35
                                                    _ = v36
                                                }
                                                _ = v33
                                            }
                                            _ = v30
                                        } else {
                                            if (1 == 152) {
                                                var v37 uint8 = 218
                                                switch 32767 {
                                                case 128:
                                                    fmt.Println(v37)
                                                }
                                                _ = v37
                                            }
                                            var v38 int = -149
                                            fmt.Println(v20, fi19)
                                            _ = v38
                                        }
                                        var v39 int64 = -790
                                        _ = v39
                                    } else {
                                        fmt.Println(p2)
                                        var v40 uint32 = 229
                                        _ = v40
                                    }
                                }
                                p1 -= p1
                            }
                        }
                        _ = v25
                    }
                    p2 = (!p2)
                    var v41 int8 = int8(fi19)
                    _ = v24
                    _ = v41
                }
            }
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    return ("hello" + p0)
}

func main() {
    var v1 string = "foo bar"
    var v2 int32 = -538
    var v3 bool = (-1 == 1)
    var v4 float64 = 0.0
    var v5 int32 = -403
    for fi42 := 0; fi42 < 1; fi42++ {
        var v43 int16 = 32767
        switch (fi42 & fi42) {
        case 359:
            switch v1 {
            case "hello":
                var v44 float32 = (float32(0.0) / float32(((v4 + ((0.1 + 0.0) / 1.0)) - float64((3.14159 / (float32(1.5) + (0.1 + -2.25)))))))
                _ = v44
            case "\n\t":
                v45 := []int{(-(fi42)), fi42, fi42}
                _ = v45
            default:
                v46 := GenS3{F0: 128}
                _ = v46
            }
        default:
            var v47 uint = 0
            _ = v47
        }
        var v48 uint64 = 192
        _ = fi42
        _ = v43
        _ = v48
    }
    var v49 uint16 = 110
    var v50 uint64 = 11
    var v51 uint32 = 2147483648
    for fi52 := 0; fi52 < 7; fi52++ {
        v53 := GenS4{F0: v49, F1: v1, F2: 42}
        for fi54 := 0; fi54 < 20; fi54++ {
            var v55 int16 = 0
            var v56 bool = true
            switch (fi54 - (^(fi52))) {
            case 127:
                v57 := GenS5{F0: fi54, F1: (true || v3)}
                _ = v57
            case -9223372036854775808:
                var v58 int16 = (-(v55))
                _ = v58
            default:
                fmt.Println(v5, v56, v55)
            }
            _ = fi54
            _ = v55
            _ = v56
        }
        _ = fi52
        _ = v53
    }
    for fi59 := 0; fi59 < 12; fi59++ {
        var v60 uint64 = v50
        for fi61 := 0; fi61 < 2; fi61++ {
            v60 = v50
            var v62 uint64 = (v60 - uint64(v51))
            v1 = (v1 + ("foo bar" + ""))
            _ = fi61
            _ = v62
        }
        _ = fi59
        _ = v60
    }
    for fi63 := 0; fi63 < 20; fi63++ {
        v4 = 1.5
        v64 := GenS3{F0: 128}
        _ = fi63
        _ = v64
    }
    v4 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
