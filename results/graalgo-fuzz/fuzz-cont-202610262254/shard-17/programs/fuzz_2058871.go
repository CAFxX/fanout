package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint8; F2 int8 }
type GenS5 struct { F0 string; F1 int16 }

func main() {
    var v1 float64 = ((0.1 - float64(0.0)) / ((3.14159 - 1e10) - 0.0))
    var v2 int64 = 42
    var v3 float64 = 0.1
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]string)
    var v6 float32 = (-2.25 * 1.5)
    v3 -= 0.0
    switch ("foo bar" + "") {
    case "":
        v7 := GenS1{F0: 0}
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        _ = len(v5)
        _ = fi8
    }
    switch 32767 {
    case 935:
        if (42 < 1) {
            var v9 uint8 = 2
            v10 := GenS1{F0: -94}
            _ = v10
            _ = v9
        }
    case 792:
        v2 = v2
    case 9223372036854775807:
        v11 := []int{}
        _ = v11
    }
    switch 9223372036854775807 {
    case 127:
        v3 += 3.14159
    }
    delete(v5, 48)
    v12 := []int{42, 32767, 32767, 42}
    if false {
        delete(v4, 65535)
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 uint = 0
            switch ("hello" + "hello") {
            case "foo bar":
                if true {
                    v15 := GenS2{F0: false}
                    for fi16 := 0; fi16 < 18; fi16++ {
                        switch ("" + ("foo bar" + "a")) {
                        case "\n\t":
                            v14 = (v14 << ((1) & 63))
                        case "":
                            if false {
                                var v17 int = (fi16 / (((fi16 & fi13)) | 1))
                                v18 := GenS3{F0: 128}
                                var v19 float32 = (0.1 + ((v6 + 0.1) - 0.0))
                                v20 := GenS1{F0: 39}
                                _ = v17
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                        case "hello":
                            var v21 uint = (v14 << ((v14) & 63))
                            _ = v21
                        }
                        if (!(((0 >= (v14 / ((0) | 1))) && true) && (32767 < 255))) {
                            delete(v4, 280)
                        }
                        fmt.Println(v3, fi16, fi16)
                        _ = fi16
                    }
                    v1 += v1
                    _ = v15
                } else {
                    var v22 int = 9223372036854775807
                    if (0 <= 42) {
                        fmt.Println(v6, v1, v1)
                        var v23 bool = (true || (-1 == -1))
                        _ = v23
                    } else {
                        var v24 float64 = (v1 * v1)
                        switch "a" {
                        case "hello":
                            for fi25 := 0; fi25 < 2; fi25++ {
                                switch fi13 {
                                case -2147483648:
                                    v26 := GenS2{F0: true}
                                    _ = v26
                                case -34:
                                    if false {
                                        v27 := GenS2{F0: false}
                                        switch (fi13 % ((fi25) | 1)) {
                                        case 0:
                                            switch "foo bar" {
                                            case "":
                                                var v28 int8 = 100
                                                _ = v28
                                            case "a":
                                                fmt.Println(fi13, v2, v24)
                                            }
                                        }
                                        var v29 int16 = -32768
                                        _ = v27
                                        _ = v29
                                    }
                                case 32768:
                                    switch v22 {
                                    case 42:
                                        var v30 uint32 = 65535
                                        if true {
                                            v12 = append(v12, (v22 ^ 128))
                                            v24 -= 1e10
                                        }
                                        _ = v30
                                    default:
                                        v31 := GenS1{F0: 127}
                                        _ = v31
                                    }
                                }
                                var v32 bool = true
                                _ = fi25
                                _ = v32
                            }
                        }
                        v22 += v22
                        _ = v24
                    }
                    _ = v22
                }
            }
            if true {
                v14 = (v14 ^ 65535)
                delete(v4, 373)
            }
            _ = fi13
            _ = v14
        }
        if false {
            var v33 uint64 = 982
            v1 += v3
            _ = v33
        }
        v2 += v2
    }
    var v34 int32 = -1
    if (2 != 69) {
        v35 := GenS4{F0: 1, F1: 255, F2: 15}
        _ = v35
    } else {
        if (v6 > (float32((0.1 / 1e10)) * (v6 + v6))) {
            if false {
                v1 = (((((v1 + float64((((0.0 + v6) + 1.5) * v6))) * 1.5) - v1) + 0.0) + (v1 + (v1 / v3)))
            } else {
                switch 4294967295 {
                case 2147483647:
                    v2 = -581
                }
                v34 = -2147483648
                var v36 int32 = (v34 >> ((162) & 63))
                _ = v36
            }
            for fi37 := 0; fi37 < 8; fi37++ {
                for fi38 := 0; fi38 < 4; fi38++ {
                    v6 -= v6
                    v2 = -1
                    if ((1 >= 127) && false) {
                        delete(v4, 0)
                    } else {
                        if false {
                            _ = len(v12)
                            v2 += -9223372036854775808
                            var v39 uint16 = 2
                            _ = v39
                        } else {
                            var v40 uint = 51
                            v41 := GenS1{F0: -128}
                            v34 += (v34 & v34)
                            _ = v40
                            _ = v41
                        }
                        if ((v6 / v6) <= -2.25) {
                            v2 = (v2 << ((1) & 63))
                        } else {
                            fmt.Println(fi38, v2, v3)
                        }
                    }
                    v1 = float64((1.5 * v6))
                    _ = fi38
                }
                _ = fi37
            }
        }
        v42 := GenS5{F0: "a", F1: 0}
        _ = v42
    }
    v34 = (v34 >> ((42) & 63))
    switch "foo bar" {
    case "a":
        var v43 int64 = 42
        _ = v43
    case "\n\t":
        if ((737 <= 2) || (2147483647 <= 4294967295)) {
            var v44 uint8 = 128
            var v45 int8 = -97
            v4[42] = (!true)
            _ = v44
            _ = v45
        }
    }
    var v46 int16 = 1000
    if ((4294967295 > 181) && (false && false)) {
        if false {
            var v47 string = "\n\t"
            if len(v12) > 0 {
                v12[0] = 1
            }
            _ = v47
        } else {
            fmt.Println(v34)
        }
        v48 := GenS5{F0: "a", F1: (^(v46))}
        _ = v48
    } else {
        fmt.Println(v46, v6)
    }
    for fi49 := 0; fi49 < 17; fi49++ {
        var v50 uint32 = 65535
        v51 := v4[65535]
        v52 := GenS5{F0: "foo bar", F1: (v46 << ((1) & 63))}
        _ = fi49
        _ = v50
        _ = v51
        _ = v52
    }
    v6 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
