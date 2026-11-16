package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 uint16; F2 int64 }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -978
    var v3 int8 = 100
    var v4 string = (("" + ("" + "foo bar")) + "hello")
    v5 := make(map[int]int)
    v6 := make(map[string]bool)
    switch "" {
    case "":
        fmt.Println(v4)
    case "hello":
        v2 += (v2 * -304)
    case "foo bar":
        v7 := GenS1{F0: 210}
        _ = v7
    default:
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := v1[(fi8 * fi8)]
            if (!(!(-1 <= 631))) {
                switch (fi8 / (((fi8 ^ fi8)) | 1)) {
                case -443:
                    v2 -= -880
                case 32768:
                    for fi10 := 0; fi10 < 6; fi10++ {
                        switch fi8 {
                        case -22:
                            var v11 bool = (42 > 375)
                            _ = v11
                        case 127:
                            fmt.Println(fi8)
                        case -537:
                            var v12 float64 = 1.5
                            _ = v12
                        default:
                            if true {
                                var v13 int16 = (v2 % ((v2) | 1))
                                var v14 int8 = 1
                                switch v4 {
                                case "\n\t":
                                    v15 := []int{fi10}
                                    _ = v15
                                }
                                _ = v13
                                _ = v14
                            }
                            v2 = (v2 - int16(v3))
                            switch (fi10 >> ((613) & 63)) {
                            case 982:
                                v16 := GenS1{F0: 232}
                                _ = v16
                            }
                            v2 -= -804
                            var v17 float32 = float32(3.14159)
                            _ = v17
                        }
                        var v18 int16 = (v2 << ((42) & 63))
                        if v9 {
                            var v19 int = (fi10 & fi10)
                            if v9 {
                                v4 = v4
                                v20 := GenS2{F0: (v4 + "a")}
                                _ = v20
                            } else {
                                var v21 int = fi10
                                _ = v21
                            }
                            var v22 int8 = v3
                            _ = v19
                            _ = v22
                        }
                        if false {
                            var v23 int16 = 32767
                            var v24 int = -706
                            _ = v23
                            _ = v24
                        }
                        _ = fi10
                        _ = v18
                    }
                case -9223372036854775808:
                    v25 := GenS3{F0: 103, F1: 15, F2: 42}
                    _ = v25
                default:
                    var v26 bool = (false && true)
                    _ = v26
                }
                if v9 {
                    var v27 float64 = 1e10
                    for fi28 := 0; fi28 < 2; fi28++ {
                        v29 := GenS4{F0: "\n\t"}
                        var v30 uint64 = 42
                        v27 -= ((1e10 / (v27 - v27)) - (3.14159 * 1e10))
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    var v31 uint64 = 1
                    _ = v27
                    _ = v31
                } else {
                    fmt.Println(v9, v4, v3)
                    v32 := GenS3{F0: 255, F1: -48, F2: 42}
                    var v33 int64 = -1
                    _ = v32
                    _ = v33
                }
                switch (v4 + v4) {
                case "a":
                    fmt.Println(v4, fi8)
                case "hello":
                    v2 -= v2
                }
            } else {
                v34 := []int{(fi8 - fi8), fi8, 32767}
                if v9 {
                    for fi35 := 0; fi35 < 9; fi35++ {
                        var v36 float64 = ((-2.25 * float64(((1.5 - 0.1) / 0.1))) + -2.25)
                        _ = fi35
                        _ = v36
                    }
                }
                _ = v34
            }
            _ = fi8
            _ = v9
        }
    }
    fmt.Println(v2, v2, v2)
    delete(v6, v4)
    var v37 uint32 = 2147483648
    if (127 != 2) {
        for fi38 := 0; fi38 < 4; fi38++ {
            for fi39 := 0; fi39 < 19; fi39++ {
                if ((!(251 > uint8(v37))) || (255 <= 2)) {
                    var v40 string = v4
                    v4 = v40
                    _ = v40
                }
                _ = fi39
            }
            fmt.Println(fi38)
            v41 := GenS5{F0: -1, F1: uint16(v37), F2: 722}
            _ = fi38
            _ = v41
        }
    } else {
        for fi42 := 0; fi42 < 3; fi42++ {
            v43 := GenS6{F0: (v2 * 970), F1: 0}
            _ = fi42
            _ = v43
        }
        v37 = v37
    }
    switch v4 {
    case "hello":
        switch (v4 + ((v4 + ("a" + (v4 + v4))) + ("a" + v4))) {
        case "":
            for fi44 := 0; fi44 < 17; fi44++ {
                for fi45 := 0; fi45 < 8; fi45++ {
                    v3 += (^(v3))
                    fmt.Println(fi44)
                    v46 := []int{}
                    _ = fi45
                    _ = v46
                }
                _ = fi44
            }
        default:
            if (42 > v3) {
                v37 = v37
                var v47 uint = 4294967295
                var v48 uint = v47
                _ = v47
                _ = v48
            } else {
                _ = len(v1)
            }
        }
    }
    if (!((!(!(true || false))) || false)) {
        v37 -= 42
        fmt.Println(v3, v3)
        delete(v1, 903)
    } else {
        var v49 float32 = 1.5
        v49 = (-2.25 / v49)
        v49 = (v49 * ((((0.1 + ((v49 + 1e10) * v49)) / -2.25) * ((v49 * v49) - v49)) - float32((float64(((v49 - v49) / (((0.0 - (v49 / (1e10 / 1.5))) + float32(-2.25)) * float32(1e10)))) / ((1e10 - 0.1) / 1.0)))))
        _ = v49
    }
    for fi50 := 0; fi50 < 1; fi50++ {
        for fi51 := 0; fi51 < 20; fi51++ {
            v52 := []int{(fi51 * fi51), (fi51 % ((-9223372036854775808) | 1)), fi51, (-(fi50))}
            if (!(!false)) {
                if false {
                    if false {
                        v52 = append(v52, 0)
                        v4 = v4
                        v53 := GenS4{F0: "a"}
                        _ = v53
                    } else {
                        var v54 uint8 = 198
                        v3 += int8(v37)
                        var v55 uint32 = v37
                        _ = v54
                        _ = v55
                    }
                } else {
                    var v56 float64 = 0.1
                    switch v4 {
                    case "a":
                        if false {
                            if true {
                                var v57 uint = 255
                                var v58 uint8 = 42
                                _ = v57
                                _ = v58
                            }
                            _ = len(v1)
                            switch ("" + "\n\t") {
                            case "a":
                                v56 -= 0.1
                            case "foo bar":
                                v4 = v4
                            default:
                                _ = len(v6)
                            }
                        } else {
                            v4 = v4
                            var v59 int = fi50
                            _ = v59
                        }
                    }
                    _ = v56
                }
            } else {
                if false {
                    fmt.Println(v37, v3, fi50)
                    v2 = v2
                } else {
                    if len(v52) > 0 {
                        v52[0] = fi51
                    }
                }
            }
            _ = fi51
            _ = v52
        }
        var v60 uint = 2
        fmt.Println(v4, fi50)
        _ = fi50
        _ = v60
    }
    v6[v4] = (482 >= 42)
    if true {
        v61 := GenS7{F0: 1, F1: 32767}
        switch -2147483648 {
        case 32768:
            v4 = v4
        }
        _ = v61
    } else {
        switch 127 {
        case 32767:
            var v62 uint64 = 1
            _ = v62
        default:
            var v63 float64 = (1e10 + -2.25)
            _ = v63
        }
    }
    var v64 uint64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
