package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
func genFn0() int64 {
    var v7 uint8 = 128
    v7 -= v7
    _ = v7
    return -1
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    var v5 int8 = 1
    var v6 uint16 = 2
    if true {
        var v8 string = "foo bar"
        switch v8 {
        case "":
            if false {
                v9 := GenS1{F0: 9223372036854775807, F1: 32767, F2: (v5 | v5)}
                for fi10 := 0; fi10 < 18; fi10++ {
                    switch fi10 {
                    case 32767:
                        for fi11 := 0; fi11 < 6; fi11++ {
                            fmt.Println(fi11)
                            _ = fi11
                        }
                    default:
                        v5 = (v5 % (((v5 / (((-(v5))) | 1))) | 1))
                    }
                    _ = fi10
                }
                _ = v9
            }
        default:
            var v12 uint64 = 1
            _ = v12
        }
        switch -9223372036854775808 {
        case -898:
            if true {
                v8 = "foo bar"
                v13 := GenS1{F0: 2147483647, F1: (v6 * v6), F2: 100}
                var v14 uint = 685
                _ = v13
                _ = v14
            } else {
                v15 := GenS1{F0: 32768, F1: (v6 / ((v6) | 1)), F2: v5}
                if (1 == 2) {
                    var v16 uint = 2
                    fmt.Println(v6, v16, v16)
                    _ = v16
                }
                _ = v15
            }
        default:
            var v17 uint16 = v6
            _ = v17
        }
        _ = v8
    } else {
        var v18 int16 = 32767
        _ = len(v3)
        _ = v18
    }
    _ = len(v2)
    v19 := GenS1{F0: 591, F1: 0, F2: v5}
    v20 := GenS2{F0: (v6 + v6)}
    for fi21 := 0; fi21 < 4; fi21++ {
        switch (^(fi21)) {
        case -2147483648:
            var v22 int64 = 524
            _ = v22
        default:
            var v23 int8 = v5
            _ = v23
        }
        v6 -= (v6 | (v6 ^ (v6 % ((v6) | 1))))
        _ = fi21
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        if (379 == 4294967295) {
            var v25 float32 = (-2.25 + 0.1)
            _ = v25
        } else {
            if (!(0.1 <= 0.1)) {
                v26 := GenS3{F0: 2147483647}
                if true {
                    var v27 int32 = 2147483647
                    _ = v27
                }
                var v28 string = ("a" + "a")
                _ = v26
                _ = v28
            } else {
                for fi29 := 0; fi29 < 7; fi29++ {
                    if true {
                        v6 = v6
                        v30 := GenS2{F0: (v6 - 0)}
                        var v31 int16 = -571
                        _ = v30
                        _ = v31
                    } else {
                        var v32 uint16 = 664
                        _ = v32
                    }
                    var v33 float32 = 3.14159
                    var v34 int16 = 32767
                    _ = fi29
                    _ = v33
                    _ = v34
                }
                for fi35 := 0; fi35 < 20; fi35++ {
                    fmt.Println(fi24, fi24, v5)
                    switch (-(fi35)) {
                    case 795:
                        switch (("hello" + ("" + (("\n\t" + (("a" + "") + (((("\n\t" + "foo bar") + ("foo bar" + "foo bar")) + (("\n\t" + ((("a" + ("foo bar" + ("foo bar" + "a"))) + "hello") + "hello")) + "\n\t")) + (("hello" + "a") + "a")))) + ("" + (("" + "\n\t") + "a"))))) + "\n\t") {
                        case "a":
                            v5 = (v5 * (v5 << ((255) & 63)))
                        case "\n\t":
                            if (((float32(((3.14159 * 3.14159) / float64((0.0 + (1e10 / 1e10))))) * 1.5) - 3.14159) == 3.14159) {
                                fmt.Println(v5)
                                switch ("foo bar" + ((("foo bar" + ("" + "a")) + "\n\t") + "hello")) {
                                case "hello":
                                    v36 := GenS4{F0: 698}
                                    _ = v36
                                default:
                                    switch "\n\t" {
                                    case "":
                                        var v37 uint8 = 127
                                        _ = v37
                                    case "hello":
                                        v38 := []int{}
                                        _ = v38
                                    default:
                                        v39 := GenS1{F0: (fi24 / ((int(v5)) | 1)), F1: 1, F2: 127}
                                        _ = v39
                                    }
                                }
                                var v40 int = (fi35 << ((1) & 63))
                                _ = v40
                            } else {
                                var v41 uint64 = 514
                                var v42 int32 = 0
                                _ = v41
                                _ = v42
                            }
                        }
                    default:
                        var v43 int16 = 32767
                        _ = v43
                    }
                    _ = fi35
                }
                v44 := v1[(-(fi24))]
                _ = v44
            }
            var v45 float64 = float64(3.14159)
            _ = v45
        }
        _ = fi24
    }
    v5 = 100
    fmt.Println(v6, v6)
    v6 += v6
    var v46 uint64 = 4294967295
    fmt.Println(v5, v5, v5)
    var v47 int64 = -290
    var v48 float64 = 3.14159
    fmt.Println(v48, v48, v46)
    var v49 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
