package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int; F1 uint8; F2 int64 }
type GenS8 struct { F0 int16; F1 uint8; F2 int64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int16; F1 uint16 }
func genFn0() uint16 {
    var v5 int16 = 863
    v6 := []int{}
    fmt.Println(v5)
    fmt.Println(v5)
    v7 := GenS1{F0: 0, F1: 222}
    _ = v5
    _ = v6
    _ = v7
    return 447
}
func genFn1() int64 {
    var v8 uint = 255
    var v9 int16 = 1
    var v10 uint64 = 992
    switch "hello" {
    case "hello":
        if ((-128 != -86) && (-751 >= 42)) {
            fmt.Println(v10)
            if true {
                var v11 int16 = 770
                fmt.Println(v11)
                _ = v11
            } else {
                v10 -= (^(v10))
                var v12 float64 = (0.1 - -2.25)
                _ = v12
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 uint8 = 45
                _ = fi13
                _ = v14
            }
        }
    case "a":
        v15 := GenS2{F0: (v9 | v9), F1: v10, F2: 2}
        _ = v15
    case "foo bar":
        v10 -= v10
    }
    v9 = v9
    _ = v10
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]string)
    var v3 int32 = -2147483648
    var v4 uint32 = 42
    var v16 string = "foo bar"
    var v17 uint16 = 320
    v18 := GenS1{F0: 100, F1: 4294967295}
    var v19 float64 = (1.5 * 0.1)
    var v20 int8 = 0
    _ = len(v2)
    var v21 uint32 = v4
    var v22 float64 = v19
    for fi23 := 0; fi23 < 2; fi23++ {
        if (42 >= uint32(v17)) {
            v3 += (v3 >> ((4294967295) & 63))
            for fi24 := 0; fi24 < 6; fi24++ {
                var v25 string = v16
                v22 = (((v22 + (v22 * 1e10)) * ((((v19 / v19) + (0.1 / (3.14159 * 0.1))) - v22) * 0.1)) / (v22 / ((v19 + v19) * v22)))
                var v26 float32 = ((3.14159 / 3.14159) / 1.0)
                _ = fi24
                _ = v25
                _ = v26
            }
        } else {
            v16 = "a"
        }
        switch ((v16 + "hello") + "foo bar") {
        case "\n\t":
            if false {
                var v27 uint = 255
                v28 := []int{fi23}
                fmt.Println(v19)
                _ = v27
                _ = v28
            } else {
                v29 := GenS3{F0: (^(fi23))}
                if false {
                    var v30 int16 = 42
                    for fi31 := 0; fi31 < 16; fi31++ {
                        var v32 int16 = -753
                        if (!((0.0 == float32(1.5)) && true)) {
                            var v33 int = fi31
                            fmt.Println(v16, v22, v3)
                            var v34 bool = (v21 != 311)
                            _ = v33
                            _ = v34
                        } else {
                            var v35 uint16 = v17
                            v36 := v2["a"]
                            _ = v35
                            _ = v36
                        }
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                } else {
                    var v37 int16 = 42
                    fmt.Println(v17)
                    _ = v37
                }
                fmt.Println(v17, v22, v17)
                _ = v29
            }
        case "":
            var v38 uint8 = 42
            _ = v38
        default:
            fmt.Println(v19, v19, v22)
        }
        _ = fi23
    }
    fmt.Println(v3, v19, v22)
    v39 := GenS4{F0: 42, F1: v4}
    v40 := []int{359, 31, 4294967295}
    if len(v40) > 0 {
        v40[0] = -886
    }
    v17 = v17
    v2[v16] = (v16 + v16)
    switch "hello" {
    case "foo bar":
        v41 := GenS5{F0: v3}
        _ = v41
    case "\n\t":
        if (-949 < 9223372036854775807) {
            for fi42 := 0; fi42 < 1; fi42++ {
                for fi43 := 0; fi43 < 19; fi43++ {
                    if true {
                        var v44 uint = 65535
                        v45 := GenS6{F0: 126, F1: v16}
                        switch "a" {
                        case "a":
                            v40 = append(v40, -1)
                            switch v16 {
                            case "foo bar":
                                switch (("foo bar" + (v16 + (v16 + v16))) + v16) {
                                case "":
                                    switch "hello" {
                                    case "a":
                                        switch ("" + v16) {
                                        case "hello":
                                            switch ((v16 + v16) + "a") {
                                            case "foo bar":
                                                v46 := GenS1{F0: 76, F1: uint(v1)}
                                                _ = len(v40)
                                                _ = v46
                                            default:
                                                v47 := GenS7{F0: fi43, F1: 42, F2: 9223372036854775807}
                                                _ = v47
                                            }
                                        case "foo bar":
                                            var v48 uint = 255
                                            _ = v48
                                        }
                                    default:
                                        var v49 int16 = 1
                                        _ = v49
                                    }
                                case "\n\t":
                                    v50 := v2[("a" + ("foo bar" + v16))]
                                    _ = v50
                                case "foo bar":
                                    fmt.Println(v21, v3, v44)
                                }
                            }
                            v16 = v16
                            v51 := GenS8{F0: -32768, F1: 216, F2: 42}
                            _ = v51
                        case "foo bar":
                            fmt.Println(v4, v4)
                        }
                        _ = v44
                        _ = v45
                    }
                    v52 := GenS9{F0: v17}
                    _ = fi43
                    _ = v52
                }
                _ = fi42
            }
            v53 := GenS10{F0: 32767, F1: v1}
            v1 -= 1
            _ = v53
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}
