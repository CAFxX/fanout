package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint64, p1 uint16, p2 uint32) int16 {
    switch 2147483647 {
    case 32767:
        p2 = (p2 << ((65535) & 63))
    case 128:
        v6 := GenS1{F0: -1, F1: 4294967295, F2: ("foo bar" + "\n\t")}
        _ = v6
    default:
        var v7 int32 = -172
        _ = v7
    }
    p2 -= p2
    var v8 float32 = (0.1 / 1e10)
    var v9 uint32 = (^(p2))
    fmt.Println(v8, v9, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint64 = 2
    v3 := make(map[string]uint32)
    var v4 float32 = 1e10
    var v5 bool = false
    v10 := GenS1{F0: -818, F1: 72, F2: "hello"}
    v5 = v5
    var v11 bool = false
    if (77 == 42) {
        switch 32768 {
        case 32767:
            v12 := []int{}
            _ = v12
        case 9223372036854775807:
            var v13 uint32 = v1
            _ = v13
        case 32768:
            v14 := GenS2{F0: 255}
            _ = v14
        default:
            v2 += 42
        }
    } else {
        var v15 uint32 = (v1 | v1)
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := v3["hello"]
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    switch ("a" + "") {
    case "":
        var v18 int64 = -1
        _ = v18
    case "\n\t":
        v19 := []int{4294967295, 32767, 2147483647}
        _ = v19
    }
    if v5 {
        v1 = v1
        if (v1 <= v1) {
            v1 = v1
            var v20 int8 = -48
            _ = v20
        }
        var v21 int64 = -435
        _ = v21
    } else {
        v22 := []int{634, 32767}
        _ = v22
    }
    v11 = v11
    delete(v3, "foo bar")
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := GenS1{F0: -325, F1: 255, F2: "a"}
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: 1, F1: 18446744073709551615, F2: ("foo bar" + (("foo bar" + "hello") + ((("hello" + "foo bar") + "foo bar") + ("" + ""))))}
    v1 = (v1 + 573)
    delete(v3, "foo bar")
    if (-37 <= 107) {
        switch -2147483648 {
        case 32767:
            delete(v3, "")
        }
        var v26 int = -1
        for fi27 := 0; fi27 < 16; fi27++ {
            for fi28 := 0; fi28 < 2; fi28++ {
                if ((v1 / ((v1) | 1)) <= 537) {
                    if v11 {
                        v11 = (v11 || (false && (((v1 >> ((18446744073709551615) & 63)) <= 2147483647) && ((true || (-217 > 0)) && (!v11)))))
                        v29 := GenS3{F0: 65535}
                        _ = v29
                    } else {
                        switch "a" {
                        case "foo bar":
                            v30 := GenS2{F0: 2}
                            _ = v30
                        case "":
                            v4 -= (float32(0.1) / v4)
                        case "hello":
                            var v31 int8 = -1
                            _ = v31
                        default:
                            if (-2147483648 != 2147483647) {
                                v32 := []int{}
                                v33 := GenS4{F0: 0}
                                _ = v32
                                _ = v33
                            } else {
                                if v11 {
                                    v34 := GenS2{F0: 1}
                                    v35 := []int{v26, (fi28 / (((v26 | fi27)) | 1))}
                                    _ = v34
                                    _ = v35
                                } else {
                                    v11 = ((("\n\t" + (("a" + "foo bar") + "\n\t")) + (("foo bar" + ("a" + "")) + "\n\t")) > "a")
                                    var v36 float64 = 1.5
                                    _ = v36
                                }
                            }
                            v37 := []int{fi27, (v26 | fi28)}
                            _ = v37
                        }
                        if (((1 != 42) && ((3.14159 - 1.5) < (v4 / (v4 * v4)))) || (212 != 189)) {
                            var v38 uint16 = 65535
                            v39 := GenS1{F0: 42, F1: 310, F2: ((("hello" + "hello") + "") + "\n\t")}
                            v5 = (!true)
                            _ = v38
                            _ = v39
                        }
                        v2 = (v2 - 18446744073709551615)
                        v40 := GenS4{F0: v1}
                        _ = v40
                    }
                    var v41 bool = v11
                    v42 := GenS4{F0: 2147483648}
                    v4 += 1e10
                    v1 += v1
                    switch v26 {
                    case -1:
                        var v43 int64 = 9223372036854775807
                        _ = v43
                    case 32767:
                        v26 -= (v26 % ((v26) | 1))
                    }
                    _ = v41
                    _ = v42
                }
                v44 := GenS5{F0: 0, F1: -1, F2: (v2 & (v2 | v2))}
                if ((v1 + v1) < 4294967295) {
                    var v45 float32 = (((0.1 * float32(1.5)) / 1.0) * -2.25)
                    _ = v45
                } else {
                    if v5 {
                        v1 = 42
                        _ = len(v3)
                        var v46 uint64 = (v2 * v2)
                        _ = v46
                    } else {
                        v47 := []int{fi28, fi28, 1, (^(fi28))}
                        var v48 int16 = genFn0((v2 % ((102) | 1)), 255, (v1 & (v1 % ((v1) | 1))))
                        if true {
                            v11 = (!(290 != 0))
                            fmt.Println(v4, v1)
                            delete(v3, "")
                        } else {
                            if (v5 || (!((v2 * 1) <= (v2 << ((uint(v2)) & 63))))) {
                                switch int(v2) {
                                case 42:
                                    v2 -= (v2 % ((v2) | 1))
                                case 0:
                                    var v49 uint32 = 335
                                    _ = v49
                                default:
                                    v1 -= 65535
                                }
                                v4 = float32(1e10)
                            } else {
                                var v50 int8 = 31
                                v51 := GenS6{F0: v1}
                                v52 := GenS4{F0: v1}
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                        }
                        _ = v47
                        _ = v48
                    }
                    fmt.Println(v1, v1, v2)
                }
                switch (v26 << ((2) & 63)) {
                case 9223372036854775807:
                    if (((v4 / v4) * v4) <= float32((((0.0 * ((0.0 * float64((float32(1.5) / float32(0.1)))) + 0.1)) - float64(v4)) - (0.0 / 1.0)))) {
                        fmt.Println(v5, v5)
                        var v53 int32 = -1
                        v4 = float32(float64(v4))
                        _ = v53
                    }
                case 32768:
                    v11 = false
                case 0:
                    fmt.Println(v11, v5, v2)
                default:
                    var v54 uint8 = 128
                    _ = v54
                }
                v55 := GenS2{F0: 255}
                _ = fi28
                _ = v44
                _ = v55
            }
            if (0 == v1) {
                v56 := v3[""]
                var v57 bool = (!(float64(0.0) <= -2.25))
                _ = v56
                _ = v57
            }
            v3[("\n\t" + "\n\t")] = (^(v1))
            _ = fi27
        }
        _ = v26
    } else {
        fmt.Println(v4)
        var v58 uint8 = 231
        switch 1 {
        case -858:
            v5 = (437 < 0)
        case 32768:
            v11 = (v11 && (255 >= 154))
        }
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
