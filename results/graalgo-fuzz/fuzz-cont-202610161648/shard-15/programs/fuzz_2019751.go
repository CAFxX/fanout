package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int16 }
type GenS2 struct { F0 string; F1 int; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 uint8; F2 int8 }
type GenS5 struct { F0 int8 }
func genFn0(p0 bool, p1 uint8, p2 uint16) uint {
    fmt.Println(p2, p2, p1)
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 uint = 1
        p0 = p0
        _ = fi4
        _ = v5
    }
    var v6 float64 = -2.25
    v6 += v6
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 226
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{32768, -758}
    v3 := make(map[string]bool)
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int64 = -481
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int64 = v8
            v11 := GenS1{F0: 1, F1: (genFn0((-842 <= 42), 0, 866) != genFn0(((v10 / ((v8) | 1)) < 0), 255, 1)), F2: -39}
            v12 := GenS2{F0: "foo bar", F1: (fi7 - 9223372036854775807), F2: false}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        if (false || false) {
            switch (fi13 & (^(fi13))) {
            case 0:
                var v14 uint16 = 65535
                _ = v14
            default:
                if true {
                    fmt.Println(fi13)
                    if (!true) {
                        if (!((!false) && true)) {
                            if ((!(true && true)) && (42 >= 4294967295)) {
                                var v15 string = ("foo bar" + "a")
                                fmt.Println(v15)
                                _ = v15
                            }
                        }
                        if true {
                            for fi16 := 0; fi16 < 8; fi16++ {
                                var v17 uint64 = 837
                                delete(v1, 922)
                                v17 += (v17 ^ (v17 ^ (v17 / ((929) | 1))))
                                if len(v2) > 0 {
                                    v2[0] = (fi16 - (fi16 << ((1) & 63)))
                                }
                                v17 -= v17
                                _ = fi16
                                _ = v17
                            }
                            if true {
                                if (859 == 42) {
                                    fmt.Println(fi13)
                                    fmt.Println(fi13)
                                } else {
                                    fmt.Println(fi13)
                                    var v18 int = 127
                                    var v19 bool = (true && (1 != 258))
                                    _ = v18
                                    _ = v19
                                }
                            } else {
                                _ = len(v2)
                            }
                            v20 := GenS3{F0: 298}
                            _ = v20
                        }
                    } else {
                        if (!(65535 != 65535)) {
                            switch fi13 {
                            case -372:
                                if (true || (-128 != 24)) {
                                    fmt.Println(fi13)
                                    _ = len(v3)
                                } else {
                                    for fi21 := 0; fi21 < 17; fi21++ {
                                        _ = len(v3)
                                        var v22 uint64 = 1
                                        var v23 int32 = 0
                                        _ = fi21
                                        _ = v22
                                        _ = v23
                                    }
                                    v3[(("foo bar" + "\n\t") + "hello")] = (true && ((false && false) && false))
                                    for fi24 := 0; fi24 < 2; fi24++ {
                                        v25 := GenS3{F0: 2}
                                        switch "" {
                                        case "":
                                            switch (fi13 + fi24) {
                                            case 127:
                                                delete(v1, 2147483648)
                                            case 627:
                                                if false {
                                                    if (true || true) {
                                                        var v26 uint64 = 725
                                                        _ = v26
                                                    }
                                                    v1[4294967295] = 127
                                                    var v27 string = "foo bar"
                                                    _ = v27
                                                } else {
                                                    _ = len(v1)
                                                    v28 := GenS1{F0: 213, F1: (2 < 143), F2: 0}
                                                    _ = v28
                                                }
                                            }
                                        default:
                                            v29 := GenS1{F0: 2147483647, F1: (false || true), F2: -976}
                                            _ = v29
                                        }
                                        fmt.Println(fi24, fi13)
                                        _ = fi24
                                        _ = v25
                                    }
                                }
                            case -249:
                                var v30 uint64 = 2
                                _ = v30
                            case -9223372036854775808:
                                _ = len(v3)
                            }
                            _ = len(v2)
                        } else {
                            if (1 > 65535) {
                                for fi31 := 0; fi31 < 5; fi31++ {
                                    var v32 uint16 = 123
                                    v32 -= (v32 + v32)
                                    _ = fi31
                                    _ = v32
                                }
                                var v33 int64 = -1
                                _ = v33
                            } else {
                                v34 := v1[2147483648]
                                _ = v34
                            }
                        }
                        var v35 int8 = 42
                        v2 = append(v2, (^(fi13)))
                        _ = v35
                    }
                } else {
                    switch ("a" + "a") {
                    case "foo bar":
                        v36 := GenS3{F0: 894}
                        _ = v36
                    default:
                        switch "a" {
                        case "a":
                            var v37 bool = (0.0 < float64(float32(-2.25)))
                            _ = v37
                        case "\n\t":
                            var v38 int16 = 1000
                            _ = v38
                        default:
                            fmt.Println(fi13)
                        }
                    }
                    delete(v3, ("a" + ("a" + "\n\t")))
                }
            }
            v39 := GenS4{F0: 0, F1: 0, F2: -1}
            _ = v39
        } else {
            delete(v3, ("" + "\n\t"))
        }
        var v40 int64 = -1
        for fi41 := 0; fi41 < 11; fi41++ {
            var v42 uint8 = 127
            _ = fi41
            _ = v42
        }
        _ = fi13
        _ = v40
    }
    v1[2] = 32768
    switch 517 {
    case 897:
        var v43 uint64 = 42
        _ = v43
    }
    var v44 bool = false
    var v45 float32 = 1.5
    v46 := GenS5{F0: -1}
    switch 2147483647 {
    case 2147483647:
        var v47 int = 27
        _ = v47
    case 0:
        v48 := GenS1{F0: 863, F1: (!v44), F2: 0}
        _ = v48
    case -622:
        fmt.Println(v44)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v46
}
