package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    var v2 string = ("a" + "a")
    v3 := make(map[int]bool)
    var v4 uint16 = 2
    var v5 uint = 255
    v6 := make(map[string]uint32)
    if false {
        var v7 int8 = -22
        _ = v7
    }
    v8 := GenS1{F0: 111}
    v9 := GenS1{F0: 248}
    if (true && false) {
        v10 := []int{9223372036854775807}
        v11 := GenS1{F0: 127}
        _ = v10
        _ = v11
    } else {
        switch 341 {
        case 1:
            fmt.Println(v5, v5)
        case 936:
            v5 += v5
        case 348:
            var v12 int8 = -2
            _ = v12
        default:
            var v13 uint8 = 255
            _ = v13
        }
        var v14 int8 = 100
        _ = v14
    }
    var v15 uint8 = 0
    switch v2 {
    case "a":
        var v16 uint64 = 0
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        v2 = "hello"
        v5 = v5
        _ = fi17
    }
    v15 = 42
    v4 = v4
    if false {
        switch 709 {
        case 9223372036854775807:
            var v18 int64 = -121
            _ = v18
        case -2147483648:
            var v19 float64 = 0.1
            _ = v19
        }
        var v20 string = "foo bar"
        var v21 bool = true
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: 2}
    for fi23 := 0; fi23 < 10; fi23++ {
        _ = len(v3)
        v4 -= uint16(v15)
        _ = fi23
    }
    v24 := []int{}
    if true {
        for fi25 := 0; fi25 < 20; fi25++ {
            if len(v24) > 0 {
                v24[0] = fi25
            }
            _ = fi25
        }
        fmt.Println(v4)
        fmt.Println(v2, v5, v15)
    } else {
        if (!(0 <= (^(v4)))) {
            v2 = "a"
            switch v2 {
            case "foo bar":
                switch -1 {
                case -9223372036854775808:
                    if false {
                        if ((1 < 842) || (1 > 127)) {
                            v26 := GenS2{F0: 42}
                            v15 = v15
                            var v27 uint = v5
                            _ = v26
                            _ = v27
                        } else {
                            var v28 uint64 = 42
                            var v29 float64 = 1.5
                            if (!false) {
                                if true {
                                    if true {
                                        v30 := GenS3{F0: 83}
                                        _ = v30
                                    }
                                    switch v2 {
                                    case "hello":
                                        if false {
                                            var v31 int16 = 32767
                                            _ = v31
                                        }
                                    case "foo bar":
                                        if true {
                                            if false {
                                                v24 = append(v24, 32767)
                                            }
                                        }
                                    case "":
                                        v4 = 65535
                                    default:
                                        v32 := v3[32767]
                                        _ = v32
                                    }
                                    v5 = v5
                                }
                                var v33 int64 = 9223372036854775807
                                _ = v33
                            }
                            _ = v28
                            _ = v29
                        }
                        var v34 int16 = -1
                        v2 = "\n\t"
                        _ = v34
                    }
                case 32767:
                    var v35 string = v2
                    _ = v35
                case 32768:
                    var v36 uint = (v5 & v5)
                    _ = v36
                }
            case "\n\t":
                if len(v24) > 0 {
                    v24[0] = 332
                }
            }
        } else {
            delete(v6, v2)
            fmt.Println(v5)
        }
        var v37 bool = false
        _ = v37
    }
    var v38 bool = false
    v15 -= v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
