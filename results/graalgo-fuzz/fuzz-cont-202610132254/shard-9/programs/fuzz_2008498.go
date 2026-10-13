package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint64; F2 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 uint16; F1 string; F2 int8 }
func genFn0(p0 bool, p1 string, p2 uint8) bool {
    v6 := GenS1{F0: 1, F1: 2, F2: 114}
    var v7 uint16 = 1
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return false
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[int]int)
    var v3 int64 = 9223372036854775807
    var v4 bool = false
    var v5 int16 = -911
    v1 = v1
    switch -598 {
    case 42:
        var v8 uint8 = 183
        _ = v8
    case 127:
        if (3.14159 != 3.14159) {
            if v4 {
                v4 = ((v3 | (^(v3))) < -9223372036854775808)
            } else {
                switch "" {
                case "a":
                    var v9 int8 = 42
                    _ = v9
                }
                var v10 int64 = v3
                _ = v10
            }
            var v11 int8 = 1
            _ = v11
        } else {
            v12 := GenS2{F0: v5}
            _ = v12
        }
    case -11:
        v13 := GenS3{F0: v3}
        _ = v13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 float32 = 0.1
        v15 = 0.1
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        if v4 {
            if v4 {
                delete(v2, (^(fi16)))
                if (-1 >= (v5 & 42)) {
                    for fi17 := 0; fi17 < 19; fi17++ {
                        if (0 < 18446744073709551615) {
                            _ = len(v2)
                            if v4 {
                                if true {
                                    v18 := GenS1{F0: 1, F1: v1, F2: 1}
                                    if ((int32(fi17) != 119) || v4) {
                                        switch fi16 {
                                        case -9223372036854775808:
                                            v19 := []int{(^(fi17)), (-(fi16)), -381, -1}
                                            _ = v19
                                        case 4294967295:
                                            if v4 {
                                                v4 = v4
                                            }
                                        case 407:
                                            if true {
                                                fmt.Println(v4)
                                                v20 := GenS4{F0: 42}
                                                _ = v20
                                            } else {
                                                v3 = v3
                                                var v21 int64 = v3
                                                _ = v21
                                            }
                                        }
                                    }
                                    switch (fi17 << ((65535) & 63)) {
                                    case 1:
                                        switch "a" {
                                        case "hello":
                                            delete(v2, fi17)
                                        }
                                    case 9223372036854775807:
                                        v3 += int64(v1)
                                    }
                                    _ = v18
                                } else {
                                    fmt.Println(fi16)
                                    v22 := GenS5{F0: "", F1: 4294967295, F2: v1}
                                    v4 = v4
                                    _ = v22
                                }
                                switch "a" {
                                case "a":
                                    if (v4 || (-120 > 97)) {
                                        fmt.Println(v1)
                                    }
                                }
                                v5 = v5
                            }
                        } else {
                            v23 := []int{(^(fi17))}
                            _ = v23
                        }
                        _ = fi17
                    }
                } else {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        switch "a" {
                        case "hello":
                            v1 = v1
                        case "":
                            switch ("foo bar" + "hello") {
                            case "":
                                v2[fi24] = (fi24 | 884)
                            case "foo bar":
                                v25 := GenS5{F0: (("" + "hello") + ((("foo bar" + "hello") + "hello") + "")), F1: 18446744073709551615, F2: (v1 >> ((904) & 63))}
                                _ = v25
                            case "a":
                                v26 := GenS6{F0: ""}
                                _ = v26
                            }
                        }
                        var v27 uint = 18446744073709551615
                        if (v4 && (!(!v4))) {
                            if v4 {
                                v3 -= (v3 + -484)
                                var v28 int64 = v3
                                var v29 uint8 = 0
                                _ = v28
                                _ = v29
                            }
                            fmt.Println(v5)
                        }
                        _ = fi24
                        _ = v27
                    }
                    v3 += (-(v3))
                }
                var v30 uint8 = 42
                _ = v30
            } else {
                v5 += -1
            }
        } else {
            fmt.Println(fi16, fi16, v4)
        }
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 uint16 = v1
            _ = fi31
            _ = v32
        }
        _ = fi16
    }
    v33 := GenS7{F0: 255}
    var v34 uint = 2
    v1 = v1
    switch "" {
    case "a":
        v35 := []int{9223372036854775807, 128, 250}
        _ = v35
    default:
        switch "hello" {
        case "":
            v34 = v34
        case "\n\t":
            v36 := GenS7{F0: 18446744073709551615}
            _ = v36
        case "a":
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 bool = (((!("\n\t" <= "a")) && ((v1 << (((^(v34))) & 63)) >= v1)) && v4)
                var v39 float64 = 0.0
                delete(v2, (fi37 % (((fi37 - fi37)) | 1)))
                _ = fi37
                _ = v38
                _ = v39
            }
        }
    }
    var v40 int16 = v5
    if (!v4) {
        v41 := v2[2147483647]
        for fi42 := 0; fi42 < 16; fi42++ {
            var v43 float32 = 1.5
            v44 := GenS7{F0: v34}
            v3 = 0
            _ = fi42
            _ = v43
            _ = v44
        }
        for fi45 := 0; fi45 < 9; fi45++ {
            for fi46 := 0; fi46 < 17; fi46++ {
                v47 := GenS8{F0: 65535, F1: v41}
                _ = fi46
                _ = v47
            }
            _ = fi45
        }
        _ = v41
    } else {
        if ((v40 * (v40 ^ (^(v5)))) == -32768) {
            for fi48 := 0; fi48 < 13; fi48++ {
                switch ("hello" + ("" + "\n\t")) {
                case "":
                    v49 := []int{fi48, (fi48 & (fi48 << ((42) & 63)))}
                    _ = v49
                case "\n\t":
                    v50 := []int{fi48, (fi48 ^ -1)}
                    _ = v50
                }
                v51 := []int{(^(fi48)), (fi48 + 4294967295), fi48, (^(fi48))}
                var v52 int = fi48
                _ = fi48
                _ = v51
                _ = v52
            }
            v53 := GenS7{F0: v34}
            _ = v53
        }
    }
    v4 = v4
    var v54 int16 = v40
    v55 := GenS9{F0: v1, F1: "a", F2: 100}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v40
    _ = v5
    _ = v54
    _ = v55
}
