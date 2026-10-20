package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 string; F2 string }
type GenS4 struct { F0 int; F1 uint32 }
func genFn0() string {
    var v6 int16 = 1000
    fmt.Println(v6)
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 354, F1: 255, F2: 351}
            var v10 string = "hello"
            var v11 float32 = float32((float64(3.14159) * 1e10))
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi7
    }
    if (2147483647 <= 2) {
        v12 := GenS2{F0: 32767}
        _ = v12
    }
    v13 := []int{344, 32767}
    _ = v13
    _ = v6
    return ""
}

func main() {
    var v1 bool = true
    var v2 uint8 = 127
    var v3 int64 = 0
    var v4 int32 = -1
    var v5 bool = v1
    fmt.Println(v3)
    var v14 uint = 4294967295
    var v15 bool = false
    v16 := []int{127}
    v4 = v4
    var v17 int8 = 42
    v1 = true
    if v1 {
        switch ("foo bar" + "\n\t") {
        case "a":
            switch 0 {
            case -243:
                if (0 > 285) {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v15 = (v2 != 0)
                        var v19 int32 = 2147483647
                        v20 := GenS2{F0: fi18}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    if false {
                        var v21 int = 9223372036854775807
                        var v22 uint = v14
                        v14 = v22
                        _ = v21
                        _ = v22
                    } else {
                        if v15 {
                            v14 += (v14 ^ (v14 + (^(v14))))
                            var v23 bool = (v15 || false)
                            _ = v23
                        } else {
                            if len(v16) > 0 {
                                v16[0] = 128
                            }
                            _ = len(v16)
                        }
                        var v24 uint = uint(v2)
                        _ = v24
                    }
                } else {
                    var v25 string = "a"
                    _ = v25
                }
            case 1:
                var v26 uint32 = 66
                _ = v26
            default:
                v27 := GenS3{F0: 0, F1: "", F2: "hello"}
                _ = v27
            }
        case "foo bar":
            v28 := GenS1{F0: 2, F1: 0, F2: 65535}
            _ = v28
        case "":
            for fi29 := 0; fi29 < 17; fi29++ {
                if ((720 <= 698) || (v2 > (v2 | (^(v2))))) {
                    var v30 uint8 = v2
                    if false {
                        if ((1e10 + (0.1 - 0.0)) == ((float64(-2.25) * 0.1) / ((((3.14159 + ((((float64(0.0) / 1.0) - 0.1) * float64(0.0)) / (1e10 + 1.5))) / 1.0) / 1.0) / 0.1))) {
                            var v31 float64 = 0.1
                            v17 += v17
                            _ = v31
                        } else {
                            var v32 uint64 = uint64(v2)
                            for fi33 := 0; fi33 < 5; fi33++ {
                                if false {
                                    if (-1 < v17) {
                                        var v34 uint16 = 2
                                        fmt.Println(v32)
                                        _ = len(v16)
                                        _ = v34
                                    }
                                    v16 = append(v16, (fi29 - 32768))
                                    if (!v1) {
                                        v5 = v1
                                        v35 := GenS1{F0: 65535, F1: 1, F2: 4294967295}
                                        var v36 int32 = (v4 / ((int32(v32)) | 1))
                                        _ = v35
                                        _ = v36
                                    } else {
                                        if ((-(fi29)) <= fi33) {
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                            v37 := GenS4{F0: 128, F1: 0}
                                            _ = v37
                                        } else {
                                            v17 -= 0
                                            var v38 float64 = 3.14159
                                            _ = v38
                                        }
                                        switch "hello" {
                                        case "hello":
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                        default:
                                            if v15 {
                                                _ = len(v16)
                                            }
                                        }
                                        fmt.Println(v32, v1)
                                    }
                                    v16 = append(v16, (fi33 & (^(fi33))))
                                } else {
                                    switch (fi33 % ((553) | 1)) {
                                    case 0:
                                        var v39 uint32 = 4294967295
                                        _ = v39
                                    }
                                }
                                v17 = v17
                                switch (fi29 - -1) {
                                case 42:
                                    fmt.Println(v1, v32, v1)
                                case 127:
                                    v40 := GenS4{F0: -2147483648, F1: 1}
                                    _ = v40
                                default:
                                    if (2147483647 >= -729) {
                                        v41 := GenS2{F0: fi29}
                                        if len(v16) > 0 {
                                            v16[0] = 42
                                        }
                                        if v5 {
                                            _ = len(v16)
                                            var v42 uint8 = (v2 << (((^(v14))) & 63))
                                            if v1 {
                                                v43 := GenS2{F0: fi33}
                                                var v44 uint64 = v32
                                                v32 -= (^(v32))
                                                _ = v43
                                                _ = v44
                                            }
                                            var v45 int32 = v4
                                            _ = v42
                                            _ = v45
                                        } else {
                                            switch (fi29 / (((fi33 + 4294967295)) | 1)) {
                                            case 1:
                                                v30 += (^(v30))
                                            }
                                            fmt.Println(v5, v30, v2)
                                        }
                                        _ = v41
                                    } else {
                                        _ = len(v16)
                                        v5 = v5
                                    }
                                }
                                _ = fi33
                            }
                            var v46 float32 = 0.1
                            _ = v32
                            _ = v46
                        }
                        v15 = (!v1)
                    } else {
                        fmt.Println(v30, v5)
                    }
                    _ = v30
                }
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = fi29
            }
        }
        fmt.Println(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
