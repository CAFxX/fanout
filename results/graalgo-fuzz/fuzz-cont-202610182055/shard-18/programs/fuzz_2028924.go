package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int64; F1 string; F2 int32 }
type GenS4 struct { F0 uint; F1 int32; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 int32 = 0
    var v2 int64 = -1
    v3 := []int{}
    var v4 int = 42
    v5 := make(map[string]bool)
    v6 := []uint32{42, 2147483648, 2147483648, 402}
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 77, F1: 100, F2: fi7}
        if (42 < 255) {
            var v9 uint32 = 0
            _ = v9
        } else {
            fmt.Println(v2)
        }
        _ = fi7
        _ = v8
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := GenS2{F0: (v2 + v2), F1: 918}
            if (42 > 127) {
                v2 += int64(fi10)
                if (2 <= 18446744073709551615) {
                    var v13 uint32 = 1
                    var v14 int = fi11
                    v15 := GenS3{F0: (^(v2)), F1: "\n\t", F2: (-(v1))}
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    fmt.Println(v4)
                }
                fmt.Println(v4, fi11)
            } else {
                v16 := v5[""]
                fmt.Println(v1, fi11, fi11)
                _ = v16
            }
            switch "" {
            case "a":
                if true {
                    v5["a"] = false
                    var v17 uint64 = 2
                    v4 = v4
                    if (false && ((("" + "a") + ("\n\t" + "")) <= "foo bar")) {
                        switch ("" + ("foo bar" + "")) {
                        case "\n\t":
                            switch "\n\t" {
                            case "hello":
                                delete(v5, "\n\t")
                            case "\n\t":
                                if true {
                                    v17 += (v17 << ((579) & 63))
                                } else {
                                    var v18 int64 = (^(v2))
                                    v19 := v5["a"]
                                    if v19 {
                                        switch ("foo bar" + (("foo bar" + "foo bar") + "foo bar")) {
                                        case "a":
                                            var v20 int8 = 127
                                            _ = v20
                                        case "foo bar":
                                            v1 = (v1 / (((v1 ^ (v1 + (v1 ^ (v1 * v1))))) | 1))
                                        }
                                        var v21 string = "a"
                                        _ = v21
                                    } else {
                                        fmt.Println(fi11, v17)
                                        v19 = v19
                                        v22 := v5[("a" + (((("foo bar" + "hello") + "\n\t") + "\n\t") + (((("hello" + ((("a" + "") + "hello") + (("foo bar" + "") + "foo bar"))) + "hello") + "hello") + "\n\t")))]
                                        v23 := v5["hello"]
                                        fmt.Println(v19, v19)
                                        _ = v22
                                        _ = v23
                                    }
                                    _ = v18
                                    _ = v19
                                }
                                switch (fi11 | (fi10 / (((-(fi11))) | 1))) {
                                case -1:
                                    var v24 int = fi10
                                    _ = v24
                                case -284:
                                    v4 = 32768
                                default:
                                    fmt.Println(v1)
                                }
                                v6 = append(v6, uint32(v17))
                            case "":
                                v4 += v4
                            }
                        case "":
                            fmt.Println(v1, v17, v1)
                        case "foo bar":
                            _ = len(v5)
                        }
                        fmt.Println(fi11)
                        if ((^(v2)) != (v2 + (v2 & v2))) {
                            var v25 int32 = v1
                            var v26 bool = ((float64(float32((1.5 / 1.0))) - 1e10) > -2.25)
                            v25 = v25
                            _ = v25
                            _ = v26
                        }
                        v1 -= v1
                        v27 := GenS4{F0: 18446744073709551615, F1: 665, F2: 1}
                        _ = v27
                    }
                    _ = v17
                }
            case "":
                v28 := GenS5{F0: 255}
                _ = v28
            }
            _ = fi11
            _ = v12
        }
        var v29 float64 = 1.5
        _ = fi10
        _ = v29
    }
    v6 = append(v6, 324)
    var v30 uint64 = 42
    for fi31 := 0; fi31 < 10; fi31++ {
        v32 := GenS2{F0: v2, F1: 2}
        _ = fi31
        _ = v32
    }
    fmt.Println(v1, v4)
    if true {
        if true {
            v3 = append(v3, (^(v4)))
            switch "a" {
            case "\n\t":
                if false {
                    if true {
                        switch (v4 + (v4 >> ((2) & 63))) {
                        case 0:
                            fmt.Println(v1)
                        }
                        v33 := GenS3{F0: v2, F1: ("hello" + ((("\n\t" + "") + "hello") + "\n\t")), F2: v1}
                        v34 := GenS5{F0: 32767}
                        _ = v33
                        _ = v34
                    }
                    var v35 float64 = -2.25
                    _ = v35
                } else {
                    if (3.14159 == float64(1e10)) {
                        if (355 == 811) {
                            v1 += v1
                            v4 = v4
                            v2 = (v2 >> ((42) & 63))
                        } else {
                            v30 = (v30 | (^(v30)))
                        }
                    } else {
                        v36 := GenS1{F0: 0, F1: 41, F2: (v4 ^ v4)}
                        if (42 != 796) {
                            var v37 int32 = (v1 >> ((299) & 63))
                            _ = v37
                        } else {
                            if (128 < 1) {
                                v38 := GenS1{F0: 105, F1: -28, F2: 32768}
                                v1 = (v1 % (((-(v1))) | 1))
                                _ = v38
                            }
                            if true {
                                v1 = (-(v1))
                                fmt.Println(v4)
                            } else {
                                if (574 <= 130) {
                                    var v39 int = (v4 ^ v4)
                                    v5[""] = true
                                    v40 := GenS6{F0: (^(v30))}
                                    _ = v39
                                    _ = v40
                                }
                            }
                            var v41 uint = 478
                            _ = v41
                        }
                        if ((!(42 == 42)) && true) {
                            var v42 string = "a"
                            _ = v42
                        }
                        _ = v36
                    }
                    v2 = v2
                }
            }
            var v43 float32 = 1.5
            _ = v43
        }
        var v44 string = "foo bar"
        if len(v6) > 0 {
            v6[0] = 0
        }
        _ = v44
    }
    var v45 float64 = 3.14159
    v3 = append(v3, (v4 | v4))
    v46 := v5[""]
    for fi47 := 0; fi47 < 10; fi47++ {
        for fi48 := 0; fi48 < 2; fi48++ {
            v49 := GenS2{F0: 1, F1: 65535}
            var v50 int64 = v2
            _ = fi48
            _ = v49
            _ = v50
        }
        _ = fi47
    }
    var v51 uint32 = 2
    v52 := v5["foo bar"]
    v53 := v5[("" + ("a" + "foo bar"))]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v6
}
