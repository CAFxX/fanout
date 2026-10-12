package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 uint; F2 bool }

func main() {
    var v1 int8 = 0
    var v2 int8 = 0
    var v3 string = (("" + "a") + "hello")
    v4 := []int{32767, 127, -44}
    if (true || (((0.1 / -2.25) + 0.0) > float32(1.5))) {
        if false {
            if (!(!(1 <= -1))) {
                v5 := GenS1{F0: -1}
                var v6 int8 = v1
                v2 -= 100
                _ = v5
                _ = v6
            }
            for fi7 := 0; fi7 < 13; fi7++ {
                v3 = v3
                switch fi7 {
                case 42:
                    _ = len(v4)
                }
                var v8 float64 = ((-2.25 / (((1e10 - 0.0) * 1e10) - 0.1)) * 1.5)
                _ = fi7
                _ = v8
            }
            switch v3 {
            case "":
                for fi9 := 0; fi9 < 10; fi9++ {
                    v10 := GenS1{F0: -1}
                    _ = fi9
                    _ = v10
                }
            case "a":
                v3 = v3
            default:
                for fi11 := 0; fi11 < 4; fi11++ {
                    if false {
                        var v12 float32 = float32(3.14159)
                        _ = v12
                    }
                    switch fi11 {
                    case 128:
                        var v13 bool = ("a" > ("hello" + (v3 + ((v3 + "a") + v3))))
                        _ = v13
                    }
                    _ = fi11
                }
            }
        } else {
            v3 = v3
            var v14 uint8 = 229
            _ = v14
        }
        for fi15 := 0; fi15 < 7; fi15++ {
            switch v3 {
            case "":
                v2 = v1
            }
            _ = fi15
        }
    }
    if false {
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := GenS1{F0: 2147483647}
            v18 := GenS1{F0: 2147483647}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 1; fi19++ {
            if false {
                v4 = append(v4, fi19)
            } else {
                for fi20 := 0; fi20 < 18; fi20++ {
                    var v21 uint = 255
                    v22 := GenS2{F0: -505, F1: 179}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
            fmt.Println(v1, v1)
            var v23 bool = false
            _ = fi19
            _ = v23
        }
    }
    var v24 int64 = -363
    fmt.Println(v3)
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 int8 = v1
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        v28 := GenS2{F0: -1, F1: -1}
        v29 := GenS3{F0: true}
        fmt.Println(v24, v3)
        _ = fi27
        _ = v28
        _ = v29
    }
    v2 -= v1
    v1 = (^(v2))
    var v30 bool = (float64(1.5) <= -2.25)
    switch (v3 + ("foo bar" + v3)) {
    case "foo bar":
        v2 = (^(v1))
    default:
        if (4294967295 == 360) {
            var v31 uint32 = 42
            _ = v31
        } else {
            switch v3 {
            case "foo bar":
                switch 4294967295 {
                case -9223372036854775808:
                    v3 = v3
                case -82:
                    if ((v1 | 42) >= -1) {
                        var v32 float32 = (3.14159 + 1e10)
                        for fi33 := 0; fi33 < 18; fi33++ {
                            var v34 uint = 65535
                            _ = fi33
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 3; fi35++ {
                            for fi36 := 0; fi36 < 9; fi36++ {
                                v24 = v24
                                var v37 uint64 = 0
                                _ = fi36
                                _ = v37
                            }
                            _ = fi35
                        }
                        _ = v32
                    } else {
                        fmt.Println(v1, v30, v3)
                        for fi38 := 0; fi38 < 3; fi38++ {
                            var v39 int32 = 655
                            switch v3 {
                            case "hello":
                                var v40 uint8 = 42
                                _ = v40
                            case "a":
                                v41 := GenS3{F0: ((v30 || v30) && (!v30))}
                                _ = v41
                            }
                            _ = fi38
                            _ = v39
                        }
                    }
                case 853:
                    switch 4294967295 {
                    case 4294967295:
                        if v30 {
                            switch 32768 {
                            case 1:
                                switch ("\n\t" + v3) {
                                case "\n\t":
                                    v24 -= (^(v24))
                                case "foo bar":
                                    v42 := GenS4{F0: v3, F1: 223}
                                    _ = v42
                                case "hello":
                                    if true {
                                        for fi43 := 0; fi43 < 16; fi43++ {
                                            var v44 uint = 4294967295
                                            _ = fi43
                                            _ = v44
                                        }
                                        var v45 float64 = ((1.5 / 3.14159) * 3.14159)
                                        _ = v45
                                    }
                                }
                            case 4294967295:
                                if v30 {
                                    var v46 float32 = (((float32(float64(((-2.25 / float32((1e10 * 1.5))) * (1.5 - 1e10)))) + 1e10) * (0.0 * -2.25)) + 1.5)
                                    _ = v46
                                }
                            default:
                                for fi47 := 0; fi47 < 19; fi47++ {
                                    for fi48 := 0; fi48 < 4; fi48++ {
                                        var v49 bool = (v30 && ((90 >= 65535) && true))
                                        var v50 uint = 838
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                    }
                                    var v51 string = "a"
                                    _ = fi47
                                    _ = v51
                                }
                            }
                            fmt.Println(v24)
                        } else {
                            var v52 int = 32767
                            _ = v52
                        }
                    case 0:
                        var v53 int8 = v2
                        _ = v53
                    case 128:
                        if (!v30) {
                            v3 = (v3 + "a")
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v54 string = "a"
                            _ = v54
                        }
                    default:
                        var v55 uint32 = 42
                        _ = v55
                    }
                default:
                    var v56 int8 = (v1 << ((348) & 63))
                    _ = v56
                }
            case "a":
                switch -124 {
                case 186:
                    switch v3 {
                    case "hello":
                        v57 := GenS5{F0: -176}
                        _ = v57
                    case "a":
                        var v58 float64 = float64(((float32(3.14159) / ((0.1 / (1.5 - float32(((((3.14159 - 1.5) - -2.25) / 1.5) + 1e10)))) * (0.1 / 1.5))) - 1.5))
                        _ = v58
                    case "foo bar":
                        var v59 bool = true
                        _ = v59
                    default:
                        v2 += 100
                    }
                case 32768:
                    for fi60 := 0; fi60 < 15; fi60++ {
                        if len(v4) > 0 {
                            v4[0] = 347
                        }
                        if ((4294967295 < 2) && (-111 >= 0)) {
                            fmt.Println(v2, v3, v1)
                            var v61 float32 = 0.1
                            v62 := GenS1{F0: 42}
                            _ = v61
                            _ = v62
                        }
                        _ = fi60
                    }
                case 9223372036854775807:
                    var v63 float64 = 0.1
                    _ = v63
                default:
                    v30 = (-1 < 0)
                }
            default:
                for fi64 := 0; fi64 < 8; fi64++ {
                    switch (fi64 - (fi64 >> ((150) & 63))) {
                    case 624:
                        _ = len(v4)
                    default:
                        v65 := GenS6{F0: v3, F1: 255, F2: v30}
                        _ = v65
                    }
                    v2 = (v1 & v2)
                    fmt.Println(v1, v1)
                    _ = fi64
                }
            }
            fmt.Println(v2, v1)
            var v66 int16 = 682
            _ = v66
        }
    }
    var v67 string = v3
    v68 := GenS4{F0: v67, F1: 1}
    var v69 uint32 = 0
    v4 = append(v4, -1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v67
    _ = v68
    _ = v69
}
