package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int; F1 uint; F2 uint32 }
type GenS6 struct { F0 int64; F1 string; F2 int64 }
func genFn0(p0 string, p1 uint) string {
    if (32767 == 523) {
        v7 := GenS1{F0: 2}
        p1 = (p1 | (^(p1)))
        p1 = (p1 % ((p1) | 1))
        _ = v7
    } else {
        if (!(!(127 < 1))) {
            var v8 int8 = int8(p1)
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 uint = (p1 + p1)
                v11 := []int{fi9, (^(fi9)), fi9}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v8
        }
        switch 32767 {
        case 218:
            v12 := GenS1{F0: 42}
            _ = v12
        case 653:
            v13 := GenS2{F0: -32768, F1: 9223372036854775807}
            _ = v13
        case 53:
            var v14 int64 = -130
            _ = v14
        default:
            p0 = ""
        }
        var v15 uint16 = 42
        _ = v15
    }
    v16 := GenS3{F0: 416, F1: 4294967295, F2: 4294967295}
    _ = p0
    _ = p1
    _ = v16
    return "hello"
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 18446744073709551615
    var v3 string = "hello"
    var v4 int32 = 42
    v5 := make(map[string]string)
    v6 := make(map[int]int)
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := []int{(fi17 * (fi17 >> (((v2 - 4294967295)) & 63)))}
        _ = fi17
        _ = v18
    }
    v5["foo bar"] = ((v3 + (("foo bar" + v3) + (v3 + v3))) + ((v3 + v3) + "hello"))
    v3 = v3
    v3 = (v3 + v3)
    v19 := GenS1{F0: 42}
    v3 = "a"
    v20 := GenS3{F0: 42, F1: 4294967295, F2: 2}
    var v21 uint = (v2 + (v2 + (v2 ^ (v2 | (v2 << (((v2 * (v2 % ((884) | 1)))) & 63))))))
    if true {
        if ((9223372036854775807 < 150) && (42 > 0)) {
            v4 += v4
            switch 97 {
            case 32768:
                var v22 uint64 = 2
                _ = v22
            case 559:
                fmt.Println(v21, v21)
            case -9223372036854775808:
                v4 -= (v4 >> ((4294967295) & 63))
            }
            if false {
                if true {
                    delete(v6, 32767)
                }
                v2 = v2
                switch (v3 + v3) {
                case "\n\t":
                    v23 := GenS4{F0: 18446744073709551615, F1: 362}
                    _ = v23
                default:
                    v1[""] = false
                }
            }
        } else {
            if ((2 != 42) || ((!(495 >= 0)) && ("" >= v3))) {
                v24 := GenS5{F0: -9223372036854775808, F1: 227, F2: 42}
                _ = v24
            }
            if true {
                for fi25 := 0; fi25 < 4; fi25++ {
                    v26 := GenS1{F0: 2}
                    for fi27 := 0; fi27 < 18; fi27++ {
                        v4 = (v4 / ((v4) | 1))
                        v28 := []int{fi27, 4294967295, fi27, fi25}
                        _ = fi27
                        _ = v28
                    }
                    if (4294967295 >= v2) {
                        v2 = 541
                        if false {
                            var v29 int64 = -1
                            v3 = v3
                            var v30 bool = (42 <= 0)
                            _ = v29
                            _ = v30
                        } else {
                            if false {
                                fmt.Println(v2, fi25)
                                v21 += 524
                            } else {
                                for fi31 := 0; fi31 < 12; fi31++ {
                                    var v32 int32 = v4
                                    _ = fi31
                                    _ = v32
                                }
                            }
                            if (v3 != "foo bar") {
                                for fi33 := 0; fi33 < 5; fi33++ {
                                    if true {
                                        v3 = v3
                                        if (false && ((0.1 / 0.1) >= 1.5)) {
                                            fmt.Println(v3, fi33, fi33)
                                            fmt.Println(fi33, v4)
                                            v5[v3] = (("" + v3) + (("" + v3) + ("" + ("\n\t" + v3))))
                                            fmt.Println(v3)
                                        }
                                    }
                                    v34 := GenS6{F0: 0, F1: v3, F2: -539}
                                    _ = fi33
                                    _ = v34
                                }
                            }
                        }
                    }
                    _ = fi25
                    _ = v26
                }
                v35 := []int{906}
                _ = v35
            }
        }
        fmt.Println(v4, v2, v2)
        var v36 float32 = 0.0
        _ = v36
    }
    v2 += (v21 / (((v2 * (v2 >> (((v21 * (v2 | 255))) & 63)))) | 1))
    var v37 uint32 = 0
    if (0 == 89) {
        v21 = (v21 << ((65535) & 63))
        var v38 int8 = 0
        v39 := []int{551}
        _ = v38
        _ = v39
    }
    var v40 int64 = -1
    var v41 float64 = 0.0
    v41 -= v41
    switch -1 {
    case 127:
        v1[v3] = (!true)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
