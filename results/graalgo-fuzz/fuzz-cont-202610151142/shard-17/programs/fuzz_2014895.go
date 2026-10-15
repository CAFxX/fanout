package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 42
    var v3 uint32 = 4294967295
    var v4 uint16 = 993
    var v5 uint64 = 1
    var v6 float64 = 1e10
    var v7 bool = true
    var v8 float32 = (-2.25 / -2.25)
    fmt.Println(v6, v5)
    v5 += (v5 + v5)
    var v9 uint64 = uint64(v3)
    for fi10 := 0; fi10 < 15; fi10++ {
        v4 = 1
        var v11 bool = v7
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 bool = ("foo bar" > "a")
            if (!(19 <= -1)) {
                v14 := []int{}
                v15 := GenS1{F0: "foo bar", F1: ("hello" + "foo bar")}
                if (v11 || (((!(44 != uint(v2))) || v11) || v11)) {
                    fmt.Println(v3, v6)
                    var v16 int8 = 1
                    v7 = v7
                    _ = v16
                }
                _ = v14
                _ = v15
            } else {
                switch 689 {
                case 32767:
                    var v17 uint64 = v9
                    _ = v17
                }
                if (v8 == ((0.1 - (-2.25 / 1.5)) * ((-2.25 + 0.0) / 1.0))) {
                    var v18 int32 = int32(v4)
                    fmt.Println(fi12, v9)
                    fmt.Println(v2)
                    _ = v18
                }
                if v11 {
                    if (v3 != (v3 << ((1) & 63))) {
                        fmt.Println(v9)
                        v5 = 42
                    } else {
                        var v19 bool = v7
                        switch ("foo bar" + "") {
                        case "foo bar":
                            fmt.Println(v7)
                        }
                        v2 -= v2
                        v3 = (v3 & (^(v3)))
                        var v20 float64 = 1e10
                        _ = v19
                        _ = v20
                    }
                    v21 := GenS1{F0: "a", F1: ("a" + "a")}
                    _ = v21
                } else {
                    if v11 {
                        switch (fi12 % ((42) | 1)) {
                        case -2147483648:
                            if v11 {
                                fmt.Println(v3, v13, v5)
                                v22 := GenS2{F0: true}
                                var v23 bool = (!((!(!true)) && (-639 == -591)))
                                _ = len(v1)
                                if true {
                                    _ = len(v1)
                                }
                                v24 := GenS2{F0: v13}
                                var v25 uint16 = 2
                                _ = v22
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        case 885:
                            if ((v5 * v5) != v9) {
                                if false {
                                    v26 := GenS3{F0: 4294967295}
                                    fmt.Println(v13, v5, fi10)
                                    v5 += 0
                                    if (1 >= 32767) {
                                        v27 := GenS1{F0: "\n\t", F1: "\n\t"}
                                        v9 -= v9
                                        fmt.Println(v8, v11)
                                        _ = v27
                                    }
                                    _ = v26
                                } else {
                                    if false {
                                        fmt.Println(fi12, v3)
                                        switch (-(fi10)) {
                                        case 128:
                                            var v28 int16 = 176
                                            _ = v28
                                        default:
                                            if v7 {
                                                fmt.Println(v8)
                                            }
                                            v29 := GenS4{F0: (^(v9))}
                                            if false {
                                                v30 := GenS3{F0: 677}
                                                _ = v30
                                            } else {
                                                v7 = v11
                                            }
                                            _ = v29
                                        }
                                        v4 += v4
                                        v7 = (127 < 125)
                                    }
                                    var v31 uint16 = v4
                                    var v32 bool = v13
                                    _ = v31
                                    _ = v32
                                }
                                v8 = 1.5
                            } else {
                                var v33 uint16 = uint16(v2)
                                _ = v33
                            }
                        default:
                            var v34 bool = (!(!false))
                            if (v11 || (!(-701 < 0))) {
                                var v35 string = (("hello" + ("hello" + "\n\t")) + ("hello" + "foo bar"))
                                _ = v35
                            } else {
                                v36 := []int{(fi12 % ((fi12) | 1)), fi10, (fi10 >> ((93) & 63))}
                                _ = v36
                            }
                            var v37 int32 = 1
                            _ = v34
                            _ = v37
                        }
                        var v38 bool = (int8(v2) < 0)
                        _ = v38
                    } else {
                        if false {
                            if (v2 <= (^(v2))) {
                                v4 = 498
                                if (((v3 <= v3) && (-1 <= -9223372036854775808)) && false) {
                                    v6 += (v6 * v6)
                                    var v39 int = fi10
                                    v40 := GenS3{F0: 2}
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v41 float32 = ((v8 / ((0.1 - 1e10) * v8)) / v8)
                                    var v42 int8 = 113
                                    _ = v41
                                    _ = v42
                                }
                                var v43 uint = 0
                                var v44 float32 = (v8 - ((v8 + v8) + float32(v6)))
                                _ = v43
                                _ = v44
                            } else {
                                v45 := v1["\n\t"]
                                v46 := GenS4{F0: v9}
                                var v47 uint64 = (^(v9))
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                            var v48 float32 = v8
                            _ = v48
                        } else {
                            v49 := GenS4{F0: v9}
                            v50 := []int{int(v5)}
                            _ = v49
                            _ = v50
                        }
                        v9 = (v5 + v5)
                        var v51 int = fi12
                        v6 = v6
                        _ = v51
                    }
                    switch (-(fi12)) {
                    case -9223372036854775808:
                        var v52 int8 = int8(v3)
                        _ = v52
                    }
                }
            }
            switch ("hello" + "\n\t") {
            case "hello":
                var v53 uint8 = v2
                _ = v53
            case "\n\t":
                v54 := GenS2{F0: v13}
                _ = v54
            case "foo bar":
                var v55 uint32 = (v3 + 212)
                _ = v55
            default:
                if true {
                    var v56 uint32 = v3
                    if true {
                        v9 += 152
                        var v57 int = (fi10 | (fi12 | -2147483648))
                        _ = v57
                    }
                    if true {
                        if (v6 < v6) {
                            var v58 bool = v11
                            _ = v58
                        }
                        var v59 int64 = -1
                        v13 = (true && v11)
                        _ = v59
                    }
                    _ = v56
                } else {
                    if v11 {
                        var v60 float32 = 1.5
                        v61 := GenS4{F0: v9}
                        _ = v60
                        _ = v61
                    }
                    if ((^(v2)) >= 138) {
                        fmt.Println(v7, v13, v8)
                    } else {
                        v6 -= v6
                        v8 = v8
                        switch fi12 {
                        case 1:
                            var v62 int = fi10
                            _ = v62
                        case -1:
                            var v63 int8 = -1
                            var v64 int64 = -1
                            _ = v63
                            _ = v64
                        case -898:
                            var v65 uint32 = v3
                            _ = v65
                        default:
                            v66 := []int{}
                            _ = v66
                        }
                    }
                }
            }
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v7 = v7
    var v67 float32 = 0.1
    v4 = 2
    v68 := GenS5{F0: -1, F1: -1}
    var v69 string = "hello"
    v70 := GenS3{F0: 841}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v69
    _ = v7
    _ = v70
    _ = v8
    _ = v9
}
