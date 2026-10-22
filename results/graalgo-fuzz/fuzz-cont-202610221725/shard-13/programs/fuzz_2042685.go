package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int16; F1 string; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int; F2 uint32 }
type GenS6 struct { F0 uint64; F1 int64; F2 uint64 }
func genFn0() bool {
    var v6 bool = (!(((1 > 2) || (!(!(42 >= 287)))) || ((true && (!((13 <= -47) && true))) || false)))
    if false {
        for fi7 := 0; fi7 < 6; fi7++ {
            if (32767 >= 1) {
                var v8 int16 = -1
                _ = v8
            } else {
                for fi9 := 0; fi9 < 1; fi9++ {
                    v10 := GenS1{F0: fi9, F1: (((1e10 / 3.14159) != 0.0) && true)}
                    v11 := []int{fi7, 4294967295, (-(fi7))}
                    switch "hello" {
                    case "foo bar":
                        _ = len(v11)
                    case "a":
                        fmt.Println(v6, fi7)
                    case "hello":
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    default:
                        v12 := GenS2{F0: -1, F1: "hello", F2: 962}
                        _ = v12
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            v13 := []int{(fi7 & (fi7 * -2147483648)), (fi7 >> ((255) & 63)), (fi7 | fi7)}
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS1{F0: 42, F1: v6}
                var v16 int8 = 127
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi7
            _ = v13
        }
    } else {
        v17 := GenS1{F0: 4294967295, F1: (("hello" + ("a" + "")) > ("foo bar" + ("foo bar" + "hello")))}
        var v18 int = 1
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := []int{fi19, 1, fi19, 32768}
        _ = fi19
        _ = v20
    }
    var v21 bool = ((v6 || (-927 != 1)) && ((!((((!false) || ((!(!(-933 != 0))) || v6)) || ((!(v6 && (0 == -472))) && v6)) || v6)) || v6))
    _ = v21
    _ = v6
    return v6
}
func genFn1(p0 int32) int64 {
    if false {
        p0 -= (p0 << ((982) & 63))
        if (!(false && ((p0 + p0) < 42))) {
            if false {
                for fi22 := 0; fi22 < 11; fi22++ {
                    fmt.Println(p0, fi22)
                    _ = fi22
                }
                p0 = 0
                v23 := []int{127, 224, -600, 781}
                _ = v23
            } else {
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := []int{fi24, (^(fi24)), 9223372036854775807, 722}
                    v26 := GenS3{F0: 0, F1: 2}
                    var v27 float64 = float64(0.1)
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                var v28 int32 = 0
                _ = v28
            }
            switch 4294967295 {
            case 68:
                v29 := []int{127, -746}
                _ = v29
            }
            p0 -= p0
        } else {
            var v30 uint16 = 255
            for fi31 := 0; fi31 < 20; fi31++ {
                if false {
                    fmt.Println(p0, v30, fi31)
                } else {
                    p0 = p0
                    var v32 uint16 = (v30 >> ((115) & 63))
                    _ = v32
                }
                var v33 uint32 = 4294967295
                _ = fi31
                _ = v33
            }
            _ = v30
        }
    } else {
        fmt.Println(p0)
        p0 = -213
    }
    v34 := []int{}
    _ = p0
    _ = v34
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 7
    var v2 uint8 = (^(v1))
    var v3 int = -9223372036854775808
    var v4 int64 = -1
    var v5 uint16 = 42
    var v35 int16 = 870
    var v36 int = 822
    v37 := GenS1{F0: 97, F1: false}
    var v38 int = (v3 / (((v36 + -599)) | 1))
    var v39 uint16 = (v5 & (v5 & 0))
    if (true || (321 >= (v39 << ((4294967295) & 63)))) {
        for fi40 := 0; fi40 < 19; fi40++ {
            var v41 float64 = 0.1
            v42 := GenS1{F0: v38, F1: false}
            for fi43 := 0; fi43 < 2; fi43++ {
                v44 := GenS4{F0: v4}
                var v45 int8 = 127
                v46 := GenS5{F0: "\n\t", F1: -1, F2: 42}
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
            }
            _ = fi40
            _ = v41
            _ = v42
        }
        fmt.Println(v2, v38)
        v47 := GenS6{F0: 18446744073709551615, F1: v4, F2: 0}
        _ = v47
    } else {
        v48 := []int{(v36 + v38), (v38 ^ (v38 * (v38 - -2147483648))), v38}
        _ = v48
    }
    switch "foo bar" {
    case "\n\t":
        var v49 uint = 1
        _ = v49
    }
    v50 := GenS2{F0: (v35 >> ((2) & 63)), F1: "foo bar", F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v50
}
