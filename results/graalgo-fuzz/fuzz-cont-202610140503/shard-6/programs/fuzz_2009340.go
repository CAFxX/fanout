package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0() uint32 {
    var v4 int = -2147483648
    var v5 uint = 1
    fmt.Println(v5, v4)
    v6 := []int{(v4 >> (((v5 * (v5 & v5))) & 63)), v4, (v4 & v4)}
    _ = len(v6)
    _ = v4
    _ = v5
    _ = v6
    return 556
}
func genFn1(p0 uint8) bool {
    fmt.Println(p0)
    var v7 float64 = 0.1
    fmt.Println(v7)
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int64 = 9223372036854775807
        for fi10 := 0; fi10 < 12; fi10++ {
            switch fi8 {
            case 2147483647:
                v7 = v7
            default:
                v11 := GenS1{F0: fi8}
                _ = v11
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    var v12 uint32 = genFn0()
    _ = p0
    _ = v12
    _ = v7
    return false
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{}
    v3 := make(map[int]bool)
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(fi13)
        _ = fi13
    }
    if true {
        v14 := GenS2{F0: 42, F1: 4294967295}
        _ = v14
    } else {
        var v15 string = ("a" + ("foo bar" + "a"))
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        if true {
            var v17 int64 = -64
            _ = len(v2)
            v17 = (v17 / ((v17) | 1))
            _ = v17
        } else {
            if true {
                var v18 uint8 = 2
                if true {
                    _ = len(v1)
                    if (false || true) {
                        v18 -= (v18 & v18)
                        v18 -= (v18 >> ((4294967295) & 63))
                        v18 = v18
                    } else {
                        v18 = 156
                    }
                } else {
                    var v19 uint32 = 2147483648
                    _ = v19
                }
                switch "a" {
                case "foo bar":
                    _ = len(v3)
                case "a":
                    if true {
                        if ((0.0 + (3.14159 / 1.5)) >= (1.5 - 1.5)) {
                            var v20 uint = 0
                            v20 -= v20
                            _ = v20
                        }
                        switch "" {
                        case "":
                            v21 := GenS1{F0: (^(fi16))}
                            _ = v21
                        }
                        v18 -= v18
                    } else {
                        switch "a" {
                        case "\n\t":
                            v22 := GenS2{F0: fi16, F1: 42}
                            _ = v22
                        case "hello":
                            fmt.Println(v18, v18)
                        case "foo bar":
                            v23 := GenS1{F0: 42}
                            _ = v23
                        }
                    }
                case "":
                    fmt.Println(v18, fi16)
                default:
                    v24 := GenS1{F0: -2147483648}
                    _ = v24
                }
                _ = v18
            }
        }
        var v25 string = ("\n\t" + "foo bar")
        v25 = (("a" + v25) + v25)
        _ = fi16
        _ = v25
    }
    if (32767 > -904) {
        if false {
            var v26 uint = 4294967295
            if false {
                if (!((-2.25 < -2.25) || (v26 < (v26 >> ((141) & 63))))) {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        fmt.Println(v26, v26)
                        _ = fi27
                    }
                } else {
                    v26 = 299
                }
                var v28 int8 = -1
                var v29 string = ("a" + "a")
                _ = v28
                _ = v29
            }
            _ = v26
        } else {
            v30 := GenS2{F0: -1, F1: 2}
            _ = len(v1)
            _ = v30
        }
    }
    var v31 int = -830
    switch "a" {
    case "\n\t":
        var v32 float64 = 3.14159
        _ = v32
    case "":
        delete(v3, -1)
    default:
        fmt.Println(v31)
    }
    var v33 int8 = 37
    v33 -= v33
    v31 -= v31
    if true {
        var v34 int64 = 148
        v33 -= v33
        if (true && (105 == 4294967295)) {
            var v35 int8 = v33
            _ = v35
        } else {
            v36 := v1[2147483647]
            if false {
                fmt.Println(v34, v31)
                v33 = (v33 << ((255) & 63))
            } else {
                switch v36 {
                case "hello":
                    var v37 int8 = (v33 % ((v33) | 1))
                    _ = v37
                case "":
                    for fi38 := 0; fi38 < 8; fi38++ {
                        if true {
                            v33 = 100
                            for fi39 := 0; fi39 < 20; fi39++ {
                                _ = len(v2)
                                if false {
                                    v40 := GenS1{F0: 2147483647}
                                    _ = v40
                                }
                                _ = fi39
                            }
                        }
                        v34 += v34
                        _ = fi38
                    }
                default:
                    for fi41 := 0; fi41 < 4; fi41++ {
                        v42 := GenS3{F0: 2}
                        v34 -= v34
                        fmt.Println(v36)
                        _ = fi41
                        _ = v42
                    }
                }
                if true {
                    var v43 float32 = 0.1
                    fmt.Println(v43)
                    var v44 int = v31
                    _ = v43
                    _ = v44
                } else {
                    var v45 uint16 = 255
                    var v46 uint16 = (v45 / ((255) | 1))
                    _ = v45
                    _ = v46
                }
            }
            _ = v36
        }
        _ = v34
    }
    var v47 int8 = (v33 & (v33 - 100))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v47
}
