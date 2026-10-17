package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint8; F1 uint8 }
func genFn0() int64 {
    v6 := GenS1{F0: 286, F1: 1, F2: 0}
    switch 216 {
    case 1:
        v7 := GenS1{F0: 42, F1: 65535, F2: -9223372036854775808}
        _ = v7
    }
    var v8 uint32 = 2147483648
    fmt.Println(v8)
    var v9 int16 = 42
    _ = v6
    _ = v8
    _ = v9
    return 42
}
func genFn1() int16 {
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint16 = 65535
            fmt.Println(v12)
            v12 -= 255
            _ = fi11
            _ = v12
        }
        var v13 uint8 = 42
        v13 = v13
        _ = fi10
        _ = v13
    }
    var v14 int16 = 32767
    if false {
        v15 := []int{}
        switch "hello" {
        case "a":
            switch 548 {
            case -2147483648:
                v14 -= (v14 & (-(v14)))
            case -814:
                if len(v15) > 0 {
                    v15[0] = 42
                }
            }
        case "":
            var v16 int8 = -77
            _ = v16
        case "foo bar":
            fmt.Println(v14)
        }
        var v17 uint32 = 65535
        _ = v15
        _ = v17
    } else {
        v14 += (v14 * 1000)
    }
    v14 = -207
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := []int{(fi18 & 9223372036854775807), -1}
        _ = fi18
        _ = v19
    }
    _ = v14
    return (v14 | v14)
}

func main() {
    var v1 int16 = -750
    var v2 uint64 = 1
    v3 := []string{"\n\t"}
    var v4 int64 = -9223372036854775808
    var v5 float32 = 3.14159
    v2 += v2
    switch 4294967295 {
    case 811:
        v20 := GenS2{F0: -1}
        _ = v20
    case 32767:
        v4 = (v4 | -613)
    case -977:
        if ((v5 >= -2.25) || ((!false) || (false || (80 == 0)))) {
            for fi21 := 0; fi21 < 7; fi21++ {
                fmt.Println(v5, v4)
                _ = len(v3)
                switch (("" + "hello") + "hello") {
                case "hello":
                    var v22 int = 9223372036854775807
                    _ = v22
                case "":
                    if true {
                        v23 := GenS1{F0: 2, F1: 42, F2: int64(fi21)}
                        if ((v1 % ((genFn1()) | 1)) == genFn1()) {
                            var v24 float64 = 1.5
                            _ = v24
                        }
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            for fi25 := 0; fi25 < 11; fi25++ {
                                var v26 uint32 = 2
                                _ = fi25
                                _ = v26
                            }
                            v4 = v4
                        } else {
                            var v27 uint32 = 2147483647
                            v4 = 42
                            _ = v27
                        }
                        _ = v23
                    } else {
                        if false {
                            fmt.Println(v4)
                            for fi28 := 0; fi28 < 8; fi28++ {
                                if (1 != 605) {
                                    var v29 float64 = (3.14159 + 1.5)
                                    var v30 bool = false
                                    if (255 < 0) {
                                        v29 = v29
                                        var v31 float64 = v29
                                        var v32 uint8 = 1
                                        _ = v31
                                        _ = v32
                                    }
                                    fmt.Println(v30, fi28)
                                    fmt.Println(v29, v1)
                                    _ = v29
                                    _ = v30
                                } else {
                                    v33 := GenS3{F0: -247}
                                    var v34 string = ""
                                    var v35 bool = (!(!(255 < 0)))
                                    if false {
                                        var v36 int8 = 1
                                        if v35 {
                                            var v37 int8 = v36
                                            v38 := GenS1{F0: 799, F1: 912, F2: (v4 + v4)}
                                            _ = v37
                                            _ = v38
                                        }
                                        v39 := GenS4{F0: 128, F1: v1}
                                        _ = v36
                                        _ = v39
                                    } else {
                                        v35 = (1 >= 42)
                                        var v40 int32 = -417
                                        _ = v40
                                    }
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                _ = fi28
                            }
                        }
                        v41 := GenS5{F0: (fi21 * 257)}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v41
                    }
                case "a":
                    v42 := GenS4{F0: 2, F1: (v1 + v1)}
                    _ = v42
                }
                _ = fi21
            }
            switch -275 {
            case 78:
                var v43 int16 = genFn1()
                _ = v43
            }
        }
    }
    v44 := GenS6{F0: 517, F1: 227}
    if (!false) {
        fmt.Println(v5, v1)
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
    } else {
        if false {
            _ = len(v3)
            var v45 uint = 179
            _ = v45
        } else {
            v4 = (-(v4))
        }
    }
    v4 = (v4 * -786)
    switch -9223372036854775808 {
    case -9223372036854775808:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 127:
        var v46 int64 = 0
        _ = v46
    case 259:
        _ = len(v3)
    default:
        if (-919 != 863) {
            switch -1 {
            case 127:
                v4 = -117
            case -275:
                var v47 int16 = genFn1()
                _ = v47
            case 265:
                var v48 uint32 = 4294967295
                _ = v48
            }
            v49 := GenS7{F0: 0, F1: 103}
            _ = v49
        }
    }
    v2 -= (v2 * v2)
    v4 = (-(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
