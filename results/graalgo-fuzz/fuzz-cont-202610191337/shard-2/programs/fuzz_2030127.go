package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 uint16; F2 string }
func genFn0(p0 int32, p1 uint16) int16 {
    var v5 int32 = p0
    var v6 bool = (false && (-43 >= 0))
    v7 := []int{}
    var v8 uint32 = 4294967295
    var v9 float32 = 0.1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -32768
}

func main() {
    var v1 float64 = 3.14159
    var v2 int8 = 42
    v3 := []int{1, 32767, 128, -602}
    var v4 int32 = -126
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v10 bool = true
    if (681 >= 42) {
        if v10 {
            fmt.Println(v10, v10, v10)
            v1 -= v1
        }
    } else {
        v10 = (!("hello" != "hello"))
        var v11 uint32 = 1
        v12 := GenS1{F0: 189, F1: 122}
        _ = v11
        _ = v12
    }
    fmt.Println(v1, v4, v2)
    if v10 {
        if (!v10) {
            for fi13 := 0; fi13 < 13; fi13++ {
                if (((32767 < 32767) || false) || (!v10)) {
                    var v14 int16 = 1000
                    for fi15 := 0; fi15 < 6; fi15++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v16 float32 = 1e10
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                } else {
                    var v17 string = ""
                    for fi18 := 0; fi18 < 13; fi18++ {
                        fmt.Println(v17)
                        v4 -= -546
                        _ = fi18
                    }
                    var v19 uint8 = 0
                    _ = v17
                    _ = v19
                }
                fmt.Println(v2, v4, v4)
                _ = fi13
            }
        }
    } else {
        v20 := GenS1{F0: 171, F1: 2147483648}
        v21 := GenS2{F0: 1, F1: 130}
        v1 = 3.14159
        _ = v20
        _ = v21
    }
    var v22 int = -9223372036854775808
    v2 -= (^(v2))
    if true {
        switch v22 {
        case 1:
            _ = len(v3)
        case 0:
            v10 = v10
        }
        if v10 {
            if (2 >= 65535) {
                switch ("a" + ("" + "foo bar")) {
                case "hello":
                    v23 := GenS3{F0: 25, F1: 42}
                    _ = v23
                }
            } else {
                switch "a" {
                case "\n\t":
                    v4 = v4
                case "a":
                    fmt.Println(v2, v10)
                default:
                    for fi24 := 0; fi24 < 16; fi24++ {
                        if len(v3) > 0 {
                            v3[0] = v22
                        }
                        v4 = v4
                        v4 += 1
                        _ = fi24
                    }
                }
            }
            v22 = 652
            switch 577 {
            case 42:
                v25 := GenS3{F0: 1, F1: 2147483648}
                _ = v25
            case 2147483647:
                switch ("foo bar" + "foo bar") {
                case "hello":
                    var v26 uint16 = 1
                    _ = v26
                case "foo bar":
                    v27 := GenS4{F0: 0, F1: 234, F2: 9223372036854775807}
                    _ = v27
                case "a":
                    for fi28 := 0; fi28 < 20; fi28++ {
                        var v29 string = "\n\t"
                        var v30 uint32 = 0
                        for fi31 := 0; fi31 < 16; fi31++ {
                            v10 = v10
                            if false {
                                if (!(v10 && ((v4 & (v4 * 1)) < -1))) {
                                    var v32 bool = v10
                                    _ = v32
                                } else {
                                    var v33 int64 = int64(v4)
                                    if v10 {
                                        v10 = v10
                                        v29 = ("" + "")
                                        if len(v3) > 0 {
                                            v3[0] = fi31
                                        }
                                        v1 = float64(float32(float64(float32(1e10))))
                                    } else {
                                        if (v4 == -764) {
                                            fmt.Println(v30, v22)
                                            var v34 bool = ((-32768 < -688) && v10)
                                            _ = v34
                                        }
                                        v3 = append(v3, (fi31 << ((694) & 63)))
                                    }
                                    _ = v33
                                }
                                if (v29 <= ((("" + v29) + v29) + "foo bar")) {
                                    var v35 uint64 = 18446744073709551615
                                    var v36 uint = 0
                                    _ = v35
                                    _ = v36
                                }
                                var v37 float64 = 1e10
                                _ = v37
                            } else {
                                switch (v22 / (((fi31 << ((42) & 63))) | 1)) {
                                case -1:
                                    v38 := GenS5{F0: fi28}
                                    _ = v38
                                case -168:
                                    v22 += (v22 >> ((919) & 63))
                                }
                                _ = len(v3)
                            }
                            _ = fi31
                        }
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
            case -980:
                if ((!false) && v10) {
                    var v39 int8 = (v2 - 100)
                    v40 := GenS6{F0: 1, F1: 2, F2: ""}
                    _ = v39
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 14; fi41++ {
                        if len(v3) > 0 {
                            v3[0] = (fi41 * (v22 * fi41))
                        }
                        _ = fi41
                    }
                    _ = len(v3)
                }
            }
        }
        v2 -= v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
