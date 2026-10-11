package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int64; F1 int8; F2 uint8 }
func genFn0() uint {
    if true {
        var v7 float64 = 1e10
        _ = v7
    } else {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 float64 = 1e10
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 uint = 2
            for fi12 := 0; fi12 < 9; fi12++ {
                v11 = v11
                if false {
                    var v13 string = "foo bar"
                    switch "a" {
                    case "":
                        fmt.Println(fi10, v13)
                    }
                    _ = v13
                }
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
    }
    if false {
        var v14 uint32 = 73
        _ = v14
    } else {
        v15 := GenS1{F0: 0}
        for fi16 := 0; fi16 < 3; fi16++ {
            if (255 == 2) {
                fmt.Println(fi16)
                if false {
                    var v17 int32 = 0
                    fmt.Println(v17)
                    switch fi16 {
                    case -2147483648:
                        var v18 bool = (false || false)
                        _ = v18
                    case -726:
                        v17 = int32(fi16)
                    default:
                        var v19 int32 = (^(v17))
                        _ = v19
                    }
                    _ = v17
                } else {
                    v20 := GenS1{F0: -469}
                    _ = v20
                }
            } else {
                v21 := []int{(^(fi16)), fi16, (fi16 & (^(fi16))), fi16}
                fmt.Println(fi16)
                for fi22 := 0; fi22 < 19; fi22++ {
                    var v23 uint8 = 203
                    if len(v21) > 0 {
                        v21[0] = 936
                    }
                    _ = fi22
                    _ = v23
                }
                _ = v21
            }
            for fi24 := 0; fi24 < 12; fi24++ {
                var v25 uint16 = 205
                v25 = 65535
                if true {
                    var v26 uint32 = 2147483647
                    _ = v26
                } else {
                    if (1 >= 2147483647) {
                        v27 := GenS2{F0: false, F1: 4}
                        var v28 int16 = 1
                        v28 -= v28
                        _ = v27
                        _ = v28
                    } else {
                        v29 := GenS3{F0: 1}
                        fmt.Println(v25)
                        _ = v29
                    }
                    var v30 int32 = -2147483648
                    switch "foo bar" {
                    case "a":
                        v25 += v25
                    }
                    _ = v30
                }
                _ = fi24
                _ = v25
            }
            _ = fi16
        }
        _ = v15
    }
    var v31 uint16 = 255
    v32 := []int{127, 32767, -911, 1}
    fmt.Println(v31)
    _ = v31
    _ = v32
    return 1
}
func genFn1(p0 bool, p1 int, p2 uint8) uint8 {
    if p0 {
        if false {
            var v33 int32 = 0
            _ = v33
        } else {
            var v34 int = p1
            if ((p0 && ((("" + "") + "\n\t") >= (("a" + "hello") + "hello"))) && (!p0)) {
                fmt.Println(p2, p2, v34)
                switch "\n\t" {
                case "hello":
                    switch p1 {
                    case 253:
                        if false {
                            var v35 int = v34
                            _ = v35
                        } else {
                            if ((v34 << ((1) & 63)) >= p1) {
                                for fi36 := 0; fi36 < 17; fi36++ {
                                    var v37 int64 = -700
                                    var v38 float64 = 3.14159
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                }
                                p1 += (p1 << ((uint(p2)) & 63))
                            } else {
                                var v39 int = p1
                                _ = v39
                            }
                            fmt.Println(v34, p0, p1)
                            switch "foo bar" {
                            case "a":
                                v40 := []int{-2147483648, 32768}
                                _ = v40
                            default:
                                p0 = false
                            }
                        }
                    }
                case "a":
                    v41 := GenS2{F0: p0, F1: 1}
                    _ = v41
                }
                for fi42 := 0; fi42 < 17; fi42++ {
                    var v43 uint64 = 42
                    v43 -= (v43 & (v43 & (^(v43))))
                    _ = fi42
                    _ = v43
                }
            }
            _ = v34
        }
    } else {
        p0 = p0
    }
    p2 = p2
    if p0 {
        v44 := GenS4{F0: 818, F1: (-125 > 32)}
        v45 := GenS5{F0: (p2 >> ((255) & 63)), F1: (42 == 2147483648)}
        p0 = (255 <= 107)
        _ = v44
        _ = v45
    }
    v46 := []int{}
    var v47 bool = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v46
    _ = v47
    return p2
}

func main() {
    var v1 float64 = 1e10
    var v2 uint64 = 42
    v3 := make(map[string]uint32)
    var v4 int16 = 1000
    var v5 uint64 = 42
    v6 := make(map[uint32]string)
    var v48 uint16 = 65535
    v1 += float64(3.14159)
    var v49 int = 127
    var v50 int8 = 57
    var v51 uint = 4294967295
    var v52 uint8 = 127
    switch (("hello" + "foo bar") + "\n\t") {
    case "a":
        switch "" {
        case "hello":
            v53 := []int{(-(v49))}
            _ = v53
        case "\n\t":
            var v54 string = "hello"
            _ = v54
        case "":
            v50 = v50
        }
    case "":
        var v55 float64 = v1
        _ = v55
    case "\n\t":
        var v56 int32 = 2147483647
        _ = v56
    }
    fmt.Println(v50)
    var v57 int = v49
    var v58 float64 = v1
    v48 = v48
    v50 = v50
    var v59 bool = true
    if (v48 >= (v48 & (v48 ^ 42))) {
        if true {
            v60 := []int{}
            v59 = v59
            var v61 uint8 = 255
            _ = v60
            _ = v61
        } else {
            var v62 int = (v49 >> (((v51 - v51)) & 63))
            var v63 string = "foo bar"
            _ = v62
            _ = v63
        }
    }
    for fi64 := 0; fi64 < 16; fi64++ {
        if v59 {
            for fi65 := 0; fi65 < 5; fi65++ {
                v66 := GenS6{F0: 9223372036854775807, F1: 127, F2: 2}
                v67 := v3["hello"]
                var v68 bool = (!(("" <= "hello") || v59))
                _ = fi65
                _ = v66
                _ = v67
                _ = v68
            }
            var v69 int64 = 9223372036854775807
            v70 := []int{(-(v49)), v57, v57}
            _ = v69
            _ = v70
        }
        _ = fi64
    }
    v71 := []int{int(v52)}
    var v72 uint16 = uint16(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v71:", len(v71), v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v71
    _ = v72
}
