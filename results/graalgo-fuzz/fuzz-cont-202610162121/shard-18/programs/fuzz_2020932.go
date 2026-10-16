package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int; F1 string; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 int32 = 42
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    switch 128 {
    case 32767:
        v6 := v5[2147483647]
        _ = v6
    case 1:
        var v7 bool = (-1 <= 385)
        _ = v7
    default:
        if (42 != -32768) {
            var v8 uint32 = 1
            v9 := GenS1{F0: 32767, F1: -1}
            _ = v8
            _ = v9
        } else {
            v1 = -2147483648
            switch 42 {
            case 127:
                if ((("\n\t" + "") + "\n\t") <= ("hello" + "\n\t")) {
                    var v10 int = 4294967295
                    v1 -= (v1 >> ((18446744073709551615) & 63))
                    _ = v10
                }
            case 9223372036854775807:
                for fi11 := 0; fi11 < 18; fi11++ {
                    v12 := GenS1{F0: 0, F1: 677}
                    var v13 uint8 = 255
                    switch fi11 {
                    case 0:
                        switch ("hello" + "\n\t") {
                        case "\n\t":
                            fmt.Println(fi11, v13, v13)
                        case "hello":
                            for fi14 := 0; fi14 < 13; fi14++ {
                                v15 := []int{630, -9223372036854775808}
                                if true {
                                    switch "\n\t" {
                                    case "":
                                        if false {
                                            var v16 uint16 = 65535
                                            v17 := GenS1{F0: 1, F1: 0}
                                            v18 := GenS1{F0: 444, F1: -929}
                                            _ = v16
                                            _ = v17
                                            _ = v18
                                        } else {
                                            fmt.Println(v13)
                                        }
                                    case "hello":
                                        var v19 bool = false
                                        _ = v19
                                    case "\n\t":
                                        v20 := GenS2{F0: 1, F1: ("a" + ("\n\t" + ("" + "foo bar"))), F2: 1}
                                        _ = v20
                                    }
                                    var v21 int64 = -847
                                    fmt.Println(fi14, fi11)
                                    var v22 string = "\n\t"
                                    var v23 string = (v22 + "foo bar")
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                } else {
                                    v24 := GenS1{F0: 1000, F1: 0}
                                    v13 = 22
                                    fmt.Println(v13)
                                    _ = v24
                                }
                                fmt.Println(fi14)
                                _ = fi14
                                _ = v15
                            }
                        }
                    case -1:
                        var v25 bool = (-128 == -128)
                        _ = v25
                    case -524:
                        switch fi11 {
                        case 949:
                            v26 := GenS1{F0: 0, F1: 1}
                            _ = v26
                        case 0:
                            for fi27 := 0; fi27 < 1; fi27++ {
                                if true {
                                    v28 := GenS1{F0: -1, F1: 9223372036854775807}
                                    delete(v5, 42)
                                    _ = v28
                                } else {
                                    var v29 uint16 = 667
                                    if (false && (false && false)) {
                                        v1 += v1
                                        v30 := GenS1{F0: 1, F1: -46}
                                        v1 += -2147483648
                                        fmt.Println(fi27, fi11)
                                        _ = v30
                                    } else {
                                        var v31 int8 = -128
                                        _ = v31
                                    }
                                    v32 := []int{1, fi27, fi11}
                                    _ = v29
                                    _ = v32
                                }
                                _ = fi27
                            }
                        default:
                            v33 := GenS3{F0: -1}
                            _ = v33
                        }
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case -2147483648:
                v3[1] = 0
            }
            v34 := GenS2{F0: -996, F1: "hello", F2: 1}
            _ = v34
        }
    }
    var v35 float32 = ((-2.25 + (3.14159 - 1e10)) - 3.14159)
    v36 := GenS2{F0: 127, F1: ("foo bar" + ""), F2: 480}
    if (("foo bar" <= "hello") || false) {
        var v37 int8 = int8(v1)
        _ = v37
    }
    var v38 int32 = v1
    for fi39 := 0; fi39 < 3; fi39++ {
        var v40 bool = false
        _ = fi39
        _ = v40
    }
    var v41 uint16 = 255
    if ((0 >= -32768) || false) {
        var v42 int64 = 42
        _ = v42
    }
    switch 42 {
    case 233:
        for fi43 := 0; fi43 < 13; fi43++ {
            v44 := GenS4{F0: 255, F1: 0}
            if (4294967295 != 792) {
                if false {
                    var v45 int = fi43
                    if false {
                        v46 := GenS4{F0: (v41 - v41), F1: 0}
                        var v47 bool = false
                        var v48 uint32 = 2
                        _ = v46
                        _ = v47
                        _ = v48
                    }
                    fmt.Println(fi43, v41)
                    _ = v45
                }
            } else {
                var v49 uint8 = 172
                _ = v49
            }
            v1 = 2147483647
            _ = fi43
            _ = v44
        }
    }
    for fi50 := 0; fi50 < 8; fi50++ {
        var v51 bool = (!(199 == 255))
        var v52 uint16 = v41
        if true {
            v52 += 255
            v52 = (^(v41))
            for fi53 := 0; fi53 < 12; fi53++ {
                var v54 uint64 = 0
                _ = fi53
                _ = v54
            }
        }
        _ = fi50
        _ = v51
        _ = v52
    }
    var v55 uint = 65535
    v3[0] = 32767
    var v56 float64 = 0.1
    v57 := []int{2147483647, 86, 32767, 285}
    var v58 int = -881
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v41
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
