package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 int16 }
func genFn0() int16 {
    switch ("hello" + ("\n\t" + "foo bar")) {
    case "\n\t":
        if false {
            var v4 uint16 = 65535
            v5 := []int{0, 32768, 42, -1}
            switch -17 {
            case 9223372036854775807:
                var v6 int32 = -1
                _ = v6
            case 42:
                _ = len(v5)
            case -497:
                for fi7 := 0; fi7 < 19; fi7++ {
                    fmt.Println(fi7, fi7)
                    v8 := GenS1{F0: ("\n\t" + (("" + "") + (("\n\t" + "hello") + "foo bar"))), F1: ("" + "\n\t"), F2: 128}
                    _ = fi7
                    _ = v8
                }
            }
            _ = v4
            _ = v5
        } else {
            for fi9 := 0; fi9 < 20; fi9++ {
                if true {
                    for fi10 := 0; fi10 < 2; fi10++ {
                        v11 := GenS1{F0: (((("\n\t" + "foo bar") + "a") + "") + ("\n\t" + "foo bar")), F1: "hello", F2: 89}
                        if false {
                            fmt.Println(fi10, fi9)
                        } else {
                            v12 := GenS2{F0: fi9, F1: 629, F2: -2147483648}
                            if true {
                                v13 := []int{145, 76, 0, -885}
                                if len(v13) > 0 {
                                    v13[0] = fi9
                                }
                                if false {
                                    if (227 != 1) {
                                        var v14 uint8 = 2
                                        v14 = 1
                                        _ = v14
                                    }
                                }
                                _ = v13
                            } else {
                                var v15 uint16 = 3
                                _ = v15
                            }
                            var v16 float64 = (3.14159 * 1e10)
                            v17 := []int{(fi10 / ((-373) | 1)), 4294967295, 127, fi9}
                            _ = v12
                            _ = v16
                            _ = v17
                        }
                        _ = fi10
                        _ = v11
                    }
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        var v18 int64 = -458
                        _ = v18
                    case "\n\t":
                        var v19 uint8 = 127
                        _ = v19
                    case "":
                        var v20 int16 = -1
                        _ = v20
                    }
                    var v21 string = ""
                    v21 = v21
                    _ = v21
                }
                _ = fi9
            }
            var v22 uint64 = 2
            var v23 int64 = -840
            _ = v22
            _ = v23
        }
    default:
        switch 1 {
        case 901:
            var v24 float64 = 3.14159
            _ = v24
        case -479:
            for fi25 := 0; fi25 < 10; fi25++ {
                var v26 uint32 = 42
                v26 += (v26 + (v26 >> ((737) & 63)))
                for fi27 := 0; fi27 < 18; fi27++ {
                    v28 := GenS3{F0: 42, F1: uint64(v26), F2: 65}
                    v26 -= (v26 / ((207) | 1))
                    _ = fi27
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
        default:
            var v29 int16 = 969
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        switch "hello" {
        case "\n\t":
            var v31 float32 = -2.25
            _ = v31
        }
        v32 := GenS1{F0: "\n\t", F1: "a", F2: 1}
        v33 := GenS1{F0: ("a" + "\n\t"), F1: "hello", F2: 16}
        _ = fi30
        _ = v32
        _ = v33
    }
    v34 := []int{2147483647, 9223372036854775807, 128}
    var v35 uint16 = 260
    fmt.Println(v35)
    _ = v34
    _ = v35
    return 473
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 float32 = (1e10 + -2.25)
    fmt.Println(v3)
    v3 = float32(3.14159)
    fmt.Println(v3)
    var v36 string = ""
    for fi37 := 0; fi37 < 15; fi37++ {
        var v38 uint8 = 69
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        var v40 bool = (1 == 100)
        for fi41 := 0; fi41 < 20; fi41++ {
            if (-1 >= 42) {
                if v40 {
                    switch fi39 {
                    case 127:
                        v42 := GenS4{F0: -32768, F1: 2}
                        _ = v42
                    }
                    v43 := GenS3{F0: 572, F1: 42, F2: 1}
                    var v44 uint8 = 255
                    var v45 int32 = 0
                    var v46 int16 = 326
                    _ = v43
                    _ = v44
                    _ = v45
                    _ = v46
                } else {
                    v47 := GenS3{F0: 540, F1: 573, F2: 1}
                    var v48 uint32 = 763
                    _ = v47
                    _ = v48
                }
            }
            var v49 int64 = 9223372036854775807
            _ = fi41
            _ = v49
        }
        _ = fi39
        _ = v40
    }
    v50 := GenS5{F0: v36}
    if false {
        v51 := []int{9223372036854775807}
        if false {
            v3 -= 1.5
        } else {
            v3 = (v3 / 3.14159)
            var v52 uint16 = 42
            v53 := GenS6{F0: -87, F1: 1000}
            _ = v52
            _ = v53
        }
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v50
}
