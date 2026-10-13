package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 uint64; F1 uint32; F2 uint16 }

func main() {
    var v1 int8 = -9
    var v2 float64 = (3.14159 / 1e10)
    var v3 int = 629
    if (false && false) {
        v4 := GenS1{F0: true, F1: 115, F2: (v1 % (((v1 ^ -128)) | 1))}
        var v5 uint64 = 703
        v5 += v5
        _ = v4
        _ = v5
    }
    switch "a" {
    case "foo bar":
        v1 = 11
    case "":
        v6 := []int{(v3 >> ((1) & 63)), v3, 4294967295}
        _ = v6
    case "hello":
        if (-459 <= -1) {
            switch "hello" {
            case "hello":
                var v7 string = ((("hello" + "\n\t") + "foo bar") + ("hello" + "\n\t"))
                _ = v7
            case "a":
                var v8 int32 = 323
                _ = v8
            case "":
                if true {
                    var v9 int64 = 399
                    var v10 int = v3
                    _ = v10
                    _ = v9
                }
            default:
                v11 := []int{}
                _ = v11
            }
            switch v3 {
            case 1:
                v2 += ((float64(-2.25) / v2) * (((((float64(-2.25) * v2) * (v2 - float64((-2.25 - (((3.14159 / -2.25) - -2.25) / 3.14159))))) / 1.0) - v2) + (v2 + 1.5)))
            default:
                v12 := []int{}
                _ = v12
            }
            var v13 int32 = -23
            _ = v13
        } else {
            var v14 uint = 0
            switch ("\n\t" + "a") {
            case "foo bar":
                v1 = (v1 + -1)
            case "\n\t":
                var v15 float32 = (1e10 / 1e10)
                _ = v15
            case "":
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 int = (v3 & v3)
                    var v18 int = -320
                    for fi19 := 0; fi19 < 18; fi19++ {
                        switch fi19 {
                        case -1:
                            var v20 uint = v14
                            _ = v20
                        case 4294967295:
                            v18 = -1
                        }
                        if true {
                            var v21 float64 = 0.0
                            if (v2 != (-2.25 * v2)) {
                                if (229 <= 2) {
                                    var v22 int = fi19
                                    _ = v22
                                } else {
                                    var v23 string = "\n\t"
                                    switch v23 {
                                    case "":
                                        v24 := []int{fi16}
                                        _ = v24
                                    }
                                    v25 := []int{}
                                    _ = v23
                                    _ = v25
                                }
                                v18 = 42
                                v26 := GenS2{F0: (v17 ^ 9223372036854775807)}
                                _ = v26
                            }
                            fmt.Println(fi19, v18, fi16)
                            switch "foo bar" {
                            case "hello":
                                v27 := []int{(^(v3))}
                                var v28 float32 = 1.5
                                fmt.Println(v2)
                                _ = v27
                                _ = v28
                            }
                            v29 := []int{}
                            _ = v21
                            _ = v29
                        } else {
                            v3 += 9223372036854775807
                            v30 := []int{66, -279}
                            if true {
                                if true {
                                    if true {
                                        var v31 float64 = ((v2 + -2.25) + v2)
                                        var v32 int8 = 119
                                        _ = v31
                                        _ = v32
                                    }
                                } else {
                                    if len(v30) > 0 {
                                        v30[0] = v18
                                    }
                                    v30 = append(v30, (v3 >> ((551) & 63)))
                                }
                            } else {
                                v33 := GenS1{F0: (("" >= ("hello" + "foo bar")) && true), F1: 113, F2: (v1 << (((^(v14))) & 63))}
                                _ = v33
                            }
                            _ = v30
                        }
                        fmt.Println(v14, fi16)
                        _ = fi19
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            default:
                v3 = (v3 >> (((v14 ^ v14)) & 63))
            }
            switch v3 {
            case 2147483647:
                if true {
                    var v34 uint32 = 2147483648
                    _ = v34
                }
            default:
                v1 += (v1 - v1)
            }
            _ = v14
        }
    default:
        switch (v3 ^ 42) {
        case 128:
            v35 := GenS2{F0: (v3 + v3)}
            _ = v35
        case -841:
            v1 = (v1 >> ((65535) & 63))
        case -914:
            var v36 float32 = (0.0 + 3.14159)
            _ = v36
        }
    }
    v37 := []int{}
    for fi38 := 0; fi38 < 16; fi38++ {
        var v39 uint16 = 1
        for fi40 := 0; fi40 < 13; fi40++ {
            v37 = append(v37, (fi40 ^ int(v39)))
            var v41 uint64 = 1
            var v42 uint = 1
            _ = fi40
            _ = v41
            _ = v42
        }
        for fi43 := 0; fi43 < 13; fi43++ {
            v39 += 1
            if ("foo bar" <= "hello") {
                if true {
                    v44 := GenS3{F0: true}
                    _ = v44
                }
                var v45 float64 = v2
                var v46 string = "foo bar"
                _ = v45
                _ = v46
            }
            v47 := GenS4{F0: (("foo bar" + "\n\t") + "foo bar")}
            v48 := GenS5{F0: -9223372036854775808, F1: 108}
            _ = fi43
            _ = v47
            _ = v48
        }
        _ = fi38
        _ = v39
    }
    v37 = append(v37, v3)
    for fi49 := 0; fi49 < 2; fi49++ {
        if len(v37) >= 2 {
            v37 = v37[1:len(v37)]
        }
        v50 := GenS6{F0: 755, F1: 399, F2: 65535}
        if true {
            if len(v37) > 0 {
                v37[0] = -1
            }
            v2 = (v2 * -2.25)
        }
        _ = fi49
        _ = v50
    }
    if len(v37) > 0 {
        v37[0] = 128
    }
    if true {
        if false {
            fmt.Println(v3, v1, v1)
        } else {
            var v51 uint = 568
            v1 = (^(v1))
            v2 += (v2 - ((v2 - 0.0) + 1.5))
            _ = v51
        }
        v3 = v3
        v1 += -128
    } else {
        for fi52 := 0; fi52 < 8; fi52++ {
            var v53 int = fi52
            fmt.Println(v3)
            _ = fi52
            _ = v53
        }
    }
    var v54 uint32 = 2147483647
    v37 = append(v37, (v3 + 2147483647))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v54
}
