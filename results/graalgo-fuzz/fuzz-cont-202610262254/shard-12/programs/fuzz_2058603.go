package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint64; F2 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    if false {
        for fi4 := 0; fi4 < 8; fi4++ {
            var v5 float64 = 1e10
            _ = fi4
            _ = v5
        }
        var v6 int16 = -173
        if false {
            var v7 int64 = int64(v6)
            if (311 == 32767) {
                switch 138 {
                case 128:
                    var v8 float64 = 3.14159
                    _ = v8
                case 42:
                    var v9 uint16 = 42
                    _ = v9
                case 4294967295:
                    var v10 int8 = 1
                    _ = v10
                default:
                    v6 += (v6 | 1000)
                }
            } else {
                v11 := GenS1{F0: 1, F1: 246}
                for fi12 := 0; fi12 < 1; fi12++ {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        if true {
                            v7 += v7
                            v7 = -1
                        } else {
                            v7 -= (v7 - (v7 % (((v7 - (v7 & (^(v7))))) | 1)))
                            var v14 int = fi13
                            v7 -= 42
                            _ = v14
                        }
                        v6 = v6
                        var v15 int16 = 0
                        _ = fi13
                        _ = v15
                    }
                    v7 -= v7
                    _ = fi12
                }
                _ = v11
            }
            _ = v7
        } else {
            _ = len(v1)
        }
        _ = v6
    } else {
        var v16 bool = false
        fmt.Println(v16)
        v16 = ((-348 <= 1) || v16)
        _ = v16
    }
    var v17 int8 = 0
    fmt.Println(v17)
    fmt.Println(v17)
    if false {
        for fi18 := 0; fi18 < 11; fi18++ {
            if false {
                v19 := []int{fi18, fi18, fi18, (fi18 + (fi18 / ((fi18) | 1)))}
                _ = v19
            }
            for fi20 := 0; fi20 < 14; fi20++ {
                if false {
                    v17 += v17
                    if (!(false || ((-9223372036854775808 < -1) && (!true)))) {
                        var v21 uint8 = 106
                        if (2 > 0) {
                            var v22 bool = true
                            _ = v22
                        }
                        var v23 string = ""
                        var v24 uint32 = 2147483648
                        var v25 string = "a"
                        delete(v1, (v24 % ((2147483647) | 1)))
                        _ = v21
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    v26 := []int{}
                    switch "hello" {
                    case "a":
                        delete(v1, 647)
                    case "\n\t":
                        var v27 string = "hello"
                        _ = v27
                    case "foo bar":
                        var v28 string = ""
                        _ = v28
                    }
                    _ = v26
                }
                var v29 float64 = float64(0.0)
                _ = fi20
                _ = v29
            }
            _ = fi18
        }
        if false {
            v30 := v3[128]
            if (0.1 < 0.1) {
                var v31 uint64 = 1
                for fi32 := 0; fi32 < 17; fi32++ {
                    if (0.0 <= ((0.0 / 1.0) / 1.5)) {
                        if ((3.14159 < 0.1) && (v30 && (((false && false) || ((v30 && false) || ((v17 >> ((1) & 63)) <= v17))) || ((516 != 4294967295) && v30)))) {
                            for fi33 := 0; fi33 < 17; fi33++ {
                                fmt.Println(fi33, v30)
                                v34 := GenS2{F0: -357}
                                _ = fi33
                                _ = v34
                            }
                        }
                        _ = len(v3)
                        var v35 uint8 = 128
                        _ = v35
                    } else {
                        v30 = (!(!false))
                    }
                    _ = fi32
                }
                _ = v31
            } else {
                v17 = -39
                if true {
                    fmt.Println(v30, v17)
                    var v36 uint8 = 2
                    _ = v36
                }
                v37 := v1[2]
                _ = v37
            }
            _ = v30
        } else {
            v38 := GenS3{F0: 1}
            switch "a" {
            case "foo bar":
                v39 := GenS1{F0: 2147483647, F1: 1}
                _ = v39
            }
            _ = v38
        }
    }
    var v40 uint8 = 1
    var v41 int = 254
    v42 := GenS3{F0: 0}
    v43 := []int{v41, v41, (^(v41))}
    if false {
        for fi44 := 0; fi44 < 12; fi44++ {
            switch fi44 {
            case 84:
                var v45 int32 = 1
                _ = v45
            }
            v41 = (v41 | (v41 * (fi44 << ((893) & 63))))
            v40 = (v40 & v40)
            _ = fi44
        }
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        var v47 uint16 = 42
        _ = fi46
        _ = v47
    }
    var v48 uint = 42
    if ((("hello" + "") + "") <= "") {
        delete(v3, (v41 >> ((v48) & 63)))
        if false {
            switch (("a" + "") + ("\n\t" + ("a" + (("a" + "foo bar") + "foo bar")))) {
            case "a":
                v49 := GenS4{F0: 2147483647, F1: 4294967295, F2: (9223372036854775807 != v41)}
                _ = v49
            }
        } else {
            switch ("\n\t" + "a") {
            case "foo bar":
                v41 = (v41 * 42)
            }
            var v50 float32 = -2.25
            var v51 int = 327
            _ = v50
            _ = v51
        }
        for fi52 := 0; fi52 < 8; fi52++ {
            switch ((("\n\t" + "\n\t") + "foo bar") + "foo bar") {
            case "foo bar":
                v17 = (v17 / ((int8(v41)) | 1))
            default:
                v3[128] = false
            }
            fmt.Println(v40)
            _ = fi52
        }
    } else {
        v41 = v41
    }
    fmt.Println(v41)
    var v53 uint64 = 230
    v41 = v41
    v43 = append(v43, 127)
    v54 := GenS4{F0: 160, F1: (v53 >> ((v48) & 63)), F2: (!(true && true))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v48:", v48)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v48
    _ = v53
    _ = v54
}
