package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 int8; F2 int16 }
type GenS5 struct { F0 uint16; F1 uint8; F2 int }
type GenS6 struct { F0 uint; F1 int64; F2 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint; F1 int16 }
func genFn0(p0 int32) int64 {
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := []int{int(p0), (fi7 >> ((926) & 63))}
            if len(v8) > 0 {
                v8[0] = (^(fi7))
            }
            v8 = append(v8, -994)
            switch "hello" {
            case "foo bar":
                p0 = (p0 / ((0) | 1))
            default:
                var v9 bool = true
                _ = v9
            }
            _ = fi7
            _ = v8
        }
        v10 := GenS1{F0: 65535, F1: 65535, F2: 447}
        if (true || true) {
            v11 := GenS2{F0: 42}
            var v12 int16 = 42
            p0 = (p0 >> ((4294967295) & 63))
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 5; fi13++ {
                switch ("a" + ("a" + ("" + "hello"))) {
                case "":
                    v14 := GenS2{F0: -2147483648}
                    _ = v14
                case "hello":
                    v15 := []int{fi6, 128}
                    _ = v15
                case "a":
                    if true {
                        v16 := []int{fi13}
                        fmt.Println(fi6, p0)
                        _ = v16
                    }
                }
                p0 -= (p0 & p0)
                if true {
                    v17 := []int{fi6}
                    var v18 int16 = -1
                    if (true && false) {
                        if true {
                            if false {
                                v18 = v18
                                fmt.Println(p0)
                            }
                            v19 := GenS3{F0: -9223372036854775808}
                            _ = v19
                        } else {
                            if len(v17) > 0 {
                                v17[0] = 2147483647
                            }
                            if false {
                                v20 := GenS4{F0: 128, F1: -1, F2: v18}
                                _ = v20
                            } else {
                                if false {
                                    if len(v17) > 0 {
                                        v17[0] = -9223372036854775808
                                    }
                                    if (!true) {
                                        if ((true && ((fi6 & fi13) != (-(fi6)))) || true) {
                                            fmt.Println(fi13)
                                            switch "a" {
                                            case "a":
                                                switch (("a" + ("hello" + "a")) + (((("\n\t" + "\n\t") + "\n\t") + ("a" + ("a" + "foo bar"))) + ((("a" + "\n\t") + "foo bar") + (("a" + "\n\t") + "")))) {
                                                case "a":
                                                    if true {
                                                        var v21 uint64 = 1
                                                        v18 += v18
                                                        _ = v21
                                                    }
                                                    var v22 int = (fi6 >> ((531) & 63))
                                                    var v23 uint32 = 4294967295
                                                    v23 -= (v23 - 42)
                                                    fmt.Println(fi6, v23)
                                                    if len(v17) > 0 {
                                                        v17[0] = 0
                                                    }
                                                    _ = v22
                                                    _ = v23
                                                case "hello":
                                                    v24 := GenS5{F0: 1, F1: 0, F2: 127}
                                                    _ = v24
                                                }
                                            case "hello":
                                                switch (^(fi13)) {
                                                case 221:
                                                    fmt.Println(fi13)
                                                }
                                            }
                                            _ = len(v17)
                                        }
                                        var v25 uint64 = 1
                                        _ = v25
                                    } else {
                                        var v26 int8 = 100
                                        v27 := GenS6{F0: 44, F1: 9223372036854775807, F2: true}
                                        _ = v26
                                        _ = v27
                                    }
                                    fmt.Println(fi6)
                                } else {
                                    if len(v17) > 0 {
                                        v17[0] = fi6
                                    }
                                }
                                fmt.Println(fi6, p0)
                            }
                            var v28 string = "foo bar"
                            _ = v28
                        }
                        fmt.Println(p0)
                        v18 -= (^(v18))
                    }
                    _ = v17
                    _ = v18
                } else {
                    switch "hello" {
                    case "\n\t":
                        var v29 int32 = (p0 / (((p0 * (^(p0)))) | 1))
                        _ = v29
                    }
                }
                _ = fi13
            }
            var v30 bool = false
            v31 := []int{(fi6 / ((fi6) | 1)), (fi6 % ((32767) | 1)), -356}
            _ = v30
            _ = v31
        }
        _ = fi6
        _ = v10
    }
    var v32 int32 = p0
    v33 := GenS1{F0: 875, F1: 607, F2: 0}
    _ = p0
    _ = v32
    _ = v33
    return 42
}
func genFn1(p0 uint32, p1 int64, p2 string) int32 {
    var v34 uint16 = 981
    if (true || false) {
        v35 := GenS2{F0: -2147483648}
        v34 += v34
        _ = v35
    }
    var v36 uint32 = 732
    fmt.Println(v36)
    switch p2 {
    case "":
        var v37 bool = (((v34 / ((v34) | 1)) < (v34 << ((255) & 63))) && (125 != 255))
        _ = v37
    case "\n\t":
        v38 := GenS7{F0: -1}
        _ = v38
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    _ = v36
    return -1
}

func main() {
    var v1 int32 = -1
    v2 := make(map[string]bool)
    var v3 int16 = 1
    v4 := make(map[string]uint32)
    var v5 int32 = (v1 & -2147483648)
    v39 := v4["a"]
    if true {
        if (555 != 784) {
            v1 = genFn1(4294967295, -39, "\n\t")
        } else {
            var v40 float64 = 3.14159
            v41 := []int{}
            var v42 int = 417
            _ = v40
            _ = v41
            _ = v42
        }
        switch -463 {
        case 4294967295:
            var v43 bool = true
            _ = v43
        default:
            fmt.Println(v3, v1, v5)
        }
    }
    var v44 int32 = v5
    var v45 bool = (!(251 == 2))
    v44 = 871
    v5 = genFn1(v39, -1, ("hello" + ("foo bar" + "\n\t")))
    v46 := GenS7{F0: v3}
    v3 -= (v3 | (v3 * v3))
    var v47 bool = (v45 && (2 == 32767))
    v48 := GenS8{F0: 42, F1: (v3 >> ((959) & 63))}
    fmt.Println(v3, v39, v47)
    if (("hello" + (((("a" + "") + "\n\t") + "foo bar") + "hello")) < "foo bar") {
        v2[("\n\t" + "\n\t")] = (!false)
    }
    fmt.Println(v45, v47, v47)
    v49 := GenS1{F0: 1, F1: 1, F2: -1}
    var v50 string = ("a" + (("foo bar" + "hello") + "\n\t"))
    var v51 int64 = -111
    if (185 >= 1) {
        if (-32768 > (v3 ^ (-(v3)))) {
            var v52 bool = v47
            var v53 uint8 = 112
            _ = v52
            _ = v53
        } else {
            v54 := GenS2{F0: (v44 ^ (v44 * v5))}
            v1 = genFn1((v39 * (v39 / (((v39 * v39)) | 1))), v51, ((v50 + v50) + (v50 + v50)))
            var v55 uint8 = 2
            _ = v54
            _ = v55
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
