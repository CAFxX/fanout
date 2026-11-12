package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int8; F1 int64; F2 uint32 }

func main() {
    var v1 uint32 = 85
    var v2 string = ("hello" + (((("foo bar" + "") + (("foo bar" + "foo bar") + (("" + "") + ""))) + ("hello" + "\n\t")) + ("foo bar" + "\n\t")))
    v3 := make(map[int]bool)
    if ((v2 >= ((((v2 + "") + v2) + "hello") + v2)) && false) {
        v1 -= 4294967295
        if true {
            if (-698 > -61) {
                if (42 > 65535) {
                    v1 += (v1 << ((65535) & 63))
                }
                if true {
                    v1 -= v1
                    for fi4 := 0; fi4 < 3; fi4++ {
                        v2 = "hello"
                        _ = fi4
                    }
                } else {
                    for fi5 := 0; fi5 < 16; fi5++ {
                        v6 := []int{(-(fi5)), fi5, fi5}
                        var v7 string = ("a" + (v2 + v2))
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi5
                        _ = v6
                        _ = v7
                    }
                    for fi8 := 0; fi8 < 11; fi8++ {
                        v2 = (v2 + v2)
                        _ = fi8
                    }
                    v2 = "foo bar"
                }
                var v9 int16 = -1
                _ = v9
            } else {
                v2 = ((("" + ("\n\t" + v2)) + "a") + (("hello" + (v2 + ("" + v2))) + (("a" + v2) + (v2 + (v2 + "a")))))
            }
            var v10 int64 = -893
            if (32768 >= 32767) {
                v1 = (v1 & 366)
            } else {
                var v11 uint32 = (^(v1))
                _ = v11
            }
            _ = v10
        }
        var v12 string = "hello"
        _ = v12
    } else {
        if true {
            var v13 uint8 = 0
            if false {
                if true {
                    var v14 uint32 = uint32(v13)
                    for fi15 := 0; fi15 < 15; fi15++ {
                        v14 += (v14 - 376)
                        if (((0.0 * 0.1) - 0.1) != -2.25) {
                            v2 = (("\n\t" + v2) + v2)
                            for fi16 := 0; fi16 < 8; fi16++ {
                                v14 = 183
                                v17 := GenS1{F0: 42, F1: v1, F2: 255}
                                _ = fi16
                                _ = v17
                            }
                        } else {
                            switch v2 {
                            case "foo bar":
                                var v18 uint32 = (v14 << ((1) & 63))
                                _ = v18
                            case "hello":
                                v1 += v1
                            default:
                                v19 := []int{2147483647, fi15}
                                _ = v19
                            }
                            if (!(!true)) {
                                var v20 float64 = -2.25
                                _ = v20
                            } else {
                                v21 := GenS1{F0: 42, F1: (v14 * v14), F2: 182}
                                v22 := []int{(fi15 & (fi15 + 2147483647)), (fi15 << ((619) & 63)), 32768}
                                _ = v21
                                _ = v22
                            }
                        }
                        var v23 bool = (!(v1 >= 138))
                        _ = fi15
                        _ = v23
                    }
                    _ = v14
                }
                v24 := GenS2{F0: v2, F1: -41}
                if false {
                    var v25 uint16 = 32767
                    fmt.Println(v13, v25, v2)
                    _ = v25
                }
                _ = v24
            } else {
                if false {
                    var v26 string = (v2 + "foo bar")
                    v27 := GenS2{F0: "", F1: 0}
                    fmt.Println(v26, v2, v2)
                    _ = v26
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 4; fi28++ {
                        fmt.Println(v1)
                        _ = fi28
                    }
                    var v29 int16 = -32768
                    if true {
                        var v30 string = "foo bar"
                        var v31 uint8 = (v13 << ((545) & 63))
                        switch 128 {
                        case -9223372036854775808:
                            var v32 uint16 = 255
                            _ = v32
                        }
                        _ = v30
                        _ = v31
                    }
                    _ = v29
                }
                for fi33 := 0; fi33 < 3; fi33++ {
                    v34 := []int{42, (fi33 + (fi33 + (fi33 & (fi33 + (fi33 ^ fi33))))), 1}
                    switch (fi33 / ((42) | 1)) {
                    case -2147483648:
                        v35 := GenS2{F0: v2, F1: 127}
                        _ = v35
                    }
                    v1 = (v1 * (v1 | 65535))
                    _ = fi33
                    _ = v34
                }
            }
            _ = v13
        }
    }
    if true {
        var v36 float64 = (1e10 + 0.1)
        v37 := []int{-901, 32767}
        if (981 >= 32767) {
            v2 = ("\n\t" + ("" + ((v2 + (v2 + (v2 + v2))) + ("foo bar" + "hello"))))
        }
        _ = v36
        _ = v37
    } else {
        if (!false) {
            v38 := GenS1{F0: 65535, F1: v1, F2: 2}
            v1 += (^(v1))
            switch (v2 + v2) {
            case "hello":
                switch -1 {
                case 42:
                    v2 = ((v2 + ("" + (v2 + v2))) + (v2 + ("hello" + (v2 + (v2 + "")))))
                case 127:
                    for fi39 := 0; fi39 < 5; fi39++ {
                        v1 = (v1 >> ((255) & 63))
                        v40 := GenS1{F0: uint16(v1), F1: v1, F2: 255}
                        _ = fi39
                        _ = v40
                    }
                case 71:
                    v1 -= (v1 << ((65535) & 63))
                default:
                    if false {
                        for fi41 := 0; fi41 < 20; fi41++ {
                            var v42 int = (-(fi41))
                            _ = fi41
                            _ = v42
                        }
                        var v43 uint = 2
                        v44 := []int{}
                        _ = v43
                        _ = v44
                    }
                }
            case "foo bar":
                var v45 int32 = 0
                _ = v45
            }
            _ = v38
        } else {
            switch v2 {
            case "hello":
                if (!(-158 != -9223372036854775808)) {
                    if (!(100 > 127)) {
                        var v46 uint64 = 18446744073709551615
                        _ = v46
                    } else {
                        if false {
                            var v47 bool = false
                            _ = v47
                        }
                        v2 = v2
                        v2 = v2
                    }
                    fmt.Println(v2, v1)
                }
            case "a":
                delete(v3, -1)
            default:
                v48 := GenS1{F0: 255, F1: (v1 - 88), F2: 40}
                _ = v48
            }
            if (!false) {
                if true {
                    var v49 uint32 = 2
                    _ = v49
                }
                var v50 float32 = 0.1
                _ = v50
            }
        }
        fmt.Println(v2)
    }
    for fi51 := 0; fi51 < 18; fi51++ {
        if false {
            v2 = "hello"
            _ = len(v3)
        }
        v52 := GenS3{F0: 92, F1: -1, F2: v1}
        v2 = ""
        _ = fi51
        _ = v52
    }
    var v53 uint = 18446744073709551615
    _ = len(v3)
    v54 := GenS1{F0: 65535, F1: 65535, F2: 785}
    v55 := GenS1{F0: 2, F1: (v1 % ((v1) | 1)), F2: 413}
    for fi56 := 0; fi56 < 2; fi56++ {
        if ((v53 << ((255) & 63)) >= (v53 | v53)) {
            if false {
                v1 += (v1 | 4294967295)
                fmt.Println(v2, v1, v2)
                v3[(fi56 ^ 2147483647)] = true
            }
            if ((true && false) || (420 == fi56)) {
                v53 += v53
                fmt.Println(v53)
            }
            var v57 float32 = ((0.0 / 1.5) * (-2.25 / float32(-2.25)))
            _ = v57
        } else {
            v3[(fi56 & 127)] = (false && true)
            v58 := GenS1{F0: 255, F1: (v1 | v1), F2: 154}
            for fi59 := 0; fi59 < 10; fi59++ {
                _ = len(v3)
                v60 := GenS2{F0: (v2 + ""), F1: 35}
                v61 := []int{(fi56 / ((665) | 1)), fi56}
                _ = fi59
                _ = v60
                _ = v61
            }
            _ = v58
        }
        v62 := v3[fi56]
        _ = fi56
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
    _ = v54
    _ = v55
}
