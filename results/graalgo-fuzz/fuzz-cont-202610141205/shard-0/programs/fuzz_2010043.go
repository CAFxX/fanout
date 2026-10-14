package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }
type GenS2 struct { F0 int16; F1 string; F2 string }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint, p1 uint32) int16 {
    fmt.Println(p1)
    fmt.Println(p1)
    var v6 int = 1
    for fi7 := 0; fi7 < 10; fi7++ {
        p0 += 0
        _ = fi7
    }
    v8 := []int{(v6 * (v6 >> ((1) & 63))), v6, v6}
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 0
}

func main() {
    var v1 uint8 = 128
    v2 := []string{"a", "foo bar", "", "\n\t"}
    var v3 float32 = 3.14159
    var v4 string = "hello"
    var v5 uint8 = (v1 | v1)
    switch "hello" {
    case "foo bar":
        var v9 int = -403
        _ = v9
    case "hello":
        var v10 int32 = -1
        _ = v10
    default:
        if true {
            for fi11 := 0; fi11 < 16; fi11++ {
                switch (fi11 + 988) {
                case -2147483648:
                    if false {
                        v3 = -2.25
                    }
                case -677:
                    v12 := GenS1{F0: -580, F1: -862, F2: int16(v1)}
                    _ = v12
                }
                switch (-(fi11)) {
                case 1:
                    switch (fi11 % ((fi11) | 1)) {
                    case -9223372036854775808:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case 4294967295:
                        fmt.Println(v3)
                    default:
                        _ = len(v2)
                    }
                }
                _ = fi11
            }
        } else {
            for fi13 := 0; fi13 < 10; fi13++ {
                v4 = (v4 + (v4 + ""))
                _ = fi13
            }
            var v14 uint = 4294967295
            for fi15 := 0; fi15 < 13; fi15++ {
                v16 := GenS1{F0: 42, F1: -1, F2: 712}
                _ = fi15
                _ = v16
            }
            _ = v14
        }
    }
    var v17 uint = 832
    var v18 uint = (^(v17))
    _ = len(v2)
    var v19 int = 127
    if false {
        for fi20 := 0; fi20 < 11; fi20++ {
            v3 = 1e10
            v21 := GenS1{F0: 1, F1: 9223372036854775807, F2: 315}
            _ = fi20
            _ = v21
        }
    } else {
        v1 = v1
        v22 := GenS2{F0: 1000, F1: (v4 + v4), F2: v4}
        var v23 string = ""
        _ = v22
        _ = v23
    }
    fmt.Println(v5, v19, v18)
    v19 += (v19 * 0)
    v2 = append(v2, (("a" + v4) + (v4 + v4)))
    v4 = "a"
    v24 := GenS1{F0: 552, F1: -44, F2: 0}
    switch v4 {
    case "":
        v25 := GenS2{F0: -32768, F1: (("" + v4) + "a"), F2: (v4 + v4)}
        _ = v25
    }
    if (v3 < v3) {
        var v26 int64 = -1
        if false {
            v18 = v17
            if true {
                v19 = 32768
                var v27 string = v4
                _ = v27
            }
            var v28 uint32 = 2147483647
            _ = v28
        } else {
            var v29 int = v19
            if true {
                for fi30 := 0; fi30 < 4; fi30++ {
                    if false {
                        switch 1 {
                        case 2147483647:
                            v19 -= v29
                        }
                        var v31 bool = true
                        switch "a" {
                        case "":
                            fmt.Println(v31, v31, v31)
                        case "a":
                            var v32 int16 = genFn0(166, 965)
                            _ = v32
                        }
                        _ = v31
                    } else {
                        v29 = (v29 << ((v18) & 63))
                    }
                    var v33 float64 = 1.5
                    _ = fi30
                    _ = v33
                }
                for fi34 := 0; fi34 < 4; fi34++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v35 := GenS2{F0: genFn0((v17 ^ v18), 0), F1: "", F2: v4}
                    _ = fi34
                    _ = v35
                }
                if (2147483647 >= 0) {
                    if false {
                        for fi36 := 0; fi36 < 20; fi36++ {
                            if true {
                                var v37 string = v4
                                _ = v37
                            }
                            var v38 bool = (true && true)
                            _ = fi36
                            _ = v38
                        }
                        var v39 int8 = -1
                        _ = v39
                    }
                } else {
                    v18 -= 42
                    for fi40 := 0; fi40 < 13; fi40++ {
                        v19 = v19
                        _ = fi40
                    }
                }
            } else {
                switch "hello" {
                case "hello":
                    for fi41 := 0; fi41 < 15; fi41++ {
                        fmt.Println(v29, v19)
                        if false {
                            v17 -= v18
                        }
                        switch fi41 {
                        case -9223372036854775808:
                            v2 = append(v2, v4)
                        }
                        _ = fi41
                    }
                case "":
                    for fi42 := 0; fi42 < 14; fi42++ {
                        v2 = append(v2, (((v4 + "hello") + (v4 + "\n\t")) + ""))
                        v1 = v5
                        _ = fi42
                    }
                default:
                    for fi43 := 0; fi43 < 10; fi43++ {
                        _ = len(v2)
                        fmt.Println(v17, v3, v26)
                        if true {
                            fmt.Println(v19, v3, fi43)
                        }
                        _ = fi43
                    }
                }
                v1 = v1
                v26 = v26
            }
            _ = v29
        }
        _ = v26
    }
    if (!false) {
        var v44 int64 = 723
        v45 := GenS3{F0: -114}
        _ = v44
        _ = v45
    } else {
        for fi46 := 0; fi46 < 14; fi46++ {
            v19 = (fi46 - (fi46 * 2147483647))
            v3 = (v3 + (1e10 * v3))
            _ = fi46
        }
        v3 -= 0.1
        for fi47 := 0; fi47 < 20; fi47++ {
            for fi48 := 0; fi48 < 20; fi48++ {
                _ = len(v2)
                v49 := GenS2{F0: -35, F1: (v4 + v4), F2: (v4 + ("" + ""))}
                _ = fi48
                _ = v49
            }
            _ = fi47
        }
    }
    v50 := GenS3{F0: -30}
    if true {
        if (((!false) && (-9223372036854775808 != -1)) || true) {
            v17 = (^(v17))
            var v51 bool = (!true)
            _ = v51
        }
        if false {
            v4 = "foo bar"
        } else {
            fmt.Println(v1, v4, v18)
            for fi52 := 0; fi52 < 2; fi52++ {
                for fi53 := 0; fi53 < 16; fi53++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = len(v2)
                    var v54 string = v4
                    v55 := GenS3{F0: 1}
                    fmt.Println(v3, v54)
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                fmt.Println(v19, v5)
                _ = fi52
            }
        }
        var v56 string = (v4 + "\n\t")
        _ = v56
    }
    var v57 int8 = -128
    v2 = append(v2, (v4 + v4))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v57
}
