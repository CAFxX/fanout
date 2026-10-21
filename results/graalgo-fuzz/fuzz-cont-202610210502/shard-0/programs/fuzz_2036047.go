package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint; F2 int16 }
type GenS4 struct { F0 int32; F1 int32; F2 int32 }
type GenS5 struct { F0 int; F1 bool; F2 bool }
func genFn0() int8 {
    if (-32768 < 42) {
        if true {
            var v6 uint16 = 32767
            var v7 int8 = -1
            _ = v6
            _ = v7
        } else {
            if true {
                var v8 uint16 = 0
                _ = v8
            } else {
                var v9 int64 = 0
                _ = v9
            }
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 int64 = 9223372036854775807
                v12 := []int{944}
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 int16 = 32767
                    v15 := GenS1{F0: 0}
                    var v16 float64 = ((0.0 - -2.25) - float64(float32((0.1 - 0.0))))
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = fi10
                _ = v11
                _ = v12
            }
            var v17 float32 = (1e10 + (1.5 - 3.14159))
            _ = v17
        }
        if (42 >= 18446744073709551615) {
            switch 42 {
            case 767:
                var v18 uint8 = 2
                _ = v18
            }
            if (((!(!(1e10 <= (3.14159 + 0.1)))) && (!(65535 >= 1))) || (!(false && true))) {
                if (42 < 1) {
                    var v19 uint32 = 42
                    _ = v19
                }
                switch ((("a" + ("hello" + "hello")) + "hello") + "\n\t") {
                case "a":
                    var v20 int64 = -34
                    _ = v20
                case "\n\t":
                    var v21 uint = 596
                    _ = v21
                case "foo bar":
                    for fi22 := 0; fi22 < 5; fi22++ {
                        if true {
                            var v23 int = (fi22 % ((-2147483648) | 1))
                            _ = v23
                        } else {
                            switch "" {
                            case "hello":
                                v24 := GenS1{F0: 2}
                                _ = v24
                            case "a":
                                v25 := GenS2{F0: 42}
                                _ = v25
                            case "\n\t":
                                v26 := []int{}
                                _ = v26
                            }
                        }
                        fmt.Println(fi22)
                        _ = fi22
                    }
                default:
                    v27 := []int{-1, 128, -2147483648, 9223372036854775807}
                    _ = v27
                }
                var v28 int = -979
                _ = v28
            } else {
                switch 127 {
                case -2147483648:
                    var v29 int16 = 656
                    _ = v29
                case -830:
                    var v30 int8 = 127
                    _ = v30
                case 0:
                    var v31 float32 = (-2.25 - 1e10)
                    _ = v31
                default:
                    v32 := GenS3{F0: "a", F1: 65535, F2: 341}
                    _ = v32
                }
            }
        } else {
            var v33 float64 = -2.25
            if true {
                v33 = float64(3.14159)
                v34 := []int{42, 0}
                v35 := GenS4{F0: -1, F1: 0, F2: -123}
                _ = v34
                _ = v35
            }
            _ = v33
        }
    } else {
        var v36 uint16 = 65535
        v37 := GenS5{F0: 32767, F1: true, F2: false}
        _ = v36
        _ = v37
    }
    v38 := GenS5{F0: 9223372036854775807, F1: (!false), F2: false}
    var v39 float32 = 0.1
    _ = v38
    _ = v39
    return 127
}
func genFn1(p0 uint8, p1 string, p2 uint64) uint64 {
    switch -2147483648 {
    case 9223372036854775807:
        var v40 uint16 = 2
        _ = v40
    default:
        v41 := []int{-902, 0, -9223372036854775808}
        _ = v41
    }
    var v42 uint64 = p2
    for fi43 := 0; fi43 < 18; fi43++ {
        p1 = "\n\t"
        _ = fi43
    }
    if true {
        fmt.Println(p1)
    } else {
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v42
    return v42
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]string)
    v3 := []string{"hello", "", ("\n\t" + "hello"), "foo bar"}
    var v4 uint16 = 255
    var v5 float64 = float64((-2.25 - -2.25))
    for fi44 := 0; fi44 < 1; fi44++ {
        _ = len(v3)
        _ = fi44
    }
    v45 := GenS5{F0: 32767, F1: (0 >= 42), F2: (269 != 131)}
    var v46 uint32 = 345
    var v47 uint32 = 2147483648
    v3 = append(v3, ("foo bar" + ("" + "a")))
    fmt.Println(v1, v4)
    switch "" {
    case "foo bar":
        _ = len(v2)
    case "hello":
        var v48 uint8 = 58
        _ = v48
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        for fi49 := 0; fi49 < 8; fi49++ {
            v47 -= (v47 + (^(v46)))
            _ = fi49
        }
    }
    var v50 int8 = 100
    for fi51 := 0; fi51 < 2; fi51++ {
        v47 += 65535
        var v52 int16 = int16(v1)
        _ = fi51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v50
}
