package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 bool; F1 int8 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint32 = 719
    var v3 int32 = 2147483647
    v3 = (-(v3))
    if false {
        v4 := []int{32767}
        for fi5 := 0; fi5 < 16; fi5++ {
            switch (("\n\t" + "\n\t") + "foo bar") {
            case "foo bar":
                for fi6 := 0; fi6 < 2; fi6++ {
                    var v7 uint16 = 42
                    fmt.Println(fi6)
                    _ = fi6
                    _ = v7
                }
            case "":
                var v8 int8 = 0
                _ = v8
            case "\n\t":
                var v9 int = fi5
                _ = v9
            default:
                v1 -= (v1 - 80)
            }
            _ = fi5
        }
        v3 -= -2147483648
        _ = v4
    } else {
        v10 := GenS1{F0: 42}
        v11 := []int{}
        v3 -= (v3 / ((v3) | 1))
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        if (!(89 < -1)) {
            v1 += v1
            v13 := []int{-2147483648, fi12, (^(fi12)), (fi12 << ((v1) & 63))}
            _ = v13
        } else {
            v14 := GenS1{F0: -123}
            _ = v14
        }
        v3 -= v3
        v15 := []int{(fi12 * fi12), fi12, fi12, (fi12 | fi12)}
        _ = fi12
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        switch -2147483648 {
        case 32767:
            v3 = (v3 - (v3 / ((42) | 1)))
        case 9223372036854775807:
            v17 := GenS1{F0: -12}
            _ = v17
        case 165:
            v18 := GenS2{F0: 2147483647, F1: (^(fi16))}
            _ = v18
        }
        _ = fi16
    }
    v19 := []int{32768, 127}
    var v20 int = -9223372036854775808
    if true {
        var v21 uint32 = (v2 >> ((v1) & 63))
        _ = v21
    } else {
        for fi22 := 0; fi22 < 8; fi22++ {
            switch "foo bar" {
            case "\n\t":
                var v23 uint64 = 645
                _ = v23
            case "":
                var v24 int8 = 93
                _ = v24
            case "hello":
                var v25 uint = uint(v2)
                _ = v25
            }
            var v26 float64 = -2.25
            _ = fi22
            _ = v26
        }
        switch ("a" + "foo bar") {
        case "":
            for fi27 := 0; fi27 < 1; fi27++ {
                for fi28 := 0; fi28 < 2; fi28++ {
                    switch "a" {
                    case "hello":
                        var v29 uint = 65535
                        _ = v29
                    case "foo bar":
                        _ = len(v19)
                    }
                    var v30 int16 = 0
                    v30 = (v30 - v30)
                    _ = fi28
                    _ = v30
                }
                for fi31 := 0; fi31 < 2; fi31++ {
                    v3 = (v3 << ((v1) & 63))
                    if false {
                        if (1e10 < 0.0) {
                            fmt.Println(v1, v20)
                            if (92 == 0) {
                                switch (v20 << (((v1 - v1)) & 63)) {
                                case -677:
                                    var v32 int8 = 100
                                    _ = v32
                                case 891:
                                    var v33 bool = (!false)
                                    _ = v33
                                }
                            } else {
                                switch 9223372036854775807 {
                                case 32768:
                                    if true {
                                        fmt.Println(v2, v20, v20)
                                    }
                                case 1:
                                    v1 = v1
                                }
                                v34 := GenS3{F0: 2, F1: 850}
                                var v35 uint8 = 2
                                _ = v34
                                _ = v35
                            }
                        } else {
                            v36 := GenS2{F0: 0, F1: (v20 - (fi27 >> ((v1) & 63)))}
                            var v37 uint16 = 200
                            _ = v36
                            _ = v37
                        }
                        if (!true) {
                            v19 = append(v19, v20)
                            var v38 int32 = v3
                            v39 := GenS3{F0: 72, F1: 193}
                            v20 -= fi27
                            _ = v38
                            _ = v39
                        } else {
                            v2 = (v2 * (v2 >> ((v1) & 63)))
                            v3 -= (v3 | v3)
                            _ = len(v19)
                        }
                        if len(v19) > 0 {
                            v19[0] = (v20 >> ((18446744073709551615) & 63))
                        }
                        v1 = (^(v1))
                    }
                    _ = fi31
                }
                fmt.Println(v20, v2)
                _ = fi27
            }
        }
    }
    v20 += 85
    v40 := GenS4{F0: 110, F1: ((("a" + "hello") + "a") + "hello")}
    if ((1.5 * 1.5) == 1.5) {
        var v41 int64 = 1
        if true {
            v42 := GenS4{F0: 1, F1: "a"}
            var v43 int8 = -104
            if false {
                if ("hello" <= ("hello" + ("a" + ("\n\t" + "\n\t")))) {
                    fmt.Println(v41)
                    v43 = v43
                    _ = len(v19)
                } else {
                    v20 += 32768
                    var v44 int16 = 1
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    _ = v44
                }
                v45 := GenS1{F0: -1}
                _ = v45
            }
            _ = v42
            _ = v43
        } else {
            switch "a" {
            case "":
                v46 := GenS5{F0: false, F1: 1}
                _ = v46
            case "hello":
                fmt.Println(v20, v20)
            }
            var v47 int16 = 32767
            _ = v47
        }
        _ = v41
    } else {
        v3 = 0
        fmt.Println(v2)
        var v48 uint16 = 0
        _ = v48
    }
    var v49 float32 = float32(-2.25)
    for fi50 := 0; fi50 < 11; fi50++ {
        switch fi50 {
        case 127:
            v19 = append(v19, v20)
        case 18:
            var v51 int32 = (v3 + (v3 ^ (v3 - 0)))
            _ = v51
        default:
            if (0 != 2) {
                var v52 bool = (v2 > v2)
                var v53 bool = v52
                _ = v52
                _ = v53
            } else {
                var v54 uint8 = 42
                _ = v54
            }
        }
        _ = fi50
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v40
    _ = v49
}
