package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint32) int {
    if true {
        var v6 string = "hello"
        _ = v6
    }
    if false {
        p0 += p0
    } else {
        p0 = p0
        v7 := GenS1{F0: "foo bar", F1: 0}
        var v8 uint16 = 688
        _ = v7
        _ = v8
    }
    p0 = (p0 | 193)
    _ = p0
    return 2147483647
}

func main() {
    var v1 int16 = 60
    v2 := make(map[string]int)
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    v9 := GenS2{F0: 2}
    for fi10 := 0; fi10 < 16; fi10++ {
        switch (fi10 - 0) {
        case 4294967295:
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 uint64 = 2
                _ = fi11
                _ = v12
            }
        case -2147483648:
            var v13 bool = (!(0 <= 848))
            _ = v13
        case -1:
            var v14 uint64 = 18446744073709551615
            _ = v14
        }
        var v15 uint64 = 4294967295
        switch fi10 {
        case -804:
            for fi16 := 0; fi16 < 12; fi16++ {
                v17 := GenS2{F0: v15}
                switch "foo bar" {
                case "a":
                    var v18 uint = 373
                    _ = v18
                case "hello":
                    switch fi10 {
                    case 1:
                        v19 := GenS1{F0: "foo bar", F1: v15}
                        _ = v19
                    default:
                        if (false && (v15 < (^(v15)))) {
                            if false {
                                var v20 int32 = 2147483647
                                fmt.Println(fi16, v1, fi16)
                                _ = v20
                            } else {
                                var v21 int32 = -677
                                v21 = 42
                                _ = v21
                            }
                        } else {
                            v15 += (v15 % ((v15) | 1))
                            v22 := GenS1{F0: "foo bar", F1: v15}
                            _ = v22
                        }
                    }
                case "foo bar":
                    v15 += v15
                }
                v1 -= v1
                _ = fi16
                _ = v17
            }
        }
        _ = fi10
        _ = v15
    }
    fmt.Println(v1)
    _ = len(v5)
    fmt.Println(v1)
    v1 = 1
    if (-1 <= 1) {
        if ((v1 & (^(v1))) > (-(v1))) {
            _ = len(v5)
            v1 += (v1 >> ((42) & 63))
        } else {
            switch "a" {
            case "hello":
                if ((238 == -643) || true) {
                    v23 := GenS1{F0: "hello", F1: 4294967295}
                    var v24 bool = false
                    _ = v23
                    _ = v24
                }
            case "foo bar":
                fmt.Println(v1)
            default:
                v25 := GenS3{F0: (^(v1)), F1: 65535}
                _ = v25
            }
            v1 = 1
        }
        for fi26 := 0; fi26 < 13; fi26++ {
            fmt.Println(v1)
            if (4294967295 <= 3) {
                v1 = (v1 >> ((0) & 63))
                var v27 uint8 = 127
                if (!(!false)) {
                    for fi28 := 0; fi28 < 12; fi28++ {
                        v29 := GenS4{F0: 0}
                        if (false && false) {
                            v1 -= v1
                        }
                        _ = fi28
                        _ = v29
                    }
                } else {
                    fmt.Println(fi26)
                }
                _ = v27
            }
            var v30 int64 = 42
            _ = fi26
            _ = v30
        }
        _ = len(v3)
    } else {
        v1 = (v1 * (v1 / ((v1) | 1)))
    }
    var v31 uint8 = 1
    for fi32 := 0; fi32 < 14; fi32++ {
        var v33 string = (("hello" + ("hello" + "hello")) + (("foo bar" + "\n\t") + "\n\t"))
        if false {
            if true {
                var v34 int64 = -247
                switch "" {
                case "\n\t":
                    var v35 int = fi32
                    _ = v35
                case "foo bar":
                    v36 := GenS5{F0: -699}
                    _ = v36
                case "hello":
                    v37 := GenS3{F0: int16(v34), F1: 1}
                    _ = v37
                }
                var v38 int32 = 77
                _ = v34
                _ = v38
            }
            if false {
                var v39 uint16 = 583
                _ = v39
            }
        }
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
