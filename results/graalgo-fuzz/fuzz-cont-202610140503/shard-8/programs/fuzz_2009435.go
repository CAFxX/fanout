package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint16; F2 int16 }
type GenS5 struct { F0 int32; F1 int32 }
func genFn0() uint {
    var v7 int32 = -1
    var v8 float64 = float64(1.5)
    if true {
        v7 += (v7 << ((1) & 63))
    } else {
        var v9 int32 = v7
        var v10 int16 = 0
        _ = v10
        _ = v9
    }
    _ = v7
    _ = v8
    return 2
}
func genFn1() int16 {
    var v11 uint16 = 1
    if (0 != 91) {
        v12 := GenS1{F0: 4294967295}
        _ = v12
    } else {
        v11 = v11
    }
    fmt.Println(v11)
    var v13 string = ""
    var v14 int64 = -9223372036854775808
    _ = v11
    _ = v13
    _ = v14
    return 1
}

func main() {
    v1 := []string{"", "a"}
    v2 := []int{-307, -2147483648, -918}
    var v3 uint = 977
    v4 := make(map[uint32]string)
    var v5 int16 = 1
    v6 := make(map[uint32]string)
    v15 := GenS2{F0: 938, F1: -761}
    var v16 int64 = 523
    var v17 int = 9223372036854775807
    for fi18 := 0; fi18 < 9; fi18++ {
        v19 := GenS3{F0: 1}
        switch ("" + "hello") {
        case "a":
            var v20 uint8 = uint8(v3)
            _ = v20
        case "hello":
            v2 = append(v2, (fi18 >> ((0) & 63)))
        }
        var v21 float32 = 1e10
        _ = fi18
        _ = v19
        _ = v21
    }
    v22 := GenS1{F0: v17}
    var v23 string = "a"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (v3 < genFn0()) {
        for fi24 := 0; fi24 < 3; fi24++ {
            if false {
                var v25 int = 32768
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = len(v2)
                _ = v25
            }
            fmt.Println(v17)
            v17 += v17
            _ = fi24
        }
        v17 += v17
    } else {
        for fi26 := 0; fi26 < 17; fi26++ {
            for fi27 := 0; fi27 < 15; fi27++ {
                v28 := v4[65535]
                if true {
                    if true {
                        v29 := GenS4{F0: (^(v3)), F1: 2, F2: v5}
                        switch fi26 {
                        case 32768:
                            var v30 float64 = float64(3.14159)
                            _ = v30
                        default:
                            var v31 int64 = v16
                            _ = v31
                        }
                        _ = v29
                    } else {
                        if (false || (!(v23 == v23))) {
                            var v32 bool = true
                            if v32 {
                                v1 = append(v1, "hello")
                                v33 := GenS3{F0: -948}
                                _ = v33
                            } else {
                                var v34 string = v28
                                v23 = "hello"
                                _ = v34
                            }
                            var v35 int8 = 127
                            _ = v32
                            _ = v35
                        }
                        switch v23 {
                        case "":
                            if true {
                                v36 := GenS3{F0: 2147483647}
                                if false {
                                    var v37 int16 = genFn1()
                                    _ = v37
                                } else {
                                    v28 = "a"
                                    v23 = (((v28 + (((v23 + v23) + ("" + v28)) + (v23 + "\n\t"))) + v28) + v28)
                                    v16 += (^(v16))
                                }
                                v38 := GenS4{F0: v3, F1: 370, F2: (v5 ^ v5)}
                                if len(v1) > 0 {
                                    v1[0] = (v28 + "a")
                                }
                                v28 = v23
                                _ = v36
                                _ = v38
                            } else {
                                _ = len(v4)
                            }
                        case "foo bar":
                            v16 = v16
                        case "a":
                            v39 := GenS1{F0: 630}
                            _ = v39
                        default:
                            v17 += (fi26 * v17)
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    v40 := GenS5{F0: -1, F1: 2147483647}
                    _ = v40
                }
                var v41 bool = false
                _ = fi27
                _ = v28
                _ = v41
            }
            _ = len(v4)
            var v42 int64 = v16
            _ = fi26
            _ = v42
        }
        var v43 bool = false
        v44 := GenS4{F0: v3, F1: 0, F2: (v5 << ((v3) & 63))}
        _ = v43
        _ = v44
    }
    v3 = (v3 * v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
