package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint8, p1 int8) uint16 {
    p0 -= 1
    var v7 float32 = 3.14159
    var v8 float64 = 3.14159
    v9 := GenS1{F0: 644, F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint16, p1 uint8, p2 int32) int64 {
    if (!("a" == "a")) {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 int16 = 1000
            var v12 uint64 = 264
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 uint64 = 18446744073709551615
    v14 := []int{813, 623, -1, 128}
    v15 := GenS2{F0: (v13 << ((2) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    return -913
}

func main() {
    var v1 int = 4294967295
    var v2 int32 = -1
    v3 := make(map[uint32]bool)
    var v4 uint32 = 2
    v5 := make(map[string]int)
    var v6 float32 = (((1.5 / 3.14159) * 1e10) + 0.1)
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 int64 = 42
        switch "" {
        case "":
            v1 = (fi16 & fi16)
        case "a":
            v18 := []int{(^(fi16))}
            _ = v18
        case "\n\t":
            var v19 string = (("a" + "hello") + "hello")
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    v4 -= v4
    var v20 float32 = (((0.0 - 0.1) * (0.1 - 1.5)) + 1e10)
    if false {
        var v21 uint8 = 92
        v22 := GenS1{F0: v4, F1: 412}
        if (!true) {
            if true {
                switch ("a" + "a") {
                case "":
                    fmt.Println(v4)
                case "a":
                    for fi23 := 0; fi23 < 15; fi23++ {
                        for fi24 := 0; fi24 < 3; fi24++ {
                            var v25 int = fi24
                            _ = fi24
                            _ = v25
                        }
                        var v26 uint8 = 2
                        var v27 uint64 = 4294967295
                        _ = fi23
                        _ = v26
                        _ = v27
                    }
                case "foo bar":
                    switch v1 {
                    case 127:
                        delete(v5, "\n\t")
                    case 1:
                        for fi28 := 0; fi28 < 1; fi28++ {
                            switch 2147483647 {
                            case -9223372036854775808:
                                v29 := []int{}
                                _ = v29
                            case 32767:
                                v21 += v21
                            case 42:
                                v2 += (v2 & (v2 * v2))
                            }
                            var v30 string = "hello"
                            _ = fi28
                            _ = v30
                        }
                    case 9223372036854775807:
                        var v31 float32 = (0.1 * (v20 + (v6 * v20)))
                        _ = v31
                    }
                default:
                    switch "foo bar" {
                    case "foo bar":
                        if true {
                            v1 = v1
                            fmt.Println(v20, v21, v21)
                            v1 = (v1 & 128)
                        }
                    case "\n\t":
                        for fi32 := 0; fi32 < 13; fi32++ {
                            var v33 int = v1
                            _ = fi32
                            _ = v33
                        }
                    case "a":
                        for fi34 := 0; fi34 < 3; fi34++ {
                            var v35 uint8 = v21
                            _ = fi34
                            _ = v35
                        }
                    default:
                        var v36 bool = (!(true || (!true)))
                        _ = v36
                    }
                }
                var v37 int64 = int64(v2)
                _ = v37
            }
            for fi38 := 0; fi38 < 16; fi38++ {
                fmt.Println(fi38)
                for fi39 := 0; fi39 < 13; fi39++ {
                    var v40 uint16 = 32767
                    _ = fi39
                    _ = v40
                }
                switch 4294967295 {
                case 127:
                    var v41 float32 = 0.0
                    _ = v41
                case 42:
                    v42 := GenS2{F0: 0}
                    _ = v42
                case 32768:
                    var v43 float64 = 1.5
                    _ = v43
                default:
                    if (true || false) {
                        var v44 string = ""
                        switch (v1 | (fi38 >> ((255) & 63))) {
                        case 32768:
                            for fi45 := 0; fi45 < 20; fi45++ {
                                v1 = (fi45 >> ((42) & 63))
                                switch ((("\n\t" + v44) + "a") + (("" + ((((((("hello" + v44) + v44) + (((((((((("foo bar" + "foo bar") + "a") + ((v44 + (v44 + v44)) + v44)) + "hello") + v44) + ("a" + (v44 + v44))) + ("a" + (v44 + (v44 + v44)))) + v44) + "hello") + v44)) + (v44 + (v44 + (v44 + ((v44 + ("foo bar" + v44)) + (v44 + v44)))))) + (v44 + (v44 + (v44 + (("hello" + "foo bar") + ""))))) + "\n\t") + "a")) + "a")) {
                                case "foo bar":
                                    if false {
                                        var v46 uint16 = genFn0(v21, 42)
                                        var v47 float32 = 3.14159
                                        var v48 uint = 42
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    } else {
                                        v20 = v20
                                        if (false && false) {
                                            v21 = (^(v21))
                                        } else {
                                            var v49 int8 = 42
                                            v50 := GenS2{F0: 18446744073709551615}
                                            v4 = v4
                                            v1 = fi38
                                            _ = v49
                                            _ = v50
                                        }
                                        var v51 uint32 = v4
                                        _ = v51
                                    }
                                case "hello":
                                    v6 -= v20
                                }
                                v1 += 9223372036854775807
                                var v52 float32 = v6
                                if false {
                                    if false {
                                        var v53 int = (fi38 - int(v2))
                                        var v54 int8 = -1
                                        v55 := GenS3{F0: 530}
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    }
                                    v1 = (v1 + 139)
                                } else {
                                    v56 := GenS3{F0: 42}
                                    v57 := v5[(v44 + v44)]
                                    v58 := GenS4{F0: 32768}
                                    v20 = v6
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                }
                                _ = fi45
                                _ = v52
                            }
                        }
                        _ = v44
                    }
                }
                _ = fi38
            }
            v59 := []int{v1, v1}
            _ = v59
        }
        _ = v21
        _ = v22
    }
    fmt.Println(v6, v1)
    v60 := GenS3{F0: 601}
    v61 := GenS5{F0: 1}
    if false {
        v62 := []int{-1, 0}
        _ = v62
    } else {
        if false {
            var v63 float64 = float64(v6)
            v2 = 0
            v64 := GenS4{F0: (v1 + v1)}
            _ = v63
            _ = v64
        } else {
            if (false && false) {
                v1 = (^(v1))
            } else {
                var v65 string = "\n\t"
                _ = v65
            }
        }
        var v66 int32 = (v2 << ((2) & 63))
        _ = v66
    }
    if false {
        if true {
            v2 -= v2
            var v67 int64 = -1
            _ = v67
        } else {
            v68 := GenS6{F0: 100}
            _ = v68
        }
        switch ("" + "\n\t") {
        case "foo bar":
            var v69 int = v1
            _ = v69
        case "a":
            for fi70 := 0; fi70 < 20; fi70++ {
                fmt.Println(v6)
                var v71 int32 = v2
                _ = fi70
                _ = v71
            }
        default:
            switch 2147483647 {
            case 32767:
                var v72 string = "hello"
                _ = v72
            case -955:
                _ = len(v5)
            case 128:
                var v73 float64 = (3.14159 - 1.5)
                _ = v73
            default:
                var v74 int64 = -1
                _ = v74
            }
        }
    } else {
        var v75 uint = 18446744073709551615
        var v76 uint64 = 0
        _ = v75
        _ = v76
    }
    v1 += (v1 / ((1) | 1))
    for fi77 := 0; fi77 < 19; fi77++ {
        v3[v4] = (true || true)
        var v78 uint64 = 0
        _ = fi77
        _ = v78
    }
    v20 = v20
    var v79 int64 = 941
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
    _ = v79
}
