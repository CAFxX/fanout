package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{42, 65535, 42, 42}
    v3 := []string{("" + "\n\t"), "hello"}
    v4 := []int{}
    var v5 int8 = 5
    if ((65535 >= 18446744073709551615) && (float32((3.14159 - 0.1)) < (float32((3.14159 + 1.5)) + 1.5))) {
        fmt.Println(v5)
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := GenS1{F0: 127, F1: -381, F2: -727}
            _ = fi6
            _ = v7
        }
        switch 32768 {
        case -351:
            for fi8 := 0; fi8 < 4; fi8++ {
                for fi9 := 0; fi9 < 20; fi9++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi9
                }
                _ = fi8
            }
        default:
            if false {
                var v10 float32 = 1e10
                v10 = -2.25
                _ = v10
            } else {
                v5 += v5
                for fi11 := 0; fi11 < 17; fi11++ {
                    switch fi11 {
                    case 42:
                        v12 := GenS1{F0: 42, F1: 0, F2: -1}
                        _ = v12
                    }
                    v5 = -1
                    _ = fi11
                }
                var v13 int64 = -745
                _ = v13
            }
        }
    } else {
        delete(v1, ("a" + "a"))
        if false {
            var v14 uint16 = 147
            v14 += v14
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        if (1 == 818) {
            var v16 float64 = 0.0
            _ = v16
        }
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := GenS2{F0: ("\n\t" + "hello")}
            v19 := GenS3{F0: "hello", F1: 128, F2: -723}
            if (!false) {
                if false {
                    _ = len(v3)
                    if true {
                        v5 = (v5 ^ v5)
                        var v20 int8 = 43
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                        _ = v20
                    } else {
                        switch "\n\t" {
                        case "":
                            switch fi17 {
                            case 9223372036854775807:
                                if true {
                                    v21 := GenS1{F0: 255, F1: -391, F2: 245}
                                    v22 := GenS1{F0: 42, F1: -643, F2: 1}
                                    _ = v21
                                    _ = v22
                                }
                            }
                        case "foo bar":
                            var v23 float32 = 1e10
                            _ = v23
                        default:
                            var v24 uint16 = 65535
                            _ = v24
                        }
                        v25 := GenS1{F0: 128, F1: -407, F2: 383}
                        _ = v25
                    }
                    v5 -= -128
                }
                if len(v4) > 0 {
                    v4[0] = (fi15 & fi17)
                }
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi15
    }
    v26 := GenS3{F0: "foo bar", F1: 0, F2: 32767}
    v5 = (v5 % (((v5 << ((850) & 63))) | 1))
    var v27 uint16 = 580
    if (!false) {
        for fi28 := 0; fi28 < 4; fi28++ {
            if ((!(false || true)) || false) {
                var v29 bool = true
                var v30 bool = v29
                var v31 string = "a"
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi28
        }
        fmt.Println(v5)
        var v32 float32 = (1e10 / (1e10 - ((((3.14159 + 1.5) / 1.0) - 0.1) / 1.0)))
        _ = v32
    }
    var v33 bool = (32768 < 32768)
    v33 = true
    v34 := v1["\n\t"]
    switch ((("a" + (v34 + v34)) + (v34 + v34)) + "a") {
    case "\n\t":
        for fi35 := 0; fi35 < 8; fi35++ {
            _ = len(v4)
            v34 = (("" + v34) + v34)
            _ = fi35
        }
    case "foo bar":
        _ = len(v3)
    case "a":
        delete(v1, "")
    }
    v36 := GenS3{F0: v34, F1: 128, F2: 42}
    v34 = ("foo bar" + "")
    switch v34 {
    case "\n\t":
        for fi37 := 0; fi37 < 3; fi37++ {
            v1["hello"] = v34
            for fi38 := 0; fi38 < 12; fi38++ {
                var v39 int8 = 0
                var v40 int32 = -1
                if false {
                    v27 -= 747
                } else {
                    v33 = (v33 || true)
                    var v41 int = (fi38 ^ (-(fi38)))
                    _ = v41
                }
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = fi37
        }
    case "foo bar":
        v42 := GenS1{F0: 127, F1: 9223372036854775807, F2: -2147483648}
        _ = v42
    }
    fmt.Println(v27)
    for fi43 := 0; fi43 < 12; fi43++ {
        var v44 string = v34
        if (!(!v33)) {
            var v45 float32 = 1e10
            if (42 > 339) {
                var v46 bool = (42 > 0)
                v27 += 422
                _ = v46
            } else {
                var v47 uint16 = 65535
                _ = v47
            }
            if (529 != 255) {
                delete(v1, "")
            } else {
                switch fi43 {
                case 52:
                    v33 = (v27 <= v27)
                case -739:
                    fmt.Println(v27, v45, v33)
                default:
                    fmt.Println(v5, v44)
                }
            }
            _ = v45
        } else {
            v48 := GenS4{F0: 368, F1: 18446744073709551615}
            v49 := v1[v44]
            _ = v48
            _ = v49
        }
        _ = fi43
        _ = v44
    }
    if v33 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi50 := 0; fi50 < 20; fi50++ {
            if (-412 > 0) {
                v51 := GenS5{F0: ((fi50 & (fi50 << ((775) & 63))) <= 0), F1: 0}
                _ = v51
            }
            _ = fi50
        }
        switch 42 {
        case -767:
            if (4294967295 >= 943) {
                v27 += (v27 - (v27 + v27))
                var v52 float64 = float64((0.0 / -2.25))
                _ = v52
            }
        default:
            var v53 uint = 4294967295
            _ = v53
        }
    }
    v34 = ("" + v34)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
}
