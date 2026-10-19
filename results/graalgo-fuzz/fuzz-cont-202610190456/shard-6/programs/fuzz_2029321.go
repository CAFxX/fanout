package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 uint16 }
func genFn0() bool {
    var v5 float64 = 0.0
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := []int{(fi6 >> ((42) & 63))}
        _ = fi6
        _ = v7
    }
    v5 -= v5
    switch 435 {
    case -1:
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := []int{fi8, 912, 32767}
            v9 = append(v9, (fi8 ^ 0))
            _ = fi8
            _ = v9
        }
    case 898:
        var v10 float32 = float32(1.5)
        _ = v10
    default:
        v5 = float64((1.5 - float32((float64(3.14159) - 0.0))))
    }
    _ = v5
    return (1.5 <= -2.25)
}
func genFn1(p0 uint16, p1 string) int32 {
    v11 := []int{}
    for fi12 := 0; fi12 < 12; fi12++ {
        switch p1 {
        case "hello":
            p0 = (p0 % (((p0 & p0)) | 1))
        }
        _ = fi12
    }
    if true {
        if false {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            var v13 int = 892
            _ = v13
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := GenS1{F0: 42}
                _ = len(v11)
                _ = fi14
                _ = v15
            }
            v16 := GenS1{F0: 857}
            p1 = p1
            _ = v16
        }
        switch "a" {
        case "hello":
            for fi17 := 0; fi17 < 20; fi17++ {
                for fi18 := 0; fi18 < 18; fi18++ {
                    _ = len(v11)
                    _ = fi18
                }
                _ = fi17
            }
        default:
            p0 = 1
        }
        var v19 uint16 = 177
        _ = v19
    } else {
        var v20 int64 = 42
        v21 := GenS2{F0: 389}
        if (!false) {
            var v22 int8 = -18
            var v23 uint32 = 0
            _ = v22
            _ = v23
        }
        _ = v20
        _ = v21
    }
    p0 += p0
    if false {
        v24 := GenS3{F0: -1}
        fmt.Println(p1)
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v11
    return 42
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]uint32)
    v3 := []string{"foo bar", ("a" + ("a" + "a"))}
    var v4 uint32 = 2147483648
    switch 42 {
    case 4294967295:
        if (true || (426 > (v4 + (v4 & v4)))) {
            var v25 uint64 = 2
            _ = v25
        }
    }
    var v26 uint = 0
    v2[("foo bar" + "\n\t")] = (v4 ^ 4294967295)
    var v27 int = 127
    if (false && true) {
        if (!false) {
            fmt.Println(v26, v26, v4)
            if false {
                switch ("" + "a") {
                case "foo bar":
                    v28 := GenS3{F0: -511}
                    _ = v28
                case "hello":
                    if (v4 <= 2147483647) {
                        v26 -= (v26 % ((35) | 1))
                    } else {
                        var v29 string = ("\n\t" + "")
                        for fi30 := 0; fi30 < 3; fi30++ {
                            v31 := GenS1{F0: 42}
                            var v32 bool = (2147483647 > 460)
                            if false {
                                var v33 int16 = 0
                                _ = v33
                            }
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        switch v27 {
                        case 127:
                            var v34 uint32 = 420
                            _ = v34
                        case 9223372036854775807:
                            if false {
                                v1 = (v1 + (v1 << ((v26) & 63)))
                                v35 := GenS1{F0: 42}
                                v3 = append(v3, v29)
                                _ = v35
                            } else {
                                v27 = 9223372036854775807
                                var v36 int32 = -2147483648
                                _ = v36
                            }
                        case 2147483647:
                            fmt.Println(v29, v26, v4)
                        }
                        _ = v29
                    }
                case "\n\t":
                    var v37 string = (((("\n\t" + "foo bar") + (("\n\t" + "hello") + "foo bar")) + "hello") + ("\n\t" + ""))
                    _ = v37
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = ("hello" + "a")
                }
                for fi38 := 0; fi38 < 15; fi38++ {
                    v39 := GenS1{F0: 1}
                    switch fi38 {
                    case -9223372036854775808:
                        v27 -= 127
                    case -594:
                        for fi40 := 0; fi40 < 1; fi40++ {
                            var v41 uint32 = v4
                            _ = fi40
                            _ = v41
                        }
                    case -1:
                        if false {
                            var v42 string = ""
                            switch (((v42 + v42) + v42) + v42) {
                            case "hello":
                                v4 = 4294967295
                            case "foo bar":
                                for fi43 := 0; fi43 < 15; fi43++ {
                                    var v44 string = v42
                                    _ = fi43
                                    _ = v44
                                }
                            }
                            _ = v42
                        }
                    default:
                        v45 := GenS2{F0: (v26 + (v26 + v26))}
                        _ = v45
                    }
                    _ = fi38
                    _ = v39
                }
                fmt.Println(v27, v1, v27)
            }
            fmt.Println(v1)
        }
    }
    v46 := GenS1{F0: 297}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        if (32767 <= 1000) {
            if true {
                v27 += (v27 >> ((v26) & 63))
            } else {
                var v47 uint32 = v4
                _ = v47
            }
            var v48 int = (v27 % ((v27) | 1))
            var v49 float64 = ((1.5 + (float64(((0.1 * (float32(1.5) / (0.1 - 1.5))) + -2.25)) + (1.5 + 1e10))) / 1.5)
            _ = v48
            _ = v49
        }
        if (true || false) {
            switch ("foo bar" + ("" + ((("\n\t" + "\n\t") + ("a" + "foo bar")) + ""))) {
            case "hello":
                v27 = v27
            }
            if ((!false) || (!(1 >= 0))) {
                for fi50 := 0; fi50 < 9; fi50++ {
                    var v51 int32 = 1
                    v26 -= v26
                    _ = fi50
                    _ = v51
                }
                var v52 float64 = (3.14159 - 1.5)
                _ = v52
            } else {
                v53 := GenS4{F0: "\n\t", F1: (v26 | v26), F2: v4}
                var v54 int = (-(v27))
                for fi55 := 0; fi55 < 20; fi55++ {
                    v4 = v4
                    for fi56 := 0; fi56 < 3; fi56++ {
                        v4 = v4
                        v2["hello"] = 2147483647
                        _ = fi56
                    }
                    v1 = v1
                    _ = fi55
                }
                _ = v53
                _ = v54
            }
        }
        v57 := GenS5{F0: v4}
        _ = v57
    }
    for fi58 := 0; fi58 < 5; fi58++ {
        v59 := GenS6{F0: false, F1: 0}
        _ = fi58
        _ = v59
    }
    if false {
        v3 = append(v3, "\n\t")
    }
    for fi60 := 0; fi60 < 15; fi60++ {
        fmt.Println(fi60, v27)
        v1 += (v1 % (((-(v1))) | 1))
        _ = fi60
    }
    v3 = append(v3, "hello")
    var v61 int32 = 0
    v62 := GenS4{F0: "", F1: 452, F2: 0}
    var v63 int8 = v1
    for fi64 := 0; fi64 < 7; fi64++ {
        v65 := GenS5{F0: v4}
        fmt.Println(v26, fi64, v63)
        _ = fi64
        _ = v65
    }
    fmt.Println(v27)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v46
    _ = v61
    _ = v62
    _ = v63
}
