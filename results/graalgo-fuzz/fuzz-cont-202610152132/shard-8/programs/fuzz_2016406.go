package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int; F1 int64; F2 bool }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32 }
func genFn0() int {
    for fi4 := 0; fi4 < 3; fi4++ {
        switch "foo bar" {
        case "a":
            for fi5 := 0; fi5 < 2; fi5++ {
                v6 := []int{}
                switch "foo bar" {
                case "foo bar":
                    if false {
                        var v7 bool = (false && (243 < 42))
                        v7 = (!(1 != 1))
                        switch (("hello" + "\n\t") + "") {
                        case "\n\t":
                            v6 = append(v6, 128)
                        }
                        _ = v7
                    }
                case "":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case "\n\t":
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        if len(v6) > 0 {
                            v6[0] = (fi4 / ((fi5) | 1))
                        }
                    default:
                        switch (("foo bar" + "a") + "a") {
                        case "foo bar":
                            var v8 string = "\n\t"
                            _ = v8
                        case "a":
                            _ = len(v6)
                        default:
                            fmt.Println(fi5, fi5)
                        }
                    }
                }
                fmt.Println(fi4, fi4)
                v9 := GenS1{F0: 1, F1: (!false)}
                _ = fi5
                _ = v6
                _ = v9
            }
        case "foo bar":
            fmt.Println(fi4)
        case "":
            v10 := []int{fi4}
            _ = v10
        }
        _ = fi4
    }
    var v11 int64 = 42
    v12 := []int{-1, 128, -1}
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint32 = 2147483647
        v11 += v11
        v15 := GenS1{F0: 0, F1: true}
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v11
    _ = v12
    return 33
}

func main() {
    var v1 int8 = -1
    v2 := []string{"a", ("" + "hello"), "\n\t", (("\n\t" + "hello") + "foo bar")}
    var v3 bool = false
    v3 = v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = true
    v16 := GenS2{F0: 0, F1: 268}
    if (581 >= 255) {
        v17 := GenS1{F0: 826, F1: (true && v3)}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 3; fi18++ {
            if ((((1e10 + float32((1e10 * -2.25))) / 3.14159) + -2.25) < 3.14159) {
                if (float32((1.5 - (1.5 + 1e10))) == 0.1) {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        var v20 uint64 = 0
                        v20 = 42
                        var v21 float32 = 1.5
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    if v3 {
                        var v22 float32 = 0.0
                        _ = v22
                    } else {
                        if ((!v3) && v3) {
                            if true {
                                v23 := GenS3{F0: fi18, F1: -694, F2: v3}
                                var v24 int16 = 232
                                v3 = ((1 != 127) || false)
                                _ = v23
                                _ = v24
                            }
                        } else {
                            v1 = (v1 - v1)
                        }
                    }
                    var v25 uint64 = 18446744073709551615
                    _ = v25
                } else {
                    v3 = false
                }
                fmt.Println(v1, v1)
            } else {
                v3 = v3
                if v3 {
                    v26 := GenS1{F0: 1, F1: v3}
                    if v3 {
                        v27 := GenS2{F0: 4294967295, F1: 2}
                        _ = v27
                    }
                    _ = v26
                } else {
                    if (((v3 && (18446744073709551615 < 2)) || v3) && v3) {
                        var v28 int32 = -1
                        _ = v28
                    }
                }
                v29 := GenS3{F0: genFn0(), F1: 154, F2: (false && true)}
                _ = v29
            }
            for fi30 := 0; fi30 < 8; fi30++ {
                v1 = (v1 & v1)
                v1 = v1
                v3 = false
                _ = len(v2)
                var v31 int64 = 9223372036854775807
                if len(v2) > 0 {
                    v2[0] = "a"
                }
                _ = fi30
                _ = v31
            }
            _ = fi18
        }
        var v32 int = 32768
        v2 = append(v2, "hello")
        _ = v32
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        if v3 {
            for fi34 := 0; fi34 < 15; fi34++ {
                if (!v3) {
                    fmt.Println(v1)
                }
                _ = fi34
            }
        } else {
            var v35 uint8 = 42
            v1 += v1
            _ = v35
        }
        v1 -= v1
        _ = fi33
    }
    var v36 int64 = -233
    var v37 float64 = -2.25
    _ = len(v2)
    var v38 int64 = (v36 | v36)
    if ((351 < 42) && (!false)) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v3 = (((-2.25 - (0.1 * (3.14159 - 3.14159))) / (0.0 - ((1.5 - 0.0) * 1.5))) < -2.25)
        v37 -= 0.1
        v39 := GenS3{F0: -977, F1: 759, F2: (1 >= 1)}
        _ = v39
    }
    fmt.Println(v1, v37)
    v40 := GenS4{F0: 2147483647, F1: 0}
    var v41 uint = 682
    switch "\n\t" {
    case "\n\t":
        v37 = v37
    case "hello":
        if (true || v3) {
            if (1 < 2147483647) {
                v36 -= v38
            }
            if v3 {
                var v42 float32 = (0.0 / 1.0)
                _ = v42
            }
        } else {
            if true {
                fmt.Println(v36)
                fmt.Println(v38, v38)
                v43 := GenS5{F0: 32767, F1: -996}
                _ = v43
            } else {
                fmt.Println(v38)
                var v44 float64 = float64(-2.25)
                if false {
                    switch 32767 {
                    case 127:
                        v45 := GenS1{F0: 1, F1: (0 == 128)}
                        _ = v45
                    case -1:
                        v41 += 0
                    default:
                        switch 1 {
                        case 4294967295:
                            v46 := GenS6{F0: 4294967295}
                            _ = v46
                        case -447:
                            var v47 int64 = v36
                            _ = v47
                        }
                    }
                    v48 := GenS6{F0: 4294967295}
                    switch "" {
                    case "a":
                        v3 = true
                    case "":
                        var v49 int16 = -1
                        _ = v49
                    case "hello":
                        v50 := GenS7{F0: 2147483647}
                        _ = v50
                    }
                    _ = v48
                } else {
                    var v51 uint8 = 157
                    var v52 int8 = (v1 * (v1 * (v1 >> ((v41) & 63))))
                    switch 758 {
                    case 4294967295:
                        var v53 int32 = 42
                        _ = v53
                    default:
                        v54 := GenS2{F0: 0, F1: 32767}
                        _ = v54
                    }
                    _ = v51
                    _ = v52
                }
                _ = v44
            }
            for fi55 := 0; fi55 < 10; fi55++ {
                v38 = v38
                var v56 int = fi55
                _ = fi55
                _ = v56
            }
        }
    case "a":
        var v57 int32 = 42
        _ = v57
    }
    var v58 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v41:", v41)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v40
    _ = v41
    _ = v58
}
