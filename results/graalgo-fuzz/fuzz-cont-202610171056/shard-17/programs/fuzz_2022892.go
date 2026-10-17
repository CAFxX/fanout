package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint64 }
func genFn0(p0 bool) int8 {
    p0 = ((1.5 * (float32(3.14159) - 1e10)) >= 0.1)
    if (!(447 != 2)) {
        var v6 int64 = 1
        v7 := []int{9223372036854775807}
        var v8 uint16 = 65535
        _ = v6
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 int32 = 1
            v11 := GenS1{F0: 542}
            var v12 int8 = 100
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    _ = p0
    return 31
}
func genFn1(p0 int16) int32 {
    if (!true) {
        p0 = 692
    } else {
        var v13 uint64 = 1
        p0 = p0
        v14 := []int{0}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        p0 = (p0 & 32767)
        switch fi15 {
        case 4294967295:
            if true {
                v16 := []int{(fi15 >> ((601) & 63)), fi15, (^(fi15))}
                switch 779 {
                case -330:
                    var v17 string = "\n\t"
                    _ = v17
                case -1:
                    var v18 bool = (true && true)
                    _ = v18
                }
                v16 = append(v16, 531)
                _ = v16
            }
        case 42:
            v19 := GenS1{F0: p0}
            _ = v19
        default:
            for fi20 := 0; fi20 < 20; fi20++ {
                if true {
                    if false {
                        var v21 int16 = (p0 % ((p0) | 1))
                        if true {
                            v22 := []int{}
                            v21 += -418
                            var v23 uint64 = 0
                            _ = v22
                            _ = v23
                        }
                        v21 += -32768
                        _ = v21
                    } else {
                        var v24 uint8 = 127
                        _ = v24
                    }
                }
                p0 = (p0 - (p0 << ((41) & 63)))
                _ = fi20
            }
        }
        _ = fi15
    }
    _ = p0
    return 2147483647
}

func main() {
    var v1 int16 = 32767
    v2 := []string{"hello", "a", ("foo bar" + "a")}
    v3 := make(map[string]string)
    var v4 int = 32767
    v5 := make(map[string]uint32)
    v4 += -273
    for fi25 := 0; fi25 < 18; fi25++ {
        if false {
            for fi26 := 0; fi26 < 7; fi26++ {
                var v27 uint = 1
                var v28 uint16 = 65535
                _ = fi26
                _ = v27
                _ = v28
            }
            switch ("\n\t" + (("" + ("" + ("" + "a"))) + (("a" + "\n\t") + "hello"))) {
            case "":
                if (false && ((!false) && true)) {
                    fmt.Println(v1)
                    fmt.Println(fi25, v1, v1)
                    if true {
                        v29 := v3[""]
                        if (!false) {
                            var v30 uint = 344
                            for fi31 := 0; fi31 < 4; fi31++ {
                                fmt.Println(v1, fi31)
                                switch "a" {
                                case "":
                                    v30 = v30
                                default:
                                    v32 := GenS1{F0: 201}
                                    _ = v32
                                }
                                _ = fi31
                            }
                            switch -2147483648 {
                            case 0:
                                switch v4 {
                                case 42:
                                    for fi33 := 0; fi33 < 19; fi33++ {
                                        var v34 int64 = -738
                                        _ = fi33
                                        _ = v34
                                    }
                                case 439:
                                    v35 := GenS2{F0: 86, F1: -472, F2: 42}
                                    _ = v35
                                case -2147483648:
                                    var v36 int16 = (v1 << (((v30 * (v30 * (^(v30))))) & 63))
                                    _ = v36
                                }
                            }
                            _ = v30
                        }
                        switch (fi25 & v4) {
                        case 32768:
                            var v37 float32 = 1.5
                            _ = v37
                        case 42:
                            v29 = "foo bar"
                        default:
                            for fi38 := 0; fi38 < 13; fi38++ {
                                var v39 uint8 = 1
                                _ = fi38
                                _ = v39
                            }
                        }
                        _ = v29
                    } else {
                        var v40 float32 = 1.5
                        for fi41 := 0; fi41 < 17; fi41++ {
                            v42 := v5[((("\n\t" + "a") + ("" + "hello")) + "\n\t")]
                            switch "" {
                            case "foo bar":
                                v42 = v42
                            }
                            _ = fi41
                            _ = v42
                        }
                        _ = v40
                    }
                } else {
                    if (!false) {
                        for fi43 := 0; fi43 < 14; fi43++ {
                            v44 := v5["foo bar"]
                            var v45 uint8 = 35
                            var v46 uint64 = 0
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                            _ = fi43
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        if len(v2) > 0 {
                            v2[0] = "\n\t"
                        }
                        for fi47 := 0; fi47 < 16; fi47++ {
                            if len(v2) > 0 {
                                v2[0] = ("a" + "foo bar")
                            }
                            v4 -= (-(fi47))
                            _ = fi47
                        }
                    }
                    _ = len(v3)
                    v1 += (v1 * v1)
                }
            case "\n\t":
                for fi48 := 0; fi48 < 3; fi48++ {
                    v4 += fi48
                    delete(v5, "\n\t")
                    var v49 uint16 = 1
                    _ = fi48
                    _ = v49
                }
            }
        } else {
            fmt.Println(v4, fi25, v1)
            if true {
                v1 += 42
                switch 42 {
                case 32767:
                    v2 = append(v2, "a")
                case -1:
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                case 930:
                    var v50 float64 = 1.5
                    _ = v50
                }
            }
            var v51 uint = 255
            _ = v51
        }
        _ = fi25
    }
    var v52 uint32 = 2
    v53 := GenS2{F0: 236, F1: -884, F2: 99}
    switch (("a" + "hello") + ("hello" + "a")) {
    case "":
        var v54 uint32 = (^(v52))
        _ = v54
    case "hello":
        var v55 int32 = -253
        _ = v55
    }
    delete(v5, "hello")
    switch (("\n\t" + "\n\t") + "") {
    case "a":
        var v56 int8 = 1
        _ = v56
    case "hello":
        var v57 float64 = (0.1 * 1.5)
        _ = v57
    case "":
        var v58 int8 = genFn0(false)
        _ = v58
    default:
        for fi59 := 0; fi59 < 19; fi59++ {
            for fi60 := 0; fi60 < 13; fi60++ {
                v5["foo bar"] = (v52 ^ (v52 << ((136) & 63)))
                _ = fi60
            }
            fmt.Println(v4, v1, v1)
            _ = fi59
        }
    }
    v52 -= (v52 / (((v52 % ((0) | 1))) | 1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        for fi61 := 0; fi61 < 20; fi61++ {
            var v62 int32 = genFn1(v1)
            var v63 uint16 = 69
            _ = fi61
            _ = v62
            _ = v63
        }
    case "a":
        var v64 string = ("a" + "\n\t")
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
}
