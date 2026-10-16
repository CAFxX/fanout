package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 bool; F1 uint8 }

func main() {
    var v1 int32 = 42
    var v2 uint32 = 4294967295
    var v3 string = ("a" + ("\n\t" + "\n\t"))
    var v4 uint64 = 4294967295
    var v5 int64 = 0
    v6 := make(map[string]bool)
    var v7 int8 = 110
    v7 += -128
    v3 = (v3 + v3)
    if (2 == 0) {
        fmt.Println(v3)
        if (!((2 >= 2) && true)) {
            for fi8 := 0; fi8 < 6; fi8++ {
                fmt.Println(v1, v5)
                var v9 uint8 = 128
                _ = fi8
                _ = v9
            }
            if true {
                var v10 string = "a"
                v11 := []int{}
                _ = v10
                _ = v11
            } else {
                v3 = ((v3 + "\n\t") + "\n\t")
                if false {
                    var v12 float64 = (0.1 / 1e10)
                    switch 0 {
                    case -13:
                        v12 -= (v12 + v12)
                    case 80:
                        var v13 float32 = float32(float64(((1e10 + 0.0) + 0.1)))
                        _ = v13
                    }
                    if false {
                        v14 := v6["a"]
                        _ = v14
                    } else {
                        var v15 bool = (v2 != 957)
                        _ = v15
                    }
                    _ = v12
                } else {
                    var v16 float32 = float32((0.0 + 0.1))
                    v3 = "\n\t"
                    _ = v16
                }
            }
        }
        v17 := GenS1{F0: -2147483648, F1: -1}
        _ = v17
    } else {
        v18 := []int{1, 900}
        _ = v18
    }
    fmt.Println(v7)
    var v19 int16 = int16(v1)
    v2 -= v2
    fmt.Println(v4)
    _ = len(v6)
    switch "hello" {
    case "":
        for fi20 := 0; fi20 < 11; fi20++ {
            v3 = ("hello" + v3)
            switch (("a" + "") + "foo bar") {
            case "hello":
                if true {
                    var v21 int8 = (-(v7))
                    v22 := GenS1{F0: v1, F1: (-(v19))}
                    if (fi20 < fi20) {
                        v23 := []int{fi20, 2147483647, (fi20 / ((-686) | 1)), fi20}
                        _ = v23
                    } else {
                        fmt.Println(v4)
                    }
                    _ = v21
                    _ = v22
                } else {
                    var v24 int16 = (v19 * 0)
                    var v25 bool = true
                    _ = v24
                    _ = v25
                }
            case "a":
                v5 = (v5 + 9223372036854775807)
            case "foo bar":
                if (true || (true && (255 < 2))) {
                    v26 := GenS1{F0: 2147483647, F1: (v19 ^ (v19 >> ((87) & 63)))}
                    if false {
                        var v27 float32 = 1e10
                        if false {
                            v28 := GenS1{F0: int32(v7), F1: (v19 - v19)}
                            _ = len(v6)
                            var v29 string = "hello"
                            _ = v28
                            _ = v29
                        }
                        switch v3 {
                        case "\n\t":
                            for fi30 := 0; fi30 < 7; fi30++ {
                                if ((v1 >> ((4294967295) & 63)) == (v1 << ((uint(v4)) & 63))) {
                                    if true {
                                        v31 := []int{(fi20 ^ (-(fi20))), (fi20 - (-(fi30))), fi20}
                                        var v32 string = ""
                                        if len(v31) >= 2 {
                                            v31 = v31[1:len(v31)]
                                        }
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v33 string = v3
                                        _ = v33
                                    }
                                    v34 := GenS1{F0: int32(v4), F1: -1}
                                    var v35 int16 = (v19 + v19)
                                    _ = v34
                                    _ = v35
                                } else {
                                    v1 = 42
                                    var v36 float32 = v27
                                    fmt.Println(v4, v5, v19)
                                    _ = v36
                                }
                                var v37 uint32 = uint32(v4)
                                _ = fi30
                                _ = v37
                            }
                        case "foo bar":
                            for fi38 := 0; fi38 < 19; fi38++ {
                                switch (-(fi20)) {
                                case 128:
                                    fmt.Println(v7, v4, v1)
                                default:
                                    v5 = (v5 << ((4294967295) & 63))
                                    v39 := GenS1{F0: (v1 % ((v1) | 1)), F1: 1}
                                    var v40 int32 = v1
                                    _ = v39
                                    _ = v40
                                }
                                _ = fi38
                            }
                        case "hello":
                            var v41 uint32 = 2147483647
                            _ = v41
                        default:
                            var v42 uint = 2
                            _ = v42
                        }
                        _ = v27
                    }
                    _ = v26
                }
            default:
                for fi43 := 0; fi43 < 14; fi43++ {
                    var v44 string = v3
                    var v45 uint32 = v2
                    v1 -= v1
                    var v46 int32 = v1
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v46
                }
            }
            var v47 int32 = 2147483647
            _ = fi20
            _ = v47
        }
    default:
        for fi48 := 0; fi48 < 2; fi48++ {
            v7 = (-(v7))
            _ = fi48
        }
    }
    var v49 int32 = int32(v2)
    var v50 uint8 = 127
    var v51 int64 = (v5 ^ (-(v5)))
    if false {
        v50 -= 128
    } else {
        v52 := GenS2{F0: ((v50 ^ (^(v50))) >= 128), F1: (v50 * v50)}
        _ = v52
    }
    var v53 float64 = 3.14159
    v54 := []int{int(v51)}
    if len(v54) >= 2 {
        v54 = v54[1:len(v54)]
    }
    for fi55 := 0; fi55 < 1; fi55++ {
        for fi56 := 0; fi56 < 14; fi56++ {
            if (true || (!(255 <= 42))) {
                var v57 float32 = 1e10
                if len(v54) >= 2 {
                    v54 = v54[1:len(v54)]
                }
                fmt.Println(v19, v4, v1)
                _ = v57
            } else {
                v58 := v6[v3]
                v50 = 46
                fmt.Println(v3, v7, fi55)
                _ = v58
            }
            _ = fi56
        }
        v59 := v6[(v3 + v3)]
        if (!(!v59)) {
            v60 := v6[("foo bar" + (v3 + v3))]
            fmt.Println(v1)
            _ = v60
        } else {
            var v61 string = "hello"
            var v62 bool = ((v51 % ((v51) | 1)) != (^(v51)))
            switch (fi55 % ((2147483647) | 1)) {
            case 2147483647:
                var v63 uint32 = v2
                _ = v63
            default:
                var v64 bool = false
                _ = v64
            }
            _ = v61
            _ = v62
        }
        _ = fi55
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v53
    _ = v54
    _ = v6
    _ = v7
}
