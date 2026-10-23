package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }
func genFn0() uint64 {
    v7 := GenS1{F0: 1, F1: -177, F2: 32767}
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := []int{fi8}
        var v10 uint64 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if true {
            var v12 uint8 = 2
            var v13 int32 = -1
            _ = v12
            _ = v13
        }
        if false {
            var v14 int8 = 100
            if (9223372036854775807 >= -9223372036854775808) {
                var v15 int8 = -1
                v16 := GenS1{F0: 0, F1: 0, F2: 989}
                _ = v15
                _ = v16
            } else {
                v17 := []int{(fi11 >> ((705) & 63)), (fi11 ^ 4294967295), 4294967295}
                _ = v17
            }
            _ = v14
        } else {
            var v18 int16 = int16(fi11)
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 int64 = -1
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        var v21 uint = 65535
        _ = fi11
        _ = v21
    }
    var v22 int64 = -691
    _ = v22
    _ = v7
    return 42
}

func main() {
    var v1 uint16 = 42
    var v2 uint8 = 14
    var v3 uint8 = (v2 % ((0) | 1))
    var v4 bool = ("\n\t" != "hello")
    v5 := make(map[string]string)
    var v6 uint32 = 2
    switch "a" {
    case "":
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS2{F0: 135}
            v1 += v1
            _ = len(v5)
            _ = fi23
            _ = v24
        }
    case "hello":
        fmt.Println(v6)
    }
    v25 := []int{}
    switch "hello" {
    case "\n\t":
        for fi26 := 0; fi26 < 16; fi26++ {
            switch (fi26 >> ((2) & 63)) {
            case 131:
                if (!(0 < (^(v6)))) {
                    switch "\n\t" {
                    case "\n\t":
                        if (((^(v1)) > (^(v1))) && true) {
                            fmt.Println(v3, v4, v2)
                            v27 := GenS1{F0: 81, F1: 42, F2: -1}
                            fmt.Println(v1, v4, v6)
                            _ = v27
                        } else {
                            for fi28 := 0; fi28 < 8; fi28++ {
                                if true {
                                    v1 -= 0
                                    switch "foo bar" {
                                    case "a":
                                        var v29 bool = (0 <= 42)
                                        _ = v29
                                    default:
                                        v2 += 42
                                    }
                                }
                                _ = fi28
                            }
                        }
                    case "":
                        if v4 {
                            if (0 > 1) {
                                if false {
                                    if false {
                                        v3 += v3
                                    }
                                    var v30 string = "a"
                                    switch v30 {
                                    case "hello":
                                        v31 := GenS1{F0: 2, F1: 9223372036854775807, F2: 506}
                                        _ = v31
                                    }
                                    _ = v30
                                }
                                var v32 uint8 = 114
                                _ = v32
                            }
                        } else {
                            v2 += (v2 - (v3 / ((v2) | 1)))
                        }
                    case "foo bar":
                        fmt.Println(v2, v6)
                    }
                }
            case -9223372036854775808:
                _ = len(v25)
            case 944:
                if len(v25) > 0 {
                    v25[0] = fi26
                }
            default:
                delete(v5, ("a" + ((("hello" + ("" + "hello")) + (("a" + "hello") + "\n\t")) + ("\n\t" + "hello"))))
            }
            for fi33 := 0; fi33 < 20; fi33++ {
                switch fi26 {
                case 32768:
                    v3 = (^(v3))
                case 636:
                    switch ("foo bar" + "\n\t") {
                    case "hello":
                        fmt.Println(v4, fi33)
                    default:
                        if true {
                            switch "\n\t" {
                            case "a":
                                var v34 string = ""
                                _ = v34
                            default:
                                v35 := GenS2{F0: -1}
                                _ = v35
                            }
                            var v36 uint32 = (v6 * v6)
                            var v37 int8 = -128
                            if v4 {
                                fmt.Println(v37, v1)
                            }
                            var v38 uint32 = (v6 ^ 604)
                            _ = v36
                            _ = v37
                            _ = v38
                        } else {
                            switch ("a" + "foo bar") {
                            case "\n\t":
                                if (!v4) {
                                    var v39 bool = (fi33 < (fi26 | 127))
                                    v40 := GenS1{F0: 122, F1: -1, F2: 32767}
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                    _ = v39
                                    _ = v40
                                } else {
                                    v4 = ((("\n\t" + "") + "a") <= (("a" + "\n\t") + ("\n\t" + "foo bar")))
                                    if false {
                                        var v41 uint16 = 32767
                                        v6 += (^(v6))
                                        _ = v41
                                    }
                                }
                            case "":
                                var v42 int = (fi26 / ((128) | 1))
                                _ = v42
                            default:
                                var v43 int32 = 0
                                _ = v43
                            }
                            var v44 float32 = 1e10
                            var v45 float64 = 0.1
                            delete(v5, ("a" + ("foo bar" + "hello")))
                            _ = v44
                            _ = v45
                        }
                    }
                default:
                    v4 = v4
                }
                _ = fi33
            }
            _ = fi26
        }
    case "foo bar":
        var v46 string = "hello"
        _ = v46
    case "":
        v3 = v2
    }
    var v47 int = 127
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v1, v6)
    case "":
        if false {
            v48 := GenS1{F0: 18446744073709551615, F1: 334, F2: -32768}
            v49 := GenS2{F0: -99}
            var v50 int64 = -764
            _ = v48
            _ = v49
            _ = v50
        } else {
            for fi51 := 0; fi51 < 7; fi51++ {
                var v52 uint16 = (v1 * 2)
                v6 = (v6 ^ 4294967295)
                _ = fi51
                _ = v52
            }
        }
    }
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    v53 := GenS1{F0: 65535, F1: 9223372036854775807, F2: 32767}
    var v54 int = (v47 * v47)
    v55 := GenS2{F0: 42}
    if v4 {
        v1 += (v1 >> ((4294967295) & 63))
        v6 = (^(v6))
    }
    for fi56 := 0; fi56 < 8; fi56++ {
        switch ("\n\t" + ((("" + "\n\t") + ("hello" + "")) + "")) {
        case "hello":
            v47 -= fi56
        case "\n\t":
            if false {
                fmt.Println(v54)
                v57 := GenS3{F0: 42}
                _ = v57
            }
        }
        v58 := GenS4{F0: "\n\t"}
        _ = fi56
        _ = v58
    }
    v4 = (v4 && ((!false) || ((-(v54)) > v47)))
    var v59 int16 = -1
    if true {
        v60 := GenS1{F0: 42, F1: 0, F2: v59}
        _ = v60
    } else {
        if (("a" + ("\n\t" + "hello")) == "a") {
            v54 += (v54 ^ int(v6))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v59
    _ = v6
}
