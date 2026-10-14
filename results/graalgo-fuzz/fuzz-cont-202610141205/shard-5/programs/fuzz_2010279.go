package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 int8 }
func genFn0() int {
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int8 = 127
        _ = fi7
        _ = v8
    }
    switch "hello" {
    case "foo bar":
        switch -948 {
        case -2147483648:
            var v9 uint16 = 591
            _ = v9
        default:
            var v10 int32 = 964
            _ = v10
        }
    case "hello":
        v11 := []int{1, 9223372036854775807, -343}
        _ = v11
    default:
        var v12 uint = 255
        _ = v12
    }
    return -49
}
func genFn1() bool {
    var v13 int16 = 0
    fmt.Println(v13)
    _ = v13
    return (true || false)
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 1
    var v3 float64 = 0.1
    var v4 uint32 = 1
    var v5 int8 = 0
    var v6 int64 = -674
    v14 := GenS1{F0: 282}
    var v15 int64 = v6
    switch "foo bar" {
    case "\n\t":
        if false {
            v16 := GenS1{F0: 1}
            v17 := GenS1{F0: -32768}
            if true {
                if ((v2 * (v2 % ((v2) | 1))) < v2) {
                    switch "foo bar" {
                    case "foo bar":
                        v5 += 100
                    default:
                        for fi18 := 0; fi18 < 5; fi18++ {
                            if true {
                                for fi19 := 0; fi19 < 18; fi19++ {
                                    v5 -= (v5 % (((-(v5))) | 1))
                                    _ = fi19
                                }
                            } else {
                                var v20 int = fi18
                                _ = v20
                            }
                            v21 := GenS2{F0: 32767, F1: 32767}
                            _ = fi18
                            _ = v21
                        }
                    }
                }
                v22 := []int{32768}
                _ = v22
            }
            _ = v16
            _ = v17
        }
    default:
        v23 := []int{}
        _ = v23
    }
    fmt.Println(v2, v3)
    var v24 uint16 = 0
    v25 := GenS3{F0: -1}
    v1[("" + ("a" + "a"))] = int(v24)
    v26 := GenS2{F0: (^(v24)), F1: 841}
    var v27 bool = true
    delete(v1, "a")
    fmt.Println(v24, v24, v4)
    if v27 {
        switch "a" {
        case "a":
            var v28 float32 = (1e10 - (3.14159 - 0.1))
            _ = v28
        case "":
            fmt.Println(v15, v3)
        }
    }
    if false {
        for fi29 := 0; fi29 < 15; fi29++ {
            v24 -= (v24 << ((4294967295) & 63))
            _ = fi29
        }
        if (9223372036854775807 <= genFn0()) {
            if (1.5 >= 0.0) {
                v4 = (v4 << ((65535) & 63))
            }
            var v30 int32 = -2147483648
            v31 := GenS3{F0: 42}
            _ = v30
            _ = v31
        } else {
            var v32 int8 = v5
            v32 = v32
            v33 := []int{2147483647}
            _ = v32
            _ = v33
        }
        if false {
            if v27 {
                v34 := []int{}
                var v35 string = "hello"
                _ = v34
                _ = v35
            } else {
                v2 = 42
            }
            switch ("a" + (("foo bar" + "hello") + "a")) {
            case "\n\t":
                var v36 int64 = (v15 & v6)
                _ = v36
            case "":
                v37 := []int{32767}
                _ = v37
            case "hello":
                for fi38 := 0; fi38 < 5; fi38++ {
                    switch fi38 {
                    case 32767:
                        fmt.Println(fi38, v27)
                    }
                    fmt.Println(v5)
                    switch 9223372036854775807 {
                    case 933:
                        v2 += 4294967295
                    case 42:
                        for fi39 := 0; fi39 < 16; fi39++ {
                            switch fi38 {
                            case 42:
                                v15 += (v6 & (v6 / ((-9223372036854775808) | 1)))
                            case -261:
                                fmt.Println(v4, fi38)
                            case 179:
                                switch (("foo bar" + ("" + "hello")) + "foo bar") {
                                case "":
                                    if v27 {
                                        switch (fi38 + (fi38 - 4294967295)) {
                                        case 42:
                                            v6 += -1
                                            v40 := v1[("a" + "\n\t")]
                                            switch (fi39 / ((fi39) | 1)) {
                                            case -2147483648:
                                                var v41 uint16 = (v24 % ((v24) | 1))
                                                _ = v41
                                            }
                                            _ = v40
                                        default:
                                            fmt.Println(fi39, fi39, v27)
                                        }
                                    }
                                case "\n\t":
                                    var v42 string = "foo bar"
                                    _ = v42
                                case "hello":
                                    switch ("" + "a") {
                                    case "":
                                        fmt.Println(v5, v3, v3)
                                    default:
                                        v43 := GenS4{F0: "hello", F1: (v5 * v5)}
                                        _ = v43
                                    }
                                }
                            }
                            if (106 == 18446744073709551615) {
                                var v44 uint32 = (^(v4))
                                switch "hello" {
                                case "":
                                    v45 := []int{fi38}
                                    _ = v45
                                case "foo bar":
                                    var v46 uint32 = v44
                                    _ = v46
                                default:
                                    v47 := GenS4{F0: "\n\t", F1: v5}
                                    v48 := GenS3{F0: 1}
                                    if v27 {
                                        if false {
                                            fmt.Println(v24, v24, v24)
                                        }
                                    }
                                    _ = v47
                                    _ = v48
                                }
                                var v49 int32 = -2147483648
                                v1["foo bar"] = (fi39 / ((int(v5)) | 1))
                                _ = v44
                                _ = v49
                            } else {
                                var v50 uint8 = 127
                                fmt.Println(v5, v5, v3)
                                var v51 float32 = 0.1
                                v52 := []int{}
                                switch fi39 {
                                case 127:
                                    v50 = (^(v50))
                                case 32767:
                                    var v53 int64 = (v6 ^ v15)
                                    _ = v53
                                case 4294967295:
                                    if v27 {
                                        v27 = v27
                                    }
                                default:
                                    var v54 int = (fi39 + genFn0())
                                    _ = v54
                                }
                                v51 -= 0.1
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                            fmt.Println(v24, v5)
                            _ = fi39
                        }
                    default:
                        for fi55 := 0; fi55 < 7; fi55++ {
                            var v56 int = (fi55 + genFn0())
                            _ = fi55
                            _ = v56
                        }
                    }
                    _ = fi38
                }
            default:
                v24 = v24
            }
        } else {
            var v57 int = -9223372036854775808
            if v27 {
                if (643 >= (^(v4))) {
                    fmt.Println(v2)
                    var v58 uint8 = 0
                    v2 -= (v2 + v2)
                    _ = v58
                } else {
                    v6 -= (v15 - v15)
                }
            }
            v4 = v4
            _ = v57
        }
    }
    var v59 float32 = (0.1 / float32(-2.25))
    v5 = v5
    if true {
        v60 := GenS1{F0: 0}
        v6 = v15
        var v61 uint16 = 409
        _ = v60
        _ = v61
    } else {
        var v62 int32 = 1
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
