package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int16 }
func genFn0() uint32 {
    var v6 int8 = 127
    switch 127 {
    case 2147483647:
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: "a", F1: 2, F2: 866}
            v9 := GenS2{F0: 42, F1: 241}
            var v10 int = (fi7 % ((fi7) | 1))
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    default:
        v6 = v6
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS3{F0: 0, F1: 990}
            var v14 bool = false
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 int = (fi11 << ((1) & 63))
        var v16 uint16 = 0
        _ = fi11
        _ = v15
        _ = v16
    }
    v6 = v6
    _ = v6
    return 2
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{("\n\t" + ("foo bar" + "hello")), "hello", "foo bar"}
    var v3 int8 = 1
    var v4 int8 = (v3 >> ((42) & 63))
    var v5 float64 = (1e10 / 1.5)
    fmt.Println(v4, v5)
    var v17 float64 = (v5 - ((float64((float32(0.0) - (1.5 - (-2.25 / float32((v5 / (v5 / ((1e10 - (1.5 / v5)) - (((v5 / v5) * ((v5 * -2.25) * v5)) + float64((((0.0 + 0.1) + 1.5) / 1e10))))))))))) + -2.25) * 0.1))
    var v18 uint64 = 2
    if (false && false) {
        for fi19 := 0; fi19 < 1; fi19++ {
            v17 = (0.0 - ((v5 / v17) * (0.0 / ((float64(1.5) - v5) / v5))))
            var v20 uint8 = 128
            v21 := v1["foo bar"]
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 3; fi22++ {
            v3 += -128
            var v23 float64 = 0.0
            if (!(true || ((!false) || ((v4 | (v3 * (v4 - v4))) >= (v3 >> ((0) & 63)))))) {
                switch (("hello" + "hello") + "hello") {
                case "":
                    v18 += (^(v18))
                case "hello":
                    switch (fi22 * int(v3)) {
                    case -1:
                        if true {
                            v17 = v17
                            v24 := GenS3{F0: 433, F1: 2}
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                            _ = v24
                        }
                    case 4294967295:
                        switch 1 {
                        case 127:
                            var v25 int16 = 1000
                            _ = v25
                        case 32767:
                            v26 := GenS1{F0: "\n\t", F1: 4294967295, F2: 65535}
                            _ = v26
                        case -7:
                            v17 += 1e10
                        default:
                            switch ("" + "\n\t") {
                            case "hello":
                                v3 = (v3 * v3)
                            case "foo bar":
                                for fi27 := 0; fi27 < 7; fi27++ {
                                    v5 += (v5 / 1.0)
                                    _ = fi27
                                }
                            case "a":
                                v3 -= v3
                            }
                        }
                    }
                case "\n\t":
                    switch ("hello" + ("\n\t" + "\n\t")) {
                    case "":
                        if false {
                            for fi28 := 0; fi28 < 19; fi28++ {
                                var v29 uint64 = 717
                                v23 -= float64(1e10)
                                _ = fi28
                                _ = v29
                            }
                        } else {
                            if (128 >= (fi22 % ((fi22) | 1))) {
                                for fi30 := 0; fi30 < 1; fi30++ {
                                    delete(v1, ("" + "foo bar"))
                                    var v31 uint = 1
                                    v4 = (v3 / (((v3 & v3)) | 1))
                                    _ = fi30
                                    _ = v31
                                }
                                var v32 bool = (0 >= -487)
                                fmt.Println(v4, v32, v4)
                                _ = v32
                            } else {
                                for fi33 := 0; fi33 < 8; fi33++ {
                                    v4 = v3
                                    _ = fi33
                                }
                            }
                            var v34 string = "\n\t"
                            v34 = ""
                            _ = v34
                        }
                    case "hello":
                        for fi35 := 0; fi35 < 18; fi35++ {
                            var v36 uint16 = 1
                            _ = fi35
                            _ = v36
                        }
                    default:
                        v2 = append(v2, "")
                    }
                }
            } else {
                if false {
                    fmt.Println(v23, fi22)
                    var v37 float32 = float32(float64(-2.25))
                    _ = v37
                } else {
                    v38 := GenS3{F0: 1, F1: 65535}
                    var v39 uint64 = (v18 & 1)
                    _ = v38
                    _ = v39
                }
                switch -2147483648 {
                case 128:
                    switch "" {
                    case "a":
                        v40 := GenS4{F0: -32768}
                        _ = v40
                    case "hello":
                        _ = len(v1)
                    case "foo bar":
                        switch "foo bar" {
                        case "hello":
                            var v41 int32 = -2147483648
                            _ = v41
                        case "foo bar":
                            for fi42 := 0; fi42 < 15; fi42++ {
                                v43 := GenS3{F0: 1000, F1: 2147483648}
                                _ = fi42
                                _ = v43
                            }
                        }
                    }
                }
            }
            _ = fi22
            _ = v23
        }
    }
    var v44 int32 = 42
    v45 := GenS1{F0: ("" + ("a" + "hello")), F1: 753, F2: 201}
    v46 := v1[("\n\t" + "a")]
    v47 := v1[("\n\t" + "\n\t")]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
