package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int; F1 uint; F2 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64 }
func genFn0(p0 string, p1 int64, p2 int8) string {
    if (false && (false || (42 > 42))) {
        var v7 float32 = (0.0 - float32((0.0 + 3.14159)))
        _ = v7
    } else {
        var v8 uint16 = 32767
        if (false && ((p1 / ((-472) | 1)) > (p1 % ((9223372036854775807) | 1)))) {
            var v9 uint8 = 0
            _ = v9
        }
        var v10 bool = (!true)
        _ = v10
        _ = v8
    }
    if false {
        switch 128 {
        case 4294967295:
            v11 := []int{0, -9223372036854775808, 128}
            _ = v11
        case 0:
            for fi12 := 0; fi12 < 1; fi12++ {
                for fi13 := 0; fi13 < 7; fi13++ {
                    var v14 float32 = -2.25
                    switch (p0 + "foo bar") {
                    case "\n\t":
                        p2 = (p2 + -125)
                    case "hello":
                        fmt.Println(fi12)
                    }
                    var v15 int64 = (^(p1))
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = fi12
            }
        case 87:
            v16 := GenS1{F0: p0, F1: -10, F2: 905}
            _ = v16
        }
    } else {
        var v17 int32 = -2147483648
        if false {
            var v18 uint = 65535
            v19 := []int{128, 387, 2147483647, 128}
            _ = v18
            _ = v19
        }
        if false {
            v20 := []int{-724, 0, 32768, -9223372036854775808}
            _ = v20
        } else {
            p2 -= p2
        }
        _ = v17
    }
    if false {
        v21 := []int{42, 2147483647, -9223372036854775808, 386}
        var v22 uint64 = 337
        p2 = -1
        _ = v21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}

func main() {
    var v1 uint64 = 769
    var v2 uint64 = (v1 | (v1 << ((65535) & 63)))
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    var v5 uint = 42
    v6 := make(map[int]string)
    v23 := GenS2{F0: 1}
    for fi24 := 0; fi24 < 2; fi24++ {
        if (("a" + "hello") != "") {
            switch ("a" + "foo bar") {
            case "":
                var v25 int16 = 42
                _ = v25
            case "hello":
                if ((fi24 * (fi24 % (((-(fi24))) | 1))) == 2147483647) {
                    v26 := []int{(fi24 & (fi24 % ((-1) | 1))), fi24, 4294967295, (fi24 / (((fi24 % ((fi24) | 1))) | 1))}
                    if false {
                        for fi27 := 0; fi27 < 5; fi27++ {
                            if ((v5 >> ((0) & 63)) <= (v5 + 894)) {
                                fmt.Println(fi27)
                                switch fi27 {
                                case -140:
                                    switch fi27 {
                                    case 127:
                                        if (!(!true)) {
                                            fmt.Println(v5, v1)
                                        }
                                    case 453:
                                        if (true && true) {
                                            if false {
                                                var v28 int16 = -32768
                                                _ = v28
                                            }
                                            fmt.Println(fi24, v5)
                                            var v29 string = "\n\t"
                                            _ = v29
                                        }
                                    }
                                case -2147483648:
                                    var v30 float64 = 1.5
                                    _ = v30
                                }
                                v26 = append(v26, (fi24 % ((32768) | 1)))
                            } else {
                                v2 -= (^(v2))
                            }
                            var v31 int = (fi27 % ((fi27) | 1))
                            _ = fi27
                            _ = v31
                        }
                        var v32 float64 = (3.14159 + (3.14159 / 1e10))
                        _ = v32
                    }
                    v33 := GenS3{F0: 65535}
                    _ = v26
                    _ = v33
                } else {
                    if false {
                        fmt.Println(fi24, fi24)
                    } else {
                        v34 := []int{42, (^(fi24)), -1}
                        if true {
                            v1 = v1
                        }
                        v2 = v1
                        _ = v34
                    }
                }
            }
            fmt.Println(fi24)
            for fi35 := 0; fi35 < 7; fi35++ {
                v2 -= 906
                v36 := GenS3{F0: 713}
                var v37 int64 = -787
                _ = fi35
                _ = v36
                _ = v37
            }
        } else {
            v5 -= 42
            var v38 int32 = 42
            var v39 uint32 = 2147483647
            _ = v38
            _ = v39
        }
        _ = fi24
    }
    v40 := GenS4{F0: 32767}
    for fi41 := 0; fi41 < 11; fi41++ {
        if false {
            fmt.Println(v5)
            for fi42 := 0; fi42 < 12; fi42++ {
                switch "" {
                case "foo bar":
                    var v43 uint8 = uint8(v1)
                    _ = v43
                case "hello":
                    var v44 float64 = float64(((-2.25 + 0.0) + (1.5 - 1.5)))
                    _ = v44
                case "":
                    if (1e10 > float64((-2.25 * (float32(-2.25) / ((((((3.14159 - (-2.25 / 0.1)) - -2.25) * (float32(float64((3.14159 * 1.5))) * (0.0 / -2.25))) / 1e10) / 1.5) - 1.5))))) {
                        fmt.Println(v5, v1, fi41)
                        switch 2147483647 {
                        case 42:
                            fmt.Println(v2, v2, fi41)
                        case 1:
                            v5 = (v5 * (v5 | v5))
                        default:
                            v2 += v2
                        }
                        v1 = (v1 % ((v1) | 1))
                    }
                }
                v1 = v1
                if true {
                    if (!(42 > 255)) {
                        delete(v4, -356)
                        v2 = v2
                        v5 -= uint(v2)
                    } else {
                        switch ("foo bar" + ("hello" + "a")) {
                        case "hello":
                            var v45 uint32 = 218
                            _ = v45
                        }
                        v46 := GenS1{F0: "a", F1: 775, F2: (v2 + (v2 - 42))}
                        v47 := []int{fi41}
                        var v48 uint16 = 451
                        if ((!(v5 == (v5 % ((v5) | 1)))) && (true && (((v5 % ((0) | 1)) == (v5 ^ uint(v48))) && (true && false)))) {
                            var v49 uint8 = 127
                            if true {
                                fmt.Println(v49)
                            } else {
                                v50 := GenS5{F0: ((v49 & uint8(v5)) != v49), F1: (v49 % ((v49) | 1))}
                                v51 := GenS2{F0: (^(v49))}
                                var v52 int = fi42
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                            if len(v47) >= 2 {
                                v47 = v47[1:len(v47)]
                            }
                            var v53 int16 = 510
                            _ = v49
                            _ = v53
                        }
                        var v54 int64 = 955
                        _ = v46
                        _ = v47
                        _ = v48
                        _ = v54
                    }
                }
                _ = fi42
            }
        }
        _ = fi41
    }
    delete(v3, uint32(v5))
    v55 := GenS3{F0: 32767}
    v5 += v5
    if true {
        var v56 bool = false
        _ = v56
    } else {
        v57 := GenS4{F0: 127}
        var v58 uint8 = 195
        v58 = 0
        _ = v57
        _ = v58
    }
    var v59 int16 = 1000
    var v60 int16 = (v59 + -560)
    var v61 int64 = 0
    v62 := GenS6{F0: 2147483647, F1: (v5 * (^(v5))), F2: (v61 / (((v61 & int64(v5))) | 1))}
    v63 := GenS7{F0: -15}
    v60 += 42
    for fi64 := 0; fi64 < 7; fi64++ {
        switch 2147483647 {
        case 127:
            v65 := GenS6{F0: fi64, F1: v5, F2: v61}
            _ = v65
        }
        v66 := GenS8{F0: v1}
        _ = fi64
        _ = v66
    }
    var v67 int16 = (v59 * (v60 << ((v5) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v55
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
    _ = v63
    _ = v67
}
