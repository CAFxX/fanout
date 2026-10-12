package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint32; F1 bool; F2 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 uint32; F2 string }
func genFn0(p0 int64, p1 uint8, p2 int16) string {
    var v6 uint16 = 65535
    var v7 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return ("foo bar" + ("foo bar" + (("\n\t" + ("a" + "hello")) + "\n\t")))
}
func genFn1(p0 uint32) int16 {
    var v8 int32 = 1
    fmt.Println(p0, p0)
    _ = p0
    _ = v8
    return 327
}

func main() {
    var v1 int32 = 739
    var v2 int8 = -1
    var v3 int = 127
    v4 := make(map[string]bool)
    var v5 int64 = 1
    switch "foo bar" {
    case "":
        v9 := GenS1{F0: genFn1(42), F1: -203, F2: 247}
        _ = v9
    case "foo bar":
        v5 = (v5 % ((v5) | 1))
    }
    switch (v3 / (((v3 * 42)) | 1)) {
    case 9223372036854775807:
        if (false && false) {
            v3 = v3
        } else {
            for fi10 := 0; fi10 < 14; fi10++ {
                var v11 uint = 231
                _ = fi10
                _ = v11
            }
            v12 := GenS1{F0: -1, F1: 127, F2: 0}
            for fi13 := 0; fi13 < 8; fi13++ {
                if true {
                    if false {
                        if true {
                            if false {
                                for fi14 := 0; fi14 < 18; fi14++ {
                                    _ = len(v4)
                                    if true {
                                        var v15 uint32 = 0
                                        v16 := GenS1{F0: -1, F1: fi13, F2: 507}
                                        v17 := GenS1{F0: 0, F1: (^(fi13)), F2: 2}
                                        switch (v3 >> ((0) & 63)) {
                                        case -1:
                                            v5 = -201
                                        case 1:
                                            if (int64(v15) != (v5 * 9223372036854775807)) {
                                                v18 := GenS2{F0: 2}
                                                var v19 int = (v3 - int(v1))
                                                fmt.Println(v3, v19)
                                                v20 := []int{0}
                                                v3 -= (^(fi13))
                                                _ = v18
                                                _ = v19
                                                _ = v20
                                            } else {
                                                v15 += v15
                                                if false {
                                                    v21 := []int{562}
                                                    _ = v21
                                                } else {
                                                    var v22 int64 = v5
                                                    var v23 int8 = -1
                                                    _ = v22
                                                    _ = v23
                                                }
                                                var v24 float64 = ((0.1 / ((float64(3.14159) * ((0.0 * 0.1) - 3.14159)) - 1.5)) + (0.1 * -2.25))
                                                if (!((false || true) && true)) {
                                                    switch (^(fi13)) {
                                                    case -819:
                                                        v25 := GenS1{F0: 1000, F1: (fi13 >> ((1) & 63)), F2: (v15 ^ (v15 >> ((42) & 63)))}
                                                        _ = v25
                                                    case 32768:
                                                        v26 := []int{(fi13 / ((v3) | 1)), (fi13 / (((^(fi13))) | 1)), 32767, -9223372036854775808}
                                                        _ = v26
                                                    case 9223372036854775807:
                                                        v27 := GenS3{F0: 1, F1: ((("a" + ("hello" + ("" + ""))) + ("" + "foo bar")) + ("a" + (("hello" + "") + ("a" + "\n\t"))))}
                                                        _ = v27
                                                    }
                                                    v28 := GenS1{F0: -794, F1: v3, F2: v15}
                                                    v24 = 0.1
                                                    _ = v28
                                                } else {
                                                    v29 := GenS2{F0: 255}
                                                    _ = v29
                                                }
                                                if ("" > "") {
                                                    switch (("a" + "\n\t") + (("\n\t" + ("hello" + "")) + "a")) {
                                                    case "foo bar":
                                                        delete(v4, "")
                                                    case "a":
                                                        v15 = v15
                                                    }
                                                    v30 := []int{-1, (fi13 / (((-(v3))) | 1)), 1, (fi14 & -1)}
                                                    if true {
                                                        if false {
                                                            v15 = (v15 * (^(v15)))
                                                        } else {
                                                            var v31 uint = 970
                                                            if (!true) {
                                                                fmt.Println(v3, v5, v2)
                                                            }
                                                            if true {
                                                                v32 := GenS3{F0: 255, F1: "\n\t"}
                                                                _ = v32
                                                            }
                                                            _ = v31
                                                        }
                                                        fmt.Println(v2)
                                                    } else {
                                                        delete(v4, "")
                                                        var v33 uint16 = 65535
                                                        _ = v33
                                                    }
                                                    _ = v30
                                                }
                                                _ = v24
                                            }
                                        case 32768:
                                            if true {
                                                var v34 int16 = -32768
                                                var v35 int8 = 42
                                                _ = v34
                                                _ = v35
                                            } else {
                                                fmt.Println(fi14)
                                            }
                                        default:
                                            switch 9223372036854775807 {
                                            case -95:
                                                v2 -= v2
                                            case 4294967295:
                                                v15 += (v15 % ((v15) | 1))
                                            }
                                        }
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                    }
                                    _ = fi14
                                }
                            }
                            var v36 string = "hello"
                            _ = v36
                        }
                        fmt.Println(v2)
                    } else {
                        fmt.Println(fi13, v1)
                        fmt.Println(v5)
                    }
                    for fi37 := 0; fi37 < 10; fi37++ {
                        v38 := []int{}
                        var v39 uint32 = 919
                        var v40 int = -1
                        _ = fi37
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                    var v41 bool = true
                    _ = v41
                }
                _ = fi13
            }
            _ = v12
        }
    case 32768:
        v5 = v5
    default:
        v42 := GenS2{F0: 1}
        _ = v42
    }
    v43 := GenS4{F0: 2147483647, F1: (42 < v1), F2: 352}
    var v44 string = ""
    var v45 int64 = -1
    var v46 bool = false
    for fi47 := 0; fi47 < 14; fi47++ {
        v1 = v1
        v48 := v4["a"]
        var v49 int32 = 0
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println(v46, v44)
    fmt.Println(v3, v5, v5)
    for fi50 := 0; fi50 < 17; fi50++ {
        for fi51 := 0; fi51 < 16; fi51++ {
            v52 := GenS5{F0: 2147483648}
            _ = fi51
            _ = v52
        }
        _ = fi50
    }
    switch (v44 + "foo bar") {
    case "\n\t":
        v1 = v1
    case "a":
        v53 := GenS6{F0: v46, F1: 65535, F2: "foo bar"}
        _ = v53
    case "":
        var v54 uint8 = 179
        _ = v54
    default:
        for fi55 := 0; fi55 < 20; fi55++ {
            var v56 string = "hello"
            var v57 float64 = 0.0
            v1 = v1
            _ = fi55
            _ = v56
            _ = v57
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
