package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint; F1 bool; F2 string }
type GenS3 struct { F0 uint; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int64; F1 string; F2 int16 }
func genFn0() int16 {
    var v4 int32 = -1
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            fmt.Println(v4, v4)
            _ = fi6
        }
        v4 = 0
        if false {
            v7 := GenS1{F0: (v4 % ((v4) | 1)), F1: v4}
            for fi8 := 0; fi8 < 4; fi8++ {
                if true {
                    var v9 int8 = 66
                    _ = v9
                } else {
                    v10 := []int{fi5, fi8, 811}
                    _ = v10
                }
                v11 := GenS2{F0: 4294967295, F1: false, F2: ""}
                _ = fi8
                _ = v11
            }
            var v12 int = fi5
            _ = v12
            _ = v7
        } else {
            for fi13 := 0; fi13 < 6; fi13++ {
                if true {
                    v4 = (v4 & v4)
                    var v14 int8 = 0
                    v4 += v4
                    _ = v14
                } else {
                    fmt.Println(fi13)
                    var v15 uint8 = 255
                    _ = v15
                }
                if (false && (false || (832 <= 32767))) {
                    v16 := []int{fi13}
                    _ = v16
                }
                _ = fi13
            }
            if (42 == 4294967295) {
                if true {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v18 := []int{32767, fi17}
                        _ = fi17
                        _ = v18
                    }
                    var v19 int = fi5
                    _ = v19
                }
            }
            if (false && false) {
                v20 := []int{32768, fi5, (fi5 % (((fi5 - fi5)) | 1))}
                _ = v20
            } else {
                for fi21 := 0; fi21 < 9; fi21++ {
                    if false {
                        v4 += 2147483647
                    } else {
                        fmt.Println(fi5)
                    }
                    fmt.Println(fi5, fi5, fi21)
                    _ = fi21
                }
                if true {
                    if (-563 == 42) {
                        v22 := []int{(fi5 | fi5), (fi5 << ((65535) & 63))}
                        fmt.Println(v4, fi5)
                        var v23 int = (fi5 << ((255) & 63))
                        _ = v22
                        _ = v23
                    } else {
                        v24 := GenS1{F0: 64, F1: (v4 % (((-(v4))) | 1))}
                        if false {
                            if (!(false && (!false))) {
                                if false {
                                    v25 := GenS2{F0: 0, F1: (2 <= 255), F2: ("" + "")}
                                    _ = v25
                                } else {
                                    var v26 int8 = 127
                                    _ = v26
                                }
                                v27 := GenS2{F0: 18446744073709551615, F1: (83 < 1), F2: "\n\t"}
                                v28 := []int{}
                                _ = v27
                                _ = v28
                            } else {
                                var v29 string = "foo bar"
                                v29 = "a"
                                v30 := GenS3{F0: 803, F1: v4, F2: 430}
                                _ = v29
                                _ = v30
                            }
                        }
                        v31 := GenS2{F0: 65535, F1: true, F2: "a"}
                        _ = v24
                        _ = v31
                    }
                    v4 = (v4 + v4)
                    v32 := GenS4{F0: 95}
                    _ = v32
                } else {
                    v4 += (v4 ^ v4)
                    var v33 uint16 = 42
                    _ = v33
                }
                for fi34 := 0; fi34 < 10; fi34++ {
                    v35 := GenS4{F0: 127}
                    v36 := []int{(fi34 >> ((0) & 63)), fi5}
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
        }
        _ = fi5
    }
    _ = v4
    return 695
}
func genFn1(p0 int32) int16 {
    var v37 uint32 = 0
    v38 := GenS1{F0: (p0 + p0), F1: (^(p0))}
    var v39 int8 = -1
    _ = p0
    _ = v37
    _ = v38
    _ = v39
    return -272
}

func main() {
    var v1 string = ("\n\t" + "a")
    var v2 string = ((v1 + v1) + "hello")
    var v3 string = "foo bar"
    v40 := GenS1{F0: 1, F1: 2147483647}
    if (false || (true || (true && false))) {
        v2 = v2
        v1 = ""
    } else {
        var v41 int64 = -9223372036854775808
        v3 = ("" + v3)
        _ = v41
    }
    fmt.Println(v1)
    var v42 int8 = 1
    fmt.Println(v2, v1)
    for fi43 := 0; fi43 < 13; fi43++ {
        if (1 != -296) {
            var v44 uint32 = 22
            v3 = v1
            _ = v44
        }
        v42 = (v42 - v42)
        if true {
            v45 := GenS4{F0: 114}
            if (!(-144 >= 0)) {
                switch v1 {
                case "\n\t":
                    switch fi43 {
                    case 819:
                        v42 -= (v42 % ((127) | 1))
                    case 530:
                        fmt.Println(v42, fi43)
                    case 127:
                        v2 = ((v2 + (("a" + "") + v3)) + v1)
                    default:
                        if false {
                            fmt.Println(v3, v1, v42)
                            switch (fi43 | 39) {
                            case -9223372036854775808:
                                v46 := GenS5{F0: 32767, F1: 194}
                                _ = v46
                            default:
                                for fi47 := 0; fi47 < 5; fi47++ {
                                    var v48 float32 = float32(float64(3.14159))
                                    var v49 int = fi47
                                    fmt.Println(v49)
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                            }
                            v1 = v3
                        } else {
                            v50 := GenS4{F0: 42}
                            v51 := []int{}
                            v2 = ("hello" + ((v1 + "a") + v3))
                            _ = v50
                            _ = v51
                        }
                    }
                case "hello":
                    if true {
                        v52 := GenS6{F0: 0}
                        if (!(42 < 1)) {
                            if (((float64(1.5) / 1.0) == 1.5) || false) {
                                v53 := []int{(fi43 - fi43), (fi43 - fi43), fi43}
                                _ = v53
                            } else {
                                if (fi43 <= (^(fi43))) {
                                    v1 = "a"
                                    if true {
                                        v54 := GenS4{F0: 31}
                                        _ = v54
                                    }
                                    v1 = (v2 + (("hello" + "foo bar") + "a"))
                                } else {
                                    v55 := []int{}
                                    if len(v55) >= 2 {
                                        v55 = v55[1:len(v55)]
                                    }
                                    _ = v55
                                }
                            }
                        }
                        _ = v52
                    } else {
                        fmt.Println(v1, fi43)
                        v56 := GenS3{F0: 1, F1: 2147483647, F2: 2}
                        switch v1 {
                        case "hello":
                            if false {
                                v42 = v42
                                for fi57 := 0; fi57 < 10; fi57++ {
                                    var v58 int32 = 0
                                    v59 := []int{1, (fi43 * (fi57 + 0))}
                                    v3 = "hello"
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                }
                                v60 := []int{127}
                                _ = v60
                            }
                        case "\n\t":
                            var v61 float32 = ((1e10 * -2.25) - float32((-2.25 * 1e10)))
                            _ = v61
                        case "":
                            fmt.Println(v42, fi43)
                        default:
                            var v62 string = ((v2 + "") + v1)
                            _ = v62
                        }
                        _ = v56
                    }
                case "foo bar":
                    v2 = (v1 + v2)
                }
            } else {
                var v63 string = (("foo bar" + "a") + "\n\t")
                _ = v63
            }
            var v64 uint = 18446744073709551615
            _ = v45
            _ = v64
        }
        _ = fi43
    }
    v2 = v2
    v42 += v42
    var v65 float32 = (1.5 - -2.25)
    v66 := []int{}
    if false {
        fmt.Println(v65, v42)
    } else {
        var v67 uint = 73
        _ = v67
    }
    v42 -= v42
    v68 := GenS7{F0: -487, F1: ((v1 + v2) + (v3 + v1)), F2: 1000}
    if len(v66) > 0 {
        v66[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v66:", len(v66), v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v42
    _ = v65
    _ = v66
    _ = v68
}
