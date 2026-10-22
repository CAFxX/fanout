package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 string; F2 int8 }
type GenS3 struct { F0 uint16; F1 int16; F2 int8 }
type GenS4 struct { F0 int; F1 uint64; F2 int16 }
type GenS5 struct { F0 int; F1 string; F2 uint64 }

func main() {
    var v1 uint8 = 40
    var v2 string = (("hello" + ("\n\t" + "hello")) + (("a" + "a") + ""))
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    if (!true) {
        v2 = (v2 + "a")
        if ((!true) && (0 == 42)) {
            if ((^(v1)) < v1) {
                if true {
                    v5 := GenS1{F0: 585}
                    var v6 uint16 = 1
                    _ = v5
                    _ = v6
                }
                v3[44] = "\n\t"
            } else {
                if (false || (true && false)) {
                    var v7 uint = 828
                    var v8 string = ((v2 + ("hello" + (v2 + v2))) + v2)
                    var v9 int = 1
                    _ = v7
                    _ = v8
                    _ = v9
                } else {
                    if ((float32(1.5) >= 3.14159) && false) {
                        if false {
                            for fi10 := 0; fi10 < 11; fi10++ {
                                for fi11 := 0; fi11 < 7; fi11++ {
                                    _ = len(v4)
                                    switch (("" + v2) + v2) {
                                    case "hello":
                                        if true {
                                            v12 := GenS1{F0: 445}
                                            _ = len(v4)
                                            _ = v12
                                        } else {
                                            var v13 uint = 919
                                            var v14 float64 = (float64(-2.25) / 1.0)
                                            var v15 int = (-(fi11))
                                            _ = v13
                                            _ = v14
                                            _ = v15
                                        }
                                    case "foo bar":
                                        if (!(2 >= 42)) {
                                            v1 += 1
                                            v3[fi11] = ((((v2 + (v2 + (v2 + v2))) + "hello") + ("" + "")) + "hello")
                                            if false {
                                                var v16 uint = 621
                                                fmt.Println(v1, fi11, v2)
                                                v1 = (v1 >> ((v16) & 63))
                                                _ = v16
                                            } else {
                                                var v17 uint16 = 0
                                                v1 = v1
                                                _ = v17
                                            }
                                        } else {
                                            v1 = 119
                                            var v18 uint = 42
                                            _ = v18
                                        }
                                    case "\n\t":
                                        fmt.Println(fi10, v2)
                                    default:
                                        var v19 uint16 = 2
                                        _ = v19
                                    }
                                    _ = fi11
                                }
                                _ = fi10
                            }
                            var v20 bool = true
                            _ = v20
                        }
                    }
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v1 = v1
                        _ = fi21
                    }
                }
                var v22 bool = false
                _ = v22
            }
            if (42 > (v1 / ((42) | 1))) {
                v2 = v2
                var v23 int = 2147483647
                _ = v23
            } else {
                v3[-221] = ("\n\t" + v2)
            }
            v2 = v2
        } else {
            var v24 uint8 = 255
            var v25 int64 = 0
            v26 := GenS1{F0: uint16(v24)}
            _ = v24
            _ = v25
            _ = v26
        }
    }
    _ = len(v3)
    v1 += (v1 & 1)
    v27 := GenS1{F0: 659}
    v3[0] = "foo bar"
    v1 = v1
    v1 = v1
    fmt.Println(v2)
    if false {
        var v28 float32 = float32((1e10 + (0.0 * 3.14159)))
        _ = v28
    } else {
        v29 := []int{9223372036854775807, -2147483648, 127}
        v30 := GenS1{F0: 53}
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        if false {
            for fi32 := 0; fi32 < 10; fi32++ {
                if true {
                    v33 := GenS1{F0: uint16(v1)}
                    _ = v33
                } else {
                    switch (("hello" + (v2 + v2)) + "foo bar") {
                    case "foo bar":
                        v34 := GenS2{F0: 1, F1: "foo bar", F2: 100}
                        _ = v34
                    case "a":
                        var v35 uint8 = 0
                        _ = v35
                    }
                    switch (fi32 >> ((418) & 63)) {
                    case -672:
                        v36 := GenS3{F0: 772, F1: 42, F2: -128}
                        _ = v36
                    case -954:
                        var v37 uint32 = 14
                        _ = v37
                    case -9223372036854775808:
                        v38 := GenS4{F0: -2147483648, F1: 643, F2: -32768}
                        _ = v38
                    }
                    v2 = (v2 + ("a" + ("" + v2)))
                }
                if (true && (true || false)) {
                    v39 := []int{(fi31 / ((fi31) | 1)), (fi32 * -979)}
                    v40 := v4[792]
                    _ = v39
                    _ = v40
                } else {
                    fmt.Println(fi31, v2)
                    delete(v3, fi32)
                }
                _ = fi32
            }
            if ((1e10 * ((0.1 / -2.25) * (3.14159 / (0.1 + 1.5)))) != 1e10) {
                fmt.Println(v1, v1)
                for fi41 := 0; fi41 < 16; fi41++ {
                    _ = len(v4)
                    var v42 int = 736
                    if false {
                        var v43 uint32 = 2147483647
                        var v44 uint = 1
                        v43 -= (v43 & v43)
                        fmt.Println(v44)
                        _ = v43
                        _ = v44
                    } else {
                        v45 := GenS5{F0: v42, F1: "hello", F2: 42}
                        v46 := GenS3{F0: 255, F1: -714, F2: 86}
                        v47 := GenS1{F0: 42}
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    _ = fi41
                    _ = v42
                }
                v48 := []int{}
                _ = v48
            } else {
                var v49 uint64 = 432
                v1 -= 11
                v49 += (v49 % ((v49) | 1))
                _ = v49
            }
        } else {
            var v50 uint8 = 1
            for fi51 := 0; fi51 < 19; fi51++ {
                var v52 float32 = (float32(float64(1e10)) + (0.0 - 3.14159))
                var v53 int16 = -32768
                _ = fi51
                _ = v52
                _ = v53
            }
            switch v2 {
            case "foo bar":
                fmt.Println(v50, fi31)
            default:
                var v54 string = v2
                _ = v54
            }
            _ = v50
        }
        _ = fi31
    }
    var v55 float64 = -2.25
    v2 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v55
}
