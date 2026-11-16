package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint16; F1 uint16; F2 bool }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint32 }

func main() {
    v1 := []uint32{218, 2147483647, 0}
    var v2 string = "a"
    v3 := make(map[string]bool)
    var v4 int = -1
    v5 := []uint32{1}
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            switch (v2 + ("" + (v2 + v2))) {
            case "hello":
                if (0 > 4294967295) {
                    var v8 int64 = 237
                    _ = v8
                } else {
                    v9 := GenS1{F0: (v4 ^ (^(v4)))}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v10 := GenS1{F0: v4}
                    v4 += fi6
                    var v11 uint16 = 2
                    _ = v10
                    _ = v11
                    _ = v9
                }
            case "foo bar":
                if (!(!(199 < 136))) {
                    switch v2 {
                    case "foo bar":
                        var v12 int8 = 127
                        _ = v12
                    case "":
                        var v13 int8 = 121
                        _ = v13
                    case "\n\t":
                        if false {
                            var v14 int32 = 2147483647
                            _ = v14
                        }
                    }
                    switch (fi6 << ((1) & 63)) {
                    case 128:
                        delete(v3, "")
                    case 127:
                        v2 = (v2 + v2)
                    case 1:
                        v2 = "foo bar"
                    }
                    fmt.Println(v4)
                } else {
                    v15 := GenS2{F0: 65535, F1: v2}
                    _ = v15
                }
            }
            switch 2147483647 {
            case 2147483647:
                if true {
                    if (167 > 0) {
                        if true {
                            var v16 int8 = 17
                            fmt.Println(v2, v16)
                            v4 = (v4 ^ (fi7 % ((fi7) | 1)))
                            switch 4294967295 {
                            case 128:
                                var v17 string = v2
                                _ = v17
                            }
                            v2 = v2
                            v4 = 1
                            _ = v16
                        }
                    }
                    if false {
                        v18 := GenS1{F0: fi7}
                        if ((4294967295 > 2147483647) && false) {
                            if true {
                                if (335 <= 42) {
                                    switch "" {
                                    case "foo bar":
                                        v19 := GenS3{F0: 255, F1: 42}
                                        _ = v19
                                    case "a":
                                        if true {
                                            switch (v2 + v2) {
                                            case "":
                                                var v20 int = (fi7 + v4)
                                                _ = v20
                                            case "\n\t":
                                                var v21 uint32 = 42
                                                _ = v21
                                            }
                                            v22 := GenS4{F0: -1}
                                            v23 := GenS4{F0: 58}
                                            _ = v22
                                            _ = v23
                                        } else {
                                            fmt.Println(fi6)
                                        }
                                    case "":
                                        var v24 float64 = 0.0
                                        _ = v24
                                    }
                                    _ = len(v5)
                                } else {
                                    v2 = v2
                                    v3["hello"] = false
                                }
                                v25 := GenS1{F0: fi7}
                                fmt.Println(fi6, fi7)
                                v26 := GenS5{F0: 162, F1: -7}
                                v5 = append(v5, 65535)
                                _ = v25
                                _ = v26
                            }
                        } else {
                            var v27 int8 = -88
                            var v28 int64 = 1
                            _ = v27
                            _ = v28
                        }
                        var v29 uint8 = 242
                        _ = v18
                        _ = v29
                    } else {
                        delete(v3, ("foo bar" + "\n\t"))
                    }
                    v30 := GenS6{F0: 1, F1: 32767, F2: true}
                    _ = v30
                } else {
                    v4 = (fi7 - v4)
                    if false {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        v4 = fi6
                        if (!((true && (-236 > 1)) || ((!true) && ((!true) && (v2 >= "\n\t"))))) {
                            switch fi7 {
                            case -103:
                                if true {
                                    v31 := GenS5{F0: 65535, F1: 450}
                                    v32 := GenS1{F0: fi6}
                                    v2 = v2
                                    _ = v31
                                    _ = v32
                                }
                            case 32768:
                                var v33 uint16 = 42
                                _ = v33
                            case 4294967295:
                                _ = len(v3)
                            default:
                                v4 = fi6
                                v2 = "a"
                            }
                        }
                        v2 = v2
                    }
                    var v34 bool = true
                    var v35 uint64 = 4294967295
                    _ = v34
                    _ = v35
                }
            }
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = len(v1)
            }
            var v36 uint8 = 255
            var v37 int16 = 32767
            v36 = 26
            v2 = v2
            _ = fi7
            _ = v36
            _ = v37
        }
        var v38 int8 = 42
        _ = fi6
        _ = v38
    }
    var v39 float32 = 0.0
    delete(v3, ("a" + ""))
    for fi40 := 0; fi40 < 1; fi40++ {
        v41 := GenS7{F0: 42, F1: 255}
        v2 = ((v2 + v2) + ("" + ""))
        if (false && false) {
            v4 = (v4 & (fi40 ^ 4294967295))
            if false {
                v4 = 2147483647
                switch (^(v4)) {
                case 0:
                    fmt.Println(fi40, v4)
                case -1:
                    v4 = v4
                case 128:
                    v42 := GenS8{F0: -881}
                    _ = v42
                }
            } else {
                var v43 bool = (!false)
                _ = v43
            }
        }
        _ = fi40
        _ = v41
    }
    if false {
        fmt.Println(v2, v2)
    } else {
        var v44 string = "\n\t"
        _ = v44
    }
    v45 := GenS6{F0: 1, F1: 32767, F2: (((false || false) && true) && false)}
    v39 += (((-2.25 - 1.5) - v39) / 0.1)
    delete(v3, ((v2 + (v2 + (v2 + v2))) + v2))
    fmt.Println(v39, v2, v39)
    if false {
        if false {
            v1 = append(v1, 4294967295)
            for fi46 := 0; fi46 < 7; fi46++ {
                v47 := GenS7{F0: 2147483647, F1: 2}
                fmt.Println(v2, fi46)
                _ = fi46
                _ = v47
            }
        } else {
            var v48 uint64 = 2
            _ = v48
        }
        v49 := GenS1{F0: v4}
        _ = v49
    }
    var v50 float32 = v39
    for fi51 := 0; fi51 < 14; fi51++ {
        v52 := GenS2{F0: 255, F1: "foo bar"}
        _ = fi51
        _ = v52
    }
    for fi53 := 0; fi53 < 19; fi53++ {
        switch "a" {
        case "hello":
            v54 := GenS9{F0: 213}
            _ = v54
        case "foo bar":
            var v55 uint = 1
            _ = v55
        }
        _ = fi53
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v45
    _ = v5
    _ = v50
}
