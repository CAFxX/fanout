package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS4 struct { F0 int64; F1 int; F2 bool }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 bool; F1 int64; F2 int16 }
type GenS7 struct { F0 int }
func genFn0(p0 int) uint {
    switch p0 {
    case 4294967295:
        v5 := GenS1{F0: "foo bar"}
        _ = v5
    case -1:
        fmt.Println(p0)
    case 32768:
        var v6 uint16 = 32767
        _ = v6
    default:
        fmt.Println(p0)
    }
    v7 := GenS2{F0: false, F1: 765}
    if false {
        var v8 int = p0
        _ = v8
    }
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 float64 = 0.0
    var v2 int32 = 1
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    switch "" {
    case "a":
        var v9 bool = (325 > -121)
        _ = v9
    case "":
        if true {
            var v10 int16 = 1
            v11 := GenS1{F0: "a"}
            _ = v10
            _ = v11
        } else {
            delete(v4, -1)
            fmt.Println(v1, v1)
            var v12 bool = false
            _ = v12
        }
    }
    v13 := GenS3{F0: 255, F1: -1, F2: 42}
    v14 := GenS4{F0: -1, F1: 573, F2: true}
    switch (("hello" + ((("hello" + ("" + ((("" + ((("" + "a") + "\n\t") + "\n\t")) + "hello") + "hello"))) + "hello") + "")) + "hello") {
    case "a":
        var v15 float32 = (0.0 - 0.0)
        _ = v15
    case "foo bar":
        v16 := GenS3{F0: 103, F1: 53, F2: 1}
        _ = v16
    case "hello":
        v17 := []int{-391, 128, 4294967295}
        _ = v17
    default:
        fmt.Println(v2, v1)
    }
    fmt.Println(v2, v2)
    v1 = v1
    var v18 uint8 = 211
    switch -488 {
    case 128:
        v18 = 42
    case 32768:
        switch -135 {
        case -67:
            _ = len(v3)
        case 109:
            switch 4294967295 {
            case -548:
                v18 += v18
            case 386:
                if ("hello" > (("" + ("hello" + "foo bar")) + "a")) {
                    v19 := GenS4{F0: -1, F1: int(v18), F2: true}
                    v20 := GenS5{F0: 0, F1: 1}
                    _ = v19
                    _ = v20
                }
            case 1:
                v21 := GenS6{F0: (1e10 != (-2.25 + -2.25)), F1: 9223372036854775807, F2: 1}
                _ = v21
            }
        case 838:
            v1 = v1
        default:
            for fi22 := 0; fi22 < 18; fi22++ {
                for fi23 := 0; fi23 < 7; fi23++ {
                    v1 = v1
                    _ = fi23
                }
                v24 := v3[""]
                for fi25 := 0; fi25 < 2; fi25++ {
                    v2 = v2
                    if true {
                        if (0.1 == (3.14159 + float64(1e10))) {
                            v26 := v3["hello"]
                            _ = v26
                        } else {
                            if (true && (!true)) {
                                if false {
                                    var v27 int16 = -1
                                    v28 := GenS4{F0: -719, F1: -672, F2: true}
                                    _ = v27
                                    _ = v28
                                }
                                fmt.Println(v24)
                                v1 += v1
                            }
                            v1 -= v1
                            fmt.Println(v2)
                        }
                        var v29 int16 = 32767
                        if false {
                            v30 := []int{}
                            _ = len(v4)
                            _ = v30
                        }
                        _ = v29
                    }
                    if (("foo bar" + (("hello" + "\n\t") + "a")) >= "hello") {
                        var v31 string = "foo bar"
                        switch v31 {
                        case "a":
                            v32 := GenS1{F0: v31}
                            _ = v32
                        case "\n\t":
                            v3[v31] = v24
                        case "":
                            v18 += (v18 >> ((uint(v18)) & 63))
                        default:
                            if (v24 != (^(v24))) {
                                v4[(^(fi25))] = ((v18 * v18) != 127)
                            } else {
                                var v33 int64 = 0
                                switch "" {
                                case "a":
                                    var v34 uint64 = 18446744073709551615
                                    _ = v34
                                case "":
                                    v24 = (v24 * (^(v24)))
                                }
                                if (18446744073709551615 > 94) {
                                    var v35 int16 = 365
                                    var v36 string = v31
                                    fmt.Println(v33)
                                    _ = v35
                                    _ = v36
                                }
                                v31 = v31
                                _ = v33
                            }
                        }
                        _ = v31
                    }
                    switch "a" {
                    case "":
                        fmt.Println(fi22, v2)
                    case "foo bar":
                        var v37 uint32 = (v24 >> ((1) & 63))
                        _ = v37
                    default:
                        var v38 uint64 = 42
                        _ = v38
                    }
                    if false {
                        var v39 uint = 42
                        fmt.Println(fi22, v1, fi22)
                        var v40 uint = genFn0((fi22 * 127))
                        _ = v39
                        _ = v40
                    }
                    var v41 uint64 = 568
                    var v42 float32 = 3.14159
                    if (!((1 > 2) || (((v24 <= 0) || true) || (((((0.0 - float32((v1 / v1))) / v42) * 0.1) - v42) < v42)))) {
                        var v43 uint8 = v18
                        _ = v43
                    } else {
                        var v44 int16 = 1
                        v45 := GenS1{F0: "\n\t"}
                        _ = v44
                        _ = v45
                    }
                    _ = fi25
                    _ = v41
                    _ = v42
                }
                _ = fi22
                _ = v24
            }
        }
    case -1:
        if (((!(true && (!(2147483647 == -1)))) && (1 < 2)) || true) {
            var v46 float32 = 0.1
            v46 += ((((1.5 + (((v46 + v46) - v46) + 1e10)) + v46) + (1e10 + 3.14159)) / ((v46 - v46) * v46))
            _ = v46
        } else {
            var v47 uint = 1
            _ = v47
        }
    }
    fmt.Println(v2)
    v4[-9223372036854775808] = ((v2 * (v2 >> ((uint(v18)) & 63))) >= v2)
    v48 := GenS5{F0: 1, F1: 127}
    v1 -= v1
    switch -1 {
    case 128:
        switch 42 {
        case 128:
            if true {
                v49 := GenS7{F0: -335}
                _ = v49
            }
        case 482:
            var v50 string = "\n\t"
            _ = v50
        }
    case 4294967295:
        switch 0 {
        case 4294967295:
            switch "" {
            case "":
                delete(v4, 500)
            }
        case -1:
            v4[128] = false
        default:
            v18 += (^(v18))
        }
    case 42:
        v51 := GenS5{F0: -616, F1: -1}
        _ = v51
    default:
        if false {
            v52 := v3["hello"]
            if (163 >= 127) {
                v52 = v52
                switch -2147483648 {
                case 4294967295:
                    var v53 uint32 = 901
                    _ = v53
                }
            } else {
                v52 = v52
            }
            _ = v52
        }
    }
    fmt.Println(v18)
    var v54 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v54
}
