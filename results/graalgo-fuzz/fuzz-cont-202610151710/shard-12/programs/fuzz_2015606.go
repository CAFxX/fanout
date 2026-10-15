package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 string }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int8; F1 int8 }

func main() {
    v1 := []string{(("foo bar" + "a") + "\n\t"), "foo bar", "", ("a" + "hello")}
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 bool = true
    v5 := make(map[string]int)
    switch "" {
    case "":
        fmt.Println(v4)
    case "a":
        var v6 int64 = 0
        _ = v6
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: (((("" + "a") + ("" + ("foo bar" + ("" + "a")))) + "a") + "\n\t"), F1: -526, F2: ("foo bar" + (("foo bar" + ("" + "foo bar")) + (("hello" + "a") + "a")))}
        v5[(("foo bar" + "foo bar") + "a")] = -842
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 2
    var v10 bool = (!(778 <= 1))
    if v4 {
        var v11 int64 = 42
        for fi12 := 0; fi12 < 5; fi12++ {
            if (v4 && (1 > 1)) {
                v13 := GenS2{F0: (v9 | 0), F1: (^(v9))}
                fmt.Println(v9)
                _ = v13
            }
            if v10 {
                var v14 float64 = 3.14159
                v14 = v14
                _ = len(v5)
                _ = v14
            } else {
                v15 := GenS1{F0: "", F1: (v11 % ((v11) | 1)), F2: ("\n\t" + "foo bar")}
                v5["foo bar"] = (fi12 * -200)
                _ = v15
            }
            var v16 int32 = int32(v11)
            _ = fi12
            _ = v16
        }
        var v17 bool = (-32768 < 1)
        _ = v11
        _ = v17
    }
    delete(v5, "")
    var v18 bool = ((1000 <= -1) && (441 <= 42))
    var v19 uint64 = v9
    var v20 float64 = (-2.25 / -2.25)
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 3; fi22++ {
            if v4 {
                if true {
                    v23 := GenS1{F0: "a", F1: -9223372036854775808, F2: "a"}
                    var v24 uint32 = 2147483648
                    _ = v23
                    _ = v24
                }
                var v25 int = fi22
                _ = v25
            } else {
                switch "hello" {
                case "":
                    v10 = (789 == -32768)
                default:
                    if v18 {
                        v26 := GenS1{F0: "foo bar", F1: 9223372036854775807, F2: "\n\t"}
                        _ = v26
                    } else {
                        v27 := GenS3{F0: 42, F1: 42}
                        v28 := GenS4{F0: 2, F1: true}
                        if v18 {
                            var v29 int32 = 466
                            fmt.Println(v29, v29, v19)
                            _ = v29
                        }
                        v18 = false
                        var v30 float64 = v20
                        var v31 int64 = -9223372036854775808
                        _ = v27
                        _ = v28
                        _ = v30
                        _ = v31
                    }
                }
                v32 := GenS4{F0: 255, F1: v4}
                _ = v32
            }
            _ = fi22
        }
        v20 = (v20 - v20)
        if v18 {
            if false {
                if v4 {
                    delete(v2, ("\n\t" + ""))
                } else {
                    if v10 {
                        var v33 int64 = 1
                        if false {
                            v34 := GenS2{F0: (v19 % (((v9 + 18446744073709551615)) | 1)), F1: 42}
                            _ = v34
                        }
                        if true {
                            fmt.Println(fi21, fi21, fi21)
                            v35 := GenS2{F0: (^(v9)), F1: (^(v9))}
                            if true {
                                v36 := GenS5{F0: 42, F1: 1}
                                fmt.Println(v20)
                                _ = v36
                            } else {
                                var v37 int32 = -408
                                for fi38 := 0; fi38 < 16; fi38++ {
                                    v4 = (2 >= 255)
                                    var v39 uint64 = (^(v9))
                                    v40 := GenS6{F0: 100, F1: 1}
                                    v37 -= (v37 + (v37 + v37))
                                    switch "\n\t" {
                                    case "\n\t":
                                        _ = len(v1)
                                    case "a":
                                        switch ("foo bar" + "foo bar") {
                                        case "foo bar":
                                            var v41 float32 = (1.5 + 0.1)
                                            _ = v41
                                        case "hello":
                                            v9 -= (v9 & (v39 ^ (v39 ^ (^(v9)))))
                                        case "\n\t":
                                            if (2147483647 > 301) {
                                                v9 += v39
                                            }
                                        }
                                    case "foo bar":
                                        var v42 bool = false
                                        _ = v42
                                    default:
                                        switch fi38 {
                                        case -946:
                                            var v43 float32 = 0.0
                                            _ = v43
                                        }
                                    }
                                    if (1000 == 420) {
                                        if false {
                                            var v44 uint64 = (v19 ^ v19)
                                            _ = v44
                                        }
                                    } else {
                                        if (v4 && false) {
                                            v20 = (v20 - v20)
                                            v18 = (v10 && (("a" + "hello") == ("a" + ("foo bar" + "foo bar"))))
                                            switch fi38 {
                                            case -2147483648:
                                                v20 += (1.5 - v20)
                                            case -9223372036854775808:
                                                v33 -= v33
                                            case 280:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v45 := GenS2{F0: 4294967295, F1: 960}
                                                    _ = v45
                                                }
                                            }
                                            if v18 {
                                                var v46 int64 = v33
                                                fmt.Println(fi21, v4)
                                                _ = v46
                                            } else {
                                                if true {
                                                    fmt.Println(v39)
                                                } else {
                                                    v47 := GenS4{F0: 36, F1: false}
                                                    _ = v47
                                                }
                                                var v48 int = fi38
                                                _ = v48
                                            }
                                            var v49 int64 = (v33 + (v33 + 0))
                                            _ = v49
                                        } else {
                                            v9 -= 145
                                            fmt.Println(v10, fi21)
                                        }
                                    }
                                    v19 = 511
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                                _ = v37
                            }
                            _ = v35
                        } else {
                            v50 := GenS4{F0: 255, F1: false}
                            _ = v50
                        }
                        _ = v33
                    } else {
                        v4 = false
                    }
                    v9 += (^(v19))
                }
            } else {
                fmt.Println(v10, v20, v10)
            }
            _ = len(v3)
            var v51 uint64 = v19
            _ = v51
        }
        _ = fi21
    }
    for fi52 := 0; fi52 < 12; fi52++ {
        var v53 int = fi52
        var v54 int64 = 1
        _ = fi52
        _ = v53
        _ = v54
    }
    fmt.Println(v10)
    var v55 int8 = 0
    v19 -= v19
    var v56 float64 = v20
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v9
}
