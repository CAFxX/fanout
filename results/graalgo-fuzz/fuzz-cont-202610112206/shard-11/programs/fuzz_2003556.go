package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int32; F1 int16 }
func genFn0() string {
    v6 := []int{1}
    var v7 uint64 = 4294967295
    _ = len(v6)
    v8 := GenS1{F0: v7}
    v9 := GenS1{F0: (v7 & v7)}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return ("a" + "")
}
func genFn1(p0 uint64, p1 uint8, p2 int16) int32 {
    var v10 uint16 = 0
    if (!true) {
        v10 -= 2
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        p0 = p0
        var v12 uint32 = 1
        for fi13 := 0; fi13 < 15; fi13++ {
            switch ("a" + (("" + "a") + "\n\t")) {
            case "":
                v10 += v10
            }
            v12 -= v12
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}

func main() {
    var v1 string = ("hello" + "a")
    var v2 uint = 571
    var v3 uint8 = 255
    v4 := make(map[string]int)
    v5 := []string{}
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 int16 = 922
            if true {
                v1 = ((v1 + v1) + v1)
            }
            v3 -= 255
            _ = fi15
            _ = v16
        }
        var v17 uint64 = 18446744073709551615
        _ = fi14
        _ = v17
    }
    v3 -= (^(v3))
    for fi18 := 0; fi18 < 10; fi18++ {
        if true {
            v19 := GenS1{F0: 2}
            v20 := GenS2{F0: 2, F1: 2, F2: 32768}
            for fi21 := 0; fi21 < 12; fi21++ {
                switch v1 {
                case "a":
                    var v22 int16 = 1
                    _ = v22
                case "foo bar":
                    if (v2 >= v2) {
                        v2 += (v2 % ((v2) | 1))
                        v1 = v1
                    } else {
                        if (2 < 206) {
                            var v23 string = (v1 + v1)
                            delete(v4, v1)
                            fmt.Println(fi21)
                            _ = v23
                        } else {
                            if true {
                                if true {
                                    v24 := GenS1{F0: 4294967295}
                                    switch (v1 + v1) {
                                    case "a":
                                        var v25 uint8 = v3
                                        _ = v25
                                    case "\n\t":
                                        v3 += v3
                                    }
                                    _ = v24
                                } else {
                                    var v26 uint8 = (v3 / ((65) | 1))
                                    var v27 uint32 = 0
                                    v3 = (^(v26))
                                    _ = v26
                                    _ = v27
                                }
                                var v28 float32 = 0.1
                                fmt.Println(v1)
                                var v29 int = 989
                                _ = v28
                                _ = v29
                            }
                        }
                    }
                case "hello":
                    if true {
                        v2 += v2
                        v1 = v1
                        if (2147483647 <= -104) {
                            _ = len(v5)
                            _ = len(v5)
                            v3 -= 1
                        } else {
                            v30 := GenS3{F0: (^(v2)), F1: 1000}
                            _ = v30
                        }
                    } else {
                        switch fi18 {
                        case 0:
                            v3 = 143
                        default:
                            _ = len(v4)
                            var v31 uint32 = 255
                            v32 := GenS1{F0: 0}
                            switch ("foo bar" + v1) {
                            case "\n\t":
                                fmt.Println(fi21, v1)
                            }
                            _ = v31
                            _ = v32
                        }
                        var v33 float64 = 1.5
                        fmt.Println(v3, v33)
                        v34 := GenS2{F0: 32767, F1: 42, F2: fi21}
                        var v35 uint32 = uint32(v2)
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    v5 = append(v5, v1)
                    if true {
                        v3 = v3
                    }
                }
                fmt.Println(fi21, v1, fi18)
                if false {
                    fmt.Println(v1, v2)
                    var v36 int64 = 42
                    v36 -= -772
                    v3 -= (v3 & 236)
                    _ = v36
                } else {
                    var v37 string = ("a" + (v1 + (v1 + "foo bar")))
                    _ = v37
                }
                _ = fi21
            }
            _ = v19
            _ = v20
        } else {
            if true {
                for fi38 := 0; fi38 < 2; fi38++ {
                    var v39 uint64 = 18446744073709551615
                    if true {
                        v3 += (v3 + 127)
                        var v40 float64 = 0.1
                        var v41 uint16 = 255
                        if (true && (1 != -1)) {
                            v5 = append(v5, "\n\t")
                            _ = len(v4)
                            v42 := GenS3{F0: (v2 - v2), F1: 1}
                            if (!true) {
                                v43 := GenS3{F0: 255, F1: 798}
                                v3 -= v3
                                _ = v43
                            } else {
                                v1 = (v1 + "hello")
                            }
                            _ = v42
                        } else {
                            var v44 int32 = -2147483648
                            _ = v44
                        }
                        v1 = ((v1 + "") + v1)
                        fmt.Println(v3)
                        _ = v40
                        _ = v41
                    }
                    _ = fi38
                    _ = v39
                }
            }
            v1 = v1
        }
        _ = fi18
    }
    v1 = v1
    var v45 string = v1
    for fi46 := 0; fi46 < 1; fi46++ {
        fmt.Println(v2, v1)
        var v47 int32 = genFn1(2, (v3 * (v3 % (((^(v3))) | 1))), 1000)
        _ = fi46
        _ = v47
    }
    var v48 int64 = 42
    switch "a" {
    case "hello":
        switch v45 {
        case "a":
            if false {
                var v49 float64 = (3.14159 - (((1e10 / 1.0) - -2.25) - float64(0.1)))
                var v50 float64 = v49
                _ = v49
                _ = v50
            }
        default:
            var v51 uint8 = 19
            _ = v51
        }
    case "":
        fmt.Println(v1, v1, v1)
    case "a":
        var v52 uint64 = 4294967295
        _ = v52
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if (!true) {
        if true {
            v53 := v4[""]
            _ = v53
        } else {
            var v54 uint16 = 490
            _ = v54
        }
        v55 := GenS4{F0: 42, F1: 584}
        _ = v55
    }
    var v56 float32 = (1e10 - 1.5)
    switch v1 {
    case "foo bar":
        var v57 int = -64
        _ = v57
    case "\n\t":
        v58 := GenS1{F0: 18446744073709551615}
        _ = v58
    case "hello":
        delete(v4, v1)
    }
    var v59 int8 = 0
    for fi60 := 0; fi60 < 9; fi60++ {
        var v61 float32 = v56
        _ = fi60
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v5
    _ = v56
    _ = v59
}
