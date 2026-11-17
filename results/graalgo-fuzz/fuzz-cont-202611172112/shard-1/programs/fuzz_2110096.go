package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int; F1 uint8; F2 int32 }
type GenS4 struct { F0 bool; F1 uint; F2 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint16) int8 {
    var v7 uint = 4294967295
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 float64 = 0.1
        if false {
            v10 := GenS1{F0: 1, F1: 232}
            _ = v10
        }
        if false {
            v11 := GenS2{F0: fi8, F1: 908}
            fmt.Println(v7)
            v7 -= 497
            _ = v11
        } else {
            var v12 int = fi8
            p0 += p0
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    switch 32768 {
    case -1:
        p0 -= (^(p0))
    case -2147483648:
        v13 := []int{326, 128}
        _ = v13
    case 9223372036854775807:
        switch ("\n\t" + ("foo bar" + "foo bar")) {
        case "foo bar":
            switch ("hello" + "") {
            case "\n\t":
                if (true && false) {
                    var v14 bool = false
                    _ = v14
                } else {
                    if false {
                        p0 += 154
                        p0 = p0
                    } else {
                        var v15 int = 32767
                        fmt.Println(v7, v15, v7)
                        _ = v15
                    }
                }
            case "foo bar":
                fmt.Println(v7, p0)
            }
        default:
            for fi16 := 0; fi16 < 6; fi16++ {
                for fi17 := 0; fi17 < 17; fi17++ {
                    p0 += 332
                    v18 := []int{(fi16 & (fi17 << ((65535) & 63)))}
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        }
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        if false {
            p0 = p0
        }
        _ = fi19
    }
    _ = p0
    _ = v7
    return -1
}

func main() {
    var v1 string = ("\n\t" + "hello")
    v2 := []string{""}
    v3 := []uint32{0}
    v4 := make(map[int]int)
    var v5 uint8 = 128
    v6 := make(map[string]uint32)
    v20 := GenS3{F0: -516, F1: (v5 * v5), F2: -1}
    var v21 int16 = 32767
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(v1, v1, v1)
        if true {
            var v23 uint = 255
            _ = v23
        }
        fmt.Println(v21, v1)
        _ = fi22
    }
    _ = len(v2)
    fmt.Println(v1, v1, v1)
    for fi24 := 0; fi24 < 20; fi24++ {
        if true {
            v25 := GenS3{F0: (fi24 ^ fi24), F1: v5, F2: 438}
            _ = v25
        } else {
            var v26 int = 870
            v26 += (fi24 & v26)
            v5 = v5
            _ = v26
        }
        if (!(false && (899 > -1))) {
            if ((v1 + v1) < (v1 + v1)) {
                for fi27 := 0; fi27 < 5; fi27++ {
                    v21 += 1000
                    switch 9223372036854775807 {
                    case -660:
                        v28 := GenS4{F0: (true && true), F1: 255, F2: -50}
                        _ = v28
                    case 127:
                        if true {
                            v29 := GenS1{F0: 2147483647, F1: v5}
                            _ = v29
                        }
                    case -708:
                        var v30 string = ("foo bar" + v1)
                        _ = v30
                    default:
                        switch (fi27 >> ((312) & 63)) {
                        case 703:
                            var v31 uint8 = 128
                            _ = v31
                        case -2147483648:
                            v5 += v5
                        default:
                            var v32 float64 = 3.14159
                            _ = v32
                        }
                    }
                    var v33 uint32 = 2
                    fmt.Println(v5)
                    switch fi27 {
                    case 471:
                        v1 = "a"
                    case 127:
                        var v34 uint = 18446744073709551615
                        _ = v34
                    case 0:
                        v35 := GenS5{F0: v33}
                        _ = v35
                    default:
                        _ = len(v2)
                    }
                    _ = fi27
                    _ = v33
                }
            }
            v1 = "hello"
            if ((3.14159 - (-2.25 - 3.14159)) >= 3.14159) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v36 uint8 = v5
                _ = v36
            }
        }
        _ = fi24
    }
    v21 = (v21 | 1)
    switch -9223372036854775808 {
    case 127:
        v6["hello"] = 65535
    default:
        if ((v5 >> ((255) & 63)) != 1) {
            v37 := GenS4{F0: false, F1: 18446744073709551615, F2: -89}
            _ = v37
        } else {
            if false {
                v38 := GenS3{F0: 128, F1: v5, F2: 167}
                v5 = 42
                _ = v38
            } else {
                for fi39 := 0; fi39 < 2; fi39++ {
                    var v40 int16 = (-(v21))
                    _ = fi39
                    _ = v40
                }
            }
        }
    }
    var v41 uint16 = 32767
    v42 := v4[4294967295]
    for fi43 := 0; fi43 < 16; fi43++ {
        switch (^(v42)) {
        case 128:
            v5 += (^(v5))
        case -2147483648:
            v1 = v1
        case 127:
            v44 := GenS6{F0: false, F1: (v41 % ((0) | 1))}
            _ = v44
        }
        if (48 == -13) {
            for fi45 := 0; fi45 < 9; fi45++ {
                var v46 string = v1
                _ = fi45
                _ = v46
            }
        }
        _ = fi43
    }
    v41 += (v41 + v41)
    fmt.Println(v5, v42, v21)
    delete(v4, v42)
    for fi47 := 0; fi47 < 12; fi47++ {
        v48 := v6[("foo bar" + v1)]
        v49 := v4[v42]
        _ = fi47
        _ = v48
        _ = v49
    }
    if false {
        var v50 uint = 691
        fmt.Println(v1, v1)
        var v51 int8 = 42
        _ = v50
        _ = v51
    }
    v1 = v1
    v52 := v6[((((((v1 + ("\n\t" + v1)) + ((("\n\t" + (v1 + ((v1 + ("\n\t" + v1)) + v1))) + "foo bar") + (v1 + (v1 + "hello")))) + (v1 + "foo bar")) + "a") + v1) + v1)]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v52
    _ = v6
}
