package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 bool }
type GenS2 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS3 struct { F0 int32; F1 uint16; F2 bool }
type GenS4 struct { F0 string; F1 int64; F2 int32 }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 int = 0
    v3 := []int{}
    v4 := []int{(-(v2)), (v2 << ((0) & 63)), v2}
    var v5 uint8 = 42
    var v6 int = (v2 * (v2 >> ((18446744073709551615) & 63)))
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 uint16 = 32767
            var v10 string = ""
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    if ((!true) && false) {
        var v11 string = "foo bar"
        _ = len(v4)
        _ = v11
    } else {
        if ((v5 % (((v5 * v5)) | 1)) > 245) {
            var v12 uint16 = 1
            v2 -= (v6 * v2)
            _ = v12
        }
        switch v6 {
        case 32767:
            for fi13 := 0; fi13 < 20; fi13++ {
                switch "" {
                case "foo bar":
                    v14 := v1["foo bar"]
                    _ = v14
                case "":
                    var v15 float64 = 0.0
                    _ = v15
                case "a":
                    switch ("" + "hello") {
                    case "hello":
                        for fi16 := 0; fi16 < 7; fi16++ {
                            _ = len(v4)
                            if false {
                                _ = len(v3)
                                var v17 int8 = 0
                                fmt.Println(fi13, v5)
                                _ = v17
                            }
                            _ = fi16
                        }
                    case "\n\t":
                        fmt.Println(v2)
                    case "foo bar":
                        fmt.Println(fi13, v5, v6)
                    default:
                        var v18 int16 = -32768
                        _ = v18
                    }
                }
                v5 += v5
                fmt.Println(v6, v2, fi13)
                _ = fi13
            }
        case 2147483647:
            v19 := GenS1{F0: -32768, F1: 42, F2: true}
            _ = v19
        }
        var v20 uint32 = 65535
        _ = v20
    }
    if false {
        switch (("a" + (("\n\t" + ("\n\t" + "a")) + "\n\t")) + "hello") {
        case "hello":
            if true {
                if (-2.25 != (1.5 - -2.25)) {
                    if false {
                        var v21 bool = true
                        if true {
                            v2 -= (v2 ^ (v6 + 0))
                            _ = len(v3)
                        } else {
                            v6 += v6
                            switch "\n\t" {
                            case "hello":
                                for fi22 := 0; fi22 < 1; fi22++ {
                                    v2 = v2
                                    v23 := GenS2{F0: 255, F1: (v5 / (((^(v5))) | 1)), F2: -2147483648}
                                    if v21 {
                                        v24 := GenS3{F0: -1, F1: 1, F2: (!((!(v21 && (!false))) && true))}
                                        for fi25 := 0; fi25 < 3; fi25++ {
                                            var v26 uint = 873
                                            _ = fi25
                                            _ = v26
                                        }
                                        var v27 bool = v21
                                        _ = v24
                                        _ = v27
                                    } else {
                                        _ = len(v3)
                                    }
                                    _ = fi22
                                    _ = v23
                                }
                            case "a":
                                v28 := v1["hello"]
                                _ = v28
                            case "foo bar":
                                v29 := GenS2{F0: 629, F1: (v5 % ((42) | 1)), F2: 2147483647}
                                _ = v29
                            }
                            var v30 uint = 18446744073709551615
                            _ = v30
                        }
                        _ = v21
                    }
                }
            }
        case "":
            for fi31 := 0; fi31 < 3; fi31++ {
                v2 -= 36
                fmt.Println(fi31, v5, v2)
                _ = fi31
            }
        }
        v2 += 32768
        v2 += -2147483648
    } else {
        v6 += (v6 | (v2 + v6))
        switch "a" {
        case "foo bar":
            var v32 float64 = -2.25
            _ = v32
        case "hello":
            if (1e10 < -2.25) {
                for fi33 := 0; fi33 < 7; fi33++ {
                    var v34 float32 = 1.5
                    _ = fi33
                    _ = v34
                }
                v6 = v2
            } else {
                var v35 int = 2147483647
                _ = v35
            }
        case "":
            if (128 == 128) {
                switch 824 {
                case 2147483647:
                    for fi36 := 0; fi36 < 12; fi36++ {
                        v37 := GenS4{F0: "foo bar", F1: 446, F2: -2147483648}
                        v6 -= (v2 - v2)
                        _ = fi36
                        _ = v37
                    }
                case 528:
                    for fi38 := 0; fi38 < 16; fi38++ {
                        fmt.Println(fi38, fi38, fi38)
                        fmt.Println(v2, v2, fi38)
                        _ = fi38
                    }
                case 128:
                    switch v6 {
                    case 32768:
                        switch "foo bar" {
                        case "\n\t":
                            for fi39 := 0; fi39 < 8; fi39++ {
                                var v40 int8 = -128
                                _ = fi39
                                _ = v40
                            }
                        }
                    case -1:
                        if (false || (-32768 < 1000)) {
                            v5 -= 0
                        }
                    default:
                        var v41 int32 = 1
                        _ = v41
                    }
                }
                for fi42 := 0; fi42 < 4; fi42++ {
                    v43 := GenS5{F0: 4294967295}
                    if len(v4) > 0 {
                        v4[0] = v6
                    }
                    var v44 float64 = float64(-2.25)
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                var v45 int32 = 0
                _ = v45
            } else {
                v1[("a" + ("\n\t" + ("\n\t" + ("hello" + ("foo bar" + "")))))] = v2
            }
        }
    }
    v2 = -1
    if false {
        var v46 int = (v6 >> ((933) & 63))
        for fi47 := 0; fi47 < 18; fi47++ {
            var v48 int = 918
            v2 = 4294967295
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    v2 = v6
    var v49 uint8 = 0
    for fi50 := 0; fi50 < 11; fi50++ {
        switch v2 {
        case 32768:
            v2 -= fi50
        case 4294967295:
            v49 -= (v5 << ((849) & 63))
        default:
            var v51 float64 = float64(-2.25)
            _ = v51
        }
        _ = fi50
    }
    v49 = v49
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
